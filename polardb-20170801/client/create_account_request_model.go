// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *CreateAccountRequest
	GetAccountDescription() *string
	SetAccountName(v string) *CreateAccountRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateAccountRequest
	GetAccountPassword() *string
	SetAccountPrivilege(v string) *CreateAccountRequest
	GetAccountPrivilege() *string
	SetAccountType(v string) *CreateAccountRequest
	GetAccountType() *string
	SetClientToken(v string) *CreateAccountRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateAccountRequest
	GetDBClusterId() *string
	SetDBName(v string) *CreateAccountRequest
	GetDBName() *string
	SetNodeType(v string) *CreateAccountRequest
	GetNodeType() *string
	SetOwnerAccount(v string) *CreateAccountRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAccountRequest
	GetOwnerId() *int64
	SetPrivForAllDB(v string) *CreateAccountRequest
	GetPrivForAllDB() *string
	SetResourceOwnerAccount(v string) *CreateAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAccountRequest
	GetResourceOwnerId() *int64
}

type CreateAccountRequest struct {
	// The description of the account. The description must meet the following requirements:
	//
	// 	- It cannot start with `http://` or `https://`.
	//
	// 	- It must be 2 to 256 characters in length.
	//
	// example:
	//
	// testdes
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the account. The name must meet the following requirements:
	//
	// 	- It must start with a lowercase letter and end with a letter or a digit.
	//
	// 	- It can contain lowercase letters, digits, and underscores (_).
	//
	// 	- It must be 2 to 16 characters in length.
	//
	// 	- It cannot be root, admin, or another username that is reserved by the system.
	//
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the account. The password must meet the following requirements:
	//
	// 	- The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// 	- Special characters include `! @ # $ % ^ & 	- ( ) _ + - =`
	//
	// This parameter is required.
	//
	// example:
	//
	// Test1111
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The permissions that are granted to the account. Valid values:
	//
	// 	- **ReadWrite**: read and write permissions.
	//
	// 	- **ReadOnly**: read-only permissions.
	//
	// 	- **DMLOnly**: the permissions to execute only DML statements.
	//
	// 	- **DDLOnly**: the permissions to execute only DDL statements.
	//
	// 	- **ReadIndex**: the read-only and index permissions.
	//
	// >
	//
	// 	- `AccountPrivilege` is valid only after you specify `DBName`.
	//
	// 	- If multiple database names are specified by the `DBName` parameter, you must grant permissions on the databases. Separate multiple permissions with commas (,), and make sure that the length of the value of `AccountPrivilege` does not exceed 900. For example, if you want to grant the account the read and write permissions on DB1 and the read-only permissions on DB2, set `DBName` to `DB1,DB2` and set `AccountPrivilege` to `ReadWrite,ReadOnly`.
	//
	// 	- This parameter is valid only for standard accounts of PolarDB for MySQL clusters.
	//
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The type of the account. Valid values:
	//
	// 	- **Normal**: standard account
	//
	// 	- **Super**: privileged account.
	//
	// >
	//
	// 	- If you leave this parameter empty, the default value **Super*	- is used.
	//
	// 	- You can create multiple privileged accounts for a PolarDB for PostgreSQL (Compatible with Oracle) cluster or a PolarDB for PostgreSQL cluster. A privileged account has more permissions than a standard account. For more information, see [Create a database account](https://help.aliyun.com/document_detail/68508.html).
	//
	// 	- You can create only one privileged account for a PolarDB for MySQL cluster. A privileged account has more permissions than a standard account. For more information, see [Create a database account](https://help.aliyun.com/document_detail/68508.html).
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. The token is case-sensitive.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database that can be accessed by the account. To enter multiple database names, separate the names with commas (,).
	//
	// >  This parameter is valid only for standard accounts of PolarDB for MySQL clusters.
	//
	// example:
	//
	// testdb
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	NodeType             *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PrivForAllDB         *string `json:"PrivForAllDB,omitempty" xml:"PrivForAllDB,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *CreateAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAccountRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateAccountRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *CreateAccountRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateAccountRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAccountRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAccountRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateAccountRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateAccountRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAccountRequest) GetPrivForAllDB() *string {
	return s.PrivForAllDB
}

func (s *CreateAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPrivilege(v string) *CreateAccountRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetClientToken(v string) *CreateAccountRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAccountRequest) SetDBClusterId(v string) *CreateAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountRequest) SetDBName(v string) *CreateAccountRequest {
	s.DBName = &v
	return s
}

func (s *CreateAccountRequest) SetNodeType(v string) *CreateAccountRequest {
	s.NodeType = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerAccount(v string) *CreateAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetPrivForAllDB(v string) *CreateAccountRequest {
	s.PrivForAllDB = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerAccount(v string) *CreateAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerId(v int64) *CreateAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	return dara.Validate(s)
}
