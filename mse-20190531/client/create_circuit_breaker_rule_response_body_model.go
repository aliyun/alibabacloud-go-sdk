// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCircuitBreakerRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCircuitBreakerRuleResponseBody
	GetCode() *string
	SetData(v *CreateCircuitBreakerRuleResponseBodyData) *CreateCircuitBreakerRuleResponseBody
	GetData() *CreateCircuitBreakerRuleResponseBodyData
	SetMessage(v string) *CreateCircuitBreakerRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCircuitBreakerRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCircuitBreakerRuleResponseBody
	GetSuccess() *bool
}

type CreateCircuitBreakerRuleResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the rule.
	Data *CreateCircuitBreakerRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE5C32A1-BC0E-4B79-817C-103E4EDF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCircuitBreakerRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCircuitBreakerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCircuitBreakerRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCircuitBreakerRuleResponseBody) GetData() *CreateCircuitBreakerRuleResponseBodyData {
	return s.Data
}

func (s *CreateCircuitBreakerRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCircuitBreakerRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCircuitBreakerRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCircuitBreakerRuleResponseBody) SetCode(v string) *CreateCircuitBreakerRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBody) SetData(v *CreateCircuitBreakerRuleResponseBodyData) *CreateCircuitBreakerRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBody) SetMessage(v string) *CreateCircuitBreakerRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBody) SetRequestId(v string) *CreateCircuitBreakerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBody) SetSuccess(v bool) *CreateCircuitBreakerRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateCircuitBreakerRuleResponseBodyData struct {
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@c3df23522bXXXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Indicates whether the rule is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The minimum number of requests that can be passed in each step after circuit breaking recovers.
	//
	// example:
	//
	// 5
	HalfOpenBaseAmountPerStep *int32 `json:"HalfOpenBaseAmountPerStep,omitempty" xml:"HalfOpenBaseAmountPerStep,omitempty"`
	// The number of circuit breaking recovery steps.
	//
	// example:
	//
	// 1
	HalfOpenRecoveryStepNum *int32 `json:"HalfOpenRecoveryStepNum,omitempty" xml:"HalfOpenRecoveryStepNum,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum response time (RT). Unit: milliseconds. If the RT of a request is greater than the value of this parameter, a slow call is counted. If you set Strategy to 0, you must specify this parameter.
	//
	// example:
	//
	// 200
	MaxAllowedRtMs *int32 `json:"MaxAllowedRtMs,omitempty" xml:"MaxAllowedRtMs,omitempty"`
	// The minimum number of requests to trigger circuit breaking. If the number of requests in the current time window is less than the value of this parameter, circuit breaking is not triggered even if the circuit breaking rule is met.
	//
	// example:
	//
	// 10
	MinRequestAmount *int32 `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	// The microservice namespace to which the application belongs.
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
	// The name of the interface to which the rule is applicable. The interface name must be the same as the name on the interface details page in the console.
	//
	// example:
	//
	// /a
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The period in which circuit breaking is implemented. Unit: milliseconds. If circuit breaking is implemented on the requests for the route, the calls to all the requests for the route fail in the configured circuit breaking period.
	//
	// example:
	//
	// 10000
	RetryTimeoutMs *int32 `json:"RetryTimeoutMs,omitempty" xml:"RetryTimeoutMs,omitempty"`
	// The length of the time window. Unit: milliseconds. The valid range is from 1 second to 120 minutes.
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
	//     Slow call ratio
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
	//     Abnormal proportion
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

func (s CreateCircuitBreakerRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCircuitBreakerRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetHalfOpenBaseAmountPerStep() *int32 {
	return s.HalfOpenBaseAmountPerStep
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetHalfOpenRecoveryStepNum() *int32 {
	return s.HalfOpenRecoveryStepNum
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetMaxAllowedRtMs() *int32 {
	return s.MaxAllowedRtMs
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetMinRequestAmount() *int32 {
	return s.MinRequestAmount
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetResource() *string {
	return s.Resource
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetRetryTimeoutMs() *int32 {
	return s.RetryTimeoutMs
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetStatIntervalMs() *int32 {
	return s.StatIntervalMs
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetStrategy() *int32 {
	return s.Strategy
}

func (s *CreateCircuitBreakerRuleResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetAppId(v string) *CreateCircuitBreakerRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetAppName(v string) *CreateCircuitBreakerRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetEnable(v bool) *CreateCircuitBreakerRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetHalfOpenBaseAmountPerStep(v int32) *CreateCircuitBreakerRuleResponseBodyData {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetHalfOpenRecoveryStepNum(v int32) *CreateCircuitBreakerRuleResponseBodyData {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetId(v int64) *CreateCircuitBreakerRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetMaxAllowedRtMs(v int32) *CreateCircuitBreakerRuleResponseBodyData {
	s.MaxAllowedRtMs = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetMinRequestAmount(v int32) *CreateCircuitBreakerRuleResponseBodyData {
	s.MinRequestAmount = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetNamespace(v string) *CreateCircuitBreakerRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetRegionId(v string) *CreateCircuitBreakerRuleResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetResource(v string) *CreateCircuitBreakerRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetRetryTimeoutMs(v int32) *CreateCircuitBreakerRuleResponseBodyData {
	s.RetryTimeoutMs = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetStatIntervalMs(v int32) *CreateCircuitBreakerRuleResponseBodyData {
	s.StatIntervalMs = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetStrategy(v int32) *CreateCircuitBreakerRuleResponseBodyData {
	s.Strategy = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) SetThreshold(v float32) *CreateCircuitBreakerRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
