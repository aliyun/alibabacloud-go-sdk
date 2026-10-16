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
	// The member account information.
	Account *ResendCreateCloudAccountEmailResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The request ID.
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
	// The account ID.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The account name.
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
	// The ID of the resource folder.
	//
	// example:
	//
	// fd-bVaRIG****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
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
	// The time when the member account was last modified. The time is displayed in UTC, as specified by RFC 3339, with Z indicating UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account record ID.
	//
	// example:
	//
	// 16950264-3f0d-4ca9-82dd-6ee7a3d3****
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-k3****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The join status of the member account. Valid values:
	//
	// - CreateSuccess: Created.
	//
	// - CreateVerifying: Pending creation confirmation.
	//
	// - CreateFailed: Creation failed.
	//
	// - CreateExpired: Creation expired.
	//
	// - CreateCancelled: Creation canceled.
	//
	// - PromoteVerifying: Pending upgrade confirmation.
	//
	// - PromoteFailed: Upgrade failed.
	//
	// - PromoteExpired: Upgrade expired.
	//
	// - PromoteCancelled: Upgrade canceled.
	//
	// - PromoteSuccess: Upgraded.
	//
	// - InviteSuccess: Invited.
	//
	// - Removed: Removed.
	//
	// example:
	//
	// CreateVerifying
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member account. The value CloudAccount indicates a cloud account.
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
