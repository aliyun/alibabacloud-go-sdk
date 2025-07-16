// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCoverUrl(v string) *QueryLiveInfoResponseBody
	GetCoverUrl() *string
	SetDuration(v int64) *QueryLiveInfoResponseBody
	GetDuration() *int64
	SetEndTime(v int64) *QueryLiveInfoResponseBody
	GetEndTime() *int64
	SetIntroduction(v string) *QueryLiveInfoResponseBody
	GetIntroduction() *string
	SetLiveId(v string) *QueryLiveInfoResponseBody
	GetLiveId() *string
	SetLivePlayUrl(v string) *QueryLiveInfoResponseBody
	GetLivePlayUrl() *string
	SetLiveStatus(v int32) *QueryLiveInfoResponseBody
	GetLiveStatus() *int32
	SetPlaybackDuration(v int64) *QueryLiveInfoResponseBody
	GetPlaybackDuration() *int64
	SetRequestId(v string) *QueryLiveInfoResponseBody
	GetRequestId() *string
	SetStartTime(v int64) *QueryLiveInfoResponseBody
	GetStartTime() *int64
	SetSubscribeCount(v int32) *QueryLiveInfoResponseBody
	GetSubscribeCount() *int32
	SetTitle(v string) *QueryLiveInfoResponseBody
	GetTitle() *string
	SetUv(v int32) *QueryLiveInfoResponseBody
	GetUv() *int32
}

type QueryLiveInfoResponseBody struct {
	// example:
	//
	// http://xxx/kk.jpg
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// example:
	//
	// 59886
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1662447951960
	EndTime      *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// example:
	//
	// 1211-3442-122
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// example:
	//
	// http://ssssss
	LivePlayUrl *string `json:"livePlayUrl,omitempty" xml:"livePlayUrl,omitempty"`
	LiveStatus  *int32  `json:"liveStatus,omitempty" xml:"liveStatus,omitempty"`
	// example:
	//
	// 13414
	PlaybackDuration *int64 `json:"playbackDuration,omitempty" xml:"playbackDuration,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1627353123000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 111
	SubscribeCount *int32  `json:"subscribeCount,omitempty" xml:"subscribeCount,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 10
	Uv *int32 `json:"uv,omitempty" xml:"uv,omitempty"`
}

func (s QueryLiveInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoResponseBody) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *QueryLiveInfoResponseBody) GetDuration() *int64 {
	return s.Duration
}

func (s *QueryLiveInfoResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryLiveInfoResponseBody) GetIntroduction() *string {
	return s.Introduction
}

func (s *QueryLiveInfoResponseBody) GetLiveId() *string {
	return s.LiveId
}

func (s *QueryLiveInfoResponseBody) GetLivePlayUrl() *string {
	return s.LivePlayUrl
}

func (s *QueryLiveInfoResponseBody) GetLiveStatus() *int32 {
	return s.LiveStatus
}

func (s *QueryLiveInfoResponseBody) GetPlaybackDuration() *int64 {
	return s.PlaybackDuration
}

func (s *QueryLiveInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLiveInfoResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryLiveInfoResponseBody) GetSubscribeCount() *int32 {
	return s.SubscribeCount
}

func (s *QueryLiveInfoResponseBody) GetTitle() *string {
	return s.Title
}

func (s *QueryLiveInfoResponseBody) GetUv() *int32 {
	return s.Uv
}

func (s *QueryLiveInfoResponseBody) SetCoverUrl(v string) *QueryLiveInfoResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetDuration(v int64) *QueryLiveInfoResponseBody {
	s.Duration = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetEndTime(v int64) *QueryLiveInfoResponseBody {
	s.EndTime = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetIntroduction(v string) *QueryLiveInfoResponseBody {
	s.Introduction = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetLiveId(v string) *QueryLiveInfoResponseBody {
	s.LiveId = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetLivePlayUrl(v string) *QueryLiveInfoResponseBody {
	s.LivePlayUrl = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetLiveStatus(v int32) *QueryLiveInfoResponseBody {
	s.LiveStatus = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetPlaybackDuration(v int64) *QueryLiveInfoResponseBody {
	s.PlaybackDuration = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetRequestId(v string) *QueryLiveInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetStartTime(v int64) *QueryLiveInfoResponseBody {
	s.StartTime = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetSubscribeCount(v int32) *QueryLiveInfoResponseBody {
	s.SubscribeCount = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetTitle(v string) *QueryLiveInfoResponseBody {
	s.Title = &v
	return s
}

func (s *QueryLiveInfoResponseBody) SetUv(v int32) *QueryLiveInfoResponseBody {
	s.Uv = &v
	return s
}

func (s *QueryLiveInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
