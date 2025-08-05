// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRuleRequest
	GetDescription() *string
	SetEventBusName(v string) *CreateRuleRequest
	GetEventBusName() *string
	SetEventTargets(v []*CreateRuleRequestEventTargets) *CreateRuleRequest
	GetEventTargets() []*CreateRuleRequestEventTargets
	SetFilterPattern(v string) *CreateRuleRequest
	GetFilterPattern() *string
	SetRuleName(v string) *CreateRuleRequest
	GetRuleName() *string
	SetStatus(v string) *CreateRuleRequest
	GetStatus() *string
}

type CreateRuleRequest struct {
	// The description of the event bus.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event targets.
	EventTargets []*CreateRuleRequestEventTargets `json:"EventTargets,omitempty" xml:"EventTargets,omitempty" type:"Repeated"`
	// The event pattern, in JSON format. Valid values: stringEqual and stringExpression. You can specify up to five expressions in the map data structure in each field.
	//
	// You can specify up to five expressions in the map data structure in each field.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"source\\": [{\\"prefix\\": \\"acs.\\"}],\\"type\\": [{\\"prefix\\":\\"oss:ObjectReplication\\"}],\\"subject\\":[{\\"prefix\\":\\"acs:oss:cn-hangzhou:123456789098****:my-movie-bucket/\\", \\"suffix\\":\\".txt\\"}]}
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The name of the event rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// MNSRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event rule. Valid values: ENABLE: enables the event rule. It is the default status of the event rule. DISABLE: disables the event rule.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRuleRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *CreateRuleRequest) GetEventTargets() []*CreateRuleRequestEventTargets {
	return s.EventTargets
}

func (s *CreateRuleRequest) GetFilterPattern() *string {
	return s.FilterPattern
}

func (s *CreateRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateRuleRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateRuleRequest) SetDescription(v string) *CreateRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateRuleRequest) SetEventBusName(v string) *CreateRuleRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateRuleRequest) SetEventTargets(v []*CreateRuleRequestEventTargets) *CreateRuleRequest {
	s.EventTargets = v
	return s
}

func (s *CreateRuleRequest) SetFilterPattern(v string) *CreateRuleRequest {
	s.FilterPattern = &v
	return s
}

func (s *CreateRuleRequest) SetRuleName(v string) *CreateRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateRuleRequest) SetStatus(v string) *CreateRuleRequest {
	s.Status = &v
	return s
}

func (s *CreateRuleRequest) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestEventTargets struct {
	// The concurrency configuration.
	ConcurrentConfig *CreateRuleRequestEventTargetsConcurrentConfig `json:"ConcurrentConfig,omitempty" xml:"ConcurrentConfig,omitempty" type:"Struct"`
	// The dead-letter queue. Events that are not processed or whose maximum number of retries is exceeded are written to the dead-letter queue. You can use queues in ApsaraMQ for RocketMQ, Simple Message Queue (SMQ, formerly MNS), and ApsaraMQ for Kafka as dead-letter queues. You can also use event buses in EventBridge as dead-letter queues.
	DeadLetterQueue *CreateRuleRequestEventTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The endpoint of the event target.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:123456789098****:queues/myqueue
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The fault tolerance policy. Valid values: ALL and NONE. The value ALL specifies that fault tolerance is allowed. If an error occurs in an event, event processing is not blocked. If the event fails to be sent after the maximum number of retries specified by the retry policy is reached, the event is delivered to the dead-letter queue or discarded based on your configurations. The value NONE specifies that fault tolerance is not allowed. If an error occurs in an event and the event fails to be sent after the maximum number of retries specified by the retry policy is reached, event processing is blocked.
	//
	// example:
	//
	// ALL
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The ID of the event target.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12021
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters that are configured for the event target.
	ParamList []*CreateRuleRequestEventTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The retry policy that you want to use to push failed events. Valid values: BACKOFF_RETRY and EXPONENTIAL_DECAY_RETRY. BACKOFF_RETRY: A failed event can be retried up to three times. The interval between two consecutive retries is a random value from 10 seconds to 20 seconds. EXPONENTIAL_DECAY_RETRY: A failed event can be retried up to 176 times. The interval between two consecutive retries exponentially increases to a maximum of 512 seconds. The total retry time is 1 day. The specific retry intervals are 1, 2, 4, 8, 16, 32, 64, 128, 256, and 512 seconds. The interval of 512 seconds is used for 167 retries.
	//
	// example:
	//
	// BACKOFF_RETRY
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
	// The type of the event target. For more information, see [Event target parameters](https://help.aliyun.com/document_detail/185887.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs.mns.queue
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRuleRequestEventTargets) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestEventTargets) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestEventTargets) GetConcurrentConfig() *CreateRuleRequestEventTargetsConcurrentConfig {
	return s.ConcurrentConfig
}

func (s *CreateRuleRequestEventTargets) GetDeadLetterQueue() *CreateRuleRequestEventTargetsDeadLetterQueue {
	return s.DeadLetterQueue
}

func (s *CreateRuleRequestEventTargets) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateRuleRequestEventTargets) GetErrorsTolerance() *string {
	return s.ErrorsTolerance
}

func (s *CreateRuleRequestEventTargets) GetId() *string {
	return s.Id
}

func (s *CreateRuleRequestEventTargets) GetParamList() []*CreateRuleRequestEventTargetsParamList {
	return s.ParamList
}

func (s *CreateRuleRequestEventTargets) GetPushRetryStrategy() *string {
	return s.PushRetryStrategy
}

func (s *CreateRuleRequestEventTargets) GetType() *string {
	return s.Type
}

func (s *CreateRuleRequestEventTargets) SetConcurrentConfig(v *CreateRuleRequestEventTargetsConcurrentConfig) *CreateRuleRequestEventTargets {
	s.ConcurrentConfig = v
	return s
}

func (s *CreateRuleRequestEventTargets) SetDeadLetterQueue(v *CreateRuleRequestEventTargetsDeadLetterQueue) *CreateRuleRequestEventTargets {
	s.DeadLetterQueue = v
	return s
}

func (s *CreateRuleRequestEventTargets) SetEndpoint(v string) *CreateRuleRequestEventTargets {
	s.Endpoint = &v
	return s
}

func (s *CreateRuleRequestEventTargets) SetErrorsTolerance(v string) *CreateRuleRequestEventTargets {
	s.ErrorsTolerance = &v
	return s
}

func (s *CreateRuleRequestEventTargets) SetId(v string) *CreateRuleRequestEventTargets {
	s.Id = &v
	return s
}

func (s *CreateRuleRequestEventTargets) SetParamList(v []*CreateRuleRequestEventTargetsParamList) *CreateRuleRequestEventTargets {
	s.ParamList = v
	return s
}

func (s *CreateRuleRequestEventTargets) SetPushRetryStrategy(v string) *CreateRuleRequestEventTargets {
	s.PushRetryStrategy = &v
	return s
}

func (s *CreateRuleRequestEventTargets) SetType(v string) *CreateRuleRequestEventTargets {
	s.Type = &v
	return s
}

func (s *CreateRuleRequestEventTargets) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestEventTargetsConcurrentConfig struct {
	// The concurrency.
	//
	// example:
	//
	// 2
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s CreateRuleRequestEventTargetsConcurrentConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestEventTargetsConcurrentConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestEventTargetsConcurrentConfig) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *CreateRuleRequestEventTargetsConcurrentConfig) SetConcurrency(v int64) *CreateRuleRequestEventTargetsConcurrentConfig {
	s.Concurrency = &v
	return s
}

func (s *CreateRuleRequestEventTargetsConcurrentConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestEventTargetsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue. Events that are not processed or whose maximum number of retries is exceeded are written to the dead-letter queue. Queues in SMQ and ApsaraMQ for RocketMQ can be used as dead-letter queues.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:123456789098****:/queues/rule-deadletterqueue
	Arn             *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateRuleRequestEventTargetsDeadLetterQueue) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestEventTargetsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) GetArn() *string {
	return s.Arn
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) GetNetwork() *string {
	return s.Network
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) SetArn(v string) *CreateRuleRequestEventTargetsDeadLetterQueue {
	s.Arn = &v
	return s
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) SetNetwork(v string) *CreateRuleRequestEventTargetsDeadLetterQueue {
	s.Network = &v
	return s
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) SetSecurityGroupId(v string) *CreateRuleRequestEventTargetsDeadLetterQueue {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) SetVSwitchIds(v string) *CreateRuleRequestEventTargetsDeadLetterQueue {
	s.VSwitchIds = &v
	return s
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) SetVpcId(v string) *CreateRuleRequestEventTargetsDeadLetterQueue {
	s.VpcId = &v
	return s
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestEventTargetsParamList struct {
	// The format of input parameters for the event target. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
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
	// The structure of the template for the event target.
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

func (s CreateRuleRequestEventTargetsParamList) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestEventTargetsParamList) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestEventTargetsParamList) GetForm() *string {
	return s.Form
}

func (s *CreateRuleRequestEventTargetsParamList) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *CreateRuleRequestEventTargetsParamList) GetTemplate() *string {
	return s.Template
}

func (s *CreateRuleRequestEventTargetsParamList) GetValue() *string {
	return s.Value
}

func (s *CreateRuleRequestEventTargetsParamList) SetForm(v string) *CreateRuleRequestEventTargetsParamList {
	s.Form = &v
	return s
}

func (s *CreateRuleRequestEventTargetsParamList) SetResourceKey(v string) *CreateRuleRequestEventTargetsParamList {
	s.ResourceKey = &v
	return s
}

func (s *CreateRuleRequestEventTargetsParamList) SetTemplate(v string) *CreateRuleRequestEventTargetsParamList {
	s.Template = &v
	return s
}

func (s *CreateRuleRequestEventTargetsParamList) SetValue(v string) *CreateRuleRequestEventTargetsParamList {
	s.Value = &v
	return s
}

func (s *CreateRuleRequestEventTargetsParamList) Validate() error {
	return dara.Validate(s)
}
