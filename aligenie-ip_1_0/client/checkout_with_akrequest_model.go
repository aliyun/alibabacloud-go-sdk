// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckoutWithAKRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *CheckoutWithAKRequest
	GetHotelId() *string
	SetRoomNo(v string) *CheckoutWithAKRequest
	GetRoomNo() *string
}

type CheckoutWithAKRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s CheckoutWithAKRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckoutWithAKRequest) GoString() string {
	return s.String()
}

func (s *CheckoutWithAKRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *CheckoutWithAKRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *CheckoutWithAKRequest) SetHotelId(v string) *CheckoutWithAKRequest {
	s.HotelId = &v
	return s
}

func (s *CheckoutWithAKRequest) SetRoomNo(v string) *CheckoutWithAKRequest {
	s.RoomNo = &v
	return s
}

func (s *CheckoutWithAKRequest) Validate() error {
	return dara.Validate(s)
}
