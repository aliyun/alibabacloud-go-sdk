// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProductsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProductsResponseBody) *DescribeProductsResponse
	GetBody() *DescribeProductsResponseBody
}

type DescribeProductsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProductsResponse) GetBody() *DescribeProductsResponseBody {
	return s.Body
}

func (s *DescribeProductsResponse) SetHeaders(v map[string]*string) *DescribeProductsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductsResponse) SetStatusCode(v int32) *DescribeProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductsResponse) SetBody(v *DescribeProductsResponseBody) *DescribeProductsResponse {
	s.Body = v
	return s
}

func (s *DescribeProductsResponse) Validate() error {
	return dara.Validate(s)
}
