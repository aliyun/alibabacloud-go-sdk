// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderCancelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *IntlFlightOrderCancelRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightOrderCancelRequest
	GetBuyerName() *string
	SetIsvName(v string) *IntlFlightOrderCancelRequest
	GetIsvName() *string
	SetOrderId(v string) *IntlFlightOrderCancelRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightOrderCancelRequest
	GetOutOrderId() *string
}

type IntlFlightOrderCancelRequest struct {
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
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s IntlFlightOrderCancelRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderCancelRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderCancelRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightOrderCancelRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightOrderCancelRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightOrderCancelRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightOrderCancelRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightOrderCancelRequest) SetBtripUserId(v string) *IntlFlightOrderCancelRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightOrderCancelRequest) SetBuyerName(v string) *IntlFlightOrderCancelRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightOrderCancelRequest) SetIsvName(v string) *IntlFlightOrderCancelRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightOrderCancelRequest) SetOrderId(v string) *IntlFlightOrderCancelRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightOrderCancelRequest) SetOutOrderId(v string) *IntlFlightOrderCancelRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightOrderCancelRequest) Validate() error {
	return dara.Validate(s)
}
