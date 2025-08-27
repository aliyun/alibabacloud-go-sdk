// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelAskingPriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdultNum(v string) *HotelAskingPriceShrinkRequest
	GetAdultNum() *string
	SetBtripUserId(v string) *HotelAskingPriceShrinkRequest
	GetBtripUserId() *string
	SetCheckInDate(v string) *HotelAskingPriceShrinkRequest
	GetCheckInDate() *string
	SetCheckOutDate(v string) *HotelAskingPriceShrinkRequest
	GetCheckOutDate() *string
	SetCityCode(v string) *HotelAskingPriceShrinkRequest
	GetCityCode() *string
	SetCityName(v string) *HotelAskingPriceShrinkRequest
	GetCityName() *string
	SetDir(v int32) *HotelAskingPriceShrinkRequest
	GetDir() *int32
	SetHotelStar(v string) *HotelAskingPriceShrinkRequest
	GetHotelStar() *string
	SetIsProtocol(v bool) *HotelAskingPriceShrinkRequest
	GetIsProtocol() *bool
	SetPaymentType(v int32) *HotelAskingPriceShrinkRequest
	GetPaymentType() *int32
	SetShidsShrink(v string) *HotelAskingPriceShrinkRequest
	GetShidsShrink() *string
	SetSortCode(v int32) *HotelAskingPriceShrinkRequest
	GetSortCode() *int32
}

type HotelAskingPriceShrinkRequest struct {
	// example:
	//
	// 1
	AdultNum *string `json:"adult_num,omitempty" xml:"adult_num,omitempty"`
	// example:
	//
	// 1000
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-22 00:00:00
	CheckInDate *string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-22 00:00:00
	CheckOutDate *string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// example:
	//
	// 330100
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 1
	Dir *int32 `json:"dir,omitempty" xml:"dir,omitempty"`
	// example:
	//
	// 1
	HotelStar *string `json:"hotel_star,omitempty" xml:"hotel_star,omitempty"`
	// example:
	//
	// true
	IsProtocol *bool `json:"is_protocol,omitempty" xml:"is_protocol,omitempty"`
	// example:
	//
	// 0
	PaymentType *int32 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// This parameter is required.
	ShidsShrink *string `json:"shids,omitempty" xml:"shids,omitempty"`
	// example:
	//
	// 0
	SortCode *int32 `json:"sort_code,omitempty" xml:"sort_code,omitempty"`
}

func (s HotelAskingPriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelAskingPriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *HotelAskingPriceShrinkRequest) GetAdultNum() *string {
	return s.AdultNum
}

func (s *HotelAskingPriceShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelAskingPriceShrinkRequest) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *HotelAskingPriceShrinkRequest) GetCheckOutDate() *string {
	return s.CheckOutDate
}

func (s *HotelAskingPriceShrinkRequest) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelAskingPriceShrinkRequest) GetCityName() *string {
	return s.CityName
}

func (s *HotelAskingPriceShrinkRequest) GetDir() *int32 {
	return s.Dir
}

func (s *HotelAskingPriceShrinkRequest) GetHotelStar() *string {
	return s.HotelStar
}

func (s *HotelAskingPriceShrinkRequest) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *HotelAskingPriceShrinkRequest) GetPaymentType() *int32 {
	return s.PaymentType
}

func (s *HotelAskingPriceShrinkRequest) GetShidsShrink() *string {
	return s.ShidsShrink
}

func (s *HotelAskingPriceShrinkRequest) GetSortCode() *int32 {
	return s.SortCode
}

func (s *HotelAskingPriceShrinkRequest) SetAdultNum(v string) *HotelAskingPriceShrinkRequest {
	s.AdultNum = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) SetBtripUserId(v string) *HotelAskingPriceShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) SetCheckInDate(v string) *HotelAskingPriceShrinkRequest {
	s.CheckInDate = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) SetCheckOutDate(v string) *HotelAskingPriceShrinkRequest {
	s.CheckOutDate = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) SetCityCode(v string) *HotelAskingPriceShrinkRequest {
	s.CityCode = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) SetCityName(v string) *HotelAskingPriceShrinkRequest {
	s.CityName = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) SetDir(v int32) *HotelAskingPriceShrinkRequest {
	s.Dir = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) SetHotelStar(v string) *HotelAskingPriceShrinkRequest {
	s.HotelStar = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) SetIsProtocol(v bool) *HotelAskingPriceShrinkRequest {
	s.IsProtocol = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) SetPaymentType(v int32) *HotelAskingPriceShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) SetShidsShrink(v string) *HotelAskingPriceShrinkRequest {
	s.ShidsShrink = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) SetSortCode(v int32) *HotelAskingPriceShrinkRequest {
	s.SortCode = &v
	return s
}

func (s *HotelAskingPriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
