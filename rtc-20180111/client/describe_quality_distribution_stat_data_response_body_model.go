// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityDistributionStatDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQualityStatDataList(v []*DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) *DescribeQualityDistributionStatDataResponseBody
	GetQualityStatDataList() []*DescribeQualityDistributionStatDataResponseBodyQualityStatDataList
	SetRequestId(v string) *DescribeQualityDistributionStatDataResponseBody
	GetRequestId() *string
}

type DescribeQualityDistributionStatDataResponseBody struct {
	// The list of quality statistics data.
	QualityStatDataList []*DescribeQualityDistributionStatDataResponseBodyQualityStatDataList `json:"QualityStatDataList,omitempty" xml:"QualityStatDataList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityDistributionStatDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataResponseBody) GetQualityStatDataList() []*DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	return s.QualityStatDataList
}

func (s *DescribeQualityDistributionStatDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQualityDistributionStatDataResponseBody) SetQualityStatDataList(v []*DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) *DescribeQualityDistributionStatDataResponseBody {
	s.QualityStatDataList = v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBody) SetRequestId(v string) *DescribeQualityDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBody) Validate() error {
	if s.QualityStatDataList != nil {
		for _, item := range s.QualityStatDataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeQualityDistributionStatDataResponseBodyQualityStatDataList struct {
	// Audio delay, in milliseconds.
	//
	// example:
	//
	// 554
	AudioDelay *int64 `json:"AudioDelay,omitempty" xml:"AudioDelay,omitempty"`
	// Audio high-quality transmission rate, represented as a four-digit decimal.
	//
	// example:
	//
	// 0.9953
	AudioHighQualityTransmissionRate *string `json:"AudioHighQualityTransmissionRate,omitempty" xml:"AudioHighQualityTransmissionRate,omitempty"`
	// The Audio Lag Rate, represented as a four-digit decimal.
	//
	// example:
	//
	// 0.0014
	AudioStuckRate *string `json:"AudioStuckRate,omitempty" xml:"AudioStuckRate,omitempty"`
	// The call duration ratio, represented as a four-digit decimal. For example, 1.0000 indicates that the call duration accounts for 100% of the total usage.
	//
	// example:
	//
	// 1.0000
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	// Five-second channel join success rate, represented as a four-digit decimal.
	//
	// example:
	//
	// 0.9560
	JoinChannelSucFiveSecRate *string `json:"JoinChannelSucFiveSecRate,omitempty" xml:"JoinChannelSucFiveSecRate,omitempty"`
	// The channel join success rate, represented as a four-digit decimal.
	//
	// example:
	//
	// 0.9575
	JoinChannelSucRate *string `json:"JoinChannelSucRate,omitempty" xml:"JoinChannelSucRate,omitempty"`
	// Statistical dimension. Valid values:
	//
	// - When StatDim is CHANNELONLINE: **ONE_TO_FIVE*	- (1–5 participants), **SIX_TO_TEN*	- (6–10 participants), **ELEVEN_TO_TWENTY*	- (11–20 participants), **TWENTYONE_TO_FIFTY*	- (21–50 participants), **ABOVE_FIFTY*	- (more than 50 participants).
	//
	// - When StatDim is NETWORK: **WiFi**, **4G**, etc.
	//
	// - When StatDim is OS: **iOS**, **Android**, etc.
	//
	// example:
	//
	// ONE_TO_FIVE
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Video delay, in milliseconds.
	//
	// example:
	//
	// 517
	VideoDelay *int64 `json:"VideoDelay,omitempty" xml:"VideoDelay,omitempty"`
	// Time to first video frame, in milliseconds.
	//
	// example:
	//
	// 1299
	VideoFirstPicDuration *int64 `json:"VideoFirstPicDuration,omitempty" xml:"VideoFirstPicDuration,omitempty"`
	// Video high-quality transmission rate, represented as a four-digit decimal.
	//
	// example:
	//
	// 0.9981
	VideoHighQualityTransmissionRate *string `json:"VideoHighQualityTransmissionRate,omitempty" xml:"VideoHighQualityTransmissionRate,omitempty"`
	// The Video Lag Rate, represented as a four-digit decimal.
	//
	// example:
	//
	// 0.0264
	VideoStuckRate *string `json:"VideoStuckRate,omitempty" xml:"VideoStuckRate,omitempty"`
}

func (s DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GetAudioDelay() *int64 {
	return s.AudioDelay
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GetAudioHighQualityTransmissionRate() *string {
	return s.AudioHighQualityTransmissionRate
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GetAudioStuckRate() *string {
	return s.AudioStuckRate
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GetCallDurationRatio() *string {
	return s.CallDurationRatio
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GetJoinChannelSucFiveSecRate() *string {
	return s.JoinChannelSucFiveSecRate
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GetJoinChannelSucRate() *string {
	return s.JoinChannelSucRate
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GetName() *string {
	return s.Name
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GetVideoDelay() *int64 {
	return s.VideoDelay
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GetVideoFirstPicDuration() *int64 {
	return s.VideoFirstPicDuration
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GetVideoHighQualityTransmissionRate() *string {
	return s.VideoHighQualityTransmissionRate
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GetVideoStuckRate() *string {
	return s.VideoStuckRate
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetAudioDelay(v int64) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioDelay = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetAudioHighQualityTransmissionRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetAudioStuckRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioStuckRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetCallDurationRatio(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucFiveSecRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucFiveSecRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetName(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.Name = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoDelay(v int64) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoDelay = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoFirstPicDuration(v int64) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoFirstPicDuration = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoHighQualityTransmissionRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoStuckRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoStuckRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) Validate() error {
	return dara.Validate(s)
}
