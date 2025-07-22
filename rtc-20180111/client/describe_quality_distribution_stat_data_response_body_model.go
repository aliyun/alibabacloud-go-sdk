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
	QualityStatDataList []*DescribeQualityDistributionStatDataResponseBodyQualityStatDataList `json:"QualityStatDataList,omitempty" xml:"QualityStatDataList,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type DescribeQualityDistributionStatDataResponseBodyQualityStatDataList struct {
	// example:
	//
	// 554
	AudioDelay *int64 `json:"AudioDelay,omitempty" xml:"AudioDelay,omitempty"`
	// example:
	//
	// 0.9953
	AudioHighQualityTransmissionRate *string `json:"AudioHighQualityTransmissionRate,omitempty" xml:"AudioHighQualityTransmissionRate,omitempty"`
	// example:
	//
	// 0.0014
	AudioStuckRate *string `json:"AudioStuckRate,omitempty" xml:"AudioStuckRate,omitempty"`
	// example:
	//
	// 1.0000
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	// example:
	//
	// 0.9560
	JoinChannelSucFiveSecRate *string `json:"JoinChannelSucFiveSecRate,omitempty" xml:"JoinChannelSucFiveSecRate,omitempty"`
	// example:
	//
	// 0.9575
	JoinChannelSucRate *string `json:"JoinChannelSucRate,omitempty" xml:"JoinChannelSucRate,omitempty"`
	// example:
	//
	// ONE_TO_FIVE
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 517
	VideoDelay *int64 `json:"VideoDelay,omitempty" xml:"VideoDelay,omitempty"`
	// example:
	//
	// 1299
	VideoFirstPicDuration *int64 `json:"VideoFirstPicDuration,omitempty" xml:"VideoFirstPicDuration,omitempty"`
	// example:
	//
	// 0.9981
	VideoHighQualityTransmissionRate *string `json:"VideoHighQualityTransmissionRate,omitempty" xml:"VideoHighQualityTransmissionRate,omitempty"`
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
