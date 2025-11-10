// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDelegatedAdministratorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v *ListDelegatedAdministratorsResponseBodyAccounts) *ListDelegatedAdministratorsResponseBody
	GetAccounts() *ListDelegatedAdministratorsResponseBodyAccounts
	SetPageNumber(v int64) *ListDelegatedAdministratorsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDelegatedAdministratorsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListDelegatedAdministratorsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDelegatedAdministratorsResponseBody
	GetTotalCount() *int64
}

type ListDelegatedAdministratorsResponseBody struct {
	// The information about the delegated administrator accounts.
	Accounts *ListDelegatedAdministratorsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 009C3A02-7D4B-416C-9CE7-548C91508F1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDelegatedAdministratorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDelegatedAdministratorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponseBody) GetAccounts() *ListDelegatedAdministratorsResponseBodyAccounts {
	return s.Accounts
}

func (s *ListDelegatedAdministratorsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDelegatedAdministratorsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDelegatedAdministratorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDelegatedAdministratorsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDelegatedAdministratorsResponseBody) SetAccounts(v *ListDelegatedAdministratorsResponseBodyAccounts) *ListDelegatedAdministratorsResponseBody {
	s.Accounts = v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetPageNumber(v int64) *ListDelegatedAdministratorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetPageSize(v int64) *ListDelegatedAdministratorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetRequestId(v string) *ListDelegatedAdministratorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetTotalCount(v int64) *ListDelegatedAdministratorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) Validate() error {
	if s.Accounts != nil {
		if err := s.Accounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDelegatedAdministratorsResponseBodyAccounts struct {
	Account []*ListDelegatedAdministratorsResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s ListDelegatedAdministratorsResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListDelegatedAdministratorsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponseBodyAccounts) GetAccount() []*ListDelegatedAdministratorsResponseBodyAccountsAccount {
	return s.Account
}

func (s *ListDelegatedAdministratorsResponseBodyAccounts) SetAccount(v []*ListDelegatedAdministratorsResponseBodyAccountsAccount) *ListDelegatedAdministratorsResponseBodyAccounts {
	s.Account = v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccounts) Validate() error {
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

type ListDelegatedAdministratorsResponseBodyAccountsAccount struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 138660628348****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the member was specified as a delegated administrator account.
	//
	// example:
	//
	// 1616652684164
	DelegationEnabledTime *string `json:"DelegationEnabledTime,omitempty" xml:"DelegationEnabledTime,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// abc
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The way in which the member joins the resource directory. Valid values:
	//
	// 	- invited: The member is invited to join the resource directory.
	//
	// 	- created: The member is directly created in the resource directory.
	//
	// example:
	//
	// created
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The identifier of the trusted service.
	//
	// example:
	//
	// cloudfw.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s ListDelegatedAdministratorsResponseBodyAccountsAccount) String() string {
	return dara.Prettify(s)
}

func (s ListDelegatedAdministratorsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) GetAccountId() *string {
	return s.AccountId
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) GetDelegationEnabledTime() *string {
	return s.DelegationEnabledTime
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) GetJoinMethod() *string {
	return s.JoinMethod
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) GetServicePrincipal() *string {
	return s.ServicePrincipal
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetAccountId(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetDelegationEnabledTime(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.DelegationEnabledTime = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetDisplayName(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetJoinMethod(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetServicePrincipal(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.ServicePrincipal = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) Validate() error {
	return dara.Validate(s)
}
