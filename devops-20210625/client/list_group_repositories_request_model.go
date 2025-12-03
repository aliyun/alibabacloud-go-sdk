// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupRepositoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListGroupRepositoriesRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListGroupRepositoriesRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListGroupRepositoriesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListGroupRepositoriesRequest
	GetPageSize() *int64
	SetSearch(v string) *ListGroupRepositoriesRequest
	GetSearch() *string
}

type ListGroupRepositoriesRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60ee8a814690c27532d412f8
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Demo
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
}

func (s ListGroupRepositoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListGroupRepositoriesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListGroupRepositoriesRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListGroupRepositoriesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListGroupRepositoriesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListGroupRepositoriesRequest) GetSearch() *string {
	return s.Search
}

func (s *ListGroupRepositoriesRequest) SetAccessToken(v string) *ListGroupRepositoriesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetOrganizationId(v string) *ListGroupRepositoriesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetPage(v int64) *ListGroupRepositoriesRequest {
	s.Page = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetPageSize(v int64) *ListGroupRepositoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetSearch(v string) *ListGroupRepositoriesRequest {
	s.Search = &v
	return s
}

func (s *ListGroupRepositoriesRequest) Validate() error {
	return dara.Validate(s)
}
