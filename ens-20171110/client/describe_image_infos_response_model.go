// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageInfosResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageInfosResponseBody) *DescribeImageInfosResponse
	GetBody() *DescribeImageInfosResponseBody
}

type DescribeImageInfosResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageInfosResponse) GetBody() *DescribeImageInfosResponseBody {
	return s.Body
}

func (s *DescribeImageInfosResponse) SetHeaders(v map[string]*string) *DescribeImageInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageInfosResponse) SetStatusCode(v int32) *DescribeImageInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageInfosResponse) SetBody(v *DescribeImageInfosResponseBody) *DescribeImageInfosResponse {
	s.Body = v
	return s
}

func (s *DescribeImageInfosResponse) Validate() error {
	return dara.Validate(s)
}
