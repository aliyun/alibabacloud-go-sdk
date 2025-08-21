// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStreamResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStreamResponseBody) *DescribeStreamResponse
	GetBody() *DescribeStreamResponseBody
}

type DescribeStreamResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStreamResponse) GetBody() *DescribeStreamResponseBody {
	return s.Body
}

func (s *DescribeStreamResponse) SetHeaders(v map[string]*string) *DescribeStreamResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamResponse) SetStatusCode(v int32) *DescribeStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStreamResponse) SetBody(v *DescribeStreamResponseBody) *DescribeStreamResponse {
	s.Body = v
	return s
}

func (s *DescribeStreamResponse) Validate() error {
	return dara.Validate(s)
}
