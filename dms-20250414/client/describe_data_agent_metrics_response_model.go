// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataAgentMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataAgentMetricsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataAgentMetricsResponseBody) *DescribeDataAgentMetricsResponse
	GetBody() *DescribeDataAgentMetricsResponseBody
}

type DescribeDataAgentMetricsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataAgentMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataAgentMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataAgentMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataAgentMetricsResponse) GetBody() *DescribeDataAgentMetricsResponseBody {
	return s.Body
}

func (s *DescribeDataAgentMetricsResponse) SetHeaders(v map[string]*string) *DescribeDataAgentMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataAgentMetricsResponse) SetStatusCode(v int32) *DescribeDataAgentMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataAgentMetricsResponse) SetBody(v *DescribeDataAgentMetricsResponseBody) *DescribeDataAgentMetricsResponse {
	s.Body = v
	return s
}

func (s *DescribeDataAgentMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
