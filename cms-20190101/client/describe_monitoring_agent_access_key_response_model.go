// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentAccessKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitoringAgentAccessKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitoringAgentAccessKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitoringAgentAccessKeyResponseBody) *DescribeMonitoringAgentAccessKeyResponse
	GetBody() *DescribeMonitoringAgentAccessKeyResponseBody
}

type DescribeMonitoringAgentAccessKeyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitoringAgentAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitoringAgentAccessKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentAccessKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitoringAgentAccessKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitoringAgentAccessKeyResponse) GetBody() *DescribeMonitoringAgentAccessKeyResponseBody {
	return s.Body
}

func (s *DescribeMonitoringAgentAccessKeyResponse) SetHeaders(v map[string]*string) *DescribeMonitoringAgentAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponse) SetStatusCode(v int32) *DescribeMonitoringAgentAccessKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponse) SetBody(v *DescribeMonitoringAgentAccessKeyResponseBody) *DescribeMonitoringAgentAccessKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
