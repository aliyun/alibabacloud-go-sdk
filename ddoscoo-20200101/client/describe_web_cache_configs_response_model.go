// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCacheConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebCacheConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebCacheConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebCacheConfigsResponseBody) *DescribeWebCacheConfigsResponse
	GetBody() *DescribeWebCacheConfigsResponseBody
}

type DescribeWebCacheConfigsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebCacheConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebCacheConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCacheConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebCacheConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebCacheConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebCacheConfigsResponse) GetBody() *DescribeWebCacheConfigsResponseBody {
	return s.Body
}

func (s *DescribeWebCacheConfigsResponse) SetHeaders(v map[string]*string) *DescribeWebCacheConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebCacheConfigsResponse) SetStatusCode(v int32) *DescribeWebCacheConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebCacheConfigsResponse) SetBody(v *DescribeWebCacheConfigsResponseBody) *DescribeWebCacheConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeWebCacheConfigsResponse) Validate() error {
	return dara.Validate(s)
}
