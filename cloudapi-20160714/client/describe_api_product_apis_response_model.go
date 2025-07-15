// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiProductApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiProductApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiProductApisResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiProductApisResponseBody) *DescribeApiProductApisResponse
	GetBody() *DescribeApiProductApisResponseBody
}

type DescribeApiProductApisResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiProductApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiProductApisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiProductApisResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiProductApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiProductApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiProductApisResponse) GetBody() *DescribeApiProductApisResponseBody {
	return s.Body
}

func (s *DescribeApiProductApisResponse) SetHeaders(v map[string]*string) *DescribeApiProductApisResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiProductApisResponse) SetStatusCode(v int32) *DescribeApiProductApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiProductApisResponse) SetBody(v *DescribeApiProductApisResponseBody) *DescribeApiProductApisResponse {
	s.Body = v
	return s
}

func (s *DescribeApiProductApisResponse) Validate() error {
	return dara.Validate(s)
}
