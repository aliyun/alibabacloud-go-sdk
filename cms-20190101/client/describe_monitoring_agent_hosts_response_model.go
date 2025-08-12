// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentHostsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitoringAgentHostsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitoringAgentHostsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitoringAgentHostsResponseBody) *DescribeMonitoringAgentHostsResponse
	GetBody() *DescribeMonitoringAgentHostsResponseBody
}

type DescribeMonitoringAgentHostsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitoringAgentHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitoringAgentHostsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentHostsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentHostsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitoringAgentHostsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitoringAgentHostsResponse) GetBody() *DescribeMonitoringAgentHostsResponseBody {
	return s.Body
}

func (s *DescribeMonitoringAgentHostsResponse) SetHeaders(v map[string]*string) *DescribeMonitoringAgentHostsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringAgentHostsResponse) SetStatusCode(v int32) *DescribeMonitoringAgentHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponse) SetBody(v *DescribeMonitoringAgentHostsResponseBody) *DescribeMonitoringAgentHostsResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitoringAgentHostsResponse) Validate() error {
	return dara.Validate(s)
}
