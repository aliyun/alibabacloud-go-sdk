// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoScalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DeleteAutoScalingConfigRequest
	GetConfigId() *string
	SetInstanceId(v string) *DeleteAutoScalingConfigRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeleteAutoScalingConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteAutoScalingConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteAutoScalingConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAutoScalingConfigRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DeleteAutoScalingConfigRequest
	GetSecurityToken() *string
}

type DeleteAutoScalingConfigRequest struct {
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

func (s DeleteAutoScalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteAutoScalingConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAutoScalingConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteAutoScalingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAutoScalingConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAutoScalingConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAutoScalingConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteAutoScalingConfigRequest) SetConfigId(v string) *DeleteAutoScalingConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetInstanceId(v string) *DeleteAutoScalingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetOwnerAccount(v string) *DeleteAutoScalingConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetOwnerId(v int64) *DeleteAutoScalingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetResourceOwnerAccount(v string) *DeleteAutoScalingConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetResourceOwnerId(v int64) *DeleteAutoScalingConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetSecurityToken(v string) *DeleteAutoScalingConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) Validate() error {
	return dara.Validate(s)
}
