// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v *CreateResourceAccountResponseBodyAccount) *CreateResourceAccountResponseBody
	GetAccount() *CreateResourceAccountResponseBodyAccount
	SetRequestId(v string) *CreateResourceAccountResponseBody
	GetRequestId() *string
}

type CreateResourceAccountResponseBody struct {
	// The information about the member.
	Account *CreateResourceAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B356A415-D860-43E5-865A-E2193D62BBD6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountResponseBody) GetAccount() *CreateResourceAccountResponseBodyAccount {
	return s.Account
}

func (s *CreateResourceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceAccountResponseBody) SetAccount(v *CreateResourceAccountResponseBodyAccount) *CreateResourceAccountResponseBody {
	s.Account = v
	return s
}

func (s *CreateResourceAccountResponseBody) SetRequestId(v string) *CreateResourceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceAccountResponseBody) Validate() error {
	if s.Account != nil {
		if err := s.Account.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateResourceAccountResponseBodyAccount struct {
	// The ID of the member.
	//
	// example:
	//
	// 112730938585****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// alice@rd-3g****.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// Dev
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-r23M55****
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
	// 2020-12-31T03:37:39.456Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-31T03:37:39.456Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-3G****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member. The value CreateSuccess indicates that the member is created.
	//
	// example:
	//
	// CreateSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member. The value ResourceAccount indicates that the member is a resource account.
	//
	// example:
	//
	// ResourceAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateResourceAccountResponseBodyAccount) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountResponseBodyAccount) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateResourceAccountResponseBodyAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateResourceAccountResponseBodyAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateResourceAccountResponseBodyAccount) GetFolderId() *string {
	return s.FolderId
}

func (s *CreateResourceAccountResponseBodyAccount) GetJoinMethod() *string {
	return s.JoinMethod
}

func (s *CreateResourceAccountResponseBodyAccount) GetJoinTime() *string {
	return s.JoinTime
}

func (s *CreateResourceAccountResponseBodyAccount) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *CreateResourceAccountResponseBodyAccount) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *CreateResourceAccountResponseBodyAccount) GetStatus() *string {
	return s.Status
}

func (s *CreateResourceAccountResponseBodyAccount) GetType() *string {
	return s.Type
}

func (s *CreateResourceAccountResponseBodyAccount) SetAccountId(v string) *CreateResourceAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetAccountName(v string) *CreateResourceAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetDisplayName(v string) *CreateResourceAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetFolderId(v string) *CreateResourceAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetJoinMethod(v string) *CreateResourceAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetJoinTime(v string) *CreateResourceAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetModifyTime(v string) *CreateResourceAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetResourceDirectoryId(v string) *CreateResourceAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetStatus(v string) *CreateResourceAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetType(v string) *CreateResourceAccountResponseBodyAccount {
	s.Type = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) Validate() error {
	return dara.Validate(s)
}
