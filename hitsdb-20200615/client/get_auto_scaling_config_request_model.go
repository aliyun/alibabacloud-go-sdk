// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *GetAutoScalingConfigRequest
	GetConfigId() *string
	SetInstanceId(v string) *GetAutoScalingConfigRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetAutoScalingConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetAutoScalingConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetAutoScalingConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetAutoScalingConfigRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetAutoScalingConfigRequest
	GetSecurityToken() *string
}

type GetAutoScalingConfigRequest struct {
	// This parameter is required.
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetAutoScalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *GetAutoScalingConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoScalingConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetAutoScalingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetAutoScalingConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAutoScalingConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAutoScalingConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetAutoScalingConfigRequest) SetConfigId(v string) *GetAutoScalingConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetInstanceId(v string) *GetAutoScalingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetOwnerAccount(v string) *GetAutoScalingConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetOwnerId(v int64) *GetAutoScalingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetResourceOwnerAccount(v string) *GetAutoScalingConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetResourceOwnerId(v int64) *GetAutoScalingConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetSecurityToken(v string) *GetAutoScalingConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetAutoScalingConfigRequest) Validate() error {
	return dara.Validate(s)
}
