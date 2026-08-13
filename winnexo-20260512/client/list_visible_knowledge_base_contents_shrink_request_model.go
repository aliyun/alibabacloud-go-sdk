// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBaseContentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetDirectoryId() *string
	SetOperatingObjectName(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetPageSize() *int64
	SetSortField(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetSortOrder() *string
	SetSourceTypesShrink(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetSourceTypesShrink() *string
	SetTenantId(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetTenantId() *string
}

type ListVisibleKnowledgeBaseContentsShrinkRequest struct {
	// 目录 ID（必传非空，必须在数字员工 linked_directories 及其子目录范围内）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 数字员工名称（运营对象 name）
	//
	// This parameter is required.
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
	// 资源类型筛选列表（有值时仅返回资源，不包含子目录）
	//
	// example:
	//
	// string_value
	SourceTypesShrink *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListVisibleKnowledgeBaseContentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBaseContentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetSourceTypesShrink() *string {
	return s.SourceTypesShrink
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetDirectoryId(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetOperatingObjectName(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetPage(v int64) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetPageSize(v int64) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetSortField(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetSortOrder(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetSourceTypesShrink(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.SourceTypesShrink = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetTenantId(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
