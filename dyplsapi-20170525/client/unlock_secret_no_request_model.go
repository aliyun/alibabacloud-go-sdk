// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockSecretNoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *UnlockSecretNoRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *UnlockSecretNoRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *UnlockSecretNoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnlockSecretNoRequest
	GetResourceOwnerId() *int64
	SetSecretNo(v string) *UnlockSecretNoRequest
	GetSecretNo() *string
}

type UnlockSecretNoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2256****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The private number that you want to unlock. You must enter a complete mobile phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1300000****
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s UnlockSecretNoRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockSecretNoRequest) GoString() string {
	return s.String()
}

func (s *UnlockSecretNoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnlockSecretNoRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *UnlockSecretNoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnlockSecretNoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnlockSecretNoRequest) GetSecretNo() *string {
	return s.SecretNo
}

func (s *UnlockSecretNoRequest) SetOwnerId(v int64) *UnlockSecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *UnlockSecretNoRequest) SetPoolKey(v string) *UnlockSecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *UnlockSecretNoRequest) SetResourceOwnerAccount(v string) *UnlockSecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnlockSecretNoRequest) SetResourceOwnerId(v int64) *UnlockSecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnlockSecretNoRequest) SetSecretNo(v string) *UnlockSecretNoRequest {
	s.SecretNo = &v
	return s
}

func (s *UnlockSecretNoRequest) Validate() error {
	return dara.Validate(s)
}
