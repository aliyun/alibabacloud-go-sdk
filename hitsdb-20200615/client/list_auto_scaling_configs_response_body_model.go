// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAutoScalingConfigsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListAutoScalingConfigsResponseBody
	GetCode() *string
	SetData(v *ListAutoScalingConfigsResponseBodyData) *ListAutoScalingConfigsResponseBody
	GetData() *ListAutoScalingConfigsResponseBodyData
	SetDynamicCode(v string) *ListAutoScalingConfigsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAutoScalingConfigsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListAutoScalingConfigsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAutoScalingConfigsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAutoScalingConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAutoScalingConfigsResponseBody
	GetSuccess() *bool
}

type ListAutoScalingConfigsResponseBody struct {
	AccessDeniedDetail *string                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *ListAutoScalingConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode        *string                                 `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAutoScalingConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAutoScalingConfigsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAutoScalingConfigsResponseBody) GetData() *ListAutoScalingConfigsResponseBodyData {
	return s.Data
}

func (s *ListAutoScalingConfigsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAutoScalingConfigsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAutoScalingConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAutoScalingConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAutoScalingConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAutoScalingConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAutoScalingConfigsResponseBody) SetAccessDeniedDetail(v string) *ListAutoScalingConfigsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetCode(v string) *ListAutoScalingConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetData(v *ListAutoScalingConfigsResponseBodyData) *ListAutoScalingConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetDynamicCode(v string) *ListAutoScalingConfigsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetDynamicMessage(v string) *ListAutoScalingConfigsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetHttpStatusCode(v int32) *ListAutoScalingConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetMessage(v string) *ListAutoScalingConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetRequestId(v string) *ListAutoScalingConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetSuccess(v bool) *ListAutoScalingConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAutoScalingConfigsResponseBodyData struct {
	ScaleConfigs []*ListAutoScalingConfigsResponseBodyDataScaleConfigs `json:"ScaleConfigs,omitempty" xml:"ScaleConfigs,omitempty" type:"Repeated"`
}

func (s ListAutoScalingConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsResponseBodyData) GetScaleConfigs() []*ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	return s.ScaleConfigs
}

func (s *ListAutoScalingConfigsResponseBodyData) SetScaleConfigs(v []*ListAutoScalingConfigsResponseBodyDataScaleConfigs) *ListAutoScalingConfigsResponseBodyData {
	s.ScaleConfigs = v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyData) Validate() error {
	if s.ScaleConfigs != nil {
		for _, item := range s.ScaleConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAutoScalingConfigsResponseBodyDataScaleConfigs struct {
	ConfigId           *string                                                            `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigName         *string                                                            `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	EffectiveTimeEnd   *string                                                            `json:"EffectiveTimeEnd,omitempty" xml:"EffectiveTimeEnd,omitempty"`
	EffectiveTimeStart *string                                                            `json:"EffectiveTimeStart,omitempty" xml:"EffectiveTimeStart,omitempty"`
	Enabled            *bool                                                              `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Engine             *string                                                            `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceId         *string                                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodesMax           *int32                                                             `json:"NodesMax,omitempty" xml:"NodesMax,omitempty"`
	NodesMin           *int32                                                             `json:"NodesMin,omitempty" xml:"NodesMin,omitempty"`
	ScaleRuleList      []*ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList `json:"ScaleRuleList,omitempty" xml:"ScaleRuleList,omitempty" type:"Repeated"`
	ScaleType          *string                                                            `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SpecId             *string                                                            `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s ListAutoScalingConfigsResponseBodyDataScaleConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingConfigsResponseBodyDataScaleConfigs) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetConfigId() *string {
	return s.ConfigId
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetConfigName() *string {
	return s.ConfigName
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetEffectiveTimeEnd() *string {
	return s.EffectiveTimeEnd
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetEffectiveTimeStart() *string {
	return s.EffectiveTimeStart
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetEngine() *string {
	return s.Engine
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetNodesMax() *int32 {
	return s.NodesMax
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetNodesMin() *int32 {
	return s.NodesMin
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetScaleRuleList() []*ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	return s.ScaleRuleList
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetScaleType() *string {
	return s.ScaleType
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) GetSpecId() *string {
	return s.SpecId
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetConfigId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetConfigName(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.ConfigName = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetEffectiveTimeEnd(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.EffectiveTimeEnd = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetEffectiveTimeStart(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.EffectiveTimeStart = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetEnabled(v bool) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.Enabled = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetEngine(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.Engine = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetInstanceId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetNodesMax(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.NodesMax = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetNodesMin(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.NodesMin = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetScaleRuleList(v []*ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.ScaleRuleList = v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetScaleType(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.ScaleType = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetSpecId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.SpecId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) Validate() error {
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

type ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList struct {
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

func (s ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetConfigId() *string {
	return s.ConfigId
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetObservationWindow() *int32 {
	return s.ObservationWindow
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetOperationType() *string {
	return s.OperationType
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetRuleId() *string {
	return s.RuleId
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetScaleInStep() *int32 {
	return s.ScaleInStep
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetScaleOutStep() *int32 {
	return s.ScaleOutStep
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetTargetMetric() *string {
	return s.TargetMetric
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetTargetNodes() *int32 {
	return s.TargetNodes
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetThresholdLower() *int32 {
	return s.ThresholdLower
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetThresholdUpper() *int32 {
	return s.ThresholdUpper
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GetTriggerCronExpr() *string {
	return s.TriggerCronExpr
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetConfigId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ConfigId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetEnabled(v bool) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.Enabled = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetEndTime(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.EndTime = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetInstanceId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetObservationWindow(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ObservationWindow = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetOperationType(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.OperationType = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetRuleId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.RuleId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetRuleName(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.RuleName = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetRuleType(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.RuleType = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetScaleInStep(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ScaleInStep = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetScaleOutStep(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ScaleOutStep = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetSilenceTime(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.SilenceTime = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetStartTime(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.StartTime = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetTargetMetric(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.TargetMetric = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetTargetNodes(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.TargetNodes = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetThresholdLower(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ThresholdLower = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetThresholdUpper(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ThresholdUpper = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetTriggerCronExpr(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.TriggerCronExpr = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) Validate() error {
	return dara.Validate(s)
}
