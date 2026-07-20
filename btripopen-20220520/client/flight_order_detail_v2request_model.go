// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderDetailV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightOrderDetailV2Request
	GetIsvName() *string
	SetOrderId(v int64) *FlightOrderDetailV2Request
	GetOrderId() *int64
	SetOutOrderId(v string) *FlightOrderDetailV2Request
	GetOutOrderId() *string
}

type FlightOrderDetailV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// cheshiapi
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 1017002195370467138
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467137
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s FlightOrderDetailV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2Request) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightOrderDetailV2Request) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightOrderDetailV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightOrderDetailV2Request) SetIsvName(v string) *FlightOrderDetailV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightOrderDetailV2Request) SetOrderId(v int64) *FlightOrderDetailV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightOrderDetailV2Request) SetOutOrderId(v string) *FlightOrderDetailV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightOrderDetailV2Request) Validate() error {
	return dara.Validate(s)
}
