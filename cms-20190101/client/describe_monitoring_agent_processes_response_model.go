// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentProcessesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitoringAgentProcessesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitoringAgentProcessesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitoringAgentProcessesResponseBody) *DescribeMonitoringAgentProcessesResponse
	GetBody() *DescribeMonitoringAgentProcessesResponseBody
}

type DescribeMonitoringAgentProcessesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitoringAgentProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitoringAgentProcessesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentProcessesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentProcessesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitoringAgentProcessesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitoringAgentProcessesResponse) GetBody() *DescribeMonitoringAgentProcessesResponseBody {
	return s.Body
}

func (s *DescribeMonitoringAgentProcessesResponse) SetHeaders(v map[string]*string) *DescribeMonitoringAgentProcessesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponse) SetStatusCode(v int32) *DescribeMonitoringAgentProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponse) SetBody(v *DescribeMonitoringAgentProcessesResponseBody) *DescribeMonitoringAgentProcessesResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponse) Validate() error {
	return dara.Validate(s)
}
