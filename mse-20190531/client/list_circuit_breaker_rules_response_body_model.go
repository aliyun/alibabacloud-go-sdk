// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCircuitBreakerRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCircuitBreakerRulesResponseBody
	GetCode() *int32
	SetData(v *ListCircuitBreakerRulesResponseBodyData) *ListCircuitBreakerRulesResponseBody
	GetData() *ListCircuitBreakerRulesResponseBodyData
	SetHttpStatusCode(v int32) *ListCircuitBreakerRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCircuitBreakerRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCircuitBreakerRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCircuitBreakerRulesResponseBody
	GetSuccess() *bool
}

type ListCircuitBreakerRulesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the rule.
	Data *ListCircuitBreakerRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// 4E9FDCFE-0738-493B-B801-82BDFBCB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCircuitBreakerRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCircuitBreakerRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCircuitBreakerRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCircuitBreakerRulesResponseBody) GetData() *ListCircuitBreakerRulesResponseBodyData {
	return s.Data
}

func (s *ListCircuitBreakerRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCircuitBreakerRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCircuitBreakerRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCircuitBreakerRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCircuitBreakerRulesResponseBody) SetCode(v int32) *ListCircuitBreakerRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBody) SetData(v *ListCircuitBreakerRulesResponseBodyData) *ListCircuitBreakerRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListCircuitBreakerRulesResponseBody) SetHttpStatusCode(v int32) *ListCircuitBreakerRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBody) SetMessage(v string) *ListCircuitBreakerRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBody) SetRequestId(v string) *ListCircuitBreakerRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBody) SetSuccess(v bool) *ListCircuitBreakerRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCircuitBreakerRulesResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned result.
	Result []*ListCircuitBreakerRulesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of pages.
	//
	// example:
	//
	// 3
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListCircuitBreakerRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCircuitBreakerRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCircuitBreakerRulesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCircuitBreakerRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCircuitBreakerRulesResponseBodyData) GetResult() []*ListCircuitBreakerRulesResponseBodyDataResult {
	return s.Result
}

func (s *ListCircuitBreakerRulesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListCircuitBreakerRulesResponseBodyData) SetPageNumber(v int32) *ListCircuitBreakerRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyData) SetPageSize(v int32) *ListCircuitBreakerRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyData) SetResult(v []*ListCircuitBreakerRulesResponseBodyDataResult) *ListCircuitBreakerRulesResponseBodyData {
	s.Result = v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyData) SetTotalSize(v int32) *ListCircuitBreakerRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCircuitBreakerRulesResponseBodyDataResult struct {
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
	// Indicates whether the rule was enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The behavior that was bound to the rule.
	//
	// example:
	//
	// "{\\"appName\\":\\"spring-cloud-a\\",\\"fallbackBehavior\\":{\\"webFallbackMode\\":0,\\"webRespContentType\\":0,\\"webRespMessage\\":\\"Blocked\\",\\"webRespStatusCode\\":429},\\"id\\":977,\\"name\\":\\"Fallback\\",\\"namespace\\":\\"default\\",\\"resourceClassification\\":1}"
	FallbackObject *string `json:"FallbackObject,omitempty" xml:"FallbackObject,omitempty"`
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
	Resource     *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceType *int32  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The period in which circuit breaking is implemented. Unit: milliseconds. If circuit breaking is implemented on the requests for the route, the calls to all the requests for the route fail in the configured circuit breaking period.
	//
	// example:
	//
	// 10000
	RetryTimeoutMs *int32 `json:"RetryTimeoutMs,omitempty" xml:"RetryTimeoutMs,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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

func (s ListCircuitBreakerRulesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListCircuitBreakerRulesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetAppId() *string {
	return s.AppId
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetAppName() *string {
	return s.AppName
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetEnable() *bool {
	return s.Enable
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetFallbackObject() *string {
	return s.FallbackObject
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetHalfOpenBaseAmountPerStep() *int32 {
	return s.HalfOpenBaseAmountPerStep
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetHalfOpenRecoveryStepNum() *int32 {
	return s.HalfOpenRecoveryStepNum
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetMaxAllowedRtMs() *int32 {
	return s.MaxAllowedRtMs
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetMinRequestAmount() *int32 {
	return s.MinRequestAmount
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetNamespace() *string {
	return s.Namespace
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetResource() *string {
	return s.Resource
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetRetryTimeoutMs() *int32 {
	return s.RetryTimeoutMs
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetStatIntervalMs() *int32 {
	return s.StatIntervalMs
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetStrategy() *int32 {
	return s.Strategy
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetAppId(v string) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.AppId = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetAppName(v string) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.AppName = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetEnable(v bool) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.Enable = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetFallbackObject(v string) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.FallbackObject = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetHalfOpenBaseAmountPerStep(v int32) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.HalfOpenBaseAmountPerStep = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetHalfOpenRecoveryStepNum(v int32) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.HalfOpenRecoveryStepNum = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetMaxAllowedRtMs(v int32) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.MaxAllowedRtMs = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetMinRequestAmount(v int32) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.MinRequestAmount = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetNamespace(v string) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.Namespace = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetRegionId(v string) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.RegionId = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetResource(v string) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.Resource = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetResourceType(v int32) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.ResourceType = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetRetryTimeoutMs(v int32) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.RetryTimeoutMs = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetRuleId(v int64) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.RuleId = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetStatIntervalMs(v int32) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.StatIntervalMs = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetStrategy(v int32) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.Strategy = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) SetThreshold(v float32) *ListCircuitBreakerRulesResponseBodyDataResult {
	s.Threshold = &v
	return s
}

func (s *ListCircuitBreakerRulesResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
