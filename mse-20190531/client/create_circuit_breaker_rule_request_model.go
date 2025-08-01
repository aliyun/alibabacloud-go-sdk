// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCircuitBreakerRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateCircuitBreakerRuleRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *CreateCircuitBreakerRuleRequest
	GetAppId() *string
	SetAppName(v string) *CreateCircuitBreakerRuleRequest
	GetAppName() *string
	SetEnable(v bool) *CreateCircuitBreakerRuleRequest
	GetEnable() *bool
	SetHalfOpenBaseAmountPerStep(v int32) *CreateCircuitBreakerRuleRequest
	GetHalfOpenBaseAmountPerStep() *int32
	SetHalfOpenRecoveryStepNum(v int32) *CreateCircuitBreakerRuleRequest
	GetHalfOpenRecoveryStepNum() *int32
	SetMaxAllowedRtMs(v int32) *CreateCircuitBreakerRuleRequest
	GetMaxAllowedRtMs() *int32
	SetMinRequestAmount(v int32) *CreateCircuitBreakerRuleRequest
	GetMinRequestAmount() *int32
	SetNamespace(v string) *CreateCircuitBreakerRuleRequest
	GetNamespace() *string
	SetRegionId(v string) *CreateCircuitBreakerRuleRequest
	GetRegionId() *string
	SetResource(v string) *CreateCircuitBreakerRuleRequest
	GetResource() *string
	SetResourceType(v int32) *CreateCircuitBreakerRuleRequest
	GetResourceType() *int32
	SetRetryTimeoutMs(v int32) *CreateCircuitBreakerRuleRequest
	GetRetryTimeoutMs() *int32
	SetStatIntervalMs(v int32) *CreateCircuitBreakerRuleRequest
	GetStatIntervalMs() *int32
	SetStrategy(v int32) *CreateCircuitBreakerRuleRequest
	GetStrategy() *int32
	SetThreshold(v float32) *CreateCircuitBreakerRuleRequest
	GetThreshold() *float32
}

type CreateCircuitBreakerRuleRequest struct {
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
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@c3df23522bXXXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
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
	// This parameter is required.
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
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the interface to which the rule applies. The interface name must be the same as the name on the interface details page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// /a
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- 0: custom interface
	//
	// 	- 1: web
	//
	// 	- 2: RPC
	//
	// 	- 3: route
	//
	// 	- 4: SQL
	//
	// example:
	//
	// 0
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The period in which circuit breaking is implemented. Unit: milliseconds. If circuit breaking is implemented on the requests for the route, the calls to all the requests for the route fail in the configured circuit breaking period. The value must be an integral multiple of 1,000. Default value: 10000. This value indicates 10 seconds.
	//
	// example:
	//
	// 10000
	RetryTimeoutMs *int32 `json:"RetryTimeoutMs,omitempty" xml:"RetryTimeoutMs,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 0
	Strategy *int32 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// A percentage threshold for triggering circuit breaking. Valid values: 0-1. These values represent 0% to 100%.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.8
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateCircuitBreakerRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCircuitBreakerRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateCircuitBreakerRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateCircuitBreakerRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateCircuitBreakerRuleRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateCircuitBreakerRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateCircuitBreakerRuleRequest) GetHalfOpenBaseAmountPerStep() *int32 {
	return s.HalfOpenBaseAmountPerStep
}

func (s *CreateCircuitBreakerRuleRequest) GetHalfOpenRecoveryStepNum() *int32 {
	return s.HalfOpenRecoveryStepNum
}

func (s *CreateCircuitBreakerRuleRequest) GetMaxAllowedRtMs() *int32 {
	return s.MaxAllowedRtMs
}

func (s *CreateCircuitBreakerRuleRequest) GetMinRequestAmount() *int32 {
	return s.MinRequestAmount
}

func (s *CreateCircuitBreakerRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateCircuitBreakerRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCircuitBreakerRuleRequest) GetResource() *string {
	return s.Resource
}

func (s *CreateCircuitBreakerRuleRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *CreateCircuitBreakerRuleRequest) GetRetryTimeoutMs() *int32 {
	return s.RetryTimeoutMs
}

func (s *CreateCircuitBreakerRuleRequest) GetStatIntervalMs() *int32 {
	return s.StatIntervalMs
}

func (s *CreateCircuitBreakerRuleRequest) GetStrategy() *int32 {
	return s.Strategy
}

func (s *CreateCircuitBreakerRuleRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreateCircuitBreakerRuleRequest) SetAcceptLanguage(v string) *CreateCircuitBreakerRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetAppId(v string) *CreateCircuitBreakerRuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetAppName(v string) *CreateCircuitBreakerRuleRequest {
	s.AppName = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetEnable(v bool) *CreateCircuitBreakerRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetHalfOpenBaseAmountPerStep(v int32) *CreateCircuitBreakerRuleRequest {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetHalfOpenRecoveryStepNum(v int32) *CreateCircuitBreakerRuleRequest {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetMaxAllowedRtMs(v int32) *CreateCircuitBreakerRuleRequest {
	s.MaxAllowedRtMs = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetMinRequestAmount(v int32) *CreateCircuitBreakerRuleRequest {
	s.MinRequestAmount = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetNamespace(v string) *CreateCircuitBreakerRuleRequest {
	s.Namespace = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetRegionId(v string) *CreateCircuitBreakerRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetResource(v string) *CreateCircuitBreakerRuleRequest {
	s.Resource = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetResourceType(v int32) *CreateCircuitBreakerRuleRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetRetryTimeoutMs(v int32) *CreateCircuitBreakerRuleRequest {
	s.RetryTimeoutMs = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetStatIntervalMs(v int32) *CreateCircuitBreakerRuleRequest {
	s.StatIntervalMs = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetStrategy(v int32) *CreateCircuitBreakerRuleRequest {
	s.Strategy = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) SetThreshold(v float32) *CreateCircuitBreakerRuleRequest {
	s.Threshold = &v
	return s
}

func (s *CreateCircuitBreakerRuleRequest) Validate() error {
	return dara.Validate(s)
}
