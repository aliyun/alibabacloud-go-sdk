// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceStatisticByRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeResourceStatisticByRegionResponseBody
	GetRequestId() *string
	SetStatistic(v *DescribeResourceStatisticByRegionResponseBodyStatistic) *DescribeResourceStatisticByRegionResponseBody
	GetStatistic() *DescribeResourceStatisticByRegionResponseBodyStatistic
}

type DescribeResourceStatisticByRegionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics information.
	Statistic *DescribeResourceStatisticByRegionResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Struct"`
}

func (s DescribeResourceStatisticByRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceStatisticByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticByRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceStatisticByRegionResponseBody) GetStatistic() *DescribeResourceStatisticByRegionResponseBodyStatistic {
	return s.Statistic
}

func (s *DescribeResourceStatisticByRegionResponseBody) SetRequestId(v string) *DescribeResourceStatisticByRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceStatisticByRegionResponseBody) SetStatistic(v *DescribeResourceStatisticByRegionResponseBodyStatistic) *DescribeResourceStatisticByRegionResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeResourceStatisticByRegionResponseBody) Validate() error {
	if s.Statistic != nil {
		if err := s.Statistic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourceStatisticByRegionResponseBodyStatistic struct {
	// The statistics by region.
	RegionStatistics []*DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics `json:"RegionStatistics,omitempty" xml:"RegionStatistics,omitempty" type:"Repeated"`
}

func (s DescribeResourceStatisticByRegionResponseBodyStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceStatisticByRegionResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatistic) GetRegionStatistics() []*DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics {
	return s.RegionStatistics
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatistic) SetRegionStatistics(v []*DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) *DescribeResourceStatisticByRegionResponseBodyStatistic {
	s.RegionStatistics = v
	return s
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatistic) Validate() error {
	if s.RegionStatistics != nil {
		for _, item := range s.RegionStatistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics struct {
	// The number of assigned users in the current region.
	//
	// example:
	//
	// 0
	AssignedEndUserCount *int64 `json:"AssignedEndUserCount,omitempty" xml:"AssignedEndUserCount,omitempty"`
	// The number of connected users in the current region.
	//
	// example:
	//
	// 0
	ConnectedEndUserCount *int64 `json:"ConnectedEndUserCount,omitempty" xml:"ConnectedEndUserCount,omitempty"`
	// The number of cloud desktops in the current region.
	//
	// example:
	//
	// 0
	DesktopCount *int64 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The number of shared cloud desktops in the current region.
	//
	// example:
	//
	// 0
	DesktopGroupCount *int64 `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	// The number of network packages in the current region.
	//
	// example:
	//
	// 0
	NetworkPackageCount *int64 `json:"NetworkPackageCount,omitempty" xml:"NetworkPackageCount,omitempty"`
	// The number of policy groups in the current region.
	//
	// example:
	//
	// 0
	PolicyGroupCount *int64 `json:"PolicyGroupCount,omitempty" xml:"PolicyGroupCount,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://next.api.aliyun.com/document/ecd/2020-09-30/DescribeRegions) operation to query the regions supported by WUYING Workspace.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) GetAssignedEndUserCount() *int64 {
	return s.AssignedEndUserCount
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) GetConnectedEndUserCount() *int64 {
	return s.ConnectedEndUserCount
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) GetDesktopCount() *int64 {
	return s.DesktopCount
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) GetDesktopGroupCount() *int64 {
	return s.DesktopGroupCount
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) GetNetworkPackageCount() *int64 {
	return s.NetworkPackageCount
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) GetPolicyGroupCount() *int64 {
	return s.PolicyGroupCount
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) SetAssignedEndUserCount(v int64) *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics {
	s.AssignedEndUserCount = &v
	return s
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) SetConnectedEndUserCount(v int64) *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics {
	s.ConnectedEndUserCount = &v
	return s
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) SetDesktopCount(v int64) *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics {
	s.DesktopCount = &v
	return s
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) SetDesktopGroupCount(v int64) *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics {
	s.DesktopGroupCount = &v
	return s
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) SetNetworkPackageCount(v int64) *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics {
	s.NetworkPackageCount = &v
	return s
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) SetPolicyGroupCount(v int64) *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics {
	s.PolicyGroupCount = &v
	return s
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) SetRegionId(v string) *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceStatisticByRegionResponseBodyStatisticRegionStatistics) Validate() error {
	return dara.Validate(s)
}
