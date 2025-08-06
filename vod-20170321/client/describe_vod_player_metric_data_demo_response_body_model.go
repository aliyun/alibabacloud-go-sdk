// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerMetricDataDemoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeVodPlayerMetricDataDemoResponseBodyDataList) *DescribeVodPlayerMetricDataDemoResponseBody
	GetDataList() []*DescribeVodPlayerMetricDataDemoResponseBodyDataList
	SetExtend(v *DescribeVodPlayerMetricDataDemoResponseBodyExtend) *DescribeVodPlayerMetricDataDemoResponseBody
	GetExtend() *DescribeVodPlayerMetricDataDemoResponseBodyExtend
	SetPageNumber(v int64) *DescribeVodPlayerMetricDataDemoResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeVodPlayerMetricDataDemoResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeVodPlayerMetricDataDemoResponseBody
	GetRequestId() *string
	SetTotalCnt(v int64) *DescribeVodPlayerMetricDataDemoResponseBody
	GetTotalCnt() *int64
}

type DescribeVodPlayerMetricDataDemoResponseBody struct {
	DataList []*DescribeVodPlayerMetricDataDemoResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Extend   *DescribeVodPlayerMetricDataDemoResponseBodyExtend     `json:"Extend,omitempty" xml:"Extend,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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
	// example:
	//
	// 500
	TotalCnt *int64 `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeVodPlayerMetricDataDemoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerMetricDataDemoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) GetDataList() []*DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	return s.DataList
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) GetExtend() *DescribeVodPlayerMetricDataDemoResponseBodyExtend {
	return s.Extend
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) GetTotalCnt() *int64 {
	return s.TotalCnt
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) SetDataList(v []*DescribeVodPlayerMetricDataDemoResponseBodyDataList) *DescribeVodPlayerMetricDataDemoResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) SetExtend(v *DescribeVodPlayerMetricDataDemoResponseBodyExtend) *DescribeVodPlayerMetricDataDemoResponseBody {
	s.Extend = v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) SetPageNumber(v int64) *DescribeVodPlayerMetricDataDemoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) SetPageSize(v int64) *DescribeVodPlayerMetricDataDemoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) SetRequestId(v string) *DescribeVodPlayerMetricDataDemoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) SetTotalCnt(v int64) *DescribeVodPlayerMetricDataDemoResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodPlayerMetricDataDemoResponseBodyDataList struct {
	// example:
	//
	// 2
	AvgPerCompletionVv *float64 `json:"AvgPerCompletionVv,omitempty" xml:"AvgPerCompletionVv,omitempty"`
	// example:
	//
	// 100000
	AvgPerPlayDuration *float64 `json:"AvgPerPlayDuration,omitempty" xml:"AvgPerPlayDuration,omitempty"`
	// example:
	//
	// 10
	AvgPerVv *float64 `json:"AvgPerVv,omitempty" xml:"AvgPerVv,omitempty"`
	// example:
	//
	// 100
	AvgPlayBitrate *float64 `json:"AvgPlayBitrate,omitempty" xml:"AvgPlayBitrate,omitempty"`
	// example:
	//
	// 100000
	AvgPlayDuration *float64 `json:"AvgPlayDuration,omitempty" xml:"AvgPlayDuration,omitempty"`
	// example:
	//
	// 100
	AvgStartBitrate *float64 `json:"AvgStartBitrate,omitempty" xml:"AvgStartBitrate,omitempty"`
	// example:
	//
	// 100000
	AvgVideoDuration *float64 `json:"AvgVideoDuration,omitempty" xml:"AvgVideoDuration,omitempty"`
	// example:
	//
	// 100
	Bps *float64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// example:
	//
	// 0.5
	CompletionDegree *float64 `json:"CompletionDegree,omitempty" xml:"CompletionDegree,omitempty"`
	// example:
	//
	// 0.8
	CompletionRate *float64 `json:"CompletionRate,omitempty" xml:"CompletionRate,omitempty"`
	// example:
	//
	// 500
	CompletionVv *float64 `json:"CompletionVv,omitempty" xml:"CompletionVv,omitempty"`
	// example:
	//
	// 100
	Cuv *float64 `json:"Cuv,omitempty" xml:"Cuv,omitempty"`
	// example:
	//
	// H265_MP4_WIFI
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// example:
	//
	// 1
	ErrorCount100s *float64 `json:"ErrorCount100s,omitempty" xml:"ErrorCount100s,omitempty"`
	// example:
	//
	// 100
	FirstFrame *float64 `json:"FirstFrame,omitempty" xml:"FirstFrame,omitempty"`
	// example:
	//
	// 0.1
	JumpRate5s *float64 `json:"JumpRate5s,omitempty" xml:"JumpRate5s,omitempty"`
	// example:
	//
	// 0.2
	PlayFailRate *float64 `json:"PlayFailRate,omitempty" xml:"PlayFailRate,omitempty"`
	// example:
	//
	// 1000
	RealVv *float64 `json:"RealVv,omitempty" xml:"RealVv,omitempty"`
	// example:
	//
	// 0.5
	ReplayRate *float64 `json:"ReplayRate,omitempty" xml:"ReplayRate,omitempty"`
	// example:
	//
	// 0.8
	SecondPlayRate *float64 `json:"SecondPlayRate,omitempty" xml:"SecondPlayRate,omitempty"`
	// example:
	//
	// 0.01
	SeedFailRate *float64 `json:"SeedFailRate,omitempty" xml:"SeedFailRate,omitempty"`
	// example:
	//
	// 100
	SeekDuration *float64 `json:"SeekDuration,omitempty" xml:"SeekDuration,omitempty"`
	// example:
	//
	// 0.2
	SlowPlayRate *float64 `json:"SlowPlayRate,omitempty" xml:"SlowPlayRate,omitempty"`
	// example:
	//
	// 2
	StuckCount100s *string `json:"StuckCount100s,omitempty" xml:"StuckCount100s,omitempty"`
	// example:
	//
	// 0.1
	StuckCountRate *float64 `json:"StuckCountRate,omitempty" xml:"StuckCountRate,omitempty"`
	// example:
	//
	// 200
	StuckDuration100s *float64 `json:"StuckDuration100s,omitempty" xml:"StuckDuration100s,omitempty"`
	// example:
	//
	// 2025-06-24T00:55:06Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// 100000
	TotalPlayDuration *float64 `json:"TotalPlayDuration,omitempty" xml:"TotalPlayDuration,omitempty"`
	// example:
	//
	// 100
	Traf *float64 `json:"Traf,omitempty" xml:"Traf,omitempty"`
	// example:
	//
	// 500
	Uv *float64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
	// example:
	//
	// 100
	VideoBitrate *float64 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// example:
	//
	// 800
	Vv *float64 `json:"Vv,omitempty" xml:"Vv,omitempty"`
}

func (s DescribeVodPlayerMetricDataDemoResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerMetricDataDemoResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetAvgPerCompletionVv() *float64 {
	return s.AvgPerCompletionVv
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetAvgPerPlayDuration() *float64 {
	return s.AvgPerPlayDuration
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetAvgPerVv() *float64 {
	return s.AvgPerVv
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetAvgPlayBitrate() *float64 {
	return s.AvgPlayBitrate
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetAvgPlayDuration() *float64 {
	return s.AvgPlayDuration
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetAvgStartBitrate() *float64 {
	return s.AvgStartBitrate
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetAvgVideoDuration() *float64 {
	return s.AvgVideoDuration
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetBps() *float64 {
	return s.Bps
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetCompletionDegree() *float64 {
	return s.CompletionDegree
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetCompletionRate() *float64 {
	return s.CompletionRate
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetCompletionVv() *float64 {
	return s.CompletionVv
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetCuv() *float64 {
	return s.Cuv
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetErrorCount100s() *float64 {
	return s.ErrorCount100s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetFirstFrame() *float64 {
	return s.FirstFrame
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetJumpRate5s() *float64 {
	return s.JumpRate5s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetPlayFailRate() *float64 {
	return s.PlayFailRate
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetRealVv() *float64 {
	return s.RealVv
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetReplayRate() *float64 {
	return s.ReplayRate
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetSecondPlayRate() *float64 {
	return s.SecondPlayRate
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetSeedFailRate() *float64 {
	return s.SeedFailRate
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetSeekDuration() *float64 {
	return s.SeekDuration
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetSlowPlayRate() *float64 {
	return s.SlowPlayRate
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetStuckCount100s() *string {
	return s.StuckCount100s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetStuckCountRate() *float64 {
	return s.StuckCountRate
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetStuckDuration100s() *float64 {
	return s.StuckDuration100s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetTotalPlayDuration() *float64 {
	return s.TotalPlayDuration
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetTraf() *float64 {
	return s.Traf
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetUv() *float64 {
	return s.Uv
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetVideoBitrate() *float64 {
	return s.VideoBitrate
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) GetVv() *float64 {
	return s.Vv
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetAvgPerCompletionVv(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.AvgPerCompletionVv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetAvgPerPlayDuration(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.AvgPerPlayDuration = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetAvgPerVv(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.AvgPerVv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetAvgPlayBitrate(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.AvgPlayBitrate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetAvgPlayDuration(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.AvgPlayDuration = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetAvgStartBitrate(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.AvgStartBitrate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetAvgVideoDuration(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.AvgVideoDuration = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetBps(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.Bps = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetCompletionDegree(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.CompletionDegree = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetCompletionRate(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.CompletionRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetCompletionVv(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.CompletionVv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetCuv(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.Cuv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetDimension(v string) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.Dimension = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetErrorCount100s(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.ErrorCount100s = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetFirstFrame(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.FirstFrame = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetJumpRate5s(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.JumpRate5s = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetPlayFailRate(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.PlayFailRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetRealVv(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.RealVv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetReplayRate(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.ReplayRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetSecondPlayRate(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.SecondPlayRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetSeedFailRate(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.SeedFailRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetSeekDuration(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.SeekDuration = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetSlowPlayRate(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.SlowPlayRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetStuckCount100s(v string) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.StuckCount100s = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetStuckCountRate(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.StuckCountRate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetStuckDuration100s(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.StuckDuration100s = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetTimeStamp(v string) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetTotalPlayDuration(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.TotalPlayDuration = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetTraf(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.Traf = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetUv(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.Uv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetVideoBitrate(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.VideoBitrate = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) SetVv(v float64) *DescribeVodPlayerMetricDataDemoResponseBodyDataList {
	s.Vv = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeVodPlayerMetricDataDemoResponseBodyExtend struct {
	// example:
	//
	// 2025-06-05T15:59:59Z
	ActualEndTime *string `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	// example:
	//
	// 2025-06-24T00:55:06Z
	ActualStartTime *string `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	// example:
	//
	// 3600
	IntervalSeconds *int64 `json:"IntervalSeconds,omitempty" xml:"IntervalSeconds,omitempty"`
}

func (s DescribeVodPlayerMetricDataDemoResponseBodyExtend) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerMetricDataDemoResponseBodyExtend) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyExtend) GetActualEndTime() *string {
	return s.ActualEndTime
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyExtend) GetActualStartTime() *string {
	return s.ActualStartTime
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyExtend) GetIntervalSeconds() *int64 {
	return s.IntervalSeconds
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyExtend) SetActualEndTime(v string) *DescribeVodPlayerMetricDataDemoResponseBodyExtend {
	s.ActualEndTime = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyExtend) SetActualStartTime(v string) *DescribeVodPlayerMetricDataDemoResponseBodyExtend {
	s.ActualStartTime = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyExtend) SetIntervalSeconds(v int64) *DescribeVodPlayerMetricDataDemoResponseBodyExtend {
	s.IntervalSeconds = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponseBodyExtend) Validate() error {
	return dara.Validate(s)
}
