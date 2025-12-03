// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeRequestsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppNameList(v []*string) *ListChangeRequestsRequest
	GetAppNameList() []*string
	SetDisplayNameKeyword(v string) *ListChangeRequestsRequest
	GetDisplayNameKeyword() *string
	SetNextToken(v string) *ListChangeRequestsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListChangeRequestsRequest
	GetOrderBy() *string
	SetOrganizationId(v string) *ListChangeRequestsRequest
	GetOrganizationId() *string
	SetOwnerIdList(v []*string) *ListChangeRequestsRequest
	GetOwnerIdList() []*string
	SetPage(v int32) *ListChangeRequestsRequest
	GetPage() *int32
	SetPagination(v string) *ListChangeRequestsRequest
	GetPagination() *string
	SetPerPage(v int32) *ListChangeRequestsRequest
	GetPerPage() *int32
	SetSort(v string) *ListChangeRequestsRequest
	GetSort() *string
	SetStateList(v []*string) *ListChangeRequestsRequest
	GetStateList() []*string
}

type ListChangeRequestsRequest struct {
	AppNameList []*string `json:"appNameList,omitempty" xml:"appNameList,omitempty" type:"Repeated"`
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
	OrganizationId *string   `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	OwnerIdList    []*string `json:"ownerIdList,omitempty" xml:"ownerIdList,omitempty" type:"Repeated"`
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
	Sort      *string   `json:"sort,omitempty" xml:"sort,omitempty"`
	StateList []*string `json:"stateList,omitempty" xml:"stateList,omitempty" type:"Repeated"`
}

func (s ListChangeRequestsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChangeRequestsRequest) GoString() string {
	return s.String()
}

func (s *ListChangeRequestsRequest) GetAppNameList() []*string {
	return s.AppNameList
}

func (s *ListChangeRequestsRequest) GetDisplayNameKeyword() *string {
	return s.DisplayNameKeyword
}

func (s *ListChangeRequestsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListChangeRequestsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListChangeRequestsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListChangeRequestsRequest) GetOwnerIdList() []*string {
	return s.OwnerIdList
}

func (s *ListChangeRequestsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListChangeRequestsRequest) GetPagination() *string {
	return s.Pagination
}

func (s *ListChangeRequestsRequest) GetPerPage() *int32 {
	return s.PerPage
}

func (s *ListChangeRequestsRequest) GetSort() *string {
	return s.Sort
}

func (s *ListChangeRequestsRequest) GetStateList() []*string {
	return s.StateList
}

func (s *ListChangeRequestsRequest) SetAppNameList(v []*string) *ListChangeRequestsRequest {
	s.AppNameList = v
	return s
}

func (s *ListChangeRequestsRequest) SetDisplayNameKeyword(v string) *ListChangeRequestsRequest {
	s.DisplayNameKeyword = &v
	return s
}

func (s *ListChangeRequestsRequest) SetNextToken(v string) *ListChangeRequestsRequest {
	s.NextToken = &v
	return s
}

func (s *ListChangeRequestsRequest) SetOrderBy(v string) *ListChangeRequestsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListChangeRequestsRequest) SetOrganizationId(v string) *ListChangeRequestsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListChangeRequestsRequest) SetOwnerIdList(v []*string) *ListChangeRequestsRequest {
	s.OwnerIdList = v
	return s
}

func (s *ListChangeRequestsRequest) SetPage(v int32) *ListChangeRequestsRequest {
	s.Page = &v
	return s
}

func (s *ListChangeRequestsRequest) SetPagination(v string) *ListChangeRequestsRequest {
	s.Pagination = &v
	return s
}

func (s *ListChangeRequestsRequest) SetPerPage(v int32) *ListChangeRequestsRequest {
	s.PerPage = &v
	return s
}

func (s *ListChangeRequestsRequest) SetSort(v string) *ListChangeRequestsRequest {
	s.Sort = &v
	return s
}

func (s *ListChangeRequestsRequest) SetStateList(v []*string) *ListChangeRequestsRequest {
	s.StateList = v
	return s
}

func (s *ListChangeRequestsRequest) Validate() error {
	return dara.Validate(s)
}
