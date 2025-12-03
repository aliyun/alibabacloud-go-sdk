// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMasterSlaveServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMasterSlaveServerGroupId(v string) *DeleteMasterSlaveServerGroupRequest
	GetMasterSlaveServerGroupId() *string
	SetOwnerAccount(v string) *DeleteMasterSlaveServerGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteMasterSlaveServerGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteMasterSlaveServerGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteMasterSlaveServerGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMasterSlaveServerGroupRequest
	GetResourceOwnerId() *int64
}

type DeleteMasterSlaveServerGroupRequest struct {
	// The primary/secondary server group ID.
	//
	// >  You cannot delete a primary/secondary server group that is in use.
	//
	// This parameter is required.
	//
	// example:
	//
	// rsp-cige6*****
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Server Load Balancer (SLB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteMasterSlaveServerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMasterSlaveServerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteMasterSlaveServerGroupRequest) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *DeleteMasterSlaveServerGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteMasterSlaveServerGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMasterSlaveServerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMasterSlaveServerGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMasterSlaveServerGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMasterSlaveServerGroupRequest) SetMasterSlaveServerGroupId(v string) *DeleteMasterSlaveServerGroupRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupRequest) SetOwnerAccount(v string) *DeleteMasterSlaveServerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupRequest) SetOwnerId(v int64) *DeleteMasterSlaveServerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupRequest) SetRegionId(v string) *DeleteMasterSlaveServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupRequest) SetResourceOwnerAccount(v string) *DeleteMasterSlaveServerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupRequest) SetResourceOwnerId(v int64) *DeleteMasterSlaveServerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupRequest) Validate() error {
	return dara.Validate(s)
}
