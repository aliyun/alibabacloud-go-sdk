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
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// **
	//
	// **This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// An array that consists of the IDs of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The interval for returning data. Unit: seconds. The interval that you can specify varies based on the time range to query. The time range to query is determined by the values of **StartTime*	- and **EndTime**.
	//
	// 	- If the time range to query is no greater than 1 hour, we recommend that you specify the interval from 60 seconds to the time range to query.
	//
	// 	- If the time range to query is greater than 1 hour but no greater than 6 hours, we recommend that you specify the interval from 600 seconds to the time range to query.
	//
	// 	- If the time range to query is greater than 6 hours but no greater than 24 hours, we recommend that you specify the interval from 1,800 seconds to the time range to query.
	//
	// 	- If the time range to query is greater than 24 hours but no greater than 7 days, we recommend that you specify the interval from 3,600 seconds to the time range to query.
	//
	// 	- If the time range to query is greater than 7 days but no greater than 15 days, we recommend that you specify the interval from 14,400 seconds to the time range to query.
	//
	// 	- If the time range to query is greater than 15 days, we recommend that you specify the interval from 43,200 seconds to the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// **
	//
	// **This UNIX timestamp must indicate a point in time that is accurate to the minute.
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
