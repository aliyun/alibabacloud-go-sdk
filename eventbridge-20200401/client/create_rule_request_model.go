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
	//
	// example:
	//
	// Filter rule for Simple Message Queue (formerly MNS) (SMQ)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The list of event targets.
	EventTargets []*CreateRuleRequestEventTargets `json:"EventTargets,omitempty" xml:"EventTargets,omitempty" type:"Repeated"`
	// The event pattern in JSON format. The following patterns are supported:
	//
	// - stringEqual pattern
	//
	// - stringExpression pattern
	//
	// Each field supports up to five expressions in a map structure.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "source": [
	//
	//     {
	//
	//       "prefix": "acs."
	//
	//     }
	//
	//   ],
	//
	//   "type": [
	//
	//     {
	//
	//       "prefix": "oss:ObjectReplication"
	//
	//     }
	//
	//   ],
	//
	//   "subject": [
	//
	//     {
	//
	//       "prefix": "acs:oss:cn-hangzhou:123456789098****:my-movie-bucket/",
	//
	//       "suffix": ".txt"
	//
	//     }
	//
	//   ]
	//
	// }
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The name of the event rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMQRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// - ENABLE: The rule is enabled. This is the default status.
	//
	// - DISABLE: The rule is disabled.
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
	if s.EventTargets != nil {
		for _, item := range s.EventTargets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRuleRequestEventTargets struct {
	// The concurrency control configuration.
	ConcurrentConfig *CreateRuleRequestEventTargetsConcurrentConfig `json:"ConcurrentConfig,omitempty" xml:"ConcurrentConfig,omitempty" type:"Struct"`
	// The dead-letter queue. Unprocessed events or events that exceed the retry count are written to the dead-letter queue. Supported queue types include ApsaraMQ for RocketMQ, Simple Message Queue (formerly MNS), ApsaraMQ for Kafka, and EventBridge event buses.
	DeadLetterQueue *CreateRuleRequestEventTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The endpoint URL for delivery.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:123456789098****:queues/myqueue
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The fault tolerance strategy. Valid values:
	//
	// - ALL: allows fault tolerance. Exceptions do not block execution. If the retry strategy is exhausted, messages are delivered to the dead-letter queue or discarded based on the configuration.
	//
	// - NONE: disables fault tolerance. Exceptions block execution when the retry strategy is exhausted.
	//
	// example:
	//
	// ALL
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The custom ID of the event target.
	//
	// This parameter is required.
	//
	// example:
	//
	// Mlm123456JHd2RsRoKw
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters of the event target.
	ParamList []*CreateRuleRequestEventTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The retry policy for pushing events. Valid values:
	//
	// - BACKOFF_RETRY: backoff retry policy. The system retries 3 times, with a random interval of 10 to 20 seconds between each retry.
	//
	// - EXPONENTIAL_DECAY_RETRY: exponential decay retry. The system retries 176 times, with the interval exponentially incrementing to 512 seconds. The total retry time is 1 day. The specific intervals are 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, and then 512 seconds for the remaining 167 retries.
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

type CreateRuleRequestEventTargetsConcurrentConfig struct {
	// The number of concurrent executions.
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
	// The ARN of the dead-letter queue. Unprocessed events or events that exceed the retry count are written to the dead-letter queue. Simple Message Queue (formerly MNS) and ApsaraMQ for RocketMQ are supported.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:123456789098****:/queues/deadletterqueue
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The network type.
	//
	// example:
	//
	// PrivateNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-2d****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-m12345****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the VPC where the dead-letter queue resides.
	//
	// example:
	//
	// vpc-2z****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	// The format of the event target parameter. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The resource parameter of the event target. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// body
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	// The parameter template style of the event target.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value of the event target parameter.
	//
	// example:
	//
	// {"key"="value"}
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
