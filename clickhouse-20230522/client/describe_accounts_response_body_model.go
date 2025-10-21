// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeAccountsResponseBodyData) *DescribeAccountsResponseBody
	GetData() *DescribeAccountsResponseBodyData
	SetRequestId(v string) *DescribeAccountsResponseBody
	GetRequestId() *string
}

type DescribeAccountsResponseBody struct {
	// The result returned.
	Data *DescribeAccountsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) GetData() *DescribeAccountsResponseBodyData {
	return s.Data
}

func (s *DescribeAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountsResponseBody) SetData(v *DescribeAccountsResponseBodyData) *DescribeAccountsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountsResponseBodyData struct {
	// The database accounts.
	Accounts []*DescribeAccountsResponseBodyDataAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccountsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyData) GetAccounts() []*DescribeAccountsResponseBodyDataAccounts {
	return s.Accounts
}

func (s *DescribeAccountsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccountsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccountsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAccountsResponseBodyData) SetAccounts(v []*DescribeAccountsResponseBodyDataAccounts) *DescribeAccountsResponseBodyData {
	s.Accounts = v
	return s
}

func (s *DescribeAccountsResponseBodyData) SetPageNumber(v int32) *DescribeAccountsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsResponseBodyData) SetPageSize(v int32) *DescribeAccountsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountsResponseBodyData) SetTotalCount(v int32) *DescribeAccountsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccountsResponseBodyData) Validate() error {
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

type DescribeAccountsResponseBodyDataAccounts struct {
	// The username of the database account.
	//
	// example:
	//
	// test
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The type of the database account. Valid values:
	//
	// 	- **1**: standard account
	//
	// 	- **6**: privileged account
	//
	// example:
	//
	// NormalAccount
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The description.
	//
	// example:
	//
	// Used for test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The state of the database account. Valid values:
	//
	// 	- **0**: The database account is being created.
	//
	// 	- **1**: The database account is in use.
	//
	// 	- **3**: The database account is being deleted.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAccountsResponseBodyDataAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyDataAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyDataAccounts) GetAccount() *string {
	return s.Account
}

func (s *DescribeAccountsResponseBodyDataAccounts) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeAccountsResponseBodyDataAccounts) GetDescription() *string {
	return s.Description
}

func (s *DescribeAccountsResponseBodyDataAccounts) GetStatus() *string {
	return s.Status
}

func (s *DescribeAccountsResponseBodyDataAccounts) SetAccount(v string) *DescribeAccountsResponseBodyDataAccounts {
	s.Account = &v
	return s
}

func (s *DescribeAccountsResponseBodyDataAccounts) SetAccountType(v string) *DescribeAccountsResponseBodyDataAccounts {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseBodyDataAccounts) SetDescription(v string) *DescribeAccountsResponseBodyDataAccounts {
	s.Description = &v
	return s
}

func (s *DescribeAccountsResponseBodyDataAccounts) SetStatus(v string) *DescribeAccountsResponseBodyDataAccounts {
	s.Status = &v
	return s
}

func (s *DescribeAccountsResponseBodyDataAccounts) Validate() error {
	return dara.Validate(s)
}
