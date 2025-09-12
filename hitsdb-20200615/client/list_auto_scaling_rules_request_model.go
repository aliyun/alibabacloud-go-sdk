// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *ListAutoScalingRulesRequest
	GetConfigId() *string
	SetInstanceId(v string) *ListAutoScalingRulesRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ListAutoScalingRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListAutoScalingRulesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListAutoScalingRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAutoScalingRulesRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ListAutoScalingRulesRequest
	GetSecurityToken() *string
}

type ListAutoScalingRulesRequest struct {
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

func (s ListAutoScalingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRulesRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ListAutoScalingRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAutoScalingRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAutoScalingRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAutoScalingRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAutoScalingRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAutoScalingRulesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListAutoScalingRulesRequest) SetConfigId(v string) *ListAutoScalingRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetInstanceId(v string) *ListAutoScalingRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetOwnerAccount(v string) *ListAutoScalingRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetOwnerId(v int64) *ListAutoScalingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetResourceOwnerAccount(v string) *ListAutoScalingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetResourceOwnerId(v int64) *ListAutoScalingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetSecurityToken(v string) *ListAutoScalingRulesRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListAutoScalingRulesRequest) Validate() error {
	return dara.Validate(s)
}
