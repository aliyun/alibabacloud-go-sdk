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
	SetCharacterType(v string) *CreateAccountRequest
	GetCharacterType() *string
	SetDBInstanceId(v string) *CreateAccountRequest
	GetDBInstanceId() *string
	SetDBName(v string) *CreateAccountRequest
	GetDBName() *string
	SetOwnerAccount(v string) *CreateAccountRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAccountRequest
	GetResourceOwnerId() *int64
	SetRoles(v []*CreateAccountRequestRoles) *CreateAccountRequest
	GetRoles() []*CreateAccountRequestRoles
}

type CreateAccountRequest struct {
	// The description of the account.
	//
	// example:
	//
	// test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The account name. The name must be 4 to 16 characters in length and can contain lowercase letters, digits, and underscores (_). It must start with a lowercase letter. The account has read-only permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// admin1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The account password. The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include ! @ # $ % ^ & 	- ( ) _ + - =
	//
	// > For information about how to handle special characters, see [What do I do if a connection fails because the username or password in the connection string contains special characters?](https://help.aliyun.com/document_detail/471568.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// Test123456!
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The type of the account to be created. Valid values:
	//
	// - **db**: shard account. This is the default value.
	//
	// - **cs**: Configserver account.
	//
	// - **normal**: replica set account.
	//
	// > Currently, only **db*	- accounts are supported.
	//
	// example:
	//
	// db
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-uf6e9433e955xxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database to which the account belongs. This is the database where the account is created or on which the account operates. Default value: admin. The name must be 1 to 63 characters in length and cannot contain /, \\, ., spaces, double quotation marks ("), or $. The name cannot be local.
	//
	// example:
	//
	// admin
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of authorized roles for the account.
	Roles []*CreateAccountRequestRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
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

func (s *CreateAccountRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *CreateAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateAccountRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateAccountRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAccountRequest) GetRoles() []*CreateAccountRequestRoles {
	return s.Roles
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

func (s *CreateAccountRequest) SetCharacterType(v string) *CreateAccountRequest {
	s.CharacterType = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceId(v string) *CreateAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetDBName(v string) *CreateAccountRequest {
	s.DBName = &v
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

func (s *CreateAccountRequest) SetResourceOwnerAccount(v string) *CreateAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerId(v int64) *CreateAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetRoles(v []*CreateAccountRequestRoles) *CreateAccountRequest {
	s.Roles = v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAccountRequestRoles struct {
	// The authorized database.
	//
	// example:
	//
	// admin
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The granted permissions.
	//
	// example:
	//
	// read
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s CreateAccountRequestRoles) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountRequestRoles) GoString() string {
	return s.String()
}

func (s *CreateAccountRequestRoles) GetDBName() *string {
	return s.DBName
}

func (s *CreateAccountRequestRoles) GetRole() *string {
	return s.Role
}

func (s *CreateAccountRequestRoles) SetDBName(v string) *CreateAccountRequestRoles {
	s.DBName = &v
	return s
}

func (s *CreateAccountRequestRoles) SetRole(v string) *CreateAccountRequestRoles {
	s.Role = &v
	return s
}

func (s *CreateAccountRequestRoles) Validate() error {
	return dara.Validate(s)
}
