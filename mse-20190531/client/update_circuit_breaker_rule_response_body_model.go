// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCircuitBreakerRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCircuitBreakerRuleResponseBody
	GetCode() *string
	SetData(v *UpdateCircuitBreakerRuleResponseBodyData) *UpdateCircuitBreakerRuleResponseBody
	GetData() *UpdateCircuitBreakerRuleResponseBodyData
	SetMessage(v string) *UpdateCircuitBreakerRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCircuitBreakerRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCircuitBreakerRuleResponseBody
	GetSuccess() *bool
}

type UpdateCircuitBreakerRuleResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the rule.
	Data *UpdateCircuitBreakerRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
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
	//     :
	//
	//     <!-- -->
	//
	//     The request was successful.
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The request failed.
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCircuitBreakerRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCircuitBreakerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCircuitBreakerRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCircuitBreakerRuleResponseBody) GetData() *UpdateCircuitBreakerRuleResponseBodyData {
	return s.Data
}

func (s *UpdateCircuitBreakerRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCircuitBreakerRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCircuitBreakerRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCircuitBreakerRuleResponseBody) SetCode(v string) *UpdateCircuitBreakerRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBody) SetData(v *UpdateCircuitBreakerRuleResponseBodyData) *UpdateCircuitBreakerRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBody) SetMessage(v string) *UpdateCircuitBreakerRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBody) SetRequestId(v string) *UpdateCircuitBreakerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBody) SetSuccess(v bool) *UpdateCircuitBreakerRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCircuitBreakerRuleResponseBodyData struct {
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
	// The maximum RT. Unit: milliseconds. If the RT of a request is greater than the value of this parameter, a slow call is counted. If you set Strategy to 0, you must specify this parameter.
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

func (s UpdateCircuitBreakerRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateCircuitBreakerRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetHalfOpenBaseAmountPerStep() *int32 {
	return s.HalfOpenBaseAmountPerStep
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetHalfOpenRecoveryStepNum() *int32 {
	return s.HalfOpenRecoveryStepNum
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetMaxAllowedRtMs() *int32 {
	return s.MaxAllowedRtMs
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetMinRequestAmount() *int32 {
	return s.MinRequestAmount
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetResource() *string {
	return s.Resource
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetRetryTimeoutMs() *int32 {
	return s.RetryTimeoutMs
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetStatIntervalMs() *int32 {
	return s.StatIntervalMs
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetStrategy() *int32 {
	return s.Strategy
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetAppId(v string) *UpdateCircuitBreakerRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetAppName(v string) *UpdateCircuitBreakerRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetEnable(v bool) *UpdateCircuitBreakerRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetHalfOpenBaseAmountPerStep(v int32) *UpdateCircuitBreakerRuleResponseBodyData {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetHalfOpenRecoveryStepNum(v int32) *UpdateCircuitBreakerRuleResponseBodyData {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetId(v int64) *UpdateCircuitBreakerRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetMaxAllowedRtMs(v int32) *UpdateCircuitBreakerRuleResponseBodyData {
	s.MaxAllowedRtMs = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetMinRequestAmount(v int32) *UpdateCircuitBreakerRuleResponseBodyData {
	s.MinRequestAmount = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetNamespace(v string) *UpdateCircuitBreakerRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetResource(v string) *UpdateCircuitBreakerRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetRetryTimeoutMs(v int32) *UpdateCircuitBreakerRuleResponseBodyData {
	s.RetryTimeoutMs = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetStatIntervalMs(v int32) *UpdateCircuitBreakerRuleResponseBodyData {
	s.StatIntervalMs = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetStrategy(v int32) *UpdateCircuitBreakerRuleResponseBodyData {
	s.Strategy = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) SetThreshold(v float32) *UpdateCircuitBreakerRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
