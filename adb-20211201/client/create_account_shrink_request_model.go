// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *CreateAccountShrinkRequest
	GetAccountDescription() *string
	SetAccountName(v string) *CreateAccountShrinkRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateAccountShrinkRequest
	GetAccountPassword() *string
	SetAccountType(v string) *CreateAccountShrinkRequest
	GetAccountType() *string
	SetDBClusterId(v string) *CreateAccountShrinkRequest
	GetDBClusterId() *string
	SetEngine(v string) *CreateAccountShrinkRequest
	GetEngine() *string
	SetRamUserListShrink(v string) *CreateAccountShrinkRequest
	GetRamUserListShrink() *string
}

type CreateAccountShrinkRequest struct {
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
	RamUserListShrink *string `json:"RamUserList,omitempty" xml:"RamUserList,omitempty"`
}

func (s CreateAccountShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountShrinkRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *CreateAccountShrinkRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAccountShrinkRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateAccountShrinkRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateAccountShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAccountShrinkRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateAccountShrinkRequest) GetRamUserListShrink() *string {
	return s.RamUserListShrink
}

func (s *CreateAccountShrinkRequest) SetAccountDescription(v string) *CreateAccountShrinkRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetAccountName(v string) *CreateAccountShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetAccountPassword(v string) *CreateAccountShrinkRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetAccountType(v string) *CreateAccountShrinkRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetDBClusterId(v string) *CreateAccountShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetEngine(v string) *CreateAccountShrinkRequest {
	s.Engine = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetRamUserListShrink(v string) *CreateAccountShrinkRequest {
	s.RamUserListShrink = &v
	return s
}

func (s *CreateAccountShrinkRequest) Validate() error {
	return dara.Validate(s)
}
