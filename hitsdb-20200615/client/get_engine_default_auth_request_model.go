// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEngineDefaultAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetEngineDefaultAuthRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetEngineDefaultAuthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetEngineDefaultAuthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetEngineDefaultAuthRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetEngineDefaultAuthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetEngineDefaultAuthRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetEngineDefaultAuthRequest
	GetSecurityToken() *string
}

type GetEngineDefaultAuthRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetEngineDefaultAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEngineDefaultAuthRequest) GoString() string {
	return s.String()
}

func (s *GetEngineDefaultAuthRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEngineDefaultAuthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetEngineDefaultAuthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetEngineDefaultAuthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEngineDefaultAuthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetEngineDefaultAuthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetEngineDefaultAuthRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetEngineDefaultAuthRequest) SetInstanceId(v string) *GetEngineDefaultAuthRequest {
	s.InstanceId = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetOwnerAccount(v string) *GetEngineDefaultAuthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetOwnerId(v int64) *GetEngineDefaultAuthRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetRegionId(v string) *GetEngineDefaultAuthRequest {
	s.RegionId = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetResourceOwnerAccount(v string) *GetEngineDefaultAuthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetResourceOwnerId(v int64) *GetEngineDefaultAuthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetSecurityToken(v string) *GetEngineDefaultAuthRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) Validate() error {
	return dara.Validate(s)
}
