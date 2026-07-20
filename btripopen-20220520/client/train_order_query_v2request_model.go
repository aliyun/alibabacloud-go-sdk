// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderQueryV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *TrainOrderQueryV2Request
	GetOrderId() *int64
	SetUserId(v string) *TrainOrderQueryV2Request
	GetUserId() *string
}

type TrainOrderQueryV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 12342123212
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 123121112
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderQueryV2Request) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2Request) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2Request) GetOrderId() *int64 {
	return s.OrderId
}

func (s *TrainOrderQueryV2Request) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderQueryV2Request) SetOrderId(v int64) *TrainOrderQueryV2Request {
	s.OrderId = &v
	return s
}

func (s *TrainOrderQueryV2Request) SetUserId(v string) *TrainOrderQueryV2Request {
	s.UserId = &v
	return s
}

func (s *TrainOrderQueryV2Request) Validate() error {
	return dara.Validate(s)
}
