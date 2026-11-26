// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateUserShrinkRequest
	GetAppToken() *string
	SetDisplayName(v string) *CreateUserShrinkRequest
	GetDisplayName() *string
	SetOpsUserCode(v string) *CreateUserShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateUserShrinkRequest
	GetOpsUserId() *string
	SetPasswordEncrypted(v string) *CreateUserShrinkRequest
	GetPasswordEncrypted() *string
	SetRoleCodesShrink(v string) *CreateUserShrinkRequest
	GetRoleCodesShrink() *string
	SetSsoProvider(v string) *CreateUserShrinkRequest
	GetSsoProvider() *string
	SetTenantId(v string) *CreateUserShrinkRequest
	GetTenantId() *string
	SetWnAccountId(v string) *CreateUserShrinkRequest
	GetWnAccountId() *string
}

type CreateUserShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The display name.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
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
	// The Base64-encoded password ciphertext encrypted by using RSA-OAEP-SHA256. This parameter is required and cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	PasswordEncrypted *string `json:"passwordEncrypted,omitempty" xml:"passwordEncrypted,omitempty"`
	// The list of new system role codes. This list fully replaces the existing roles and must contain at least one role. Valid values: SUPER_ADMIN, SYSTEM_ADMIN, SEMANTIC_ADMIN, SKILL_ADMIN, KB_ADMIN, AGENT_ADMIN, and APPLICATION_USER.
	//
	// example:
	//
	// ["SYSTEM_ADMIN"]
	RoleCodesShrink *string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty"`
	// The type of the SSO provider. This parameter is optional if the tenant has only one external logon method, and is required if the tenant has multiple external logon methods. Currently, createUser supports BUILD_IN and AGENT_ONE.
	//
	// example:
	//
	// AGENT_ONE
	SsoProvider *string `json:"ssoProvider,omitempty" xml:"ssoProvider,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The WINNEXO logon account. This is a unique identifier and cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleAccountId
	WnAccountId *string `json:"wnAccountId,omitempty" xml:"wnAccountId,omitempty"`
}

func (s CreateUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateUserShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateUserShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateUserShrinkRequest) GetPasswordEncrypted() *string {
	return s.PasswordEncrypted
}

func (s *CreateUserShrinkRequest) GetRoleCodesShrink() *string {
	return s.RoleCodesShrink
}

func (s *CreateUserShrinkRequest) GetSsoProvider() *string {
	return s.SsoProvider
}

func (s *CreateUserShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateUserShrinkRequest) GetWnAccountId() *string {
	return s.WnAccountId
}

func (s *CreateUserShrinkRequest) SetAppToken(v string) *CreateUserShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *CreateUserShrinkRequest) SetDisplayName(v string) *CreateUserShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserShrinkRequest) SetOpsUserCode(v string) *CreateUserShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateUserShrinkRequest) SetOpsUserId(v string) *CreateUserShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateUserShrinkRequest) SetPasswordEncrypted(v string) *CreateUserShrinkRequest {
	s.PasswordEncrypted = &v
	return s
}

func (s *CreateUserShrinkRequest) SetRoleCodesShrink(v string) *CreateUserShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *CreateUserShrinkRequest) SetSsoProvider(v string) *CreateUserShrinkRequest {
	s.SsoProvider = &v
	return s
}

func (s *CreateUserShrinkRequest) SetTenantId(v string) *CreateUserShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateUserShrinkRequest) SetWnAccountId(v string) *CreateUserShrinkRequest {
	s.WnAccountId = &v
	return s
}

func (s *CreateUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
