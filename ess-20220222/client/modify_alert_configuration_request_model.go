// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAlertConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ModifyAlertConfigurationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyAlertConfigurationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyAlertConfigurationRequest
	GetResourceOwnerAccount() *string
	SetScaleStatuses(v []*string) *ModifyAlertConfigurationRequest
	GetScaleStatuses() []*string
	SetScalingGroupId(v string) *ModifyAlertConfigurationRequest
	GetScalingGroupId() *string
}

type ModifyAlertConfigurationRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The status of the scaling activities that prompt text message or email notifications.
	ScaleStatuses []*string `json:"ScaleStatuses,omitempty" xml:"ScaleStatuses,omitempty" type:"Repeated"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1eyv4qn8ssgv43****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s ModifyAlertConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAlertConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyAlertConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAlertConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAlertConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAlertConfigurationRequest) GetScaleStatuses() []*string {
	return s.ScaleStatuses
}

func (s *ModifyAlertConfigurationRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ModifyAlertConfigurationRequest) SetOwnerId(v int64) *ModifyAlertConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAlertConfigurationRequest) SetRegionId(v string) *ModifyAlertConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAlertConfigurationRequest) SetResourceOwnerAccount(v string) *ModifyAlertConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAlertConfigurationRequest) SetScaleStatuses(v []*string) *ModifyAlertConfigurationRequest {
	s.ScaleStatuses = v
	return s
}

func (s *ModifyAlertConfigurationRequest) SetScalingGroupId(v string) *ModifyAlertConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ModifyAlertConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
