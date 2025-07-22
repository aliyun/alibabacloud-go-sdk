// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageOsSdkVersionDistributionStatDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBody
	GetRequestId() *string
	SetUsageOsSdkVersionStatList(v []*DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) *DescribeUsageOsSdkVersionDistributionStatDataResponseBody
	GetUsageOsSdkVersionStatList() []*DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList
}

type DescribeUsageOsSdkVersionDistributionStatDataResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId                 *string                                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageOsSdkVersionStatList []*DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList `json:"UsageOsSdkVersionStatList,omitempty" xml:"UsageOsSdkVersionStatList,omitempty" type:"Repeated"`
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) GetUsageOsSdkVersionStatList() []*DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	return s.UsageOsSdkVersionStatList
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) SetRequestId(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) SetUsageOsSdkVersionStatList(v []*DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) *DescribeUsageOsSdkVersionDistributionStatDataResponseBody {
	s.UsageOsSdkVersionStatList = v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList struct {
	// example:
	//
	// 3
	AudioCallDuration *int64 `json:"AudioCallDuration,omitempty" xml:"AudioCallDuration,omitempty"`
	// example:
	//
	// 0.0984
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	// example:
	//
	// 1.0.0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// macOS
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 476
	TotalCallDuration *int64 `json:"TotalCallDuration,omitempty" xml:"TotalCallDuration,omitempty"`
	// example:
	//
	// 473
	VideoCallDuration *int64 `json:"VideoCallDuration,omitempty" xml:"VideoCallDuration,omitempty"`
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) GetAudioCallDuration() *int64 {
	return s.AudioCallDuration
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) GetCallDurationRatio() *string {
	return s.CallDurationRatio
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) GetName() *string {
	return s.Name
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) GetOs() *string {
	return s.Os
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) GetTotalCallDuration() *int64 {
	return s.TotalCallDuration
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) GetVideoCallDuration() *int64 {
	return s.VideoCallDuration
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetAudioCallDuration(v int64) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.AudioCallDuration = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetCallDurationRatio(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetName(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.Name = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetOs(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.Os = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetTotalCallDuration(v int64) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.TotalCallDuration = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetVideoCallDuration(v int64) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.VideoCallDuration = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) Validate() error {
	return dara.Validate(s)
}
