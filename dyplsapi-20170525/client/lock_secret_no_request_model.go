// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockSecretNoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *LockSecretNoRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *LockSecretNoRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *LockSecretNoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *LockSecretNoRequest
	GetResourceOwnerId() *int64
	SetSecretNo(v string) *LockSecretNoRequest
	GetSecretNo() *string
}

type LockSecretNoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC123****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The private number that you want to lock. You must enter a complete mobile phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1300000****
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s LockSecretNoRequest) String() string {
	return dara.Prettify(s)
}

func (s LockSecretNoRequest) GoString() string {
	return s.String()
}

func (s *LockSecretNoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *LockSecretNoRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *LockSecretNoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *LockSecretNoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *LockSecretNoRequest) GetSecretNo() *string {
	return s.SecretNo
}

func (s *LockSecretNoRequest) SetOwnerId(v int64) *LockSecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *LockSecretNoRequest) SetPoolKey(v string) *LockSecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *LockSecretNoRequest) SetResourceOwnerAccount(v string) *LockSecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *LockSecretNoRequest) SetResourceOwnerId(v int64) *LockSecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *LockSecretNoRequest) SetSecretNo(v string) *LockSecretNoRequest {
	s.SecretNo = &v
	return s
}

func (s *LockSecretNoRequest) Validate() error {
	return dara.Validate(s)
}
