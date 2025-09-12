// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultOlapComputeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *SetDefaultOlapComputeGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *SetDefaultOlapComputeGroupRequest
	GetInstanceId() *string
	SetIsDefault(v bool) *SetDefaultOlapComputeGroupRequest
	GetIsDefault() *bool
	SetOwnerAccount(v string) *SetDefaultOlapComputeGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetDefaultOlapComputeGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetDefaultOlapComputeGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetDefaultOlapComputeGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetDefaultOlapComputeGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *SetDefaultOlapComputeGroupRequest
	GetSecurityToken() *string
}

type SetDefaultOlapComputeGroupRequest struct {
	// This parameter is required.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsDefault            *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetDefaultOlapComputeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultOlapComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultOlapComputeGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *SetDefaultOlapComputeGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetDefaultOlapComputeGroupRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *SetDefaultOlapComputeGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetDefaultOlapComputeGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetDefaultOlapComputeGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDefaultOlapComputeGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetDefaultOlapComputeGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetDefaultOlapComputeGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetDefaultOlapComputeGroupRequest) SetGroupName(v string) *SetDefaultOlapComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetInstanceId(v string) *SetDefaultOlapComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetIsDefault(v bool) *SetDefaultOlapComputeGroupRequest {
	s.IsDefault = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetOwnerAccount(v string) *SetDefaultOlapComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetOwnerId(v int64) *SetDefaultOlapComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetRegionId(v string) *SetDefaultOlapComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetResourceOwnerAccount(v string) *SetDefaultOlapComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetResourceOwnerId(v int64) *SetDefaultOlapComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetSecurityToken(v string) *SetDefaultOlapComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) Validate() error {
	return dara.Validate(s)
}
