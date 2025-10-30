// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySubsIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QuerySubsIdRequest
	GetOwnerId() *int64
	SetPhoneNoX(v string) *QuerySubsIdRequest
	GetPhoneNoX() *string
	SetPoolKey(v string) *QuerySubsIdRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *QuerySubsIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySubsIdRequest
	GetResourceOwnerId() *int64
}

type QuerySubsIdRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private number in the binding, that is, phone number X.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dyplsnext.console.aliyun.com/overview) and view the key of the phone number pool on the Number Pool Management page.
	//
	// example:
	//
	// FC123456
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySubsIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySubsIdRequest) GoString() string {
	return s.String()
}

func (s *QuerySubsIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySubsIdRequest) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *QuerySubsIdRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *QuerySubsIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySubsIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySubsIdRequest) SetOwnerId(v int64) *QuerySubsIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySubsIdRequest) SetPhoneNoX(v string) *QuerySubsIdRequest {
	s.PhoneNoX = &v
	return s
}

func (s *QuerySubsIdRequest) SetPoolKey(v string) *QuerySubsIdRequest {
	s.PoolKey = &v
	return s
}

func (s *QuerySubsIdRequest) SetResourceOwnerAccount(v string) *QuerySubsIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySubsIdRequest) SetResourceOwnerId(v int64) *QuerySubsIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySubsIdRequest) Validate() error {
	return dara.Validate(s)
}
