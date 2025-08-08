// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiMeteringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiMeteringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiMeteringResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiMeteringResponseBody) *DescribeApiMeteringResponse
	GetBody() *DescribeApiMeteringResponseBody
}

type DescribeApiMeteringResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiMeteringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiMeteringResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiMeteringResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiMeteringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiMeteringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiMeteringResponse) GetBody() *DescribeApiMeteringResponseBody {
	return s.Body
}

func (s *DescribeApiMeteringResponse) SetHeaders(v map[string]*string) *DescribeApiMeteringResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiMeteringResponse) SetStatusCode(v int32) *DescribeApiMeteringResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiMeteringResponse) SetBody(v *DescribeApiMeteringResponseBody) *DescribeApiMeteringResponse {
	s.Body = v
	return s
}

func (s *DescribeApiMeteringResponse) Validate() error {
	return dara.Validate(s)
}
