// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendPromoteResourceAccountEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v *ResendPromoteResourceAccountEmailResponseBodyAccount) *ResendPromoteResourceAccountEmailResponseBody
	GetAccount() *ResendPromoteResourceAccountEmailResponseBodyAccount
	SetRequestId(v string) *ResendPromoteResourceAccountEmailResponseBody
	GetRequestId() *string
}

type ResendPromoteResourceAccountEmailResponseBody struct {
	// The information of the member account.
	Account *ResendPromoteResourceAccountEmailResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResendPromoteResourceAccountEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResendPromoteResourceAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ResendPromoteResourceAccountEmailResponseBody) GetAccount() *ResendPromoteResourceAccountEmailResponseBodyAccount {
	return s.Account
}

func (s *ResendPromoteResourceAccountEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResendPromoteResourceAccountEmailResponseBody) SetAccount(v *ResendPromoteResourceAccountEmailResponseBodyAccount) *ResendPromoteResourceAccountEmailResponseBody {
	s.Account = v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBody) SetRequestId(v string) *ResendPromoteResourceAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBody) Validate() error {
	if s.Account != nil {
		if err := s.Account.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResendPromoteResourceAccountEmailResponseBodyAccount struct {
	// The ID of the account.
	//
	// example:
	//
	// 12323344****
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
	// 16950264-3f0d-4ca9-82dd-6ee7a3d33d6b
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
	// 	- CloudAccount: cloud account
	//
	// 	- ResourceAccount: resource account
	//
	// example:
	//
	// ResourceAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ResendPromoteResourceAccountEmailResponseBodyAccount) String() string {
	return dara.Prettify(s)
}

func (s ResendPromoteResourceAccountEmailResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) GetAccountId() *string {
	return s.AccountId
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) GetFolderId() *string {
	return s.FolderId
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) GetJoinMethod() *string {
	return s.JoinMethod
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) GetJoinTime() *string {
	return s.JoinTime
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) GetRecordId() *string {
	return s.RecordId
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) GetStatus() *string {
	return s.Status
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) GetType() *string {
	return s.Type
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetAccountId(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetAccountName(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetDisplayName(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetFolderId(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetJoinMethod(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetJoinTime(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetModifyTime(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetRecordId(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.RecordId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetResourceDirectoryId(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetStatus(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) SetType(v string) *ResendPromoteResourceAccountEmailResponseBodyAccount {
	s.Type = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseBodyAccount) Validate() error {
	return dara.Validate(s)
}
