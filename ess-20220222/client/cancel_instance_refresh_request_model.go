// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelInstanceRefreshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRefreshTaskId(v string) *CancelInstanceRefreshRequest
	GetInstanceRefreshTaskId() *string
	SetOwnerId(v int64) *CancelInstanceRefreshRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelInstanceRefreshRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelInstanceRefreshRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *CancelInstanceRefreshRequest
	GetScalingGroupId() *string
}

type CancelInstanceRefreshRequest struct {
	// The ID of the instance refresh task.
	//
	// This parameter is required.
	//
	// example:
	//
	// ir-aca123sf****
	InstanceRefreshTaskId *string `json:"InstanceRefreshTaskId,omitempty" xml:"InstanceRefreshTaskId,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s CancelInstanceRefreshRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelInstanceRefreshRequest) GoString() string {
	return s.String()
}

func (s *CancelInstanceRefreshRequest) GetInstanceRefreshTaskId() *string {
	return s.InstanceRefreshTaskId
}

func (s *CancelInstanceRefreshRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelInstanceRefreshRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelInstanceRefreshRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelInstanceRefreshRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CancelInstanceRefreshRequest) SetInstanceRefreshTaskId(v string) *CancelInstanceRefreshRequest {
	s.InstanceRefreshTaskId = &v
	return s
}

func (s *CancelInstanceRefreshRequest) SetOwnerId(v int64) *CancelInstanceRefreshRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelInstanceRefreshRequest) SetRegionId(v string) *CancelInstanceRefreshRequest {
	s.RegionId = &v
	return s
}

func (s *CancelInstanceRefreshRequest) SetResourceOwnerAccount(v string) *CancelInstanceRefreshRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelInstanceRefreshRequest) SetScalingGroupId(v string) *CancelInstanceRefreshRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CancelInstanceRefreshRequest) Validate() error {
	return dara.Validate(s)
}
