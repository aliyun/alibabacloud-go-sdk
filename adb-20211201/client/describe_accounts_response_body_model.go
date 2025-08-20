// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountList(v *DescribeAccountsResponseBodyAccountList) *DescribeAccountsResponseBody
	GetAccountList() *DescribeAccountsResponseBodyAccountList
	SetRequestId(v string) *DescribeAccountsResponseBody
	GetRequestId() *string
}

type DescribeAccountsResponseBody struct {
	// The queried database accounts.
	AccountList *DescribeAccountsResponseBodyAccountList `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9CCFAAB4-97B7-5800-B9F2-685EB596E3EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) GetAccountList() *DescribeAccountsResponseBodyAccountList {
	return s.AccountList
}

func (s *DescribeAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountsResponseBody) SetAccountList(v *DescribeAccountsResponseBodyAccountList) *DescribeAccountsResponseBody {
	s.AccountList = v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAccountsResponseBodyAccountList struct {
	DBAccount []*DescribeAccountsResponseBodyAccountListDBAccount `json:"DBAccount,omitempty" xml:"DBAccount,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccountList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountList) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountList) GetDBAccount() []*DescribeAccountsResponseBodyAccountListDBAccount {
	return s.DBAccount
}

func (s *DescribeAccountsResponseBodyAccountList) SetDBAccount(v []*DescribeAccountsResponseBodyAccountListDBAccount) *DescribeAccountsResponseBodyAccountList {
	s.DBAccount = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountList) Validate() error {
	return dara.Validate(s)
}

type DescribeAccountsResponseBodyAccountListDBAccount struct {
	// The description of the database account.
	//
	// example:
	//
	// test_accout_des
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// test_accout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The status of the database account. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Available**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Available
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The type of the database account. Valid values:
	//
	// 	- **Normal**: standard account.
	//
	// 	- **Super**: privileged account.
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The database engine of the cluster. Valid values:
	//
	// 	- **AnalyticDB**: the AnalyticDB for MySQL engine.
	//
	// 	- **Clickhouse**: the wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The ID of the Resource Access Management (RAM) user.
	//
	// example:
	//
	// 1958134230****
	RamUsers *string `json:"RamUsers,omitempty" xml:"RamUsers,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountListDBAccount) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountListDBAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) GetRamUsers() *string {
	return s.RamUsers
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountType(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetEngine(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.Engine = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetRamUsers(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.RamUsers = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) Validate() error {
	return dara.Validate(s)
}
