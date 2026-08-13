// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdminKnowledgeBasesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetDirectoryId() *string
	SetKeyword(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetKeyword() *string
	SetPage(v int64) *ListAdminKnowledgeBasesShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListAdminKnowledgeBasesShrinkRequest
	GetPageSize() *int64
	SetSortField(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetSortOrder() *string
	SetSourceTypesShrink(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetSourceTypesShrink() *string
	SetTenantId(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetTenantId() *string
}

type ListAdminKnowledgeBasesShrinkRequest struct {
	// 目录 ID；为空或 \"root\" 时返回 KB 顶层列表，传具体值时下钻返回该目录的子目录 + 资源（混合分页，由 itemType 区分）
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 搜索关键词，仅在 directoryId 为空/root 时生效，模糊匹配 KB 名称或描述（忽略大小写）
	//
	// example:
	//
	// 示例关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 页码，从 1 开始
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
	// 排序字段：name / gmtCreate / gmtModified；非法值回退为 name
	//
	// example:
	//
	// name
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// 排序方向：asc / desc；非法值回退为 asc
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// 资源类型过滤，仅在下钻（directoryId 非空）时生效；命中时仅返回匹配类型的资源，不含子目录
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

func (s ListAdminKnowledgeBasesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAdminKnowledgeBasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetSourceTypesShrink() *string {
	return s.SourceTypesShrink
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetDirectoryId(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetKeyword(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetPage(v int64) *ListAdminKnowledgeBasesShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetPageSize(v int64) *ListAdminKnowledgeBasesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetSortField(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetSortOrder(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetSourceTypesShrink(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.SourceTypesShrink = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetTenantId(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
