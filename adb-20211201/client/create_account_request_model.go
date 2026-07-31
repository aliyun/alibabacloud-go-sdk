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
	SetRamUserList(v []*string) *CreateAccountRequest
	GetRamUserList() []*string
}

type CreateAccountRequest struct {
	// The description of the account.
	//
	// - Cannot start with `http://` or `https://`.
	//
	// - Cannot exceed 256 characters in length.
	//
	// example:
	//
	// 数据库连接测试账号
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account. The name must meet the following requirements:
	//
	// - Starts with a lowercase letter and ends with a lowercase letter or digit.
	//
	// - Contains only lowercase letters, digits, or underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_accout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// - Must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - Special characters include: `!@#$%^&*()_+-=`
	//
	// - Must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test_accout1
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The type of the account. Valid values:
	//
	// - **Normal**: standard account.
	//
	// - **Super**: privileged account.
	//
	// This parameter is required.
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine. Valid values:
	//
	// - **AnalyticDB*	- (default): AnalyticDB for MySQL engine.
	//
	// - **Clickhouse**: wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The list of Alibaba Cloud Resource Access Management (RAM) user IDs to attach. Currently, only one RAM user can be attached.
	RamUserList []*string `json:"RamUserList,omitempty" xml:"RamUserList,omitempty" type:"Repeated"`
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

func (s *CreateAccountRequest) GetRamUserList() []*string {
	return s.RamUserList
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

func (s *CreateAccountRequest) SetRamUserList(v []*string) *CreateAccountRequest {
	s.RamUserList = v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	return dara.Validate(s)
}
