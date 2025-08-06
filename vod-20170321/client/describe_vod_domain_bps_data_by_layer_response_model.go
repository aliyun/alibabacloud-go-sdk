// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainBpsDataByLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainBpsDataByLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainBpsDataByLayerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainBpsDataByLayerResponseBody) *DescribeVodDomainBpsDataByLayerResponse
	GetBody() *DescribeVodDomainBpsDataByLayerResponseBody
}

type DescribeVodDomainBpsDataByLayerResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainBpsDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainBpsDataByLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainBpsDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataByLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainBpsDataByLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainBpsDataByLayerResponse) GetBody() *DescribeVodDomainBpsDataByLayerResponseBody {
	return s.Body
}

func (s *DescribeVodDomainBpsDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeVodDomainBpsDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerResponse) SetStatusCode(v int32) *DescribeVodDomainBpsDataByLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerResponse) SetBody(v *DescribeVodDomainBpsDataByLayerResponseBody) *DescribeVodDomainBpsDataByLayerResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerResponse) Validate() error {
	return dara.Validate(s)
}
