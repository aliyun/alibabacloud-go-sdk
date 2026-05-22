// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSitesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *ListSitesShrinkRequest
	GetAccessType() *string
	SetCoverage(v string) *ListSitesShrinkRequest
	GetCoverage() *string
	SetOnlyEnterprise(v bool) *ListSitesShrinkRequest
	GetOnlyEnterprise() *bool
	SetOrderBy(v string) *ListSitesShrinkRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *ListSitesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSitesShrinkRequest
	GetPageSize() *int32
	SetPlanSubscribeType(v string) *ListSitesShrinkRequest
	GetPlanSubscribeType() *string
	SetResourceGroupId(v string) *ListSitesShrinkRequest
	GetResourceGroupId() *string
	SetSiteName(v string) *ListSitesShrinkRequest
	GetSiteName() *string
	SetSiteSearchType(v string) *ListSitesShrinkRequest
	GetSiteSearchType() *string
	SetStatus(v string) *ListSitesShrinkRequest
	GetStatus() *string
	SetTagFilterShrink(v string) *ListSitesShrinkRequest
	GetTagFilterShrink() *string
}

type ListSitesShrinkRequest struct {
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
	TagFilterShrink *string `json:"TagFilter,omitempty" xml:"TagFilter,omitempty"`
}

func (s ListSitesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSitesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSitesShrinkRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *ListSitesShrinkRequest) GetCoverage() *string {
	return s.Coverage
}

func (s *ListSitesShrinkRequest) GetOnlyEnterprise() *bool {
	return s.OnlyEnterprise
}

func (s *ListSitesShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListSitesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSitesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSitesShrinkRequest) GetPlanSubscribeType() *string {
	return s.PlanSubscribeType
}

func (s *ListSitesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSitesShrinkRequest) GetSiteName() *string {
	return s.SiteName
}

func (s *ListSitesShrinkRequest) GetSiteSearchType() *string {
	return s.SiteSearchType
}

func (s *ListSitesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSitesShrinkRequest) GetTagFilterShrink() *string {
	return s.TagFilterShrink
}

func (s *ListSitesShrinkRequest) SetAccessType(v string) *ListSitesShrinkRequest {
	s.AccessType = &v
	return s
}

func (s *ListSitesShrinkRequest) SetCoverage(v string) *ListSitesShrinkRequest {
	s.Coverage = &v
	return s
}

func (s *ListSitesShrinkRequest) SetOnlyEnterprise(v bool) *ListSitesShrinkRequest {
	s.OnlyEnterprise = &v
	return s
}

func (s *ListSitesShrinkRequest) SetOrderBy(v string) *ListSitesShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListSitesShrinkRequest) SetPageNumber(v int32) *ListSitesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSitesShrinkRequest) SetPageSize(v int32) *ListSitesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSitesShrinkRequest) SetPlanSubscribeType(v string) *ListSitesShrinkRequest {
	s.PlanSubscribeType = &v
	return s
}

func (s *ListSitesShrinkRequest) SetResourceGroupId(v string) *ListSitesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSitesShrinkRequest) SetSiteName(v string) *ListSitesShrinkRequest {
	s.SiteName = &v
	return s
}

func (s *ListSitesShrinkRequest) SetSiteSearchType(v string) *ListSitesShrinkRequest {
	s.SiteSearchType = &v
	return s
}

func (s *ListSitesShrinkRequest) SetStatus(v string) *ListSitesShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListSitesShrinkRequest) SetTagFilterShrink(v string) *ListSitesShrinkRequest {
	s.TagFilterShrink = &v
	return s
}

func (s *ListSitesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
