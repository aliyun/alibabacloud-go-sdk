// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKFirstFrameDelayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKFirstFrameDelayResponseBody
	GetDataInterval() *string
	SetEndTime(v string) *DescribeRTSNativeSDKFirstFrameDelayResponseBody
	GetEndTime() *string
	SetFrameDelayData(v []*DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData) *DescribeRTSNativeSDKFirstFrameDelayResponseBody
	GetFrameDelayData() []*DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData
	SetRequestId(v string) *DescribeRTSNativeSDKFirstFrameDelayResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeRTSNativeSDKFirstFrameDelayResponseBody
	GetStartTime() *string
}

type DescribeRTSNativeSDKFirstFrameDelayResponseBody struct {
	// The time granularity.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The end of the time range for which the data was queried.
	//
	// example:
	//
	// 2021-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The average latency of first frames at each interval. Unit: milliseconds.
	FrameDelayData []*DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData `json:"FrameDelayData,omitempty" xml:"FrameDelayData,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range for which the data was queried.
	//
	// example:
	//
	// 2021-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRTSNativeSDKFirstFrameDelayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKFirstFrameDelayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBody) GetFrameDelayData() []*DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData {
	return s.FrameDelayData
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBody) SetDataInterval(v string) *DescribeRTSNativeSDKFirstFrameDelayResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBody) SetEndTime(v string) *DescribeRTSNativeSDKFirstFrameDelayResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBody) SetFrameDelayData(v []*DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData) *DescribeRTSNativeSDKFirstFrameDelayResponseBody {
	s.FrameDelayData = v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBody) SetRequestId(v string) *DescribeRTSNativeSDKFirstFrameDelayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBody) SetStartTime(v string) *DescribeRTSNativeSDKFirstFrameDelayResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBody) Validate() error {
	if s.FrameDelayData != nil {
		for _, item := range s.FrameDelayData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData struct {
	// The average latency of first frames within the period of time.
	//
	// example:
	//
	// 400
	FrameDelay *string `json:"FrameDelay,omitempty" xml:"FrameDelay,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2021-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData) GetFrameDelay() *string {
	return s.FrameDelay
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData) SetFrameDelay(v string) *DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData {
	s.FrameDelay = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData) SetTimeStamp(v string) *DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponseBodyFrameDelayData) Validate() error {
	return dara.Validate(s)
}
