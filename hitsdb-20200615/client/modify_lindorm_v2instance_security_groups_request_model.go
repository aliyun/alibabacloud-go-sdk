// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLindormV2InstanceSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyLindormV2InstanceSecurityGroupsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyLindormV2InstanceSecurityGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLindormV2InstanceSecurityGroupsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyLindormV2InstanceSecurityGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLindormV2InstanceSecurityGroupsRequest
	GetResourceOwnerId() *int64
	SetSecurityGroups(v string) *ModifyLindormV2InstanceSecurityGroupsRequest
	GetSecurityGroups() *string
	SetSecurityToken(v string) *ModifyLindormV2InstanceSecurityGroupsRequest
	GetSecurityToken() *string
}

type ModifyLindormV2InstanceSecurityGroupsRequest struct {
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

func (s ModifyLindormV2InstanceSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLindormV2InstanceSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) GetSecurityGroups() *string {
	return s.SecurityGroups
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) SetInstanceId(v string) *ModifyLindormV2InstanceSecurityGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) SetOwnerAccount(v string) *ModifyLindormV2InstanceSecurityGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) SetOwnerId(v int64) *ModifyLindormV2InstanceSecurityGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) SetResourceOwnerAccount(v string) *ModifyLindormV2InstanceSecurityGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) SetResourceOwnerId(v int64) *ModifyLindormV2InstanceSecurityGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) SetSecurityGroups(v string) *ModifyLindormV2InstanceSecurityGroupsRequest {
	s.SecurityGroups = &v
	return s
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) SetSecurityToken(v string) *ModifyLindormV2InstanceSecurityGroupsRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyLindormV2InstanceSecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}
