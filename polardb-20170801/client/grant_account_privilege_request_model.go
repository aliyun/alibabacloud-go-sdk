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
	SetDBClusterId(v string) *GrantAccountPrivilegeRequest
	GetDBClusterId() *string
	SetDBName(v string) *GrantAccountPrivilegeRequest
	GetDBName() *string
	SetOwnerAccount(v string) *GrantAccountPrivilegeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GrantAccountPrivilegeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GrantAccountPrivilegeRequest
	GetResourceOwnerAccount() *string
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
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The permissions that are granted to the account. Valid values:
	//
	// 	- **ReadWrite**: read and write permissions
	//
	// 	- **ReadOnly**: read-only permissions
	//
	// 	- **DMLOnly**: The account is granted the permissions to execute only DML statements on the database.
	//
	// 	- **DDLOnly**: The account is granted the permissions to execute only DDL statements on the database.
	//
	// 	- **ReadIndex**: The account has the read and index permissions on the database.
	//
	// > The number of **AccountPrivilege*	- values must be the consistent with the number of **DBName*	- values. Each account permission must correspond to a database name in sequence. For example, you can set **DBName*	- to `testdb_1,testdb_2` and set **AccountPrivilege*	- to `ReadWrite,ReadOnly`. In this case, the specified standard account is granted the **read and write*	- permissions on the **testdb_1*	- database and the **read*	- permission on the **testdb_2*	- database.
	//
	// This parameter is required.
	//
	// example:
	//
	// ReadWrite,ReadOnly
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The names of the databases that the account can access. You can grant the access permissions on one or more databases to the specified standard account. If you need to specify multiple database names, separate the database names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb_1,testdb_2
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *GrantAccountPrivilegeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GrantAccountPrivilegeRequest) GetDBName() *string {
	return s.DBName
}

func (s *GrantAccountPrivilegeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GrantAccountPrivilegeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GrantAccountPrivilegeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
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

func (s *GrantAccountPrivilegeRequest) SetDBClusterId(v string) *GrantAccountPrivilegeRequest {
	s.DBClusterId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetDBName(v string) *GrantAccountPrivilegeRequest {
	s.DBName = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetOwnerAccount(v string) *GrantAccountPrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerAccount(v string) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
