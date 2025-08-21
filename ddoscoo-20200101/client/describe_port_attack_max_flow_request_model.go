// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortAttackMaxFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribePortAttackMaxFlowRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribePortAttackMaxFlowRequest
	GetInstanceIds() []*string
	SetResourceGroupId(v string) *DescribePortAttackMaxFlowRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribePortAttackMaxFlowRequest
	GetStartTime() *int64
}

type DescribePortAttackMaxFlowRequest struct {
	// The end of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IDs of the Anti-DDoS Proxy instances to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582992000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePortAttackMaxFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortAttackMaxFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribePortAttackMaxFlowRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePortAttackMaxFlowRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribePortAttackMaxFlowRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePortAttackMaxFlowRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePortAttackMaxFlowRequest) SetEndTime(v int64) *DescribePortAttackMaxFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortAttackMaxFlowRequest) SetInstanceIds(v []*string) *DescribePortAttackMaxFlowRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePortAttackMaxFlowRequest) SetResourceGroupId(v string) *DescribePortAttackMaxFlowRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortAttackMaxFlowRequest) SetStartTime(v int64) *DescribePortAttackMaxFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortAttackMaxFlowRequest) Validate() error {
	return dara.Validate(s)
}
