// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarOrderQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *CarOrderQueryRequest
	GetOrderId() *int64
	SetSubOrderId(v int64) *CarOrderQueryRequest
	GetSubOrderId() *int64
}

type CarOrderQueryRequest struct {
	// example:
	//
	// 1012000000000000
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1000000
	SubOrderId *int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

func (s CarOrderQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s CarOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *CarOrderQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CarOrderQueryRequest) GetSubOrderId() *int64 {
	return s.SubOrderId
}

func (s *CarOrderQueryRequest) SetOrderId(v int64) *CarOrderQueryRequest {
	s.OrderId = &v
	return s
}

func (s *CarOrderQueryRequest) SetSubOrderId(v int64) *CarOrderQueryRequest {
	s.SubOrderId = &v
	return s
}

func (s *CarOrderQueryRequest) Validate() error {
	return dara.Validate(s)
}
