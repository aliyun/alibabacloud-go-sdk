// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAutoScalingRulesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListAutoScalingRulesResponseBody
	GetCode() *string
	SetData(v *ListAutoScalingRulesResponseBodyData) *ListAutoScalingRulesResponseBody
	GetData() *ListAutoScalingRulesResponseBodyData
	SetDynamicCode(v string) *ListAutoScalingRulesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAutoScalingRulesResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListAutoScalingRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAutoScalingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAutoScalingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAutoScalingRulesResponseBody
	GetSuccess() *bool
}

type ListAutoScalingRulesResponseBody struct {
	AccessDeniedDetail *string                               `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *ListAutoScalingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode        *string                               `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAutoScalingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRulesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAutoScalingRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAutoScalingRulesResponseBody) GetData() *ListAutoScalingRulesResponseBodyData {
	return s.Data
}

func (s *ListAutoScalingRulesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAutoScalingRulesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAutoScalingRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAutoScalingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAutoScalingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAutoScalingRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAutoScalingRulesResponseBody) SetAccessDeniedDetail(v string) *ListAutoScalingRulesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetCode(v string) *ListAutoScalingRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetData(v *ListAutoScalingRulesResponseBodyData) *ListAutoScalingRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetDynamicCode(v string) *ListAutoScalingRulesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetDynamicMessage(v string) *ListAutoScalingRulesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetHttpStatusCode(v int32) *ListAutoScalingRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetMessage(v string) *ListAutoScalingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetRequestId(v string) *ListAutoScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetSuccess(v bool) *ListAutoScalingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAutoScalingRulesResponseBodyData struct {
	ScaleRules []*ListAutoScalingRulesResponseBodyDataScaleRules `json:"ScaleRules,omitempty" xml:"ScaleRules,omitempty" type:"Repeated"`
}

func (s ListAutoScalingRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRulesResponseBodyData) GetScaleRules() []*ListAutoScalingRulesResponseBodyDataScaleRules {
	return s.ScaleRules
}

func (s *ListAutoScalingRulesResponseBodyData) SetScaleRules(v []*ListAutoScalingRulesResponseBodyDataScaleRules) *ListAutoScalingRulesResponseBodyData {
	s.ScaleRules = v
	return s
}

func (s *ListAutoScalingRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAutoScalingRulesResponseBodyDataScaleRules struct {
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

func (s ListAutoScalingRulesResponseBodyDataScaleRules) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingRulesResponseBodyDataScaleRules) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetConfigId() *string {
	return s.ConfigId
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetObservationWindow() *int32 {
	return s.ObservationWindow
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetOperationType() *string {
	return s.OperationType
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetRuleId() *string {
	return s.RuleId
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetRuleType() *string {
	return s.RuleType
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetScaleInStep() *int32 {
	return s.ScaleInStep
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetScaleOutStep() *int32 {
	return s.ScaleOutStep
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetTargetMetric() *string {
	return s.TargetMetric
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetTargetNodes() *int32 {
	return s.TargetNodes
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetThresholdLower() *int32 {
	return s.ThresholdLower
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetThresholdUpper() *int32 {
	return s.ThresholdUpper
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) GetTriggerCronExpr() *string {
	return s.TriggerCronExpr
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetConfigId(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ConfigId = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetEnabled(v bool) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.Enabled = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetEndTime(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.EndTime = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetInstanceId(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetObservationWindow(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ObservationWindow = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetOperationType(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.OperationType = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetRuleId(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.RuleId = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetRuleName(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.RuleName = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetRuleType(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.RuleType = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetScaleInStep(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ScaleInStep = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetScaleOutStep(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ScaleOutStep = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetSilenceTime(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.SilenceTime = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetStartTime(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.StartTime = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetTargetMetric(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.TargetMetric = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetTargetNodes(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.TargetNodes = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetThresholdLower(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ThresholdLower = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetThresholdUpper(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ThresholdUpper = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetTriggerCronExpr(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.TriggerCronExpr = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) Validate() error {
	return dara.Validate(s)
}
