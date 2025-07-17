// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstancesProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *SetInstancesProtectionRequest
	GetInstanceIds() []*string
	SetOwnerId(v int64) *SetInstancesProtectionRequest
	GetOwnerId() *int64
	SetProtectedFromScaleIn(v bool) *SetInstancesProtectionRequest
	GetProtectedFromScaleIn() *bool
	SetResourceOwnerAccount(v string) *SetInstancesProtectionRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *SetInstancesProtectionRequest
	GetScalingGroupId() *string
}

type SetInstancesProtectionRequest struct {
	// The IDs of the ECS instances.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	OwnerId     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to protect ECS instances from being stopped or removed from the scaling group during scale-ins. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ProtectedFromScaleIn *bool   `json:"ProtectedFromScaleIn,omitempty" xml:"ProtectedFromScaleIn,omitempty"`
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

func (s SetInstancesProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetInstancesProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetInstancesProtectionRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *SetInstancesProtectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetInstancesProtectionRequest) GetProtectedFromScaleIn() *bool {
	return s.ProtectedFromScaleIn
}

func (s *SetInstancesProtectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetInstancesProtectionRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *SetInstancesProtectionRequest) SetInstanceIds(v []*string) *SetInstancesProtectionRequest {
	s.InstanceIds = v
	return s
}

func (s *SetInstancesProtectionRequest) SetOwnerId(v int64) *SetInstancesProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *SetInstancesProtectionRequest) SetProtectedFromScaleIn(v bool) *SetInstancesProtectionRequest {
	s.ProtectedFromScaleIn = &v
	return s
}

func (s *SetInstancesProtectionRequest) SetResourceOwnerAccount(v string) *SetInstancesProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetInstancesProtectionRequest) SetScalingGroupId(v string) *SetInstancesProtectionRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *SetInstancesProtectionRequest) Validate() error {
	return dara.Validate(s)
}
