// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoProvisioningGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoProvisioningGroupId(v string) *DeleteAutoProvisioningGroupRequest
	GetAutoProvisioningGroupId() *string
	SetOwnerAccount(v string) *DeleteAutoProvisioningGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteAutoProvisioningGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteAutoProvisioningGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteAutoProvisioningGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAutoProvisioningGroupRequest
	GetResourceOwnerId() *int64
	SetTerminateInstances(v bool) *DeleteAutoProvisioningGroupRequest
	GetTerminateInstances() *bool
}

type DeleteAutoProvisioningGroupRequest struct {
	// The ID of the auto provisioning group.
	//
	// This parameter is required.
	//
	// example:
	//
	// apg-bpuf6jel2bbl62wh13****
	AutoProvisioningGroupId *string `json:"AutoProvisioningGroupId,omitempty" xml:"AutoProvisioningGroupId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the auto provisioning group resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to release the instances in the auto-provisioning group when the group is deleted. Valid values:
	//
	// - true: Releases the instances auto provisioning group.
	//
	// - false: The instances auto provisioning group continue to run.
	//
	// >The default value of this parameter is inherited from the TerminateInstances parameter that you specified when you called the CreateAutoProvisioningGroup operation to create the auto-provisioning group. You can also set the TerminateInstances parameter to a new value when you call this operation to delete the auto-provisioning group.
	//
	// example:
	//
	// true
	TerminateInstances *bool `json:"TerminateInstances,omitempty" xml:"TerminateInstances,omitempty"`
}

func (s DeleteAutoProvisioningGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoProvisioningGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoProvisioningGroupRequest) GetAutoProvisioningGroupId() *string {
	return s.AutoProvisioningGroupId
}

func (s *DeleteAutoProvisioningGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteAutoProvisioningGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAutoProvisioningGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAutoProvisioningGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAutoProvisioningGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAutoProvisioningGroupRequest) GetTerminateInstances() *bool {
	return s.TerminateInstances
}

func (s *DeleteAutoProvisioningGroupRequest) SetAutoProvisioningGroupId(v string) *DeleteAutoProvisioningGroupRequest {
	s.AutoProvisioningGroupId = &v
	return s
}

func (s *DeleteAutoProvisioningGroupRequest) SetOwnerAccount(v string) *DeleteAutoProvisioningGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAutoProvisioningGroupRequest) SetOwnerId(v int64) *DeleteAutoProvisioningGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAutoProvisioningGroupRequest) SetRegionId(v string) *DeleteAutoProvisioningGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAutoProvisioningGroupRequest) SetResourceOwnerAccount(v string) *DeleteAutoProvisioningGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAutoProvisioningGroupRequest) SetResourceOwnerId(v int64) *DeleteAutoProvisioningGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAutoProvisioningGroupRequest) SetTerminateInstances(v bool) *DeleteAutoProvisioningGroupRequest {
	s.TerminateInstances = &v
	return s
}

func (s *DeleteAutoProvisioningGroupRequest) Validate() error {
	return dara.Validate(s)
}
