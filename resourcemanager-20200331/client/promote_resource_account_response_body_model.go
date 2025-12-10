// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromoteResourceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v *PromoteResourceAccountResponseBodyAccount) *PromoteResourceAccountResponseBody
	GetAccount() *PromoteResourceAccountResponseBodyAccount
	SetRequestId(v string) *PromoteResourceAccountResponseBody
	GetRequestId() *string
}

type PromoteResourceAccountResponseBody struct {
	// The information of the member account.
	Account *PromoteResourceAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PromoteResourceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PromoteResourceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *PromoteResourceAccountResponseBody) GetAccount() *PromoteResourceAccountResponseBodyAccount {
	return s.Account
}

func (s *PromoteResourceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PromoteResourceAccountResponseBody) SetAccount(v *PromoteResourceAccountResponseBodyAccount) *PromoteResourceAccountResponseBody {
	s.Account = v
	return s
}

func (s *PromoteResourceAccountResponseBody) SetRequestId(v string) *PromoteResourceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *PromoteResourceAccountResponseBody) Validate() error {
	if s.Account != nil {
		if err := s.Account.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PromoteResourceAccountResponseBodyAccount struct {
	// The ID of the member account.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member account.
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
	// The ID of the folder.
	//
	// example:
	//
	// fd-bVaRIG****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member account joined the resource directory. Valid values:
	//
	// 	- invited: The member account is invited to join the resource directory.
	//
	// 	- created: The member account is directly created in the resource directory.
	//
	// example:
	//
	// created
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member account joined the resource directory.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member account was modified.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account record ID.
	//
	// example:
	//
	// 06950264-3f0d-4ca9-82dd-6ee7a3d33d6b
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-k3****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member account. Valid values:
	//
	// 	- CreateSuccess: The member account is created.
	//
	// 	- CreateVerifying: The creation of the member account is under confirmation.
	//
	// 	- CreateFailed: The member account failed to be created.
	//
	// 	- CreateExpired: The creation of the member account expired.
	//
	// 	- CreateCancelled: The creation of the member account is canceled.
	//
	// 	- PromoteVerifying: The upgrade of the member account is under confirmation.
	//
	// 	- PromoteFailed: The member account failed to be upgraded.
	//
	// 	- PromoteExpired: The upgrade of the member account expired.
	//
	// 	- PromoteCancelled: The upgrade of the member account is canceled.
	//
	// 	- PromoteSuccess: The member account is upgraded.
	//
	// 	- InviteSuccess: The owner of the member account accepted the invitation.
	//
	// 	- Removed: The member account is removed from the resource directory.
	//
	// example:
	//
	// PromoteVerifying
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member account. Valid values:
	//
	// 	- CloudAccount
	//
	// 	- ResourceAccount
	//
	// example:
	//
	// ResourceAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PromoteResourceAccountResponseBodyAccount) String() string {
	return dara.Prettify(s)
}

func (s PromoteResourceAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *PromoteResourceAccountResponseBodyAccount) GetAccountId() *string {
	return s.AccountId
}

func (s *PromoteResourceAccountResponseBodyAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *PromoteResourceAccountResponseBodyAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *PromoteResourceAccountResponseBodyAccount) GetFolderId() *string {
	return s.FolderId
}

func (s *PromoteResourceAccountResponseBodyAccount) GetJoinMethod() *string {
	return s.JoinMethod
}

func (s *PromoteResourceAccountResponseBodyAccount) GetJoinTime() *string {
	return s.JoinTime
}

func (s *PromoteResourceAccountResponseBodyAccount) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *PromoteResourceAccountResponseBodyAccount) GetRecordId() *string {
	return s.RecordId
}

func (s *PromoteResourceAccountResponseBodyAccount) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *PromoteResourceAccountResponseBodyAccount) GetStatus() *string {
	return s.Status
}

func (s *PromoteResourceAccountResponseBodyAccount) GetType() *string {
	return s.Type
}

func (s *PromoteResourceAccountResponseBodyAccount) SetAccountId(v string) *PromoteResourceAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetAccountName(v string) *PromoteResourceAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetDisplayName(v string) *PromoteResourceAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetFolderId(v string) *PromoteResourceAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetJoinMethod(v string) *PromoteResourceAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetJoinTime(v string) *PromoteResourceAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetModifyTime(v string) *PromoteResourceAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetRecordId(v string) *PromoteResourceAccountResponseBodyAccount {
	s.RecordId = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetResourceDirectoryId(v string) *PromoteResourceAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetStatus(v string) *PromoteResourceAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) SetType(v string) *PromoteResourceAccountResponseBodyAccount {
	s.Type = &v
	return s
}

func (s *PromoteResourceAccountResponseBodyAccount) Validate() error {
	return dara.Validate(s)
}
