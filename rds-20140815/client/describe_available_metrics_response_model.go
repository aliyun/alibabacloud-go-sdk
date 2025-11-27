// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableMetricsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableMetricsResponseBody) *DescribeAvailableMetricsResponse
	GetBody() *DescribeAvailableMetricsResponseBody
}

type DescribeAvailableMetricsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableMetricsResponse) GetBody() *DescribeAvailableMetricsResponseBody {
	return s.Body
}

func (s *DescribeAvailableMetricsResponse) SetHeaders(v map[string]*string) *DescribeAvailableMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableMetricsResponse) SetStatusCode(v int32) *DescribeAvailableMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableMetricsResponse) SetBody(v *DescribeAvailableMetricsResponseBody) *DescribeAvailableMetricsResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
