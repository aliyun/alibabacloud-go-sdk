// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderRefundDetailQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCooperatorOrderId(v string) *OrderRefundDetailQueryRequest
	GetCooperatorOrderId() *string
	SetOrderId(v string) *OrderRefundDetailQueryRequest
	GetOrderId() *string
}

type OrderRefundDetailQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ZLJD12241231000002
	CooperatorOrderId *string `json:"cooperator_order_id,omitempty" xml:"cooperator_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1012000000000000
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s OrderRefundDetailQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s OrderRefundDetailQueryRequest) GoString() string {
	return s.String()
}

func (s *OrderRefundDetailQueryRequest) GetCooperatorOrderId() *string {
	return s.CooperatorOrderId
}

func (s *OrderRefundDetailQueryRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *OrderRefundDetailQueryRequest) SetCooperatorOrderId(v string) *OrderRefundDetailQueryRequest {
	s.CooperatorOrderId = &v
	return s
}

func (s *OrderRefundDetailQueryRequest) SetOrderId(v string) *OrderRefundDetailQueryRequest {
	s.OrderId = &v
	return s
}

func (s *OrderRefundDetailQueryRequest) Validate() error {
	return dara.Validate(s)
}
