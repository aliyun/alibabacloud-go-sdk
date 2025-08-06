// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainSrcBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainSrcBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainSrcBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainSrcBpsDataResponseBody) *DescribeVodDomainSrcBpsDataResponse
	GetBody() *DescribeVodDomainSrcBpsDataResponseBody
}

type DescribeVodDomainSrcBpsDataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainSrcBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainSrcBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainSrcBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainSrcBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainSrcBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainSrcBpsDataResponse) GetBody() *DescribeVodDomainSrcBpsDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainSrcBpsDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainSrcBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponse) SetStatusCode(v int32) *DescribeVodDomainSrcBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponse) SetBody(v *DescribeVodDomainSrcBpsDataResponseBody) *DescribeVodDomainSrcBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainSrcBpsDataResponse) Validate() error {
	return dara.Validate(s)
}
