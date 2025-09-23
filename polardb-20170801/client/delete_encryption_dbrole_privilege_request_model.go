// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEncryptionDBRolePrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteEncryptionDBRolePrivilegeRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeleteEncryptionDBRolePrivilegeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteEncryptionDBRolePrivilegeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteEncryptionDBRolePrivilegeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteEncryptionDBRolePrivilegeRequest
	GetResourceOwnerId() *int64
	SetRolePrivilegeNameList(v string) *DeleteEncryptionDBRolePrivilegeRequest
	GetRolePrivilegeNameList() *string
}

type DeleteEncryptionDBRolePrivilegeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// test
	RolePrivilegeNameList *string `json:"RolePrivilegeNameList,omitempty" xml:"RolePrivilegeNameList,omitempty"`
}

func (s DeleteEncryptionDBRolePrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEncryptionDBRolePrivilegeRequest) GoString() string {
	return s.String()
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) GetRolePrivilegeNameList() *string {
	return s.RolePrivilegeNameList
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) SetDBClusterId(v string) *DeleteEncryptionDBRolePrivilegeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) SetOwnerAccount(v string) *DeleteEncryptionDBRolePrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) SetOwnerId(v int64) *DeleteEncryptionDBRolePrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) SetResourceOwnerAccount(v string) *DeleteEncryptionDBRolePrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) SetResourceOwnerId(v int64) *DeleteEncryptionDBRolePrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) SetRolePrivilegeNameList(v string) *DeleteEncryptionDBRolePrivilegeRequest {
	s.RolePrivilegeNameList = &v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
