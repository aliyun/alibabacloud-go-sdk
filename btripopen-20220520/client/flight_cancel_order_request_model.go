// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCancelOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisOrderId(v string) *FlightCancelOrderRequest
	GetDisOrderId() *string
}

type FlightCancelOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
}

func (s FlightCancelOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightCancelOrderRequest) GoString() string {
	return s.String()
}

func (s *FlightCancelOrderRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightCancelOrderRequest) SetDisOrderId(v string) *FlightCancelOrderRequest {
	s.DisOrderId = &v
	return s
}

func (s *FlightCancelOrderRequest) Validate() error {
	return dara.Validate(s)
}
