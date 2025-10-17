// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcPeakChannelCntDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPeakChannelCntDataPerInterval(v *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval) *DescribeRtcPeakChannelCntDataResponseBody
	GetPeakChannelCntDataPerInterval() *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval
	SetRequestId(v string) *DescribeRtcPeakChannelCntDataResponseBody
	GetRequestId() *string
}

type DescribeRtcPeakChannelCntDataResponseBody struct {
	PeakChannelCntDataPerInterval *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval `json:"PeakChannelCntDataPerInterval,omitempty" xml:"PeakChannelCntDataPerInterval,omitempty" type:"Struct"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRtcPeakChannelCntDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataResponseBody) GetPeakChannelCntDataPerInterval() *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval {
	return s.PeakChannelCntDataPerInterval
}

func (s *DescribeRtcPeakChannelCntDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRtcPeakChannelCntDataResponseBody) SetPeakChannelCntDataPerInterval(v *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval) *DescribeRtcPeakChannelCntDataResponseBody {
	s.PeakChannelCntDataPerInterval = v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponseBody) SetRequestId(v string) *DescribeRtcPeakChannelCntDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponseBody) Validate() error {
	if s.PeakChannelCntDataPerInterval != nil {
		if err := s.PeakChannelCntDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval struct {
	PeakChannelCntModule []*DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule `json:"PeakChannelCntModule,omitempty" xml:"PeakChannelCntModule,omitempty" type:"Repeated"`
}

func (s DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval) GetPeakChannelCntModule() []*DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule {
	return s.PeakChannelCntModule
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval) SetPeakChannelCntModule(v []*DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval {
	s.PeakChannelCntModule = v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval) Validate() error {
	if s.PeakChannelCntModule != nil {
		for _, item := range s.PeakChannelCntModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule struct {
	// example:
	//
	// 10
	ActiveChannelPeak *int64 `json:"ActiveChannelPeak,omitempty" xml:"ActiveChannelPeak,omitempty"`
	// example:
	//
	// 2018-01-29T00:01:00Z
	ActiveChannelPeakTime *string `json:"ActiveChannelPeakTime,omitempty" xml:"ActiveChannelPeakTime,omitempty"`
	// example:
	//
	// 2018-01-29T00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) GetActiveChannelPeak() *int64 {
	return s.ActiveChannelPeak
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) GetActiveChannelPeakTime() *string {
	return s.ActiveChannelPeakTime
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) SetActiveChannelPeak(v int64) *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule {
	s.ActiveChannelPeak = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) SetActiveChannelPeakTime(v string) *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule {
	s.ActiveChannelPeakTime = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) SetTimeStamp(v string) *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) Validate() error {
	return dara.Validate(s)
}
