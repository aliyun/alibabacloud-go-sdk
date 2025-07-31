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
	SetCharacterType(v string) *ResetAccountPasswordRequest
	GetCharacterType() *string
	SetDBInstanceId(v string) *ResetAccountPasswordRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ResetAccountPasswordRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResetAccountPasswordRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ResetAccountPasswordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResetAccountPasswordRequest
	GetResourceOwnerId() *int64
}

type ResetAccountPasswordRequest struct {
	// The account whose password needs to be reset. Set the value to **root**.
	//
	// This parameter is required.
	//
	// example:
	//
	// root
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The new password.
	//
	// 	- The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! # $ % ^ & 	- ( ) _ + - =`
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ali!123456
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The role of the instance.
	//
	// 	- If the instance is a sharded cluster instance, this parameter is required. Valid values: db and cs.
	//
	// 	- If the instance is a replica set instance, you can leave this parameter empty or set the parameter to normal.
	//
	// example:
	//
	// db
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *ResetAccountPasswordRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *ResetAccountPasswordRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
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

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetCharacterType(v string) *ResetAccountPasswordRequest {
	s.CharacterType = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetDBInstanceId(v string) *ResetAccountPasswordRequest {
	s.DBInstanceId = &v
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

func (s *ResetAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
