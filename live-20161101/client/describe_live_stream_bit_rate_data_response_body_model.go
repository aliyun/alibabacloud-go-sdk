// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamBitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFrameRateAndBitRateInfos(v *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos) *DescribeLiveStreamBitRateDataResponseBody
	GetFrameRateAndBitRateInfos() *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos
	SetRequestId(v string) *DescribeLiveStreamBitRateDataResponseBody
	GetRequestId() *string
}

type DescribeLiveStreamBitRateDataResponseBody struct {
	// The frame rates and bitrates of the live stream.
	FrameRateAndBitRateInfos *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos `json:"FrameRateAndBitRateInfos,omitempty" xml:"FrameRateAndBitRateInfos,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveStreamBitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamBitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamBitRateDataResponseBody) GetFrameRateAndBitRateInfos() *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos {
	return s.FrameRateAndBitRateInfos
}

func (s *DescribeLiveStreamBitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamBitRateDataResponseBody) SetFrameRateAndBitRateInfos(v *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos) *DescribeLiveStreamBitRateDataResponseBody {
	s.FrameRateAndBitRateInfos = v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseBody) SetRequestId(v string) *DescribeLiveStreamBitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos struct {
	FrameRateAndBitRateInfo []*DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo `json:"FrameRateAndBitRateInfo,omitempty" xml:"FrameRateAndBitRateInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos) GetFrameRateAndBitRateInfo() []*DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	return s.FrameRateAndBitRateInfo
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos) SetFrameRateAndBitRateInfo(v []*DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos {
	s.FrameRateAndBitRateInfo = v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo struct {
	// The audio frame rate of the live stream. Unit: FPS.
	//
	// example:
	//
	// 100
	AudioFrameRate *float32 `json:"AudioFrameRate,omitempty" xml:"AudioFrameRate,omitempty"`
	// The bitrate of the live stream. Unit: bit/s.
	//
	// example:
	//
	// 600
	BitRate *float32 `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	// The URL of the live stream.
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
	// The time when the data was collected. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-09-13T16:04:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The video frame rate of the live stream. Unit: frames per second (FPS).
	//
	// example:
	//
	// 30
	VideoFrameRate *float32 `json:"VideoFrameRate,omitempty" xml:"VideoFrameRate,omitempty"`
}

func (s DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GetAudioFrameRate() *float32 {
	return s.AudioFrameRate
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GetBitRate() *float32 {
	return s.BitRate
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GetTime() *string {
	return s.Time
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GetVideoFrameRate() *float32 {
	return s.VideoFrameRate
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetAudioFrameRate(v float32) *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.AudioFrameRate = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetBitRate(v float32) *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.BitRate = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetStreamUrl(v string) *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.StreamUrl = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetTime(v string) *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.Time = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetVideoFrameRate(v float32) *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.VideoFrameRate = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) Validate() error {
	return dara.Validate(s)
}
