// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppReleaseStageExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListAppReleaseStageExecutionsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListAppReleaseStageExecutionsRequest
	GetOrderBy() *string
	SetOrganizationId(v string) *ListAppReleaseStageExecutionsRequest
	GetOrganizationId() *string
	SetPage(v int32) *ListAppReleaseStageExecutionsRequest
	GetPage() *int32
	SetPagination(v string) *ListAppReleaseStageExecutionsRequest
	GetPagination() *string
	SetPerPage(v int32) *ListAppReleaseStageExecutionsRequest
	GetPerPage() *int32
	SetSort(v string) *ListAppReleaseStageExecutionsRequest
	GetSort() *string
}

type ListAppReleaseStageExecutionsRequest struct {
	// example:
	//
	// ""
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// gmtCreate
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// keyset
	Pagination *string `json:"pagination,omitempty" xml:"pagination,omitempty"`
	// example:
	//
	// 20
	PerPage *int32 `json:"perPage,omitempty" xml:"perPage,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s ListAppReleaseStageExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppReleaseStageExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListAppReleaseStageExecutionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppReleaseStageExecutionsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListAppReleaseStageExecutionsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListAppReleaseStageExecutionsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListAppReleaseStageExecutionsRequest) GetPagination() *string {
	return s.Pagination
}

func (s *ListAppReleaseStageExecutionsRequest) GetPerPage() *int32 {
	return s.PerPage
}

func (s *ListAppReleaseStageExecutionsRequest) GetSort() *string {
	return s.Sort
}

func (s *ListAppReleaseStageExecutionsRequest) SetNextToken(v string) *ListAppReleaseStageExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppReleaseStageExecutionsRequest) SetOrderBy(v string) *ListAppReleaseStageExecutionsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListAppReleaseStageExecutionsRequest) SetOrganizationId(v string) *ListAppReleaseStageExecutionsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListAppReleaseStageExecutionsRequest) SetPage(v int32) *ListAppReleaseStageExecutionsRequest {
	s.Page = &v
	return s
}

func (s *ListAppReleaseStageExecutionsRequest) SetPagination(v string) *ListAppReleaseStageExecutionsRequest {
	s.Pagination = &v
	return s
}

func (s *ListAppReleaseStageExecutionsRequest) SetPerPage(v int32) *ListAppReleaseStageExecutionsRequest {
	s.PerPage = &v
	return s
}

func (s *ListAppReleaseStageExecutionsRequest) SetSort(v string) *ListAppReleaseStageExecutionsRequest {
	s.Sort = &v
	return s
}

func (s *ListAppReleaseStageExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
