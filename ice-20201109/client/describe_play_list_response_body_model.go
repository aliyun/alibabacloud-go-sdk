// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int64) *DescribePlayListResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *DescribePlayListResponseBody
	GetPageSize() *int64
	SetPlayList(v []*DescribePlayListResponseBodyPlayList) *DescribePlayListResponseBody
	GetPlayList() []*DescribePlayListResponseBodyPlayList
	SetRequestId(v string) *DescribePlayListResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *DescribePlayListResponseBody
	GetTotalNum() *int64
}

type DescribePlayListResponseBody struct {
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlayList []*DescribePlayListResponseBodyPlayList `json:"PlayList,omitempty" xml:"PlayList,omitempty" type:"Repeated"`
	// Id
	//
	// example:
	//
	// B960580D-26FA-5547-8AFC-3CDC812DBF27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 49
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s DescribePlayListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayListResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribePlayListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePlayListResponseBody) GetPlayList() []*DescribePlayListResponseBodyPlayList {
	return s.PlayList
}

func (s *DescribePlayListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayListResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribePlayListResponseBody) SetPageNum(v int64) *DescribePlayListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribePlayListResponseBody) SetPageSize(v int64) *DescribePlayListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePlayListResponseBody) SetPlayList(v []*DescribePlayListResponseBodyPlayList) *DescribePlayListResponseBody {
	s.PlayList = v
	return s
}

func (s *DescribePlayListResponseBody) SetRequestId(v string) *DescribePlayListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayListResponseBody) SetTotalNum(v int64) *DescribePlayListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribePlayListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePlayListResponseBodyPlayList struct {
	// example:
	//
	// 200
	FirstFrameDuration *string `json:"FirstFrameDuration,omitempty" xml:"FirstFrameDuration,omitempty"`
	// example:
	//
	// 1000
	PlayDuration *string `json:"PlayDuration,omitempty" xml:"PlayDuration,omitempty"`
	// example:
	//
	// 1675922209572
	PlayTs *string `json:"PlayTs,omitempty" xml:"PlayTs,omitempty"`
	// example:
	//
	// vod
	PlayType *string `json:"PlayType,omitempty" xml:"PlayType,omitempty"`
	// example:
	//
	// 91488be2-8381-40c9-8494-e8afe22c4a2d
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 20
	StuckDuration *string `json:"StuckDuration,omitempty" xml:"StuckDuration,omitempty"`
	// example:
	//
	// 0b736abf16724820210842673d9543
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// example:
	//
	// 2000
	VideoDuration *string `json:"VideoDuration,omitempty" xml:"VideoDuration,omitempty"`
	// example:
	//
	// 250314203f0171eebff17035d0b20102
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DescribePlayListResponseBodyPlayList) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayListResponseBodyPlayList) GoString() string {
	return s.String()
}

func (s *DescribePlayListResponseBodyPlayList) GetFirstFrameDuration() *string {
	return s.FirstFrameDuration
}

func (s *DescribePlayListResponseBodyPlayList) GetPlayDuration() *string {
	return s.PlayDuration
}

func (s *DescribePlayListResponseBodyPlayList) GetPlayTs() *string {
	return s.PlayTs
}

func (s *DescribePlayListResponseBodyPlayList) GetPlayType() *string {
	return s.PlayType
}

func (s *DescribePlayListResponseBodyPlayList) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePlayListResponseBodyPlayList) GetStatus() *string {
	return s.Status
}

func (s *DescribePlayListResponseBodyPlayList) GetStuckDuration() *string {
	return s.StuckDuration
}

func (s *DescribePlayListResponseBodyPlayList) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribePlayListResponseBodyPlayList) GetVideoDuration() *string {
	return s.VideoDuration
}

func (s *DescribePlayListResponseBodyPlayList) GetVideoId() *string {
	return s.VideoId
}

func (s *DescribePlayListResponseBodyPlayList) SetFirstFrameDuration(v string) *DescribePlayListResponseBodyPlayList {
	s.FirstFrameDuration = &v
	return s
}

func (s *DescribePlayListResponseBodyPlayList) SetPlayDuration(v string) *DescribePlayListResponseBodyPlayList {
	s.PlayDuration = &v
	return s
}

func (s *DescribePlayListResponseBodyPlayList) SetPlayTs(v string) *DescribePlayListResponseBodyPlayList {
	s.PlayTs = &v
	return s
}

func (s *DescribePlayListResponseBodyPlayList) SetPlayType(v string) *DescribePlayListResponseBodyPlayList {
	s.PlayType = &v
	return s
}

func (s *DescribePlayListResponseBodyPlayList) SetSessionId(v string) *DescribePlayListResponseBodyPlayList {
	s.SessionId = &v
	return s
}

func (s *DescribePlayListResponseBodyPlayList) SetStatus(v string) *DescribePlayListResponseBodyPlayList {
	s.Status = &v
	return s
}

func (s *DescribePlayListResponseBodyPlayList) SetStuckDuration(v string) *DescribePlayListResponseBodyPlayList {
	s.StuckDuration = &v
	return s
}

func (s *DescribePlayListResponseBodyPlayList) SetTraceId(v string) *DescribePlayListResponseBodyPlayList {
	s.TraceId = &v
	return s
}

func (s *DescribePlayListResponseBodyPlayList) SetVideoDuration(v string) *DescribePlayListResponseBodyPlayList {
	s.VideoDuration = &v
	return s
}

func (s *DescribePlayListResponseBodyPlayList) SetVideoId(v string) *DescribePlayListResponseBodyPlayList {
	s.VideoId = &v
	return s
}

func (s *DescribePlayListResponseBodyPlayList) Validate() error {
	return dara.Validate(s)
}
