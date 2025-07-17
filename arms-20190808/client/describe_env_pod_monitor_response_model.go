// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvPodMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnvPodMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnvPodMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnvPodMonitorResponseBody) *DescribeEnvPodMonitorResponse
	GetBody() *DescribeEnvPodMonitorResponseBody
}

type DescribeEnvPodMonitorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnvPodMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnvPodMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvPodMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnvPodMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnvPodMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnvPodMonitorResponse) GetBody() *DescribeEnvPodMonitorResponseBody {
	return s.Body
}

func (s *DescribeEnvPodMonitorResponse) SetHeaders(v map[string]*string) *DescribeEnvPodMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnvPodMonitorResponse) SetStatusCode(v int32) *DescribeEnvPodMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnvPodMonitorResponse) SetBody(v *DescribeEnvPodMonitorResponseBody) *DescribeEnvPodMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeEnvPodMonitorResponse) Validate() error {
	return dara.Validate(s)
}
