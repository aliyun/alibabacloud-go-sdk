// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateInstanceSecurityGroupsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *UpdateInstanceSecurityGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateInstanceSecurityGroupsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateInstanceSecurityGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateInstanceSecurityGroupsRequest
	GetResourceOwnerId() *int64
	SetSecurityGroups(v string) *UpdateInstanceSecurityGroupsRequest
	GetSecurityGroups() *string
	SetSecurityToken(v string) *UpdateInstanceSecurityGroupsRequest
	GetSecurityToken() *string
}

type UpdateInstanceSecurityGroupsRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SecurityGroups *string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateInstanceSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSecurityGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceSecurityGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateInstanceSecurityGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateInstanceSecurityGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateInstanceSecurityGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateInstanceSecurityGroupsRequest) GetSecurityGroups() *string {
	return s.SecurityGroups
}

func (s *UpdateInstanceSecurityGroupsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateInstanceSecurityGroupsRequest) SetInstanceId(v string) *UpdateInstanceSecurityGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetOwnerAccount(v string) *UpdateInstanceSecurityGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetOwnerId(v int64) *UpdateInstanceSecurityGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetResourceOwnerAccount(v string) *UpdateInstanceSecurityGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetResourceOwnerId(v int64) *UpdateInstanceSecurityGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetSecurityGroups(v string) *UpdateInstanceSecurityGroupsRequest {
	s.SecurityGroups = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetSecurityToken(v string) *UpdateInstanceSecurityGroupsRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}
