// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *CreateAccountZonalRequest
	GetAccountDescription() *string
	SetAccountName(v string) *CreateAccountZonalRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateAccountZonalRequest
	GetAccountPassword() *string
	SetAccountPrivilege(v string) *CreateAccountZonalRequest
	GetAccountPrivilege() *string
	SetAccountType(v string) *CreateAccountZonalRequest
	GetAccountType() *string
	SetClientToken(v string) *CreateAccountZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateAccountZonalRequest
	GetDBClusterId() *string
	SetDBName(v string) *CreateAccountZonalRequest
	GetDBName() *string
	SetNodeType(v string) *CreateAccountZonalRequest
	GetNodeType() *string
	SetOwnerAccount(v string) *CreateAccountZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAccountZonalRequest
	GetOwnerId() *int64
	SetPrivForAllDB(v string) *CreateAccountZonalRequest
	GetPrivForAllDB() *string
	SetResourceOwnerAccount(v string) *CreateAccountZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAccountZonalRequest
	GetResourceOwnerId() *int64
}

type CreateAccountZonalRequest struct {
	// The description of the account. The description must meet the following requirements:
	//
	// - Cannot start with `http://` or `https://`.
	//
	// - Be 2 to 256 characters in length.
	//
	// example:
	//
	// testdes
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The account name. The name must meet the following requirements:
	//
	// - Start with a lowercase letter and end with a letter or a digit.
	//
	// - Contain only lowercase letters, digits, and underscores (_).
	//
	// - Be 2 to 16 characters in length.
	//
	// - Cannot be a reserved username, such as root or admin.
	//
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The account password. The password must meet the following requirements:
	//
	// - Contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - Be 8 to 32 characters in length.
	//
	// - Special characters are `!@#$%^&*()_+-=`.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test1111
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The permissions of the account. Valid values:
	//
	// - ReadWrite: read and write permissions.
	//
	// - ReadOnly: read-only permissions.
	//
	// - DMLOnly: DML permissions only.
	//
	// - DDLOnly: DDL permissions only.
	//
	// - ReadIndex: read and index permissions.
	//
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The account type. Valid values:
	//
	// - Normal: a standard account.
	//
	// - Super: a privileged account.
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token is case-sensitive and can contain a maximum of 64 ASCII characters.
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
	// The name of the database that the destination account can access. You can specify multiple database names. Separate them with commas (,).
	//
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The node type.
	//
	// example:
	//
	// Normal
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to grant permissions on all current and future databases in the cluster. Valid values:
	//
	// - 0 or empty: Does not grant permissions.
	//
	// - 1: Grants permissions.
	//
	// example:
	//
	// 0
	PrivForAllDB         *string `json:"PrivForAllDB,omitempty" xml:"PrivForAllDB,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateAccountZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountZonalRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountZonalRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *CreateAccountZonalRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAccountZonalRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateAccountZonalRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *CreateAccountZonalRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateAccountZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAccountZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAccountZonalRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateAccountZonalRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateAccountZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAccountZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAccountZonalRequest) GetPrivForAllDB() *string {
	return s.PrivForAllDB
}

func (s *CreateAccountZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAccountZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAccountZonalRequest) SetAccountDescription(v string) *CreateAccountZonalRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountZonalRequest) SetAccountName(v string) *CreateAccountZonalRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountZonalRequest) SetAccountPassword(v string) *CreateAccountZonalRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountZonalRequest) SetAccountPrivilege(v string) *CreateAccountZonalRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateAccountZonalRequest) SetAccountType(v string) *CreateAccountZonalRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountZonalRequest) SetClientToken(v string) *CreateAccountZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAccountZonalRequest) SetDBClusterId(v string) *CreateAccountZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountZonalRequest) SetDBName(v string) *CreateAccountZonalRequest {
	s.DBName = &v
	return s
}

func (s *CreateAccountZonalRequest) SetNodeType(v string) *CreateAccountZonalRequest {
	s.NodeType = &v
	return s
}

func (s *CreateAccountZonalRequest) SetOwnerAccount(v string) *CreateAccountZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountZonalRequest) SetOwnerId(v int64) *CreateAccountZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountZonalRequest) SetPrivForAllDB(v string) *CreateAccountZonalRequest {
	s.PrivForAllDB = &v
	return s
}

func (s *CreateAccountZonalRequest) SetResourceOwnerAccount(v string) *CreateAccountZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountZonalRequest) SetResourceOwnerId(v int64) *CreateAccountZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountZonalRequest) Validate() error {
	return dara.Validate(s)
}
