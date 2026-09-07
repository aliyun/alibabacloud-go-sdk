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
	SetPromqlInsertPrivilegesShrink(v string) *CreateAccountShrinkRequest
	GetPromqlInsertPrivilegesShrink() *string
	SetPromqlSelectNodePercentage(v float64) *CreateAccountShrinkRequest
	GetPromqlSelectNodePercentage() *float64
	SetPromqlSelectPrivilegesShrink(v string) *CreateAccountShrinkRequest
	GetPromqlSelectPrivilegesShrink() *string
	SetRamUserListShrink(v string) *CreateAccountShrinkRequest
	GetRamUserListShrink() *string
	SetResourceGroupName(v string) *CreateAccountShrinkRequest
	GetResourceGroupName() *string
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
	// Database connection test account
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
	// - **AnalyticDB*	- (default): the AnalyticDB for MySQL engine.
	//
	// - **Clickhouse**: the wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine                       *string  `json:"Engine,omitempty" xml:"Engine,omitempty"`
	PromqlInsertPrivilegesShrink *string  `json:"PromqlInsertPrivileges,omitempty" xml:"PromqlInsertPrivileges,omitempty"`
	PromqlSelectNodePercentage   *float64 `json:"PromqlSelectNodePercentage,omitempty" xml:"PromqlSelectNodePercentage,omitempty"`
	PromqlSelectPrivilegesShrink *string  `json:"PromqlSelectPrivileges,omitempty" xml:"PromqlSelectPrivileges,omitempty"`
	// The list of Alibaba Cloud RAM user IDs to bind. Currently, only one RAM user can be bound.
	RamUserListShrink *string `json:"RamUserList,omitempty" xml:"RamUserList,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
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

func (s *CreateAccountShrinkRequest) GetPromqlInsertPrivilegesShrink() *string {
	return s.PromqlInsertPrivilegesShrink
}

func (s *CreateAccountShrinkRequest) GetPromqlSelectNodePercentage() *float64 {
	return s.PromqlSelectNodePercentage
}

func (s *CreateAccountShrinkRequest) GetPromqlSelectPrivilegesShrink() *string {
	return s.PromqlSelectPrivilegesShrink
}

func (s *CreateAccountShrinkRequest) GetRamUserListShrink() *string {
	return s.RamUserListShrink
}

func (s *CreateAccountShrinkRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
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

func (s *CreateAccountShrinkRequest) SetPromqlInsertPrivilegesShrink(v string) *CreateAccountShrinkRequest {
	s.PromqlInsertPrivilegesShrink = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetPromqlSelectNodePercentage(v float64) *CreateAccountShrinkRequest {
	s.PromqlSelectNodePercentage = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetPromqlSelectPrivilegesShrink(v string) *CreateAccountShrinkRequest {
	s.PromqlSelectPrivilegesShrink = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetRamUserListShrink(v string) *CreateAccountShrinkRequest {
	s.RamUserListShrink = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetResourceGroupName(v string) *CreateAccountShrinkRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *CreateAccountShrinkRequest) Validate() error {
	return dara.Validate(s)
}
