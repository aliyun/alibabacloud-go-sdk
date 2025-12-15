// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDistributeCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalDistributeCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalDistributeCacheResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalDistributeCacheResponseBody) *DescribeGlobalDistributeCacheResponse
	GetBody() *DescribeGlobalDistributeCacheResponseBody
}

type DescribeGlobalDistributeCacheResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalDistributeCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalDistributeCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDistributeCacheResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDistributeCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalDistributeCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalDistributeCacheResponse) GetBody() *DescribeGlobalDistributeCacheResponseBody {
	return s.Body
}

func (s *DescribeGlobalDistributeCacheResponse) SetHeaders(v map[string]*string) *DescribeGlobalDistributeCacheResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalDistributeCacheResponse) SetStatusCode(v int32) *DescribeGlobalDistributeCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponse) SetBody(v *DescribeGlobalDistributeCacheResponseBody) *DescribeGlobalDistributeCacheResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalDistributeCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
