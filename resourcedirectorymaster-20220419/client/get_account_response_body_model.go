// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v *GetAccountResponseBodyAccount) *GetAccountResponseBody
	GetAccount() *GetAccountResponseBodyAccount
	SetRequestId(v string) *GetAccountResponseBody
	GetRequestId() *string
}

type GetAccountResponseBody struct {
	// The information about the member.
	Account *GetAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountResponseBody) GetAccount() *GetAccountResponseBodyAccount {
	return s.Account
}

func (s *GetAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountResponseBody) SetAccount(v *GetAccountResponseBodyAccount) *GetAccountResponseBody {
	s.Account = v
	return s
}

func (s *GetAccountResponseBody) SetRequestId(v string) *GetAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAccountResponseBodyAccount struct {
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
	// someone@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// admin
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The status of the modification for the email address bound to the member. Valid values:
	//
	// 	- If the value of this parameter is empty, no modification is performed for the email address.
	//
	// 	- WAIT_MODIFY: The modification is being performed.
	//
	// 	- CANCELLED: The modification is canceled.
	//
	// 	- EXPIRED: The modification expires.
	//
	// example:
	//
	// WAIT_MODIFY
	EmailStatus *string `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-bVaRIG****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// Indicates whether a mobile phone number is bound to the member for security purposes. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	HasSecureMobilePhone *bool `json:"HasSecureMobilePhone,omitempty" xml:"HasSecureMobilePhone,omitempty"`
	// The real-name verification information.
	//
	// example:
	//
	// aliyun-admin
	IdentityInformation *string `json:"IdentityInformation,omitempty" xml:"IdentityInformation,omitempty"`
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
	// The time when the member joined the resource directory.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The location of the member in the resource directory.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The time when the member was modified.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-k3****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The path of the member in the resource directory.
	ResourceDirectoryPath *string `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
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
	// The tags of the member.
	Tags []*GetAccountResponseBodyAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s GetAccountResponseBodyAccount) String() string {
	return dara.Prettify(s)
}

func (s GetAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *GetAccountResponseBodyAccount) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAccountResponseBodyAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *GetAccountResponseBodyAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetAccountResponseBodyAccount) GetEmailStatus() *string {
	return s.EmailStatus
}

func (s *GetAccountResponseBodyAccount) GetFolderId() *string {
	return s.FolderId
}

func (s *GetAccountResponseBodyAccount) GetHasSecureMobilePhone() *bool {
	return s.HasSecureMobilePhone
}

func (s *GetAccountResponseBodyAccount) GetIdentityInformation() *string {
	return s.IdentityInformation
}

func (s *GetAccountResponseBodyAccount) GetJoinMethod() *string {
	return s.JoinMethod
}

func (s *GetAccountResponseBodyAccount) GetJoinTime() *string {
	return s.JoinTime
}

func (s *GetAccountResponseBodyAccount) GetLocation() *string {
	return s.Location
}

func (s *GetAccountResponseBodyAccount) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetAccountResponseBodyAccount) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *GetAccountResponseBodyAccount) GetResourceDirectoryPath() *string {
	return s.ResourceDirectoryPath
}

func (s *GetAccountResponseBodyAccount) GetStatus() *string {
	return s.Status
}

func (s *GetAccountResponseBodyAccount) GetTags() []*GetAccountResponseBodyAccountTags {
	return s.Tags
}

func (s *GetAccountResponseBodyAccount) GetType() *string {
	return s.Type
}

func (s *GetAccountResponseBodyAccount) SetAccountId(v string) *GetAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetAccountName(v string) *GetAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetDisplayName(v string) *GetAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetEmailStatus(v string) *GetAccountResponseBodyAccount {
	s.EmailStatus = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetFolderId(v string) *GetAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetHasSecureMobilePhone(v bool) *GetAccountResponseBodyAccount {
	s.HasSecureMobilePhone = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetIdentityInformation(v string) *GetAccountResponseBodyAccount {
	s.IdentityInformation = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetJoinMethod(v string) *GetAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetJoinTime(v string) *GetAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetLocation(v string) *GetAccountResponseBodyAccount {
	s.Location = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetModifyTime(v string) *GetAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetResourceDirectoryId(v string) *GetAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetResourceDirectoryPath(v string) *GetAccountResponseBodyAccount {
	s.ResourceDirectoryPath = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetStatus(v string) *GetAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetTags(v []*GetAccountResponseBodyAccountTags) *GetAccountResponseBodyAccount {
	s.Tags = v
	return s
}

func (s *GetAccountResponseBodyAccount) SetType(v string) *GetAccountResponseBodyAccount {
	s.Type = &v
	return s
}

func (s *GetAccountResponseBodyAccount) Validate() error {
	return dara.Validate(s)
}

type GetAccountResponseBodyAccountTags struct {
	// The tag key.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAccountResponseBodyAccountTags) String() string {
	return dara.Prettify(s)
}

func (s GetAccountResponseBodyAccountTags) GoString() string {
	return s.String()
}

func (s *GetAccountResponseBodyAccountTags) GetKey() *string {
	return s.Key
}

func (s *GetAccountResponseBodyAccountTags) GetValue() *string {
	return s.Value
}

func (s *GetAccountResponseBodyAccountTags) SetKey(v string) *GetAccountResponseBodyAccountTags {
	s.Key = &v
	return s
}

func (s *GetAccountResponseBodyAccountTags) SetValue(v string) *GetAccountResponseBodyAccountTags {
	s.Value = &v
	return s
}

func (s *GetAccountResponseBodyAccountTags) Validate() error {
	return dara.Validate(s)
}
