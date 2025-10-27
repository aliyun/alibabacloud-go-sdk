// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoScalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigName(v string) *CreateAutoScalingConfigRequest
	GetConfigName() *string
	SetEffectiveTimeEnd(v string) *CreateAutoScalingConfigRequest
	GetEffectiveTimeEnd() *string
	SetEffectiveTimeStart(v string) *CreateAutoScalingConfigRequest
	GetEffectiveTimeStart() *string
	SetEnabled(v bool) *CreateAutoScalingConfigRequest
	GetEnabled() *bool
	SetEngine(v string) *CreateAutoScalingConfigRequest
	GetEngine() *string
	SetInstanceId(v string) *CreateAutoScalingConfigRequest
	GetInstanceId() *string
	SetNodesMax(v int32) *CreateAutoScalingConfigRequest
	GetNodesMax() *int32
	SetNodesMin(v int32) *CreateAutoScalingConfigRequest
	GetNodesMin() *int32
	SetOwnerAccount(v string) *CreateAutoScalingConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAutoScalingConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateAutoScalingConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAutoScalingConfigRequest
	GetResourceOwnerId() *int64
	SetScaleRuleList(v []*CreateAutoScalingConfigRequestScaleRuleList) *CreateAutoScalingConfigRequest
	GetScaleRuleList() []*CreateAutoScalingConfigRequestScaleRuleList
	SetScaleType(v string) *CreateAutoScalingConfigRequest
	GetScaleType() *string
	SetSecurityToken(v string) *CreateAutoScalingConfigRequest
	GetSecurityToken() *string
	SetSpecId(v string) *CreateAutoScalingConfigRequest
	GetSpecId() *string
}

type CreateAutoScalingConfigRequest struct {
	// This parameter is required.
	ConfigName         *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	EffectiveTimeEnd   *string `json:"EffectiveTimeEnd,omitempty" xml:"EffectiveTimeEnd,omitempty"`
	EffectiveTimeStart *string `json:"EffectiveTimeStart,omitempty" xml:"EffectiveTimeStart,omitempty"`
	Enabled            *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	InstanceId           *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodesMax             *int32                                         `json:"NodesMax,omitempty" xml:"NodesMax,omitempty"`
	NodesMin             *int32                                         `json:"NodesMin,omitempty" xml:"NodesMin,omitempty"`
	OwnerAccount         *string                                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleRuleList        []*CreateAutoScalingConfigRequestScaleRuleList `json:"ScaleRuleList,omitempty" xml:"ScaleRuleList,omitempty" type:"Repeated"`
	// This parameter is required.
	ScaleType     *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s CreateAutoScalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingConfigRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *CreateAutoScalingConfigRequest) GetEffectiveTimeEnd() *string {
	return s.EffectiveTimeEnd
}

func (s *CreateAutoScalingConfigRequest) GetEffectiveTimeStart() *string {
	return s.EffectiveTimeStart
}

func (s *CreateAutoScalingConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAutoScalingConfigRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateAutoScalingConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAutoScalingConfigRequest) GetNodesMax() *int32 {
	return s.NodesMax
}

func (s *CreateAutoScalingConfigRequest) GetNodesMin() *int32 {
	return s.NodesMin
}

func (s *CreateAutoScalingConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAutoScalingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAutoScalingConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAutoScalingConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAutoScalingConfigRequest) GetScaleRuleList() []*CreateAutoScalingConfigRequestScaleRuleList {
	return s.ScaleRuleList
}

func (s *CreateAutoScalingConfigRequest) GetScaleType() *string {
	return s.ScaleType
}

func (s *CreateAutoScalingConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateAutoScalingConfigRequest) GetSpecId() *string {
	return s.SpecId
}

func (s *CreateAutoScalingConfigRequest) SetConfigName(v string) *CreateAutoScalingConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetEffectiveTimeEnd(v string) *CreateAutoScalingConfigRequest {
	s.EffectiveTimeEnd = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetEffectiveTimeStart(v string) *CreateAutoScalingConfigRequest {
	s.EffectiveTimeStart = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetEnabled(v bool) *CreateAutoScalingConfigRequest {
	s.Enabled = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetEngine(v string) *CreateAutoScalingConfigRequest {
	s.Engine = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetInstanceId(v string) *CreateAutoScalingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetNodesMax(v int32) *CreateAutoScalingConfigRequest {
	s.NodesMax = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetNodesMin(v int32) *CreateAutoScalingConfigRequest {
	s.NodesMin = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetOwnerAccount(v string) *CreateAutoScalingConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetOwnerId(v int64) *CreateAutoScalingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetResourceOwnerAccount(v string) *CreateAutoScalingConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetResourceOwnerId(v int64) *CreateAutoScalingConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetScaleRuleList(v []*CreateAutoScalingConfigRequestScaleRuleList) *CreateAutoScalingConfigRequest {
	s.ScaleRuleList = v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetScaleType(v string) *CreateAutoScalingConfigRequest {
	s.ScaleType = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetSecurityToken(v string) *CreateAutoScalingConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetSpecId(v string) *CreateAutoScalingConfigRequest {
	s.SpecId = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) Validate() error {
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

type CreateAutoScalingConfigRequestScaleRuleList struct {
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

func (s CreateAutoScalingConfigRequestScaleRuleList) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoScalingConfigRequestScaleRuleList) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetConfigId() *string {
	return s.ConfigId
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetObservationWindow() *int32 {
	return s.ObservationWindow
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetScaleInStep() *int32 {
	return s.ScaleInStep
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetScaleOutStep() *int32 {
	return s.ScaleOutStep
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetTargetMetric() *string {
	return s.TargetMetric
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetTargetNodes() *int32 {
	return s.TargetNodes
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetThresholdLower() *int32 {
	return s.ThresholdLower
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetThresholdUpper() *int32 {
	return s.ThresholdUpper
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) GetTriggerCronExpr() *string {
	return s.TriggerCronExpr
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetConfigId(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ConfigId = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetEnabled(v bool) *CreateAutoScalingConfigRequestScaleRuleList {
	s.Enabled = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetEndTime(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.EndTime = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetInstanceId(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.InstanceId = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetObservationWindow(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ObservationWindow = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetOperationType(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.OperationType = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetRuleId(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.RuleId = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetRuleName(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.RuleName = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetRuleType(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.RuleType = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetScaleInStep(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ScaleInStep = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetScaleOutStep(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ScaleOutStep = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetSilenceTime(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.SilenceTime = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetStartTime(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.StartTime = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetTargetMetric(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.TargetMetric = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetTargetNodes(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.TargetNodes = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetThresholdLower(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ThresholdLower = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetThresholdUpper(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ThresholdUpper = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetTriggerCronExpr(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.TriggerCronExpr = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) Validate() error {
	return dara.Validate(s)
}
