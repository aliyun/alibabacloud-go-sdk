// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderPayCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *IntlFlightOrderPayCheckRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightOrderPayCheckRequest
	GetBuyerName() *string
	SetIsvName(v string) *IntlFlightOrderPayCheckRequest
	GetIsvName() *string
	SetOrderId(v string) *IntlFlightOrderPayCheckRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightOrderPayCheckRequest
	GetOutOrderId() *string
}

type IntlFlightOrderPayCheckRequest struct {
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
	// 1002094194679957528
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// F11374007131319304192
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s IntlFlightOrderPayCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderPayCheckRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderPayCheckRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightOrderPayCheckRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightOrderPayCheckRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightOrderPayCheckRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightOrderPayCheckRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightOrderPayCheckRequest) SetBtripUserId(v string) *IntlFlightOrderPayCheckRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightOrderPayCheckRequest) SetBuyerName(v string) *IntlFlightOrderPayCheckRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightOrderPayCheckRequest) SetIsvName(v string) *IntlFlightOrderPayCheckRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightOrderPayCheckRequest) SetOrderId(v string) *IntlFlightOrderPayCheckRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightOrderPayCheckRequest) SetOutOrderId(v string) *IntlFlightOrderPayCheckRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightOrderPayCheckRequest) Validate() error {
	return dara.Validate(s)
}
