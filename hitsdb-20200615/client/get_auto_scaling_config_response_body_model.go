// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAutoScalingConfigResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetAutoScalingConfigResponseBody
	GetCode() *string
	SetData(v *GetAutoScalingConfigResponseBodyData) *GetAutoScalingConfigResponseBody
	GetData() *GetAutoScalingConfigResponseBodyData
	SetDynamicCode(v string) *GetAutoScalingConfigResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAutoScalingConfigResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetAutoScalingConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAutoScalingConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAutoScalingConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAutoScalingConfigResponseBody
	GetSuccess() *bool
}

type GetAutoScalingConfigResponseBody struct {
	AccessDeniedDetail *string                               `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *GetAutoScalingConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode        *string                               `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutoScalingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAutoScalingConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAutoScalingConfigResponseBody) GetData() *GetAutoScalingConfigResponseBodyData {
	return s.Data
}

func (s *GetAutoScalingConfigResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAutoScalingConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAutoScalingConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAutoScalingConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAutoScalingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoScalingConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAutoScalingConfigResponseBody) SetAccessDeniedDetail(v string) *GetAutoScalingConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetCode(v string) *GetAutoScalingConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetData(v *GetAutoScalingConfigResponseBodyData) *GetAutoScalingConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetDynamicCode(v string) *GetAutoScalingConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetDynamicMessage(v string) *GetAutoScalingConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetHttpStatusCode(v int32) *GetAutoScalingConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetMessage(v string) *GetAutoScalingConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetRequestId(v string) *GetAutoScalingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetSuccess(v bool) *GetAutoScalingConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoScalingConfigResponseBodyData struct {
	ConfigId           *string                                              `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigName         *string                                              `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	EffectiveTimeEnd   *string                                              `json:"EffectiveTimeEnd,omitempty" xml:"EffectiveTimeEnd,omitempty"`
	EffectiveTimeStart *string                                              `json:"EffectiveTimeStart,omitempty" xml:"EffectiveTimeStart,omitempty"`
	Enabled            *bool                                                `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Engine             *string                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceId         *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodesMax           *int32                                               `json:"NodesMax,omitempty" xml:"NodesMax,omitempty"`
	NodesMin           *int32                                               `json:"NodesMin,omitempty" xml:"NodesMin,omitempty"`
	ScaleRuleList      []*GetAutoScalingConfigResponseBodyDataScaleRuleList `json:"ScaleRuleList,omitempty" xml:"ScaleRuleList,omitempty" type:"Repeated"`
	ScaleType          *string                                              `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SpecId             *string                                              `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s GetAutoScalingConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigResponseBodyData) GetConfigId() *string {
	return s.ConfigId
}

func (s *GetAutoScalingConfigResponseBodyData) GetConfigName() *string {
	return s.ConfigName
}

func (s *GetAutoScalingConfigResponseBodyData) GetEffectiveTimeEnd() *string {
	return s.EffectiveTimeEnd
}

func (s *GetAutoScalingConfigResponseBodyData) GetEffectiveTimeStart() *string {
	return s.EffectiveTimeStart
}

func (s *GetAutoScalingConfigResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetAutoScalingConfigResponseBodyData) GetEngine() *string {
	return s.Engine
}

func (s *GetAutoScalingConfigResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoScalingConfigResponseBodyData) GetNodesMax() *int32 {
	return s.NodesMax
}

func (s *GetAutoScalingConfigResponseBodyData) GetNodesMin() *int32 {
	return s.NodesMin
}

func (s *GetAutoScalingConfigResponseBodyData) GetScaleRuleList() []*GetAutoScalingConfigResponseBodyDataScaleRuleList {
	return s.ScaleRuleList
}

func (s *GetAutoScalingConfigResponseBodyData) GetScaleType() *string {
	return s.ScaleType
}

func (s *GetAutoScalingConfigResponseBodyData) GetSpecId() *string {
	return s.SpecId
}

func (s *GetAutoScalingConfigResponseBodyData) SetConfigId(v string) *GetAutoScalingConfigResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetConfigName(v string) *GetAutoScalingConfigResponseBodyData {
	s.ConfigName = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetEffectiveTimeEnd(v string) *GetAutoScalingConfigResponseBodyData {
	s.EffectiveTimeEnd = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetEffectiveTimeStart(v string) *GetAutoScalingConfigResponseBodyData {
	s.EffectiveTimeStart = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetEnabled(v bool) *GetAutoScalingConfigResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetEngine(v string) *GetAutoScalingConfigResponseBodyData {
	s.Engine = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetInstanceId(v string) *GetAutoScalingConfigResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetNodesMax(v int32) *GetAutoScalingConfigResponseBodyData {
	s.NodesMax = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetNodesMin(v int32) *GetAutoScalingConfigResponseBodyData {
	s.NodesMin = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetScaleRuleList(v []*GetAutoScalingConfigResponseBodyDataScaleRuleList) *GetAutoScalingConfigResponseBodyData {
	s.ScaleRuleList = v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetScaleType(v string) *GetAutoScalingConfigResponseBodyData {
	s.ScaleType = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetSpecId(v string) *GetAutoScalingConfigResponseBodyData {
	s.SpecId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) Validate() error {
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

type GetAutoScalingConfigResponseBodyDataScaleRuleList struct {
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

func (s GetAutoScalingConfigResponseBodyDataScaleRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigResponseBodyDataScaleRuleList) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetConfigId() *string {
	return s.ConfigId
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetObservationWindow() *int32 {
	return s.ObservationWindow
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetOperationType() *string {
	return s.OperationType
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetRuleId() *string {
	return s.RuleId
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetScaleInStep() *int32 {
	return s.ScaleInStep
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetScaleOutStep() *int32 {
	return s.ScaleOutStep
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetTargetMetric() *string {
	return s.TargetMetric
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetTargetNodes() *int32 {
	return s.TargetNodes
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetThresholdLower() *int32 {
	return s.ThresholdLower
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetThresholdUpper() *int32 {
	return s.ThresholdUpper
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) GetTriggerCronExpr() *string {
	return s.TriggerCronExpr
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetConfigId(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ConfigId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetEnabled(v bool) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.Enabled = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetEndTime(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.EndTime = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetInstanceId(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetObservationWindow(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ObservationWindow = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetOperationType(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.OperationType = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetRuleId(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.RuleId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetRuleName(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.RuleName = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetRuleType(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.RuleType = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetScaleInStep(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ScaleInStep = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetScaleOutStep(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ScaleOutStep = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetSilenceTime(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.SilenceTime = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetStartTime(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.StartTime = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetTargetMetric(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.TargetMetric = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetTargetNodes(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.TargetNodes = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetThresholdLower(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ThresholdLower = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetThresholdUpper(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ThresholdUpper = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetTriggerCronExpr(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.TriggerCronExpr = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) Validate() error {
	return dara.Validate(s)
}
