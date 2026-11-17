// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageSummaryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *QuerySourceUsageSummaryShrinkRequest
	GetCollaborationGroupId() *string
	SetEndTime(v string) *QuerySourceUsageSummaryShrinkRequest
	GetEndTime() *string
	SetEventType(v string) *QuerySourceUsageSummaryShrinkRequest
	GetEventType() *string
	SetKnowledgeBaseIdsShrink(v string) *QuerySourceUsageSummaryShrinkRequest
	GetKnowledgeBaseIdsShrink() *string
	SetScopeType(v string) *QuerySourceUsageSummaryShrinkRequest
	GetScopeType() *string
	SetSourceIdsShrink(v string) *QuerySourceUsageSummaryShrinkRequest
	GetSourceIdsShrink() *string
	SetSourceName(v string) *QuerySourceUsageSummaryShrinkRequest
	GetSourceName() *string
	SetStartTime(v string) *QuerySourceUsageSummaryShrinkRequest
	GetStartTime() *string
	SetTenantId(v string) *QuerySourceUsageSummaryShrinkRequest
	GetTenantId() *string
}

type QuerySourceUsageSummaryShrinkRequest struct {
	// The ID of the collaboration group to which the resource belongs, such as cg_101. If you specify this parameter, a group workspace task is created, and the caller must be a valid group member. If you leave this parameter empty, a personal task is created.
	//
	// example:
	//
	// exampleCollaborationGroupId
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
	// UserOperator
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// The ID of the knowledge base folder.
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
	// Report perspective: TENANT for enterprise knowledge, GROUP for a single workspace, and ALL_GROUPS for all workspaces of the current tenant (platform super administrator only). Valid values: TENANT, GROUP, and ALL_GROUPS
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// The source ID.
	SourceIdsShrink *string `json:"sourceIds,omitempty" xml:"sourceIds,omitempty"`
	// The resume channel source.
	//
	// example:
	//
	// The name of the source
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-01T00:00:00+08:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The tenant ID, which is a common parameter. If you use winnexo-cli, pass this parameter explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QuerySourceUsageSummaryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySourceUsageSummaryShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySourceUsageSummaryShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *QuerySourceUsageSummaryShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySourceUsageSummaryShrinkRequest) GetEventType() *string {
	return s.EventType
}

func (s *QuerySourceUsageSummaryShrinkRequest) GetKnowledgeBaseIdsShrink() *string {
	return s.KnowledgeBaseIdsShrink
}

func (s *QuerySourceUsageSummaryShrinkRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *QuerySourceUsageSummaryShrinkRequest) GetSourceIdsShrink() *string {
	return s.SourceIdsShrink
}

func (s *QuerySourceUsageSummaryShrinkRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *QuerySourceUsageSummaryShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySourceUsageSummaryShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySourceUsageSummaryShrinkRequest) SetCollaborationGroupId(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *QuerySourceUsageSummaryShrinkRequest) SetEndTime(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySourceUsageSummaryShrinkRequest) SetEventType(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.EventType = &v
	return s
}

func (s *QuerySourceUsageSummaryShrinkRequest) SetKnowledgeBaseIdsShrink(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.KnowledgeBaseIdsShrink = &v
	return s
}

func (s *QuerySourceUsageSummaryShrinkRequest) SetScopeType(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.ScopeType = &v
	return s
}

func (s *QuerySourceUsageSummaryShrinkRequest) SetSourceIdsShrink(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.SourceIdsShrink = &v
	return s
}

func (s *QuerySourceUsageSummaryShrinkRequest) SetSourceName(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.SourceName = &v
	return s
}

func (s *QuerySourceUsageSummaryShrinkRequest) SetStartTime(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySourceUsageSummaryShrinkRequest) SetTenantId(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *QuerySourceUsageSummaryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
