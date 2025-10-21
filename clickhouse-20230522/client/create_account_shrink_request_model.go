// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *CreateAccountShrinkRequest
	GetAccount() *string
	SetAccountType(v string) *CreateAccountShrinkRequest
	GetAccountType() *string
	SetDBInstanceId(v string) *CreateAccountShrinkRequest
	GetDBInstanceId() *string
	SetDescription(v string) *CreateAccountShrinkRequest
	GetDescription() *string
	SetDmlAuthSettingShrink(v string) *CreateAccountShrinkRequest
	GetDmlAuthSettingShrink() *string
	SetPassword(v string) *CreateAccountShrinkRequest
	GetPassword() *string
	SetProduct(v string) *CreateAccountShrinkRequest
	GetProduct() *string
	SetRegionId(v string) *CreateAccountShrinkRequest
	GetRegionId() *string
}

type CreateAccountShrinkRequest struct {
	// The name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The type of the database account. Valid values:
	//
	// 	- **NormalAccount**: standard account
	//
	// 	- **SuperAccount**: privileged account
	//
	// This parameter is required.
	//
	// example:
	//
	// NormalAccount
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description of the account.
	//
	// example:
	//
	// Used for account
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about permissions.
	DmlAuthSettingShrink *string `json:"DmlAuthSetting,omitempty" xml:"DmlAuthSetting,omitempty"`
	// The password of the database account. The password must meet the following requirements:
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - The following special characters are supported: ! @ # $ % ^ & 	- ( ) _ + - =
	//
	// - The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1b2c3d4@
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAccountShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountShrinkRequest) GetAccount() *string {
	return s.Account
}

func (s *CreateAccountShrinkRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateAccountShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateAccountShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAccountShrinkRequest) GetDmlAuthSettingShrink() *string {
	return s.DmlAuthSettingShrink
}

func (s *CreateAccountShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateAccountShrinkRequest) GetProduct() *string {
	return s.Product
}

func (s *CreateAccountShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAccountShrinkRequest) SetAccount(v string) *CreateAccountShrinkRequest {
	s.Account = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetAccountType(v string) *CreateAccountShrinkRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetDBInstanceId(v string) *CreateAccountShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetDescription(v string) *CreateAccountShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetDmlAuthSettingShrink(v string) *CreateAccountShrinkRequest {
	s.DmlAuthSettingShrink = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetPassword(v string) *CreateAccountShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetProduct(v string) *CreateAccountShrinkRequest {
	s.Product = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetRegionId(v string) *CreateAccountShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAccountShrinkRequest) Validate() error {
	return dara.Validate(s)
}
