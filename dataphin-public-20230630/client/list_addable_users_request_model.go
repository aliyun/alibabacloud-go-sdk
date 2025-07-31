// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddableUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListAddableUsersRequestListQuery) *ListAddableUsersRequest
	GetListQuery() *ListAddableUsersRequestListQuery
	SetOpTenantId(v int64) *ListAddableUsersRequest
	GetOpTenantId() *int64
}

type ListAddableUsersRequest struct {
	// This parameter is required.
	ListQuery *ListAddableUsersRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListAddableUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAddableUsersRequest) GoString() string {
	return s.String()
}

func (s *ListAddableUsersRequest) GetListQuery() *ListAddableUsersRequestListQuery {
	return s.ListQuery
}

func (s *ListAddableUsersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAddableUsersRequest) SetListQuery(v *ListAddableUsersRequestListQuery) *ListAddableUsersRequest {
	s.ListQuery = v
	return s
}

func (s *ListAddableUsersRequest) SetOpTenantId(v int64) *ListAddableUsersRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAddableUsersRequest) Validate() error {
	return dara.Validate(s)
}

type ListAddableUsersRequestListQuery struct {
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
}

func (s ListAddableUsersRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListAddableUsersRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListAddableUsersRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListAddableUsersRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAddableUsersRequestListQuery) GetSearchText() *string {
	return s.SearchText
}

func (s *ListAddableUsersRequestListQuery) SetPage(v int32) *ListAddableUsersRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListAddableUsersRequestListQuery) SetPageSize(v int32) *ListAddableUsersRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListAddableUsersRequestListQuery) SetSearchText(v string) *ListAddableUsersRequestListQuery {
	s.SearchText = &v
	return s
}

func (s *ListAddableUsersRequestListQuery) Validate() error {
	return dara.Validate(s)
}
