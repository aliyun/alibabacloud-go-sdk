// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListResourcePermissionsRequestListQuery) *ListResourcePermissionsRequest
	GetListQuery() *ListResourcePermissionsRequestListQuery
	SetOpTenantId(v int64) *ListResourcePermissionsRequest
	GetOpTenantId() *int64
}

type ListResourcePermissionsRequest struct {
	// This parameter is required.
	ListQuery *ListResourcePermissionsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListResourcePermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsRequest) GetListQuery() *ListResourcePermissionsRequestListQuery {
	return s.ListQuery
}

func (s *ListResourcePermissionsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListResourcePermissionsRequest) SetListQuery(v *ListResourcePermissionsRequestListQuery) *ListResourcePermissionsRequest {
	s.ListQuery = v
	return s
}

func (s *ListResourcePermissionsRequest) SetOpTenantId(v int64) *ListResourcePermissionsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListResourcePermissionsRequest) Validate() error {
	return dara.Validate(s)
}

type ListResourcePermissionsRequestListQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx测试
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	TabType *string `json:"TabType,omitempty" xml:"TabType,omitempty"`
}

func (s ListResourcePermissionsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListResourcePermissionsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcePermissionsRequestListQuery) GetSearchText() *string {
	return s.SearchText
}

func (s *ListResourcePermissionsRequestListQuery) GetTabType() *string {
	return s.TabType
}

func (s *ListResourcePermissionsRequestListQuery) SetPage(v int32) *ListResourcePermissionsRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListResourcePermissionsRequestListQuery) SetPageSize(v int32) *ListResourcePermissionsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListResourcePermissionsRequestListQuery) SetSearchText(v string) *ListResourcePermissionsRequestListQuery {
	s.SearchText = &v
	return s
}

func (s *ListResourcePermissionsRequestListQuery) SetTabType(v string) *ListResourcePermissionsRequestListQuery {
	s.TabType = &v
	return s
}

func (s *ListResourcePermissionsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
