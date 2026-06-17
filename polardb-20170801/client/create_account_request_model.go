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
	// - It cannot start with `http://` or `https://`.
	//
	// - It must be 2 to 256 characters in length.
	//
	// example:
	//
	// testdes
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account. The name must meet the following requirements:
	//
	// - It must start with a lowercase letter and end with a letter or a digit.
	//
	// - It can contain lowercase letters, digits, and underscores (_).
	//
	// - It must be 2 to 16 characters in length.
	//
	// - It cannot be a reserved keyword, such as root or admin.
	//
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account. The password must meet the following requirements:
	//
	// - It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - It must be 8 to 32 characters in length.
	//
	// - The special characters are `!@#$%^&*()_+-=`.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test1111
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The privilege level to grant on the specified databases. Valid values:
	//
	// - **ReadWrite**: read and write permissions
	//
	// - **ReadOnly**: read-only permissions
	//
	// - **DMLOnly**: DML permissions only
	//
	// - **DDLOnly**: DDL permissions only
	//
	// - **ReadIndex**: read-only and index permissions
	//
	// > 	- This parameter takes effect only when you specify the `DBName` parameter.
	//
	// >
	//
	// > 	- If you specify multiple databases in `DBName`, you must specify a corresponding permission for each in `AccountPrivilege`, separated by commas. The `AccountPrivilege` string cannot exceed 900 characters. For example, to grant read and write permissions to database DB1 and read-only permissions to database DB2, set `DBName` to `DB1,DB2` and set `AccountPrivilege` to `ReadWrite,ReadOnly`.
	//
	// >
	//
	// > 	- This parameter applies only to standard accounts on PolarDB for MySQL clusters.
	//
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The type of the account. Valid values:
	//
	// - **Normal**: a standard account.
	//
	// - **Super**: a privileged account.
	//
	// > 	- If you do not specify this parameter, the system creates a **Super*	- account by default.
	//
	// >
	//
	// > 	- You can create multiple privileged accounts on PolarDB for PostgreSQL (Oracle-Compatible) and PolarDB for PostgreSQL clusters. A privileged account has more permissions than a standard account. For more information, see [Create database accounts](https://help.aliyun.com/document_detail/68508.html).
	//
	// >
	//
	// > 	- For a PolarDB for MySQL cluster, you can create only one privileged account. For more information, see [Create database accounts](https://help.aliyun.com/document_detail/68508.html).
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// A client-generated token to ensure request idempotency. The token must be unique across requests. It is case-sensitive and can be up to 64 ASCII characters long.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database that the account can access. To specify multiple databases, separate the database names with a comma (,).
	//
	// > This parameter applies only to standard accounts on PolarDB for MySQL clusters.
	//
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The type of the node. Valid values:
	//
	// - **Search**: For creating an account on a PolarDB Search node.
	//
	// example:
	//
	// Search
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to grant the account permissions on all current and future databases in the cluster. Valid values:
	//
	// - **0 or do not specify**: The specified permissions are not granted to all databases.
	//
	// - **1**: Grants the specified permissions to all current and future databases.
	//
	// > 	- This parameter takes effect only when you specify the `AccountPrivilege` parameter.
	//
	// >
	//
	// > 	- If you set this parameter to `1`, the permissions specified in `AccountPrivilege` are granted to all databases.
	//
	// example:
	//
	// 0
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
