// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityOsSdkVersionDistributionStatDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQualityOsSdkVersionStatDataList(v []*DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) *DescribeQualityOsSdkVersionDistributionStatDataResponseBody
	GetQualityOsSdkVersionStatDataList() []*DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList
	SetRequestId(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBody
	GetRequestId() *string
}

type DescribeQualityOsSdkVersionDistributionStatDataResponseBody struct {
	QualityOsSdkVersionStatDataList []*DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList `json:"QualityOsSdkVersionStatDataList,omitempty" xml:"QualityOsSdkVersionStatDataList,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) GetQualityOsSdkVersionStatDataList() []*DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	return s.QualityOsSdkVersionStatDataList
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) SetQualityOsSdkVersionStatDataList(v []*DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) *DescribeQualityOsSdkVersionDistributionStatDataResponseBody {
	s.QualityOsSdkVersionStatDataList = v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) SetRequestId(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) Validate() error {
	if s.QualityOsSdkVersionStatDataList != nil {
		for _, item := range s.QualityOsSdkVersionStatDataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList struct {
	// example:
	//
	// 248
	AudioDelay *int64 `json:"AudioDelay,omitempty" xml:"AudioDelay,omitempty"`
	// example:
	//
	// 0.9987
	AudioHighQualityTransmissionRate *string `json:"AudioHighQualityTransmissionRate,omitempty" xml:"AudioHighQualityTransmissionRate,omitempty"`
	// example:
	//
	// 0.0011
	AudioStuckRate *string `json:"AudioStuckRate,omitempty" xml:"AudioStuckRate,omitempty"`
	// example:
	//
	// 0.0984
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	// example:
	//
	// 0.8276
	JoinChannelSucFiveSecRate *string `json:"JoinChannelSucFiveSecRate,omitempty" xml:"JoinChannelSucFiveSecRate,omitempty"`
	// example:
	//
	// 0.8276
	JoinChannelSucRate *string `json:"JoinChannelSucRate,omitempty" xml:"JoinChannelSucRate,omitempty"`
	// example:
	//
	// 2.1.0.210316.dev--release/rtcsdk_v2.1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// macOS
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 333
	VideoDelay *int64 `json:"VideoDelay,omitempty" xml:"VideoDelay,omitempty"`
	// example:
	//
	// 5643
	VideoFirstPicDuration *int64 `json:"VideoFirstPicDuration,omitempty" xml:"VideoFirstPicDuration,omitempty"`
	// example:
	//
	// 0.9997
	VideoHighQualityTransmissionRate *string `json:"VideoHighQualityTransmissionRate,omitempty" xml:"VideoHighQualityTransmissionRate,omitempty"`
	// example:
	//
	// 0.0054
	VideoStuckRate *string `json:"VideoStuckRate,omitempty" xml:"VideoStuckRate,omitempty"`
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetAudioDelay() *int64 {
	return s.AudioDelay
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetAudioHighQualityTransmissionRate() *string {
	return s.AudioHighQualityTransmissionRate
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetAudioStuckRate() *string {
	return s.AudioStuckRate
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetCallDurationRatio() *string {
	return s.CallDurationRatio
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetJoinChannelSucFiveSecRate() *string {
	return s.JoinChannelSucFiveSecRate
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetJoinChannelSucRate() *string {
	return s.JoinChannelSucRate
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetName() *string {
	return s.Name
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetOs() *string {
	return s.Os
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetVideoDelay() *int64 {
	return s.VideoDelay
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetVideoFirstPicDuration() *int64 {
	return s.VideoFirstPicDuration
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetVideoHighQualityTransmissionRate() *string {
	return s.VideoHighQualityTransmissionRate
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GetVideoStuckRate() *string {
	return s.VideoStuckRate
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetAudioDelay(v int64) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.AudioDelay = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetAudioHighQualityTransmissionRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.AudioHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetAudioStuckRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.AudioStuckRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetCallDurationRatio(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetJoinChannelSucFiveSecRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.JoinChannelSucFiveSecRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetJoinChannelSucRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.JoinChannelSucRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetName(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.Name = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetOs(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.Os = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoDelay(v int64) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoDelay = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoFirstPicDuration(v int64) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoFirstPicDuration = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoHighQualityTransmissionRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoStuckRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoStuckRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) Validate() error {
	return dara.Validate(s)
}
