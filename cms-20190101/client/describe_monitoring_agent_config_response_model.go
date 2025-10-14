// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitoringAgentConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitoringAgentConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitoringAgentConfigResponseBody) *DescribeMonitoringAgentConfigResponse
	GetBody() *DescribeMonitoringAgentConfigResponseBody
}

type DescribeMonitoringAgentConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitoringAgentConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitoringAgentConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitoringAgentConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitoringAgentConfigResponse) GetBody() *DescribeMonitoringAgentConfigResponseBody {
	return s.Body
}

func (s *DescribeMonitoringAgentConfigResponse) SetHeaders(v map[string]*string) *DescribeMonitoringAgentConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringAgentConfigResponse) SetStatusCode(v int32) *DescribeMonitoringAgentConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponse) SetBody(v *DescribeMonitoringAgentConfigResponseBody) *DescribeMonitoringAgentConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitoringAgentConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
