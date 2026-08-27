// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserWithGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *CreateUserWithGroupsRequest
	GetDisplayName() *string
	SetPasswordEncrypted(v string) *CreateUserWithGroupsRequest
	GetPasswordEncrypted() *string
	SetRoleCodes(v []*string) *CreateUserWithGroupsRequest
	GetRoleCodes() []*string
	SetTenantId(v string) *CreateUserWithGroupsRequest
	GetTenantId() *string
	SetUserGroupIds(v []*string) *CreateUserWithGroupsRequest
	GetUserGroupIds() []*string
	SetWnAccountId(v string) *CreateUserWithGroupsRequest
	GetWnAccountId() *string
}

type CreateUserWithGroupsRequest struct {
	// The display name of the user (unique within the tenant, required, up to 100 characters).
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The Base64-encoded password ciphertext encrypted by RSA-OAEP-SHA256 (required).
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	PasswordEncrypted *string `json:"passwordEncrypted,omitempty" xml:"passwordEncrypted,omitempty"`
	// The list of system role codes. Valid values: SUPER_ADMIN, SYSTEM_ADMIN, SEMANTIC_ADMIN, SKILL_ADMIN, KB_ADMIN, AGENT_ADMIN, and APPLICATION_USER. Default value: APPLICATION_USER.
	//
	// example:
	//
	// string_value
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of initial user group IDs. This parameter is optional. All user groups must belong to the current tenant.
	//
	// example:
	//
	// string_value
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// The WINNEXO logon account (unique identifier, required).
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleAccountId
	WnAccountId *string `json:"wnAccountId,omitempty" xml:"wnAccountId,omitempty"`
}

func (s CreateUserWithGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserWithGroupsRequest) GoString() string {
	return s.String()
}

func (s *CreateUserWithGroupsRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserWithGroupsRequest) GetPasswordEncrypted() *string {
	return s.PasswordEncrypted
}

func (s *CreateUserWithGroupsRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *CreateUserWithGroupsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateUserWithGroupsRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateUserWithGroupsRequest) GetWnAccountId() *string {
	return s.WnAccountId
}

func (s *CreateUserWithGroupsRequest) SetDisplayName(v string) *CreateUserWithGroupsRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserWithGroupsRequest) SetPasswordEncrypted(v string) *CreateUserWithGroupsRequest {
	s.PasswordEncrypted = &v
	return s
}

func (s *CreateUserWithGroupsRequest) SetRoleCodes(v []*string) *CreateUserWithGroupsRequest {
	s.RoleCodes = v
	return s
}

func (s *CreateUserWithGroupsRequest) SetTenantId(v string) *CreateUserWithGroupsRequest {
	s.TenantId = &v
	return s
}

func (s *CreateUserWithGroupsRequest) SetUserGroupIds(v []*string) *CreateUserWithGroupsRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateUserWithGroupsRequest) SetWnAccountId(v string) *CreateUserWithGroupsRequest {
	s.WnAccountId = &v
	return s
}

func (s *CreateUserWithGroupsRequest) Validate() error {
	return dara.Validate(s)
}
