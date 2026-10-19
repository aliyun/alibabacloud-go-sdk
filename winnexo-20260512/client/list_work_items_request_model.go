// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatuses(v []*string) *ListWorkItemsRequest
	GetBusinessStatuses() []*string
	SetCollaborationGroupId(v string) *ListWorkItemsRequest
	GetCollaborationGroupId() *string
	SetItemTypes(v []*string) *ListWorkItemsRequest
	GetItemTypes() []*string
	SetKeyword(v string) *ListWorkItemsRequest
	GetKeyword() *string
	SetMaxResults(v int64) *ListWorkItemsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListWorkItemsRequest
	GetNextToken() *string
	SetPage(v int64) *ListWorkItemsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListWorkItemsRequest
	GetPageSize() *int64
	SetPriorities(v []*string) *ListWorkItemsRequest
	GetPriorities() []*string
	SetScope(v string) *ListWorkItemsRequest
	GetScope() *string
	SetTenantId(v string) *ListWorkItemsRequest
	GetTenantId() *string
}

type ListWorkItemsRequest struct {
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
	// 单页最大返回数量（1~100）；传入时优先于 pageSize
	//
	// example:
	//
	// 1
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
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListWorkItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkItemsRequest) GetBusinessStatuses() []*string {
	return s.BusinessStatuses
}

func (s *ListWorkItemsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkItemsRequest) GetItemTypes() []*string {
	return s.ItemTypes
}

func (s *ListWorkItemsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListWorkItemsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListWorkItemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkItemsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListWorkItemsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWorkItemsRequest) GetPriorities() []*string {
	return s.Priorities
}

func (s *ListWorkItemsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListWorkItemsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkItemsRequest) SetBusinessStatuses(v []*string) *ListWorkItemsRequest {
	s.BusinessStatuses = v
	return s
}

func (s *ListWorkItemsRequest) SetCollaborationGroupId(v string) *ListWorkItemsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkItemsRequest) SetItemTypes(v []*string) *ListWorkItemsRequest {
	s.ItemTypes = v
	return s
}

func (s *ListWorkItemsRequest) SetKeyword(v string) *ListWorkItemsRequest {
	s.Keyword = &v
	return s
}

func (s *ListWorkItemsRequest) SetMaxResults(v int64) *ListWorkItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkItemsRequest) SetNextToken(v string) *ListWorkItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkItemsRequest) SetPage(v int64) *ListWorkItemsRequest {
	s.Page = &v
	return s
}

func (s *ListWorkItemsRequest) SetPageSize(v int64) *ListWorkItemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkItemsRequest) SetPriorities(v []*string) *ListWorkItemsRequest {
	s.Priorities = v
	return s
}

func (s *ListWorkItemsRequest) SetScope(v string) *ListWorkItemsRequest {
	s.Scope = &v
	return s
}

func (s *ListWorkItemsRequest) SetTenantId(v string) *ListWorkItemsRequest {
	s.TenantId = &v
	return s
}

func (s *ListWorkItemsRequest) Validate() error {
	return dara.Validate(s)
}
