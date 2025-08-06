// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainUsageDataResponseBody) *DescribeVodDomainUsageDataResponse
	GetBody() *DescribeVodDomainUsageDataResponseBody
}

type DescribeVodDomainUsageDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainUsageDataResponse) GetBody() *DescribeVodDomainUsageDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainUsageDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainUsageDataResponse) SetStatusCode(v int32) *DescribeVodDomainUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponse) SetBody(v *DescribeVodDomainUsageDataResponseBody) *DescribeVodDomainUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainUsageDataResponse) Validate() error {
	return dara.Validate(s)
}
