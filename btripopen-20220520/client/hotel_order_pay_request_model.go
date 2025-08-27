// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderPayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripOrderId(v int64) *HotelOrderPayRequest
	GetBtripOrderId() *int64
	SetBtripUserId(v string) *HotelOrderPayRequest
	GetBtripUserId() *string
	SetCompanyPayFee(v int64) *HotelOrderPayRequest
	GetCompanyPayFee() *int64
	SetPersonPayFee(v int64) *HotelOrderPayRequest
	GetPersonPayFee() *int64
	SetThirdPayAccount(v string) *HotelOrderPayRequest
	GetThirdPayAccount() *string
	SetThirdTradeNo(v string) *HotelOrderPayRequest
	GetThirdTradeNo() *string
	SetTotalPrice(v int64) *HotelOrderPayRequest
	GetTotalPrice() *int64
}

type HotelOrderPayRequest struct {
	// 供应商订单号（取自创单返回的订单号）
	//
	// This parameter is required.
	//
	// example:
	//
	// 1002202194207077022
	BtripOrderId *int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23918781
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	CompanyPayFee *int64 `json:"company_pay_fee,omitempty" xml:"company_pay_fee,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	PersonPayFee *int64 `json:"person_pay_fee,omitempty" xml:"person_pay_fee,omitempty"`
	// example:
	//
	// demo
	ThirdPayAccount *string `json:"third_pay_account,omitempty" xml:"third_pay_account,omitempty"`
	// example:
	//
	// demo
	ThirdTradeNo *string `json:"third_trade_no,omitempty" xml:"third_trade_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	TotalPrice *int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
}

func (s HotelOrderPayRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPayRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderPayRequest) GetBtripOrderId() *int64 {
	return s.BtripOrderId
}

func (s *HotelOrderPayRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelOrderPayRequest) GetCompanyPayFee() *int64 {
	return s.CompanyPayFee
}

func (s *HotelOrderPayRequest) GetPersonPayFee() *int64 {
	return s.PersonPayFee
}

func (s *HotelOrderPayRequest) GetThirdPayAccount() *string {
	return s.ThirdPayAccount
}

func (s *HotelOrderPayRequest) GetThirdTradeNo() *string {
	return s.ThirdTradeNo
}

func (s *HotelOrderPayRequest) GetTotalPrice() *int64 {
	return s.TotalPrice
}

func (s *HotelOrderPayRequest) SetBtripOrderId(v int64) *HotelOrderPayRequest {
	s.BtripOrderId = &v
	return s
}

func (s *HotelOrderPayRequest) SetBtripUserId(v string) *HotelOrderPayRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelOrderPayRequest) SetCompanyPayFee(v int64) *HotelOrderPayRequest {
	s.CompanyPayFee = &v
	return s
}

func (s *HotelOrderPayRequest) SetPersonPayFee(v int64) *HotelOrderPayRequest {
	s.PersonPayFee = &v
	return s
}

func (s *HotelOrderPayRequest) SetThirdPayAccount(v string) *HotelOrderPayRequest {
	s.ThirdPayAccount = &v
	return s
}

func (s *HotelOrderPayRequest) SetThirdTradeNo(v string) *HotelOrderPayRequest {
	s.ThirdTradeNo = &v
	return s
}

func (s *HotelOrderPayRequest) SetTotalPrice(v int64) *HotelOrderPayRequest {
	s.TotalPrice = &v
	return s
}

func (s *HotelOrderPayRequest) Validate() error {
	return dara.Validate(s)
}
