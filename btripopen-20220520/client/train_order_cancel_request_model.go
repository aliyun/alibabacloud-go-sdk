// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderCancelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *TrainOrderCancelRequest
	GetChangeOrderId() *string
	SetOrderId(v string) *TrainOrderCancelRequest
	GetOrderId() *string
	SetOutChangeOrderId(v string) *TrainOrderCancelRequest
	GetOutChangeOrderId() *string
	SetOutOrderId(v string) *TrainOrderCancelRequest
	GetOutOrderId() *string
}

type TrainOrderCancelRequest struct {
	// example:
	//
	// 1234223
	ChangeOrderId *string `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1002145190081005400
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 123321245534
	OutChangeOrderId *string `json:"out_change_order_id,omitempty" xml:"out_change_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s TrainOrderCancelRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCancelRequest) GoString() string {
	return s.String()
}

func (s *TrainOrderCancelRequest) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *TrainOrderCancelRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainOrderCancelRequest) GetOutChangeOrderId() *string {
	return s.OutChangeOrderId
}

func (s *TrainOrderCancelRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainOrderCancelRequest) SetChangeOrderId(v string) *TrainOrderCancelRequest {
	s.ChangeOrderId = &v
	return s
}

func (s *TrainOrderCancelRequest) SetOrderId(v string) *TrainOrderCancelRequest {
	s.OrderId = &v
	return s
}

func (s *TrainOrderCancelRequest) SetOutChangeOrderId(v string) *TrainOrderCancelRequest {
	s.OutChangeOrderId = &v
	return s
}

func (s *TrainOrderCancelRequest) SetOutOrderId(v string) *TrainOrderCancelRequest {
	s.OutOrderId = &v
	return s
}

func (s *TrainOrderCancelRequest) Validate() error {
	return dara.Validate(s)
}
