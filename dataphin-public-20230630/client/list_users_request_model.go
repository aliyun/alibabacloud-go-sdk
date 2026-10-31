// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListUsersRequestListQuery) *ListUsersRequest
	GetListQuery() *ListUsersRequestListQuery
	SetOpTenantId(v int64) *ListUsersRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListUsersRequest
	GetOpUserId() *string
}

type ListUsersRequest struct {
	// The paged query for paging.
	//
	// This parameter is required.
	ListQuery *ListUsersRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetListQuery() *ListUsersRequestListQuery {
	return s.ListQuery
}

func (s *ListUsersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListUsersRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListUsersRequest) SetListQuery(v *ListUsersRequestListQuery) *ListUsersRequest {
	s.ListQuery = v
	return s
}

func (s *ListUsersRequest) SetOpTenantId(v int64) *ListUsersRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListUsersRequest) SetOpUserId(v string) *ListUsersRequest {
	s.OpUserId = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUsersRequestListQuery struct {
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of records per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// John Doe
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
}

func (s ListUsersRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListUsersRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListUsersRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUsersRequestListQuery) GetSearchText() *string {
	return s.SearchText
}

func (s *ListUsersRequestListQuery) SetPage(v int32) *ListUsersRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListUsersRequestListQuery) SetPageSize(v int32) *ListUsersRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequestListQuery) SetSearchText(v string) *ListUsersRequestListQuery {
	s.SearchText = &v
	return s
}

func (s *ListUsersRequestListQuery) Validate() error {
	return dara.Validate(s)
}
