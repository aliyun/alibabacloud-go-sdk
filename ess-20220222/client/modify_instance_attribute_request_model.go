// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntrusted(v bool) *ModifyInstanceAttributeRequest
	GetEntrusted() *bool
	SetInstanceId(v string) *ModifyInstanceAttributeRequest
	GetInstanceId() *string
	SetInstanceIds(v []*string) *ModifyInstanceAttributeRequest
	GetInstanceIds() []*string
	SetOwnerId(v int64) *ModifyInstanceAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *ModifyInstanceAttributeRequest
	GetScalingGroupId() *string
}

type ModifyInstanceAttributeRequest struct {
	// Specifies whether to allow the scaling group to manage the lifecycles of the manually added ECS instances. Valid values:
	//
	// 	- true: allows the scaling group to manage the lifecycles of the manually added ECS instances. The scaling group manages the lifecycles of manually added instances and automatically created instances in the same manner. In this case, Auto Scaling releases the instances when they are removed from the scaling group. This rule does not apply to instances that are removed by calling the DetachInstances operation.
	//
	// 	- false: does not allow the scaling group to manage the lifecycles of the manually added ECS instances. In this case, Auto Scaling does not release the instances when they are removed from the scaling group.
	//
	// >  You can specify this parameter only for ECS instances that are manually added to the scaling group.
	//
	// example:
	//
	// true
	Entrusted *bool `json:"Entrusted,omitempty" xml:"Entrusted,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-bp109k5j3dum1ce6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	OwnerId     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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

func (s ModifyInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) GetEntrusted() *bool {
	return s.Entrusted
}

func (s *ModifyInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAttributeRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceAttributeRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ModifyInstanceAttributeRequest) SetEntrusted(v bool) *ModifyInstanceAttributeRequest {
	s.Entrusted = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceIds(v []*string) *ModifyInstanceAttributeRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetOwnerId(v int64) *ModifyInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetRegionId(v string) *ModifyInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetScalingGroupId(v string) *ModifyInstanceAttributeRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
