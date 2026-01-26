// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvServiceMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnvServiceMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnvServiceMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnvServiceMonitorResponseBody) *DescribeEnvServiceMonitorResponse
	GetBody() *DescribeEnvServiceMonitorResponseBody
}

type DescribeEnvServiceMonitorResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnvServiceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnvServiceMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvServiceMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnvServiceMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnvServiceMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnvServiceMonitorResponse) GetBody() *DescribeEnvServiceMonitorResponseBody {
	return s.Body
}

func (s *DescribeEnvServiceMonitorResponse) SetHeaders(v map[string]*string) *DescribeEnvServiceMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnvServiceMonitorResponse) SetStatusCode(v int32) *DescribeEnvServiceMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnvServiceMonitorResponse) SetBody(v *DescribeEnvServiceMonitorResponseBody) *DescribeEnvServiceMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeEnvServiceMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
