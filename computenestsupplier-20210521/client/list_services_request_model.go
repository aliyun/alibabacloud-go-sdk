// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllVersions(v bool) *ListServicesRequest
	GetAllVersions() *bool
	SetFilter(v []*ListServicesRequestFilter) *ListServicesRequest
	GetFilter() []*ListServicesRequestFilter
	SetMaxResults(v int32) *ListServicesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServicesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServicesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListServicesRequest
	GetResourceGroupId() *string
	SetTag(v []*ListServicesRequestTag) *ListServicesRequest
	GetTag() []*ListServicesRequestTag
}

type ListServicesRequest struct {
	// Specifies whether to return all versions of a service. Default value: false, which specifies that only the default version of a service is returned.
	//
	// example:
	//
	// false
	AllVersions *bool `json:"AllVersions,omitempty" xml:"AllVersions,omitempty"`
	// The filters.
	Filter []*ListServicesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzkt5buxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The custom tags.
	Tag []*ListServicesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) GetAllVersions() *bool {
	return s.AllVersions
}

func (s *ListServicesRequest) GetFilter() []*ListServicesRequestFilter {
	return s.Filter
}

func (s *ListServicesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServicesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServicesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServicesRequest) GetTag() []*ListServicesRequestTag {
	return s.Tag
}

func (s *ListServicesRequest) SetAllVersions(v bool) *ListServicesRequest {
	s.AllVersions = &v
	return s
}

func (s *ListServicesRequest) SetFilter(v []*ListServicesRequestFilter) *ListServicesRequest {
	s.Filter = v
	return s
}

func (s *ListServicesRequest) SetMaxResults(v int32) *ListServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServicesRequest) SetNextToken(v string) *ListServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServicesRequest) SetRegionId(v string) *ListServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServicesRequest) SetResourceGroupId(v string) *ListServicesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServicesRequest) SetTag(v []*ListServicesRequestTag) *ListServicesRequest {
	s.Tag = v
	return s
}

func (s *ListServicesRequest) Validate() error {
	return dara.Validate(s)
}

type ListServicesRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// 	- ServiceId: the ID of the service.
	//
	// 	- Name: the name of the service.
	//
	// 	- Status: the state of the service.
	//
	// 	- SupplierName: the name of the service provider.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServicesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServicesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListServicesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListServicesRequestFilter) SetName(v string) *ListServicesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServicesRequestFilter) SetValue(v []*string) *ListServicesRequestFilter {
	s.Value = v
	return s
}

func (s *ListServicesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type ListServicesRequestTag struct {
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

func (s ListServicesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServicesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListServicesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListServicesRequestTag) SetKey(v string) *ListServicesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServicesRequestTag) SetValue(v string) *ListServicesRequestTag {
	s.Value = &v
	return s
}

func (s *ListServicesRequestTag) Validate() error {
	return dara.Validate(s)
}
