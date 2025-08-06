// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainTrafficDataResponseBody) *DescribeVodDomainTrafficDataResponse
	GetBody() *DescribeVodDomainTrafficDataResponseBody
}

type DescribeVodDomainTrafficDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainTrafficDataResponse) GetBody() *DescribeVodDomainTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainTrafficDataResponse) SetStatusCode(v int32) *DescribeVodDomainTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponse) SetBody(v *DescribeVodDomainTrafficDataResponseBody) *DescribeVodDomainTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
