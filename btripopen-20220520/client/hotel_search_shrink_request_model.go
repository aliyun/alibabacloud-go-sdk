// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelSearchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdultNum(v string) *HotelSearchShrinkRequest
	GetAdultNum() *string
	SetBrandCodeShrink(v string) *HotelSearchShrinkRequest
	GetBrandCodeShrink() *string
	SetBtripUserId(v string) *HotelSearchShrinkRequest
	GetBtripUserId() *string
	SetCheckInDate(v string) *HotelSearchShrinkRequest
	GetCheckInDate() *string
	SetCheckOutDate(v string) *HotelSearchShrinkRequest
	GetCheckOutDate() *string
	SetCityCode(v string) *HotelSearchShrinkRequest
	GetCityCode() *string
	SetDir(v int32) *HotelSearchShrinkRequest
	GetDir() *int32
	SetDistance(v int32) *HotelSearchShrinkRequest
	GetDistance() *int32
	SetDistrictCode(v string) *HotelSearchShrinkRequest
	GetDistrictCode() *string
	SetHotelStar(v string) *HotelSearchShrinkRequest
	GetHotelStar() *string
	SetIsProtocol(v bool) *HotelSearchShrinkRequest
	GetIsProtocol() *bool
	SetKeyWords(v string) *HotelSearchShrinkRequest
	GetKeyWords() *string
	SetLocation(v string) *HotelSearchShrinkRequest
	GetLocation() *string
	SetMaxPrice(v float64) *HotelSearchShrinkRequest
	GetMaxPrice() *float64
	SetMinPrice(v float64) *HotelSearchShrinkRequest
	GetMinPrice() *float64
	SetPageNo(v int32) *HotelSearchShrinkRequest
	GetPageNo() *int32
	SetPageSize(v int32) *HotelSearchShrinkRequest
	GetPageSize() *int32
	SetPayOverType(v int32) *HotelSearchShrinkRequest
	GetPayOverType() *int32
	SetPaymentType(v int32) *HotelSearchShrinkRequest
	GetPaymentType() *int32
	SetShidsShrink(v string) *HotelSearchShrinkRequest
	GetShidsShrink() *string
	SetSortCode(v int32) *HotelSearchShrinkRequest
	GetSortCode() *int32
	SetSuperMan(v int32) *HotelSearchShrinkRequest
	GetSuperMan() *int32
}

type HotelSearchShrinkRequest struct {
	// example:
	//
	// 1
	AdultNum        *string `json:"adult_num,omitempty" xml:"adult_num,omitempty"`
	BrandCodeShrink *string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// example:
	//
	// 1000
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-22
	CheckInDate *string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-22
	CheckOutDate *string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// example:
	//
	// 330100
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 1
	Dir *int32 `json:"dir,omitempty" xml:"dir,omitempty"`
	// example:
	//
	// 100
	Distance *int32 `json:"distance,omitempty" xml:"distance,omitempty"`
	// example:
	//
	// 330000
	DistrictCode *string `json:"district_code,omitempty" xml:"district_code,omitempty"`
	// example:
	//
	// 0
	HotelStar *string `json:"hotel_star,omitempty" xml:"hotel_star,omitempty"`
	// example:
	//
	// true
	IsProtocol *bool   `json:"is_protocol,omitempty" xml:"is_protocol,omitempty"`
	KeyWords   *string `json:"key_words,omitempty" xml:"key_words,omitempty"`
	// example:
	//
	// 120.010059, 30.284666
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// 1000
	MaxPrice *float64 `json:"max_price,omitempty" xml:"max_price,omitempty"`
	// example:
	//
	// 100
	MinPrice *float64 `json:"min_price,omitempty" xml:"min_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 1
	PayOverType *int32 `json:"pay_over_type,omitempty" xml:"pay_over_type,omitempty"`
	// example:
	//
	// 0
	PaymentType *int32  `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	ShidsShrink *string `json:"shids,omitempty" xml:"shids,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	SortCode *int32 `json:"sort_code,omitempty" xml:"sort_code,omitempty"`
	// example:
	//
	// 0
	SuperMan *int32 `json:"super_man,omitempty" xml:"super_man,omitempty"`
}

func (s HotelSearchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelSearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *HotelSearchShrinkRequest) GetAdultNum() *string {
	return s.AdultNum
}

func (s *HotelSearchShrinkRequest) GetBrandCodeShrink() *string {
	return s.BrandCodeShrink
}

func (s *HotelSearchShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelSearchShrinkRequest) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *HotelSearchShrinkRequest) GetCheckOutDate() *string {
	return s.CheckOutDate
}

func (s *HotelSearchShrinkRequest) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelSearchShrinkRequest) GetDir() *int32 {
	return s.Dir
}

func (s *HotelSearchShrinkRequest) GetDistance() *int32 {
	return s.Distance
}

func (s *HotelSearchShrinkRequest) GetDistrictCode() *string {
	return s.DistrictCode
}

func (s *HotelSearchShrinkRequest) GetHotelStar() *string {
	return s.HotelStar
}

func (s *HotelSearchShrinkRequest) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *HotelSearchShrinkRequest) GetKeyWords() *string {
	return s.KeyWords
}

func (s *HotelSearchShrinkRequest) GetLocation() *string {
	return s.Location
}

func (s *HotelSearchShrinkRequest) GetMaxPrice() *float64 {
	return s.MaxPrice
}

func (s *HotelSearchShrinkRequest) GetMinPrice() *float64 {
	return s.MinPrice
}

func (s *HotelSearchShrinkRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *HotelSearchShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *HotelSearchShrinkRequest) GetPayOverType() *int32 {
	return s.PayOverType
}

func (s *HotelSearchShrinkRequest) GetPaymentType() *int32 {
	return s.PaymentType
}

func (s *HotelSearchShrinkRequest) GetShidsShrink() *string {
	return s.ShidsShrink
}

func (s *HotelSearchShrinkRequest) GetSortCode() *int32 {
	return s.SortCode
}

func (s *HotelSearchShrinkRequest) GetSuperMan() *int32 {
	return s.SuperMan
}

func (s *HotelSearchShrinkRequest) SetAdultNum(v string) *HotelSearchShrinkRequest {
	s.AdultNum = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetBrandCodeShrink(v string) *HotelSearchShrinkRequest {
	s.BrandCodeShrink = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetBtripUserId(v string) *HotelSearchShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetCheckInDate(v string) *HotelSearchShrinkRequest {
	s.CheckInDate = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetCheckOutDate(v string) *HotelSearchShrinkRequest {
	s.CheckOutDate = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetCityCode(v string) *HotelSearchShrinkRequest {
	s.CityCode = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetDir(v int32) *HotelSearchShrinkRequest {
	s.Dir = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetDistance(v int32) *HotelSearchShrinkRequest {
	s.Distance = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetDistrictCode(v string) *HotelSearchShrinkRequest {
	s.DistrictCode = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetHotelStar(v string) *HotelSearchShrinkRequest {
	s.HotelStar = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetIsProtocol(v bool) *HotelSearchShrinkRequest {
	s.IsProtocol = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetKeyWords(v string) *HotelSearchShrinkRequest {
	s.KeyWords = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetLocation(v string) *HotelSearchShrinkRequest {
	s.Location = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetMaxPrice(v float64) *HotelSearchShrinkRequest {
	s.MaxPrice = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetMinPrice(v float64) *HotelSearchShrinkRequest {
	s.MinPrice = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetPageNo(v int32) *HotelSearchShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetPageSize(v int32) *HotelSearchShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetPayOverType(v int32) *HotelSearchShrinkRequest {
	s.PayOverType = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetPaymentType(v int32) *HotelSearchShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetShidsShrink(v string) *HotelSearchShrinkRequest {
	s.ShidsShrink = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetSortCode(v int32) *HotelSearchShrinkRequest {
	s.SortCode = &v
	return s
}

func (s *HotelSearchShrinkRequest) SetSuperMan(v int32) *HotelSearchShrinkRequest {
	s.SuperMan = &v
	return s
}

func (s *HotelSearchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
