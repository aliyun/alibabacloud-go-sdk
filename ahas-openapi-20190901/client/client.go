// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckExperimentRunnableRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	NameSpace    *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
}

func (s CheckExperimentRunnableRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckExperimentRunnableRequest) GoString() string {
	return s.String()
}

func (s *CheckExperimentRunnableRequest) SetAhasRegionId(v string) *CheckExperimentRunnableRequest {
	s.AhasRegionId = &v
	return s
}

func (s *CheckExperimentRunnableRequest) SetExperimentId(v string) *CheckExperimentRunnableRequest {
	s.ExperimentId = &v
	return s
}

func (s *CheckExperimentRunnableRequest) SetNameSpace(v string) *CheckExperimentRunnableRequest {
	s.NameSpace = &v
	return s
}

type CheckExperimentRunnableResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckExperimentRunnableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckExperimentRunnableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckExperimentRunnableResponseBody) SetCode(v string) *CheckExperimentRunnableResponseBody {
	s.Code = &v
	return s
}

func (s *CheckExperimentRunnableResponseBody) SetMessage(v string) *CheckExperimentRunnableResponseBody {
	s.Message = &v
	return s
}

func (s *CheckExperimentRunnableResponseBody) SetRequestId(v string) *CheckExperimentRunnableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckExperimentRunnableResponseBody) SetSuccess(v bool) *CheckExperimentRunnableResponseBody {
	s.Success = &v
	return s
}

type CheckExperimentRunnableResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckExperimentRunnableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckExperimentRunnableResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckExperimentRunnableResponse) GoString() string {
	return s.String()
}

func (s *CheckExperimentRunnableResponse) SetHeaders(v map[string]*string) *CheckExperimentRunnableResponse {
	s.Headers = v
	return s
}

func (s *CheckExperimentRunnableResponse) SetStatusCode(v int32) *CheckExperimentRunnableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckExperimentRunnableResponse) SetBody(v *CheckExperimentRunnableResponseBody) *CheckExperimentRunnableResponse {
	s.Body = v
	return s
}

type CreateDegradeRuleRequest struct {
	AhasRegionId              *string  `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName                   *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable                    *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HalfOpenBaseAmountPerStep *int32   `json:"HalfOpenBaseAmountPerStep,omitempty" xml:"HalfOpenBaseAmountPerStep,omitempty"`
	HalfOpenRecoveryStepNum   *int32   `json:"HalfOpenRecoveryStepNum,omitempty" xml:"HalfOpenRecoveryStepNum,omitempty"`
	MinRequestAmount          *int32   `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	Namespace                 *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RecoveryTimeoutMs         *int32   `json:"RecoveryTimeoutMs,omitempty" xml:"RecoveryTimeoutMs,omitempty"`
	Resource                  *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	SlowRtMs                  *int32   `json:"SlowRtMs,omitempty" xml:"SlowRtMs,omitempty"`
	StatDurationMs            *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Strategy                  *int32   `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Threshold                 *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateDegradeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDegradeRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDegradeRuleRequest) SetAhasRegionId(v string) *CreateDegradeRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetAppName(v string) *CreateDegradeRuleRequest {
	s.AppName = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetEnable(v bool) *CreateDegradeRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetHalfOpenBaseAmountPerStep(v int32) *CreateDegradeRuleRequest {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetHalfOpenRecoveryStepNum(v int32) *CreateDegradeRuleRequest {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetMinRequestAmount(v int32) *CreateDegradeRuleRequest {
	s.MinRequestAmount = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetNamespace(v string) *CreateDegradeRuleRequest {
	s.Namespace = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetRecoveryTimeoutMs(v int32) *CreateDegradeRuleRequest {
	s.RecoveryTimeoutMs = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetResource(v string) *CreateDegradeRuleRequest {
	s.Resource = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetSlowRtMs(v int32) *CreateDegradeRuleRequest {
	s.SlowRtMs = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetStatDurationMs(v int32) *CreateDegradeRuleRequest {
	s.StatDurationMs = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetStrategy(v int32) *CreateDegradeRuleRequest {
	s.Strategy = &v
	return s
}

func (s *CreateDegradeRuleRequest) SetThreshold(v float32) *CreateDegradeRuleRequest {
	s.Threshold = &v
	return s
}

type CreateDegradeRuleResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateDegradeRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDegradeRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDegradeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDegradeRuleResponseBody) SetCode(v string) *CreateDegradeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDegradeRuleResponseBody) SetData(v *CreateDegradeRuleResponseBodyData) *CreateDegradeRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateDegradeRuleResponseBody) SetMessage(v string) *CreateDegradeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDegradeRuleResponseBody) SetRequestId(v string) *CreateDegradeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDegradeRuleResponseBody) SetSuccess(v bool) *CreateDegradeRuleResponseBody {
	s.Success = &v
	return s
}

type CreateDegradeRuleResponseBodyData struct {
	AppName                   *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable                    *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HalfOpenBaseAmountPerStep *int32   `json:"HalfOpenBaseAmountPerStep,omitempty" xml:"HalfOpenBaseAmountPerStep,omitempty"`
	HalfOpenRecoveryStepNum   *int32   `json:"HalfOpenRecoveryStepNum,omitempty" xml:"HalfOpenRecoveryStepNum,omitempty"`
	MinRequestAmount          *int32   `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	Namespace                 *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RecoveryTimeoutMs         *int32   `json:"RecoveryTimeoutMs,omitempty" xml:"RecoveryTimeoutMs,omitempty"`
	Resource                  *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                    *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SlowRtMs                  *int32   `json:"SlowRtMs,omitempty" xml:"SlowRtMs,omitempty"`
	StatDurationMs            *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Strategy                  *int32   `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Threshold                 *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateDegradeRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDegradeRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDegradeRuleResponseBodyData) SetAppName(v string) *CreateDegradeRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetEnable(v bool) *CreateDegradeRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetHalfOpenBaseAmountPerStep(v int32) *CreateDegradeRuleResponseBodyData {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetHalfOpenRecoveryStepNum(v int32) *CreateDegradeRuleResponseBodyData {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetMinRequestAmount(v int32) *CreateDegradeRuleResponseBodyData {
	s.MinRequestAmount = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetNamespace(v string) *CreateDegradeRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetRecoveryTimeoutMs(v int32) *CreateDegradeRuleResponseBodyData {
	s.RecoveryTimeoutMs = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetResource(v string) *CreateDegradeRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetRuleId(v int64) *CreateDegradeRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetSlowRtMs(v int32) *CreateDegradeRuleResponseBodyData {
	s.SlowRtMs = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetStatDurationMs(v int32) *CreateDegradeRuleResponseBodyData {
	s.StatDurationMs = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetStrategy(v int32) *CreateDegradeRuleResponseBodyData {
	s.Strategy = &v
	return s
}

func (s *CreateDegradeRuleResponseBodyData) SetThreshold(v float32) *CreateDegradeRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type CreateDegradeRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDegradeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDegradeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDegradeRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDegradeRuleResponse) SetHeaders(v map[string]*string) *CreateDegradeRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDegradeRuleResponse) SetStatusCode(v int32) *CreateDegradeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDegradeRuleResponse) SetBody(v *CreateDegradeRuleResponseBody) *CreateDegradeRuleResponse {
	s.Body = v
	return s
}

type CreateExperimentRequest struct {
	AhasRegionId *string   `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	Definition   *string   `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description  *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NameSpace    *string   `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	Tags         []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentRequest) SetAhasRegionId(v string) *CreateExperimentRequest {
	s.AhasRegionId = &v
	return s
}

func (s *CreateExperimentRequest) SetDefinition(v string) *CreateExperimentRequest {
	s.Definition = &v
	return s
}

func (s *CreateExperimentRequest) SetDescription(v string) *CreateExperimentRequest {
	s.Description = &v
	return s
}

func (s *CreateExperimentRequest) SetName(v string) *CreateExperimentRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentRequest) SetNameSpace(v string) *CreateExperimentRequest {
	s.NameSpace = &v
	return s
}

func (s *CreateExperimentRequest) SetTags(v []*string) *CreateExperimentRequest {
	s.Tags = v
	return s
}

type CreateExperimentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ExperimentId   *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentResponseBody) SetCode(v string) *CreateExperimentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExperimentResponseBody) SetExperimentId(v string) *CreateExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *CreateExperimentResponseBody) SetHttpStatusCode(v int32) *CreateExperimentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateExperimentResponseBody) SetMessage(v string) *CreateExperimentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExperimentResponseBody) SetRequestId(v string) *CreateExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExperimentResponseBody) SetSuccess(v bool) *CreateExperimentResponseBody {
	s.Success = &v
	return s
}

type CreateExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentResponse) SetHeaders(v map[string]*string) *CreateExperimentResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentResponse) SetStatusCode(v int32) *CreateExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentResponse) SetBody(v *CreateExperimentResponseBody) *CreateExperimentResponse {
	s.Body = v
	return s
}

type CreateFlowRuleRequest struct {
	AhasRegionId      *string  `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName           *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ControlBehavior   *int32   `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable            *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin       *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	MaxQueueingTimeMs *int32   `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	Namespace         *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource       *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy  *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource          *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	Threshold         *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WarmUpPeriodSec   *int32   `json:"WarmUpPeriodSec,omitempty" xml:"WarmUpPeriodSec,omitempty"`
}

func (s CreateFlowRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRuleRequest) SetAhasRegionId(v string) *CreateFlowRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *CreateFlowRuleRequest) SetAppName(v string) *CreateFlowRuleRequest {
	s.AppName = &v
	return s
}

func (s *CreateFlowRuleRequest) SetControlBehavior(v int32) *CreateFlowRuleRequest {
	s.ControlBehavior = &v
	return s
}

func (s *CreateFlowRuleRequest) SetEnable(v bool) *CreateFlowRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateFlowRuleRequest) SetLimitOrigin(v string) *CreateFlowRuleRequest {
	s.LimitOrigin = &v
	return s
}

func (s *CreateFlowRuleRequest) SetMaxQueueingTimeMs(v int32) *CreateFlowRuleRequest {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *CreateFlowRuleRequest) SetNamespace(v string) *CreateFlowRuleRequest {
	s.Namespace = &v
	return s
}

func (s *CreateFlowRuleRequest) SetRefResource(v string) *CreateFlowRuleRequest {
	s.RefResource = &v
	return s
}

func (s *CreateFlowRuleRequest) SetRelationStrategy(v int32) *CreateFlowRuleRequest {
	s.RelationStrategy = &v
	return s
}

func (s *CreateFlowRuleRequest) SetResource(v string) *CreateFlowRuleRequest {
	s.Resource = &v
	return s
}

func (s *CreateFlowRuleRequest) SetThreshold(v float32) *CreateFlowRuleRequest {
	s.Threshold = &v
	return s
}

func (s *CreateFlowRuleRequest) SetWarmUpPeriodSec(v int32) *CreateFlowRuleRequest {
	s.WarmUpPeriodSec = &v
	return s
}

type CreateFlowRuleResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateFlowRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFlowRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowRuleResponseBody) SetCode(v string) *CreateFlowRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFlowRuleResponseBody) SetData(v *CreateFlowRuleResponseBodyData) *CreateFlowRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateFlowRuleResponseBody) SetMessage(v string) *CreateFlowRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFlowRuleResponseBody) SetRequestId(v string) *CreateFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowRuleResponseBody) SetSuccess(v bool) *CreateFlowRuleResponseBody {
	s.Success = &v
	return s
}

type CreateFlowRuleResponseBodyData struct {
	AppName                  *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ClusterEstimatedMaxQps   *float32 `json:"ClusterEstimatedMaxQps,omitempty" xml:"ClusterEstimatedMaxQps,omitempty"`
	ClusterFallbackStrategy  *int32   `json:"ClusterFallbackStrategy,omitempty" xml:"ClusterFallbackStrategy,omitempty"`
	ClusterFallbackThreshold *int32   `json:"ClusterFallbackThreshold,omitempty" xml:"ClusterFallbackThreshold,omitempty"`
	ClusterMode              *bool    `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	ClusterThresholdType     *int32   `json:"ClusterThresholdType,omitempty" xml:"ClusterThresholdType,omitempty"`
	ControlBehavior          *int32   `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable                   *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin              *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	MaxQueueingTimeMs        *int32   `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	Namespace                *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource              *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy         *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource                 *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                   *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationMs           *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Threshold                *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WarmUpPeriodSec          *int32   `json:"WarmUpPeriodSec,omitempty" xml:"WarmUpPeriodSec,omitempty"`
}

func (s CreateFlowRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateFlowRuleResponseBodyData) SetAppName(v string) *CreateFlowRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetClusterEstimatedMaxQps(v float32) *CreateFlowRuleResponseBodyData {
	s.ClusterEstimatedMaxQps = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetClusterFallbackStrategy(v int32) *CreateFlowRuleResponseBodyData {
	s.ClusterFallbackStrategy = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetClusterFallbackThreshold(v int32) *CreateFlowRuleResponseBodyData {
	s.ClusterFallbackThreshold = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetClusterMode(v bool) *CreateFlowRuleResponseBodyData {
	s.ClusterMode = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetClusterThresholdType(v int32) *CreateFlowRuleResponseBodyData {
	s.ClusterThresholdType = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetControlBehavior(v int32) *CreateFlowRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetEnable(v bool) *CreateFlowRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetLimitOrigin(v string) *CreateFlowRuleResponseBodyData {
	s.LimitOrigin = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *CreateFlowRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetNamespace(v string) *CreateFlowRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetRefResource(v string) *CreateFlowRuleResponseBodyData {
	s.RefResource = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetRelationStrategy(v int32) *CreateFlowRuleResponseBodyData {
	s.RelationStrategy = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetResource(v string) *CreateFlowRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetRuleId(v int64) *CreateFlowRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetStatDurationMs(v int32) *CreateFlowRuleResponseBodyData {
	s.StatDurationMs = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetThreshold(v float32) *CreateFlowRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetWarmUpPeriodSec(v int32) *CreateFlowRuleResponseBodyData {
	s.WarmUpPeriodSec = &v
	return s
}

type CreateFlowRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowRuleResponse) SetHeaders(v map[string]*string) *CreateFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowRuleResponse) SetStatusCode(v int32) *CreateFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowRuleResponse) SetBody(v *CreateFlowRuleResponseBody) *CreateFlowRuleResponse {
	s.Body = v
	return s
}

type CreateHotParamItemsRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	Items        *string `json:"Items,omitempty" xml:"Items,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateHotParamItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHotParamItemsRequest) GoString() string {
	return s.String()
}

func (s *CreateHotParamItemsRequest) SetAhasRegionId(v string) *CreateHotParamItemsRequest {
	s.AhasRegionId = &v
	return s
}

func (s *CreateHotParamItemsRequest) SetItems(v string) *CreateHotParamItemsRequest {
	s.Items = &v
	return s
}

func (s *CreateHotParamItemsRequest) SetRuleId(v int64) *CreateHotParamItemsRequest {
	s.RuleId = &v
	return s
}

type CreateHotParamItemsResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateHotParamItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHotParamItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHotParamItemsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHotParamItemsResponseBody) SetCode(v string) *CreateHotParamItemsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHotParamItemsResponseBody) SetData(v *CreateHotParamItemsResponseBodyData) *CreateHotParamItemsResponseBody {
	s.Data = v
	return s
}

func (s *CreateHotParamItemsResponseBody) SetMessage(v string) *CreateHotParamItemsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHotParamItemsResponseBody) SetRequestId(v string) *CreateHotParamItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHotParamItemsResponseBody) SetSuccess(v bool) *CreateHotParamItemsResponseBody {
	s.Success = &v
	return s
}

type CreateHotParamItemsResponseBodyData struct {
	AppName           *string                                                 `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BurstCount        *int32                                                  `json:"BurstCount,omitempty" xml:"BurstCount,omitempty"`
	ControlBehavior   *int32                                                  `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable            *bool                                                   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MaxQueueingTimeMs *int32                                                  `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	MetricType        *int32                                                  `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Namespace         *string                                                 `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParamFlowItemList []*CreateHotParamItemsResponseBodyDataParamFlowItemList `json:"ParamFlowItemList,omitempty" xml:"ParamFlowItemList,omitempty" type:"Repeated"`
	ParamIdx          *int32                                                  `json:"ParamIdx,omitempty" xml:"ParamIdx,omitempty"`
	Resource          *string                                                 `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId            *int64                                                  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationSec   *int64                                                  `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	Threshold         *float32                                                `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateHotParamItemsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateHotParamItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateHotParamItemsResponseBodyData) SetAppName(v string) *CreateHotParamItemsResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetBurstCount(v int32) *CreateHotParamItemsResponseBodyData {
	s.BurstCount = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetControlBehavior(v int32) *CreateHotParamItemsResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetEnable(v bool) *CreateHotParamItemsResponseBodyData {
	s.Enable = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetMaxQueueingTimeMs(v int32) *CreateHotParamItemsResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetMetricType(v int32) *CreateHotParamItemsResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetNamespace(v string) *CreateHotParamItemsResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetParamFlowItemList(v []*CreateHotParamItemsResponseBodyDataParamFlowItemList) *CreateHotParamItemsResponseBodyData {
	s.ParamFlowItemList = v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetParamIdx(v int32) *CreateHotParamItemsResponseBodyData {
	s.ParamIdx = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetResource(v string) *CreateHotParamItemsResponseBodyData {
	s.Resource = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetRuleId(v int64) *CreateHotParamItemsResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetStatDurationSec(v int64) *CreateHotParamItemsResponseBodyData {
	s.StatDurationSec = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyData) SetThreshold(v float32) *CreateHotParamItemsResponseBodyData {
	s.Threshold = &v
	return s
}

type CreateHotParamItemsResponseBodyDataParamFlowItemList struct {
	ItemType  *string  `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	ItemValue *string  `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateHotParamItemsResponseBodyDataParamFlowItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateHotParamItemsResponseBodyDataParamFlowItemList) GoString() string {
	return s.String()
}

func (s *CreateHotParamItemsResponseBodyDataParamFlowItemList) SetItemType(v string) *CreateHotParamItemsResponseBodyDataParamFlowItemList {
	s.ItemType = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyDataParamFlowItemList) SetItemValue(v string) *CreateHotParamItemsResponseBodyDataParamFlowItemList {
	s.ItemValue = &v
	return s
}

func (s *CreateHotParamItemsResponseBodyDataParamFlowItemList) SetThreshold(v float32) *CreateHotParamItemsResponseBodyDataParamFlowItemList {
	s.Threshold = &v
	return s
}

type CreateHotParamItemsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHotParamItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHotParamItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHotParamItemsResponse) GoString() string {
	return s.String()
}

func (s *CreateHotParamItemsResponse) SetHeaders(v map[string]*string) *CreateHotParamItemsResponse {
	s.Headers = v
	return s
}

func (s *CreateHotParamItemsResponse) SetStatusCode(v int32) *CreateHotParamItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHotParamItemsResponse) SetBody(v *CreateHotParamItemsResponseBody) *CreateHotParamItemsResponse {
	s.Body = v
	return s
}

type CreateHotParamRuleRequest struct {
	AhasRegionId      *string  `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName           *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BurstCount        *int32   `json:"BurstCount,omitempty" xml:"BurstCount,omitempty"`
	ControlBehavior   *int32   `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable            *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MaxQueueingTimeMs *int32   `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	MetricType        *int32   `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Namespace         *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParamIdx          *int32   `json:"ParamIdx,omitempty" xml:"ParamIdx,omitempty"`
	Resource          *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	StatDurationSec   *int64   `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	Threshold         *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateHotParamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHotParamRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateHotParamRuleRequest) SetAhasRegionId(v string) *CreateHotParamRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *CreateHotParamRuleRequest) SetAppName(v string) *CreateHotParamRuleRequest {
	s.AppName = &v
	return s
}

func (s *CreateHotParamRuleRequest) SetBurstCount(v int32) *CreateHotParamRuleRequest {
	s.BurstCount = &v
	return s
}

func (s *CreateHotParamRuleRequest) SetControlBehavior(v int32) *CreateHotParamRuleRequest {
	s.ControlBehavior = &v
	return s
}

func (s *CreateHotParamRuleRequest) SetEnable(v bool) *CreateHotParamRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateHotParamRuleRequest) SetMaxQueueingTimeMs(v int32) *CreateHotParamRuleRequest {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *CreateHotParamRuleRequest) SetMetricType(v int32) *CreateHotParamRuleRequest {
	s.MetricType = &v
	return s
}

func (s *CreateHotParamRuleRequest) SetNamespace(v string) *CreateHotParamRuleRequest {
	s.Namespace = &v
	return s
}

func (s *CreateHotParamRuleRequest) SetParamIdx(v int32) *CreateHotParamRuleRequest {
	s.ParamIdx = &v
	return s
}

func (s *CreateHotParamRuleRequest) SetResource(v string) *CreateHotParamRuleRequest {
	s.Resource = &v
	return s
}

func (s *CreateHotParamRuleRequest) SetStatDurationSec(v int64) *CreateHotParamRuleRequest {
	s.StatDurationSec = &v
	return s
}

func (s *CreateHotParamRuleRequest) SetThreshold(v float32) *CreateHotParamRuleRequest {
	s.Threshold = &v
	return s
}

type CreateHotParamRuleResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateHotParamRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHotParamRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHotParamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHotParamRuleResponseBody) SetCode(v string) *CreateHotParamRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHotParamRuleResponseBody) SetData(v *CreateHotParamRuleResponseBodyData) *CreateHotParamRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateHotParamRuleResponseBody) SetMessage(v string) *CreateHotParamRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHotParamRuleResponseBody) SetRequestId(v string) *CreateHotParamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHotParamRuleResponseBody) SetSuccess(v bool) *CreateHotParamRuleResponseBody {
	s.Success = &v
	return s
}

type CreateHotParamRuleResponseBodyData struct {
	AppName           *string                                                `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BurstCount        *int32                                                 `json:"BurstCount,omitempty" xml:"BurstCount,omitempty"`
	ControlBehavior   *int32                                                 `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable            *bool                                                  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MaxQueueingTimeMs *int32                                                 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	MetricType        *int32                                                 `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Namespace         *string                                                `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParamFlowItemList []*CreateHotParamRuleResponseBodyDataParamFlowItemList `json:"ParamFlowItemList,omitempty" xml:"ParamFlowItemList,omitempty" type:"Repeated"`
	ParamIdx          *int32                                                 `json:"ParamIdx,omitempty" xml:"ParamIdx,omitempty"`
	Resource          *string                                                `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId            *int64                                                 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationSec   *int64                                                 `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	Threshold         *float32                                               `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateHotParamRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateHotParamRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateHotParamRuleResponseBodyData) SetAppName(v string) *CreateHotParamRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetBurstCount(v int32) *CreateHotParamRuleResponseBodyData {
	s.BurstCount = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetControlBehavior(v int32) *CreateHotParamRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetEnable(v bool) *CreateHotParamRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *CreateHotParamRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetMetricType(v int32) *CreateHotParamRuleResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetNamespace(v string) *CreateHotParamRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetParamFlowItemList(v []*CreateHotParamRuleResponseBodyDataParamFlowItemList) *CreateHotParamRuleResponseBodyData {
	s.ParamFlowItemList = v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetParamIdx(v int32) *CreateHotParamRuleResponseBodyData {
	s.ParamIdx = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetResource(v string) *CreateHotParamRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetRuleId(v int64) *CreateHotParamRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetStatDurationSec(v int64) *CreateHotParamRuleResponseBodyData {
	s.StatDurationSec = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyData) SetThreshold(v float32) *CreateHotParamRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type CreateHotParamRuleResponseBodyDataParamFlowItemList struct {
	ItemType  *string  `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	ItemValue *string  `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateHotParamRuleResponseBodyDataParamFlowItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateHotParamRuleResponseBodyDataParamFlowItemList) GoString() string {
	return s.String()
}

func (s *CreateHotParamRuleResponseBodyDataParamFlowItemList) SetItemType(v string) *CreateHotParamRuleResponseBodyDataParamFlowItemList {
	s.ItemType = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyDataParamFlowItemList) SetItemValue(v string) *CreateHotParamRuleResponseBodyDataParamFlowItemList {
	s.ItemValue = &v
	return s
}

func (s *CreateHotParamRuleResponseBodyDataParamFlowItemList) SetThreshold(v float32) *CreateHotParamRuleResponseBodyDataParamFlowItemList {
	s.Threshold = &v
	return s
}

type CreateHotParamRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHotParamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHotParamRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHotParamRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateHotParamRuleResponse) SetHeaders(v map[string]*string) *CreateHotParamRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateHotParamRuleResponse) SetStatusCode(v int32) *CreateHotParamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHotParamRuleResponse) SetBody(v *CreateHotParamRuleResponseBody) *CreateHotParamRuleResponse {
	s.Body = v
	return s
}

type CreateIsolationRuleRequest struct {
	AhasRegionId     *string  `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName          *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable           *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin      *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	Namespace        *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource      *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource         *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	Threshold        *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateIsolationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIsolationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateIsolationRuleRequest) SetAhasRegionId(v string) *CreateIsolationRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetAppName(v string) *CreateIsolationRuleRequest {
	s.AppName = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetEnable(v bool) *CreateIsolationRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetLimitOrigin(v string) *CreateIsolationRuleRequest {
	s.LimitOrigin = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetNamespace(v string) *CreateIsolationRuleRequest {
	s.Namespace = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetRefResource(v string) *CreateIsolationRuleRequest {
	s.RefResource = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetRelationStrategy(v int32) *CreateIsolationRuleRequest {
	s.RelationStrategy = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetResource(v string) *CreateIsolationRuleRequest {
	s.Resource = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetThreshold(v float32) *CreateIsolationRuleRequest {
	s.Threshold = &v
	return s
}

type CreateIsolationRuleResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateIsolationRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateIsolationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIsolationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIsolationRuleResponseBody) SetCode(v string) *CreateIsolationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIsolationRuleResponseBody) SetData(v *CreateIsolationRuleResponseBodyData) *CreateIsolationRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateIsolationRuleResponseBody) SetMessage(v string) *CreateIsolationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIsolationRuleResponseBody) SetRequestId(v string) *CreateIsolationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIsolationRuleResponseBody) SetSuccess(v bool) *CreateIsolationRuleResponseBody {
	s.Success = &v
	return s
}

type CreateIsolationRuleResponseBodyData struct {
	AppName          *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable           *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin      *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	Namespace        *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource      *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource         *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId           *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold        *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateIsolationRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateIsolationRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIsolationRuleResponseBodyData) SetAppName(v string) *CreateIsolationRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetEnable(v bool) *CreateIsolationRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetLimitOrigin(v string) *CreateIsolationRuleResponseBodyData {
	s.LimitOrigin = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetNamespace(v string) *CreateIsolationRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetRefResource(v string) *CreateIsolationRuleResponseBodyData {
	s.RefResource = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetRelationStrategy(v int32) *CreateIsolationRuleResponseBodyData {
	s.RelationStrategy = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetResource(v string) *CreateIsolationRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetRuleId(v int64) *CreateIsolationRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetThreshold(v float32) *CreateIsolationRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type CreateIsolationRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIsolationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIsolationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIsolationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateIsolationRuleResponse) SetHeaders(v map[string]*string) *CreateIsolationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateIsolationRuleResponse) SetStatusCode(v int32) *CreateIsolationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIsolationRuleResponse) SetBody(v *CreateIsolationRuleResponseBody) *CreateIsolationRuleResponse {
	s.Body = v
	return s
}

type CreateSystemRuleRequest struct {
	AhasRegionId *string  `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable       *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MetricType   *int32   `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Namespace    *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Threshold    *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateSystemRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSystemRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateSystemRuleRequest) SetAhasRegionId(v string) *CreateSystemRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *CreateSystemRuleRequest) SetAppName(v string) *CreateSystemRuleRequest {
	s.AppName = &v
	return s
}

func (s *CreateSystemRuleRequest) SetEnable(v bool) *CreateSystemRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateSystemRuleRequest) SetMetricType(v int32) *CreateSystemRuleRequest {
	s.MetricType = &v
	return s
}

func (s *CreateSystemRuleRequest) SetNamespace(v string) *CreateSystemRuleRequest {
	s.Namespace = &v
	return s
}

func (s *CreateSystemRuleRequest) SetThreshold(v float32) *CreateSystemRuleRequest {
	s.Threshold = &v
	return s
}

type CreateSystemRuleResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateSystemRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSystemRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSystemRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSystemRuleResponseBody) SetCode(v string) *CreateSystemRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSystemRuleResponseBody) SetData(v *CreateSystemRuleResponseBodyData) *CreateSystemRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateSystemRuleResponseBody) SetMessage(v string) *CreateSystemRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSystemRuleResponseBody) SetRequestId(v string) *CreateSystemRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSystemRuleResponseBody) SetSuccess(v bool) *CreateSystemRuleResponseBody {
	s.Success = &v
	return s
}

type CreateSystemRuleResponseBodyData struct {
	Enable     *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MetricType *int32   `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	RuleId     *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold  *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateSystemRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSystemRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSystemRuleResponseBodyData) SetEnable(v bool) *CreateSystemRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *CreateSystemRuleResponseBodyData) SetMetricType(v int32) *CreateSystemRuleResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *CreateSystemRuleResponseBodyData) SetRuleId(v int64) *CreateSystemRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *CreateSystemRuleResponseBodyData) SetThreshold(v float32) *CreateSystemRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type CreateSystemRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSystemRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSystemRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSystemRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateSystemRuleResponse) SetHeaders(v map[string]*string) *CreateSystemRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateSystemRuleResponse) SetStatusCode(v int32) *CreateSystemRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSystemRuleResponse) SetBody(v *CreateSystemRuleResponseBody) *CreateSystemRuleResponse {
	s.Body = v
	return s
}

type DeleteDegradeRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteDegradeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDegradeRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDegradeRuleRequest) SetAhasRegionId(v string) *DeleteDegradeRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *DeleteDegradeRuleRequest) SetRuleId(v int64) *DeleteDegradeRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteDegradeRuleResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteDegradeRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDegradeRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDegradeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDegradeRuleResponseBody) SetCode(v string) *DeleteDegradeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDegradeRuleResponseBody) SetData(v *DeleteDegradeRuleResponseBodyData) *DeleteDegradeRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDegradeRuleResponseBody) SetMessage(v string) *DeleteDegradeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDegradeRuleResponseBody) SetRequestId(v string) *DeleteDegradeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDegradeRuleResponseBody) SetSuccess(v bool) *DeleteDegradeRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteDegradeRuleResponseBodyData struct {
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteDegradeRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteDegradeRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDegradeRuleResponseBodyData) SetRuleId(v int64) *DeleteDegradeRuleResponseBodyData {
	s.RuleId = &v
	return s
}

type DeleteDegradeRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDegradeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDegradeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDegradeRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDegradeRuleResponse) SetHeaders(v map[string]*string) *DeleteDegradeRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDegradeRuleResponse) SetStatusCode(v int32) *DeleteDegradeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDegradeRuleResponse) SetBody(v *DeleteDegradeRuleResponseBody) *DeleteDegradeRuleResponse {
	s.Body = v
	return s
}

type DeleteFlowRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteFlowRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRuleRequest) SetAhasRegionId(v string) *DeleteFlowRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *DeleteFlowRuleRequest) SetRuleId(v int64) *DeleteFlowRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteFlowRuleResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteFlowRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFlowRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowRuleResponseBody) SetCode(v string) *DeleteFlowRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFlowRuleResponseBody) SetData(v *DeleteFlowRuleResponseBodyData) *DeleteFlowRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteFlowRuleResponseBody) SetMessage(v string) *DeleteFlowRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFlowRuleResponseBody) SetRequestId(v string) *DeleteFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowRuleResponseBody) SetSuccess(v bool) *DeleteFlowRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteFlowRuleResponseBodyData struct {
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteFlowRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteFlowRuleResponseBodyData) SetRuleId(v int64) *DeleteFlowRuleResponseBodyData {
	s.RuleId = &v
	return s
}

type DeleteFlowRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowRuleResponse) SetHeaders(v map[string]*string) *DeleteFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowRuleResponse) SetStatusCode(v int32) *DeleteFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowRuleResponse) SetBody(v *DeleteFlowRuleResponseBody) *DeleteFlowRuleResponse {
	s.Body = v
	return s
}

type DeleteHotParamRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteHotParamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotParamRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotParamRuleRequest) SetAhasRegionId(v string) *DeleteHotParamRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *DeleteHotParamRuleRequest) SetRuleId(v int64) *DeleteHotParamRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteHotParamRuleResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteHotParamRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHotParamRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotParamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHotParamRuleResponseBody) SetCode(v string) *DeleteHotParamRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHotParamRuleResponseBody) SetData(v *DeleteHotParamRuleResponseBodyData) *DeleteHotParamRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteHotParamRuleResponseBody) SetMessage(v string) *DeleteHotParamRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHotParamRuleResponseBody) SetRequestId(v string) *DeleteHotParamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHotParamRuleResponseBody) SetSuccess(v bool) *DeleteHotParamRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteHotParamRuleResponseBodyData struct {
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteHotParamRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotParamRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteHotParamRuleResponseBodyData) SetRuleId(v int64) *DeleteHotParamRuleResponseBodyData {
	s.RuleId = &v
	return s
}

type DeleteHotParamRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteHotParamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHotParamRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHotParamRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteHotParamRuleResponse) SetHeaders(v map[string]*string) *DeleteHotParamRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteHotParamRuleResponse) SetStatusCode(v int32) *DeleteHotParamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHotParamRuleResponse) SetBody(v *DeleteHotParamRuleResponseBody) *DeleteHotParamRuleResponse {
	s.Body = v
	return s
}

type DeleteIsolationRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteIsolationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIsolationRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteIsolationRuleRequest) SetAhasRegionId(v string) *DeleteIsolationRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *DeleteIsolationRuleRequest) SetRuleId(v int64) *DeleteIsolationRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteIsolationRuleResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteIsolationRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteIsolationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIsolationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIsolationRuleResponseBody) SetCode(v string) *DeleteIsolationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIsolationRuleResponseBody) SetData(v *DeleteIsolationRuleResponseBodyData) *DeleteIsolationRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteIsolationRuleResponseBody) SetMessage(v string) *DeleteIsolationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIsolationRuleResponseBody) SetRequestId(v string) *DeleteIsolationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIsolationRuleResponseBody) SetSuccess(v bool) *DeleteIsolationRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteIsolationRuleResponseBodyData struct {
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteIsolationRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteIsolationRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteIsolationRuleResponseBodyData) SetRuleId(v int64) *DeleteIsolationRuleResponseBodyData {
	s.RuleId = &v
	return s
}

type DeleteIsolationRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIsolationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIsolationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIsolationRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteIsolationRuleResponse) SetHeaders(v map[string]*string) *DeleteIsolationRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteIsolationRuleResponse) SetStatusCode(v int32) *DeleteIsolationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIsolationRuleResponse) SetBody(v *DeleteIsolationRuleResponseBody) *DeleteIsolationRuleResponse {
	s.Body = v
	return s
}

type DeleteSystemRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteSystemRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSystemRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteSystemRuleRequest) SetAhasRegionId(v string) *DeleteSystemRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *DeleteSystemRuleRequest) SetRuleId(v int64) *DeleteSystemRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteSystemRuleResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteSystemRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSystemRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSystemRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSystemRuleResponseBody) SetCode(v string) *DeleteSystemRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSystemRuleResponseBody) SetData(v *DeleteSystemRuleResponseBodyData) *DeleteSystemRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSystemRuleResponseBody) SetMessage(v string) *DeleteSystemRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSystemRuleResponseBody) SetRequestId(v string) *DeleteSystemRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSystemRuleResponseBody) SetSuccess(v bool) *DeleteSystemRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteSystemRuleResponseBodyData struct {
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteSystemRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteSystemRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSystemRuleResponseBodyData) SetRuleId(v int64) *DeleteSystemRuleResponseBodyData {
	s.RuleId = &v
	return s
}

type DeleteSystemRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSystemRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSystemRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSystemRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteSystemRuleResponse) SetHeaders(v map[string]*string) *DeleteSystemRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteSystemRuleResponse) SetStatusCode(v int32) *DeleteSystemRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSystemRuleResponse) SetBody(v *DeleteSystemRuleResponseBody) *DeleteSystemRuleResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DisableDegradeRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DisableDegradeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableDegradeRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableDegradeRuleRequest) SetAhasRegionId(v string) *DisableDegradeRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *DisableDegradeRuleRequest) SetRuleId(v int64) *DisableDegradeRuleRequest {
	s.RuleId = &v
	return s
}

type DisableDegradeRuleResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DisableDegradeRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableDegradeRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableDegradeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDegradeRuleResponseBody) SetCode(v string) *DisableDegradeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableDegradeRuleResponseBody) SetData(v *DisableDegradeRuleResponseBodyData) *DisableDegradeRuleResponseBody {
	s.Data = v
	return s
}

func (s *DisableDegradeRuleResponseBody) SetMessage(v string) *DisableDegradeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableDegradeRuleResponseBody) SetRequestId(v string) *DisableDegradeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDegradeRuleResponseBody) SetSuccess(v bool) *DisableDegradeRuleResponseBody {
	s.Success = &v
	return s
}

type DisableDegradeRuleResponseBodyData struct {
	AppName                   *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable                    *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HalfOpenBaseAmountPerStep *int32   `json:"HalfOpenBaseAmountPerStep,omitempty" xml:"HalfOpenBaseAmountPerStep,omitempty"`
	HalfOpenRecoveryStepNum   *int32   `json:"HalfOpenRecoveryStepNum,omitempty" xml:"HalfOpenRecoveryStepNum,omitempty"`
	MinRequestAmount          *int32   `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	Namespace                 *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RecoveryTimeoutMs         *int32   `json:"RecoveryTimeoutMs,omitempty" xml:"RecoveryTimeoutMs,omitempty"`
	Resource                  *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                    *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SlowRtMs                  *int32   `json:"SlowRtMs,omitempty" xml:"SlowRtMs,omitempty"`
	StatDurationMs            *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Strategy                  *int32   `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Threshold                 *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DisableDegradeRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DisableDegradeRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DisableDegradeRuleResponseBodyData) SetAppName(v string) *DisableDegradeRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetEnable(v bool) *DisableDegradeRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetHalfOpenBaseAmountPerStep(v int32) *DisableDegradeRuleResponseBodyData {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetHalfOpenRecoveryStepNum(v int32) *DisableDegradeRuleResponseBodyData {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetMinRequestAmount(v int32) *DisableDegradeRuleResponseBodyData {
	s.MinRequestAmount = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetNamespace(v string) *DisableDegradeRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetRecoveryTimeoutMs(v int32) *DisableDegradeRuleResponseBodyData {
	s.RecoveryTimeoutMs = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetResource(v string) *DisableDegradeRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetRuleId(v int64) *DisableDegradeRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetSlowRtMs(v int32) *DisableDegradeRuleResponseBodyData {
	s.SlowRtMs = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetStatDurationMs(v int32) *DisableDegradeRuleResponseBodyData {
	s.StatDurationMs = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetStrategy(v int32) *DisableDegradeRuleResponseBodyData {
	s.Strategy = &v
	return s
}

func (s *DisableDegradeRuleResponseBodyData) SetThreshold(v float32) *DisableDegradeRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type DisableDegradeRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableDegradeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableDegradeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableDegradeRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableDegradeRuleResponse) SetHeaders(v map[string]*string) *DisableDegradeRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableDegradeRuleResponse) SetStatusCode(v int32) *DisableDegradeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDegradeRuleResponse) SetBody(v *DisableDegradeRuleResponseBody) *DisableDegradeRuleResponse {
	s.Body = v
	return s
}

type DisableFlowRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DisableFlowRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableFlowRuleRequest) SetAhasRegionId(v string) *DisableFlowRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *DisableFlowRuleRequest) SetRuleId(v int64) *DisableFlowRuleRequest {
	s.RuleId = &v
	return s
}

type DisableFlowRuleResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DisableFlowRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableFlowRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableFlowRuleResponseBody) SetCode(v string) *DisableFlowRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableFlowRuleResponseBody) SetData(v *DisableFlowRuleResponseBodyData) *DisableFlowRuleResponseBody {
	s.Data = v
	return s
}

func (s *DisableFlowRuleResponseBody) SetMessage(v string) *DisableFlowRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableFlowRuleResponseBody) SetRequestId(v string) *DisableFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableFlowRuleResponseBody) SetSuccess(v bool) *DisableFlowRuleResponseBody {
	s.Success = &v
	return s
}

type DisableFlowRuleResponseBodyData struct {
	AppName                  *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ClusterEstimatedMaxQps   *float32 `json:"ClusterEstimatedMaxQps,omitempty" xml:"ClusterEstimatedMaxQps,omitempty"`
	ClusterFallbackStrategy  *int32   `json:"ClusterFallbackStrategy,omitempty" xml:"ClusterFallbackStrategy,omitempty"`
	ClusterFallbackThreshold *int32   `json:"ClusterFallbackThreshold,omitempty" xml:"ClusterFallbackThreshold,omitempty"`
	ClusterMode              *bool    `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	ClusterThresholdType     *int32   `json:"ClusterThresholdType,omitempty" xml:"ClusterThresholdType,omitempty"`
	ControlBehavior          *int32   `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable                   *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin              *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	MaxQueueingTimeMs        *int32   `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	Namespace                *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource              *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy         *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource                 *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                   *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationMs           *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Threshold                *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WarmUpPeriodSec          *int32   `json:"WarmUpPeriodSec,omitempty" xml:"WarmUpPeriodSec,omitempty"`
}

func (s DisableFlowRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DisableFlowRuleResponseBodyData) SetAppName(v string) *DisableFlowRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetClusterEstimatedMaxQps(v float32) *DisableFlowRuleResponseBodyData {
	s.ClusterEstimatedMaxQps = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetClusterFallbackStrategy(v int32) *DisableFlowRuleResponseBodyData {
	s.ClusterFallbackStrategy = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetClusterFallbackThreshold(v int32) *DisableFlowRuleResponseBodyData {
	s.ClusterFallbackThreshold = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetClusterMode(v bool) *DisableFlowRuleResponseBodyData {
	s.ClusterMode = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetClusterThresholdType(v int32) *DisableFlowRuleResponseBodyData {
	s.ClusterThresholdType = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetControlBehavior(v int32) *DisableFlowRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetEnable(v bool) *DisableFlowRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetLimitOrigin(v string) *DisableFlowRuleResponseBodyData {
	s.LimitOrigin = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *DisableFlowRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetNamespace(v string) *DisableFlowRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetRefResource(v string) *DisableFlowRuleResponseBodyData {
	s.RefResource = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetRelationStrategy(v int32) *DisableFlowRuleResponseBodyData {
	s.RelationStrategy = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetResource(v string) *DisableFlowRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetRuleId(v int64) *DisableFlowRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetStatDurationMs(v int32) *DisableFlowRuleResponseBodyData {
	s.StatDurationMs = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetThreshold(v float32) *DisableFlowRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *DisableFlowRuleResponseBodyData) SetWarmUpPeriodSec(v int32) *DisableFlowRuleResponseBodyData {
	s.WarmUpPeriodSec = &v
	return s
}

type DisableFlowRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableFlowRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableFlowRuleResponse) SetHeaders(v map[string]*string) *DisableFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableFlowRuleResponse) SetStatusCode(v int32) *DisableFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableFlowRuleResponse) SetBody(v *DisableFlowRuleResponseBody) *DisableFlowRuleResponse {
	s.Body = v
	return s
}

type DisableHotParamRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DisableHotParamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableHotParamRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableHotParamRuleRequest) SetAhasRegionId(v string) *DisableHotParamRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *DisableHotParamRuleRequest) SetRuleId(v int64) *DisableHotParamRuleRequest {
	s.RuleId = &v
	return s
}

type DisableHotParamRuleResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DisableHotParamRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableHotParamRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableHotParamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableHotParamRuleResponseBody) SetCode(v string) *DisableHotParamRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableHotParamRuleResponseBody) SetData(v *DisableHotParamRuleResponseBodyData) *DisableHotParamRuleResponseBody {
	s.Data = v
	return s
}

func (s *DisableHotParamRuleResponseBody) SetMessage(v string) *DisableHotParamRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableHotParamRuleResponseBody) SetRequestId(v string) *DisableHotParamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableHotParamRuleResponseBody) SetSuccess(v bool) *DisableHotParamRuleResponseBody {
	s.Success = &v
	return s
}

type DisableHotParamRuleResponseBodyData struct {
	AppName           *string                                                 `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BurstCount        *int32                                                  `json:"BurstCount,omitempty" xml:"BurstCount,omitempty"`
	ControlBehavior   *int32                                                  `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable            *bool                                                   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MaxQueueingTimeMs *int32                                                  `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	MetricType        *int32                                                  `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Namespace         *string                                                 `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParamFlowItemList []*DisableHotParamRuleResponseBodyDataParamFlowItemList `json:"ParamFlowItemList,omitempty" xml:"ParamFlowItemList,omitempty" type:"Repeated"`
	ParamIdx          *int32                                                  `json:"ParamIdx,omitempty" xml:"ParamIdx,omitempty"`
	Resource          *string                                                 `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId            *int64                                                  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationSec   *int64                                                  `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	Threshold         *float32                                                `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DisableHotParamRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DisableHotParamRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DisableHotParamRuleResponseBodyData) SetAppName(v string) *DisableHotParamRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetBurstCount(v int32) *DisableHotParamRuleResponseBodyData {
	s.BurstCount = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetControlBehavior(v int32) *DisableHotParamRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetEnable(v bool) *DisableHotParamRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *DisableHotParamRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetMetricType(v int32) *DisableHotParamRuleResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetNamespace(v string) *DisableHotParamRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetParamFlowItemList(v []*DisableHotParamRuleResponseBodyDataParamFlowItemList) *DisableHotParamRuleResponseBodyData {
	s.ParamFlowItemList = v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetParamIdx(v int32) *DisableHotParamRuleResponseBodyData {
	s.ParamIdx = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetResource(v string) *DisableHotParamRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetRuleId(v int64) *DisableHotParamRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetStatDurationSec(v int64) *DisableHotParamRuleResponseBodyData {
	s.StatDurationSec = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyData) SetThreshold(v float32) *DisableHotParamRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type DisableHotParamRuleResponseBodyDataParamFlowItemList struct {
	ItemType  *string  `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	ItemValue *string  `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DisableHotParamRuleResponseBodyDataParamFlowItemList) String() string {
	return tea.Prettify(s)
}

func (s DisableHotParamRuleResponseBodyDataParamFlowItemList) GoString() string {
	return s.String()
}

func (s *DisableHotParamRuleResponseBodyDataParamFlowItemList) SetItemType(v string) *DisableHotParamRuleResponseBodyDataParamFlowItemList {
	s.ItemType = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyDataParamFlowItemList) SetItemValue(v string) *DisableHotParamRuleResponseBodyDataParamFlowItemList {
	s.ItemValue = &v
	return s
}

func (s *DisableHotParamRuleResponseBodyDataParamFlowItemList) SetThreshold(v float32) *DisableHotParamRuleResponseBodyDataParamFlowItemList {
	s.Threshold = &v
	return s
}

type DisableHotParamRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableHotParamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableHotParamRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableHotParamRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableHotParamRuleResponse) SetHeaders(v map[string]*string) *DisableHotParamRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableHotParamRuleResponse) SetStatusCode(v int32) *DisableHotParamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableHotParamRuleResponse) SetBody(v *DisableHotParamRuleResponseBody) *DisableHotParamRuleResponse {
	s.Body = v
	return s
}

type DisableIsolationRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DisableIsolationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableIsolationRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableIsolationRuleRequest) SetAhasRegionId(v string) *DisableIsolationRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *DisableIsolationRuleRequest) SetRuleId(v int64) *DisableIsolationRuleRequest {
	s.RuleId = &v
	return s
}

type DisableIsolationRuleResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DisableIsolationRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableIsolationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableIsolationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableIsolationRuleResponseBody) SetCode(v string) *DisableIsolationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableIsolationRuleResponseBody) SetData(v *DisableIsolationRuleResponseBodyData) *DisableIsolationRuleResponseBody {
	s.Data = v
	return s
}

func (s *DisableIsolationRuleResponseBody) SetMessage(v string) *DisableIsolationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableIsolationRuleResponseBody) SetRequestId(v string) *DisableIsolationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableIsolationRuleResponseBody) SetSuccess(v bool) *DisableIsolationRuleResponseBody {
	s.Success = &v
	return s
}

type DisableIsolationRuleResponseBodyData struct {
	AppName          *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable           *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin      *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	Namespace        *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource      *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource         *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId           *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold        *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DisableIsolationRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DisableIsolationRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DisableIsolationRuleResponseBodyData) SetAppName(v string) *DisableIsolationRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DisableIsolationRuleResponseBodyData) SetEnable(v bool) *DisableIsolationRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *DisableIsolationRuleResponseBodyData) SetLimitOrigin(v string) *DisableIsolationRuleResponseBodyData {
	s.LimitOrigin = &v
	return s
}

func (s *DisableIsolationRuleResponseBodyData) SetNamespace(v string) *DisableIsolationRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *DisableIsolationRuleResponseBodyData) SetRefResource(v string) *DisableIsolationRuleResponseBodyData {
	s.RefResource = &v
	return s
}

func (s *DisableIsolationRuleResponseBodyData) SetRelationStrategy(v int32) *DisableIsolationRuleResponseBodyData {
	s.RelationStrategy = &v
	return s
}

func (s *DisableIsolationRuleResponseBodyData) SetResource(v string) *DisableIsolationRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *DisableIsolationRuleResponseBodyData) SetRuleId(v int64) *DisableIsolationRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *DisableIsolationRuleResponseBodyData) SetThreshold(v float32) *DisableIsolationRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type DisableIsolationRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableIsolationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableIsolationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableIsolationRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableIsolationRuleResponse) SetHeaders(v map[string]*string) *DisableIsolationRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableIsolationRuleResponse) SetStatusCode(v int32) *DisableIsolationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableIsolationRuleResponse) SetBody(v *DisableIsolationRuleResponseBody) *DisableIsolationRuleResponse {
	s.Body = v
	return s
}

type DisableSystemRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DisableSystemRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableSystemRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableSystemRuleRequest) SetAhasRegionId(v string) *DisableSystemRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *DisableSystemRuleRequest) SetRuleId(v int64) *DisableSystemRuleRequest {
	s.RuleId = &v
	return s
}

type DisableSystemRuleResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DisableSystemRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSystemRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableSystemRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSystemRuleResponseBody) SetCode(v string) *DisableSystemRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableSystemRuleResponseBody) SetData(v *DisableSystemRuleResponseBodyData) *DisableSystemRuleResponseBody {
	s.Data = v
	return s
}

func (s *DisableSystemRuleResponseBody) SetMessage(v string) *DisableSystemRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableSystemRuleResponseBody) SetRequestId(v string) *DisableSystemRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSystemRuleResponseBody) SetSuccess(v bool) *DisableSystemRuleResponseBody {
	s.Success = &v
	return s
}

type DisableSystemRuleResponseBodyData struct {
	Enable     *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MetricType *int32   `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	RuleId     *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold  *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DisableSystemRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DisableSystemRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DisableSystemRuleResponseBodyData) SetEnable(v bool) *DisableSystemRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *DisableSystemRuleResponseBodyData) SetMetricType(v int32) *DisableSystemRuleResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *DisableSystemRuleResponseBodyData) SetRuleId(v int64) *DisableSystemRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *DisableSystemRuleResponseBodyData) SetThreshold(v float32) *DisableSystemRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type DisableSystemRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableSystemRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableSystemRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableSystemRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableSystemRuleResponse) SetHeaders(v map[string]*string) *DisableSystemRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableSystemRuleResponse) SetStatusCode(v int32) *DisableSystemRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSystemRuleResponse) SetBody(v *DisableSystemRuleResponseBody) *DisableSystemRuleResponse {
	s.Body = v
	return s
}

type EnableDegradeRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s EnableDegradeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableDegradeRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableDegradeRuleRequest) SetAhasRegionId(v string) *EnableDegradeRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *EnableDegradeRuleRequest) SetRuleId(v int64) *EnableDegradeRuleRequest {
	s.RuleId = &v
	return s
}

type EnableDegradeRuleResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *EnableDegradeRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableDegradeRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableDegradeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableDegradeRuleResponseBody) SetCode(v string) *EnableDegradeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *EnableDegradeRuleResponseBody) SetData(v *EnableDegradeRuleResponseBodyData) *EnableDegradeRuleResponseBody {
	s.Data = v
	return s
}

func (s *EnableDegradeRuleResponseBody) SetMessage(v string) *EnableDegradeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *EnableDegradeRuleResponseBody) SetRequestId(v string) *EnableDegradeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableDegradeRuleResponseBody) SetSuccess(v bool) *EnableDegradeRuleResponseBody {
	s.Success = &v
	return s
}

type EnableDegradeRuleResponseBodyData struct {
	AppName                   *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable                    *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HalfOpenBaseAmountPerStep *int32   `json:"HalfOpenBaseAmountPerStep,omitempty" xml:"HalfOpenBaseAmountPerStep,omitempty"`
	HalfOpenRecoveryStepNum   *int32   `json:"HalfOpenRecoveryStepNum,omitempty" xml:"HalfOpenRecoveryStepNum,omitempty"`
	MinRequestAmount          *int32   `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	Namespace                 *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RecoveryTimeoutMs         *int32   `json:"RecoveryTimeoutMs,omitempty" xml:"RecoveryTimeoutMs,omitempty"`
	Resource                  *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                    *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SlowRtMs                  *int32   `json:"SlowRtMs,omitempty" xml:"SlowRtMs,omitempty"`
	StatDurationMs            *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Strategy                  *int32   `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Threshold                 *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s EnableDegradeRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnableDegradeRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnableDegradeRuleResponseBodyData) SetAppName(v string) *EnableDegradeRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetEnable(v bool) *EnableDegradeRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetHalfOpenBaseAmountPerStep(v int32) *EnableDegradeRuleResponseBodyData {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetHalfOpenRecoveryStepNum(v int32) *EnableDegradeRuleResponseBodyData {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetMinRequestAmount(v int32) *EnableDegradeRuleResponseBodyData {
	s.MinRequestAmount = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetNamespace(v string) *EnableDegradeRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetRecoveryTimeoutMs(v int32) *EnableDegradeRuleResponseBodyData {
	s.RecoveryTimeoutMs = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetResource(v string) *EnableDegradeRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetRuleId(v int64) *EnableDegradeRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetSlowRtMs(v int32) *EnableDegradeRuleResponseBodyData {
	s.SlowRtMs = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetStatDurationMs(v int32) *EnableDegradeRuleResponseBodyData {
	s.StatDurationMs = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetStrategy(v int32) *EnableDegradeRuleResponseBodyData {
	s.Strategy = &v
	return s
}

func (s *EnableDegradeRuleResponseBodyData) SetThreshold(v float32) *EnableDegradeRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type EnableDegradeRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableDegradeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableDegradeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableDegradeRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableDegradeRuleResponse) SetHeaders(v map[string]*string) *EnableDegradeRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableDegradeRuleResponse) SetStatusCode(v int32) *EnableDegradeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableDegradeRuleResponse) SetBody(v *EnableDegradeRuleResponseBody) *EnableDegradeRuleResponse {
	s.Body = v
	return s
}

type EnableFlowRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s EnableFlowRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableFlowRuleRequest) SetAhasRegionId(v string) *EnableFlowRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *EnableFlowRuleRequest) SetRuleId(v int64) *EnableFlowRuleRequest {
	s.RuleId = &v
	return s
}

type EnableFlowRuleResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *EnableFlowRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableFlowRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableFlowRuleResponseBody) SetCode(v string) *EnableFlowRuleResponseBody {
	s.Code = &v
	return s
}

func (s *EnableFlowRuleResponseBody) SetData(v *EnableFlowRuleResponseBodyData) *EnableFlowRuleResponseBody {
	s.Data = v
	return s
}

func (s *EnableFlowRuleResponseBody) SetMessage(v string) *EnableFlowRuleResponseBody {
	s.Message = &v
	return s
}

func (s *EnableFlowRuleResponseBody) SetRequestId(v string) *EnableFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableFlowRuleResponseBody) SetSuccess(v bool) *EnableFlowRuleResponseBody {
	s.Success = &v
	return s
}

type EnableFlowRuleResponseBodyData struct {
	AppName                  *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ClusterEstimatedMaxQps   *float32 `json:"ClusterEstimatedMaxQps,omitempty" xml:"ClusterEstimatedMaxQps,omitempty"`
	ClusterFallbackStrategy  *int32   `json:"ClusterFallbackStrategy,omitempty" xml:"ClusterFallbackStrategy,omitempty"`
	ClusterFallbackThreshold *int32   `json:"ClusterFallbackThreshold,omitempty" xml:"ClusterFallbackThreshold,omitempty"`
	ClusterMode              *bool    `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	ClusterThresholdType     *int32   `json:"ClusterThresholdType,omitempty" xml:"ClusterThresholdType,omitempty"`
	ControlBehavior          *int32   `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable                   *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin              *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	MaxQueueingTimeMs        *int32   `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	Namespace                *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource              *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy         *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource                 *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                   *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationMs           *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Threshold                *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WarmUpPeriodSec          *int32   `json:"WarmUpPeriodSec,omitempty" xml:"WarmUpPeriodSec,omitempty"`
}

func (s EnableFlowRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnableFlowRuleResponseBodyData) SetAppName(v string) *EnableFlowRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetClusterEstimatedMaxQps(v float32) *EnableFlowRuleResponseBodyData {
	s.ClusterEstimatedMaxQps = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetClusterFallbackStrategy(v int32) *EnableFlowRuleResponseBodyData {
	s.ClusterFallbackStrategy = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetClusterFallbackThreshold(v int32) *EnableFlowRuleResponseBodyData {
	s.ClusterFallbackThreshold = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetClusterMode(v bool) *EnableFlowRuleResponseBodyData {
	s.ClusterMode = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetClusterThresholdType(v int32) *EnableFlowRuleResponseBodyData {
	s.ClusterThresholdType = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetControlBehavior(v int32) *EnableFlowRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetEnable(v bool) *EnableFlowRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetLimitOrigin(v string) *EnableFlowRuleResponseBodyData {
	s.LimitOrigin = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *EnableFlowRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetNamespace(v string) *EnableFlowRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetRefResource(v string) *EnableFlowRuleResponseBodyData {
	s.RefResource = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetRelationStrategy(v int32) *EnableFlowRuleResponseBodyData {
	s.RelationStrategy = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetResource(v string) *EnableFlowRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetRuleId(v int64) *EnableFlowRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetStatDurationMs(v int32) *EnableFlowRuleResponseBodyData {
	s.StatDurationMs = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetThreshold(v float32) *EnableFlowRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *EnableFlowRuleResponseBodyData) SetWarmUpPeriodSec(v int32) *EnableFlowRuleResponseBodyData {
	s.WarmUpPeriodSec = &v
	return s
}

type EnableFlowRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableFlowRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableFlowRuleResponse) SetHeaders(v map[string]*string) *EnableFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableFlowRuleResponse) SetStatusCode(v int32) *EnableFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableFlowRuleResponse) SetBody(v *EnableFlowRuleResponseBody) *EnableFlowRuleResponse {
	s.Body = v
	return s
}

type EnableHotParamRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s EnableHotParamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableHotParamRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableHotParamRuleRequest) SetAhasRegionId(v string) *EnableHotParamRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *EnableHotParamRuleRequest) SetRuleId(v int64) *EnableHotParamRuleRequest {
	s.RuleId = &v
	return s
}

type EnableHotParamRuleResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *EnableHotParamRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableHotParamRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableHotParamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableHotParamRuleResponseBody) SetCode(v string) *EnableHotParamRuleResponseBody {
	s.Code = &v
	return s
}

func (s *EnableHotParamRuleResponseBody) SetData(v *EnableHotParamRuleResponseBodyData) *EnableHotParamRuleResponseBody {
	s.Data = v
	return s
}

func (s *EnableHotParamRuleResponseBody) SetMessage(v string) *EnableHotParamRuleResponseBody {
	s.Message = &v
	return s
}

func (s *EnableHotParamRuleResponseBody) SetRequestId(v string) *EnableHotParamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableHotParamRuleResponseBody) SetSuccess(v bool) *EnableHotParamRuleResponseBody {
	s.Success = &v
	return s
}

type EnableHotParamRuleResponseBodyData struct {
	AppName           *string                                                `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BurstCount        *int32                                                 `json:"BurstCount,omitempty" xml:"BurstCount,omitempty"`
	ControlBehavior   *int32                                                 `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable            *bool                                                  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MaxQueueingTimeMs *int32                                                 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	MetricType        *int32                                                 `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Namespace         *string                                                `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParamFlowItemList []*EnableHotParamRuleResponseBodyDataParamFlowItemList `json:"ParamFlowItemList,omitempty" xml:"ParamFlowItemList,omitempty" type:"Repeated"`
	ParamIdx          *int32                                                 `json:"ParamIdx,omitempty" xml:"ParamIdx,omitempty"`
	Resource          *string                                                `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId            *int64                                                 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationSec   *int64                                                 `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	Threshold         *float32                                               `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s EnableHotParamRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnableHotParamRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnableHotParamRuleResponseBodyData) SetAppName(v string) *EnableHotParamRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetBurstCount(v int32) *EnableHotParamRuleResponseBodyData {
	s.BurstCount = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetControlBehavior(v int32) *EnableHotParamRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetEnable(v bool) *EnableHotParamRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *EnableHotParamRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetMetricType(v int32) *EnableHotParamRuleResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetNamespace(v string) *EnableHotParamRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetParamFlowItemList(v []*EnableHotParamRuleResponseBodyDataParamFlowItemList) *EnableHotParamRuleResponseBodyData {
	s.ParamFlowItemList = v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetParamIdx(v int32) *EnableHotParamRuleResponseBodyData {
	s.ParamIdx = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetResource(v string) *EnableHotParamRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetRuleId(v int64) *EnableHotParamRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetStatDurationSec(v int64) *EnableHotParamRuleResponseBodyData {
	s.StatDurationSec = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyData) SetThreshold(v float32) *EnableHotParamRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type EnableHotParamRuleResponseBodyDataParamFlowItemList struct {
	ItemType  *string  `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	ItemValue *string  `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s EnableHotParamRuleResponseBodyDataParamFlowItemList) String() string {
	return tea.Prettify(s)
}

func (s EnableHotParamRuleResponseBodyDataParamFlowItemList) GoString() string {
	return s.String()
}

func (s *EnableHotParamRuleResponseBodyDataParamFlowItemList) SetItemType(v string) *EnableHotParamRuleResponseBodyDataParamFlowItemList {
	s.ItemType = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyDataParamFlowItemList) SetItemValue(v string) *EnableHotParamRuleResponseBodyDataParamFlowItemList {
	s.ItemValue = &v
	return s
}

func (s *EnableHotParamRuleResponseBodyDataParamFlowItemList) SetThreshold(v float32) *EnableHotParamRuleResponseBodyDataParamFlowItemList {
	s.Threshold = &v
	return s
}

type EnableHotParamRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableHotParamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableHotParamRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableHotParamRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableHotParamRuleResponse) SetHeaders(v map[string]*string) *EnableHotParamRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableHotParamRuleResponse) SetStatusCode(v int32) *EnableHotParamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableHotParamRuleResponse) SetBody(v *EnableHotParamRuleResponseBody) *EnableHotParamRuleResponse {
	s.Body = v
	return s
}

type EnableIsolationRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s EnableIsolationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableIsolationRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableIsolationRuleRequest) SetAhasRegionId(v string) *EnableIsolationRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *EnableIsolationRuleRequest) SetRuleId(v int64) *EnableIsolationRuleRequest {
	s.RuleId = &v
	return s
}

type EnableIsolationRuleResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *EnableIsolationRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableIsolationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableIsolationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableIsolationRuleResponseBody) SetCode(v string) *EnableIsolationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *EnableIsolationRuleResponseBody) SetData(v *EnableIsolationRuleResponseBodyData) *EnableIsolationRuleResponseBody {
	s.Data = v
	return s
}

func (s *EnableIsolationRuleResponseBody) SetMessage(v string) *EnableIsolationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *EnableIsolationRuleResponseBody) SetRequestId(v string) *EnableIsolationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableIsolationRuleResponseBody) SetSuccess(v bool) *EnableIsolationRuleResponseBody {
	s.Success = &v
	return s
}

type EnableIsolationRuleResponseBodyData struct {
	AppName          *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable           *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin      *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	Namespace        *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource      *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource         *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId           *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold        *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s EnableIsolationRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnableIsolationRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnableIsolationRuleResponseBodyData) SetAppName(v string) *EnableIsolationRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *EnableIsolationRuleResponseBodyData) SetEnable(v bool) *EnableIsolationRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *EnableIsolationRuleResponseBodyData) SetLimitOrigin(v string) *EnableIsolationRuleResponseBodyData {
	s.LimitOrigin = &v
	return s
}

func (s *EnableIsolationRuleResponseBodyData) SetNamespace(v string) *EnableIsolationRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *EnableIsolationRuleResponseBodyData) SetRefResource(v string) *EnableIsolationRuleResponseBodyData {
	s.RefResource = &v
	return s
}

func (s *EnableIsolationRuleResponseBodyData) SetRelationStrategy(v int32) *EnableIsolationRuleResponseBodyData {
	s.RelationStrategy = &v
	return s
}

func (s *EnableIsolationRuleResponseBodyData) SetResource(v string) *EnableIsolationRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *EnableIsolationRuleResponseBodyData) SetRuleId(v int64) *EnableIsolationRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *EnableIsolationRuleResponseBodyData) SetThreshold(v float32) *EnableIsolationRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type EnableIsolationRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableIsolationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableIsolationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableIsolationRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableIsolationRuleResponse) SetHeaders(v map[string]*string) *EnableIsolationRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableIsolationRuleResponse) SetStatusCode(v int32) *EnableIsolationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableIsolationRuleResponse) SetBody(v *EnableIsolationRuleResponseBody) *EnableIsolationRuleResponse {
	s.Body = v
	return s
}

type EnableSystemRuleRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s EnableSystemRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSystemRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableSystemRuleRequest) SetAhasRegionId(v string) *EnableSystemRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *EnableSystemRuleRequest) SetRuleId(v int64) *EnableSystemRuleRequest {
	s.RuleId = &v
	return s
}

type EnableSystemRuleResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *EnableSystemRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSystemRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSystemRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSystemRuleResponseBody) SetCode(v string) *EnableSystemRuleResponseBody {
	s.Code = &v
	return s
}

func (s *EnableSystemRuleResponseBody) SetData(v *EnableSystemRuleResponseBodyData) *EnableSystemRuleResponseBody {
	s.Data = v
	return s
}

func (s *EnableSystemRuleResponseBody) SetMessage(v string) *EnableSystemRuleResponseBody {
	s.Message = &v
	return s
}

func (s *EnableSystemRuleResponseBody) SetRequestId(v string) *EnableSystemRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableSystemRuleResponseBody) SetSuccess(v bool) *EnableSystemRuleResponseBody {
	s.Success = &v
	return s
}

type EnableSystemRuleResponseBodyData struct {
	Enable     *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MetricType *int32   `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	RuleId     *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold  *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s EnableSystemRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnableSystemRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnableSystemRuleResponseBodyData) SetEnable(v bool) *EnableSystemRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *EnableSystemRuleResponseBodyData) SetMetricType(v int32) *EnableSystemRuleResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *EnableSystemRuleResponseBodyData) SetRuleId(v int64) *EnableSystemRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *EnableSystemRuleResponseBodyData) SetThreshold(v float32) *EnableSystemRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type EnableSystemRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableSystemRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSystemRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSystemRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableSystemRuleResponse) SetHeaders(v map[string]*string) *EnableSystemRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableSystemRuleResponse) SetStatusCode(v int32) *EnableSystemRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableSystemRuleResponse) SetBody(v *EnableSystemRuleResponseBody) *EnableSystemRuleResponse {
	s.Body = v
	return s
}

type ExecuteExperimentRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	Definition   *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	NameSpace    *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
}

func (s ExecuteExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteExperimentRequest) GoString() string {
	return s.String()
}

func (s *ExecuteExperimentRequest) SetAhasRegionId(v string) *ExecuteExperimentRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ExecuteExperimentRequest) SetDefinition(v string) *ExecuteExperimentRequest {
	s.Definition = &v
	return s
}

func (s *ExecuteExperimentRequest) SetExperimentId(v string) *ExecuteExperimentRequest {
	s.ExperimentId = &v
	return s
}

func (s *ExecuteExperimentRequest) SetNameSpace(v string) *ExecuteExperimentRequest {
	s.NameSpace = &v
	return s
}

type ExecuteExperimentResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExecuteExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteExperimentResponseBody) SetCode(v string) *ExecuteExperimentResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteExperimentResponseBody) SetMessage(v string) *ExecuteExperimentResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteExperimentResponseBody) SetRequestId(v string) *ExecuteExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteExperimentResponseBody) SetSuccess(v bool) *ExecuteExperimentResponseBody {
	s.Success = &v
	return s
}

func (s *ExecuteExperimentResponseBody) SetTaskId(v string) *ExecuteExperimentResponseBody {
	s.TaskId = &v
	return s
}

type ExecuteExperimentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecuteExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteExperimentResponse) GoString() string {
	return s.String()
}

func (s *ExecuteExperimentResponse) SetHeaders(v map[string]*string) *ExecuteExperimentResponse {
	s.Headers = v
	return s
}

func (s *ExecuteExperimentResponse) SetStatusCode(v int32) *ExecuteExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteExperimentResponse) SetBody(v *ExecuteExperimentResponseBody) *ExecuteExperimentResponse {
	s.Body = v
	return s
}

type FinishExperimentTaskRequest struct {
	AhasRegionId     *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	ExperimentTaskId *string `json:"ExperimentTaskId,omitempty" xml:"ExperimentTaskId,omitempty"`
	NameSpace        *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
}

func (s FinishExperimentTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishExperimentTaskRequest) GoString() string {
	return s.String()
}

func (s *FinishExperimentTaskRequest) SetAhasRegionId(v string) *FinishExperimentTaskRequest {
	s.AhasRegionId = &v
	return s
}

func (s *FinishExperimentTaskRequest) SetExperimentTaskId(v string) *FinishExperimentTaskRequest {
	s.ExperimentTaskId = &v
	return s
}

func (s *FinishExperimentTaskRequest) SetNameSpace(v string) *FinishExperimentTaskRequest {
	s.NameSpace = &v
	return s
}

type FinishExperimentTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FinishExperimentTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishExperimentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *FinishExperimentTaskResponseBody) SetCode(v string) *FinishExperimentTaskResponseBody {
	s.Code = &v
	return s
}

func (s *FinishExperimentTaskResponseBody) SetMessage(v string) *FinishExperimentTaskResponseBody {
	s.Message = &v
	return s
}

func (s *FinishExperimentTaskResponseBody) SetRequestId(v string) *FinishExperimentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishExperimentTaskResponseBody) SetSuccess(v bool) *FinishExperimentTaskResponseBody {
	s.Success = &v
	return s
}

type FinishExperimentTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FinishExperimentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FinishExperimentTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishExperimentTaskResponse) GoString() string {
	return s.String()
}

func (s *FinishExperimentTaskResponse) SetHeaders(v map[string]*string) *FinishExperimentTaskResponse {
	s.Headers = v
	return s
}

func (s *FinishExperimentTaskResponse) SetStatusCode(v int32) *FinishExperimentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishExperimentTaskResponse) SetBody(v *FinishExperimentTaskResponseBody) *FinishExperimentTaskResponse {
	s.Body = v
	return s
}

type GetActivityTaskRequest struct {
	ActivityTaskId   *string `json:"ActivityTaskId,omitempty" xml:"ActivityTaskId,omitempty"`
	AhasRegionId     *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	ExperimentTaskId *string `json:"ExperimentTaskId,omitempty" xml:"ExperimentTaskId,omitempty"`
	NameSpace        *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
}

func (s GetActivityTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetActivityTaskRequest) GoString() string {
	return s.String()
}

func (s *GetActivityTaskRequest) SetActivityTaskId(v string) *GetActivityTaskRequest {
	s.ActivityTaskId = &v
	return s
}

func (s *GetActivityTaskRequest) SetAhasRegionId(v string) *GetActivityTaskRequest {
	s.AhasRegionId = &v
	return s
}

func (s *GetActivityTaskRequest) SetExperimentTaskId(v string) *GetActivityTaskRequest {
	s.ExperimentTaskId = &v
	return s
}

func (s *GetActivityTaskRequest) SetNameSpace(v string) *GetActivityTaskRequest {
	s.NameSpace = &v
	return s
}

type GetActivityTaskResponseBody struct {
	ActivityId       *string                             `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	ActivityName     *string                             `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	EndTime          *int64                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExperimentTaskId *string                             `json:"ExperimentTaskId,omitempty" xml:"ExperimentTaskId,omitempty"`
	Hosts            []*GetActivityTaskResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	HttpStatusCode   *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Phase            *string                             `json:"Phase,omitempty" xml:"Phase,omitempty"`
	RequestId        *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunResult        *string                             `json:"RunResult,omitempty" xml:"RunResult,omitempty"`
	StartTime        *int64                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State            *string                             `json:"State,omitempty" xml:"State,omitempty"`
	Success          *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetActivityTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetActivityTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetActivityTaskResponseBody) SetActivityId(v string) *GetActivityTaskResponseBody {
	s.ActivityId = &v
	return s
}

func (s *GetActivityTaskResponseBody) SetActivityName(v string) *GetActivityTaskResponseBody {
	s.ActivityName = &v
	return s
}

func (s *GetActivityTaskResponseBody) SetEndTime(v int64) *GetActivityTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetActivityTaskResponseBody) SetExperimentTaskId(v string) *GetActivityTaskResponseBody {
	s.ExperimentTaskId = &v
	return s
}

func (s *GetActivityTaskResponseBody) SetHosts(v []*GetActivityTaskResponseBodyHosts) *GetActivityTaskResponseBody {
	s.Hosts = v
	return s
}

func (s *GetActivityTaskResponseBody) SetHttpStatusCode(v int32) *GetActivityTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetActivityTaskResponseBody) SetPhase(v string) *GetActivityTaskResponseBody {
	s.Phase = &v
	return s
}

func (s *GetActivityTaskResponseBody) SetRequestId(v string) *GetActivityTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetActivityTaskResponseBody) SetRunResult(v string) *GetActivityTaskResponseBody {
	s.RunResult = &v
	return s
}

func (s *GetActivityTaskResponseBody) SetStartTime(v int64) *GetActivityTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetActivityTaskResponseBody) SetState(v string) *GetActivityTaskResponseBody {
	s.State = &v
	return s
}

func (s *GetActivityTaskResponseBody) SetSuccess(v bool) *GetActivityTaskResponseBody {
	s.Success = &v
	return s
}

type GetActivityTaskResponseBodyHosts struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ExpId        *string `json:"ExpId,omitempty" xml:"ExpId,omitempty"`
	HostIp       *string `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetActivityTaskResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s GetActivityTaskResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *GetActivityTaskResponseBodyHosts) SetData(v string) *GetActivityTaskResponseBodyHosts {
	s.Data = &v
	return s
}

func (s *GetActivityTaskResponseBodyHosts) SetEndTime(v int64) *GetActivityTaskResponseBodyHosts {
	s.EndTime = &v
	return s
}

func (s *GetActivityTaskResponseBodyHosts) SetErrorMessage(v string) *GetActivityTaskResponseBodyHosts {
	s.ErrorMessage = &v
	return s
}

func (s *GetActivityTaskResponseBodyHosts) SetExpId(v string) *GetActivityTaskResponseBodyHosts {
	s.ExpId = &v
	return s
}

func (s *GetActivityTaskResponseBodyHosts) SetHostIp(v string) *GetActivityTaskResponseBodyHosts {
	s.HostIp = &v
	return s
}

func (s *GetActivityTaskResponseBodyHosts) SetResult(v string) *GetActivityTaskResponseBodyHosts {
	s.Result = &v
	return s
}

func (s *GetActivityTaskResponseBodyHosts) SetStartTime(v int64) *GetActivityTaskResponseBodyHosts {
	s.StartTime = &v
	return s
}

func (s *GetActivityTaskResponseBodyHosts) SetState(v string) *GetActivityTaskResponseBodyHosts {
	s.State = &v
	return s
}

func (s *GetActivityTaskResponseBodyHosts) SetTaskId(v string) *GetActivityTaskResponseBodyHosts {
	s.TaskId = &v
	return s
}

type GetActivityTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetActivityTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetActivityTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetActivityTaskResponse) GoString() string {
	return s.String()
}

func (s *GetActivityTaskResponse) SetHeaders(v map[string]*string) *GetActivityTaskResponse {
	s.Headers = v
	return s
}

func (s *GetActivityTaskResponse) SetStatusCode(v int32) *GetActivityTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetActivityTaskResponse) SetBody(v *GetActivityTaskResponseBody) *GetActivityTaskResponse {
	s.Body = v
	return s
}

type GetExperimentMetaRequest struct {
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	NameSpace    *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
}

func (s GetExperimentMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentMetaRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentMetaRequest) SetExperimentId(v string) *GetExperimentMetaRequest {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentMetaRequest) SetNameSpace(v string) *GetExperimentMetaRequest {
	s.NameSpace = &v
	return s
}

type GetExperimentMetaResponseBody struct {
	Code         *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateTime   *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExperimentId *string   `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Message      *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State        *string   `json:"State,omitempty" xml:"State,omitempty"`
	Success      *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	Tags         []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetExperimentMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentMetaResponseBody) SetCode(v string) *GetExperimentMetaResponseBody {
	s.Code = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetCreateTime(v string) *GetExperimentMetaResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetExperimentId(v string) *GetExperimentMetaResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetMessage(v string) *GetExperimentMetaResponseBody {
	s.Message = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetName(v string) *GetExperimentMetaResponseBody {
	s.Name = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetRequestId(v string) *GetExperimentMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetState(v string) *GetExperimentMetaResponseBody {
	s.State = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetSuccess(v bool) *GetExperimentMetaResponseBody {
	s.Success = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetTags(v []*string) *GetExperimentMetaResponseBody {
	s.Tags = v
	return s
}

type GetExperimentMetaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExperimentMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExperimentMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentMetaResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentMetaResponse) SetHeaders(v map[string]*string) *GetExperimentMetaResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentMetaResponse) SetStatusCode(v int32) *GetExperimentMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentMetaResponse) SetBody(v *GetExperimentMetaResponseBody) *GetExperimentMetaResponse {
	s.Body = v
	return s
}

type GetExperimentTaskRequest struct {
	AhasRegionId     *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	ExperimentTaskId *string `json:"ExperimentTaskId,omitempty" xml:"ExperimentTaskId,omitempty"`
	NameSpace        *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
}

func (s GetExperimentTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentTaskRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentTaskRequest) SetAhasRegionId(v string) *GetExperimentTaskRequest {
	s.AhasRegionId = &v
	return s
}

func (s *GetExperimentTaskRequest) SetExperimentTaskId(v string) *GetExperimentTaskRequest {
	s.ExperimentTaskId = &v
	return s
}

func (s *GetExperimentTaskRequest) SetNameSpace(v string) *GetExperimentTaskRequest {
	s.NameSpace = &v
	return s
}

type GetExperimentTaskResponseBody struct {
	Activities     []*GetExperimentTaskResponseBodyActivities `json:"Activities,omitempty" xml:"Activities,omitempty" type:"Repeated"`
	ExperimentId   *string                                    `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ExperimentName *string                                    `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Namespace      *string                                    `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result         *string                                    `json:"Result,omitempty" xml:"Result,omitempty"`
	StartTime      *int64                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State          *string                                    `json:"State,omitempty" xml:"State,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId         *string                                    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetExperimentTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentTaskResponseBody) SetActivities(v []*GetExperimentTaskResponseBodyActivities) *GetExperimentTaskResponseBody {
	s.Activities = v
	return s
}

func (s *GetExperimentTaskResponseBody) SetExperimentId(v string) *GetExperimentTaskResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentTaskResponseBody) SetExperimentName(v string) *GetExperimentTaskResponseBody {
	s.ExperimentName = &v
	return s
}

func (s *GetExperimentTaskResponseBody) SetHttpStatusCode(v int32) *GetExperimentTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetExperimentTaskResponseBody) SetNamespace(v string) *GetExperimentTaskResponseBody {
	s.Namespace = &v
	return s
}

func (s *GetExperimentTaskResponseBody) SetRequestId(v string) *GetExperimentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentTaskResponseBody) SetResult(v string) *GetExperimentTaskResponseBody {
	s.Result = &v
	return s
}

func (s *GetExperimentTaskResponseBody) SetStartTime(v int64) *GetExperimentTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetExperimentTaskResponseBody) SetState(v string) *GetExperimentTaskResponseBody {
	s.State = &v
	return s
}

func (s *GetExperimentTaskResponseBody) SetSuccess(v bool) *GetExperimentTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetExperimentTaskResponseBody) SetTaskId(v string) *GetExperimentTaskResponseBody {
	s.TaskId = &v
	return s
}

type GetExperimentTaskResponseBodyActivities struct {
	ActivityId       *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	ActivityName     *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	CheckState       *string `json:"CheckState,omitempty" xml:"CheckState,omitempty"`
	EndTime          *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExperimentTaskId *string `json:"ExperimentTaskId,omitempty" xml:"ExperimentTaskId,omitempty"`
	Phase            *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	RunResult        *string `json:"RunResult,omitempty" xml:"RunResult,omitempty"`
	StartTime        *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	TaskId           *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetExperimentTaskResponseBodyActivities) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentTaskResponseBodyActivities) GoString() string {
	return s.String()
}

func (s *GetExperimentTaskResponseBodyActivities) SetActivityId(v string) *GetExperimentTaskResponseBodyActivities {
	s.ActivityId = &v
	return s
}

func (s *GetExperimentTaskResponseBodyActivities) SetActivityName(v string) *GetExperimentTaskResponseBodyActivities {
	s.ActivityName = &v
	return s
}

func (s *GetExperimentTaskResponseBodyActivities) SetCheckState(v string) *GetExperimentTaskResponseBodyActivities {
	s.CheckState = &v
	return s
}

func (s *GetExperimentTaskResponseBodyActivities) SetEndTime(v int64) *GetExperimentTaskResponseBodyActivities {
	s.EndTime = &v
	return s
}

func (s *GetExperimentTaskResponseBodyActivities) SetExperimentTaskId(v string) *GetExperimentTaskResponseBodyActivities {
	s.ExperimentTaskId = &v
	return s
}

func (s *GetExperimentTaskResponseBodyActivities) SetPhase(v string) *GetExperimentTaskResponseBodyActivities {
	s.Phase = &v
	return s
}

func (s *GetExperimentTaskResponseBodyActivities) SetRunResult(v string) *GetExperimentTaskResponseBodyActivities {
	s.RunResult = &v
	return s
}

func (s *GetExperimentTaskResponseBodyActivities) SetStartTime(v int64) *GetExperimentTaskResponseBodyActivities {
	s.StartTime = &v
	return s
}

func (s *GetExperimentTaskResponseBodyActivities) SetState(v string) *GetExperimentTaskResponseBodyActivities {
	s.State = &v
	return s
}

func (s *GetExperimentTaskResponseBodyActivities) SetTaskId(v string) *GetExperimentTaskResponseBodyActivities {
	s.TaskId = &v
	return s
}

type GetExperimentTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExperimentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExperimentTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentTaskResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentTaskResponse) SetHeaders(v map[string]*string) *GetExperimentTaskResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentTaskResponse) SetStatusCode(v int32) *GetExperimentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentTaskResponse) SetBody(v *GetExperimentTaskResponseBody) *GetExperimentTaskResponse {
	s.Body = v
	return s
}

type GetLicenseKeyRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetLicenseKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseKeyRequest) GoString() string {
	return s.String()
}

func (s *GetLicenseKeyRequest) SetAhasRegionId(v string) *GetLicenseKeyRequest {
	s.AhasRegionId = &v
	return s
}

func (s *GetLicenseKeyRequest) SetNamespace(v string) *GetLicenseKeyRequest {
	s.Namespace = &v
	return s
}

type GetLicenseKeyResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLicenseKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetLicenseKeyResponseBody) SetCode(v string) *GetLicenseKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GetLicenseKeyResponseBody) SetData(v string) *GetLicenseKeyResponseBody {
	s.Data = &v
	return s
}

func (s *GetLicenseKeyResponseBody) SetMessage(v string) *GetLicenseKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GetLicenseKeyResponseBody) SetRequestId(v string) *GetLicenseKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLicenseKeyResponseBody) SetSuccess(v bool) *GetLicenseKeyResponseBody {
	s.Success = &v
	return s
}

type GetLicenseKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLicenseKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLicenseKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseKeyResponse) GoString() string {
	return s.String()
}

func (s *GetLicenseKeyResponse) SetHeaders(v map[string]*string) *GetLicenseKeyResponse {
	s.Headers = v
	return s
}

func (s *GetLicenseKeyResponse) SetStatusCode(v int32) *GetLicenseKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLicenseKeyResponse) SetBody(v *GetLicenseKeyResponseBody) *GetLicenseKeyResponse {
	s.Body = v
	return s
}

type GetMetricsOfAppRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetMetricsOfAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMetricsOfAppRequest) GoString() string {
	return s.String()
}

func (s *GetMetricsOfAppRequest) SetAhasRegionId(v string) *GetMetricsOfAppRequest {
	s.AhasRegionId = &v
	return s
}

func (s *GetMetricsOfAppRequest) SetAppName(v string) *GetMetricsOfAppRequest {
	s.AppName = &v
	return s
}

func (s *GetMetricsOfAppRequest) SetEndTime(v int64) *GetMetricsOfAppRequest {
	s.EndTime = &v
	return s
}

func (s *GetMetricsOfAppRequest) SetNamespace(v string) *GetMetricsOfAppRequest {
	s.Namespace = &v
	return s
}

func (s *GetMetricsOfAppRequest) SetStartTime(v int64) *GetMetricsOfAppRequest {
	s.StartTime = &v
	return s
}

type GetMetricsOfAppResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetMetricsOfAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetricsOfAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMetricsOfAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetricsOfAppResponseBody) SetCode(v string) *GetMetricsOfAppResponseBody {
	s.Code = &v
	return s
}

func (s *GetMetricsOfAppResponseBody) SetData(v *GetMetricsOfAppResponseBodyData) *GetMetricsOfAppResponseBody {
	s.Data = v
	return s
}

func (s *GetMetricsOfAppResponseBody) SetMessage(v string) *GetMetricsOfAppResponseBody {
	s.Message = &v
	return s
}

func (s *GetMetricsOfAppResponseBody) SetRequestId(v string) *GetMetricsOfAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetricsOfAppResponseBody) SetSuccess(v bool) *GetMetricsOfAppResponseBody {
	s.Success = &v
	return s
}

type GetMetricsOfAppResponseBodyData struct {
	AppName      *string                                        `json:"AppName,omitempty" xml:"AppName,omitempty"`
	InnerMetrics []*GetMetricsOfAppResponseBodyDataInnerMetrics `json:"InnerMetrics,omitempty" xml:"InnerMetrics,omitempty" type:"Repeated"`
	Namespace    *string                                        `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetMetricsOfAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMetricsOfAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetricsOfAppResponseBodyData) SetAppName(v string) *GetMetricsOfAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyData) SetInnerMetrics(v []*GetMetricsOfAppResponseBodyDataInnerMetrics) *GetMetricsOfAppResponseBodyData {
	s.InnerMetrics = v
	return s
}

func (s *GetMetricsOfAppResponseBodyData) SetNamespace(v string) *GetMetricsOfAppResponseBodyData {
	s.Namespace = &v
	return s
}

type GetMetricsOfAppResponseBodyDataInnerMetrics struct {
	BlockedQps    *float32 `json:"BlockedQps,omitempty" xml:"BlockedQps,omitempty"`
	BlockedQpsAvg *float32 `json:"BlockedQpsAvg,omitempty" xml:"BlockedQpsAvg,omitempty"`
	BlockedQpsMax *float32 `json:"BlockedQpsMax,omitempty" xml:"BlockedQpsMax,omitempty"`
	BlockedQpsMin *float32 `json:"BlockedQpsMin,omitempty" xml:"BlockedQpsMin,omitempty"`
	BlockedQpsP75 *float32 `json:"BlockedQpsP75,omitempty" xml:"BlockedQpsP75,omitempty"`
	BlockedQpsP95 *float32 `json:"BlockedQpsP95,omitempty" xml:"BlockedQpsP95,omitempty"`
	BlockedQpsP99 *float32 `json:"BlockedQpsP99,omitempty" xml:"BlockedQpsP99,omitempty"`
	BlockedQpsStd *float32 `json:"BlockedQpsStd,omitempty" xml:"BlockedQpsStd,omitempty"`
	Count         *int32   `json:"Count,omitempty" xml:"Count,omitempty"`
	Exception     *float32 `json:"Exception,omitempty" xml:"Exception,omitempty"`
	ExceptionAvg  *float32 `json:"ExceptionAvg,omitempty" xml:"ExceptionAvg,omitempty"`
	ExceptionMax  *float32 `json:"ExceptionMax,omitempty" xml:"ExceptionMax,omitempty"`
	ExceptionMin  *float32 `json:"ExceptionMin,omitempty" xml:"ExceptionMin,omitempty"`
	ExceptionP75  *float32 `json:"ExceptionP75,omitempty" xml:"ExceptionP75,omitempty"`
	ExceptionP95  *float32 `json:"ExceptionP95,omitempty" xml:"ExceptionP95,omitempty"`
	ExceptionP99  *float32 `json:"ExceptionP99,omitempty" xml:"ExceptionP99,omitempty"`
	ExceptionStd  *float32 `json:"ExceptionStd,omitempty" xml:"ExceptionStd,omitempty"`
	PassedQps     *float32 `json:"PassedQps,omitempty" xml:"PassedQps,omitempty"`
	PassedQpsAvg  *float32 `json:"PassedQpsAvg,omitempty" xml:"PassedQpsAvg,omitempty"`
	PassedQpsMax  *float32 `json:"PassedQpsMax,omitempty" xml:"PassedQpsMax,omitempty"`
	PassedQpsMin  *float32 `json:"PassedQpsMin,omitempty" xml:"PassedQpsMin,omitempty"`
	PassedQpsP75  *float32 `json:"PassedQpsP75,omitempty" xml:"PassedQpsP75,omitempty"`
	PassedQpsP95  *float32 `json:"PassedQpsP95,omitempty" xml:"PassedQpsP95,omitempty"`
	PassedQpsP99  *float32 `json:"PassedQpsP99,omitempty" xml:"PassedQpsP99,omitempty"`
	PassedQpsStd  *float32 `json:"PassedQpsStd,omitempty" xml:"PassedQpsStd,omitempty"`
	Rt            *float32 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	RtAvg         *float32 `json:"RtAvg,omitempty" xml:"RtAvg,omitempty"`
	RtMax         *float32 `json:"RtMax,omitempty" xml:"RtMax,omitempty"`
	RtMin         *float32 `json:"RtMin,omitempty" xml:"RtMin,omitempty"`
	RtP75         *float32 `json:"RtP75,omitempty" xml:"RtP75,omitempty"`
	RtP95         *float32 `json:"RtP95,omitempty" xml:"RtP95,omitempty"`
	RtP99         *float32 `json:"RtP99,omitempty" xml:"RtP99,omitempty"`
	RtStd         *float32 `json:"RtStd,omitempty" xml:"RtStd,omitempty"`
	SuccessQps    *float32 `json:"SuccessQps,omitempty" xml:"SuccessQps,omitempty"`
	SuccessQpsAvg *float32 `json:"SuccessQpsAvg,omitempty" xml:"SuccessQpsAvg,omitempty"`
	SuccessQpsMax *float32 `json:"SuccessQpsMax,omitempty" xml:"SuccessQpsMax,omitempty"`
	SuccessQpsMin *float32 `json:"SuccessQpsMin,omitempty" xml:"SuccessQpsMin,omitempty"`
	SuccessQpsP75 *float32 `json:"SuccessQpsP75,omitempty" xml:"SuccessQpsP75,omitempty"`
	SuccessQpsP95 *float32 `json:"SuccessQpsP95,omitempty" xml:"SuccessQpsP95,omitempty"`
	SuccessQpsP99 *float32 `json:"SuccessQpsP99,omitempty" xml:"SuccessQpsP99,omitempty"`
	SuccessQpsStd *float32 `json:"SuccessQpsStd,omitempty" xml:"SuccessQpsStd,omitempty"`
	Thread        *float32 `json:"Thread,omitempty" xml:"Thread,omitempty"`
	ThreadAvg     *float32 `json:"ThreadAvg,omitempty" xml:"ThreadAvg,omitempty"`
	ThreadMax     *float32 `json:"ThreadMax,omitempty" xml:"ThreadMax,omitempty"`
	ThreadMin     *float32 `json:"ThreadMin,omitempty" xml:"ThreadMin,omitempty"`
	ThreadP75     *float32 `json:"ThreadP75,omitempty" xml:"ThreadP75,omitempty"`
	ThreadP95     *float32 `json:"ThreadP95,omitempty" xml:"ThreadP95,omitempty"`
	ThreadP99     *float32 `json:"ThreadP99,omitempty" xml:"ThreadP99,omitempty"`
	ThreadStd     *float32 `json:"ThreadStd,omitempty" xml:"ThreadStd,omitempty"`
	Timestamp     *int64   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMetricsOfAppResponseBodyDataInnerMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetMetricsOfAppResponseBodyDataInnerMetrics) GoString() string {
	return s.String()
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetBlockedQps(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.BlockedQps = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetBlockedQpsAvg(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.BlockedQpsAvg = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetBlockedQpsMax(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.BlockedQpsMax = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetBlockedQpsMin(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.BlockedQpsMin = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetBlockedQpsP75(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.BlockedQpsP75 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetBlockedQpsP95(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.BlockedQpsP95 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetBlockedQpsP99(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.BlockedQpsP99 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetBlockedQpsStd(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.BlockedQpsStd = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetCount(v int32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.Count = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetException(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.Exception = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetExceptionAvg(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ExceptionAvg = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetExceptionMax(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ExceptionMax = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetExceptionMin(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ExceptionMin = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetExceptionP75(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ExceptionP75 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetExceptionP95(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ExceptionP95 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetExceptionP99(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ExceptionP99 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetExceptionStd(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ExceptionStd = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetPassedQps(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.PassedQps = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetPassedQpsAvg(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.PassedQpsAvg = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetPassedQpsMax(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.PassedQpsMax = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetPassedQpsMin(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.PassedQpsMin = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetPassedQpsP75(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.PassedQpsP75 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetPassedQpsP95(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.PassedQpsP95 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetPassedQpsP99(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.PassedQpsP99 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetPassedQpsStd(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.PassedQpsStd = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetRt(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.Rt = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetRtAvg(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.RtAvg = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetRtMax(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.RtMax = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetRtMin(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.RtMin = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetRtP75(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.RtP75 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetRtP95(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.RtP95 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetRtP99(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.RtP99 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetRtStd(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.RtStd = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetSuccessQps(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.SuccessQps = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetSuccessQpsAvg(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.SuccessQpsAvg = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetSuccessQpsMax(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.SuccessQpsMax = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetSuccessQpsMin(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.SuccessQpsMin = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetSuccessQpsP75(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.SuccessQpsP75 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetSuccessQpsP95(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.SuccessQpsP95 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetSuccessQpsP99(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.SuccessQpsP99 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetSuccessQpsStd(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.SuccessQpsStd = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetThread(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.Thread = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetThreadAvg(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ThreadAvg = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetThreadMax(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ThreadMax = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetThreadMin(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ThreadMin = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetThreadP75(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ThreadP75 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetThreadP95(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ThreadP95 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetThreadP99(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ThreadP99 = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetThreadStd(v float32) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.ThreadStd = &v
	return s
}

func (s *GetMetricsOfAppResponseBodyDataInnerMetrics) SetTimestamp(v int64) *GetMetricsOfAppResponseBodyDataInnerMetrics {
	s.Timestamp = &v
	return s
}

type GetMetricsOfAppResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMetricsOfAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMetricsOfAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMetricsOfAppResponse) GoString() string {
	return s.String()
}

func (s *GetMetricsOfAppResponse) SetHeaders(v map[string]*string) *GetMetricsOfAppResponse {
	s.Headers = v
	return s
}

func (s *GetMetricsOfAppResponse) SetStatusCode(v int32) *GetMetricsOfAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetricsOfAppResponse) SetBody(v *GetMetricsOfAppResponseBody) *GetMetricsOfAppResponse {
	s.Body = v
	return s
}

type GetMetricsOfResourceRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Resource     *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetMetricsOfResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMetricsOfResourceRequest) GoString() string {
	return s.String()
}

func (s *GetMetricsOfResourceRequest) SetAhasRegionId(v string) *GetMetricsOfResourceRequest {
	s.AhasRegionId = &v
	return s
}

func (s *GetMetricsOfResourceRequest) SetAppName(v string) *GetMetricsOfResourceRequest {
	s.AppName = &v
	return s
}

func (s *GetMetricsOfResourceRequest) SetEndTime(v int64) *GetMetricsOfResourceRequest {
	s.EndTime = &v
	return s
}

func (s *GetMetricsOfResourceRequest) SetNamespace(v string) *GetMetricsOfResourceRequest {
	s.Namespace = &v
	return s
}

func (s *GetMetricsOfResourceRequest) SetResource(v string) *GetMetricsOfResourceRequest {
	s.Resource = &v
	return s
}

func (s *GetMetricsOfResourceRequest) SetStartTime(v int64) *GetMetricsOfResourceRequest {
	s.StartTime = &v
	return s
}

type GetMetricsOfResourceResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetMetricsOfResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetricsOfResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMetricsOfResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetricsOfResourceResponseBody) SetCode(v string) *GetMetricsOfResourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetMetricsOfResourceResponseBody) SetData(v *GetMetricsOfResourceResponseBodyData) *GetMetricsOfResourceResponseBody {
	s.Data = v
	return s
}

func (s *GetMetricsOfResourceResponseBody) SetMessage(v string) *GetMetricsOfResourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetMetricsOfResourceResponseBody) SetRequestId(v string) *GetMetricsOfResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetricsOfResourceResponseBody) SetSuccess(v bool) *GetMetricsOfResourceResponseBody {
	s.Success = &v
	return s
}

type GetMetricsOfResourceResponseBodyData struct {
	AppName      *string                                             `json:"AppName,omitempty" xml:"AppName,omitempty"`
	InnerMetrics []*GetMetricsOfResourceResponseBodyDataInnerMetrics `json:"InnerMetrics,omitempty" xml:"InnerMetrics,omitempty" type:"Repeated"`
	Namespace    *string                                             `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Resource     *string                                             `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s GetMetricsOfResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMetricsOfResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetricsOfResourceResponseBodyData) SetAppName(v string) *GetMetricsOfResourceResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyData) SetInnerMetrics(v []*GetMetricsOfResourceResponseBodyDataInnerMetrics) *GetMetricsOfResourceResponseBodyData {
	s.InnerMetrics = v
	return s
}

func (s *GetMetricsOfResourceResponseBodyData) SetNamespace(v string) *GetMetricsOfResourceResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyData) SetResource(v string) *GetMetricsOfResourceResponseBodyData {
	s.Resource = &v
	return s
}

type GetMetricsOfResourceResponseBodyDataInnerMetrics struct {
	BlockedQps    *float32 `json:"BlockedQps,omitempty" xml:"BlockedQps,omitempty"`
	BlockedQpsAvg *float32 `json:"BlockedQpsAvg,omitempty" xml:"BlockedQpsAvg,omitempty"`
	BlockedQpsMax *float32 `json:"BlockedQpsMax,omitempty" xml:"BlockedQpsMax,omitempty"`
	BlockedQpsMin *float32 `json:"BlockedQpsMin,omitempty" xml:"BlockedQpsMin,omitempty"`
	BlockedQpsP75 *float32 `json:"BlockedQpsP75,omitempty" xml:"BlockedQpsP75,omitempty"`
	BlockedQpsP95 *float32 `json:"BlockedQpsP95,omitempty" xml:"BlockedQpsP95,omitempty"`
	BlockedQpsP99 *float32 `json:"BlockedQpsP99,omitempty" xml:"BlockedQpsP99,omitempty"`
	BlockedQpsStd *float32 `json:"BlockedQpsStd,omitempty" xml:"BlockedQpsStd,omitempty"`
	Count         *int32   `json:"Count,omitempty" xml:"Count,omitempty"`
	Exception     *float32 `json:"Exception,omitempty" xml:"Exception,omitempty"`
	ExceptionAvg  *float32 `json:"ExceptionAvg,omitempty" xml:"ExceptionAvg,omitempty"`
	ExceptionMax  *float32 `json:"ExceptionMax,omitempty" xml:"ExceptionMax,omitempty"`
	ExceptionMin  *float32 `json:"ExceptionMin,omitempty" xml:"ExceptionMin,omitempty"`
	ExceptionP75  *float32 `json:"ExceptionP75,omitempty" xml:"ExceptionP75,omitempty"`
	ExceptionP95  *float32 `json:"ExceptionP95,omitempty" xml:"ExceptionP95,omitempty"`
	ExceptionP99  *float32 `json:"ExceptionP99,omitempty" xml:"ExceptionP99,omitempty"`
	ExceptionStd  *float32 `json:"ExceptionStd,omitempty" xml:"ExceptionStd,omitempty"`
	PassedQps     *float32 `json:"PassedQps,omitempty" xml:"PassedQps,omitempty"`
	PassedQpsAvg  *float32 `json:"PassedQpsAvg,omitempty" xml:"PassedQpsAvg,omitempty"`
	PassedQpsMax  *float32 `json:"PassedQpsMax,omitempty" xml:"PassedQpsMax,omitempty"`
	PassedQpsMin  *float32 `json:"PassedQpsMin,omitempty" xml:"PassedQpsMin,omitempty"`
	PassedQpsP75  *float32 `json:"PassedQpsP75,omitempty" xml:"PassedQpsP75,omitempty"`
	PassedQpsP95  *float32 `json:"PassedQpsP95,omitempty" xml:"PassedQpsP95,omitempty"`
	PassedQpsP99  *float32 `json:"PassedQpsP99,omitempty" xml:"PassedQpsP99,omitempty"`
	PassedQpsStd  *float32 `json:"PassedQpsStd,omitempty" xml:"PassedQpsStd,omitempty"`
	Rt            *float32 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	RtAvg         *float32 `json:"RtAvg,omitempty" xml:"RtAvg,omitempty"`
	RtMax         *float32 `json:"RtMax,omitempty" xml:"RtMax,omitempty"`
	RtMin         *float32 `json:"RtMin,omitempty" xml:"RtMin,omitempty"`
	RtP75         *float32 `json:"RtP75,omitempty" xml:"RtP75,omitempty"`
	RtP95         *float32 `json:"RtP95,omitempty" xml:"RtP95,omitempty"`
	RtP99         *float32 `json:"RtP99,omitempty" xml:"RtP99,omitempty"`
	RtStd         *float32 `json:"RtStd,omitempty" xml:"RtStd,omitempty"`
	SuccessQps    *float32 `json:"SuccessQps,omitempty" xml:"SuccessQps,omitempty"`
	SuccessQpsAvg *float32 `json:"SuccessQpsAvg,omitempty" xml:"SuccessQpsAvg,omitempty"`
	SuccessQpsMax *float32 `json:"SuccessQpsMax,omitempty" xml:"SuccessQpsMax,omitempty"`
	SuccessQpsMin *float32 `json:"SuccessQpsMin,omitempty" xml:"SuccessQpsMin,omitempty"`
	SuccessQpsP75 *float32 `json:"SuccessQpsP75,omitempty" xml:"SuccessQpsP75,omitempty"`
	SuccessQpsP95 *float32 `json:"SuccessQpsP95,omitempty" xml:"SuccessQpsP95,omitempty"`
	SuccessQpsP99 *float32 `json:"SuccessQpsP99,omitempty" xml:"SuccessQpsP99,omitempty"`
	SuccessQpsStd *float32 `json:"SuccessQpsStd,omitempty" xml:"SuccessQpsStd,omitempty"`
	Thread        *float32 `json:"Thread,omitempty" xml:"Thread,omitempty"`
	ThreadAvg     *float32 `json:"ThreadAvg,omitempty" xml:"ThreadAvg,omitempty"`
	ThreadMax     *float32 `json:"ThreadMax,omitempty" xml:"ThreadMax,omitempty"`
	ThreadMin     *float32 `json:"ThreadMin,omitempty" xml:"ThreadMin,omitempty"`
	ThreadP75     *float32 `json:"ThreadP75,omitempty" xml:"ThreadP75,omitempty"`
	ThreadP95     *float32 `json:"ThreadP95,omitempty" xml:"ThreadP95,omitempty"`
	ThreadP99     *float32 `json:"ThreadP99,omitempty" xml:"ThreadP99,omitempty"`
	ThreadStd     *float32 `json:"ThreadStd,omitempty" xml:"ThreadStd,omitempty"`
	Timestamp     *int64   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMetricsOfResourceResponseBodyDataInnerMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetMetricsOfResourceResponseBodyDataInnerMetrics) GoString() string {
	return s.String()
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetBlockedQps(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.BlockedQps = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetBlockedQpsAvg(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.BlockedQpsAvg = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetBlockedQpsMax(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.BlockedQpsMax = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetBlockedQpsMin(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.BlockedQpsMin = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetBlockedQpsP75(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.BlockedQpsP75 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetBlockedQpsP95(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.BlockedQpsP95 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetBlockedQpsP99(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.BlockedQpsP99 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetBlockedQpsStd(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.BlockedQpsStd = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetCount(v int32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.Count = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetException(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.Exception = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetExceptionAvg(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ExceptionAvg = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetExceptionMax(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ExceptionMax = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetExceptionMin(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ExceptionMin = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetExceptionP75(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ExceptionP75 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetExceptionP95(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ExceptionP95 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetExceptionP99(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ExceptionP99 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetExceptionStd(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ExceptionStd = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetPassedQps(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.PassedQps = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetPassedQpsAvg(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.PassedQpsAvg = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetPassedQpsMax(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.PassedQpsMax = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetPassedQpsMin(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.PassedQpsMin = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetPassedQpsP75(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.PassedQpsP75 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetPassedQpsP95(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.PassedQpsP95 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetPassedQpsP99(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.PassedQpsP99 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetPassedQpsStd(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.PassedQpsStd = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetRt(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.Rt = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetRtAvg(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.RtAvg = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetRtMax(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.RtMax = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetRtMin(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.RtMin = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetRtP75(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.RtP75 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetRtP95(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.RtP95 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetRtP99(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.RtP99 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetRtStd(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.RtStd = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetSuccessQps(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.SuccessQps = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetSuccessQpsAvg(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.SuccessQpsAvg = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetSuccessQpsMax(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.SuccessQpsMax = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetSuccessQpsMin(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.SuccessQpsMin = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetSuccessQpsP75(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.SuccessQpsP75 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetSuccessQpsP95(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.SuccessQpsP95 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetSuccessQpsP99(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.SuccessQpsP99 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetSuccessQpsStd(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.SuccessQpsStd = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetThread(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.Thread = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetThreadAvg(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ThreadAvg = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetThreadMax(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ThreadMax = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetThreadMin(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ThreadMin = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetThreadP75(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ThreadP75 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetThreadP95(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ThreadP95 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetThreadP99(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ThreadP99 = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetThreadStd(v float32) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.ThreadStd = &v
	return s
}

func (s *GetMetricsOfResourceResponseBodyDataInnerMetrics) SetTimestamp(v int64) *GetMetricsOfResourceResponseBodyDataInnerMetrics {
	s.Timestamp = &v
	return s
}

type GetMetricsOfResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMetricsOfResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMetricsOfResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMetricsOfResourceResponse) GoString() string {
	return s.String()
}

func (s *GetMetricsOfResourceResponse) SetHeaders(v map[string]*string) *GetMetricsOfResourceResponse {
	s.Headers = v
	return s
}

func (s *GetMetricsOfResourceResponse) SetStatusCode(v int32) *GetMetricsOfResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetricsOfResourceResponse) SetBody(v *GetMetricsOfResourceResponseBody) *GetMetricsOfResourceResponse {
	s.Body = v
	return s
}

type GetSentinelAppSumMetricRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AhasRegionId   *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Namespace      *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetSentinelAppSumMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSentinelAppSumMetricRequest) GoString() string {
	return s.String()
}

func (s *GetSentinelAppSumMetricRequest) SetAcceptLanguage(v string) *GetSentinelAppSumMetricRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetSentinelAppSumMetricRequest) SetAhasRegionId(v string) *GetSentinelAppSumMetricRequest {
	s.AhasRegionId = &v
	return s
}

func (s *GetSentinelAppSumMetricRequest) SetAppName(v string) *GetSentinelAppSumMetricRequest {
	s.AppName = &v
	return s
}

func (s *GetSentinelAppSumMetricRequest) SetEndTime(v string) *GetSentinelAppSumMetricRequest {
	s.EndTime = &v
	return s
}

func (s *GetSentinelAppSumMetricRequest) SetNamespace(v string) *GetSentinelAppSumMetricRequest {
	s.Namespace = &v
	return s
}

func (s *GetSentinelAppSumMetricRequest) SetStartTime(v string) *GetSentinelAppSumMetricRequest {
	s.StartTime = &v
	return s
}

type GetSentinelAppSumMetricResponseBody struct {
	Code       *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	MetricData *GetSentinelAppSumMetricResponseBodyMetricData `json:"MetricData,omitempty" xml:"MetricData,omitempty" type:"Struct"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSentinelAppSumMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSentinelAppSumMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetSentinelAppSumMetricResponseBody) SetCode(v string) *GetSentinelAppSumMetricResponseBody {
	s.Code = &v
	return s
}

func (s *GetSentinelAppSumMetricResponseBody) SetMessage(v string) *GetSentinelAppSumMetricResponseBody {
	s.Message = &v
	return s
}

func (s *GetSentinelAppSumMetricResponseBody) SetMetricData(v *GetSentinelAppSumMetricResponseBodyMetricData) *GetSentinelAppSumMetricResponseBody {
	s.MetricData = v
	return s
}

func (s *GetSentinelAppSumMetricResponseBody) SetRequestId(v string) *GetSentinelAppSumMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSentinelAppSumMetricResponseBody) SetSuccess(v bool) *GetSentinelAppSumMetricResponseBody {
	s.Success = &v
	return s
}

type GetSentinelAppSumMetricResponseBodyMetricData struct {
	AppName      *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AvgRt        *float32 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	BlockCount   *float32 `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	MachineCount *int64   `json:"MachineCount,omitempty" xml:"MachineCount,omitempty"`
	Namespace    *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PassCount    *float32 `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	TotalCount   *float32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserId       *string  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSentinelAppSumMetricResponseBodyMetricData) String() string {
	return tea.Prettify(s)
}

func (s GetSentinelAppSumMetricResponseBodyMetricData) GoString() string {
	return s.String()
}

func (s *GetSentinelAppSumMetricResponseBodyMetricData) SetAppName(v string) *GetSentinelAppSumMetricResponseBodyMetricData {
	s.AppName = &v
	return s
}

func (s *GetSentinelAppSumMetricResponseBodyMetricData) SetAvgRt(v float32) *GetSentinelAppSumMetricResponseBodyMetricData {
	s.AvgRt = &v
	return s
}

func (s *GetSentinelAppSumMetricResponseBodyMetricData) SetBlockCount(v float32) *GetSentinelAppSumMetricResponseBodyMetricData {
	s.BlockCount = &v
	return s
}

func (s *GetSentinelAppSumMetricResponseBodyMetricData) SetMachineCount(v int64) *GetSentinelAppSumMetricResponseBodyMetricData {
	s.MachineCount = &v
	return s
}

func (s *GetSentinelAppSumMetricResponseBodyMetricData) SetNamespace(v string) *GetSentinelAppSumMetricResponseBodyMetricData {
	s.Namespace = &v
	return s
}

func (s *GetSentinelAppSumMetricResponseBodyMetricData) SetPassCount(v float32) *GetSentinelAppSumMetricResponseBodyMetricData {
	s.PassCount = &v
	return s
}

func (s *GetSentinelAppSumMetricResponseBodyMetricData) SetTotalCount(v float32) *GetSentinelAppSumMetricResponseBodyMetricData {
	s.TotalCount = &v
	return s
}

func (s *GetSentinelAppSumMetricResponseBodyMetricData) SetUserId(v string) *GetSentinelAppSumMetricResponseBodyMetricData {
	s.UserId = &v
	return s
}

type GetSentinelAppSumMetricResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSentinelAppSumMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSentinelAppSumMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSentinelAppSumMetricResponse) GoString() string {
	return s.String()
}

func (s *GetSentinelAppSumMetricResponse) SetHeaders(v map[string]*string) *GetSentinelAppSumMetricResponse {
	s.Headers = v
	return s
}

func (s *GetSentinelAppSumMetricResponse) SetStatusCode(v int32) *GetSentinelAppSumMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSentinelAppSumMetricResponse) SetBody(v *GetSentinelAppSumMetricResponseBody) *GetSentinelAppSumMetricResponse {
	s.Body = v
	return s
}

type GetUserApplicationGroupsRequest struct {
	AhasRegionId  *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	NameSpace     *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
}

func (s GetUserApplicationGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserApplicationGroupsRequest) GoString() string {
	return s.String()
}

func (s *GetUserApplicationGroupsRequest) SetAhasRegionId(v string) *GetUserApplicationGroupsRequest {
	s.AhasRegionId = &v
	return s
}

func (s *GetUserApplicationGroupsRequest) SetApplicationId(v string) *GetUserApplicationGroupsRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetUserApplicationGroupsRequest) SetNameSpace(v string) *GetUserApplicationGroupsRequest {
	s.NameSpace = &v
	return s
}

type GetUserApplicationGroupsResponseBody struct {
	AppGroups      []*string `json:"AppGroups,omitempty" xml:"AppGroups,omitempty" type:"Repeated"`
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserApplicationGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserApplicationGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserApplicationGroupsResponseBody) SetAppGroups(v []*string) *GetUserApplicationGroupsResponseBody {
	s.AppGroups = v
	return s
}

func (s *GetUserApplicationGroupsResponseBody) SetCode(v string) *GetUserApplicationGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserApplicationGroupsResponseBody) SetHttpStatusCode(v int32) *GetUserApplicationGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserApplicationGroupsResponseBody) SetMessage(v string) *GetUserApplicationGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserApplicationGroupsResponseBody) SetRequestId(v string) *GetUserApplicationGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserApplicationGroupsResponseBody) SetSuccess(v bool) *GetUserApplicationGroupsResponseBody {
	s.Success = &v
	return s
}

type GetUserApplicationGroupsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserApplicationGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserApplicationGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserApplicationGroupsResponse) GoString() string {
	return s.String()
}

func (s *GetUserApplicationGroupsResponse) SetHeaders(v map[string]*string) *GetUserApplicationGroupsResponse {
	s.Headers = v
	return s
}

func (s *GetUserApplicationGroupsResponse) SetStatusCode(v int32) *GetUserApplicationGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserApplicationGroupsResponse) SetBody(v *GetUserApplicationGroupsResponseBody) *GetUserApplicationGroupsResponse {
	s.Body = v
	return s
}

type GetUserApplicationsRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetUserApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserApplicationsRequest) GoString() string {
	return s.String()
}

func (s *GetUserApplicationsRequest) SetAhasRegionId(v string) *GetUserApplicationsRequest {
	s.AhasRegionId = &v
	return s
}

func (s *GetUserApplicationsRequest) SetNamespace(v string) *GetUserApplicationsRequest {
	s.Namespace = &v
	return s
}

type GetUserApplicationsResponseBody struct {
	AppNameAndIdPairs []*GetUserApplicationsResponseBodyAppNameAndIdPairs `json:"AppNameAndIdPairs,omitempty" xml:"AppNameAndIdPairs,omitempty" type:"Repeated"`
	Code              *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode    *int32                                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message           *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserApplicationsResponseBody) SetAppNameAndIdPairs(v []*GetUserApplicationsResponseBodyAppNameAndIdPairs) *GetUserApplicationsResponseBody {
	s.AppNameAndIdPairs = v
	return s
}

func (s *GetUserApplicationsResponseBody) SetCode(v string) *GetUserApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserApplicationsResponseBody) SetHttpStatusCode(v int32) *GetUserApplicationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserApplicationsResponseBody) SetMessage(v string) *GetUserApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserApplicationsResponseBody) SetRequestId(v string) *GetUserApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserApplicationsResponseBody) SetSuccess(v bool) *GetUserApplicationsResponseBody {
	s.Success = &v
	return s
}

type GetUserApplicationsResponseBodyAppNameAndIdPairs struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName   *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType   *int32  `json:"AppType,omitempty" xml:"AppType,omitempty"`
	ScopeType *int32  `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s GetUserApplicationsResponseBodyAppNameAndIdPairs) String() string {
	return tea.Prettify(s)
}

func (s GetUserApplicationsResponseBodyAppNameAndIdPairs) GoString() string {
	return s.String()
}

func (s *GetUserApplicationsResponseBodyAppNameAndIdPairs) SetAppId(v string) *GetUserApplicationsResponseBodyAppNameAndIdPairs {
	s.AppId = &v
	return s
}

func (s *GetUserApplicationsResponseBodyAppNameAndIdPairs) SetAppName(v string) *GetUserApplicationsResponseBodyAppNameAndIdPairs {
	s.AppName = &v
	return s
}

func (s *GetUserApplicationsResponseBodyAppNameAndIdPairs) SetAppType(v int32) *GetUserApplicationsResponseBodyAppNameAndIdPairs {
	s.AppType = &v
	return s
}

func (s *GetUserApplicationsResponseBodyAppNameAndIdPairs) SetScopeType(v int32) *GetUserApplicationsResponseBodyAppNameAndIdPairs {
	s.ScopeType = &v
	return s
}

type GetUserApplicationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserApplicationsResponse) GoString() string {
	return s.String()
}

func (s *GetUserApplicationsResponse) SetHeaders(v map[string]*string) *GetUserApplicationsResponse {
	s.Headers = v
	return s
}

func (s *GetUserApplicationsResponse) SetStatusCode(v int32) *GetUserApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserApplicationsResponse) SetBody(v *GetUserApplicationsResponseBody) *GetUserApplicationsResponse {
	s.Body = v
	return s
}

type GetUserWorkspaceRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetUserWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetUserWorkspaceRequest) SetAhasRegionId(v string) *GetUserWorkspaceRequest {
	s.AhasRegionId = &v
	return s
}

func (s *GetUserWorkspaceRequest) SetNamespace(v string) *GetUserWorkspaceRequest {
	s.Namespace = &v
	return s
}

type GetUserWorkspaceResponseBody struct {
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	WorkspaceList  []*GetUserWorkspaceResponseBodyWorkspaceList `json:"WorkspaceList,omitempty" xml:"WorkspaceList,omitempty" type:"Repeated"`
}

func (s GetUserWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserWorkspaceResponseBody) SetCode(v string) *GetUserWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserWorkspaceResponseBody) SetHttpStatusCode(v int32) *GetUserWorkspaceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserWorkspaceResponseBody) SetMessage(v string) *GetUserWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserWorkspaceResponseBody) SetRequestId(v string) *GetUserWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserWorkspaceResponseBody) SetSuccess(v bool) *GetUserWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserWorkspaceResponseBody) SetWorkspaceList(v []*GetUserWorkspaceResponseBodyWorkspaceList) *GetUserWorkspaceResponseBody {
	s.WorkspaceList = v
	return s
}

type GetUserWorkspaceResponseBodyWorkspaceList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetUserWorkspaceResponseBodyWorkspaceList) String() string {
	return tea.Prettify(s)
}

func (s GetUserWorkspaceResponseBodyWorkspaceList) GoString() string {
	return s.String()
}

func (s *GetUserWorkspaceResponseBodyWorkspaceList) SetDescription(v string) *GetUserWorkspaceResponseBodyWorkspaceList {
	s.Description = &v
	return s
}

func (s *GetUserWorkspaceResponseBodyWorkspaceList) SetName(v string) *GetUserWorkspaceResponseBodyWorkspaceList {
	s.Name = &v
	return s
}

func (s *GetUserWorkspaceResponseBodyWorkspaceList) SetType(v int32) *GetUserWorkspaceResponseBodyWorkspaceList {
	s.Type = &v
	return s
}

func (s *GetUserWorkspaceResponseBodyWorkspaceList) SetWorkspaceId(v string) *GetUserWorkspaceResponseBodyWorkspaceList {
	s.WorkspaceId = &v
	return s
}

type GetUserWorkspaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetUserWorkspaceResponse) SetHeaders(v map[string]*string) *GetUserWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetUserWorkspaceResponse) SetStatusCode(v int32) *GetUserWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserWorkspaceResponse) SetBody(v *GetUserWorkspaceResponseBody) *GetUserWorkspaceResponse {
	s.Body = v
	return s
}

type ListActiveAppsRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppType      *int32  `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListActiveAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListActiveAppsRequest) GoString() string {
	return s.String()
}

func (s *ListActiveAppsRequest) SetAhasRegionId(v string) *ListActiveAppsRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ListActiveAppsRequest) SetAppType(v int32) *ListActiveAppsRequest {
	s.AppType = &v
	return s
}

func (s *ListActiveAppsRequest) SetNamespace(v string) *ListActiveAppsRequest {
	s.Namespace = &v
	return s
}

type ListActiveAppsResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListActiveAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListActiveAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListActiveAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListActiveAppsResponseBody) SetCode(v string) *ListActiveAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListActiveAppsResponseBody) SetData(v []*ListActiveAppsResponseBodyData) *ListActiveAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListActiveAppsResponseBody) SetMessage(v string) *ListActiveAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListActiveAppsResponseBody) SetRequestId(v string) *ListActiveAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListActiveAppsResponseBody) SetSuccess(v bool) *ListActiveAppsResponseBody {
	s.Success = &v
	return s
}

type ListActiveAppsResponseBodyData struct {
	AhasAppName        *string `json:"AhasAppName,omitempty" xml:"AhasAppName,omitempty"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType            *int32  `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CurrentLevel       *int32  `json:"CurrentLevel,omitempty" xml:"CurrentLevel,omitempty"`
	DirtyLevel         *int32  `json:"DirtyLevel,omitempty" xml:"DirtyLevel,omitempty"`
	LastHealthPingTime *int64  `json:"LastHealthPingTime,omitempty" xml:"LastHealthPingTime,omitempty"`
	Namespace          *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListActiveAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListActiveAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListActiveAppsResponseBodyData) SetAhasAppName(v string) *ListActiveAppsResponseBodyData {
	s.AhasAppName = &v
	return s
}

func (s *ListActiveAppsResponseBodyData) SetAppName(v string) *ListActiveAppsResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListActiveAppsResponseBodyData) SetAppType(v int32) *ListActiveAppsResponseBodyData {
	s.AppType = &v
	return s
}

func (s *ListActiveAppsResponseBodyData) SetCurrentLevel(v int32) *ListActiveAppsResponseBodyData {
	s.CurrentLevel = &v
	return s
}

func (s *ListActiveAppsResponseBodyData) SetDirtyLevel(v int32) *ListActiveAppsResponseBodyData {
	s.DirtyLevel = &v
	return s
}

func (s *ListActiveAppsResponseBodyData) SetLastHealthPingTime(v int64) *ListActiveAppsResponseBodyData {
	s.LastHealthPingTime = &v
	return s
}

func (s *ListActiveAppsResponseBodyData) SetNamespace(v string) *ListActiveAppsResponseBodyData {
	s.Namespace = &v
	return s
}

type ListActiveAppsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListActiveAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListActiveAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListActiveAppsResponse) GoString() string {
	return s.String()
}

func (s *ListActiveAppsResponse) SetHeaders(v map[string]*string) *ListActiveAppsResponse {
	s.Headers = v
	return s
}

func (s *ListActiveAppsResponse) SetStatusCode(v int32) *ListActiveAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActiveAppsResponse) SetBody(v *ListActiveAppsResponseBody) *ListActiveAppsResponse {
	s.Body = v
	return s
}

type ListDegradeRulesOfAppRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDegradeRulesOfAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeRulesOfAppRequest) GoString() string {
	return s.String()
}

func (s *ListDegradeRulesOfAppRequest) SetAhasRegionId(v string) *ListDegradeRulesOfAppRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ListDegradeRulesOfAppRequest) SetAppName(v string) *ListDegradeRulesOfAppRequest {
	s.AppName = &v
	return s
}

func (s *ListDegradeRulesOfAppRequest) SetNamespace(v string) *ListDegradeRulesOfAppRequest {
	s.Namespace = &v
	return s
}

func (s *ListDegradeRulesOfAppRequest) SetPageIndex(v int32) *ListDegradeRulesOfAppRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDegradeRulesOfAppRequest) SetPageSize(v int32) *ListDegradeRulesOfAppRequest {
	s.PageSize = &v
	return s
}

type ListDegradeRulesOfAppResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListDegradeRulesOfAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDegradeRulesOfAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeRulesOfAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListDegradeRulesOfAppResponseBody) SetCode(v string) *ListDegradeRulesOfAppResponseBody {
	s.Code = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBody) SetData(v *ListDegradeRulesOfAppResponseBodyData) *ListDegradeRulesOfAppResponseBody {
	s.Data = v
	return s
}

func (s *ListDegradeRulesOfAppResponseBody) SetMessage(v string) *ListDegradeRulesOfAppResponseBody {
	s.Message = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBody) SetRequestId(v string) *ListDegradeRulesOfAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBody) SetSuccess(v bool) *ListDegradeRulesOfAppResponseBody {
	s.Success = &v
	return s
}

type ListDegradeRulesOfAppResponseBodyData struct {
	Datas      []*ListDegradeRulesOfAppResponseBodyDataDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	PageIndex  *int32                                        `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListDegradeRulesOfAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeRulesOfAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDegradeRulesOfAppResponseBodyData) SetDatas(v []*ListDegradeRulesOfAppResponseBodyDataDatas) *ListDegradeRulesOfAppResponseBodyData {
	s.Datas = v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyData) SetPageIndex(v int32) *ListDegradeRulesOfAppResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyData) SetPageSize(v int32) *ListDegradeRulesOfAppResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyData) SetTotalCount(v int32) *ListDegradeRulesOfAppResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyData) SetTotalPage(v int32) *ListDegradeRulesOfAppResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListDegradeRulesOfAppResponseBodyDataDatas struct {
	AppName                   *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable                    *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HalfOpenBaseAmountPerStep *int32   `json:"HalfOpenBaseAmountPerStep,omitempty" xml:"HalfOpenBaseAmountPerStep,omitempty"`
	HalfOpenRecoveryStepNum   *int32   `json:"HalfOpenRecoveryStepNum,omitempty" xml:"HalfOpenRecoveryStepNum,omitempty"`
	MinRequestAmount          *int32   `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	Namespace                 *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RecoveryTimeoutMs         *int32   `json:"RecoveryTimeoutMs,omitempty" xml:"RecoveryTimeoutMs,omitempty"`
	Resource                  *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                    *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SlowRtMs                  *int32   `json:"SlowRtMs,omitempty" xml:"SlowRtMs,omitempty"`
	StatDurationMs            *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Strategy                  *int32   `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Threshold                 *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListDegradeRulesOfAppResponseBodyDataDatas) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeRulesOfAppResponseBodyDataDatas) GoString() string {
	return s.String()
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetAppName(v string) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.AppName = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetEnable(v bool) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.Enable = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetHalfOpenBaseAmountPerStep(v int32) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetHalfOpenRecoveryStepNum(v int32) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetMinRequestAmount(v int32) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.MinRequestAmount = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetNamespace(v string) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.Namespace = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetRecoveryTimeoutMs(v int32) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.RecoveryTimeoutMs = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetResource(v string) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.Resource = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetRuleId(v int64) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.RuleId = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetSlowRtMs(v int32) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.SlowRtMs = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetStatDurationMs(v int32) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.StatDurationMs = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetStrategy(v int32) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.Strategy = &v
	return s
}

func (s *ListDegradeRulesOfAppResponseBodyDataDatas) SetThreshold(v float32) *ListDegradeRulesOfAppResponseBodyDataDatas {
	s.Threshold = &v
	return s
}

type ListDegradeRulesOfAppResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDegradeRulesOfAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDegradeRulesOfAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeRulesOfAppResponse) GoString() string {
	return s.String()
}

func (s *ListDegradeRulesOfAppResponse) SetHeaders(v map[string]*string) *ListDegradeRulesOfAppResponse {
	s.Headers = v
	return s
}

func (s *ListDegradeRulesOfAppResponse) SetStatusCode(v int32) *ListDegradeRulesOfAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDegradeRulesOfAppResponse) SetBody(v *ListDegradeRulesOfAppResponseBody) *ListDegradeRulesOfAppResponse {
	s.Body = v
	return s
}

type ListDegradeRulesOfResourceRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Resource     *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s ListDegradeRulesOfResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeRulesOfResourceRequest) GoString() string {
	return s.String()
}

func (s *ListDegradeRulesOfResourceRequest) SetAhasRegionId(v string) *ListDegradeRulesOfResourceRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ListDegradeRulesOfResourceRequest) SetAppName(v string) *ListDegradeRulesOfResourceRequest {
	s.AppName = &v
	return s
}

func (s *ListDegradeRulesOfResourceRequest) SetNamespace(v string) *ListDegradeRulesOfResourceRequest {
	s.Namespace = &v
	return s
}

func (s *ListDegradeRulesOfResourceRequest) SetPageIndex(v int32) *ListDegradeRulesOfResourceRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDegradeRulesOfResourceRequest) SetPageSize(v int32) *ListDegradeRulesOfResourceRequest {
	s.PageSize = &v
	return s
}

func (s *ListDegradeRulesOfResourceRequest) SetResource(v string) *ListDegradeRulesOfResourceRequest {
	s.Resource = &v
	return s
}

type ListDegradeRulesOfResourceResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListDegradeRulesOfResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDegradeRulesOfResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeRulesOfResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListDegradeRulesOfResourceResponseBody) SetCode(v string) *ListDegradeRulesOfResourceResponseBody {
	s.Code = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBody) SetData(v *ListDegradeRulesOfResourceResponseBodyData) *ListDegradeRulesOfResourceResponseBody {
	s.Data = v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBody) SetMessage(v string) *ListDegradeRulesOfResourceResponseBody {
	s.Message = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBody) SetRequestId(v string) *ListDegradeRulesOfResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBody) SetSuccess(v bool) *ListDegradeRulesOfResourceResponseBody {
	s.Success = &v
	return s
}

type ListDegradeRulesOfResourceResponseBodyData struct {
	Datas      []*ListDegradeRulesOfResourceResponseBodyDataDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	PageIndex  *int32                                             `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                             `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListDegradeRulesOfResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeRulesOfResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDegradeRulesOfResourceResponseBodyData) SetDatas(v []*ListDegradeRulesOfResourceResponseBodyDataDatas) *ListDegradeRulesOfResourceResponseBodyData {
	s.Datas = v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyData) SetPageIndex(v int32) *ListDegradeRulesOfResourceResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyData) SetPageSize(v int32) *ListDegradeRulesOfResourceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyData) SetTotalCount(v int32) *ListDegradeRulesOfResourceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyData) SetTotalPage(v int32) *ListDegradeRulesOfResourceResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListDegradeRulesOfResourceResponseBodyDataDatas struct {
	AppName                   *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable                    *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HalfOpenBaseAmountPerStep *int32   `json:"HalfOpenBaseAmountPerStep,omitempty" xml:"HalfOpenBaseAmountPerStep,omitempty"`
	HalfOpenRecoveryStepNum   *int32   `json:"HalfOpenRecoveryStepNum,omitempty" xml:"HalfOpenRecoveryStepNum,omitempty"`
	MinRequestAmount          *int32   `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	Namespace                 *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RecoveryTimeoutMs         *int32   `json:"RecoveryTimeoutMs,omitempty" xml:"RecoveryTimeoutMs,omitempty"`
	Resource                  *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                    *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SlowRtMs                  *int32   `json:"SlowRtMs,omitempty" xml:"SlowRtMs,omitempty"`
	StatDurationMs            *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Strategy                  *int32   `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Threshold                 *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListDegradeRulesOfResourceResponseBodyDataDatas) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeRulesOfResourceResponseBodyDataDatas) GoString() string {
	return s.String()
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetAppName(v string) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.AppName = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetEnable(v bool) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.Enable = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetHalfOpenBaseAmountPerStep(v int32) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetHalfOpenRecoveryStepNum(v int32) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetMinRequestAmount(v int32) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.MinRequestAmount = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetNamespace(v string) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.Namespace = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetRecoveryTimeoutMs(v int32) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.RecoveryTimeoutMs = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetResource(v string) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.Resource = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetRuleId(v int64) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.RuleId = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetSlowRtMs(v int32) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.SlowRtMs = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetStatDurationMs(v int32) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.StatDurationMs = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetStrategy(v int32) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.Strategy = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponseBodyDataDatas) SetThreshold(v float32) *ListDegradeRulesOfResourceResponseBodyDataDatas {
	s.Threshold = &v
	return s
}

type ListDegradeRulesOfResourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDegradeRulesOfResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDegradeRulesOfResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeRulesOfResourceResponse) GoString() string {
	return s.String()
}

func (s *ListDegradeRulesOfResourceResponse) SetHeaders(v map[string]*string) *ListDegradeRulesOfResourceResponse {
	s.Headers = v
	return s
}

func (s *ListDegradeRulesOfResourceResponse) SetStatusCode(v int32) *ListDegradeRulesOfResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDegradeRulesOfResourceResponse) SetBody(v *ListDegradeRulesOfResourceResponseBody) *ListDegradeRulesOfResourceResponse {
	s.Body = v
	return s
}

type ListExperimentMetasRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	NameSpace    *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	Page         *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListExperimentMetasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentMetasRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentMetasRequest) SetAhasRegionId(v string) *ListExperimentMetasRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ListExperimentMetasRequest) SetNameSpace(v string) *ListExperimentMetasRequest {
	s.NameSpace = &v
	return s
}

func (s *ListExperimentMetasRequest) SetPage(v int32) *ListExperimentMetasRequest {
	s.Page = &v
	return s
}

func (s *ListExperimentMetasRequest) SetSize(v int32) *ListExperimentMetasRequest {
	s.Size = &v
	return s
}

type ListExperimentMetasResponseBody struct {
	Code        *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Content     []*ListExperimentMetasResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	CurrentPage *int32                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Message     *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize    *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pages       *int32                                    `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Total       *int32                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListExperimentMetasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentMetasResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentMetasResponseBody) SetCode(v string) *ListExperimentMetasResponseBody {
	s.Code = &v
	return s
}

func (s *ListExperimentMetasResponseBody) SetContent(v []*ListExperimentMetasResponseBodyContent) *ListExperimentMetasResponseBody {
	s.Content = v
	return s
}

func (s *ListExperimentMetasResponseBody) SetCurrentPage(v int32) *ListExperimentMetasResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListExperimentMetasResponseBody) SetMessage(v string) *ListExperimentMetasResponseBody {
	s.Message = &v
	return s
}

func (s *ListExperimentMetasResponseBody) SetPageSize(v int32) *ListExperimentMetasResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListExperimentMetasResponseBody) SetPages(v int32) *ListExperimentMetasResponseBody {
	s.Pages = &v
	return s
}

func (s *ListExperimentMetasResponseBody) SetRequestId(v string) *ListExperimentMetasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentMetasResponseBody) SetSuccess(v bool) *ListExperimentMetasResponseBody {
	s.Success = &v
	return s
}

func (s *ListExperimentMetasResponseBody) SetTotal(v int32) *ListExperimentMetasResponseBody {
	s.Total = &v
	return s
}

type ListExperimentMetasResponseBodyContent struct {
	CreateTime   *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExperimentId *string   `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	State        *string   `json:"State,omitempty" xml:"State,omitempty"`
	Tags         []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListExperimentMetasResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentMetasResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListExperimentMetasResponseBodyContent) SetCreateTime(v int64) *ListExperimentMetasResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *ListExperimentMetasResponseBodyContent) SetExperimentId(v string) *ListExperimentMetasResponseBodyContent {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentMetasResponseBodyContent) SetName(v string) *ListExperimentMetasResponseBodyContent {
	s.Name = &v
	return s
}

func (s *ListExperimentMetasResponseBodyContent) SetState(v string) *ListExperimentMetasResponseBodyContent {
	s.State = &v
	return s
}

func (s *ListExperimentMetasResponseBodyContent) SetTags(v []*string) *ListExperimentMetasResponseBodyContent {
	s.Tags = v
	return s
}

type ListExperimentMetasResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListExperimentMetasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExperimentMetasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentMetasResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentMetasResponse) SetHeaders(v map[string]*string) *ListExperimentMetasResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentMetasResponse) SetStatusCode(v int32) *ListExperimentMetasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentMetasResponse) SetBody(v *ListExperimentMetasResponseBody) *ListExperimentMetasResponse {
	s.Body = v
	return s
}

type ListFlowRulesOfAppRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFlowRulesOfAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRulesOfAppRequest) GoString() string {
	return s.String()
}

func (s *ListFlowRulesOfAppRequest) SetAhasRegionId(v string) *ListFlowRulesOfAppRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ListFlowRulesOfAppRequest) SetAppName(v string) *ListFlowRulesOfAppRequest {
	s.AppName = &v
	return s
}

func (s *ListFlowRulesOfAppRequest) SetNamespace(v string) *ListFlowRulesOfAppRequest {
	s.Namespace = &v
	return s
}

func (s *ListFlowRulesOfAppRequest) SetPageIndex(v int32) *ListFlowRulesOfAppRequest {
	s.PageIndex = &v
	return s
}

func (s *ListFlowRulesOfAppRequest) SetPageSize(v int32) *ListFlowRulesOfAppRequest {
	s.PageSize = &v
	return s
}

type ListFlowRulesOfAppResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListFlowRulesOfAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFlowRulesOfAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRulesOfAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowRulesOfAppResponseBody) SetCode(v string) *ListFlowRulesOfAppResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBody) SetData(v *ListFlowRulesOfAppResponseBodyData) *ListFlowRulesOfAppResponseBody {
	s.Data = v
	return s
}

func (s *ListFlowRulesOfAppResponseBody) SetMessage(v string) *ListFlowRulesOfAppResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBody) SetRequestId(v string) *ListFlowRulesOfAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBody) SetSuccess(v bool) *ListFlowRulesOfAppResponseBody {
	s.Success = &v
	return s
}

type ListFlowRulesOfAppResponseBodyData struct {
	Datas      []*ListFlowRulesOfAppResponseBodyDataDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	PageIndex  *int32                                     `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                     `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListFlowRulesOfAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRulesOfAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlowRulesOfAppResponseBodyData) SetDatas(v []*ListFlowRulesOfAppResponseBodyDataDatas) *ListFlowRulesOfAppResponseBodyData {
	s.Datas = v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyData) SetPageIndex(v int32) *ListFlowRulesOfAppResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyData) SetPageSize(v int32) *ListFlowRulesOfAppResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyData) SetTotalCount(v int32) *ListFlowRulesOfAppResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyData) SetTotalPage(v int32) *ListFlowRulesOfAppResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListFlowRulesOfAppResponseBodyDataDatas struct {
	AppName                  *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ClusterEstimatedMaxQps   *float32 `json:"ClusterEstimatedMaxQps,omitempty" xml:"ClusterEstimatedMaxQps,omitempty"`
	ClusterFallbackStrategy  *int32   `json:"ClusterFallbackStrategy,omitempty" xml:"ClusterFallbackStrategy,omitempty"`
	ClusterFallbackThreshold *int32   `json:"ClusterFallbackThreshold,omitempty" xml:"ClusterFallbackThreshold,omitempty"`
	ClusterMode              *bool    `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	ClusterThresholdType     *int32   `json:"ClusterThresholdType,omitempty" xml:"ClusterThresholdType,omitempty"`
	ControlBehavior          *int32   `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable                   *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin              *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	MaxQueueingTimeMs        *int32   `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	Namespace                *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource              *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy         *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource                 *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                   *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationMs           *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Threshold                *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WarmUpPeriodSec          *int32   `json:"WarmUpPeriodSec,omitempty" xml:"WarmUpPeriodSec,omitempty"`
}

func (s ListFlowRulesOfAppResponseBodyDataDatas) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRulesOfAppResponseBodyDataDatas) GoString() string {
	return s.String()
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetAppName(v string) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.AppName = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetClusterEstimatedMaxQps(v float32) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.ClusterEstimatedMaxQps = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetClusterFallbackStrategy(v int32) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.ClusterFallbackStrategy = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetClusterFallbackThreshold(v int32) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.ClusterFallbackThreshold = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetClusterMode(v bool) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.ClusterMode = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetClusterThresholdType(v int32) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.ClusterThresholdType = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetControlBehavior(v int32) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.ControlBehavior = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetEnable(v bool) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.Enable = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetLimitOrigin(v string) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.LimitOrigin = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetMaxQueueingTimeMs(v int32) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetNamespace(v string) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.Namespace = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetRefResource(v string) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.RefResource = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetRelationStrategy(v int32) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.RelationStrategy = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetResource(v string) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.Resource = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetRuleId(v int64) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.RuleId = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetStatDurationMs(v int32) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.StatDurationMs = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetThreshold(v float32) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.Threshold = &v
	return s
}

func (s *ListFlowRulesOfAppResponseBodyDataDatas) SetWarmUpPeriodSec(v int32) *ListFlowRulesOfAppResponseBodyDataDatas {
	s.WarmUpPeriodSec = &v
	return s
}

type ListFlowRulesOfAppResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFlowRulesOfAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowRulesOfAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRulesOfAppResponse) GoString() string {
	return s.String()
}

func (s *ListFlowRulesOfAppResponse) SetHeaders(v map[string]*string) *ListFlowRulesOfAppResponse {
	s.Headers = v
	return s
}

func (s *ListFlowRulesOfAppResponse) SetStatusCode(v int32) *ListFlowRulesOfAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowRulesOfAppResponse) SetBody(v *ListFlowRulesOfAppResponseBody) *ListFlowRulesOfAppResponse {
	s.Body = v
	return s
}

type ListFlowRulesOfResourceRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Resource     *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s ListFlowRulesOfResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRulesOfResourceRequest) GoString() string {
	return s.String()
}

func (s *ListFlowRulesOfResourceRequest) SetAhasRegionId(v string) *ListFlowRulesOfResourceRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ListFlowRulesOfResourceRequest) SetAppName(v string) *ListFlowRulesOfResourceRequest {
	s.AppName = &v
	return s
}

func (s *ListFlowRulesOfResourceRequest) SetNamespace(v string) *ListFlowRulesOfResourceRequest {
	s.Namespace = &v
	return s
}

func (s *ListFlowRulesOfResourceRequest) SetPageIndex(v int32) *ListFlowRulesOfResourceRequest {
	s.PageIndex = &v
	return s
}

func (s *ListFlowRulesOfResourceRequest) SetPageSize(v int32) *ListFlowRulesOfResourceRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowRulesOfResourceRequest) SetResource(v string) *ListFlowRulesOfResourceRequest {
	s.Resource = &v
	return s
}

type ListFlowRulesOfResourceResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListFlowRulesOfResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFlowRulesOfResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRulesOfResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowRulesOfResourceResponseBody) SetCode(v string) *ListFlowRulesOfResourceResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBody) SetData(v *ListFlowRulesOfResourceResponseBodyData) *ListFlowRulesOfResourceResponseBody {
	s.Data = v
	return s
}

func (s *ListFlowRulesOfResourceResponseBody) SetMessage(v string) *ListFlowRulesOfResourceResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBody) SetRequestId(v string) *ListFlowRulesOfResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBody) SetSuccess(v bool) *ListFlowRulesOfResourceResponseBody {
	s.Success = &v
	return s
}

type ListFlowRulesOfResourceResponseBodyData struct {
	Datas      []*ListFlowRulesOfResourceResponseBodyDataDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	PageIndex  *int32                                          `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                          `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListFlowRulesOfResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRulesOfResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlowRulesOfResourceResponseBodyData) SetDatas(v []*ListFlowRulesOfResourceResponseBodyDataDatas) *ListFlowRulesOfResourceResponseBodyData {
	s.Datas = v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyData) SetPageIndex(v int32) *ListFlowRulesOfResourceResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyData) SetPageSize(v int32) *ListFlowRulesOfResourceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyData) SetTotalCount(v int32) *ListFlowRulesOfResourceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyData) SetTotalPage(v int32) *ListFlowRulesOfResourceResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListFlowRulesOfResourceResponseBodyDataDatas struct {
	AppName                  *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ClusterEstimatedMaxQps   *float32 `json:"ClusterEstimatedMaxQps,omitempty" xml:"ClusterEstimatedMaxQps,omitempty"`
	ClusterFallbackStrategy  *int32   `json:"ClusterFallbackStrategy,omitempty" xml:"ClusterFallbackStrategy,omitempty"`
	ClusterFallbackThreshold *int32   `json:"ClusterFallbackThreshold,omitempty" xml:"ClusterFallbackThreshold,omitempty"`
	ClusterMode              *bool    `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	ClusterThresholdType     *int32   `json:"ClusterThresholdType,omitempty" xml:"ClusterThresholdType,omitempty"`
	ControlBehavior          *int32   `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable                   *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin              *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	MaxQueueingTimeMs        *int32   `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	Namespace                *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource              *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy         *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource                 *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                   *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationMs           *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Threshold                *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WarmUpPeriodSec          *int32   `json:"WarmUpPeriodSec,omitempty" xml:"WarmUpPeriodSec,omitempty"`
}

func (s ListFlowRulesOfResourceResponseBodyDataDatas) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRulesOfResourceResponseBodyDataDatas) GoString() string {
	return s.String()
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetAppName(v string) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.AppName = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetClusterEstimatedMaxQps(v float32) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.ClusterEstimatedMaxQps = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetClusterFallbackStrategy(v int32) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.ClusterFallbackStrategy = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetClusterFallbackThreshold(v int32) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.ClusterFallbackThreshold = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetClusterMode(v bool) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.ClusterMode = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetClusterThresholdType(v int32) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.ClusterThresholdType = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetControlBehavior(v int32) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.ControlBehavior = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetEnable(v bool) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.Enable = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetLimitOrigin(v string) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.LimitOrigin = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetMaxQueueingTimeMs(v int32) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetNamespace(v string) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.Namespace = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetRefResource(v string) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.RefResource = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetRelationStrategy(v int32) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.RelationStrategy = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetResource(v string) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.Resource = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetRuleId(v int64) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.RuleId = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetStatDurationMs(v int32) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.StatDurationMs = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetThreshold(v float32) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.Threshold = &v
	return s
}

func (s *ListFlowRulesOfResourceResponseBodyDataDatas) SetWarmUpPeriodSec(v int32) *ListFlowRulesOfResourceResponseBodyDataDatas {
	s.WarmUpPeriodSec = &v
	return s
}

type ListFlowRulesOfResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFlowRulesOfResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowRulesOfResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRulesOfResourceResponse) GoString() string {
	return s.String()
}

func (s *ListFlowRulesOfResourceResponse) SetHeaders(v map[string]*string) *ListFlowRulesOfResourceResponse {
	s.Headers = v
	return s
}

func (s *ListFlowRulesOfResourceResponse) SetStatusCode(v int32) *ListFlowRulesOfResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowRulesOfResourceResponse) SetBody(v *ListFlowRulesOfResourceResponseBody) *ListFlowRulesOfResourceResponse {
	s.Body = v
	return s
}

type ListHotParamRulesOfAppRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHotParamRulesOfAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfAppRequest) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfAppRequest) SetAhasRegionId(v string) *ListHotParamRulesOfAppRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ListHotParamRulesOfAppRequest) SetAppName(v string) *ListHotParamRulesOfAppRequest {
	s.AppName = &v
	return s
}

func (s *ListHotParamRulesOfAppRequest) SetNamespace(v string) *ListHotParamRulesOfAppRequest {
	s.Namespace = &v
	return s
}

func (s *ListHotParamRulesOfAppRequest) SetPageIndex(v int32) *ListHotParamRulesOfAppRequest {
	s.PageIndex = &v
	return s
}

func (s *ListHotParamRulesOfAppRequest) SetPageSize(v int32) *ListHotParamRulesOfAppRequest {
	s.PageSize = &v
	return s
}

type ListHotParamRulesOfAppResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListHotParamRulesOfAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHotParamRulesOfAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfAppResponseBody) SetCode(v string) *ListHotParamRulesOfAppResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBody) SetData(v *ListHotParamRulesOfAppResponseBodyData) *ListHotParamRulesOfAppResponseBody {
	s.Data = v
	return s
}

func (s *ListHotParamRulesOfAppResponseBody) SetMessage(v string) *ListHotParamRulesOfAppResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBody) SetRequestId(v string) *ListHotParamRulesOfAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBody) SetSuccess(v bool) *ListHotParamRulesOfAppResponseBody {
	s.Success = &v
	return s
}

type ListHotParamRulesOfAppResponseBodyData struct {
	Datas      []*ListHotParamRulesOfAppResponseBodyDataDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	PageIndex  *int32                                         `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                         `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotParamRulesOfAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfAppResponseBodyData) SetDatas(v []*ListHotParamRulesOfAppResponseBodyDataDatas) *ListHotParamRulesOfAppResponseBodyData {
	s.Datas = v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyData) SetPageIndex(v int32) *ListHotParamRulesOfAppResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyData) SetPageSize(v int32) *ListHotParamRulesOfAppResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyData) SetTotalCount(v int32) *ListHotParamRulesOfAppResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyData) SetTotalPage(v int32) *ListHotParamRulesOfAppResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListHotParamRulesOfAppResponseBodyDataDatas struct {
	AppName           *string                                                         `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BurstCount        *int32                                                          `json:"BurstCount,omitempty" xml:"BurstCount,omitempty"`
	ControlBehavior   *int32                                                          `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable            *bool                                                           `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MaxQueueingTimeMs *int32                                                          `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	MetricType        *int32                                                          `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Namespace         *string                                                         `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParamFlowItemList []*ListHotParamRulesOfAppResponseBodyDataDatasParamFlowItemList `json:"ParamFlowItemList,omitempty" xml:"ParamFlowItemList,omitempty" type:"Repeated"`
	ParamIdx          *int32                                                          `json:"ParamIdx,omitempty" xml:"ParamIdx,omitempty"`
	Resource          *string                                                         `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId            *int64                                                          `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationSec   *int64                                                          `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	Threshold         *float32                                                        `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListHotParamRulesOfAppResponseBodyDataDatas) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfAppResponseBodyDataDatas) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetAppName(v string) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.AppName = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetBurstCount(v int32) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.BurstCount = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetControlBehavior(v int32) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.ControlBehavior = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetEnable(v bool) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.Enable = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetMaxQueueingTimeMs(v int32) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetMetricType(v int32) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.MetricType = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetNamespace(v string) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.Namespace = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetParamFlowItemList(v []*ListHotParamRulesOfAppResponseBodyDataDatasParamFlowItemList) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.ParamFlowItemList = v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetParamIdx(v int32) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.ParamIdx = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetResource(v string) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.Resource = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetRuleId(v int64) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.RuleId = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetStatDurationSec(v int64) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.StatDurationSec = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatas) SetThreshold(v float32) *ListHotParamRulesOfAppResponseBodyDataDatas {
	s.Threshold = &v
	return s
}

type ListHotParamRulesOfAppResponseBodyDataDatasParamFlowItemList struct {
	ItemType  *string  `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	ItemValue *string  `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListHotParamRulesOfAppResponseBodyDataDatasParamFlowItemList) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfAppResponseBodyDataDatasParamFlowItemList) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatasParamFlowItemList) SetItemType(v string) *ListHotParamRulesOfAppResponseBodyDataDatasParamFlowItemList {
	s.ItemType = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatasParamFlowItemList) SetItemValue(v string) *ListHotParamRulesOfAppResponseBodyDataDatasParamFlowItemList {
	s.ItemValue = &v
	return s
}

func (s *ListHotParamRulesOfAppResponseBodyDataDatasParamFlowItemList) SetThreshold(v float32) *ListHotParamRulesOfAppResponseBodyDataDatasParamFlowItemList {
	s.Threshold = &v
	return s
}

type ListHotParamRulesOfAppResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHotParamRulesOfAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotParamRulesOfAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfAppResponse) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfAppResponse) SetHeaders(v map[string]*string) *ListHotParamRulesOfAppResponse {
	s.Headers = v
	return s
}

func (s *ListHotParamRulesOfAppResponse) SetStatusCode(v int32) *ListHotParamRulesOfAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotParamRulesOfAppResponse) SetBody(v *ListHotParamRulesOfAppResponseBody) *ListHotParamRulesOfAppResponse {
	s.Body = v
	return s
}

type ListHotParamRulesOfResourceRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Resource     *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s ListHotParamRulesOfResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfResourceRequest) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfResourceRequest) SetAhasRegionId(v string) *ListHotParamRulesOfResourceRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ListHotParamRulesOfResourceRequest) SetAppName(v string) *ListHotParamRulesOfResourceRequest {
	s.AppName = &v
	return s
}

func (s *ListHotParamRulesOfResourceRequest) SetNamespace(v string) *ListHotParamRulesOfResourceRequest {
	s.Namespace = &v
	return s
}

func (s *ListHotParamRulesOfResourceRequest) SetPageIndex(v int32) *ListHotParamRulesOfResourceRequest {
	s.PageIndex = &v
	return s
}

func (s *ListHotParamRulesOfResourceRequest) SetPageSize(v int32) *ListHotParamRulesOfResourceRequest {
	s.PageSize = &v
	return s
}

func (s *ListHotParamRulesOfResourceRequest) SetResource(v string) *ListHotParamRulesOfResourceRequest {
	s.Resource = &v
	return s
}

type ListHotParamRulesOfResourceResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListHotParamRulesOfResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHotParamRulesOfResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfResourceResponseBody) SetCode(v string) *ListHotParamRulesOfResourceResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBody) SetData(v *ListHotParamRulesOfResourceResponseBodyData) *ListHotParamRulesOfResourceResponseBody {
	s.Data = v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBody) SetMessage(v string) *ListHotParamRulesOfResourceResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBody) SetRequestId(v string) *ListHotParamRulesOfResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBody) SetSuccess(v bool) *ListHotParamRulesOfResourceResponseBody {
	s.Success = &v
	return s
}

type ListHotParamRulesOfResourceResponseBodyData struct {
	Datas      []*ListHotParamRulesOfResourceResponseBodyDataDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	PageIndex  *int32                                              `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotParamRulesOfResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfResourceResponseBodyData) SetDatas(v []*ListHotParamRulesOfResourceResponseBodyDataDatas) *ListHotParamRulesOfResourceResponseBodyData {
	s.Datas = v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyData) SetPageIndex(v int32) *ListHotParamRulesOfResourceResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyData) SetPageSize(v int32) *ListHotParamRulesOfResourceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyData) SetTotalCount(v int32) *ListHotParamRulesOfResourceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyData) SetTotalPage(v int32) *ListHotParamRulesOfResourceResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListHotParamRulesOfResourceResponseBodyDataDatas struct {
	AppName           *string                                                              `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BurstCount        *int32                                                               `json:"BurstCount,omitempty" xml:"BurstCount,omitempty"`
	ControlBehavior   *int32                                                               `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable            *bool                                                                `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MaxQueueingTimeMs *int32                                                               `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	MetricType        *int32                                                               `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Namespace         *string                                                              `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParamFlowItemList []*ListHotParamRulesOfResourceResponseBodyDataDatasParamFlowItemList `json:"ParamFlowItemList,omitempty" xml:"ParamFlowItemList,omitempty" type:"Repeated"`
	ParamIdx          *int32                                                               `json:"ParamIdx,omitempty" xml:"ParamIdx,omitempty"`
	Resource          *string                                                              `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId            *int64                                                               `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationSec   *int64                                                               `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	Threshold         *float32                                                             `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListHotParamRulesOfResourceResponseBodyDataDatas) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfResourceResponseBodyDataDatas) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetAppName(v string) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.AppName = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetBurstCount(v int32) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.BurstCount = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetControlBehavior(v int32) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.ControlBehavior = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetEnable(v bool) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.Enable = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetMaxQueueingTimeMs(v int32) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetMetricType(v int32) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.MetricType = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetNamespace(v string) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.Namespace = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetParamFlowItemList(v []*ListHotParamRulesOfResourceResponseBodyDataDatasParamFlowItemList) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.ParamFlowItemList = v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetParamIdx(v int32) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.ParamIdx = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetResource(v string) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.Resource = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetRuleId(v int64) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.RuleId = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetStatDurationSec(v int64) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.StatDurationSec = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatas) SetThreshold(v float32) *ListHotParamRulesOfResourceResponseBodyDataDatas {
	s.Threshold = &v
	return s
}

type ListHotParamRulesOfResourceResponseBodyDataDatasParamFlowItemList struct {
	ItemType  *string  `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	ItemValue *string  `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListHotParamRulesOfResourceResponseBodyDataDatasParamFlowItemList) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfResourceResponseBodyDataDatasParamFlowItemList) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatasParamFlowItemList) SetItemType(v string) *ListHotParamRulesOfResourceResponseBodyDataDatasParamFlowItemList {
	s.ItemType = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatasParamFlowItemList) SetItemValue(v string) *ListHotParamRulesOfResourceResponseBodyDataDatasParamFlowItemList {
	s.ItemValue = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponseBodyDataDatasParamFlowItemList) SetThreshold(v float32) *ListHotParamRulesOfResourceResponseBodyDataDatasParamFlowItemList {
	s.Threshold = &v
	return s
}

type ListHotParamRulesOfResourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHotParamRulesOfResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotParamRulesOfResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotParamRulesOfResourceResponse) GoString() string {
	return s.String()
}

func (s *ListHotParamRulesOfResourceResponse) SetHeaders(v map[string]*string) *ListHotParamRulesOfResourceResponse {
	s.Headers = v
	return s
}

func (s *ListHotParamRulesOfResourceResponse) SetStatusCode(v int32) *ListHotParamRulesOfResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotParamRulesOfResourceResponse) SetBody(v *ListHotParamRulesOfResourceResponseBody) *ListHotParamRulesOfResourceResponse {
	s.Body = v
	return s
}

type ListIsolationRulesOfAppRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIsolationRulesOfAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIsolationRulesOfAppRequest) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesOfAppRequest) SetAhasRegionId(v string) *ListIsolationRulesOfAppRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ListIsolationRulesOfAppRequest) SetAppName(v string) *ListIsolationRulesOfAppRequest {
	s.AppName = &v
	return s
}

func (s *ListIsolationRulesOfAppRequest) SetNamespace(v string) *ListIsolationRulesOfAppRequest {
	s.Namespace = &v
	return s
}

func (s *ListIsolationRulesOfAppRequest) SetPageIndex(v int32) *ListIsolationRulesOfAppRequest {
	s.PageIndex = &v
	return s
}

func (s *ListIsolationRulesOfAppRequest) SetPageSize(v int32) *ListIsolationRulesOfAppRequest {
	s.PageSize = &v
	return s
}

type ListIsolationRulesOfAppResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListIsolationRulesOfAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIsolationRulesOfAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIsolationRulesOfAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesOfAppResponseBody) SetCode(v string) *ListIsolationRulesOfAppResponseBody {
	s.Code = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBody) SetData(v *ListIsolationRulesOfAppResponseBodyData) *ListIsolationRulesOfAppResponseBody {
	s.Data = v
	return s
}

func (s *ListIsolationRulesOfAppResponseBody) SetMessage(v string) *ListIsolationRulesOfAppResponseBody {
	s.Message = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBody) SetRequestId(v string) *ListIsolationRulesOfAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBody) SetSuccess(v bool) *ListIsolationRulesOfAppResponseBody {
	s.Success = &v
	return s
}

type ListIsolationRulesOfAppResponseBodyData struct {
	Datas      []*ListIsolationRulesOfAppResponseBodyDataDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	PageIndex  *int32                                          `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                          `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListIsolationRulesOfAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIsolationRulesOfAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesOfAppResponseBodyData) SetDatas(v []*ListIsolationRulesOfAppResponseBodyDataDatas) *ListIsolationRulesOfAppResponseBodyData {
	s.Datas = v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyData) SetPageIndex(v int32) *ListIsolationRulesOfAppResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyData) SetPageSize(v int32) *ListIsolationRulesOfAppResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyData) SetTotalCount(v int32) *ListIsolationRulesOfAppResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyData) SetTotalPage(v int32) *ListIsolationRulesOfAppResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListIsolationRulesOfAppResponseBodyDataDatas struct {
	AppName          *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable           *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin      *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	Namespace        *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource      *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource         *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId           *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold        *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListIsolationRulesOfAppResponseBodyDataDatas) String() string {
	return tea.Prettify(s)
}

func (s ListIsolationRulesOfAppResponseBodyDataDatas) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesOfAppResponseBodyDataDatas) SetAppName(v string) *ListIsolationRulesOfAppResponseBodyDataDatas {
	s.AppName = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyDataDatas) SetEnable(v bool) *ListIsolationRulesOfAppResponseBodyDataDatas {
	s.Enable = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyDataDatas) SetLimitOrigin(v string) *ListIsolationRulesOfAppResponseBodyDataDatas {
	s.LimitOrigin = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyDataDatas) SetNamespace(v string) *ListIsolationRulesOfAppResponseBodyDataDatas {
	s.Namespace = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyDataDatas) SetRefResource(v string) *ListIsolationRulesOfAppResponseBodyDataDatas {
	s.RefResource = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyDataDatas) SetRelationStrategy(v int32) *ListIsolationRulesOfAppResponseBodyDataDatas {
	s.RelationStrategy = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyDataDatas) SetResource(v string) *ListIsolationRulesOfAppResponseBodyDataDatas {
	s.Resource = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyDataDatas) SetRuleId(v int64) *ListIsolationRulesOfAppResponseBodyDataDatas {
	s.RuleId = &v
	return s
}

func (s *ListIsolationRulesOfAppResponseBodyDataDatas) SetThreshold(v float32) *ListIsolationRulesOfAppResponseBodyDataDatas {
	s.Threshold = &v
	return s
}

type ListIsolationRulesOfAppResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIsolationRulesOfAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIsolationRulesOfAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIsolationRulesOfAppResponse) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesOfAppResponse) SetHeaders(v map[string]*string) *ListIsolationRulesOfAppResponse {
	s.Headers = v
	return s
}

func (s *ListIsolationRulesOfAppResponse) SetStatusCode(v int32) *ListIsolationRulesOfAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIsolationRulesOfAppResponse) SetBody(v *ListIsolationRulesOfAppResponseBody) *ListIsolationRulesOfAppResponse {
	s.Body = v
	return s
}

type ListIsolationRulesOfResourceRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Resource     *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s ListIsolationRulesOfResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIsolationRulesOfResourceRequest) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesOfResourceRequest) SetAhasRegionId(v string) *ListIsolationRulesOfResourceRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ListIsolationRulesOfResourceRequest) SetAppName(v string) *ListIsolationRulesOfResourceRequest {
	s.AppName = &v
	return s
}

func (s *ListIsolationRulesOfResourceRequest) SetNamespace(v string) *ListIsolationRulesOfResourceRequest {
	s.Namespace = &v
	return s
}

func (s *ListIsolationRulesOfResourceRequest) SetPageIndex(v int32) *ListIsolationRulesOfResourceRequest {
	s.PageIndex = &v
	return s
}

func (s *ListIsolationRulesOfResourceRequest) SetPageSize(v int32) *ListIsolationRulesOfResourceRequest {
	s.PageSize = &v
	return s
}

func (s *ListIsolationRulesOfResourceRequest) SetResource(v string) *ListIsolationRulesOfResourceRequest {
	s.Resource = &v
	return s
}

type ListIsolationRulesOfResourceResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListIsolationRulesOfResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIsolationRulesOfResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIsolationRulesOfResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesOfResourceResponseBody) SetCode(v string) *ListIsolationRulesOfResourceResponseBody {
	s.Code = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBody) SetData(v *ListIsolationRulesOfResourceResponseBodyData) *ListIsolationRulesOfResourceResponseBody {
	s.Data = v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBody) SetMessage(v string) *ListIsolationRulesOfResourceResponseBody {
	s.Message = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBody) SetRequestId(v string) *ListIsolationRulesOfResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBody) SetSuccess(v bool) *ListIsolationRulesOfResourceResponseBody {
	s.Success = &v
	return s
}

type ListIsolationRulesOfResourceResponseBodyData struct {
	Datas      []*ListIsolationRulesOfResourceResponseBodyDataDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	PageIndex  *int32                                               `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListIsolationRulesOfResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIsolationRulesOfResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesOfResourceResponseBodyData) SetDatas(v []*ListIsolationRulesOfResourceResponseBodyDataDatas) *ListIsolationRulesOfResourceResponseBodyData {
	s.Datas = v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyData) SetPageIndex(v int32) *ListIsolationRulesOfResourceResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyData) SetPageSize(v int32) *ListIsolationRulesOfResourceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyData) SetTotalCount(v int32) *ListIsolationRulesOfResourceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyData) SetTotalPage(v int32) *ListIsolationRulesOfResourceResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListIsolationRulesOfResourceResponseBodyDataDatas struct {
	AppName          *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable           *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin      *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	Namespace        *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource      *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource         *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId           *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold        *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListIsolationRulesOfResourceResponseBodyDataDatas) String() string {
	return tea.Prettify(s)
}

func (s ListIsolationRulesOfResourceResponseBodyDataDatas) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesOfResourceResponseBodyDataDatas) SetAppName(v string) *ListIsolationRulesOfResourceResponseBodyDataDatas {
	s.AppName = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyDataDatas) SetEnable(v bool) *ListIsolationRulesOfResourceResponseBodyDataDatas {
	s.Enable = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyDataDatas) SetLimitOrigin(v string) *ListIsolationRulesOfResourceResponseBodyDataDatas {
	s.LimitOrigin = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyDataDatas) SetNamespace(v string) *ListIsolationRulesOfResourceResponseBodyDataDatas {
	s.Namespace = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyDataDatas) SetRefResource(v string) *ListIsolationRulesOfResourceResponseBodyDataDatas {
	s.RefResource = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyDataDatas) SetRelationStrategy(v int32) *ListIsolationRulesOfResourceResponseBodyDataDatas {
	s.RelationStrategy = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyDataDatas) SetResource(v string) *ListIsolationRulesOfResourceResponseBodyDataDatas {
	s.Resource = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyDataDatas) SetRuleId(v int64) *ListIsolationRulesOfResourceResponseBodyDataDatas {
	s.RuleId = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponseBodyDataDatas) SetThreshold(v float32) *ListIsolationRulesOfResourceResponseBodyDataDatas {
	s.Threshold = &v
	return s
}

type ListIsolationRulesOfResourceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIsolationRulesOfResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIsolationRulesOfResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIsolationRulesOfResourceResponse) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesOfResourceResponse) SetHeaders(v map[string]*string) *ListIsolationRulesOfResourceResponse {
	s.Headers = v
	return s
}

func (s *ListIsolationRulesOfResourceResponse) SetStatusCode(v int32) *ListIsolationRulesOfResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIsolationRulesOfResourceResponse) SetBody(v *ListIsolationRulesOfResourceResponseBody) *ListIsolationRulesOfResourceResponse {
	s.Body = v
	return s
}

type ListSystemRulesRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSystemRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSystemRulesRequest) GoString() string {
	return s.String()
}

func (s *ListSystemRulesRequest) SetAhasRegionId(v string) *ListSystemRulesRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ListSystemRulesRequest) SetAppName(v string) *ListSystemRulesRequest {
	s.AppName = &v
	return s
}

func (s *ListSystemRulesRequest) SetNamespace(v string) *ListSystemRulesRequest {
	s.Namespace = &v
	return s
}

func (s *ListSystemRulesRequest) SetPageIndex(v int32) *ListSystemRulesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListSystemRulesRequest) SetPageSize(v int32) *ListSystemRulesRequest {
	s.PageSize = &v
	return s
}

type ListSystemRulesResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListSystemRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSystemRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSystemRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemRulesResponseBody) SetCode(v string) *ListSystemRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListSystemRulesResponseBody) SetData(v *ListSystemRulesResponseBodyData) *ListSystemRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListSystemRulesResponseBody) SetMessage(v string) *ListSystemRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListSystemRulesResponseBody) SetRequestId(v string) *ListSystemRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemRulesResponseBody) SetSuccess(v bool) *ListSystemRulesResponseBody {
	s.Success = &v
	return s
}

type ListSystemRulesResponseBodyData struct {
	Datas      []*ListSystemRulesResponseBodyDataDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	PageIndex  *int32                                  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                  `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListSystemRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSystemRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSystemRulesResponseBodyData) SetDatas(v []*ListSystemRulesResponseBodyDataDatas) *ListSystemRulesResponseBodyData {
	s.Datas = v
	return s
}

func (s *ListSystemRulesResponseBodyData) SetPageIndex(v int32) *ListSystemRulesResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListSystemRulesResponseBodyData) SetPageSize(v int32) *ListSystemRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSystemRulesResponseBodyData) SetTotalCount(v int32) *ListSystemRulesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSystemRulesResponseBodyData) SetTotalPage(v int32) *ListSystemRulesResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListSystemRulesResponseBodyDataDatas struct {
	AppName    *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable     *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MetricType *int32   `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Namespace  *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RuleId     *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold  *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListSystemRulesResponseBodyDataDatas) String() string {
	return tea.Prettify(s)
}

func (s ListSystemRulesResponseBodyDataDatas) GoString() string {
	return s.String()
}

func (s *ListSystemRulesResponseBodyDataDatas) SetAppName(v string) *ListSystemRulesResponseBodyDataDatas {
	s.AppName = &v
	return s
}

func (s *ListSystemRulesResponseBodyDataDatas) SetEnable(v bool) *ListSystemRulesResponseBodyDataDatas {
	s.Enable = &v
	return s
}

func (s *ListSystemRulesResponseBodyDataDatas) SetMetricType(v int32) *ListSystemRulesResponseBodyDataDatas {
	s.MetricType = &v
	return s
}

func (s *ListSystemRulesResponseBodyDataDatas) SetNamespace(v string) *ListSystemRulesResponseBodyDataDatas {
	s.Namespace = &v
	return s
}

func (s *ListSystemRulesResponseBodyDataDatas) SetRuleId(v int64) *ListSystemRulesResponseBodyDataDatas {
	s.RuleId = &v
	return s
}

func (s *ListSystemRulesResponseBodyDataDatas) SetThreshold(v float32) *ListSystemRulesResponseBodyDataDatas {
	s.Threshold = &v
	return s
}

type ListSystemRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSystemRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSystemRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSystemRulesResponse) GoString() string {
	return s.String()
}

func (s *ListSystemRulesResponse) SetHeaders(v map[string]*string) *ListSystemRulesResponse {
	s.Headers = v
	return s
}

func (s *ListSystemRulesResponse) SetStatusCode(v int32) *ListSystemRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemRulesResponse) SetBody(v *ListSystemRulesResponseBody) *ListSystemRulesResponse {
	s.Body = v
	return s
}

type ModifyDegradeRuleRequest struct {
	AhasRegionId              *string  `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	HalfOpenBaseAmountPerStep *int32   `json:"HalfOpenBaseAmountPerStep,omitempty" xml:"HalfOpenBaseAmountPerStep,omitempty"`
	HalfOpenRecoveryStepNum   *int32   `json:"HalfOpenRecoveryStepNum,omitempty" xml:"HalfOpenRecoveryStepNum,omitempty"`
	MinRequestAmount          *int32   `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	RecoveryTimeoutMs         *int32   `json:"RecoveryTimeoutMs,omitempty" xml:"RecoveryTimeoutMs,omitempty"`
	RuleId                    *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SlowRtMs                  *int32   `json:"SlowRtMs,omitempty" xml:"SlowRtMs,omitempty"`
	StatDurationMs            *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Strategy                  *int32   `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Threshold                 *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyDegradeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDegradeRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyDegradeRuleRequest) SetAhasRegionId(v string) *ModifyDegradeRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ModifyDegradeRuleRequest) SetHalfOpenBaseAmountPerStep(v int32) *ModifyDegradeRuleRequest {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *ModifyDegradeRuleRequest) SetHalfOpenRecoveryStepNum(v int32) *ModifyDegradeRuleRequest {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *ModifyDegradeRuleRequest) SetMinRequestAmount(v int32) *ModifyDegradeRuleRequest {
	s.MinRequestAmount = &v
	return s
}

func (s *ModifyDegradeRuleRequest) SetRecoveryTimeoutMs(v int32) *ModifyDegradeRuleRequest {
	s.RecoveryTimeoutMs = &v
	return s
}

func (s *ModifyDegradeRuleRequest) SetRuleId(v int64) *ModifyDegradeRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyDegradeRuleRequest) SetSlowRtMs(v int32) *ModifyDegradeRuleRequest {
	s.SlowRtMs = &v
	return s
}

func (s *ModifyDegradeRuleRequest) SetStatDurationMs(v int32) *ModifyDegradeRuleRequest {
	s.StatDurationMs = &v
	return s
}

func (s *ModifyDegradeRuleRequest) SetStrategy(v int32) *ModifyDegradeRuleRequest {
	s.Strategy = &v
	return s
}

func (s *ModifyDegradeRuleRequest) SetThreshold(v float32) *ModifyDegradeRuleRequest {
	s.Threshold = &v
	return s
}

type ModifyDegradeRuleResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ModifyDegradeRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDegradeRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDegradeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDegradeRuleResponseBody) SetCode(v string) *ModifyDegradeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDegradeRuleResponseBody) SetData(v *ModifyDegradeRuleResponseBodyData) *ModifyDegradeRuleResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDegradeRuleResponseBody) SetMessage(v string) *ModifyDegradeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDegradeRuleResponseBody) SetRequestId(v string) *ModifyDegradeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDegradeRuleResponseBody) SetSuccess(v bool) *ModifyDegradeRuleResponseBody {
	s.Success = &v
	return s
}

type ModifyDegradeRuleResponseBodyData struct {
	AppName                   *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable                    *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HalfOpenBaseAmountPerStep *int32   `json:"HalfOpenBaseAmountPerStep,omitempty" xml:"HalfOpenBaseAmountPerStep,omitempty"`
	HalfOpenRecoveryStepNum   *int32   `json:"HalfOpenRecoveryStepNum,omitempty" xml:"HalfOpenRecoveryStepNum,omitempty"`
	MinRequestAmount          *int32   `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	Namespace                 *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RecoveryTimeoutMs         *int32   `json:"RecoveryTimeoutMs,omitempty" xml:"RecoveryTimeoutMs,omitempty"`
	Resource                  *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                    *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SlowRtMs                  *int32   `json:"SlowRtMs,omitempty" xml:"SlowRtMs,omitempty"`
	StatDurationMs            *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Strategy                  *int32   `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Threshold                 *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyDegradeRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyDegradeRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDegradeRuleResponseBodyData) SetAppName(v string) *ModifyDegradeRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetEnable(v bool) *ModifyDegradeRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetHalfOpenBaseAmountPerStep(v int32) *ModifyDegradeRuleResponseBodyData {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetHalfOpenRecoveryStepNum(v int32) *ModifyDegradeRuleResponseBodyData {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetMinRequestAmount(v int32) *ModifyDegradeRuleResponseBodyData {
	s.MinRequestAmount = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetNamespace(v string) *ModifyDegradeRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetRecoveryTimeoutMs(v int32) *ModifyDegradeRuleResponseBodyData {
	s.RecoveryTimeoutMs = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetResource(v string) *ModifyDegradeRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetRuleId(v int64) *ModifyDegradeRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetSlowRtMs(v int32) *ModifyDegradeRuleResponseBodyData {
	s.SlowRtMs = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetStatDurationMs(v int32) *ModifyDegradeRuleResponseBodyData {
	s.StatDurationMs = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetStrategy(v int32) *ModifyDegradeRuleResponseBodyData {
	s.Strategy = &v
	return s
}

func (s *ModifyDegradeRuleResponseBodyData) SetThreshold(v float32) *ModifyDegradeRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type ModifyDegradeRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDegradeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDegradeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDegradeRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyDegradeRuleResponse) SetHeaders(v map[string]*string) *ModifyDegradeRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyDegradeRuleResponse) SetStatusCode(v int32) *ModifyDegradeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDegradeRuleResponse) SetBody(v *ModifyDegradeRuleResponseBody) *ModifyDegradeRuleResponse {
	s.Body = v
	return s
}

type ModifyFlowRuleRequest struct {
	AhasRegionId      *string  `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	ControlBehavior   *string  `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	LimitOrigin       *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	MaxQueueingTimeMs *int32   `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	Namespace         *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource       *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy  *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	RuleId            *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold         *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WarmUpPeriodSec   *int32   `json:"WarmUpPeriodSec,omitempty" xml:"WarmUpPeriodSec,omitempty"`
}

func (s ModifyFlowRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowRuleRequest) SetAhasRegionId(v string) *ModifyFlowRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ModifyFlowRuleRequest) SetControlBehavior(v string) *ModifyFlowRuleRequest {
	s.ControlBehavior = &v
	return s
}

func (s *ModifyFlowRuleRequest) SetLimitOrigin(v string) *ModifyFlowRuleRequest {
	s.LimitOrigin = &v
	return s
}

func (s *ModifyFlowRuleRequest) SetMaxQueueingTimeMs(v int32) *ModifyFlowRuleRequest {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *ModifyFlowRuleRequest) SetNamespace(v string) *ModifyFlowRuleRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyFlowRuleRequest) SetRefResource(v string) *ModifyFlowRuleRequest {
	s.RefResource = &v
	return s
}

func (s *ModifyFlowRuleRequest) SetRelationStrategy(v int32) *ModifyFlowRuleRequest {
	s.RelationStrategy = &v
	return s
}

func (s *ModifyFlowRuleRequest) SetRuleId(v int64) *ModifyFlowRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyFlowRuleRequest) SetThreshold(v float32) *ModifyFlowRuleRequest {
	s.Threshold = &v
	return s
}

func (s *ModifyFlowRuleRequest) SetWarmUpPeriodSec(v int32) *ModifyFlowRuleRequest {
	s.WarmUpPeriodSec = &v
	return s
}

type ModifyFlowRuleResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ModifyFlowRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyFlowRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowRuleResponseBody) SetCode(v string) *ModifyFlowRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyFlowRuleResponseBody) SetData(v *ModifyFlowRuleResponseBodyData) *ModifyFlowRuleResponseBody {
	s.Data = v
	return s
}

func (s *ModifyFlowRuleResponseBody) SetMessage(v string) *ModifyFlowRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyFlowRuleResponseBody) SetRequestId(v string) *ModifyFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFlowRuleResponseBody) SetSuccess(v bool) *ModifyFlowRuleResponseBody {
	s.Success = &v
	return s
}

type ModifyFlowRuleResponseBodyData struct {
	AppName                  *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ClusterEstimatedMaxQps   *float32 `json:"ClusterEstimatedMaxQps,omitempty" xml:"ClusterEstimatedMaxQps,omitempty"`
	ClusterFallbackStrategy  *int32   `json:"ClusterFallbackStrategy,omitempty" xml:"ClusterFallbackStrategy,omitempty"`
	ClusterFallbackThreshold *int32   `json:"ClusterFallbackThreshold,omitempty" xml:"ClusterFallbackThreshold,omitempty"`
	ClusterMode              *bool    `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	ClusterThresholdType     *int32   `json:"ClusterThresholdType,omitempty" xml:"ClusterThresholdType,omitempty"`
	ControlBehavior          *int32   `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable                   *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin              *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	MaxQueueingTimeMs        *int32   `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	Namespace                *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource              *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy         *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource                 *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId                   *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationMs           *int32   `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	Threshold                *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WarmUpPeriodSec          *int32   `json:"WarmUpPeriodSec,omitempty" xml:"WarmUpPeriodSec,omitempty"`
}

func (s ModifyFlowRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyFlowRuleResponseBodyData) SetAppName(v string) *ModifyFlowRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetClusterEstimatedMaxQps(v float32) *ModifyFlowRuleResponseBodyData {
	s.ClusterEstimatedMaxQps = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetClusterFallbackStrategy(v int32) *ModifyFlowRuleResponseBodyData {
	s.ClusterFallbackStrategy = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetClusterFallbackThreshold(v int32) *ModifyFlowRuleResponseBodyData {
	s.ClusterFallbackThreshold = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetClusterMode(v bool) *ModifyFlowRuleResponseBodyData {
	s.ClusterMode = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetClusterThresholdType(v int32) *ModifyFlowRuleResponseBodyData {
	s.ClusterThresholdType = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetControlBehavior(v int32) *ModifyFlowRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetEnable(v bool) *ModifyFlowRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetLimitOrigin(v string) *ModifyFlowRuleResponseBodyData {
	s.LimitOrigin = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *ModifyFlowRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetNamespace(v string) *ModifyFlowRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetRefResource(v string) *ModifyFlowRuleResponseBodyData {
	s.RefResource = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetRelationStrategy(v int32) *ModifyFlowRuleResponseBodyData {
	s.RelationStrategy = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetResource(v string) *ModifyFlowRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetRuleId(v int64) *ModifyFlowRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetStatDurationMs(v int32) *ModifyFlowRuleResponseBodyData {
	s.StatDurationMs = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetThreshold(v float32) *ModifyFlowRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *ModifyFlowRuleResponseBodyData) SetWarmUpPeriodSec(v int32) *ModifyFlowRuleResponseBodyData {
	s.WarmUpPeriodSec = &v
	return s
}

type ModifyFlowRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowRuleResponse) SetHeaders(v map[string]*string) *ModifyFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowRuleResponse) SetStatusCode(v int32) *ModifyFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFlowRuleResponse) SetBody(v *ModifyFlowRuleResponseBody) *ModifyFlowRuleResponse {
	s.Body = v
	return s
}

type ModifyHotParamRuleRequest struct {
	AhasRegionId      *string  `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	BurstCount        *int32   `json:"BurstCount,omitempty" xml:"BurstCount,omitempty"`
	ControlBehavior   *int32   `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable            *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MaxQueueingTimeMs *int32   `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	MetricType        *int32   `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	ParamIdx          *int32   `json:"ParamIdx,omitempty" xml:"ParamIdx,omitempty"`
	RuleId            *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationSec   *int64   `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	Threshold         *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyHotParamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHotParamRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyHotParamRuleRequest) SetAhasRegionId(v string) *ModifyHotParamRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ModifyHotParamRuleRequest) SetBurstCount(v int32) *ModifyHotParamRuleRequest {
	s.BurstCount = &v
	return s
}

func (s *ModifyHotParamRuleRequest) SetControlBehavior(v int32) *ModifyHotParamRuleRequest {
	s.ControlBehavior = &v
	return s
}

func (s *ModifyHotParamRuleRequest) SetEnable(v bool) *ModifyHotParamRuleRequest {
	s.Enable = &v
	return s
}

func (s *ModifyHotParamRuleRequest) SetMaxQueueingTimeMs(v int32) *ModifyHotParamRuleRequest {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *ModifyHotParamRuleRequest) SetMetricType(v int32) *ModifyHotParamRuleRequest {
	s.MetricType = &v
	return s
}

func (s *ModifyHotParamRuleRequest) SetParamIdx(v int32) *ModifyHotParamRuleRequest {
	s.ParamIdx = &v
	return s
}

func (s *ModifyHotParamRuleRequest) SetRuleId(v int64) *ModifyHotParamRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyHotParamRuleRequest) SetStatDurationSec(v int64) *ModifyHotParamRuleRequest {
	s.StatDurationSec = &v
	return s
}

func (s *ModifyHotParamRuleRequest) SetThreshold(v float32) *ModifyHotParamRuleRequest {
	s.Threshold = &v
	return s
}

type ModifyHotParamRuleResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ModifyHotParamRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyHotParamRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHotParamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHotParamRuleResponseBody) SetCode(v string) *ModifyHotParamRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyHotParamRuleResponseBody) SetData(v *ModifyHotParamRuleResponseBodyData) *ModifyHotParamRuleResponseBody {
	s.Data = v
	return s
}

func (s *ModifyHotParamRuleResponseBody) SetMessage(v string) *ModifyHotParamRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyHotParamRuleResponseBody) SetRequestId(v string) *ModifyHotParamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHotParamRuleResponseBody) SetSuccess(v bool) *ModifyHotParamRuleResponseBody {
	s.Success = &v
	return s
}

type ModifyHotParamRuleResponseBodyData struct {
	AppName           *string                                                `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BurstCount        *int32                                                 `json:"BurstCount,omitempty" xml:"BurstCount,omitempty"`
	ControlBehavior   *int32                                                 `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	Enable            *bool                                                  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MaxQueueingTimeMs *int32                                                 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	MetricType        *int32                                                 `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Namespace         *string                                                `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParamFlowItemList []*ModifyHotParamRuleResponseBodyDataParamFlowItemList `json:"ParamFlowItemList,omitempty" xml:"ParamFlowItemList,omitempty" type:"Repeated"`
	ParamIdx          *int32                                                 `json:"ParamIdx,omitempty" xml:"ParamIdx,omitempty"`
	Resource          *string                                                `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId            *int64                                                 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StatDurationSec   *int64                                                 `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	Threshold         *float32                                               `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyHotParamRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyHotParamRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyHotParamRuleResponseBodyData) SetAppName(v string) *ModifyHotParamRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetBurstCount(v int32) *ModifyHotParamRuleResponseBodyData {
	s.BurstCount = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetControlBehavior(v int32) *ModifyHotParamRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetEnable(v bool) *ModifyHotParamRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *ModifyHotParamRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetMetricType(v int32) *ModifyHotParamRuleResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetNamespace(v string) *ModifyHotParamRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetParamFlowItemList(v []*ModifyHotParamRuleResponseBodyDataParamFlowItemList) *ModifyHotParamRuleResponseBodyData {
	s.ParamFlowItemList = v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetParamIdx(v int32) *ModifyHotParamRuleResponseBodyData {
	s.ParamIdx = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetResource(v string) *ModifyHotParamRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetRuleId(v int64) *ModifyHotParamRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetStatDurationSec(v int64) *ModifyHotParamRuleResponseBodyData {
	s.StatDurationSec = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyData) SetThreshold(v float32) *ModifyHotParamRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type ModifyHotParamRuleResponseBodyDataParamFlowItemList struct {
	ItemType  *string  `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	ItemValue *string  `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyHotParamRuleResponseBodyDataParamFlowItemList) String() string {
	return tea.Prettify(s)
}

func (s ModifyHotParamRuleResponseBodyDataParamFlowItemList) GoString() string {
	return s.String()
}

func (s *ModifyHotParamRuleResponseBodyDataParamFlowItemList) SetItemType(v string) *ModifyHotParamRuleResponseBodyDataParamFlowItemList {
	s.ItemType = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyDataParamFlowItemList) SetItemValue(v string) *ModifyHotParamRuleResponseBodyDataParamFlowItemList {
	s.ItemValue = &v
	return s
}

func (s *ModifyHotParamRuleResponseBodyDataParamFlowItemList) SetThreshold(v float32) *ModifyHotParamRuleResponseBodyDataParamFlowItemList {
	s.Threshold = &v
	return s
}

type ModifyHotParamRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyHotParamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHotParamRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHotParamRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyHotParamRuleResponse) SetHeaders(v map[string]*string) *ModifyHotParamRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyHotParamRuleResponse) SetStatusCode(v int32) *ModifyHotParamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHotParamRuleResponse) SetBody(v *ModifyHotParamRuleResponseBody) *ModifyHotParamRuleResponse {
	s.Body = v
	return s
}

type ModifyIsolationRuleRequest struct {
	AhasRegionId     *string  `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	LimitOrigin      *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	RefResource      *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	RuleId           *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold        *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyIsolationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIsolationRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyIsolationRuleRequest) SetAhasRegionId(v string) *ModifyIsolationRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ModifyIsolationRuleRequest) SetLimitOrigin(v string) *ModifyIsolationRuleRequest {
	s.LimitOrigin = &v
	return s
}

func (s *ModifyIsolationRuleRequest) SetRefResource(v string) *ModifyIsolationRuleRequest {
	s.RefResource = &v
	return s
}

func (s *ModifyIsolationRuleRequest) SetRelationStrategy(v int32) *ModifyIsolationRuleRequest {
	s.RelationStrategy = &v
	return s
}

func (s *ModifyIsolationRuleRequest) SetRuleId(v int64) *ModifyIsolationRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyIsolationRuleRequest) SetThreshold(v float32) *ModifyIsolationRuleRequest {
	s.Threshold = &v
	return s
}

type ModifyIsolationRuleResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ModifyIsolationRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyIsolationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIsolationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIsolationRuleResponseBody) SetCode(v string) *ModifyIsolationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyIsolationRuleResponseBody) SetData(v *ModifyIsolationRuleResponseBodyData) *ModifyIsolationRuleResponseBody {
	s.Data = v
	return s
}

func (s *ModifyIsolationRuleResponseBody) SetMessage(v string) *ModifyIsolationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyIsolationRuleResponseBody) SetRequestId(v string) *ModifyIsolationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIsolationRuleResponseBody) SetSuccess(v bool) *ModifyIsolationRuleResponseBody {
	s.Success = &v
	return s
}

type ModifyIsolationRuleResponseBodyData struct {
	AppName          *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Enable           *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitOrigin      *string  `json:"LimitOrigin,omitempty" xml:"LimitOrigin,omitempty"`
	Namespace        *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RefResource      *string  `json:"RefResource,omitempty" xml:"RefResource,omitempty"`
	RelationStrategy *int32   `json:"RelationStrategy,omitempty" xml:"RelationStrategy,omitempty"`
	Resource         *string  `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId           *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold        *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyIsolationRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyIsolationRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyIsolationRuleResponseBodyData) SetAppName(v string) *ModifyIsolationRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ModifyIsolationRuleResponseBodyData) SetEnable(v bool) *ModifyIsolationRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *ModifyIsolationRuleResponseBodyData) SetLimitOrigin(v string) *ModifyIsolationRuleResponseBodyData {
	s.LimitOrigin = &v
	return s
}

func (s *ModifyIsolationRuleResponseBodyData) SetNamespace(v string) *ModifyIsolationRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ModifyIsolationRuleResponseBodyData) SetRefResource(v string) *ModifyIsolationRuleResponseBodyData {
	s.RefResource = &v
	return s
}

func (s *ModifyIsolationRuleResponseBodyData) SetRelationStrategy(v int32) *ModifyIsolationRuleResponseBodyData {
	s.RelationStrategy = &v
	return s
}

func (s *ModifyIsolationRuleResponseBodyData) SetResource(v string) *ModifyIsolationRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *ModifyIsolationRuleResponseBodyData) SetRuleId(v int64) *ModifyIsolationRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *ModifyIsolationRuleResponseBodyData) SetThreshold(v float32) *ModifyIsolationRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type ModifyIsolationRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyIsolationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyIsolationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIsolationRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyIsolationRuleResponse) SetHeaders(v map[string]*string) *ModifyIsolationRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyIsolationRuleResponse) SetStatusCode(v int32) *ModifyIsolationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIsolationRuleResponse) SetBody(v *ModifyIsolationRuleResponseBody) *ModifyIsolationRuleResponse {
	s.Body = v
	return s
}

type ModifySystemRuleRequest struct {
	AhasRegionId *string  `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	RuleId       *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold    *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifySystemRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySystemRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifySystemRuleRequest) SetAhasRegionId(v string) *ModifySystemRuleRequest {
	s.AhasRegionId = &v
	return s
}

func (s *ModifySystemRuleRequest) SetRuleId(v int64) *ModifySystemRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifySystemRuleRequest) SetThreshold(v float32) *ModifySystemRuleRequest {
	s.Threshold = &v
	return s
}

type ModifySystemRuleResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ModifySystemRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySystemRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySystemRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySystemRuleResponseBody) SetCode(v string) *ModifySystemRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySystemRuleResponseBody) SetData(v *ModifySystemRuleResponseBodyData) *ModifySystemRuleResponseBody {
	s.Data = v
	return s
}

func (s *ModifySystemRuleResponseBody) SetMessage(v string) *ModifySystemRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySystemRuleResponseBody) SetRequestId(v string) *ModifySystemRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySystemRuleResponseBody) SetSuccess(v bool) *ModifySystemRuleResponseBody {
	s.Success = &v
	return s
}

type ModifySystemRuleResponseBodyData struct {
	Enable     *bool    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MetricType *int32   `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	RuleId     *int64   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Threshold  *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifySystemRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifySystemRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifySystemRuleResponseBodyData) SetEnable(v bool) *ModifySystemRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *ModifySystemRuleResponseBodyData) SetMetricType(v int32) *ModifySystemRuleResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *ModifySystemRuleResponseBodyData) SetRuleId(v int64) *ModifySystemRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *ModifySystemRuleResponseBodyData) SetThreshold(v float32) *ModifySystemRuleResponseBodyData {
	s.Threshold = &v
	return s
}

type ModifySystemRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySystemRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySystemRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySystemRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifySystemRuleResponse) SetHeaders(v map[string]*string) *ModifySystemRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifySystemRuleResponse) SetStatusCode(v int32) *ModifySystemRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySystemRuleResponse) SetBody(v *ModifySystemRuleResponseBody) *ModifySystemRuleResponse {
	s.Body = v
	return s
}

type OpenAhasServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s OpenAhasServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenAhasServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenAhasServiceRequest) SetOwnerId(v int64) *OpenAhasServiceRequest {
	s.OwnerId = &v
	return s
}

type OpenAhasServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenAhasServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenAhasServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenAhasServiceResponseBody) SetOrderId(v string) *OpenAhasServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenAhasServiceResponseBody) SetRequestId(v string) *OpenAhasServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenAhasServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenAhasServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenAhasServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenAhasServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenAhasServiceResponse) SetHeaders(v map[string]*string) *OpenAhasServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenAhasServiceResponse) SetStatusCode(v int32) *OpenAhasServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenAhasServiceResponse) SetBody(v *OpenAhasServiceResponseBody) *OpenAhasServiceResponse {
	s.Body = v
	return s
}

type PageableQueryExperimentTaskByClusterIdRequest struct {
	AhasRegionId       *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IncludeInvalidHost *bool   `json:"IncludeInvalidHost,omitempty" xml:"IncludeInvalidHost,omitempty"`
	Namespace          *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Page               *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Size               *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s PageableQueryExperimentTaskByClusterIdRequest) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByClusterIdRequest) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByClusterIdRequest) SetAhasRegionId(v string) *PageableQueryExperimentTaskByClusterIdRequest {
	s.AhasRegionId = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdRequest) SetClusterId(v string) *PageableQueryExperimentTaskByClusterIdRequest {
	s.ClusterId = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdRequest) SetIncludeInvalidHost(v bool) *PageableQueryExperimentTaskByClusterIdRequest {
	s.IncludeInvalidHost = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdRequest) SetNamespace(v string) *PageableQueryExperimentTaskByClusterIdRequest {
	s.Namespace = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdRequest) SetPage(v int32) *PageableQueryExperimentTaskByClusterIdRequest {
	s.Page = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdRequest) SetRegionId(v string) *PageableQueryExperimentTaskByClusterIdRequest {
	s.RegionId = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdRequest) SetSize(v int32) *PageableQueryExperimentTaskByClusterIdRequest {
	s.Size = &v
	return s
}

type PageableQueryExperimentTaskByClusterIdResponseBody struct {
	Code            *string                                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage     *int32                                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ExperimentTasks []*PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks `json:"ExperimentTasks,omitempty" xml:"ExperimentTasks,omitempty" type:"Repeated"`
	HttpStatusCode  *int32                                                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message         *string                                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize        *int32                                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pages           *int32                                                               `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId       *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool                                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Total           *int32                                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PageableQueryExperimentTaskByClusterIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByClusterIdResponseBody) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBody) SetCode(v string) *PageableQueryExperimentTaskByClusterIdResponseBody {
	s.Code = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBody) SetCurrentPage(v int32) *PageableQueryExperimentTaskByClusterIdResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBody) SetExperimentTasks(v []*PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) *PageableQueryExperimentTaskByClusterIdResponseBody {
	s.ExperimentTasks = v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBody) SetHttpStatusCode(v int32) *PageableQueryExperimentTaskByClusterIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBody) SetMessage(v string) *PageableQueryExperimentTaskByClusterIdResponseBody {
	s.Message = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBody) SetPageSize(v int32) *PageableQueryExperimentTaskByClusterIdResponseBody {
	s.PageSize = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBody) SetPages(v int32) *PageableQueryExperimentTaskByClusterIdResponseBody {
	s.Pages = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBody) SetRequestId(v string) *PageableQueryExperimentTaskByClusterIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBody) SetSuccess(v bool) *PageableQueryExperimentTaskByClusterIdResponseBody {
	s.Success = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBody) SetTotal(v int32) *PageableQueryExperimentTaskByClusterIdResponseBody {
	s.Total = &v
	return s
}

type PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks struct {
	Creator        *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	CurrentPhase   *string                                                                   `json:"CurrentPhase,omitempty" xml:"CurrentPhase,omitempty"`
	EndTime        *int64                                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExperimentId   *string                                                                   `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ExperimentName *string                                                                   `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
	ExtInfo        *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfo `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty" type:"Struct"`
	Message        *string                                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespace      *string                                                                   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Result         *string                                                                   `json:"Result,omitempty" xml:"Result,omitempty"`
	StartTime      *int64                                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State          *string                                                                   `json:"State,omitempty" xml:"State,omitempty"`
	TaskId         *string                                                                   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetCreator(v *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksCreator) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.Creator = v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetCurrentPhase(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.CurrentPhase = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetEndTime(v int64) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.EndTime = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetExperimentId(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.ExperimentId = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetExperimentName(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.ExperimentName = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetExtInfo(v *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfo) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.ExtInfo = v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetMessage(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.Message = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetNamespace(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.Namespace = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetResult(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.Result = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetStartTime(v int64) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.StartTime = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetState(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.State = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks) SetTaskId(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasks {
	s.TaskId = &v
	return s
}

type PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksCreator struct {
	SubUserId *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksCreator) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksCreator) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksCreator) SetSubUserId(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksCreator {
	s.SubUserId = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksCreator) SetUserId(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksCreator {
	s.UserId = &v
	return s
}

type PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfo struct {
	SchedulerConfig *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfoSchedulerConfig `json:"SchedulerConfig,omitempty" xml:"SchedulerConfig,omitempty" type:"Struct"`
}

func (s PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfo) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfo) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfo) SetSchedulerConfig(v *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfoSchedulerConfig) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfo {
	s.SchedulerConfig = v
	return s
}

type PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfoSchedulerConfig struct {
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	FixedTime      *string `json:"FixedTime,omitempty" xml:"FixedTime,omitempty"`
}

func (s PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfoSchedulerConfig) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfoSchedulerConfig) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfoSchedulerConfig) SetCronExpression(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfoSchedulerConfig {
	s.CronExpression = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfoSchedulerConfig) SetFixedTime(v string) *PageableQueryExperimentTaskByClusterIdResponseBodyExperimentTasksExtInfoSchedulerConfig {
	s.FixedTime = &v
	return s
}

type PageableQueryExperimentTaskByClusterIdResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PageableQueryExperimentTaskByClusterIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageableQueryExperimentTaskByClusterIdResponse) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByClusterIdResponse) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByClusterIdResponse) SetHeaders(v map[string]*string) *PageableQueryExperimentTaskByClusterIdResponse {
	s.Headers = v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponse) SetStatusCode(v int32) *PageableQueryExperimentTaskByClusterIdResponse {
	s.StatusCode = &v
	return s
}

func (s *PageableQueryExperimentTaskByClusterIdResponse) SetBody(v *PageableQueryExperimentTaskByClusterIdResponseBody) *PageableQueryExperimentTaskByClusterIdResponse {
	s.Body = v
	return s
}

type PageableQueryExperimentTaskByExperimentIdRequest struct {
	AhasRegionId *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Page         *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s PageableQueryExperimentTaskByExperimentIdRequest) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByExperimentIdRequest) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByExperimentIdRequest) SetAhasRegionId(v string) *PageableQueryExperimentTaskByExperimentIdRequest {
	s.AhasRegionId = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdRequest) SetExperimentId(v string) *PageableQueryExperimentTaskByExperimentIdRequest {
	s.ExperimentId = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdRequest) SetNamespace(v string) *PageableQueryExperimentTaskByExperimentIdRequest {
	s.Namespace = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdRequest) SetPage(v int32) *PageableQueryExperimentTaskByExperimentIdRequest {
	s.Page = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdRequest) SetRegionId(v string) *PageableQueryExperimentTaskByExperimentIdRequest {
	s.RegionId = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdRequest) SetSize(v int32) *PageableQueryExperimentTaskByExperimentIdRequest {
	s.Size = &v
	return s
}

type PageableQueryExperimentTaskByExperimentIdResponseBody struct {
	Code            *string                                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage     *int32                                                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ExperimentTasks []*PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks `json:"ExperimentTasks,omitempty" xml:"ExperimentTasks,omitempty" type:"Repeated"`
	HttpStatusCode  *int32                                                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message         *string                                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize        *int32                                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pages           *int32                                                                  `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId       *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool                                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Total           *int32                                                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PageableQueryExperimentTaskByExperimentIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByExperimentIdResponseBody) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBody) SetCode(v string) *PageableQueryExperimentTaskByExperimentIdResponseBody {
	s.Code = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBody) SetCurrentPage(v int32) *PageableQueryExperimentTaskByExperimentIdResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBody) SetExperimentTasks(v []*PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) *PageableQueryExperimentTaskByExperimentIdResponseBody {
	s.ExperimentTasks = v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBody) SetHttpStatusCode(v int32) *PageableQueryExperimentTaskByExperimentIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBody) SetMessage(v string) *PageableQueryExperimentTaskByExperimentIdResponseBody {
	s.Message = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBody) SetPageSize(v int32) *PageableQueryExperimentTaskByExperimentIdResponseBody {
	s.PageSize = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBody) SetPages(v int32) *PageableQueryExperimentTaskByExperimentIdResponseBody {
	s.Pages = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBody) SetRequestId(v string) *PageableQueryExperimentTaskByExperimentIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBody) SetSuccess(v bool) *PageableQueryExperimentTaskByExperimentIdResponseBody {
	s.Success = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBody) SetTotal(v int32) *PageableQueryExperimentTaskByExperimentIdResponseBody {
	s.Total = &v
	return s
}

type PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks struct {
	Creator        *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	CurrentPhase   *string                                                                      `json:"CurrentPhase,omitempty" xml:"CurrentPhase,omitempty"`
	EndTime        *int64                                                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExperimentId   *string                                                                      `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ExperimentName *string                                                                      `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
	ExtInfo        *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfo `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty" type:"Struct"`
	Message        *string                                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespace      *string                                                                      `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Result         *string                                                                      `json:"Result,omitempty" xml:"Result,omitempty"`
	StartTime      *int64                                                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State          *string                                                                      `json:"State,omitempty" xml:"State,omitempty"`
	TaskId         *string                                                                      `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetCreator(v *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksCreator) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.Creator = v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetCurrentPhase(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.CurrentPhase = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetEndTime(v int64) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.EndTime = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetExperimentId(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.ExperimentId = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetExperimentName(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.ExperimentName = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetExtInfo(v *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfo) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.ExtInfo = v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetMessage(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.Message = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetNamespace(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.Namespace = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetResult(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.Result = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetStartTime(v int64) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.StartTime = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetState(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.State = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks) SetTaskId(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasks {
	s.TaskId = &v
	return s
}

type PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksCreator struct {
	SubUserId *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksCreator) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksCreator) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksCreator) SetSubUserId(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksCreator {
	s.SubUserId = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksCreator) SetUserId(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksCreator {
	s.UserId = &v
	return s
}

type PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfo struct {
	SchedulerConfig *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfoSchedulerConfig `json:"SchedulerConfig,omitempty" xml:"SchedulerConfig,omitempty" type:"Struct"`
}

func (s PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfo) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfo) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfo) SetSchedulerConfig(v *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfoSchedulerConfig) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfo {
	s.SchedulerConfig = v
	return s
}

type PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfoSchedulerConfig struct {
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	FixedTime      *string `json:"FixedTime,omitempty" xml:"FixedTime,omitempty"`
}

func (s PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfoSchedulerConfig) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfoSchedulerConfig) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfoSchedulerConfig) SetCronExpression(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfoSchedulerConfig {
	s.CronExpression = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfoSchedulerConfig) SetFixedTime(v string) *PageableQueryExperimentTaskByExperimentIdResponseBodyExperimentTasksExtInfoSchedulerConfig {
	s.FixedTime = &v
	return s
}

type PageableQueryExperimentTaskByExperimentIdResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PageableQueryExperimentTaskByExperimentIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageableQueryExperimentTaskByExperimentIdResponse) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryExperimentTaskByExperimentIdResponse) GoString() string {
	return s.String()
}

func (s *PageableQueryExperimentTaskByExperimentIdResponse) SetHeaders(v map[string]*string) *PageableQueryExperimentTaskByExperimentIdResponse {
	s.Headers = v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponse) SetStatusCode(v int32) *PageableQueryExperimentTaskByExperimentIdResponse {
	s.StatusCode = &v
	return s
}

func (s *PageableQueryExperimentTaskByExperimentIdResponse) SetBody(v *PageableQueryExperimentTaskByExperimentIdResponseBody) *PageableQueryExperimentTaskByExperimentIdResponse {
	s.Body = v
	return s
}

type PageableQueryUserExperimentRequest struct {
	AhasRegionId *string   `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	Namespace    *string   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Page         *int32    `json:"Page,omitempty" xml:"Page,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Results      []*string `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	SearchKey    *string   `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	Size         *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	State        *string   `json:"State,omitempty" xml:"State,omitempty"`
	Tags         []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	WorkspaceId  *string   `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PageableQueryUserExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryUserExperimentRequest) GoString() string {
	return s.String()
}

func (s *PageableQueryUserExperimentRequest) SetAhasRegionId(v string) *PageableQueryUserExperimentRequest {
	s.AhasRegionId = &v
	return s
}

func (s *PageableQueryUserExperimentRequest) SetNamespace(v string) *PageableQueryUserExperimentRequest {
	s.Namespace = &v
	return s
}

func (s *PageableQueryUserExperimentRequest) SetPage(v int32) *PageableQueryUserExperimentRequest {
	s.Page = &v
	return s
}

func (s *PageableQueryUserExperimentRequest) SetRegionId(v string) *PageableQueryUserExperimentRequest {
	s.RegionId = &v
	return s
}

func (s *PageableQueryUserExperimentRequest) SetResults(v []*string) *PageableQueryUserExperimentRequest {
	s.Results = v
	return s
}

func (s *PageableQueryUserExperimentRequest) SetSearchKey(v string) *PageableQueryUserExperimentRequest {
	s.SearchKey = &v
	return s
}

func (s *PageableQueryUserExperimentRequest) SetSize(v int32) *PageableQueryUserExperimentRequest {
	s.Size = &v
	return s
}

func (s *PageableQueryUserExperimentRequest) SetState(v string) *PageableQueryUserExperimentRequest {
	s.State = &v
	return s
}

func (s *PageableQueryUserExperimentRequest) SetTags(v []*string) *PageableQueryUserExperimentRequest {
	s.Tags = v
	return s
}

func (s *PageableQueryUserExperimentRequest) SetWorkspaceId(v string) *PageableQueryUserExperimentRequest {
	s.WorkspaceId = &v
	return s
}

type PageableQueryUserExperimentResponseBody struct {
	Code           *string                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage    *int32                                                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ExperimentList []*PageableQueryUserExperimentResponseBodyExperimentList `json:"ExperimentList,omitempty" xml:"ExperimentList,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize       *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pages          *int32                                                   `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId      *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Total          *int32                                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PageableQueryUserExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryUserExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *PageableQueryUserExperimentResponseBody) SetCode(v string) *PageableQueryUserExperimentResponseBody {
	s.Code = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBody) SetCurrentPage(v int32) *PageableQueryUserExperimentResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBody) SetExperimentList(v []*PageableQueryUserExperimentResponseBodyExperimentList) *PageableQueryUserExperimentResponseBody {
	s.ExperimentList = v
	return s
}

func (s *PageableQueryUserExperimentResponseBody) SetHttpStatusCode(v int32) *PageableQueryUserExperimentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBody) SetMessage(v string) *PageableQueryUserExperimentResponseBody {
	s.Message = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBody) SetPageSize(v int32) *PageableQueryUserExperimentResponseBody {
	s.PageSize = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBody) SetPages(v int32) *PageableQueryUserExperimentResponseBody {
	s.Pages = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBody) SetRequestId(v string) *PageableQueryUserExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBody) SetSuccess(v bool) *PageableQueryUserExperimentResponseBody {
	s.Success = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBody) SetTotal(v int32) *PageableQueryUserExperimentResponseBody {
	s.Total = &v
	return s
}

type PageableQueryUserExperimentResponseBodyExperimentList struct {
	CreateTime   *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator      *string   `json:"Creator,omitempty" xml:"Creator,omitempty"`
	ExperimentId *string   `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	MiniApps     []*string `json:"MiniApps,omitempty" xml:"MiniApps,omitempty" type:"Repeated"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Permission   *int32    `json:"Permission,omitempty" xml:"Permission,omitempty"`
	Result       *string   `json:"Result,omitempty" xml:"Result,omitempty"`
	State        *string   `json:"State,omitempty" xml:"State,omitempty"`
	Tags         []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s PageableQueryUserExperimentResponseBodyExperimentList) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryUserExperimentResponseBodyExperimentList) GoString() string {
	return s.String()
}

func (s *PageableQueryUserExperimentResponseBodyExperimentList) SetCreateTime(v int64) *PageableQueryUserExperimentResponseBodyExperimentList {
	s.CreateTime = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBodyExperimentList) SetCreator(v string) *PageableQueryUserExperimentResponseBodyExperimentList {
	s.Creator = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBodyExperimentList) SetExperimentId(v string) *PageableQueryUserExperimentResponseBodyExperimentList {
	s.ExperimentId = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBodyExperimentList) SetMiniApps(v []*string) *PageableQueryUserExperimentResponseBodyExperimentList {
	s.MiniApps = v
	return s
}

func (s *PageableQueryUserExperimentResponseBodyExperimentList) SetName(v string) *PageableQueryUserExperimentResponseBodyExperimentList {
	s.Name = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBodyExperimentList) SetPermission(v int32) *PageableQueryUserExperimentResponseBodyExperimentList {
	s.Permission = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBodyExperimentList) SetResult(v string) *PageableQueryUserExperimentResponseBodyExperimentList {
	s.Result = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBodyExperimentList) SetState(v string) *PageableQueryUserExperimentResponseBodyExperimentList {
	s.State = &v
	return s
}

func (s *PageableQueryUserExperimentResponseBodyExperimentList) SetTags(v []*string) *PageableQueryUserExperimentResponseBodyExperimentList {
	s.Tags = v
	return s
}

type PageableQueryUserExperimentResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PageableQueryUserExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageableQueryUserExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s PageableQueryUserExperimentResponse) GoString() string {
	return s.String()
}

func (s *PageableQueryUserExperimentResponse) SetHeaders(v map[string]*string) *PageableQueryUserExperimentResponse {
	s.Headers = v
	return s
}

func (s *PageableQueryUserExperimentResponse) SetStatusCode(v int32) *PageableQueryUserExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *PageableQueryUserExperimentResponse) SetBody(v *PageableQueryUserExperimentResponseBody) *PageableQueryUserExperimentResponse {
	s.Body = v
	return s
}

type PushExperimentTaskRequest struct {
	AhasRegionId     *string `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	ExperimentTaskId *string `json:"ExperimentTaskId,omitempty" xml:"ExperimentTaskId,omitempty"`
	NameSpace        *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
}

func (s PushExperimentTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s PushExperimentTaskRequest) GoString() string {
	return s.String()
}

func (s *PushExperimentTaskRequest) SetAhasRegionId(v string) *PushExperimentTaskRequest {
	s.AhasRegionId = &v
	return s
}

func (s *PushExperimentTaskRequest) SetExperimentTaskId(v string) *PushExperimentTaskRequest {
	s.ExperimentTaskId = &v
	return s
}

func (s *PushExperimentTaskRequest) SetNameSpace(v string) *PushExperimentTaskRequest {
	s.NameSpace = &v
	return s
}

type PushExperimentTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushExperimentTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushExperimentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PushExperimentTaskResponseBody) SetCode(v string) *PushExperimentTaskResponseBody {
	s.Code = &v
	return s
}

func (s *PushExperimentTaskResponseBody) SetMessage(v string) *PushExperimentTaskResponseBody {
	s.Message = &v
	return s
}

func (s *PushExperimentTaskResponseBody) SetRequestId(v string) *PushExperimentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushExperimentTaskResponseBody) SetSuccess(v bool) *PushExperimentTaskResponseBody {
	s.Success = &v
	return s
}

type PushExperimentTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushExperimentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushExperimentTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s PushExperimentTaskResponse) GoString() string {
	return s.String()
}

func (s *PushExperimentTaskResponse) SetHeaders(v map[string]*string) *PushExperimentTaskResponse {
	s.Headers = v
	return s
}

func (s *PushExperimentTaskResponse) SetStatusCode(v int32) *PushExperimentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PushExperimentTaskResponse) SetBody(v *PushExperimentTaskResponseBody) *PushExperimentTaskResponse {
	s.Body = v
	return s
}

type UpdateExperimentRequest struct {
	AhasRegionId *string   `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	Definition   *string   `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description  *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ExperimentId *string   `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NameSpace    *string   `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	Tags         []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdateExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentRequest) SetAhasRegionId(v string) *UpdateExperimentRequest {
	s.AhasRegionId = &v
	return s
}

func (s *UpdateExperimentRequest) SetDefinition(v string) *UpdateExperimentRequest {
	s.Definition = &v
	return s
}

func (s *UpdateExperimentRequest) SetDescription(v string) *UpdateExperimentRequest {
	s.Description = &v
	return s
}

func (s *UpdateExperimentRequest) SetExperimentId(v string) *UpdateExperimentRequest {
	s.ExperimentId = &v
	return s
}

func (s *UpdateExperimentRequest) SetName(v string) *UpdateExperimentRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentRequest) SetNameSpace(v string) *UpdateExperimentRequest {
	s.NameSpace = &v
	return s
}

func (s *UpdateExperimentRequest) SetTags(v []*string) *UpdateExperimentRequest {
	s.Tags = v
	return s
}

type UpdateExperimentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentResponseBody) SetCode(v string) *UpdateExperimentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateExperimentResponseBody) SetHttpStatusCode(v int32) *UpdateExperimentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateExperimentResponseBody) SetMessage(v string) *UpdateExperimentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateExperimentResponseBody) SetRequestId(v string) *UpdateExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentResponseBody) SetSuccess(v bool) *UpdateExperimentResponseBody {
	s.Success = &v
	return s
}

type UpdateExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentResponse) SetHeaders(v map[string]*string) *UpdateExperimentResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentResponse) SetStatusCode(v int32) *UpdateExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentResponse) SetBody(v *UpdateExperimentResponseBody) *UpdateExperimentResponse {
	s.Body = v
	return s
}

type UpdateExperimentBasicInfoRequest struct {
	AhasRegionId *string   `json:"AhasRegionId,omitempty" xml:"AhasRegionId,omitempty"`
	Description  *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ExperimentId *string   `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NameSpace    *string   `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	Tags         []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdateExperimentBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentBasicInfoRequest) SetAhasRegionId(v string) *UpdateExperimentBasicInfoRequest {
	s.AhasRegionId = &v
	return s
}

func (s *UpdateExperimentBasicInfoRequest) SetDescription(v string) *UpdateExperimentBasicInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateExperimentBasicInfoRequest) SetExperimentId(v string) *UpdateExperimentBasicInfoRequest {
	s.ExperimentId = &v
	return s
}

func (s *UpdateExperimentBasicInfoRequest) SetName(v string) *UpdateExperimentBasicInfoRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentBasicInfoRequest) SetNameSpace(v string) *UpdateExperimentBasicInfoRequest {
	s.NameSpace = &v
	return s
}

func (s *UpdateExperimentBasicInfoRequest) SetTags(v []*string) *UpdateExperimentBasicInfoRequest {
	s.Tags = v
	return s
}

type UpdateExperimentBasicInfoResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateExperimentBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentBasicInfoResponseBody) SetCode(v string) *UpdateExperimentBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateExperimentBasicInfoResponseBody) SetHttpStatusCode(v int32) *UpdateExperimentBasicInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateExperimentBasicInfoResponseBody) SetMessage(v string) *UpdateExperimentBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateExperimentBasicInfoResponseBody) SetRequestId(v string) *UpdateExperimentBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentBasicInfoResponseBody) SetSuccess(v bool) *UpdateExperimentBasicInfoResponseBody {
	s.Success = &v
	return s
}

type UpdateExperimentBasicInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateExperimentBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateExperimentBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentBasicInfoResponse) SetHeaders(v map[string]*string) *UpdateExperimentBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentBasicInfoResponse) SetStatusCode(v int32) *UpdateExperimentBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentBasicInfoResponse) SetBody(v *UpdateExperimentBasicInfoResponseBody) *UpdateExperimentBasicInfoResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing":     tea.String("ahas.cn-beijing.aliyuncs.com"),
		"cn-zhangjiakou": tea.String("ahas.cn-zhangjiakou.aliyuncs.com"),
		"cn-hangzhou":    tea.String("ahas.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":    tea.String("ahas.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":    tea.String("ahas.cn-shenzhen.aliyuncs.com"),
		"ap-southeast-1": tea.String("ahas.ap-southeast-1.aliyuncs.com"),
		"cn-hongkong":    tea.String("ahas.cn-hongkong.aliyuncs.com"),
		"eu-central-1":   tea.String("ahas.eu-central-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ahas-openapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckExperimentRunnableWithOptions(request *CheckExperimentRunnableRequest, runtime *util.RuntimeOptions) (_result *CheckExperimentRunnableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckExperimentRunnable"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckExperimentRunnableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckExperimentRunnable(request *CheckExperimentRunnableRequest) (_result *CheckExperimentRunnableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckExperimentRunnableResponse{}
	_body, _err := client.CheckExperimentRunnableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDegradeRuleWithOptions(request *CreateDegradeRuleRequest, runtime *util.RuntimeOptions) (_result *CreateDegradeRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.HalfOpenBaseAmountPerStep)) {
		query["HalfOpenBaseAmountPerStep"] = request.HalfOpenBaseAmountPerStep
	}

	if !tea.BoolValue(util.IsUnset(request.HalfOpenRecoveryStepNum)) {
		query["HalfOpenRecoveryStepNum"] = request.HalfOpenRecoveryStepNum
	}

	if !tea.BoolValue(util.IsUnset(request.MinRequestAmount)) {
		query["MinRequestAmount"] = request.MinRequestAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveryTimeoutMs)) {
		query["RecoveryTimeoutMs"] = request.RecoveryTimeoutMs
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.SlowRtMs)) {
		query["SlowRtMs"] = request.SlowRtMs
	}

	if !tea.BoolValue(util.IsUnset(request.StatDurationMs)) {
		query["StatDurationMs"] = request.StatDurationMs
	}

	if !tea.BoolValue(util.IsUnset(request.Strategy)) {
		query["Strategy"] = request.Strategy
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDegradeRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDegradeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDegradeRule(request *CreateDegradeRuleRequest) (_result *CreateDegradeRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDegradeRuleResponse{}
	_body, _err := client.CreateDegradeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateExperimentWithOptions(request *CreateExperimentRequest, runtime *util.RuntimeOptions) (_result *CreateExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Definition)) {
		query["Definition"] = request.Definition
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExperiment"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateExperiment(request *CreateExperimentRequest) (_result *CreateExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateExperimentResponse{}
	_body, _err := client.CreateExperimentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowRuleWithOptions(request *CreateFlowRuleRequest, runtime *util.RuntimeOptions) (_result *CreateFlowRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ControlBehavior)) {
		query["ControlBehavior"] = request.ControlBehavior
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.LimitOrigin)) {
		query["LimitOrigin"] = request.LimitOrigin
	}

	if !tea.BoolValue(util.IsUnset(request.MaxQueueingTimeMs)) {
		query["MaxQueueingTimeMs"] = request.MaxQueueingTimeMs
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.RefResource)) {
		query["RefResource"] = request.RefResource
	}

	if !tea.BoolValue(util.IsUnset(request.RelationStrategy)) {
		query["RelationStrategy"] = request.RelationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	if !tea.BoolValue(util.IsUnset(request.WarmUpPeriodSec)) {
		query["WarmUpPeriodSec"] = request.WarmUpPeriodSec
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlowRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowRule(request *CreateFlowRuleRequest) (_result *CreateFlowRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowRuleResponse{}
	_body, _err := client.CreateFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHotParamItemsWithOptions(request *CreateHotParamItemsRequest, runtime *util.RuntimeOptions) (_result *CreateHotParamItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Items)) {
		query["Items"] = request.Items
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHotParamItems"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHotParamItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHotParamItems(request *CreateHotParamItemsRequest) (_result *CreateHotParamItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHotParamItemsResponse{}
	_body, _err := client.CreateHotParamItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHotParamRuleWithOptions(request *CreateHotParamRuleRequest, runtime *util.RuntimeOptions) (_result *CreateHotParamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BurstCount)) {
		query["BurstCount"] = request.BurstCount
	}

	if !tea.BoolValue(util.IsUnset(request.ControlBehavior)) {
		query["ControlBehavior"] = request.ControlBehavior
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.MaxQueueingTimeMs)) {
		query["MaxQueueingTimeMs"] = request.MaxQueueingTimeMs
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ParamIdx)) {
		query["ParamIdx"] = request.ParamIdx
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StatDurationSec)) {
		query["StatDurationSec"] = request.StatDurationSec
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHotParamRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHotParamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHotParamRule(request *CreateHotParamRuleRequest) (_result *CreateHotParamRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHotParamRuleResponse{}
	_body, _err := client.CreateHotParamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIsolationRuleWithOptions(request *CreateIsolationRuleRequest, runtime *util.RuntimeOptions) (_result *CreateIsolationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.LimitOrigin)) {
		query["LimitOrigin"] = request.LimitOrigin
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.RefResource)) {
		query["RefResource"] = request.RefResource
	}

	if !tea.BoolValue(util.IsUnset(request.RelationStrategy)) {
		query["RelationStrategy"] = request.RelationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIsolationRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIsolationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIsolationRule(request *CreateIsolationRuleRequest) (_result *CreateIsolationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIsolationRuleResponse{}
	_body, _err := client.CreateIsolationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSystemRuleWithOptions(request *CreateSystemRuleRequest, runtime *util.RuntimeOptions) (_result *CreateSystemRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSystemRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSystemRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSystemRule(request *CreateSystemRuleRequest) (_result *CreateSystemRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSystemRuleResponse{}
	_body, _err := client.CreateSystemRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDegradeRuleWithOptions(request *DeleteDegradeRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteDegradeRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDegradeRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDegradeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDegradeRule(request *DeleteDegradeRuleRequest) (_result *DeleteDegradeRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDegradeRuleResponse{}
	_body, _err := client.DeleteDegradeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowRuleWithOptions(request *DeleteFlowRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowRule(request *DeleteFlowRuleRequest) (_result *DeleteFlowRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowRuleResponse{}
	_body, _err := client.DeleteFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHotParamRuleWithOptions(request *DeleteHotParamRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteHotParamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHotParamRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHotParamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHotParamRule(request *DeleteHotParamRuleRequest) (_result *DeleteHotParamRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHotParamRuleResponse{}
	_body, _err := client.DeleteHotParamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIsolationRuleWithOptions(request *DeleteIsolationRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteIsolationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIsolationRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIsolationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIsolationRule(request *DeleteIsolationRuleRequest) (_result *DeleteIsolationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIsolationRuleResponse{}
	_body, _err := client.DeleteIsolationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSystemRuleWithOptions(request *DeleteSystemRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteSystemRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSystemRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSystemRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSystemRule(request *DeleteSystemRuleRequest) (_result *DeleteSystemRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSystemRuleResponse{}
	_body, _err := client.DeleteSystemRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableDegradeRuleWithOptions(request *DisableDegradeRuleRequest, runtime *util.RuntimeOptions) (_result *DisableDegradeRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableDegradeRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableDegradeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableDegradeRule(request *DisableDegradeRuleRequest) (_result *DisableDegradeRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableDegradeRuleResponse{}
	_body, _err := client.DisableDegradeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableFlowRuleWithOptions(request *DisableFlowRuleRequest, runtime *util.RuntimeOptions) (_result *DisableFlowRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableFlowRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableFlowRule(request *DisableFlowRuleRequest) (_result *DisableFlowRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableFlowRuleResponse{}
	_body, _err := client.DisableFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableHotParamRuleWithOptions(request *DisableHotParamRuleRequest, runtime *util.RuntimeOptions) (_result *DisableHotParamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableHotParamRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableHotParamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableHotParamRule(request *DisableHotParamRuleRequest) (_result *DisableHotParamRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableHotParamRuleResponse{}
	_body, _err := client.DisableHotParamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableIsolationRuleWithOptions(request *DisableIsolationRuleRequest, runtime *util.RuntimeOptions) (_result *DisableIsolationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableIsolationRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableIsolationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableIsolationRule(request *DisableIsolationRuleRequest) (_result *DisableIsolationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableIsolationRuleResponse{}
	_body, _err := client.DisableIsolationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableSystemRuleWithOptions(request *DisableSystemRuleRequest, runtime *util.RuntimeOptions) (_result *DisableSystemRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableSystemRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableSystemRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableSystemRule(request *DisableSystemRuleRequest) (_result *DisableSystemRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableSystemRuleResponse{}
	_body, _err := client.DisableSystemRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableDegradeRuleWithOptions(request *EnableDegradeRuleRequest, runtime *util.RuntimeOptions) (_result *EnableDegradeRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableDegradeRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableDegradeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableDegradeRule(request *EnableDegradeRuleRequest) (_result *EnableDegradeRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableDegradeRuleResponse{}
	_body, _err := client.EnableDegradeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableFlowRuleWithOptions(request *EnableFlowRuleRequest, runtime *util.RuntimeOptions) (_result *EnableFlowRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableFlowRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableFlowRule(request *EnableFlowRuleRequest) (_result *EnableFlowRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableFlowRuleResponse{}
	_body, _err := client.EnableFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableHotParamRuleWithOptions(request *EnableHotParamRuleRequest, runtime *util.RuntimeOptions) (_result *EnableHotParamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableHotParamRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableHotParamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableHotParamRule(request *EnableHotParamRuleRequest) (_result *EnableHotParamRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableHotParamRuleResponse{}
	_body, _err := client.EnableHotParamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableIsolationRuleWithOptions(request *EnableIsolationRuleRequest, runtime *util.RuntimeOptions) (_result *EnableIsolationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableIsolationRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableIsolationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableIsolationRule(request *EnableIsolationRuleRequest) (_result *EnableIsolationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableIsolationRuleResponse{}
	_body, _err := client.EnableIsolationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSystemRuleWithOptions(request *EnableSystemRuleRequest, runtime *util.RuntimeOptions) (_result *EnableSystemRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableSystemRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSystemRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSystemRule(request *EnableSystemRuleRequest) (_result *EnableSystemRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSystemRuleResponse{}
	_body, _err := client.EnableSystemRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteExperimentWithOptions(request *ExecuteExperimentRequest, runtime *util.RuntimeOptions) (_result *ExecuteExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Definition)) {
		query["Definition"] = request.Definition
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteExperiment"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteExperiment(request *ExecuteExperimentRequest) (_result *ExecuteExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteExperimentResponse{}
	_body, _err := client.ExecuteExperimentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FinishExperimentTaskWithOptions(request *FinishExperimentTaskRequest, runtime *util.RuntimeOptions) (_result *FinishExperimentTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentTaskId)) {
		query["ExperimentTaskId"] = request.ExperimentTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FinishExperimentTask"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishExperimentTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FinishExperimentTask(request *FinishExperimentTaskRequest) (_result *FinishExperimentTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FinishExperimentTaskResponse{}
	_body, _err := client.FinishExperimentTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetActivityTaskWithOptions(request *GetActivityTaskRequest, runtime *util.RuntimeOptions) (_result *GetActivityTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityTaskId)) {
		query["ActivityTaskId"] = request.ActivityTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentTaskId)) {
		query["ExperimentTaskId"] = request.ExperimentTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetActivityTask"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetActivityTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetActivityTask(request *GetActivityTaskRequest) (_result *GetActivityTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetActivityTaskResponse{}
	_body, _err := client.GetActivityTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExperimentMetaWithOptions(request *GetExperimentMetaRequest, runtime *util.RuntimeOptions) (_result *GetExperimentMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentMeta"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExperimentMeta(request *GetExperimentMetaRequest) (_result *GetExperimentMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExperimentMetaResponse{}
	_body, _err := client.GetExperimentMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExperimentTaskWithOptions(request *GetExperimentTaskRequest, runtime *util.RuntimeOptions) (_result *GetExperimentTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentTaskId)) {
		query["ExperimentTaskId"] = request.ExperimentTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentTask"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExperimentTask(request *GetExperimentTaskRequest) (_result *GetExperimentTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExperimentTaskResponse{}
	_body, _err := client.GetExperimentTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLicenseKeyWithOptions(request *GetLicenseKeyRequest, runtime *util.RuntimeOptions) (_result *GetLicenseKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLicenseKey"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLicenseKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLicenseKey(request *GetLicenseKeyRequest) (_result *GetLicenseKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLicenseKeyResponse{}
	_body, _err := client.GetLicenseKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMetricsOfAppWithOptions(request *GetMetricsOfAppRequest, runtime *util.RuntimeOptions) (_result *GetMetricsOfAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMetricsOfApp"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMetricsOfAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMetricsOfApp(request *GetMetricsOfAppRequest) (_result *GetMetricsOfAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMetricsOfAppResponse{}
	_body, _err := client.GetMetricsOfAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMetricsOfResourceWithOptions(request *GetMetricsOfResourceRequest, runtime *util.RuntimeOptions) (_result *GetMetricsOfResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMetricsOfResource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMetricsOfResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMetricsOfResource(request *GetMetricsOfResourceRequest) (_result *GetMetricsOfResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMetricsOfResourceResponse{}
	_body, _err := client.GetMetricsOfResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSentinelAppSumMetricWithOptions(request *GetSentinelAppSumMetricRequest, runtime *util.RuntimeOptions) (_result *GetSentinelAppSumMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSentinelAppSumMetric"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSentinelAppSumMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSentinelAppSumMetric(request *GetSentinelAppSumMetricRequest) (_result *GetSentinelAppSumMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSentinelAppSumMetricResponse{}
	_body, _err := client.GetSentinelAppSumMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserApplicationGroupsWithOptions(request *GetUserApplicationGroupsRequest, runtime *util.RuntimeOptions) (_result *GetUserApplicationGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserApplicationGroups"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserApplicationGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserApplicationGroups(request *GetUserApplicationGroupsRequest) (_result *GetUserApplicationGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserApplicationGroupsResponse{}
	_body, _err := client.GetUserApplicationGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserApplicationsWithOptions(request *GetUserApplicationsRequest, runtime *util.RuntimeOptions) (_result *GetUserApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserApplications"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserApplications(request *GetUserApplicationsRequest) (_result *GetUserApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserApplicationsResponse{}
	_body, _err := client.GetUserApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWorkspaceWithOptions(request *GetUserWorkspaceRequest, runtime *util.RuntimeOptions) (_result *GetUserWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserWorkspace"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserWorkspace(request *GetUserWorkspaceRequest) (_result *GetUserWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserWorkspaceResponse{}
	_body, _err := client.GetUserWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListActiveAppsWithOptions(request *ListActiveAppsRequest, runtime *util.RuntimeOptions) (_result *ListActiveAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListActiveApps"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListActiveAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListActiveApps(request *ListActiveAppsRequest) (_result *ListActiveAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListActiveAppsResponse{}
	_body, _err := client.ListActiveAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDegradeRulesOfAppWithOptions(request *ListDegradeRulesOfAppRequest, runtime *util.RuntimeOptions) (_result *ListDegradeRulesOfAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDegradeRulesOfApp"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDegradeRulesOfAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDegradeRulesOfApp(request *ListDegradeRulesOfAppRequest) (_result *ListDegradeRulesOfAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDegradeRulesOfAppResponse{}
	_body, _err := client.ListDegradeRulesOfAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDegradeRulesOfResourceWithOptions(request *ListDegradeRulesOfResourceRequest, runtime *util.RuntimeOptions) (_result *ListDegradeRulesOfResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDegradeRulesOfResource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDegradeRulesOfResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDegradeRulesOfResource(request *ListDegradeRulesOfResourceRequest) (_result *ListDegradeRulesOfResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDegradeRulesOfResourceResponse{}
	_body, _err := client.ListDegradeRulesOfResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExperimentMetasWithOptions(request *ListExperimentMetasRequest, runtime *util.RuntimeOptions) (_result *ListExperimentMetasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExperimentMetas"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExperimentMetasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExperimentMetas(request *ListExperimentMetasRequest) (_result *ListExperimentMetasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExperimentMetasResponse{}
	_body, _err := client.ListExperimentMetasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowRulesOfAppWithOptions(request *ListFlowRulesOfAppRequest, runtime *util.RuntimeOptions) (_result *ListFlowRulesOfAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowRulesOfApp"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowRulesOfAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowRulesOfApp(request *ListFlowRulesOfAppRequest) (_result *ListFlowRulesOfAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowRulesOfAppResponse{}
	_body, _err := client.ListFlowRulesOfAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowRulesOfResourceWithOptions(request *ListFlowRulesOfResourceRequest, runtime *util.RuntimeOptions) (_result *ListFlowRulesOfResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowRulesOfResource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowRulesOfResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowRulesOfResource(request *ListFlowRulesOfResourceRequest) (_result *ListFlowRulesOfResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowRulesOfResourceResponse{}
	_body, _err := client.ListFlowRulesOfResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotParamRulesOfAppWithOptions(request *ListHotParamRulesOfAppRequest, runtime *util.RuntimeOptions) (_result *ListHotParamRulesOfAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotParamRulesOfApp"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotParamRulesOfAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotParamRulesOfApp(request *ListHotParamRulesOfAppRequest) (_result *ListHotParamRulesOfAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotParamRulesOfAppResponse{}
	_body, _err := client.ListHotParamRulesOfAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotParamRulesOfResourceWithOptions(request *ListHotParamRulesOfResourceRequest, runtime *util.RuntimeOptions) (_result *ListHotParamRulesOfResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotParamRulesOfResource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotParamRulesOfResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotParamRulesOfResource(request *ListHotParamRulesOfResourceRequest) (_result *ListHotParamRulesOfResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotParamRulesOfResourceResponse{}
	_body, _err := client.ListHotParamRulesOfResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIsolationRulesOfAppWithOptions(request *ListIsolationRulesOfAppRequest, runtime *util.RuntimeOptions) (_result *ListIsolationRulesOfAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIsolationRulesOfApp"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIsolationRulesOfAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIsolationRulesOfApp(request *ListIsolationRulesOfAppRequest) (_result *ListIsolationRulesOfAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIsolationRulesOfAppResponse{}
	_body, _err := client.ListIsolationRulesOfAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIsolationRulesOfResourceWithOptions(request *ListIsolationRulesOfResourceRequest, runtime *util.RuntimeOptions) (_result *ListIsolationRulesOfResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIsolationRulesOfResource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIsolationRulesOfResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIsolationRulesOfResource(request *ListIsolationRulesOfResourceRequest) (_result *ListIsolationRulesOfResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIsolationRulesOfResourceResponse{}
	_body, _err := client.ListIsolationRulesOfResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSystemRulesWithOptions(request *ListSystemRulesRequest, runtime *util.RuntimeOptions) (_result *ListSystemRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSystemRules"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSystemRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSystemRules(request *ListSystemRulesRequest) (_result *ListSystemRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSystemRulesResponse{}
	_body, _err := client.ListSystemRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDegradeRuleWithOptions(request *ModifyDegradeRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyDegradeRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.HalfOpenBaseAmountPerStep)) {
		query["HalfOpenBaseAmountPerStep"] = request.HalfOpenBaseAmountPerStep
	}

	if !tea.BoolValue(util.IsUnset(request.HalfOpenRecoveryStepNum)) {
		query["HalfOpenRecoveryStepNum"] = request.HalfOpenRecoveryStepNum
	}

	if !tea.BoolValue(util.IsUnset(request.MinRequestAmount)) {
		query["MinRequestAmount"] = request.MinRequestAmount
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveryTimeoutMs)) {
		query["RecoveryTimeoutMs"] = request.RecoveryTimeoutMs
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SlowRtMs)) {
		query["SlowRtMs"] = request.SlowRtMs
	}

	if !tea.BoolValue(util.IsUnset(request.StatDurationMs)) {
		query["StatDurationMs"] = request.StatDurationMs
	}

	if !tea.BoolValue(util.IsUnset(request.Strategy)) {
		query["Strategy"] = request.Strategy
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDegradeRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDegradeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDegradeRule(request *ModifyDegradeRuleRequest) (_result *ModifyDegradeRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDegradeRuleResponse{}
	_body, _err := client.ModifyDegradeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowRuleWithOptions(request *ModifyFlowRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ControlBehavior)) {
		query["ControlBehavior"] = request.ControlBehavior
	}

	if !tea.BoolValue(util.IsUnset(request.LimitOrigin)) {
		query["LimitOrigin"] = request.LimitOrigin
	}

	if !tea.BoolValue(util.IsUnset(request.MaxQueueingTimeMs)) {
		query["MaxQueueingTimeMs"] = request.MaxQueueingTimeMs
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.RefResource)) {
		query["RefResource"] = request.RefResource
	}

	if !tea.BoolValue(util.IsUnset(request.RelationStrategy)) {
		query["RelationStrategy"] = request.RelationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	if !tea.BoolValue(util.IsUnset(request.WarmUpPeriodSec)) {
		query["WarmUpPeriodSec"] = request.WarmUpPeriodSec
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFlowRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowRule(request *ModifyFlowRuleRequest) (_result *ModifyFlowRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowRuleResponse{}
	_body, _err := client.ModifyFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHotParamRuleWithOptions(request *ModifyHotParamRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyHotParamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.BurstCount)) {
		query["BurstCount"] = request.BurstCount
	}

	if !tea.BoolValue(util.IsUnset(request.ControlBehavior)) {
		query["ControlBehavior"] = request.ControlBehavior
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.MaxQueueingTimeMs)) {
		query["MaxQueueingTimeMs"] = request.MaxQueueingTimeMs
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.ParamIdx)) {
		query["ParamIdx"] = request.ParamIdx
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.StatDurationSec)) {
		query["StatDurationSec"] = request.StatDurationSec
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyHotParamRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyHotParamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHotParamRule(request *ModifyHotParamRuleRequest) (_result *ModifyHotParamRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHotParamRuleResponse{}
	_body, _err := client.ModifyHotParamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyIsolationRuleWithOptions(request *ModifyIsolationRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyIsolationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LimitOrigin)) {
		query["LimitOrigin"] = request.LimitOrigin
	}

	if !tea.BoolValue(util.IsUnset(request.RefResource)) {
		query["RefResource"] = request.RefResource
	}

	if !tea.BoolValue(util.IsUnset(request.RelationStrategy)) {
		query["RelationStrategy"] = request.RelationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIsolationRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyIsolationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyIsolationRule(request *ModifyIsolationRuleRequest) (_result *ModifyIsolationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyIsolationRuleResponse{}
	_body, _err := client.ModifyIsolationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySystemRuleWithOptions(request *ModifySystemRuleRequest, runtime *util.RuntimeOptions) (_result *ModifySystemRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySystemRule"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySystemRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySystemRule(request *ModifySystemRuleRequest) (_result *ModifySystemRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySystemRuleResponse{}
	_body, _err := client.ModifySystemRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenAhasServiceWithOptions(request *OpenAhasServiceRequest, runtime *util.RuntimeOptions) (_result *OpenAhasServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenAhasService"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenAhasServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenAhasService(request *OpenAhasServiceRequest) (_result *OpenAhasServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenAhasServiceResponse{}
	_body, _err := client.OpenAhasServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageableQueryExperimentTaskByClusterIdWithOptions(request *PageableQueryExperimentTaskByClusterIdRequest, runtime *util.RuntimeOptions) (_result *PageableQueryExperimentTaskByClusterIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeInvalidHost)) {
		query["IncludeInvalidHost"] = request.IncludeInvalidHost
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PageableQueryExperimentTaskByClusterId"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageableQueryExperimentTaskByClusterIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageableQueryExperimentTaskByClusterId(request *PageableQueryExperimentTaskByClusterIdRequest) (_result *PageableQueryExperimentTaskByClusterIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageableQueryExperimentTaskByClusterIdResponse{}
	_body, _err := client.PageableQueryExperimentTaskByClusterIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageableQueryExperimentTaskByExperimentIdWithOptions(request *PageableQueryExperimentTaskByExperimentIdRequest, runtime *util.RuntimeOptions) (_result *PageableQueryExperimentTaskByExperimentIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PageableQueryExperimentTaskByExperimentId"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageableQueryExperimentTaskByExperimentIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageableQueryExperimentTaskByExperimentId(request *PageableQueryExperimentTaskByExperimentIdRequest) (_result *PageableQueryExperimentTaskByExperimentIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageableQueryExperimentTaskByExperimentIdResponse{}
	_body, _err := client.PageableQueryExperimentTaskByExperimentIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageableQueryUserExperimentWithOptions(request *PageableQueryUserExperimentRequest, runtime *util.RuntimeOptions) (_result *PageableQueryUserExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Results)) {
		query["Results"] = request.Results
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PageableQueryUserExperiment"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageableQueryUserExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageableQueryUserExperiment(request *PageableQueryUserExperimentRequest) (_result *PageableQueryUserExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageableQueryUserExperimentResponse{}
	_body, _err := client.PageableQueryUserExperimentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushExperimentTaskWithOptions(request *PushExperimentTaskRequest, runtime *util.RuntimeOptions) (_result *PushExperimentTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentTaskId)) {
		query["ExperimentTaskId"] = request.ExperimentTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushExperimentTask"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushExperimentTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushExperimentTask(request *PushExperimentTaskRequest) (_result *PushExperimentTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushExperimentTaskResponse{}
	_body, _err := client.PushExperimentTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateExperimentWithOptions(request *UpdateExperimentRequest, runtime *util.RuntimeOptions) (_result *UpdateExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Definition)) {
		query["Definition"] = request.Definition
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExperiment"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateExperiment(request *UpdateExperimentRequest) (_result *UpdateExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateExperimentResponse{}
	_body, _err := client.UpdateExperimentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateExperimentBasicInfoWithOptions(request *UpdateExperimentBasicInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateExperimentBasicInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AhasRegionId)) {
		query["AhasRegionId"] = request.AhasRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExperimentBasicInfo"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExperimentBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateExperimentBasicInfo(request *UpdateExperimentBasicInfoRequest) (_result *UpdateExperimentBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateExperimentBasicInfoResponse{}
	_body, _err := client.UpdateExperimentBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
