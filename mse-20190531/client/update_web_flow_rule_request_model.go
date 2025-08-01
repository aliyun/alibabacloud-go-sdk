// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebFlowRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateWebFlowRuleRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *UpdateWebFlowRuleRequest
	GetAppId() *string
	SetAppName(v string) *UpdateWebFlowRuleRequest
	GetAppName() *string
	SetBurst(v int32) *UpdateWebFlowRuleRequest
	GetBurst() *int32
	SetControlBehavior(v int32) *UpdateWebFlowRuleRequest
	GetControlBehavior() *int32
	SetEnable(v bool) *UpdateWebFlowRuleRequest
	GetEnable() *bool
	SetMaxQueueingTimeMs(v int32) *UpdateWebFlowRuleRequest
	GetMaxQueueingTimeMs() *int32
	SetMetricType(v int32) *UpdateWebFlowRuleRequest
	GetMetricType() *int32
	SetNamespace(v string) *UpdateWebFlowRuleRequest
	GetNamespace() *string
	SetParamItem(v string) *UpdateWebFlowRuleRequest
	GetParamItem() *string
	SetRegionId(v string) *UpdateWebFlowRuleRequest
	GetRegionId() *string
	SetResourceMode(v int32) *UpdateWebFlowRuleRequest
	GetResourceMode() *int32
	SetRuleId(v int64) *UpdateWebFlowRuleRequest
	GetRuleId() *int64
	SetStatIntervalMs(v int32) *UpdateWebFlowRuleRequest
	GetStatIntervalMs() *int32
	SetThreshold(v float32) *UpdateWebFlowRuleRequest
	GetThreshold() *float32
}

type UpdateWebFlowRuleRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// hkhon1****@c3df23522******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 0
	Burst *int32 `json:"Burst,omitempty" xml:"Burst,omitempty"`
	// example:
	//
	// 0
	ControlBehavior *int32 `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 200
	MaxQueueingTimeMs *int32 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	// example:
	//
	// 1
	MetricType *int32 `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// {"fieldName":"key","matchStrategy":2,"parseStrategy":2,"pattern":"value"}
	ParamItem *string `json:"ParamItem,omitempty" xml:"ParamItem,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 0
	ResourceMode *int32 `json:"ResourceMode,omitempty" xml:"ResourceMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 1000
	StatIntervalMs *int32 `json:"StatIntervalMs,omitempty" xml:"StatIntervalMs,omitempty"`
	// example:
	//
	// 20
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateWebFlowRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebFlowRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateWebFlowRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateWebFlowRuleRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateWebFlowRuleRequest) GetBurst() *int32 {
	return s.Burst
}

func (s *UpdateWebFlowRuleRequest) GetControlBehavior() *int32 {
	return s.ControlBehavior
}

func (s *UpdateWebFlowRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateWebFlowRuleRequest) GetMaxQueueingTimeMs() *int32 {
	return s.MaxQueueingTimeMs
}

func (s *UpdateWebFlowRuleRequest) GetMetricType() *int32 {
	return s.MetricType
}

func (s *UpdateWebFlowRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateWebFlowRuleRequest) GetParamItem() *string {
	return s.ParamItem
}

func (s *UpdateWebFlowRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWebFlowRuleRequest) GetResourceMode() *int32 {
	return s.ResourceMode
}

func (s *UpdateWebFlowRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *UpdateWebFlowRuleRequest) GetStatIntervalMs() *int32 {
	return s.StatIntervalMs
}

func (s *UpdateWebFlowRuleRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateWebFlowRuleRequest) SetAcceptLanguage(v string) *UpdateWebFlowRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetAppId(v string) *UpdateWebFlowRuleRequest {
	s.AppId = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetAppName(v string) *UpdateWebFlowRuleRequest {
	s.AppName = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetBurst(v int32) *UpdateWebFlowRuleRequest {
	s.Burst = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetControlBehavior(v int32) *UpdateWebFlowRuleRequest {
	s.ControlBehavior = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetEnable(v bool) *UpdateWebFlowRuleRequest {
	s.Enable = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetMaxQueueingTimeMs(v int32) *UpdateWebFlowRuleRequest {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetMetricType(v int32) *UpdateWebFlowRuleRequest {
	s.MetricType = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetNamespace(v string) *UpdateWebFlowRuleRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetParamItem(v string) *UpdateWebFlowRuleRequest {
	s.ParamItem = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetRegionId(v string) *UpdateWebFlowRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetResourceMode(v int32) *UpdateWebFlowRuleRequest {
	s.ResourceMode = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetRuleId(v int64) *UpdateWebFlowRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetStatIntervalMs(v int32) *UpdateWebFlowRuleRequest {
	s.StatIntervalMs = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) SetThreshold(v float32) *UpdateWebFlowRuleRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateWebFlowRuleRequest) Validate() error {
	return dara.Validate(s)
}
