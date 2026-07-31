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
	if s.AccountList != nil {
		if err := s.AccountList.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.DBAccount != nil {
		for _, item := range s.DBAccount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccountListDBAccount struct {
	AccountDescription *string                                                      `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string                                                      `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountStatus      *string                                                      `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AccountType        *string                                                      `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Engine             *string                                                      `json:"Engine,omitempty" xml:"Engine,omitempty"`
	RamUserList        *DescribeAccountsResponseBodyAccountListDBAccountRamUserList `json:"RamUserList,omitempty" xml:"RamUserList,omitempty" type:"Struct"`
	RamUsers           *string                                                      `json:"RamUsers,omitempty" xml:"RamUsers,omitempty"`
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

func (s *DescribeAccountsResponseBodyAccountListDBAccount) GetRamUserList() *DescribeAccountsResponseBodyAccountListDBAccountRamUserList {
	return s.RamUserList
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

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetRamUserList(v *DescribeAccountsResponseBodyAccountListDBAccountRamUserList) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.RamUserList = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetRamUsers(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.RamUsers = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) Validate() error {
	if s.RamUserList != nil {
		if err := s.RamUserList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccountListDBAccountRamUserList struct {
	RamUserList []*string `json:"RamUserList,omitempty" xml:"RamUserList,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccountListDBAccountRamUserList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountListDBAccountRamUserList) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountListDBAccountRamUserList) GetRamUserList() []*string {
	return s.RamUserList
}

func (s *DescribeAccountsResponseBodyAccountListDBAccountRamUserList) SetRamUserList(v []*string) *DescribeAccountsResponseBodyAccountListDBAccountRamUserList {
	s.RamUserList = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccountRamUserList) Validate() error {
	return dara.Validate(s)
}
