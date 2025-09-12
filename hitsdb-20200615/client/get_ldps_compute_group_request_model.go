// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLdpsComputeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *GetLdpsComputeGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *GetLdpsComputeGroupRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLdpsComputeGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLdpsComputeGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetLdpsComputeGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetLdpsComputeGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLdpsComputeGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLdpsComputeGroupRequest
	GetSecurityToken() *string
}

type GetLdpsComputeGroupRequest struct {
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

func (s GetLdpsComputeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLdpsComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *GetLdpsComputeGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetLdpsComputeGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLdpsComputeGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLdpsComputeGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLdpsComputeGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLdpsComputeGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLdpsComputeGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLdpsComputeGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLdpsComputeGroupRequest) SetGroupName(v string) *GetLdpsComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetInstanceId(v string) *GetLdpsComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetOwnerAccount(v string) *GetLdpsComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetOwnerId(v int64) *GetLdpsComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetRegionId(v string) *GetLdpsComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetResourceOwnerAccount(v string) *GetLdpsComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetResourceOwnerId(v int64) *GetLdpsComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetSecurityToken(v string) *GetLdpsComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) Validate() error {
	return dara.Validate(s)
}
