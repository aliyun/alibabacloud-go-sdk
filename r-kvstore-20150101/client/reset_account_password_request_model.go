// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ResetAccountPasswordRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ResetAccountPasswordRequest
	GetAccountPassword() *string
	SetInstanceId(v string) *ResetAccountPasswordRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ResetAccountPasswordRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResetAccountPasswordRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ResetAccountPasswordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResetAccountPasswordRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ResetAccountPasswordRequest
	GetSecurityToken() *string
	SetSourceBiz(v string) *ResetAccountPasswordRequest
	GetSourceBiz() *string
}

type ResetAccountPasswordRequest struct {
	// The name of the account. You can call the [DescribeAccounts](~~DescribeAccounts~~) operation to obtain the name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The new password for the account. The password must be 8 to 32 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!@ # $ % ^ & 	- ( ) _ + - =`
	//
	// This parameter is required.
	//
	// example:
	//
	// uWonno_221****
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The ID of the instance to which the account belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is used only for internal maintenance. You do not need to specify this parameter.
	//
	// example:
	//
	// SDK
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ResetAccountPasswordRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ResetAccountPasswordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetAccountPasswordRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResetAccountPasswordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResetAccountPasswordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResetAccountPasswordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResetAccountPasswordRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ResetAccountPasswordRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetInstanceId(v string) *ResetAccountPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetOwnerAccount(v string) *ResetAccountPasswordRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetOwnerId(v int64) *ResetAccountPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetResourceOwnerAccount(v string) *ResetAccountPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetResourceOwnerId(v int64) *ResetAccountPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetSecurityToken(v string) *ResetAccountPasswordRequest {
	s.SecurityToken = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetSourceBiz(v string) *ResetAccountPasswordRequest {
	s.SourceBiz = &v
	return s
}

func (s *ResetAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
