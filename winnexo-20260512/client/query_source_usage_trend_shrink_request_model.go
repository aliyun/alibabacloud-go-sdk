// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageTrendShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *QuerySourceUsageTrendShrinkRequest
	GetCollaborationGroupId() *string
	SetEndTime(v string) *QuerySourceUsageTrendShrinkRequest
	GetEndTime() *string
	SetEventType(v string) *QuerySourceUsageTrendShrinkRequest
	GetEventType() *string
	SetGranularity(v string) *QuerySourceUsageTrendShrinkRequest
	GetGranularity() *string
	SetKnowledgeBaseIdsShrink(v string) *QuerySourceUsageTrendShrinkRequest
	GetKnowledgeBaseIdsShrink() *string
	SetScopeType(v string) *QuerySourceUsageTrendShrinkRequest
	GetScopeType() *string
	SetSourceIdsShrink(v string) *QuerySourceUsageTrendShrinkRequest
	GetSourceIdsShrink() *string
	SetSourceName(v string) *QuerySourceUsageTrendShrinkRequest
	GetSourceName() *string
	SetStartTime(v string) *QuerySourceUsageTrendShrinkRequest
	GetStartTime() *string
	SetTenantId(v string) *QuerySourceUsageTrendShrinkRequest
	GetTenantId() *string
	SetTimeZone(v string) *QuerySourceUsageTrendShrinkRequest
	GetTimeZone() *string
}

type QuerySourceUsageTrendShrinkRequest struct {
	// The ID of the collaboration group to which the resource belongs, such as cg_101. If this parameter is specified, a group workspace task is created, and the caller must be a valid group member. If this parameter is left empty, a personal task is created.
	//
	// example:
	//
	// example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The actual end timestamp of the live stream, in milliseconds.
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
	// MAIN
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// The aggregation granularity. Valid values: hourly and daily.
	//
	// example:
	//
	// daily
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The IDs of the knowledge base folders.
	//
	// example:
	//
	// ["kb_9401"]
	KnowledgeBaseIdsShrink *string `json:"knowledgeBaseIds,omitempty" xml:"knowledgeBaseIds,omitempty"`
	// The role scope type. Valid values:
	//
	// - **ALL_VISIBLE**: All members.
	//
	// - **PART_VISIBLE**: Partial members.
	//
	// This parameter is required.
	//
	// example:
	//
	// Report perspective: TENANT for enterprise knowledge, GROUP for a single workspace, ALL_GROUPS for all workspaces of the current tenant (platform super administrator only). Valid values: TENANT, GROUP, and ALL_GROUPS
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// The IDs of the resource sources.
	//
	// example:
	//
	// ["src_9401"]
	SourceIdsShrink *string `json:"sourceIds,omitempty" xml:"sourceIds,omitempty"`
	// The resume channel source.
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
	// 2026-09-01T00:00:00+08:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s QuerySourceUsageTrendShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySourceUsageTrendShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySourceUsageTrendShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *QuerySourceUsageTrendShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySourceUsageTrendShrinkRequest) GetEventType() *string {
	return s.EventType
}

func (s *QuerySourceUsageTrendShrinkRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *QuerySourceUsageTrendShrinkRequest) GetKnowledgeBaseIdsShrink() *string {
	return s.KnowledgeBaseIdsShrink
}

func (s *QuerySourceUsageTrendShrinkRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *QuerySourceUsageTrendShrinkRequest) GetSourceIdsShrink() *string {
	return s.SourceIdsShrink
}

func (s *QuerySourceUsageTrendShrinkRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *QuerySourceUsageTrendShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySourceUsageTrendShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySourceUsageTrendShrinkRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *QuerySourceUsageTrendShrinkRequest) SetCollaborationGroupId(v string) *QuerySourceUsageTrendShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) SetEndTime(v string) *QuerySourceUsageTrendShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) SetEventType(v string) *QuerySourceUsageTrendShrinkRequest {
	s.EventType = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) SetGranularity(v string) *QuerySourceUsageTrendShrinkRequest {
	s.Granularity = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) SetKnowledgeBaseIdsShrink(v string) *QuerySourceUsageTrendShrinkRequest {
	s.KnowledgeBaseIdsShrink = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) SetScopeType(v string) *QuerySourceUsageTrendShrinkRequest {
	s.ScopeType = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) SetSourceIdsShrink(v string) *QuerySourceUsageTrendShrinkRequest {
	s.SourceIdsShrink = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) SetSourceName(v string) *QuerySourceUsageTrendShrinkRequest {
	s.SourceName = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) SetStartTime(v string) *QuerySourceUsageTrendShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) SetTenantId(v string) *QuerySourceUsageTrendShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) SetTimeZone(v string) *QuerySourceUsageTrendShrinkRequest {
	s.TimeZone = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) Validate() error {
	return dara.Validate(s)
}
