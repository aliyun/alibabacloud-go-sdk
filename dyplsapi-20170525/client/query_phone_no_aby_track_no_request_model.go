// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneNoAByTrackNoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCabinetNo(v string) *QueryPhoneNoAByTrackNoRequest
	GetCabinetNo() *string
	SetOwnerId(v int64) *QueryPhoneNoAByTrackNoRequest
	GetOwnerId() *int64
	SetPhoneNoX(v string) *QueryPhoneNoAByTrackNoRequest
	GetPhoneNoX() *string
	SetResourceOwnerAccount(v string) *QueryPhoneNoAByTrackNoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPhoneNoAByTrackNoRequest
	GetResourceOwnerId() *int64
	SetTrackNo(v string) *QueryPhoneNoAByTrackNoRequest
	GetTrackNo() *string
}

type QueryPhoneNoAByTrackNoRequest struct {
	// The cabinet number.
	//
	// example:
	//
	// 25689****
	CabinetNo *string `json:"CabinetNo,omitempty" xml:"CabinetNo,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Phone number X returned by the API operation for creating a binding.
	//
	// example:
	//
	// 1710000****
	PhoneNoX             *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tracking number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22573****
	TrackNo *string `json:"trackNo,omitempty" xml:"trackNo,omitempty"`
}

func (s QueryPhoneNoAByTrackNoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoRequest) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoRequest) GetCabinetNo() *string {
	return s.CabinetNo
}

func (s *QueryPhoneNoAByTrackNoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPhoneNoAByTrackNoRequest) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *QueryPhoneNoAByTrackNoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPhoneNoAByTrackNoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPhoneNoAByTrackNoRequest) GetTrackNo() *string {
	return s.TrackNo
}

func (s *QueryPhoneNoAByTrackNoRequest) SetCabinetNo(v string) *QueryPhoneNoAByTrackNoRequest {
	s.CabinetNo = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetOwnerId(v int64) *QueryPhoneNoAByTrackNoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetPhoneNoX(v string) *QueryPhoneNoAByTrackNoRequest {
	s.PhoneNoX = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetResourceOwnerAccount(v string) *QueryPhoneNoAByTrackNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetResourceOwnerId(v int64) *QueryPhoneNoAByTrackNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetTrackNo(v string) *QueryPhoneNoAByTrackNoRequest {
	s.TrackNo = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) Validate() error {
	return dara.Validate(s)
}
