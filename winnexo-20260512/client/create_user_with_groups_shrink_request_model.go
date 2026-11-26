// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserWithGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateUserWithGroupsShrinkRequest
	GetAppToken() *string
	SetDisplayName(v string) *CreateUserWithGroupsShrinkRequest
	GetDisplayName() *string
	SetOpsUserCode(v string) *CreateUserWithGroupsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateUserWithGroupsShrinkRequest
	GetOpsUserId() *string
	SetPasswordEncrypted(v string) *CreateUserWithGroupsShrinkRequest
	GetPasswordEncrypted() *string
	SetRoleCodesShrink(v string) *CreateUserWithGroupsShrinkRequest
	GetRoleCodesShrink() *string
	SetTenantId(v string) *CreateUserWithGroupsShrinkRequest
	GetTenantId() *string
	SetUserGroupIdsShrink(v string) *CreateUserWithGroupsShrinkRequest
	GetUserGroupIdsShrink() *string
	SetWnAccountId(v string) *CreateUserWithGroupsShrinkRequest
	GetWnAccountId() *string
}

type CreateUserWithGroupsShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The display name of the user. The name must be unique within the tenant and cannot exceed 100 characters in length.
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
	// The Base64-encoded password ciphertext encrypted by using the RSA-OAEP-SHA256 algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	PasswordEncrypted *string `json:"passwordEncrypted,omitempty" xml:"passwordEncrypted,omitempty"`
	// The list of initial system role codes. If this parameter is not specified, the `APPLICATION_USER` role is assigned by default.
	//
	// example:
	//
	// string_value
	RoleCodesShrink *string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this parameter explicitly by using `--tenant-id`.
	//
	// example:
	//
	// string_value
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of initial user group IDs. A maximum of 100 user group IDs can be specified. All user groups must belong to the current tenant.
	//
	// example:
	//
	// string_value
	UserGroupIdsShrink *string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty"`
	// The WINNEXO logon account. This parameter is a unique identifier and cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleAccountId
	WnAccountId *string `json:"wnAccountId,omitempty" xml:"wnAccountId,omitempty"`
}

func (s CreateUserWithGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserWithGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserWithGroupsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateUserWithGroupsShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserWithGroupsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateUserWithGroupsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateUserWithGroupsShrinkRequest) GetPasswordEncrypted() *string {
	return s.PasswordEncrypted
}

func (s *CreateUserWithGroupsShrinkRequest) GetRoleCodesShrink() *string {
	return s.RoleCodesShrink
}

func (s *CreateUserWithGroupsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateUserWithGroupsShrinkRequest) GetUserGroupIdsShrink() *string {
	return s.UserGroupIdsShrink
}

func (s *CreateUserWithGroupsShrinkRequest) GetWnAccountId() *string {
	return s.WnAccountId
}

func (s *CreateUserWithGroupsShrinkRequest) SetAppToken(v string) *CreateUserWithGroupsShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *CreateUserWithGroupsShrinkRequest) SetDisplayName(v string) *CreateUserWithGroupsShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserWithGroupsShrinkRequest) SetOpsUserCode(v string) *CreateUserWithGroupsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateUserWithGroupsShrinkRequest) SetOpsUserId(v string) *CreateUserWithGroupsShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateUserWithGroupsShrinkRequest) SetPasswordEncrypted(v string) *CreateUserWithGroupsShrinkRequest {
	s.PasswordEncrypted = &v
	return s
}

func (s *CreateUserWithGroupsShrinkRequest) SetRoleCodesShrink(v string) *CreateUserWithGroupsShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *CreateUserWithGroupsShrinkRequest) SetTenantId(v string) *CreateUserWithGroupsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateUserWithGroupsShrinkRequest) SetUserGroupIdsShrink(v string) *CreateUserWithGroupsShrinkRequest {
	s.UserGroupIdsShrink = &v
	return s
}

func (s *CreateUserWithGroupsShrinkRequest) SetWnAccountId(v string) *CreateUserWithGroupsShrinkRequest {
	s.WnAccountId = &v
	return s
}

func (s *CreateUserWithGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
