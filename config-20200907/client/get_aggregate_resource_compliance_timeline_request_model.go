// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceTimelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateResourceComplianceTimelineRequest
	GetAggregatorId() *string
	SetEndTime(v int64) *GetAggregateResourceComplianceTimelineRequest
	GetEndTime() *int64
	SetMaxResults(v int32) *GetAggregateResourceComplianceTimelineRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetAggregateResourceComplianceTimelineRequest
	GetNextToken() *string
	SetRegion(v string) *GetAggregateResourceComplianceTimelineRequest
	GetRegion() *string
	SetResourceAccountId(v int64) *GetAggregateResourceComplianceTimelineRequest
	GetResourceAccountId() *int64
	SetResourceId(v string) *GetAggregateResourceComplianceTimelineRequest
	GetResourceId() *string
	SetResourceOwnerId(v int64) *GetAggregateResourceComplianceTimelineRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *GetAggregateResourceComplianceTimelineRequest
	GetResourceType() *string
	SetStartTime(v int64) *GetAggregateResourceComplianceTimelineRequest
	GetStartTime() *int64
}

type GetAggregateResourceComplianceTimelineRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-5885626622af0008****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The end timestamp. By default, data up to the current time is queried. Unit: milliseconds.
	//
	// example:
	//
	// 1625821156000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If the output of a request is truncated, you can use this token to query the next page of results.
	//
	// example:
	//
	// 5OVS5J4I1/UKTkHV5oNs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the resource resides.
	//
	// For more information about how to obtain the ID of the region where a resource resides, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource in the account group belongs.
	//
	// > Set either the ResourceAccountId or ResourceOwnerId parameter. This parameter is recommended.
	//
	// example:
	//
	// 100931896542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// The resource ID.
	//
	// For more information about how to obtain the resource ID, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// new-bucket
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Deprecated
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type.
	//
	// For more information about how to obtain the resource type, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::OSS::Bucket
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start timestamp. By default, data from the last 30 days is queried. Unit: milliseconds.
	//
	// example:
	//
	// 1623211156000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAggregateResourceComplianceTimelineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceTimelineRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateResourceComplianceTimelineRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetAggregateResourceComplianceTimelineRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetAggregateResourceComplianceTimelineRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetAggregateResourceComplianceTimelineRequest) GetRegion() *string {
	return s.Region
}

func (s *GetAggregateResourceComplianceTimelineRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *GetAggregateResourceComplianceTimelineRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetAggregateResourceComplianceTimelineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAggregateResourceComplianceTimelineRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAggregateResourceComplianceTimelineRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetAggregatorId(v string) *GetAggregateResourceComplianceTimelineRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetEndTime(v int64) *GetAggregateResourceComplianceTimelineRequest {
	s.EndTime = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetMaxResults(v int32) *GetAggregateResourceComplianceTimelineRequest {
	s.MaxResults = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetNextToken(v string) *GetAggregateResourceComplianceTimelineRequest {
	s.NextToken = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetRegion(v string) *GetAggregateResourceComplianceTimelineRequest {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetResourceAccountId(v int64) *GetAggregateResourceComplianceTimelineRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetResourceId(v string) *GetAggregateResourceComplianceTimelineRequest {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetResourceOwnerId(v int64) *GetAggregateResourceComplianceTimelineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetResourceType(v string) *GetAggregateResourceComplianceTimelineRequest {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetStartTime(v int64) *GetAggregateResourceComplianceTimelineRequest {
	s.StartTime = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) Validate() error {
	return dara.Validate(s)
}
