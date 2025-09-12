// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *GetAutoScalingRuleRequest
	GetConfigId() *string
	SetInstanceId(v string) *GetAutoScalingRuleRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetAutoScalingRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetAutoScalingRuleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetAutoScalingRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetAutoScalingRuleRequest
	GetResourceOwnerId() *int64
	SetRuleId(v string) *GetAutoScalingRuleRequest
	GetRuleId() *string
	SetSecurityToken(v string) *GetAutoScalingRuleRequest
	GetSecurityToken() *string
}

type GetAutoScalingRuleRequest struct {
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

func (s GetAutoScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScalingRuleRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *GetAutoScalingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoScalingRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetAutoScalingRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetAutoScalingRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAutoScalingRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAutoScalingRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *GetAutoScalingRuleRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetAutoScalingRuleRequest) SetConfigId(v string) *GetAutoScalingRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetInstanceId(v string) *GetAutoScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetOwnerAccount(v string) *GetAutoScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetOwnerId(v int64) *GetAutoScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetResourceOwnerAccount(v string) *GetAutoScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetResourceOwnerId(v int64) *GetAutoScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetRuleId(v string) *GetAutoScalingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetSecurityToken(v string) *GetAutoScalingRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetAutoScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
