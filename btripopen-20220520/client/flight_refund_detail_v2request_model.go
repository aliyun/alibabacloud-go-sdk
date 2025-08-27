// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundDetailV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightRefundDetailV2Request
	GetIsvName() *string
	SetOrderId(v string) *FlightRefundDetailV2Request
	GetOrderId() *string
	SetOutOrderId(v string) *FlightRefundDetailV2Request
	GetOutOrderId() *string
	SetOutRefundApplyId(v string) *FlightRefundDetailV2Request
	GetOutRefundApplyId() *string
	SetRefundApplyId(v string) *FlightRefundDetailV2Request
	GetRefundApplyId() *string
}

type FlightRefundDetailV2Request struct {
	// example:
	//
	// cheshiapi
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 1002039195025156784
	OrderId          *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutOrderId       *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	OutRefundApplyId *string `json:"out_refund_apply_id,omitempty" xml:"out_refund_apply_id,omitempty"`
	// example:
	//
	// 1002039195025156700
	RefundApplyId *string `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
}

func (s FlightRefundDetailV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailV2Request) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightRefundDetailV2Request) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightRefundDetailV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightRefundDetailV2Request) GetOutRefundApplyId() *string {
	return s.OutRefundApplyId
}

func (s *FlightRefundDetailV2Request) GetRefundApplyId() *string {
	return s.RefundApplyId
}

func (s *FlightRefundDetailV2Request) SetIsvName(v string) *FlightRefundDetailV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightRefundDetailV2Request) SetOrderId(v string) *FlightRefundDetailV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightRefundDetailV2Request) SetOutOrderId(v string) *FlightRefundDetailV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightRefundDetailV2Request) SetOutRefundApplyId(v string) *FlightRefundDetailV2Request {
	s.OutRefundApplyId = &v
	return s
}

func (s *FlightRefundDetailV2Request) SetRefundApplyId(v string) *FlightRefundDetailV2Request {
	s.RefundApplyId = &v
	return s
}

func (s *FlightRefundDetailV2Request) Validate() error {
	return dara.Validate(s)
}
