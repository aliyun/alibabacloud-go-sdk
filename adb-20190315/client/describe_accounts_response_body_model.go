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
	// 64E37E6F-C363-41F3-867A-70EF5DC60EA4
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
	// The description of the database account.
	//
	// example:
	//
	// C@test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// test1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The state of the database account. Valid values:
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
	AccountType *string                                               `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Tags        *DescribeAccountsResponseBodyAccountListDBAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
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

func (s *DescribeAccountsResponseBodyAccountListDBAccount) GetTags() *DescribeAccountsResponseBodyAccountListDBAccountTags {
	return s.Tags
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

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetTags(v *DescribeAccountsResponseBodyAccountListDBAccountTags) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.Tags = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccountListDBAccountTags struct {
	Tag []*DescribeAccountsResponseBodyAccountListDBAccountTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccountListDBAccountTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountListDBAccountTags) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountListDBAccountTags) GetTag() []*DescribeAccountsResponseBodyAccountListDBAccountTagsTag {
	return s.Tag
}

func (s *DescribeAccountsResponseBodyAccountListDBAccountTags) SetTag(v []*DescribeAccountsResponseBodyAccountListDBAccountTagsTag) *DescribeAccountsResponseBodyAccountListDBAccountTags {
	s.Tag = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccountTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccountListDBAccountTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountListDBAccountTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountListDBAccountTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountListDBAccountTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAccountsResponseBodyAccountListDBAccountTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAccountsResponseBodyAccountListDBAccountTagsTag) SetKey(v string) *DescribeAccountsResponseBodyAccountListDBAccountTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccountTagsTag) SetValue(v string) *DescribeAccountsResponseBodyAccountListDBAccountTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccountTagsTag) Validate() error {
	return dara.Validate(s)
}
