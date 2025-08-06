// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRegionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainRegionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainRegionDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainRegionDataResponseBody) *DescribeVodDomainRegionDataResponse
	GetBody() *DescribeVodDomainRegionDataResponseBody
}

type DescribeVodDomainRegionDataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainRegionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainRegionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRegionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRegionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainRegionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainRegionDataResponse) GetBody() *DescribeVodDomainRegionDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainRegionDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRegionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRegionDataResponse) SetStatusCode(v int32) *DescribeVodDomainRegionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponse) SetBody(v *DescribeVodDomainRegionDataResponseBody) *DescribeVodDomainRegionDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainRegionDataResponse) Validate() error {
	return dara.Validate(s)
}
