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
	// The display name of the user. The name must be unique within the tenant and cannot exceed 100 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
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
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
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
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// The WINNEXO logon account. This parameter is a unique identifier and cannot be empty.
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
