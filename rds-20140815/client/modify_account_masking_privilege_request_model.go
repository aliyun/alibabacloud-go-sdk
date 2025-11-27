// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountMaskingPrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyAccountMaskingPrivilegeRequest
	GetDBInstanceName() *string
	SetDBName(v string) *ModifyAccountMaskingPrivilegeRequest
	GetDBName() *string
	SetExpireTime(v string) *ModifyAccountMaskingPrivilegeRequest
	GetExpireTime() *string
	SetOwnerId(v string) *ModifyAccountMaskingPrivilegeRequest
	GetOwnerId() *string
	SetPrivilege(v string) *ModifyAccountMaskingPrivilegeRequest
	GetPrivilege() *string
	SetRegionId(v string) *ModifyAccountMaskingPrivilegeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyAccountMaskingPrivilegeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountMaskingPrivilegeRequest
	GetResourceOwnerId() *int64
	SetUserName(v string) *ModifyAccountMaskingPrivilegeRequest
	GetUserName() *string
}

type ModifyAccountMaskingPrivilegeRequest struct {
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBName         *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExpireTime     *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Privilege            *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifyAccountMaskingPrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountMaskingPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountMaskingPrivilegeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyAccountMaskingPrivilegeRequest) GetDBName() *string {
	return s.DBName
}

func (s *ModifyAccountMaskingPrivilegeRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ModifyAccountMaskingPrivilegeRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyAccountMaskingPrivilegeRequest) GetPrivilege() *string {
	return s.Privilege
}

func (s *ModifyAccountMaskingPrivilegeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountMaskingPrivilegeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountMaskingPrivilegeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountMaskingPrivilegeRequest) GetUserName() *string {
	return s.UserName
}

func (s *ModifyAccountMaskingPrivilegeRequest) SetDBInstanceName(v string) *ModifyAccountMaskingPrivilegeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeRequest) SetDBName(v string) *ModifyAccountMaskingPrivilegeRequest {
	s.DBName = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeRequest) SetExpireTime(v string) *ModifyAccountMaskingPrivilegeRequest {
	s.ExpireTime = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeRequest) SetOwnerId(v string) *ModifyAccountMaskingPrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeRequest) SetPrivilege(v string) *ModifyAccountMaskingPrivilegeRequest {
	s.Privilege = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeRequest) SetRegionId(v string) *ModifyAccountMaskingPrivilegeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeRequest) SetResourceOwnerAccount(v string) *ModifyAccountMaskingPrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeRequest) SetResourceOwnerId(v int64) *ModifyAccountMaskingPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeRequest) SetUserName(v string) *ModifyAccountMaskingPrivilegeRequest {
	s.UserName = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
