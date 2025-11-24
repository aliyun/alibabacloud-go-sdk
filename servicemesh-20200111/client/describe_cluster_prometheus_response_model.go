// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterPrometheusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterPrometheusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterPrometheusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterPrometheusResponseBody) *DescribeClusterPrometheusResponse
	GetBody() *DescribeClusterPrometheusResponseBody
}

type DescribeClusterPrometheusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterPrometheusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterPrometheusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterPrometheusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterPrometheusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterPrometheusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterPrometheusResponse) GetBody() *DescribeClusterPrometheusResponseBody {
	return s.Body
}

func (s *DescribeClusterPrometheusResponse) SetHeaders(v map[string]*string) *DescribeClusterPrometheusResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterPrometheusResponse) SetStatusCode(v int32) *DescribeClusterPrometheusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterPrometheusResponse) SetBody(v *DescribeClusterPrometheusResponseBody) *DescribeClusterPrometheusResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterPrometheusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
