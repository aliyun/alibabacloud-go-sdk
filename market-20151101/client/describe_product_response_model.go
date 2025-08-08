// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProductResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProductResponseBody) *DescribeProductResponse
	GetBody() *DescribeProductResponseBody
}

type DescribeProductResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProductResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProductResponse) GetBody() *DescribeProductResponseBody {
	return s.Body
}

func (s *DescribeProductResponse) SetHeaders(v map[string]*string) *DescribeProductResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductResponse) SetStatusCode(v int32) *DescribeProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductResponse) SetBody(v *DescribeProductResponseBody) *DescribeProductResponse {
	s.Body = v
	return s
}

func (s *DescribeProductResponse) Validate() error {
	return dara.Validate(s)
}
