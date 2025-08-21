// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse
	GetBody() *DescribeInstanceResponseBody
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceResponse) GetBody() *DescribeInstanceResponseBody {
	return s.Body
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceResponse) Validate() error {
	return dara.Validate(s)
}
