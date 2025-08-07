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
	SetDBClusterId(v string) *ModifyAccountPasswordRequest
	GetDBClusterId() *string
	SetNewAccountPassword(v string) *ModifyAccountPasswordRequest
	GetNewAccountPassword() *string
	SetOwnerAccount(v string) *ModifyAccountPasswordRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAccountPasswordRequest
	GetOwnerId() *int64
	SetPasswordType(v string) *ModifyAccountPasswordRequest
	GetPasswordType() *string
	SetResourceOwnerAccount(v string) *ModifyAccountPasswordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountPasswordRequest
	GetResourceOwnerId() *int64
}

type ModifyAccountPasswordRequest struct {
	// The username of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The new password of the account. The new password must meet the following requirements:
	//
	// 	- It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- It must be 8 to 32 characters in length.
	//
	// 	- Special characters include `! @ # $ % ^ & 	- ( ) _ + - =`
	//
	// This parameter is required.
	//
	// example:
	//
	// Pw123456
	NewAccountPassword *string `json:"NewAccountPassword,omitempty" xml:"NewAccountPassword,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password type.
	//
	// example:
	//
	// Tair
	PasswordType         *string `json:"PasswordType,omitempty" xml:"PasswordType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *ModifyAccountPasswordRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAccountPasswordRequest) GetNewAccountPassword() *string {
	return s.NewAccountPassword
}

func (s *ModifyAccountPasswordRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAccountPasswordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAccountPasswordRequest) GetPasswordType() *string {
	return s.PasswordType
}

func (s *ModifyAccountPasswordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountPasswordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountPasswordRequest) SetAccountName(v string) *ModifyAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetDBClusterId(v string) *ModifyAccountPasswordRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetNewAccountPassword(v string) *ModifyAccountPasswordRequest {
	s.NewAccountPassword = &v
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

func (s *ModifyAccountPasswordRequest) SetPasswordType(v string) *ModifyAccountPasswordRequest {
	s.PasswordType = &v
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

func (s *ModifyAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
