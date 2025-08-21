// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHealthCheckStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHealthCheckStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHealthCheckStatusResponseBody) *DescribeHealthCheckStatusResponse
	GetBody() *DescribeHealthCheckStatusResponseBody
}

type DescribeHealthCheckStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHealthCheckStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHealthCheckStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHealthCheckStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHealthCheckStatusResponse) GetBody() *DescribeHealthCheckStatusResponseBody {
	return s.Body
}

func (s *DescribeHealthCheckStatusResponse) SetHeaders(v map[string]*string) *DescribeHealthCheckStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthCheckStatusResponse) SetStatusCode(v int32) *DescribeHealthCheckStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHealthCheckStatusResponse) SetBody(v *DescribeHealthCheckStatusResponseBody) *DescribeHealthCheckStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeHealthCheckStatusResponse) Validate() error {
	return dara.Validate(s)
}
