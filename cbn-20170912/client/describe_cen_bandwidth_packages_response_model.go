// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenBandwidthPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenBandwidthPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenBandwidthPackagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenBandwidthPackagesResponseBody) *DescribeCenBandwidthPackagesResponse
	GetBody() *DescribeCenBandwidthPackagesResponseBody
}

type DescribeCenBandwidthPackagesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenBandwidthPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenBandwidthPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenBandwidthPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenBandwidthPackagesResponse) GetBody() *DescribeCenBandwidthPackagesResponseBody {
	return s.Body
}

func (s *DescribeCenBandwidthPackagesResponse) SetHeaders(v map[string]*string) *DescribeCenBandwidthPackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponse) SetStatusCode(v int32) *DescribeCenBandwidthPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponse) SetBody(v *DescribeCenBandwidthPackagesResponseBody) *DescribeCenBandwidthPackagesResponse {
	s.Body = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponse) Validate() error {
	return dara.Validate(s)
}
