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
	// The access type. Valid values:
	//
	// - **NS**: NS access.
	//
	// - **CNAME**: CNAME access.
	//
	// example:
	//
	// NS
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The acceleration region. Valid values:
	//
	// - **domestic**: Chinese mainland only.
	//
	// - **global**: Global.
	//
	// - **overseas**: Global (excluding the Chinese mainland).
	//
	// example:
	//
	// global
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// Specifies whether to return only sites that use the Enterprise Edition.
	//
	// example:
	//
	// false
	OnlyEnterprise *bool `json:"OnlyEnterprise,omitempty" xml:"OnlyEnterprise,omitempty"`
	// The field to sort the results by. By default, results are sorted by creation time (gmtCreate). Supported values:
	//
	// - `gmtCreate`: site creation time
	//
	// - `visitTime`: site access time
	//
	// example:
	//
	// visitTime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of the page to return. The default value is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The default value is **500**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The subscription plan type. Valid values:
	//
	// - **basicplan**: Basic Edition.
	//
	// - **standardplan**: Standard Edition.
	//
	// - **advancedplan**: Advanced Edition.
	//
	// - **enterpriseplan**: Enterprise Edition.
	//
	// example:
	//
	// basicplan
	PlanSubscribeType *string `json:"PlanSubscribeType,omitempty" xml:"PlanSubscribeType,omitempty"`
	// The resource group ID, used to filter query results.
	//
	// example:
	//
	// rg-aekzd3styujvyei
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The site name, used to filter query results.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The match mode for the `SiteName` parameter. The default value is `exact`. Valid values:
	//
	// - **prefix**: prefix match.
	//
	// - **suffix**: suffix match.
	//
	// - **exact**: exact match.
	//
	// - **fuzzy**: fuzzy match.
	//
	// example:
	//
	// fuzzy
	SiteSearchType *string `json:"SiteSearchType,omitempty" xml:"SiteSearchType,omitempty"`
	// The site status, used to filter query results.
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A list of tags to use for filtering sites.
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
