// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageResponseBody) *DescribeImageResponse
	GetBody() *DescribeImageResponseBody
}

type DescribeImageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageResponse) GetBody() *DescribeImageResponseBody {
	return s.Body
}

func (s *DescribeImageResponse) SetHeaders(v map[string]*string) *DescribeImageResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageResponse) SetStatusCode(v int32) *DescribeImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageResponse) SetBody(v *DescribeImageResponseBody) *DescribeImageResponse {
	s.Body = v
	return s
}

func (s *DescribeImageResponse) Validate() error {
	return dara.Validate(s)
}
