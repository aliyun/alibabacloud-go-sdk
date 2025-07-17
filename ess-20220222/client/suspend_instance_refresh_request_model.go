// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendInstanceRefreshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRefreshTaskId(v string) *SuspendInstanceRefreshRequest
	GetInstanceRefreshTaskId() *string
	SetOwnerId(v int64) *SuspendInstanceRefreshRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SuspendInstanceRefreshRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SuspendInstanceRefreshRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *SuspendInstanceRefreshRequest
	GetScalingGroupId() *string
}

type SuspendInstanceRefreshRequest struct {
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

func (s SuspendInstanceRefreshRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendInstanceRefreshRequest) GoString() string {
	return s.String()
}

func (s *SuspendInstanceRefreshRequest) GetInstanceRefreshTaskId() *string {
	return s.InstanceRefreshTaskId
}

func (s *SuspendInstanceRefreshRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SuspendInstanceRefreshRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SuspendInstanceRefreshRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SuspendInstanceRefreshRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *SuspendInstanceRefreshRequest) SetInstanceRefreshTaskId(v string) *SuspendInstanceRefreshRequest {
	s.InstanceRefreshTaskId = &v
	return s
}

func (s *SuspendInstanceRefreshRequest) SetOwnerId(v int64) *SuspendInstanceRefreshRequest {
	s.OwnerId = &v
	return s
}

func (s *SuspendInstanceRefreshRequest) SetRegionId(v string) *SuspendInstanceRefreshRequest {
	s.RegionId = &v
	return s
}

func (s *SuspendInstanceRefreshRequest) SetResourceOwnerAccount(v string) *SuspendInstanceRefreshRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SuspendInstanceRefreshRequest) SetScalingGroupId(v string) *SuspendInstanceRefreshRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *SuspendInstanceRefreshRequest) Validate() error {
	return dara.Validate(s)
}
