// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageFixTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageFixTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageFixTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageFixTaskResponseBody) *DescribeImageFixTaskResponse
	GetBody() *DescribeImageFixTaskResponseBody
}

type DescribeImageFixTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageFixTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageFixTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFixTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageFixTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageFixTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageFixTaskResponse) GetBody() *DescribeImageFixTaskResponseBody {
	return s.Body
}

func (s *DescribeImageFixTaskResponse) SetHeaders(v map[string]*string) *DescribeImageFixTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageFixTaskResponse) SetStatusCode(v int32) *DescribeImageFixTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageFixTaskResponse) SetBody(v *DescribeImageFixTaskResponseBody) *DescribeImageFixTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeImageFixTaskResponse) Validate() error {
	return dara.Validate(s)
}
