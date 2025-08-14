// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenInterRegionBandwidthLimitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenInterRegionBandwidthLimitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenInterRegionBandwidthLimitsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenInterRegionBandwidthLimitsResponseBody) *DescribeCenInterRegionBandwidthLimitsResponse
	GetBody() *DescribeCenInterRegionBandwidthLimitsResponseBody
}

type DescribeCenInterRegionBandwidthLimitsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenInterRegionBandwidthLimitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenInterRegionBandwidthLimitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenInterRegionBandwidthLimitsResponse) GetBody() *DescribeCenInterRegionBandwidthLimitsResponseBody {
	return s.Body
}

func (s *DescribeCenInterRegionBandwidthLimitsResponse) SetHeaders(v map[string]*string) *DescribeCenInterRegionBandwidthLimitsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponse) SetStatusCode(v int32) *DescribeCenInterRegionBandwidthLimitsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponse) SetBody(v *DescribeCenInterRegionBandwidthLimitsResponseBody) *DescribeCenInterRegionBandwidthLimitsResponse {
	s.Body = v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponse) Validate() error {
	return dara.Validate(s)
}
