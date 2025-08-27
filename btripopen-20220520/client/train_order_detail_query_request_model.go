// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderDetailQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *TrainOrderDetailQueryRequest
	GetOrderId() *string
	SetOutOrderId(v string) *TrainOrderDetailQueryRequest
	GetOutOrderId() *string
}

type TrainOrderDetailQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2083528200659337994
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3703190607180169216
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s TrainOrderDetailQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryRequest) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainOrderDetailQueryRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainOrderDetailQueryRequest) SetOrderId(v string) *TrainOrderDetailQueryRequest {
	s.OrderId = &v
	return s
}

func (s *TrainOrderDetailQueryRequest) SetOutOrderId(v string) *TrainOrderDetailQueryRequest {
	s.OutOrderId = &v
	return s
}

func (s *TrainOrderDetailQueryRequest) Validate() error {
	return dara.Validate(s)
}
