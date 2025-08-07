// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v []*DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody
	GetAccounts() []*DescribeAccountsResponseBodyAccounts
	SetPageNumber(v int32) *DescribeAccountsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeAccountsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeAccountsResponseBody
	GetRequestId() *string
}

type DescribeAccountsResponseBody struct {
	// The details of the account.
	Accounts []*DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The page number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 155462B9-205F-4FFC-BB43-4855FE******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) GetAccounts() []*DescribeAccountsResponseBodyAccounts {
	return s.Accounts
}

func (s *DescribeAccountsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccountsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountsResponseBody) SetAccounts(v []*DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *DescribeAccountsResponseBody) SetPageNumber(v int32) *DescribeAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetPageRecordCount(v int32) *DescribeAccountsResponseBody {
	s.PageRecordCount = &v
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
	// The description of the account.
	//
	// example:
	//
	// test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The locking state of the account. Valid values:
	//
	// 	- **UnLock**
	//
	// 	- **Lock**
	//
	// example:
	//
	// UnLock
	AccountLockState *string `json:"AccountLockState,omitempty" xml:"AccountLockState,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// test_acc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The validity period of the password.
	//
	// example:
	//
	// undefined
	AccountPasswordValidTime *string `json:"AccountPasswordValidTime,omitempty" xml:"AccountPasswordValidTime,omitempty"`
	// The state of the account. Valid values:
	//
	// 	- **Creating**: The account is being created.
	//
	// 	- **Available**: The account is available.
	//
	// 	- **Deleting**: The account is being deleted
	//
	// example:
	//
	// Available
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The type of the account. Valid values:
	//
	// 	- **Normal**: standard account.
	//
	// 	- **Super**: privileged account.
	//
	// 	- **ReadOnly**: global read-only account.
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The list of database permissions that are granted to the account.
	DatabasePrivileges []*DescribeAccountsResponseBodyAccountsDatabasePrivileges `json:"DatabasePrivileges,omitempty" xml:"DatabasePrivileges,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccounts) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *DescribeAccountsResponseBodyAccounts) GetAccountLockState() *string {
	return s.AccountLockState
}

func (s *DescribeAccountsResponseBodyAccounts) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountsResponseBodyAccounts) GetAccountPasswordValidTime() *string {
	return s.AccountPasswordValidTime
}

func (s *DescribeAccountsResponseBodyAccounts) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeAccountsResponseBodyAccounts) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeAccountsResponseBodyAccounts) GetDatabasePrivileges() []*DescribeAccountsResponseBodyAccountsDatabasePrivileges {
	return s.DatabasePrivileges
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountLockState(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountLockState = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountName(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountPasswordValidTime(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountPasswordValidTime = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountType(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetDatabasePrivileges(v []*DescribeAccountsResponseBodyAccountsDatabasePrivileges) *DescribeAccountsResponseBodyAccounts {
	s.DatabasePrivileges = v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) Validate() error {
	return dara.Validate(s)
}

type DescribeAccountsResponseBodyAccountsDatabasePrivileges struct {
	// The permissions that the account is granted on the database. Valid values:
	//
	// example:
	//
	// ReadOnly
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// DBtest
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsDatabasePrivileges) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsDatabasePrivileges) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsDatabasePrivileges) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *DescribeAccountsResponseBodyAccountsDatabasePrivileges) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAccountsResponseBodyAccountsDatabasePrivileges) SetAccountPrivilege(v string) *DescribeAccountsResponseBodyAccountsDatabasePrivileges {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDatabasePrivileges) SetDBName(v string) *DescribeAccountsResponseBodyAccountsDatabasePrivileges {
	s.DBName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDatabasePrivileges) Validate() error {
	return dara.Validate(s)
}
