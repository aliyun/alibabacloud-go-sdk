// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *CreateAccountRequest
	GetAccount() *string
	SetAccountType(v string) *CreateAccountRequest
	GetAccountType() *string
	SetDBInstanceId(v string) *CreateAccountRequest
	GetDBInstanceId() *string
	SetDescription(v string) *CreateAccountRequest
	GetDescription() *string
	SetDmlAuthSetting(v *CreateAccountRequestDmlAuthSetting) *CreateAccountRequest
	GetDmlAuthSetting() *CreateAccountRequestDmlAuthSetting
	SetPassword(v string) *CreateAccountRequest
	GetPassword() *string
	SetProduct(v string) *CreateAccountRequest
	GetProduct() *string
	SetRegionId(v string) *CreateAccountRequest
	GetRegionId() *string
}

type CreateAccountRequest struct {
	// The database account. The account name must meet the following requirements:
	//
	// - The name must be unique.
	//
	// - The name can contain lowercase letters, digits, and underscores (_).
	//
	// - The name must start with a lowercase letter and end with a lowercase letter or a digit.
	//
	// - The name must be 2 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The type of the database account. Valid values:
	//
	// - **NormalAccount**: Standard account.
	//
	// - **SuperAccount**: Privileged account.
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
	// The description.
	//
	// example:
	//
	// For testing purposes
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The authorization information.
	DmlAuthSetting *CreateAccountRequestDmlAuthSetting `json:"DmlAuthSetting,omitempty" xml:"DmlAuthSetting,omitempty" type:"Struct"`
	// The password of the database account. The password must meet the following requirements:
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - Special characters include !@#$%^&*()_+-=.
	//
	// - The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1b2c3d4@
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The product code.
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

func (s CreateAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) GetAccount() *string {
	return s.Account
}

func (s *CreateAccountRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateAccountRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAccountRequest) GetDmlAuthSetting() *CreateAccountRequestDmlAuthSetting {
	return s.DmlAuthSetting
}

func (s *CreateAccountRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateAccountRequest) GetProduct() *string {
	return s.Product
}

func (s *CreateAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAccountRequest) SetAccount(v string) *CreateAccountRequest {
	s.Account = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceId(v string) *CreateAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetDescription(v string) *CreateAccountRequest {
	s.Description = &v
	return s
}

func (s *CreateAccountRequest) SetDmlAuthSetting(v *CreateAccountRequestDmlAuthSetting) *CreateAccountRequest {
	s.DmlAuthSetting = v
	return s
}

func (s *CreateAccountRequest) SetPassword(v string) *CreateAccountRequest {
	s.Password = &v
	return s
}

func (s *CreateAccountRequest) SetProduct(v string) *CreateAccountRequest {
	s.Product = &v
	return s
}

func (s *CreateAccountRequest) SetRegionId(v string) *CreateAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	if s.DmlAuthSetting != nil {
		if err := s.DmlAuthSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAccountRequestDmlAuthSetting struct {
	// The list of databases to be authorized. Separate multiple databases with commas (,).
	AllowDatabases []*string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty" type:"Repeated"`
	// The list of dictionaries to be authorized. Separate multiple dictionaries with commas (,).
	AllowDictionaries []*string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty" type:"Repeated"`
	// Specifies whether to grant DDL permissions to the database account. Valid values:
	//
	// - **true**: Allows DDL operations.
	//
	// - **false**: Prohibits DDL operations.
	//
	// example:
	//
	// true
	DdlAuthority *bool `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	// Specifies whether to grant DML permissions to the database account. Valid values:
	//
	// - **0**: Allows queries for reading, writing, and changing settings.
	//
	// - **1**: Allows only queries for reading data.
	//
	// - **2**: Allows queries for reading data and changing settings.
	//
	// example:
	//
	// 0
	DmlAuthority *int32 `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
}

func (s CreateAccountRequestDmlAuthSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountRequestDmlAuthSetting) GoString() string {
	return s.String()
}

func (s *CreateAccountRequestDmlAuthSetting) GetAllowDatabases() []*string {
	return s.AllowDatabases
}

func (s *CreateAccountRequestDmlAuthSetting) GetAllowDictionaries() []*string {
	return s.AllowDictionaries
}

func (s *CreateAccountRequestDmlAuthSetting) GetDdlAuthority() *bool {
	return s.DdlAuthority
}

func (s *CreateAccountRequestDmlAuthSetting) GetDmlAuthority() *int32 {
	return s.DmlAuthority
}

func (s *CreateAccountRequestDmlAuthSetting) SetAllowDatabases(v []*string) *CreateAccountRequestDmlAuthSetting {
	s.AllowDatabases = v
	return s
}

func (s *CreateAccountRequestDmlAuthSetting) SetAllowDictionaries(v []*string) *CreateAccountRequestDmlAuthSetting {
	s.AllowDictionaries = v
	return s
}

func (s *CreateAccountRequestDmlAuthSetting) SetDdlAuthority(v bool) *CreateAccountRequestDmlAuthSetting {
	s.DdlAuthority = &v
	return s
}

func (s *CreateAccountRequestDmlAuthSetting) SetDmlAuthority(v int32) *CreateAccountRequestDmlAuthSetting {
	s.DmlAuthority = &v
	return s
}

func (s *CreateAccountRequestDmlAuthSetting) Validate() error {
	return dara.Validate(s)
}
