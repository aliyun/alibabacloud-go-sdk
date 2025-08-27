// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisOrderId(v string) *FlightRefundDetailRequest
	GetDisOrderId() *string
	SetDisSubOrderId(v string) *FlightRefundDetailRequest
	GetDisSubOrderId() *string
}

type FlightRefundDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// refun123
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
}

func (s FlightRefundDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailRequest) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightRefundDetailRequest) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *FlightRefundDetailRequest) SetDisOrderId(v string) *FlightRefundDetailRequest {
	s.DisOrderId = &v
	return s
}

func (s *FlightRefundDetailRequest) SetDisSubOrderId(v string) *FlightRefundDetailRequest {
	s.DisSubOrderId = &v
	return s
}

func (s *FlightRefundDetailRequest) Validate() error {
	return dara.Validate(s)
}
