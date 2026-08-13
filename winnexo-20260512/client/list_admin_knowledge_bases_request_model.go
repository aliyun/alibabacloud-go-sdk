// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdminKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListAdminKnowledgeBasesRequest
	GetDirectoryId() *string
	SetKeyword(v string) *ListAdminKnowledgeBasesRequest
	GetKeyword() *string
	SetPage(v int64) *ListAdminKnowledgeBasesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListAdminKnowledgeBasesRequest
	GetPageSize() *int64
	SetSortField(v string) *ListAdminKnowledgeBasesRequest
	GetSortField() *string
	SetSortOrder(v string) *ListAdminKnowledgeBasesRequest
	GetSortOrder() *string
	SetSourceTypes(v []*string) *ListAdminKnowledgeBasesRequest
	GetSourceTypes() []*string
	SetTenantId(v string) *ListAdminKnowledgeBasesRequest
	GetTenantId() *string
}

type ListAdminKnowledgeBasesRequest struct {
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
	SourceTypes []*string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty" type:"Repeated"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAdminKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAdminKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListAdminKnowledgeBasesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListAdminKnowledgeBasesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAdminKnowledgeBasesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListAdminKnowledgeBasesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAdminKnowledgeBasesRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListAdminKnowledgeBasesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListAdminKnowledgeBasesRequest) GetSourceTypes() []*string {
	return s.SourceTypes
}

func (s *ListAdminKnowledgeBasesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAdminKnowledgeBasesRequest) SetDirectoryId(v string) *ListAdminKnowledgeBasesRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetKeyword(v string) *ListAdminKnowledgeBasesRequest {
	s.Keyword = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetPage(v int64) *ListAdminKnowledgeBasesRequest {
	s.Page = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetPageSize(v int64) *ListAdminKnowledgeBasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetSortField(v string) *ListAdminKnowledgeBasesRequest {
	s.SortField = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetSortOrder(v string) *ListAdminKnowledgeBasesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetSourceTypes(v []*string) *ListAdminKnowledgeBasesRequest {
	s.SourceTypes = v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetTenantId(v string) *ListAdminKnowledgeBasesRequest {
	s.TenantId = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
