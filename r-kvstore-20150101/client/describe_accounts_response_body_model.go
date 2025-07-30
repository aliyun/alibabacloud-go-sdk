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
	// Details about returned accounts of the instance.
	Accounts *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6C9E114C-217C-4118-83C0-B4070222****
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeAccountsResponseBodyAccountsAccount struct {
	// The description of the account.
	//
	// example:
	//
	// testdec
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// demoaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The state of the account. Valid values:
	//
	// 	- **Unavailable**: The account is unavailable.
	//
	// 	- **Available**: The account is available.
	//
	// example:
	//
	// Available
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The type of the account. Valid values:
	//
	// 	- **Normal**: standard account
	//
	// 	- **Super**: super account
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// Details about account permissions.
	DatabasePrivileges *DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges `json:"DatabasePrivileges,omitempty" xml:"DatabasePrivileges,omitempty" type:"Struct"`
	// The ID of the instance.
	//
	// example:
	//
	// r-bp10noxlhcoim2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *DescribeAccountsResponseBodyAccountsAccount) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeAccountsResponseBodyAccountsAccount) GetDatabasePrivileges() *DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges {
	return s.DatabasePrivileges
}

func (s *DescribeAccountsResponseBodyAccountsAccount) GetInstanceId() *string {
	return s.InstanceId
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

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountType(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetDatabasePrivileges(v *DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges) *DescribeAccountsResponseBodyAccountsAccount {
	s.DatabasePrivileges = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetInstanceId(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.InstanceId = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) Validate() error {
	return dara.Validate(s)
}

type DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges struct {
	DatabasePrivilege []*DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege `json:"DatabasePrivilege,omitempty" xml:"DatabasePrivilege,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges) GetDatabasePrivilege() []*DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege {
	return s.DatabasePrivilege
}

func (s *DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges) SetDatabasePrivilege(v []*DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege) *DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges {
	s.DatabasePrivilege = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccountDatabasePrivileges) Validate() error {
	return dara.Validate(s)
}

type DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege struct {
	// The permission of the account. Default value: RoleReadWrite. Valid values:
	//
	// 	- **RoleReadOnly**: The account has the read-only permissions.
	//
	// 	- **RoleReadWrite**: The account has the read and write permissions.
	//
	// example:
	//
	// RoleReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege) SetAccountPrivilege(v string) *DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccountDatabasePrivilegesDatabasePrivilege) Validate() error {
	return dara.Validate(s)
}
