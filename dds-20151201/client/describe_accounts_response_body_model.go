// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v *DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody
	GetAccounts() *DescribeAccountsResponseBodyAccounts
	SetRequestId(v string) *DescribeAccountsResponseBody
	GetRequestId() *string
}

type DescribeAccountsResponseBody struct {
	Accounts *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B562A65B-39AB-4EE8-8635-5A222054FB35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) GetAccounts() *DescribeAccountsResponseBodyAccounts {
	return s.Accounts
}

func (s *DescribeAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountsResponseBody) SetAccounts(v *DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponseBody) Validate() error {
	if s.Accounts != nil {
		if err := s.Accounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccounts struct {
	Account []*DescribeAccountsResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccounts) GetAccount() []*DescribeAccountsResponseBodyAccountsAccount {
	return s.Account
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccount(v []*DescribeAccountsResponseBodyAccountsAccount) *DescribeAccountsResponseBodyAccounts {
	s.Account = v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) Validate() error {
	if s.Account != nil {
		for _, item := range s.Account {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccountsAccount struct {
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountStatus      *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	CharacterType      *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// admin
	DBName *string                                           `json:"DBName,omitempty" xml:"DBName,omitempty"`
	Roles  *DescribeAccountsResponseBodyAccountsAccountRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
}

func (s DescribeAccountsResponseBodyAccountsAccount) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsAccount) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *DescribeAccountsResponseBodyAccountsAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountsResponseBodyAccountsAccount) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeAccountsResponseBodyAccountsAccount) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeAccountsResponseBodyAccountsAccount) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAccountsResponseBodyAccountsAccount) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAccountsResponseBodyAccountsAccount) GetRoles() *DescribeAccountsResponseBodyAccountsAccountRoles {
	return s.Roles
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetCharacterType(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.CharacterType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetDBInstanceId(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetDBName(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.DBName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetRoles(v *DescribeAccountsResponseBodyAccountsAccountRoles) *DescribeAccountsResponseBodyAccountsAccount {
	s.Roles = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) Validate() error {
	if s.Roles != nil {
		if err := s.Roles.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccountsAccountRoles struct {
	Role []*DescribeAccountsResponseBodyAccountsAccountRolesRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccountsAccountRoles) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsAccountRoles) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsAccountRoles) GetRole() []*DescribeAccountsResponseBodyAccountsAccountRolesRole {
	return s.Role
}

func (s *DescribeAccountsResponseBodyAccountsAccountRoles) SetRole(v []*DescribeAccountsResponseBodyAccountsAccountRolesRole) *DescribeAccountsResponseBodyAccountsAccountRoles {
	s.Role = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccountRoles) Validate() error {
	if s.Role != nil {
		for _, item := range s.Role {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccountsAccountRolesRole struct {
	// example:
	//
	// admin
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// read
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsAccountRolesRole) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsAccountRolesRole) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsAccountRolesRole) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAccountsResponseBodyAccountsAccountRolesRole) GetRole() *string {
	return s.Role
}

func (s *DescribeAccountsResponseBodyAccountsAccountRolesRole) SetDBName(v string) *DescribeAccountsResponseBodyAccountsAccountRolesRole {
	s.DBName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccountRolesRole) SetRole(v string) *DescribeAccountsResponseBodyAccountsAccountRolesRole {
	s.Role = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccountRolesRole) Validate() error {
	return dara.Validate(s)
}
