// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderPayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *IntlFlightOrderPayRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightOrderPayRequest
	GetBuyerName() *string
	SetIsvName(v string) *IntlFlightOrderPayRequest
	GetIsvName() *string
	SetOrderId(v string) *IntlFlightOrderPayRequest
	GetOrderId() *string
	SetOrderPrice(v int64) *IntlFlightOrderPayRequest
	GetOrderPrice() *int64
	SetOutOrderId(v string) *IntlFlightOrderPayRequest
	GetOutOrderId() *string
}

type IntlFlightOrderPayRequest struct {
	// example:
	//
	// 10001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// example:
	//
	// ZHANG/SAN
	BuyerName *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	IsvName   *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// This parameter is required.
	OrderId    *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderPrice *int64  `json:"order_price,omitempty" xml:"order_price,omitempty"`
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s IntlFlightOrderPayRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderPayRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderPayRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightOrderPayRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightOrderPayRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightOrderPayRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightOrderPayRequest) GetOrderPrice() *int64 {
	return s.OrderPrice
}

func (s *IntlFlightOrderPayRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightOrderPayRequest) SetBtripUserId(v string) *IntlFlightOrderPayRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightOrderPayRequest) SetBuyerName(v string) *IntlFlightOrderPayRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightOrderPayRequest) SetIsvName(v string) *IntlFlightOrderPayRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightOrderPayRequest) SetOrderId(v string) *IntlFlightOrderPayRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightOrderPayRequest) SetOrderPrice(v int64) *IntlFlightOrderPayRequest {
	s.OrderPrice = &v
	return s
}

func (s *IntlFlightOrderPayRequest) SetOutOrderId(v string) *IntlFlightOrderPayRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightOrderPayRequest) Validate() error {
	return dara.Validate(s)
}
