// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnnouncementShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateAnnouncementShrinkRequest
	GetAppToken() *string
	SetContent(v string) *CreateAnnouncementShrinkRequest
	GetContent() *string
	SetDisplayPage(v string) *CreateAnnouncementShrinkRequest
	GetDisplayPage() *string
	SetDisplayType(v string) *CreateAnnouncementShrinkRequest
	GetDisplayType() *string
	SetEffectiveEnd(v string) *CreateAnnouncementShrinkRequest
	GetEffectiveEnd() *string
	SetEffectiveStart(v string) *CreateAnnouncementShrinkRequest
	GetEffectiveStart() *string
	SetOpsUserCode(v string) *CreateAnnouncementShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateAnnouncementShrinkRequest
	GetOpsUserId() *string
	SetPriority(v string) *CreateAnnouncementShrinkRequest
	GetPriority() *string
	SetPublishNow(v bool) *CreateAnnouncementShrinkRequest
	GetPublishNow() *bool
	SetTargetRoleCodesShrink(v string) *CreateAnnouncementShrinkRequest
	GetTargetRoleCodesShrink() *string
	SetTargetRoleMode(v string) *CreateAnnouncementShrinkRequest
	GetTargetRoleMode() *string
	SetTargetTenantIdsShrink(v string) *CreateAnnouncementShrinkRequest
	GetTargetTenantIdsShrink() *string
	SetTargetTenantMode(v string) *CreateAnnouncementShrinkRequest
	GetTargetTenantMode() *string
	SetTenantId(v string) *CreateAnnouncementShrinkRequest
	GetTenantId() *string
	SetTitle(v string) *CreateAnnouncementShrinkRequest
	GetTitle() *string
}

type CreateAnnouncementShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The content of the notice.
	//
	// This parameter is required.
	//
	// example:
	//
	// The system will undergo maintenance tonight
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The display page. Valid values: ALL, FRONTEND, and BACKEND.
	//
	// example:
	//
	// ALL
	DisplayPage *string `json:"displayPage,omitempty" xml:"displayPage,omitempty"`
	// The display type and group label.
	//
	// example:
	//
	// LIST
	DisplayType *string `json:"displayType,omitempty" xml:"displayType,omitempty"`
	// The effective end time.
	//
	// example:
	//
	// 2026-08-21T14:00:00+08:00
	EffectiveEnd *string `json:"effectiveEnd,omitempty" xml:"effectiveEnd,omitempty"`
	// The effective start time in ISO 8601 format with time zone. If this parameter is not specified, the notice takes effect immediately.
	//
	// example:
	//
	// 2026-08-20T14:00:00+08:00
	EffectiveStart *string `json:"effectiveStart,omitempty" xml:"effectiveStart,omitempty"`
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
	// The priority. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// GENERAL
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// Specifies whether to publish the notice immediately after creation.
	//
	// example:
	//
	// false
	PublishNow *bool `json:"publishNow,omitempty" xml:"publishNow,omitempty"`
	// The list of system role codes. This parameter is used when targetRoleMode is set to SPECIFIED.
	//
	// example:
	//
	// ["SYSTEM_ADMIN"]
	TargetRoleCodesShrink *string `json:"targetRoleCodes,omitempty" xml:"targetRoleCodes,omitempty"`
	// The target role scope. Valid values: ALL and SPECIFIED.
	//
	// example:
	//
	// ALL
	TargetRoleMode *string `json:"targetRoleMode,omitempty" xml:"targetRoleMode,omitempty"`
	// The list of target tenant IDs. This parameter is used when targetTenantMode is set to SPECIFIED.
	//
	// example:
	//
	// [10000]
	TargetTenantIdsShrink *string `json:"targetTenantIds,omitempty" xml:"targetTenantIds,omitempty"`
	// The target tenant scope. Valid values: ALL and SPECIFIED.
	//
	// example:
	//
	// ALL
	TargetTenantMode *string `json:"targetTenantMode,omitempty" xml:"targetTenantMode,omitempty"`
	// The ID of the tenant for which the notice takes effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The title of the notice.
	//
	// This parameter is required.
	//
	// example:
	//
	// System Maintenance Notification
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateAnnouncementShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnouncementShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAnnouncementShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateAnnouncementShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *CreateAnnouncementShrinkRequest) GetDisplayPage() *string {
	return s.DisplayPage
}

func (s *CreateAnnouncementShrinkRequest) GetDisplayType() *string {
	return s.DisplayType
}

func (s *CreateAnnouncementShrinkRequest) GetEffectiveEnd() *string {
	return s.EffectiveEnd
}

func (s *CreateAnnouncementShrinkRequest) GetEffectiveStart() *string {
	return s.EffectiveStart
}

func (s *CreateAnnouncementShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateAnnouncementShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateAnnouncementShrinkRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateAnnouncementShrinkRequest) GetPublishNow() *bool {
	return s.PublishNow
}

func (s *CreateAnnouncementShrinkRequest) GetTargetRoleCodesShrink() *string {
	return s.TargetRoleCodesShrink
}

func (s *CreateAnnouncementShrinkRequest) GetTargetRoleMode() *string {
	return s.TargetRoleMode
}

func (s *CreateAnnouncementShrinkRequest) GetTargetTenantIdsShrink() *string {
	return s.TargetTenantIdsShrink
}

func (s *CreateAnnouncementShrinkRequest) GetTargetTenantMode() *string {
	return s.TargetTenantMode
}

func (s *CreateAnnouncementShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateAnnouncementShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateAnnouncementShrinkRequest) SetAppToken(v string) *CreateAnnouncementShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetContent(v string) *CreateAnnouncementShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetDisplayPage(v string) *CreateAnnouncementShrinkRequest {
	s.DisplayPage = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetDisplayType(v string) *CreateAnnouncementShrinkRequest {
	s.DisplayType = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetEffectiveEnd(v string) *CreateAnnouncementShrinkRequest {
	s.EffectiveEnd = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetEffectiveStart(v string) *CreateAnnouncementShrinkRequest {
	s.EffectiveStart = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetOpsUserCode(v string) *CreateAnnouncementShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetOpsUserId(v string) *CreateAnnouncementShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetPriority(v string) *CreateAnnouncementShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetPublishNow(v bool) *CreateAnnouncementShrinkRequest {
	s.PublishNow = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetTargetRoleCodesShrink(v string) *CreateAnnouncementShrinkRequest {
	s.TargetRoleCodesShrink = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetTargetRoleMode(v string) *CreateAnnouncementShrinkRequest {
	s.TargetRoleMode = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetTargetTenantIdsShrink(v string) *CreateAnnouncementShrinkRequest {
	s.TargetTenantIdsShrink = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetTargetTenantMode(v string) *CreateAnnouncementShrinkRequest {
	s.TargetTenantMode = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetTenantId(v string) *CreateAnnouncementShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) SetTitle(v string) *CreateAnnouncementShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateAnnouncementShrinkRequest) Validate() error {
	return dara.Validate(s)
}
