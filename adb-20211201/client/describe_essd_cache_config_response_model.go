// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEssdCacheConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEssdCacheConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEssdCacheConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEssdCacheConfigResponseBody) *DescribeEssdCacheConfigResponse
	GetBody() *DescribeEssdCacheConfigResponseBody
}

type DescribeEssdCacheConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEssdCacheConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEssdCacheConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEssdCacheConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeEssdCacheConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEssdCacheConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEssdCacheConfigResponse) GetBody() *DescribeEssdCacheConfigResponseBody {
	return s.Body
}

func (s *DescribeEssdCacheConfigResponse) SetHeaders(v map[string]*string) *DescribeEssdCacheConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeEssdCacheConfigResponse) SetStatusCode(v int32) *DescribeEssdCacheConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEssdCacheConfigResponse) SetBody(v *DescribeEssdCacheConfigResponseBody) *DescribeEssdCacheConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeEssdCacheConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
