// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageResultExtResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageResultExtResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageResultExtResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageResultExtResponseBody) *DescribeImageResultExtResponse
	GetBody() *DescribeImageResultExtResponseBody
}

type DescribeImageResultExtResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageResultExtResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageResultExtResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageResultExtResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageResultExtResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageResultExtResponse) GetBody() *DescribeImageResultExtResponseBody {
	return s.Body
}

func (s *DescribeImageResultExtResponse) SetHeaders(v map[string]*string) *DescribeImageResultExtResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageResultExtResponse) SetStatusCode(v int32) *DescribeImageResultExtResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageResultExtResponse) SetBody(v *DescribeImageResultExtResponseBody) *DescribeImageResultExtResponse {
	s.Body = v
	return s
}

func (s *DescribeImageResultExtResponse) Validate() error {
	return dara.Validate(s)
}
