// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapacityReservationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCapacityReservationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCapacityReservationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCapacityReservationResponseBody) *DescribeCapacityReservationResponse
	GetBody() *DescribeCapacityReservationResponseBody
}

type DescribeCapacityReservationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCapacityReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCapacityReservationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationResponse) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCapacityReservationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCapacityReservationResponse) GetBody() *DescribeCapacityReservationResponseBody {
	return s.Body
}

func (s *DescribeCapacityReservationResponse) SetHeaders(v map[string]*string) *DescribeCapacityReservationResponse {
	s.Headers = v
	return s
}

func (s *DescribeCapacityReservationResponse) SetStatusCode(v int32) *DescribeCapacityReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCapacityReservationResponse) SetBody(v *DescribeCapacityReservationResponseBody) *DescribeCapacityReservationResponse {
	s.Body = v
	return s
}

func (s *DescribeCapacityReservationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
