// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientSourceIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetClientSourceIpRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetClientSourceIpRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetClientSourceIpRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetClientSourceIpRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetClientSourceIpRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetClientSourceIpRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetClientSourceIpRequest
	GetSecurityToken() *string
}

type GetClientSourceIpRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetClientSourceIpRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientSourceIpRequest) GoString() string {
	return s.String()
}

func (s *GetClientSourceIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetClientSourceIpRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetClientSourceIpRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetClientSourceIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetClientSourceIpRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetClientSourceIpRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetClientSourceIpRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetClientSourceIpRequest) SetInstanceId(v string) *GetClientSourceIpRequest {
	s.InstanceId = &v
	return s
}

func (s *GetClientSourceIpRequest) SetOwnerAccount(v string) *GetClientSourceIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetClientSourceIpRequest) SetOwnerId(v int64) *GetClientSourceIpRequest {
	s.OwnerId = &v
	return s
}

func (s *GetClientSourceIpRequest) SetRegionId(v string) *GetClientSourceIpRequest {
	s.RegionId = &v
	return s
}

func (s *GetClientSourceIpRequest) SetResourceOwnerAccount(v string) *GetClientSourceIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetClientSourceIpRequest) SetResourceOwnerId(v int64) *GetClientSourceIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetClientSourceIpRequest) SetSecurityToken(v string) *GetClientSourceIpRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetClientSourceIpRequest) Validate() error {
	return dara.Validate(s)
}
