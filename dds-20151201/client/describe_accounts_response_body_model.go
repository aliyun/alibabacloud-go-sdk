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
	// The details of the account.
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
	// The description of the account.
	//
	// example:
	//
	// Admin
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// root
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The status of the account. Valid values:
	//
	// 	- **Unavailable**
	//
	// 	- **Available**
	//
	// example:
	//
	// Available
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The role of the account. Valid values:
	//
	// 	- **db**: shard node
	//
	// 	- **cs**: Configserver node
	//
	// 	- **mongos**: mongos node
	//
	// 	- **logic**: sharded cluster instance
	//
	// 	- **normal**: replica set instance
	//
	// example:
	//
	// mongos
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The ID of the instance to which the account belongs.
	//
	// example:
	//
	// dds-bp1fd530f271****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

func (s *DescribeAccountsResponseBodyAccountsAccount) Validate() error {
	return dara.Validate(s)
}
