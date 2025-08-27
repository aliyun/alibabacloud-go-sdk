// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightRefundDetailRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightRefundDetailRequest
	GetOutOrderId() *string
	SetOutRefundApplyId(v string) *IntlFlightRefundDetailRequest
	GetOutRefundApplyId() *string
	SetRefundApplyId(v string) *IntlFlightRefundDetailRequest
	GetRefundApplyId() *string
}

type IntlFlightRefundDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1002145190081005400
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 40820102379649052
	OutOrderId       *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	OutRefundApplyId *string `json:"out_refund_apply_id,omitempty" xml:"out_refund_apply_id,omitempty"`
	// This parameter is required.
	RefundApplyId *string `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
}

func (s IntlFlightRefundDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightRefundDetailRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightRefundDetailRequest) GetOutRefundApplyId() *string {
	return s.OutRefundApplyId
}

func (s *IntlFlightRefundDetailRequest) GetRefundApplyId() *string {
	return s.RefundApplyId
}

func (s *IntlFlightRefundDetailRequest) SetOrderId(v string) *IntlFlightRefundDetailRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightRefundDetailRequest) SetOutOrderId(v string) *IntlFlightRefundDetailRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightRefundDetailRequest) SetOutRefundApplyId(v string) *IntlFlightRefundDetailRequest {
	s.OutRefundApplyId = &v
	return s
}

func (s *IntlFlightRefundDetailRequest) SetRefundApplyId(v string) *IntlFlightRefundDetailRequest {
	s.RefundApplyId = &v
	return s
}

func (s *IntlFlightRefundDetailRequest) Validate() error {
	return dara.Validate(s)
}
