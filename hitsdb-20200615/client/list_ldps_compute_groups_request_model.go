// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLdpsComputeGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListLdpsComputeGroupsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ListLdpsComputeGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListLdpsComputeGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListLdpsComputeGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListLdpsComputeGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListLdpsComputeGroupsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ListLdpsComputeGroupsRequest
	GetSecurityToken() *string
}

type ListLdpsComputeGroupsRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListLdpsComputeGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLdpsComputeGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListLdpsComputeGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLdpsComputeGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListLdpsComputeGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListLdpsComputeGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLdpsComputeGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListLdpsComputeGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListLdpsComputeGroupsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListLdpsComputeGroupsRequest) SetInstanceId(v string) *ListLdpsComputeGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetOwnerAccount(v string) *ListLdpsComputeGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetOwnerId(v int64) *ListLdpsComputeGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetRegionId(v string) *ListLdpsComputeGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetResourceOwnerAccount(v string) *ListLdpsComputeGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetResourceOwnerId(v int64) *ListLdpsComputeGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetSecurityToken(v string) *ListLdpsComputeGroupsRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) Validate() error {
	return dara.Validate(s)
}
