// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLiveInputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *GetMediaLiveInputResponseBodyInput) *GetMediaLiveInputResponseBody
	GetInput() *GetMediaLiveInputResponseBodyInput
	SetRequestId(v string) *GetMediaLiveInputResponseBody
	GetRequestId() *string
}

type GetMediaLiveInputResponseBody struct {
	// The input information.
	Input *GetMediaLiveInputResponseBodyInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaLiveInputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveInputResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaLiveInputResponseBody) GetInput() *GetMediaLiveInputResponseBodyInput {
	return s.Input
}

func (s *GetMediaLiveInputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaLiveInputResponseBody) SetInput(v *GetMediaLiveInputResponseBodyInput) *GetMediaLiveInputResponseBody {
	s.Input = v
	return s
}

func (s *GetMediaLiveInputResponseBody) SetRequestId(v string) *GetMediaLiveInputResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaLiveInputResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveInputResponseBodyInput struct {
	// The IDs of the channels associated with the input.
	ChannelIds []*string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty" type:"Repeated"`
	// The time when the input was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-12-03T06:56:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the input.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	InputId *string `json:"InputId,omitempty" xml:"InputId,omitempty"`
	// The input configurations.
	InputInfos []*GetMediaLiveInputResponseBodyInputInputInfos `json:"InputInfos,omitempty" xml:"InputInfos,omitempty" type:"Repeated"`
	// The name of the input.
	//
	// example:
	//
	// myinput
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of the security groups associated with the input.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The input type.
	//
	// example:
	//
	// RTMP_PUSH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMediaLiveInputResponseBodyInput) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveInputResponseBodyInput) GoString() string {
	return s.String()
}

func (s *GetMediaLiveInputResponseBodyInput) GetChannelIds() []*string {
	return s.ChannelIds
}

func (s *GetMediaLiveInputResponseBodyInput) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaLiveInputResponseBodyInput) GetInputId() *string {
	return s.InputId
}

func (s *GetMediaLiveInputResponseBodyInput) GetInputInfos() []*GetMediaLiveInputResponseBodyInputInputInfos {
	return s.InputInfos
}

func (s *GetMediaLiveInputResponseBodyInput) GetName() *string {
	return s.Name
}

func (s *GetMediaLiveInputResponseBodyInput) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *GetMediaLiveInputResponseBodyInput) GetType() *string {
	return s.Type
}

func (s *GetMediaLiveInputResponseBodyInput) SetChannelIds(v []*string) *GetMediaLiveInputResponseBodyInput {
	s.ChannelIds = v
	return s
}

func (s *GetMediaLiveInputResponseBodyInput) SetCreateTime(v string) *GetMediaLiveInputResponseBodyInput {
	s.CreateTime = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInput) SetInputId(v string) *GetMediaLiveInputResponseBodyInput {
	s.InputId = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInput) SetInputInfos(v []*GetMediaLiveInputResponseBodyInputInputInfos) *GetMediaLiveInputResponseBodyInput {
	s.InputInfos = v
	return s
}

func (s *GetMediaLiveInputResponseBodyInput) SetName(v string) *GetMediaLiveInputResponseBodyInput {
	s.Name = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInput) SetSecurityGroupIds(v []*string) *GetMediaLiveInputResponseBodyInput {
	s.SecurityGroupIds = v
	return s
}

func (s *GetMediaLiveInputResponseBodyInput) SetType(v string) *GetMediaLiveInputResponseBodyInput {
	s.Type = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInput) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveInputResponseBodyInputInputInfos struct {
	// The endpoint that the stream is pushed to. This parameter is returned for PUSH inputs.
	//
	// example:
	//
	// rtmp://domain/app/stream
	DestHost *string `json:"DestHost,omitempty" xml:"DestHost,omitempty"`
	// The ID of the flow from MediaConnect.
	//
	// example:
	//
	// ******81-9693-40dc-bbab-db5e49******
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The output name of the MediaConnect flow.
	//
	// example:
	//
	// myFlowOutputName
	FlowOutputName *string `json:"FlowOutputName,omitempty" xml:"FlowOutputName,omitempty"`
	// The URL for input monitoring.
	//
	// example:
	//
	// rtmp://domain/app/stream_for_monitor
	MonitorUrl *string `json:"MonitorUrl,omitempty" xml:"MonitorUrl,omitempty"`
	// The source URL where the stream is pulled from. This parameter is returned for PULL inputs.
	//
	// example:
	//
	// rtmp://domain/app/stream
	SourceUrl     *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	SrtLatency    *int32  `json:"SrtLatency,omitempty" xml:"SrtLatency,omitempty"`
	SrtMaxBitrate *int32  `json:"SrtMaxBitrate,omitempty" xml:"SrtMaxBitrate,omitempty"`
	SrtPassphrase *string `json:"SrtPassphrase,omitempty" xml:"SrtPassphrase,omitempty"`
	SrtPbKeyLen   *int32  `json:"SrtPbKeyLen,omitempty" xml:"SrtPbKeyLen,omitempty"`
	// The name of the pushed stream. This parameter is returned for PUSH inputs.
	//
	// example:
	//
	// mystream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s GetMediaLiveInputResponseBodyInputInputInfos) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveInputResponseBodyInputInputInfos) GoString() string {
	return s.String()
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) GetDestHost() *string {
	return s.DestHost
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) GetFlowId() *string {
	return s.FlowId
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) GetFlowOutputName() *string {
	return s.FlowOutputName
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) GetMonitorUrl() *string {
	return s.MonitorUrl
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) GetSrtLatency() *int32 {
	return s.SrtLatency
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) GetSrtMaxBitrate() *int32 {
	return s.SrtMaxBitrate
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) GetSrtPassphrase() *string {
	return s.SrtPassphrase
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) GetSrtPbKeyLen() *int32 {
	return s.SrtPbKeyLen
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) GetStreamName() *string {
	return s.StreamName
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) SetDestHost(v string) *GetMediaLiveInputResponseBodyInputInputInfos {
	s.DestHost = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) SetFlowId(v string) *GetMediaLiveInputResponseBodyInputInputInfos {
	s.FlowId = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) SetFlowOutputName(v string) *GetMediaLiveInputResponseBodyInputInputInfos {
	s.FlowOutputName = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) SetMonitorUrl(v string) *GetMediaLiveInputResponseBodyInputInputInfos {
	s.MonitorUrl = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) SetSourceUrl(v string) *GetMediaLiveInputResponseBodyInputInputInfos {
	s.SourceUrl = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) SetSrtLatency(v int32) *GetMediaLiveInputResponseBodyInputInputInfos {
	s.SrtLatency = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) SetSrtMaxBitrate(v int32) *GetMediaLiveInputResponseBodyInputInputInfos {
	s.SrtMaxBitrate = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) SetSrtPassphrase(v string) *GetMediaLiveInputResponseBodyInputInputInfos {
	s.SrtPassphrase = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) SetSrtPbKeyLen(v int32) *GetMediaLiveInputResponseBodyInputInputInfos {
	s.SrtPbKeyLen = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) SetStreamName(v string) *GetMediaLiveInputResponseBodyInputInputInfos {
	s.StreamName = &v
	return s
}

func (s *GetMediaLiveInputResponseBodyInputInputInfos) Validate() error {
	return dara.Validate(s)
}
