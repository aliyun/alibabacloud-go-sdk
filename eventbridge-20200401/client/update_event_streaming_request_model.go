// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventStreamingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateEventStreamingRequest
	GetDescription() *string
	SetEventStreamingName(v string) *UpdateEventStreamingRequest
	GetEventStreamingName() *string
	SetFilterPattern(v string) *UpdateEventStreamingRequest
	GetFilterPattern() *string
	SetMetadata(v string) *UpdateEventStreamingRequest
	GetMetadata() *string
	SetRunOptions(v *UpdateEventStreamingRequestRunOptions) *UpdateEventStreamingRequest
	GetRunOptions() *UpdateEventStreamingRequestRunOptions
	SetSink(v *UpdateEventStreamingRequestSink) *UpdateEventStreamingRequest
	GetSink() *UpdateEventStreamingRequestSink
	SetSource(v *UpdateEventStreamingRequestSource) *UpdateEventStreamingRequest
	GetSource() *UpdateEventStreamingRequestSource
	SetTransforms(v []*UpdateEventStreamingRequestTransforms) *UpdateEventStreamingRequest
	GetTransforms() []*UpdateEventStreamingRequestTransforms
}

type UpdateEventStreamingRequest struct {
	// The description of the event stream.
	//
	// example:
	//
	// rocketmq2mns
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// myeventstreaming
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	// The event filtering rule. If you do not specify this parameter, all events are matched. For more information, see [https://www.alibabacloud.com/help/en/eventbridge/user-guide/event-patterns](https://www.alibabacloud.com/help/en/eventbridge/user-guide/event-patterns)
	//
	// example:
	//
	// {
	//
	//     "source": [
	//
	//         {
	//
	//             "prefix": "acs:mns"
	//
	//         }
	//
	//     ],
	//
	//     "type": [
	//
	//         {
	//
	//             "prefix": "mns:Queue"
	//
	//         }
	//
	//     ],
	//
	//     "subject": [
	//
	//         {
	//
	//             "prefix": "acs:mns:cn-hangzhou:123456789098****:queues/zeus"
	//
	//         }
	//
	//     ]
	//
	// }
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	Metadata      *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The runtime parameters.
	RunOptions *UpdateEventStreamingRequestRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	// The event target. You must select one and only one Sink type.
	Sink *UpdateEventStreamingRequestSink `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	// The event provider. You must select one and only one Source type.
	Source *UpdateEventStreamingRequestSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// The Transform-related configurations.
	Transforms []*UpdateEventStreamingRequestTransforms `json:"Transforms,omitempty" xml:"Transforms,omitempty" type:"Repeated"`
}

func (s UpdateEventStreamingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEventStreamingRequest) GetEventStreamingName() *string {
	return s.EventStreamingName
}

func (s *UpdateEventStreamingRequest) GetFilterPattern() *string {
	return s.FilterPattern
}

func (s *UpdateEventStreamingRequest) GetMetadata() *string {
	return s.Metadata
}

func (s *UpdateEventStreamingRequest) GetRunOptions() *UpdateEventStreamingRequestRunOptions {
	return s.RunOptions
}

func (s *UpdateEventStreamingRequest) GetSink() *UpdateEventStreamingRequestSink {
	return s.Sink
}

func (s *UpdateEventStreamingRequest) GetSource() *UpdateEventStreamingRequestSource {
	return s.Source
}

func (s *UpdateEventStreamingRequest) GetTransforms() []*UpdateEventStreamingRequestTransforms {
	return s.Transforms
}

func (s *UpdateEventStreamingRequest) SetDescription(v string) *UpdateEventStreamingRequest {
	s.Description = &v
	return s
}

func (s *UpdateEventStreamingRequest) SetEventStreamingName(v string) *UpdateEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

func (s *UpdateEventStreamingRequest) SetFilterPattern(v string) *UpdateEventStreamingRequest {
	s.FilterPattern = &v
	return s
}

func (s *UpdateEventStreamingRequest) SetMetadata(v string) *UpdateEventStreamingRequest {
	s.Metadata = &v
	return s
}

func (s *UpdateEventStreamingRequest) SetRunOptions(v *UpdateEventStreamingRequestRunOptions) *UpdateEventStreamingRequest {
	s.RunOptions = v
	return s
}

func (s *UpdateEventStreamingRequest) SetSink(v *UpdateEventStreamingRequestSink) *UpdateEventStreamingRequest {
	s.Sink = v
	return s
}

func (s *UpdateEventStreamingRequest) SetSource(v *UpdateEventStreamingRequestSource) *UpdateEventStreamingRequest {
	s.Source = v
	return s
}

func (s *UpdateEventStreamingRequest) SetTransforms(v []*UpdateEventStreamingRequestTransforms) *UpdateEventStreamingRequest {
	s.Transforms = v
	return s
}

func (s *UpdateEventStreamingRequest) Validate() error {
	if s.RunOptions != nil {
		if err := s.RunOptions.Validate(); err != nil {
			return err
		}
	}
	if s.Sink != nil {
		if err := s.Sink.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	if s.Transforms != nil {
		for _, item := range s.Transforms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateEventStreamingRequestRunOptions struct {
	// The batching window.
	BatchWindow    *UpdateEventStreamingRequestRunOptionsBatchWindow    `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	BusinessOption *UpdateEventStreamingRequestRunOptionsBusinessOption `json:"BusinessOption,omitempty" xml:"BusinessOption,omitempty" type:"Struct"`
	// Specifies whether to enable the dead-letter queue. The dead-letter queue is disabled by default. Messages that exceed the retry policy are discarded.
	DeadLetterQueue *UpdateEventStreamingRequestRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The error tolerance policy: NONE (no error tolerance) or ALL (tolerate all errors).
	//
	// example:
	//
	// ALL
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The concurrency.
	//
	// example:
	//
	// 2
	MaximumTasks *int64 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The retry strategy when event push fails.
	RetryStrategy *UpdateEventStreamingRequestRunOptionsRetryStrategy `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
	Throttling    *int32                                              `json:"Throttling,omitempty" xml:"Throttling,omitempty"`
}

func (s UpdateEventStreamingRequestRunOptions) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestRunOptions) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestRunOptions) GetBatchWindow() *UpdateEventStreamingRequestRunOptionsBatchWindow {
	return s.BatchWindow
}

func (s *UpdateEventStreamingRequestRunOptions) GetBusinessOption() *UpdateEventStreamingRequestRunOptionsBusinessOption {
	return s.BusinessOption
}

func (s *UpdateEventStreamingRequestRunOptions) GetDeadLetterQueue() *UpdateEventStreamingRequestRunOptionsDeadLetterQueue {
	return s.DeadLetterQueue
}

func (s *UpdateEventStreamingRequestRunOptions) GetErrorsTolerance() *string {
	return s.ErrorsTolerance
}

func (s *UpdateEventStreamingRequestRunOptions) GetMaximumTasks() *int64 {
	return s.MaximumTasks
}

func (s *UpdateEventStreamingRequestRunOptions) GetRetryStrategy() *UpdateEventStreamingRequestRunOptionsRetryStrategy {
	return s.RetryStrategy
}

func (s *UpdateEventStreamingRequestRunOptions) GetThrottling() *int32 {
	return s.Throttling
}

func (s *UpdateEventStreamingRequestRunOptions) SetBatchWindow(v *UpdateEventStreamingRequestRunOptionsBatchWindow) *UpdateEventStreamingRequestRunOptions {
	s.BatchWindow = v
	return s
}

func (s *UpdateEventStreamingRequestRunOptions) SetBusinessOption(v *UpdateEventStreamingRequestRunOptionsBusinessOption) *UpdateEventStreamingRequestRunOptions {
	s.BusinessOption = v
	return s
}

func (s *UpdateEventStreamingRequestRunOptions) SetDeadLetterQueue(v *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) *UpdateEventStreamingRequestRunOptions {
	s.DeadLetterQueue = v
	return s
}

func (s *UpdateEventStreamingRequestRunOptions) SetErrorsTolerance(v string) *UpdateEventStreamingRequestRunOptions {
	s.ErrorsTolerance = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptions) SetMaximumTasks(v int64) *UpdateEventStreamingRequestRunOptions {
	s.MaximumTasks = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptions) SetRetryStrategy(v *UpdateEventStreamingRequestRunOptionsRetryStrategy) *UpdateEventStreamingRequestRunOptions {
	s.RetryStrategy = v
	return s
}

func (s *UpdateEventStreamingRequestRunOptions) SetThrottling(v int32) *UpdateEventStreamingRequestRunOptions {
	s.Throttling = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptions) Validate() error {
	if s.BatchWindow != nil {
		if err := s.BatchWindow.Validate(); err != nil {
			return err
		}
	}
	if s.BusinessOption != nil {
		if err := s.BusinessOption.Validate(); err != nil {
			return err
		}
	}
	if s.DeadLetterQueue != nil {
		if err := s.DeadLetterQueue.Validate(); err != nil {
			return err
		}
	}
	if s.RetryStrategy != nil {
		if err := s.RetryStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestRunOptionsBatchWindow struct {
	// The maximum number of events that can be contained in the window. When this threshold is reached, the data in the window is pushed downstream. If multiple windows exist, a push is triggered when any window meets the threshold.
	//
	// example:
	//
	// 100
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	// The maximum time range (in seconds) for events in the window. When this threshold is reached, the data in the window is pushed downstream. If multiple windows exist, a push is triggered when any window meets the threshold.
	//
	// example:
	//
	// 10
	TimeBasedWindow *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
}

func (s UpdateEventStreamingRequestRunOptionsBatchWindow) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestRunOptionsBatchWindow) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestRunOptionsBatchWindow) GetCountBasedWindow() *int32 {
	return s.CountBasedWindow
}

func (s *UpdateEventStreamingRequestRunOptionsBatchWindow) GetTimeBasedWindow() *int32 {
	return s.TimeBasedWindow
}

func (s *UpdateEventStreamingRequestRunOptionsBatchWindow) SetCountBasedWindow(v int32) *UpdateEventStreamingRequestRunOptionsBatchWindow {
	s.CountBasedWindow = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsBatchWindow) SetTimeBasedWindow(v int32) *UpdateEventStreamingRequestRunOptionsBatchWindow {
	s.TimeBasedWindow = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsBatchWindow) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestRunOptionsBusinessOption struct {
	BusinessMode         *string `json:"BusinessMode,omitempty" xml:"BusinessMode,omitempty"`
	MaxCapacityUnitCount *int64  `json:"MaxCapacityUnitCount,omitempty" xml:"MaxCapacityUnitCount,omitempty"`
	MinCapacityUnitCount *int64  `json:"MinCapacityUnitCount,omitempty" xml:"MinCapacityUnitCount,omitempty"`
}

func (s UpdateEventStreamingRequestRunOptionsBusinessOption) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestRunOptionsBusinessOption) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestRunOptionsBusinessOption) GetBusinessMode() *string {
	return s.BusinessMode
}

func (s *UpdateEventStreamingRequestRunOptionsBusinessOption) GetMaxCapacityUnitCount() *int64 {
	return s.MaxCapacityUnitCount
}

func (s *UpdateEventStreamingRequestRunOptionsBusinessOption) GetMinCapacityUnitCount() *int64 {
	return s.MinCapacityUnitCount
}

func (s *UpdateEventStreamingRequestRunOptionsBusinessOption) SetBusinessMode(v string) *UpdateEventStreamingRequestRunOptionsBusinessOption {
	s.BusinessMode = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsBusinessOption) SetMaxCapacityUnitCount(v int64) *UpdateEventStreamingRequestRunOptionsBusinessOption {
	s.MaxCapacityUnitCount = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsBusinessOption) SetMinCapacityUnitCount(v int64) *UpdateEventStreamingRequestRunOptionsBusinessOption {
	s.MinCapacityUnitCount = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsBusinessOption) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestRunOptionsDeadLetterQueue struct {
	// The ARN of the dead-letter queue.
	//
	// example:
	//
	// acs:ram::131733464781****:role/rdstoecsassumekms
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The network type of the dead-letter queue. Valid values:
	//
	// - PrivateNetwork
	//
	// - PublicNetwork
	//
	// example:
	//
	// PrivateNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The security group ID of the dead-letter queue instance.
	//
	// example:
	//
	// sg-2vcgdxz7o1n9zapp****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID of the dead-letter queue.
	//
	// example:
	//
	// vsw-m5ev8asdc6h12345****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The VPC ID of the dead-letter queue.
	//
	// example:
	//
	// vpc-2zehizpoendb3****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventStreamingRequestRunOptionsDeadLetterQueue) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestRunOptionsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) GetArn() *string {
	return s.Arn
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) GetNetwork() *string {
	return s.Network
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) SetArn(v string) *UpdateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.Arn = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) SetNetwork(v string) *UpdateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.Network = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) SetSecurityGroupId(v string) *UpdateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) SetVSwitchIds(v string) *UpdateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.VSwitchIds = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) SetVpcId(v string) *UpdateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.VpcId = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestRunOptionsRetryStrategy struct {
	// The maximum retry time.
	//
	// example:
	//
	// 512
	MaximumEventAgeInSeconds *int64 `json:"MaximumEventAgeInSeconds,omitempty" xml:"MaximumEventAgeInSeconds,omitempty"`
	// The maximum number of retry attempts.
	//
	// example:
	//
	// 2
	MaximumRetryAttempts *int64 `json:"MaximumRetryAttempts,omitempty" xml:"MaximumRetryAttempts,omitempty"`
	// The retry strategy: BACKOFF_RETRY (backoff retry) or EXPONENTIAL_DECAY_RETRY (exponential decay retry).
	//
	// example:
	//
	// BACKOFF_RETRY
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
}

func (s UpdateEventStreamingRequestRunOptionsRetryStrategy) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestRunOptionsRetryStrategy) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestRunOptionsRetryStrategy) GetMaximumEventAgeInSeconds() *int64 {
	return s.MaximumEventAgeInSeconds
}

func (s *UpdateEventStreamingRequestRunOptionsRetryStrategy) GetMaximumRetryAttempts() *int64 {
	return s.MaximumRetryAttempts
}

func (s *UpdateEventStreamingRequestRunOptionsRetryStrategy) GetPushRetryStrategy() *string {
	return s.PushRetryStrategy
}

func (s *UpdateEventStreamingRequestRunOptionsRetryStrategy) SetMaximumEventAgeInSeconds(v int64) *UpdateEventStreamingRequestRunOptionsRetryStrategy {
	s.MaximumEventAgeInSeconds = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsRetryStrategy) SetMaximumRetryAttempts(v int64) *UpdateEventStreamingRequestRunOptionsRetryStrategy {
	s.MaximumRetryAttempts = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsRetryStrategy) SetPushRetryStrategy(v string) *UpdateEventStreamingRequestRunOptionsRetryStrategy {
	s.PushRetryStrategy = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsRetryStrategy) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSink struct {
	SinkAgentRunParameters *SinkAgentRunParameters `json:"SinkAgentRunParameters,omitempty" xml:"SinkAgentRunParameters,omitempty"`
	// The description.
	SinkApacheKafkaParameters *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters `json:"SinkApacheKafkaParameters,omitempty" xml:"SinkApacheKafkaParameters,omitempty" type:"Struct"`
	// Sink Apache RocketMQ Checkpoint Parameters
	SinkApacheRocketMQCheckpointParameters *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters `json:"SinkApacheRocketMQCheckpointParameters,omitempty" xml:"SinkApacheRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The ApiDestination target parameters.
	SinkApiDestinationParameters *SinkApiDestinationParameters `json:"SinkApiDestinationParameters,omitempty" xml:"SinkApiDestinationParameters,omitempty"`
	// Sink BaiLian Parameters
	SinkBaiLianParameters *SinkBaiLianParameters `json:"SinkBaiLianParameters,omitempty" xml:"SinkBaiLianParameters,omitempty"`
	// The Sink Kafka connector parameters.
	SinkCustomizedKafkaConnectorParameters *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters `json:"SinkCustomizedKafkaConnectorParameters,omitempty" xml:"SinkCustomizedKafkaConnectorParameters,omitempty" type:"Struct"`
	// The Sink Kafka parameters.
	SinkCustomizedKafkaParameters *UpdateEventStreamingRequestSinkSinkCustomizedKafkaParameters `json:"SinkCustomizedKafkaParameters,omitempty" xml:"SinkCustomizedKafkaParameters,omitempty" type:"Struct"`
	// The Sink DashVector parameters.
	SinkDashVectorParameters *UpdateEventStreamingRequestSinkSinkDashVectorParameters `json:"SinkDashVectorParameters,omitempty" xml:"SinkDashVectorParameters,omitempty" type:"Struct"`
	// The Sink DataHub parameters.
	SinkDataHubParameters          *UpdateEventStreamingRequestSinkSinkDataHubParameters `json:"SinkDataHubParameters,omitempty" xml:"SinkDataHubParameters,omitempty" type:"Struct"`
	SinkDataWorksTriggerParameters *SinkDataWorksTriggerParameters                       `json:"SinkDataWorksTriggerParameters,omitempty" xml:"SinkDataWorksTriggerParameters,omitempty"`
	// The event source type.
	SinkDorisParameters *UpdateEventStreamingRequestSinkSinkDorisParameters `json:"SinkDorisParameters,omitempty" xml:"SinkDorisParameters,omitempty" type:"Struct"`
	// The event target name.
	SinkEventHouseParameters *UpdateEventStreamingRequestSinkSinkEventHouseParameters `json:"SinkEventHouseParameters,omitempty" xml:"SinkEventHouseParameters,omitempty" type:"Struct"`
	// The function target.
	SinkFcParameters *UpdateEventStreamingRequestSinkSinkFcParameters `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	// The Sink Fnf parameters.
	SinkFnfParameters *UpdateEventStreamingRequestSinkSinkFnfParameters `json:"SinkFnfParameters,omitempty" xml:"SinkFnfParameters,omitempty" type:"Struct"`
	// The HTTPS target parameters.
	SinkHttpsParameters *SinkHttpsParameters `json:"SinkHttpsParameters,omitempty" xml:"SinkHttpsParameters,omitempty"`
	// The Sink Kafka parameters.
	SinkKafkaParameters *UpdateEventStreamingRequestSinkSinkKafkaParameters `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	// The Simple Message Queue (formerly MNS) event target.
	SinkMNSParameters  *UpdateEventStreamingRequestSinkSinkMNSParameters `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
	SinkMQTTParameters *SinkMQTTParameters                               `json:"SinkMQTTParameters,omitempty" xml:"SinkMQTTParameters,omitempty"`
	SinkOSSParameters  *SinkOSSParameters                                `json:"SinkOSSParameters,omitempty" xml:"SinkOSSParameters,omitempty"`
	// Sink Open Source RabbitMQ Parameters
	SinkOpenSourceRabbitMQParameters *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters `json:"SinkOpenSourceRabbitMQParameters,omitempty" xml:"SinkOpenSourceRabbitMQParameters,omitempty" type:"Struct"`
	// The Sink Prometheus parameters.
	SinkPrometheusParameters      *UpdateEventStreamingRequestSinkSinkPrometheusParameters `json:"SinkPrometheusParameters,omitempty" xml:"SinkPrometheusParameters,omitempty" type:"Struct"`
	SinkRabbitMQMetaParameters    *SinkRabbitMQMetaParameters                              `json:"SinkRabbitMQMetaParameters,omitempty" xml:"SinkRabbitMQMetaParameters,omitempty"`
	SinkRabbitMQMsgSyncParameters *SinkRabbitMQMsgSyncParameters                           `json:"SinkRabbitMQMsgSyncParameters,omitempty" xml:"SinkRabbitMQMsgSyncParameters,omitempty"`
	// The Sink RabbitMQ parameters.
	SinkRabbitMQParameters *UpdateEventStreamingRequestSinkSinkRabbitMQParameters `json:"SinkRabbitMQParameters,omitempty" xml:"SinkRabbitMQParameters,omitempty" type:"Struct"`
	// Sink RocketMQ Checkpoint Parameters
	SinkRocketMQCheckpointParameters *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters `json:"SinkRocketMQCheckpointParameters,omitempty" xml:"SinkRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// Sink RocketMQ Parameters
	SinkRocketMQParameters *UpdateEventStreamingRequestSinkSinkRocketMQParameters `json:"SinkRocketMQParameters,omitempty" xml:"SinkRocketMQParameters,omitempty" type:"Struct"`
	// Sink SLS Parameters
	SinkSLSParameters *UpdateEventStreamingRequestSinkSinkSLSParameters `json:"SinkSLSParameters,omitempty" xml:"SinkSLSParameters,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSink) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSink) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSink) GetSinkAgentRunParameters() *SinkAgentRunParameters {
	return s.SinkAgentRunParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkApacheKafkaParameters() *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	return s.SinkApacheKafkaParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkApacheRocketMQCheckpointParameters() *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	return s.SinkApacheRocketMQCheckpointParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkApiDestinationParameters() *SinkApiDestinationParameters {
	return s.SinkApiDestinationParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkBaiLianParameters() *SinkBaiLianParameters {
	return s.SinkBaiLianParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkCustomizedKafkaConnectorParameters() *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters {
	return s.SinkCustomizedKafkaConnectorParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkCustomizedKafkaParameters() *UpdateEventStreamingRequestSinkSinkCustomizedKafkaParameters {
	return s.SinkCustomizedKafkaParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkDashVectorParameters() *UpdateEventStreamingRequestSinkSinkDashVectorParameters {
	return s.SinkDashVectorParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkDataHubParameters() *UpdateEventStreamingRequestSinkSinkDataHubParameters {
	return s.SinkDataHubParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkDataWorksTriggerParameters() *SinkDataWorksTriggerParameters {
	return s.SinkDataWorksTriggerParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkDorisParameters() *UpdateEventStreamingRequestSinkSinkDorisParameters {
	return s.SinkDorisParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkEventHouseParameters() *UpdateEventStreamingRequestSinkSinkEventHouseParameters {
	return s.SinkEventHouseParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkFcParameters() *UpdateEventStreamingRequestSinkSinkFcParameters {
	return s.SinkFcParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkFnfParameters() *UpdateEventStreamingRequestSinkSinkFnfParameters {
	return s.SinkFnfParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkHttpsParameters() *SinkHttpsParameters {
	return s.SinkHttpsParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkKafkaParameters() *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	return s.SinkKafkaParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkMNSParameters() *UpdateEventStreamingRequestSinkSinkMNSParameters {
	return s.SinkMNSParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkMQTTParameters() *SinkMQTTParameters {
	return s.SinkMQTTParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkOSSParameters() *SinkOSSParameters {
	return s.SinkOSSParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkOpenSourceRabbitMQParameters() *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	return s.SinkOpenSourceRabbitMQParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkPrometheusParameters() *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	return s.SinkPrometheusParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkRabbitMQMetaParameters() *SinkRabbitMQMetaParameters {
	return s.SinkRabbitMQMetaParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkRabbitMQMsgSyncParameters() *SinkRabbitMQMsgSyncParameters {
	return s.SinkRabbitMQMsgSyncParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkRabbitMQParameters() *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	return s.SinkRabbitMQParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkRocketMQCheckpointParameters() *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	return s.SinkRocketMQCheckpointParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkRocketMQParameters() *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	return s.SinkRocketMQParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkSLSParameters() *UpdateEventStreamingRequestSinkSinkSLSParameters {
	return s.SinkSLSParameters
}

func (s *UpdateEventStreamingRequestSink) SetSinkAgentRunParameters(v *SinkAgentRunParameters) *UpdateEventStreamingRequestSink {
	s.SinkAgentRunParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkApacheKafkaParameters(v *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) *UpdateEventStreamingRequestSink {
	s.SinkApacheKafkaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkApacheRocketMQCheckpointParameters(v *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) *UpdateEventStreamingRequestSink {
	s.SinkApacheRocketMQCheckpointParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkApiDestinationParameters(v *SinkApiDestinationParameters) *UpdateEventStreamingRequestSink {
	s.SinkApiDestinationParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkBaiLianParameters(v *SinkBaiLianParameters) *UpdateEventStreamingRequestSink {
	s.SinkBaiLianParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkCustomizedKafkaConnectorParameters(v *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) *UpdateEventStreamingRequestSink {
	s.SinkCustomizedKafkaConnectorParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkCustomizedKafkaParameters(v *UpdateEventStreamingRequestSinkSinkCustomizedKafkaParameters) *UpdateEventStreamingRequestSink {
	s.SinkCustomizedKafkaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkDashVectorParameters(v *UpdateEventStreamingRequestSinkSinkDashVectorParameters) *UpdateEventStreamingRequestSink {
	s.SinkDashVectorParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkDataHubParameters(v *UpdateEventStreamingRequestSinkSinkDataHubParameters) *UpdateEventStreamingRequestSink {
	s.SinkDataHubParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkDataWorksTriggerParameters(v *SinkDataWorksTriggerParameters) *UpdateEventStreamingRequestSink {
	s.SinkDataWorksTriggerParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkDorisParameters(v *UpdateEventStreamingRequestSinkSinkDorisParameters) *UpdateEventStreamingRequestSink {
	s.SinkDorisParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkEventHouseParameters(v *UpdateEventStreamingRequestSinkSinkEventHouseParameters) *UpdateEventStreamingRequestSink {
	s.SinkEventHouseParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkFcParameters(v *UpdateEventStreamingRequestSinkSinkFcParameters) *UpdateEventStreamingRequestSink {
	s.SinkFcParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkFnfParameters(v *UpdateEventStreamingRequestSinkSinkFnfParameters) *UpdateEventStreamingRequestSink {
	s.SinkFnfParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkHttpsParameters(v *SinkHttpsParameters) *UpdateEventStreamingRequestSink {
	s.SinkHttpsParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkKafkaParameters(v *UpdateEventStreamingRequestSinkSinkKafkaParameters) *UpdateEventStreamingRequestSink {
	s.SinkKafkaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkMNSParameters(v *UpdateEventStreamingRequestSinkSinkMNSParameters) *UpdateEventStreamingRequestSink {
	s.SinkMNSParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkMQTTParameters(v *SinkMQTTParameters) *UpdateEventStreamingRequestSink {
	s.SinkMQTTParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkOSSParameters(v *SinkOSSParameters) *UpdateEventStreamingRequestSink {
	s.SinkOSSParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkOpenSourceRabbitMQParameters(v *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) *UpdateEventStreamingRequestSink {
	s.SinkOpenSourceRabbitMQParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkPrometheusParameters(v *UpdateEventStreamingRequestSinkSinkPrometheusParameters) *UpdateEventStreamingRequestSink {
	s.SinkPrometheusParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkRabbitMQMetaParameters(v *SinkRabbitMQMetaParameters) *UpdateEventStreamingRequestSink {
	s.SinkRabbitMQMetaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkRabbitMQMsgSyncParameters(v *SinkRabbitMQMsgSyncParameters) *UpdateEventStreamingRequestSink {
	s.SinkRabbitMQMsgSyncParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkRabbitMQParameters(v *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) *UpdateEventStreamingRequestSink {
	s.SinkRabbitMQParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkRocketMQCheckpointParameters(v *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) *UpdateEventStreamingRequestSink {
	s.SinkRocketMQCheckpointParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkRocketMQParameters(v *UpdateEventStreamingRequestSinkSinkRocketMQParameters) *UpdateEventStreamingRequestSink {
	s.SinkRocketMQParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkSLSParameters(v *UpdateEventStreamingRequestSinkSinkSLSParameters) *UpdateEventStreamingRequestSink {
	s.SinkSLSParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) Validate() error {
	if s.SinkAgentRunParameters != nil {
		if err := s.SinkAgentRunParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkApacheKafkaParameters != nil {
		if err := s.SinkApacheKafkaParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkApacheRocketMQCheckpointParameters != nil {
		if err := s.SinkApacheRocketMQCheckpointParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkApiDestinationParameters != nil {
		if err := s.SinkApiDestinationParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkBaiLianParameters != nil {
		if err := s.SinkBaiLianParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkCustomizedKafkaConnectorParameters != nil {
		if err := s.SinkCustomizedKafkaConnectorParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkCustomizedKafkaParameters != nil {
		if err := s.SinkCustomizedKafkaParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkDashVectorParameters != nil {
		if err := s.SinkDashVectorParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkDataHubParameters != nil {
		if err := s.SinkDataHubParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkDataWorksTriggerParameters != nil {
		if err := s.SinkDataWorksTriggerParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkDorisParameters != nil {
		if err := s.SinkDorisParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkEventHouseParameters != nil {
		if err := s.SinkEventHouseParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkFcParameters != nil {
		if err := s.SinkFcParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkFnfParameters != nil {
		if err := s.SinkFnfParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkHttpsParameters != nil {
		if err := s.SinkHttpsParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkKafkaParameters != nil {
		if err := s.SinkKafkaParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkMNSParameters != nil {
		if err := s.SinkMNSParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkMQTTParameters != nil {
		if err := s.SinkMQTTParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkOSSParameters != nil {
		if err := s.SinkOSSParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkOpenSourceRabbitMQParameters != nil {
		if err := s.SinkOpenSourceRabbitMQParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkPrometheusParameters != nil {
		if err := s.SinkPrometheusParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkRabbitMQMetaParameters != nil {
		if err := s.SinkRabbitMQMetaParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkRabbitMQMsgSyncParameters != nil {
		if err := s.SinkRabbitMQMsgSyncParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkRabbitMQParameters != nil {
		if err := s.SinkRabbitMQParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkRocketMQCheckpointParameters != nil {
		if err := s.SinkRocketMQCheckpointParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkRocketMQParameters != nil {
		if err := s.SinkRocketMQParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SinkSLSParameters != nil {
		if err := s.SinkSLSParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkApacheKafkaParameters struct {
	Acks            *string `json:"Acks,omitempty" xml:"Acks,omitempty"`
	Bootstraps      *string `json:"Bootstraps,omitempty" xml:"Bootstraps,omitempty"`
	CompressionType *string `json:"CompressionType,omitempty" xml:"CompressionType,omitempty"`
	// Specifies the target Topic routing strategy for messages. If both the Topic parameter and the DynamicTopic parameter are specified, the DynamicTopic parameter takes precedence. Two configuration modes are supported:
	//
	//     1. **Static constant mode**: Specify a fixed Topic name string (for example, "order_created"). All messages are sent to this Topic.
	//
	//     2. **Dynamic extraction mode**: Specify a standard JSONPath expression (for example, "$.user.id" or "$.metadata.category"). The system parses the upstream message body and extracts the matching field value as the target Topic name.
	DynamicTopic     *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic    `json:"DynamicTopic,omitempty" xml:"DynamicTopic,omitempty" type:"Struct"`
	Headers          *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders         `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	Key              *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey             `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	NetworkType      *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType     `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	SaslMechanism    *string                                                                  `json:"SaslMechanism,omitempty" xml:"SaslMechanism,omitempty"`
	SaslPassword     *string                                                                  `json:"SaslPassword,omitempty" xml:"SaslPassword,omitempty"`
	SaslUser         *string                                                                  `json:"SaslUser,omitempty" xml:"SaslUser,omitempty"`
	SecurityGroupId  *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	SecurityProtocol *string                                                                  `json:"SecurityProtocol,omitempty" xml:"SecurityProtocol,omitempty"`
	// [Required for encrypted private key] The Kafka client private key password. This parameter is required when the client private key is password-protected (the PEM file contains \\"Proc-Type: 4,ENCRYPTED\\" or \\"ENCRYPTED\\" markers). Leave empty if the private key is not encrypted. Note: This password is only used to decrypt the private key and is unrelated to Kafka authentication.
	SslKeyPassword *string `json:"SslKeyPassword,omitempty" xml:"SslKeyPassword,omitempty"`
	// [Required for mutual authentication] The Kafka client certificate chain. This parameter is required when the Kafka server enables mutual SSL authentication (ssl.client.auth=required). Format: Base64-encoded PEM format containing the client certificate and the complete certificate chain (client certificate first, intermediate CA certificate next, root CA certificate optional). Note: Ensure each PEM file content starts with \\"-----BEGIN CERTIFICATE-----\\" and ends with \\"-----END CERTIFICATE-----\\", then Base64-encode the concatenated content.
	SslKeystoreCertificateChain *string `json:"SslKeystoreCertificateChain,omitempty" xml:"SslKeystoreCertificateChain,omitempty"`
	// [Required for mutual authentication] The SSL private key configuration object. This parameter is required when the Kafka server enables mutual SSL authentication. Only KMS mode is supported: specify the Key Management Service resource that stores the private key through KmsArn. The system retrieves the private key content from KMS only in memory for higher security. Configuration example: {\\"KmsArn\\": \\"acs:kms:cn-hangzhou:123456789:secret/ssl-key-xxxx\\", \\"KmsSecretValueKey\\": \\"keystore_private_key\\"}
	SslKeystoreKey *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey `json:"SslKeystoreKey,omitempty" xml:"SslKeystoreKey,omitempty" type:"Struct"`
	// [Required for SSL] The Kafka server trust certificate. Used to authenticate the SSL certificate of the Kafka Broker to prevent man-in-the-middle attacks. Format: Base64 encoding of PEM format, typically containing the CA certificate or the server certificate of the Kafka server. Example: Base64-encode the PEM file content of the CA certificate (ensure it starts with \\"-----BEGIN CERTIFICATE-----\\" and ends with \\"-----END CERTIFICATE-----\\"). If Kafka uses a self-signed certificate, provide the CA certificate that issued the certificate.
	SslTruststoreCertificates *string                                                             `json:"SslTruststoreCertificates,omitempty" xml:"SslTruststoreCertificates,omitempty"`
	Topic                     *string                                                             `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds                *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	Value                     *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue      `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
	VpcId                     *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId      `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetAcks() *string {
	return s.Acks
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetBootstraps() *string {
	return s.Bootstraps
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetCompressionType() *string {
	return s.CompressionType
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetDynamicTopic() *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic {
	return s.DynamicTopic
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetHeaders() *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders {
	return s.Headers
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetKey() *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey {
	return s.Key
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetNetworkType() *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType {
	return s.NetworkType
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSaslMechanism() *string {
	return s.SaslMechanism
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSaslPassword() *string {
	return s.SaslPassword
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSaslUser() *string {
	return s.SaslUser
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSecurityGroupId() *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSecurityProtocol() *string {
	return s.SecurityProtocol
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSslKeyPassword() *string {
	return s.SslKeyPassword
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSslKeystoreCertificateChain() *string {
	return s.SslKeystoreCertificateChain
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSslKeystoreKey() *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey {
	return s.SslKeystoreKey
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSslTruststoreCertificates() *string {
	return s.SslTruststoreCertificates
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetVSwitchIds() *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetValue() *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) GetVpcId() *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetAcks(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Acks = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetBootstraps(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Bootstraps = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetCompressionType(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.CompressionType = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetDynamicTopic(v *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.DynamicTopic = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetHeaders(v *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Headers = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetKey(v *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Key = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetNetworkType(v *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.NetworkType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSaslMechanism(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SaslMechanism = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSaslPassword(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SaslPassword = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSaslUser(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SaslUser = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSecurityGroupId(v *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SecurityGroupId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSecurityProtocol(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SecurityProtocol = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSslKeyPassword(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SslKeyPassword = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSslKeystoreCertificateChain(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SslKeystoreCertificateChain = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSslKeystoreKey(v *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SslKeystoreKey = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSslTruststoreCertificates(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SslTruststoreCertificates = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetTopic(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetVSwitchIds(v *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.VSwitchIds = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetValue(v *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Value = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) SetVpcId(v *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.VpcId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParameters) Validate() error {
	if s.DynamicTopic != nil {
		if err := s.DynamicTopic.Validate(); err != nil {
			return err
		}
	}
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	if s.Key != nil {
		if err := s.Key.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkType != nil {
		if err := s.NetworkType.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroupId != nil {
		if err := s.SecurityGroupId.Validate(); err != nil {
			return err
		}
	}
	if s.SslKeystoreKey != nil {
		if err := s.SslKeystoreKey.Validate(); err != nil {
			return err
		}
	}
	if s.VSwitchIds != nil {
		if err := s.VSwitchIds.Validate(); err != nil {
			return err
		}
	}
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	if s.VpcId != nil {
		if err := s.VpcId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic struct {
	// The transformation type. Valid values:
	//
	// - CONSTANT: fixed value
	//
	// - JSONPATH: extracted from upstream based on path
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic) SetForm(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic) SetValue(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersDynamicTopic) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) SetForm(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) SetValue(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey) SetForm(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey) SetValue(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersKey) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey struct {
	// [Required] The KMS resource ARN that stores the SSL private key. Used to locate the Key Management Service instance that stores the client private key. Format example: \\"acs:kms:cn-hangzhou:123456789:secret/ssl-keystore-key-xxxx\\". Obtain this value from the ARN information of the corresponding key in the KMS console.
	KmsArn *string `json:"KmsArn,omitempty" xml:"KmsArn,omitempty"`
	// [KMS KV mode] The key name in the KMS credential. When the KMS credential is stored as a key-value (KV) structure, specify this parameter to indicate the key corresponding to the SSL private key. Example: if the KMS credential is \\"{"ssl_keystore_key":"-----BEGIN PRIVATE KEY-----...","ssl_truststore_key":"..."}\\", enter \\"ssl_keystore_key\\". Leave empty if the KMS credential is in plain text mode (directly stores the PEM content of the private key).
	KmsSecretValueKey *string `json:"KmsSecretValueKey,omitempty" xml:"KmsSecretValueKey,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey) GetKmsArn() *string {
	return s.KmsArn
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey) GetKmsSecretValueKey() *string {
	return s.KmsSecretValueKey
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey) SetKmsArn(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey {
	s.KmsArn = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey) SetKmsSecretValueKey(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey {
	s.KmsSecretValueKey = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersSslKeystoreKey) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) SetForm(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) SetValue(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue) SetForm(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue) SetValue(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersValue) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters struct {
	// The timestamp of message consumption.
	ConsumeTimestamp *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp `json:"ConsumeTimestamp,omitempty" xml:"ConsumeTimestamp,omitempty" type:"Struct"`
	// The Group ID of the consumer group.
	Group *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The instance endpoint.
	//
	// example:
	//
	// 192.168.1.1:9876
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The password of the username.
	//
	// example:
	//
	// ****
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The username required for authentication.
	//
	// example:
	//
	// admin
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The network type. Valid values:
	//
	// - PublicNetwork
	//
	// - PrivateNetwork
	//
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-2ze5bmpw6adn0q******
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The topic of the RocketMQ instance.
	Topic *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf62oqt1twuevrt******
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-2zeccak5pb0j3ay******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetConsumeTimestamp() *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	return s.ConsumeTimestamp
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetGroup() *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup {
	return s.Group
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetTopic() *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetConsumeTimestamp(v *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.ConsumeTimestamp = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetGroup(v *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.Group = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetInstanceEndpoint(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetInstancePassword(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.InstancePassword = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetInstanceUsername(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.InstanceUsername = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetNetworkType(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.NetworkType = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetSecurityGroupId(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetTopic(v *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.Topic = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetVSwitchId(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.VSwitchId = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetVpcId(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.VpcId = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) Validate() error {
	if s.ConsumeTimestamp != nil {
		if err := s.ConsumeTimestamp.Validate(); err != nil {
			return err
		}
	}
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	if s.Topic != nil {
		if err := s.Topic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 1570761026400
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetForm(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetValue(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Group ID
	//
	// example:
	//
	// GID_EVENTBRIDGE_1736234******
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) SetForm(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) SetValue(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic name of the RocketMQ instance.
	//
	// example:
	//
	// Mytopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) SetForm(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) SetValue(v string) *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters struct {
	// The OSS file download URL.
	//
	// example:
	//
	// "https://examplebucket.oss-cn-hangzhou.aliyuncs.com/testDoc/Old_Homebrew/2024-06-26%2022%3A34%3A08/opt/homebrew/homebrew/Library/Homebrew/test/support/fixtures/cask/AppWithBinary.zip?OSSAccessKeyId=ri&Expires=1725539627&Signature=rb8q3OpV2i3gZJ"
	ConnectorPackageUrl *string `json:"ConnectorPackageUrl,omitempty" xml:"ConnectorPackageUrl,omitempty"`
	// Parses the properties file in the current ZIP package.
	ConnectorParameters *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters `json:"ConnectorParameters,omitempty" xml:"ConnectorParameters,omitempty" type:"Struct"`
	// The instance configuration.
	//
	// example:
	//
	// {
	//
	//         "group.id": "connect-eb-cluster-KAFKA_CONNECTORC",
	//
	//         "offset.storage.topic": "connect-eb-offset-KAFKA_CONNECTOR_yjqC8K5ewC",
	//
	//         "config.storage.topic": "connect-eb-config-KAFKA_CONNECTOR_yjqC8K5ewC",
	//
	//         "status.storage.topic": "connect-eb-status-KAFKA_CONNECTOR_yjqC8K5ewC",
	//
	//         "consumer.group.id": "connector-eb-cluster-KAFKA_CONNECTOR_yjqC8K5ewC-mongo-sink",
	//
	//         "bootstrap.servers": "alikafka-post:9092"
	//
	//       }
	WorkerParameters map[string]interface{} `json:"WorkerParameters,omitempty" xml:"WorkerParameters,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) GetConnectorPackageUrl() *string {
	return s.ConnectorPackageUrl
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) GetConnectorParameters() *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	return s.ConnectorParameters
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) GetWorkerParameters() map[string]interface{} {
	return s.WorkerParameters
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) SetConnectorPackageUrl(v string) *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters {
	s.ConnectorPackageUrl = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) SetConnectorParameters(v *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters {
	s.ConnectorParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) SetWorkerParameters(v map[string]interface{}) *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters {
	s.WorkerParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) Validate() error {
	if s.ConnectorParameters != nil {
		if err := s.ConnectorParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters struct {
	// The connector configuration.
	//
	// example:
	//
	// {
	//
	//           "connector.class": "com.mongodb.kafka.connect.MongoSinkConnector",
	//
	//           "tasks.max": "1",
	//
	//           "topics": "sourceA,sourceB"
	//
	//         }
	Config map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	// The connector name.
	//
	// example:
	//
	// mongo-sink
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GetName() *string {
	return s.Name
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) SetConfig(v map[string]interface{}) *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	s.Config = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) SetName(v string) *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	s.Name = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkCustomizedKafkaParameters struct {
	// The instance ID of the ApsaraMQ for Kafka instance.
	//
	// example:
	//
	// 90be1f96-4229-4535-bb76-34b4f6fb2b71
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkCustomizedKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkCustomizedKafkaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSinkSinkCustomizedKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkCustomizedKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDashVectorParameters struct {
	// The API key created in the DashVector console.
	//
	// example:
	//
	// Q34nExQH7sQ****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The collection name.
	//
	// example:
	//
	// collection1
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The Schema field definition of the table entry when inserting into DashVector. The event content transformation result must be in JSON format.
	DashVectorSchemaParameters []*UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters `json:"DashVectorSchemaParameters,omitempty" xml:"DashVectorSchemaParameters,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// vrs-cn-lbj3ru1***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type. Valid values:
	//
	// - PrivateNetwork
	//
	// - PublicNetwork
	//
	// example:
	//
	// PublicNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The DashVector database operation type. Valid values:
	//
	// - Delete
	//
	// - Upsert
	//
	// example:
	//
	// Upsert
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The partition. Default value: default.
	Partition *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition `json:"Partition,omitempty" xml:"Partition,omitempty" type:"Struct"`
	// The primary key ID when inserting or deleting records. If this field is not specified, a random primary key ID is used.
	PrimaryKeyId *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId `json:"PrimaryKeyId,omitempty" xml:"PrimaryKeyId,omitempty" type:"Struct"`
	// The vector of the record inserted into DashVector.
	Vector *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) GetCollection() *string {
	return s.Collection
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) GetDashVectorSchemaParameters() []*UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
	return s.DashVectorSchemaParameters
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) GetNetwork() *string {
	return s.Network
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) GetOperation() *string {
	return s.Operation
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) GetPartition() *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition {
	return s.Partition
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) GetPrimaryKeyId() *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId {
	return s.PrimaryKeyId
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) GetVector() *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector {
	return s.Vector
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) SetApiKey(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParameters {
	s.ApiKey = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) SetCollection(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParameters {
	s.Collection = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) SetDashVectorSchemaParameters(v []*UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) *UpdateEventStreamingRequestSinkSinkDashVectorParameters {
	s.DashVectorSchemaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) SetNetwork(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParameters {
	s.Network = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) SetOperation(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParameters {
	s.Operation = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) SetPartition(v *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition) *UpdateEventStreamingRequestSinkSinkDashVectorParameters {
	s.Partition = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) SetPrimaryKeyId(v *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) *UpdateEventStreamingRequestSinkSinkDashVectorParameters {
	s.PrimaryKeyId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) SetVector(v *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector) *UpdateEventStreamingRequestSinkSinkDashVectorParameters {
	s.Vector = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) Validate() error {
	if s.DashVectorSchemaParameters != nil {
		for _, item := range s.DashVectorSchemaParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Partition != nil {
		if err := s.Partition.Validate(); err != nil {
			return err
		}
	}
	if s.PrimaryKeyId != nil {
		if err := s.PrimaryKeyId.Validate(); err != nil {
			return err
		}
	}
	if s.Vector != nil {
		if err := s.Vector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters struct {
	Name  *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName  `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	Type  *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType  `json:"Type,omitempty" xml:"Type,omitempty" type:"Struct"`
	Value *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GetName() *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	return s.Name
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GetType() *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	return s.Type
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GetValue() *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) SetName(v *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Name = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) SetType(v *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Type = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) SetValue(v *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Value = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	if s.Type != nil {
		if err := s.Type.Validate(); err != nil {
			return err
		}
	}
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition struct {
	// The transformation format. Valid values:
	//
	// - JSONPATH
	//
	// - CONSTANT
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value. Valid values:
	//
	// - If Form is CONSTANT: a constant value
	//
	// - If Form is JSONPATH: JSONPath extraction content
	//
	// > The Value field cannot exceed 10240 characters.
	//
	// example:
	//
	// default
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId struct {
	// The transformation format. Valid values:
	//
	// - JSONPATH
	//
	// - TEMPLATE
	//
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The primary key ID template. This parameter is required only when Form is set to TEMPLATE.
	//
	// example:
	//
	// ${ID}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value. Valid values:
	//
	// - If Form is JSONPATH: JSONPath extraction content
	//
	// - If Form is TEMPLATE: a template variable
	//
	// > The Value field cannot exceed 10240 characters.
	//
	// example:
	//
	// $.data.requestId
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDashVectorParametersVector struct {
	// The transformation format.
	//
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The JSONPath extraction content.
	//
	// > The Value field cannot exceed 10240 characters.
	//
	// example:
	//
	// $.data.messageBody
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersVector) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersVector) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersVector) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDataHubParameters struct {
	// The BLOB type Record content template.
	Body *UpdateEventStreamingRequestSinkSinkDataHubParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The custom log key-value pairs. This parameter takes effect only when ContentType is set to KeyValue. Each key-value pair is represented by Key_n and Value_n.
	ContentSchema *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema `json:"ContentSchema,omitempty" xml:"ContentSchema,omitempty" type:"Struct"`
	// The data format. You can select the default format or configure specified key-value pairs. Valid values:
	//
	// - JSON
	//
	// - KeyValue
	ContentType *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType `json:"ContentType,omitempty" xml:"ContentType,omitempty" type:"Struct"`
	// The DataHub project name.
	Project *UpdateEventStreamingRequestSinkSinkDataHubParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The task role name.
	RoleName *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The DataHub topic name.
	Topic *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The TUPLE type topic content schema.
	TopicSchema *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema `json:"TopicSchema,omitempty" xml:"TopicSchema,omitempty" type:"Struct"`
	// The topic type. Valid values:
	//
	// - TUPLE
	//
	// - BLOB
	TopicType *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType `json:"TopicType,omitempty" xml:"TopicType,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) GetBody() *UpdateEventStreamingRequestSinkSinkDataHubParametersBody {
	return s.Body
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) GetContentSchema() *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema {
	return s.ContentSchema
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) GetContentType() *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType {
	return s.ContentType
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) GetProject() *UpdateEventStreamingRequestSinkSinkDataHubParametersProject {
	return s.Project
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) GetRoleName() *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName {
	return s.RoleName
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) GetTopic() *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) GetTopicSchema() *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema {
	return s.TopicSchema
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) GetTopicType() *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType {
	return s.TopicType
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkDataHubParametersBody) *UpdateEventStreamingRequestSinkSinkDataHubParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) SetContentSchema(v *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema) *UpdateEventStreamingRequestSinkSinkDataHubParameters {
	s.ContentSchema = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) SetContentType(v *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType) *UpdateEventStreamingRequestSinkSinkDataHubParameters {
	s.ContentType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) SetProject(v *UpdateEventStreamingRequestSinkSinkDataHubParametersProject) *UpdateEventStreamingRequestSinkSinkDataHubParameters {
	s.Project = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) SetRoleName(v *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName) *UpdateEventStreamingRequestSinkSinkDataHubParameters {
	s.RoleName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) SetTopic(v *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic) *UpdateEventStreamingRequestSinkSinkDataHubParameters {
	s.Topic = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) SetTopicSchema(v *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) *UpdateEventStreamingRequestSinkSinkDataHubParameters {
	s.TopicSchema = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) SetTopicType(v *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType) *UpdateEventStreamingRequestSinkSinkDataHubParameters {
	s.TopicType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParameters) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.ContentSchema != nil {
		if err := s.ContentSchema.Validate(); err != nil {
			return err
		}
	}
	if s.ContentType != nil {
		if err := s.ContentType.Validate(); err != nil {
			return err
		}
	}
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	if s.RoleName != nil {
		if err := s.RoleName.Validate(); err != nil {
			return err
		}
	}
	if s.Topic != nil {
		if err := s.Topic.Validate(); err != nil {
			return err
		}
	}
	if s.TopicSchema != nil {
		if err := s.TopicSchema.Validate(); err != nil {
			return err
		}
	}
	if s.TopicType != nil {
		if err := s.TopicType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkDataHubParametersBody struct {
	// The transformation format.
	//
	// example:
	//
	// ORIGINAL
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The BLOB type Record content template.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersBody) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersBody) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersBody {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema struct {
	// The template style.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {"Key_1":{"form":"CONSTANT","value":"demoKey"},"Value_1":{"form":"JSONPATH","value":"$.data.value"}}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDataHubParametersContentType struct {
	// The transformation format.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// JSON
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersContentType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersContentType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDataHubParametersProject struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The DataHub project name.
	//
	// example:
	//
	// demo-project
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersProject) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersProject) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersProject) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersProject) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersProject) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersProject) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersProject {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersProject) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersProject {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersProject) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersProject {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersProject) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The task role name.
	//
	// example:
	//
	// test-role
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDataHubParametersTopic struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The DataHub topic name.
	//
	// example:
	//
	// demo-topic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersTopic) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// {"k1":"${k1}","k2":"${k2}"}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The TUPLE type topic content schema.
	//
	// example:
	//
	// {"k1":"value1","k2":"value2"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic type. Valid values:
	//
	// - TUPLE
	//
	// - BLOB
	//
	// example:
	//
	// TUPLE
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParameters struct {
	BeHttpEndpoint  *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint  `json:"BeHttpEndpoint,omitempty" xml:"BeHttpEndpoint,omitempty" type:"Struct"`
	Body            *UpdateEventStreamingRequestSinkSinkDorisParametersBody            `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Database        *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase        `json:"Database,omitempty" xml:"Database,omitempty" type:"Struct"`
	FeHttpEndpoint  *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint  `json:"FeHttpEndpoint,omitempty" xml:"FeHttpEndpoint,omitempty" type:"Struct"`
	NetworkType     *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType     `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	Password        *UpdateEventStreamingRequestSinkSinkDorisParametersPassword        `json:"Password,omitempty" xml:"Password,omitempty" type:"Struct"`
	QueryEndpoint   *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint   `json:"QueryEndpoint,omitempty" xml:"QueryEndpoint,omitempty" type:"Struct"`
	SecurityGroupId *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	Table           *UpdateEventStreamingRequestSinkSinkDorisParametersTable           `json:"Table,omitempty" xml:"Table,omitempty" type:"Struct"`
	UserName        *UpdateEventStreamingRequestSinkSinkDorisParametersUserName        `json:"UserName,omitempty" xml:"UserName,omitempty" type:"Struct"`
	VSwitchIds      *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds      `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	VpcId           *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId           `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetBeHttpEndpoint() *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint {
	return s.BeHttpEndpoint
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetBody() *UpdateEventStreamingRequestSinkSinkDorisParametersBody {
	return s.Body
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetDatabase() *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase {
	return s.Database
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetFeHttpEndpoint() *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint {
	return s.FeHttpEndpoint
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetNetworkType() *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType {
	return s.NetworkType
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetPassword() *UpdateEventStreamingRequestSinkSinkDorisParametersPassword {
	return s.Password
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetQueryEndpoint() *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint {
	return s.QueryEndpoint
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetSecurityGroupId() *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetTable() *UpdateEventStreamingRequestSinkSinkDorisParametersTable {
	return s.Table
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetUserName() *UpdateEventStreamingRequestSinkSinkDorisParametersUserName {
	return s.UserName
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetVSwitchIds() *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) GetVpcId() *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetBeHttpEndpoint(v *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.BeHttpEndpoint = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkDorisParametersBody) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetDatabase(v *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.Database = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetFeHttpEndpoint(v *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.FeHttpEndpoint = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetNetworkType(v *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.NetworkType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetPassword(v *UpdateEventStreamingRequestSinkSinkDorisParametersPassword) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.Password = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetQueryEndpoint(v *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.QueryEndpoint = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetSecurityGroupId(v *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.SecurityGroupId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetTable(v *UpdateEventStreamingRequestSinkSinkDorisParametersTable) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.Table = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetUserName(v *UpdateEventStreamingRequestSinkSinkDorisParametersUserName) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.UserName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetVSwitchIds(v *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.VSwitchIds = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) SetVpcId(v *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId) *UpdateEventStreamingRequestSinkSinkDorisParameters {
	s.VpcId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParameters) Validate() error {
	if s.BeHttpEndpoint != nil {
		if err := s.BeHttpEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.Database != nil {
		if err := s.Database.Validate(); err != nil {
			return err
		}
	}
	if s.FeHttpEndpoint != nil {
		if err := s.FeHttpEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkType != nil {
		if err := s.NetworkType.Validate(); err != nil {
			return err
		}
	}
	if s.Password != nil {
		if err := s.Password.Validate(); err != nil {
			return err
		}
	}
	if s.QueryEndpoint != nil {
		if err := s.QueryEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroupId != nil {
		if err := s.SecurityGroupId.Validate(); err != nil {
			return err
		}
	}
	if s.Table != nil {
		if err := s.Table.Validate(); err != nil {
			return err
		}
	}
	if s.UserName != nil {
		if err := s.UserName.Validate(); err != nil {
			return err
		}
	}
	if s.VSwitchIds != nil {
		if err := s.VSwitchIds.Validate(); err != nil {
			return err
		}
	}
	if s.VpcId != nil {
		if err := s.VpcId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParametersBody struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBody) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBody) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersBody {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParametersDatabase struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersDatabase) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersDatabase) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersDatabase) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersNetworkType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParametersPassword struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersPassword) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersPassword) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersPassword) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersPassword) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersPassword) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersPassword) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersPassword {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersPassword) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersPassword {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersPassword) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersPassword {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersPassword) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParametersTable struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersTable) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersTable) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersTable) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersTable) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersTable) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersTable) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersTable {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersTable) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersTable {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersTable) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersTable {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersTable) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParametersUserName struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersUserName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersUserName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersUserName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersUserName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersUserName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersUserName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersUserName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersUserName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersUserName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersUserName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersUserName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersUserName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDorisParametersVpcId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDorisParametersVpcId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDorisParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkEventHouseParameters struct {
	// The catalog name.
	//
	// example:
	//
	// demo
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The name of the target table.
	//
	// example:
	//
	// demo-table
	EventTableName *string `json:"EventTableName,omitempty" xml:"EventTableName,omitempty"`
	// The field mapping rules.
	MappingRules []*UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules `json:"MappingRules,omitempty" xml:"MappingRules,omitempty" type:"Repeated"`
	// The namespace of the target table.
	//
	// example:
	//
	// name1
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkEventHouseParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkEventHouseParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParameters) GetCatalogName() *string {
	return s.CatalogName
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParameters) GetEventTableName() *string {
	return s.EventTableName
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParameters) GetMappingRules() []*UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules {
	return s.MappingRules
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParameters) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParameters) SetCatalogName(v string) *UpdateEventStreamingRequestSinkSinkEventHouseParameters {
	s.CatalogName = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParameters) SetEventTableName(v string) *UpdateEventStreamingRequestSinkSinkEventHouseParameters {
	s.EventTableName = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParameters) SetMappingRules(v []*UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules) *UpdateEventStreamingRequestSinkSinkEventHouseParameters {
	s.MappingRules = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParameters) SetNamespaceName(v string) *UpdateEventStreamingRequestSinkSinkEventHouseParameters {
	s.NamespaceName = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParameters) Validate() error {
	if s.MappingRules != nil {
		for _, item := range s.MappingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules struct {
	// The column name.
	//
	// example:
	//
	// age
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The column type.
	//
	// example:
	//
	// text
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// The column value extraction rule.
	ColumnValue *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue `json:"ColumnValue,omitempty" xml:"ColumnValue,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules) GetColumnName() *string {
	return s.ColumnName
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules) GetColumnType() *string {
	return s.ColumnType
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules) GetColumnValue() *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue {
	return s.ColumnValue
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules) SetColumnName(v string) *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules {
	s.ColumnName = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules) SetColumnType(v string) *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules {
	s.ColumnType = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules) SetColumnValue(v *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue) *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules {
	s.ColumnValue = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRules) Validate() error {
	if s.ColumnValue != nil {
		if err := s.ColumnValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue struct {
	// The transformation method, such as JSONPATH.
	//
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template configuration.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The extraction path, such as $.data.value.name.
	//
	// example:
	//
	// $.data.value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue) SetForm(v string) *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue) SetValue(v string) *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkEventHouseParametersMappingRulesColumnValue) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFcParameters struct {
	// The content body sent to the function.
	Body *UpdateEventStreamingRequestSinkSinkFcParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The delivery concurrency. Minimum value: 1.
	Concurrency *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	// The format conversion rule for event content.
	DataFormat *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" type:"Struct"`
	// The function name.
	FunctionName *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	// The invocation type. Valid values:
	//
	// - Sync: synchronous.
	//
	// - Async: asynchronous.
	InvocationType *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	// The alias of the service to which the function belongs.
	Qualifier *UpdateEventStreamingRequestSinkSinkFcParametersQualifier `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	// The name of the service.
	ServiceName *UpdateEventStreamingRequestSinkSinkFcParametersServiceName `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) GetBody() *UpdateEventStreamingRequestSinkSinkFcParametersBody {
	return s.Body
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) GetConcurrency() *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency {
	return s.Concurrency
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) GetDataFormat() *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat {
	return s.DataFormat
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) GetFunctionName() *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName {
	return s.FunctionName
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) GetInvocationType() *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType {
	return s.InvocationType
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) GetQualifier() *UpdateEventStreamingRequestSinkSinkFcParametersQualifier {
	return s.Qualifier
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) GetServiceName() *UpdateEventStreamingRequestSinkSinkFcParametersServiceName {
	return s.ServiceName
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkFcParametersBody) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetConcurrency(v *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.Concurrency = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetDataFormat(v *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.DataFormat = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetFunctionName(v *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.FunctionName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetInvocationType(v *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.InvocationType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetQualifier(v *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.Qualifier = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetServiceName(v *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.ServiceName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.Concurrency != nil {
		if err := s.Concurrency.Validate(); err != nil {
			return err
		}
	}
	if s.DataFormat != nil {
		if err := s.DataFormat.Validate(); err != nil {
			return err
		}
	}
	if s.FunctionName != nil {
		if err := s.FunctionName.Validate(); err != nil {
			return err
		}
	}
	if s.InvocationType != nil {
		if err := s.InvocationType.Validate(); err != nil {
			return err
		}
	}
	if s.Qualifier != nil {
		if err := s.Qualifier.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceName != nil {
		if err := s.ServiceName.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkFcParametersBody struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersBody) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersBody) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersBody {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFcParametersConcurrency struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The delivery concurrency. Minimum value: 1.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFcParametersDataFormat struct {
	// The transformation format. Valid values:
	//
	// - ORIGINAL: complete event
	//
	// - JSONPATH: partial event
	//
	// - CONSTANT: constant
	//
	// - TEMPLATE: template
	//
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// $.data.key
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersDataFormat) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersDataFormat) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFcParametersFunctionName struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The function name.
	//
	// example:
	//
	// mFunction
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFcParametersInvocationType struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The invocation type.
	//
	// example:
	//
	// Async
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFcParametersQualifier struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The alias of the service to which the function belongs.
	//
	// example:
	//
	// LATEST
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersQualifier) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersQualifier) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFcParametersServiceName struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// myService
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersServiceName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersServiceName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFnfParameters struct {
	// The execution name.
	ExecutionName *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty" type:"Struct"`
	// The flow name.
	FlowName *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName `json:"FlowName,omitempty" xml:"FlowName,omitempty" type:"Struct"`
	// The execution input information.
	Input *UpdateEventStreamingRequestSinkSinkFnfParametersInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The role configuration.
	RoleName *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkFnfParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFnfParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) GetExecutionName() *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	return s.ExecutionName
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) GetFlowName() *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName {
	return s.FlowName
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) GetInput() *UpdateEventStreamingRequestSinkSinkFnfParametersInput {
	return s.Input
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) GetRoleName() *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName {
	return s.RoleName
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) SetExecutionName(v *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) *UpdateEventStreamingRequestSinkSinkFnfParameters {
	s.ExecutionName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) SetFlowName(v *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) *UpdateEventStreamingRequestSinkSinkFnfParameters {
	s.FlowName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) SetInput(v *UpdateEventStreamingRequestSinkSinkFnfParametersInput) *UpdateEventStreamingRequestSinkSinkFnfParameters {
	s.Input = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) SetRoleName(v *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) *UpdateEventStreamingRequestSinkSinkFnfParameters {
	s.RoleName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) Validate() error {
	if s.ExecutionName != nil {
		if err := s.ExecutionName.Validate(); err != nil {
			return err
		}
	}
	if s.FlowName != nil {
		if err := s.FlowName.Validate(); err != nil {
			return err
		}
	}
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.RoleName != nil {
		if err := s.RoleName.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The execution name.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFnfParametersFlowName struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The flow name.
	//
	// example:
	//
	// test-streaming-fnf
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFnfParametersInput struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The execution input information.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersInput) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersInput) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersInput) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersInput) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersInput) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersInput) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersInput) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersInput) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFnfParametersRoleName struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role configuration.
	//
	// example:
	//
	// Al****FNF-x****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkKafkaParameters struct {
	// The Kafka write acknowledgment mode. Valid values:
	//
	// - acks=0: No response is required from the server. Performance is high, but the risk of data loss is high.
	//
	// - acks=1: A response is returned after the primary node writes successfully. Performance is moderate, and the risk of data loss is moderate. Data loss may occur if the primary node goes down.
	//
	// - acks=all: A response is returned after the primary node writes successfully and the secondary nodes synchronize successfully. Performance is low, but data is more secure. Data loss occurs only if both the primary and secondary nodes go down.
	Acks            *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	CompressionType *string                                                 `json:"CompressionType,omitempty" xml:"CompressionType,omitempty"`
	// Specifies the target Topic routing strategy for messages. If both the Topic parameter and the DynamicTopic parameter are specified, the DynamicTopic parameter takes precedence. Two configuration modes are supported:
	//
	//     1. **Static constant mode**: Specify a fixed Topic name string (for example, "order_created"). All messages are sent to this Topic.
	//
	//     2. **Dynamic extraction mode**: Specify a standard JSONPath expression (for example, "$.user.id" or "$.metadata.category"). The system parses the upstream message body and extracts the matching field value as the target Topic name.
	DynamicTopic *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic `json:"DynamicTopic,omitempty" xml:"DynamicTopic,omitempty" type:"Struct"`
	// The additional metadata of the Kafka message.
	Headers *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The target service type is ApsaraMQ for Kafka.
	InstanceId *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message identifier.
	Key *UpdateEventStreamingRequestSinkSinkKafkaParametersKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	// The topic name.
	Topic *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The message body content.
	Value *UpdateEventStreamingRequestSinkSinkKafkaParametersValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) GetAcks() *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks {
	return s.Acks
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) GetCompressionType() *string {
	return s.CompressionType
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) GetDynamicTopic() *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic {
	return s.DynamicTopic
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) GetHeaders() *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders {
	return s.Headers
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) GetInstanceId() *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) GetKey() *UpdateEventStreamingRequestSinkSinkKafkaParametersKey {
	return s.Key
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) GetTopic() *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) GetValue() *UpdateEventStreamingRequestSinkSinkKafkaParametersValue {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetAcks(v *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.Acks = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetCompressionType(v string) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.CompressionType = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetDynamicTopic(v *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.DynamicTopic = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetHeaders(v *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.Headers = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetInstanceId(v *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.InstanceId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetKey(v *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.Key = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetTopic(v *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.Topic = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetValue(v *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.Value = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) Validate() error {
	if s.Acks != nil {
		if err := s.Acks.Validate(); err != nil {
			return err
		}
	}
	if s.DynamicTopic != nil {
		if err := s.DynamicTopic.Validate(); err != nil {
			return err
		}
	}
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceId != nil {
		if err := s.InstanceId.Validate(); err != nil {
			return err
		}
	}
	if s.Key != nil {
		if err := s.Key.Validate(); err != nil {
			return err
		}
	}
	if s.Topic != nil {
		if err := s.Topic.Validate(); err != nil {
			return err
		}
	}
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersAcks struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Kafka write acknowledgment mode. Valid values:
	//
	// - acks=0: No response is required from the server. Performance is high, but the risk of data loss is high.
	//
	// - acks=1: A response is returned after the primary node writes successfully. Performance is moderate, and the risk of data loss is moderate. Data loss may occur if the primary node goes down.
	//
	// - acks=all: A response is returned after the primary node writes successfully and the secondary nodes synchronize successfully. Performance is low, but data is more secure. Data loss occurs only if both the primary and secondary nodes go down.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic struct {
	// The transformation type. Valid values:
	//
	// - CONSTANT: fixed value
	//
	// - JSONPATH: extracted from upstream based on path
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersDynamicTopic) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders struct {
	// The transformation format. Valid values:
	//
	// - ORIGINAL: complete event
	//
	// - JSONPATH: partial event
	//
	// - CONSTANT: constant
	//
	// - TEMPLATE: template
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// Defaut_1283278472_s****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersKey struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The message identifier.
	//
	// example:
	//
	// key
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersKey) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersTopic struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic name.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersValue struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersValue) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkMNSParameters struct {
	// The message content.
	Body *UpdateEventStreamingRequestSinkSinkMNSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// Specifies whether to enable Base64 encoding.
	IsBase64Encode *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	// The target service type is Simple Message Queue (formerly MNS).
	QueueName *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkMNSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParameters) GetBody() *UpdateEventStreamingRequestSinkSinkMNSParametersBody {
	return s.Body
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParameters) GetIsBase64Encode() *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	return s.IsBase64Encode
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParameters) GetQueueName() *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName {
	return s.QueueName
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkMNSParametersBody) *UpdateEventStreamingRequestSinkSinkMNSParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParameters) SetIsBase64Encode(v *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) *UpdateEventStreamingRequestSinkSinkMNSParameters {
	s.IsBase64Encode = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParameters) SetQueueName(v *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) *UpdateEventStreamingRequestSinkSinkMNSParameters {
	s.QueueName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParameters) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.IsBase64Encode != nil {
		if err := s.IsBase64Encode.Validate(); err != nil {
			return err
		}
	}
	if s.QueueName != nil {
		if err := s.QueueName.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkMNSParametersBody struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersBody) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersBody) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode struct {
	// The event transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Specifies whether to enable Base64 encoding.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetForm(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetValue(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkMNSParametersQueueName struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the queue in Simple Message Queue (formerly MNS).
	//
	// example:
	//
	// MyQueue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters struct {
	// The authentication type. Valid values:
	//
	// - ACL
	//
	// - No configuration required
	//
	// example:
	//
	// ACL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The message body content.
	Body *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The endpoint.
	//
	// example:
	//
	// 192.168.1.1:9876
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The Exchange name in RabbitMQ. This parameter takes effect when TargetType is set to Exchange.
	//
	// example:
	//
	// my-exchange
	Exchange *string `json:"Exchange,omitempty" xml:"Exchange,omitempty"`
	// The unique identifier of the message.
	MessageId *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// The network type. Valid values:
	//
	// - PrivateNetwork
	//
	// - PublicNetwork
	//
	// example:
	//
	// PublicNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The password used to access the RabbitMQ instance.
	//
	// example:
	//
	// ****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The additional properties of the message.
	Properties *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The queue name in RabbitMQ. This parameter takes effect only when TargetType is set to Queue.
	//
	// example:
	//
	// my-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The message routing key.
	RoutingKey *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// The security group ID.
	//
	// example:
	//
	// sg-uf6of9452b2pba82c ****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The target type for message delivery. Valid values:
	//
	// - **Exchange:*	- Messages are routed through an exchange.
	//
	// - **Queue:*	- Messages are delivered directly to the specified queue.
	//
	// example:
	//
	// Exchange
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The username used to access the RabbitMQ instance.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6of9452b2pba82c ****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The virtual host name of the RabbitMQ instance.
	//
	// example:
	//
	// Vhost1
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf6of9452b2pba82c ****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetBody() *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody {
	return s.Body
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetExchange() *string {
	return s.Exchange
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetMessageId() *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId {
	return s.MessageId
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetPassword() *string {
	return s.Password
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetProperties() *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties {
	return s.Properties
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetRoutingKey() *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	return s.RoutingKey
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetUsername() *string {
	return s.Username
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetAuthType(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.AuthType = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetEndpoint(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Endpoint = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetExchange(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Exchange = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetMessageId(v *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetNetworkType(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.NetworkType = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetPassword(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Password = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetProperties(v *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetQueueName(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetRoutingKey(v *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetSecurityGroupId(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetTargetType(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.TargetType = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetUsername(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Username = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetVSwitchIds(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetVirtualHostName(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetVpcId(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.VpcId = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.MessageId != nil {
		if err := s.MessageId.Validate(); err != nil {
			return err
		}
	}
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	if s.RoutingKey != nil {
		if err := s.RoutingKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The raw data value.
	//
	// example:
	//
	// {"key": "value"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The message ID value.
	//
	// example:
	//
	// 12345
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The property content.
	//
	// example:
	//
	// {"env": "prod"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) SetForm(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) SetValue(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The routing key value.
	//
	// example:
	//
	// {"Form": "CONSTANT", "Value": "my-routing-key"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetForm(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetValue(v string) *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkPrometheusParameters struct {
	// The authorization type.
	AuthorizationType *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty" type:"Struct"`
	// The metric content.
	Data *UpdateEventStreamingRequestSinkSinkPrometheusParametersData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The data structure of the request header parameters.
	HeaderParameters *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Struct"`
	// The network type.
	NetworkType *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	// The password.
	Password *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword `json:"Password,omitempty" xml:"Password,omitempty" type:"Struct"`
	// The security group ID.
	SecurityGroupId *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	// The Prometheus Remote Write URL.
	URL *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL `json:"URL,omitempty" xml:"URL,omitempty" type:"Struct"`
	// The username.
	Username *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername `json:"Username,omitempty" xml:"Username,omitempty" type:"Struct"`
	// The vSwitch ID.
	VSwitchId *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Struct"`
	// VPC ID。
	VpcId *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) GetAuthorizationType() *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	return s.AuthorizationType
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) GetData() *UpdateEventStreamingRequestSinkSinkPrometheusParametersData {
	return s.Data
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) GetHeaderParameters() *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters {
	return s.HeaderParameters
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) GetNetworkType() *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	return s.NetworkType
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) GetPassword() *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	return s.Password
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) GetSecurityGroupId() *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) GetURL() *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL {
	return s.URL
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) GetUsername() *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	return s.Username
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) GetVSwitchId() *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	return s.VSwitchId
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) GetVpcId() *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetAuthorizationType(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.AuthorizationType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetData(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Data = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetHeaderParameters(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.HeaderParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetNetworkType(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.NetworkType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetPassword(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Password = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetSecurityGroupId(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.SecurityGroupId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetURL(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.URL = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetUsername(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Username = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetVSwitchId(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.VSwitchId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetVpcId(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.VpcId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) Validate() error {
	if s.AuthorizationType != nil {
		if err := s.AuthorizationType.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.HeaderParameters != nil {
		if err := s.HeaderParameters.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkType != nil {
		if err := s.NetworkType.Validate(); err != nil {
			return err
		}
	}
	if s.Password != nil {
		if err := s.Password.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroupId != nil {
		if err := s.SecurityGroupId.Validate(); err != nil {
			return err
		}
	}
	if s.URL != nil {
		if err := s.URL.Validate(); err != nil {
			return err
		}
	}
	if s.Username != nil {
		if err := s.Username.Validate(); err != nil {
			return err
		}
	}
	if s.VSwitchId != nil {
		if err := s.VSwitchId.Validate(); err != nil {
			return err
		}
	}
	if s.VpcId != nil {
		if err := s.VpcId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The authorization type.
	//
	// example:
	//
	// BASIC_AUTH
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersData struct {
	// The transformation format. Default value: JSONPATH.
	//
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The metric content.
	//
	// example:
	//
	// $.data
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersData) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersData) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters struct {
	// The transformation format. Valid values:
	//
	// - JSONPATH
	//
	// - CONSTANT
	//
	// - TEMPLATE
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The HTTP request header template style. This parameter is required when Form is set to TEMPLATE. The event content transformation result must be in JSON format.
	//
	// example:
	//
	// {
	//
	//     "user_name":"${name}"
	//
	// }
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value. Valid values:
	//
	// - If Form is CONSTANT: a constant value
	//
	// - If Form is JSONPATH: JSONPath extraction content
	//
	// - If Form is TEMPLATE: a template variable
	//
	// Note: The Value field cannot exceed 10240 characters.
	//
	// example:
	//
	// name
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The network type. Valid values:
	//
	// - PublicNetwork
	//
	// - PrivateNetwork
	//
	// example:
	//
	// PrivateNetwork
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The password.
	//
	// example:
	//
	// abc
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-mw43*****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersURL struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Prometheus Remote Write URL.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The username.
	//
	// example:
	//
	// ***admin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-dwaafds****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-adw1awdw*****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParameters struct {
	// The message content.
	Body *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The Exchange mode. This parameter is required only when TargetType is set to Exchange.
	Exchange *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	// The target service type is ApsaraMQ for RabbitMQ.
	InstanceId *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message ID.
	MessageId *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// The filter properties.
	Properties *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The Queue mode. This parameter is required only when TargetType is set to Queue.
	QueueName *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	// The routing rule of the message. This parameter is required only when TargetType is set to Exchange.
	RoutingKey *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// The target type.
	TargetType *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
	// The name of the vhost of the ApsaraMQ for RabbitMQ instance.
	VirtualHostName *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) GetBody() *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	return s.Body
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) GetExchange() *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	return s.Exchange
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) GetInstanceId() *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) GetMessageId() *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	return s.MessageId
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) GetProperties() *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	return s.Properties
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) GetQueueName() *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	return s.QueueName
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) GetRoutingKey() *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	return s.RoutingKey
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) GetTargetType() *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	return s.TargetType
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) GetVirtualHostName() *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	return s.VirtualHostName
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetExchange(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Exchange = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetInstanceId(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.InstanceId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetMessageId(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetProperties(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetQueueName(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.QueueName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetRoutingKey(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetTargetType(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.TargetType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetVirtualHostName(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.VirtualHostName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.Exchange != nil {
		if err := s.Exchange.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceId != nil {
		if err := s.InstanceId.Validate(); err != nil {
			return err
		}
	}
	if s.MessageId != nil {
		if err := s.MessageId.Validate(); err != nil {
			return err
		}
	}
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	if s.QueueName != nil {
		if err := s.QueueName.Validate(); err != nil {
			return err
		}
	}
	if s.RoutingKey != nil {
		if err := s.RoutingKey.Validate(); err != nil {
			return err
		}
	}
	if s.TargetType != nil {
		if err := s.TargetType.Validate(); err != nil {
			return err
		}
	}
	if s.VirtualHostName != nil {
		if err := s.VirtualHostName.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the Exchange of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// a_exchange
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance ID of ApsaraMQ for RabbitMQ.
	//
	// example:
	//
	// amqp-cn-2r42e73o****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the queue of the instance.
	//
	// example:
	//
	// MyQueue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The routing rule of the message.
	//
	// example:
	//
	// housekeeping
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The target type. Valid values:
	//
	// - Exchange: Exchange mode.
	//
	// - Queue: Queue mode.
	//
	// example:
	//
	// Exchange/Queue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the vhost of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// rabbit-host
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters struct {
	// The timestamp of message consumption.
	ConsumeTimestamp *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp `json:"ConsumeTimestamp,omitempty" xml:"ConsumeTimestamp,omitempty" type:"Struct"`
	// The Group ID of the consumer group.
	Group *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The instance ID.
	//
	// example:
	//
	// MQ_INST_164901546557****_BAAN****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type. Valid values:
	//
	// - Cloud_4: Alibaba Cloud RocketMQ 4.0 instance
	//
	// - Cloud_5: Alibaba Cloud RocketMQ 5.0 instance
	//
	// example:
	//
	// Cloud_4
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The topic of the ApsaraMQ for RocketMQ instance.
	Topic *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GetConsumeTimestamp() *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	return s.ConsumeTimestamp
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GetGroup() *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup {
	return s.Group
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GetTopic() *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) SetConsumeTimestamp(v *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	s.ConsumeTimestamp = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) SetGroup(v *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	s.Group = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) SetInstanceType(v string) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	s.InstanceType = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) SetTopic(v *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	s.Topic = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) Validate() error {
	if s.ConsumeTimestamp != nil {
		if err := s.ConsumeTimestamp.Validate(); err != nil {
			return err
		}
	}
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	if s.Topic != nil {
		if err := s.Topic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 1570761026400
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Group ID
	//
	// example:
	//
	// GID_EVENTBRIDGE_1736234******
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic name of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// Mytopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParameters struct {
	// The message content.
	Body *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The delivery order type of the message. This parameter is optional. Default value: concurrent delivery.
	DeliveryOrderType *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType `json:"DeliveryOrderType,omitempty" xml:"DeliveryOrderType,omitempty" type:"Struct"`
	// The instance endpoint.
	InstanceEndpoint *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty" type:"Struct"`
	// The target service type is ApsaraMQ for RocketMQ.
	InstanceId *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The instance password.
	InstancePassword *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty" type:"Struct"`
	// The instance type.
	InstanceType *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Struct"`
	// The instance username.
	InstanceUsername *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty" type:"Struct"`
	// The filter properties.
	Keys *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	// The network type. Valid values:
	//
	// - PublicNetwork
	//
	// - PrivateNetwork
	Network *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The filter properties.
	Properties *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The security group ID.
	SecurityGroupId *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	// The sharding key of the message.
	//
	// > When DeliveryOrderType is set to Orderly, this parameter specifies the event content transformation rule for the ShardingKey property when writing messages downstream. When Source is RocketMQ, ShardingKey can be empty. In this case, the upstream BrokerName and QueueId are concatenated to generate the message ShardingKey.
	ShardingKey *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey `json:"ShardingKey,omitempty" xml:"ShardingKey,omitempty" type:"Struct"`
	// The filter properties.
	Tags *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The topic of the ApsaraMQ for RocketMQ instance.
	Topic *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The vSwitch ID.
	VSwitchIds *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	// The VPC ID.
	VpcId *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetBody() *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody {
	return s.Body
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetDeliveryOrderType() *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType {
	return s.DeliveryOrderType
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetInstanceEndpoint() *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint {
	return s.InstanceEndpoint
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetInstanceId() *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetInstancePassword() *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword {
	return s.InstancePassword
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetInstanceType() *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType {
	return s.InstanceType
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetInstanceUsername() *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername {
	return s.InstanceUsername
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetKeys() *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	return s.Keys
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetNetwork() *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork {
	return s.Network
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetProperties() *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	return s.Properties
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetSecurityGroupId() *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetShardingKey() *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey {
	return s.ShardingKey
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetTags() *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags {
	return s.Tags
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetTopic() *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetVSwitchIds() *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) GetVpcId() *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetDeliveryOrderType(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.DeliveryOrderType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceEndpoint(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceEndpoint = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceId(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetInstancePassword(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstancePassword = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceType(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceUsername(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceUsername = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetKeys(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Keys = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetNetwork(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Network = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetProperties(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Properties = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetSecurityGroupId(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.SecurityGroupId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetShardingKey(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.ShardingKey = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetTags(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Tags = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetTopic(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Topic = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetVSwitchIds(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.VSwitchIds = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetVpcId(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.VpcId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.DeliveryOrderType != nil {
		if err := s.DeliveryOrderType.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceEndpoint != nil {
		if err := s.InstanceEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceId != nil {
		if err := s.InstanceId.Validate(); err != nil {
			return err
		}
	}
	if s.InstancePassword != nil {
		if err := s.InstancePassword.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceType != nil {
		if err := s.InstanceType.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceUsername != nil {
		if err := s.InstanceUsername.Validate(); err != nil {
			return err
		}
	}
	if s.Keys != nil {
		if err := s.Keys.Validate(); err != nil {
			return err
		}
	}
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroupId != nil {
		if err := s.SecurityGroupId.Validate(); err != nil {
			return err
		}
	}
	if s.ShardingKey != nil {
		if err := s.ShardingKey.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.Topic != nil {
		if err := s.Topic.Validate(); err != nil {
			return err
		}
	}
	if s.VSwitchIds != nil {
		if err := s.VSwitchIds.Validate(); err != nil {
			return err
		}
	}
	if s.VpcId != nil {
		if err := s.VpcId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersBody struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The delivery order type. Valid values:
	//
	// - **Orderly:*	- ordered delivery
	//
	// - **Concurrently:*	- concurrent delivery
	//
	// example:
	//
	// Concurrently
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance endpoint.
	//
	// example:
	//
	// vbr-8vbsvkkbpf3vb0zef****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance ID of ApsaraMQ for RocketMQ.
	//
	// example:
	//
	// MQ_INST_164901546557****_BAAN****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance password.
	//
	// example:
	//
	// admin****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance type. Valid values:
	//
	// - Cloud_4: Alibaba Cloud RocketMQ 4.0 instance (default)
	//
	// - Cloud_5: Alibaba Cloud RocketMQ 5.0 instance
	//
	// - SelfBuilt: self-managed Apache RocketMQ cluster
	//
	// example:
	//
	// Cloud_4
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance username.
	//
	// example:
	//
	// admin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The network type. Valid values:
	//
	// - PublicNetwork
	//
	// - PrivateNetwork
	//
	// example:
	//
	// PublicNetwork
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// b4bf375515f6440f942e3a20c33d****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The sharding key value.
	//
	// example:
	//
	// order_id
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersTags struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// Mytopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vbr-8vb835n3zf9shwl****mp
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId struct {
	// The event transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vbr-8vb835n3zf9shwlvb****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkSLSParameters struct {
	// The content sent to SLS.
	Body *UpdateEventStreamingRequestSinkSinkSLSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The custom log key-value pairs. This parameter takes effect only when ContentType is set to KeyValue. Each key-value pair is represented by Key_n and Value_n.
	ContentSchema *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema `json:"ContentSchema,omitempty" xml:"ContentSchema,omitempty" type:"Struct"`
	// The SLS data format. You can select the default format or configure specified key-value pairs. Valid values:
	//
	// - JSON
	//
	// - KeyValue
	ContentType *UpdateEventStreamingRequestSinkSinkSLSParametersContentType `json:"ContentType,omitempty" xml:"ContentType,omitempty" type:"Struct"`
	// The Logstore of Simple Log Service (SLS).
	LogStore *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	// The log project of Simple Log Service (SLS).
	Project *UpdateEventStreamingRequestSinkSinkSLSParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The role name used to authorize the event bus EventBridge to read SLS log content. When creating the role in the Resource Access Management (RAM) console, select "Alibaba Cloud Service" and set "Trusted Service" to "EventBridge".
	RoleName *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The topic where the log resides, corresponding to the SLS reserved field "topic".
	Topic *UpdateEventStreamingRequestSinkSinkSLSParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) GetBody() *UpdateEventStreamingRequestSinkSinkSLSParametersBody {
	return s.Body
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) GetContentSchema() *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema {
	return s.ContentSchema
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) GetContentType() *UpdateEventStreamingRequestSinkSinkSLSParametersContentType {
	return s.ContentType
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) GetLogStore() *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore {
	return s.LogStore
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) GetProject() *UpdateEventStreamingRequestSinkSinkSLSParametersProject {
	return s.Project
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) GetRoleName() *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName {
	return s.RoleName
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) GetTopic() *UpdateEventStreamingRequestSinkSinkSLSParametersTopic {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkSLSParametersBody) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetContentSchema(v *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.ContentSchema = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetContentType(v *UpdateEventStreamingRequestSinkSinkSLSParametersContentType) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.ContentType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetLogStore(v *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.LogStore = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetProject(v *UpdateEventStreamingRequestSinkSinkSLSParametersProject) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.Project = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetRoleName(v *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.RoleName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetTopic(v *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.Topic = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.ContentSchema != nil {
		if err := s.ContentSchema.Validate(); err != nil {
			return err
		}
	}
	if s.ContentType != nil {
		if err := s.ContentType.Validate(); err != nil {
			return err
		}
	}
	if s.LogStore != nil {
		if err := s.LogStore.Validate(); err != nil {
			return err
		}
	}
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	if s.RoleName != nil {
		if err := s.RoleName.Validate(); err != nil {
			return err
		}
	}
	if s.Topic != nil {
		if err := s.Topic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSinkSinkSLSParametersBody struct {
	// The transformation format.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersBody) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersBody) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema struct {
	// The transformation format.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The custom key-value pairs.
	//
	// example:
	//
	// {"Key_1":{"form":"CONSTANT","value":"demoKey"},"Value_1":{"form":"JSONPATH","value":"$.data.value"}}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkSLSParametersContentType struct {
	// The transformation format.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The SLS data format.
	//
	// example:
	//
	// JSON
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersContentType) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersContentType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentType) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentType) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentType) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersContentType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersContentType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersContentType {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersContentType) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkSLSParametersLogStore struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Logstore of Simple Log Service (SLS).
	//
	// example:
	//
	// test-logstore
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkSLSParametersProject struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The log project of Simple Log Service (SLS).
	//
	// example:
	//
	// test-project
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersProject) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersProject) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersProject) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersProject) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersProject) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersProject) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersProject) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersProject) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersProject) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkSLSParametersRoleName struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name used to authorize the event bus EventBridge to read SLS log content. When creating the role in the Resource Access Management (RAM) console, select "Alibaba Cloud Service" and set "Trusted Service" to "EventBridge".
	//
	// example:
	//
	// testRole
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkSLSParametersTopic struct {
	// The transformation format. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic where the log resides, corresponding to the SLS reserved field "topic".
	//
	// example:
	//
	// testTopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersTopic) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSource struct {
	SourceApacheKafkaParameters *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters `json:"SourceApacheKafkaParameters,omitempty" xml:"SourceApacheKafkaParameters,omitempty" type:"Struct"`
	// The Source RocketMQ Checkpoint source.
	SourceApacheRocketMQCheckpointParameters *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters `json:"SourceApacheRocketMQCheckpointParameters,omitempty" xml:"SourceApacheRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The custom connector Apache Kafka event source.
	SourceCustomizedKafkaConnectorParameters *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters `json:"SourceCustomizedKafkaConnectorParameters,omitempty" xml:"SourceCustomizedKafkaConnectorParameters,omitempty" type:"Struct"`
	// The custom Kafka event source.
	SourceCustomizedKafkaParameters *UpdateEventStreamingRequestSourceSourceCustomizedKafkaParameters `json:"SourceCustomizedKafkaParameters,omitempty" xml:"SourceCustomizedKafkaParameters,omitempty" type:"Struct"`
	// The Source DTS source.
	SourceDTSParameters *UpdateEventStreamingRequestSourceSourceDTSParameters `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	// The Source EventBus source.
	SourceEventBusParameters   *UpdateEventStreamingRequestSourceSourceEventBusParameters `json:"SourceEventBusParameters,omitempty" xml:"SourceEventBusParameters,omitempty" type:"Struct"`
	SourceFeiShuDocsParameters *SourceFeiShuDocsParameters                                `json:"SourceFeiShuDocsParameters,omitempty" xml:"SourceFeiShuDocsParameters,omitempty"`
	SourceJDBCParameters       *SourceJDBCParameters                                      `json:"SourceJDBCParameters,omitempty" xml:"SourceJDBCParameters,omitempty"`
	// The Source Kafka source.
	SourceKafkaParameters *UpdateEventStreamingRequestSourceSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The Source Simple Message Queue (formerly
	SourceMNSParameters *UpdateEventStreamingRequestSourceSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The Source MQTT source.
	SourceMQTTParameters *UpdateEventStreamingRequestSourceSourceMQTTParameters `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	// The Source MySQL source.
	SourceMySQLParameters *SourceMySQLParameters `json:"SourceMySQLParameters,omitempty" xml:"SourceMySQLParameters,omitempty"`
	// The Source OSS event source.
	SourceOSSParameters *UpdateEventStreamingRequestSourceSourceOSSParameters `json:"SourceOSSParameters,omitempty" xml:"SourceOSSParameters,omitempty" type:"Struct"`
	// The Source Open Source RabbitMQ source.
	SourceOpenSourceRabbitMQParameters *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters `json:"SourceOpenSourceRabbitMQParameters,omitempty" xml:"SourceOpenSourceRabbitMQParameters,omitempty" type:"Struct"`
	SourcePostgreSQLParameters         *SourcePostgreSQLParameters                                          `json:"SourcePostgreSQLParameters,omitempty" xml:"SourcePostgreSQLParameters,omitempty"`
	// The Source Prometheus event source.
	SourcePrometheusParameters      *UpdateEventStreamingRequestSourceSourcePrometheusParameters `json:"SourcePrometheusParameters,omitempty" xml:"SourcePrometheusParameters,omitempty" type:"Struct"`
	SourceRabbitMQMetaParameters    *SourceRabbitMQMetaParameters                                `json:"SourceRabbitMQMetaParameters,omitempty" xml:"SourceRabbitMQMetaParameters,omitempty"`
	SourceRabbitMQMsgSyncParameters *SourceRabbitMQMsgSyncParameters                             `json:"SourceRabbitMQMsgSyncParameters,omitempty" xml:"SourceRabbitMQMsgSyncParameters,omitempty"`
	// The Source RabbitMQ source.
	SourceRabbitMQParameters *UpdateEventStreamingRequestSourceSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// The Source RocketMQ Checkpoint source.
	SourceRocketMQCheckpointParameters *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters `json:"SourceRocketMQCheckpointParameters,omitempty" xml:"SourceRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The Source RocketMQ source.
	SourceRocketMQParameters *UpdateEventStreamingRequestSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The Source SLS source.
	SourceSLSParameters *UpdateEventStreamingRequestSourceSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSource) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSource) GetSourceApacheKafkaParameters() *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	return s.SourceApacheKafkaParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceApacheRocketMQCheckpointParameters() *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	return s.SourceApacheRocketMQCheckpointParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceCustomizedKafkaConnectorParameters() *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters {
	return s.SourceCustomizedKafkaConnectorParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceCustomizedKafkaParameters() *UpdateEventStreamingRequestSourceSourceCustomizedKafkaParameters {
	return s.SourceCustomizedKafkaParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceDTSParameters() *UpdateEventStreamingRequestSourceSourceDTSParameters {
	return s.SourceDTSParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceEventBusParameters() *UpdateEventStreamingRequestSourceSourceEventBusParameters {
	return s.SourceEventBusParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceFeiShuDocsParameters() *SourceFeiShuDocsParameters {
	return s.SourceFeiShuDocsParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceJDBCParameters() *SourceJDBCParameters {
	return s.SourceJDBCParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceKafkaParameters() *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	return s.SourceKafkaParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceMNSParameters() *UpdateEventStreamingRequestSourceSourceMNSParameters {
	return s.SourceMNSParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceMQTTParameters() *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	return s.SourceMQTTParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceMySQLParameters() *SourceMySQLParameters {
	return s.SourceMySQLParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceOSSParameters() *UpdateEventStreamingRequestSourceSourceOSSParameters {
	return s.SourceOSSParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceOpenSourceRabbitMQParameters() *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	return s.SourceOpenSourceRabbitMQParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourcePostgreSQLParameters() *SourcePostgreSQLParameters {
	return s.SourcePostgreSQLParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourcePrometheusParameters() *UpdateEventStreamingRequestSourceSourcePrometheusParameters {
	return s.SourcePrometheusParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceRabbitMQMetaParameters() *SourceRabbitMQMetaParameters {
	return s.SourceRabbitMQMetaParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceRabbitMQMsgSyncParameters() *SourceRabbitMQMsgSyncParameters {
	return s.SourceRabbitMQMsgSyncParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceRabbitMQParameters() *UpdateEventStreamingRequestSourceSourceRabbitMQParameters {
	return s.SourceRabbitMQParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceRocketMQCheckpointParameters() *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters {
	return s.SourceRocketMQCheckpointParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceRocketMQParameters() *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	return s.SourceRocketMQParameters
}

func (s *UpdateEventStreamingRequestSource) GetSourceSLSParameters() *UpdateEventStreamingRequestSourceSourceSLSParameters {
	return s.SourceSLSParameters
}

func (s *UpdateEventStreamingRequestSource) SetSourceApacheKafkaParameters(v *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) *UpdateEventStreamingRequestSource {
	s.SourceApacheKafkaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceApacheRocketMQCheckpointParameters(v *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) *UpdateEventStreamingRequestSource {
	s.SourceApacheRocketMQCheckpointParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceCustomizedKafkaConnectorParameters(v *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) *UpdateEventStreamingRequestSource {
	s.SourceCustomizedKafkaConnectorParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceCustomizedKafkaParameters(v *UpdateEventStreamingRequestSourceSourceCustomizedKafkaParameters) *UpdateEventStreamingRequestSource {
	s.SourceCustomizedKafkaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceDTSParameters(v *UpdateEventStreamingRequestSourceSourceDTSParameters) *UpdateEventStreamingRequestSource {
	s.SourceDTSParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceEventBusParameters(v *UpdateEventStreamingRequestSourceSourceEventBusParameters) *UpdateEventStreamingRequestSource {
	s.SourceEventBusParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceFeiShuDocsParameters(v *SourceFeiShuDocsParameters) *UpdateEventStreamingRequestSource {
	s.SourceFeiShuDocsParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceJDBCParameters(v *SourceJDBCParameters) *UpdateEventStreamingRequestSource {
	s.SourceJDBCParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceKafkaParameters(v *UpdateEventStreamingRequestSourceSourceKafkaParameters) *UpdateEventStreamingRequestSource {
	s.SourceKafkaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceMNSParameters(v *UpdateEventStreamingRequestSourceSourceMNSParameters) *UpdateEventStreamingRequestSource {
	s.SourceMNSParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceMQTTParameters(v *UpdateEventStreamingRequestSourceSourceMQTTParameters) *UpdateEventStreamingRequestSource {
	s.SourceMQTTParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceMySQLParameters(v *SourceMySQLParameters) *UpdateEventStreamingRequestSource {
	s.SourceMySQLParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceOSSParameters(v *UpdateEventStreamingRequestSourceSourceOSSParameters) *UpdateEventStreamingRequestSource {
	s.SourceOSSParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceOpenSourceRabbitMQParameters(v *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) *UpdateEventStreamingRequestSource {
	s.SourceOpenSourceRabbitMQParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourcePostgreSQLParameters(v *SourcePostgreSQLParameters) *UpdateEventStreamingRequestSource {
	s.SourcePostgreSQLParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourcePrometheusParameters(v *UpdateEventStreamingRequestSourceSourcePrometheusParameters) *UpdateEventStreamingRequestSource {
	s.SourcePrometheusParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceRabbitMQMetaParameters(v *SourceRabbitMQMetaParameters) *UpdateEventStreamingRequestSource {
	s.SourceRabbitMQMetaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceRabbitMQMsgSyncParameters(v *SourceRabbitMQMsgSyncParameters) *UpdateEventStreamingRequestSource {
	s.SourceRabbitMQMsgSyncParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceRabbitMQParameters(v *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) *UpdateEventStreamingRequestSource {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceRocketMQCheckpointParameters(v *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) *UpdateEventStreamingRequestSource {
	s.SourceRocketMQCheckpointParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceRocketMQParameters(v *UpdateEventStreamingRequestSourceSourceRocketMQParameters) *UpdateEventStreamingRequestSource {
	s.SourceRocketMQParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceSLSParameters(v *UpdateEventStreamingRequestSourceSourceSLSParameters) *UpdateEventStreamingRequestSource {
	s.SourceSLSParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) Validate() error {
	if s.SourceApacheKafkaParameters != nil {
		if err := s.SourceApacheKafkaParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceApacheRocketMQCheckpointParameters != nil {
		if err := s.SourceApacheRocketMQCheckpointParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceCustomizedKafkaConnectorParameters != nil {
		if err := s.SourceCustomizedKafkaConnectorParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceCustomizedKafkaParameters != nil {
		if err := s.SourceCustomizedKafkaParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceDTSParameters != nil {
		if err := s.SourceDTSParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceEventBusParameters != nil {
		if err := s.SourceEventBusParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceFeiShuDocsParameters != nil {
		if err := s.SourceFeiShuDocsParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceJDBCParameters != nil {
		if err := s.SourceJDBCParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceKafkaParameters != nil {
		if err := s.SourceKafkaParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceMNSParameters != nil {
		if err := s.SourceMNSParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceMQTTParameters != nil {
		if err := s.SourceMQTTParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceMySQLParameters != nil {
		if err := s.SourceMySQLParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceOSSParameters != nil {
		if err := s.SourceOSSParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceOpenSourceRabbitMQParameters != nil {
		if err := s.SourceOpenSourceRabbitMQParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourcePostgreSQLParameters != nil {
		if err := s.SourcePostgreSQLParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourcePrometheusParameters != nil {
		if err := s.SourcePrometheusParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceRabbitMQMetaParameters != nil {
		if err := s.SourceRabbitMQMetaParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceRabbitMQMsgSyncParameters != nil {
		if err := s.SourceRabbitMQMsgSyncParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceRabbitMQParameters != nil {
		if err := s.SourceRabbitMQParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceRocketMQCheckpointParameters != nil {
		if err := s.SourceRocketMQCheckpointParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceRocketMQParameters != nil {
		if err := s.SourceRocketMQParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceSLSParameters != nil {
		if err := s.SourceSLSParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSourceSourceApacheKafkaParameters struct {
	Bootstraps       *string `json:"Bootstraps,omitempty" xml:"Bootstraps,omitempty"`
	ConsumerGroup    *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OffsetReset      *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	SaslMechanism    *string `json:"SaslMechanism,omitempty" xml:"SaslMechanism,omitempty"`
	SaslPassword     *string `json:"SaslPassword,omitempty" xml:"SaslPassword,omitempty"`
	SaslUser         *string `json:"SaslUser,omitempty" xml:"SaslUser,omitempty"`
	SecurityGroupId  *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityProtocol *string `json:"SecurityProtocol,omitempty" xml:"SecurityProtocol,omitempty"`
	// [Required for encrypted private key] The Kafka client private key password. This parameter is required when the client private key is password-protected (the PEM file contains \\"Proc-Type: 4,ENCRYPTED\\" or \\"ENCRYPTED\\" markers). Leave empty if the private key is not encrypted. Note: This password is only used to decrypt the private key and is unrelated to Kafka authentication.
	SslKeyPassword *string `json:"SslKeyPassword,omitempty" xml:"SslKeyPassword,omitempty"`
	// [Required for mutual authentication] The Kafka client certificate chain. This parameter is required when the Kafka server enables mutual SSL authentication (ssl.client.auth=required). Format: Base64-encoded PEM format containing the client certificate and the complete certificate chain (client certificate first, intermediate CA certificate next, root CA certificate optional). Note: Ensure each PEM file content starts with \\"-----BEGIN CERTIFICATE-----\\" and ends with \\"-----END CERTIFICATE-----\\", then Base64-encode the concatenated content.
	SslKeystoreCertificateChain *string `json:"SslKeystoreCertificateChain,omitempty" xml:"SslKeystoreCertificateChain,omitempty"`
	// [Required for bidirectional authentication] The SSL private key configuration object. When the Kafka server enables bidirectional SSL authentication, you must provide the client private key. Only KMS pattern is supported for the key: specify the Key Management EPS resource that stores the private key by using KmsArn. The system retrieves the private key content from KMS only in memory, which provides higher security. Configuration example: {\\"KmsArn\\": \\"acs:kms:ap-southeast-1:123456789:secret/ssl-key-xxxx\\", \\"KmsSecretValueKey\\": \\"keystore_private_key\\"}\\n"
	SslKeystoreKey *UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey `json:"SslKeystoreKey,omitempty" xml:"SslKeystoreKey,omitempty" type:"Struct"`
	// [Required for SSL] The Kafka server trusted certificate. Used to authenticate the validity of the Kafka Broker SSL certificate and prevent man-in-the-middle attacks. Format: Base64 encoding of PEM format, typically containing the CA certificate or the server certificate of the Kafka server. Example: Base64-encode the PEM file content of the CA certificate (ensure the content starts with \\"-----BEGIN CERTIFICATE-----\\" and ends with \\"-----END CERTIFICATE-----\\"). If Kafka uses a self-signed certificate, provide the CA certificate that issued the certificate.
	SslTruststoreCertificates *string `json:"SslTruststoreCertificates,omitempty" xml:"SslTruststoreCertificates,omitempty"`
	Topic                     *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds                *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	ValueDataType             *string `json:"ValueDataType,omitempty" xml:"ValueDataType,omitempty"`
	VpcId                     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetBootstraps() *string {
	return s.Bootstraps
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSaslMechanism() *string {
	return s.SaslMechanism
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSaslPassword() *string {
	return s.SaslPassword
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSaslUser() *string {
	return s.SaslUser
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSecurityProtocol() *string {
	return s.SecurityProtocol
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSslKeyPassword() *string {
	return s.SslKeyPassword
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSslKeystoreCertificateChain() *string {
	return s.SslKeystoreCertificateChain
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSslKeystoreKey() *UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey {
	return s.SslKeystoreKey
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSslTruststoreCertificates() *string {
	return s.SslTruststoreCertificates
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetValueDataType() *string {
	return s.ValueDataType
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetBootstraps(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.Bootstraps = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetConsumerGroup(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetNetworkType(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.NetworkType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetOffsetReset(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSaslMechanism(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SaslMechanism = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSaslPassword(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SaslPassword = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSaslUser(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SaslUser = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSecurityGroupId(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSecurityProtocol(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SecurityProtocol = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSslKeyPassword(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SslKeyPassword = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSslKeystoreCertificateChain(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SslKeystoreCertificateChain = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSslKeystoreKey(v *UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SslKeystoreKey = v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSslTruststoreCertificates(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SslTruststoreCertificates = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetTopic(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetVSwitchIds(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetValueDataType(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.ValueDataType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) SetVpcId(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParameters) Validate() error {
	if s.SslKeystoreKey != nil {
		if err := s.SslKeystoreKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey struct {
	// [Required] The KMS resource ARN that stores the SSL private key. Used to locate the Key Management Service instance that stores the client private key. Format example: \\"acs:kms:cn-hangzhou:123456789:secret/ssl-keystore-key-xxxx\\". Obtain this value from the ARN information of the corresponding key in the KMS console.
	KmsArn *string `json:"KmsArn,omitempty" xml:"KmsArn,omitempty"`
	// [KMS KV mode] The key name in the KMS credential. When the KMS credential is stored as a key-value (KV) structure, specify this parameter to indicate the key corresponding to the SSL private key. Example: if the KMS credential is \\"{"ssl_keystore_key":"-----BEGIN PRIVATE KEY-----...","ssl_truststore_key":"..."}\\", enter \\"ssl_keystore_key\\". Leave empty if the KMS credential is in plain text mode (directly stores the PEM content of the private key).
	KmsSecretValueKey *string `json:"KmsSecretValueKey,omitempty" xml:"KmsSecretValueKey,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey) GetKmsArn() *string {
	return s.KmsArn
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey) GetKmsSecretValueKey() *string {
	return s.KmsSecretValueKey
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey) SetKmsArn(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey {
	s.KmsArn = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey) SetKmsSecretValueKey(v string) *UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey {
	s.KmsSecretValueKey = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheKafkaParametersSslKeystoreKey) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters struct {
	// The endpoint of the Apache RocketMQ instance.
	//
	// example:
	//
	// 192.168.1.1:9876
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The password of the Apache RocketMQ instance.
	//
	// example:
	//
	// ****
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The username of the Apache RocketMQ instance.
	//
	// example:
	//
	// admin
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The network type. Valid values:
	//
	// - PublicNetwork
	//
	// - PrivateNetwork
	//
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-mw43*****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The topic of the Apache RocketMQ instance.
	Topics []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-dwaafds****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-adw1awdw*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetTopics() []*string {
	return s.Topics
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetInstanceEndpoint(v string) *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetInstancePassword(v string) *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.InstancePassword = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetInstanceUsername(v string) *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.InstanceUsername = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetNetworkType(v string) *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.NetworkType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetSecurityGroupId(v string) *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetTopics(v []*string) *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.Topics = v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetVSwitchId(v string) *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.VSwitchId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetVpcId(v string) *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.VpcId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters struct {
	// The download URL of the OSS resource ZIP package.
	//
	// example:
	//
	// "https://examplebucket.oss-cn-hangzhou.aliyuncs.com/testDoc/Old_Homebrew/2024-06-26%2022%3A34%3A08/opt/homebrew/homebrew/Library/Homebrew/test/support/fixtures/cask/AppWithBinary.zip?OSSAccessKeyId=ri&Expires=1725539627&Signature=rb8q3OpV2i3gZJ"
	ConnectorPackageUrl *string `json:"ConnectorPackageUrl,omitempty" xml:"ConnectorPackageUrl,omitempty"`
	// The connector parameters.
	ConnectorParameters *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters `json:"ConnectorParameters,omitempty" xml:"ConnectorParameters,omitempty" type:"Struct"`
	// The instance configuration.
	//
	// example:
	//
	// {
	//
	//         "group.id": "connect-eb-cluster-KAFKA_CONNECTORC",
	//
	//         "offset.storage.topic": "connect-eb-offset-KAFKA_CONNECTOR_yjqC8K5ewC",
	//
	//         "config.storage.topic": "connect-eb-config-KAFKA_CONNECTOR_yjqC8K5ewC",
	//
	//         "status.storage.topic": "connect-eb-status-KAFKA_CONNECTOR_yjqC8K5ewC",
	//
	//         "consumer.group.id": "connector-eb-cluster-KAFKA_CONNECTOR_yjqC8K5ewC-mongo-sink",
	//
	//         "bootstrap.servers": "alikafka-post:9092"
	//
	//       }
	WorkerParameters map[string]interface{} `json:"WorkerParameters,omitempty" xml:"WorkerParameters,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) GetConnectorPackageUrl() *string {
	return s.ConnectorPackageUrl
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) GetConnectorParameters() *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	return s.ConnectorParameters
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) GetWorkerParameters() map[string]interface{} {
	return s.WorkerParameters
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) SetConnectorPackageUrl(v string) *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters {
	s.ConnectorPackageUrl = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) SetConnectorParameters(v *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters {
	s.ConnectorParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) SetWorkerParameters(v map[string]interface{}) *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters {
	s.WorkerParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) Validate() error {
	if s.ConnectorParameters != nil {
		if err := s.ConnectorParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters struct {
	// The connector configuration.
	//
	// example:
	//
	// {
	//
	//           "connector.class": "com.mongodb.kafka.connect.MongoSinkConnector",
	//
	//           "tasks.max": "1",
	//
	//           "topics": "sourceA,sourceB"
	//
	//         }
	Config map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	// The connector name.
	//
	// example:
	//
	// test-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GetName() *string {
	return s.Name
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) SetConfig(v map[string]interface{}) *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	s.Config = v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) SetName(v string) *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	s.Name = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceCustomizedKafkaParameters struct {
	// The instance ID of the ApsaraMQ for Kafka instance.
	//
	// example:
	//
	// r-8vb64581862c****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceCustomizedKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceCustomizedKafkaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceCustomizedKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceCustomizedKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceDTSParameters struct {
	// The network address and port number of the data subscription channel.
	BrokerUrl *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	// The consumption offset, which is the timestamp when the SDK client consumes the first data record. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1620962769
	InitCheckPoint *int64 `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	// The password of the consumer group account.
	//
	// example:
	//
	// admin
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The consumer group ID.
	//
	// example:
	//
	// hkprdb
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The task ID.
	//
	// example:
	//
	// f86e5814-b223-482c-b768-3b873297****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The subscription topic of the data subscription channel.
	//
	// example:
	//
	// LTC_CACHE_PRD
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The account of the consumer group.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceDTSParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceDTSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) GetBrokerUrl() *string {
	return s.BrokerUrl
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) GetInitCheckPoint() *int64 {
	return s.InitCheckPoint
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) GetPassword() *string {
	return s.Password
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) GetSid() *string {
	return s.Sid
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) GetTopic() *string {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) GetUsername() *string {
	return s.Username
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetBrokerUrl(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.BrokerUrl = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetInitCheckPoint(v int64) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.InitCheckPoint = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetPassword(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.Password = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetSid(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.Sid = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetTaskId(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.TaskId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetTopic(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetUsername(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.Username = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceEventBusParameters struct {
	// The event bus name.
	//
	// example:
	//
	// my-event-bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event rule name.
	//
	// example:
	//
	// my-event-rule
	EventRuleName *string `json:"EventRuleName,omitempty" xml:"EventRuleName,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceEventBusParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceEventBusParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceEventBusParameters) GetEventBusName() *string {
	return s.EventBusName
}

func (s *UpdateEventStreamingRequestSourceSourceEventBusParameters) GetEventRuleName() *string {
	return s.EventRuleName
}

func (s *UpdateEventStreamingRequestSourceSourceEventBusParameters) SetEventBusName(v string) *UpdateEventStreamingRequestSourceSourceEventBusParameters {
	s.EventBusName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceEventBusParameters) SetEventRuleName(v string) *UpdateEventStreamingRequestSourceSourceEventBusParameters {
	s.EventRuleName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceEventBusParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceKafkaParameters struct {
	// The Group ID of the consumer that subscribes to the topic.
	//
	// example:
	//
	// DEFAULT_GROUP
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-8vbh4a5b9yfhgkkzm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network configuration. Default value: Default. Set to PublicNetwork for VPC networks.
	//
	// example:
	//
	// Default
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The offset.
	//
	// example:
	//
	// latest
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-uf6jcm3y5hcs7hkl****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The topic name.
	//
	// example:
	//
	// topic_empower_1641539400786
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-wz9t1l1e8eu2om****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The encoding and decoding method. Valid values:
	//
	// - Json: decodes bytes into a string by using UTF-8 encoding and parses the string into JSON format.
	//
	// - Text: decodes bytes into a string by using UTF-8 encoding and directly places the string into the payload.
	//
	// - Binary: encodes bytes into a string by using Base64 encoding and places the string into the payload.
	//
	// example:
	//
	// Text
	ValueDataType *string `json:"ValueDataType,omitempty" xml:"ValueDataType,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2ze6p0o345nykmekxt****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) GetNetwork() *string {
	return s.Network
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) GetValueDataType() *string {
	return s.ValueDataType
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetConsumerGroup(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetNetwork(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetOffsetReset(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetSecurityGroupId(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetTopic(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetVSwitchIds(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetValueDataType(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.ValueDataType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetVpcId(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceMNSParameters struct {
	// Specifies whether to enable Base64 decoding. Default value: true.
	//
	// example:
	//
	// true
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The queue name.
	//
	// example:
	//
	// queue_api_1642474203601
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceMNSParameters) GetIsBase64Decode() *bool {
	return s.IsBase64Decode
}

func (s *UpdateEventStreamingRequestSourceSourceMNSParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateEventStreamingRequestSourceSourceMNSParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventStreamingRequestSourceSourceMNSParameters) SetIsBase64Decode(v bool) *UpdateEventStreamingRequestSourceSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMNSParameters) SetQueueName(v string) *UpdateEventStreamingRequestSourceSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMNSParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceMNSParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMNSParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceMQTTParameters struct {
	// The message encoding format. Valid values:
	//
	// - JSON
	//
	// - Text
	//
	// - Binary
	//
	// example:
	//
	// JSON
	BodyDataType *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1dsudbecqwt61j****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The topic name.
	//
	// example:
	//
	// topic_empower_1642400400779
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The vSwitch ID.
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// VPC ID。
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceMQTTParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceMQTTParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) GetTopic() *string {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetBodyDataType(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.BodyDataType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetNetworkType(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.NetworkType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetSecurityGroupId(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetTopic(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetVSwitchIds(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.VSwitchIds = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetVpcId(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.VpcId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceOSSParameters struct {
	// The name of the bucket in Object Storage Service (OSS).
	//
	// example:
	//
	// bucket_abc
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The delimiter. In chunked loading mode, this delimiter is used as the text chunking identifier. The default delimiter is the newline character
	//
	// .
	//
	// example:
	//
	// \\n
	Delimiter *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	// The document loader.
	//
	// example:
	//
	// TextLoader
	LoadFormat *string `json:"LoadFormat,omitempty" xml:"LoadFormat,omitempty"`
	// The data loading mode. Valid values: single (single document loading) and element (chunked loading). Default value: single.
	//
	// example:
	//
	// single
	LoadMode *string `json:"LoadMode,omitempty" xml:"LoadMode,omitempty"`
	// The file path prefix.
	//
	// example:
	//
	// fun/document/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The role name used for authorization to allow the event bus EventBridge to read OSS files. The role must have at least read-only permissions on OSS.
	//
	// example:
	//
	// eventbridge_oss_role
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceOSSParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceOSSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) GetBucketName() *string {
	return s.BucketName
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) GetDelimiter() *string {
	return s.Delimiter
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) GetLoadFormat() *string {
	return s.LoadFormat
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) GetLoadMode() *string {
	return s.LoadMode
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) GetPrefix() *string {
	return s.Prefix
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) SetBucketName(v string) *UpdateEventStreamingRequestSourceSourceOSSParameters {
	s.BucketName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) SetDelimiter(v string) *UpdateEventStreamingRequestSourceSourceOSSParameters {
	s.Delimiter = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) SetLoadFormat(v string) *UpdateEventStreamingRequestSourceSourceOSSParameters {
	s.LoadFormat = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) SetLoadMode(v string) *UpdateEventStreamingRequestSourceSourceOSSParameters {
	s.LoadMode = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) SetPrefix(v string) *UpdateEventStreamingRequestSourceSourceOSSParameters {
	s.Prefix = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) SetRoleName(v string) *UpdateEventStreamingRequestSourceSourceOSSParameters {
	s.RoleName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOSSParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters struct {
	// The authentication type.
	//
	// example:
	//
	// ACL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The message body data type.
	//
	// example:
	//
	// Json
	BodyDataType *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	// The instance endpoint.
	//
	// example:
	//
	// 192.168.1.1:9876
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The network type. Valid values:
	//
	// - PublicNetwork
	//
	// - PrivateNetwork
	//
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The password used to connect to the open source RabbitMQ instance.
	//
	// example:
	//
	// ****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The queue name of the open source RabbitMQ instance.
	//
	// example:
	//
	// demo
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-m5edtu24f12345****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The username used to connect to the open source RabbitMQ instance.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-m5ev8asdc6h12345****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The virtual host name of the open source RabbitMQ instance.
	//
	// example:
	//
	// Vhost1
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-m5e3sv4b12345****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetPassword() *string {
	return s.Password
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetUsername() *string {
	return s.Username
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetAuthType(v string) *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.AuthType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetBodyDataType(v string) *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.BodyDataType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetEndpoint(v string) *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.Endpoint = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetNetworkType(v string) *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.NetworkType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetPassword(v string) *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.Password = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetQueueName(v string) *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetSecurityGroupId(v string) *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetUsername(v string) *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.Username = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetVSwitchIds(v string) *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetVirtualHostName(v string) *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetVpcId(v string) *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.VpcId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourcePrometheusParameters struct {
	// The cluster ID.
	//
	// example:
	//
	// c83555068b6******ad213f565f209
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The data type.
	//
	// example:
	//
	// Json
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The external labels appended to the event stream.
	//
	// example:
	//
	// {"env":"test"}
	ExternalLabels *string `json:"ExternalLabels,omitempty" xml:"ExternalLabels,omitempty"`
	// The labels.
	//
	// example:
	//
	// __name__=.*
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task role name.
	//
	// example:
	//
	// test-role
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourcePrometheusParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourcePrometheusParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) GetDataType() *string {
	return s.DataType
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) GetExternalLabels() *string {
	return s.ExternalLabels
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) GetLabels() *string {
	return s.Labels
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) SetClusterId(v string) *UpdateEventStreamingRequestSourceSourcePrometheusParameters {
	s.ClusterId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) SetDataType(v string) *UpdateEventStreamingRequestSourceSourcePrometheusParameters {
	s.DataType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) SetExternalLabels(v string) *UpdateEventStreamingRequestSourceSourcePrometheusParameters {
	s.ExternalLabels = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) SetLabels(v string) *UpdateEventStreamingRequestSourceSourcePrometheusParameters {
	s.Labels = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourcePrometheusParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) SetRoleName(v string) *UpdateEventStreamingRequestSourceSourcePrometheusParameters {
	s.RoleName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceRabbitMQParameters struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// i-f8z9lqkldlb4oxsxwwub
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// demo
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the vhost of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// eb-connect
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) SetQueueName(v string) *UpdateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) SetVirtualHostName(v string) *UpdateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters struct {
	// The instance ID of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// rmq-cn-jte3w******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// Cloud_5
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic of the ApsaraMQ for RocketMQ instance.
	Topics []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
}

func (s UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) GetTopics() []*string {
	return s.Topics
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) SetInstanceType(v string) *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters {
	s.InstanceType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) SetTopics(v []*string) *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters {
	s.Topics = v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceRocketMQParameters struct {
	// The authentication type.
	//
	// example:
	//
	// ACL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The message encoding format. Valid values:
	//
	// - Json
	//
	// - Text
	//
	// - Binary
	//
	// example:
	//
	// Json
	BodyDataType *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	// The SQL filter statement.
	//
	// example:
	//
	// index > 10
	FilterSql *string `json:"FilterSql,omitempty" xml:"FilterSql,omitempty"`
	// The message filter type.
	//
	// example:
	//
	// Tag
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// The group ID of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// GID_test
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// The instance endpoint.
	//
	// example:
	//
	// reg****-vpc.cn-zhangjiakou.aliyuncs.com
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The instance ID of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// i-f8z9a9mcgwri1c1id****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network information of the instance. Valid values:
	//
	// - PublicNetwork
	//
	// - PrivateNetwork
	//
	// example:
	//
	// PublicNetwork
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The instance password.
	//
	// example:
	//
	// admin
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The security group information of the instance.
	//
	// example:
	//
	// sg-m5edtu24f12345****
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// The instance type. Valid values:
	//
	// - Cloud_4 (default): Alibaba Cloud RocketMQ 4.0 instance
	//
	// - Cloud_5: Alibaba Cloud RocketMQ 5.0 instance
	//
	// - SelfBuilt: self-managed Apache RocketMQ instance
	//
	// example:
	//
	// Cloud_5
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance username.
	//
	// example:
	//
	// admin
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The vSwitch information of the instance.
	//
	// example:
	//
	// vsw-m5ev8asdc6h12****
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The VPC information of the instance.
	//
	// example:
	//
	// vpc-m5e3sv4b12345****
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The network type. Valid values:
	//
	// - PublicNetwork
	//
	// - PrivateNetwork
	//
	// example:
	//
	// PublicNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The consumption offset of the message. Valid values:
	//
	// - CONSUMEFROMLASTOFFSET: Consumption starts from the latest offset.
	//
	// - CONSUMEFROMFIRSTOFFSET: Consumption starts from the earliest offset.
	//
	// - CONSUMEFROMTIMESTAMP: Consumption starts from the offset at the specified time.
	//
	// Default value: CONSUMEFROMLAST_OFFSET.
	//
	// example:
	//
	// CONSUMEFROMLASTOFFSET
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The region ID of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-m5edtu24f12345****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The filter tag of the message.
	//
	// example:
	//
	// test
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp. This parameter is valid only when the Offset parameter is set to CONSUMEFROMTIMESTAMP.
	//
	// example:
	//
	// 1670656652009
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The topic name.
	//
	// example:
	//
	// TOPIC-cainiao-pcs-order-process-inBoundConditionCheck
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-m5ev8asdc6h12345****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The VPC ID of the instance.
	//
	// example:
	//
	// vpc-m5e3sv4b12345****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetFilterSql() *string {
	return s.FilterSql
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetFilterType() *string {
	return s.FilterType
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetGroupID() *string {
	return s.GroupID
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceNetwork() *string {
	return s.InstanceNetwork
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceSecurityGroupId() *string {
	return s.InstanceSecurityGroupId
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceVSwitchIds() *string {
	return s.InstanceVSwitchIds
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceVpcId() *string {
	return s.InstanceVpcId
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetNetwork() *string {
	return s.Network
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetOffset() *string {
	return s.Offset
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetTag() *string {
	return s.Tag
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetTopic() *string {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetAuthType(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetBodyDataType(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.BodyDataType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetFilterSql(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.FilterSql = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetFilterType(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.FilterType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetGroupID(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceEndpoint(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceNetwork(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstancePassword(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceType(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceUsername(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceVpcId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetNetwork(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Network = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetOffset(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetSecurityGroupId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetTag(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetTimestamp(v int64) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetTopic(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetVSwitchIds(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetVpcId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.VpcId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceSLSParameters struct {
	// The role name used for authorization to allow the event bus EventBridge to read Simple Log Service log content. When you create the role in the Resource Access Management (RAM) console, select "Alibaba Cloud Service" and set "Trusted Service" to "event bus".
	//
	// example:
	//
	// testRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceSLSParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceSLSParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateEventStreamingRequestSourceSourceSLSParameters) SetRoleName(v string) *UpdateEventStreamingRequestSourceSourceSLSParameters {
	s.RoleName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceSLSParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestTransforms struct {
	// The ARN of the cloud product, such as the ARN of a function in Function Compute.
	//
	// example:
	//
	// acs:fc:cn-hangzhou:*****:services/demo-service.LATEST/functions/demo-func
	Arn                             *string                          `json:"Arn,omitempty" xml:"Arn,omitempty"`
	BaiLianAgentTransformParameters *BaiLianAgentTransformParameters `json:"BaiLianAgentTransformParameters,omitempty" xml:"BaiLianAgentTransformParameters,omitempty"`
	DashScopeTransformParameters    *DashScopeTransformParameters    `json:"DashScopeTransformParameters,omitempty" xml:"DashScopeTransformParameters,omitempty"`
}

func (s UpdateEventStreamingRequestTransforms) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestTransforms) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestTransforms) GetArn() *string {
	return s.Arn
}

func (s *UpdateEventStreamingRequestTransforms) GetBaiLianAgentTransformParameters() *BaiLianAgentTransformParameters {
	return s.BaiLianAgentTransformParameters
}

func (s *UpdateEventStreamingRequestTransforms) GetDashScopeTransformParameters() *DashScopeTransformParameters {
	return s.DashScopeTransformParameters
}

func (s *UpdateEventStreamingRequestTransforms) SetArn(v string) *UpdateEventStreamingRequestTransforms {
	s.Arn = &v
	return s
}

func (s *UpdateEventStreamingRequestTransforms) SetBaiLianAgentTransformParameters(v *BaiLianAgentTransformParameters) *UpdateEventStreamingRequestTransforms {
	s.BaiLianAgentTransformParameters = v
	return s
}

func (s *UpdateEventStreamingRequestTransforms) SetDashScopeTransformParameters(v *DashScopeTransformParameters) *UpdateEventStreamingRequestTransforms {
	s.DashScopeTransformParameters = v
	return s
}

func (s *UpdateEventStreamingRequestTransforms) Validate() error {
	if s.BaiLianAgentTransformParameters != nil {
		if err := s.BaiLianAgentTransformParameters.Validate(); err != nil {
			return err
		}
	}
	if s.DashScopeTransformParameters != nil {
		if err := s.DashScopeTransformParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}
