// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalDirectoryContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListPersonalDirectoryContentsRequest
	GetDirectoryId() *string
	SetOperatingObjectName(v string) *ListPersonalDirectoryContentsRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListPersonalDirectoryContentsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListPersonalDirectoryContentsRequest
	GetPageSize() *int64
	SetSortField(v string) *ListPersonalDirectoryContentsRequest
	GetSortField() *string
	SetSortOrder(v string) *ListPersonalDirectoryContentsRequest
	GetSortOrder() *string
	SetSourceTypes(v []*string) *ListPersonalDirectoryContentsRequest
	GetSourceTypes() []*string
	SetTenantId(v string) *ListPersonalDirectoryContentsRequest
	GetTenantId() *string
}

type ListPersonalDirectoryContentsRequest struct {
	// 目录 ID（必传非空）；传 \"root\" 时 service 自动解析当前数字员工的默认根目录并返回其下内容（首屏知识库卡片场景），传具体目录 ID 时返回该目录下子目录与资源
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 数字员工名称（已废弃：不再作为个人资源隔离条件，仅保留用于来源追溯）
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 页码（从 1 开始）
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量，范围 1-100
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 排序字段，可选 name / gmt_create / gmt_modified
	//
	// example:
	//
	// name
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// 排序方向，可选 asc / desc
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// 资源类型筛选列表（有值时仅返回资源，不包含子目录）；支持虚拟类型 OUTPUT（产出保存的资源，service 自动反查关联表）
	//
	// example:
	//
	// string_value
	SourceTypes []*string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty" type:"Repeated"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListPersonalDirectoryContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalDirectoryContentsRequest) GoString() string {
	return s.String()
}

func (s *ListPersonalDirectoryContentsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListPersonalDirectoryContentsRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListPersonalDirectoryContentsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListPersonalDirectoryContentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPersonalDirectoryContentsRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListPersonalDirectoryContentsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListPersonalDirectoryContentsRequest) GetSourceTypes() []*string {
	return s.SourceTypes
}

func (s *ListPersonalDirectoryContentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListPersonalDirectoryContentsRequest) SetDirectoryId(v string) *ListPersonalDirectoryContentsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetOperatingObjectName(v string) *ListPersonalDirectoryContentsRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetPage(v int64) *ListPersonalDirectoryContentsRequest {
	s.Page = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetPageSize(v int64) *ListPersonalDirectoryContentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetSortField(v string) *ListPersonalDirectoryContentsRequest {
	s.SortField = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetSortOrder(v string) *ListPersonalDirectoryContentsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetSourceTypes(v []*string) *ListPersonalDirectoryContentsRequest {
	s.SourceTypes = v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetTenantId(v string) *ListPersonalDirectoryContentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) Validate() error {
	return dara.Validate(s)
}
