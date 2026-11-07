// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebFlowRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListWebFlowRulesResponseBody
	GetCode() *int32
	SetData(v *ListWebFlowRulesResponseBodyData) *ListWebFlowRulesResponseBody
	GetData() *ListWebFlowRulesResponseBodyData
	SetHttpStatusCode(v int32) *ListWebFlowRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListWebFlowRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWebFlowRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWebFlowRulesResponseBody
	GetSuccess() *bool
}

type ListWebFlowRulesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListWebFlowRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 4AE73569-304C-5AA9-AE11-C1D99C7D1689
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWebFlowRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWebFlowRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWebFlowRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListWebFlowRulesResponseBody) GetData() *ListWebFlowRulesResponseBodyData {
	return s.Data
}

func (s *ListWebFlowRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListWebFlowRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWebFlowRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWebFlowRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWebFlowRulesResponseBody) SetCode(v int32) *ListWebFlowRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListWebFlowRulesResponseBody) SetData(v *ListWebFlowRulesResponseBodyData) *ListWebFlowRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListWebFlowRulesResponseBody) SetHttpStatusCode(v int32) *ListWebFlowRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListWebFlowRulesResponseBody) SetMessage(v string) *ListWebFlowRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListWebFlowRulesResponseBody) SetRequestId(v string) *ListWebFlowRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWebFlowRulesResponseBody) SetSuccess(v bool) *ListWebFlowRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListWebFlowRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWebFlowRulesResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned results.
	Result []*ListWebFlowRulesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of queried rules.
	//
	// example:
	//
	// 36
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListWebFlowRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWebFlowRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWebFlowRulesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWebFlowRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWebFlowRulesResponseBodyData) GetResult() []*ListWebFlowRulesResponseBodyDataResult {
	return s.Result
}

func (s *ListWebFlowRulesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListWebFlowRulesResponseBodyData) SetPageNumber(v int32) *ListWebFlowRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyData) SetPageSize(v int32) *ListWebFlowRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyData) SetResult(v []*ListWebFlowRulesResponseBodyDataResult) *ListWebFlowRulesResponseBodyData {
	s.Result = v
	return s
}

func (s *ListWebFlowRulesResponseBodyData) SetTotalSize(v int32) *ListWebFlowRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyData) Validate() error {
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

type ListWebFlowRulesResponseBodyDataResult struct {
	// The application ID.
	//
	// example:
	//
	// hkhon1XXXX@c3df23522bXXXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of additional requests allowed for burst requests. The effect is equivalent to the maximum number of tokens for the token bucket algorithm.
	//
	// example:
	//
	// 0
	Burst *int32 `json:"Burst,omitempty" xml:"Burst,omitempty"`
	// The flow control method. Valid values:
	//
	// - 0: fast failure
	//
	// - 2: wait in queue
	//
	// example:
	//
	// 0
	ControlBehavior *int32 `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	// Indicates whether the rule is enabled or disabled.
	//
	// - `true`: enabled.
	//
	// - `false`: disabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The behavior that was bound to the rule.
	//
	// example:
	//
	// {\\"appName\\":\\"spring-cloud-a\\",\\"fallbackBehavior\\":{\\"webFallbackMode\\":0,\\"webRespContentType\\":0,\\"webRespMessage\\":\\"Blocked\\",\\"webRespStatusCode\\":429},\\"id\\":977,\\"name\\":\\"Fallback\\",\\"namespace\\":\\"default\\",\\"resourceClassification\\":1}
	FallbackObject *string `json:"FallbackObject,omitempty" xml:"FallbackObject,omitempty"`
	// The timeout period for queuing when the ControlBehavior parameter is set to 2. Unit: milliseconds.
	//
	// example:
	//
	// 200
	MaxQueueingTimeMs *int32 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	// The threshold type. Valid values:
	//
	// - 1: number of requests
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
	// The parameter description item, which is a JSON string containing the following key-value pairs:
	//
	// - parseStrategy: the parameter property (0: client IP address. 1: remote host. 2: header. 3: URL parameter)
	//
	// - fieldName: the header name or URL parameter name. You must specify this key only if parseStrategy is set to 2 or 3
	//
	// - matchStrategy: the method to match the property value. You must specify this key only if parseStrategy is set to 2 or 3 and the property value is matched. Valid values: 0 (exact match), 2 (regular expression match), and 3 (substring match)
	//
	// - pattern: the match string. You must specify this key only if parseStrategy is set to 2 or 3 and the property value is matched
	//
	// example:
	//
	// {"fieldName":"testKey","matchStrategy":2,"parseStrategy":2,"pattern":"testValue"}
	ParamItem *string `json:"ParamItem,omitempty" xml:"ParamItem,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the interface resource.
	//
	// example:
	//
	// /flow
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The resource mode. Valid values:
	//
	// - 0: resource name mode
	//
	// example:
	//
	// 0
	ResourceMode *int32 `json:"ResourceMode,omitempty" xml:"ResourceMode,omitempty"`
	// The resource type
	//
	// - 0: custom interface
	//
	// - 1: WEB
	//
	// - 2: RPC
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 1
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The length of the time window. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	StatIntervalMs *int32 `json:"StatIntervalMs,omitempty" xml:"StatIntervalMs,omitempty"`
	// The threshold.
	//
	// example:
	//
	// 20
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListWebFlowRulesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListWebFlowRulesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetAppId() *string {
	return s.AppId
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetAppName() *string {
	return s.AppName
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetBurst() *int32 {
	return s.Burst
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetControlBehavior() *int32 {
	return s.ControlBehavior
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetEnable() *bool {
	return s.Enable
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetFallbackObject() *string {
	return s.FallbackObject
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetMaxQueueingTimeMs() *int32 {
	return s.MaxQueueingTimeMs
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetMetricType() *int32 {
	return s.MetricType
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetNamespace() *string {
	return s.Namespace
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetParamItem() *string {
	return s.ParamItem
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetResource() *string {
	return s.Resource
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetResourceMode() *int32 {
	return s.ResourceMode
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetRuleId() *string {
	return s.RuleId
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetStatIntervalMs() *int32 {
	return s.StatIntervalMs
}

func (s *ListWebFlowRulesResponseBodyDataResult) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetAppId(v string) *ListWebFlowRulesResponseBodyDataResult {
	s.AppId = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetAppName(v string) *ListWebFlowRulesResponseBodyDataResult {
	s.AppName = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetBurst(v int32) *ListWebFlowRulesResponseBodyDataResult {
	s.Burst = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetControlBehavior(v int32) *ListWebFlowRulesResponseBodyDataResult {
	s.ControlBehavior = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetEnable(v bool) *ListWebFlowRulesResponseBodyDataResult {
	s.Enable = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetFallbackObject(v string) *ListWebFlowRulesResponseBodyDataResult {
	s.FallbackObject = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetMaxQueueingTimeMs(v int32) *ListWebFlowRulesResponseBodyDataResult {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetMetricType(v int32) *ListWebFlowRulesResponseBodyDataResult {
	s.MetricType = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetNamespace(v string) *ListWebFlowRulesResponseBodyDataResult {
	s.Namespace = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetParamItem(v string) *ListWebFlowRulesResponseBodyDataResult {
	s.ParamItem = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetRegionId(v string) *ListWebFlowRulesResponseBodyDataResult {
	s.RegionId = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetResource(v string) *ListWebFlowRulesResponseBodyDataResult {
	s.Resource = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetResourceMode(v int32) *ListWebFlowRulesResponseBodyDataResult {
	s.ResourceMode = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetResourceType(v int32) *ListWebFlowRulesResponseBodyDataResult {
	s.ResourceType = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetRuleId(v string) *ListWebFlowRulesResponseBodyDataResult {
	s.RuleId = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetStatIntervalMs(v int32) *ListWebFlowRulesResponseBodyDataResult {
	s.StatIntervalMs = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) SetThreshold(v float32) *ListWebFlowRulesResponseBodyDataResult {
	s.Threshold = &v
	return s
}

func (s *ListWebFlowRulesResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
