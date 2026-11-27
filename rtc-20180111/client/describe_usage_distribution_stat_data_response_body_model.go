// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageDistributionStatDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUsageDistributionStatDataResponseBody
	GetRequestId() *string
	SetUsageStatList(v []*DescribeUsageDistributionStatDataResponseBodyUsageStatList) *DescribeUsageDistributionStatDataResponseBody
	GetUsageStatList() []*DescribeUsageDistributionStatDataResponseBodyUsageStatList
}

type DescribeUsageDistributionStatDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of usage statistics.
	UsageStatList []*DescribeUsageDistributionStatDataResponseBodyUsageStatList `json:"UsageStatList,omitempty" xml:"UsageStatList,omitempty" type:"Repeated"`
}

func (s DescribeUsageDistributionStatDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUsageDistributionStatDataResponseBody) GetUsageStatList() []*DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	return s.UsageStatList
}

func (s *DescribeUsageDistributionStatDataResponseBody) SetRequestId(v string) *DescribeUsageDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBody) SetUsageStatList(v []*DescribeUsageDistributionStatDataResponseBodyUsageStatList) *DescribeUsageDistributionStatDataResponseBody {
	s.UsageStatList = v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBody) Validate() error {
	if s.UsageStatList != nil {
		for _, item := range s.UsageStatList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUsageDistributionStatDataResponseBodyUsageStatList struct {
	// The duration of audio calls, in minutes.
	//
	// example:
	//
	// 51
	AudioCallDuration *int64 `json:"AudioCallDuration,omitempty" xml:"AudioCallDuration,omitempty"`
	// The ratio of the call duration, represented by four decimal places, such as 0.9782.
	//
	// example:
	//
	// 0.9782
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	// The statistical dimension. Valid values:
	//
	// - When StatDim is set to CHANNEL_ONLINE, valid values are **ONE_TO_FIVE*	- (1 to 5 users), **SIX_TO_TEN*	- (6 to 10 users), **ELEVEN_TO_TWENTY*	- (11 to 20 users), **TWENTY_ONE_TO_FIFTY*	- (21 to 50 users), and **ABOVE_FIFTY*	- (more than 50 users).
	//
	// - When StatDim is set to NETWORK, valid values are **WiFi**, **4G**, and others.
	//
	// - When StatDim is set to OS, valid values are **iOS**, **Android**, and others.
	//
	// example:
	//
	// ONE_TO_FIVE
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The total duration of calls, in minutes.
	//
	// example:
	//
	// 10636
	TotalCallDuration *int64 `json:"TotalCallDuration,omitempty" xml:"TotalCallDuration,omitempty"`
	// The duration of video calls, in minutes.
	//
	// example:
	//
	// 10585
	VideoCallDuration *int64 `json:"VideoCallDuration,omitempty" xml:"VideoCallDuration,omitempty"`
}

func (s DescribeUsageDistributionStatDataResponseBodyUsageStatList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageDistributionStatDataResponseBodyUsageStatList) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) GetAudioCallDuration() *int64 {
	return s.AudioCallDuration
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) GetCallDurationRatio() *string {
	return s.CallDurationRatio
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) GetName() *string {
	return s.Name
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) GetTotalCallDuration() *int64 {
	return s.TotalCallDuration
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) GetVideoCallDuration() *int64 {
	return s.VideoCallDuration
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetAudioCallDuration(v int64) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.AudioCallDuration = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetCallDurationRatio(v string) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetName(v string) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.Name = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetTotalCallDuration(v int64) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.TotalCallDuration = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetVideoCallDuration(v int64) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.VideoCallDuration = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) Validate() error {
	return dara.Validate(s)
}
