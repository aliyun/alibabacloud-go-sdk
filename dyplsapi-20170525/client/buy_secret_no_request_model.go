// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuySecretNoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *BuySecretNoRequest
	GetCity() *string
	SetDisplayPool(v bool) *BuySecretNoRequest
	GetDisplayPool() *bool
	SetOwnerId(v int64) *BuySecretNoRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *BuySecretNoRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *BuySecretNoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BuySecretNoRequest
	GetResourceOwnerId() *int64
	SetSecretNo(v string) *BuySecretNoRequest
	GetSecretNo() *string
	SetSpecId(v int64) *BuySecretNoRequest
	GetSpecId() *int64
}

type BuySecretNoRequest struct {
	// Specifies the home location of the phone number.
	//
	// >
	//
	// 	- The home location can be set only to a location in the Chinese mainland.
	//
	// 	- A phone number that starts with 95 does not have a home location. If you purchase a phone number that starts with 95, set this parameter to **Nationwide**.
	//
	// This parameter is required.
	//
	// example:
	//
	// hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// Specifies whether to add the phone number to the pool of numbers that will be displayed during calls.
	//
	// >  This parameter takes effect only for customers who have enabled the number display feature.
	//
	// example:
	//
	// true
	DisplayPool *bool  `json:"DisplayPool,omitempty" xml:"DisplayPool,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The prefix of the phone number. If you specify the value of **SecretNo*	- when you purchase a phone number, a phone number starting with the specified prefix is selected.
	//
	// >  You can specify up to 18 digits of the phone number prefix.
	//
	// example:
	//
	// 130
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The type of the phone number. Valid values:
	//
	// 	- **1**: a phone number assigned by a virtual network operator, that is, a phone number that belongs to the 170 or 171 number segment.
	//
	// 	- **2**: a phone number provided by a carrier.
	//
	// 	- **3**: a phone number that starts with 95.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SpecId *int64 `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s BuySecretNoRequest) String() string {
	return dara.Prettify(s)
}

func (s BuySecretNoRequest) GoString() string {
	return s.String()
}

func (s *BuySecretNoRequest) GetCity() *string {
	return s.City
}

func (s *BuySecretNoRequest) GetDisplayPool() *bool {
	return s.DisplayPool
}

func (s *BuySecretNoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BuySecretNoRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *BuySecretNoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BuySecretNoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BuySecretNoRequest) GetSecretNo() *string {
	return s.SecretNo
}

func (s *BuySecretNoRequest) GetSpecId() *int64 {
	return s.SpecId
}

func (s *BuySecretNoRequest) SetCity(v string) *BuySecretNoRequest {
	s.City = &v
	return s
}

func (s *BuySecretNoRequest) SetDisplayPool(v bool) *BuySecretNoRequest {
	s.DisplayPool = &v
	return s
}

func (s *BuySecretNoRequest) SetOwnerId(v int64) *BuySecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *BuySecretNoRequest) SetPoolKey(v string) *BuySecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *BuySecretNoRequest) SetResourceOwnerAccount(v string) *BuySecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BuySecretNoRequest) SetResourceOwnerId(v int64) *BuySecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BuySecretNoRequest) SetSecretNo(v string) *BuySecretNoRequest {
	s.SecretNo = &v
	return s
}

func (s *BuySecretNoRequest) SetSpecId(v int64) *BuySecretNoRequest {
	s.SpecId = &v
	return s
}

func (s *BuySecretNoRequest) Validate() error {
	return dara.Validate(s)
}
