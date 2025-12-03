// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListRepositoryTagsRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListRepositoryTagsRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListRepositoryTagsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListRepositoryTagsRequest
	GetPageSize() *int64
	SetSearch(v string) *ListRepositoryTagsRequest
	GetSearch() *string
	SetSort(v string) *ListRepositoryTagsRequest
	GetSort() *string
}

type ListRepositoryTagsRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 611b75680fc7bf0dbe1dce55
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 2
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Demo
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// example:
	//
	// updated_desc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s ListRepositoryTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryTagsRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListRepositoryTagsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRepositoryTagsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListRepositoryTagsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRepositoryTagsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListRepositoryTagsRequest) GetSort() *string {
	return s.Sort
}

func (s *ListRepositoryTagsRequest) SetAccessToken(v string) *ListRepositoryTagsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetOrganizationId(v string) *ListRepositoryTagsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetPage(v int64) *ListRepositoryTagsRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetPageSize(v int64) *ListRepositoryTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetSearch(v string) *ListRepositoryTagsRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetSort(v string) *ListRepositoryTagsRequest {
	s.Sort = &v
	return s
}

func (s *ListRepositoryTagsRequest) Validate() error {
	return dara.Validate(s)
}
