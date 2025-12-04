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
	SetDBClusterId(v string) *ResetAccountPasswordRequest
	GetDBClusterId() *string
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
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The new password for the database account.
	//
	// >
	//
	// 	- The password must contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- The password can contain the following special characters: ! @ # $ % ^ & \\	- ( ) _ + - =
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Ff
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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

func (s *ResetAccountPasswordRequest) GetDBClusterId() *string {
	return s.DBClusterId
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

func (s *ResetAccountPasswordRequest) SetDBClusterId(v string) *ResetAccountPasswordRequest {
	s.DBClusterId = &v
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
