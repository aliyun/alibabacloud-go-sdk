// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnhealthyHostAvailabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUnhealthyHostAvailabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUnhealthyHostAvailabilityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUnhealthyHostAvailabilityResponseBody) *DescribeUnhealthyHostAvailabilityResponse
	GetBody() *DescribeUnhealthyHostAvailabilityResponseBody
}

type DescribeUnhealthyHostAvailabilityResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUnhealthyHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUnhealthyHostAvailabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUnhealthyHostAvailabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUnhealthyHostAvailabilityResponse) GetBody() *DescribeUnhealthyHostAvailabilityResponseBody {
	return s.Body
}

func (s *DescribeUnhealthyHostAvailabilityResponse) SetHeaders(v map[string]*string) *DescribeUnhealthyHostAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponse) SetStatusCode(v int32) *DescribeUnhealthyHostAvailabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponse) SetBody(v *DescribeUnhealthyHostAvailabilityResponseBody) *DescribeUnhealthyHostAvailabilityResponse {
	s.Body = v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponse) Validate() error {
	return dara.Validate(s)
}
