// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventTopNMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityEventTopNMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityEventTopNMetricResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityEventTopNMetricResponseBody) *DescribeSecurityEventTopNMetricResponse
	GetBody() *DescribeSecurityEventTopNMetricResponseBody
}

type DescribeSecurityEventTopNMetricResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityEventTopNMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityEventTopNMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTopNMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTopNMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityEventTopNMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityEventTopNMetricResponse) GetBody() *DescribeSecurityEventTopNMetricResponseBody {
	return s.Body
}

func (s *DescribeSecurityEventTopNMetricResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventTopNMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponse) SetStatusCode(v int32) *DescribeSecurityEventTopNMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponse) SetBody(v *DescribeSecurityEventTopNMetricResponseBody) *DescribeSecurityEventTopNMetricResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponse) Validate() error {
	return dara.Validate(s)
}
