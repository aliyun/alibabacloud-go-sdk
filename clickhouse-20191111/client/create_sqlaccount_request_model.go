// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSQLAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *CreateSQLAccountRequest
	GetAccountDescription() *string
	SetAccountName(v string) *CreateSQLAccountRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateSQLAccountRequest
	GetAccountPassword() *string
	SetAccountType(v string) *CreateSQLAccountRequest
	GetAccountType() *string
	SetDBClusterId(v string) *CreateSQLAccountRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CreateSQLAccountRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSQLAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateSQLAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSQLAccountRequest
	GetResourceOwnerId() *int64
}

type CreateSQLAccountRequest struct {
	// The description of the database account.
	//
	// 	- The description cannot start with http:// or https://.
	//
	// 	- The description can be up to 256 characters in length or be an empty string.
	//
	// example:
	//
	// Used for test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// 	- The name must be unique in the cluster.
	//
	// 	- The name can contain lowercase letters, digits, or underscores (_).
	//
	// 	- The name must start with a lowercase letter and end with a lowercase letter or digit.
	//
	// 	- The name must be 2 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// 	- The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- Special characters include ! @ # $ % ^ & \\	- ( ) _ + - =
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test1234
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The type of the database account. Valid values:
	//
	// 	- **Super**: privileged account.
	//
	// 	- **Normal**: standard account.
	//
	// This parameter is required.
	//
	// example:
	//
	// Super
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1p816075e21****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateSQLAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateSQLAccountRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *CreateSQLAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateSQLAccountRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateSQLAccountRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateSQLAccountRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateSQLAccountRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSQLAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSQLAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSQLAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSQLAccountRequest) SetAccountDescription(v string) *CreateSQLAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateSQLAccountRequest) SetAccountName(v string) *CreateSQLAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateSQLAccountRequest) SetAccountPassword(v string) *CreateSQLAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateSQLAccountRequest) SetAccountType(v string) *CreateSQLAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateSQLAccountRequest) SetDBClusterId(v string) *CreateSQLAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateSQLAccountRequest) SetOwnerAccount(v string) *CreateSQLAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSQLAccountRequest) SetOwnerId(v int64) *CreateSQLAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSQLAccountRequest) SetResourceOwnerAccount(v string) *CreateSQLAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSQLAccountRequest) SetResourceOwnerId(v int64) *CreateSQLAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSQLAccountRequest) Validate() error {
	return dara.Validate(s)
}
