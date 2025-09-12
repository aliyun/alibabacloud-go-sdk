// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *ModifyAutoScalingRuleRequest
	GetConfigId() *string
	SetEnabled(v bool) *ModifyAutoScalingRuleRequest
	GetEnabled() *bool
	SetEndTime(v string) *ModifyAutoScalingRuleRequest
	GetEndTime() *string
	SetInstanceId(v string) *ModifyAutoScalingRuleRequest
	GetInstanceId() *string
	SetObservationWindow(v int32) *ModifyAutoScalingRuleRequest
	GetObservationWindow() *int32
	SetOperationType(v string) *ModifyAutoScalingRuleRequest
	GetOperationType() *string
	SetOwnerAccount(v string) *ModifyAutoScalingRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAutoScalingRuleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAutoScalingRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAutoScalingRuleRequest
	GetResourceOwnerId() *int64
	SetRuleId(v string) *ModifyAutoScalingRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *ModifyAutoScalingRuleRequest
	GetRuleName() *string
	SetRuleType(v string) *ModifyAutoScalingRuleRequest
	GetRuleType() *string
	SetScaleInStep(v int32) *ModifyAutoScalingRuleRequest
	GetScaleInStep() *int32
	SetScaleOutStep(v int32) *ModifyAutoScalingRuleRequest
	GetScaleOutStep() *int32
	SetSecurityToken(v string) *ModifyAutoScalingRuleRequest
	GetSecurityToken() *string
	SetSilenceTime(v int32) *ModifyAutoScalingRuleRequest
	GetSilenceTime() *int32
	SetStartTime(v string) *ModifyAutoScalingRuleRequest
	GetStartTime() *string
	SetTargetMetric(v string) *ModifyAutoScalingRuleRequest
	GetTargetMetric() *string
	SetTargetNodes(v int32) *ModifyAutoScalingRuleRequest
	GetTargetNodes() *int32
	SetThresholdLower(v int32) *ModifyAutoScalingRuleRequest
	GetThresholdLower() *int32
	SetThresholdUpper(v int32) *ModifyAutoScalingRuleRequest
	GetThresholdUpper() *int32
	SetTriggerCronExpr(v string) *ModifyAutoScalingRuleRequest
	GetTriggerCronExpr() *string
}

type ModifyAutoScalingRuleRequest struct {
	// This parameter is required.
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enabled  *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ObservationWindow    *int32  `json:"ObservationWindow,omitempty" xml:"ObservationWindow,omitempty"`
	OperationType        *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RuleId          *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType        *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScaleInStep     *int32  `json:"ScaleInStep,omitempty" xml:"ScaleInStep,omitempty"`
	ScaleOutStep    *int32  `json:"ScaleOutStep,omitempty" xml:"ScaleOutStep,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SilenceTime     *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TargetMetric    *string `json:"TargetMetric,omitempty" xml:"TargetMetric,omitempty"`
	TargetNodes     *int32  `json:"TargetNodes,omitempty" xml:"TargetNodes,omitempty"`
	ThresholdLower  *int32  `json:"ThresholdLower,omitempty" xml:"ThresholdLower,omitempty"`
	ThresholdUpper  *int32  `json:"ThresholdUpper,omitempty" xml:"ThresholdUpper,omitempty"`
	TriggerCronExpr *string `json:"TriggerCronExpr,omitempty" xml:"TriggerCronExpr,omitempty"`
}

func (s ModifyAutoScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingRuleRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ModifyAutoScalingRuleRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyAutoScalingRuleRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyAutoScalingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAutoScalingRuleRequest) GetObservationWindow() *int32 {
	return s.ObservationWindow
}

func (s *ModifyAutoScalingRuleRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ModifyAutoScalingRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAutoScalingRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAutoScalingRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAutoScalingRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAutoScalingRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *ModifyAutoScalingRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyAutoScalingRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ModifyAutoScalingRuleRequest) GetScaleInStep() *int32 {
	return s.ScaleInStep
}

func (s *ModifyAutoScalingRuleRequest) GetScaleOutStep() *int32 {
	return s.ScaleOutStep
}

func (s *ModifyAutoScalingRuleRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyAutoScalingRuleRequest) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *ModifyAutoScalingRuleRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyAutoScalingRuleRequest) GetTargetMetric() *string {
	return s.TargetMetric
}

func (s *ModifyAutoScalingRuleRequest) GetTargetNodes() *int32 {
	return s.TargetNodes
}

func (s *ModifyAutoScalingRuleRequest) GetThresholdLower() *int32 {
	return s.ThresholdLower
}

func (s *ModifyAutoScalingRuleRequest) GetThresholdUpper() *int32 {
	return s.ThresholdUpper
}

func (s *ModifyAutoScalingRuleRequest) GetTriggerCronExpr() *string {
	return s.TriggerCronExpr
}

func (s *ModifyAutoScalingRuleRequest) SetConfigId(v string) *ModifyAutoScalingRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetEnabled(v bool) *ModifyAutoScalingRuleRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetEndTime(v string) *ModifyAutoScalingRuleRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetInstanceId(v string) *ModifyAutoScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetObservationWindow(v int32) *ModifyAutoScalingRuleRequest {
	s.ObservationWindow = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetOperationType(v string) *ModifyAutoScalingRuleRequest {
	s.OperationType = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetOwnerAccount(v string) *ModifyAutoScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetOwnerId(v int64) *ModifyAutoScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetResourceOwnerAccount(v string) *ModifyAutoScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetResourceOwnerId(v int64) *ModifyAutoScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetRuleId(v string) *ModifyAutoScalingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetRuleName(v string) *ModifyAutoScalingRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetRuleType(v string) *ModifyAutoScalingRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetScaleInStep(v int32) *ModifyAutoScalingRuleRequest {
	s.ScaleInStep = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetScaleOutStep(v int32) *ModifyAutoScalingRuleRequest {
	s.ScaleOutStep = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetSecurityToken(v string) *ModifyAutoScalingRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetSilenceTime(v int32) *ModifyAutoScalingRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetStartTime(v string) *ModifyAutoScalingRuleRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetTargetMetric(v string) *ModifyAutoScalingRuleRequest {
	s.TargetMetric = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetTargetNodes(v int32) *ModifyAutoScalingRuleRequest {
	s.TargetNodes = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetThresholdLower(v int32) *ModifyAutoScalingRuleRequest {
	s.ThresholdLower = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetThresholdUpper(v int32) *ModifyAutoScalingRuleRequest {
	s.ThresholdUpper = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetTriggerCronExpr(v string) *ModifyAutoScalingRuleRequest {
	s.TriggerCronExpr = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
