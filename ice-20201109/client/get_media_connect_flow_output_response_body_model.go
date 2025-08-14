// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *GetMediaConnectFlowOutputResponseBodyContent) *GetMediaConnectFlowOutputResponseBody
	GetContent() *GetMediaConnectFlowOutputResponseBodyContent
	SetDescription(v string) *GetMediaConnectFlowOutputResponseBody
	GetDescription() *string
	SetRequestId(v string) *GetMediaConnectFlowOutputResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *GetMediaConnectFlowOutputResponseBody
	GetRetCode() *int32
}

type GetMediaConnectFlowOutputResponseBody struct {
	// The response body.
	Content *GetMediaConnectFlowOutputResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The call description.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0DB23DCE-0D69-598B-AA7C-7268D55E2F89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s GetMediaConnectFlowOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowOutputResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowOutputResponseBody) GetContent() *GetMediaConnectFlowOutputResponseBodyContent {
	return s.Content
}

func (s *GetMediaConnectFlowOutputResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetMediaConnectFlowOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaConnectFlowOutputResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *GetMediaConnectFlowOutputResponseBody) SetContent(v *GetMediaConnectFlowOutputResponseBodyContent) *GetMediaConnectFlowOutputResponseBody {
	s.Content = v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBody) SetDescription(v string) *GetMediaConnectFlowOutputResponseBody {
	s.Description = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBody) SetRequestId(v string) *GetMediaConnectFlowOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBody) SetRetCode(v int32) *GetMediaConnectFlowOutputResponseBody {
	s.RetCode = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMediaConnectFlowOutputResponseBodyContent struct {
	// The IP address whitelist in CIDR format. CIDR blocks are separated with commas (,).
	//
	// example:
	//
	// 10.211.0.0/17
	Cidrs *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty"`
	// The time when the flow was created.
	//
	// example:
	//
	// 2024-07-18T01:29:24Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Forbid     *string `json:"Forbid,omitempty" xml:"Forbid,omitempty"`
	// The output name.
	//
	// example:
	//
	// AliTestInput
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
	// example:
	//
	// SRT-PULL
	OutputProtocol *string `json:"OutputProtocol,omitempty" xml:"OutputProtocol,omitempty"`
	// The output URL.
	//
	// example:
	//
	// srt://1.2.3.4:1025
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// The ID of the destination flow. This parameter is returned when the output type is Flow.
	//
	// example:
	//
	// 805fbdd0-575e-4146-b35d-ec7f63937b20
	PairFlowId *string `json:"PairFlowId,omitempty" xml:"PairFlowId,omitempty"`
	// The source name of the destination flow. This parameter is returned when the output type is Flow.
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
	// The latency for the SRT stream. Unit: milliseconds. This parameter is returned when the source type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// 1000
	SrtLatency *int32 `json:"SrtLatency,omitempty" xml:"SrtLatency,omitempty"`
	// The SRT key. This parameter is returned when the source type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// FICUBPX4Q77DYHRF
	SrtPassphrase *string `json:"SrtPassphrase,omitempty" xml:"SrtPassphrase,omitempty"`
	// The encryption key length. This parameter is returned when the source type is SRT-Listener or SRT-Caller.
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
	SrtPbkeyLen *int32 `json:"SrtPbkeyLen,omitempty" xml:"SrtPbkeyLen,omitempty"`
}

func (s GetMediaConnectFlowOutputResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowOutputResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetCidrs() *string {
	return s.Cidrs
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetForbid() *string {
	return s.Forbid
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetOutputName() *string {
	return s.OutputName
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetOutputProtocol() *string {
	return s.OutputProtocol
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetPairFlowId() *string {
	return s.PairFlowId
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetPairInputName() *string {
	return s.PairInputName
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetPlayerLimit() *int32 {
	return s.PlayerLimit
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetSrtLatency() *int32 {
	return s.SrtLatency
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetSrtPassphrase() *string {
	return s.SrtPassphrase
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) GetSrtPbkeyLen() *int32 {
	return s.SrtPbkeyLen
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetCidrs(v string) *GetMediaConnectFlowOutputResponseBodyContent {
	s.Cidrs = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetCreateTime(v string) *GetMediaConnectFlowOutputResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetForbid(v string) *GetMediaConnectFlowOutputResponseBodyContent {
	s.Forbid = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetOutputName(v string) *GetMediaConnectFlowOutputResponseBodyContent {
	s.OutputName = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetOutputProtocol(v string) *GetMediaConnectFlowOutputResponseBodyContent {
	s.OutputProtocol = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetOutputUrl(v string) *GetMediaConnectFlowOutputResponseBodyContent {
	s.OutputUrl = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetPairFlowId(v string) *GetMediaConnectFlowOutputResponseBodyContent {
	s.PairFlowId = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetPairInputName(v string) *GetMediaConnectFlowOutputResponseBodyContent {
	s.PairInputName = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetPlayerLimit(v int32) *GetMediaConnectFlowOutputResponseBodyContent {
	s.PlayerLimit = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetSrtLatency(v int32) *GetMediaConnectFlowOutputResponseBodyContent {
	s.SrtLatency = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetSrtPassphrase(v string) *GetMediaConnectFlowOutputResponseBodyContent {
	s.SrtPassphrase = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) SetSrtPbkeyLen(v int32) *GetMediaConnectFlowOutputResponseBodyContent {
	s.SrtPbkeyLen = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
