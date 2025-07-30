// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountPasswordRequest
	GetAccountName() *string
	SetInstanceId(v string) *ModifyAccountPasswordRequest
	GetInstanceId() *string
	SetNewAccountPassword(v string) *ModifyAccountPasswordRequest
	GetNewAccountPassword() *string
	SetOldAccountPassword(v string) *ModifyAccountPasswordRequest
	GetOldAccountPassword() *string
	SetOwnerAccount(v string) *ModifyAccountPasswordRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAccountPasswordRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAccountPasswordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountPasswordRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyAccountPasswordRequest
	GetSecurityToken() *string
	SetSourceBiz(v string) *ModifyAccountPasswordRequest
	GetSourceBiz() *string
}

type ModifyAccountPasswordRequest struct {
	// The username of the account for which you want to change the password. You can call the [DescribeAccounts](https://help.aliyun.com/document_detail/473816.html) operation to query the username of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new password to be set for the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and specific special characters. These special characters include `! @ # $ % ^ & 	- ( ) _ + - =`
	//
	// This parameter is required.
	//
	// example:
	//
	// newPassWd888****
	NewAccountPassword *string `json:"NewAccountPassword,omitempty" xml:"NewAccountPassword,omitempty"`
	// The current password of the account.
	//
	//
	//
	// > If you forget your password, you can call the [ResetAccountPassword](https://help.aliyun.com/document_detail/473815.html) operation to reset your password.
	//
	// This parameter is required.
	//
	// example:
	//
	// oldPassWd999****
	OldAccountPassword   *string `json:"OldAccountPassword,omitempty" xml:"OldAccountPassword,omitempty"`
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

func (s ModifyAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountPasswordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAccountPasswordRequest) GetNewAccountPassword() *string {
	return s.NewAccountPassword
}

func (s *ModifyAccountPasswordRequest) GetOldAccountPassword() *string {
	return s.OldAccountPassword
}

func (s *ModifyAccountPasswordRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAccountPasswordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAccountPasswordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountPasswordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountPasswordRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyAccountPasswordRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *ModifyAccountPasswordRequest) SetAccountName(v string) *ModifyAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetInstanceId(v string) *ModifyAccountPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetNewAccountPassword(v string) *ModifyAccountPasswordRequest {
	s.NewAccountPassword = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetOldAccountPassword(v string) *ModifyAccountPasswordRequest {
	s.OldAccountPassword = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetOwnerAccount(v string) *ModifyAccountPasswordRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetOwnerId(v int64) *ModifyAccountPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetResourceOwnerAccount(v string) *ModifyAccountPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetResourceOwnerId(v int64) *ModifyAccountPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetSecurityToken(v string) *ModifyAccountPasswordRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetSourceBiz(v string) *ModifyAccountPasswordRequest {
	s.SourceBiz = &v
	return s
}

func (s *ModifyAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
