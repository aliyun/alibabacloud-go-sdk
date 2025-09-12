// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLindormV2InstanceSecurityGroupsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLindormV2InstanceSecurityGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormV2InstanceSecurityGroupsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetLindormV2InstanceSecurityGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormV2InstanceSecurityGroupsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormV2InstanceSecurityGroupsRequest
	GetSecurityToken() *string
}

type GetLindormV2InstanceSecurityGroupsRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormV2InstanceSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) SetInstanceId(v string) *GetLindormV2InstanceSecurityGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) SetOwnerAccount(v string) *GetLindormV2InstanceSecurityGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) SetOwnerId(v int64) *GetLindormV2InstanceSecurityGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) SetResourceOwnerAccount(v string) *GetLindormV2InstanceSecurityGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) SetResourceOwnerId(v int64) *GetLindormV2InstanceSecurityGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) SetSecurityToken(v string) *GetLindormV2InstanceSecurityGroupsRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}
