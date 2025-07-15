// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTrafficResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceTrafficResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceTrafficResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceTrafficResponseBody) *DescribeInstanceTrafficResponse
	GetBody() *DescribeInstanceTrafficResponseBody
}

type DescribeInstanceTrafficResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceTrafficResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTrafficResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceTrafficResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceTrafficResponse) GetBody() *DescribeInstanceTrafficResponseBody {
	return s.Body
}

func (s *DescribeInstanceTrafficResponse) SetHeaders(v map[string]*string) *DescribeInstanceTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTrafficResponse) SetStatusCode(v int32) *DescribeInstanceTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceTrafficResponse) SetBody(v *DescribeInstanceTrafficResponseBody) *DescribeInstanceTrafficResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceTrafficResponse) Validate() error {
	return dara.Validate(s)
}
