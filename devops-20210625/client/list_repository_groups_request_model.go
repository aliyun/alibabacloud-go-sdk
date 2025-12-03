// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListRepositoryGroupsRequest
	GetAccessToken() *string
	SetIncludePersonal(v bool) *ListRepositoryGroupsRequest
	GetIncludePersonal() *bool
	SetOrderBy(v string) *ListRepositoryGroupsRequest
	GetOrderBy() *string
	SetOrganizationId(v string) *ListRepositoryGroupsRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListRepositoryGroupsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListRepositoryGroupsRequest
	GetPageSize() *int64
	SetParentId(v int64) *ListRepositoryGroupsRequest
	GetParentId() *int64
	SetSearch(v string) *ListRepositoryGroupsRequest
	GetSearch() *string
	SetSort(v string) *ListRepositoryGroupsRequest
	GetSort() *string
}

type ListRepositoryGroupsRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// false
	IncludePersonal *bool `json:"includePersonal,omitempty" xml:"includePersonal,omitempty"`
	// example:
	//
	// updated_at
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 26842
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// Demo
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s ListRepositoryGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryGroupsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListRepositoryGroupsRequest) GetIncludePersonal() *bool {
	return s.IncludePersonal
}

func (s *ListRepositoryGroupsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListRepositoryGroupsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRepositoryGroupsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListRepositoryGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRepositoryGroupsRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *ListRepositoryGroupsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListRepositoryGroupsRequest) GetSort() *string {
	return s.Sort
}

func (s *ListRepositoryGroupsRequest) SetAccessToken(v string) *ListRepositoryGroupsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryGroupsRequest) SetIncludePersonal(v bool) *ListRepositoryGroupsRequest {
	s.IncludePersonal = &v
	return s
}

func (s *ListRepositoryGroupsRequest) SetOrderBy(v string) *ListRepositoryGroupsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListRepositoryGroupsRequest) SetOrganizationId(v string) *ListRepositoryGroupsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryGroupsRequest) SetPage(v int64) *ListRepositoryGroupsRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryGroupsRequest) SetPageSize(v int64) *ListRepositoryGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryGroupsRequest) SetParentId(v int64) *ListRepositoryGroupsRequest {
	s.ParentId = &v
	return s
}

func (s *ListRepositoryGroupsRequest) SetSearch(v string) *ListRepositoryGroupsRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoryGroupsRequest) SetSort(v string) *ListRepositoryGroupsRequest {
	s.Sort = &v
	return s
}

func (s *ListRepositoryGroupsRequest) Validate() error {
	return dara.Validate(s)
}
