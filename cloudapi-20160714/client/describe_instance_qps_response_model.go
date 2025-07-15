// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceQpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceQpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceQpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceQpsResponseBody) *DescribeInstanceQpsResponse
	GetBody() *DescribeInstanceQpsResponseBody
}

type DescribeInstanceQpsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceQpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceQpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceQpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceQpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceQpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceQpsResponse) GetBody() *DescribeInstanceQpsResponseBody {
	return s.Body
}

func (s *DescribeInstanceQpsResponse) SetHeaders(v map[string]*string) *DescribeInstanceQpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceQpsResponse) SetStatusCode(v int32) *DescribeInstanceQpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceQpsResponse) SetBody(v *DescribeInstanceQpsResponseBody) *DescribeInstanceQpsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceQpsResponse) Validate() error {
	return dara.Validate(s)
}
