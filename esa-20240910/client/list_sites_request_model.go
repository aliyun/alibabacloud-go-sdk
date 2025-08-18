// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSitesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *ListSitesRequest
	GetAccessType() *string
	SetCoverage(v string) *ListSitesRequest
	GetCoverage() *string
	SetOnlyEnterprise(v bool) *ListSitesRequest
	GetOnlyEnterprise() *bool
	SetOrderBy(v string) *ListSitesRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *ListSitesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSitesRequest
	GetPageSize() *int32
	SetPlanSubscribeType(v string) *ListSitesRequest
	GetPlanSubscribeType() *string
	SetResourceGroupId(v string) *ListSitesRequest
	GetResourceGroupId() *string
	SetSiteName(v string) *ListSitesRequest
	GetSiteName() *string
	SetSiteSearchType(v string) *ListSitesRequest
	GetSiteSearchType() *string
	SetStatus(v string) *ListSitesRequest
	GetStatus() *string
	SetTagFilter(v []*ListSitesRequestTagFilter) *ListSitesRequest
	GetTagFilter() []*ListSitesRequestTagFilter
}

type ListSitesRequest struct {
	// The DNS setup. Valid values:
	//
	// 	- **NS**
	//
	// 	- **CNAME**
	//
	// example:
	//
	// NS
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The service location. Valid values:
	//
	// 	- **domestic**: the Chinese mainland
	//
	// 	- **global**: global
	//
	// 	- **overseas**: outside the Chinese mainland
	//
	// example:
	//
	// global
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// Specifies whether to query only websites on Enterprise plans. Valid values: **true and false**.
	//
	// example:
	//
	// false
	OnlyEnterprise *bool `json:"OnlyEnterprise,omitempty" xml:"OnlyEnterprise,omitempty"`
	// Sorting field. By default, it sorts by creation time, supporting the following options:
	//
	// - gmtCreate: website creation time
	//
	// - visitTime: website visit time
	//
	// example:
	//
	// visitTime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **500**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The plan type. Valid values:
	//
	// 	- **basicplan**: Entrance
	//
	// 	- **standardplan**: Pro
	//
	// 	- **advancedplan**: Premium
	//
	// 	- **enterpriseplan**: Enterprise
	//
	// example:
	//
	// basicplan
	PlanSubscribeType *string `json:"PlanSubscribeType,omitempty" xml:"PlanSubscribeType,omitempty"`
	// The ID of the resource group. This parameter specifies a filter condition for the query.
	//
	// example:
	//
	// rg-aekzd3styujvyei
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The website name. This parameter specifies a filter condition for the query.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The match mode to search for the website name. Default value: exact. Valid values:
	//
	// 	- **prefix**: match by prefix.
	//
	// 	- **suffix**: match by suffix.
	//
	// 	- **exact**: exact match.
	//
	// 	- **fuzzy**: fuzzy match.
	//
	// example:
	//
	// fuzzy
	SiteSearchType *string `json:"SiteSearchType,omitempty" xml:"SiteSearchType,omitempty"`
	// The website status. This parameter specifies a filter condition for the query.
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag filtering rule.
	TagFilter []*ListSitesRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Repeated"`
}

func (s ListSitesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSitesRequest) GoString() string {
	return s.String()
}

func (s *ListSitesRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *ListSitesRequest) GetCoverage() *string {
	return s.Coverage
}

func (s *ListSitesRequest) GetOnlyEnterprise() *bool {
	return s.OnlyEnterprise
}

func (s *ListSitesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListSitesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSitesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSitesRequest) GetPlanSubscribeType() *string {
	return s.PlanSubscribeType
}

func (s *ListSitesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSitesRequest) GetSiteName() *string {
	return s.SiteName
}

func (s *ListSitesRequest) GetSiteSearchType() *string {
	return s.SiteSearchType
}

func (s *ListSitesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSitesRequest) GetTagFilter() []*ListSitesRequestTagFilter {
	return s.TagFilter
}

func (s *ListSitesRequest) SetAccessType(v string) *ListSitesRequest {
	s.AccessType = &v
	return s
}

func (s *ListSitesRequest) SetCoverage(v string) *ListSitesRequest {
	s.Coverage = &v
	return s
}

func (s *ListSitesRequest) SetOnlyEnterprise(v bool) *ListSitesRequest {
	s.OnlyEnterprise = &v
	return s
}

func (s *ListSitesRequest) SetOrderBy(v string) *ListSitesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListSitesRequest) SetPageNumber(v int32) *ListSitesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSitesRequest) SetPageSize(v int32) *ListSitesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSitesRequest) SetPlanSubscribeType(v string) *ListSitesRequest {
	s.PlanSubscribeType = &v
	return s
}

func (s *ListSitesRequest) SetResourceGroupId(v string) *ListSitesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSitesRequest) SetSiteName(v string) *ListSitesRequest {
	s.SiteName = &v
	return s
}

func (s *ListSitesRequest) SetSiteSearchType(v string) *ListSitesRequest {
	s.SiteSearchType = &v
	return s
}

func (s *ListSitesRequest) SetStatus(v string) *ListSitesRequest {
	s.Status = &v
	return s
}

func (s *ListSitesRequest) SetTagFilter(v []*ListSitesRequestTagFilter) *ListSitesRequest {
	s.TagFilter = v
	return s
}

func (s *ListSitesRequest) Validate() error {
	return dara.Validate(s)
}

type ListSitesRequestTagFilter struct {
	// The tag key. This parameter specifies a filter condition for the query.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. This parameter specifies a filter condition for the query.
	//
	// example:
	//
	// aaa
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSitesRequestTagFilter) String() string {
	return dara.Prettify(s)
}

func (s ListSitesRequestTagFilter) GoString() string {
	return s.String()
}

func (s *ListSitesRequestTagFilter) GetKey() *string {
	return s.Key
}

func (s *ListSitesRequestTagFilter) GetValue() *string {
	return s.Value
}

func (s *ListSitesRequestTagFilter) SetKey(v string) *ListSitesRequestTagFilter {
	s.Key = &v
	return s
}

func (s *ListSitesRequestTagFilter) SetValue(v string) *ListSitesRequestTagFilter {
	s.Value = &v
	return s
}

func (s *ListSitesRequestTagFilter) Validate() error {
	return dara.Validate(s)
}
