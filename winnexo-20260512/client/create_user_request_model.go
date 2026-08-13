// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *CreateUserRequest
	GetDisplayName() *string
	SetPasswordEncrypted(v string) *CreateUserRequest
	GetPasswordEncrypted() *string
	SetRoleCodes(v []*string) *CreateUserRequest
	GetRoleCodes() []*string
	SetTenantId(v string) *CreateUserRequest
	GetTenantId() *string
	SetWnAccountId(v string) *CreateUserRequest
	GetWnAccountId() *string
}

type CreateUserRequest struct {
	// 用户显示名称（租户内唯一，不可为空，最多100字）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// RSA-OAEP-SHA256 加密后的 base64 密码密文（必填，不可为空）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	PasswordEncrypted *string `json:"passwordEncrypted,omitempty" xml:"passwordEncrypted,omitempty"`
	// 系统角色 code 列表，可选值: SUPER_ADMIN / SYSTEM_ADMIN / SEMANTIC_ADMIN / SKILL_ADMIN / KB_ADMIN / AGENT_ADMIN / APPLICATION_USER。不传默认 APPLICATION_USER
	//
	// example:
	//
	// string_value
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// WINNEXO 登录账号（唯一标识，不可为空）
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

func (s *CreateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserRequest) GetPasswordEncrypted() *string {
	return s.PasswordEncrypted
}

func (s *CreateUserRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *CreateUserRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateUserRequest) GetWnAccountId() *string {
	return s.WnAccountId
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
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
