// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCancelOrderV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightCancelOrderV2Request
	GetIsvName() *string
	SetOrderId(v string) *FlightCancelOrderV2Request
	GetOrderId() *string
	SetOutOrderId(v string) *FlightCancelOrderV2Request
	GetOutOrderId() *string
}

type FlightCancelOrderV2Request struct {
	// This parameter is required.
	IsvName    *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	OrderId    *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s FlightCancelOrderV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightCancelOrderV2Request) GoString() string {
	return s.String()
}

func (s *FlightCancelOrderV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightCancelOrderV2Request) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightCancelOrderV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightCancelOrderV2Request) SetIsvName(v string) *FlightCancelOrderV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightCancelOrderV2Request) SetOrderId(v string) *FlightCancelOrderV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightCancelOrderV2Request) SetOutOrderId(v string) *FlightCancelOrderV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightCancelOrderV2Request) Validate() error {
	return dara.Validate(s)
}
