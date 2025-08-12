// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupMonitoringAgentProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupMonitoringAgentProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupMonitoringAgentProcessResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupMonitoringAgentProcessResponseBody) *DescribeGroupMonitoringAgentProcessResponse
	GetBody() *DescribeGroupMonitoringAgentProcessResponseBody
}

type DescribeGroupMonitoringAgentProcessResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupMonitoringAgentProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupMonitoringAgentProcessResponse) GetBody() *DescribeGroupMonitoringAgentProcessResponseBody {
	return s.Body
}

func (s *DescribeGroupMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *DescribeGroupMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponse) SetStatusCode(v int32) *DescribeGroupMonitoringAgentProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponse) SetBody(v *DescribeGroupMonitoringAgentProcessResponseBody) *DescribeGroupMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponse) Validate() error {
	return dara.Validate(s)
}
