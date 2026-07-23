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
	// The description of the event rule.
	//
	// example:
	//
	// SMQ filter rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// A list of event targets.
	EventTargets []*CreateRuleRequestEventTargets `json:"EventTargets,omitempty" xml:"EventTargets,omitempty" type:"Repeated"`
	// The event pattern, in JSON format. Supported pattern types are `stringEqual` and `stringExpression`. Each field can contain a maximum of five expressions in a map structure.
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
	// The status of the event rule. Valid values: `ENABLE`: The rule is enabled. This is the default value. `DISABLE`: The rule is disabled.
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
	// The dead-letter queue. If an event fails to be processed or exceeds the retry limit, it is sent to the dead-letter queue. Supported services for the dead-letter queue include Message Queue for Apache RocketMQ, Message Service (MNS), Message Queue for Apache Kafka, and EventBridge event buses.
	DeadLetterQueue *CreateRuleRequestEventTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The delivery endpoint for events.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:123456789098****:queues/myqueue
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The fault tolerance policy. Valid values:<br>`ALL`: Enables fault tolerance. Execution continues even if an error occurs. After all retry attempts fail, the event is sent to the dead-letter queue (if configured) or discarded.<br>`NONE`: Disables fault tolerance. Execution is blocked if an error occurs and all retry attempts fail.<br><br>
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
	// The parameters for the event target.
	ParamList []*CreateRuleRequestEventTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The push retry strategy. Valid values:<br>`BACKOFF_RETRY`: A backoff retry strategy where the system makes three retry attempts at random intervals of 10 to 20 seconds.<br>`EXPONENTIAL_DECAY_RETRY`: An exponential decay retry strategy where the system makes 176 retry attempts over 24 hours. The interval starts at 1 second and doubles with each of the first 10 attempts (up to 512 seconds). Subsequent retries occur every 512 seconds.<br><br>
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
	// The maximum number of concurrent executions for the event target.
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
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue. Events that fail to be processed or exceed the retry limit are sent to this ARN. Supported services for this parameter include Message Service (MNS) and Message Queue for Apache RocketMQ.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:123456789098****:/queues/deadletterqueue
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The network type.
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The VPC ID.
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
	// The name of the target parameter. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// body
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	// The template for the event target parameter.
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
