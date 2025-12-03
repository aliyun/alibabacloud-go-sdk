// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListApplicationsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListApplicationsRequest
	GetOrderBy() *string
	SetOrganizationId(v string) *ListApplicationsRequest
	GetOrganizationId() *string
	SetPagination(v string) *ListApplicationsRequest
	GetPagination() *string
	SetPerPage(v int32) *ListApplicationsRequest
	GetPerPage() *int32
	SetSort(v string) *ListApplicationsRequest
	GetSort() *string
}

type ListApplicationsRequest struct {
	// example:
	//
	// vxc2341gfssad12
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// asc
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
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
	// id
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListApplicationsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListApplicationsRequest) GetPagination() *string {
	return s.Pagination
}

func (s *ListApplicationsRequest) GetPerPage() *int32 {
	return s.PerPage
}

func (s *ListApplicationsRequest) GetSort() *string {
	return s.Sort
}

func (s *ListApplicationsRequest) SetNextToken(v string) *ListApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsRequest) SetOrderBy(v string) *ListApplicationsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListApplicationsRequest) SetOrganizationId(v string) *ListApplicationsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListApplicationsRequest) SetPagination(v string) *ListApplicationsRequest {
	s.Pagination = &v
	return s
}

func (s *ListApplicationsRequest) SetPerPage(v int32) *ListApplicationsRequest {
	s.PerPage = &v
	return s
}

func (s *ListApplicationsRequest) SetSort(v string) *ListApplicationsRequest {
	s.Sort = &v
	return s
}

func (s *ListApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
