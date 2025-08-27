// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelPricePullShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *HotelPricePullShrinkRequest
	GetBtripUserId() *string
	SetCheckIn(v string) *HotelPricePullShrinkRequest
	GetCheckIn() *string
	SetCheckOut(v string) *HotelPricePullShrinkRequest
	GetCheckOut() *string
	SetCityCode(v int32) *HotelPricePullShrinkRequest
	GetCityCode() *int32
	SetHotelIdsShrink(v string) *HotelPricePullShrinkRequest
	GetHotelIdsShrink() *string
	SetPaymentType(v int32) *HotelPricePullShrinkRequest
	GetPaymentType() *int32
}

type HotelPricePullShrinkRequest struct {
	// example:
	//
	// 1000
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// example:
	//
	// 2022-05-15
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// example:
	//
	// 2022-05-15
	CheckOut *string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// example:
	//
	// 330100
	CityCode       *int32  `json:"city_code,omitempty" xml:"city_code,omitempty"`
	HotelIdsShrink *string `json:"hotel_ids,omitempty" xml:"hotel_ids,omitempty"`
	// example:
	//
	// 0
	PaymentType *int32 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
}

func (s HotelPricePullShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullShrinkRequest) GoString() string {
	return s.String()
}

func (s *HotelPricePullShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelPricePullShrinkRequest) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelPricePullShrinkRequest) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelPricePullShrinkRequest) GetCityCode() *int32 {
	return s.CityCode
}

func (s *HotelPricePullShrinkRequest) GetHotelIdsShrink() *string {
	return s.HotelIdsShrink
}

func (s *HotelPricePullShrinkRequest) GetPaymentType() *int32 {
	return s.PaymentType
}

func (s *HotelPricePullShrinkRequest) SetBtripUserId(v string) *HotelPricePullShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelPricePullShrinkRequest) SetCheckIn(v string) *HotelPricePullShrinkRequest {
	s.CheckIn = &v
	return s
}

func (s *HotelPricePullShrinkRequest) SetCheckOut(v string) *HotelPricePullShrinkRequest {
	s.CheckOut = &v
	return s
}

func (s *HotelPricePullShrinkRequest) SetCityCode(v int32) *HotelPricePullShrinkRequest {
	s.CityCode = &v
	return s
}

func (s *HotelPricePullShrinkRequest) SetHotelIdsShrink(v string) *HotelPricePullShrinkRequest {
	s.HotelIdsShrink = &v
	return s
}

func (s *HotelPricePullShrinkRequest) SetPaymentType(v int32) *HotelPricePullShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *HotelPricePullShrinkRequest) Validate() error {
	return dara.Validate(s)
}
