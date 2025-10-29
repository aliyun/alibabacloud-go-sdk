// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeXpackMonitorConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeXpackMonitorConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeXpackMonitorConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeXpackMonitorConfigResponseBody) *DescribeXpackMonitorConfigResponse
	GetBody() *DescribeXpackMonitorConfigResponseBody
}

type DescribeXpackMonitorConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeXpackMonitorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeXpackMonitorConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeXpackMonitorConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeXpackMonitorConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeXpackMonitorConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeXpackMonitorConfigResponse) GetBody() *DescribeXpackMonitorConfigResponseBody {
	return s.Body
}

func (s *DescribeXpackMonitorConfigResponse) SetHeaders(v map[string]*string) *DescribeXpackMonitorConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeXpackMonitorConfigResponse) SetStatusCode(v int32) *DescribeXpackMonitorConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeXpackMonitorConfigResponse) SetBody(v *DescribeXpackMonitorConfigResponseBody) *DescribeXpackMonitorConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeXpackMonitorConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
