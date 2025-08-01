// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateFlowRuleRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *UpdateFlowRuleRequest
	GetAppId() *string
	SetAppName(v string) *UpdateFlowRuleRequest
	GetAppName() *string
	SetControlBehavior(v int32) *UpdateFlowRuleRequest
	GetControlBehavior() *int32
	SetEnable(v bool) *UpdateFlowRuleRequest
	GetEnable() *bool
	SetLimitApp(v string) *UpdateFlowRuleRequest
	GetLimitApp() *string
	SetMaxQueueingTimeMs(v int32) *UpdateFlowRuleRequest
	GetMaxQueueingTimeMs() *int32
	SetNamespace(v string) *UpdateFlowRuleRequest
	GetNamespace() *string
	SetRuleId(v int64) *UpdateFlowRuleRequest
	GetRuleId() *int64
	SetThreshold(v int32) *UpdateFlowRuleRequest
	GetThreshold() *int32
}

type UpdateFlowRuleRequest struct {
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
	// The application ID.
	//
	// example:
	//
	// hkhon1po62@c3df23522******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The throttling effect.
	//
	// Valid values:
	//
	// 	- 0
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     fast failure
	//
	//     <!-- -->
	//
	// 	- 2
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     in queue
	//
	//     <!-- -->
	//
	// example:
	//
	// 0
	ControlBehavior *int32 `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	// Specifies whether to enable the rule.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Enable   *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitApp *string `json:"LimitApp,omitempty" xml:"LimitApp,omitempty"`
	// The timeout period. Unit: milliseconds. This parameter is required when the value of ControlBehavior is set to 2.
	//
	// example:
	//
	// 500
	MaxQueueingTimeMs *int32 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	// The namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The throttling threshold.
	//
	// example:
	//
	// 30
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateFlowRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateFlowRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateFlowRuleRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateFlowRuleRequest) GetControlBehavior() *int32 {
	return s.ControlBehavior
}

func (s *UpdateFlowRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateFlowRuleRequest) GetLimitApp() *string {
	return s.LimitApp
}

func (s *UpdateFlowRuleRequest) GetMaxQueueingTimeMs() *int32 {
	return s.MaxQueueingTimeMs
}

func (s *UpdateFlowRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateFlowRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *UpdateFlowRuleRequest) GetThreshold() *int32 {
	return s.Threshold
}

func (s *UpdateFlowRuleRequest) SetAcceptLanguage(v string) *UpdateFlowRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateFlowRuleRequest) SetAppId(v string) *UpdateFlowRuleRequest {
	s.AppId = &v
	return s
}

func (s *UpdateFlowRuleRequest) SetAppName(v string) *UpdateFlowRuleRequest {
	s.AppName = &v
	return s
}

func (s *UpdateFlowRuleRequest) SetControlBehavior(v int32) *UpdateFlowRuleRequest {
	s.ControlBehavior = &v
	return s
}

func (s *UpdateFlowRuleRequest) SetEnable(v bool) *UpdateFlowRuleRequest {
	s.Enable = &v
	return s
}

func (s *UpdateFlowRuleRequest) SetLimitApp(v string) *UpdateFlowRuleRequest {
	s.LimitApp = &v
	return s
}

func (s *UpdateFlowRuleRequest) SetMaxQueueingTimeMs(v int32) *UpdateFlowRuleRequest {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *UpdateFlowRuleRequest) SetNamespace(v string) *UpdateFlowRuleRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateFlowRuleRequest) SetRuleId(v int64) *UpdateFlowRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateFlowRuleRequest) SetThreshold(v int32) *UpdateFlowRuleRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateFlowRuleRequest) Validate() error {
	return dara.Validate(s)
}
