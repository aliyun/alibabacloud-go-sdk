// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTargetsResponseBody
	GetCode() *string
	SetData(v *ListTargetsResponseBodyData) *ListTargetsResponseBody
	GetData() *ListTargetsResponseBodyData
	SetMessage(v string) *ListTargetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTargetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTargetsResponseBody
	GetSuccess() *bool
}

type ListTargetsResponseBody struct {
	// The returned response code. Valid values:
	//
	// 	- Success: The request is successful.
	//
	// 	- Other codes: The request failed. For a list of error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListTargetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// EventRuleNotExisted
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5DAF96FB-A4B6-548C-B999-0BFDCB2261B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTargetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTargetsResponseBody) GetData() *ListTargetsResponseBodyData {
	return s.Data
}

func (s *ListTargetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTargetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTargetsResponseBody) SetCode(v string) *ListTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTargetsResponseBody) SetData(v *ListTargetsResponseBodyData) *ListTargetsResponseBody {
	s.Data = v
	return s
}

func (s *ListTargetsResponseBody) SetMessage(v string) *ListTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTargetsResponseBody) SetRequestId(v string) *ListTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTargetsResponseBody) SetSuccess(v bool) *ListTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTargetsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTargetsResponseBodyData struct {
	// If excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The event targets.
	Targets []*ListTargetsResponseBodyDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 18
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTargetsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTargetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTargetsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTargetsResponseBodyData) GetTargets() []*ListTargetsResponseBodyDataTargets {
	return s.Targets
}

func (s *ListTargetsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListTargetsResponseBodyData) SetNextToken(v string) *ListTargetsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTargetsResponseBodyData) SetTargets(v []*ListTargetsResponseBodyDataTargets) *ListTargetsResponseBodyData {
	s.Targets = v
	return s
}

func (s *ListTargetsResponseBodyData) SetTotal(v int32) *ListTargetsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListTargetsResponseBodyData) Validate() error {
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTargetsResponseBodyDataTargets struct {
	// The concurrency configuration.
	ConcurrentConfig *ListTargetsResponseBodyDataTargetsConcurrentConfig `json:"ConcurrentConfig,omitempty" xml:"ConcurrentConfig,omitempty" type:"Struct"`
	// The endpoint of the event target.
	//
	// example:
	//
	// acs:fc:cn-hangzhou:123456789098****:services/guide.LATEST/functions/HelloFC
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The fault tolerance policy. Valid values: ALL and NONE.
	//
	// - **ALL**: Fault tolerance is allowed. If an error occurs in an event, event processing is not blocked. If the event fails to be sent after the maximum number of retries specified by the retry policy is reached, the event is delivered to the dead-letter queue or discarded based on your configurations.
	//
	// - **NONE**: Fault tolerance is not allowed. If an error occurs in an event and the event fails to be sent after the maximum number of retries specified by the retry policy is reached, event processing is blocked.
	//
	// example:
	//
	// ALL
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The name of the event bus.
	//
	// example:
	//
	// sls-beijing-one1-tf
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The ID of the event target.
	//
	// example:
	//
	// 1453
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters that are configured for the event target.
	ParamList []*ListTargetsResponseBodyDataTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The name of the event rule.
	//
	// example:
	//
	// rule-uKAK2
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the event target. For more information, see [Event target parameters](https://help.aliyun.com/document_detail/183698.html).
	//
	// example:
	//
	// acs.fc.function
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTargetsResponseBodyDataTargets) String() string {
	return dara.Prettify(s)
}

func (s ListTargetsResponseBodyDataTargets) GoString() string {
	return s.String()
}

func (s *ListTargetsResponseBodyDataTargets) GetConcurrentConfig() *ListTargetsResponseBodyDataTargetsConcurrentConfig {
	return s.ConcurrentConfig
}

func (s *ListTargetsResponseBodyDataTargets) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListTargetsResponseBodyDataTargets) GetErrorsTolerance() *string {
	return s.ErrorsTolerance
}

func (s *ListTargetsResponseBodyDataTargets) GetEventBusName() *string {
	return s.EventBusName
}

func (s *ListTargetsResponseBodyDataTargets) GetId() *string {
	return s.Id
}

func (s *ListTargetsResponseBodyDataTargets) GetParamList() []*ListTargetsResponseBodyDataTargetsParamList {
	return s.ParamList
}

func (s *ListTargetsResponseBodyDataTargets) GetRuleName() *string {
	return s.RuleName
}

func (s *ListTargetsResponseBodyDataTargets) GetType() *string {
	return s.Type
}

func (s *ListTargetsResponseBodyDataTargets) SetConcurrentConfig(v *ListTargetsResponseBodyDataTargetsConcurrentConfig) *ListTargetsResponseBodyDataTargets {
	s.ConcurrentConfig = v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetEndpoint(v string) *ListTargetsResponseBodyDataTargets {
	s.Endpoint = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetErrorsTolerance(v string) *ListTargetsResponseBodyDataTargets {
	s.ErrorsTolerance = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetEventBusName(v string) *ListTargetsResponseBodyDataTargets {
	s.EventBusName = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetId(v string) *ListTargetsResponseBodyDataTargets {
	s.Id = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetParamList(v []*ListTargetsResponseBodyDataTargetsParamList) *ListTargetsResponseBodyDataTargets {
	s.ParamList = v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetRuleName(v string) *ListTargetsResponseBodyDataTargets {
	s.RuleName = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetType(v string) *ListTargetsResponseBodyDataTargets {
	s.Type = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) Validate() error {
	if s.ConcurrentConfig != nil {
		if err := s.ConcurrentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ParamList != nil {
		for _, item := range s.ParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTargetsResponseBodyDataTargetsConcurrentConfig struct {
	// The maximum number of concurrent events allowed in the event target.
	//
	// example:
	//
	// 10
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s ListTargetsResponseBodyDataTargetsConcurrentConfig) String() string {
	return dara.Prettify(s)
}

func (s ListTargetsResponseBodyDataTargetsConcurrentConfig) GoString() string {
	return s.String()
}

func (s *ListTargetsResponseBodyDataTargetsConcurrentConfig) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *ListTargetsResponseBodyDataTargetsConcurrentConfig) SetConcurrency(v int64) *ListTargetsResponseBodyDataTargetsConcurrentConfig {
	s.Concurrency = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargetsConcurrentConfig) Validate() error {
	return dara.Validate(s)
}

type ListTargetsResponseBodyDataTargetsParamList struct {
	// The format that is used by the event target parameter.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The resource parameter of the event target.
	//
	// example:
	//
	// body
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	// The template that is used by the event target parameter.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value of the event target parameter.
	//
	// example:
	//
	// {\\"key\\"=\\"value\\"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTargetsResponseBodyDataTargetsParamList) String() string {
	return dara.Prettify(s)
}

func (s ListTargetsResponseBodyDataTargetsParamList) GoString() string {
	return s.String()
}

func (s *ListTargetsResponseBodyDataTargetsParamList) GetForm() *string {
	return s.Form
}

func (s *ListTargetsResponseBodyDataTargetsParamList) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *ListTargetsResponseBodyDataTargetsParamList) GetTemplate() *string {
	return s.Template
}

func (s *ListTargetsResponseBodyDataTargetsParamList) GetValue() *string {
	return s.Value
}

func (s *ListTargetsResponseBodyDataTargetsParamList) SetForm(v string) *ListTargetsResponseBodyDataTargetsParamList {
	s.Form = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargetsParamList) SetResourceKey(v string) *ListTargetsResponseBodyDataTargetsParamList {
	s.ResourceKey = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargetsParamList) SetTemplate(v string) *ListTargetsResponseBodyDataTargetsParamList {
	s.Template = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargetsParamList) SetValue(v string) *ListTargetsResponseBodyDataTargetsParamList {
	s.Value = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargetsParamList) Validate() error {
	return dara.Validate(s)
}
