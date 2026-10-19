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
	// Indicates whether auto-renewal is enabled for the purchased bandwidth. Valid values:
	//
	// 	- **true**: Enabled.
	//
	// 	- **false**: Disabled.
	//
	// > This parameter is not returned if no bandwidth has been purchased.
	//
	// example:
	//
	// true
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The expiration time of the purchased bandwidth, in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format.
	//
	// > This parameter is not returned if no additional bandwidth has been purchased for the instance.
	//
	// example:
	//
	// 2021-03-06T16:00:00Z
	BandwidthExpireTime *string `json:"BandwidthExpireTime,omitempty" xml:"BandwidthExpireTime,omitempty"`
	// The billing method of the bandwidth plan. Valid values:
	//
	// - **0**: pay-as-you-go.
	//
	// - **1**: subscription.
	//
	// example:
	//
	// 0
	BandwidthPrePaid *string `json:"BandwidthPrePaid,omitempty" xml:"BandwidthPrePaid,omitempty"`
	// The expiration time of the temporary bandwidth, in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format.
	//
	// > If no temporary bandwidth adjustment has been performed on the instance or the temporary bandwidth has expired, this parameter returns **0**.
	//
	// example:
	//
	// 0
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the instance has an unexpired bandwidth plan. Valid values:
	//
	// 	- **true**: An unexpired bandwidth plan exists.
	//
	// 	- **false**: No unexpired bandwidth plan exists.
	//
	// > This parameter is not returned if no additional bandwidth has been purchased for the instance.
	//
	// example:
	//
	// true
	HasPrePaidBandWidthOrderRunning *bool  `json:"HasPrePaidBandWidthOrderRunning,omitempty" xml:"HasPrePaidBandWidthOrderRunning,omitempty"`
	IntranetBandWidthBurst          *int32 `json:"IntranetBandWidthBurst,omitempty" xml:"IntranetBandWidthBurst,omitempty"`
	// The current internal network bandwidth of the instance. This parameter indicates the total bandwidth of all shards in the instance. Unit: MB/s.
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
