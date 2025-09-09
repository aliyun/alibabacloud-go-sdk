// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*ListServiceInstanceResourcesRequestFilters) *ListServiceInstanceResourcesRequest
	GetFilters() []*ListServiceInstanceResourcesRequestFilters
	SetMaxResults(v int32) *ListServiceInstanceResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstanceResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServiceInstanceResourcesRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *ListServiceInstanceResourcesRequest
	GetServiceInstanceId() *string
	SetServiceInstanceResourceType(v string) *ListServiceInstanceResourcesRequest
	GetServiceInstanceResourceType() *string
	SetTag(v []*ListServiceInstanceResourcesRequestTag) *ListServiceInstanceResourcesRequest
	GetTag() []*ListServiceInstanceResourcesRequestTag
}

type ListServiceInstanceResourcesRequest struct {
	// The filter.
	Filters []*ListServiceInstanceResourcesRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID where the service instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d8a0cc2a1ee04dce****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// Service Instance resource type，include AliyunResource and ContainerResource.
	//
	// example:
	//
	// AliyunResource
	ServiceInstanceResourceType *string `json:"ServiceInstanceResourceType,omitempty" xml:"ServiceInstanceResourceType,omitempty"`
	// The tags.
	Tag []*ListServiceInstanceResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequest) GetFilters() []*ListServiceInstanceResourcesRequestFilters {
	return s.Filters
}

func (s *ListServiceInstanceResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstanceResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstanceResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceInstanceResourcesRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ListServiceInstanceResourcesRequest) GetServiceInstanceResourceType() *string {
	return s.ServiceInstanceResourceType
}

func (s *ListServiceInstanceResourcesRequest) GetTag() []*ListServiceInstanceResourcesRequestTag {
	return s.Tag
}

func (s *ListServiceInstanceResourcesRequest) SetFilters(v []*ListServiceInstanceResourcesRequestFilters) *ListServiceInstanceResourcesRequest {
	s.Filters = v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetMaxResults(v int32) *ListServiceInstanceResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetNextToken(v string) *ListServiceInstanceResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetRegionId(v string) *ListServiceInstanceResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetServiceInstanceId(v string) *ListServiceInstanceResourcesRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetServiceInstanceResourceType(v string) *ListServiceInstanceResourcesRequest {
	s.ServiceInstanceResourceType = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetTag(v []*ListServiceInstanceResourcesRequestTag) *ListServiceInstanceResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListServiceInstanceResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstanceResourcesRequestFilters struct {
	// Vaild values:
	//
	// - ExpireTimeStart
	//
	// - ExpireTimeEnd
	//
	// - PayType
	//
	// - ResourceARN
	//
	// example:
	//
	// ResourceARN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A value of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceResourcesRequestFilters) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequestFilters) GetName() *string {
	return s.Name
}

func (s *ListServiceInstanceResourcesRequestFilters) GetValues() []*string {
	return s.Values
}

func (s *ListServiceInstanceResourcesRequestFilters) SetName(v string) *ListServiceInstanceResourcesRequestFilters {
	s.Name = &v
	return s
}

func (s *ListServiceInstanceResourcesRequestFilters) SetValues(v []*string) *ListServiceInstanceResourcesRequestFilters {
	s.Values = v
	return s
}

func (s *ListServiceInstanceResourcesRequestFilters) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstanceResourcesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstanceResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListServiceInstanceResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListServiceInstanceResourcesRequestTag) SetKey(v string) *ListServiceInstanceResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstanceResourcesRequestTag) SetValue(v string) *ListServiceInstanceResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *ListServiceInstanceResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
