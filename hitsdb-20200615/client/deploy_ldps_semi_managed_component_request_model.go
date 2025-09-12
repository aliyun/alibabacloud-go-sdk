// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployLdpsSemiManagedComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentName(v string) *DeployLdpsSemiManagedComponentRequest
	GetComponentName() *string
	SetInstanceId(v string) *DeployLdpsSemiManagedComponentRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeployLdpsSemiManagedComponentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeployLdpsSemiManagedComponentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeployLdpsSemiManagedComponentRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeployLdpsSemiManagedComponentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeployLdpsSemiManagedComponentRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DeployLdpsSemiManagedComponentRequest
	GetSecurityToken() *string
}

type DeployLdpsSemiManagedComponentRequest struct {
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeployLdpsSemiManagedComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployLdpsSemiManagedComponentRequest) GoString() string {
	return s.String()
}

func (s *DeployLdpsSemiManagedComponentRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *DeployLdpsSemiManagedComponentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeployLdpsSemiManagedComponentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeployLdpsSemiManagedComponentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeployLdpsSemiManagedComponentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeployLdpsSemiManagedComponentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeployLdpsSemiManagedComponentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeployLdpsSemiManagedComponentRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeployLdpsSemiManagedComponentRequest) SetComponentName(v string) *DeployLdpsSemiManagedComponentRequest {
	s.ComponentName = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetInstanceId(v string) *DeployLdpsSemiManagedComponentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetOwnerAccount(v string) *DeployLdpsSemiManagedComponentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetOwnerId(v int64) *DeployLdpsSemiManagedComponentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetRegionId(v string) *DeployLdpsSemiManagedComponentRequest {
	s.RegionId = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetResourceOwnerAccount(v string) *DeployLdpsSemiManagedComponentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetResourceOwnerId(v int64) *DeployLdpsSemiManagedComponentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetSecurityToken(v string) *DeployLdpsSemiManagedComponentRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) Validate() error {
	return dara.Validate(s)
}
