// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourceUsageStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListSourceUsageStatisticsRequest
	GetCollaborationGroupId() *string
	SetEndTime(v string) *ListSourceUsageStatisticsRequest
	GetEndTime() *string
	SetEventType(v string) *ListSourceUsageStatisticsRequest
	GetEventType() *string
	SetKnowledgeBaseIds(v []*string) *ListSourceUsageStatisticsRequest
	GetKnowledgeBaseIds() []*string
	SetPage(v int64) *ListSourceUsageStatisticsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListSourceUsageStatisticsRequest
	GetPageSize() *int64
	SetScopeType(v string) *ListSourceUsageStatisticsRequest
	GetScopeType() *string
	SetSortBy(v string) *ListSourceUsageStatisticsRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListSourceUsageStatisticsRequest
	GetSortOrder() *string
	SetSourceIds(v []*string) *ListSourceUsageStatisticsRequest
	GetSourceIds() []*string
	SetSourceName(v string) *ListSourceUsageStatisticsRequest
	GetSourceName() *string
	SetStartTime(v string) *ListSourceUsageStatisticsRequest
	GetStartTime() *string
	SetTenantId(v string) *ListSourceUsageStatisticsRequest
	GetTenantId() *string
}

type ListSourceUsageStatisticsRequest struct {
	// The ID of the collaboration group to which the resource belongs, such as cg_101. If you specify this parameter, a group workspace task is created, and the caller must be a valid group member. If you leave this parameter empty, a personal task is created.
	//
	// example:
	//
	// 1112
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The actual end timestamp, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-08T23:59:59+08:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The event type.
	//
	// example:
	//
	// HealthCheck
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// The root IDs of the enterprise knowledge bases, which are available only for TENANT. The knowledge bases are filtered based on current directory bindings, not historical ownership. Maximum number of items: 100.
	//
	// example:
	//
	// ["kb_9401"]
	KnowledgeBaseIds []*string `json:"knowledgeBaseIds,omitempty" xml:"knowledgeBaseIds,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The report perspective. Valid values: TENANT (enterprise knowledge), GROUP (a single workspace), and ALL_GROUPS (all workspaces of the current tenant, only for platform super administrators).
	//
	// This parameter is required.
	//
	// example:
	//
	// Report perspective: TENANT for enterprise knowledge, GROUP for a single workspace, ALL_GROUPS for all workspaces of the current tenant (platform super administrator only). Valid values: TENANT, GROUP, ALL_GROUPS
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// The content of the sort clause.
	//
	// example:
	//
	// 1
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The sort order. This parameter takes effect only when sortBy is specified. Valid values: ASC and DESC. The values are case-insensitive.
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The IDs of the associated sources.
	//
	// example:
	//
	// ["src_9401"]
	SourceIds []*string `json:"sourceIds,omitempty" xml:"sourceIds,omitempty" type:"Repeated"`
	// The name of the source.
	//
	// example:
	//
	// The naming content of the source
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListSourceUsageStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSourceUsageStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListSourceUsageStatisticsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListSourceUsageStatisticsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSourceUsageStatisticsRequest) GetEventType() *string {
	return s.EventType
}

func (s *ListSourceUsageStatisticsRequest) GetKnowledgeBaseIds() []*string {
	return s.KnowledgeBaseIds
}

func (s *ListSourceUsageStatisticsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListSourceUsageStatisticsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSourceUsageStatisticsRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *ListSourceUsageStatisticsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSourceUsageStatisticsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListSourceUsageStatisticsRequest) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *ListSourceUsageStatisticsRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *ListSourceUsageStatisticsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSourceUsageStatisticsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListSourceUsageStatisticsRequest) SetCollaborationGroupId(v string) *ListSourceUsageStatisticsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetEndTime(v string) *ListSourceUsageStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetEventType(v string) *ListSourceUsageStatisticsRequest {
	s.EventType = &v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetKnowledgeBaseIds(v []*string) *ListSourceUsageStatisticsRequest {
	s.KnowledgeBaseIds = v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetPage(v int64) *ListSourceUsageStatisticsRequest {
	s.Page = &v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetPageSize(v int64) *ListSourceUsageStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetScopeType(v string) *ListSourceUsageStatisticsRequest {
	s.ScopeType = &v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetSortBy(v string) *ListSourceUsageStatisticsRequest {
	s.SortBy = &v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetSortOrder(v string) *ListSourceUsageStatisticsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetSourceIds(v []*string) *ListSourceUsageStatisticsRequest {
	s.SourceIds = v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetSourceName(v string) *ListSourceUsageStatisticsRequest {
	s.SourceName = &v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetStartTime(v string) *ListSourceUsageStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *ListSourceUsageStatisticsRequest) SetTenantId(v string) *ListSourceUsageStatisticsRequest {
	s.TenantId = &v
	return s
}

func (s *ListSourceUsageStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
