// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkItemFeedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatuses(v []*string) *QueryWorkItemFeedRequest
	GetBusinessStatuses() []*string
	SetCollaborationGroupId(v string) *QueryWorkItemFeedRequest
	GetCollaborationGroupId() *string
	SetCursor(v string) *QueryWorkItemFeedRequest
	GetCursor() *string
	SetItemTypes(v []*string) *QueryWorkItemFeedRequest
	GetItemTypes() []*string
	SetKeyword(v string) *QueryWorkItemFeedRequest
	GetKeyword() *string
	SetPageSize(v int64) *QueryWorkItemFeedRequest
	GetPageSize() *int64
	SetPriorities(v []*string) *QueryWorkItemFeedRequest
	GetPriorities() []*string
	SetScope(v string) *QueryWorkItemFeedRequest
	GetScope() *string
	SetSourceTaskId(v string) *QueryWorkItemFeedRequest
	GetSourceTaskId() *string
	SetTenantId(v string) *QueryWorkItemFeedRequest
	GetTenantId() *string
	SetView(v string) *QueryWorkItemFeedRequest
	GetView() *string
}

type QueryWorkItemFeedRequest struct {
	// 业务状态筛选（七态）
	//
	// example:
	//
	// example
	BusinessStatuses []*string `json:"businessStatuses,omitempty" xml:"businessStatuses,omitempty" type:"Repeated"`
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
	ItemTypes []*string `json:"itemTypes,omitempty" xml:"itemTypes,omitempty" type:"Repeated"`
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
	Priorities []*string `json:"priorities,omitempty" xml:"priorities,omitempty" type:"Repeated"`
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

func (s QueryWorkItemFeedRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkItemFeedRequest) GoString() string {
	return s.String()
}

func (s *QueryWorkItemFeedRequest) GetBusinessStatuses() []*string {
	return s.BusinessStatuses
}

func (s *QueryWorkItemFeedRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *QueryWorkItemFeedRequest) GetCursor() *string {
	return s.Cursor
}

func (s *QueryWorkItemFeedRequest) GetItemTypes() []*string {
	return s.ItemTypes
}

func (s *QueryWorkItemFeedRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryWorkItemFeedRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryWorkItemFeedRequest) GetPriorities() []*string {
	return s.Priorities
}

func (s *QueryWorkItemFeedRequest) GetScope() *string {
	return s.Scope
}

func (s *QueryWorkItemFeedRequest) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *QueryWorkItemFeedRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryWorkItemFeedRequest) GetView() *string {
	return s.View
}

func (s *QueryWorkItemFeedRequest) SetBusinessStatuses(v []*string) *QueryWorkItemFeedRequest {
	s.BusinessStatuses = v
	return s
}

func (s *QueryWorkItemFeedRequest) SetCollaborationGroupId(v string) *QueryWorkItemFeedRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetCursor(v string) *QueryWorkItemFeedRequest {
	s.Cursor = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetItemTypes(v []*string) *QueryWorkItemFeedRequest {
	s.ItemTypes = v
	return s
}

func (s *QueryWorkItemFeedRequest) SetKeyword(v string) *QueryWorkItemFeedRequest {
	s.Keyword = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetPageSize(v int64) *QueryWorkItemFeedRequest {
	s.PageSize = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetPriorities(v []*string) *QueryWorkItemFeedRequest {
	s.Priorities = v
	return s
}

func (s *QueryWorkItemFeedRequest) SetScope(v string) *QueryWorkItemFeedRequest {
	s.Scope = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetSourceTaskId(v string) *QueryWorkItemFeedRequest {
	s.SourceTaskId = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetTenantId(v string) *QueryWorkItemFeedRequest {
	s.TenantId = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetView(v string) *QueryWorkItemFeedRequest {
	s.View = &v
	return s
}

func (s *QueryWorkItemFeedRequest) Validate() error {
	return dara.Validate(s)
}
