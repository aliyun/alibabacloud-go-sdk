// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainFeeCalculateChangeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeTrainDetailsShrink(v string) *TrainFeeCalculateChangeShrinkRequest
	GetChangeTrainDetailsShrink() *string
	SetDistributeOrderId(v string) *TrainFeeCalculateChangeShrinkRequest
	GetDistributeOrderId() *string
	SetOrderId(v string) *TrainFeeCalculateChangeShrinkRequest
	GetOrderId() *string
}

type TrainFeeCalculateChangeShrinkRequest struct {
	// This parameter is required.
	ChangeTrainDetailsShrink *string `json:"change_train_details,omitempty" xml:"change_train_details,omitempty"`
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
	// 2627694109810885616
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s TrainFeeCalculateChangeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateChangeShrinkRequest) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateChangeShrinkRequest) GetChangeTrainDetailsShrink() *string {
	return s.ChangeTrainDetailsShrink
}

func (s *TrainFeeCalculateChangeShrinkRequest) GetDistributeOrderId() *string {
	return s.DistributeOrderId
}

func (s *TrainFeeCalculateChangeShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainFeeCalculateChangeShrinkRequest) SetChangeTrainDetailsShrink(v string) *TrainFeeCalculateChangeShrinkRequest {
	s.ChangeTrainDetailsShrink = &v
	return s
}

func (s *TrainFeeCalculateChangeShrinkRequest) SetDistributeOrderId(v string) *TrainFeeCalculateChangeShrinkRequest {
	s.DistributeOrderId = &v
	return s
}

func (s *TrainFeeCalculateChangeShrinkRequest) SetOrderId(v string) *TrainFeeCalculateChangeShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *TrainFeeCalculateChangeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
