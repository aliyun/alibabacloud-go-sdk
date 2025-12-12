// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoScalingConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *ModifyAutoScalingConfigShrinkRequest
	GetConfigId() *string
	SetConfigName(v string) *ModifyAutoScalingConfigShrinkRequest
	GetConfigName() *string
	SetEffectiveTimeEnd(v string) *ModifyAutoScalingConfigShrinkRequest
	GetEffectiveTimeEnd() *string
	SetEffectiveTimeStart(v string) *ModifyAutoScalingConfigShrinkRequest
	GetEffectiveTimeStart() *string
	SetEnabled(v bool) *ModifyAutoScalingConfigShrinkRequest
	GetEnabled() *bool
	SetEngine(v string) *ModifyAutoScalingConfigShrinkRequest
	GetEngine() *string
	SetInstanceId(v string) *ModifyAutoScalingConfigShrinkRequest
	GetInstanceId() *string
	SetNodesMax(v int32) *ModifyAutoScalingConfigShrinkRequest
	GetNodesMax() *int32
	SetNodesMin(v int32) *ModifyAutoScalingConfigShrinkRequest
	GetNodesMin() *int32
	SetOwnerAccount(v string) *ModifyAutoScalingConfigShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAutoScalingConfigShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAutoScalingConfigShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAutoScalingConfigShrinkRequest
	GetResourceOwnerId() *int64
	SetScaleRuleListShrink(v string) *ModifyAutoScalingConfigShrinkRequest
	GetScaleRuleListShrink() *string
	SetScaleType(v string) *ModifyAutoScalingConfigShrinkRequest
	GetScaleType() *string
	SetSecurityToken(v string) *ModifyAutoScalingConfigShrinkRequest
	GetSecurityToken() *string
	SetSpecId(v string) *ModifyAutoScalingConfigShrinkRequest
	GetSpecId() *string
	SetStorageCapacityMax(v int64) *ModifyAutoScalingConfigShrinkRequest
	GetStorageCapacityMax() *int64
}

type ModifyAutoScalingConfigShrinkRequest struct {
	// This parameter is required.
	ConfigId           *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigName         *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	EffectiveTimeEnd   *string `json:"EffectiveTimeEnd,omitempty" xml:"EffectiveTimeEnd,omitempty"`
	EffectiveTimeStart *string `json:"EffectiveTimeStart,omitempty" xml:"EffectiveTimeStart,omitempty"`
	Enabled            *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Engine             *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodesMax             *int32  `json:"NodesMax,omitempty" xml:"NodesMax,omitempty"`
	NodesMin             *int32  `json:"NodesMin,omitempty" xml:"NodesMin,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleRuleListShrink  *string `json:"ScaleRuleList,omitempty" xml:"ScaleRuleList,omitempty"`
	ScaleType            *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SpecId               *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	StorageCapacityMax   *int64  `json:"StorageCapacityMax,omitempty" xml:"StorageCapacityMax,omitempty"`
}

func (s ModifyAutoScalingConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetEffectiveTimeEnd() *string {
	return s.EffectiveTimeEnd
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetEffectiveTimeStart() *string {
	return s.EffectiveTimeStart
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetEngine() *string {
	return s.Engine
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetNodesMax() *int32 {
	return s.NodesMax
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetNodesMin() *int32 {
	return s.NodesMin
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetScaleRuleListShrink() *string {
	return s.ScaleRuleListShrink
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetScaleType() *string {
	return s.ScaleType
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetSpecId() *string {
	return s.SpecId
}

func (s *ModifyAutoScalingConfigShrinkRequest) GetStorageCapacityMax() *int64 {
	return s.StorageCapacityMax
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetConfigId(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.ConfigId = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetConfigName(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.ConfigName = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetEffectiveTimeEnd(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.EffectiveTimeEnd = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetEffectiveTimeStart(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.EffectiveTimeStart = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetEnabled(v bool) *ModifyAutoScalingConfigShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetEngine(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.Engine = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetInstanceId(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetNodesMax(v int32) *ModifyAutoScalingConfigShrinkRequest {
	s.NodesMax = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetNodesMin(v int32) *ModifyAutoScalingConfigShrinkRequest {
	s.NodesMin = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetOwnerAccount(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetOwnerId(v int64) *ModifyAutoScalingConfigShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetResourceOwnerAccount(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetResourceOwnerId(v int64) *ModifyAutoScalingConfigShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetScaleRuleListShrink(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.ScaleRuleListShrink = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetScaleType(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.ScaleType = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetSecurityToken(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetSpecId(v string) *ModifyAutoScalingConfigShrinkRequest {
	s.SpecId = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) SetStorageCapacityMax(v int64) *ModifyAutoScalingConfigShrinkRequest {
	s.StorageCapacityMax = &v
	return s
}

func (s *ModifyAutoScalingConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
