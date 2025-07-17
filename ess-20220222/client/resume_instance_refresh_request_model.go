// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeInstanceRefreshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRefreshTaskId(v string) *ResumeInstanceRefreshRequest
	GetInstanceRefreshTaskId() *string
	SetOwnerId(v int64) *ResumeInstanceRefreshRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ResumeInstanceRefreshRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ResumeInstanceRefreshRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *ResumeInstanceRefreshRequest
	GetScalingGroupId() *string
}

type ResumeInstanceRefreshRequest struct {
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

func (s ResumeInstanceRefreshRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeInstanceRefreshRequest) GoString() string {
	return s.String()
}

func (s *ResumeInstanceRefreshRequest) GetInstanceRefreshTaskId() *string {
	return s.InstanceRefreshTaskId
}

func (s *ResumeInstanceRefreshRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResumeInstanceRefreshRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResumeInstanceRefreshRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResumeInstanceRefreshRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ResumeInstanceRefreshRequest) SetInstanceRefreshTaskId(v string) *ResumeInstanceRefreshRequest {
	s.InstanceRefreshTaskId = &v
	return s
}

func (s *ResumeInstanceRefreshRequest) SetOwnerId(v int64) *ResumeInstanceRefreshRequest {
	s.OwnerId = &v
	return s
}

func (s *ResumeInstanceRefreshRequest) SetRegionId(v string) *ResumeInstanceRefreshRequest {
	s.RegionId = &v
	return s
}

func (s *ResumeInstanceRefreshRequest) SetResourceOwnerAccount(v string) *ResumeInstanceRefreshRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResumeInstanceRefreshRequest) SetScalingGroupId(v string) *ResumeInstanceRefreshRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ResumeInstanceRefreshRequest) Validate() error {
	return dara.Validate(s)
}
