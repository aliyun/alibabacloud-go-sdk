// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpDdosThresholdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeIpDdosThresholdResponseBody
	GetRequestId() *string
	SetThreshold(v *DescribeIpDdosThresholdResponseBodyThreshold) *DescribeIpDdosThresholdResponseBody
	GetThreshold() *DescribeIpDdosThresholdResponseBodyThreshold
}

type DescribeIpDdosThresholdResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 025F1688-680B-551A-9A8E-1A0D796315CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the details of the threshold.
	Threshold *DescribeIpDdosThresholdResponseBodyThreshold `json:"Threshold,omitempty" xml:"Threshold,omitempty" type:"Struct"`
}

func (s DescribeIpDdosThresholdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpDdosThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpDdosThresholdResponseBody) GetThreshold() *DescribeIpDdosThresholdResponseBodyThreshold {
	return s.Threshold
}

func (s *DescribeIpDdosThresholdResponseBody) SetRequestId(v string) *DescribeIpDdosThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBody) SetThreshold(v *DescribeIpDdosThresholdResponseBodyThreshold) *DescribeIpDdosThresholdResponseBody {
	s.Threshold = v
	return s
}

func (s *DescribeIpDdosThresholdResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeIpDdosThresholdResponseBodyThreshold struct {
	// If the value of the **DdosType*	- parameter is **defense**, the Bps parameter indicates the current traffic scrubbing threshold. Unit: Mbit/s.
	//
	// If the value of the **DdosType*	- parameter is **blackhole**, the Bps parameter indicates the basic protection threshold. Unit: Mbit/s.
	//
	// example:
	//
	// 7500
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The type of the threshold. Valid values:
	//
	// 	- **defense**: traffic scrubbing threshold
	//
	// 	- **blackhole**: DDoS mitigation threshold
	//
	// example:
	//
	// defense
	DdosType *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	// The burstable protection threshold (the maximum DDoS mitigation threshold). Unit: Mbit/s.
	//
	// > This parameter is returned only when the value of the **DdosType*	- parameter is **blackhole**.
	//
	// example:
	//
	// 12310
	ElasticBps *int32 `json:"ElasticBps,omitempty" xml:"ElasticBps,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1i88rqjza51s****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Indicates whether the threshold is automatically adjusted. Valid values:
	//
	// 	- **true**: The scrubbing thresholds are automatically adjusted based on the traffic load on the asset.
	//
	// 	- **false**: The scrubbing thresholds are not automatically adjusted. You must manually specify the scrubbing thresholds.
	//
	// example:
	//
	// false
	IsAuto *bool `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	// The maximum traffic scrubbing threshold. Unit: Mbit/s.
	//
	// example:
	//
	// 7500
	MaxBps *int32 `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	// The maximum packet scrubbing threshold. Unit: pps.
	//
	// example:
	//
	// 5000000
	MaxPps *int32 `json:"MaxPps,omitempty" xml:"MaxPps,omitempty"`
	// The packet scrubbing threshold. Unit: packets per second (pps).
	//
	// > This parameter is returned only when the value of the **DdosType*	- parameter is **defense**.
	//
	// example:
	//
	// 5000000
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeIpDdosThresholdResponseBodyThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpDdosThresholdResponseBodyThreshold) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) GetBps() *int32 {
	return s.Bps
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) GetDdosType() *string {
	return s.DdosType
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) GetElasticBps() *int32 {
	return s.ElasticBps
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) GetIsAuto() *bool {
	return s.IsAuto
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) GetMaxBps() *int32 {
	return s.MaxBps
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) GetMaxPps() *int32 {
	return s.MaxPps
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) GetPps() *int32 {
	return s.Pps
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetBps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.Bps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetDdosType(v string) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.DdosType = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetElasticBps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.ElasticBps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetInstanceId(v string) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetInternetIp(v string) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.InternetIp = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetIsAuto(v bool) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.IsAuto = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetMaxBps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.MaxBps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetMaxPps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.MaxPps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetPps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.Pps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) Validate() error {
	return dara.Validate(s)
}
