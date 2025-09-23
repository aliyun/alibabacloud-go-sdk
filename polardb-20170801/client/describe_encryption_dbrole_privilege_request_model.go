// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionDBRolePrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeEncryptionDBRolePrivilegeRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeEncryptionDBRolePrivilegeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeEncryptionDBRolePrivilegeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeEncryptionDBRolePrivilegeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEncryptionDBRolePrivilegeRequest
	GetResourceOwnerId() *int64
	SetRolePrivilegeNameList(v string) *DescribeEncryptionDBRolePrivilegeRequest
	GetRolePrivilegeNameList() *string
}

type DescribeEncryptionDBRolePrivilegeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-wz9fb5nn44u1d****
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

func (s DescribeEncryptionDBRolePrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionDBRolePrivilegeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) GetRolePrivilegeNameList() *string {
	return s.RolePrivilegeNameList
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) SetDBClusterId(v string) *DescribeEncryptionDBRolePrivilegeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) SetOwnerAccount(v string) *DescribeEncryptionDBRolePrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) SetOwnerId(v int64) *DescribeEncryptionDBRolePrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) SetResourceOwnerAccount(v string) *DescribeEncryptionDBRolePrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) SetResourceOwnerId(v int64) *DescribeEncryptionDBRolePrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) SetRolePrivilegeNameList(v string) *DescribeEncryptionDBRolePrivilegeRequest {
	s.RolePrivilegeNameList = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
