// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelPricePullRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *HotelPricePullRequest
	GetBtripUserId() *string
	SetCheckIn(v string) *HotelPricePullRequest
	GetCheckIn() *string
	SetCheckOut(v string) *HotelPricePullRequest
	GetCheckOut() *string
	SetCityCode(v int32) *HotelPricePullRequest
	GetCityCode() *int32
	SetHotelIds(v []*string) *HotelPricePullRequest
	GetHotelIds() []*string
	SetPaymentType(v int32) *HotelPricePullRequest
	GetPaymentType() *int32
}

type HotelPricePullRequest struct {
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
	CityCode *int32    `json:"city_code,omitempty" xml:"city_code,omitempty"`
	HotelIds []*string `json:"hotel_ids,omitempty" xml:"hotel_ids,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	PaymentType *int32 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
}

func (s HotelPricePullRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullRequest) GoString() string {
	return s.String()
}

func (s *HotelPricePullRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelPricePullRequest) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelPricePullRequest) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelPricePullRequest) GetCityCode() *int32 {
	return s.CityCode
}

func (s *HotelPricePullRequest) GetHotelIds() []*string {
	return s.HotelIds
}

func (s *HotelPricePullRequest) GetPaymentType() *int32 {
	return s.PaymentType
}

func (s *HotelPricePullRequest) SetBtripUserId(v string) *HotelPricePullRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelPricePullRequest) SetCheckIn(v string) *HotelPricePullRequest {
	s.CheckIn = &v
	return s
}

func (s *HotelPricePullRequest) SetCheckOut(v string) *HotelPricePullRequest {
	s.CheckOut = &v
	return s
}

func (s *HotelPricePullRequest) SetCityCode(v int32) *HotelPricePullRequest {
	s.CityCode = &v
	return s
}

func (s *HotelPricePullRequest) SetHotelIds(v []*string) *HotelPricePullRequest {
	s.HotelIds = v
	return s
}

func (s *HotelPricePullRequest) SetPaymentType(v int32) *HotelPricePullRequest {
	s.PaymentType = &v
	return s
}

func (s *HotelPricePullRequest) Validate() error {
	return dara.Validate(s)
}
