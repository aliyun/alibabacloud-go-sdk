// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v *ListAuthorizedAccountsResponseBodyAccounts) *ListAuthorizedAccountsResponseBody
	GetAccounts() *ListAuthorizedAccountsResponseBodyAccounts
	SetPageNumber(v int32) *ListAuthorizedAccountsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedAccountsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAuthorizedAccountsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAuthorizedAccountsResponseBody
	GetTotalCount() *int32
}

type ListAuthorizedAccountsResponseBody struct {
	// The information about the member.
	Accounts *ListAuthorizedAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
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
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthorizedAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAccountsResponseBody) GetAccounts() *ListAuthorizedAccountsResponseBodyAccounts {
	return s.Accounts
}

func (s *ListAuthorizedAccountsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedAccountsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedAccountsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAuthorizedAccountsResponseBody) SetAccounts(v *ListAuthorizedAccountsResponseBodyAccounts) *ListAuthorizedAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *ListAuthorizedAccountsResponseBody) SetPageNumber(v int32) *ListAuthorizedAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBody) SetPageSize(v int32) *ListAuthorizedAccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBody) SetRequestId(v string) *ListAuthorizedAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBody) SetTotalCount(v int32) *ListAuthorizedAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBody) Validate() error {
	if s.Accounts != nil {
		if err := s.Accounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAuthorizedAccountsResponseBodyAccounts struct {
	Account []*ListAuthorizedAccountsResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s ListAuthorizedAccountsResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAccountsResponseBodyAccounts) GetAccount() []*ListAuthorizedAccountsResponseBodyAccountsAccount {
	return s.Account
}

func (s *ListAuthorizedAccountsResponseBodyAccounts) SetAccount(v []*ListAuthorizedAccountsResponseBodyAccountsAccount) *ListAuthorizedAccountsResponseBodyAccounts {
	s.Account = v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccounts) Validate() error {
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

type ListAuthorizedAccountsResponseBodyAccountsAccount struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 184311716100****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member.
	//
	// example:
	//
	// someone@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// admin
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-bVaRIG****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
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
	// The time when the member joined the resource directory. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-k4****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The RDPath of the member.
	ResourceDirectoryPath *string `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
	// The status of the member. Valid values:
	//
	// 	- CreateSuccess: The member is created.
	//
	// 	- PromoteVerifying: The upgrade of the member is under confirmation.
	//
	// 	- PromoteFailed: The upgrade of the member fails.
	//
	// 	- PromoteExpired: The upgrade of the member expires.
	//
	// 	- PromoteCancelled: The upgrade of the member is canceled.
	//
	// 	- PromoteSuccess: The member is upgraded.
	//
	// 	- InviteSuccess: The member accepts the invitation.
	//
	// example:
	//
	// CreateSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member. Valid values:
	//
	// 	- CloudAccount: cloud account
	//
	// 	- ResourceAccount: resource account
	//
	// example:
	//
	// ResourceAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAuthorizedAccountsResponseBodyAccountsAccount) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAccountsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) GetAccountId() *string {
	return s.AccountId
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) GetFolderId() *string {
	return s.FolderId
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) GetJoinMethod() *string {
	return s.JoinMethod
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) GetJoinTime() *string {
	return s.JoinTime
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) GetResourceDirectoryPath() *string {
	return s.ResourceDirectoryPath
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) GetStatus() *string {
	return s.Status
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) GetType() *string {
	return s.Type
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) SetAccountId(v string) *ListAuthorizedAccountsResponseBodyAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) SetAccountName(v string) *ListAuthorizedAccountsResponseBodyAccountsAccount {
	s.AccountName = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) SetDisplayName(v string) *ListAuthorizedAccountsResponseBodyAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) SetFolderId(v string) *ListAuthorizedAccountsResponseBodyAccountsAccount {
	s.FolderId = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) SetJoinMethod(v string) *ListAuthorizedAccountsResponseBodyAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) SetJoinTime(v string) *ListAuthorizedAccountsResponseBodyAccountsAccount {
	s.JoinTime = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) SetModifyTime(v string) *ListAuthorizedAccountsResponseBodyAccountsAccount {
	s.ModifyTime = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) SetResourceDirectoryId(v string) *ListAuthorizedAccountsResponseBodyAccountsAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) SetResourceDirectoryPath(v string) *ListAuthorizedAccountsResponseBodyAccountsAccount {
	s.ResourceDirectoryPath = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) SetStatus(v string) *ListAuthorizedAccountsResponseBodyAccountsAccount {
	s.Status = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) SetType(v string) *ListAuthorizedAccountsResponseBodyAccountsAccount {
	s.Type = &v
	return s
}

func (s *ListAuthorizedAccountsResponseBodyAccountsAccount) Validate() error {
	return dara.Validate(s)
}
