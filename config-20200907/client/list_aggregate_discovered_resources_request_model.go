// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateDiscoveredResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateDiscoveredResourcesRequest
	GetAggregatorId() *string
	SetEndUpdateTimestamp(v int64) *ListAggregateDiscoveredResourcesRequest
	GetEndUpdateTimestamp() *int64
	SetExcludeResourceTypes(v string) *ListAggregateDiscoveredResourcesRequest
	GetExcludeResourceTypes() *string
	SetMaxResults(v int32) *ListAggregateDiscoveredResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAggregateDiscoveredResourcesRequest
	GetNextToken() *string
	SetRegions(v string) *ListAggregateDiscoveredResourcesRequest
	GetRegions() *string
	SetResourceAccountId(v int64) *ListAggregateDiscoveredResourcesRequest
	GetResourceAccountId() *int64
	SetResourceDeleted(v int32) *ListAggregateDiscoveredResourcesRequest
	GetResourceDeleted() *int32
	SetResourceId(v string) *ListAggregateDiscoveredResourcesRequest
	GetResourceId() *string
	SetResourceName(v string) *ListAggregateDiscoveredResourcesRequest
	GetResourceName() *string
	SetResourceOwnerId(v int64) *ListAggregateDiscoveredResourcesRequest
	GetResourceOwnerId() *int64
	SetResourceTypes(v string) *ListAggregateDiscoveredResourcesRequest
	GetResourceTypes() *string
	SetStartUpdateTimestamp(v int64) *ListAggregateDiscoveredResourcesRequest
	GetStartUpdateTimestamp() *int64
}

type ListAggregateDiscoveredResourcesRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-c560626622af0005****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The end of the time range to query. This is a standard UTC timestamp. The following limits apply:
	//
	// - The value must be a timestamp in milliseconds.
	//
	// - The value cannot be earlier than StartUpdateTimestamp. The interval between StartUpdateTimestamp and EndUpdateTimestamp cannot exceed 30 days.
	//
	// - You must specify both StartUpdateTimestamp and EndUpdateTimestamp, or leave both empty.
	//
	// example:
	//
	// 1724947200000
	EndUpdateTimestamp *int64 `json:"EndUpdateTimestamp,omitempty" xml:"EndUpdateTimestamp,omitempty"`
	// The resource types to exclude. Separate multiple resource types with commas (,). This parameter has a higher priority than the ResourceTypes parameter.
	//
	// example:
	//
	// ACS::ECS::Instance,ACS::ECS::NetworkInterface
	ExcludeResourceTypes *string `json:"ExcludeResourceTypes,omitempty" xml:"ExcludeResourceTypes,omitempty"`
	// The maximum number of entries to return for a single request. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If the response is truncated, use the `NextToken` to retrieve the next page of results.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the resource resides. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-huhehaote
	Regions *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// The ID of the Alibaba Cloud account to which the resources to be queried belong. The account is a member of the account group.
	//
	// example:
	//
	// 100931896542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// The status of the resource. Valid values:
	//
	// - 0: The resource is deleted. A resource is displayed as Deleted in Cloud Config after it is deleted from the source Alibaba Cloud service.
	//
	// - 1 (Default): The resource is active. A resource is displayed as Active in Cloud Config if it is properly managed.
	//
	// example:
	//
	// 1
	ResourceDeleted *int32 `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// eni-hp31cqoba96jagtz****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// launch-advisor-20200330
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// Deprecated
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Separate multiple resource types with commas (,).
	//
	// example:
	//
	// ACS::ECS::NetworkInterface
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// The start of the time range to query. This is a standard UTC timestamp. The following limits apply:
	//
	// - The value must be a timestamp in milliseconds.
	//
	// - The value cannot be later than EndUpdateTimestamp. The interval between StartUpdateTimestamp and EndUpdateTimestamp cannot exceed 30 days.
	//
	// - You must specify both StartUpdateTimestamp and EndUpdateTimestamp, or leave both empty.
	//
	// example:
	//
	// 1722441600000
	StartUpdateTimestamp *int64 `json:"StartUpdateTimestamp,omitempty" xml:"StartUpdateTimestamp,omitempty"`
}

func (s ListAggregateDiscoveredResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateDiscoveredResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateDiscoveredResourcesRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateDiscoveredResourcesRequest) GetEndUpdateTimestamp() *int64 {
	return s.EndUpdateTimestamp
}

func (s *ListAggregateDiscoveredResourcesRequest) GetExcludeResourceTypes() *string {
	return s.ExcludeResourceTypes
}

func (s *ListAggregateDiscoveredResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggregateDiscoveredResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregateDiscoveredResourcesRequest) GetRegions() *string {
	return s.Regions
}

func (s *ListAggregateDiscoveredResourcesRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *ListAggregateDiscoveredResourcesRequest) GetResourceDeleted() *int32 {
	return s.ResourceDeleted
}

func (s *ListAggregateDiscoveredResourcesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListAggregateDiscoveredResourcesRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListAggregateDiscoveredResourcesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAggregateDiscoveredResourcesRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *ListAggregateDiscoveredResourcesRequest) GetStartUpdateTimestamp() *int64 {
	return s.StartUpdateTimestamp
}

func (s *ListAggregateDiscoveredResourcesRequest) SetAggregatorId(v string) *ListAggregateDiscoveredResourcesRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetEndUpdateTimestamp(v int64) *ListAggregateDiscoveredResourcesRequest {
	s.EndUpdateTimestamp = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetExcludeResourceTypes(v string) *ListAggregateDiscoveredResourcesRequest {
	s.ExcludeResourceTypes = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetMaxResults(v int32) *ListAggregateDiscoveredResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetNextToken(v string) *ListAggregateDiscoveredResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetRegions(v string) *ListAggregateDiscoveredResourcesRequest {
	s.Regions = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetResourceAccountId(v int64) *ListAggregateDiscoveredResourcesRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetResourceDeleted(v int32) *ListAggregateDiscoveredResourcesRequest {
	s.ResourceDeleted = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetResourceId(v string) *ListAggregateDiscoveredResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetResourceName(v string) *ListAggregateDiscoveredResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetResourceOwnerId(v int64) *ListAggregateDiscoveredResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetResourceTypes(v string) *ListAggregateDiscoveredResourcesRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetStartUpdateTimestamp(v int64) *ListAggregateDiscoveredResourcesRequest {
	s.StartUpdateTimestamp = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) Validate() error {
	return dara.Validate(s)
}
