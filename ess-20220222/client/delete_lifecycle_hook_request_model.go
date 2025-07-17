// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLifecycleHookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLifecycleHookId(v string) *DeleteLifecycleHookRequest
	GetLifecycleHookId() *string
	SetLifecycleHookName(v string) *DeleteLifecycleHookRequest
	GetLifecycleHookName() *string
	SetOwnerAccount(v string) *DeleteLifecycleHookRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteLifecycleHookRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLifecycleHookRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteLifecycleHookRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DeleteLifecycleHookRequest
	GetScalingGroupId() *string
}

type DeleteLifecycleHookRequest struct {
	// The ID of the lifecycle hook.
	//
	// example:
	//
	// ash-bp14g3ee6bt3sc98****
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	// The name of the lifecycle hook.
	//
	// example:
	//
	// lifecyclehook****
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DeleteLifecycleHookRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLifecycleHookRequest) GoString() string {
	return s.String()
}

func (s *DeleteLifecycleHookRequest) GetLifecycleHookId() *string {
	return s.LifecycleHookId
}

func (s *DeleteLifecycleHookRequest) GetLifecycleHookName() *string {
	return s.LifecycleHookName
}

func (s *DeleteLifecycleHookRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteLifecycleHookRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLifecycleHookRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLifecycleHookRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteLifecycleHookRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DeleteLifecycleHookRequest) SetLifecycleHookId(v string) *DeleteLifecycleHookRequest {
	s.LifecycleHookId = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetLifecycleHookName(v string) *DeleteLifecycleHookRequest {
	s.LifecycleHookName = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetOwnerAccount(v string) *DeleteLifecycleHookRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetOwnerId(v int64) *DeleteLifecycleHookRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetRegionId(v string) *DeleteLifecycleHookRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetResourceOwnerAccount(v string) *DeleteLifecycleHookRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetScalingGroupId(v string) *DeleteLifecycleHookRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DeleteLifecycleHookRequest) Validate() error {
	return dara.Validate(s)
}
