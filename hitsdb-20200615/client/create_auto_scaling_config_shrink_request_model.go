// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoScalingConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigName(v string) *CreateAutoScalingConfigShrinkRequest
	GetConfigName() *string
	SetEffectiveTimeEnd(v string) *CreateAutoScalingConfigShrinkRequest
	GetEffectiveTimeEnd() *string
	SetEffectiveTimeStart(v string) *CreateAutoScalingConfigShrinkRequest
	GetEffectiveTimeStart() *string
	SetEnabled(v bool) *CreateAutoScalingConfigShrinkRequest
	GetEnabled() *bool
	SetEngine(v string) *CreateAutoScalingConfigShrinkRequest
	GetEngine() *string
	SetInstanceId(v string) *CreateAutoScalingConfigShrinkRequest
	GetInstanceId() *string
	SetNodesMax(v int32) *CreateAutoScalingConfigShrinkRequest
	GetNodesMax() *int32
	SetNodesMin(v int32) *CreateAutoScalingConfigShrinkRequest
	GetNodesMin() *int32
	SetOwnerAccount(v string) *CreateAutoScalingConfigShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAutoScalingConfigShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateAutoScalingConfigShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAutoScalingConfigShrinkRequest
	GetResourceOwnerId() *int64
	SetScaleRuleListShrink(v string) *CreateAutoScalingConfigShrinkRequest
	GetScaleRuleListShrink() *string
	SetScaleType(v string) *CreateAutoScalingConfigShrinkRequest
	GetScaleType() *string
	SetSecurityToken(v string) *CreateAutoScalingConfigShrinkRequest
	GetSecurityToken() *string
	SetSpecId(v string) *CreateAutoScalingConfigShrinkRequest
	GetSpecId() *string
}

type CreateAutoScalingConfigShrinkRequest struct {
	// This parameter is required.
	ConfigName         *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	EffectiveTimeEnd   *string `json:"EffectiveTimeEnd,omitempty" xml:"EffectiveTimeEnd,omitempty"`
	EffectiveTimeStart *string `json:"EffectiveTimeStart,omitempty" xml:"EffectiveTimeStart,omitempty"`
	Enabled            *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodesMax             *int32  `json:"NodesMax,omitempty" xml:"NodesMax,omitempty"`
	NodesMin             *int32  `json:"NodesMin,omitempty" xml:"NodesMin,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleRuleListShrink  *string `json:"ScaleRuleList,omitempty" xml:"ScaleRuleList,omitempty"`
	// This parameter is required.
	ScaleType     *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s CreateAutoScalingConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoScalingConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingConfigShrinkRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *CreateAutoScalingConfigShrinkRequest) GetEffectiveTimeEnd() *string {
	return s.EffectiveTimeEnd
}

func (s *CreateAutoScalingConfigShrinkRequest) GetEffectiveTimeStart() *string {
	return s.EffectiveTimeStart
}

func (s *CreateAutoScalingConfigShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAutoScalingConfigShrinkRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateAutoScalingConfigShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAutoScalingConfigShrinkRequest) GetNodesMax() *int32 {
	return s.NodesMax
}

func (s *CreateAutoScalingConfigShrinkRequest) GetNodesMin() *int32 {
	return s.NodesMin
}

func (s *CreateAutoScalingConfigShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAutoScalingConfigShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAutoScalingConfigShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAutoScalingConfigShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAutoScalingConfigShrinkRequest) GetScaleRuleListShrink() *string {
	return s.ScaleRuleListShrink
}

func (s *CreateAutoScalingConfigShrinkRequest) GetScaleType() *string {
	return s.ScaleType
}

func (s *CreateAutoScalingConfigShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateAutoScalingConfigShrinkRequest) GetSpecId() *string {
	return s.SpecId
}

func (s *CreateAutoScalingConfigShrinkRequest) SetConfigName(v string) *CreateAutoScalingConfigShrinkRequest {
	s.ConfigName = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetEffectiveTimeEnd(v string) *CreateAutoScalingConfigShrinkRequest {
	s.EffectiveTimeEnd = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetEffectiveTimeStart(v string) *CreateAutoScalingConfigShrinkRequest {
	s.EffectiveTimeStart = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetEnabled(v bool) *CreateAutoScalingConfigShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetEngine(v string) *CreateAutoScalingConfigShrinkRequest {
	s.Engine = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetInstanceId(v string) *CreateAutoScalingConfigShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetNodesMax(v int32) *CreateAutoScalingConfigShrinkRequest {
	s.NodesMax = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetNodesMin(v int32) *CreateAutoScalingConfigShrinkRequest {
	s.NodesMin = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetOwnerAccount(v string) *CreateAutoScalingConfigShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetOwnerId(v int64) *CreateAutoScalingConfigShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetResourceOwnerAccount(v string) *CreateAutoScalingConfigShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetResourceOwnerId(v int64) *CreateAutoScalingConfigShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetScaleRuleListShrink(v string) *CreateAutoScalingConfigShrinkRequest {
	s.ScaleRuleListShrink = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetScaleType(v string) *CreateAutoScalingConfigShrinkRequest {
	s.ScaleType = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetSecurityToken(v string) *CreateAutoScalingConfigShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetSpecId(v string) *CreateAutoScalingConfigShrinkRequest {
	s.SpecId = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
