// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayQosListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *DescribePlayQosListResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *DescribePlayQosListResponseBody
	GetPageSize() *int64
	SetQosInfoList(v []*DescribePlayQosListResponseBodyQosInfoList) *DescribePlayQosListResponseBody
	GetQosInfoList() []*DescribePlayQosListResponseBodyQosInfoList
	SetRequestId(v string) *DescribePlayQosListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePlayQosListResponseBody
	GetTotalCount() *int64
}

type DescribePlayQosListResponseBody struct {
	PageNo      *int64                                        `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int64                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QosInfoList []*DescribePlayQosListResponseBodyQosInfoList `json:"QosInfoList,omitempty" xml:"QosInfoList,omitempty" type:"Repeated"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePlayQosListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayQosListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayQosListResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribePlayQosListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePlayQosListResponseBody) GetQosInfoList() []*DescribePlayQosListResponseBodyQosInfoList {
	return s.QosInfoList
}

func (s *DescribePlayQosListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayQosListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePlayQosListResponseBody) SetPageNo(v int64) *DescribePlayQosListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribePlayQosListResponseBody) SetPageSize(v int64) *DescribePlayQosListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePlayQosListResponseBody) SetQosInfoList(v []*DescribePlayQosListResponseBodyQosInfoList) *DescribePlayQosListResponseBody {
	s.QosInfoList = v
	return s
}

func (s *DescribePlayQosListResponseBody) SetRequestId(v string) *DescribePlayQosListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayQosListResponseBody) SetTotalCount(v int64) *DescribePlayQosListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePlayQosListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePlayQosListResponseBodyQosInfoList struct {
	QosFirstFrame        *float32 `json:"QosFirstFrame,omitempty" xml:"QosFirstFrame,omitempty"`
	QosKbps              *float32 `json:"QosKbps,omitempty" xml:"QosKbps,omitempty"`
	QosPlay              *float32 `json:"QosPlay,omitempty" xml:"QosPlay,omitempty"`
	QosPlayFail          *float32 `json:"QosPlayFail,omitempty" xml:"QosPlayFail,omitempty"`
	QosRealPlay          *float32 `json:"QosRealPlay,omitempty" xml:"QosRealPlay,omitempty"`
	QosSecondPlayRate    *float32 `json:"QosSecondPlayRate,omitempty" xml:"QosSecondPlayRate,omitempty"`
	QosSeedFailRate      *float32 `json:"QosSeedFailRate,omitempty" xml:"QosSeedFailRate,omitempty"`
	QosSeekDuration      *float64 `json:"QosSeekDuration,omitempty" xml:"QosSeekDuration,omitempty"`
	QosSlowPlayRate      *float32 `json:"QosSlowPlayRate,omitempty" xml:"QosSlowPlayRate,omitempty"`
	QosStuckDuration100s *float64 `json:"QosStuckDuration100s,omitempty" xml:"QosStuckDuration100s,omitempty"`
	QosStuckRate         *float32 `json:"QosStuckRate,omitempty" xml:"QosStuckRate,omitempty"`
	TraceId              *string  `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribePlayQosListResponseBodyQosInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayQosListResponseBodyQosInfoList) GoString() string {
	return s.String()
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetQosFirstFrame() *float32 {
	return s.QosFirstFrame
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetQosKbps() *float32 {
	return s.QosKbps
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetQosPlay() *float32 {
	return s.QosPlay
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetQosPlayFail() *float32 {
	return s.QosPlayFail
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetQosRealPlay() *float32 {
	return s.QosRealPlay
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetQosSecondPlayRate() *float32 {
	return s.QosSecondPlayRate
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetQosSeedFailRate() *float32 {
	return s.QosSeedFailRate
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetQosSeekDuration() *float64 {
	return s.QosSeekDuration
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetQosSlowPlayRate() *float32 {
	return s.QosSlowPlayRate
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetQosStuckDuration100s() *float64 {
	return s.QosStuckDuration100s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetQosStuckRate() *float32 {
	return s.QosStuckRate
}

func (s *DescribePlayQosListResponseBodyQosInfoList) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetQosFirstFrame(v float32) *DescribePlayQosListResponseBodyQosInfoList {
	s.QosFirstFrame = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetQosKbps(v float32) *DescribePlayQosListResponseBodyQosInfoList {
	s.QosKbps = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetQosPlay(v float32) *DescribePlayQosListResponseBodyQosInfoList {
	s.QosPlay = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetQosPlayFail(v float32) *DescribePlayQosListResponseBodyQosInfoList {
	s.QosPlayFail = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetQosRealPlay(v float32) *DescribePlayQosListResponseBodyQosInfoList {
	s.QosRealPlay = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetQosSecondPlayRate(v float32) *DescribePlayQosListResponseBodyQosInfoList {
	s.QosSecondPlayRate = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetQosSeedFailRate(v float32) *DescribePlayQosListResponseBodyQosInfoList {
	s.QosSeedFailRate = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetQosSeekDuration(v float64) *DescribePlayQosListResponseBodyQosInfoList {
	s.QosSeekDuration = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetQosSlowPlayRate(v float32) *DescribePlayQosListResponseBodyQosInfoList {
	s.QosSlowPlayRate = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetQosStuckDuration100s(v float64) *DescribePlayQosListResponseBodyQosInfoList {
	s.QosStuckDuration100s = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetQosStuckRate(v float32) *DescribePlayQosListResponseBodyQosInfoList {
	s.QosStuckRate = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) SetTraceId(v string) *DescribePlayQosListResponseBodyQosInfoList {
	s.TraceId = &v
	return s
}

func (s *DescribePlayQosListResponseBodyQosInfoList) Validate() error {
	return dara.Validate(s)
}
