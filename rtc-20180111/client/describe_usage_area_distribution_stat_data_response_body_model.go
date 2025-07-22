// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageAreaDistributionStatDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUsageAreaDistributionStatDataResponseBody
	GetRequestId() *string
	SetUsageAreaStatList(v []*DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) *DescribeUsageAreaDistributionStatDataResponseBody
	GetUsageAreaStatList() []*DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList
}

type DescribeUsageAreaDistributionStatDataResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId         *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageAreaStatList []*DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList `json:"UsageAreaStatList,omitempty" xml:"UsageAreaStatList,omitempty" type:"Repeated"`
}

func (s DescribeUsageAreaDistributionStatDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUsageAreaDistributionStatDataResponseBody) GetUsageAreaStatList() []*DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	return s.UsageAreaStatList
}

func (s *DescribeUsageAreaDistributionStatDataResponseBody) SetRequestId(v string) *DescribeUsageAreaDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBody) SetUsageAreaStatList(v []*DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) *DescribeUsageAreaDistributionStatDataResponseBody {
	s.UsageAreaStatList = v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList struct {
	// example:
	//
	// 45
	AudioCallDuration *int32 `json:"AudioCallDuration,omitempty" xml:"AudioCallDuration,omitempty"`
	// example:
	//
	// 中国
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4821
	TotalCallDuration *int32 `json:"TotalCallDuration,omitempty" xml:"TotalCallDuration,omitempty"`
	// example:
	//
	// 4776
	VideoCallDuration *int32 `json:"VideoCallDuration,omitempty" xml:"VideoCallDuration,omitempty"`
}

func (s DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) GetAudioCallDuration() *int32 {
	return s.AudioCallDuration
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) GetName() *string {
	return s.Name
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) GetTotalCallDuration() *int32 {
	return s.TotalCallDuration
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) GetVideoCallDuration() *int32 {
	return s.VideoCallDuration
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetAudioCallDuration(v int32) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.AudioCallDuration = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetName(v string) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.Name = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetTotalCallDuration(v int32) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.TotalCallDuration = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetVideoCallDuration(v int32) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.VideoCallDuration = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) Validate() error {
	return dara.Validate(s)
}
