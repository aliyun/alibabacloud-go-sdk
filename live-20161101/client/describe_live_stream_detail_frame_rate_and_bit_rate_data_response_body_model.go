// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFrameRateAndBitRateInfos(v []*DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody
	GetFrameRateAndBitRateInfos() []*DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos
	SetRequestId(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody
	GetRequestId() *string
}

type DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody struct {
	// The audio and video frame rates and bitrates at each time granularity.
	FrameRateAndBitRateInfos []*DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos `json:"FrameRateAndBitRateInfos,omitempty" xml:"FrameRateAndBitRateInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BC858082-736F-4A25-867B-E5B67C85ACF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody) GetFrameRateAndBitRateInfos() []*DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos {
	return s.FrameRateAndBitRateInfos
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody) SetFrameRateAndBitRateInfos(v []*DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody {
	s.FrameRateAndBitRateInfos = v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody) SetRequestId(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody) Validate() error {
	if s.FrameRateAndBitRateInfos != nil {
		for _, item := range s.FrameRateAndBitRateInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos struct {
	// The audio bitrate of the live stream. Unit: bit/s.
	//
	// example:
	//
	// 200
	AudioBitRate *float32 `json:"AudioBitRate,omitempty" xml:"AudioBitRate,omitempty"`
	// The audio frame rate of the live stream. Unit: FPS.
	//
	// example:
	//
	// 60
	AudioFrameRate *float32 `json:"AudioFrameRate,omitempty" xml:"AudioFrameRate,omitempty"`
	// The bitrate of the live stream. Unit: bit/s.
	//
	// example:
	//
	// 1420
	BitRate *float32 `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	// The URL of the live stream.
	//
	// example:
	//
	// rtmp://example.com/AppName/exampleStreamName
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
	// The time when the data was collected. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-09-13T16:04:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The video bitrate of the live stream. Unit: bit/s.
	//
	// example:
	//
	// 1200
	VideoBitRate *float32 `json:"VideoBitRate,omitempty" xml:"VideoBitRate,omitempty"`
	// The video frame rate of the live stream. Unit: frames per second (FPS).
	//
	// example:
	//
	// 30
	VideoFrameRate *float32 `json:"VideoFrameRate,omitempty" xml:"VideoFrameRate,omitempty"`
}

func (s DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) GetAudioBitRate() *float32 {
	return s.AudioBitRate
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) GetAudioFrameRate() *float32 {
	return s.AudioFrameRate
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) GetBitRate() *float32 {
	return s.BitRate
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) GetTime() *string {
	return s.Time
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) GetVideoBitRate() *float32 {
	return s.VideoBitRate
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) GetVideoFrameRate() *float32 {
	return s.VideoFrameRate
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) SetAudioBitRate(v float32) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos {
	s.AudioBitRate = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) SetAudioFrameRate(v float32) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos {
	s.AudioFrameRate = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) SetBitRate(v float32) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos {
	s.BitRate = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) SetStreamUrl(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos {
	s.StreamUrl = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) SetTime(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos {
	s.Time = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) SetVideoBitRate(v float32) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos {
	s.VideoBitRate = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) SetVideoFrameRate(v float32) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos {
	s.VideoFrameRate = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) Validate() error {
	return dara.Validate(s)
}
