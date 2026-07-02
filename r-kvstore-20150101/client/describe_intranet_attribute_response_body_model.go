// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntranetAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewal(v bool) *DescribeIntranetAttributeResponseBody
	GetAutoRenewal() *bool
	SetBandwidthExpireTime(v string) *DescribeIntranetAttributeResponseBody
	GetBandwidthExpireTime() *string
	SetBandwidthPrePaid(v string) *DescribeIntranetAttributeResponseBody
	GetBandwidthPrePaid() *string
	SetExpireTime(v string) *DescribeIntranetAttributeResponseBody
	GetExpireTime() *string
	SetHasPrePaidBandWidthOrderRunning(v bool) *DescribeIntranetAttributeResponseBody
	GetHasPrePaidBandWidthOrderRunning() *bool
	SetIntranetBandWidthBurst(v int32) *DescribeIntranetAttributeResponseBody
	GetIntranetBandWidthBurst() *int32
	SetIntranetBandwidth(v int32) *DescribeIntranetAttributeResponseBody
	GetIntranetBandwidth() *int32
	SetRequestId(v string) *DescribeIntranetAttributeResponseBody
	GetRequestId() *string
}

type DescribeIntranetAttributeResponseBody struct {
	// Indicates whether auto-renewal is enabled for the bandwidth package. Valid values:
	//
	// - **true**: Auto-renewal is enabled.
	//
	// - **false**: Auto-renewal is disabled.
	//
	// > This parameter is not returned if no additional bandwidth is purchased.
	//
	// example:
	//
	// true
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The expiration time of the bandwidth package. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format.
	//
	// > This parameter is not returned if no additional bandwidth is purchased.
	//
	// example:
	//
	// 2021-03-06T16:00:00Z
	BandwidthExpireTime *string `json:"BandwidthExpireTime,omitempty" xml:"BandwidthExpireTime,omitempty"`
	// The billing method of the bandwidth package. Valid values:
	//
	// - **0**: pay-as-you-go.
	//
	// - **1**: subscription.
	//
	// example:
	//
	// 0
	BandwidthPrePaid *string `json:"BandwidthPrePaid,omitempty" xml:"BandwidthPrePaid,omitempty"`
	// The expiration time of the temporary bandwidth. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format.
	//
	// > This parameter returns **0*	- if the instance has no temporary bandwidth or if the temporary bandwidth has expired.
	//
	// example:
	//
	// 0
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the instance has an unexpired bandwidth package. Valid values:
	//
	// - **true**: An unexpired bandwidth package exists.
	//
	// - **false**: No unexpired bandwidth package exists.
	//
	// > This parameter is not returned if no additional bandwidth is purchased.
	//
	// example:
	//
	// true
	HasPrePaidBandWidthOrderRunning *bool  `json:"HasPrePaidBandWidthOrderRunning,omitempty" xml:"HasPrePaidBandWidthOrderRunning,omitempty"`
	IntranetBandWidthBurst          *int32 `json:"IntranetBandWidthBurst,omitempty" xml:"IntranetBandWidthBurst,omitempty"`
	// The total intranet bandwidth across all shards in the instance, in MB/s.
	//
	// example:
	//
	// 102
	IntranetBandwidth *int32 `json:"IntranetBandwidth,omitempty" xml:"IntranetBandwidth,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25D42CC3-FBA1-4AEC-BCE2-B8DD3137****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIntranetAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntranetAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIntranetAttributeResponseBody) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DescribeIntranetAttributeResponseBody) GetBandwidthExpireTime() *string {
	return s.BandwidthExpireTime
}

func (s *DescribeIntranetAttributeResponseBody) GetBandwidthPrePaid() *string {
	return s.BandwidthPrePaid
}

func (s *DescribeIntranetAttributeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeIntranetAttributeResponseBody) GetHasPrePaidBandWidthOrderRunning() *bool {
	return s.HasPrePaidBandWidthOrderRunning
}

func (s *DescribeIntranetAttributeResponseBody) GetIntranetBandWidthBurst() *int32 {
	return s.IntranetBandWidthBurst
}

func (s *DescribeIntranetAttributeResponseBody) GetIntranetBandwidth() *int32 {
	return s.IntranetBandwidth
}

func (s *DescribeIntranetAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIntranetAttributeResponseBody) SetAutoRenewal(v bool) *DescribeIntranetAttributeResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeIntranetAttributeResponseBody) SetBandwidthExpireTime(v string) *DescribeIntranetAttributeResponseBody {
	s.BandwidthExpireTime = &v
	return s
}

func (s *DescribeIntranetAttributeResponseBody) SetBandwidthPrePaid(v string) *DescribeIntranetAttributeResponseBody {
	s.BandwidthPrePaid = &v
	return s
}

func (s *DescribeIntranetAttributeResponseBody) SetExpireTime(v string) *DescribeIntranetAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeIntranetAttributeResponseBody) SetHasPrePaidBandWidthOrderRunning(v bool) *DescribeIntranetAttributeResponseBody {
	s.HasPrePaidBandWidthOrderRunning = &v
	return s
}

func (s *DescribeIntranetAttributeResponseBody) SetIntranetBandWidthBurst(v int32) *DescribeIntranetAttributeResponseBody {
	s.IntranetBandWidthBurst = &v
	return s
}

func (s *DescribeIntranetAttributeResponseBody) SetIntranetBandwidth(v int32) *DescribeIntranetAttributeResponseBody {
	s.IntranetBandwidth = &v
	return s
}

func (s *DescribeIntranetAttributeResponseBody) SetRequestId(v string) *DescribeIntranetAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIntranetAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
