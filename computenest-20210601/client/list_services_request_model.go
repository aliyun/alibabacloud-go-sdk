// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListServicesRequestFilter) *ListServicesRequest
	GetFilter() []*ListServicesRequestFilter
	SetFuzzyKeyword(v string) *ListServicesRequest
	GetFuzzyKeyword() *string
	SetInUsed(v bool) *ListServicesRequest
	GetInUsed() *bool
	SetMaxResults(v int32) *ListServicesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServicesRequest
	GetNextToken() *string
	SetOrderByType(v string) *ListServicesRequest
	GetOrderByType() *string
	SetRegionId(v string) *ListServicesRequest
	GetRegionId() *string
	SetServiceAccessType(v string) *ListServicesRequest
	GetServiceAccessType() *string
	SetTag(v []*ListServicesRequestTag) *ListServicesRequest
	GetTag() []*ListServicesRequestTag
}

type ListServicesRequest struct {
	// The filter.
	Filter []*ListServicesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Keyword fuzzy query.
	//
	// example:
	//
	// name
	FuzzyKeyword *string `json:"FuzzyKeyword,omitempty" xml:"FuzzyKeyword,omitempty"`
	// Whether it is used. Optional values:
	//
	//
	//
	// - false: not being used.
	//
	//
	//
	// - true: already in use.
	//
	// example:
	//
	// false
	InUsed *bool `json:"InUsed,omitempty" xml:"InUsed,omitempty"`
	// The number of entries page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Service ordering type.
	//
	// example:
	//
	// UpdateTime
	OrderByType *string `json:"OrderByType,omitempty" xml:"OrderByType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service access type.
	//
	// example:
	//
	// All
	ServiceAccessType *string `json:"ServiceAccessType,omitempty" xml:"ServiceAccessType,omitempty"`
	// The tags.
	Tag []*ListServicesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) GetFilter() []*ListServicesRequestFilter {
	return s.Filter
}

func (s *ListServicesRequest) GetFuzzyKeyword() *string {
	return s.FuzzyKeyword
}

func (s *ListServicesRequest) GetInUsed() *bool {
	return s.InUsed
}

func (s *ListServicesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServicesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServicesRequest) GetOrderByType() *string {
	return s.OrderByType
}

func (s *ListServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServicesRequest) GetServiceAccessType() *string {
	return s.ServiceAccessType
}

func (s *ListServicesRequest) GetTag() []*ListServicesRequestTag {
	return s.Tag
}

func (s *ListServicesRequest) SetFilter(v []*ListServicesRequestFilter) *ListServicesRequest {
	s.Filter = v
	return s
}

func (s *ListServicesRequest) SetFuzzyKeyword(v string) *ListServicesRequest {
	s.FuzzyKeyword = &v
	return s
}

func (s *ListServicesRequest) SetInUsed(v bool) *ListServicesRequest {
	s.InUsed = &v
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

func (s *ListServicesRequest) SetOrderByType(v string) *ListServicesRequest {
	s.OrderByType = &v
	return s
}

func (s *ListServicesRequest) SetRegionId(v string) *ListServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServicesRequest) SetServiceAccessType(v string) *ListServicesRequest {
	s.ServiceAccessType = &v
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
	// A value of the filter condition.
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
