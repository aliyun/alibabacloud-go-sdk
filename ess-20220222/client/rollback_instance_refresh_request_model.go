// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackInstanceRefreshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRefreshTaskId(v string) *RollbackInstanceRefreshRequest
	GetInstanceRefreshTaskId() *string
	SetOwnerId(v int64) *RollbackInstanceRefreshRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RollbackInstanceRefreshRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RollbackInstanceRefreshRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *RollbackInstanceRefreshRequest
	GetScalingGroupId() *string
}

type RollbackInstanceRefreshRequest struct {
	// The ID of the instance refresh task.
	//
	// This parameter is required.
	//
	// example:
	//
	// ir-a12ds234fasd*****
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
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s RollbackInstanceRefreshRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackInstanceRefreshRequest) GoString() string {
	return s.String()
}

func (s *RollbackInstanceRefreshRequest) GetInstanceRefreshTaskId() *string {
	return s.InstanceRefreshTaskId
}

func (s *RollbackInstanceRefreshRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RollbackInstanceRefreshRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RollbackInstanceRefreshRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RollbackInstanceRefreshRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *RollbackInstanceRefreshRequest) SetInstanceRefreshTaskId(v string) *RollbackInstanceRefreshRequest {
	s.InstanceRefreshTaskId = &v
	return s
}

func (s *RollbackInstanceRefreshRequest) SetOwnerId(v int64) *RollbackInstanceRefreshRequest {
	s.OwnerId = &v
	return s
}

func (s *RollbackInstanceRefreshRequest) SetRegionId(v string) *RollbackInstanceRefreshRequest {
	s.RegionId = &v
	return s
}

func (s *RollbackInstanceRefreshRequest) SetResourceOwnerAccount(v string) *RollbackInstanceRefreshRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RollbackInstanceRefreshRequest) SetScalingGroupId(v string) *RollbackInstanceRefreshRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *RollbackInstanceRefreshRequest) Validate() error {
	return dara.Validate(s)
}
