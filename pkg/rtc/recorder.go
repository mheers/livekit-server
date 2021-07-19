package rtc

import (
	"fmt"
	"net"
	"time"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3"
)

type udpConn struct {
	conn        *net.UDPConn
	port        int
	payloadType uint8
}

// Prepare udp conns
// Also update incoming packets with expected PayloadType, the browser may use
// a different value. We have to modify so our stream matches what rtp-forwarder.sdp expects
var udpConns = map[string]*udpConn{
	"audio": {port: 4000, payloadType: 111},
	"video": {port: 4002, payloadType: 96},
}

var laddr *net.UDPAddr

var ConnectionToStream = "Marcello"

func init() {
	var err error

	// Create a local addr
	if laddr, err = net.ResolveUDPAddr("udp", "127.0.0.1:"); err != nil {
		panic(err)
	}

	go startUDPListener()
}

func startUDPListener() {
	var err error
	for k, c := range udpConns {
		// Create remote addr
		var raddr *net.UDPAddr
		if raddr, err = net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", c.port)); err != nil {
			panic(err)
		}

		// Dial udp
		if c.conn, err = net.DialUDP("udp", laddr, raddr); err != nil {
			panic(err)
		}
		defer func(conn net.PacketConn) {
			if closeErr := conn.Close(); closeErr != nil {
				panic(closeErr)
			}
		}(c.conn)
		fmt.Printf("Started udp for %s on port %d\n", k, c.port)
	}

	// Block forever
	select {}
}

func record(peerConnection *webrtc.PeerConnection, id string) {
	var err error

	fmt.Println(id)

	// Set a handler for when a new remote track starts, this handler will forward data to
	// our UDP listeners.
	// In your application this is where you would handle/process audio/video
	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		fmt.Printf("peerConnection.OnTrack\n")
		// Retrieve udp connection
		c, ok := udpConns[track.Kind().String()]
		if !ok {
			return
		}

		// Send a PLI on an interval so that the publisher is pushing a keyframe every rtcpPLIInterval
		go func() {
			ticker := time.NewTicker(time.Second * 2)
			for range ticker.C {
				if rtcpErr := peerConnection.WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: uint32(track.SSRC())}}); rtcpErr != nil {
					fmt.Println(rtcpErr)
				}
			}
		}()

		b := make([]byte, 1500)
		rtpPacket := &rtp.Packet{}
		for {
			// Read
			n, _, readErr := track.Read(b)
			if readErr != nil {
				return
			}

			// Unmarshal the packet and update the PayloadType
			if err = rtpPacket.Unmarshal(b[:n]); err != nil {
				return
			}
			rtpPacket.PayloadType = c.payloadType

			// Marshal into original buffer with updated PayloadType
			if n, err = rtpPacket.MarshalTo(b); err != nil {
				return
			}

			if id == ConnectionToStream {
				// Write
				if _, err = c.conn.Write(b[:n]); err != nil {
					// For this particular example, third party applications usually timeout after a short
					// amount of time during which the user doesn't have enough time to provide the answer
					// to the browser.
					// That's why, for this particular example, the user first needs to provide the answer
					// to the browser then open the third party application. Therefore we must not kill
					// the forward on "connection refused" errors
					if opError, ok := err.(*net.OpError); ok && opError.Err.Error() == "write: connection refused" {
						continue
					}
					return
				}
			}

		}
	})
}
