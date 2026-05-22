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
	AccessType        *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Coverage          *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	OnlyEnterprise    *bool   `json:"OnlyEnterprise,omitempty" xml:"OnlyEnterprise,omitempty"`
	OrderBy           *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlanSubscribeType *string `json:"PlanSubscribeType,omitempty" xml:"PlanSubscribeType,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SiteName          *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	SiteSearchType    *string `json:"SiteSearchType,omitempty" xml:"SiteSearchType,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TagFilterShrink   *string `json:"TagFilter,omitempty" xml:"TagFilter,omitempty"`
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
