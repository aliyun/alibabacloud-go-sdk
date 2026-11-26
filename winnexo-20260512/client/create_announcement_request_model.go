// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnnouncementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateAnnouncementRequest
	GetAppToken() *string
	SetContent(v string) *CreateAnnouncementRequest
	GetContent() *string
	SetDisplayPage(v string) *CreateAnnouncementRequest
	GetDisplayPage() *string
	SetDisplayType(v string) *CreateAnnouncementRequest
	GetDisplayType() *string
	SetEffectiveEnd(v string) *CreateAnnouncementRequest
	GetEffectiveEnd() *string
	SetEffectiveStart(v string) *CreateAnnouncementRequest
	GetEffectiveStart() *string
	SetOpsUserCode(v string) *CreateAnnouncementRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateAnnouncementRequest
	GetOpsUserId() *string
	SetPriority(v string) *CreateAnnouncementRequest
	GetPriority() *string
	SetPublishNow(v bool) *CreateAnnouncementRequest
	GetPublishNow() *bool
	SetTargetRoleCodes(v []*string) *CreateAnnouncementRequest
	GetTargetRoleCodes() []*string
	SetTargetRoleMode(v string) *CreateAnnouncementRequest
	GetTargetRoleMode() *string
	SetTargetTenantIds(v []*int64) *CreateAnnouncementRequest
	GetTargetTenantIds() []*int64
	SetTargetTenantMode(v string) *CreateAnnouncementRequest
	GetTargetTenantMode() *string
	SetTenantId(v string) *CreateAnnouncementRequest
	GetTenantId() *string
	SetTitle(v string) *CreateAnnouncementRequest
	GetTitle() *string
}

type CreateAnnouncementRequest struct {
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
	TargetRoleCodes []*string `json:"targetRoleCodes,omitempty" xml:"targetRoleCodes,omitempty" type:"Repeated"`
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
	TargetTenantIds []*int64 `json:"targetTenantIds,omitempty" xml:"targetTenantIds,omitempty" type:"Repeated"`
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

func (s CreateAnnouncementRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnouncementRequest) GoString() string {
	return s.String()
}

func (s *CreateAnnouncementRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateAnnouncementRequest) GetContent() *string {
	return s.Content
}

func (s *CreateAnnouncementRequest) GetDisplayPage() *string {
	return s.DisplayPage
}

func (s *CreateAnnouncementRequest) GetDisplayType() *string {
	return s.DisplayType
}

func (s *CreateAnnouncementRequest) GetEffectiveEnd() *string {
	return s.EffectiveEnd
}

func (s *CreateAnnouncementRequest) GetEffectiveStart() *string {
	return s.EffectiveStart
}

func (s *CreateAnnouncementRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateAnnouncementRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateAnnouncementRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateAnnouncementRequest) GetPublishNow() *bool {
	return s.PublishNow
}

func (s *CreateAnnouncementRequest) GetTargetRoleCodes() []*string {
	return s.TargetRoleCodes
}

func (s *CreateAnnouncementRequest) GetTargetRoleMode() *string {
	return s.TargetRoleMode
}

func (s *CreateAnnouncementRequest) GetTargetTenantIds() []*int64 {
	return s.TargetTenantIds
}

func (s *CreateAnnouncementRequest) GetTargetTenantMode() *string {
	return s.TargetTenantMode
}

func (s *CreateAnnouncementRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateAnnouncementRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateAnnouncementRequest) SetAppToken(v string) *CreateAnnouncementRequest {
	s.AppToken = &v
	return s
}

func (s *CreateAnnouncementRequest) SetContent(v string) *CreateAnnouncementRequest {
	s.Content = &v
	return s
}

func (s *CreateAnnouncementRequest) SetDisplayPage(v string) *CreateAnnouncementRequest {
	s.DisplayPage = &v
	return s
}

func (s *CreateAnnouncementRequest) SetDisplayType(v string) *CreateAnnouncementRequest {
	s.DisplayType = &v
	return s
}

func (s *CreateAnnouncementRequest) SetEffectiveEnd(v string) *CreateAnnouncementRequest {
	s.EffectiveEnd = &v
	return s
}

func (s *CreateAnnouncementRequest) SetEffectiveStart(v string) *CreateAnnouncementRequest {
	s.EffectiveStart = &v
	return s
}

func (s *CreateAnnouncementRequest) SetOpsUserCode(v string) *CreateAnnouncementRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateAnnouncementRequest) SetOpsUserId(v string) *CreateAnnouncementRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateAnnouncementRequest) SetPriority(v string) *CreateAnnouncementRequest {
	s.Priority = &v
	return s
}

func (s *CreateAnnouncementRequest) SetPublishNow(v bool) *CreateAnnouncementRequest {
	s.PublishNow = &v
	return s
}

func (s *CreateAnnouncementRequest) SetTargetRoleCodes(v []*string) *CreateAnnouncementRequest {
	s.TargetRoleCodes = v
	return s
}

func (s *CreateAnnouncementRequest) SetTargetRoleMode(v string) *CreateAnnouncementRequest {
	s.TargetRoleMode = &v
	return s
}

func (s *CreateAnnouncementRequest) SetTargetTenantIds(v []*int64) *CreateAnnouncementRequest {
	s.TargetTenantIds = v
	return s
}

func (s *CreateAnnouncementRequest) SetTargetTenantMode(v string) *CreateAnnouncementRequest {
	s.TargetTenantMode = &v
	return s
}

func (s *CreateAnnouncementRequest) SetTenantId(v string) *CreateAnnouncementRequest {
	s.TenantId = &v
	return s
}

func (s *CreateAnnouncementRequest) SetTitle(v string) *CreateAnnouncementRequest {
	s.Title = &v
	return s
}

func (s *CreateAnnouncementRequest) Validate() error {
	return dara.Validate(s)
}
