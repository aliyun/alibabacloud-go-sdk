// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoScalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *ModifyAutoScalingConfigRequest
	GetConfigId() *string
	SetConfigName(v string) *ModifyAutoScalingConfigRequest
	GetConfigName() *string
	SetEffectiveTimeEnd(v string) *ModifyAutoScalingConfigRequest
	GetEffectiveTimeEnd() *string
	SetEffectiveTimeStart(v string) *ModifyAutoScalingConfigRequest
	GetEffectiveTimeStart() *string
	SetEnabled(v bool) *ModifyAutoScalingConfigRequest
	GetEnabled() *bool
	SetEngine(v string) *ModifyAutoScalingConfigRequest
	GetEngine() *string
	SetInstanceId(v string) *ModifyAutoScalingConfigRequest
	GetInstanceId() *string
	SetNodesMax(v int32) *ModifyAutoScalingConfigRequest
	GetNodesMax() *int32
	SetNodesMin(v int32) *ModifyAutoScalingConfigRequest
	GetNodesMin() *int32
	SetOwnerAccount(v string) *ModifyAutoScalingConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAutoScalingConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAutoScalingConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAutoScalingConfigRequest
	GetResourceOwnerId() *int64
	SetScaleType(v string) *ModifyAutoScalingConfigRequest
	GetScaleType() *string
	SetSecurityToken(v string) *ModifyAutoScalingConfigRequest
	GetSecurityToken() *string
	SetSpecId(v string) *ModifyAutoScalingConfigRequest
	GetSpecId() *string
}

type ModifyAutoScalingConfigRequest struct {
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
	ScaleType            *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SpecId               *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s ModifyAutoScalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ModifyAutoScalingConfigRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *ModifyAutoScalingConfigRequest) GetEffectiveTimeEnd() *string {
	return s.EffectiveTimeEnd
}

func (s *ModifyAutoScalingConfigRequest) GetEffectiveTimeStart() *string {
	return s.EffectiveTimeStart
}

func (s *ModifyAutoScalingConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyAutoScalingConfigRequest) GetEngine() *string {
	return s.Engine
}

func (s *ModifyAutoScalingConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAutoScalingConfigRequest) GetNodesMax() *int32 {
	return s.NodesMax
}

func (s *ModifyAutoScalingConfigRequest) GetNodesMin() *int32 {
	return s.NodesMin
}

func (s *ModifyAutoScalingConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAutoScalingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAutoScalingConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAutoScalingConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAutoScalingConfigRequest) GetScaleType() *string {
	return s.ScaleType
}

func (s *ModifyAutoScalingConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyAutoScalingConfigRequest) GetSpecId() *string {
	return s.SpecId
}

func (s *ModifyAutoScalingConfigRequest) SetConfigId(v string) *ModifyAutoScalingConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetConfigName(v string) *ModifyAutoScalingConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetEffectiveTimeEnd(v string) *ModifyAutoScalingConfigRequest {
	s.EffectiveTimeEnd = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetEffectiveTimeStart(v string) *ModifyAutoScalingConfigRequest {
	s.EffectiveTimeStart = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetEnabled(v bool) *ModifyAutoScalingConfigRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetEngine(v string) *ModifyAutoScalingConfigRequest {
	s.Engine = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetInstanceId(v string) *ModifyAutoScalingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetNodesMax(v int32) *ModifyAutoScalingConfigRequest {
	s.NodesMax = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetNodesMin(v int32) *ModifyAutoScalingConfigRequest {
	s.NodesMin = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetOwnerAccount(v string) *ModifyAutoScalingConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetOwnerId(v int64) *ModifyAutoScalingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetResourceOwnerAccount(v string) *ModifyAutoScalingConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetResourceOwnerId(v int64) *ModifyAutoScalingConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetScaleType(v string) *ModifyAutoScalingConfigRequest {
	s.ScaleType = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetSecurityToken(v string) *ModifyAutoScalingConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetSpecId(v string) *ModifyAutoScalingConfigRequest {
	s.SpecId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) Validate() error {
	return dara.Validate(s)
}
