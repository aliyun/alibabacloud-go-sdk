// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *QuerySourceUsageTrendRequest
	GetAppToken() *string
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
	SetOpsUserCode(v string) *QuerySourceUsageTrendRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *QuerySourceUsageTrendRequest
	GetOpsUserId() *string
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
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the associated collaboration group, such as cg_101. If you specify this parameter, a group workspace task is created, and the caller must be a valid group member. If you leave this parameter empty, a personal task is created.
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
	// ALL
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// The trend aggregation granularity. Valid values: DAY (daily), WEEK (weekly), and MONTH (monthly).
	//
	// example:
	//
	// DAY
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The IDs of the knowledge base folders.
	//
	// example:
	//
	// ["kb_9401"]
	KnowledgeBaseIds []*string `json:"knowledgeBaseIds,omitempty" xml:"knowledgeBaseIds,omitempty" type:"Repeated"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
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
	// Naming content of the source
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

func (s *QuerySourceUsageTrendRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *QuerySourceUsageTrendRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *QuerySourceUsageTrendRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *QuerySourceUsageTrendRequest) SetAppToken(v string) *QuerySourceUsageTrendRequest {
	s.AppToken = &v
	return s
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

func (s *QuerySourceUsageTrendRequest) SetOpsUserCode(v string) *QuerySourceUsageTrendRequest {
	s.OpsUserCode = &v
	return s
}

func (s *QuerySourceUsageTrendRequest) SetOpsUserId(v string) *QuerySourceUsageTrendRequest {
	s.OpsUserId = &v
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
