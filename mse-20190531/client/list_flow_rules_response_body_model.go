// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListFlowRulesResponseBody
	GetCode() *int32
	SetData(v *ListFlowRulesResponseBodyData) *ListFlowRulesResponseBody
	GetData() *ListFlowRulesResponseBodyData
	SetHttpStatusCode(v int32) *ListFlowRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListFlowRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFlowRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFlowRulesResponseBody
	GetSuccess() *bool
}

type ListFlowRulesResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	Data *ListFlowRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 316F5F64-F73D-42DC-8632-01E308B6****
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

func (s ListFlowRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlowRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListFlowRulesResponseBody) GetData() *ListFlowRulesResponseBodyData {
	return s.Data
}

func (s *ListFlowRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFlowRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFlowRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFlowRulesResponseBody) SetCode(v int32) *ListFlowRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlowRulesResponseBody) SetData(v *ListFlowRulesResponseBodyData) *ListFlowRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListFlowRulesResponseBody) SetHttpStatusCode(v int32) *ListFlowRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFlowRulesResponseBody) SetMessage(v string) *ListFlowRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlowRulesResponseBody) SetRequestId(v string) *ListFlowRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowRulesResponseBody) SetSuccess(v bool) *ListFlowRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListFlowRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFlowRulesResponseBodyData struct {
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
	// The data returned.
	Result []*ListFlowRulesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 11
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListFlowRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFlowRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlowRulesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlowRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlowRulesResponseBodyData) GetResult() []*ListFlowRulesResponseBodyDataResult {
	return s.Result
}

func (s *ListFlowRulesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListFlowRulesResponseBodyData) SetPageNumber(v int32) *ListFlowRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListFlowRulesResponseBodyData) SetPageSize(v int32) *ListFlowRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFlowRulesResponseBodyData) SetResult(v []*ListFlowRulesResponseBodyDataResult) *ListFlowRulesResponseBodyData {
	s.Result = v
	return s
}

func (s *ListFlowRulesResponseBodyData) SetTotalSize(v int32) *ListFlowRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListFlowRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListFlowRulesResponseBodyDataResult struct {
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@5f1b08becb*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// finance
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
	// Indicates whether the throttling rule was enabled.
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
	// The throttling behavior.
	//
	// example:
	//
	// "{\\"appName\\":\\"spring-cloud-a\\",\\"fallbackBehavior\\":{\\"webFallbackMode\\":0,\\"webRespContentType\\":0,\\"webRespMessage\\":\\"Blocked\\",\\"webRespStatusCode\\":429},\\"id\\":977,\\"name\\":\\"Fallback\\",\\"namespace\\":\\"default\\",\\"resourceClassification\\":1}"
	FallbackObject *string `json:"FallbackObject,omitempty" xml:"FallbackObject,omitempty"`
	// Requests source application.
	//
	// example:
	//
	// spring-cloud-a
	LimitApp *string `json:"LimitApp,omitempty" xml:"LimitApp,omitempty"`
	// The timeout period for queuing when the value of ControlBehavior is set to 2. Unit: milliseconds.
	//
	// example:
	//
	// 500
	MaxQueueingTimeMs *int32 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	// The statistical dimension. The value 1 indicates that the number of accepted requests is used for statistics.
	//
	// example:
	//
	// 1
	MetricType *int32 `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the interface resource.
	//
	// example:
	//
	// /a
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// Interface resource type.
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 2
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The throttling threshold.
	//
	// example:
	//
	// 10
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// {\\"origin\\":\\"default\\"}
	TrafficTags map[string]interface{} `json:"TrafficTags,omitempty" xml:"TrafficTags,omitempty"`
}

func (s ListFlowRulesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListFlowRulesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListFlowRulesResponseBodyDataResult) GetAppId() *string {
	return s.AppId
}

func (s *ListFlowRulesResponseBodyDataResult) GetAppName() *string {
	return s.AppName
}

func (s *ListFlowRulesResponseBodyDataResult) GetControlBehavior() *int32 {
	return s.ControlBehavior
}

func (s *ListFlowRulesResponseBodyDataResult) GetEnable() *bool {
	return s.Enable
}

func (s *ListFlowRulesResponseBodyDataResult) GetFallbackObject() *string {
	return s.FallbackObject
}

func (s *ListFlowRulesResponseBodyDataResult) GetLimitApp() *string {
	return s.LimitApp
}

func (s *ListFlowRulesResponseBodyDataResult) GetMaxQueueingTimeMs() *int32 {
	return s.MaxQueueingTimeMs
}

func (s *ListFlowRulesResponseBodyDataResult) GetMetricType() *int32 {
	return s.MetricType
}

func (s *ListFlowRulesResponseBodyDataResult) GetNamespace() *string {
	return s.Namespace
}

func (s *ListFlowRulesResponseBodyDataResult) GetRegionId() *string {
	return s.RegionId
}

func (s *ListFlowRulesResponseBodyDataResult) GetResource() *string {
	return s.Resource
}

func (s *ListFlowRulesResponseBodyDataResult) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *ListFlowRulesResponseBodyDataResult) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListFlowRulesResponseBodyDataResult) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ListFlowRulesResponseBodyDataResult) GetTrafficTags() map[string]interface{} {
	return s.TrafficTags
}

func (s *ListFlowRulesResponseBodyDataResult) SetAppId(v string) *ListFlowRulesResponseBodyDataResult {
	s.AppId = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetAppName(v string) *ListFlowRulesResponseBodyDataResult {
	s.AppName = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetControlBehavior(v int32) *ListFlowRulesResponseBodyDataResult {
	s.ControlBehavior = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetEnable(v bool) *ListFlowRulesResponseBodyDataResult {
	s.Enable = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetFallbackObject(v string) *ListFlowRulesResponseBodyDataResult {
	s.FallbackObject = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetLimitApp(v string) *ListFlowRulesResponseBodyDataResult {
	s.LimitApp = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetMaxQueueingTimeMs(v int32) *ListFlowRulesResponseBodyDataResult {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetMetricType(v int32) *ListFlowRulesResponseBodyDataResult {
	s.MetricType = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetNamespace(v string) *ListFlowRulesResponseBodyDataResult {
	s.Namespace = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetRegionId(v string) *ListFlowRulesResponseBodyDataResult {
	s.RegionId = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetResource(v string) *ListFlowRulesResponseBodyDataResult {
	s.Resource = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetResourceType(v int32) *ListFlowRulesResponseBodyDataResult {
	s.ResourceType = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetRuleId(v int64) *ListFlowRulesResponseBodyDataResult {
	s.RuleId = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetThreshold(v float32) *ListFlowRulesResponseBodyDataResult {
	s.Threshold = &v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) SetTrafficTags(v map[string]interface{}) *ListFlowRulesResponseBodyDataResult {
	s.TrafficTags = v
	return s
}

func (s *ListFlowRulesResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
