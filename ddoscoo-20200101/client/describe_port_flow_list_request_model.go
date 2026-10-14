// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortFlowListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribePortFlowListRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribePortFlowListRequest
	GetInstanceIds() []*string
	SetInterval(v int32) *DescribePortFlowListRequest
	GetInterval() *int32
	SetResourceGroupId(v string) *DescribePortFlowListRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribePortFlowListRequest
	GetStartTime() *int64
}

type DescribePortFlowListRequest struct {
	// The end time of the query. Specify the value as a UNIX timestamp. Unit: seconds.
	//
	// > The value must be on the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Anti-DDoS Pro and Anti-DDoS Premium instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The interval at which data is returned. Unit: seconds. The supported interval values vary based on the query time range (determined by **StartTime*	- and **EndTime**). Recommended settings:
	//
	// - Query time range ≤ 1 hour: minimum interval of 60 seconds, maximum not exceeding the query time range.
	//
	// - Query time range > 1 hour and ≤ 6 hours: minimum interval of 600 seconds, maximum not exceeding the query time range.
	//
	// - Query time range > 6 hours and ≤ 24 hours: minimum interval of 1800 seconds, maximum not exceeding the query time range.
	//
	// - Query time range > 24 hours and ≤ 7 days: minimum interval of 3600 seconds, maximum not exceeding the query time range.
	//
	// - Query time range > 7 days and ≤ 15 days: minimum interval of 14400 seconds, maximum not exceeding the query time range.
	//
	// - Query time range > 15 days: minimum interval of 43200 seconds, maximum not exceeding the query time range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro and Anti-DDoS Premium instance belongs in Resource Management. By default, this parameter is empty, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time of the query. Specify the value as a UNIX timestamp. Unit: seconds.
	//
	// > The value must be on the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582992000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePortFlowListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortFlowListRequest) GoString() string {
	return s.String()
}

func (s *DescribePortFlowListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePortFlowListRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribePortFlowListRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribePortFlowListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePortFlowListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePortFlowListRequest) SetEndTime(v int64) *DescribePortFlowListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortFlowListRequest) SetInstanceIds(v []*string) *DescribePortFlowListRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePortFlowListRequest) SetInterval(v int32) *DescribePortFlowListRequest {
	s.Interval = &v
	return s
}

func (s *DescribePortFlowListRequest) SetResourceGroupId(v string) *DescribePortFlowListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortFlowListRequest) SetStartTime(v int64) *DescribePortFlowListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortFlowListRequest) Validate() error {
	return dara.Validate(s)
}
