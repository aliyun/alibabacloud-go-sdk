// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenGeographicSpanRemainingBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemainingBandwidth(v int64) *DescribeCenGeographicSpanRemainingBandwidthResponseBody
	GetRemainingBandwidth() *int64
	SetRequestId(v string) *DescribeCenGeographicSpanRemainingBandwidthResponseBody
	GetRequestId() *string
}

type DescribeCenGeographicSpanRemainingBandwidthResponseBody struct {
	// The remaining bandwidth of the bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	RemainingBandwidth *int64 `json:"RemainingBandwidth,omitempty" xml:"RemainingBandwidth,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E4B345CD-2CBA-4881-AF6D-E5D9BAE1CA7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCenGeographicSpanRemainingBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenGeographicSpanRemainingBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponseBody) GetRemainingBandwidth() *int64 {
	return s.RemainingBandwidth
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponseBody) SetRemainingBandwidth(v int64) *DescribeCenGeographicSpanRemainingBandwidthResponseBody {
	s.RemainingBandwidth = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponseBody) SetRequestId(v string) *DescribeCenGeographicSpanRemainingBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
