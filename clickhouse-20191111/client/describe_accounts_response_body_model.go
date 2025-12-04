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
	SetPageNumber(v int32) *DescribeAccountsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccountsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAccountsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAccountsResponseBody
	GetTotalCount() *int32
}

type DescribeAccountsResponseBody struct {
	// The database accounts.
	Accounts *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeAccountsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccountsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAccountsResponseBody) SetAccounts(v *DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *DescribeAccountsResponseBody) SetPageNumber(v int32) *DescribeAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetPageSize(v int32) *DescribeAccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetTotalCount(v int32) *DescribeAccountsResponseBody {
	s.TotalCount = &v
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
	// The description of the database account.
	//
	// example:
	//
	// test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The state of the database account. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Available**
	//
	// 	- **AuthorityModifying**
	//
	// 	- **Deleting**
	//
	// >  Only XML-configured database accounts can be in the **AuthorityModifying*	- state.
	//
	// example:
	//
	// Creating
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The type of the database account. Valid values:
	//
	// 	- **Super**: privileged account
	//
	// 	- **Normal**: standard account
	//
	// example:
	//
	// Super
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The configuration mode of the database account.
	//
	// example:
	//
	// SQL
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
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

func (s *DescribeAccountsResponseBodyAccountsAccount) GetConfigType() *string {
	return s.ConfigType
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

func (s *DescribeAccountsResponseBodyAccountsAccount) SetConfigType(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.ConfigType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) Validate() error {
	return dara.Validate(s)
}
