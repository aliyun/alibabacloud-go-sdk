// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v *CreateCloudAccountResponseBodyAccount) *CreateCloudAccountResponseBody
	GetAccount() *CreateCloudAccountResponseBodyAccount
	SetRequestId(v string) *CreateCloudAccountResponseBody
	GetRequestId() *string
}

type CreateCloudAccountResponseBody struct {
	// The information of the member account.
	Account *CreateCloudAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountResponseBody) GetAccount() *CreateCloudAccountResponseBodyAccount {
	return s.Account
}

func (s *CreateCloudAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudAccountResponseBody) SetAccount(v *CreateCloudAccountResponseBodyAccount) *CreateCloudAccountResponseBody {
	s.Account = v
	return s
}

func (s *CreateCloudAccountResponseBody) SetRequestId(v string) *CreateCloudAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudAccountResponseBody) Validate() error {
	if s.Account != nil {
		if err := s.Account.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCloudAccountResponseBodyAccount struct {
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
	// admin-****
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
	// 06950264-3f0d-4ca9-82dd-6ee7a3d3****
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

func (s CreateCloudAccountResponseBodyAccount) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountResponseBodyAccount) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateCloudAccountResponseBodyAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateCloudAccountResponseBodyAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateCloudAccountResponseBodyAccount) GetFolderId() *string {
	return s.FolderId
}

func (s *CreateCloudAccountResponseBodyAccount) GetJoinMethod() *string {
	return s.JoinMethod
}

func (s *CreateCloudAccountResponseBodyAccount) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *CreateCloudAccountResponseBodyAccount) GetRecordId() *string {
	return s.RecordId
}

func (s *CreateCloudAccountResponseBodyAccount) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *CreateCloudAccountResponseBodyAccount) GetStatus() *string {
	return s.Status
}

func (s *CreateCloudAccountResponseBodyAccount) GetType() *string {
	return s.Type
}

func (s *CreateCloudAccountResponseBodyAccount) SetAccountId(v string) *CreateCloudAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetAccountName(v string) *CreateCloudAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetDisplayName(v string) *CreateCloudAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetFolderId(v string) *CreateCloudAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetJoinMethod(v string) *CreateCloudAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetModifyTime(v string) *CreateCloudAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetRecordId(v string) *CreateCloudAccountResponseBodyAccount {
	s.RecordId = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetResourceDirectoryId(v string) *CreateCloudAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetStatus(v string) *CreateCloudAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) SetType(v string) *CreateCloudAccountResponseBodyAccount {
	s.Type = &v
	return s
}

func (s *CreateCloudAccountResponseBodyAccount) Validate() error {
	return dara.Validate(s)
}
