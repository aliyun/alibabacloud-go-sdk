// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateUserRequest
	GetAppToken() *string
	SetDisplayName(v string) *CreateUserRequest
	GetDisplayName() *string
	SetOpsUserCode(v string) *CreateUserRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateUserRequest
	GetOpsUserId() *string
	SetPasswordEncrypted(v string) *CreateUserRequest
	GetPasswordEncrypted() *string
	SetRoleCodes(v []*string) *CreateUserRequest
	GetRoleCodes() []*string
	SetSsoProvider(v string) *CreateUserRequest
	GetSsoProvider() *string
	SetTenantId(v string) *CreateUserRequest
	GetTenantId() *string
	SetWnAccountId(v string) *CreateUserRequest
	GetWnAccountId() *string
}

type CreateUserRequest struct {
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
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
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

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateUserRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateUserRequest) GetPasswordEncrypted() *string {
	return s.PasswordEncrypted
}

func (s *CreateUserRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *CreateUserRequest) GetSsoProvider() *string {
	return s.SsoProvider
}

func (s *CreateUserRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateUserRequest) GetWnAccountId() *string {
	return s.WnAccountId
}

func (s *CreateUserRequest) SetAppToken(v string) *CreateUserRequest {
	s.AppToken = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetOpsUserCode(v string) *CreateUserRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateUserRequest) SetOpsUserId(v string) *CreateUserRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateUserRequest) SetPasswordEncrypted(v string) *CreateUserRequest {
	s.PasswordEncrypted = &v
	return s
}

func (s *CreateUserRequest) SetRoleCodes(v []*string) *CreateUserRequest {
	s.RoleCodes = v
	return s
}

func (s *CreateUserRequest) SetSsoProvider(v string) *CreateUserRequest {
	s.SsoProvider = &v
	return s
}

func (s *CreateUserRequest) SetTenantId(v string) *CreateUserRequest {
	s.TenantId = &v
	return s
}

func (s *CreateUserRequest) SetWnAccountId(v string) *CreateUserRequest {
	s.WnAccountId = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	return dara.Validate(s)
}
