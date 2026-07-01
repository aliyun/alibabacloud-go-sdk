// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaConnectFlowInputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrs(v string) *AddMediaConnectFlowInputRequest
	GetCidrs() *string
	SetFlowId(v string) *AddMediaConnectFlowInputRequest
	GetFlowId() *string
	SetInputFromUrl(v string) *AddMediaConnectFlowInputRequest
	GetInputFromUrl() *string
	SetInputName(v string) *AddMediaConnectFlowInputRequest
	GetInputName() *string
	SetInputProtocol(v string) *AddMediaConnectFlowInputRequest
	GetInputProtocol() *string
	SetMaxBitrate(v int32) *AddMediaConnectFlowInputRequest
	GetMaxBitrate() *int32
	SetPairFlowId(v string) *AddMediaConnectFlowInputRequest
	GetPairFlowId() *string
	SetPairOutputName(v string) *AddMediaConnectFlowInputRequest
	GetPairOutputName() *string
	SetSrtLatency(v int32) *AddMediaConnectFlowInputRequest
	GetSrtLatency() *int32
	SetSrtPassphrase(v string) *AddMediaConnectFlowInputRequest
	GetSrtPassphrase() *string
	SetSrtPbkeyLen(v string) *AddMediaConnectFlowInputRequest
	GetSrtPbkeyLen() *string
	SetWithInternalVip(v string) *AddMediaConnectFlowInputRequest
	GetWithInternalVip() *string
}

type AddMediaConnectFlowInputRequest struct {
	// IP address whitelist in CIDR notation. Separate multiple CIDR blocks with commas.
	//
	// example:
	//
	// 19.168.1.1/32,18.168.1.1/16
	Cidrs *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty"`
	// Flow instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// Input URL. Required only when the input type is RTMP-PULL or SRT-Listener.
	//
	// example:
	//
	// rtmp://pull.test.alivecdn.com/live/alitest
	InputFromUrl *string `json:"InputFromUrl,omitempty" xml:"InputFromUrl,omitempty"`
	// Input name
	//
	// This parameter is required.
	//
	// example:
	//
	// AliTestInput
	InputName *string `json:"InputName,omitempty" xml:"InputName,omitempty"`
	// Input type
	//
	// This parameter is required.
	//
	// example:
	//
	// RTMP-PUSH
	InputProtocol *string `json:"InputProtocol,omitempty" xml:"InputProtocol,omitempty"`
	// Maximum bitrate in bits per second (bps)
	//
	// example:
	//
	// 2000000
	MaxBitrate *int32 `json:"MaxBitrate,omitempty" xml:"MaxBitrate,omitempty"`
	// Upstream Flow ID. Required only when the input type is Flow.
	//
	// example:
	//
	// 805fbdd0-575e-4146-b35d-ec7f63937b20
	PairFlowId *string `json:"PairFlowId,omitempty" xml:"PairFlowId,omitempty"`
	// Upstream Flow output name. Required only when the input type is Flow.
	//
	// example:
	//
	// AliTestOutput
	PairOutputName *string `json:"PairOutputName,omitempty" xml:"PairOutputName,omitempty"`
	// SRT latency in milliseconds. Required only when the input type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// 1000
	SrtLatency *int32 `json:"SrtLatency,omitempty" xml:"SrtLatency,omitempty"`
	// SRT encryption key. Required only when the input type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// BETTERG08S01
	SrtPassphrase *string `json:"SrtPassphrase,omitempty" xml:"SrtPassphrase,omitempty"`
	// SRT encryption key length in bytes. Required only when the input type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// 32
	SrtPbkeyLen     *string `json:"SrtPbkeyLen,omitempty" xml:"SrtPbkeyLen,omitempty"`
	WithInternalVip *string `json:"WithInternalVip,omitempty" xml:"WithInternalVip,omitempty"`
}

func (s AddMediaConnectFlowInputRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMediaConnectFlowInputRequest) GoString() string {
	return s.String()
}

func (s *AddMediaConnectFlowInputRequest) GetCidrs() *string {
	return s.Cidrs
}

func (s *AddMediaConnectFlowInputRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *AddMediaConnectFlowInputRequest) GetInputFromUrl() *string {
	return s.InputFromUrl
}

func (s *AddMediaConnectFlowInputRequest) GetInputName() *string {
	return s.InputName
}

func (s *AddMediaConnectFlowInputRequest) GetInputProtocol() *string {
	return s.InputProtocol
}

func (s *AddMediaConnectFlowInputRequest) GetMaxBitrate() *int32 {
	return s.MaxBitrate
}

func (s *AddMediaConnectFlowInputRequest) GetPairFlowId() *string {
	return s.PairFlowId
}

func (s *AddMediaConnectFlowInputRequest) GetPairOutputName() *string {
	return s.PairOutputName
}

func (s *AddMediaConnectFlowInputRequest) GetSrtLatency() *int32 {
	return s.SrtLatency
}

func (s *AddMediaConnectFlowInputRequest) GetSrtPassphrase() *string {
	return s.SrtPassphrase
}

func (s *AddMediaConnectFlowInputRequest) GetSrtPbkeyLen() *string {
	return s.SrtPbkeyLen
}

func (s *AddMediaConnectFlowInputRequest) GetWithInternalVip() *string {
	return s.WithInternalVip
}

func (s *AddMediaConnectFlowInputRequest) SetCidrs(v string) *AddMediaConnectFlowInputRequest {
	s.Cidrs = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) SetFlowId(v string) *AddMediaConnectFlowInputRequest {
	s.FlowId = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) SetInputFromUrl(v string) *AddMediaConnectFlowInputRequest {
	s.InputFromUrl = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) SetInputName(v string) *AddMediaConnectFlowInputRequest {
	s.InputName = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) SetInputProtocol(v string) *AddMediaConnectFlowInputRequest {
	s.InputProtocol = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) SetMaxBitrate(v int32) *AddMediaConnectFlowInputRequest {
	s.MaxBitrate = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) SetPairFlowId(v string) *AddMediaConnectFlowInputRequest {
	s.PairFlowId = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) SetPairOutputName(v string) *AddMediaConnectFlowInputRequest {
	s.PairOutputName = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) SetSrtLatency(v int32) *AddMediaConnectFlowInputRequest {
	s.SrtLatency = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) SetSrtPassphrase(v string) *AddMediaConnectFlowInputRequest {
	s.SrtPassphrase = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) SetSrtPbkeyLen(v string) *AddMediaConnectFlowInputRequest {
	s.SrtPbkeyLen = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) SetWithInternalVip(v string) *AddMediaConnectFlowInputRequest {
	s.WithInternalVip = &v
	return s
}

func (s *AddMediaConnectFlowInputRequest) Validate() error {
	return dara.Validate(s)
}
