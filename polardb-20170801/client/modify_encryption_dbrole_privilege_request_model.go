// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEncryptionDBRolePrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyEncryptionDBRolePrivilegeRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyEncryptionDBRolePrivilegeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyEncryptionDBRolePrivilegeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyEncryptionDBRolePrivilegeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyEncryptionDBRolePrivilegeRequest
	GetResourceOwnerId() *int64
	SetRolePrivilegeConfig(v string) *ModifyEncryptionDBRolePrivilegeRequest
	GetRolePrivilegeConfig() *string
	SetRolePrivilegeName(v string) *ModifyEncryptionDBRolePrivilegeRequest
	GetRolePrivilegeName() *string
}

type ModifyEncryptionDBRolePrivilegeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
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

func (s ModifyEncryptionDBRolePrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEncryptionDBRolePrivilegeRequest) GoString() string {
	return s.String()
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) GetRolePrivilegeConfig() *string {
	return s.RolePrivilegeConfig
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) GetRolePrivilegeName() *string {
	return s.RolePrivilegeName
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) SetDBClusterId(v string) *ModifyEncryptionDBRolePrivilegeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) SetOwnerAccount(v string) *ModifyEncryptionDBRolePrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) SetOwnerId(v int64) *ModifyEncryptionDBRolePrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) SetResourceOwnerAccount(v string) *ModifyEncryptionDBRolePrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) SetResourceOwnerId(v int64) *ModifyEncryptionDBRolePrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) SetRolePrivilegeConfig(v string) *ModifyEncryptionDBRolePrivilegeRequest {
	s.RolePrivilegeConfig = &v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) SetRolePrivilegeName(v string) *ModifyEncryptionDBRolePrivilegeRequest {
	s.RolePrivilegeName = &v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
