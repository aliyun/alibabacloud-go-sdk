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
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	if s.Thresholds != nil {
		if err := s.Thresholds.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.Threshold != nil {
		for _, item := range s.Threshold {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDdosThresholdResponseBodyThresholdsThreshold struct {
	Bps        *int32  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	DdosType   *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	ElasticBps *int32  `json:"ElasticBps,omitempty" xml:"ElasticBps,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	IsAuto     *bool   `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	MaxBps     *int32  `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxPps     *int32  `json:"MaxPps,omitempty" xml:"MaxPps,omitempty"`
	Pps        *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
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
