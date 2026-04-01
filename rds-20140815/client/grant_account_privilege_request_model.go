// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAccountPrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GrantAccountPrivilegeRequest
	GetAccountName() *string
	SetAccountPrivilege(v string) *GrantAccountPrivilegeRequest
	GetAccountPrivilege() *string
	SetDBInstanceId(v string) *GrantAccountPrivilegeRequest
	GetDBInstanceId() *string
	SetDBName(v string) *GrantAccountPrivilegeRequest
	GetDBName() *string
	SetResourceOwnerId(v int64) *GrantAccountPrivilegeRequest
	GetResourceOwnerId() *int64
}

type GrantAccountPrivilegeRequest struct {
	// The username of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The permissions that you want to grant to the account. The number of permissions must be the same as the number of databases that you specify for the DBName parameter. You can specify this parameter based on your business requirements. Valid values:
	//
	// 	- **ReadWrite**: read and write permissions
	//
	// 	- **ReadOnly**: read-only permissions
	//
	// 	- **DDLOnly**: DDL-only permissions
	//
	// 	- **DMLOnly**: DML-only permissions
	//
	// 	- **DBOwner**: database owner permissions
	//
	// >
	//
	// 	- If the instance runs MySQL or MariaDB, you can set this parameter to **ReadWrite**, **ReadOnly**, **DDLOnly**, or **DMLOnly**.
	//
	// 	- If the instance runs SQL Server, you can set this parameter to **ReadWrite**, **ReadOnly**, or **DBOwner**.
	//
	// 	- If the instance runs PostgreSQL and uses cloud disks, you can set this parameter to **DBOwner**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database on which you want to grant permissions. Separate multiple database names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// testDB1
	DBName          *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GrantAccountPrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GrantAccountPrivilegeRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *GrantAccountPrivilegeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GrantAccountPrivilegeRequest) GetDBName() *string {
	return s.DBName
}

func (s *GrantAccountPrivilegeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GrantAccountPrivilegeRequest) SetAccountName(v string) *GrantAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetAccountPrivilege(v string) *GrantAccountPrivilegeRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetDBInstanceId(v string) *GrantAccountPrivilegeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetDBName(v string) *GrantAccountPrivilegeRequest {
	s.DBName = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
