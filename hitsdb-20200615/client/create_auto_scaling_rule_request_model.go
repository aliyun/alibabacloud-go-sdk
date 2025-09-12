// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *CreateAutoScalingRuleRequest
	GetConfigId() *string
	SetEnabled(v bool) *CreateAutoScalingRuleRequest
	GetEnabled() *bool
	SetEndTime(v string) *CreateAutoScalingRuleRequest
	GetEndTime() *string
	SetInstanceId(v string) *CreateAutoScalingRuleRequest
	GetInstanceId() *string
	SetObservationWindow(v int32) *CreateAutoScalingRuleRequest
	GetObservationWindow() *int32
	SetOperationType(v string) *CreateAutoScalingRuleRequest
	GetOperationType() *string
	SetOwnerAccount(v string) *CreateAutoScalingRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAutoScalingRuleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateAutoScalingRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAutoScalingRuleRequest
	GetResourceOwnerId() *int64
	SetRuleName(v string) *CreateAutoScalingRuleRequest
	GetRuleName() *string
	SetRuleType(v string) *CreateAutoScalingRuleRequest
	GetRuleType() *string
	SetScaleInStep(v int32) *CreateAutoScalingRuleRequest
	GetScaleInStep() *int32
	SetScaleOutStep(v int32) *CreateAutoScalingRuleRequest
	GetScaleOutStep() *int32
	SetSecurityToken(v string) *CreateAutoScalingRuleRequest
	GetSecurityToken() *string
	SetSilenceTime(v int32) *CreateAutoScalingRuleRequest
	GetSilenceTime() *int32
	SetStartTime(v string) *CreateAutoScalingRuleRequest
	GetStartTime() *string
	SetTargetMetric(v string) *CreateAutoScalingRuleRequest
	GetTargetMetric() *string
	SetTargetNodes(v int32) *CreateAutoScalingRuleRequest
	GetTargetNodes() *int32
	SetThresholdLower(v int32) *CreateAutoScalingRuleRequest
	GetThresholdLower() *int32
	SetThresholdUpper(v int32) *CreateAutoScalingRuleRequest
	GetThresholdUpper() *int32
	SetTriggerCronExpr(v string) *CreateAutoScalingRuleRequest
	GetTriggerCronExpr() *string
}

type CreateAutoScalingRuleRequest struct {
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
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
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

func (s CreateAutoScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingRuleRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *CreateAutoScalingRuleRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAutoScalingRuleRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateAutoScalingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAutoScalingRuleRequest) GetObservationWindow() *int32 {
	return s.ObservationWindow
}

func (s *CreateAutoScalingRuleRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateAutoScalingRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAutoScalingRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAutoScalingRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAutoScalingRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAutoScalingRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateAutoScalingRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *CreateAutoScalingRuleRequest) GetScaleInStep() *int32 {
	return s.ScaleInStep
}

func (s *CreateAutoScalingRuleRequest) GetScaleOutStep() *int32 {
	return s.ScaleOutStep
}

func (s *CreateAutoScalingRuleRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateAutoScalingRuleRequest) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *CreateAutoScalingRuleRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateAutoScalingRuleRequest) GetTargetMetric() *string {
	return s.TargetMetric
}

func (s *CreateAutoScalingRuleRequest) GetTargetNodes() *int32 {
	return s.TargetNodes
}

func (s *CreateAutoScalingRuleRequest) GetThresholdLower() *int32 {
	return s.ThresholdLower
}

func (s *CreateAutoScalingRuleRequest) GetThresholdUpper() *int32 {
	return s.ThresholdUpper
}

func (s *CreateAutoScalingRuleRequest) GetTriggerCronExpr() *string {
	return s.TriggerCronExpr
}

func (s *CreateAutoScalingRuleRequest) SetConfigId(v string) *CreateAutoScalingRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetEnabled(v bool) *CreateAutoScalingRuleRequest {
	s.Enabled = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetEndTime(v string) *CreateAutoScalingRuleRequest {
	s.EndTime = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetInstanceId(v string) *CreateAutoScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetObservationWindow(v int32) *CreateAutoScalingRuleRequest {
	s.ObservationWindow = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetOperationType(v string) *CreateAutoScalingRuleRequest {
	s.OperationType = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetOwnerAccount(v string) *CreateAutoScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetOwnerId(v int64) *CreateAutoScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetResourceOwnerAccount(v string) *CreateAutoScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetResourceOwnerId(v int64) *CreateAutoScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetRuleName(v string) *CreateAutoScalingRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetRuleType(v string) *CreateAutoScalingRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetScaleInStep(v int32) *CreateAutoScalingRuleRequest {
	s.ScaleInStep = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetScaleOutStep(v int32) *CreateAutoScalingRuleRequest {
	s.ScaleOutStep = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetSecurityToken(v string) *CreateAutoScalingRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetSilenceTime(v int32) *CreateAutoScalingRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetStartTime(v string) *CreateAutoScalingRuleRequest {
	s.StartTime = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetTargetMetric(v string) *CreateAutoScalingRuleRequest {
	s.TargetMetric = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetTargetNodes(v int32) *CreateAutoScalingRuleRequest {
	s.TargetNodes = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetThresholdLower(v int32) *CreateAutoScalingRuleRequest {
	s.ThresholdLower = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetThresholdUpper(v int32) *CreateAutoScalingRuleRequest {
	s.ThresholdUpper = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetTriggerCronExpr(v string) *CreateAutoScalingRuleRequest {
	s.TriggerCronExpr = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
