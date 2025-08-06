// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainSrcTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainSrcTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainSrcTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainSrcTrafficDataResponseBody) *DescribeVodDomainSrcTrafficDataResponse
	GetBody() *DescribeVodDomainSrcTrafficDataResponseBody
}

type DescribeVodDomainSrcTrafficDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainSrcTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainSrcTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainSrcTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainSrcTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainSrcTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainSrcTrafficDataResponse) GetBody() *DescribeVodDomainSrcTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainSrcTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainSrcTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponse) SetStatusCode(v int32) *DescribeVodDomainSrcTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponse) SetBody(v *DescribeVodDomainSrcTrafficDataResponseBody) *DescribeVodDomainSrcTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainSrcTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
