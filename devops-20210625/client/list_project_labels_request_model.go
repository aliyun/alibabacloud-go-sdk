// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListProjectLabelsRequest
	GetAccessToken() *string
	SetOrderBy(v string) *ListProjectLabelsRequest
	GetOrderBy() *string
	SetOrganizationId(v string) *ListProjectLabelsRequest
	GetOrganizationId() *string
	SetPage(v int32) *ListProjectLabelsRequest
	GetPage() *int32
	SetPageSize(v int64) *ListProjectLabelsRequest
	GetPageSize() *int64
	SetRepositoryIdentity(v string) *ListProjectLabelsRequest
	GetRepositoryIdentity() *string
	SetSearch(v string) *ListProjectLabelsRequest
	GetSearch() *string
	SetSort(v string) *ListProjectLabelsRequest
	GetSort() *string
	SetWithCounts(v bool) *ListProjectLabelsRequest
	GetWithCounts() *bool
}

type ListProjectLabelsRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// label_name
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
	// example:
	//
	// TEST
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// example:
	//
	// false
	WithCounts *bool `json:"withCounts,omitempty" xml:"withCounts,omitempty"`
}

func (s ListProjectLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectLabelsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectLabelsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListProjectLabelsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListProjectLabelsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListProjectLabelsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListProjectLabelsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProjectLabelsRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *ListProjectLabelsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListProjectLabelsRequest) GetSort() *string {
	return s.Sort
}

func (s *ListProjectLabelsRequest) GetWithCounts() *bool {
	return s.WithCounts
}

func (s *ListProjectLabelsRequest) SetAccessToken(v string) *ListProjectLabelsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListProjectLabelsRequest) SetOrderBy(v string) *ListProjectLabelsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListProjectLabelsRequest) SetOrganizationId(v string) *ListProjectLabelsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListProjectLabelsRequest) SetPage(v int32) *ListProjectLabelsRequest {
	s.Page = &v
	return s
}

func (s *ListProjectLabelsRequest) SetPageSize(v int64) *ListProjectLabelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectLabelsRequest) SetRepositoryIdentity(v string) *ListProjectLabelsRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *ListProjectLabelsRequest) SetSearch(v string) *ListProjectLabelsRequest {
	s.Search = &v
	return s
}

func (s *ListProjectLabelsRequest) SetSort(v string) *ListProjectLabelsRequest {
	s.Sort = &v
	return s
}

func (s *ListProjectLabelsRequest) SetWithCounts(v bool) *ListProjectLabelsRequest {
	s.WithCounts = &v
	return s
}

func (s *ListProjectLabelsRequest) Validate() error {
	return dara.Validate(s)
}
