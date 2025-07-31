// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePermissionOperationLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListResourcePermissionOperationLogRequestListQuery) *ListResourcePermissionOperationLogRequest
	GetListQuery() *ListResourcePermissionOperationLogRequestListQuery
	SetOpTenantId(v int64) *ListResourcePermissionOperationLogRequest
	GetOpTenantId() *int64
}

type ListResourcePermissionOperationLogRequest struct {
	// This parameter is required.
	ListQuery *ListResourcePermissionOperationLogRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListResourcePermissionOperationLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogRequest) GetListQuery() *ListResourcePermissionOperationLogRequestListQuery {
	return s.ListQuery
}

func (s *ListResourcePermissionOperationLogRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListResourcePermissionOperationLogRequest) SetListQuery(v *ListResourcePermissionOperationLogRequestListQuery) *ListResourcePermissionOperationLogRequest {
	s.ListQuery = v
	return s
}

func (s *ListResourcePermissionOperationLogRequest) SetOpTenantId(v int64) *ListResourcePermissionOperationLogRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListResourcePermissionOperationLogRequest) Validate() error {
	return dara.Validate(s)
}

type ListResourcePermissionOperationLogRequestListQuery struct {
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

func (s ListResourcePermissionOperationLogRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListResourcePermissionOperationLogRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcePermissionOperationLogRequestListQuery) GetSearchText() *string {
	return s.SearchText
}

func (s *ListResourcePermissionOperationLogRequestListQuery) GetTabType() *string {
	return s.TabType
}

func (s *ListResourcePermissionOperationLogRequestListQuery) SetPage(v int32) *ListResourcePermissionOperationLogRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListResourcePermissionOperationLogRequestListQuery) SetPageSize(v int32) *ListResourcePermissionOperationLogRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListResourcePermissionOperationLogRequestListQuery) SetSearchText(v string) *ListResourcePermissionOperationLogRequestListQuery {
	s.SearchText = &v
	return s
}

func (s *ListResourcePermissionOperationLogRequestListQuery) SetTabType(v string) *ListResourcePermissionOperationLogRequestListQuery {
	s.TabType = &v
	return s
}

func (s *ListResourcePermissionOperationLogRequestListQuery) Validate() error {
	return dara.Validate(s)
}
