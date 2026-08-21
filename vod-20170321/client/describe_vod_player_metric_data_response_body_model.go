// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeVodPlayerMetricDataResponseBodyDataList) *DescribeVodPlayerMetricDataResponseBody
	GetDataList() []*DescribeVodPlayerMetricDataResponseBodyDataList
	SetExtend(v *DescribeVodPlayerMetricDataResponseBodyExtend) *DescribeVodPlayerMetricDataResponseBody
	GetExtend() *DescribeVodPlayerMetricDataResponseBodyExtend
	SetPageNumber(v int64) *DescribeVodPlayerMetricDataResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeVodPlayerMetricDataResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeVodPlayerMetricDataResponseBody
	GetRequestId() *string
	SetTotalCnt(v int64) *DescribeVodPlayerMetricDataResponseBody
	GetTotalCnt() *int64
}

type DescribeVodPlayerMetricDataResponseBody struct {
	// The list of returned data.
	DataList []*DescribeVodPlayerMetricDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The extended information.
	Extend *DescribeVodPlayerMetricDataResponseBodyExtend `json:"Extend,omitempty" xml:"Extend,omitempty" type:"Struct"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 5000
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of data points.
	//
	// example:
	//
	// 500
	TotalCnt *int64 `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeVodPlayerMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerMetricDataResponseBody) GetDataList() []*DescribeVodPlayerMetricDataResponseBodyDataList {
	return s.DataList
}

func (s *DescribeVodPlayerMetricDataResponseBody) GetExtend() *DescribeVodPlayerMetricDataResponseBodyExtend {
	return s.Extend
}

func (s *DescribeVodPlayerMetricDataResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeVodPlayerMetricDataResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodPlayerMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodPlayerMetricDataResponseBody) GetTotalCnt() *int64 {
	return s.TotalCnt
}

func (s *DescribeVodPlayerMetricDataResponseBody) SetDataList(v []*DescribeVodPlayerMetricDataResponseBodyDataList) *DescribeVodPlayerMetricDataResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBody) SetExtend(v *DescribeVodPlayerMetricDataResponseBodyExtend) *DescribeVodPlayerMetricDataResponseBody {
	s.Extend = v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBody) SetPageNumber(v int64) *DescribeVodPlayerMetricDataResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBody) SetPageSize(v int64) *DescribeVodPlayerMetricDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBody) SetRequestId(v string) *DescribeVodPlayerMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBody) SetTotalCnt(v int64) *DescribeVodPlayerMetricDataResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Extend != nil {
		if err := s.Extend.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodPlayerMetricDataResponseBodyDataList struct {
	// The average completion plays per user.
	//
	// example:
	//
	// 2
	AvgPerCompletionVv *float64 `json:"AvgPerCompletionVv,omitempty" xml:"AvgPerCompletionVv,omitempty"`
	// The average play duration per user, in milliseconds (ms).
	//
	// example:
	//
	// 100000
	AvgPerPlayDuration *float64 `json:"AvgPerPlayDuration,omitempty" xml:"AvgPerPlayDuration,omitempty"`
	// The average plays per user.
	//
	// example:
	//
	// 10
	AvgPerVv *float64 `json:"AvgPerVv,omitempty" xml:"AvgPerVv,omitempty"`
	// The average playback bitrate, in bps.
	//
	// example:
	//
	// 100
	AvgPlayBitrate *float64 `json:"AvgPlayBitrate,omitempty" xml:"AvgPlayBitrate,omitempty"`
	// The average play duration, in milliseconds (ms).
	//
	// example:
	//
	// 100000
	AvgPlayDuration *float64 `json:"AvgPlayDuration,omitempty" xml:"AvgPlayDuration,omitempty"`
	// The average start bitrate, in bps.
	//
	// example:
	//
	// 100
	AvgStartBitrate *float64 `json:"AvgStartBitrate,omitempty" xml:"AvgStartBitrate,omitempty"`
	// The average video duration, in milliseconds (ms).
	//
	// example:
	//
	// 100000
	AvgVideoDuration *float64 `json:"AvgVideoDuration,omitempty" xml:"AvgVideoDuration,omitempty"`
	// The completion rate.
	//
	// example:
	//
	// 0.8
	CompletionRate *float64 `json:"CompletionRate,omitempty" xml:"CompletionRate,omitempty"`
	// The completion count.
	//
	// example:
	//
	// 500
	CompletionVv *float64 `json:"CompletionVv,omitempty" xml:"CompletionVv,omitempty"`
	// The dimension.
	//
	// example:
	//
	// H265_MP4_WIFI
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The error count per 100 seconds.
	//
	// example:
	//
	// 1
	ErrorCount100s *float64 `json:"ErrorCount100s,omitempty" xml:"ErrorCount100s,omitempty"`
	// The first frame time, in milliseconds (ms).
	//
	// example:
	//
	// 100
	FirstFrame *float64 `json:"FirstFrame,omitempty" xml:"FirstFrame,omitempty"`
	// The 5-second bounce rate.
	//
	// example:
	//
	// 0.1
	JumpRate5s *float64 `json:"JumpRate5s,omitempty" xml:"JumpRate5s,omitempty"`
	// The play failure rate.
	//
	// example:
	//
	// 0.2
	PlayFailRate *float64 `json:"PlayFailRate,omitempty" xml:"PlayFailRate,omitempty"`
	// The actual play count.
	//
	// example:
	//
	// 1000
	RealVv *float64 `json:"RealVv,omitempty" xml:"RealVv,omitempty"`
	// The instant play rate.
	//
	// example:
	//
	// 0.8
	SecondPlayRate *float64 `json:"SecondPlayRate,omitempty" xml:"SecondPlayRate,omitempty"`
	// The non-play rate.
	//
	// example:
	//
	// 0.01
	SeedFailRate *float64 `json:"SeedFailRate,omitempty" xml:"SeedFailRate,omitempty"`
	// The seek duration, in milliseconds (ms).
	//
	// example:
	//
	// 100
	SeekDuration *float64 `json:"SeekDuration,omitempty" xml:"SeekDuration,omitempty"`
	// The slow play rate.
	//
	// example:
	//
	// 0.2
	SlowPlayRate *float64 `json:"SlowPlayRate,omitempty" xml:"SlowPlayRate,omitempty"`
	// The stuttering count per 100 seconds.
	//
	// example:
	//
	// 2
	StuckCount100s *string `json:"StuckCount100s,omitempty" xml:"StuckCount100s,omitempty"`
	// The stuttering rate by count.
	//
	// example:
	//
	// 0.1
	StuckCountRate *float64 `json:"StuckCountRate,omitempty" xml:"StuckCountRate,omitempty"`
	// The stuttering duration per 100 seconds, in milliseconds (ms).
	//
	// example:
	//
	// 200
	StuckDuration100s *float64 `json:"StuckDuration100s,omitempty" xml:"StuckDuration100s,omitempty"`
	// The timestamp of the data. Format: yyyy-mm-ddthh:mm:ssz (UTC).
	//
	// example:
	//
	// 2025-06-24T00:55:06Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total play duration, in milliseconds (ms).
	//
	// example:
	//
	// 100000
	TotalPlayDuration *float64 `json:"TotalPlayDuration,omitempty" xml:"TotalPlayDuration,omitempty"`
	// The number of playback users.
	//
	// example:
	//
	// 500
	Uv *float64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
	// The play count.
	//
	// example:
	//
	// 800
	Vv *float64 `json:"Vv,omitempty" xml:"Vv,omitempty"`
}

func (s DescribeVodPlayerMetricDataResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerMetricDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetAvgPerCompletionVv() *float64 {
	return s.AvgPerCompletionVv
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetAvgPerPlayDuration() *float64 {
	return s.AvgPerPlayDuration
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetAvgPerVv() *float64 {
	return s.AvgPerVv
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetAvgPlayBitrate() *float64 {
	return s.AvgPlayBitrate
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetAvgPlayDuration() *float64 {
	return s.AvgPlayDuration
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetAvgStartBitrate() *float64 {
	return s.AvgStartBitrate
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetAvgVideoDuration() *float64 {
	return s.AvgVideoDuration
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetCompletionRate() *float64 {
	return s.CompletionRate
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetCompletionVv() *float64 {
	return s.CompletionVv
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetErrorCount100s() *float64 {
	return s.ErrorCount100s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetFirstFrame() *float64 {
	return s.FirstFrame
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetJumpRate5s() *float64 {
	return s.JumpRate5s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetPlayFailRate() *float64 {
	return s.PlayFailRate
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetRealVv() *float64 {
	return s.RealVv
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetSecondPlayRate() *float64 {
	return s.SecondPlayRate
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetSeedFailRate() *float64 {
	return s.SeedFailRate
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetSeekDuration() *float64 {
	return s.SeekDuration
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetSlowPlayRate() *float64 {
	return s.SlowPlayRate
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetStuckCount100s() *string {
	return s.StuckCount100s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetStuckCountRate() *float64 {
	return s.StuckCountRate
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetStuckDuration100s() *float64 {
	return s.StuckDuration100s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetTotalPlayDuration() *float64 {
	return s.TotalPlayDuration
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetUv() *float64 {
	return s.Uv
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) GetVv() *float64 {
	return s.Vv
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetAvgPerCompletionVv(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.AvgPerCompletionVv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetAvgPerPlayDuration(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.AvgPerPlayDuration = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetAvgPerVv(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.AvgPerVv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetAvgPlayBitrate(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.AvgPlayBitrate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetAvgPlayDuration(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.AvgPlayDuration = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetAvgStartBitrate(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.AvgStartBitrate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetAvgVideoDuration(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.AvgVideoDuration = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetCompletionRate(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.CompletionRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetCompletionVv(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.CompletionVv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetDimension(v string) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.Dimension = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetErrorCount100s(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.ErrorCount100s = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetFirstFrame(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.FirstFrame = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetJumpRate5s(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.JumpRate5s = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetPlayFailRate(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.PlayFailRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetRealVv(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.RealVv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetSecondPlayRate(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.SecondPlayRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetSeedFailRate(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.SeedFailRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetSeekDuration(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.SeekDuration = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetSlowPlayRate(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.SlowPlayRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetStuckCount100s(v string) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.StuckCount100s = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetStuckCountRate(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.StuckCountRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetStuckDuration100s(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.StuckDuration100s = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetTimeStamp(v string) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetTotalPlayDuration(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.TotalPlayDuration = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetUv(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.Uv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) SetVv(v float64) *DescribeVodPlayerMetricDataResponseBodyDataList {
	s.Vv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeVodPlayerMetricDataResponseBodyExtend struct {
	// The latest time with data points within the specified time range. Format: yyyy-mm-ddthh:mm:ssz (UTC).
	//
	// example:
	//
	// 2025-06-05T15:59:59Z
	ActualEndTime *string `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	// The earliest time with data points within the specified time range. Format: yyyy-mm-ddthh:mm:ssz (UTC).
	//
	// example:
	//
	// 2025-06-24T00:55:06Z
	ActualStartTime *string `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	// The time granularity of the queried data, in seconds (s).
	//
	// example:
	//
	// 3600
	IntervalSeconds *int64 `json:"IntervalSeconds,omitempty" xml:"IntervalSeconds,omitempty"`
}

func (s DescribeVodPlayerMetricDataResponseBodyExtend) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerMetricDataResponseBodyExtend) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerMetricDataResponseBodyExtend) GetActualEndTime() *string {
	return s.ActualEndTime
}

func (s *DescribeVodPlayerMetricDataResponseBodyExtend) GetActualStartTime() *string {
	return s.ActualStartTime
}

func (s *DescribeVodPlayerMetricDataResponseBodyExtend) GetIntervalSeconds() *int64 {
	return s.IntervalSeconds
}

func (s *DescribeVodPlayerMetricDataResponseBodyExtend) SetActualEndTime(v string) *DescribeVodPlayerMetricDataResponseBodyExtend {
	s.ActualEndTime = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyExtend) SetActualStartTime(v string) *DescribeVodPlayerMetricDataResponseBodyExtend {
	s.ActualStartTime = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyExtend) SetIntervalSeconds(v int64) *DescribeVodPlayerMetricDataResponseBodyExtend {
	s.IntervalSeconds = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponseBodyExtend) Validate() error {
	return dara.Validate(s)
}
