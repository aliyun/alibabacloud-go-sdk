// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsForParentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v *ListAccountsForParentResponseBodyAccounts) *ListAccountsForParentResponseBody
	GetAccounts() *ListAccountsForParentResponseBodyAccounts
	SetPageNumber(v int32) *ListAccountsForParentResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAccountsForParentResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAccountsForParentResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAccountsForParentResponseBody
	GetTotalCount() *int32
}

type ListAccountsForParentResponseBody struct {
	// The information of the members.
	Accounts *ListAccountsForParentResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
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

func (s ListAccountsForParentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsForParentResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBody) GetAccounts() *ListAccountsForParentResponseBodyAccounts {
	return s.Accounts
}

func (s *ListAccountsForParentResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAccountsForParentResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAccountsForParentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccountsForParentResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAccountsForParentResponseBody) SetAccounts(v *ListAccountsForParentResponseBodyAccounts) *ListAccountsForParentResponseBody {
	s.Accounts = v
	return s
}

func (s *ListAccountsForParentResponseBody) SetPageNumber(v int32) *ListAccountsForParentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsForParentResponseBody) SetPageSize(v int32) *ListAccountsForParentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAccountsForParentResponseBody) SetRequestId(v string) *ListAccountsForParentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsForParentResponseBody) SetTotalCount(v int32) *ListAccountsForParentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAccountsForParentResponseBody) Validate() error {
	if s.Accounts != nil {
		if err := s.Accounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAccountsForParentResponseBodyAccounts struct {
	Account []*ListAccountsForParentResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s ListAccountsForParentResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccounts) GetAccount() []*ListAccountsForParentResponseBodyAccountsAccount {
	return s.Account
}

func (s *ListAccountsForParentResponseBodyAccounts) SetAccount(v []*ListAccountsForParentResponseBodyAccountsAccount) *ListAccountsForParentResponseBodyAccounts {
	s.Account = v
	return s
}

func (s *ListAccountsForParentResponseBodyAccounts) Validate() error {
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

type ListAccountsForParentResponseBodyAccountsAccount struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 184311716100****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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
	// The way in which the member joins the resource directory.
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
	// The status of the member. Valid values:
	//
	// 	- CreateSuccess: The member is created.
	//
	// 	- PromoteVerifying: The upgrade of the member is being confirmed.
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
	// The tags that are added to the member.
	Tags *ListAccountsForParentResponseBodyAccountsAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The type of the member.
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

func (s ListAccountsForParentResponseBodyAccountsAccount) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) GetAccountId() *string {
	return s.AccountId
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) GetFolderId() *string {
	return s.FolderId
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) GetJoinMethod() *string {
	return s.JoinMethod
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) GetJoinTime() *string {
	return s.JoinTime
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) GetStatus() *string {
	return s.Status
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) GetTags() *ListAccountsForParentResponseBodyAccountsAccountTags {
	return s.Tags
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) GetType() *string {
	return s.Type
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetAccountId(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetDisplayName(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetFolderId(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.FolderId = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetJoinMethod(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetJoinTime(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.JoinTime = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetModifyTime(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.ModifyTime = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetResourceDirectoryId(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetStatus(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.Status = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetTags(v *ListAccountsForParentResponseBodyAccountsAccountTags) *ListAccountsForParentResponseBodyAccountsAccount {
	s.Tags = v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetType(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.Type = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAccountsForParentResponseBodyAccountsAccountTags struct {
	Tag []*ListAccountsForParentResponseBodyAccountsAccountTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsForParentResponseBodyAccountsAccountTags) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccountsAccountTags) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTags) GetTag() []*ListAccountsForParentResponseBodyAccountsAccountTagsTag {
	return s.Tag
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTags) SetTag(v []*ListAccountsForParentResponseBodyAccountsAccountTagsTag) *ListAccountsForParentResponseBodyAccountsAccountTags {
	s.Tag = v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTags) Validate() error {
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

type ListAccountsForParentResponseBodyAccountsAccountTagsTag struct {
	// A tag key.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsForParentResponseBodyAccountsAccountTagsTag) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccountsAccountTagsTag) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTagsTag) GetKey() *string {
	return s.Key
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTagsTag) GetValue() *string {
	return s.Value
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTagsTag) SetKey(v string) *ListAccountsForParentResponseBodyAccountsAccountTagsTag {
	s.Key = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTagsTag) SetValue(v string) *ListAccountsForParentResponseBodyAccountsAccountTagsTag {
	s.Value = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTagsTag) Validate() error {
	return dara.Validate(s)
}
