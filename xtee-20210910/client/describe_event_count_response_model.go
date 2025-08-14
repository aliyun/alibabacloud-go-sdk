// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventCountResponseBody) *DescribeEventCountResponse
	GetBody() *DescribeEventCountResponseBody
}

type DescribeEventCountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventCountResponse) GetBody() *DescribeEventCountResponseBody {
	return s.Body
}

func (s *DescribeEventCountResponse) SetHeaders(v map[string]*string) *DescribeEventCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventCountResponse) SetStatusCode(v int32) *DescribeEventCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventCountResponse) SetBody(v *DescribeEventCountResponseBody) *DescribeEventCountResponse {
	s.Body = v
	return s
}

func (s *DescribeEventCountResponse) Validate() error {
	return dara.Validate(s)
}
