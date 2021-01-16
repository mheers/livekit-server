// Code generated by counterfeiter. DO NOT EDIT.
package routingfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/routing"
	"github.com/livekit/livekit-server/proto/livekit"
)

type FakeRouter struct {
	GetNodeStub        func(string) (*livekit.Node, error)
	getNodeMutex       sync.RWMutex
	getNodeArgsForCall []struct {
		arg1 string
	}
	getNodeReturns struct {
		result1 *livekit.Node
		result2 error
	}
	getNodeReturnsOnCall map[int]struct {
		result1 *livekit.Node
		result2 error
	}
	GetNodeIdForRoomStub        func(string) (string, error)
	getNodeIdForRoomMutex       sync.RWMutex
	getNodeIdForRoomArgsForCall []struct {
		arg1 string
	}
	getNodeIdForRoomReturns struct {
		result1 string
		result2 error
	}
	getNodeIdForRoomReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetRequestSinkStub        func(string) routing.MessageSink
	getRequestSinkMutex       sync.RWMutex
	getRequestSinkArgsForCall []struct {
		arg1 string
	}
	getRequestSinkReturns struct {
		result1 routing.MessageSink
	}
	getRequestSinkReturnsOnCall map[int]struct {
		result1 routing.MessageSink
	}
	GetResponseSourceStub        func(string) routing.MessageSource
	getResponseSourceMutex       sync.RWMutex
	getResponseSourceArgsForCall []struct {
		arg1 string
	}
	getResponseSourceReturns struct {
		result1 routing.MessageSource
	}
	getResponseSourceReturnsOnCall map[int]struct {
		result1 routing.MessageSource
	}
	OnNewParticipantStub        func(routing.ParticipantCallback)
	onNewParticipantMutex       sync.RWMutex
	onNewParticipantArgsForCall []struct {
		arg1 routing.ParticipantCallback
	}
	RegisterNodeStub        func(*livekit.Node) error
	registerNodeMutex       sync.RWMutex
	registerNodeArgsForCall []struct {
		arg1 *livekit.Node
	}
	registerNodeReturns struct {
		result1 error
	}
	registerNodeReturnsOnCall map[int]struct {
		result1 error
	}
	SetRTCNodeStub        func(string, string) error
	setRTCNodeMutex       sync.RWMutex
	setRTCNodeArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setRTCNodeReturns struct {
		result1 error
	}
	setRTCNodeReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StartParticipantStub        func(string, string, string, string) error
	startParticipantMutex       sync.RWMutex
	startParticipantArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	startParticipantReturns struct {
		result1 error
	}
	startParticipantReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func()
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRouter) GetNode(arg1 string) (*livekit.Node, error) {
	fake.getNodeMutex.Lock()
	ret, specificReturn := fake.getNodeReturnsOnCall[len(fake.getNodeArgsForCall)]
	fake.getNodeArgsForCall = append(fake.getNodeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetNodeStub
	fakeReturns := fake.getNodeReturns
	fake.recordInvocation("GetNode", []interface{}{arg1})
	fake.getNodeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRouter) GetNodeCallCount() int {
	fake.getNodeMutex.RLock()
	defer fake.getNodeMutex.RUnlock()
	return len(fake.getNodeArgsForCall)
}

func (fake *FakeRouter) GetNodeCalls(stub func(string) (*livekit.Node, error)) {
	fake.getNodeMutex.Lock()
	defer fake.getNodeMutex.Unlock()
	fake.GetNodeStub = stub
}

func (fake *FakeRouter) GetNodeArgsForCall(i int) string {
	fake.getNodeMutex.RLock()
	defer fake.getNodeMutex.RUnlock()
	argsForCall := fake.getNodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) GetNodeReturns(result1 *livekit.Node, result2 error) {
	fake.getNodeMutex.Lock()
	defer fake.getNodeMutex.Unlock()
	fake.GetNodeStub = nil
	fake.getNodeReturns = struct {
		result1 *livekit.Node
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) GetNodeReturnsOnCall(i int, result1 *livekit.Node, result2 error) {
	fake.getNodeMutex.Lock()
	defer fake.getNodeMutex.Unlock()
	fake.GetNodeStub = nil
	if fake.getNodeReturnsOnCall == nil {
		fake.getNodeReturnsOnCall = make(map[int]struct {
			result1 *livekit.Node
			result2 error
		})
	}
	fake.getNodeReturnsOnCall[i] = struct {
		result1 *livekit.Node
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) GetNodeIdForRoom(arg1 string) (string, error) {
	fake.getNodeIdForRoomMutex.Lock()
	ret, specificReturn := fake.getNodeIdForRoomReturnsOnCall[len(fake.getNodeIdForRoomArgsForCall)]
	fake.getNodeIdForRoomArgsForCall = append(fake.getNodeIdForRoomArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetNodeIdForRoomStub
	fakeReturns := fake.getNodeIdForRoomReturns
	fake.recordInvocation("GetNodeIdForRoom", []interface{}{arg1})
	fake.getNodeIdForRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRouter) GetNodeIdForRoomCallCount() int {
	fake.getNodeIdForRoomMutex.RLock()
	defer fake.getNodeIdForRoomMutex.RUnlock()
	return len(fake.getNodeIdForRoomArgsForCall)
}

func (fake *FakeRouter) GetNodeIdForRoomCalls(stub func(string) (string, error)) {
	fake.getNodeIdForRoomMutex.Lock()
	defer fake.getNodeIdForRoomMutex.Unlock()
	fake.GetNodeIdForRoomStub = stub
}

func (fake *FakeRouter) GetNodeIdForRoomArgsForCall(i int) string {
	fake.getNodeIdForRoomMutex.RLock()
	defer fake.getNodeIdForRoomMutex.RUnlock()
	argsForCall := fake.getNodeIdForRoomArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) GetNodeIdForRoomReturns(result1 string, result2 error) {
	fake.getNodeIdForRoomMutex.Lock()
	defer fake.getNodeIdForRoomMutex.Unlock()
	fake.GetNodeIdForRoomStub = nil
	fake.getNodeIdForRoomReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) GetNodeIdForRoomReturnsOnCall(i int, result1 string, result2 error) {
	fake.getNodeIdForRoomMutex.Lock()
	defer fake.getNodeIdForRoomMutex.Unlock()
	fake.GetNodeIdForRoomStub = nil
	if fake.getNodeIdForRoomReturnsOnCall == nil {
		fake.getNodeIdForRoomReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getNodeIdForRoomReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) GetRequestSink(arg1 string) routing.MessageSink {
	fake.getRequestSinkMutex.Lock()
	ret, specificReturn := fake.getRequestSinkReturnsOnCall[len(fake.getRequestSinkArgsForCall)]
	fake.getRequestSinkArgsForCall = append(fake.getRequestSinkArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetRequestSinkStub
	fakeReturns := fake.getRequestSinkReturns
	fake.recordInvocation("GetRequestSink", []interface{}{arg1})
	fake.getRequestSinkMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) GetRequestSinkCallCount() int {
	fake.getRequestSinkMutex.RLock()
	defer fake.getRequestSinkMutex.RUnlock()
	return len(fake.getRequestSinkArgsForCall)
}

func (fake *FakeRouter) GetRequestSinkCalls(stub func(string) routing.MessageSink) {
	fake.getRequestSinkMutex.Lock()
	defer fake.getRequestSinkMutex.Unlock()
	fake.GetRequestSinkStub = stub
}

func (fake *FakeRouter) GetRequestSinkArgsForCall(i int) string {
	fake.getRequestSinkMutex.RLock()
	defer fake.getRequestSinkMutex.RUnlock()
	argsForCall := fake.getRequestSinkArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) GetRequestSinkReturns(result1 routing.MessageSink) {
	fake.getRequestSinkMutex.Lock()
	defer fake.getRequestSinkMutex.Unlock()
	fake.GetRequestSinkStub = nil
	fake.getRequestSinkReturns = struct {
		result1 routing.MessageSink
	}{result1}
}

func (fake *FakeRouter) GetRequestSinkReturnsOnCall(i int, result1 routing.MessageSink) {
	fake.getRequestSinkMutex.Lock()
	defer fake.getRequestSinkMutex.Unlock()
	fake.GetRequestSinkStub = nil
	if fake.getRequestSinkReturnsOnCall == nil {
		fake.getRequestSinkReturnsOnCall = make(map[int]struct {
			result1 routing.MessageSink
		})
	}
	fake.getRequestSinkReturnsOnCall[i] = struct {
		result1 routing.MessageSink
	}{result1}
}

func (fake *FakeRouter) GetResponseSource(arg1 string) routing.MessageSource {
	fake.getResponseSourceMutex.Lock()
	ret, specificReturn := fake.getResponseSourceReturnsOnCall[len(fake.getResponseSourceArgsForCall)]
	fake.getResponseSourceArgsForCall = append(fake.getResponseSourceArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetResponseSourceStub
	fakeReturns := fake.getResponseSourceReturns
	fake.recordInvocation("GetResponseSource", []interface{}{arg1})
	fake.getResponseSourceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) GetResponseSourceCallCount() int {
	fake.getResponseSourceMutex.RLock()
	defer fake.getResponseSourceMutex.RUnlock()
	return len(fake.getResponseSourceArgsForCall)
}

func (fake *FakeRouter) GetResponseSourceCalls(stub func(string) routing.MessageSource) {
	fake.getResponseSourceMutex.Lock()
	defer fake.getResponseSourceMutex.Unlock()
	fake.GetResponseSourceStub = stub
}

func (fake *FakeRouter) GetResponseSourceArgsForCall(i int) string {
	fake.getResponseSourceMutex.RLock()
	defer fake.getResponseSourceMutex.RUnlock()
	argsForCall := fake.getResponseSourceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) GetResponseSourceReturns(result1 routing.MessageSource) {
	fake.getResponseSourceMutex.Lock()
	defer fake.getResponseSourceMutex.Unlock()
	fake.GetResponseSourceStub = nil
	fake.getResponseSourceReturns = struct {
		result1 routing.MessageSource
	}{result1}
}

func (fake *FakeRouter) GetResponseSourceReturnsOnCall(i int, result1 routing.MessageSource) {
	fake.getResponseSourceMutex.Lock()
	defer fake.getResponseSourceMutex.Unlock()
	fake.GetResponseSourceStub = nil
	if fake.getResponseSourceReturnsOnCall == nil {
		fake.getResponseSourceReturnsOnCall = make(map[int]struct {
			result1 routing.MessageSource
		})
	}
	fake.getResponseSourceReturnsOnCall[i] = struct {
		result1 routing.MessageSource
	}{result1}
}

func (fake *FakeRouter) OnNewParticipant(arg1 routing.ParticipantCallback) {
	fake.onNewParticipantMutex.Lock()
	fake.onNewParticipantArgsForCall = append(fake.onNewParticipantArgsForCall, struct {
		arg1 routing.ParticipantCallback
	}{arg1})
	stub := fake.OnNewParticipantStub
	fake.recordInvocation("OnNewParticipant", []interface{}{arg1})
	fake.onNewParticipantMutex.Unlock()
	if stub != nil {
		fake.OnNewParticipantStub(arg1)
	}
}

func (fake *FakeRouter) OnNewParticipantCallCount() int {
	fake.onNewParticipantMutex.RLock()
	defer fake.onNewParticipantMutex.RUnlock()
	return len(fake.onNewParticipantArgsForCall)
}

func (fake *FakeRouter) OnNewParticipantCalls(stub func(routing.ParticipantCallback)) {
	fake.onNewParticipantMutex.Lock()
	defer fake.onNewParticipantMutex.Unlock()
	fake.OnNewParticipantStub = stub
}

func (fake *FakeRouter) OnNewParticipantArgsForCall(i int) routing.ParticipantCallback {
	fake.onNewParticipantMutex.RLock()
	defer fake.onNewParticipantMutex.RUnlock()
	argsForCall := fake.onNewParticipantArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) RegisterNode(arg1 *livekit.Node) error {
	fake.registerNodeMutex.Lock()
	ret, specificReturn := fake.registerNodeReturnsOnCall[len(fake.registerNodeArgsForCall)]
	fake.registerNodeArgsForCall = append(fake.registerNodeArgsForCall, struct {
		arg1 *livekit.Node
	}{arg1})
	stub := fake.RegisterNodeStub
	fakeReturns := fake.registerNodeReturns
	fake.recordInvocation("RegisterNode", []interface{}{arg1})
	fake.registerNodeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) RegisterNodeCallCount() int {
	fake.registerNodeMutex.RLock()
	defer fake.registerNodeMutex.RUnlock()
	return len(fake.registerNodeArgsForCall)
}

func (fake *FakeRouter) RegisterNodeCalls(stub func(*livekit.Node) error) {
	fake.registerNodeMutex.Lock()
	defer fake.registerNodeMutex.Unlock()
	fake.RegisterNodeStub = stub
}

func (fake *FakeRouter) RegisterNodeArgsForCall(i int) *livekit.Node {
	fake.registerNodeMutex.RLock()
	defer fake.registerNodeMutex.RUnlock()
	argsForCall := fake.registerNodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) RegisterNodeReturns(result1 error) {
	fake.registerNodeMutex.Lock()
	defer fake.registerNodeMutex.Unlock()
	fake.RegisterNodeStub = nil
	fake.registerNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) RegisterNodeReturnsOnCall(i int, result1 error) {
	fake.registerNodeMutex.Lock()
	defer fake.registerNodeMutex.Unlock()
	fake.RegisterNodeStub = nil
	if fake.registerNodeReturnsOnCall == nil {
		fake.registerNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) SetRTCNode(arg1 string, arg2 string) error {
	fake.setRTCNodeMutex.Lock()
	ret, specificReturn := fake.setRTCNodeReturnsOnCall[len(fake.setRTCNodeArgsForCall)]
	fake.setRTCNodeArgsForCall = append(fake.setRTCNodeArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.SetRTCNodeStub
	fakeReturns := fake.setRTCNodeReturns
	fake.recordInvocation("SetRTCNode", []interface{}{arg1, arg2})
	fake.setRTCNodeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) SetRTCNodeCallCount() int {
	fake.setRTCNodeMutex.RLock()
	defer fake.setRTCNodeMutex.RUnlock()
	return len(fake.setRTCNodeArgsForCall)
}

func (fake *FakeRouter) SetRTCNodeCalls(stub func(string, string) error) {
	fake.setRTCNodeMutex.Lock()
	defer fake.setRTCNodeMutex.Unlock()
	fake.SetRTCNodeStub = stub
}

func (fake *FakeRouter) SetRTCNodeArgsForCall(i int) (string, string) {
	fake.setRTCNodeMutex.RLock()
	defer fake.setRTCNodeMutex.RUnlock()
	argsForCall := fake.setRTCNodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRouter) SetRTCNodeReturns(result1 error) {
	fake.setRTCNodeMutex.Lock()
	defer fake.setRTCNodeMutex.Unlock()
	fake.SetRTCNodeStub = nil
	fake.setRTCNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) SetRTCNodeReturnsOnCall(i int, result1 error) {
	fake.setRTCNodeMutex.Lock()
	defer fake.setRTCNodeMutex.Unlock()
	fake.SetRTCNodeStub = nil
	if fake.setRTCNodeReturnsOnCall == nil {
		fake.setRTCNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setRTCNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) Start() error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
	}{})
	stub := fake.StartStub
	fakeReturns := fake.startReturns
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeRouter) StartCalls(stub func() error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeRouter) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) StartParticipant(arg1 string, arg2 string, arg3 string, arg4 string) error {
	fake.startParticipantMutex.Lock()
	ret, specificReturn := fake.startParticipantReturnsOnCall[len(fake.startParticipantArgsForCall)]
	fake.startParticipantArgsForCall = append(fake.startParticipantArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.StartParticipantStub
	fakeReturns := fake.startParticipantReturns
	fake.recordInvocation("StartParticipant", []interface{}{arg1, arg2, arg3, arg4})
	fake.startParticipantMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) StartParticipantCallCount() int {
	fake.startParticipantMutex.RLock()
	defer fake.startParticipantMutex.RUnlock()
	return len(fake.startParticipantArgsForCall)
}

func (fake *FakeRouter) StartParticipantCalls(stub func(string, string, string, string) error) {
	fake.startParticipantMutex.Lock()
	defer fake.startParticipantMutex.Unlock()
	fake.StartParticipantStub = stub
}

func (fake *FakeRouter) StartParticipantArgsForCall(i int) (string, string, string, string) {
	fake.startParticipantMutex.RLock()
	defer fake.startParticipantMutex.RUnlock()
	argsForCall := fake.startParticipantArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeRouter) StartParticipantReturns(result1 error) {
	fake.startParticipantMutex.Lock()
	defer fake.startParticipantMutex.Unlock()
	fake.StartParticipantStub = nil
	fake.startParticipantReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) StartParticipantReturnsOnCall(i int, result1 error) {
	fake.startParticipantMutex.Lock()
	defer fake.startParticipantMutex.Unlock()
	fake.StartParticipantStub = nil
	if fake.startParticipantReturnsOnCall == nil {
		fake.startParticipantReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startParticipantReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	stub := fake.StopStub
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if stub != nil {
		fake.StopStub()
	}
}

func (fake *FakeRouter) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeRouter) StopCalls(stub func()) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeRouter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getNodeMutex.RLock()
	defer fake.getNodeMutex.RUnlock()
	fake.getNodeIdForRoomMutex.RLock()
	defer fake.getNodeIdForRoomMutex.RUnlock()
	fake.getRequestSinkMutex.RLock()
	defer fake.getRequestSinkMutex.RUnlock()
	fake.getResponseSourceMutex.RLock()
	defer fake.getResponseSourceMutex.RUnlock()
	fake.onNewParticipantMutex.RLock()
	defer fake.onNewParticipantMutex.RUnlock()
	fake.registerNodeMutex.RLock()
	defer fake.registerNodeMutex.RUnlock()
	fake.setRTCNodeMutex.RLock()
	defer fake.setRTCNodeMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.startParticipantMutex.RLock()
	defer fake.startParticipantMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRouter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ routing.Router = new(FakeRouter)
