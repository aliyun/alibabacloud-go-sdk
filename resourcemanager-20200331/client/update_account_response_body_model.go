// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v *UpdateAccountResponseBodyAccount) *UpdateAccountResponseBody
	GetAccount() *UpdateAccountResponseBodyAccount
	SetRequestId(v string) *UpdateAccountResponseBody
	GetRequestId() *string
}

type UpdateAccountResponseBody struct {
	// The information of the member.
	Account *UpdateAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountResponseBody) GetAccount() *UpdateAccountResponseBodyAccount {
	return s.Account
}

func (s *UpdateAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAccountResponseBody) SetAccount(v *UpdateAccountResponseBodyAccount) *UpdateAccountResponseBody {
	s.Account = v
	return s
}

func (s *UpdateAccountResponseBody) SetRequestId(v string) *UpdateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccountResponseBody) Validate() error {
	if s.Account != nil {
		if err := s.Account.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAccountResponseBodyAccount struct {
	// The ID of the Alibaba Cloud account that corresponds to the member.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the Alibaba Cloud account that corresponds to the member.
	//
	// example:
	//
	// ecs-manager@aliyun.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
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
	// rd-k3****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member. Valid values:
	//
	// 	- CreateSuccess: The member is created.
	//
	// 	- InviteSuccess: The member accepts the invitation.
	//
	// 	- Removed: The member is removed.
	//
	// 	- SwitchSuccess: The type of the member is switched.
	//
	// example:
	//
	// CreateSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s UpdateAccountResponseBodyAccount) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *UpdateAccountResponseBodyAccount) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateAccountResponseBodyAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *UpdateAccountResponseBodyAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateAccountResponseBodyAccount) GetFolderId() *string {
	return s.FolderId
}

func (s *UpdateAccountResponseBodyAccount) GetJoinMethod() *string {
	return s.JoinMethod
}

func (s *UpdateAccountResponseBodyAccount) GetJoinTime() *string {
	return s.JoinTime
}

func (s *UpdateAccountResponseBodyAccount) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *UpdateAccountResponseBodyAccount) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *UpdateAccountResponseBodyAccount) GetStatus() *string {
	return s.Status
}

func (s *UpdateAccountResponseBodyAccount) GetType() *string {
	return s.Type
}

func (s *UpdateAccountResponseBodyAccount) SetAccountId(v string) *UpdateAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetAccountName(v string) *UpdateAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetDisplayName(v string) *UpdateAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetFolderId(v string) *UpdateAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetJoinMethod(v string) *UpdateAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetJoinTime(v string) *UpdateAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetModifyTime(v string) *UpdateAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetResourceDirectoryId(v string) *UpdateAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetStatus(v string) *UpdateAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetType(v string) *UpdateAccountResponseBodyAccount {
	s.Type = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) Validate() error {
	return dara.Validate(s)
}
