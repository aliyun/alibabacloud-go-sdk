// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTodosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListTodosRequest
	GetCollaborationGroupId() *string
	SetKeyword(v string) *ListTodosRequest
	GetKeyword() *string
	SetMaxResults(v int64) *ListTodosRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListTodosRequest
	GetNextToken() *string
	SetPage(v int64) *ListTodosRequest
	GetPage() *int64
	SetPageSize(v int64) *ListTodosRequest
	GetPageSize() *int64
	SetPriorities(v []*string) *ListTodosRequest
	GetPriorities() []*string
	SetReadStatus(v []*string) *ListTodosRequest
	GetReadStatus() []*string
	SetSortBy(v string) *ListTodosRequest
	GetSortBy() *string
	SetSourceTypes(v []*string) *ListTodosRequest
	GetSourceTypes() []*string
	SetStatuses(v []*string) *ListTodosRequest
	GetStatuses() []*string
	SetTenantId(v string) *ListTodosRequest
	GetTenantId() *string
}

type ListTodosRequest struct {
	// 协作空间 ID（可选过滤；不传 = 全部空间）
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 标题/摘要模糊搜索关键词
	//
	// example:
	//
	// 审批
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 单页最大返回数量（1~100）；传入时优先于 pageSize
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 翻页令牌，取上次响应返回的 nextToken；传入时优先于 page
	//
	// example:
	//
	// example
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 页码；传入 nextToken 时忽略
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数（1~100）
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 优先级筛选（HIGH/MEDIUM/LOW）
	//
	// example:
	//
	// HIGH
	Priorities []*string `json:"priorities,omitempty" xml:"priorities,omitempty" type:"Repeated"`
	// 读取状态筛选（UNREAD/READ，按处理人 readAt 判定）
	//
	// example:
	//
	// UNREAD
	ReadStatus []*string `json:"readStatus,omitempty" xml:"readStatus,omitempty" type:"Repeated"`
	// 排序方式
	//
	// example:
	//
	// DEFAULT
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// 待办类型筛选（WORKFLOW/OPEN_API）
	//
	// example:
	//
	// WORKFLOW
	SourceTypes []*string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty" type:"Repeated"`
	// 待办状态筛选（PENDING/COMPLETED/CANCELLED）
	//
	// example:
	//
	// PENDING
	Statuses []*string `json:"statuses,omitempty" xml:"statuses,omitempty" type:"Repeated"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListTodosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTodosRequest) GoString() string {
	return s.String()
}

func (s *ListTodosRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListTodosRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTodosRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListTodosRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTodosRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListTodosRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTodosRequest) GetPriorities() []*string {
	return s.Priorities
}

func (s *ListTodosRequest) GetReadStatus() []*string {
	return s.ReadStatus
}

func (s *ListTodosRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTodosRequest) GetSourceTypes() []*string {
	return s.SourceTypes
}

func (s *ListTodosRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListTodosRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListTodosRequest) SetCollaborationGroupId(v string) *ListTodosRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListTodosRequest) SetKeyword(v string) *ListTodosRequest {
	s.Keyword = &v
	return s
}

func (s *ListTodosRequest) SetMaxResults(v int64) *ListTodosRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTodosRequest) SetNextToken(v string) *ListTodosRequest {
	s.NextToken = &v
	return s
}

func (s *ListTodosRequest) SetPage(v int64) *ListTodosRequest {
	s.Page = &v
	return s
}

func (s *ListTodosRequest) SetPageSize(v int64) *ListTodosRequest {
	s.PageSize = &v
	return s
}

func (s *ListTodosRequest) SetPriorities(v []*string) *ListTodosRequest {
	s.Priorities = v
	return s
}

func (s *ListTodosRequest) SetReadStatus(v []*string) *ListTodosRequest {
	s.ReadStatus = v
	return s
}

func (s *ListTodosRequest) SetSortBy(v string) *ListTodosRequest {
	s.SortBy = &v
	return s
}

func (s *ListTodosRequest) SetSourceTypes(v []*string) *ListTodosRequest {
	s.SourceTypes = v
	return s
}

func (s *ListTodosRequest) SetStatuses(v []*string) *ListTodosRequest {
	s.Statuses = v
	return s
}

func (s *ListTodosRequest) SetTenantId(v string) *ListTodosRequest {
	s.TenantId = &v
	return s
}

func (s *ListTodosRequest) Validate() error {
	return dara.Validate(s)
}
