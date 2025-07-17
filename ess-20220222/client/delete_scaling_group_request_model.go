// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceDelete(v bool) *DeleteScalingGroupRequest
	GetForceDelete() *bool
	SetOwnerAccount(v string) *DeleteScalingGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteScalingGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteScalingGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteScalingGroupRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DeleteScalingGroupRequest
	GetScalingGroupId() *string
}

type DeleteScalingGroupRequest struct {
	// Specifies whether to enforce the deletion of the scaling group, including the removal of the current ECS instances or elastic container instances from the scaling group and their subsequent release, even if the scaling group is actively undergoing scaling activities. Valid values:
	//
	// 	- true: enforces the deletion of the scaling group. In this case, the scaling group first enters the Disabled state, ceasing acceptance of new scaling requests. Auto Scaling awaits the conclusion of all ongoing scaling activities in the scaling group before it automatically removes the current ECS instances or elastic container instances from the scaling group and enforces the deletion operation. Note that manually added instances are merely removed from the scaling group, whereas auto-provisioned instances are removed and deleted.
	//
	// 	- false: does not enforce the deletion of the scaling group. The scaling group will be disabled and then deleted once all the following requirements are satisfied:
	//
	//     	- The scaling group has no ongoing scaling activities.
	//
	//     	- The scaling group contains no ECS instances or elastic container instances (Total Capacity=0).
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceDelete  *bool   `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DeleteScalingGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingGroupRequest) GetForceDelete() *bool {
	return s.ForceDelete
}

func (s *DeleteScalingGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteScalingGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteScalingGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteScalingGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteScalingGroupRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DeleteScalingGroupRequest) SetForceDelete(v bool) *DeleteScalingGroupRequest {
	s.ForceDelete = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetOwnerAccount(v string) *DeleteScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetOwnerId(v int64) *DeleteScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetRegionId(v string) *DeleteScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetResourceOwnerAccount(v string) *DeleteScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetScalingGroupId(v string) *DeleteScalingGroupRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DeleteScalingGroupRequest) Validate() error {
	return dara.Validate(s)
}
