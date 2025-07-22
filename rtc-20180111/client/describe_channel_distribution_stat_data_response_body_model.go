// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelDistributionStatDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeChannelDistributionStatDataResponseBody
	GetRequestId() *string
	SetStatList(v []*DescribeChannelDistributionStatDataResponseBodyStatList) *DescribeChannelDistributionStatDataResponseBody
	GetStatList() []*DescribeChannelDistributionStatDataResponseBodyStatList
}

type DescribeChannelDistributionStatDataResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatList  []*DescribeChannelDistributionStatDataResponseBodyStatList `json:"StatList,omitempty" xml:"StatList,omitempty" type:"Repeated"`
}

func (s DescribeChannelDistributionStatDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelDistributionStatDataResponseBody) GetStatList() []*DescribeChannelDistributionStatDataResponseBodyStatList {
	return s.StatList
}

func (s *DescribeChannelDistributionStatDataResponseBody) SetRequestId(v string) *DescribeChannelDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponseBody) SetStatList(v []*DescribeChannelDistributionStatDataResponseBodyStatList) *DescribeChannelDistributionStatDataResponseBody {
	s.StatList = v
	return s
}

func (s *DescribeChannelDistributionStatDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeChannelDistributionStatDataResponseBodyStatList struct {
	// example:
	//
	// 1
	CallUserCount *int32 `json:"CallUserCount,omitempty" xml:"CallUserCount,omitempty"`
	// example:
	//
	// 1.0000
	CallUserRatio *string `json:"CallUserRatio,omitempty" xml:"CallUserRatio,omitempty"`
	// example:
	//
	// OS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeChannelDistributionStatDataResponseBodyStatList) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelDistributionStatDataResponseBodyStatList) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) GetCallUserCount() *int32 {
	return s.CallUserCount
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) GetCallUserRatio() *string {
	return s.CallUserRatio
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) GetName() *string {
	return s.Name
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) SetCallUserCount(v int32) *DescribeChannelDistributionStatDataResponseBodyStatList {
	s.CallUserCount = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) SetCallUserRatio(v string) *DescribeChannelDistributionStatDataResponseBodyStatList {
	s.CallUserRatio = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) SetName(v string) *DescribeChannelDistributionStatDataResponseBodyStatList {
	s.Name = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) Validate() error {
	return dara.Validate(s)
}
