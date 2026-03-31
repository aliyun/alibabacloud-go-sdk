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
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-5885626622af0008****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The end of the time range to query. The default value indicates the time when the GetAggregateResourceConfigurationTimeline operation is called. Unit: milliseconds.
	//
	// example:
	//
	// 1625821156000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries to return for a single request. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request if the response of the current request is truncated. You can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which the resource resides.
	//
	// For more information about how to obtain the ID of a region, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Required. The ID of the Alibaba Cloud account to which the specified resource belongs in the account group.
	//
	// example:
	//
	// 100931896542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// The resource ID.
	//
	// For more information about how to query the ID of a resource, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
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
	// For more information about how to obtain the type of a resource, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::OSS::Bucket
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The beginning of the time range to query. By default, Cloud Config retrieves the configuration changes in the last 30 days for the specified resource. Unit: milliseconds.
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
