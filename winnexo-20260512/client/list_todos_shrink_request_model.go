// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTodosShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListTodosShrinkRequest
	GetCollaborationGroupId() *string
	SetKeyword(v string) *ListTodosShrinkRequest
	GetKeyword() *string
	SetMaxResults(v int64) *ListTodosShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListTodosShrinkRequest
	GetNextToken() *string
	SetPage(v int64) *ListTodosShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListTodosShrinkRequest
	GetPageSize() *int64
	SetPrioritiesShrink(v string) *ListTodosShrinkRequest
	GetPrioritiesShrink() *string
	SetReadStatusShrink(v string) *ListTodosShrinkRequest
	GetReadStatusShrink() *string
	SetSortBy(v string) *ListTodosShrinkRequest
	GetSortBy() *string
	SetSourceTypesShrink(v string) *ListTodosShrinkRequest
	GetSourceTypesShrink() *string
	SetStatusesShrink(v string) *ListTodosShrinkRequest
	GetStatusesShrink() *string
	SetTenantId(v string) *ListTodosShrinkRequest
	GetTenantId() *string
}

type ListTodosShrinkRequest struct {
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
	PrioritiesShrink *string `json:"priorities,omitempty" xml:"priorities,omitempty"`
	// 读取状态筛选（UNREAD/READ，按处理人 readAt 判定）
	//
	// example:
	//
	// UNREAD
	ReadStatusShrink *string `json:"readStatus,omitempty" xml:"readStatus,omitempty"`
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
	SourceTypesShrink *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// 待办状态筛选（PENDING/COMPLETED/CANCELLED）
	//
	// example:
	//
	// PENDING
	StatusesShrink *string `json:"statuses,omitempty" xml:"statuses,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListTodosShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTodosShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTodosShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListTodosShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTodosShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListTodosShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTodosShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListTodosShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTodosShrinkRequest) GetPrioritiesShrink() *string {
	return s.PrioritiesShrink
}

func (s *ListTodosShrinkRequest) GetReadStatusShrink() *string {
	return s.ReadStatusShrink
}

func (s *ListTodosShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTodosShrinkRequest) GetSourceTypesShrink() *string {
	return s.SourceTypesShrink
}

func (s *ListTodosShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *ListTodosShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListTodosShrinkRequest) SetCollaborationGroupId(v string) *ListTodosShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListTodosShrinkRequest) SetKeyword(v string) *ListTodosShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListTodosShrinkRequest) SetMaxResults(v int64) *ListTodosShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTodosShrinkRequest) SetNextToken(v string) *ListTodosShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTodosShrinkRequest) SetPage(v int64) *ListTodosShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListTodosShrinkRequest) SetPageSize(v int64) *ListTodosShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTodosShrinkRequest) SetPrioritiesShrink(v string) *ListTodosShrinkRequest {
	s.PrioritiesShrink = &v
	return s
}

func (s *ListTodosShrinkRequest) SetReadStatusShrink(v string) *ListTodosShrinkRequest {
	s.ReadStatusShrink = &v
	return s
}

func (s *ListTodosShrinkRequest) SetSortBy(v string) *ListTodosShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListTodosShrinkRequest) SetSourceTypesShrink(v string) *ListTodosShrinkRequest {
	s.SourceTypesShrink = &v
	return s
}

func (s *ListTodosShrinkRequest) SetStatusesShrink(v string) *ListTodosShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *ListTodosShrinkRequest) SetTenantId(v string) *ListTodosShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListTodosShrinkRequest) Validate() error {
	return dara.Validate(s)
}
