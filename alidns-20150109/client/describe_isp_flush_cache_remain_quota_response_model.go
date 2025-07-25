// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheRemainQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIspFlushCacheRemainQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIspFlushCacheRemainQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIspFlushCacheRemainQuotaResponseBody) *DescribeIspFlushCacheRemainQuotaResponse
	GetBody() *DescribeIspFlushCacheRemainQuotaResponseBody
}

type DescribeIspFlushCacheRemainQuotaResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIspFlushCacheRemainQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIspFlushCacheRemainQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheRemainQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheRemainQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIspFlushCacheRemainQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIspFlushCacheRemainQuotaResponse) GetBody() *DescribeIspFlushCacheRemainQuotaResponseBody {
	return s.Body
}

func (s *DescribeIspFlushCacheRemainQuotaResponse) SetHeaders(v map[string]*string) *DescribeIspFlushCacheRemainQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeIspFlushCacheRemainQuotaResponse) SetStatusCode(v int32) *DescribeIspFlushCacheRemainQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIspFlushCacheRemainQuotaResponse) SetBody(v *DescribeIspFlushCacheRemainQuotaResponseBody) *DescribeIspFlushCacheRemainQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeIspFlushCacheRemainQuotaResponse) Validate() error {
	return dara.Validate(s)
}
