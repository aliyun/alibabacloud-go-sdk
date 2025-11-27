// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ResetAccountRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ResetAccountRequest
	GetAccountPassword() *string
	SetDBInstanceId(v string) *ResetAccountRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ResetAccountRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResetAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ResetAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResetAccountRequest
	GetResourceOwnerId() *int64
}

type ResetAccountRequest struct {
	// The name of the privileged account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The new password of the privileged account.
	//
	// >
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// 	- It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- It can contain the following special characters: `! @ # $ & ; % ^ 	- () _ + - =`.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test123456
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResetAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ResetAccountRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ResetAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ResetAccountRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResetAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResetAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResetAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResetAccountRequest) SetAccountName(v string) *ResetAccountRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountRequest) SetAccountPassword(v string) *ResetAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountRequest) SetDBInstanceId(v string) *ResetAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResetAccountRequest) SetOwnerAccount(v string) *ResetAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetAccountRequest) SetOwnerId(v int64) *ResetAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetAccountRequest) SetResourceOwnerAccount(v string) *ResetAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetAccountRequest) SetResourceOwnerId(v int64) *ResetAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetAccountRequest) Validate() error {
	return dara.Validate(s)
}
