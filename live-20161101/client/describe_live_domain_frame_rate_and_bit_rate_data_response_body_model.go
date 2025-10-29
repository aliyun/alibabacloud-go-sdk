// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainFrameRateAndBitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFrameRateAndBitRateInfos(v *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) *DescribeLiveDomainFrameRateAndBitRateDataResponseBody
	GetFrameRateAndBitRateInfos() *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos
	SetRequestId(v string) *DescribeLiveDomainFrameRateAndBitRateDataResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainFrameRateAndBitRateDataResponseBody struct {
	// The frame rates and bitrates of the live streams that were queried.
	FrameRateAndBitRateInfos *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos `json:"FrameRateAndBitRateInfos,omitempty" xml:"FrameRateAndBitRateInfos,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C3F2C2C4-59BB-4B62-81FF-345BE557E3E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBody) GetFrameRateAndBitRateInfos() *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos {
	return s.FrameRateAndBitRateInfos
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBody) SetFrameRateAndBitRateInfos(v *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) *DescribeLiveDomainFrameRateAndBitRateDataResponseBody {
	s.FrameRateAndBitRateInfos = v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBody) SetRequestId(v string) *DescribeLiveDomainFrameRateAndBitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBody) Validate() error {
	if s.FrameRateAndBitRateInfos != nil {
		if err := s.FrameRateAndBitRateInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos struct {
	FrameRateAndBitRateInfo []*DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo `json:"FrameRateAndBitRateInfo,omitempty" xml:"FrameRateAndBitRateInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) GetFrameRateAndBitRateInfo() []*DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	return s.FrameRateAndBitRateInfo
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) SetFrameRateAndBitRateInfo(v []*DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos {
	s.FrameRateAndBitRateInfo = v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfos) Validate() error {
	if s.FrameRateAndBitRateInfo != nil {
		for _, item := range s.FrameRateAndBitRateInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo struct {
	// The audio frame rate of the live stream. Unit: FPS.
	//
	// example:
	//
	// 42.9
	AudioFrameRate *float32 `json:"AudioFrameRate,omitempty" xml:"AudioFrameRate,omitempty"`
	// The bitrate of the live stream. Unit: bit/s.
	//
	// example:
	//
	// 30693.96
	BitRate *float32 `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	// The URL of the live stream.
	//
	// example:
	//
	// rtmp://demo.aliyundoc.com/test/liveStream****_3_4
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
	// The video frame rate of the live stream. Unit: frames per second (FPS).
	//
	// example:
	//
	// 24.9
	VideoFrameRate *float32 `json:"VideoFrameRate,omitempty" xml:"VideoFrameRate,omitempty"`
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GetAudioFrameRate() *float32 {
	return s.AudioFrameRate
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GetBitRate() *float32 {
	return s.BitRate
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GetVideoFrameRate() *float32 {
	return s.VideoFrameRate
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetAudioFrameRate(v float32) *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.AudioFrameRate = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetBitRate(v float32) *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.BitRate = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetStreamUrl(v string) *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.StreamUrl = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetVideoFrameRate(v float32) *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.VideoFrameRate = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseBodyFrameRateAndBitRateInfosFrameRateAndBitRateInfo) Validate() error {
	return dara.Validate(s)
}
