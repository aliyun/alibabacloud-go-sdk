// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretNoDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QuerySecretNoDetailRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *QuerySecretNoDetailRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *QuerySecretNoDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySecretNoDetailRequest
	GetResourceOwnerId() *int64
	SetSecretNo(v string) *QuerySecretNoDetailRequest
	GetSecretNo() *string
}

type QuerySecretNoDetailRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2258****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The private number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s QuerySecretNoDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretNoDetailRequest) GoString() string {
	return s.String()
}

func (s *QuerySecretNoDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySecretNoDetailRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *QuerySecretNoDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySecretNoDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySecretNoDetailRequest) GetSecretNo() *string {
	return s.SecretNo
}

func (s *QuerySecretNoDetailRequest) SetOwnerId(v int64) *QuerySecretNoDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySecretNoDetailRequest) SetPoolKey(v string) *QuerySecretNoDetailRequest {
	s.PoolKey = &v
	return s
}

func (s *QuerySecretNoDetailRequest) SetResourceOwnerAccount(v string) *QuerySecretNoDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySecretNoDetailRequest) SetResourceOwnerId(v int64) *QuerySecretNoDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySecretNoDetailRequest) SetSecretNo(v string) *QuerySecretNoDetailRequest {
	s.SecretNo = &v
	return s
}

func (s *QuerySecretNoDetailRequest) Validate() error {
	return dara.Validate(s)
}
