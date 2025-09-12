// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLdpsComputeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *CreateLdpsComputeGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *CreateLdpsComputeGroupRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *CreateLdpsComputeGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLdpsComputeGroupRequest
	GetOwnerId() *int64
	SetProperties(v string) *CreateLdpsComputeGroupRequest
	GetProperties() *string
	SetRegionId(v string) *CreateLdpsComputeGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateLdpsComputeGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLdpsComputeGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateLdpsComputeGroupRequest
	GetSecurityToken() *string
}

type CreateLdpsComputeGroupRequest struct {
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

func (s CreateLdpsComputeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLdpsComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateLdpsComputeGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateLdpsComputeGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLdpsComputeGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLdpsComputeGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLdpsComputeGroupRequest) GetProperties() *string {
	return s.Properties
}

func (s *CreateLdpsComputeGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLdpsComputeGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLdpsComputeGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLdpsComputeGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateLdpsComputeGroupRequest) SetGroupName(v string) *CreateLdpsComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetInstanceId(v string) *CreateLdpsComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetOwnerAccount(v string) *CreateLdpsComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetOwnerId(v int64) *CreateLdpsComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetProperties(v string) *CreateLdpsComputeGroupRequest {
	s.Properties = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetRegionId(v string) *CreateLdpsComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetResourceOwnerAccount(v string) *CreateLdpsComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetResourceOwnerId(v int64) *CreateLdpsComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetSecurityToken(v string) *CreateLdpsComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) Validate() error {
	return dara.Validate(s)
}
