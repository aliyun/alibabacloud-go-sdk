// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourceUsageStatisticsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListSourceUsageStatisticsShrinkRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *ListSourceUsageStatisticsShrinkRequest
	GetCollaborationGroupId() *string
	SetEndTime(v string) *ListSourceUsageStatisticsShrinkRequest
	GetEndTime() *string
	SetEventType(v string) *ListSourceUsageStatisticsShrinkRequest
	GetEventType() *string
	SetKnowledgeBaseIdsShrink(v string) *ListSourceUsageStatisticsShrinkRequest
	GetKnowledgeBaseIdsShrink() *string
	SetOpsUserCode(v string) *ListSourceUsageStatisticsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListSourceUsageStatisticsShrinkRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListSourceUsageStatisticsShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListSourceUsageStatisticsShrinkRequest
	GetPageSize() *int64
	SetScopeType(v string) *ListSourceUsageStatisticsShrinkRequest
	GetScopeType() *string
	SetSortBy(v string) *ListSourceUsageStatisticsShrinkRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListSourceUsageStatisticsShrinkRequest
	GetSortOrder() *string
	SetSourceIdsShrink(v string) *ListSourceUsageStatisticsShrinkRequest
	GetSourceIdsShrink() *string
	SetSourceName(v string) *ListSourceUsageStatisticsShrinkRequest
	GetSourceName() *string
	SetStartTime(v string) *ListSourceUsageStatisticsShrinkRequest
	GetStartTime() *string
	SetTenantId(v string) *ListSourceUsageStatisticsShrinkRequest
	GetTenantId() *string
}

type ListSourceUsageStatisticsShrinkRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration group to which the resource belongs, such as cg_101. If specified, a group workspace task is created, and the caller must be a valid group member. If left empty, a personal task is created.
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
	// ALL
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// The root IDs of enterprise knowledge bases, available only for TENANT. The IDs are filtered based on current directory bindings rather than historical ownership. A maximum of 100 items are supported.
	//
	// example:
	//
	// ["kb_9401"]
	KnowledgeBaseIdsShrink *string `json:"knowledgeBaseIds,omitempty" xml:"knowledgeBaseIds,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
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
	// The report perspective. Valid values: TENANT for enterprise knowledge, GROUP for a single workspace, and ALL_GROUPS for all workspaces of the current tenant (available only to platform super administrators).
	//
	// This parameter is required.
	//
	// example:
	//
	// TENANT
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// The content of the sort clause.
	//
	// example:
	//
	// retrievalCount
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// Specifies the sort order. This parameter takes effect only when sortBy is specified. Valid values: ASC and DESC. The values are case-insensitive.
	//
	// example:
	//
	// ASC
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The associated source IDs.
	//
	// example:
	//
	// ["src_9401"]
	SourceIdsShrink *string `json:"sourceIds,omitempty" xml:"sourceIds,omitempty"`
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

func (s ListSourceUsageStatisticsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSourceUsageStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetEventType() *string {
	return s.EventType
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetKnowledgeBaseIdsShrink() *string {
	return s.KnowledgeBaseIdsShrink
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetSourceIdsShrink() *string {
	return s.SourceIdsShrink
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSourceUsageStatisticsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetAppToken(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetCollaborationGroupId(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetEndTime(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetEventType(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.EventType = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetKnowledgeBaseIdsShrink(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.KnowledgeBaseIdsShrink = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetOpsUserCode(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetOpsUserId(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetPage(v int64) *ListSourceUsageStatisticsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetPageSize(v int64) *ListSourceUsageStatisticsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetScopeType(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.ScopeType = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetSortBy(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetSortOrder(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetSourceIdsShrink(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.SourceIdsShrink = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetSourceName(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.SourceName = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetStartTime(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) SetTenantId(v string) *ListSourceUsageStatisticsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListSourceUsageStatisticsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
