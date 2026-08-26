// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKPlayTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKPlayTimeResponseBody
	GetDataInterval() *string
	SetEndTime(v string) *DescribeRTSNativeSDKPlayTimeResponseBody
	GetEndTime() *string
	SetPlayTimeData(v []*DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData) *DescribeRTSNativeSDKPlayTimeResponseBody
	GetPlayTimeData() []*DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData
	SetRequestId(v string) *DescribeRTSNativeSDKPlayTimeResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeRTSNativeSDKPlayTimeResponseBody
	GetStartTime() *string
}

type DescribeRTSNativeSDKPlayTimeResponseBody struct {
	// The time granularity.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The end time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The average playback duration and average stall duration for each time interval. Unit: milliseconds.
	PlayTimeData []*DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData `json:"PlayTimeData,omitempty" xml:"PlayTimeData,omitempty" type:"Repeated"`
	// Id
	//
	// example:
	//
	// 7BF95F2A-3B24-4CDE-9346-7F6FA86697A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRTSNativeSDKPlayTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKPlayTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBody) GetPlayTimeData() []*DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData {
	return s.PlayTimeData
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBody) SetDataInterval(v string) *DescribeRTSNativeSDKPlayTimeResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBody) SetEndTime(v string) *DescribeRTSNativeSDKPlayTimeResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBody) SetPlayTimeData(v []*DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData) *DescribeRTSNativeSDKPlayTimeResponseBody {
	s.PlayTimeData = v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBody) SetRequestId(v string) *DescribeRTSNativeSDKPlayTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBody) SetStartTime(v string) *DescribeRTSNativeSDKPlayTimeResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBody) Validate() error {
	if s.PlayTimeData != nil {
		for _, item := range s.PlayTimeData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData struct {
	// The average playback duration during the time interval. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	PlayTime *string `json:"PlayTime,omitempty" xml:"PlayTime,omitempty"`
	// The average stall duration during the time interval. Unit: milliseconds.
	//
	// example:
	//
	// 100
	StallTime *string `json:"StallTime,omitempty" xml:"StallTime,omitempty"`
	// The start time of the time interval. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData) GetPlayTime() *string {
	return s.PlayTime
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData) GetStallTime() *string {
	return s.StallTime
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData) SetPlayTime(v string) *DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData {
	s.PlayTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData) SetStallTime(v string) *DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData {
	s.StallTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData) SetTimeStamp(v string) *DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeResponseBodyPlayTimeData) Validate() error {
	return dara.Validate(s)
}
