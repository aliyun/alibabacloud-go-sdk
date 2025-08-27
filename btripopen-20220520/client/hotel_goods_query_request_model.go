// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelGoodsQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdultNum(v string) *HotelGoodsQueryRequest
	GetAdultNum() *string
	SetAgreementPrice(v bool) *HotelGoodsQueryRequest
	GetAgreementPrice() *bool
	SetBeginDate(v string) *HotelGoodsQueryRequest
	GetBeginDate() *string
	SetBreakfastIncluded(v bool) *HotelGoodsQueryRequest
	GetBreakfastIncluded() *bool
	SetBtripUserId(v string) *HotelGoodsQueryRequest
	GetBtripUserId() *string
	SetCityCode(v string) *HotelGoodsQueryRequest
	GetCityCode() *string
	SetEndDate(v string) *HotelGoodsQueryRequest
	GetEndDate() *string
	SetHotelId(v string) *HotelGoodsQueryRequest
	GetHotelId() *string
	SetPayOverType(v int32) *HotelGoodsQueryRequest
	GetPayOverType() *int32
	SetPaymentType(v int32) *HotelGoodsQueryRequest
	GetPaymentType() *int32
	SetSpecialInvoice(v bool) *HotelGoodsQueryRequest
	GetSpecialInvoice() *bool
	SetSuperMan(v int32) *HotelGoodsQueryRequest
	GetSuperMan() *int32
}

type HotelGoodsQueryRequest struct {
	// example:
	//
	// 1
	AdultNum *string `json:"adult_num,omitempty" xml:"adult_num,omitempty"`
	// example:
	//
	// false
	AgreementPrice *bool `json:"agreement_price,omitempty" xml:"agreement_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-27
	BeginDate *string `json:"begin_date,omitempty" xml:"begin_date,omitempty"`
	// example:
	//
	// false
	BreakfastIncluded *bool `json:"breakfast_included,omitempty" xml:"breakfast_included,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// example:
	//
	// 330100
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-28
	EndDate *string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 182873
	HotelId *string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// example:
	//
	// 0
	PayOverType *int32 `json:"pay_over_type,omitempty" xml:"pay_over_type,omitempty"`
	// example:
	//
	// 0
	PaymentType *int32 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// example:
	//
	// false
	SpecialInvoice *bool `json:"special_invoice,omitempty" xml:"special_invoice,omitempty"`
	// example:
	//
	// 0
	SuperMan *int32 `json:"super_man,omitempty" xml:"super_man,omitempty"`
}

func (s HotelGoodsQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryRequest) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryRequest) GetAdultNum() *string {
	return s.AdultNum
}

func (s *HotelGoodsQueryRequest) GetAgreementPrice() *bool {
	return s.AgreementPrice
}

func (s *HotelGoodsQueryRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *HotelGoodsQueryRequest) GetBreakfastIncluded() *bool {
	return s.BreakfastIncluded
}

func (s *HotelGoodsQueryRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelGoodsQueryRequest) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelGoodsQueryRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *HotelGoodsQueryRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *HotelGoodsQueryRequest) GetPayOverType() *int32 {
	return s.PayOverType
}

func (s *HotelGoodsQueryRequest) GetPaymentType() *int32 {
	return s.PaymentType
}

func (s *HotelGoodsQueryRequest) GetSpecialInvoice() *bool {
	return s.SpecialInvoice
}

func (s *HotelGoodsQueryRequest) GetSuperMan() *int32 {
	return s.SuperMan
}

func (s *HotelGoodsQueryRequest) SetAdultNum(v string) *HotelGoodsQueryRequest {
	s.AdultNum = &v
	return s
}

func (s *HotelGoodsQueryRequest) SetAgreementPrice(v bool) *HotelGoodsQueryRequest {
	s.AgreementPrice = &v
	return s
}

func (s *HotelGoodsQueryRequest) SetBeginDate(v string) *HotelGoodsQueryRequest {
	s.BeginDate = &v
	return s
}

func (s *HotelGoodsQueryRequest) SetBreakfastIncluded(v bool) *HotelGoodsQueryRequest {
	s.BreakfastIncluded = &v
	return s
}

func (s *HotelGoodsQueryRequest) SetBtripUserId(v string) *HotelGoodsQueryRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelGoodsQueryRequest) SetCityCode(v string) *HotelGoodsQueryRequest {
	s.CityCode = &v
	return s
}

func (s *HotelGoodsQueryRequest) SetEndDate(v string) *HotelGoodsQueryRequest {
	s.EndDate = &v
	return s
}

func (s *HotelGoodsQueryRequest) SetHotelId(v string) *HotelGoodsQueryRequest {
	s.HotelId = &v
	return s
}

func (s *HotelGoodsQueryRequest) SetPayOverType(v int32) *HotelGoodsQueryRequest {
	s.PayOverType = &v
	return s
}

func (s *HotelGoodsQueryRequest) SetPaymentType(v int32) *HotelGoodsQueryRequest {
	s.PaymentType = &v
	return s
}

func (s *HotelGoodsQueryRequest) SetSpecialInvoice(v bool) *HotelGoodsQueryRequest {
	s.SpecialInvoice = &v
	return s
}

func (s *HotelGoodsQueryRequest) SetSuperMan(v int32) *HotelGoodsQueryRequest {
	s.SuperMan = &v
	return s
}

func (s *HotelGoodsQueryRequest) Validate() error {
	return dara.Validate(s)
}
