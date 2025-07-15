// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCenterStreamRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRateDatas(v []*DescribeLiveCenterStreamRateDataResponseBodyRateDatas) *DescribeLiveCenterStreamRateDataResponseBody
	GetRateDatas() []*DescribeLiveCenterStreamRateDataResponseBodyRateDatas
	SetRequestId(v string) *DescribeLiveCenterStreamRateDataResponseBody
	GetRequestId() *string
}

type DescribeLiveCenterStreamRateDataResponseBody struct {
	// The list of frame rates and bitrates.
	RateDatas []*DescribeLiveCenterStreamRateDataResponseBodyRateDatas `json:"RateDatas,omitempty" xml:"RateDatas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B6E125BE-E9B8-1103-8684-A3585CB632F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveCenterStreamRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCenterStreamRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveCenterStreamRateDataResponseBody) GetRateDatas() []*DescribeLiveCenterStreamRateDataResponseBodyRateDatas {
	return s.RateDatas
}

func (s *DescribeLiveCenterStreamRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveCenterStreamRateDataResponseBody) SetRateDatas(v []*DescribeLiveCenterStreamRateDataResponseBodyRateDatas) *DescribeLiveCenterStreamRateDataResponseBody {
	s.RateDatas = v
	return s
}

func (s *DescribeLiveCenterStreamRateDataResponseBody) SetRequestId(v string) *DescribeLiveCenterStreamRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveCenterStreamRateDataResponseBodyRateDatas struct {
	// The audio frame rate.
	//
	// example:
	//
	// 47
	AudioFps *string `json:"AudioFps,omitempty" xml:"AudioFps,omitempty"`
	// The audio bitrate.
	//
	// example:
	//
	// 600
	AudioRate *string `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	// The time when the data was collected. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-03-05T18:00:53Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The video frame rate.
	//
	// example:
	//
	// 30
	VideoFps *string `json:"VideoFps,omitempty" xml:"VideoFps,omitempty"`
	// The video bitrate.
	//
	// example:
	//
	// 1953584
	VideoRate *string `json:"VideoRate,omitempty" xml:"VideoRate,omitempty"`
}

func (s DescribeLiveCenterStreamRateDataResponseBodyRateDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCenterStreamRateDataResponseBodyRateDatas) GoString() string {
	return s.String()
}

func (s *DescribeLiveCenterStreamRateDataResponseBodyRateDatas) GetAudioFps() *string {
	return s.AudioFps
}

func (s *DescribeLiveCenterStreamRateDataResponseBodyRateDatas) GetAudioRate() *string {
	return s.AudioRate
}

func (s *DescribeLiveCenterStreamRateDataResponseBodyRateDatas) GetTime() *string {
	return s.Time
}

func (s *DescribeLiveCenterStreamRateDataResponseBodyRateDatas) GetVideoFps() *string {
	return s.VideoFps
}

func (s *DescribeLiveCenterStreamRateDataResponseBodyRateDatas) GetVideoRate() *string {
	return s.VideoRate
}

func (s *DescribeLiveCenterStreamRateDataResponseBodyRateDatas) SetAudioFps(v string) *DescribeLiveCenterStreamRateDataResponseBodyRateDatas {
	s.AudioFps = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataResponseBodyRateDatas) SetAudioRate(v string) *DescribeLiveCenterStreamRateDataResponseBodyRateDatas {
	s.AudioRate = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataResponseBodyRateDatas) SetTime(v string) *DescribeLiveCenterStreamRateDataResponseBodyRateDatas {
	s.Time = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataResponseBodyRateDatas) SetVideoFps(v string) *DescribeLiveCenterStreamRateDataResponseBodyRateDatas {
	s.VideoFps = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataResponseBodyRateDatas) SetVideoRate(v string) *DescribeLiveCenterStreamRateDataResponseBodyRateDatas {
	s.VideoRate = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataResponseBodyRateDatas) Validate() error {
	return dara.Validate(s)
}
