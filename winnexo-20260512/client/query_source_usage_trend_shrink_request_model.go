// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageTrendShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *QuerySourceUsageTrendShrinkRequest
	GetAppToken() *string
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
	SetOpsUserCode(v string) *QuerySourceUsageTrendShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *QuerySourceUsageTrendShrinkRequest
	GetOpsUserId() *string
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
	KnowledgeBaseIdsShrink *string `json:"knowledgeBaseIds,omitempty" xml:"knowledgeBaseIds,omitempty"`
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
	SourceIdsShrink *string `json:"sourceIds,omitempty" xml:"sourceIds,omitempty"`
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

func (s QuerySourceUsageTrendShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySourceUsageTrendShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySourceUsageTrendShrinkRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *QuerySourceUsageTrendShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *QuerySourceUsageTrendShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *QuerySourceUsageTrendShrinkRequest) SetAppToken(v string) *QuerySourceUsageTrendShrinkRequest {
	s.AppToken = &v
	return s
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

func (s *QuerySourceUsageTrendShrinkRequest) SetOpsUserCode(v string) *QuerySourceUsageTrendShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *QuerySourceUsageTrendShrinkRequest) SetOpsUserId(v string) *QuerySourceUsageTrendShrinkRequest {
	s.OpsUserId = &v
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
