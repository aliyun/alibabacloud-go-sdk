// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaConnectFlowInputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrs(v string) *UpdateMediaConnectFlowInputRequest
	GetCidrs() *string
	SetFlowId(v string) *UpdateMediaConnectFlowInputRequest
	GetFlowId() *string
	SetInputFromUrl(v string) *UpdateMediaConnectFlowInputRequest
	GetInputFromUrl() *string
	SetInputName(v string) *UpdateMediaConnectFlowInputRequest
	GetInputName() *string
	SetMaxBitrate(v int32) *UpdateMediaConnectFlowInputRequest
	GetMaxBitrate() *int32
	SetSrtLatency(v int32) *UpdateMediaConnectFlowInputRequest
	GetSrtLatency() *int32
	SetSrtPassphrase(v string) *UpdateMediaConnectFlowInputRequest
	GetSrtPassphrase() *string
	SetSrtPbkeyLen(v int32) *UpdateMediaConnectFlowInputRequest
	GetSrtPbkeyLen() *int32
}

type UpdateMediaConnectFlowInputRequest struct {
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
	// The source URL. You can modify this parameter only when the source type is RTMP-PULL or SRT-Listener.
	//
	// example:
	//
	// rtmp://pull.test.alivecdn.com/live/alitest
	InputFromUrl *string `json:"InputFromUrl,omitempty" xml:"InputFromUrl,omitempty"`
	InputName    *string `json:"InputName,omitempty" xml:"InputName,omitempty"`
	// The maximum bitrate. Unit: bit/s.
	//
	// example:
	//
	// 2000000
	MaxBitrate *int32 `json:"MaxBitrate,omitempty" xml:"MaxBitrate,omitempty"`
	// The latency for the SRT stream. You can modify this parameter only when the source type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// 1000
	SrtLatency *int32 `json:"SrtLatency,omitempty" xml:"SrtLatency,omitempty"`
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
	SrtPbkeyLen *int32 `json:"SrtPbkeyLen,omitempty" xml:"SrtPbkeyLen,omitempty"`
}

func (s UpdateMediaConnectFlowInputRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaConnectFlowInputRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaConnectFlowInputRequest) GetCidrs() *string {
	return s.Cidrs
}

func (s *UpdateMediaConnectFlowInputRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *UpdateMediaConnectFlowInputRequest) GetInputFromUrl() *string {
	return s.InputFromUrl
}

func (s *UpdateMediaConnectFlowInputRequest) GetInputName() *string {
	return s.InputName
}

func (s *UpdateMediaConnectFlowInputRequest) GetMaxBitrate() *int32 {
	return s.MaxBitrate
}

func (s *UpdateMediaConnectFlowInputRequest) GetSrtLatency() *int32 {
	return s.SrtLatency
}

func (s *UpdateMediaConnectFlowInputRequest) GetSrtPassphrase() *string {
	return s.SrtPassphrase
}

func (s *UpdateMediaConnectFlowInputRequest) GetSrtPbkeyLen() *int32 {
	return s.SrtPbkeyLen
}

func (s *UpdateMediaConnectFlowInputRequest) SetCidrs(v string) *UpdateMediaConnectFlowInputRequest {
	s.Cidrs = &v
	return s
}

func (s *UpdateMediaConnectFlowInputRequest) SetFlowId(v string) *UpdateMediaConnectFlowInputRequest {
	s.FlowId = &v
	return s
}

func (s *UpdateMediaConnectFlowInputRequest) SetInputFromUrl(v string) *UpdateMediaConnectFlowInputRequest {
	s.InputFromUrl = &v
	return s
}

func (s *UpdateMediaConnectFlowInputRequest) SetInputName(v string) *UpdateMediaConnectFlowInputRequest {
	s.InputName = &v
	return s
}

func (s *UpdateMediaConnectFlowInputRequest) SetMaxBitrate(v int32) *UpdateMediaConnectFlowInputRequest {
	s.MaxBitrate = &v
	return s
}

func (s *UpdateMediaConnectFlowInputRequest) SetSrtLatency(v int32) *UpdateMediaConnectFlowInputRequest {
	s.SrtLatency = &v
	return s
}

func (s *UpdateMediaConnectFlowInputRequest) SetSrtPassphrase(v string) *UpdateMediaConnectFlowInputRequest {
	s.SrtPassphrase = &v
	return s
}

func (s *UpdateMediaConnectFlowInputRequest) SetSrtPbkeyLen(v int32) *UpdateMediaConnectFlowInputRequest {
	s.SrtPbkeyLen = &v
	return s
}

func (s *UpdateMediaConnectFlowInputRequest) Validate() error {
	return dara.Validate(s)
}
