// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopConsultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightReShopConsultRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightReShopConsultRequest
	GetOutOrderId() *string
}

type IntlFlightReShopConsultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1003038202430742196
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// JP2024072600000006
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s IntlFlightReShopConsultRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopConsultRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopConsultRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopConsultRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopConsultRequest) SetOrderId(v string) *IntlFlightReShopConsultRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopConsultRequest) SetOutOrderId(v string) *IntlFlightReShopConsultRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopConsultRequest) Validate() error {
	return dara.Validate(s)
}
