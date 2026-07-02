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
	// - Cannot start with `http://` or `https://`.
	//
	// - Is 2 to 256 characters in length.
	//
	// example:
	//
	// testdes
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The account name. The name must meet the following requirements:
	//
	// 	- Starts with a lowercase letter and ends with a letter or digit.
	//
	// 	- Contains only lowercase letters, digits, or underscores (_).
	//
	// 	- Is 2 to 16 characters in length.
	//
	// 	- Cannot use certain reserved usernames such as root or admin.
	//
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The account password. The password must meet the following requirements:
	//
	// 	- Contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- Is 8 to 32 characters in length.
	//
	// 	- Special characters include `!@#$%^&*()_+-=`.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test1111
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The permissions of the account. Valid values:
	//
	// 	- **ReadWrite**: read and write
	//
	// 	- **ReadOnly**: read-only
	//
	// 	- **DMLOnly**: DML only
	//
	// 	- **DDLOnly**: DDL only
	//
	// 	- **ReadIndex**: read-only and index
	//
	// >	- The DBName parameter must be specified for AccountPrivilege to take effect.
	//
	// >	- If you specify multiple database names for the DBName parameter, you must grant the corresponding permissions to each database. Separate multiple permissions with commas (,) and make sure that the total length of the AccountPrivilege string does not exceed 900 characters. For example, to grant read and write permissions on database DB1 and read-only permissions on database DB2, set DBName to `DB1,DB2` and set AccountPrivilege to `ReadWrite,ReadOnly`.
	//
	// > 	- This parameter is supported only for standard accounts of PolarDB for MySQL clusters.
	//
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The account type. Valid values:
	//
	// - **Normal**: standard account.
	//
	// - **Super**: privileged account.
	//
	// - **DynamoDB**: DynamoDB account.
	//
	//
	//
	//
	// > 	- If this parameter is left empty, a **Super*	- account is created by default.
	//
	// > 	- If the cluster is a PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL cluster, you can create multiple privileged accounts for each cluster. Privileged accounts have more permissions than standard accounts. For more information, see [Create a database account](https://help.aliyun.com/document_detail/68508.html).
	//
	// > 	- If the cluster is a PolarDB for MySQL cluster, you can create at most one privileged account for each cluster. Privileged accounts have more permissions than standard accounts. For more information, see [Create a database account](https://help.aliyun.com/document_detail/68508.html).
	//
	// > 	- DynamoDB accounts are dedicated accounts created for the DynamoDB compatibility feature of PolarDB for PostgreSQL. For more information, see [DynamoDB usage instructions](https://help.aliyun.com/document_detail/2979941.html).
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value. Make sure that the value is unique among different requests. The token is case-sensitive and cannot exceed 64 ASCII characters in length.
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
	// The name of the database that the account is authorized to access. You can specify multiple database names separated by commas (,).
	//
	// > This parameter is supported only for standard accounts of PolarDB for MySQL clusters.
	//
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The node type. Valid values:
	//
	// - Search: required when creating an account for a PolarDB Search node
	//
	// example:
	//
	// Search
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to grant permissions on all existing databases and all new databases in the current cluster. Valid values:
	//
	// - **0 or empty**: does not grant permissions.
	//
	// - **1**: grants permissions.
	//
	// >	- The AccountPrivilege parameter must be specified for this parameter to take effect.
	//
	// >	- If this parameter is set to `1`, the permissions specified by AccountPrivilege are granted on all databases.
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
