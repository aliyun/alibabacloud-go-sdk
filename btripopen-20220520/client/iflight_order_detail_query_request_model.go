// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIFlightOrderDetailQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *IFlightOrderDetailQueryRequest
	GetOrderId() *int64
}

type IFlightOrderDetailQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1017035199702438072
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s IFlightOrderDetailQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryRequest) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *IFlightOrderDetailQueryRequest) SetOrderId(v int64) *IFlightOrderDetailQueryRequest {
	s.OrderId = &v
	return s
}

func (s *IFlightOrderDetailQueryRequest) Validate() error {
	return dara.Validate(s)
}
