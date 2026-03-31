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
	// The end time of the time range for querying resources. The value is a timestamp in the UTC format. When you specify this parameter, take note of the following limits:
	//
	// 	- The value must be a timestamp in milliseconds.
	//
	// 	- The value cannot be less than the value of the StartUpdateTimestamp parameter. The interval between the value and the value of the StartUpdateTimestamp parameter must be less than or equal to 30 days.
	//
	// 	- The StartUpdateTimestamp and EndUpdateTimestamp parameters must be specified at the same time or left empty at the same time.
	//
	// example:
	//
	// 1724947200000
	EndUpdateTimestamp *int64 `json:"EndUpdateTimestamp,omitempty" xml:"EndUpdateTimestamp,omitempty"`
	// The types of resources that are excluded. Separate multiple values with commas (,). If this parameter conflicts with the ResourceTypes parameter, this parameter prevails.
	//
	// example:
	//
	// ACS::ECS::Instance,ACS::ECS::NetworkInterface
	ExcludeResourceTypes *string `json:"ExcludeResourceTypes,omitempty" xml:"ExcludeResourceTypes,omitempty"`
	// The maximum number of entries returned for a single request. Valid values: 1 to 100.
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
	// The `token` that you want to use to initiate the current request. If the response of the previous request is truncated, you can use this token to initiate another request and obtain the remaining entries.
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
	// 	- 0: The resource is deleted. If a resource is deleted from the desired cloud service, **Deleted*	- is displayed in the resource list in the Cloud Config console.
	//
	// 	- 1 (default): The resource is retained. If a resource is managed as expected, **Active*	- is displayed in the resource list in the Cloud Config console.
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource. Separate multiple resource types with commas (,).
	//
	// example:
	//
	// ACS::ECS::NetworkInterface
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// The start time of the time range for querying resources. The value is a timestamp in the UTC format. When you specify this parameter, take note of the following limits:
	//
	// 	- The value must be a timestamp in milliseconds.
	//
	// 	- The value cannot be greater than the value of the EndUpdateTimestamp parameter. The interval between the value and the value of the EndUpdateTimestamp parameter must be less than or equal to 30 days.
	//
	// 	- The StartUpdateTimestamp and EndUpdateTimestamp parameters must be specified at the same time or left blank at the same time.
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
