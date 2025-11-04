// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaLiveInputsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInputs(v []*ListMediaLiveInputsResponseBodyInputs) *ListMediaLiveInputsResponseBody
	GetInputs() []*ListMediaLiveInputsResponseBodyInputs
	SetMaxResults(v int32) *ListMediaLiveInputsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListMediaLiveInputsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMediaLiveInputsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListMediaLiveInputsResponseBody
	GetTotalCount() *int32
}

type ListMediaLiveInputsResponseBody struct {
	// The inputs.
	Inputs []*ListMediaLiveInputsResponseBodyInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMediaLiveInputsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveInputsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaLiveInputsResponseBody) GetInputs() []*ListMediaLiveInputsResponseBodyInputs {
	return s.Inputs
}

func (s *ListMediaLiveInputsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMediaLiveInputsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMediaLiveInputsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaLiveInputsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMediaLiveInputsResponseBody) SetInputs(v []*ListMediaLiveInputsResponseBodyInputs) *ListMediaLiveInputsResponseBody {
	s.Inputs = v
	return s
}

func (s *ListMediaLiveInputsResponseBody) SetMaxResults(v int32) *ListMediaLiveInputsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMediaLiveInputsResponseBody) SetNextToken(v string) *ListMediaLiveInputsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMediaLiveInputsResponseBody) SetRequestId(v string) *ListMediaLiveInputsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaLiveInputsResponseBody) SetTotalCount(v int32) *ListMediaLiveInputsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMediaLiveInputsResponseBody) Validate() error {
	if s.Inputs != nil {
		for _, item := range s.Inputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMediaLiveInputsResponseBodyInputs struct {
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
	InputInfos []*ListMediaLiveInputsResponseBodyInputsInputInfos `json:"InputInfos,omitempty" xml:"InputInfos,omitempty" type:"Repeated"`
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

func (s ListMediaLiveInputsResponseBodyInputs) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveInputsResponseBodyInputs) GoString() string {
	return s.String()
}

func (s *ListMediaLiveInputsResponseBodyInputs) GetChannelIds() []*string {
	return s.ChannelIds
}

func (s *ListMediaLiveInputsResponseBodyInputs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMediaLiveInputsResponseBodyInputs) GetInputId() *string {
	return s.InputId
}

func (s *ListMediaLiveInputsResponseBodyInputs) GetInputInfos() []*ListMediaLiveInputsResponseBodyInputsInputInfos {
	return s.InputInfos
}

func (s *ListMediaLiveInputsResponseBodyInputs) GetName() *string {
	return s.Name
}

func (s *ListMediaLiveInputsResponseBodyInputs) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ListMediaLiveInputsResponseBodyInputs) GetType() *string {
	return s.Type
}

func (s *ListMediaLiveInputsResponseBodyInputs) SetChannelIds(v []*string) *ListMediaLiveInputsResponseBodyInputs {
	s.ChannelIds = v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputs) SetCreateTime(v string) *ListMediaLiveInputsResponseBodyInputs {
	s.CreateTime = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputs) SetInputId(v string) *ListMediaLiveInputsResponseBodyInputs {
	s.InputId = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputs) SetInputInfos(v []*ListMediaLiveInputsResponseBodyInputsInputInfos) *ListMediaLiveInputsResponseBodyInputs {
	s.InputInfos = v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputs) SetName(v string) *ListMediaLiveInputsResponseBodyInputs {
	s.Name = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputs) SetSecurityGroupIds(v []*string) *ListMediaLiveInputsResponseBodyInputs {
	s.SecurityGroupIds = v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputs) SetType(v string) *ListMediaLiveInputsResponseBodyInputs {
	s.Type = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputs) Validate() error {
	if s.InputInfos != nil {
		for _, item := range s.InputInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMediaLiveInputsResponseBodyInputsInputInfos struct {
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

func (s ListMediaLiveInputsResponseBodyInputsInputInfos) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveInputsResponseBodyInputsInputInfos) GoString() string {
	return s.String()
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) GetDestHost() *string {
	return s.DestHost
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) GetFlowId() *string {
	return s.FlowId
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) GetFlowOutputName() *string {
	return s.FlowOutputName
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) GetMonitorUrl() *string {
	return s.MonitorUrl
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) GetSrtLatency() *int32 {
	return s.SrtLatency
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) GetSrtMaxBitrate() *int32 {
	return s.SrtMaxBitrate
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) GetSrtPassphrase() *string {
	return s.SrtPassphrase
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) GetSrtPbKeyLen() *int32 {
	return s.SrtPbKeyLen
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) GetStreamName() *string {
	return s.StreamName
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) SetDestHost(v string) *ListMediaLiveInputsResponseBodyInputsInputInfos {
	s.DestHost = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) SetFlowId(v string) *ListMediaLiveInputsResponseBodyInputsInputInfos {
	s.FlowId = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) SetFlowOutputName(v string) *ListMediaLiveInputsResponseBodyInputsInputInfos {
	s.FlowOutputName = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) SetMonitorUrl(v string) *ListMediaLiveInputsResponseBodyInputsInputInfos {
	s.MonitorUrl = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) SetSourceUrl(v string) *ListMediaLiveInputsResponseBodyInputsInputInfos {
	s.SourceUrl = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) SetSrtLatency(v int32) *ListMediaLiveInputsResponseBodyInputsInputInfos {
	s.SrtLatency = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) SetSrtMaxBitrate(v int32) *ListMediaLiveInputsResponseBodyInputsInputInfos {
	s.SrtMaxBitrate = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) SetSrtPassphrase(v string) *ListMediaLiveInputsResponseBodyInputsInputInfos {
	s.SrtPassphrase = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) SetSrtPbKeyLen(v int32) *ListMediaLiveInputsResponseBodyInputsInputInfos {
	s.SrtPbKeyLen = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) SetStreamName(v string) *ListMediaLiveInputsResponseBodyInputsInputInfos {
	s.StreamName = &v
	return s
}

func (s *ListMediaLiveInputsResponseBodyInputsInputInfos) Validate() error {
	return dara.Validate(s)
}
