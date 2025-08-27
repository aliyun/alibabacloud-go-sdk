// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelAskingPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdultNum(v string) *HotelAskingPriceRequest
	GetAdultNum() *string
	SetBtripUserId(v string) *HotelAskingPriceRequest
	GetBtripUserId() *string
	SetCheckInDate(v string) *HotelAskingPriceRequest
	GetCheckInDate() *string
	SetCheckOutDate(v string) *HotelAskingPriceRequest
	GetCheckOutDate() *string
	SetCityCode(v string) *HotelAskingPriceRequest
	GetCityCode() *string
	SetCityName(v string) *HotelAskingPriceRequest
	GetCityName() *string
	SetDir(v int32) *HotelAskingPriceRequest
	GetDir() *int32
	SetHotelStar(v string) *HotelAskingPriceRequest
	GetHotelStar() *string
	SetIsProtocol(v bool) *HotelAskingPriceRequest
	GetIsProtocol() *bool
	SetPaymentType(v int32) *HotelAskingPriceRequest
	GetPaymentType() *int32
	SetShids(v []*int64) *HotelAskingPriceRequest
	GetShids() []*int64
	SetSortCode(v int32) *HotelAskingPriceRequest
	GetSortCode() *int32
}

type HotelAskingPriceRequest struct {
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
	Shids []*int64 `json:"shids,omitempty" xml:"shids,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	SortCode *int32 `json:"sort_code,omitempty" xml:"sort_code,omitempty"`
}

func (s HotelAskingPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelAskingPriceRequest) GoString() string {
	return s.String()
}

func (s *HotelAskingPriceRequest) GetAdultNum() *string {
	return s.AdultNum
}

func (s *HotelAskingPriceRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelAskingPriceRequest) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *HotelAskingPriceRequest) GetCheckOutDate() *string {
	return s.CheckOutDate
}

func (s *HotelAskingPriceRequest) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelAskingPriceRequest) GetCityName() *string {
	return s.CityName
}

func (s *HotelAskingPriceRequest) GetDir() *int32 {
	return s.Dir
}

func (s *HotelAskingPriceRequest) GetHotelStar() *string {
	return s.HotelStar
}

func (s *HotelAskingPriceRequest) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *HotelAskingPriceRequest) GetPaymentType() *int32 {
	return s.PaymentType
}

func (s *HotelAskingPriceRequest) GetShids() []*int64 {
	return s.Shids
}

func (s *HotelAskingPriceRequest) GetSortCode() *int32 {
	return s.SortCode
}

func (s *HotelAskingPriceRequest) SetAdultNum(v string) *HotelAskingPriceRequest {
	s.AdultNum = &v
	return s
}

func (s *HotelAskingPriceRequest) SetBtripUserId(v string) *HotelAskingPriceRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelAskingPriceRequest) SetCheckInDate(v string) *HotelAskingPriceRequest {
	s.CheckInDate = &v
	return s
}

func (s *HotelAskingPriceRequest) SetCheckOutDate(v string) *HotelAskingPriceRequest {
	s.CheckOutDate = &v
	return s
}

func (s *HotelAskingPriceRequest) SetCityCode(v string) *HotelAskingPriceRequest {
	s.CityCode = &v
	return s
}

func (s *HotelAskingPriceRequest) SetCityName(v string) *HotelAskingPriceRequest {
	s.CityName = &v
	return s
}

func (s *HotelAskingPriceRequest) SetDir(v int32) *HotelAskingPriceRequest {
	s.Dir = &v
	return s
}

func (s *HotelAskingPriceRequest) SetHotelStar(v string) *HotelAskingPriceRequest {
	s.HotelStar = &v
	return s
}

func (s *HotelAskingPriceRequest) SetIsProtocol(v bool) *HotelAskingPriceRequest {
	s.IsProtocol = &v
	return s
}

func (s *HotelAskingPriceRequest) SetPaymentType(v int32) *HotelAskingPriceRequest {
	s.PaymentType = &v
	return s
}

func (s *HotelAskingPriceRequest) SetShids(v []*int64) *HotelAskingPriceRequest {
	s.Shids = v
	return s
}

func (s *HotelAskingPriceRequest) SetSortCode(v int32) *HotelAskingPriceRequest {
	s.SortCode = &v
	return s
}

func (s *HotelAskingPriceRequest) Validate() error {
	return dara.Validate(s)
}
