// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceSecurityGroupsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetInstanceSecurityGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetInstanceSecurityGroupsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetInstanceSecurityGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetInstanceSecurityGroupsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetInstanceSecurityGroupsRequest
	GetSecurityToken() *string
}

type GetInstanceSecurityGroupsRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetInstanceSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceSecurityGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceSecurityGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetInstanceSecurityGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetInstanceSecurityGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetInstanceSecurityGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetInstanceSecurityGroupsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetInstanceSecurityGroupsRequest) SetInstanceId(v string) *GetInstanceSecurityGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceSecurityGroupsRequest) SetOwnerAccount(v string) *GetInstanceSecurityGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetInstanceSecurityGroupsRequest) SetOwnerId(v int64) *GetInstanceSecurityGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetInstanceSecurityGroupsRequest) SetResourceOwnerAccount(v string) *GetInstanceSecurityGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetInstanceSecurityGroupsRequest) SetResourceOwnerId(v int64) *GetInstanceSecurityGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetInstanceSecurityGroupsRequest) SetSecurityToken(v string) *GetInstanceSecurityGroupsRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetInstanceSecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}
