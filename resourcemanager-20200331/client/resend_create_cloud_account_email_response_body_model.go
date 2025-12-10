// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendCreateCloudAccountEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v *ResendCreateCloudAccountEmailResponseBodyAccount) *ResendCreateCloudAccountEmailResponseBody
	GetAccount() *ResendCreateCloudAccountEmailResponseBodyAccount
	SetRequestId(v string) *ResendCreateCloudAccountEmailResponseBody
	GetRequestId() *string
}

type ResendCreateCloudAccountEmailResponseBody struct {
	// The information of the member account.
	Account *ResendCreateCloudAccountEmailResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResendCreateCloudAccountEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResendCreateCloudAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ResendCreateCloudAccountEmailResponseBody) GetAccount() *ResendCreateCloudAccountEmailResponseBodyAccount {
	return s.Account
}

func (s *ResendCreateCloudAccountEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResendCreateCloudAccountEmailResponseBody) SetAccount(v *ResendCreateCloudAccountEmailResponseBodyAccount) *ResendCreateCloudAccountEmailResponseBody {
	s.Account = v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBody) SetRequestId(v string) *ResendCreateCloudAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBody) Validate() error {
	if s.Account != nil {
		if err := s.Account.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResendCreateCloudAccountEmailResponseBodyAccount struct {
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
	// CreateVerifying
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member account. The value CloudAccount indicates that the member account is a cloud account.
	//
	// example:
	//
	// CloudAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ResendCreateCloudAccountEmailResponseBodyAccount) String() string {
	return dara.Prettify(s)
}

func (s ResendCreateCloudAccountEmailResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) GetAccountId() *string {
	return s.AccountId
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) GetFolderId() *string {
	return s.FolderId
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) GetJoinMethod() *string {
	return s.JoinMethod
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) GetJoinTime() *string {
	return s.JoinTime
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) GetRecordId() *string {
	return s.RecordId
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) GetStatus() *string {
	return s.Status
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) GetType() *string {
	return s.Type
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetAccountId(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetAccountName(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetDisplayName(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetFolderId(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetJoinMethod(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetJoinTime(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetModifyTime(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetRecordId(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.RecordId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetResourceDirectoryId(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetStatus(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) SetType(v string) *ResendCreateCloudAccountEmailResponseBodyAccount {
	s.Type = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseBodyAccount) Validate() error {
	return dara.Validate(s)
}
