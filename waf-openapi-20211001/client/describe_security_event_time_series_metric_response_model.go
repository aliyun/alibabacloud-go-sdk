// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventTimeSeriesMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityEventTimeSeriesMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityEventTimeSeriesMetricResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityEventTimeSeriesMetricResponseBody) *DescribeSecurityEventTimeSeriesMetricResponse
	GetBody() *DescribeSecurityEventTimeSeriesMetricResponseBody
}

type DescribeSecurityEventTimeSeriesMetricResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityEventTimeSeriesMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityEventTimeSeriesMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTimeSeriesMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTimeSeriesMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityEventTimeSeriesMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityEventTimeSeriesMetricResponse) GetBody() *DescribeSecurityEventTimeSeriesMetricResponseBody {
	return s.Body
}

func (s *DescribeSecurityEventTimeSeriesMetricResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventTimeSeriesMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponse) SetStatusCode(v int32) *DescribeSecurityEventTimeSeriesMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponse) SetBody(v *DescribeSecurityEventTimeSeriesMetricResponseBody) *DescribeSecurityEventTimeSeriesMetricResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
