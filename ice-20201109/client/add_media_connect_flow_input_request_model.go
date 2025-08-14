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
}

type AddMediaConnectFlowInputRequest struct {
	// The IP address whitelist in CIDR format. Separate multiple CIDR blocks with commas (,).
	//
	// example:
	//
	// 19.168.1.1/32,18.168.1.1/16
	Cidrs *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty"`
	// The flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The source URL. This parameter is required when the source type is RTMP-PULL or SRT-Listener.
	//
	// example:
	//
	// rtmp://pull.test.alivecdn.com/live/alitest
	InputFromUrl *string `json:"InputFromUrl,omitempty" xml:"InputFromUrl,omitempty"`
	// The source name.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliTestInput
	InputName *string `json:"InputName,omitempty" xml:"InputName,omitempty"`
	// The source type.
	//
	// Valid values:
	//
	// 	- RTMP-PUSH
	//
	// 	- SRT-Caller
	//
	// 	- RTMP-PULL
	//
	// 	- SRT-Listener
	//
	// 	- Flow
	//
	// This parameter is required.
	//
	// example:
	//
	// RTMP-PUSH
	InputProtocol *string `json:"InputProtocol,omitempty" xml:"InputProtocol,omitempty"`
	// The maximum bitrate. Unit: bit/s.
	//
	// example:
	//
	// 2000000
	MaxBitrate *int32 `json:"MaxBitrate,omitempty" xml:"MaxBitrate,omitempty"`
	// The ID of the source flow. This parameter is required when the source type is Flow.
	//
	// example:
	//
	// 805fbdd0-575e-4146-b35d-ec7f63937b20
	PairFlowId *string `json:"PairFlowId,omitempty" xml:"PairFlowId,omitempty"`
	// The output of the source flow. This parameter is required when the source type is Flow.
	//
	// example:
	//
	// AliTestOutput
	PairOutputName *string `json:"PairOutputName,omitempty" xml:"PairOutputName,omitempty"`
	// The latency for the SRT stream. This parameter is required the source type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// 1000
	SrtLatency *int32 `json:"SrtLatency,omitempty" xml:"SrtLatency,omitempty"`
	// The SRT key. This parameter is required when the source type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// BETTERG08S01
	SrtPassphrase *string `json:"SrtPassphrase,omitempty" xml:"SrtPassphrase,omitempty"`
	// The encryption key length. This parameter is required when the source type is SRT-Listener or SRT-Caller.
	//
	// Valid values:
	//
	// 	- 0
	//
	// 	- 16
	//
	// 	- 24
	//
	// 	- 32
	//
	// example:
	//
	// 32
	SrtPbkeyLen *string `json:"SrtPbkeyLen,omitempty" xml:"SrtPbkeyLen,omitempty"`
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

func (s *AddMediaConnectFlowInputRequest) Validate() error {
	return dara.Validate(s)
}
