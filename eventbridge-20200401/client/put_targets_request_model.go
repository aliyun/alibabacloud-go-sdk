// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *PutTargetsRequest
	GetEventBusName() *string
	SetRuleName(v string) *PutTargetsRequest
	GetRuleName() *string
	SetTargets(v []*PutTargetsRequestTargets) *PutTargetsRequest
	GetTargets() []*PutTargetsRequestTargets
}

type PutTargetsRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// eventTest
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ssr-send-to-vendor-test01
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// A list of event targets to create or update. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// This parameter is required.
	Targets []*PutTargetsRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s PutTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s PutTargetsRequest) GoString() string {
	return s.String()
}

func (s *PutTargetsRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *PutTargetsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PutTargetsRequest) GetTargets() []*PutTargetsRequestTargets {
	return s.Targets
}

func (s *PutTargetsRequest) SetEventBusName(v string) *PutTargetsRequest {
	s.EventBusName = &v
	return s
}

func (s *PutTargetsRequest) SetRuleName(v string) *PutTargetsRequest {
	s.RuleName = &v
	return s
}

func (s *PutTargetsRequest) SetTargets(v []*PutTargetsRequestTargets) *PutTargetsRequest {
	s.Targets = v
	return s
}

func (s *PutTargetsRequest) Validate() error {
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

type PutTargetsRequestTargets struct {
	// The concurrency control settings.
	ConcurrentConfig *PutTargetsRequestTargetsConcurrentConfig `json:"ConcurrentConfig,omitempty" xml:"ConcurrentConfig,omitempty" type:"Struct"`
	// The dead-letter queue (DLQ) to which events are sent after all retry attempts fail. Supported DLQ types include Message Queue for Apache RocketMQ, Message Service (MNS), Message Queue for Apache Kafka, and EventBridge.
	DeadLetterQueue *PutTargetsRequestTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The endpoint of the event target.
	//
	// example:
	//
	// acs:fc:cn-hangzhou:123456789098****:services/guide.LATEST/functions/HelloFC
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The fault tolerance policy. Valid values:
	//
	// - **ALL**: Enables fault tolerance. If an error occurs, execution continues. After the retry attempts defined by the retry strategy are exhausted, the event is sent to the configured dead-letter queue or discarded.
	//
	// - **NONE**: Disables fault tolerance. If an error persists after all retry attempts fail, execution is blocked.
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
	// A list of parameters for the event target.
	ParamList []*PutTargetsRequestTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The retry strategy for pushing events. Valid values:
	//
	// - **BACKOFF_RETRY**: The event is retried up to three times at random intervals between 10 and 20 seconds.
	//
	// - **EXPONENTIAL_DECAY_RETRY**: The event is retried up to 176 times over 24 hours. The retry interval starts at 1 second, doubles with each attempt (1, 2, 4, ..., 256 seconds), and is capped at 512 seconds for all subsequent retries.
	//
	// example:
	//
	// BACKOFFRETRY
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
	// The type of the event target. For more information, see [Event target parameters](https://help.aliyun.com/document_detail/185887.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs.fc.function
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PutTargetsRequestTargets) String() string {
	return dara.Prettify(s)
}

func (s PutTargetsRequestTargets) GoString() string {
	return s.String()
}

func (s *PutTargetsRequestTargets) GetConcurrentConfig() *PutTargetsRequestTargetsConcurrentConfig {
	return s.ConcurrentConfig
}

func (s *PutTargetsRequestTargets) GetDeadLetterQueue() *PutTargetsRequestTargetsDeadLetterQueue {
	return s.DeadLetterQueue
}

func (s *PutTargetsRequestTargets) GetEndpoint() *string {
	return s.Endpoint
}

func (s *PutTargetsRequestTargets) GetErrorsTolerance() *string {
	return s.ErrorsTolerance
}

func (s *PutTargetsRequestTargets) GetId() *string {
	return s.Id
}

func (s *PutTargetsRequestTargets) GetParamList() []*PutTargetsRequestTargetsParamList {
	return s.ParamList
}

func (s *PutTargetsRequestTargets) GetPushRetryStrategy() *string {
	return s.PushRetryStrategy
}

func (s *PutTargetsRequestTargets) GetType() *string {
	return s.Type
}

func (s *PutTargetsRequestTargets) SetConcurrentConfig(v *PutTargetsRequestTargetsConcurrentConfig) *PutTargetsRequestTargets {
	s.ConcurrentConfig = v
	return s
}

func (s *PutTargetsRequestTargets) SetDeadLetterQueue(v *PutTargetsRequestTargetsDeadLetterQueue) *PutTargetsRequestTargets {
	s.DeadLetterQueue = v
	return s
}

func (s *PutTargetsRequestTargets) SetEndpoint(v string) *PutTargetsRequestTargets {
	s.Endpoint = &v
	return s
}

func (s *PutTargetsRequestTargets) SetErrorsTolerance(v string) *PutTargetsRequestTargets {
	s.ErrorsTolerance = &v
	return s
}

func (s *PutTargetsRequestTargets) SetId(v string) *PutTargetsRequestTargets {
	s.Id = &v
	return s
}

func (s *PutTargetsRequestTargets) SetParamList(v []*PutTargetsRequestTargetsParamList) *PutTargetsRequestTargets {
	s.ParamList = v
	return s
}

func (s *PutTargetsRequestTargets) SetPushRetryStrategy(v string) *PutTargetsRequestTargets {
	s.PushRetryStrategy = &v
	return s
}

func (s *PutTargetsRequestTargets) SetType(v string) *PutTargetsRequestTargets {
	s.Type = &v
	return s
}

func (s *PutTargetsRequestTargets) Validate() error {
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

type PutTargetsRequestTargetsConcurrentConfig struct {
	// The maximum number of concurrent executions for the event target.
	//
	// example:
	//
	// 34
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s PutTargetsRequestTargetsConcurrentConfig) String() string {
	return dara.Prettify(s)
}

func (s PutTargetsRequestTargetsConcurrentConfig) GoString() string {
	return s.String()
}

func (s *PutTargetsRequestTargetsConcurrentConfig) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *PutTargetsRequestTargetsConcurrentConfig) SetConcurrency(v int64) *PutTargetsRequestTargetsConcurrentConfig {
	s.Concurrency = &v
	return s
}

func (s *PutTargetsRequestTargetsConcurrentConfig) Validate() error {
	return dara.Validate(s)
}

type PutTargetsRequestTargetsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue.
	//
	// example:
	//
	// Acs:mns:cn-hangzhou:123456789098****:/queues/deadletterqueue
	//
	// or
	//
	// acs:mq:cn-hangzhou:123456789098****:/instances/MQ_INST_123456789098****_BX8QbBPL/topic/deadlettertopic
	//
	// or
	//
	// acs:alikafka:cn-hangzhou:123456789098****:instance/alikafka_post-cn-123456/topic/deadlettertopic
	//
	// or
	//
	// acs:eventbridge:cn-hangzhou:123456789098****:eventbus/deadletterbus
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The network type of the dead-letter queue.
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The VSwitch IDs.
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s PutTargetsRequestTargetsDeadLetterQueue) String() string {
	return dara.Prettify(s)
}

func (s PutTargetsRequestTargetsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) GetArn() *string {
	return s.Arn
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) GetNetwork() *string {
	return s.Network
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) GetVpcId() *string {
	return s.VpcId
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) SetArn(v string) *PutTargetsRequestTargetsDeadLetterQueue {
	s.Arn = &v
	return s
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) SetNetwork(v string) *PutTargetsRequestTargetsDeadLetterQueue {
	s.Network = &v
	return s
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) SetSecurityGroupId(v string) *PutTargetsRequestTargetsDeadLetterQueue {
	s.SecurityGroupId = &v
	return s
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) SetVSwitchIds(v string) *PutTargetsRequestTargetsDeadLetterQueue {
	s.VSwitchIds = &v
	return s
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) SetVpcId(v string) *PutTargetsRequestTargetsDeadLetterQueue {
	s.VpcId = &v
	return s
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) Validate() error {
	return dara.Validate(s)
}

type PutTargetsRequestTargetsParamList struct {
	// The format of the parameter value. For more information, see [Event target parameters](https://help.aliyun.com/document_detail/185887.html).
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The key of the parameter. For more information, see [Event target parameters](https://help.aliyun.com/document_detail/185887.html).
	//
	// example:
	//
	// body
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	// The template for the parameter value. This parameter applies only when `Form` is set to `TEMPLATE`.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// {\\"key\\"=\\"value\\"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutTargetsRequestTargetsParamList) String() string {
	return dara.Prettify(s)
}

func (s PutTargetsRequestTargetsParamList) GoString() string {
	return s.String()
}

func (s *PutTargetsRequestTargetsParamList) GetForm() *string {
	return s.Form
}

func (s *PutTargetsRequestTargetsParamList) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *PutTargetsRequestTargetsParamList) GetTemplate() *string {
	return s.Template
}

func (s *PutTargetsRequestTargetsParamList) GetValue() *string {
	return s.Value
}

func (s *PutTargetsRequestTargetsParamList) SetForm(v string) *PutTargetsRequestTargetsParamList {
	s.Form = &v
	return s
}

func (s *PutTargetsRequestTargetsParamList) SetResourceKey(v string) *PutTargetsRequestTargetsParamList {
	s.ResourceKey = &v
	return s
}

func (s *PutTargetsRequestTargetsParamList) SetTemplate(v string) *PutTargetsRequestTargetsParamList {
	s.Template = &v
	return s
}

func (s *PutTargetsRequestTargetsParamList) SetValue(v string) *PutTargetsRequestTargetsParamList {
	s.Value = &v
	return s
}

func (s *PutTargetsRequestTargetsParamList) Validate() error {
	return dara.Validate(s)
}
