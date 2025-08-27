// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *FlightOrderQueryRequest
	GetOrderId() *int64
	SetUserId(v string) *FlightOrderQueryRequest
	GetUserId() *string
}

type FlightOrderQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 246584390
	OrderId *int64  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	UserId  *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightOrderQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderQueryRequest) SetOrderId(v int64) *FlightOrderQueryRequest {
	s.OrderId = &v
	return s
}

func (s *FlightOrderQueryRequest) SetUserId(v string) *FlightOrderQueryRequest {
	s.UserId = &v
	return s
}

func (s *FlightOrderQueryRequest) Validate() error {
	return dara.Validate(s)
}
