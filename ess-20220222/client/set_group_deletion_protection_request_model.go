// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGroupDeletionProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupDeletionProtection(v bool) *SetGroupDeletionProtectionRequest
	GetGroupDeletionProtection() *bool
	SetOwnerId(v int64) *SetGroupDeletionProtectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetGroupDeletionProtectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetGroupDeletionProtectionRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *SetGroupDeletionProtectionRequest
	GetScalingGroupId() *string
}

type SetGroupDeletionProtectionRequest struct {
	// Specifies whether to enable deletion protection for the scaling group. Valid values:
	//
	// 	- true: enables deletion protection. In this case, you cannot delete the scaling group by using the Auto Scaling console or calling an API operation. You must disable deletion protection before you can delete the scaling group.
	//
	// 	- false: disables deletion protection.
	//
	// Default value: false.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	GroupDeletionProtection *bool  `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	OwnerId                 *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-qingdao
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

func (s SetGroupDeletionProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetGroupDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetGroupDeletionProtectionRequest) GetGroupDeletionProtection() *bool {
	return s.GroupDeletionProtection
}

func (s *SetGroupDeletionProtectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetGroupDeletionProtectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetGroupDeletionProtectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetGroupDeletionProtectionRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *SetGroupDeletionProtectionRequest) SetGroupDeletionProtection(v bool) *SetGroupDeletionProtectionRequest {
	s.GroupDeletionProtection = &v
	return s
}

func (s *SetGroupDeletionProtectionRequest) SetOwnerId(v int64) *SetGroupDeletionProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *SetGroupDeletionProtectionRequest) SetRegionId(v string) *SetGroupDeletionProtectionRequest {
	s.RegionId = &v
	return s
}

func (s *SetGroupDeletionProtectionRequest) SetResourceOwnerAccount(v string) *SetGroupDeletionProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetGroupDeletionProtectionRequest) SetScalingGroupId(v string) *SetGroupDeletionProtectionRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *SetGroupDeletionProtectionRequest) Validate() error {
	return dara.Validate(s)
}
