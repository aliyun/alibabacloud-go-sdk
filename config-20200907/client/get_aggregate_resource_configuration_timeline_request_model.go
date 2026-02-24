// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceConfigurationTimelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateResourceConfigurationTimelineRequest
	GetAggregatorId() *string
	SetEndTime(v int64) *GetAggregateResourceConfigurationTimelineRequest
	GetEndTime() *int64
	SetMaxResults(v int32) *GetAggregateResourceConfigurationTimelineRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetAggregateResourceConfigurationTimelineRequest
	GetNextToken() *string
	SetRegion(v string) *GetAggregateResourceConfigurationTimelineRequest
	GetRegion() *string
	SetResourceAccountId(v int64) *GetAggregateResourceConfigurationTimelineRequest
	GetResourceAccountId() *int64
	SetResourceId(v string) *GetAggregateResourceConfigurationTimelineRequest
	GetResourceId() *string
	SetResourceOwnerId(v int64) *GetAggregateResourceConfigurationTimelineRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *GetAggregateResourceConfigurationTimelineRequest
	GetResourceType() *string
	SetStartTime(v int64) *GetAggregateResourceConfigurationTimelineRequest
	GetStartTime() *int64
}

type GetAggregateResourceConfigurationTimelineRequest struct {
	// The ID of the account group.
	//
	// For more information, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-5885626622af0008****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The end time of the query. This is a UNIX timestamp in milliseconds. By default, data up to the current time is queried.
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
	// If the output is truncated, you can use the `NextToken` to start the next query from the truncation point.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the resource resides.
	//
	// For more information, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the Alibaba Cloud account that owns the resource in the account group.
	//
	// example:
	//
	// 100931896542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// The ID of the resource.
	//
	// For more information, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// new-bucket
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Deprecated
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource.
	//
	// For more information, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::OSS::Bucket
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start time of the query. This is a UNIX timestamp in milliseconds. By default, data from the last 30 days is queried.
	//
	// example:
	//
	// 1623211156000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAggregateResourceConfigurationTimelineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceConfigurationTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceConfigurationTimelineRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateResourceConfigurationTimelineRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetAggregateResourceConfigurationTimelineRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetAggregateResourceConfigurationTimelineRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetAggregateResourceConfigurationTimelineRequest) GetRegion() *string {
	return s.Region
}

func (s *GetAggregateResourceConfigurationTimelineRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *GetAggregateResourceConfigurationTimelineRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetAggregateResourceConfigurationTimelineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAggregateResourceConfigurationTimelineRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAggregateResourceConfigurationTimelineRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetAggregatorId(v string) *GetAggregateResourceConfigurationTimelineRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetEndTime(v int64) *GetAggregateResourceConfigurationTimelineRequest {
	s.EndTime = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetMaxResults(v int32) *GetAggregateResourceConfigurationTimelineRequest {
	s.MaxResults = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetNextToken(v string) *GetAggregateResourceConfigurationTimelineRequest {
	s.NextToken = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetRegion(v string) *GetAggregateResourceConfigurationTimelineRequest {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetResourceAccountId(v int64) *GetAggregateResourceConfigurationTimelineRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetResourceId(v string) *GetAggregateResourceConfigurationTimelineRequest {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetResourceOwnerId(v int64) *GetAggregateResourceConfigurationTimelineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetResourceType(v string) *GetAggregateResourceConfigurationTimelineRequest {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetStartTime(v int64) *GetAggregateResourceConfigurationTimelineRequest {
	s.StartTime = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) Validate() error {
	return dara.Validate(s)
}
