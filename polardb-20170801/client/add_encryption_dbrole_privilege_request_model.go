// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEncryptionDBRolePrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AddEncryptionDBRolePrivilegeRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *AddEncryptionDBRolePrivilegeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddEncryptionDBRolePrivilegeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddEncryptionDBRolePrivilegeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddEncryptionDBRolePrivilegeRequest
	GetResourceOwnerId() *int64
	SetRolePrivilegeConfig(v string) *AddEncryptionDBRolePrivilegeRequest
	GetRolePrivilegeConfig() *string
	SetRolePrivilegeName(v string) *AddEncryptionDBRolePrivilegeRequest
	GetRolePrivilegeName() *string
}

type AddEncryptionDBRolePrivilegeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-******************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// {"notEncryption":["alton"],"encryption":["alton01"]}
	RolePrivilegeConfig *string `json:"RolePrivilegeConfig,omitempty" xml:"RolePrivilegeConfig,omitempty"`
	// example:
	//
	// test
	RolePrivilegeName *string `json:"RolePrivilegeName,omitempty" xml:"RolePrivilegeName,omitempty"`
}

func (s AddEncryptionDBRolePrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddEncryptionDBRolePrivilegeRequest) GoString() string {
	return s.String()
}

func (s *AddEncryptionDBRolePrivilegeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AddEncryptionDBRolePrivilegeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddEncryptionDBRolePrivilegeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddEncryptionDBRolePrivilegeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddEncryptionDBRolePrivilegeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddEncryptionDBRolePrivilegeRequest) GetRolePrivilegeConfig() *string {
	return s.RolePrivilegeConfig
}

func (s *AddEncryptionDBRolePrivilegeRequest) GetRolePrivilegeName() *string {
	return s.RolePrivilegeName
}

func (s *AddEncryptionDBRolePrivilegeRequest) SetDBClusterId(v string) *AddEncryptionDBRolePrivilegeRequest {
	s.DBClusterId = &v
	return s
}

func (s *AddEncryptionDBRolePrivilegeRequest) SetOwnerAccount(v string) *AddEncryptionDBRolePrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddEncryptionDBRolePrivilegeRequest) SetOwnerId(v int64) *AddEncryptionDBRolePrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *AddEncryptionDBRolePrivilegeRequest) SetResourceOwnerAccount(v string) *AddEncryptionDBRolePrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddEncryptionDBRolePrivilegeRequest) SetResourceOwnerId(v int64) *AddEncryptionDBRolePrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddEncryptionDBRolePrivilegeRequest) SetRolePrivilegeConfig(v string) *AddEncryptionDBRolePrivilegeRequest {
	s.RolePrivilegeConfig = &v
	return s
}

func (s *AddEncryptionDBRolePrivilegeRequest) SetRolePrivilegeName(v string) *AddEncryptionDBRolePrivilegeRequest {
	s.RolePrivilegeName = &v
	return s
}

func (s *AddEncryptionDBRolePrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
