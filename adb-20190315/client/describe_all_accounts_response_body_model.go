// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountList(v []*DescribeAllAccountsResponseBodyAccountList) *DescribeAllAccountsResponseBody
	GetAccountList() []*DescribeAllAccountsResponseBodyAccountList
	SetRequestId(v string) *DescribeAllAccountsResponseBody
	GetRequestId() *string
}

type DescribeAllAccountsResponseBody struct {
	// The queried database accounts.
	AccountList []*DescribeAllAccountsResponseBodyAccountList `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAllAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsResponseBody) GetAccountList() []*DescribeAllAccountsResponseBodyAccountList {
	return s.AccountList
}

func (s *DescribeAllAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllAccountsResponseBody) SetAccountList(v []*DescribeAllAccountsResponseBodyAccountList) *DescribeAllAccountsResponseBody {
	s.AccountList = v
	return s
}

func (s *DescribeAllAccountsResponseBody) SetRequestId(v string) *DescribeAllAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllAccountsResponseBody) Validate() error {
	if s.AccountList != nil {
		for _, item := range s.AccountList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAllAccountsResponseBodyAccountList struct {
	// The name of the database account.
	//
	// example:
	//
	// rdsdt_dts_adb
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAllAccountsResponseBodyAccountList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllAccountsResponseBodyAccountList) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsResponseBodyAccountList) GetUser() *string {
	return s.User
}

func (s *DescribeAllAccountsResponseBodyAccountList) SetUser(v string) *DescribeAllAccountsResponseBodyAccountList {
	s.User = &v
	return s
}

func (s *DescribeAllAccountsResponseBodyAccountList) Validate() error {
	return dara.Validate(s)
}
