// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *IntlFlightOrderDetailRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightOrderDetailRequest
	GetBuyerName() *string
	SetIsvName(v string) *IntlFlightOrderDetailRequest
	GetIsvName() *string
	SetOrderId(v string) *IntlFlightOrderDetailRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightOrderDetailRequest
	GetOutOrderId() *string
}

type IntlFlightOrderDetailRequest struct {
	// example:
	//
	// 10001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// example:
	//
	// ZHANG/SAN
	BuyerName *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// example:
	//
	// TRAVEL
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1012000000000000
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// F11374007131319304192
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s IntlFlightOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightOrderDetailRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightOrderDetailRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightOrderDetailRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightOrderDetailRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightOrderDetailRequest) SetBtripUserId(v string) *IntlFlightOrderDetailRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightOrderDetailRequest) SetBuyerName(v string) *IntlFlightOrderDetailRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightOrderDetailRequest) SetIsvName(v string) *IntlFlightOrderDetailRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightOrderDetailRequest) SetOrderId(v string) *IntlFlightOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightOrderDetailRequest) SetOutOrderId(v string) *IntlFlightOrderDetailRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
