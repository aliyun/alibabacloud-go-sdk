// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiProductsByAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiProductsByAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiProductsByAppResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiProductsByAppResponseBody) *DescribeApiProductsByAppResponse
	GetBody() *DescribeApiProductsByAppResponseBody
}

type DescribeApiProductsByAppResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiProductsByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiProductsByAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiProductsByAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiProductsByAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiProductsByAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiProductsByAppResponse) GetBody() *DescribeApiProductsByAppResponseBody {
	return s.Body
}

func (s *DescribeApiProductsByAppResponse) SetHeaders(v map[string]*string) *DescribeApiProductsByAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiProductsByAppResponse) SetStatusCode(v int32) *DescribeApiProductsByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiProductsByAppResponse) SetBody(v *DescribeApiProductsByAppResponseBody) *DescribeApiProductsByAppResponse {
	s.Body = v
	return s
}

func (s *DescribeApiProductsByAppResponse) Validate() error {
	return dara.Validate(s)
}
