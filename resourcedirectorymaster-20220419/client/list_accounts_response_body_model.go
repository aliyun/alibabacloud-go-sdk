// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v *ListAccountsResponseBodyAccounts) *ListAccountsResponseBody
	GetAccounts() *ListAccountsResponseBodyAccounts
	SetNextToken(v string) *ListAccountsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListAccountsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAccountsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAccountsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAccountsResponseBody
	GetTotalCount() *int32
}

type ListAccountsResponseBody struct {
	// The information about the members.
	Accounts *ListAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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

func (s ListAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBody) GetAccounts() *ListAccountsResponseBodyAccounts {
	return s.Accounts
}

func (s *ListAccountsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccountsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAccountsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccountsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAccountsResponseBody) SetAccounts(v *ListAccountsResponseBodyAccounts) *ListAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *ListAccountsResponseBody) SetNextToken(v string) *ListAccountsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccountsResponseBody) SetPageNumber(v int32) *ListAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsResponseBody) SetPageSize(v int32) *ListAccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAccountsResponseBody) SetRequestId(v string) *ListAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsResponseBody) SetTotalCount(v int32) *ListAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAccountsResponseBodyAccounts struct {
	Account []*ListAccountsResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s ListAccountsResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccounts) GetAccount() []*ListAccountsResponseBodyAccountsAccount {
	return s.Account
}

func (s *ListAccountsResponseBodyAccounts) SetAccount(v []*ListAccountsResponseBodyAccountsAccount) *ListAccountsResponseBodyAccounts {
	s.Account = v
	return s
}

func (s *ListAccountsResponseBodyAccounts) Validate() error {
	return dara.Validate(s)
}

type ListAccountsResponseBodyAccountsAccount struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The Alibaba Cloud account name of the member.
	//
	// example:
	//
	// oxy01414357@alibaba-inc.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The deletion status of the member. Valid values:
	//
	// 	- Checking: A deletion check is being performed for the member.
	//
	// 	- Deleting: The member is being deleted.
	//
	// 	- CheckFailed: The deletion check for the member fails.
	//
	// 	- DeleteFailed: The member fails to be deleted.
	//
	// >  If deletion is not performed for the member, the value of this parameter is empty.
	//
	// example:
	//
	// Checking
	DeletionStatus *string `json:"DeletionStatus,omitempty" xml:"DeletionStatus,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-QRzuim****
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
	// 2021-01-18T08:01:50.522Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-18T08:04:37.668Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-3G****
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
	// The tags that are added to the member.
	Tags *ListAccountsResponseBodyAccountsAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
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

func (s ListAccountsResponseBodyAccountsAccount) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccountsAccount) GetAccountId() *string {
	return s.AccountId
}

func (s *ListAccountsResponseBodyAccountsAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *ListAccountsResponseBodyAccountsAccount) GetDeletionStatus() *string {
	return s.DeletionStatus
}

func (s *ListAccountsResponseBodyAccountsAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAccountsResponseBodyAccountsAccount) GetFolderId() *string {
	return s.FolderId
}

func (s *ListAccountsResponseBodyAccountsAccount) GetJoinMethod() *string {
	return s.JoinMethod
}

func (s *ListAccountsResponseBodyAccountsAccount) GetJoinTime() *string {
	return s.JoinTime
}

func (s *ListAccountsResponseBodyAccountsAccount) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListAccountsResponseBodyAccountsAccount) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *ListAccountsResponseBodyAccountsAccount) GetResourceDirectoryPath() *string {
	return s.ResourceDirectoryPath
}

func (s *ListAccountsResponseBodyAccountsAccount) GetStatus() *string {
	return s.Status
}

func (s *ListAccountsResponseBodyAccountsAccount) GetTags() *ListAccountsResponseBodyAccountsAccountTags {
	return s.Tags
}

func (s *ListAccountsResponseBodyAccountsAccount) GetType() *string {
	return s.Type
}

func (s *ListAccountsResponseBodyAccountsAccount) SetAccountId(v string) *ListAccountsResponseBodyAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetAccountName(v string) *ListAccountsResponseBodyAccountsAccount {
	s.AccountName = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetDeletionStatus(v string) *ListAccountsResponseBodyAccountsAccount {
	s.DeletionStatus = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetDisplayName(v string) *ListAccountsResponseBodyAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetFolderId(v string) *ListAccountsResponseBodyAccountsAccount {
	s.FolderId = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetJoinMethod(v string) *ListAccountsResponseBodyAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetJoinTime(v string) *ListAccountsResponseBodyAccountsAccount {
	s.JoinTime = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetModifyTime(v string) *ListAccountsResponseBodyAccountsAccount {
	s.ModifyTime = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetResourceDirectoryId(v string) *ListAccountsResponseBodyAccountsAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetResourceDirectoryPath(v string) *ListAccountsResponseBodyAccountsAccount {
	s.ResourceDirectoryPath = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetStatus(v string) *ListAccountsResponseBodyAccountsAccount {
	s.Status = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetTags(v *ListAccountsResponseBodyAccountsAccountTags) *ListAccountsResponseBodyAccountsAccount {
	s.Tags = v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetType(v string) *ListAccountsResponseBodyAccountsAccount {
	s.Type = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) Validate() error {
	return dara.Validate(s)
}

type ListAccountsResponseBodyAccountsAccountTags struct {
	Tag []*ListAccountsResponseBodyAccountsAccountTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsResponseBodyAccountsAccountTags) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsResponseBodyAccountsAccountTags) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccountsAccountTags) GetTag() []*ListAccountsResponseBodyAccountsAccountTagsTag {
	return s.Tag
}

func (s *ListAccountsResponseBodyAccountsAccountTags) SetTag(v []*ListAccountsResponseBodyAccountsAccountTagsTag) *ListAccountsResponseBodyAccountsAccountTags {
	s.Tag = v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccountTags) Validate() error {
	return dara.Validate(s)
}

type ListAccountsResponseBodyAccountsAccountTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsResponseBodyAccountsAccountTagsTag) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsResponseBodyAccountsAccountTagsTag) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccountsAccountTagsTag) GetKey() *string {
	return s.Key
}

func (s *ListAccountsResponseBodyAccountsAccountTagsTag) GetValue() *string {
	return s.Value
}

func (s *ListAccountsResponseBodyAccountsAccountTagsTag) SetKey(v string) *ListAccountsResponseBodyAccountsAccountTagsTag {
	s.Key = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccountTagsTag) SetValue(v string) *ListAccountsResponseBodyAccountsAccountTagsTag {
	s.Value = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccountTagsTag) Validate() error {
	return dara.Validate(s)
}
