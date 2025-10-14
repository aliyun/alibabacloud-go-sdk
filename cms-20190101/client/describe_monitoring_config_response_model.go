// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitoringConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitoringConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitoringConfigResponseBody) *DescribeMonitoringConfigResponse
	GetBody() *DescribeMonitoringConfigResponseBody
}

type DescribeMonitoringConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitoringConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitoringConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitoringConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitoringConfigResponse) GetBody() *DescribeMonitoringConfigResponseBody {
	return s.Body
}

func (s *DescribeMonitoringConfigResponse) SetHeaders(v map[string]*string) *DescribeMonitoringConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringConfigResponse) SetStatusCode(v int32) *DescribeMonitoringConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitoringConfigResponse) SetBody(v *DescribeMonitoringConfigResponseBody) *DescribeMonitoringConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitoringConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
