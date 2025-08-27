// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderPayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *TrainOrderPayRequest
	GetOrderId() *string
	SetOutOrderId(v string) *TrainOrderPayRequest
	GetOutOrderId() *string
	SetPayAmount(v int64) *TrainOrderPayRequest
	GetPayAmount() *int64
}

type TrainOrderPayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1379598062646
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
	// 100
	PayAmount *int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
}

func (s TrainOrderPayRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderPayRequest) GoString() string {
	return s.String()
}

func (s *TrainOrderPayRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainOrderPayRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainOrderPayRequest) GetPayAmount() *int64 {
	return s.PayAmount
}

func (s *TrainOrderPayRequest) SetOrderId(v string) *TrainOrderPayRequest {
	s.OrderId = &v
	return s
}

func (s *TrainOrderPayRequest) SetOutOrderId(v string) *TrainOrderPayRequest {
	s.OutOrderId = &v
	return s
}

func (s *TrainOrderPayRequest) SetPayAmount(v int64) *TrainOrderPayRequest {
	s.PayAmount = &v
	return s
}

func (s *TrainOrderPayRequest) Validate() error {
	return dara.Validate(s)
}
