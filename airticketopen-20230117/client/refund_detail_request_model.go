// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRefundOrderNum(v int64) *RefundDetailRequest
	GetRefundOrderNum() *int64
}

type RefundDetailRequest struct {
	// refund order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 4966***617732
	RefundOrderNum *int64 `json:"refund_order_num,omitempty" xml:"refund_order_num,omitempty"`
}

func (s RefundDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailRequest) GoString() string {
	return s.String()
}

func (s *RefundDetailRequest) GetRefundOrderNum() *int64 {
	return s.RefundOrderNum
}

func (s *RefundDetailRequest) SetRefundOrderNum(v int64) *RefundDetailRequest {
	s.RefundOrderNum = &v
	return s
}

func (s *RefundDetailRequest) Validate() error {
	return dara.Validate(s)
}
