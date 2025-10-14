// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightChangeOfOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderNum(v int64) *FlightChangeOfOrderRequest
	GetOrderNum() *int64
}

type FlightChangeOfOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s FlightChangeOfOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightChangeOfOrderRequest) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *FlightChangeOfOrderRequest) SetOrderNum(v int64) *FlightChangeOfOrderRequest {
	s.OrderNum = &v
	return s
}

func (s *FlightChangeOfOrderRequest) Validate() error {
	return dara.Validate(s)
}
