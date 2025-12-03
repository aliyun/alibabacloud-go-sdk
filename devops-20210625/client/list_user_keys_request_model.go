// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListUserKeysRequest
	GetAccessToken() *string
	SetOrderBy(v string) *ListUserKeysRequest
	GetOrderBy() *string
	SetOrganizationId(v string) *ListUserKeysRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListUserKeysRequest
	GetPage() *int64
	SetPageSize(v int64) *ListUserKeysRequest
	GetPageSize() *int64
	SetSort(v string) *ListUserKeysRequest
	GetSort() *string
}

type ListUserKeysRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// created_at
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
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s ListUserKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserKeysRequest) GoString() string {
	return s.String()
}

func (s *ListUserKeysRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListUserKeysRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListUserKeysRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListUserKeysRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListUserKeysRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUserKeysRequest) GetSort() *string {
	return s.Sort
}

func (s *ListUserKeysRequest) SetAccessToken(v string) *ListUserKeysRequest {
	s.AccessToken = &v
	return s
}

func (s *ListUserKeysRequest) SetOrderBy(v string) *ListUserKeysRequest {
	s.OrderBy = &v
	return s
}

func (s *ListUserKeysRequest) SetOrganizationId(v string) *ListUserKeysRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListUserKeysRequest) SetPage(v int64) *ListUserKeysRequest {
	s.Page = &v
	return s
}

func (s *ListUserKeysRequest) SetPageSize(v int64) *ListUserKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserKeysRequest) SetSort(v string) *ListUserKeysRequest {
	s.Sort = &v
	return s
}

func (s *ListUserKeysRequest) Validate() error {
	return dara.Validate(s)
}
