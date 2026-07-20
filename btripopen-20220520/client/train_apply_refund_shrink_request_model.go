// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainApplyRefundShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *TrainApplyRefundShrinkRequest
	GetOrderId() *string
	SetOutOrderId(v string) *TrainApplyRefundShrinkRequest
	GetOutOrderId() *string
	SetOutRefundId(v string) *TrainApplyRefundShrinkRequest
	GetOutRefundId() *string
	SetRefundTrainInfosShrink(v string) *TrainApplyRefundShrinkRequest
	GetRefundTrainInfosShrink() *string
}

type TrainApplyRefundShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1002145190081005400
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456778
	OutRefundId *string `json:"out_refund_id,omitempty" xml:"out_refund_id,omitempty"`
	// This parameter is required.
	RefundTrainInfosShrink *string `json:"refund_train_infos,omitempty" xml:"refund_train_infos,omitempty"`
}

func (s TrainApplyRefundShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyRefundShrinkRequest) GoString() string {
	return s.String()
}

func (s *TrainApplyRefundShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainApplyRefundShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainApplyRefundShrinkRequest) GetOutRefundId() *string {
	return s.OutRefundId
}

func (s *TrainApplyRefundShrinkRequest) GetRefundTrainInfosShrink() *string {
	return s.RefundTrainInfosShrink
}

func (s *TrainApplyRefundShrinkRequest) SetOrderId(v string) *TrainApplyRefundShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *TrainApplyRefundShrinkRequest) SetOutOrderId(v string) *TrainApplyRefundShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *TrainApplyRefundShrinkRequest) SetOutRefundId(v string) *TrainApplyRefundShrinkRequest {
	s.OutRefundId = &v
	return s
}

func (s *TrainApplyRefundShrinkRequest) SetRefundTrainInfosShrink(v string) *TrainApplyRefundShrinkRequest {
	s.RefundTrainInfosShrink = &v
	return s
}

func (s *TrainApplyRefundShrinkRequest) Validate() error {
	return dara.Validate(s)
}
