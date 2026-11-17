// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatGatewayCapacityReservationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNatGatewayCapacityReservationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNatGatewayCapacityReservationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNatGatewayCapacityReservationResponseBody) *ModifyNatGatewayCapacityReservationResponse
	GetBody() *ModifyNatGatewayCapacityReservationResponseBody
}

type ModifyNatGatewayCapacityReservationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNatGatewayCapacityReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNatGatewayCapacityReservationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayCapacityReservationResponse) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayCapacityReservationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNatGatewayCapacityReservationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNatGatewayCapacityReservationResponse) GetBody() *ModifyNatGatewayCapacityReservationResponseBody {
	return s.Body
}

func (s *ModifyNatGatewayCapacityReservationResponse) SetHeaders(v map[string]*string) *ModifyNatGatewayCapacityReservationResponse {
	s.Headers = v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponse) SetStatusCode(v int32) *ModifyNatGatewayCapacityReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponse) SetBody(v *ModifyNatGatewayCapacityReservationResponseBody) *ModifyNatGatewayCapacityReservationResponse {
	s.Body = v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
