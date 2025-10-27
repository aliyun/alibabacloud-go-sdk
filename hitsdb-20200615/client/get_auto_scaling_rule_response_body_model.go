// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAutoScalingRuleResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetAutoScalingRuleResponseBody
	GetCode() *string
	SetData(v *GetAutoScalingRuleResponseBodyData) *GetAutoScalingRuleResponseBody
	GetData() *GetAutoScalingRuleResponseBodyData
	SetDynamicCode(v string) *GetAutoScalingRuleResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAutoScalingRuleResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetAutoScalingRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAutoScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAutoScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAutoScalingRuleResponseBody
	GetSuccess() *bool
}

type GetAutoScalingRuleResponseBody struct {
	AccessDeniedDetail *string                             `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *GetAutoScalingRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode        *string                             `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                             `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutoScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScalingRuleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAutoScalingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAutoScalingRuleResponseBody) GetData() *GetAutoScalingRuleResponseBodyData {
	return s.Data
}

func (s *GetAutoScalingRuleResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAutoScalingRuleResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAutoScalingRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAutoScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAutoScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAutoScalingRuleResponseBody) SetAccessDeniedDetail(v string) *GetAutoScalingRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetCode(v string) *GetAutoScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetData(v *GetAutoScalingRuleResponseBodyData) *GetAutoScalingRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetDynamicCode(v string) *GetAutoScalingRuleResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetDynamicMessage(v string) *GetAutoScalingRuleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetHttpStatusCode(v int32) *GetAutoScalingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetMessage(v string) *GetAutoScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetRequestId(v string) *GetAutoScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetSuccess(v bool) *GetAutoScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoScalingRuleResponseBodyData struct {
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

func (s GetAutoScalingRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutoScalingRuleResponseBodyData) GetConfigId() *string {
	return s.ConfigId
}

func (s *GetAutoScalingRuleResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetAutoScalingRuleResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAutoScalingRuleResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoScalingRuleResponseBodyData) GetObservationWindow() *int32 {
	return s.ObservationWindow
}

func (s *GetAutoScalingRuleResponseBodyData) GetOperationType() *string {
	return s.OperationType
}

func (s *GetAutoScalingRuleResponseBodyData) GetRuleId() *string {
	return s.RuleId
}

func (s *GetAutoScalingRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *GetAutoScalingRuleResponseBodyData) GetRuleType() *string {
	return s.RuleType
}

func (s *GetAutoScalingRuleResponseBodyData) GetScaleInStep() *int32 {
	return s.ScaleInStep
}

func (s *GetAutoScalingRuleResponseBodyData) GetScaleOutStep() *int32 {
	return s.ScaleOutStep
}

func (s *GetAutoScalingRuleResponseBodyData) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *GetAutoScalingRuleResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAutoScalingRuleResponseBodyData) GetTargetMetric() *string {
	return s.TargetMetric
}

func (s *GetAutoScalingRuleResponseBodyData) GetTargetNodes() *int32 {
	return s.TargetNodes
}

func (s *GetAutoScalingRuleResponseBodyData) GetThresholdLower() *int32 {
	return s.ThresholdLower
}

func (s *GetAutoScalingRuleResponseBodyData) GetThresholdUpper() *int32 {
	return s.ThresholdUpper
}

func (s *GetAutoScalingRuleResponseBodyData) GetTriggerCronExpr() *string {
	return s.TriggerCronExpr
}

func (s *GetAutoScalingRuleResponseBodyData) SetConfigId(v string) *GetAutoScalingRuleResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetEnabled(v bool) *GetAutoScalingRuleResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetEndTime(v string) *GetAutoScalingRuleResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetInstanceId(v string) *GetAutoScalingRuleResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetObservationWindow(v int32) *GetAutoScalingRuleResponseBodyData {
	s.ObservationWindow = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetOperationType(v string) *GetAutoScalingRuleResponseBodyData {
	s.OperationType = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetRuleId(v string) *GetAutoScalingRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetRuleName(v string) *GetAutoScalingRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetRuleType(v string) *GetAutoScalingRuleResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetScaleInStep(v int32) *GetAutoScalingRuleResponseBodyData {
	s.ScaleInStep = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetScaleOutStep(v int32) *GetAutoScalingRuleResponseBodyData {
	s.ScaleOutStep = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetSilenceTime(v int32) *GetAutoScalingRuleResponseBodyData {
	s.SilenceTime = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetStartTime(v string) *GetAutoScalingRuleResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetTargetMetric(v string) *GetAutoScalingRuleResponseBodyData {
	s.TargetMetric = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetTargetNodes(v int32) *GetAutoScalingRuleResponseBodyData {
	s.TargetNodes = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetThresholdLower(v int32) *GetAutoScalingRuleResponseBodyData {
	s.ThresholdLower = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetThresholdUpper(v int32) *GetAutoScalingRuleResponseBodyData {
	s.ThresholdUpper = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetTriggerCronExpr(v string) *GetAutoScalingRuleResponseBodyData {
	s.TriggerCronExpr = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
