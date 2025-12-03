// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListRepositoriesRequest
	GetAccessToken() *string
	SetArchived(v bool) *ListRepositoriesRequest
	GetArchived() *bool
	SetMinAccessLevel(v int32) *ListRepositoriesRequest
	GetMinAccessLevel() *int32
	SetOrderBy(v string) *ListRepositoriesRequest
	GetOrderBy() *string
	SetOrganizationId(v string) *ListRepositoriesRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListRepositoriesRequest
	GetPage() *int64
	SetPerPage(v int64) *ListRepositoriesRequest
	GetPerPage() *int64
	SetSearch(v string) *ListRepositoriesRequest
	GetSearch() *string
	SetSort(v string) *ListRepositoriesRequest
	GetSort() *string
}

type ListRepositoriesRequest struct {
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// false
	Archived       *bool  `json:"archived,omitempty" xml:"archived,omitempty"`
	MinAccessLevel *int32 `json:"minAccessLevel,omitempty" xml:"minAccessLevel,omitempty"`
	// example:
	//
	// created_at
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// This parameter is required.
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 2
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PerPage *int64 `json:"perPage,omitempty" xml:"perPage,omitempty"`
	// example:
	//
	// Demo
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s ListRepositoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoriesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListRepositoriesRequest) GetArchived() *bool {
	return s.Archived
}

func (s *ListRepositoriesRequest) GetMinAccessLevel() *int32 {
	return s.MinAccessLevel
}

func (s *ListRepositoriesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListRepositoriesRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRepositoriesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListRepositoriesRequest) GetPerPage() *int64 {
	return s.PerPage
}

func (s *ListRepositoriesRequest) GetSearch() *string {
	return s.Search
}

func (s *ListRepositoriesRequest) GetSort() *string {
	return s.Sort
}

func (s *ListRepositoriesRequest) SetAccessToken(v string) *ListRepositoriesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoriesRequest) SetArchived(v bool) *ListRepositoriesRequest {
	s.Archived = &v
	return s
}

func (s *ListRepositoriesRequest) SetMinAccessLevel(v int32) *ListRepositoriesRequest {
	s.MinAccessLevel = &v
	return s
}

func (s *ListRepositoriesRequest) SetOrderBy(v string) *ListRepositoriesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListRepositoriesRequest) SetOrganizationId(v string) *ListRepositoriesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoriesRequest) SetPage(v int64) *ListRepositoriesRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoriesRequest) SetPerPage(v int64) *ListRepositoriesRequest {
	s.PerPage = &v
	return s
}

func (s *ListRepositoriesRequest) SetSearch(v string) *ListRepositoriesRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoriesRequest) SetSort(v string) *ListRepositoriesRequest {
	s.Sort = &v
	return s
}

func (s *ListRepositoriesRequest) Validate() error {
	return dara.Validate(s)
}
