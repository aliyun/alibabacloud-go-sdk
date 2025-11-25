// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRdAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v []*DescribeInstanceRdAccountsResponseBodyAccounts) *DescribeInstanceRdAccountsResponseBody
	GetAccounts() []*DescribeInstanceRdAccountsResponseBodyAccounts
	SetRequestId(v string) *DescribeInstanceRdAccountsResponseBody
	GetRequestId() *string
}

type DescribeInstanceRdAccountsResponseBody struct {
	Accounts []*DescribeInstanceRdAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// example:
	//
	// 06D1EC07-C9EB-58AC-A750-C87C9A0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceRdAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRdAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRdAccountsResponseBody) GetAccounts() []*DescribeInstanceRdAccountsResponseBodyAccounts {
	return s.Accounts
}

func (s *DescribeInstanceRdAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceRdAccountsResponseBody) SetAccounts(v []*DescribeInstanceRdAccountsResponseBodyAccounts) *DescribeInstanceRdAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *DescribeInstanceRdAccountsResponseBody) SetRequestId(v string) *DescribeInstanceRdAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceRdAccountsResponseBody) Validate() error {
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceRdAccountsResponseBodyAccounts struct {
	// example:
	//
	// 171054237268****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
}

func (s DescribeInstanceRdAccountsResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRdAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRdAccountsResponseBodyAccounts) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeInstanceRdAccountsResponseBodyAccounts) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeInstanceRdAccountsResponseBodyAccounts) SetAccountId(v string) *DescribeInstanceRdAccountsResponseBodyAccounts {
	s.AccountId = &v
	return s
}

func (s *DescribeInstanceRdAccountsResponseBodyAccounts) SetDisplayName(v string) *DescribeInstanceRdAccountsResponseBodyAccounts {
	s.DisplayName = &v
	return s
}

func (s *DescribeInstanceRdAccountsResponseBodyAccounts) Validate() error {
	return dara.Validate(s)
}
