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
	// Specifies whether the lifecycle (epoch) of an instance manually added to the scaling group is managed by the scaling group. Valid values:
	//
	// - true: Managed. The lifecycle (epoch) of the instance is managed by Auto Scaling, which is consistent with that of instances automatically created in the scaling group. The instance is automatically released when it is removed from the scaling group, excluding instances removed when you invoke the DetachInstances operation.
	//
	// - false: Unmanaged. The instance is not released when it is removed from the scaling group.
	//
	// > This parameter applies only to instances manually added to the scaling group, not to instances created within the scaling group.
	//
	// example:
	//
	// true
	Entrusted *bool `json:"Entrusted,omitempty" xml:"Entrusted,omitempty"`
	// The ID of the ECS instance in the scaling group.
	//
	// example:
	//
	// i-bp109k5j3dum1ce6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	OwnerId     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the scaling group belongs.
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
