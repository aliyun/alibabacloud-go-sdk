// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *QuerySourceUsageSummaryRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *QuerySourceUsageSummaryRequest
	GetCollaborationGroupId() *string
	SetEndTime(v string) *QuerySourceUsageSummaryRequest
	GetEndTime() *string
	SetEventType(v string) *QuerySourceUsageSummaryRequest
	GetEventType() *string
	SetKnowledgeBaseIds(v []*string) *QuerySourceUsageSummaryRequest
	GetKnowledgeBaseIds() []*string
	SetOpsUserCode(v string) *QuerySourceUsageSummaryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *QuerySourceUsageSummaryRequest
	GetOpsUserId() *string
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
	// The platform application token issued by the console. If specified, the request is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration group to which the resource belongs, such as cg_101. If specified, a group workspace task is created, and the caller must be a valid group member. If left empty, a personal task is created.
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
	// ALL
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// The IDs of the knowledge base folders.
	KnowledgeBaseIds []*string `json:"knowledgeBaseIds,omitempty" xml:"knowledgeBaseIds,omitempty" type:"Repeated"`
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
	// TENANT
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// The source IDs.
	SourceIds []*string `json:"sourceIds,omitempty" xml:"sourceIds,omitempty" type:"Repeated"`
	// The resume channel source.
	//
	// example:
	//
	// Source name content
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-01T00:00:00+08:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The tenant ID. This is a common parameter. For winnexo-cli, this parameter is explicitly passed by using --tenant-id.
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

func (s *QuerySourceUsageSummaryRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *QuerySourceUsageSummaryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *QuerySourceUsageSummaryRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *QuerySourceUsageSummaryRequest) SetAppToken(v string) *QuerySourceUsageSummaryRequest {
	s.AppToken = &v
	return s
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

func (s *QuerySourceUsageSummaryRequest) SetOpsUserCode(v string) *QuerySourceUsageSummaryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *QuerySourceUsageSummaryRequest) SetOpsUserId(v string) *QuerySourceUsageSummaryRequest {
	s.OpsUserId = &v
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
