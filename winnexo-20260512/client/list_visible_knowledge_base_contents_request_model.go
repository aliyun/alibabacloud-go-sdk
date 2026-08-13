// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBaseContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListVisibleKnowledgeBaseContentsRequest
	GetDirectoryId() *string
	SetOperatingObjectName(v string) *ListVisibleKnowledgeBaseContentsRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListVisibleKnowledgeBaseContentsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListVisibleKnowledgeBaseContentsRequest
	GetPageSize() *int64
	SetSortField(v string) *ListVisibleKnowledgeBaseContentsRequest
	GetSortField() *string
	SetSortOrder(v string) *ListVisibleKnowledgeBaseContentsRequest
	GetSortOrder() *string
	SetSourceTypes(v []*string) *ListVisibleKnowledgeBaseContentsRequest
	GetSourceTypes() []*string
	SetTenantId(v string) *ListVisibleKnowledgeBaseContentsRequest
	GetTenantId() *string
}

type ListVisibleKnowledgeBaseContentsRequest struct {
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
	SourceTypes []*string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty" type:"Repeated"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListVisibleKnowledgeBaseContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBaseContentsRequest) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetSourceTypes() []*string {
	return s.SourceTypes
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetDirectoryId(v string) *ListVisibleKnowledgeBaseContentsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetOperatingObjectName(v string) *ListVisibleKnowledgeBaseContentsRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetPage(v int64) *ListVisibleKnowledgeBaseContentsRequest {
	s.Page = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetPageSize(v int64) *ListVisibleKnowledgeBaseContentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetSortField(v string) *ListVisibleKnowledgeBaseContentsRequest {
	s.SortField = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetSortOrder(v string) *ListVisibleKnowledgeBaseContentsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetSourceTypes(v []*string) *ListVisibleKnowledgeBaseContentsRequest {
	s.SourceTypes = v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetTenantId(v string) *ListVisibleKnowledgeBaseContentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) Validate() error {
	return dara.Validate(s)
}
