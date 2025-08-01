// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCircuitBreakerRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCircuitBreakerRuleRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *UpdateCircuitBreakerRuleRequest
	GetAppId() *string
	SetAppName(v string) *UpdateCircuitBreakerRuleRequest
	GetAppName() *string
	SetEnable(v bool) *UpdateCircuitBreakerRuleRequest
	GetEnable() *bool
	SetHalfOpenBaseAmountPerStep(v int32) *UpdateCircuitBreakerRuleRequest
	GetHalfOpenBaseAmountPerStep() *int32
	SetHalfOpenRecoveryStepNum(v int32) *UpdateCircuitBreakerRuleRequest
	GetHalfOpenRecoveryStepNum() *int32
	SetMaxAllowedRtMs(v int32) *UpdateCircuitBreakerRuleRequest
	GetMaxAllowedRtMs() *int32
	SetMinRequestAmount(v int32) *UpdateCircuitBreakerRuleRequest
	GetMinRequestAmount() *int32
	SetNamespace(v string) *UpdateCircuitBreakerRuleRequest
	GetNamespace() *string
	SetRetryTimeoutMs(v int32) *UpdateCircuitBreakerRuleRequest
	GetRetryTimeoutMs() *int32
	SetRuleId(v int64) *UpdateCircuitBreakerRuleRequest
	GetRuleId() *int64
	SetStatIntervalMs(v int32) *UpdateCircuitBreakerRuleRequest
	GetStatIntervalMs() *int32
	SetStrategy(v int32) *UpdateCircuitBreakerRuleRequest
	GetStrategy() *int32
	SetThreshold(v float32) *UpdateCircuitBreakerRuleRequest
	GetThreshold() *float32
}

type UpdateCircuitBreakerRuleRequest struct {
	// The language of the response. Valid values: zh-CN and en-US. Default value: zh-CN. The value zh-CN indicates Chinese, and the value en-US indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@c3df23522bXXXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
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
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The minimum number of requests that can be passed in each step after circuit breaking recovers. Default value: 1.
	//
	// example:
	//
	// 5
	HalfOpenBaseAmountPerStep *int32 `json:"HalfOpenBaseAmountPerStep,omitempty" xml:"HalfOpenBaseAmountPerStep,omitempty"`
	// The number of circuit breaking recovery steps. Default value: 1.
	//
	// example:
	//
	// 1
	HalfOpenRecoveryStepNum *int32 `json:"HalfOpenRecoveryStepNum,omitempty" xml:"HalfOpenRecoveryStepNum,omitempty"`
	// The maximum response time (RT). Unit: milliseconds. If the RT of a request is greater than the value of this parameter, a slow call is counted. If you set Strategy to 0, you must specify this parameter.
	//
	// example:
	//
	// 200
	MaxAllowedRtMs *int32 `json:"MaxAllowedRtMs,omitempty" xml:"MaxAllowedRtMs,omitempty"`
	// The minimum number of requests to trigger circuit breaking. If the number of requests in the current time window is less than the value of this parameter, circuit breaking is not triggered even if the circuit breaking rule is met. Default value: 10.
	//
	// example:
	//
	// 10
	MinRequestAmount *int32 `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	// The microservice namespace to which the application belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The period in which circuit breaking is implemented. Unit: milliseconds. If circuit breaking is implemented on the requests for the route, the calls to all the requests for the route fail in the configured circuit breaking period. The value must be an integral multiple of 1,000. Default value: 10000. This value indicates 10 seconds.
	//
	// example:
	//
	// 10000
	RetryTimeoutMs *int32 `json:"RetryTimeoutMs,omitempty" xml:"RetryTimeoutMs,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The length of the time window. Unit: milliseconds. The valid range is from 1 second to 120 minutes. The default value is 20000. This value indicates 20 seconds.
	//
	// example:
	//
	// 20000
	StatIntervalMs *int32 `json:"StatIntervalMs,omitempty" xml:"StatIntervalMs,omitempty"`
	// The threshold type.
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
	//     slow call proportion
	//
	//     <!-- -->
	//
	// 	- 1
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     abnormal proportion
	//
	//     <!-- -->
	//
	// example:
	//
	// 0
	Strategy *int32 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// A percentage threshold for triggering circuit breaking. Valid values: 0-1. These values represent 0% to 100%.
	//
	// example:
	//
	// 0.8
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateCircuitBreakerRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCircuitBreakerRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateCircuitBreakerRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCircuitBreakerRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateCircuitBreakerRuleRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateCircuitBreakerRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateCircuitBreakerRuleRequest) GetHalfOpenBaseAmountPerStep() *int32 {
	return s.HalfOpenBaseAmountPerStep
}

func (s *UpdateCircuitBreakerRuleRequest) GetHalfOpenRecoveryStepNum() *int32 {
	return s.HalfOpenRecoveryStepNum
}

func (s *UpdateCircuitBreakerRuleRequest) GetMaxAllowedRtMs() *int32 {
	return s.MaxAllowedRtMs
}

func (s *UpdateCircuitBreakerRuleRequest) GetMinRequestAmount() *int32 {
	return s.MinRequestAmount
}

func (s *UpdateCircuitBreakerRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateCircuitBreakerRuleRequest) GetRetryTimeoutMs() *int32 {
	return s.RetryTimeoutMs
}

func (s *UpdateCircuitBreakerRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *UpdateCircuitBreakerRuleRequest) GetStatIntervalMs() *int32 {
	return s.StatIntervalMs
}

func (s *UpdateCircuitBreakerRuleRequest) GetStrategy() *int32 {
	return s.Strategy
}

func (s *UpdateCircuitBreakerRuleRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateCircuitBreakerRuleRequest) SetAcceptLanguage(v string) *UpdateCircuitBreakerRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetAppId(v string) *UpdateCircuitBreakerRuleRequest {
	s.AppId = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetAppName(v string) *UpdateCircuitBreakerRuleRequest {
	s.AppName = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetEnable(v bool) *UpdateCircuitBreakerRuleRequest {
	s.Enable = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetHalfOpenBaseAmountPerStep(v int32) *UpdateCircuitBreakerRuleRequest {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetHalfOpenRecoveryStepNum(v int32) *UpdateCircuitBreakerRuleRequest {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetMaxAllowedRtMs(v int32) *UpdateCircuitBreakerRuleRequest {
	s.MaxAllowedRtMs = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetMinRequestAmount(v int32) *UpdateCircuitBreakerRuleRequest {
	s.MinRequestAmount = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetNamespace(v string) *UpdateCircuitBreakerRuleRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetRetryTimeoutMs(v int32) *UpdateCircuitBreakerRuleRequest {
	s.RetryTimeoutMs = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetRuleId(v int64) *UpdateCircuitBreakerRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetStatIntervalMs(v int32) *UpdateCircuitBreakerRuleRequest {
	s.StatIntervalMs = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetStrategy(v int32) *UpdateCircuitBreakerRuleRequest {
	s.Strategy = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) SetThreshold(v float32) *UpdateCircuitBreakerRuleRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateCircuitBreakerRuleRequest) Validate() error {
	return dara.Validate(s)
}
