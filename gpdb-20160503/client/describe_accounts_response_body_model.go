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
	// The queried database accounts.
	Accounts *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7565770E-7C45-462D-BA4A-8A5396F2CAD1
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
	DBInstanceAccount []*DescribeAccountsResponseBodyAccountsDBInstanceAccount `json:"DBInstanceAccount,omitempty" xml:"DBInstanceAccount,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccounts) GetDBInstanceAccount() []*DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	return s.DBInstanceAccount
}

func (s *DescribeAccountsResponseBodyAccounts) SetDBInstanceAccount(v []*DescribeAccountsResponseBodyAccountsDBInstanceAccount) *DescribeAccountsResponseBodyAccounts {
	s.DBInstanceAccount = v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) Validate() error {
	return dara.Validate(s)
}

type DescribeAccountsResponseBodyAccountsDBInstanceAccount struct {
	// The description of the account.
	//
	// example:
	//
	// testuser
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// testuser
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The state of the account.
	//
	// 	- **0**: The account is being created.
	//
	// 	- **1**: The account is in use.
	//
	// 	- **3**: The account is being deleted.
	//
	// example:
	//
	// 1
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The type of the database account. Valid values: Super and Normal. Super indicates a privileged account and Normal indicates a standard account.
	//
	// example:
	//
	// Super
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccount) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountType(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetDBInstanceId(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) Validate() error {
	return dara.Validate(s)
}
