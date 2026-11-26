// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateShareShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateShareShrinkRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *UpdateShareShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateShareShrinkRequest
	GetOpsUserId() *string
	SetOutputItemId(v string) *UpdateShareShrinkRequest
	GetOutputItemId() *string
	SetOutputItemName(v string) *UpdateShareShrinkRequest
	GetOutputItemName() *string
	SetRefreshShareLink(v bool) *UpdateShareShrinkRequest
	GetRefreshShareLink() *bool
	SetSessionId(v string) *UpdateShareShrinkRequest
	GetSessionId() *string
	SetShareAccessType(v string) *UpdateShareShrinkRequest
	GetShareAccessType() *string
	SetShareEnabled(v bool) *UpdateShareShrinkRequest
	GetShareEnabled() *bool
	SetShareUserGroupIdsShrink(v string) *UpdateShareShrinkRequest
	GetShareUserGroupIdsShrink() *string
	SetShareUserIdsShrink(v string) *UpdateShareShrinkRequest
	GetShareUserIdsShrink() *string
	SetTargetType(v string) *UpdateShareShrinkRequest
	GetTargetType() *string
	SetTenantId(v string) *UpdateShareShrinkRequest
	GetTenantId() *string
}

type UpdateShareShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
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
	// The output item ID.
	//
	// example:
	//
	// exampleOutputItemId
	OutputItemId *string `json:"outputItemId,omitempty" xml:"outputItemId,omitempty"`
	// The name of the output item. This parameter must be used together with sessionId.
	//
	// example:
	//
	// Quarterly Review.pptx
	OutputItemName *string `json:"outputItemName,omitempty" xml:"outputItemName,omitempty"`
	// Specifies whether to refresh the sharing link. This is supported only for output items that have sharing enabled or are being enabled in the current request.
	//
	// example:
	//
	// false
	RefreshShareLink *bool `json:"refreshShareLink,omitempty" xml:"refreshShareLink,omitempty"`
	// The session ID.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The sharing scope. Valid values:
	//
	// - PUBLIC: Accessible to all users.
	//
	// - SPECIFIED: Accessible to specified users or user groups.
	//
	// example:
	//
	// SPECIFIED
	ShareAccessType *string `json:"shareAccessType,omitempty" xml:"shareAccessType,omitempty"`
	// Specifies whether sharing is enabled.
	//
	// example:
	//
	// true
	ShareEnabled *bool `json:"shareEnabled,omitempty" xml:"shareEnabled,omitempty"`
	// The list of user group IDs within the specified scope. Pass an empty array to clear the list.
	//
	// example:
	//
	// ["group-1"]
	ShareUserGroupIdsShrink *string `json:"shareUserGroupIds,omitempty" xml:"shareUserGroupIds,omitempty"`
	// The list of platform user IDs within the specified scope. Pass an empty array to clear the list.
	//
	// example:
	//
	// [10001]
	ShareUserIdsShrink *string `json:"shareUserIds,omitempty" xml:"shareUserIds,omitempty"`
	// The type of the current permission. Valid values:
	//
	// - period: Period type.
	//
	// - objective: Objective type.
	//
	// - kr: Key result type.
	//
	// This parameter is required.
	//
	// example:
	//
	// outputItem
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateShareShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateShareShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateShareShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateShareShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateShareShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateShareShrinkRequest) GetOutputItemId() *string {
	return s.OutputItemId
}

func (s *UpdateShareShrinkRequest) GetOutputItemName() *string {
	return s.OutputItemName
}

func (s *UpdateShareShrinkRequest) GetRefreshShareLink() *bool {
	return s.RefreshShareLink
}

func (s *UpdateShareShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *UpdateShareShrinkRequest) GetShareAccessType() *string {
	return s.ShareAccessType
}

func (s *UpdateShareShrinkRequest) GetShareEnabled() *bool {
	return s.ShareEnabled
}

func (s *UpdateShareShrinkRequest) GetShareUserGroupIdsShrink() *string {
	return s.ShareUserGroupIdsShrink
}

func (s *UpdateShareShrinkRequest) GetShareUserIdsShrink() *string {
	return s.ShareUserIdsShrink
}

func (s *UpdateShareShrinkRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateShareShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateShareShrinkRequest) SetAppToken(v string) *UpdateShareShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetOpsUserCode(v string) *UpdateShareShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetOpsUserId(v string) *UpdateShareShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetOutputItemId(v string) *UpdateShareShrinkRequest {
	s.OutputItemId = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetOutputItemName(v string) *UpdateShareShrinkRequest {
	s.OutputItemName = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetRefreshShareLink(v bool) *UpdateShareShrinkRequest {
	s.RefreshShareLink = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetSessionId(v string) *UpdateShareShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetShareAccessType(v string) *UpdateShareShrinkRequest {
	s.ShareAccessType = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetShareEnabled(v bool) *UpdateShareShrinkRequest {
	s.ShareEnabled = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetShareUserGroupIdsShrink(v string) *UpdateShareShrinkRequest {
	s.ShareUserGroupIdsShrink = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetShareUserIdsShrink(v string) *UpdateShareShrinkRequest {
	s.ShareUserIdsShrink = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetTargetType(v string) *UpdateShareShrinkRequest {
	s.TargetType = &v
	return s
}

func (s *UpdateShareShrinkRequest) SetTenantId(v string) *UpdateShareShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateShareShrinkRequest) Validate() error {
	return dara.Validate(s)
}
