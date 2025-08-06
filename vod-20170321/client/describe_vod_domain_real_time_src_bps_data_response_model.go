// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeSrcBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeSrcBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainRealTimeSrcBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainRealTimeSrcBpsDataResponseBody) *DescribeVodDomainRealTimeSrcBpsDataResponse
	GetBody() *DescribeVodDomainRealTimeSrcBpsDataResponseBody
}

type DescribeVodDomainRealTimeSrcBpsDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainRealTimeSrcBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainRealTimeSrcBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeSrcBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponse) GetBody() *DescribeVodDomainRealTimeSrcBpsDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeSrcBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponse) SetStatusCode(v int32) *DescribeVodDomainRealTimeSrcBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponse) SetBody(v *DescribeVodDomainRealTimeSrcBpsDataResponseBody) *DescribeVodDomainRealTimeSrcBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataResponse) Validate() error {
	return dara.Validate(s)
}
