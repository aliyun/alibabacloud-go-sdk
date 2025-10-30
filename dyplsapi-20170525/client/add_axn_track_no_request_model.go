// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAxnTrackNoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *AddAxnTrackNoRequest
	GetOwnerId() *int64
	SetPhoneNoX(v string) *AddAxnTrackNoRequest
	GetPhoneNoX() *string
	SetPoolKey(v string) *AddAxnTrackNoRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *AddAxnTrackNoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddAxnTrackNoRequest
	GetResourceOwnerId() *int64
	SetSubsId(v string) *AddAxnTrackNoRequest
	GetSubsId() *string
	SetTrackNo(v string) *AddAxnTrackNoRequest
	GetTrackNo() *string
}

type AddAxnTrackNoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private number in the AXN binding, that is, phone number X.
	//
	// You can call the [BindAxn](https://help.aliyun.com/document_detail/110258.html) operation to obtain the value of PhoneNoX.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700000****
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2235****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The binding ID.
	//
	// You can call the [BindAxn](https://help.aliyun.com/document_detail/110258.html) operation to obtain the value of SubsId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15678890****
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	// The tracking number.
	//
	// This parameter is required.
	//
	// example:
	//
	// abcde*****
	TrackNo *string `json:"trackNo,omitempty" xml:"trackNo,omitempty"`
}

func (s AddAxnTrackNoRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAxnTrackNoRequest) GoString() string {
	return s.String()
}

func (s *AddAxnTrackNoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddAxnTrackNoRequest) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *AddAxnTrackNoRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *AddAxnTrackNoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddAxnTrackNoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddAxnTrackNoRequest) GetSubsId() *string {
	return s.SubsId
}

func (s *AddAxnTrackNoRequest) GetTrackNo() *string {
	return s.TrackNo
}

func (s *AddAxnTrackNoRequest) SetOwnerId(v int64) *AddAxnTrackNoRequest {
	s.OwnerId = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetPhoneNoX(v string) *AddAxnTrackNoRequest {
	s.PhoneNoX = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetPoolKey(v string) *AddAxnTrackNoRequest {
	s.PoolKey = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetResourceOwnerAccount(v string) *AddAxnTrackNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetResourceOwnerId(v int64) *AddAxnTrackNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetSubsId(v string) *AddAxnTrackNoRequest {
	s.SubsId = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetTrackNo(v string) *AddAxnTrackNoRequest {
	s.TrackNo = &v
	return s
}

func (s *AddAxnTrackNoRequest) Validate() error {
	return dara.Validate(s)
}
