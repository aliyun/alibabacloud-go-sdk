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
	SetAccountType(v string) *CreateAccountRequest
	GetAccountType() *string
	SetDBClusterId(v string) *CreateAccountRequest
	GetDBClusterId() *string
	SetEngine(v string) *CreateAccountRequest
	GetEngine() *string
}

type CreateAccountRequest struct {
	// The description of the account.
	//
	// 	- The description cannot start with `http://` or `https://`.
	//
	// 	- The description can be up to 256 characters in length.
	//
	// example:
	//
	// test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// 	- The name must start with a lowercase letter and end with a lowercase letter or a digit.
	//
	// 	- The name can contain lowercase letters, digits, and underscores (_).
	//
	// 	- The name must be 2 to 16 characters in length.
	//
	// 	- Reserved account names such as root, admin, and opsadmin cannot be used.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_accout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// 	- The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- Special characters include `! @ # $ % ^ & 	- ( ) _ + - =`
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test_accout1
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The type of the database account. Valid values:
	//
	// 	- **Normal**: standard account.
	//
	// 	- **Super**: privileged account.
	//
	// This parameter is required.
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine of the cluster. Valid values:
	//
	// 	- **AnalyticDB*	- (default): the AnalyticDB for MySQL engine.
	//
	// 	- **Clickhouse**: the wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
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

func (s *CreateAccountRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateAccountRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAccountRequest) GetEngine() *string {
	return s.Engine
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

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetDBClusterId(v string) *CreateAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountRequest) SetEngine(v string) *CreateAccountRequest {
	s.Engine = &v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	return dara.Validate(s)
}
