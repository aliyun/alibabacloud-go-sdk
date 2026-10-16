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
	// The member information.
	Account *CreateCloudAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The request ID.
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
	// The member name.
	//
	// example:
	//
	// admin-****
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// fd-bVaRIG****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The method used by the member to join the resource directory. Valid values:
	//
	// - invited: The member is invited to join.
	//
	// - created: The member is created in the resource directory.
	//
	// example:
	//
	// created
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member was last modified. The time is displayed in UTC, as specified by RFC 3339, with Z indicating UTC.
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
	// The resource directory ID.
	//
	// example:
	//
	// rd-k3****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member. Valid values:
	//
	// - CreateSuccess: The member is created.
	//
	// - CreateVerifying: The creation is being confirmed.
	//
	// - CreateFailed: The member failed to be created.
	//
	// - CreateExpired: The creation has expired.
	//
	// - CreateCancelled: The creation is canceled.
	//
	// - PromoteVerifying: The upgrade is being confirmed.
	//
	// - PromoteFailed: The member failed to be upgraded.
	//
	// - PromoteExpired: The upgrade has expired.
	//
	// - PromoteCancelled: The upgrade is canceled.
	//
	// - PromoteSuccess: The member is upgraded.
	//
	// - InviteSuccess: The member is invited.
	//
	// - Removed: The member is removed.
	//
	// example:
	//
	// CreateVerifying
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The member type. The value CloudAccount indicates a cloud account.
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
