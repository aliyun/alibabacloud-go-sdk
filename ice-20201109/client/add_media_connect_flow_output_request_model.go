// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaConnectFlowOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrs(v string) *AddMediaConnectFlowOutputRequest
	GetCidrs() *string
	SetFlowId(v string) *AddMediaConnectFlowOutputRequest
	GetFlowId() *string
	SetOutputName(v string) *AddMediaConnectFlowOutputRequest
	GetOutputName() *string
	SetOutputProtocol(v string) *AddMediaConnectFlowOutputRequest
	GetOutputProtocol() *string
	SetOutputToUrl(v string) *AddMediaConnectFlowOutputRequest
	GetOutputToUrl() *string
	SetPairFlowId(v string) *AddMediaConnectFlowOutputRequest
	GetPairFlowId() *string
	SetPairInputName(v string) *AddMediaConnectFlowOutputRequest
	GetPairInputName() *string
	SetPlayerLimit(v int32) *AddMediaConnectFlowOutputRequest
	GetPlayerLimit() *int32
	SetSrtLatency(v int32) *AddMediaConnectFlowOutputRequest
	GetSrtLatency() *int32
	SetSrtPassphrase(v string) *AddMediaConnectFlowOutputRequest
	GetSrtPassphrase() *string
	SetSrtPbkeyLen(v string) *AddMediaConnectFlowOutputRequest
	GetSrtPbkeyLen() *string
}

type AddMediaConnectFlowOutputRequest struct {
	// The IP address whitelist in CIDR format. Separate multiple CIDR blocks with commas (,).
	//
	// example:
	//
	// 83.17.231.31/32
	Cidrs *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty"`
	// The flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The output name.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliTestOutput
	OutputName *string `json:"OutputName,omitempty" xml:"OutputName,omitempty"`
	// The output type.
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
	// RTMP-PULL
	OutputProtocol *string `json:"OutputProtocol,omitempty" xml:"OutputProtocol,omitempty"`
	// The output URL. This parameter is required when OutputProtocol is set to RTMP-PUSH or SRT-Caller.
	//
	// example:
	//
	// rtmp://push.test.alivecdn.com/live/alitest
	OutputToUrl *string `json:"OutputToUrl,omitempty" xml:"OutputToUrl,omitempty"`
	// The ID of the destination flow. This parameter is required when OutputProtocol is set to Flow.
	//
	// example:
	//
	// 8666ec062190f00e263012666319a5be
	PairFlowId *string `json:"PairFlowId,omitempty" xml:"PairFlowId,omitempty"`
	// The source name of the destination flow. This parameter is required when OutputProtocol is set to Flow.
	//
	// example:
	//
	// AliTestInput
	PairInputName *string `json:"PairInputName,omitempty" xml:"PairInputName,omitempty"`
	// The maximum number of viewers.
	//
	// example:
	//
	// 5
	PlayerLimit *int32 `json:"PlayerLimit,omitempty" xml:"PlayerLimit,omitempty"`
	// The latency for the SRT stream. This parameter is required when the source type is SRT-Listener or SRT-Caller.
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
	// example:
	//
	// 32
	SrtPbkeyLen *string `json:"SrtPbkeyLen,omitempty" xml:"SrtPbkeyLen,omitempty"`
}

func (s AddMediaConnectFlowOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMediaConnectFlowOutputRequest) GoString() string {
	return s.String()
}

func (s *AddMediaConnectFlowOutputRequest) GetCidrs() *string {
	return s.Cidrs
}

func (s *AddMediaConnectFlowOutputRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *AddMediaConnectFlowOutputRequest) GetOutputName() *string {
	return s.OutputName
}

func (s *AddMediaConnectFlowOutputRequest) GetOutputProtocol() *string {
	return s.OutputProtocol
}

func (s *AddMediaConnectFlowOutputRequest) GetOutputToUrl() *string {
	return s.OutputToUrl
}

func (s *AddMediaConnectFlowOutputRequest) GetPairFlowId() *string {
	return s.PairFlowId
}

func (s *AddMediaConnectFlowOutputRequest) GetPairInputName() *string {
	return s.PairInputName
}

func (s *AddMediaConnectFlowOutputRequest) GetPlayerLimit() *int32 {
	return s.PlayerLimit
}

func (s *AddMediaConnectFlowOutputRequest) GetSrtLatency() *int32 {
	return s.SrtLatency
}

func (s *AddMediaConnectFlowOutputRequest) GetSrtPassphrase() *string {
	return s.SrtPassphrase
}

func (s *AddMediaConnectFlowOutputRequest) GetSrtPbkeyLen() *string {
	return s.SrtPbkeyLen
}

func (s *AddMediaConnectFlowOutputRequest) SetCidrs(v string) *AddMediaConnectFlowOutputRequest {
	s.Cidrs = &v
	return s
}

func (s *AddMediaConnectFlowOutputRequest) SetFlowId(v string) *AddMediaConnectFlowOutputRequest {
	s.FlowId = &v
	return s
}

func (s *AddMediaConnectFlowOutputRequest) SetOutputName(v string) *AddMediaConnectFlowOutputRequest {
	s.OutputName = &v
	return s
}

func (s *AddMediaConnectFlowOutputRequest) SetOutputProtocol(v string) *AddMediaConnectFlowOutputRequest {
	s.OutputProtocol = &v
	return s
}

func (s *AddMediaConnectFlowOutputRequest) SetOutputToUrl(v string) *AddMediaConnectFlowOutputRequest {
	s.OutputToUrl = &v
	return s
}

func (s *AddMediaConnectFlowOutputRequest) SetPairFlowId(v string) *AddMediaConnectFlowOutputRequest {
	s.PairFlowId = &v
	return s
}

func (s *AddMediaConnectFlowOutputRequest) SetPairInputName(v string) *AddMediaConnectFlowOutputRequest {
	s.PairInputName = &v
	return s
}

func (s *AddMediaConnectFlowOutputRequest) SetPlayerLimit(v int32) *AddMediaConnectFlowOutputRequest {
	s.PlayerLimit = &v
	return s
}

func (s *AddMediaConnectFlowOutputRequest) SetSrtLatency(v int32) *AddMediaConnectFlowOutputRequest {
	s.SrtLatency = &v
	return s
}

func (s *AddMediaConnectFlowOutputRequest) SetSrtPassphrase(v string) *AddMediaConnectFlowOutputRequest {
	s.SrtPassphrase = &v
	return s
}

func (s *AddMediaConnectFlowOutputRequest) SetSrtPbkeyLen(v string) *AddMediaConnectFlowOutputRequest {
	s.SrtPbkeyLen = &v
	return s
}

func (s *AddMediaConnectFlowOutputRequest) Validate() error {
	return dara.Validate(s)
}
