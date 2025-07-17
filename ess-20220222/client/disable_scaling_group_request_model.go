// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableScalingGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DisableScalingGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisableScalingGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DisableScalingGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisableScalingGroupRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *DisableScalingGroupRequest
	GetScalingGroupId() *string
}

type DisableScalingGroupRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DisableScalingGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *DisableScalingGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisableScalingGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableScalingGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableScalingGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableScalingGroupRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DisableScalingGroupRequest) SetOwnerAccount(v string) *DisableScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableScalingGroupRequest) SetOwnerId(v int64) *DisableScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableScalingGroupRequest) SetResourceOwnerAccount(v string) *DisableScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableScalingGroupRequest) SetResourceOwnerId(v int64) *DisableScalingGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableScalingGroupRequest) SetScalingGroupId(v string) *DisableScalingGroupRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DisableScalingGroupRequest) Validate() error {
	return dara.Validate(s)
}
