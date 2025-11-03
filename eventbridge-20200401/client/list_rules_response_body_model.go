// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRulesResponseBody
	GetCode() *string
	SetData(v *ListRulesResponseBodyData) *ListRulesResponseBody
	GetData() *ListRulesResponseBodyData
	SetMessage(v string) *ListRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRulesResponseBody
	GetSuccess() *bool
}

type ListRulesResponseBody struct {
	// The response code.
	//
	// 	- **Success**: The request was successful.
	//
	// 	- **Other codes**: The request failed. For more information about error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// Specified parameter Limit is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C7043799-F4DA-5290-9249-97C359876D97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRulesResponseBody) GetData() *ListRulesResponseBodyData {
	return s.Data
}

func (s *ListRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRulesResponseBody) SetCode(v string) *ListRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRulesResponseBody) SetData(v *ListRulesResponseBodyData) *ListRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListRulesResponseBody) SetMessage(v string) *ListRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetSuccess(v bool) *ListRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRulesResponseBodyData struct {
	// If excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 1000
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The event rules.
	Rules []*ListRulesResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 6
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRulesResponseBodyData) GetRules() []*ListRulesResponseBodyDataRules {
	return s.Rules
}

func (s *ListRulesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListRulesResponseBodyData) SetNextToken(v string) *ListRulesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListRulesResponseBodyData) SetRules(v []*ListRulesResponseBodyDataRules) *ListRulesResponseBodyData {
	s.Rules = v
	return s
}

func (s *ListRulesResponseBodyData) SetTotal(v int32) *ListRulesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListRulesResponseBodyData) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRulesResponseBodyDataRules struct {
	// The creation timestamp.
	//
	// example:
	//
	// 1607071602000
	CreatedTimestamp *int64 `json:"CreatedTimestamp,omitempty" xml:"CreatedTimestamp,omitempty"`
	// The rule description.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the event rule.
	DetailMap map[string]interface{} `json:"DetailMap,omitempty" xml:"DetailMap,omitempty"`
	// The name of the event bus.
	//
	// example:
	//
	// demo
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event pattern, in JSON format. Valid values:
	//
	// 	- **stringEqual**: Up to five expressions in the map data structure can be specified in each field.
	//
	// 	- **stringExpression**: Up to five expressions in the map data structure can be specified in each field.
	//
	// example:
	//
	// {\\"source\\":[\\"acs.oss\\"],\\"type\\":[\\"oss:BucketQueried:GetBucketStat\\"]}
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the rule.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:123456789098****:eventbus/default/rule/myRule3
	RuleARN *string `json:"RuleARN,omitempty" xml:"RuleARN,omitempty"`
	// The name of the event rule.
	//
	// example:
	//
	// tf-testacc-rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event rule. Valid values:
	//
	// 	- **ENABLE*	- (default)
	//
	// 	- **DISABLE**
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The event targets.
	Targets []*ListRulesResponseBodyDataRulesTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyDataRules) GetCreatedTimestamp() *int64 {
	return s.CreatedTimestamp
}

func (s *ListRulesResponseBodyDataRules) GetDescription() *string {
	return s.Description
}

func (s *ListRulesResponseBodyDataRules) GetDetailMap() map[string]interface{} {
	return s.DetailMap
}

func (s *ListRulesResponseBodyDataRules) GetEventBusName() *string {
	return s.EventBusName
}

func (s *ListRulesResponseBodyDataRules) GetFilterPattern() *string {
	return s.FilterPattern
}

func (s *ListRulesResponseBodyDataRules) GetRuleARN() *string {
	return s.RuleARN
}

func (s *ListRulesResponseBodyDataRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRulesResponseBodyDataRules) GetStatus() *string {
	return s.Status
}

func (s *ListRulesResponseBodyDataRules) GetTargets() []*ListRulesResponseBodyDataRulesTargets {
	return s.Targets
}

func (s *ListRulesResponseBodyDataRules) SetCreatedTimestamp(v int64) *ListRulesResponseBodyDataRules {
	s.CreatedTimestamp = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetDescription(v string) *ListRulesResponseBodyDataRules {
	s.Description = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetDetailMap(v map[string]interface{}) *ListRulesResponseBodyDataRules {
	s.DetailMap = v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetEventBusName(v string) *ListRulesResponseBodyDataRules {
	s.EventBusName = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetFilterPattern(v string) *ListRulesResponseBodyDataRules {
	s.FilterPattern = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetRuleARN(v string) *ListRulesResponseBodyDataRules {
	s.RuleARN = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetRuleName(v string) *ListRulesResponseBodyDataRules {
	s.RuleName = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetStatus(v string) *ListRulesResponseBodyDataRules {
	s.Status = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetTargets(v []*ListRulesResponseBodyDataRulesTargets) *ListRulesResponseBodyDataRules {
	s.Targets = v
	return s
}

func (s *ListRulesResponseBodyDataRules) Validate() error {
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

type ListRulesResponseBodyDataRulesTargets struct {
	// The endpoint of the event target.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:123456789098****:queues/myqueue
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The fault tolerance policy. Valid values:
	//
	// 	- **ALL**: Fault tolerance is allowed. If an error occurs in an event, event processing is not blocked. If the event fails to be sent after the maximum number of retries specified by the retry policy is reached, the event is delivered to the dead-letter queue or discarded based on your configurations.
	//
	// 	- **NONE**: Fault tolerance is prohibited. If an error occurs in an event and the event fails to be sent after the maximum number of retries specified by the retry policy is reached, event processing is blocked.
	//
	// example:
	//
	// ALL
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The ID of the custom event target.
	//
	// example:
	//
	// 177
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The transformer that is used to push events.
	//
	// example:
	//
	// MATCHED_EVENT
	PushSelector *string `json:"PushSelector,omitempty" xml:"PushSelector,omitempty"`
	// The type of the event target. For more information, see [Event target parameters.](https://www.alibabacloud.com/help/en/eventbridge/latest/event-target-parameters)
	//
	// example:
	//
	// acs.mns.queue
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRulesResponseBodyDataRulesTargets) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyDataRulesTargets) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyDataRulesTargets) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListRulesResponseBodyDataRulesTargets) GetErrorsTolerance() *string {
	return s.ErrorsTolerance
}

func (s *ListRulesResponseBodyDataRulesTargets) GetId() *string {
	return s.Id
}

func (s *ListRulesResponseBodyDataRulesTargets) GetPushSelector() *string {
	return s.PushSelector
}

func (s *ListRulesResponseBodyDataRulesTargets) GetType() *string {
	return s.Type
}

func (s *ListRulesResponseBodyDataRulesTargets) SetEndpoint(v string) *ListRulesResponseBodyDataRulesTargets {
	s.Endpoint = &v
	return s
}

func (s *ListRulesResponseBodyDataRulesTargets) SetErrorsTolerance(v string) *ListRulesResponseBodyDataRulesTargets {
	s.ErrorsTolerance = &v
	return s
}

func (s *ListRulesResponseBodyDataRulesTargets) SetId(v string) *ListRulesResponseBodyDataRulesTargets {
	s.Id = &v
	return s
}

func (s *ListRulesResponseBodyDataRulesTargets) SetPushSelector(v string) *ListRulesResponseBodyDataRulesTargets {
	s.PushSelector = &v
	return s
}

func (s *ListRulesResponseBodyDataRulesTargets) SetType(v string) *ListRulesResponseBodyDataRulesTargets {
	s.Type = &v
	return s
}

func (s *ListRulesResponseBodyDataRulesTargets) Validate() error {
	return dara.Validate(s)
}
