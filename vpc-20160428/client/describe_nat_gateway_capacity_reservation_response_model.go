// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewayCapacityReservationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatGatewayCapacityReservationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatGatewayCapacityReservationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatGatewayCapacityReservationResponseBody) *DescribeNatGatewayCapacityReservationResponse
	GetBody() *DescribeNatGatewayCapacityReservationResponseBody
}

type DescribeNatGatewayCapacityReservationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatGatewayCapacityReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatGatewayCapacityReservationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayCapacityReservationResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayCapacityReservationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatGatewayCapacityReservationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatGatewayCapacityReservationResponse) GetBody() *DescribeNatGatewayCapacityReservationResponseBody {
	return s.Body
}

func (s *DescribeNatGatewayCapacityReservationResponse) SetHeaders(v map[string]*string) *DescribeNatGatewayCapacityReservationResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponse) SetStatusCode(v int32) *DescribeNatGatewayCapacityReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponse) SetBody(v *DescribeNatGatewayCapacityReservationResponseBody) *DescribeNatGatewayCapacityReservationResponse {
	s.Body = v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
