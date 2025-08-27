// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightPayOrderV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightPayOrderV2Request
	GetIsvName() *string
	SetOrderId(v string) *FlightPayOrderV2Request
	GetOrderId() *string
	SetOutOrderId(v string) *FlightPayOrderV2Request
	GetOutOrderId() *string
	SetTotalPrice(v int64) *FlightPayOrderV2Request
	GetTotalPrice() *int64
	SetTotalServiceFeePrice(v int64) *FlightPayOrderV2Request
	GetTotalServiceFeePrice() *int64
}

type FlightPayOrderV2Request struct {
	// example:
	//
	// cheshiapi
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// cheshiapi002kwl
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195798359400
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 5100
	TotalPrice *int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// example:
	//
	// 10
	TotalServiceFeePrice *int64 `json:"total_service_fee_price,omitempty" xml:"total_service_fee_price,omitempty"`
}

func (s FlightPayOrderV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightPayOrderV2Request) GoString() string {
	return s.String()
}

func (s *FlightPayOrderV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightPayOrderV2Request) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightPayOrderV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightPayOrderV2Request) GetTotalPrice() *int64 {
	return s.TotalPrice
}

func (s *FlightPayOrderV2Request) GetTotalServiceFeePrice() *int64 {
	return s.TotalServiceFeePrice
}

func (s *FlightPayOrderV2Request) SetIsvName(v string) *FlightPayOrderV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightPayOrderV2Request) SetOrderId(v string) *FlightPayOrderV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightPayOrderV2Request) SetOutOrderId(v string) *FlightPayOrderV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightPayOrderV2Request) SetTotalPrice(v int64) *FlightPayOrderV2Request {
	s.TotalPrice = &v
	return s
}

func (s *FlightPayOrderV2Request) SetTotalServiceFeePrice(v int64) *FlightPayOrderV2Request {
	s.TotalServiceFeePrice = &v
	return s
}

func (s *FlightPayOrderV2Request) Validate() error {
	return dara.Validate(s)
}
