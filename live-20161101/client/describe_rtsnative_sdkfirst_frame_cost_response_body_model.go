// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKFirstFrameCostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBody
	GetDataInterval() *string
	SetEndTime(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBody
	GetEndTime() *string
	SetFirstFrameCostData(v []*DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) *DescribeRTSNativeSDKFirstFrameCostResponseBody
	GetFirstFrameCostData() []*DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData
	SetRequestId(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBody
	GetStartTime() *string
}

type DescribeRTSNativeSDKFirstFrameCostResponseBody struct {
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
	FirstFrameCostData []*DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData `json:"FirstFrameCostData,omitempty" xml:"FirstFrameCostData,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// BC858082-736F-4A25-867B-E5B67C85ACF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range for which the data was queried.
	//
	// example:
	//
	// 2021-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRTSNativeSDKFirstFrameCostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKFirstFrameCostResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBody) GetFirstFrameCostData() []*DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData {
	return s.FirstFrameCostData
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBody) SetDataInterval(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBody) SetEndTime(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBody) SetFirstFrameCostData(v []*DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) *DescribeRTSNativeSDKFirstFrameCostResponseBody {
	s.FirstFrameCostData = v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBody) SetRequestId(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBody) SetStartTime(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData struct {
	// The time elapsed from initialization to connection establishment.
	//
	// example:
	//
	// 100
	Connected *string `json:"Connected,omitempty" xml:"Connected,omitempty"`
	// The time elapsed from connection establishment to subscription.
	//
	// example:
	//
	// 89
	FinishGetStreamInfo *string `json:"FinishGetStreamInfo,omitempty" xml:"FinishGetStreamInfo,omitempty"`
	// The time elapsed from first packet processing to display of the first frame.
	//
	// example:
	//
	// 32
	FirstFrameComplete *string `json:"FirstFrameComplete,omitempty" xml:"FirstFrameComplete,omitempty"`
	// The time elapsed from subscription to first packet processing.
	//
	// example:
	//
	// 132
	FirstPacket *string `json:"FirstPacket,omitempty" xml:"FirstPacket,omitempty"`
	// The time consumed by initialization.
	//
	// example:
	//
	// 1100
	Initialized *string `json:"Initialized,omitempty" xml:"Initialized,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2021-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) GetConnected() *string {
	return s.Connected
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) GetFinishGetStreamInfo() *string {
	return s.FinishGetStreamInfo
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) GetFirstFrameComplete() *string {
	return s.FirstFrameComplete
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) GetFirstPacket() *string {
	return s.FirstPacket
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) GetInitialized() *string {
	return s.Initialized
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) SetConnected(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData {
	s.Connected = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) SetFinishGetStreamInfo(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData {
	s.FinishGetStreamInfo = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) SetFirstFrameComplete(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData {
	s.FirstFrameComplete = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) SetFirstPacket(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData {
	s.FirstPacket = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) SetInitialized(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData {
	s.Initialized = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) SetTimeStamp(v string) *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponseBodyFirstFrameCostData) Validate() error {
	return dara.Validate(s)
}
