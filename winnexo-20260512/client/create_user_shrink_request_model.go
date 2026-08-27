// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *CreateUserShrinkRequest
	GetDisplayName() *string
	SetPasswordEncrypted(v string) *CreateUserShrinkRequest
	GetPasswordEncrypted() *string
	SetRoleCodesShrink(v string) *CreateUserShrinkRequest
	GetRoleCodesShrink() *string
	SetTenantId(v string) *CreateUserShrinkRequest
	GetTenantId() *string
	SetWnAccountId(v string) *CreateUserShrinkRequest
	GetWnAccountId() *string
}

type CreateUserShrinkRequest struct {
	// The cluster name.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The base64-encoded password ciphertext encrypted by RSA-OAEP-SHA256 (required).
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	PasswordEncrypted *string `json:"passwordEncrypted,omitempty" xml:"passwordEncrypted,omitempty"`
	// The list of new system role codes (full replacement, must contain at least one role). Valid values: SUPER_ADMIN, SYSTEM_ADMIN, SEMANTIC_ADMIN, SKILL_ADMIN, KB_ADMIN, AGENT_ADMIN, and APPLICATION_USER.
	//
	// example:
	//
	// string_value
	RoleCodesShrink *string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty"`
	// The ID of the tenant in which the operation takes effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The WINNEXO logon account (unique identifier, required).
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

func (s *CreateUserShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserShrinkRequest) GetPasswordEncrypted() *string {
	return s.PasswordEncrypted
}

func (s *CreateUserShrinkRequest) GetRoleCodesShrink() *string {
	return s.RoleCodesShrink
}

func (s *CreateUserShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateUserShrinkRequest) GetWnAccountId() *string {
	return s.WnAccountId
}

func (s *CreateUserShrinkRequest) SetDisplayName(v string) *CreateUserShrinkRequest {
	s.DisplayName = &v
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
