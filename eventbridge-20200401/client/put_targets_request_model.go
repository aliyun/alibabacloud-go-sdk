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
	// The event targets to be created or updated. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
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
	// The concurrency configuration.
	ConcurrentConfig *PutTargetsRequestTargetsConcurrentConfig `json:"ConcurrentConfig,omitempty" xml:"ConcurrentConfig,omitempty" type:"Struct"`
	// The dead-letter queue. Events that are not processed or whose maximum retries are exceeded are written to the dead-letter queue. You can use queues in ApsaraMQ for RocketMQ, Simple Message Queue (SMQ, formerly MNS), and ApsaraMQ for Kafka as dead-letter queues. You can also use event buses in EventBridge as dead-letter queues.
	DeadLetterQueue *PutTargetsRequestTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The endpoint of the event target.
	//
	// example:
	//
	// acs:fc:cn-hangzhou:123456789098****:services/guide.LATEST/functions/HelloFC
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The fault tolerance policy. Valid values:
	//
	// 	- **ALL**: allows fault tolerance. If an error occurs, event processing is not blocked. If the message exceeds the number of retries specified by the retry policy, the message is delivered to a dead-letter queue or discarded based on your configurations.
	//
	// 	- **NONE**: prohibits fault tolerance. If an error occurs and the message exceeds the number of retries specified by the retry policy, event processing is blocked.
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
	// Mlm123456JHd2RsRoKw
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters that are configured for the event target.
	ParamList []*PutTargetsRequestTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The retry policy to be used to push events. Valid values:
	//
	// 	- **BACKOFF_RETRY**: backoff retry. A failed event can be retried up to three times. The interval between two consecutive retries is a random value from 10 seconds to 20 seconds.
	//
	// 	- **EXPONENTIAL_DECAY_RETRY**: exponential decay retry. A failed event can be retried up to 176 times. The interval between two consecutive retries exponentially increases to a maximum of 512 seconds. The total retry time is 1 day. The specific retry intervals are 1, 2, 4, 8, 16, 32, 64, 128, 256, and 512 seconds. The interval of 512 seconds is used for 167 retries.
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
	// The concurrency.
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
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue. Events that are not processed or whose maximum retries are exceeded are written to the dead-letter queue.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:123456789098****:/queues/deadletterqueue or acs:mq:cn-hangzhou:123456789098****:/instances/MQ_INST_123456789098****_BX8QbBPL/topic/deadlettertopic or acs:alikafka:cn-hangzhou:123456789098****:instance/alikafka_post-cn-123456/topic/deadlettertopic or acs:eventbridge:cn-hangzhou:123456789098****:eventbus/deadletterbus
	Arn             *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	// The format of input parameters for the event target. For more information, see [Event target parameters](https://help.aliyun.com/document_detail/185887.html).
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The resource key of the event target. For more information, see [Event target parameters](https://help.aliyun.com/document_detail/185887.html).
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
	// The event target.
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
