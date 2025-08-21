// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosThresholdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDdosThresholdResponseBody
	GetRequestId() *string
	SetThresholds(v *DescribeDdosThresholdResponseBodyThresholds) *DescribeDdosThresholdResponseBody
	GetThresholds() *DescribeDdosThresholdResponseBodyThresholds
}

type DescribeDdosThresholdResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// E9B3C090-55AD-59C6-979E-FCFD81E7D9E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the details of the threshold.
	Thresholds *DescribeDdosThresholdResponseBodyThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
}

func (s DescribeDdosThresholdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDdosThresholdResponseBody) GetThresholds() *DescribeDdosThresholdResponseBodyThresholds {
	return s.Thresholds
}

func (s *DescribeDdosThresholdResponseBody) SetRequestId(v string) *DescribeDdosThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosThresholdResponseBody) SetThresholds(v *DescribeDdosThresholdResponseBodyThresholds) *DescribeDdosThresholdResponseBody {
	s.Thresholds = v
	return s
}

func (s *DescribeDdosThresholdResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDdosThresholdResponseBodyThresholds struct {
	Threshold []*DescribeDdosThresholdResponseBodyThresholdsThreshold `json:"Threshold,omitempty" xml:"Threshold,omitempty" type:"Repeated"`
}

func (s DescribeDdosThresholdResponseBodyThresholds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosThresholdResponseBodyThresholds) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponseBodyThresholds) GetThreshold() []*DescribeDdosThresholdResponseBodyThresholdsThreshold {
	return s.Threshold
}

func (s *DescribeDdosThresholdResponseBodyThresholds) SetThreshold(v []*DescribeDdosThresholdResponseBodyThresholdsThreshold) *DescribeDdosThresholdResponseBodyThresholds {
	s.Threshold = v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholds) Validate() error {
	return dara.Validate(s)
}

type DescribeDdosThresholdResponseBodyThresholdsThreshold struct {
	// If the value of the **DdosType*	- parameter is **defense**, the Bps parameter indicates the current traffic scrubbing threshold. Unit: Mbit/s.
	//
	// If the value of the **DdosType*	- parameter is **blackhole**, the Bps parameter indicates the basic protection threshold. Unit: Mbit/s.
	//
	// example:
	//
	// 500
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
	// i-bp10bclrt56fblts****
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
	// 1024
	MaxBps *int32 `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	// The maximum packet scrubbing threshold. Unit: pps.
	//
	// example:
	//
	// 150000
	MaxPps *int32 `json:"MaxPps,omitempty" xml:"MaxPps,omitempty"`
	// The packet scrubbing threshold. Unit: pps.
	//
	// > This parameter is returned only when the value of the **DdosType*	- parameter is **defense**.
	//
	// example:
	//
	// 150000
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeDdosThresholdResponseBodyThresholdsThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosThresholdResponseBodyThresholdsThreshold) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) GetBps() *int32 {
	return s.Bps
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) GetDdosType() *string {
	return s.DdosType
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) GetElasticBps() *int32 {
	return s.ElasticBps
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) GetIsAuto() *bool {
	return s.IsAuto
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) GetMaxBps() *int32 {
	return s.MaxBps
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) GetMaxPps() *int32 {
	return s.MaxPps
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) GetPps() *int32 {
	return s.Pps
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetBps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.Bps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetDdosType(v string) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetElasticBps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.ElasticBps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetInstanceId(v string) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetInternetIp(v string) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.InternetIp = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetIsAuto(v bool) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.IsAuto = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetMaxBps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.MaxBps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetMaxPps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.MaxPps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetPps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.Pps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) Validate() error {
	return dara.Validate(s)
}
