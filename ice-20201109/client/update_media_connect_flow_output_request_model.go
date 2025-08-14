// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaConnectFlowOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrs(v string) *UpdateMediaConnectFlowOutputRequest
	GetCidrs() *string
	SetFlowId(v string) *UpdateMediaConnectFlowOutputRequest
	GetFlowId() *string
	SetOutputName(v string) *UpdateMediaConnectFlowOutputRequest
	GetOutputName() *string
	SetOutputToUrl(v string) *UpdateMediaConnectFlowOutputRequest
	GetOutputToUrl() *string
	SetPlayerLimit(v string) *UpdateMediaConnectFlowOutputRequest
	GetPlayerLimit() *string
	SetSrtLatency(v string) *UpdateMediaConnectFlowOutputRequest
	GetSrtLatency() *string
	SetSrtPassphrase(v string) *UpdateMediaConnectFlowOutputRequest
	GetSrtPassphrase() *string
	SetSrtPbkeyLen(v string) *UpdateMediaConnectFlowOutputRequest
	GetSrtPbkeyLen() *string
}

type UpdateMediaConnectFlowOutputRequest struct {
	// The IP address whitelist.
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
	// The output name.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliTestOutput
	OutputName *string `json:"OutputName,omitempty" xml:"OutputName,omitempty"`
	// The output URL. You can modify this parameter only when the output type is RTMP-PUSH or SRT-Caller.
	//
	// example:
	//
	// rtmp://push.test.alivecdn.com/live/alitest
	OutputToUrl *string `json:"OutputToUrl,omitempty" xml:"OutputToUrl,omitempty"`
	// The maximum number of viewers.
	//
	// example:
	//
	// 5
	PlayerLimit *string `json:"PlayerLimit,omitempty" xml:"PlayerLimit,omitempty"`
	// The latency for the SRT stream. You can modify this parameter only when the source type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// 1000
	SrtLatency *string `json:"SrtLatency,omitempty" xml:"SrtLatency,omitempty"`
	// The SRT key. You can modify this parameter only when the source type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// FICUBPX4Q77DYHRF
	SrtPassphrase *string `json:"SrtPassphrase,omitempty" xml:"SrtPassphrase,omitempty"`
	// The encryption key length. You can modify this parameter only when the source type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// 32
	SrtPbkeyLen *string `json:"SrtPbkeyLen,omitempty" xml:"SrtPbkeyLen,omitempty"`
}

func (s UpdateMediaConnectFlowOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaConnectFlowOutputRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaConnectFlowOutputRequest) GetCidrs() *string {
	return s.Cidrs
}

func (s *UpdateMediaConnectFlowOutputRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *UpdateMediaConnectFlowOutputRequest) GetOutputName() *string {
	return s.OutputName
}

func (s *UpdateMediaConnectFlowOutputRequest) GetOutputToUrl() *string {
	return s.OutputToUrl
}

func (s *UpdateMediaConnectFlowOutputRequest) GetPlayerLimit() *string {
	return s.PlayerLimit
}

func (s *UpdateMediaConnectFlowOutputRequest) GetSrtLatency() *string {
	return s.SrtLatency
}

func (s *UpdateMediaConnectFlowOutputRequest) GetSrtPassphrase() *string {
	return s.SrtPassphrase
}

func (s *UpdateMediaConnectFlowOutputRequest) GetSrtPbkeyLen() *string {
	return s.SrtPbkeyLen
}

func (s *UpdateMediaConnectFlowOutputRequest) SetCidrs(v string) *UpdateMediaConnectFlowOutputRequest {
	s.Cidrs = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputRequest) SetFlowId(v string) *UpdateMediaConnectFlowOutputRequest {
	s.FlowId = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputRequest) SetOutputName(v string) *UpdateMediaConnectFlowOutputRequest {
	s.OutputName = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputRequest) SetOutputToUrl(v string) *UpdateMediaConnectFlowOutputRequest {
	s.OutputToUrl = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputRequest) SetPlayerLimit(v string) *UpdateMediaConnectFlowOutputRequest {
	s.PlayerLimit = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputRequest) SetSrtLatency(v string) *UpdateMediaConnectFlowOutputRequest {
	s.SrtLatency = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputRequest) SetSrtPassphrase(v string) *UpdateMediaConnectFlowOutputRequest {
	s.SrtPassphrase = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputRequest) SetSrtPbkeyLen(v string) *UpdateMediaConnectFlowOutputRequest {
	s.SrtPbkeyLen = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputRequest) Validate() error {
	return dara.Validate(s)
}
