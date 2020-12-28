// Code generated by counterfeiter. DO NOT EDIT.
package rtcfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc"
	"github.com/pion/rtcp"
	webrtc "github.com/pion/webrtc/v3"
)

type FakePeerConnection struct {
	AddICECandidateStub        func(webrtc.ICECandidateInit) error
	addICECandidateMutex       sync.RWMutex
	addICECandidateArgsForCall []struct {
		arg1 webrtc.ICECandidateInit
	}
	addICECandidateReturns struct {
		result1 error
	}
	addICECandidateReturnsOnCall map[int]struct {
		result1 error
	}
	AddTransceiverFromTrackStub        func(webrtc.TrackLocal, ...webrtc.RTPTransceiverInit) (*webrtc.RTPTransceiver, error)
	addTransceiverFromTrackMutex       sync.RWMutex
	addTransceiverFromTrackArgsForCall []struct {
		arg1 webrtc.TrackLocal
		arg2 []webrtc.RTPTransceiverInit
	}
	addTransceiverFromTrackReturns struct {
		result1 *webrtc.RTPTransceiver
		result2 error
	}
	addTransceiverFromTrackReturnsOnCall map[int]struct {
		result1 *webrtc.RTPTransceiver
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	ConnectionStateStub        func() webrtc.PeerConnectionState
	connectionStateMutex       sync.RWMutex
	connectionStateArgsForCall []struct {
	}
	connectionStateReturns struct {
		result1 webrtc.PeerConnectionState
	}
	connectionStateReturnsOnCall map[int]struct {
		result1 webrtc.PeerConnectionState
	}
	CreateAnswerStub        func(*webrtc.AnswerOptions) (webrtc.SessionDescription, error)
	createAnswerMutex       sync.RWMutex
	createAnswerArgsForCall []struct {
		arg1 *webrtc.AnswerOptions
	}
	createAnswerReturns struct {
		result1 webrtc.SessionDescription
		result2 error
	}
	createAnswerReturnsOnCall map[int]struct {
		result1 webrtc.SessionDescription
		result2 error
	}
	CreateDataChannelStub        func(string, *webrtc.DataChannelInit) (*webrtc.DataChannel, error)
	createDataChannelMutex       sync.RWMutex
	createDataChannelArgsForCall []struct {
		arg1 string
		arg2 *webrtc.DataChannelInit
	}
	createDataChannelReturns struct {
		result1 *webrtc.DataChannel
		result2 error
	}
	createDataChannelReturnsOnCall map[int]struct {
		result1 *webrtc.DataChannel
		result2 error
	}
	CreateOfferStub        func(*webrtc.OfferOptions) (webrtc.SessionDescription, error)
	createOfferMutex       sync.RWMutex
	createOfferArgsForCall []struct {
		arg1 *webrtc.OfferOptions
	}
	createOfferReturns struct {
		result1 webrtc.SessionDescription
		result2 error
	}
	createOfferReturnsOnCall map[int]struct {
		result1 webrtc.SessionDescription
		result2 error
	}
	OnDataChannelStub        func(func(d *webrtc.DataChannel))
	onDataChannelMutex       sync.RWMutex
	onDataChannelArgsForCall []struct {
		arg1 func(d *webrtc.DataChannel)
	}
	OnICECandidateStub        func(func(*webrtc.ICECandidate))
	onICECandidateMutex       sync.RWMutex
	onICECandidateArgsForCall []struct {
		arg1 func(*webrtc.ICECandidate)
	}
	OnICEConnectionStateChangeStub        func(func(webrtc.ICEConnectionState))
	onICEConnectionStateChangeMutex       sync.RWMutex
	onICEConnectionStateChangeArgsForCall []struct {
		arg1 func(webrtc.ICEConnectionState)
	}
	OnNegotiationNeededStub        func(func())
	onNegotiationNeededMutex       sync.RWMutex
	onNegotiationNeededArgsForCall []struct {
		arg1 func()
	}
	OnTrackStub        func(func(*webrtc.TrackRemote, *webrtc.RTPReceiver))
	onTrackMutex       sync.RWMutex
	onTrackArgsForCall []struct {
		arg1 func(*webrtc.TrackRemote, *webrtc.RTPReceiver)
	}
	RemoveTrackStub        func(*webrtc.RTPSender) error
	removeTrackMutex       sync.RWMutex
	removeTrackArgsForCall []struct {
		arg1 *webrtc.RTPSender
	}
	removeTrackReturns struct {
		result1 error
	}
	removeTrackReturnsOnCall map[int]struct {
		result1 error
	}
	SetLocalDescriptionStub        func(webrtc.SessionDescription) error
	setLocalDescriptionMutex       sync.RWMutex
	setLocalDescriptionArgsForCall []struct {
		arg1 webrtc.SessionDescription
	}
	setLocalDescriptionReturns struct {
		result1 error
	}
	setLocalDescriptionReturnsOnCall map[int]struct {
		result1 error
	}
	SetRemoteDescriptionStub        func(webrtc.SessionDescription) error
	setRemoteDescriptionMutex       sync.RWMutex
	setRemoteDescriptionArgsForCall []struct {
		arg1 webrtc.SessionDescription
	}
	setRemoteDescriptionReturns struct {
		result1 error
	}
	setRemoteDescriptionReturnsOnCall map[int]struct {
		result1 error
	}
	WriteRTCPStub        func([]rtcp.Packet) error
	writeRTCPMutex       sync.RWMutex
	writeRTCPArgsForCall []struct {
		arg1 []rtcp.Packet
	}
	writeRTCPReturns struct {
		result1 error
	}
	writeRTCPReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePeerConnection) AddICECandidate(arg1 webrtc.ICECandidateInit) error {
	fake.addICECandidateMutex.Lock()
	ret, specificReturn := fake.addICECandidateReturnsOnCall[len(fake.addICECandidateArgsForCall)]
	fake.addICECandidateArgsForCall = append(fake.addICECandidateArgsForCall, struct {
		arg1 webrtc.ICECandidateInit
	}{arg1})
	stub := fake.AddICECandidateStub
	fakeReturns := fake.addICECandidateReturns
	fake.recordInvocation("AddICECandidate", []interface{}{arg1})
	fake.addICECandidateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePeerConnection) AddICECandidateCallCount() int {
	fake.addICECandidateMutex.RLock()
	defer fake.addICECandidateMutex.RUnlock()
	return len(fake.addICECandidateArgsForCall)
}

func (fake *FakePeerConnection) AddICECandidateCalls(stub func(webrtc.ICECandidateInit) error) {
	fake.addICECandidateMutex.Lock()
	defer fake.addICECandidateMutex.Unlock()
	fake.AddICECandidateStub = stub
}

func (fake *FakePeerConnection) AddICECandidateArgsForCall(i int) webrtc.ICECandidateInit {
	fake.addICECandidateMutex.RLock()
	defer fake.addICECandidateMutex.RUnlock()
	argsForCall := fake.addICECandidateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) AddICECandidateReturns(result1 error) {
	fake.addICECandidateMutex.Lock()
	defer fake.addICECandidateMutex.Unlock()
	fake.AddICECandidateStub = nil
	fake.addICECandidateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) AddICECandidateReturnsOnCall(i int, result1 error) {
	fake.addICECandidateMutex.Lock()
	defer fake.addICECandidateMutex.Unlock()
	fake.AddICECandidateStub = nil
	if fake.addICECandidateReturnsOnCall == nil {
		fake.addICECandidateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addICECandidateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) AddTransceiverFromTrack(arg1 webrtc.TrackLocal, arg2 ...webrtc.RTPTransceiverInit) (*webrtc.RTPTransceiver, error) {
	fake.addTransceiverFromTrackMutex.Lock()
	ret, specificReturn := fake.addTransceiverFromTrackReturnsOnCall[len(fake.addTransceiverFromTrackArgsForCall)]
	fake.addTransceiverFromTrackArgsForCall = append(fake.addTransceiverFromTrackArgsForCall, struct {
		arg1 webrtc.TrackLocal
		arg2 []webrtc.RTPTransceiverInit
	}{arg1, arg2})
	stub := fake.AddTransceiverFromTrackStub
	fakeReturns := fake.addTransceiverFromTrackReturns
	fake.recordInvocation("AddTransceiverFromTrack", []interface{}{arg1, arg2})
	fake.addTransceiverFromTrackMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePeerConnection) AddTransceiverFromTrackCallCount() int {
	fake.addTransceiverFromTrackMutex.RLock()
	defer fake.addTransceiverFromTrackMutex.RUnlock()
	return len(fake.addTransceiverFromTrackArgsForCall)
}

func (fake *FakePeerConnection) AddTransceiverFromTrackCalls(stub func(webrtc.TrackLocal, ...webrtc.RTPTransceiverInit) (*webrtc.RTPTransceiver, error)) {
	fake.addTransceiverFromTrackMutex.Lock()
	defer fake.addTransceiverFromTrackMutex.Unlock()
	fake.AddTransceiverFromTrackStub = stub
}

func (fake *FakePeerConnection) AddTransceiverFromTrackArgsForCall(i int) (webrtc.TrackLocal, []webrtc.RTPTransceiverInit) {
	fake.addTransceiverFromTrackMutex.RLock()
	defer fake.addTransceiverFromTrackMutex.RUnlock()
	argsForCall := fake.addTransceiverFromTrackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePeerConnection) AddTransceiverFromTrackReturns(result1 *webrtc.RTPTransceiver, result2 error) {
	fake.addTransceiverFromTrackMutex.Lock()
	defer fake.addTransceiverFromTrackMutex.Unlock()
	fake.AddTransceiverFromTrackStub = nil
	fake.addTransceiverFromTrackReturns = struct {
		result1 *webrtc.RTPTransceiver
		result2 error
	}{result1, result2}
}

func (fake *FakePeerConnection) AddTransceiverFromTrackReturnsOnCall(i int, result1 *webrtc.RTPTransceiver, result2 error) {
	fake.addTransceiverFromTrackMutex.Lock()
	defer fake.addTransceiverFromTrackMutex.Unlock()
	fake.AddTransceiverFromTrackStub = nil
	if fake.addTransceiverFromTrackReturnsOnCall == nil {
		fake.addTransceiverFromTrackReturnsOnCall = make(map[int]struct {
			result1 *webrtc.RTPTransceiver
			result2 error
		})
	}
	fake.addTransceiverFromTrackReturnsOnCall[i] = struct {
		result1 *webrtc.RTPTransceiver
		result2 error
	}{result1, result2}
}

func (fake *FakePeerConnection) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePeerConnection) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakePeerConnection) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakePeerConnection) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) ConnectionState() webrtc.PeerConnectionState {
	fake.connectionStateMutex.Lock()
	ret, specificReturn := fake.connectionStateReturnsOnCall[len(fake.connectionStateArgsForCall)]
	fake.connectionStateArgsForCall = append(fake.connectionStateArgsForCall, struct {
	}{})
	stub := fake.ConnectionStateStub
	fakeReturns := fake.connectionStateReturns
	fake.recordInvocation("ConnectionState", []interface{}{})
	fake.connectionStateMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePeerConnection) ConnectionStateCallCount() int {
	fake.connectionStateMutex.RLock()
	defer fake.connectionStateMutex.RUnlock()
	return len(fake.connectionStateArgsForCall)
}

func (fake *FakePeerConnection) ConnectionStateCalls(stub func() webrtc.PeerConnectionState) {
	fake.connectionStateMutex.Lock()
	defer fake.connectionStateMutex.Unlock()
	fake.ConnectionStateStub = stub
}

func (fake *FakePeerConnection) ConnectionStateReturns(result1 webrtc.PeerConnectionState) {
	fake.connectionStateMutex.Lock()
	defer fake.connectionStateMutex.Unlock()
	fake.ConnectionStateStub = nil
	fake.connectionStateReturns = struct {
		result1 webrtc.PeerConnectionState
	}{result1}
}

func (fake *FakePeerConnection) ConnectionStateReturnsOnCall(i int, result1 webrtc.PeerConnectionState) {
	fake.connectionStateMutex.Lock()
	defer fake.connectionStateMutex.Unlock()
	fake.ConnectionStateStub = nil
	if fake.connectionStateReturnsOnCall == nil {
		fake.connectionStateReturnsOnCall = make(map[int]struct {
			result1 webrtc.PeerConnectionState
		})
	}
	fake.connectionStateReturnsOnCall[i] = struct {
		result1 webrtc.PeerConnectionState
	}{result1}
}

func (fake *FakePeerConnection) CreateAnswer(arg1 *webrtc.AnswerOptions) (webrtc.SessionDescription, error) {
	fake.createAnswerMutex.Lock()
	ret, specificReturn := fake.createAnswerReturnsOnCall[len(fake.createAnswerArgsForCall)]
	fake.createAnswerArgsForCall = append(fake.createAnswerArgsForCall, struct {
		arg1 *webrtc.AnswerOptions
	}{arg1})
	stub := fake.CreateAnswerStub
	fakeReturns := fake.createAnswerReturns
	fake.recordInvocation("CreateAnswer", []interface{}{arg1})
	fake.createAnswerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePeerConnection) CreateAnswerCallCount() int {
	fake.createAnswerMutex.RLock()
	defer fake.createAnswerMutex.RUnlock()
	return len(fake.createAnswerArgsForCall)
}

func (fake *FakePeerConnection) CreateAnswerCalls(stub func(*webrtc.AnswerOptions) (webrtc.SessionDescription, error)) {
	fake.createAnswerMutex.Lock()
	defer fake.createAnswerMutex.Unlock()
	fake.CreateAnswerStub = stub
}

func (fake *FakePeerConnection) CreateAnswerArgsForCall(i int) *webrtc.AnswerOptions {
	fake.createAnswerMutex.RLock()
	defer fake.createAnswerMutex.RUnlock()
	argsForCall := fake.createAnswerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) CreateAnswerReturns(result1 webrtc.SessionDescription, result2 error) {
	fake.createAnswerMutex.Lock()
	defer fake.createAnswerMutex.Unlock()
	fake.CreateAnswerStub = nil
	fake.createAnswerReturns = struct {
		result1 webrtc.SessionDescription
		result2 error
	}{result1, result2}
}

func (fake *FakePeerConnection) CreateAnswerReturnsOnCall(i int, result1 webrtc.SessionDescription, result2 error) {
	fake.createAnswerMutex.Lock()
	defer fake.createAnswerMutex.Unlock()
	fake.CreateAnswerStub = nil
	if fake.createAnswerReturnsOnCall == nil {
		fake.createAnswerReturnsOnCall = make(map[int]struct {
			result1 webrtc.SessionDescription
			result2 error
		})
	}
	fake.createAnswerReturnsOnCall[i] = struct {
		result1 webrtc.SessionDescription
		result2 error
	}{result1, result2}
}

func (fake *FakePeerConnection) CreateDataChannel(arg1 string, arg2 *webrtc.DataChannelInit) (*webrtc.DataChannel, error) {
	fake.createDataChannelMutex.Lock()
	ret, specificReturn := fake.createDataChannelReturnsOnCall[len(fake.createDataChannelArgsForCall)]
	fake.createDataChannelArgsForCall = append(fake.createDataChannelArgsForCall, struct {
		arg1 string
		arg2 *webrtc.DataChannelInit
	}{arg1, arg2})
	stub := fake.CreateDataChannelStub
	fakeReturns := fake.createDataChannelReturns
	fake.recordInvocation("CreateDataChannel", []interface{}{arg1, arg2})
	fake.createDataChannelMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePeerConnection) CreateDataChannelCallCount() int {
	fake.createDataChannelMutex.RLock()
	defer fake.createDataChannelMutex.RUnlock()
	return len(fake.createDataChannelArgsForCall)
}

func (fake *FakePeerConnection) CreateDataChannelCalls(stub func(string, *webrtc.DataChannelInit) (*webrtc.DataChannel, error)) {
	fake.createDataChannelMutex.Lock()
	defer fake.createDataChannelMutex.Unlock()
	fake.CreateDataChannelStub = stub
}

func (fake *FakePeerConnection) CreateDataChannelArgsForCall(i int) (string, *webrtc.DataChannelInit) {
	fake.createDataChannelMutex.RLock()
	defer fake.createDataChannelMutex.RUnlock()
	argsForCall := fake.createDataChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePeerConnection) CreateDataChannelReturns(result1 *webrtc.DataChannel, result2 error) {
	fake.createDataChannelMutex.Lock()
	defer fake.createDataChannelMutex.Unlock()
	fake.CreateDataChannelStub = nil
	fake.createDataChannelReturns = struct {
		result1 *webrtc.DataChannel
		result2 error
	}{result1, result2}
}

func (fake *FakePeerConnection) CreateDataChannelReturnsOnCall(i int, result1 *webrtc.DataChannel, result2 error) {
	fake.createDataChannelMutex.Lock()
	defer fake.createDataChannelMutex.Unlock()
	fake.CreateDataChannelStub = nil
	if fake.createDataChannelReturnsOnCall == nil {
		fake.createDataChannelReturnsOnCall = make(map[int]struct {
			result1 *webrtc.DataChannel
			result2 error
		})
	}
	fake.createDataChannelReturnsOnCall[i] = struct {
		result1 *webrtc.DataChannel
		result2 error
	}{result1, result2}
}

func (fake *FakePeerConnection) CreateOffer(arg1 *webrtc.OfferOptions) (webrtc.SessionDescription, error) {
	fake.createOfferMutex.Lock()
	ret, specificReturn := fake.createOfferReturnsOnCall[len(fake.createOfferArgsForCall)]
	fake.createOfferArgsForCall = append(fake.createOfferArgsForCall, struct {
		arg1 *webrtc.OfferOptions
	}{arg1})
	stub := fake.CreateOfferStub
	fakeReturns := fake.createOfferReturns
	fake.recordInvocation("CreateOffer", []interface{}{arg1})
	fake.createOfferMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePeerConnection) CreateOfferCallCount() int {
	fake.createOfferMutex.RLock()
	defer fake.createOfferMutex.RUnlock()
	return len(fake.createOfferArgsForCall)
}

func (fake *FakePeerConnection) CreateOfferCalls(stub func(*webrtc.OfferOptions) (webrtc.SessionDescription, error)) {
	fake.createOfferMutex.Lock()
	defer fake.createOfferMutex.Unlock()
	fake.CreateOfferStub = stub
}

func (fake *FakePeerConnection) CreateOfferArgsForCall(i int) *webrtc.OfferOptions {
	fake.createOfferMutex.RLock()
	defer fake.createOfferMutex.RUnlock()
	argsForCall := fake.createOfferArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) CreateOfferReturns(result1 webrtc.SessionDescription, result2 error) {
	fake.createOfferMutex.Lock()
	defer fake.createOfferMutex.Unlock()
	fake.CreateOfferStub = nil
	fake.createOfferReturns = struct {
		result1 webrtc.SessionDescription
		result2 error
	}{result1, result2}
}

func (fake *FakePeerConnection) CreateOfferReturnsOnCall(i int, result1 webrtc.SessionDescription, result2 error) {
	fake.createOfferMutex.Lock()
	defer fake.createOfferMutex.Unlock()
	fake.CreateOfferStub = nil
	if fake.createOfferReturnsOnCall == nil {
		fake.createOfferReturnsOnCall = make(map[int]struct {
			result1 webrtc.SessionDescription
			result2 error
		})
	}
	fake.createOfferReturnsOnCall[i] = struct {
		result1 webrtc.SessionDescription
		result2 error
	}{result1, result2}
}

func (fake *FakePeerConnection) OnDataChannel(arg1 func(d *webrtc.DataChannel)) {
	fake.onDataChannelMutex.Lock()
	fake.onDataChannelArgsForCall = append(fake.onDataChannelArgsForCall, struct {
		arg1 func(d *webrtc.DataChannel)
	}{arg1})
	stub := fake.OnDataChannelStub
	fake.recordInvocation("OnDataChannel", []interface{}{arg1})
	fake.onDataChannelMutex.Unlock()
	if stub != nil {
		fake.OnDataChannelStub(arg1)
	}
}

func (fake *FakePeerConnection) OnDataChannelCallCount() int {
	fake.onDataChannelMutex.RLock()
	defer fake.onDataChannelMutex.RUnlock()
	return len(fake.onDataChannelArgsForCall)
}

func (fake *FakePeerConnection) OnDataChannelCalls(stub func(func(d *webrtc.DataChannel))) {
	fake.onDataChannelMutex.Lock()
	defer fake.onDataChannelMutex.Unlock()
	fake.OnDataChannelStub = stub
}

func (fake *FakePeerConnection) OnDataChannelArgsForCall(i int) func(d *webrtc.DataChannel) {
	fake.onDataChannelMutex.RLock()
	defer fake.onDataChannelMutex.RUnlock()
	argsForCall := fake.onDataChannelArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) OnICECandidate(arg1 func(*webrtc.ICECandidate)) {
	fake.onICECandidateMutex.Lock()
	fake.onICECandidateArgsForCall = append(fake.onICECandidateArgsForCall, struct {
		arg1 func(*webrtc.ICECandidate)
	}{arg1})
	stub := fake.OnICECandidateStub
	fake.recordInvocation("OnICECandidate", []interface{}{arg1})
	fake.onICECandidateMutex.Unlock()
	if stub != nil {
		fake.OnICECandidateStub(arg1)
	}
}

func (fake *FakePeerConnection) OnICECandidateCallCount() int {
	fake.onICECandidateMutex.RLock()
	defer fake.onICECandidateMutex.RUnlock()
	return len(fake.onICECandidateArgsForCall)
}

func (fake *FakePeerConnection) OnICECandidateCalls(stub func(func(*webrtc.ICECandidate))) {
	fake.onICECandidateMutex.Lock()
	defer fake.onICECandidateMutex.Unlock()
	fake.OnICECandidateStub = stub
}

func (fake *FakePeerConnection) OnICECandidateArgsForCall(i int) func(*webrtc.ICECandidate) {
	fake.onICECandidateMutex.RLock()
	defer fake.onICECandidateMutex.RUnlock()
	argsForCall := fake.onICECandidateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) OnICEConnectionStateChange(arg1 func(webrtc.ICEConnectionState)) {
	fake.onICEConnectionStateChangeMutex.Lock()
	fake.onICEConnectionStateChangeArgsForCall = append(fake.onICEConnectionStateChangeArgsForCall, struct {
		arg1 func(webrtc.ICEConnectionState)
	}{arg1})
	stub := fake.OnICEConnectionStateChangeStub
	fake.recordInvocation("OnICEConnectionStateChange", []interface{}{arg1})
	fake.onICEConnectionStateChangeMutex.Unlock()
	if stub != nil {
		fake.OnICEConnectionStateChangeStub(arg1)
	}
}

func (fake *FakePeerConnection) OnICEConnectionStateChangeCallCount() int {
	fake.onICEConnectionStateChangeMutex.RLock()
	defer fake.onICEConnectionStateChangeMutex.RUnlock()
	return len(fake.onICEConnectionStateChangeArgsForCall)
}

func (fake *FakePeerConnection) OnICEConnectionStateChangeCalls(stub func(func(webrtc.ICEConnectionState))) {
	fake.onICEConnectionStateChangeMutex.Lock()
	defer fake.onICEConnectionStateChangeMutex.Unlock()
	fake.OnICEConnectionStateChangeStub = stub
}

func (fake *FakePeerConnection) OnICEConnectionStateChangeArgsForCall(i int) func(webrtc.ICEConnectionState) {
	fake.onICEConnectionStateChangeMutex.RLock()
	defer fake.onICEConnectionStateChangeMutex.RUnlock()
	argsForCall := fake.onICEConnectionStateChangeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) OnNegotiationNeeded(arg1 func()) {
	fake.onNegotiationNeededMutex.Lock()
	fake.onNegotiationNeededArgsForCall = append(fake.onNegotiationNeededArgsForCall, struct {
		arg1 func()
	}{arg1})
	stub := fake.OnNegotiationNeededStub
	fake.recordInvocation("OnNegotiationNeeded", []interface{}{arg1})
	fake.onNegotiationNeededMutex.Unlock()
	if stub != nil {
		fake.OnNegotiationNeededStub(arg1)
	}
}

func (fake *FakePeerConnection) OnNegotiationNeededCallCount() int {
	fake.onNegotiationNeededMutex.RLock()
	defer fake.onNegotiationNeededMutex.RUnlock()
	return len(fake.onNegotiationNeededArgsForCall)
}

func (fake *FakePeerConnection) OnNegotiationNeededCalls(stub func(func())) {
	fake.onNegotiationNeededMutex.Lock()
	defer fake.onNegotiationNeededMutex.Unlock()
	fake.OnNegotiationNeededStub = stub
}

func (fake *FakePeerConnection) OnNegotiationNeededArgsForCall(i int) func() {
	fake.onNegotiationNeededMutex.RLock()
	defer fake.onNegotiationNeededMutex.RUnlock()
	argsForCall := fake.onNegotiationNeededArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) OnTrack(arg1 func(*webrtc.TrackRemote, *webrtc.RTPReceiver)) {
	fake.onTrackMutex.Lock()
	fake.onTrackArgsForCall = append(fake.onTrackArgsForCall, struct {
		arg1 func(*webrtc.TrackRemote, *webrtc.RTPReceiver)
	}{arg1})
	stub := fake.OnTrackStub
	fake.recordInvocation("OnTrack", []interface{}{arg1})
	fake.onTrackMutex.Unlock()
	if stub != nil {
		fake.OnTrackStub(arg1)
	}
}

func (fake *FakePeerConnection) OnTrackCallCount() int {
	fake.onTrackMutex.RLock()
	defer fake.onTrackMutex.RUnlock()
	return len(fake.onTrackArgsForCall)
}

func (fake *FakePeerConnection) OnTrackCalls(stub func(func(*webrtc.TrackRemote, *webrtc.RTPReceiver))) {
	fake.onTrackMutex.Lock()
	defer fake.onTrackMutex.Unlock()
	fake.OnTrackStub = stub
}

func (fake *FakePeerConnection) OnTrackArgsForCall(i int) func(*webrtc.TrackRemote, *webrtc.RTPReceiver) {
	fake.onTrackMutex.RLock()
	defer fake.onTrackMutex.RUnlock()
	argsForCall := fake.onTrackArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) RemoveTrack(arg1 *webrtc.RTPSender) error {
	fake.removeTrackMutex.Lock()
	ret, specificReturn := fake.removeTrackReturnsOnCall[len(fake.removeTrackArgsForCall)]
	fake.removeTrackArgsForCall = append(fake.removeTrackArgsForCall, struct {
		arg1 *webrtc.RTPSender
	}{arg1})
	stub := fake.RemoveTrackStub
	fakeReturns := fake.removeTrackReturns
	fake.recordInvocation("RemoveTrack", []interface{}{arg1})
	fake.removeTrackMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePeerConnection) RemoveTrackCallCount() int {
	fake.removeTrackMutex.RLock()
	defer fake.removeTrackMutex.RUnlock()
	return len(fake.removeTrackArgsForCall)
}

func (fake *FakePeerConnection) RemoveTrackCalls(stub func(*webrtc.RTPSender) error) {
	fake.removeTrackMutex.Lock()
	defer fake.removeTrackMutex.Unlock()
	fake.RemoveTrackStub = stub
}

func (fake *FakePeerConnection) RemoveTrackArgsForCall(i int) *webrtc.RTPSender {
	fake.removeTrackMutex.RLock()
	defer fake.removeTrackMutex.RUnlock()
	argsForCall := fake.removeTrackArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) RemoveTrackReturns(result1 error) {
	fake.removeTrackMutex.Lock()
	defer fake.removeTrackMutex.Unlock()
	fake.RemoveTrackStub = nil
	fake.removeTrackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) RemoveTrackReturnsOnCall(i int, result1 error) {
	fake.removeTrackMutex.Lock()
	defer fake.removeTrackMutex.Unlock()
	fake.RemoveTrackStub = nil
	if fake.removeTrackReturnsOnCall == nil {
		fake.removeTrackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeTrackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) SetLocalDescription(arg1 webrtc.SessionDescription) error {
	fake.setLocalDescriptionMutex.Lock()
	ret, specificReturn := fake.setLocalDescriptionReturnsOnCall[len(fake.setLocalDescriptionArgsForCall)]
	fake.setLocalDescriptionArgsForCall = append(fake.setLocalDescriptionArgsForCall, struct {
		arg1 webrtc.SessionDescription
	}{arg1})
	stub := fake.SetLocalDescriptionStub
	fakeReturns := fake.setLocalDescriptionReturns
	fake.recordInvocation("SetLocalDescription", []interface{}{arg1})
	fake.setLocalDescriptionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePeerConnection) SetLocalDescriptionCallCount() int {
	fake.setLocalDescriptionMutex.RLock()
	defer fake.setLocalDescriptionMutex.RUnlock()
	return len(fake.setLocalDescriptionArgsForCall)
}

func (fake *FakePeerConnection) SetLocalDescriptionCalls(stub func(webrtc.SessionDescription) error) {
	fake.setLocalDescriptionMutex.Lock()
	defer fake.setLocalDescriptionMutex.Unlock()
	fake.SetLocalDescriptionStub = stub
}

func (fake *FakePeerConnection) SetLocalDescriptionArgsForCall(i int) webrtc.SessionDescription {
	fake.setLocalDescriptionMutex.RLock()
	defer fake.setLocalDescriptionMutex.RUnlock()
	argsForCall := fake.setLocalDescriptionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) SetLocalDescriptionReturns(result1 error) {
	fake.setLocalDescriptionMutex.Lock()
	defer fake.setLocalDescriptionMutex.Unlock()
	fake.SetLocalDescriptionStub = nil
	fake.setLocalDescriptionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) SetLocalDescriptionReturnsOnCall(i int, result1 error) {
	fake.setLocalDescriptionMutex.Lock()
	defer fake.setLocalDescriptionMutex.Unlock()
	fake.SetLocalDescriptionStub = nil
	if fake.setLocalDescriptionReturnsOnCall == nil {
		fake.setLocalDescriptionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setLocalDescriptionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) SetRemoteDescription(arg1 webrtc.SessionDescription) error {
	fake.setRemoteDescriptionMutex.Lock()
	ret, specificReturn := fake.setRemoteDescriptionReturnsOnCall[len(fake.setRemoteDescriptionArgsForCall)]
	fake.setRemoteDescriptionArgsForCall = append(fake.setRemoteDescriptionArgsForCall, struct {
		arg1 webrtc.SessionDescription
	}{arg1})
	stub := fake.SetRemoteDescriptionStub
	fakeReturns := fake.setRemoteDescriptionReturns
	fake.recordInvocation("SetRemoteDescription", []interface{}{arg1})
	fake.setRemoteDescriptionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePeerConnection) SetRemoteDescriptionCallCount() int {
	fake.setRemoteDescriptionMutex.RLock()
	defer fake.setRemoteDescriptionMutex.RUnlock()
	return len(fake.setRemoteDescriptionArgsForCall)
}

func (fake *FakePeerConnection) SetRemoteDescriptionCalls(stub func(webrtc.SessionDescription) error) {
	fake.setRemoteDescriptionMutex.Lock()
	defer fake.setRemoteDescriptionMutex.Unlock()
	fake.SetRemoteDescriptionStub = stub
}

func (fake *FakePeerConnection) SetRemoteDescriptionArgsForCall(i int) webrtc.SessionDescription {
	fake.setRemoteDescriptionMutex.RLock()
	defer fake.setRemoteDescriptionMutex.RUnlock()
	argsForCall := fake.setRemoteDescriptionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) SetRemoteDescriptionReturns(result1 error) {
	fake.setRemoteDescriptionMutex.Lock()
	defer fake.setRemoteDescriptionMutex.Unlock()
	fake.SetRemoteDescriptionStub = nil
	fake.setRemoteDescriptionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) SetRemoteDescriptionReturnsOnCall(i int, result1 error) {
	fake.setRemoteDescriptionMutex.Lock()
	defer fake.setRemoteDescriptionMutex.Unlock()
	fake.SetRemoteDescriptionStub = nil
	if fake.setRemoteDescriptionReturnsOnCall == nil {
		fake.setRemoteDescriptionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setRemoteDescriptionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) WriteRTCP(arg1 []rtcp.Packet) error {
	var arg1Copy []rtcp.Packet
	if arg1 != nil {
		arg1Copy = make([]rtcp.Packet, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.writeRTCPMutex.Lock()
	ret, specificReturn := fake.writeRTCPReturnsOnCall[len(fake.writeRTCPArgsForCall)]
	fake.writeRTCPArgsForCall = append(fake.writeRTCPArgsForCall, struct {
		arg1 []rtcp.Packet
	}{arg1Copy})
	stub := fake.WriteRTCPStub
	fakeReturns := fake.writeRTCPReturns
	fake.recordInvocation("WriteRTCP", []interface{}{arg1Copy})
	fake.writeRTCPMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePeerConnection) WriteRTCPCallCount() int {
	fake.writeRTCPMutex.RLock()
	defer fake.writeRTCPMutex.RUnlock()
	return len(fake.writeRTCPArgsForCall)
}

func (fake *FakePeerConnection) WriteRTCPCalls(stub func([]rtcp.Packet) error) {
	fake.writeRTCPMutex.Lock()
	defer fake.writeRTCPMutex.Unlock()
	fake.WriteRTCPStub = stub
}

func (fake *FakePeerConnection) WriteRTCPArgsForCall(i int) []rtcp.Packet {
	fake.writeRTCPMutex.RLock()
	defer fake.writeRTCPMutex.RUnlock()
	argsForCall := fake.writeRTCPArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePeerConnection) WriteRTCPReturns(result1 error) {
	fake.writeRTCPMutex.Lock()
	defer fake.writeRTCPMutex.Unlock()
	fake.WriteRTCPStub = nil
	fake.writeRTCPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) WriteRTCPReturnsOnCall(i int, result1 error) {
	fake.writeRTCPMutex.Lock()
	defer fake.writeRTCPMutex.Unlock()
	fake.WriteRTCPStub = nil
	if fake.writeRTCPReturnsOnCall == nil {
		fake.writeRTCPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeRTCPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePeerConnection) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addICECandidateMutex.RLock()
	defer fake.addICECandidateMutex.RUnlock()
	fake.addTransceiverFromTrackMutex.RLock()
	defer fake.addTransceiverFromTrackMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.connectionStateMutex.RLock()
	defer fake.connectionStateMutex.RUnlock()
	fake.createAnswerMutex.RLock()
	defer fake.createAnswerMutex.RUnlock()
	fake.createDataChannelMutex.RLock()
	defer fake.createDataChannelMutex.RUnlock()
	fake.createOfferMutex.RLock()
	defer fake.createOfferMutex.RUnlock()
	fake.onDataChannelMutex.RLock()
	defer fake.onDataChannelMutex.RUnlock()
	fake.onICECandidateMutex.RLock()
	defer fake.onICECandidateMutex.RUnlock()
	fake.onICEConnectionStateChangeMutex.RLock()
	defer fake.onICEConnectionStateChangeMutex.RUnlock()
	fake.onNegotiationNeededMutex.RLock()
	defer fake.onNegotiationNeededMutex.RUnlock()
	fake.onTrackMutex.RLock()
	defer fake.onTrackMutex.RUnlock()
	fake.removeTrackMutex.RLock()
	defer fake.removeTrackMutex.RUnlock()
	fake.setLocalDescriptionMutex.RLock()
	defer fake.setLocalDescriptionMutex.RUnlock()
	fake.setRemoteDescriptionMutex.RLock()
	defer fake.setRemoteDescriptionMutex.RUnlock()
	fake.writeRTCPMutex.RLock()
	defer fake.writeRTCPMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePeerConnection) recordInvocation(key string, args []interface{}) {
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

var _ rtc.PeerConnection = new(FakePeerConnection)