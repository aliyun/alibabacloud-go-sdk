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
	AccessType        *string                      `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Coverage          *string                      `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	OnlyEnterprise    *bool                        `json:"OnlyEnterprise,omitempty" xml:"OnlyEnterprise,omitempty"`
	OrderBy           *string                      `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNumber        *int32                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlanSubscribeType *string                      `json:"PlanSubscribeType,omitempty" xml:"PlanSubscribeType,omitempty"`
	ResourceGroupId   *string                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SiteName          *string                      `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	SiteSearchType    *string                      `json:"SiteSearchType,omitempty" xml:"SiteSearchType,omitempty"`
	Status            *string                      `json:"Status,omitempty" xml:"Status,omitempty"`
	TagFilter         []*ListSitesRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Repeated"`
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
	if s.TagFilter != nil {
		for _, item := range s.TagFilter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSitesRequestTagFilter struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
