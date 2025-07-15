// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterLiveBypassDurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioSummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody
	GetAudioSummaryDuration() *int64
	SetData(v []*DescribeMeterLiveBypassDurationResponseBodyData) *DescribeMeterLiveBypassDurationResponseBody
	GetData() []*DescribeMeterLiveBypassDurationResponseBodyData
	SetRequestId(v string) *DescribeMeterLiveBypassDurationResponseBody
	GetRequestId() *string
	SetSingleAudioSummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody
	GetSingleAudioSummaryDuration() *int64
	SetSingleVideoSummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody
	GetSingleVideoSummaryDuration() *int64
	SetTotalSummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody
	GetTotalSummaryDuration() *int64
	SetV1080SummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody
	GetV1080SummaryDuration() *int64
	SetV480SummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody
	GetV480SummaryDuration() *int64
	SetV720SummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody
	GetV720SummaryDuration() *int64
}

type DescribeMeterLiveBypassDurationResponseBody struct {
	// The total audio-only duration. Audio-only is a basic specification. Unit: minutes.
	//
	// example:
	//
	// 20
	AudioSummaryDuration *int64 `json:"AudioSummaryDuration,omitempty" xml:"AudioSummaryDuration,omitempty"`
	// The usage statistics for each time granularity.
	Data []*DescribeMeterLiveBypassDurationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4B460F8B-993C-4F48-B98A-910811DEBFEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total single-stream relay duration for audio. Unit: minutes.
	//
	// example:
	//
	// 20
	SingleAudioSummaryDuration *int64 `json:"SingleAudioSummaryDuration,omitempty" xml:"SingleAudioSummaryDuration,omitempty"`
	// The total single-stream relay duration for video. Unit: minutes.
	//
	// example:
	//
	// 30
	SingleVideoSummaryDuration *int64 `json:"SingleVideoSummaryDuration,omitempty" xml:"SingleVideoSummaryDuration,omitempty"`
	// The total duration. Unit: minutes.
	//
	// example:
	//
	// 150
	TotalSummaryDuration *int64 `json:"TotalSummaryDuration,omitempty" xml:"TotalSummaryDuration,omitempty"`
	// The total Full HD duration. The video resolution is 1920 × 1080 or lower. Unit: minutes.
	//
	// example:
	//
	// 10
	V1080SummaryDuration *int64 `json:"V1080SummaryDuration,omitempty" xml:"V1080SummaryDuration,omitempty"`
	// The total standard definition (SD) duration. The video resolution is 640 × 480 or lower. Unit: minutes.
	//
	// example:
	//
	// 30
	V480SummaryDuration *int64 `json:"V480SummaryDuration,omitempty" xml:"V480SummaryDuration,omitempty"`
	// The total high definition (HD) duration. The video resolution is 1280 × 720 or lower. Unit: minutes.
	//
	// example:
	//
	// 40
	V720SummaryDuration *int64 `json:"V720SummaryDuration,omitempty" xml:"V720SummaryDuration,omitempty"`
}

func (s DescribeMeterLiveBypassDurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterLiveBypassDurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterLiveBypassDurationResponseBody) GetAudioSummaryDuration() *int64 {
	return s.AudioSummaryDuration
}

func (s *DescribeMeterLiveBypassDurationResponseBody) GetData() []*DescribeMeterLiveBypassDurationResponseBodyData {
	return s.Data
}

func (s *DescribeMeterLiveBypassDurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMeterLiveBypassDurationResponseBody) GetSingleAudioSummaryDuration() *int64 {
	return s.SingleAudioSummaryDuration
}

func (s *DescribeMeterLiveBypassDurationResponseBody) GetSingleVideoSummaryDuration() *int64 {
	return s.SingleVideoSummaryDuration
}

func (s *DescribeMeterLiveBypassDurationResponseBody) GetTotalSummaryDuration() *int64 {
	return s.TotalSummaryDuration
}

func (s *DescribeMeterLiveBypassDurationResponseBody) GetV1080SummaryDuration() *int64 {
	return s.V1080SummaryDuration
}

func (s *DescribeMeterLiveBypassDurationResponseBody) GetV480SummaryDuration() *int64 {
	return s.V480SummaryDuration
}

func (s *DescribeMeterLiveBypassDurationResponseBody) GetV720SummaryDuration() *int64 {
	return s.V720SummaryDuration
}

func (s *DescribeMeterLiveBypassDurationResponseBody) SetAudioSummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody {
	s.AudioSummaryDuration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBody) SetData(v []*DescribeMeterLiveBypassDurationResponseBodyData) *DescribeMeterLiveBypassDurationResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBody) SetRequestId(v string) *DescribeMeterLiveBypassDurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBody) SetSingleAudioSummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody {
	s.SingleAudioSummaryDuration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBody) SetSingleVideoSummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody {
	s.SingleVideoSummaryDuration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBody) SetTotalSummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody {
	s.TotalSummaryDuration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBody) SetV1080SummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody {
	s.V1080SummaryDuration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBody) SetV480SummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody {
	s.V480SummaryDuration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBody) SetV720SummaryDuration(v int64) *DescribeMeterLiveBypassDurationResponseBody {
	s.V720SummaryDuration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMeterLiveBypassDurationResponseBodyData struct {
	// The audio-only duration. Audio-only is a basic specification. Unit: minutes.
	//
	// example:
	//
	// 20
	AudioDuration *int64 `json:"AudioDuration,omitempty" xml:"AudioDuration,omitempty"`
	// The single-stream relay duration for audio. Unit: minutes.
	//
	// example:
	//
	// 20
	SingleAudio *int64 `json:"Single_Audio,omitempty" xml:"Single_Audio,omitempty"`
	// The single-stream relay duration for video. Unit: minutes.
	//
	// example:
	//
	// 30
	SingleVideo *int64 `json:"Single_Video,omitempty" xml:"Single_Video,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2022-10-10T20:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The duration. Unit: minutes.
	//
	// example:
	//
	// 150
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// The Full HD duration. The video resolution is 1920 × 1080 or lower. Unit: minutes.
	//
	// example:
	//
	// 10
	V1080Duration *int64 `json:"V1080Duration,omitempty" xml:"V1080Duration,omitempty"`
	// The SD duration. The video resolution is 640 × 480 or lower. Unit: minutes.
	//
	// example:
	//
	// 30
	V480Duration *int64 `json:"V480Duration,omitempty" xml:"V480Duration,omitempty"`
	// The HD duration. The video resolution is 1280 × 720 or lower. Unit: minutes.
	//
	// example:
	//
	// 40
	V720Duration *int64 `json:"V720Duration,omitempty" xml:"V720Duration,omitempty"`
}

func (s DescribeMeterLiveBypassDurationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterLiveBypassDurationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) GetAudioDuration() *int64 {
	return s.AudioDuration
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) GetSingleAudio() *int64 {
	return s.SingleAudio
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) GetSingleVideo() *int64 {
	return s.SingleVideo
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) GetV1080Duration() *int64 {
	return s.V1080Duration
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) GetV480Duration() *int64 {
	return s.V480Duration
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) GetV720Duration() *int64 {
	return s.V720Duration
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) SetAudioDuration(v int64) *DescribeMeterLiveBypassDurationResponseBodyData {
	s.AudioDuration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) SetSingleAudio(v int64) *DescribeMeterLiveBypassDurationResponseBodyData {
	s.SingleAudio = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) SetSingleVideo(v int64) *DescribeMeterLiveBypassDurationResponseBodyData {
	s.SingleVideo = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) SetTimestamp(v string) *DescribeMeterLiveBypassDurationResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) SetTotalDuration(v int64) *DescribeMeterLiveBypassDurationResponseBodyData {
	s.TotalDuration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) SetV1080Duration(v int64) *DescribeMeterLiveBypassDurationResponseBodyData {
	s.V1080Duration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) SetV480Duration(v int64) *DescribeMeterLiveBypassDurationResponseBodyData {
	s.V480Duration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) SetV720Duration(v int64) *DescribeMeterLiveBypassDurationResponseBodyData {
	s.V720Duration = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
