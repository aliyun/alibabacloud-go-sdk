// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSecretNoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ReleaseSecretNoRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *ReleaseSecretNoRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *ReleaseSecretNoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseSecretNoRequest
	GetResourceOwnerId() *int64
	SetSecretNo(v string) *ReleaseSecretNoRequest
	GetSecretNo() *string
}

type ReleaseSecretNoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC123456
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The prefix of phone numbers. When you call the ReleaseSecretNo operation with **SecretNo*	- specified, the system performs fuzzy matching against phone numbers based on the prefix.
	//
	// >  Up to 18 digits of a phone number prefix can be specified.
	//
	// This parameter is required.
	//
	// example:
	//
	// 130
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s ReleaseSecretNoRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSecretNoRequest) GoString() string {
	return s.String()
}

func (s *ReleaseSecretNoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseSecretNoRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *ReleaseSecretNoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseSecretNoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseSecretNoRequest) GetSecretNo() *string {
	return s.SecretNo
}

func (s *ReleaseSecretNoRequest) SetOwnerId(v int64) *ReleaseSecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetPoolKey(v string) *ReleaseSecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetResourceOwnerAccount(v string) *ReleaseSecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetResourceOwnerId(v int64) *ReleaseSecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetSecretNo(v string) *ReleaseSecretNoRequest {
	s.SecretNo = &v
	return s
}

func (s *ReleaseSecretNoRequest) Validate() error {
	return dara.Validate(s)
}
