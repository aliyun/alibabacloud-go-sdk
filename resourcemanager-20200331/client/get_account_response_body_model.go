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
	// The member account information.
	Account *GetAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The request ID.
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
	if s.Account != nil {
		if err := s.Account.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccountResponseBodyAccount struct {
	// The member account ID.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// someone@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member account.
	//
	// example:
	//
	// admin
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email status. Valid values:
	//
	// - Empty: No email modification record exists.
	//
	// - WAIT_MODIFY: The modification is in progress.
	//
	// - CANCELLED: The modification is canceled.
	//
	// - EXPIRED: The modification has expired.
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
	// The real-name verification information.
	//
	// example:
	//
	// aliyun-admin
	IdentityInformation *string `json:"IdentityInformation,omitempty" xml:"IdentityInformation,omitempty"`
	// The method used by the member account to join the resource directory. Valid values:
	//
	// - invited: Invited.
	//
	// - created: Created.
	//
	// example:
	//
	// created
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member account joined the resource directory. The time is displayed in UTC, as specified by RFC 3339, with Z indicating UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The location of the member account in the resource directory.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The time when the member account was last modified. The time is displayed in UTC, as specified by RFC 3339, with Z indicating UTC.
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
	// The path ID (RDPath) of the member account in the resource directory.
	ResourceDirectoryPath *string `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
	// The join status of the member account. Valid values:
	//
	// - CreateSuccess: Created.
	//
	// - PromoteVerifying: Upgrade is being confirmed.
	//
	// - PromoteFailed: Upgrade failed.
	//
	// - PromoteExpired: Upgrade expired.
	//
	// - PromoteCancelled: Upgrade canceled.
	//
	// - PromoteSuccess: Upgrade succeeded.
	//
	// - InviteSuccess: Invitation succeeded.
	//
	// example:
	//
	// CreateSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*GetAccountResponseBodyAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the member account. Valid values:
	//
	// - CloudAccount: Cloud account.
	//
	// - ResourceAccount: Resource account.
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
