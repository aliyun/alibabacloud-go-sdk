// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckStatusListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHealthCheckStatusListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHealthCheckStatusListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHealthCheckStatusListResponseBody) *DescribeHealthCheckStatusListResponse
	GetBody() *DescribeHealthCheckStatusListResponseBody
}

type DescribeHealthCheckStatusListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHealthCheckStatusListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHealthCheckStatusListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHealthCheckStatusListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHealthCheckStatusListResponse) GetBody() *DescribeHealthCheckStatusListResponseBody {
	return s.Body
}

func (s *DescribeHealthCheckStatusListResponse) SetHeaders(v map[string]*string) *DescribeHealthCheckStatusListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthCheckStatusListResponse) SetStatusCode(v int32) *DescribeHealthCheckStatusListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponse) SetBody(v *DescribeHealthCheckStatusListResponseBody) *DescribeHealthCheckStatusListResponse {
	s.Body = v
	return s
}

func (s *DescribeHealthCheckStatusListResponse) Validate() error {
	return dara.Validate(s)
}
