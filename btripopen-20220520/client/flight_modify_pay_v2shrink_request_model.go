// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyPayV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtParamsShrink(v string) *FlightModifyPayV2ShrinkRequest
	GetExtParamsShrink() *string
	SetIsvName(v string) *FlightModifyPayV2ShrinkRequest
	GetIsvName() *string
	SetModifyPayAmount(v int64) *FlightModifyPayV2ShrinkRequest
	GetModifyPayAmount() *int64
	SetOrderId(v int64) *FlightModifyPayV2ShrinkRequest
	GetOrderId() *int64
	SetOutOrderId(v string) *FlightModifyPayV2ShrinkRequest
	GetOutOrderId() *string
	SetOutSubOrderId(v string) *FlightModifyPayV2ShrinkRequest
	GetOutSubOrderId() *string
	SetSubOrderId(v int64) *FlightModifyPayV2ShrinkRequest
	GetSubOrderId() *int64
}

type FlightModifyPayV2ShrinkRequest struct {
	ExtParamsShrink *string `json:"ext_params,omitempty" xml:"ext_params,omitempty"`
	IsvName         *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	ModifyPayAmount *int64  `json:"modify_pay_amount,omitempty" xml:"modify_pay_amount,omitempty"`
	OrderId         *int64  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutOrderId      *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	OutSubOrderId   *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	SubOrderId      *int64  `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

func (s FlightModifyPayV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyPayV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightModifyPayV2ShrinkRequest) GetExtParamsShrink() *string {
	return s.ExtParamsShrink
}

func (s *FlightModifyPayV2ShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightModifyPayV2ShrinkRequest) GetModifyPayAmount() *int64 {
	return s.ModifyPayAmount
}

func (s *FlightModifyPayV2ShrinkRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightModifyPayV2ShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyPayV2ShrinkRequest) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *FlightModifyPayV2ShrinkRequest) GetSubOrderId() *int64 {
	return s.SubOrderId
}

func (s *FlightModifyPayV2ShrinkRequest) SetExtParamsShrink(v string) *FlightModifyPayV2ShrinkRequest {
	s.ExtParamsShrink = &v
	return s
}

func (s *FlightModifyPayV2ShrinkRequest) SetIsvName(v string) *FlightModifyPayV2ShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *FlightModifyPayV2ShrinkRequest) SetModifyPayAmount(v int64) *FlightModifyPayV2ShrinkRequest {
	s.ModifyPayAmount = &v
	return s
}

func (s *FlightModifyPayV2ShrinkRequest) SetOrderId(v int64) *FlightModifyPayV2ShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *FlightModifyPayV2ShrinkRequest) SetOutOrderId(v string) *FlightModifyPayV2ShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyPayV2ShrinkRequest) SetOutSubOrderId(v string) *FlightModifyPayV2ShrinkRequest {
	s.OutSubOrderId = &v
	return s
}

func (s *FlightModifyPayV2ShrinkRequest) SetSubOrderId(v int64) *FlightModifyPayV2ShrinkRequest {
	s.SubOrderId = &v
	return s
}

func (s *FlightModifyPayV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
