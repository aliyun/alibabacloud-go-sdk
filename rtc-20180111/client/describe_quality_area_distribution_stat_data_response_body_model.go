// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityAreaDistributionStatDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQualityStatDataList(v []*DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) *DescribeQualityAreaDistributionStatDataResponseBody
	GetQualityStatDataList() []*DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList
	SetRequestId(v string) *DescribeQualityAreaDistributionStatDataResponseBody
	GetRequestId() *string
}

type DescribeQualityAreaDistributionStatDataResponseBody struct {
	// The list of quality distribution data.
	QualityStatDataList []*DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList `json:"QualityStatDataList,omitempty" xml:"QualityStatDataList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataResponseBody) GetQualityStatDataList() []*DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	return s.QualityStatDataList
}

func (s *DescribeQualityAreaDistributionStatDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQualityAreaDistributionStatDataResponseBody) SetQualityStatDataList(v []*DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) *DescribeQualityAreaDistributionStatDataResponseBody {
	s.QualityStatDataList = v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBody) SetRequestId(v string) *DescribeQualityAreaDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBody) Validate() error {
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

type DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList struct {
	// The audio delay, in milliseconds.
	//
	// example:
	//
	// 347
	AudioDelay *int64 `json:"AudioDelay,omitempty" xml:"AudioDelay,omitempty"`
	// The high-quality transmission rate of audio, represented by a four-decimal number.
	//
	// example:
	//
	// 0.9933
	AudioHighQualityTransmissionRate *string `json:"AudioHighQualityTransmissionRate,omitempty" xml:"AudioHighQualityTransmissionRate,omitempty"`
	// The audio stuttering rate, represented by a four-decimal number.
	//
	// example:
	//
	// 0.0021
	AudioStuckRate *string `json:"AudioStuckRate,omitempty" xml:"AudioStuckRate,omitempty"`
	// The ratio of the call duration, represented by a four-decimal number. For example, 1.0000 indicates that the call duration ratio is 100%.
	//
	// example:
	//
	// 0.6654
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	// The success rate of joining a channel within 5 seconds, represented by a four-decimal number.
	//
	// example:
	//
	// 0.9338
	JoinChannelSucFiveSecRate *string `json:"JoinChannelSucFiveSecRate,omitempty" xml:"JoinChannelSucFiveSecRate,omitempty"`
	// The success rate of joining a channel, represented by a four-decimal number.
	//
	// example:
	//
	// 0.9356
	JoinChannelSucRate *string `json:"JoinChannelSucRate,omitempty" xml:"JoinChannelSucRate,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// China_Zhejiang Province
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The video delay, in milliseconds.
	//
	// example:
	//
	// 291
	VideoDelay *int64 `json:"VideoDelay,omitempty" xml:"VideoDelay,omitempty"`
	// The time taken to render the first video frame, in milliseconds.
	//
	// example:
	//
	// 1363
	VideoFirstPicDuration *int64 `json:"VideoFirstPicDuration,omitempty" xml:"VideoFirstPicDuration,omitempty"`
	// The high-quality transmission rate of video, represented by a four-decimal number.
	//
	// example:
	//
	// 0.9967
	VideoHighQualityTransmissionRate *string `json:"VideoHighQualityTransmissionRate,omitempty" xml:"VideoHighQualityTransmissionRate,omitempty"`
	// The video stuttering rate, represented by a four-decimal number.
	//
	// example:
	//
	// 0.0058
	VideoStuckRate *string `json:"VideoStuckRate,omitempty" xml:"VideoStuckRate,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GetAudioDelay() *int64 {
	return s.AudioDelay
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GetAudioHighQualityTransmissionRate() *string {
	return s.AudioHighQualityTransmissionRate
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GetAudioStuckRate() *string {
	return s.AudioStuckRate
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GetCallDurationRatio() *string {
	return s.CallDurationRatio
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GetJoinChannelSucFiveSecRate() *string {
	return s.JoinChannelSucFiveSecRate
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GetJoinChannelSucRate() *string {
	return s.JoinChannelSucRate
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GetName() *string {
	return s.Name
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GetVideoDelay() *int64 {
	return s.VideoDelay
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GetVideoFirstPicDuration() *int64 {
	return s.VideoFirstPicDuration
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GetVideoHighQualityTransmissionRate() *string {
	return s.VideoHighQualityTransmissionRate
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GetVideoStuckRate() *string {
	return s.VideoStuckRate
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetAudioDelay(v int64) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioDelay = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetAudioHighQualityTransmissionRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetAudioStuckRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioStuckRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetCallDurationRatio(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucFiveSecRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucFiveSecRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetName(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.Name = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoDelay(v int64) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoDelay = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoFirstPicDuration(v int64) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoFirstPicDuration = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoHighQualityTransmissionRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoStuckRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoStuckRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) Validate() error {
	return dara.Validate(s)
}
