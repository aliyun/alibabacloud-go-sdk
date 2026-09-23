// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomAgentMonitorMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomAgentMonitorMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomAgentMonitorMetricsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomAgentMonitorMetricsResponseBody) *DescribeCustomAgentMonitorMetricsResponse
	GetBody() *DescribeCustomAgentMonitorMetricsResponseBody
}

type DescribeCustomAgentMonitorMetricsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomAgentMonitorMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomAgentMonitorMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentMonitorMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentMonitorMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomAgentMonitorMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomAgentMonitorMetricsResponse) GetBody() *DescribeCustomAgentMonitorMetricsResponseBody {
	return s.Body
}

func (s *DescribeCustomAgentMonitorMetricsResponse) SetHeaders(v map[string]*string) *DescribeCustomAgentMonitorMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponse) SetStatusCode(v int32) *DescribeCustomAgentMonitorMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponse) SetBody(v *DescribeCustomAgentMonitorMetricsResponseBody) *DescribeCustomAgentMonitorMetricsResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
