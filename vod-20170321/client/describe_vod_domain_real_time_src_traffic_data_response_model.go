// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeSrcTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeSrcTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainRealTimeSrcTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) *DescribeVodDomainRealTimeSrcTrafficDataResponse
	GetBody() *DescribeVodDomainRealTimeSrcTrafficDataResponseBody
}

type DescribeVodDomainRealTimeSrcTrafficDataResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainRealTimeSrcTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainRealTimeSrcTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeSrcTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponse) GetBody() *DescribeVodDomainRealTimeSrcTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeSrcTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponse) SetStatusCode(v int32) *DescribeVodDomainRealTimeSrcTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponse) SetBody(v *DescribeVodDomainRealTimeSrcTrafficDataResponseBody) *DescribeVodDomainRealTimeSrcTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
