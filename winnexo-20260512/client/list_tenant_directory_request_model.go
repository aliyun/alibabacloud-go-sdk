// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListTenantDirectoryRequest
	GetDirectoryId() *string
	SetPage(v int64) *ListTenantDirectoryRequest
	GetPage() *int64
	SetPageSize(v int64) *ListTenantDirectoryRequest
	GetPageSize() *int64
	SetSortField(v string) *ListTenantDirectoryRequest
	GetSortField() *string
	SetSortOrder(v string) *ListTenantDirectoryRequest
	GetSortOrder() *string
	SetSourceTypes(v string) *ListTenantDirectoryRequest
	GetSourceTypes() *string
	SetTenantId(v string) *ListTenantDirectoryRequest
	GetTenantId() *string
}

type ListTenantDirectoryRequest struct {
	// 目录唯一标识；不传或传 root 时查询知识库根目录列表
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 排序字段
	//
	// example:
	//
	// name
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// 排序方向
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// 资源类型过滤，多个类型使用逗号分隔；传入后只返回资源
	//
	// example:
	//
	// string_value
	SourceTypes *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListTenantDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTenantDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ListTenantDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListTenantDirectoryRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListTenantDirectoryRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTenantDirectoryRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListTenantDirectoryRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListTenantDirectoryRequest) GetSourceTypes() *string {
	return s.SourceTypes
}

func (s *ListTenantDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListTenantDirectoryRequest) SetDirectoryId(v string) *ListTenantDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetPage(v int64) *ListTenantDirectoryRequest {
	s.Page = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetPageSize(v int64) *ListTenantDirectoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetSortField(v string) *ListTenantDirectoryRequest {
	s.SortField = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetSortOrder(v string) *ListTenantDirectoryRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetSourceTypes(v string) *ListTenantDirectoryRequest {
	s.SourceTypes = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetTenantId(v string) *ListTenantDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *ListTenantDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
