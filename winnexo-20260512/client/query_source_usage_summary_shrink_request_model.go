// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageSummaryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *QuerySourceUsageSummaryShrinkRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *QuerySourceUsageSummaryShrinkRequest
	GetCollaborationGroupId() *string
	SetEndTime(v string) *QuerySourceUsageSummaryShrinkRequest
	GetEndTime() *string
	SetEventType(v string) *QuerySourceUsageSummaryShrinkRequest
	GetEventType() *string
	SetKnowledgeBaseIdsShrink(v string) *QuerySourceUsageSummaryShrinkRequest
	GetKnowledgeBaseIdsShrink() *string
	SetOpsUserCode(v string) *QuerySourceUsageSummaryShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *QuerySourceUsageSummaryShrinkRequest
	GetOpsUserId() *string
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
	SourceIdsShrink *string `json:"sourceIds,omitempty" xml:"sourceIds,omitempty"`
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

func (s QuerySourceUsageSummaryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySourceUsageSummaryShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySourceUsageSummaryShrinkRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *QuerySourceUsageSummaryShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *QuerySourceUsageSummaryShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *QuerySourceUsageSummaryShrinkRequest) SetAppToken(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.AppToken = &v
	return s
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

func (s *QuerySourceUsageSummaryShrinkRequest) SetOpsUserCode(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *QuerySourceUsageSummaryShrinkRequest) SetOpsUserId(v string) *QuerySourceUsageSummaryShrinkRequest {
	s.OpsUserId = &v
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
