// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryBranchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListRepositoryBranchesRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListRepositoryBranchesRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListRepositoryBranchesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListRepositoryBranchesRequest
	GetPageSize() *int64
	SetSearch(v string) *ListRepositoryBranchesRequest
	GetSearch() *string
	SetSort(v string) *ListRepositoryBranchesRequest
	GetSort() *string
}

type ListRepositoryBranchesRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
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
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Demo
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s ListRepositoryBranchesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryBranchesRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListRepositoryBranchesRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRepositoryBranchesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListRepositoryBranchesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRepositoryBranchesRequest) GetSearch() *string {
	return s.Search
}

func (s *ListRepositoryBranchesRequest) GetSort() *string {
	return s.Sort
}

func (s *ListRepositoryBranchesRequest) SetAccessToken(v string) *ListRepositoryBranchesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetOrganizationId(v string) *ListRepositoryBranchesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetPage(v int64) *ListRepositoryBranchesRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetPageSize(v int64) *ListRepositoryBranchesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetSearch(v string) *ListRepositoryBranchesRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetSort(v string) *ListRepositoryBranchesRequest {
	s.Sort = &v
	return s
}

func (s *ListRepositoryBranchesRequest) Validate() error {
	return dara.Validate(s)
}
