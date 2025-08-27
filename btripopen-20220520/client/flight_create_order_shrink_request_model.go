// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCreateOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrAirportCode(v string) *FlightCreateOrderShrinkRequest
	GetArrAirportCode() *string
	SetArrCityCode(v string) *FlightCreateOrderShrinkRequest
	GetArrCityCode() *string
	SetAutoPay(v int32) *FlightCreateOrderShrinkRequest
	GetAutoPay() *int32
	SetBuyerName(v string) *FlightCreateOrderShrinkRequest
	GetBuyerName() *string
	SetBuyerUniqueKey(v string) *FlightCreateOrderShrinkRequest
	GetBuyerUniqueKey() *string
	SetContactInfoShrink(v string) *FlightCreateOrderShrinkRequest
	GetContactInfoShrink() *string
	SetDepAirportCode(v string) *FlightCreateOrderShrinkRequest
	GetDepAirportCode() *string
	SetDepCityCode(v string) *FlightCreateOrderShrinkRequest
	GetDepCityCode() *string
	SetDepDate(v string) *FlightCreateOrderShrinkRequest
	GetDepDate() *string
	SetDisOrderId(v string) *FlightCreateOrderShrinkRequest
	GetDisOrderId() *string
	SetOrderAttrShrink(v string) *FlightCreateOrderShrinkRequest
	GetOrderAttrShrink() *string
	SetOrderParams(v string) *FlightCreateOrderShrinkRequest
	GetOrderParams() *string
	SetOtaItemId(v string) *FlightCreateOrderShrinkRequest
	GetOtaItemId() *string
	SetPrice(v int64) *FlightCreateOrderShrinkRequest
	GetPrice() *int64
	SetReceiptAddress(v string) *FlightCreateOrderShrinkRequest
	GetReceiptAddress() *string
	SetReceiptTarget(v int32) *FlightCreateOrderShrinkRequest
	GetReceiptTarget() *int32
	SetReceiptTitle(v string) *FlightCreateOrderShrinkRequest
	GetReceiptTitle() *string
	SetTravelerInfoListShrink(v string) *FlightCreateOrderShrinkRequest
	GetTravelerInfoListShrink() *string
	SetTripType(v int32) *FlightCreateOrderShrinkRequest
	GetTripType() *int32
}

type FlightCreateOrderShrinkRequest struct {
	// example:
	//
	// HGH
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// 0
	AutoPay   *int32  `json:"auto_pay,omitempty" xml:"auto_pay,omitempty"`
	BuyerName *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	BuyerUniqueKey *string `json:"buyer_unique_key,omitempty" xml:"buyer_unique_key,omitempty"`
	// This parameter is required.
	ContactInfoShrink *string `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
	// example:
	//
	// PEK
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2000-00-00 00:00:00
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId      *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	OrderAttrShrink *string `json:"order_attr,omitempty" xml:"order_attr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000_1_0
	OrderParams *string `json:"order_params,omitempty" xml:"order_params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7fb731deeb4510b86c17e8c8c25740_11
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Price          *int64  `json:"price,omitempty" xml:"price,omitempty"`
	ReceiptAddress *string `json:"receipt_address,omitempty" xml:"receipt_address,omitempty"`
	// example:
	//
	// 1
	ReceiptTarget *int32  `json:"receipt_target,omitempty" xml:"receipt_target,omitempty"`
	ReceiptTitle  *string `json:"receipt_title,omitempty" xml:"receipt_title,omitempty"`
	// This parameter is required.
	TravelerInfoListShrink *string `json:"traveler_info_list,omitempty" xml:"traveler_info_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s FlightCreateOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderShrinkRequest) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightCreateOrderShrinkRequest) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightCreateOrderShrinkRequest) GetAutoPay() *int32 {
	return s.AutoPay
}

func (s *FlightCreateOrderShrinkRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *FlightCreateOrderShrinkRequest) GetBuyerUniqueKey() *string {
	return s.BuyerUniqueKey
}

func (s *FlightCreateOrderShrinkRequest) GetContactInfoShrink() *string {
	return s.ContactInfoShrink
}

func (s *FlightCreateOrderShrinkRequest) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightCreateOrderShrinkRequest) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightCreateOrderShrinkRequest) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightCreateOrderShrinkRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightCreateOrderShrinkRequest) GetOrderAttrShrink() *string {
	return s.OrderAttrShrink
}

func (s *FlightCreateOrderShrinkRequest) GetOrderParams() *string {
	return s.OrderParams
}

func (s *FlightCreateOrderShrinkRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *FlightCreateOrderShrinkRequest) GetPrice() *int64 {
	return s.Price
}

func (s *FlightCreateOrderShrinkRequest) GetReceiptAddress() *string {
	return s.ReceiptAddress
}

func (s *FlightCreateOrderShrinkRequest) GetReceiptTarget() *int32 {
	return s.ReceiptTarget
}

func (s *FlightCreateOrderShrinkRequest) GetReceiptTitle() *string {
	return s.ReceiptTitle
}

func (s *FlightCreateOrderShrinkRequest) GetTravelerInfoListShrink() *string {
	return s.TravelerInfoListShrink
}

func (s *FlightCreateOrderShrinkRequest) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightCreateOrderShrinkRequest) SetArrAirportCode(v string) *FlightCreateOrderShrinkRequest {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetArrCityCode(v string) *FlightCreateOrderShrinkRequest {
	s.ArrCityCode = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetAutoPay(v int32) *FlightCreateOrderShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetBuyerName(v string) *FlightCreateOrderShrinkRequest {
	s.BuyerName = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetBuyerUniqueKey(v string) *FlightCreateOrderShrinkRequest {
	s.BuyerUniqueKey = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetContactInfoShrink(v string) *FlightCreateOrderShrinkRequest {
	s.ContactInfoShrink = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetDepAirportCode(v string) *FlightCreateOrderShrinkRequest {
	s.DepAirportCode = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetDepCityCode(v string) *FlightCreateOrderShrinkRequest {
	s.DepCityCode = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetDepDate(v string) *FlightCreateOrderShrinkRequest {
	s.DepDate = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetDisOrderId(v string) *FlightCreateOrderShrinkRequest {
	s.DisOrderId = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetOrderAttrShrink(v string) *FlightCreateOrderShrinkRequest {
	s.OrderAttrShrink = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetOrderParams(v string) *FlightCreateOrderShrinkRequest {
	s.OrderParams = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetOtaItemId(v string) *FlightCreateOrderShrinkRequest {
	s.OtaItemId = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetPrice(v int64) *FlightCreateOrderShrinkRequest {
	s.Price = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetReceiptAddress(v string) *FlightCreateOrderShrinkRequest {
	s.ReceiptAddress = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetReceiptTarget(v int32) *FlightCreateOrderShrinkRequest {
	s.ReceiptTarget = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetReceiptTitle(v string) *FlightCreateOrderShrinkRequest {
	s.ReceiptTitle = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetTravelerInfoListShrink(v string) *FlightCreateOrderShrinkRequest {
	s.TravelerInfoListShrink = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) SetTripType(v int32) *FlightCreateOrderShrinkRequest {
	s.TripType = &v
	return s
}

func (s *FlightCreateOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
