// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvgWatchTime(v int64) *QueryLiveWatchDetailResponseBody
	GetAvgWatchTime() *int64
	SetLiveUv(v int32) *QueryLiveWatchDetailResponseBody
	GetLiveUv() *int32
	SetMsgCount(v int32) *QueryLiveWatchDetailResponseBody
	GetMsgCount() *int32
	SetPlaybackUv(v int32) *QueryLiveWatchDetailResponseBody
	GetPlaybackUv() *int32
	SetPraiseCount(v int32) *QueryLiveWatchDetailResponseBody
	GetPraiseCount() *int32
	SetPv(v int32) *QueryLiveWatchDetailResponseBody
	GetPv() *int32
	SetRequestId(v string) *QueryLiveWatchDetailResponseBody
	GetRequestId() *string
	SetTotalWatchTime(v int64) *QueryLiveWatchDetailResponseBody
	GetTotalWatchTime() *int64
	SetUv(v int32) *QueryLiveWatchDetailResponseBody
	GetUv() *int32
}

type QueryLiveWatchDetailResponseBody struct {
	// example:
	//
	// 84600
	AvgWatchTime *int64 `json:"avgWatchTime,omitempty" xml:"avgWatchTime,omitempty"`
	// example:
	//
	// 100
	LiveUv *int32 `json:"liveUv,omitempty" xml:"liveUv,omitempty"`
	// example:
	//
	// 10000
	MsgCount *int32 `json:"msgCount,omitempty" xml:"msgCount,omitempty"`
	// example:
	//
	// 20
	PlaybackUv *int32 `json:"playbackUv,omitempty" xml:"playbackUv,omitempty"`
	// example:
	//
	// 30
	PraiseCount *int32 `json:"praiseCount,omitempty" xml:"praiseCount,omitempty"`
	// example:
	//
	// 1000
	Pv *int32 `json:"pv,omitempty" xml:"pv,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1222000
	TotalWatchTime *int64 `json:"totalWatchTime,omitempty" xml:"totalWatchTime,omitempty"`
	// example:
	//
	// 10
	Uv *int32 `json:"uv,omitempty" xml:"uv,omitempty"`
}

func (s QueryLiveWatchDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailResponseBody) GetAvgWatchTime() *int64 {
	return s.AvgWatchTime
}

func (s *QueryLiveWatchDetailResponseBody) GetLiveUv() *int32 {
	return s.LiveUv
}

func (s *QueryLiveWatchDetailResponseBody) GetMsgCount() *int32 {
	return s.MsgCount
}

func (s *QueryLiveWatchDetailResponseBody) GetPlaybackUv() *int32 {
	return s.PlaybackUv
}

func (s *QueryLiveWatchDetailResponseBody) GetPraiseCount() *int32 {
	return s.PraiseCount
}

func (s *QueryLiveWatchDetailResponseBody) GetPv() *int32 {
	return s.Pv
}

func (s *QueryLiveWatchDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLiveWatchDetailResponseBody) GetTotalWatchTime() *int64 {
	return s.TotalWatchTime
}

func (s *QueryLiveWatchDetailResponseBody) GetUv() *int32 {
	return s.Uv
}

func (s *QueryLiveWatchDetailResponseBody) SetAvgWatchTime(v int64) *QueryLiveWatchDetailResponseBody {
	s.AvgWatchTime = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBody) SetLiveUv(v int32) *QueryLiveWatchDetailResponseBody {
	s.LiveUv = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBody) SetMsgCount(v int32) *QueryLiveWatchDetailResponseBody {
	s.MsgCount = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBody) SetPlaybackUv(v int32) *QueryLiveWatchDetailResponseBody {
	s.PlaybackUv = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBody) SetPraiseCount(v int32) *QueryLiveWatchDetailResponseBody {
	s.PraiseCount = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBody) SetPv(v int32) *QueryLiveWatchDetailResponseBody {
	s.Pv = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBody) SetRequestId(v string) *QueryLiveWatchDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBody) SetTotalWatchTime(v int64) *QueryLiveWatchDetailResponseBody {
	s.TotalWatchTime = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBody) SetUv(v int32) *QueryLiveWatchDetailResponseBody {
	s.Uv = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
