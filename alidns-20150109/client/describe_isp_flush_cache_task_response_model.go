// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIspFlushCacheTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIspFlushCacheTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIspFlushCacheTaskResponseBody) *DescribeIspFlushCacheTaskResponse
	GetBody() *DescribeIspFlushCacheTaskResponseBody
}

type DescribeIspFlushCacheTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIspFlushCacheTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIspFlushCacheTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIspFlushCacheTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIspFlushCacheTaskResponse) GetBody() *DescribeIspFlushCacheTaskResponseBody {
	return s.Body
}

func (s *DescribeIspFlushCacheTaskResponse) SetHeaders(v map[string]*string) *DescribeIspFlushCacheTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeIspFlushCacheTaskResponse) SetStatusCode(v int32) *DescribeIspFlushCacheTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponse) SetBody(v *DescribeIspFlushCacheTaskResponseBody) *DescribeIspFlushCacheTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeIspFlushCacheTaskResponse) Validate() error {
	return dara.Validate(s)
}
