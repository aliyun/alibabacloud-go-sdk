// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTotalAttackMaxFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeTotalAttackMaxFlowRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribeTotalAttackMaxFlowRequest
	GetInstanceIds() []*string
	SetResourceGroupId(v string) *DescribeTotalAttackMaxFlowRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeTotalAttackMaxFlowRequest
	GetStartTime() *int64
}

type DescribeTotalAttackMaxFlowRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1659697200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IDs of the Anti-DDoS Proxy instances. Separate multiple instance IDs with commas (,). Example: InstanceIds.1, InstanceIds.2, InstanceIds.3.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the resource group to which the instance belongs in Resource Management. If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1669240800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTotalAttackMaxFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTotalAttackMaxFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeTotalAttackMaxFlowRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeTotalAttackMaxFlowRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeTotalAttackMaxFlowRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTotalAttackMaxFlowRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeTotalAttackMaxFlowRequest) SetEndTime(v int64) *DescribeTotalAttackMaxFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTotalAttackMaxFlowRequest) SetInstanceIds(v []*string) *DescribeTotalAttackMaxFlowRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeTotalAttackMaxFlowRequest) SetResourceGroupId(v string) *DescribeTotalAttackMaxFlowRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTotalAttackMaxFlowRequest) SetStartTime(v int64) *DescribeTotalAttackMaxFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTotalAttackMaxFlowRequest) Validate() error {
	return dara.Validate(s)
}
