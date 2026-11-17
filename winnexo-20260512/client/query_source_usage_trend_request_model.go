// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *QuerySourceUsageTrendRequest
	GetCollaborationGroupId() *string
	SetEndTime(v string) *QuerySourceUsageTrendRequest
	GetEndTime() *string
	SetEventType(v string) *QuerySourceUsageTrendRequest
	GetEventType() *string
	SetGranularity(v string) *QuerySourceUsageTrendRequest
	GetGranularity() *string
	SetKnowledgeBaseIds(v []*string) *QuerySourceUsageTrendRequest
	GetKnowledgeBaseIds() []*string
	SetScopeType(v string) *QuerySourceUsageTrendRequest
	GetScopeType() *string
	SetSourceIds(v []*string) *QuerySourceUsageTrendRequest
	GetSourceIds() []*string
	SetSourceName(v string) *QuerySourceUsageTrendRequest
	GetSourceName() *string
	SetStartTime(v string) *QuerySourceUsageTrendRequest
	GetStartTime() *string
	SetTenantId(v string) *QuerySourceUsageTrendRequest
	GetTenantId() *string
	SetTimeZone(v string) *QuerySourceUsageTrendRequest
	GetTimeZone() *string
}

type QuerySourceUsageTrendRequest struct {
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
	KnowledgeBaseIds []*string `json:"knowledgeBaseIds,omitempty" xml:"knowledgeBaseIds,omitempty" type:"Repeated"`
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
	SourceIds []*string `json:"sourceIds,omitempty" xml:"sourceIds,omitempty" type:"Repeated"`
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

func (s QuerySourceUsageTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySourceUsageTrendRequest) GoString() string {
	return s.String()
}

func (s *QuerySourceUsageTrendRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *QuerySourceUsageTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySourceUsageTrendRequest) GetEventType() *string {
	return s.EventType
}

func (s *QuerySourceUsageTrendRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *QuerySourceUsageTrendRequest) GetKnowledgeBaseIds() []*string {
	return s.KnowledgeBaseIds
}

func (s *QuerySourceUsageTrendRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *QuerySourceUsageTrendRequest) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *QuerySourceUsageTrendRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *QuerySourceUsageTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySourceUsageTrendRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySourceUsageTrendRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *QuerySourceUsageTrendRequest) SetCollaborationGroupId(v string) *QuerySourceUsageTrendRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *QuerySourceUsageTrendRequest) SetEndTime(v string) *QuerySourceUsageTrendRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySourceUsageTrendRequest) SetEventType(v string) *QuerySourceUsageTrendRequest {
	s.EventType = &v
	return s
}

func (s *QuerySourceUsageTrendRequest) SetGranularity(v string) *QuerySourceUsageTrendRequest {
	s.Granularity = &v
	return s
}

func (s *QuerySourceUsageTrendRequest) SetKnowledgeBaseIds(v []*string) *QuerySourceUsageTrendRequest {
	s.KnowledgeBaseIds = v
	return s
}

func (s *QuerySourceUsageTrendRequest) SetScopeType(v string) *QuerySourceUsageTrendRequest {
	s.ScopeType = &v
	return s
}

func (s *QuerySourceUsageTrendRequest) SetSourceIds(v []*string) *QuerySourceUsageTrendRequest {
	s.SourceIds = v
	return s
}

func (s *QuerySourceUsageTrendRequest) SetSourceName(v string) *QuerySourceUsageTrendRequest {
	s.SourceName = &v
	return s
}

func (s *QuerySourceUsageTrendRequest) SetStartTime(v string) *QuerySourceUsageTrendRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySourceUsageTrendRequest) SetTenantId(v string) *QuerySourceUsageTrendRequest {
	s.TenantId = &v
	return s
}

func (s *QuerySourceUsageTrendRequest) SetTimeZone(v string) *QuerySourceUsageTrendRequest {
	s.TimeZone = &v
	return s
}

func (s *QuerySourceUsageTrendRequest) Validate() error {
	return dara.Validate(s)
}
