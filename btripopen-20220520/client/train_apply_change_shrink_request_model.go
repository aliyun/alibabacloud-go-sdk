// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainApplyChangeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptNoSeat(v string) *TrainApplyChangeShrinkRequest
	GetAcceptNoSeat() *string
	SetChangeTrainInfoSShrink(v string) *TrainApplyChangeShrinkRequest
	GetChangeTrainInfoSShrink() *string
	SetForceMatch(v string) *TrainApplyChangeShrinkRequest
	GetForceMatch() *string
	SetIsPayNow(v bool) *TrainApplyChangeShrinkRequest
	GetIsPayNow() *bool
	SetOrderId(v string) *TrainApplyChangeShrinkRequest
	GetOrderId() *string
	SetOutChangeApplyId(v string) *TrainApplyChangeShrinkRequest
	GetOutChangeApplyId() *string
	SetOutOrderId(v string) *TrainApplyChangeShrinkRequest
	GetOutOrderId() *string
}

type TrainApplyChangeShrinkRequest struct {
	// example:
	//
	// 0
	AcceptNoSeat *string `json:"accept_no_seat,omitempty" xml:"accept_no_seat,omitempty"`
	// This parameter is required.
	ChangeTrainInfoSShrink *string `json:"change_train_info_s,omitempty" xml:"change_train_info_s,omitempty"`
	// example:
	//
	// 0
	ForceMatch *string `json:"force_match,omitempty" xml:"force_match,omitempty"`
	// example:
	//
	// false
	IsPayNow *bool `json:"is_pay_now,omitempty" xml:"is_pay_now,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1017028198411054446
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	OutChangeApplyId *string `json:"out_change_apply_id,omitempty" xml:"out_change_apply_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s TrainApplyChangeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyChangeShrinkRequest) GoString() string {
	return s.String()
}

func (s *TrainApplyChangeShrinkRequest) GetAcceptNoSeat() *string {
	return s.AcceptNoSeat
}

func (s *TrainApplyChangeShrinkRequest) GetChangeTrainInfoSShrink() *string {
	return s.ChangeTrainInfoSShrink
}

func (s *TrainApplyChangeShrinkRequest) GetForceMatch() *string {
	return s.ForceMatch
}

func (s *TrainApplyChangeShrinkRequest) GetIsPayNow() *bool {
	return s.IsPayNow
}

func (s *TrainApplyChangeShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainApplyChangeShrinkRequest) GetOutChangeApplyId() *string {
	return s.OutChangeApplyId
}

func (s *TrainApplyChangeShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainApplyChangeShrinkRequest) SetAcceptNoSeat(v string) *TrainApplyChangeShrinkRequest {
	s.AcceptNoSeat = &v
	return s
}

func (s *TrainApplyChangeShrinkRequest) SetChangeTrainInfoSShrink(v string) *TrainApplyChangeShrinkRequest {
	s.ChangeTrainInfoSShrink = &v
	return s
}

func (s *TrainApplyChangeShrinkRequest) SetForceMatch(v string) *TrainApplyChangeShrinkRequest {
	s.ForceMatch = &v
	return s
}

func (s *TrainApplyChangeShrinkRequest) SetIsPayNow(v bool) *TrainApplyChangeShrinkRequest {
	s.IsPayNow = &v
	return s
}

func (s *TrainApplyChangeShrinkRequest) SetOrderId(v string) *TrainApplyChangeShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *TrainApplyChangeShrinkRequest) SetOutChangeApplyId(v string) *TrainApplyChangeShrinkRequest {
	s.OutChangeApplyId = &v
	return s
}

func (s *TrainApplyChangeShrinkRequest) SetOutOrderId(v string) *TrainApplyChangeShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *TrainApplyChangeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
