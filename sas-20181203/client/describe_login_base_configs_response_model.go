// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoginBaseConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoginBaseConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoginBaseConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoginBaseConfigsResponseBody) *DescribeLoginBaseConfigsResponse
	GetBody() *DescribeLoginBaseConfigsResponseBody
}

type DescribeLoginBaseConfigsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoginBaseConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoginBaseConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginBaseConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoginBaseConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoginBaseConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoginBaseConfigsResponse) GetBody() *DescribeLoginBaseConfigsResponseBody {
	return s.Body
}

func (s *DescribeLoginBaseConfigsResponse) SetHeaders(v map[string]*string) *DescribeLoginBaseConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoginBaseConfigsResponse) SetStatusCode(v int32) *DescribeLoginBaseConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoginBaseConfigsResponse) SetBody(v *DescribeLoginBaseConfigsResponseBody) *DescribeLoginBaseConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeLoginBaseConfigsResponse) Validate() error {
	return dara.Validate(s)
}
