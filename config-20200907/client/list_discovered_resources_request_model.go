// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiscoveredResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUpdateTimestamp(v int64) *ListDiscoveredResourcesRequest
	GetEndUpdateTimestamp() *int64
	SetExcludeResourceTypes(v string) *ListDiscoveredResourcesRequest
	GetExcludeResourceTypes() *string
	SetMaxResults(v int32) *ListDiscoveredResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDiscoveredResourcesRequest
	GetNextToken() *string
	SetRegions(v string) *ListDiscoveredResourcesRequest
	GetRegions() *string
	SetResourceDeleted(v int32) *ListDiscoveredResourcesRequest
	GetResourceDeleted() *int32
	SetResourceId(v string) *ListDiscoveredResourcesRequest
	GetResourceId() *string
	SetResourceName(v string) *ListDiscoveredResourcesRequest
	GetResourceName() *string
	SetResourceTypes(v string) *ListDiscoveredResourcesRequest
	GetResourceTypes() *string
	SetStartUpdateTimestamp(v int64) *ListDiscoveredResourcesRequest
	GetStartUpdateTimestamp() *int64
}

type ListDiscoveredResourcesRequest struct {
	// The end of the time range to query resources, specified as a UNIX timestamp in milliseconds. Note:
	//
	// - The value cannot be earlier than StartUpdateTimestamp.
	//
	// - The time interval between StartUpdateTimestamp and EndUpdateTimestamp cannot exceed 30 days.
	//
	// - Specify both StartUpdateTimestamp and EndUpdateTimestamp, or leave both blank.
	//
	// example:
	//
	// 1724947200000
	EndUpdateTimestamp *int64 `json:"EndUpdateTimestamp,omitempty" xml:"EndUpdateTimestamp,omitempty"`
	// The resource types to exclude. Separate multiple resource types with commas (,). This parameter takes precedence over the ResourceTypes parameter.
	//
	// example:
	//
	// ACS::ECS::Instance,ACS::ECS::NetworkInterface
	ExcludeResourceTypes *string `json:"ExcludeResourceTypes,omitempty" xml:"ExcludeResourceTypes,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. If the response is truncated, use this token in a subsequent request to retrieve the next page of results.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the resource resides. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	Regions *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// The status of the resource. Valid values:
	//
	// - 0: The resource is deleted. If you delete a resource in the corresponding Alibaba Cloud service, Cloud Config displays the resource as **Deleted**.
	//
	// - 1 (Default): The resource is active. If a resource is managed, Cloud Config displays the resource as **Active**.
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
	// test-resource-name
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type. Separate multiple resource types with commas (,).
	//
	// example:
	//
	// ACS::ECS::NetworkInterface
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// The start of the time range to query resources, specified as a UNIX timestamp in milliseconds. Note:
	//
	// - The value cannot be later than EndUpdateTimestamp.
	//
	// - The time interval between StartUpdateTimestamp and EndUpdateTimestamp cannot exceed 30 days.
	//
	// - Specify both StartUpdateTimestamp and EndUpdateTimestamp, or leave both blank.
	//
	// example:
	//
	// 1722441600000
	StartUpdateTimestamp *int64 `json:"StartUpdateTimestamp,omitempty" xml:"StartUpdateTimestamp,omitempty"`
}

func (s ListDiscoveredResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDiscoveredResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesRequest) GetEndUpdateTimestamp() *int64 {
	return s.EndUpdateTimestamp
}

func (s *ListDiscoveredResourcesRequest) GetExcludeResourceTypes() *string {
	return s.ExcludeResourceTypes
}

func (s *ListDiscoveredResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDiscoveredResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDiscoveredResourcesRequest) GetRegions() *string {
	return s.Regions
}

func (s *ListDiscoveredResourcesRequest) GetResourceDeleted() *int32 {
	return s.ResourceDeleted
}

func (s *ListDiscoveredResourcesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListDiscoveredResourcesRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListDiscoveredResourcesRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *ListDiscoveredResourcesRequest) GetStartUpdateTimestamp() *int64 {
	return s.StartUpdateTimestamp
}

func (s *ListDiscoveredResourcesRequest) SetEndUpdateTimestamp(v int64) *ListDiscoveredResourcesRequest {
	s.EndUpdateTimestamp = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetExcludeResourceTypes(v string) *ListDiscoveredResourcesRequest {
	s.ExcludeResourceTypes = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetMaxResults(v int32) *ListDiscoveredResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetNextToken(v string) *ListDiscoveredResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetRegions(v string) *ListDiscoveredResourcesRequest {
	s.Regions = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetResourceDeleted(v int32) *ListDiscoveredResourcesRequest {
	s.ResourceDeleted = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetResourceId(v string) *ListDiscoveredResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetResourceName(v string) *ListDiscoveredResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetResourceTypes(v string) *ListDiscoveredResourcesRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetStartUpdateTimestamp(v int64) *ListDiscoveredResourcesRequest {
	s.StartUpdateTimestamp = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) Validate() error {
	return dara.Validate(s)
}
