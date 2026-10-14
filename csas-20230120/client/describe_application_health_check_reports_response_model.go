// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationHealthCheckReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationHealthCheckReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationHealthCheckReportsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationHealthCheckReportsResponseBody) *DescribeApplicationHealthCheckReportsResponse
	GetBody() *DescribeApplicationHealthCheckReportsResponseBody
}

type DescribeApplicationHealthCheckReportsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationHealthCheckReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationHealthCheckReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationHealthCheckReportsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationHealthCheckReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationHealthCheckReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationHealthCheckReportsResponse) GetBody() *DescribeApplicationHealthCheckReportsResponseBody {
	return s.Body
}

func (s *DescribeApplicationHealthCheckReportsResponse) SetHeaders(v map[string]*string) *DescribeApplicationHealthCheckReportsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponse) SetStatusCode(v int32) *DescribeApplicationHealthCheckReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponse) SetBody(v *DescribeApplicationHealthCheckReportsResponseBody) *DescribeApplicationHealthCheckReportsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
