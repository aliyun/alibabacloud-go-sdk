// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyPayV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetExtParams(v map[string]interface{}) *FlightModifyPayV2Request
	GetExtParams() map[string]interface{}
	SetIsvName(v string) *FlightModifyPayV2Request
	GetIsvName() *string
	SetModifyPayAmount(v int64) *FlightModifyPayV2Request
	GetModifyPayAmount() *int64
	SetOrderId(v int64) *FlightModifyPayV2Request
	GetOrderId() *int64
	SetOutOrderId(v string) *FlightModifyPayV2Request
	GetOutOrderId() *string
	SetOutSubOrderId(v string) *FlightModifyPayV2Request
	GetOutSubOrderId() *string
	SetSubOrderId(v int64) *FlightModifyPayV2Request
	GetSubOrderId() *int64
}

type FlightModifyPayV2Request struct {
	ExtParams map[string]interface{} `json:"ext_params,omitempty" xml:"ext_params,omitempty"`
	// example:
	//
	// name
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 5100
	ModifyPayAmount *int64 `json:"modify_pay_amount,omitempty" xml:"modify_pay_amount,omitempty"`
	// example:
	//
	// 1017002195370467200
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467200
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 1019195786853020
	OutSubOrderId *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// example:
	//
	// 1019195786853020
	SubOrderId *int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

func (s FlightModifyPayV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyPayV2Request) GoString() string {
	return s.String()
}

func (s *FlightModifyPayV2Request) GetExtParams() map[string]interface{} {
	return s.ExtParams
}

func (s *FlightModifyPayV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightModifyPayV2Request) GetModifyPayAmount() *int64 {
	return s.ModifyPayAmount
}

func (s *FlightModifyPayV2Request) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightModifyPayV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyPayV2Request) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *FlightModifyPayV2Request) GetSubOrderId() *int64 {
	return s.SubOrderId
}

func (s *FlightModifyPayV2Request) SetExtParams(v map[string]interface{}) *FlightModifyPayV2Request {
	s.ExtParams = v
	return s
}

func (s *FlightModifyPayV2Request) SetIsvName(v string) *FlightModifyPayV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightModifyPayV2Request) SetModifyPayAmount(v int64) *FlightModifyPayV2Request {
	s.ModifyPayAmount = &v
	return s
}

func (s *FlightModifyPayV2Request) SetOrderId(v int64) *FlightModifyPayV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightModifyPayV2Request) SetOutOrderId(v string) *FlightModifyPayV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyPayV2Request) SetOutSubOrderId(v string) *FlightModifyPayV2Request {
	s.OutSubOrderId = &v
	return s
}

func (s *FlightModifyPayV2Request) SetSubOrderId(v int64) *FlightModifyPayV2Request {
	s.SubOrderId = &v
	return s
}

func (s *FlightModifyPayV2Request) Validate() error {
	return dara.Validate(s)
}
