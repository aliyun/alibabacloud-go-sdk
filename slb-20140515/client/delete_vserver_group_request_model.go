// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteVServerGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVServerGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVServerGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVServerGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVServerGroupRequest
	GetResourceOwnerId() *int64
	SetVServerGroupId(v string) *DeleteVServerGroupRequest
	GetVServerGroupId() *string
}

type DeleteVServerGroupRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the associated Server Load Balancer (SLB) instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VServer group to be deleted.
	//
	// >  If the VServer group is in use, it cannot be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// rsp-cige6j*****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DeleteVServerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVServerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteVServerGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVServerGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVServerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVServerGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVServerGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVServerGroupRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DeleteVServerGroupRequest) SetOwnerAccount(v string) *DeleteVServerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVServerGroupRequest) SetOwnerId(v int64) *DeleteVServerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVServerGroupRequest) SetRegionId(v string) *DeleteVServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVServerGroupRequest) SetResourceOwnerAccount(v string) *DeleteVServerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVServerGroupRequest) SetResourceOwnerId(v int64) *DeleteVServerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVServerGroupRequest) SetVServerGroupId(v string) *DeleteVServerGroupRequest {
	s.VServerGroupId = &v
	return s
}

func (s *DeleteVServerGroupRequest) Validate() error {
	return dara.Validate(s)
}
