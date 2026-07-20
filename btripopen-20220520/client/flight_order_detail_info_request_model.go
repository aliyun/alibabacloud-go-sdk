// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderDetailInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisOrderId(v string) *FlightOrderDetailInfoRequest
	GetDisOrderId() *string
}

type FlightOrderDetailInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
}

func (s FlightOrderDetailInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailInfoRequest) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailInfoRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightOrderDetailInfoRequest) SetDisOrderId(v string) *FlightOrderDetailInfoRequest {
	s.DisOrderId = &v
	return s
}

func (s *FlightOrderDetailInfoRequest) Validate() error {
	return dara.Validate(s)
}
