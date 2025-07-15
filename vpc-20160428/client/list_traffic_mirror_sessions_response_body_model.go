// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficMirrorSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTrafficMirrorSessionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTrafficMirrorSessionsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListTrafficMirrorSessionsResponseBody
	GetTotalCount() *string
	SetTrafficMirrorSessions(v []*ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) *ListTrafficMirrorSessionsResponseBody
	GetTrafficMirrorSessions() []*ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions
}

type ListTrafficMirrorSessionsResponseBody struct {
	// The token that is used for the next query. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next queries are sent.
	//
	// 	- If a value of **NextToken*	- is returned, the value is the token that is used for the subsequent query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 739CA01C-92EB-4C69-BCC0-280149C6F41E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details about the traffic mirror session.
	TrafficMirrorSessions []*ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions `json:"TrafficMirrorSessions,omitempty" xml:"TrafficMirrorSessions,omitempty" type:"Repeated"`
}

func (s ListTrafficMirrorSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorSessionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTrafficMirrorSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrafficMirrorSessionsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListTrafficMirrorSessionsResponseBody) GetTrafficMirrorSessions() []*ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	return s.TrafficMirrorSessions
}

func (s *ListTrafficMirrorSessionsResponseBody) SetNextToken(v string) *ListTrafficMirrorSessionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBody) SetRequestId(v string) *ListTrafficMirrorSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBody) SetTotalCount(v string) *ListTrafficMirrorSessionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBody) SetTrafficMirrorSessions(v []*ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) *ListTrafficMirrorSessionsResponseBody {
	s.TrafficMirrorSessions = v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions struct {
	// The time when the session is created.
	//
	// example:
	//
	// 2023-09-05T15:26Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the traffic mirror session was enabled.
	//
	// 	- **false*	-
	//
	// 	- **true**
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum transmission unit.
	//
	// example:
	//
	// 1500
	PacketLength *int32 `json:"PacketLength,omitempty" xml:"PacketLength,omitempty"`
	// The priority of the traffic mirror session.
	//
	// A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the resource group to which the traffic mirror session belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag list.
	Tags []*ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the filter.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
	// The status of the traffic mirror session.
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// example:
	//
	// Normal
	TrafficMirrorSessionBusinessStatus *string `json:"TrafficMirrorSessionBusinessStatus,omitempty" xml:"TrafficMirrorSessionBusinessStatus,omitempty"`
	// The description of the traffic mirror session.
	//
	// example:
	//
	// This is a session.
	TrafficMirrorSessionDescription *string `json:"TrafficMirrorSessionDescription,omitempty" xml:"TrafficMirrorSessionDescription,omitempty"`
	// The ID of the traffic mirror session.
	//
	// example:
	//
	// tms-j6cla50buc44ap8tu****
	TrafficMirrorSessionId *string `json:"TrafficMirrorSessionId,omitempty" xml:"TrafficMirrorSessionId,omitempty"`
	// The name of the traffic mirror session.
	//
	// example:
	//
	// abc
	TrafficMirrorSessionName *string `json:"TrafficMirrorSessionName,omitempty" xml:"TrafficMirrorSessionName,omitempty"`
	// The status of the traffic mirror session. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Created
	TrafficMirrorSessionStatus *string `json:"TrafficMirrorSessionStatus,omitempty" xml:"TrafficMirrorSessionStatus,omitempty"`
	// The ID of the traffic mirror source.
	TrafficMirrorSourceIds []*string `json:"TrafficMirrorSourceIds,omitempty" xml:"TrafficMirrorSourceIds,omitempty" type:"Repeated"`
	// The ID of the traffic mirror destination.
	//
	// example:
	//
	// eni-j6c2fp57q8rr47rp****
	TrafficMirrorTargetId *string `json:"TrafficMirrorTargetId,omitempty" xml:"TrafficMirrorTargetId,omitempty"`
	// The type of the traffic mirror destination. Valid values:
	//
	// 	- **NetworkInterface**: an elastic network interface (ENI)
	//
	// 	- **SLB**: an internal-facing Server Load Balancer (SLB) instance
	//
	// example:
	//
	// NetworkInterface
	TrafficMirrorTargetType *string `json:"TrafficMirrorTargetType,omitempty" xml:"TrafficMirrorTargetType,omitempty"`
	// You can specify VNIs to distinguish different mirrored traffic.
	//
	// example:
	//
	// 10
	VirtualNetworkId *int32 `json:"VirtualNetworkId,omitempty" xml:"VirtualNetworkId,omitempty"`
}

func (s ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetPacketLength() *int32 {
	return s.PacketLength
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetPriority() *int32 {
	return s.Priority
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetTags() []*ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags {
	return s.Tags
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetTrafficMirrorSessionBusinessStatus() *string {
	return s.TrafficMirrorSessionBusinessStatus
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetTrafficMirrorSessionDescription() *string {
	return s.TrafficMirrorSessionDescription
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetTrafficMirrorSessionId() *string {
	return s.TrafficMirrorSessionId
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetTrafficMirrorSessionName() *string {
	return s.TrafficMirrorSessionName
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetTrafficMirrorSessionStatus() *string {
	return s.TrafficMirrorSessionStatus
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetTrafficMirrorSourceIds() []*string {
	return s.TrafficMirrorSourceIds
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetTrafficMirrorTargetId() *string {
	return s.TrafficMirrorTargetId
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetTrafficMirrorTargetType() *string {
	return s.TrafficMirrorTargetType
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) GetVirtualNetworkId() *int32 {
	return s.VirtualNetworkId
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetCreationTime(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.CreationTime = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetEnabled(v bool) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.Enabled = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetPacketLength(v int32) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.PacketLength = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetPriority(v int32) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.Priority = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetResourceGroupId(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetTags(v []*ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.Tags = v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetTrafficMirrorFilterId(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetTrafficMirrorSessionBusinessStatus(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.TrafficMirrorSessionBusinessStatus = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetTrafficMirrorSessionDescription(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.TrafficMirrorSessionDescription = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetTrafficMirrorSessionId(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.TrafficMirrorSessionId = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetTrafficMirrorSessionName(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.TrafficMirrorSessionName = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetTrafficMirrorSessionStatus(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.TrafficMirrorSessionStatus = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetTrafficMirrorSourceIds(v []*string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.TrafficMirrorSourceIds = v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetTrafficMirrorTargetId(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.TrafficMirrorTargetId = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetTrafficMirrorTargetType(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.TrafficMirrorTargetType = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) SetVirtualNetworkId(v int32) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions {
	s.VirtualNetworkId = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessions) Validate() error {
	return dara.Validate(s)
}

type ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags) GetKey() *string {
	return s.Key
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags) GetValue() *string {
	return s.Value
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags) SetKey(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags {
	s.Key = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags) SetValue(v string) *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags {
	s.Value = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponseBodyTrafficMirrorSessionsTags) Validate() error {
	return dara.Validate(s)
}
