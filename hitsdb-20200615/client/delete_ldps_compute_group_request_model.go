// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLdpsComputeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DeleteLdpsComputeGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *DeleteLdpsComputeGroupRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeleteLdpsComputeGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteLdpsComputeGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLdpsComputeGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteLdpsComputeGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteLdpsComputeGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DeleteLdpsComputeGroupRequest
	GetSecurityToken() *string
}

type DeleteLdpsComputeGroupRequest struct {
	// This parameter is required.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteLdpsComputeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLdpsComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteLdpsComputeGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteLdpsComputeGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteLdpsComputeGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteLdpsComputeGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLdpsComputeGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLdpsComputeGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteLdpsComputeGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteLdpsComputeGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteLdpsComputeGroupRequest) SetGroupName(v string) *DeleteLdpsComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetInstanceId(v string) *DeleteLdpsComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetOwnerAccount(v string) *DeleteLdpsComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetOwnerId(v int64) *DeleteLdpsComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetRegionId(v string) *DeleteLdpsComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetResourceOwnerAccount(v string) *DeleteLdpsComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetResourceOwnerId(v int64) *DeleteLdpsComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetSecurityToken(v string) *DeleteLdpsComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) Validate() error {
	return dara.Validate(s)
}
