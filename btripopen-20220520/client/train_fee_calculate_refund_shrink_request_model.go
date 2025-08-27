// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainFeeCalculateRefundShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDistributeOrderId(v string) *TrainFeeCalculateRefundShrinkRequest
	GetDistributeOrderId() *string
	SetOrderId(v string) *TrainFeeCalculateRefundShrinkRequest
	GetOrderId() *string
	SetRefundTrainInfosShrink(v string) *TrainFeeCalculateRefundShrinkRequest
	GetRefundTrainInfosShrink() *string
}

type TrainFeeCalculateRefundShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DistributeOrderId *string `json:"distribute_order_id,omitempty" xml:"distribute_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1002145190081005400
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	RefundTrainInfosShrink *string `json:"refund_train_infos,omitempty" xml:"refund_train_infos,omitempty"`
}

func (s TrainFeeCalculateRefundShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateRefundShrinkRequest) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateRefundShrinkRequest) GetDistributeOrderId() *string {
	return s.DistributeOrderId
}

func (s *TrainFeeCalculateRefundShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainFeeCalculateRefundShrinkRequest) GetRefundTrainInfosShrink() *string {
	return s.RefundTrainInfosShrink
}

func (s *TrainFeeCalculateRefundShrinkRequest) SetDistributeOrderId(v string) *TrainFeeCalculateRefundShrinkRequest {
	s.DistributeOrderId = &v
	return s
}

func (s *TrainFeeCalculateRefundShrinkRequest) SetOrderId(v string) *TrainFeeCalculateRefundShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *TrainFeeCalculateRefundShrinkRequest) SetRefundTrainInfosShrink(v string) *TrainFeeCalculateRefundShrinkRequest {
	s.RefundTrainInfosShrink = &v
	return s
}

func (s *TrainFeeCalculateRefundShrinkRequest) Validate() error {
	return dara.Validate(s)
}
