// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderChangeConfirmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeApplyId(v string) *TrainOrderChangeConfirmRequest
	GetChangeApplyId() *string
	SetChangeSettleAmount(v int64) *TrainOrderChangeConfirmRequest
	GetChangeSettleAmount() *int64
	SetOrderId(v string) *TrainOrderChangeConfirmRequest
	GetOrderId() *string
	SetOutChangeApplyId(v string) *TrainOrderChangeConfirmRequest
	GetOutChangeApplyId() *string
	SetOutOrderId(v string) *TrainOrderChangeConfirmRequest
	GetOutOrderId() *string
}

type TrainOrderChangeConfirmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	ChangeApplyId *string `json:"change_apply_id,omitempty" xml:"change_apply_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	ChangeSettleAmount *int64 `json:"change_settle_amount,omitempty" xml:"change_settle_amount,omitempty"`
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
	// 12345
	OutChangeApplyId *string `json:"out_change_apply_id,omitempty" xml:"out_change_apply_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s TrainOrderChangeConfirmRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderChangeConfirmRequest) GoString() string {
	return s.String()
}

func (s *TrainOrderChangeConfirmRequest) GetChangeApplyId() *string {
	return s.ChangeApplyId
}

func (s *TrainOrderChangeConfirmRequest) GetChangeSettleAmount() *int64 {
	return s.ChangeSettleAmount
}

func (s *TrainOrderChangeConfirmRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainOrderChangeConfirmRequest) GetOutChangeApplyId() *string {
	return s.OutChangeApplyId
}

func (s *TrainOrderChangeConfirmRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainOrderChangeConfirmRequest) SetChangeApplyId(v string) *TrainOrderChangeConfirmRequest {
	s.ChangeApplyId = &v
	return s
}

func (s *TrainOrderChangeConfirmRequest) SetChangeSettleAmount(v int64) *TrainOrderChangeConfirmRequest {
	s.ChangeSettleAmount = &v
	return s
}

func (s *TrainOrderChangeConfirmRequest) SetOrderId(v string) *TrainOrderChangeConfirmRequest {
	s.OrderId = &v
	return s
}

func (s *TrainOrderChangeConfirmRequest) SetOutChangeApplyId(v string) *TrainOrderChangeConfirmRequest {
	s.OutChangeApplyId = &v
	return s
}

func (s *TrainOrderChangeConfirmRequest) SetOutOrderId(v string) *TrainOrderChangeConfirmRequest {
	s.OutOrderId = &v
	return s
}

func (s *TrainOrderChangeConfirmRequest) Validate() error {
	return dara.Validate(s)
}
