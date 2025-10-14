// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentStatusesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitoringAgentStatusesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitoringAgentStatusesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitoringAgentStatusesResponseBody) *DescribeMonitoringAgentStatusesResponse
	GetBody() *DescribeMonitoringAgentStatusesResponseBody
}

type DescribeMonitoringAgentStatusesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitoringAgentStatusesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitoringAgentStatusesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentStatusesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentStatusesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitoringAgentStatusesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitoringAgentStatusesResponse) GetBody() *DescribeMonitoringAgentStatusesResponseBody {
	return s.Body
}

func (s *DescribeMonitoringAgentStatusesResponse) SetHeaders(v map[string]*string) *DescribeMonitoringAgentStatusesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponse) SetStatusCode(v int32) *DescribeMonitoringAgentStatusesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponse) SetBody(v *DescribeMonitoringAgentStatusesResponseBody) *DescribeMonitoringAgentStatusesResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
