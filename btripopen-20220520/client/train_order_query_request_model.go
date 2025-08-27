// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *TrainOrderQueryRequest
	GetOrderId() *int64
	SetUserId(v string) *TrainOrderQueryRequest
	GetUserId() *string
}

type TrainOrderQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2627694109810885616
	OrderId *int64  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	UserId  *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *TrainOrderQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderQueryRequest) SetOrderId(v int64) *TrainOrderQueryRequest {
	s.OrderId = &v
	return s
}

func (s *TrainOrderQueryRequest) SetUserId(v string) *TrainOrderQueryRequest {
	s.UserId = &v
	return s
}

func (s *TrainOrderQueryRequest) Validate() error {
	return dara.Validate(s)
}
