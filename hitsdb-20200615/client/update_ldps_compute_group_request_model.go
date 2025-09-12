// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLdpsComputeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *UpdateLdpsComputeGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *UpdateLdpsComputeGroupRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *UpdateLdpsComputeGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateLdpsComputeGroupRequest
	GetOwnerId() *int64
	SetProperties(v string) *UpdateLdpsComputeGroupRequest
	GetProperties() *string
	SetRegionId(v string) *UpdateLdpsComputeGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateLdpsComputeGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateLdpsComputeGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *UpdateLdpsComputeGroupRequest
	GetSecurityToken() *string
}

type UpdateLdpsComputeGroupRequest struct {
	// This parameter is required.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Properties           *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateLdpsComputeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLdpsComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateLdpsComputeGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateLdpsComputeGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLdpsComputeGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateLdpsComputeGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLdpsComputeGroupRequest) GetProperties() *string {
	return s.Properties
}

func (s *UpdateLdpsComputeGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLdpsComputeGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateLdpsComputeGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateLdpsComputeGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateLdpsComputeGroupRequest) SetGroupName(v string) *UpdateLdpsComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetInstanceId(v string) *UpdateLdpsComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetOwnerAccount(v string) *UpdateLdpsComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetOwnerId(v int64) *UpdateLdpsComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetProperties(v string) *UpdateLdpsComputeGroupRequest {
	s.Properties = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetRegionId(v string) *UpdateLdpsComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetResourceOwnerAccount(v string) *UpdateLdpsComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetResourceOwnerId(v int64) *UpdateLdpsComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetSecurityToken(v string) *UpdateLdpsComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) Validate() error {
	return dara.Validate(s)
}
