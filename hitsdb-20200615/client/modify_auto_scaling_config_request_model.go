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
	SetScaleRuleList(v []*ModifyAutoScalingConfigRequestScaleRuleList) *ModifyAutoScalingConfigRequest
	GetScaleRuleList() []*ModifyAutoScalingConfigRequestScaleRuleList
	SetScaleType(v string) *ModifyAutoScalingConfigRequest
	GetScaleType() *string
	SetSecurityToken(v string) *ModifyAutoScalingConfigRequest
	GetSecurityToken() *string
	SetSpecId(v string) *ModifyAutoScalingConfigRequest
	GetSpecId() *string
	SetStorageCapacityMax(v int64) *ModifyAutoScalingConfigRequest
	GetStorageCapacityMax() *int64
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
	InstanceId           *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodesMax             *int32                                         `json:"NodesMax,omitempty" xml:"NodesMax,omitempty"`
	NodesMin             *int32                                         `json:"NodesMin,omitempty" xml:"NodesMin,omitempty"`
	OwnerAccount         *string                                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleRuleList        []*ModifyAutoScalingConfigRequestScaleRuleList `json:"ScaleRuleList,omitempty" xml:"ScaleRuleList,omitempty" type:"Repeated"`
	ScaleType            *string                                        `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SecurityToken        *string                                        `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SpecId               *string                                        `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	StorageCapacityMax   *int64                                         `json:"StorageCapacityMax,omitempty" xml:"StorageCapacityMax,omitempty"`
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

func (s *ModifyAutoScalingConfigRequest) GetScaleRuleList() []*ModifyAutoScalingConfigRequestScaleRuleList {
	return s.ScaleRuleList
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

func (s *ModifyAutoScalingConfigRequest) GetStorageCapacityMax() *int64 {
	return s.StorageCapacityMax
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

func (s *ModifyAutoScalingConfigRequest) SetScaleRuleList(v []*ModifyAutoScalingConfigRequestScaleRuleList) *ModifyAutoScalingConfigRequest {
	s.ScaleRuleList = v
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

func (s *ModifyAutoScalingConfigRequest) SetStorageCapacityMax(v int64) *ModifyAutoScalingConfigRequest {
	s.StorageCapacityMax = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) Validate() error {
	if s.ScaleRuleList != nil {
		for _, item := range s.ScaleRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyAutoScalingConfigRequestScaleRuleList struct {
	ConfigId          *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enabled           *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ObservationWindow *int32  `json:"ObservationWindow,omitempty" xml:"ObservationWindow,omitempty"`
	OperationType     *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	RuleId            *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType          *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScaleInStep       *int32  `json:"ScaleInStep,omitempty" xml:"ScaleInStep,omitempty"`
	ScaleOutStep      *int32  `json:"ScaleOutStep,omitempty" xml:"ScaleOutStep,omitempty"`
	SilenceTime       *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TargetMetric      *string `json:"TargetMetric,omitempty" xml:"TargetMetric,omitempty"`
	TargetNodes       *int32  `json:"TargetNodes,omitempty" xml:"TargetNodes,omitempty"`
	ThresholdLower    *int32  `json:"ThresholdLower,omitempty" xml:"ThresholdLower,omitempty"`
	ThresholdUpper    *int32  `json:"ThresholdUpper,omitempty" xml:"ThresholdUpper,omitempty"`
	TriggerCronExpr   *string `json:"TriggerCronExpr,omitempty" xml:"TriggerCronExpr,omitempty"`
}

func (s ModifyAutoScalingConfigRequestScaleRuleList) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigRequestScaleRuleList) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetConfigId() *string {
	return s.ConfigId
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetObservationWindow() *int32 {
	return s.ObservationWindow
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetOperationType() *string {
	return s.OperationType
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetRuleId() *string {
	return s.RuleId
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetScaleInStep() *int32 {
	return s.ScaleInStep
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetScaleOutStep() *int32 {
	return s.ScaleOutStep
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetTargetMetric() *string {
	return s.TargetMetric
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetTargetNodes() *int32 {
	return s.TargetNodes
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetThresholdLower() *int32 {
	return s.ThresholdLower
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetThresholdUpper() *int32 {
	return s.ThresholdUpper
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) GetTriggerCronExpr() *string {
	return s.TriggerCronExpr
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetConfigId(v string) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.ConfigId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetEnabled(v bool) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.Enabled = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetEndTime(v string) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.EndTime = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetInstanceId(v string) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.InstanceId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetObservationWindow(v int32) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.ObservationWindow = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetOperationType(v string) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.OperationType = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetRuleId(v string) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.RuleId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetRuleName(v string) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.RuleName = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetRuleType(v string) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.RuleType = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetScaleInStep(v int32) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.ScaleInStep = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetScaleOutStep(v int32) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.ScaleOutStep = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetSilenceTime(v int32) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.SilenceTime = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetStartTime(v string) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.StartTime = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetTargetMetric(v string) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.TargetMetric = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetTargetNodes(v int32) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.TargetNodes = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetThresholdLower(v int32) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.ThresholdLower = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetThresholdUpper(v int32) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.ThresholdUpper = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) SetTriggerCronExpr(v string) *ModifyAutoScalingConfigRequestScaleRuleList {
	s.TriggerCronExpr = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestScaleRuleList) Validate() error {
	return dara.Validate(s)
}
