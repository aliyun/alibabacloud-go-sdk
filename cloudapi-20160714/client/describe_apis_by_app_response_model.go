// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisByAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisByAppResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisByAppResponseBody) *DescribeApisByAppResponse
	GetBody() *DescribeApisByAppResponseBody
}

type DescribeApisByAppResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisByAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisByAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisByAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisByAppResponse) GetBody() *DescribeApisByAppResponseBody {
	return s.Body
}

func (s *DescribeApisByAppResponse) SetHeaders(v map[string]*string) *DescribeApisByAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisByAppResponse) SetStatusCode(v int32) *DescribeApisByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisByAppResponse) SetBody(v *DescribeApisByAppResponseBody) *DescribeApisByAppResponse {
	s.Body = v
	return s
}

func (s *DescribeApisByAppResponse) Validate() error {
	return dara.Validate(s)
}
