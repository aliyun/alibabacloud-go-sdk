// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeRequestsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppNameListShrink(v string) *ListChangeRequestsShrinkRequest
	GetAppNameListShrink() *string
	SetDisplayNameKeyword(v string) *ListChangeRequestsShrinkRequest
	GetDisplayNameKeyword() *string
	SetNextToken(v string) *ListChangeRequestsShrinkRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListChangeRequestsShrinkRequest
	GetOrderBy() *string
	SetOrganizationId(v string) *ListChangeRequestsShrinkRequest
	GetOrganizationId() *string
	SetOwnerIdListShrink(v string) *ListChangeRequestsShrinkRequest
	GetOwnerIdListShrink() *string
	SetPage(v int32) *ListChangeRequestsShrinkRequest
	GetPage() *int32
	SetPagination(v string) *ListChangeRequestsShrinkRequest
	GetPagination() *string
	SetPerPage(v int32) *ListChangeRequestsShrinkRequest
	GetPerPage() *int32
	SetSort(v string) *ListChangeRequestsShrinkRequest
	GetSort() *string
	SetStateListShrink(v string) *ListChangeRequestsShrinkRequest
	GetStateListShrink() *string
}

type ListChangeRequestsShrinkRequest struct {
	AppNameListShrink *string `json:"appNameList,omitempty" xml:"appNameList,omitempty"`
	// example:
	//
	// change1
	DisplayNameKeyword *string `json:"displayNameKeyword,omitempty" xml:"displayNameKeyword,omitempty"`
	// example:
	//
	// 4dc150725770510122396e2476
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// id
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId    *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	OwnerIdListShrink *string `json:"ownerIdList,omitempty" xml:"ownerIdList,omitempty"`
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
	Sort            *string `json:"sort,omitempty" xml:"sort,omitempty"`
	StateListShrink *string `json:"stateList,omitempty" xml:"stateList,omitempty"`
}

func (s ListChangeRequestsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChangeRequestsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListChangeRequestsShrinkRequest) GetAppNameListShrink() *string {
	return s.AppNameListShrink
}

func (s *ListChangeRequestsShrinkRequest) GetDisplayNameKeyword() *string {
	return s.DisplayNameKeyword
}

func (s *ListChangeRequestsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListChangeRequestsShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListChangeRequestsShrinkRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListChangeRequestsShrinkRequest) GetOwnerIdListShrink() *string {
	return s.OwnerIdListShrink
}

func (s *ListChangeRequestsShrinkRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListChangeRequestsShrinkRequest) GetPagination() *string {
	return s.Pagination
}

func (s *ListChangeRequestsShrinkRequest) GetPerPage() *int32 {
	return s.PerPage
}

func (s *ListChangeRequestsShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *ListChangeRequestsShrinkRequest) GetStateListShrink() *string {
	return s.StateListShrink
}

func (s *ListChangeRequestsShrinkRequest) SetAppNameListShrink(v string) *ListChangeRequestsShrinkRequest {
	s.AppNameListShrink = &v
	return s
}

func (s *ListChangeRequestsShrinkRequest) SetDisplayNameKeyword(v string) *ListChangeRequestsShrinkRequest {
	s.DisplayNameKeyword = &v
	return s
}

func (s *ListChangeRequestsShrinkRequest) SetNextToken(v string) *ListChangeRequestsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListChangeRequestsShrinkRequest) SetOrderBy(v string) *ListChangeRequestsShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListChangeRequestsShrinkRequest) SetOrganizationId(v string) *ListChangeRequestsShrinkRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListChangeRequestsShrinkRequest) SetOwnerIdListShrink(v string) *ListChangeRequestsShrinkRequest {
	s.OwnerIdListShrink = &v
	return s
}

func (s *ListChangeRequestsShrinkRequest) SetPage(v int32) *ListChangeRequestsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListChangeRequestsShrinkRequest) SetPagination(v string) *ListChangeRequestsShrinkRequest {
	s.Pagination = &v
	return s
}

func (s *ListChangeRequestsShrinkRequest) SetPerPage(v int32) *ListChangeRequestsShrinkRequest {
	s.PerPage = &v
	return s
}

func (s *ListChangeRequestsShrinkRequest) SetSort(v string) *ListChangeRequestsShrinkRequest {
	s.Sort = &v
	return s
}

func (s *ListChangeRequestsShrinkRequest) SetStateListShrink(v string) *ListChangeRequestsShrinkRequest {
	s.StateListShrink = &v
	return s
}

func (s *ListChangeRequestsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
