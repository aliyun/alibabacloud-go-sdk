// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIspFlushCacheTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIspFlushCacheTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIspFlushCacheTasksResponseBody) *DescribeIspFlushCacheTasksResponse
	GetBody() *DescribeIspFlushCacheTasksResponseBody
}

type DescribeIspFlushCacheTasksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIspFlushCacheTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIspFlushCacheTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIspFlushCacheTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIspFlushCacheTasksResponse) GetBody() *DescribeIspFlushCacheTasksResponseBody {
	return s.Body
}

func (s *DescribeIspFlushCacheTasksResponse) SetHeaders(v map[string]*string) *DescribeIspFlushCacheTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeIspFlushCacheTasksResponse) SetStatusCode(v int32) *DescribeIspFlushCacheTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIspFlushCacheTasksResponse) SetBody(v *DescribeIspFlushCacheTasksResponseBody) *DescribeIspFlushCacheTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeIspFlushCacheTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
