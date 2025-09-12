// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DeleteAutoScalingRuleRequest
	GetConfigId() *string
	SetInstanceId(v string) *DeleteAutoScalingRuleRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeleteAutoScalingRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteAutoScalingRuleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteAutoScalingRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAutoScalingRuleRequest
	GetResourceOwnerId() *int64
	SetRuleId(v string) *DeleteAutoScalingRuleRequest
	GetRuleId() *string
	SetSecurityToken(v string) *DeleteAutoScalingRuleRequest
	GetSecurityToken() *string
}

type DeleteAutoScalingRuleRequest struct {
	// This parameter is required.
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RuleId        *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteAutoScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingRuleRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteAutoScalingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAutoScalingRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteAutoScalingRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAutoScalingRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAutoScalingRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAutoScalingRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteAutoScalingRuleRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteAutoScalingRuleRequest) SetConfigId(v string) *DeleteAutoScalingRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetInstanceId(v string) *DeleteAutoScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetOwnerAccount(v string) *DeleteAutoScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetOwnerId(v int64) *DeleteAutoScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetResourceOwnerAccount(v string) *DeleteAutoScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetResourceOwnerId(v int64) *DeleteAutoScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetRuleId(v string) *DeleteAutoScalingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetSecurityToken(v string) *DeleteAutoScalingRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
