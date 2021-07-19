# Recording

```bash
# Record
ffmpeg -protocol_whitelist file,udp,rtp -i recorder.sdp -max_muxing_queue_size 9999 /tmp/test.webm

# Live in VLC
vlc recorder.sdp
```
