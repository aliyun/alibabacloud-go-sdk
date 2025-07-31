// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailabilityZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailabilityZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailabilityZonesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailabilityZonesResponseBody) *DescribeAvailabilityZonesResponse
	GetBody() *DescribeAvailabilityZonesResponseBody
}

type DescribeAvailabilityZonesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailabilityZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailabilityZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailabilityZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailabilityZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailabilityZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailabilityZonesResponse) GetBody() *DescribeAvailabilityZonesResponseBody {
	return s.Body
}

func (s *DescribeAvailabilityZonesResponse) SetHeaders(v map[string]*string) *DescribeAvailabilityZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailabilityZonesResponse) SetStatusCode(v int32) *DescribeAvailabilityZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailabilityZonesResponse) SetBody(v *DescribeAvailabilityZonesResponseBody) *DescribeAvailabilityZonesResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailabilityZonesResponse) Validate() error {
	return dara.Validate(s)
}
