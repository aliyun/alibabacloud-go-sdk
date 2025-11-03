// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRuleResponseBody
	GetCode() *string
	SetData(v *GetRuleResponseBodyData) *GetRuleResponseBody
	GetData() *GetRuleResponseBodyData
	SetMessage(v string) *GetRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRuleResponseBody
	GetSuccess() *bool
}

type GetRuleResponseBody struct {
	// The response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For a list of error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned parameters.
	Data *GetRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// The event rule not existed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2BC1857D-E633-5E79-B2C2-43EF5F7730D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. If the operation is successful, the value true is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRuleResponseBody) GetData() *GetRuleResponseBodyData {
	return s.Data
}

func (s *GetRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRuleResponseBody) SetCode(v string) *GetRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleResponseBody) SetData(v *GetRuleResponseBodyData) *GetRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleResponseBody) SetMessage(v string) *GetRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleResponseBody) SetRequestId(v string) *GetRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleResponseBody) SetSuccess(v bool) *GetRuleResponseBody {
	s.Success = &v
	return s
}

func (s *GetRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuleResponseBodyData struct {
	// The timestamp that indicates when the event rule was created.
	//
	// example:
	//
	// 1607071602000
	CreatedTimestamp *int64 `json:"CreatedTimestamp,omitempty" xml:"CreatedTimestamp,omitempty"`
	// The description of the event rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	//
	// example:
	//
	// Housekeeping-Bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event pattern, in JSON format. Valid values: stringEqual and stringExpression. You can specify up to five expressions in the map data structure in each field.
	//
	// You can specify up to five expressions in the map data structure in each field.
	//
	// example:
	//
	// {\\"source\\":[\\"acs.oss\\"],\\"type\\":[\\"oss:BucketQueried:GetBucketStat\\"]}
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the event rule.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:123456789098****:eventbus/default/rule/myRule3
	RuleARN *string `json:"RuleARN,omitempty" xml:"RuleARN,omitempty"`
	// The name of the event rule.
	//
	// example:
	//
	// ramrolechange-fc
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event rule. Valid values: ENABLE (default): The event rule is enabled. DISABLE: The event rule is disabled.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The event targets.
	Targets []*GetRuleResponseBodyDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s GetRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyData) GetCreatedTimestamp() *int64 {
	return s.CreatedTimestamp
}

func (s *GetRuleResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetRuleResponseBodyData) GetEventBusName() *string {
	return s.EventBusName
}

func (s *GetRuleResponseBodyData) GetFilterPattern() *string {
	return s.FilterPattern
}

func (s *GetRuleResponseBodyData) GetRuleARN() *string {
	return s.RuleARN
}

func (s *GetRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *GetRuleResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetRuleResponseBodyData) GetTargets() []*GetRuleResponseBodyDataTargets {
	return s.Targets
}

func (s *GetRuleResponseBodyData) SetCreatedTimestamp(v int64) *GetRuleResponseBodyData {
	s.CreatedTimestamp = &v
	return s
}

func (s *GetRuleResponseBodyData) SetDescription(v string) *GetRuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetRuleResponseBodyData) SetEventBusName(v string) *GetRuleResponseBodyData {
	s.EventBusName = &v
	return s
}

func (s *GetRuleResponseBodyData) SetFilterPattern(v string) *GetRuleResponseBodyData {
	s.FilterPattern = &v
	return s
}

func (s *GetRuleResponseBodyData) SetRuleARN(v string) *GetRuleResponseBodyData {
	s.RuleARN = &v
	return s
}

func (s *GetRuleResponseBodyData) SetRuleName(v string) *GetRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetRuleResponseBodyData) SetStatus(v string) *GetRuleResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRuleResponseBodyData) SetTargets(v []*GetRuleResponseBodyDataTargets) *GetRuleResponseBodyData {
	s.Targets = v
	return s
}

func (s *GetRuleResponseBodyData) Validate() error {
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

type GetRuleResponseBodyDataTargets struct {
	ConcurrentConfig *GetRuleResponseBodyDataTargetsConcurrentConfig `json:"ConcurrentConfig,omitempty" xml:"ConcurrentConfig,omitempty" type:"Struct"`
	// The dead-letter queue.
	DeadLetterQueue *GetRuleResponseBodyDataTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The information about the event target.
	DetailMap map[string]interface{} `json:"DetailMap,omitempty" xml:"DetailMap,omitempty"`
	// The endpoint of the event target.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:123456789098****:queues/myqueue
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The fault tolerance policy. Valid values: ALL and NONE. ALL: Fault tolerance is allowed. If an error occurs in an event, event processing is not blocked. If the event fails to be sent after the maximum number of retries specified by the retry policy is reached, the event is delivered to the dead-letter queue or discarded based on your configurations. NONE: Fault tolerance is not allowed. If an error occurs in an event and the event fails to be sent after the maximum number of retries specified by the retry policy is reached, event processing is blocked.
	//
	// example:
	//
	// ALL
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The ID of the event target.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters that are configured for the event target.
	ParamList []*GetRuleResponseBodyDataTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The retry policy that is used to push failed events. Valid values: BACKOFF_RETRY and EXPONENTIAL_DECAY_RETRY. BACKOFF_RETRY: backoff retry. A failed event can be retried up to three times. The interval between two consecutive retries is a random value between 10 seconds and 20 seconds. EXPONENTIAL_DECAY_RETRY: exponential decay retry. A failed event can be retried up to 176 times. The interval between two consecutive retries exponentially increases to a maximum of 512 seconds. The total retry time is 1 day. The specific retry intervals are 1, 2, 4, 8, 16, 32, 64, 128, 256, and 512 seconds. The interval of 512 seconds is used for 167 retries.
	//
	// example:
	//
	// BACKOFF_RETRY
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
	// The transformer that is used to push events.
	//
	// example:
	//
	// MATCHED_EVENT
	PushSelector *string `json:"PushSelector,omitempty" xml:"PushSelector,omitempty"`
	// The type of the event target. For more information, see [Event target parameters](https://help.aliyun.com/document_detail/185887.html).
	//
	// example:
	//
	// acs.mns.queue
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRuleResponseBodyDataTargets) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyDataTargets) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataTargets) GetConcurrentConfig() *GetRuleResponseBodyDataTargetsConcurrentConfig {
	return s.ConcurrentConfig
}

func (s *GetRuleResponseBodyDataTargets) GetDeadLetterQueue() *GetRuleResponseBodyDataTargetsDeadLetterQueue {
	return s.DeadLetterQueue
}

func (s *GetRuleResponseBodyDataTargets) GetDetailMap() map[string]interface{} {
	return s.DetailMap
}

func (s *GetRuleResponseBodyDataTargets) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetRuleResponseBodyDataTargets) GetErrorsTolerance() *string {
	return s.ErrorsTolerance
}

func (s *GetRuleResponseBodyDataTargets) GetId() *string {
	return s.Id
}

func (s *GetRuleResponseBodyDataTargets) GetParamList() []*GetRuleResponseBodyDataTargetsParamList {
	return s.ParamList
}

func (s *GetRuleResponseBodyDataTargets) GetPushRetryStrategy() *string {
	return s.PushRetryStrategy
}

func (s *GetRuleResponseBodyDataTargets) GetPushSelector() *string {
	return s.PushSelector
}

func (s *GetRuleResponseBodyDataTargets) GetType() *string {
	return s.Type
}

func (s *GetRuleResponseBodyDataTargets) SetConcurrentConfig(v *GetRuleResponseBodyDataTargetsConcurrentConfig) *GetRuleResponseBodyDataTargets {
	s.ConcurrentConfig = v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetDeadLetterQueue(v *GetRuleResponseBodyDataTargetsDeadLetterQueue) *GetRuleResponseBodyDataTargets {
	s.DeadLetterQueue = v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetDetailMap(v map[string]interface{}) *GetRuleResponseBodyDataTargets {
	s.DetailMap = v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetEndpoint(v string) *GetRuleResponseBodyDataTargets {
	s.Endpoint = &v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetErrorsTolerance(v string) *GetRuleResponseBodyDataTargets {
	s.ErrorsTolerance = &v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetId(v string) *GetRuleResponseBodyDataTargets {
	s.Id = &v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetParamList(v []*GetRuleResponseBodyDataTargetsParamList) *GetRuleResponseBodyDataTargets {
	s.ParamList = v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetPushRetryStrategy(v string) *GetRuleResponseBodyDataTargets {
	s.PushRetryStrategy = &v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetPushSelector(v string) *GetRuleResponseBodyDataTargets {
	s.PushSelector = &v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetType(v string) *GetRuleResponseBodyDataTargets {
	s.Type = &v
	return s
}

func (s *GetRuleResponseBodyDataTargets) Validate() error {
	if s.ConcurrentConfig != nil {
		if err := s.ConcurrentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DeadLetterQueue != nil {
		if err := s.DeadLetterQueue.Validate(); err != nil {
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

type GetRuleResponseBodyDataTargetsConcurrentConfig struct {
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s GetRuleResponseBodyDataTargetsConcurrentConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyDataTargetsConcurrentConfig) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataTargetsConcurrentConfig) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *GetRuleResponseBodyDataTargetsConcurrentConfig) SetConcurrency(v int64) *GetRuleResponseBodyDataTargetsConcurrentConfig {
	s.Concurrency = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsConcurrentConfig) Validate() error {
	return dara.Validate(s)
}

type GetRuleResponseBodyDataTargetsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:164901546557****:eventbus/my-event-bus/eventsource/myRocketMQ.source
	Arn             *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetRuleResponseBodyDataTargetsDeadLetterQueue) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyDataTargetsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) GetArn() *string {
	return s.Arn
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) GetNetwork() *string {
	return s.Network
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) GetVpcId() *string {
	return s.VpcId
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) SetArn(v string) *GetRuleResponseBodyDataTargetsDeadLetterQueue {
	s.Arn = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) SetNetwork(v string) *GetRuleResponseBodyDataTargetsDeadLetterQueue {
	s.Network = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) SetSecurityGroupId(v string) *GetRuleResponseBodyDataTargetsDeadLetterQueue {
	s.SecurityGroupId = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) SetVSwitchIds(v string) *GetRuleResponseBodyDataTargetsDeadLetterQueue {
	s.VSwitchIds = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) SetVpcId(v string) *GetRuleResponseBodyDataTargetsDeadLetterQueue {
	s.VpcId = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) Validate() error {
	return dara.Validate(s)
}

type GetRuleResponseBodyDataTargetsParamList struct {
	// The format that is used by the event target parameter. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The resource key of the event target. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// body
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	// The template based on which events are delivered to the event target.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The event target.
	//
	// example:
	//
	// {\\"key\\"=\\"value\\"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRuleResponseBodyDataTargetsParamList) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyDataTargetsParamList) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataTargetsParamList) GetForm() *string {
	return s.Form
}

func (s *GetRuleResponseBodyDataTargetsParamList) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *GetRuleResponseBodyDataTargetsParamList) GetTemplate() *string {
	return s.Template
}

func (s *GetRuleResponseBodyDataTargetsParamList) GetValue() *string {
	return s.Value
}

func (s *GetRuleResponseBodyDataTargetsParamList) SetForm(v string) *GetRuleResponseBodyDataTargetsParamList {
	s.Form = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsParamList) SetResourceKey(v string) *GetRuleResponseBodyDataTargetsParamList {
	s.ResourceKey = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsParamList) SetTemplate(v string) *GetRuleResponseBodyDataTargetsParamList {
	s.Template = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsParamList) SetValue(v string) *GetRuleResponseBodyDataTargetsParamList {
	s.Value = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsParamList) Validate() error {
	return dara.Validate(s)
}
