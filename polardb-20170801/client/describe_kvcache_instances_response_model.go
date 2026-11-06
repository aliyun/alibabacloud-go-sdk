// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKVCacheInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKVCacheInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKVCacheInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKVCacheInstancesResponseBody) *DescribeKVCacheInstancesResponse
	GetBody() *DescribeKVCacheInstancesResponseBody
}

type DescribeKVCacheInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKVCacheInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKVCacheInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKVCacheInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeKVCacheInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKVCacheInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKVCacheInstancesResponse) GetBody() *DescribeKVCacheInstancesResponseBody {
	return s.Body
}

func (s *DescribeKVCacheInstancesResponse) SetHeaders(v map[string]*string) *DescribeKVCacheInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeKVCacheInstancesResponse) SetStatusCode(v int32) *DescribeKVCacheInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKVCacheInstancesResponse) SetBody(v *DescribeKVCacheInstancesResponseBody) *DescribeKVCacheInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeKVCacheInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
