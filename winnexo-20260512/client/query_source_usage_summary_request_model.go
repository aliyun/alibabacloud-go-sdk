// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *QuerySourceUsageSummaryRequest
	GetCollaborationGroupId() *string
	SetEndTime(v string) *QuerySourceUsageSummaryRequest
	GetEndTime() *string
	SetEventType(v string) *QuerySourceUsageSummaryRequest
	GetEventType() *string
	SetKnowledgeBaseIds(v []*string) *QuerySourceUsageSummaryRequest
	GetKnowledgeBaseIds() []*string
	SetScopeType(v string) *QuerySourceUsageSummaryRequest
	GetScopeType() *string
	SetSourceIds(v []*string) *QuerySourceUsageSummaryRequest
	GetSourceIds() []*string
	SetSourceName(v string) *QuerySourceUsageSummaryRequest
	GetSourceName() *string
	SetStartTime(v string) *QuerySourceUsageSummaryRequest
	GetStartTime() *string
	SetTenantId(v string) *QuerySourceUsageSummaryRequest
	GetTenantId() *string
}

type QuerySourceUsageSummaryRequest struct {
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
	// Report perspective: TENANT for enterprise knowledge, GROUP for a single workspace, and ALL_GROUPS for all workspaces of the current tenant (platform super administrator only). Valid values: TENANT, GROUP, and ALL_GROUPS
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// The source ID.
	SourceIds []*string `json:"sourceIds,omitempty" xml:"sourceIds,omitempty" type:"Repeated"`
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

func (s QuerySourceUsageSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySourceUsageSummaryRequest) GoString() string {
	return s.String()
}

func (s *QuerySourceUsageSummaryRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *QuerySourceUsageSummaryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySourceUsageSummaryRequest) GetEventType() *string {
	return s.EventType
}

func (s *QuerySourceUsageSummaryRequest) GetKnowledgeBaseIds() []*string {
	return s.KnowledgeBaseIds
}

func (s *QuerySourceUsageSummaryRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *QuerySourceUsageSummaryRequest) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *QuerySourceUsageSummaryRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *QuerySourceUsageSummaryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySourceUsageSummaryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySourceUsageSummaryRequest) SetCollaborationGroupId(v string) *QuerySourceUsageSummaryRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *QuerySourceUsageSummaryRequest) SetEndTime(v string) *QuerySourceUsageSummaryRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySourceUsageSummaryRequest) SetEventType(v string) *QuerySourceUsageSummaryRequest {
	s.EventType = &v
	return s
}

func (s *QuerySourceUsageSummaryRequest) SetKnowledgeBaseIds(v []*string) *QuerySourceUsageSummaryRequest {
	s.KnowledgeBaseIds = v
	return s
}

func (s *QuerySourceUsageSummaryRequest) SetScopeType(v string) *QuerySourceUsageSummaryRequest {
	s.ScopeType = &v
	return s
}

func (s *QuerySourceUsageSummaryRequest) SetSourceIds(v []*string) *QuerySourceUsageSummaryRequest {
	s.SourceIds = v
	return s
}

func (s *QuerySourceUsageSummaryRequest) SetSourceName(v string) *QuerySourceUsageSummaryRequest {
	s.SourceName = &v
	return s
}

func (s *QuerySourceUsageSummaryRequest) SetStartTime(v string) *QuerySourceUsageSummaryRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySourceUsageSummaryRequest) SetTenantId(v string) *QuerySourceUsageSummaryRequest {
	s.TenantId = &v
	return s
}

func (s *QuerySourceUsageSummaryRequest) Validate() error {
	return dara.Validate(s)
}
