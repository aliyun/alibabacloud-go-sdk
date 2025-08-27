// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdultNum(v string) *HotelSearchRequest
	GetAdultNum() *string
	SetBrandCode(v []*string) *HotelSearchRequest
	GetBrandCode() []*string
	SetBtripUserId(v string) *HotelSearchRequest
	GetBtripUserId() *string
	SetCheckInDate(v string) *HotelSearchRequest
	GetCheckInDate() *string
	SetCheckOutDate(v string) *HotelSearchRequest
	GetCheckOutDate() *string
	SetCityCode(v string) *HotelSearchRequest
	GetCityCode() *string
	SetDir(v int32) *HotelSearchRequest
	GetDir() *int32
	SetDistance(v int32) *HotelSearchRequest
	GetDistance() *int32
	SetDistrictCode(v string) *HotelSearchRequest
	GetDistrictCode() *string
	SetHotelStar(v string) *HotelSearchRequest
	GetHotelStar() *string
	SetIsProtocol(v bool) *HotelSearchRequest
	GetIsProtocol() *bool
	SetKeyWords(v string) *HotelSearchRequest
	GetKeyWords() *string
	SetLocation(v string) *HotelSearchRequest
	GetLocation() *string
	SetMaxPrice(v float64) *HotelSearchRequest
	GetMaxPrice() *float64
	SetMinPrice(v float64) *HotelSearchRequest
	GetMinPrice() *float64
	SetPageNo(v int32) *HotelSearchRequest
	GetPageNo() *int32
	SetPageSize(v int32) *HotelSearchRequest
	GetPageSize() *int32
	SetPayOverType(v int32) *HotelSearchRequest
	GetPayOverType() *int32
	SetPaymentType(v int32) *HotelSearchRequest
	GetPaymentType() *int32
	SetShids(v []*int64) *HotelSearchRequest
	GetShids() []*int64
	SetSortCode(v int32) *HotelSearchRequest
	GetSortCode() *int32
	SetSuperMan(v int32) *HotelSearchRequest
	GetSuperMan() *int32
}

type HotelSearchRequest struct {
	// example:
	//
	// 1
	AdultNum  *string   `json:"adult_num,omitempty" xml:"adult_num,omitempty"`
	BrandCode []*string `json:"brand_code,omitempty" xml:"brand_code,omitempty" type:"Repeated"`
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
	PaymentType *int32   `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	Shids       []*int64 `json:"shids,omitempty" xml:"shids,omitempty" type:"Repeated"`
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

func (s HotelSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelSearchRequest) GoString() string {
	return s.String()
}

func (s *HotelSearchRequest) GetAdultNum() *string {
	return s.AdultNum
}

func (s *HotelSearchRequest) GetBrandCode() []*string {
	return s.BrandCode
}

func (s *HotelSearchRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelSearchRequest) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *HotelSearchRequest) GetCheckOutDate() *string {
	return s.CheckOutDate
}

func (s *HotelSearchRequest) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelSearchRequest) GetDir() *int32 {
	return s.Dir
}

func (s *HotelSearchRequest) GetDistance() *int32 {
	return s.Distance
}

func (s *HotelSearchRequest) GetDistrictCode() *string {
	return s.DistrictCode
}

func (s *HotelSearchRequest) GetHotelStar() *string {
	return s.HotelStar
}

func (s *HotelSearchRequest) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *HotelSearchRequest) GetKeyWords() *string {
	return s.KeyWords
}

func (s *HotelSearchRequest) GetLocation() *string {
	return s.Location
}

func (s *HotelSearchRequest) GetMaxPrice() *float64 {
	return s.MaxPrice
}

func (s *HotelSearchRequest) GetMinPrice() *float64 {
	return s.MinPrice
}

func (s *HotelSearchRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *HotelSearchRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *HotelSearchRequest) GetPayOverType() *int32 {
	return s.PayOverType
}

func (s *HotelSearchRequest) GetPaymentType() *int32 {
	return s.PaymentType
}

func (s *HotelSearchRequest) GetShids() []*int64 {
	return s.Shids
}

func (s *HotelSearchRequest) GetSortCode() *int32 {
	return s.SortCode
}

func (s *HotelSearchRequest) GetSuperMan() *int32 {
	return s.SuperMan
}

func (s *HotelSearchRequest) SetAdultNum(v string) *HotelSearchRequest {
	s.AdultNum = &v
	return s
}

func (s *HotelSearchRequest) SetBrandCode(v []*string) *HotelSearchRequest {
	s.BrandCode = v
	return s
}

func (s *HotelSearchRequest) SetBtripUserId(v string) *HotelSearchRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelSearchRequest) SetCheckInDate(v string) *HotelSearchRequest {
	s.CheckInDate = &v
	return s
}

func (s *HotelSearchRequest) SetCheckOutDate(v string) *HotelSearchRequest {
	s.CheckOutDate = &v
	return s
}

func (s *HotelSearchRequest) SetCityCode(v string) *HotelSearchRequest {
	s.CityCode = &v
	return s
}

func (s *HotelSearchRequest) SetDir(v int32) *HotelSearchRequest {
	s.Dir = &v
	return s
}

func (s *HotelSearchRequest) SetDistance(v int32) *HotelSearchRequest {
	s.Distance = &v
	return s
}

func (s *HotelSearchRequest) SetDistrictCode(v string) *HotelSearchRequest {
	s.DistrictCode = &v
	return s
}

func (s *HotelSearchRequest) SetHotelStar(v string) *HotelSearchRequest {
	s.HotelStar = &v
	return s
}

func (s *HotelSearchRequest) SetIsProtocol(v bool) *HotelSearchRequest {
	s.IsProtocol = &v
	return s
}

func (s *HotelSearchRequest) SetKeyWords(v string) *HotelSearchRequest {
	s.KeyWords = &v
	return s
}

func (s *HotelSearchRequest) SetLocation(v string) *HotelSearchRequest {
	s.Location = &v
	return s
}

func (s *HotelSearchRequest) SetMaxPrice(v float64) *HotelSearchRequest {
	s.MaxPrice = &v
	return s
}

func (s *HotelSearchRequest) SetMinPrice(v float64) *HotelSearchRequest {
	s.MinPrice = &v
	return s
}

func (s *HotelSearchRequest) SetPageNo(v int32) *HotelSearchRequest {
	s.PageNo = &v
	return s
}

func (s *HotelSearchRequest) SetPageSize(v int32) *HotelSearchRequest {
	s.PageSize = &v
	return s
}

func (s *HotelSearchRequest) SetPayOverType(v int32) *HotelSearchRequest {
	s.PayOverType = &v
	return s
}

func (s *HotelSearchRequest) SetPaymentType(v int32) *HotelSearchRequest {
	s.PaymentType = &v
	return s
}

func (s *HotelSearchRequest) SetShids(v []*int64) *HotelSearchRequest {
	s.Shids = v
	return s
}

func (s *HotelSearchRequest) SetSortCode(v int32) *HotelSearchRequest {
	s.SortCode = &v
	return s
}

func (s *HotelSearchRequest) SetSuperMan(v int32) *HotelSearchRequest {
	s.SuperMan = &v
	return s
}

func (s *HotelSearchRequest) Validate() error {
	return dara.Validate(s)
}
