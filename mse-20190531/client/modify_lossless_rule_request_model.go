// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLosslessRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ModifyLosslessRuleRequest
	GetAcceptLanguage() *string
	SetAligned(v bool) *ModifyLosslessRuleRequest
	GetAligned() *bool
	SetAppId(v string) *ModifyLosslessRuleRequest
	GetAppId() *string
	SetAppName(v string) *ModifyLosslessRuleRequest
	GetAppName() *string
	SetDelayTime(v int32) *ModifyLosslessRuleRequest
	GetDelayTime() *int32
	SetEnable(v bool) *ModifyLosslessRuleRequest
	GetEnable() *bool
	SetFuncType(v int32) *ModifyLosslessRuleRequest
	GetFuncType() *int32
	SetLossLessDetail(v bool) *ModifyLosslessRuleRequest
	GetLossLessDetail() *bool
	SetNamespace(v string) *ModifyLosslessRuleRequest
	GetNamespace() *string
	SetNotice(v bool) *ModifyLosslessRuleRequest
	GetNotice() *bool
	SetRegionId(v string) *ModifyLosslessRuleRequest
	GetRegionId() *string
	SetRelated(v bool) *ModifyLosslessRuleRequest
	GetRelated() *bool
	SetWarmupTime(v int32) *ModifyLosslessRuleRequest
	GetWarmupTime() *int32
}

type ModifyLosslessRuleRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Specifies whether to align the lifecycle of the application in the Kubernetes cluster with that of the microservice.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Aligned *bool `json:"Aligned,omitempty" xml:"Aligned,omitempty"`
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// c644n5frmc@3e75f25fd4*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// wx-work-api
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The registration latency.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// Specifies whether to enable the alert rule. Valid values:
	//
	// 	- `true`: enables the rule.
	//
	// 	- `false`: disables the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The slope of the prefetching curve.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	FuncType *int32 `json:"FuncType,omitempty" xml:"FuncType,omitempty"`
	// Specifies whether to display online and offline processing details.
	//
	// example:
	//
	// false
	LossLessDetail *bool `json:"LossLessDetail,omitempty" xml:"LossLessDetail,omitempty"`
	// The microservice namespace to which the rule applies.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Specifies whether to enable notification.
	//
	// example:
	//
	// false
	Notice *bool `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to associate with service prefetching.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Related *bool `json:"Related,omitempty" xml:"Related,omitempty"`
	// The prefetching duration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	WarmupTime *int32 `json:"WarmupTime,omitempty" xml:"WarmupTime,omitempty"`
}

func (s ModifyLosslessRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLosslessRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyLosslessRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ModifyLosslessRuleRequest) GetAligned() *bool {
	return s.Aligned
}

func (s *ModifyLosslessRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLosslessRuleRequest) GetAppName() *string {
	return s.AppName
}

func (s *ModifyLosslessRuleRequest) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *ModifyLosslessRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyLosslessRuleRequest) GetFuncType() *int32 {
	return s.FuncType
}

func (s *ModifyLosslessRuleRequest) GetLossLessDetail() *bool {
	return s.LossLessDetail
}

func (s *ModifyLosslessRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyLosslessRuleRequest) GetNotice() *bool {
	return s.Notice
}

func (s *ModifyLosslessRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLosslessRuleRequest) GetRelated() *bool {
	return s.Related
}

func (s *ModifyLosslessRuleRequest) GetWarmupTime() *int32 {
	return s.WarmupTime
}

func (s *ModifyLosslessRuleRequest) SetAcceptLanguage(v string) *ModifyLosslessRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetAligned(v bool) *ModifyLosslessRuleRequest {
	s.Aligned = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetAppId(v string) *ModifyLosslessRuleRequest {
	s.AppId = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetAppName(v string) *ModifyLosslessRuleRequest {
	s.AppName = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetDelayTime(v int32) *ModifyLosslessRuleRequest {
	s.DelayTime = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetEnable(v bool) *ModifyLosslessRuleRequest {
	s.Enable = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetFuncType(v int32) *ModifyLosslessRuleRequest {
	s.FuncType = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetLossLessDetail(v bool) *ModifyLosslessRuleRequest {
	s.LossLessDetail = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetNamespace(v string) *ModifyLosslessRuleRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetNotice(v bool) *ModifyLosslessRuleRequest {
	s.Notice = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetRegionId(v string) *ModifyLosslessRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetRelated(v bool) *ModifyLosslessRuleRequest {
	s.Related = &v
	return s
}

func (s *ModifyLosslessRuleRequest) SetWarmupTime(v int32) *ModifyLosslessRuleRequest {
	s.WarmupTime = &v
	return s
}

func (s *ModifyLosslessRuleRequest) Validate() error {
	return dara.Validate(s)
}
