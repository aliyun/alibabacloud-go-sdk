// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL7RsPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeL7RsPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeL7RsPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeL7RsPolicyResponseBody) *DescribeL7RsPolicyResponse
	GetBody() *DescribeL7RsPolicyResponseBody
}

type DescribeL7RsPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeL7RsPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeL7RsPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7RsPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeL7RsPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeL7RsPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeL7RsPolicyResponse) GetBody() *DescribeL7RsPolicyResponseBody {
	return s.Body
}

func (s *DescribeL7RsPolicyResponse) SetHeaders(v map[string]*string) *DescribeL7RsPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeL7RsPolicyResponse) SetStatusCode(v int32) *DescribeL7RsPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeL7RsPolicyResponse) SetBody(v *DescribeL7RsPolicyResponseBody) *DescribeL7RsPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeL7RsPolicyResponse) Validate() error {
	return dara.Validate(s)
}
