// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIspFlushCacheInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIspFlushCacheInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIspFlushCacheInstancesResponseBody) *DescribeIspFlushCacheInstancesResponse
	GetBody() *DescribeIspFlushCacheInstancesResponseBody
}

type DescribeIspFlushCacheInstancesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIspFlushCacheInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIspFlushCacheInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIspFlushCacheInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIspFlushCacheInstancesResponse) GetBody() *DescribeIspFlushCacheInstancesResponseBody {
	return s.Body
}

func (s *DescribeIspFlushCacheInstancesResponse) SetHeaders(v map[string]*string) *DescribeIspFlushCacheInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponse) SetStatusCode(v int32) *DescribeIspFlushCacheInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponse) SetBody(v *DescribeIspFlushCacheInstancesResponseBody) *DescribeIspFlushCacheInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponse) Validate() error {
	return dara.Validate(s)
}
