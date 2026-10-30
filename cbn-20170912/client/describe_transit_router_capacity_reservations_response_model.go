// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterCapacityReservationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTransitRouterCapacityReservationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTransitRouterCapacityReservationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTransitRouterCapacityReservationsResponseBody) *DescribeTransitRouterCapacityReservationsResponse
	GetBody() *DescribeTransitRouterCapacityReservationsResponseBody
}

type DescribeTransitRouterCapacityReservationsResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransitRouterCapacityReservationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTransitRouterCapacityReservationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterCapacityReservationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterCapacityReservationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTransitRouterCapacityReservationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTransitRouterCapacityReservationsResponse) GetBody() *DescribeTransitRouterCapacityReservationsResponseBody {
	return s.Body
}

func (s *DescribeTransitRouterCapacityReservationsResponse) SetHeaders(v map[string]*string) *DescribeTransitRouterCapacityReservationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponse) SetStatusCode(v int32) *DescribeTransitRouterCapacityReservationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponse) SetBody(v *DescribeTransitRouterCapacityReservationsResponseBody) *DescribeTransitRouterCapacityReservationsResponse {
	s.Body = v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
