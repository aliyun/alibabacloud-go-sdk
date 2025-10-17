// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v []*DescribeAccountsZonalResponseBodyAccounts) *DescribeAccountsZonalResponseBody
	GetAccounts() []*DescribeAccountsZonalResponseBodyAccounts
	SetMaxResults(v int32) *DescribeAccountsZonalResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAccountsZonalResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeAccountsZonalResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeAccountsZonalResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeAccountsZonalResponseBody
	GetRequestId() *string
}

type DescribeAccountsZonalResponseBody struct {
	Accounts []*DescribeAccountsZonalResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 155462B9-205F-4FFC-BB43-4855FE******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountsZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsZonalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsZonalResponseBody) GetAccounts() []*DescribeAccountsZonalResponseBodyAccounts {
	return s.Accounts
}

func (s *DescribeAccountsZonalResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAccountsZonalResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAccountsZonalResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccountsZonalResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeAccountsZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountsZonalResponseBody) SetAccounts(v []*DescribeAccountsZonalResponseBodyAccounts) *DescribeAccountsZonalResponseBody {
	s.Accounts = v
	return s
}

func (s *DescribeAccountsZonalResponseBody) SetMaxResults(v int32) *DescribeAccountsZonalResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeAccountsZonalResponseBody) SetNextToken(v string) *DescribeAccountsZonalResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAccountsZonalResponseBody) SetPageNumber(v int32) *DescribeAccountsZonalResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsZonalResponseBody) SetPageRecordCount(v int32) *DescribeAccountsZonalResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAccountsZonalResponseBody) SetRequestId(v string) *DescribeAccountsZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsZonalResponseBody) Validate() error {
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

type DescribeAccountsZonalResponseBodyAccounts struct {
	// example:
	//
	// test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// example:
	//
	// UnLock
	AccountLockState *string `json:"AccountLockState,omitempty" xml:"AccountLockState,omitempty"`
	// example:
	//
	// test_acc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// undefined
	AccountPasswordValidTime *string `json:"AccountPasswordValidTime,omitempty" xml:"AccountPasswordValidTime,omitempty"`
	// example:
	//
	// Available
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// example:
	//
	// Normal
	AccountType        *string                                                        `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	DatabasePrivileges []*DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges `json:"DatabasePrivileges,omitempty" xml:"DatabasePrivileges,omitempty" type:"Repeated"`
}

func (s DescribeAccountsZonalResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsZonalResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsZonalResponseBodyAccounts) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *DescribeAccountsZonalResponseBodyAccounts) GetAccountLockState() *string {
	return s.AccountLockState
}

func (s *DescribeAccountsZonalResponseBodyAccounts) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountsZonalResponseBodyAccounts) GetAccountPasswordValidTime() *string {
	return s.AccountPasswordValidTime
}

func (s *DescribeAccountsZonalResponseBodyAccounts) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeAccountsZonalResponseBodyAccounts) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeAccountsZonalResponseBodyAccounts) GetDatabasePrivileges() []*DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges {
	return s.DatabasePrivileges
}

func (s *DescribeAccountsZonalResponseBodyAccounts) SetAccountDescription(v string) *DescribeAccountsZonalResponseBodyAccounts {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsZonalResponseBodyAccounts) SetAccountLockState(v string) *DescribeAccountsZonalResponseBodyAccounts {
	s.AccountLockState = &v
	return s
}

func (s *DescribeAccountsZonalResponseBodyAccounts) SetAccountName(v string) *DescribeAccountsZonalResponseBodyAccounts {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsZonalResponseBodyAccounts) SetAccountPasswordValidTime(v string) *DescribeAccountsZonalResponseBodyAccounts {
	s.AccountPasswordValidTime = &v
	return s
}

func (s *DescribeAccountsZonalResponseBodyAccounts) SetAccountStatus(v string) *DescribeAccountsZonalResponseBodyAccounts {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsZonalResponseBodyAccounts) SetAccountType(v string) *DescribeAccountsZonalResponseBodyAccounts {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsZonalResponseBodyAccounts) SetDatabasePrivileges(v []*DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges) *DescribeAccountsZonalResponseBodyAccounts {
	s.DatabasePrivileges = v
	return s
}

func (s *DescribeAccountsZonalResponseBodyAccounts) Validate() error {
	if s.DatabasePrivileges != nil {
		for _, item := range s.DatabasePrivileges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges struct {
	// example:
	//
	// ReadOnly
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// example:
	//
	// DBtest
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges) GoString() string {
	return s.String()
}

func (s *DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges) SetAccountPrivilege(v string) *DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges) SetDBName(v string) *DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges {
	s.DBName = &v
	return s
}

func (s *DescribeAccountsZonalResponseBodyAccountsDatabasePrivileges) Validate() error {
	return dara.Validate(s)
}
