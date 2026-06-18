// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountPrivilegeRequest
	GetAccountName() *string
	SetAccountPrivilege(v string) *ModifyAccountPrivilegeRequest
	GetAccountPrivilege() *string
	SetDBInstanceName(v string) *ModifyAccountPrivilegeRequest
	GetDBInstanceName() *string
	SetDbName(v string) *ModifyAccountPrivilegeRequest
	GetDbName() *string
	SetRegionId(v string) *ModifyAccountPrivilegeRequest
	GetRegionId() *string
	SetSecurityAccountName(v string) *ModifyAccountPrivilegeRequest
	GetSecurityAccountName() *string
	SetSecurityAccountPassword(v string) *ModifyAccountPrivilegeRequest
	GetSecurityAccountPassword() *string
}

type ModifyAccountPrivilegeRequest struct {
	// The account name.
	//
	// This parameter is required.
	//
	// example:
	//
	// account_sec
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The account permissions. Valid values:
	//
	// - **ReadWrite**
	//
	// - **ReadOnly**
	//
	// - **DMLOnly**
	//
	// - **DDLOnly**.
	//
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0ori2r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The database name.
	//
	// example:
	//
	// sbtest
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the security administrator account.
	//
	// example:
	//
	// account_audit
	SecurityAccountName *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	// The password of the security administrator account.
	//
	// example:
	//
	// *****
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
}

func (s ModifyAccountPrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountPrivilegeRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *ModifyAccountPrivilegeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyAccountPrivilegeRequest) GetDbName() *string {
	return s.DbName
}

func (s *ModifyAccountPrivilegeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountPrivilegeRequest) GetSecurityAccountName() *string {
	return s.SecurityAccountName
}

func (s *ModifyAccountPrivilegeRequest) GetSecurityAccountPassword() *string {
	return s.SecurityAccountPassword
}

func (s *ModifyAccountPrivilegeRequest) SetAccountName(v string) *ModifyAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetAccountPrivilege(v string) *ModifyAccountPrivilegeRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetDBInstanceName(v string) *ModifyAccountPrivilegeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetDbName(v string) *ModifyAccountPrivilegeRequest {
	s.DbName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetRegionId(v string) *ModifyAccountPrivilegeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetSecurityAccountName(v string) *ModifyAccountPrivilegeRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetSecurityAccountPassword(v string) *ModifyAccountPrivilegeRequest {
	s.SecurityAccountPassword = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
