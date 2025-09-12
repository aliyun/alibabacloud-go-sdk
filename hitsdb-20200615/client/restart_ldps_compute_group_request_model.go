// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartLdpsComputeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *RestartLdpsComputeGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *RestartLdpsComputeGroupRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RestartLdpsComputeGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartLdpsComputeGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RestartLdpsComputeGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RestartLdpsComputeGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartLdpsComputeGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *RestartLdpsComputeGroupRequest
	GetSecurityToken() *string
}

type RestartLdpsComputeGroupRequest struct {
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

func (s RestartLdpsComputeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartLdpsComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *RestartLdpsComputeGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *RestartLdpsComputeGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartLdpsComputeGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartLdpsComputeGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartLdpsComputeGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartLdpsComputeGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartLdpsComputeGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartLdpsComputeGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RestartLdpsComputeGroupRequest) SetGroupName(v string) *RestartLdpsComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetInstanceId(v string) *RestartLdpsComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetOwnerAccount(v string) *RestartLdpsComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetOwnerId(v int64) *RestartLdpsComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetRegionId(v string) *RestartLdpsComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetResourceOwnerAccount(v string) *RestartLdpsComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetResourceOwnerId(v int64) *RestartLdpsComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetSecurityToken(v string) *RestartLdpsComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) Validate() error {
	return dara.Validate(s)
}
