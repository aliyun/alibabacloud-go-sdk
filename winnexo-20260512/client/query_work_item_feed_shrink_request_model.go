// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkItemFeedShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatusesShrink(v string) *QueryWorkItemFeedShrinkRequest
	GetBusinessStatusesShrink() *string
	SetCollaborationGroupId(v string) *QueryWorkItemFeedShrinkRequest
	GetCollaborationGroupId() *string
	SetCursor(v string) *QueryWorkItemFeedShrinkRequest
	GetCursor() *string
	SetItemTypesShrink(v string) *QueryWorkItemFeedShrinkRequest
	GetItemTypesShrink() *string
	SetKeyword(v string) *QueryWorkItemFeedShrinkRequest
	GetKeyword() *string
	SetPageSize(v int64) *QueryWorkItemFeedShrinkRequest
	GetPageSize() *int64
	SetPrioritiesShrink(v string) *QueryWorkItemFeedShrinkRequest
	GetPrioritiesShrink() *string
	SetScope(v string) *QueryWorkItemFeedShrinkRequest
	GetScope() *string
	SetSourceTaskId(v string) *QueryWorkItemFeedShrinkRequest
	GetSourceTaskId() *string
	SetTenantId(v string) *QueryWorkItemFeedShrinkRequest
	GetTenantId() *string
	SetView(v string) *QueryWorkItemFeedShrinkRequest
	GetView() *string
}

type QueryWorkItemFeedShrinkRequest struct {
	// 业务状态筛选（七态）
	//
	// example:
	//
	// example
	BusinessStatusesShrink *string `json:"businessStatuses,omitempty" xml:"businessStatuses,omitempty"`
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 翻页游标，取上次响应的 nextCursor
	//
	// example:
	//
	// example
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// 工作项类型筛选
	//
	// example:
	//
	// example
	ItemTypesShrink *string `json:"itemTypes,omitempty" xml:"itemTypes,omitempty"`
	// 标题/描述模糊搜索关键词
	//
	// example:
	//
	// example
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 单页条数（1~100）
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 优先级筛选
	//
	// example:
	//
	// example
	PrioritiesShrink *string `json:"priorities,omitempty" xml:"priorities,omitempty"`
	// 个人范围筛选
	//
	// example:
	//
	// ASSIGNED_TO_ME
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 来源任务ID；指定时返回该任务平铺子项
	//
	// example:
	//
	// example
	SourceTaskId *string `json:"sourceTaskId,omitempty" xml:"sourceTaskId,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 视图：FLAT 平铺 / TREE 任务分组
	//
	// example:
	//
	// FLAT
	View *string `json:"view,omitempty" xml:"view,omitempty"`
}

func (s QueryWorkItemFeedShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkItemFeedShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryWorkItemFeedShrinkRequest) GetBusinessStatusesShrink() *string {
	return s.BusinessStatusesShrink
}

func (s *QueryWorkItemFeedShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *QueryWorkItemFeedShrinkRequest) GetCursor() *string {
	return s.Cursor
}

func (s *QueryWorkItemFeedShrinkRequest) GetItemTypesShrink() *string {
	return s.ItemTypesShrink
}

func (s *QueryWorkItemFeedShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryWorkItemFeedShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryWorkItemFeedShrinkRequest) GetPrioritiesShrink() *string {
	return s.PrioritiesShrink
}

func (s *QueryWorkItemFeedShrinkRequest) GetScope() *string {
	return s.Scope
}

func (s *QueryWorkItemFeedShrinkRequest) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *QueryWorkItemFeedShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryWorkItemFeedShrinkRequest) GetView() *string {
	return s.View
}

func (s *QueryWorkItemFeedShrinkRequest) SetBusinessStatusesShrink(v string) *QueryWorkItemFeedShrinkRequest {
	s.BusinessStatusesShrink = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetCollaborationGroupId(v string) *QueryWorkItemFeedShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetCursor(v string) *QueryWorkItemFeedShrinkRequest {
	s.Cursor = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetItemTypesShrink(v string) *QueryWorkItemFeedShrinkRequest {
	s.ItemTypesShrink = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetKeyword(v string) *QueryWorkItemFeedShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetPageSize(v int64) *QueryWorkItemFeedShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetPrioritiesShrink(v string) *QueryWorkItemFeedShrinkRequest {
	s.PrioritiesShrink = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetScope(v string) *QueryWorkItemFeedShrinkRequest {
	s.Scope = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetSourceTaskId(v string) *QueryWorkItemFeedShrinkRequest {
	s.SourceTaskId = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetTenantId(v string) *QueryWorkItemFeedShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetView(v string) *QueryWorkItemFeedShrinkRequest {
	s.View = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) Validate() error {
	return dara.Validate(s)
}
