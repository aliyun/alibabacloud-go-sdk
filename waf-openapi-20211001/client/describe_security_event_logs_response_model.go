// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityEventLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityEventLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityEventLogsResponseBody) *DescribeSecurityEventLogsResponse
	GetBody() *DescribeSecurityEventLogsResponseBody
}

type DescribeSecurityEventLogsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityEventLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityEventLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityEventLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityEventLogsResponse) GetBody() *DescribeSecurityEventLogsResponseBody {
	return s.Body
}

func (s *DescribeSecurityEventLogsResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventLogsResponse) SetStatusCode(v int32) *DescribeSecurityEventLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityEventLogsResponse) SetBody(v *DescribeSecurityEventLogsResponseBody) *DescribeSecurityEventLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityEventLogsResponse) Validate() error {
	return dara.Validate(s)
}
