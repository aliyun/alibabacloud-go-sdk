// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventStreamingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateEventStreamingRequest
	GetDescription() *string
	SetEventStreamingName(v string) *CreateEventStreamingRequest
	GetEventStreamingName() *string
	SetFilterPattern(v string) *CreateEventStreamingRequest
	GetFilterPattern() *string
	SetRunOptions(v *CreateEventStreamingRequestRunOptions) *CreateEventStreamingRequest
	GetRunOptions() *CreateEventStreamingRequestRunOptions
	SetSink(v *CreateEventStreamingRequestSink) *CreateEventStreamingRequest
	GetSink() *CreateEventStreamingRequestSink
	SetSource(v *CreateEventStreamingRequestSource) *CreateEventStreamingRequest
	GetSource() *CreateEventStreamingRequestSource
	SetTags(v []*CreateEventStreamingRequestTags) *CreateEventStreamingRequest
	GetTags() []*CreateEventStreamingRequestTags
	SetTransforms(v []*CreateEventStreamingRequestTransforms) *CreateEventStreamingRequest
	GetTransforms() []*CreateEventStreamingRequestTransforms
}

type CreateEventStreamingRequest struct {
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
	// The rule that is used to filter events. If you leave this parameter empty, all events are matched.
	//
	// This parameter is required.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The parameters that are configured for the runtime environment.
	RunOptions *CreateEventStreamingRequestRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	// The event target. You must and can specify only one event target.
	//
	// This parameter is required.
	Sink *CreateEventStreamingRequestSink `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	// The event provider, which is also known as the event source. You must and can specify only one event source.
	//
	// This parameter is required.
	Source     *CreateEventStreamingRequestSource       `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Tags       []*CreateEventStreamingRequestTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Transforms []*CreateEventStreamingRequestTransforms `json:"Transforms,omitempty" xml:"Transforms,omitempty" type:"Repeated"`
}

func (s CreateEventStreamingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEventStreamingRequest) GetEventStreamingName() *string {
	return s.EventStreamingName
}

func (s *CreateEventStreamingRequest) GetFilterPattern() *string {
	return s.FilterPattern
}

func (s *CreateEventStreamingRequest) GetRunOptions() *CreateEventStreamingRequestRunOptions {
	return s.RunOptions
}

func (s *CreateEventStreamingRequest) GetSink() *CreateEventStreamingRequestSink {
	return s.Sink
}

func (s *CreateEventStreamingRequest) GetSource() *CreateEventStreamingRequestSource {
	return s.Source
}

func (s *CreateEventStreamingRequest) GetTags() []*CreateEventStreamingRequestTags {
	return s.Tags
}

func (s *CreateEventStreamingRequest) GetTransforms() []*CreateEventStreamingRequestTransforms {
	return s.Transforms
}

func (s *CreateEventStreamingRequest) SetDescription(v string) *CreateEventStreamingRequest {
	s.Description = &v
	return s
}

func (s *CreateEventStreamingRequest) SetEventStreamingName(v string) *CreateEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

func (s *CreateEventStreamingRequest) SetFilterPattern(v string) *CreateEventStreamingRequest {
	s.FilterPattern = &v
	return s
}

func (s *CreateEventStreamingRequest) SetRunOptions(v *CreateEventStreamingRequestRunOptions) *CreateEventStreamingRequest {
	s.RunOptions = v
	return s
}

func (s *CreateEventStreamingRequest) SetSink(v *CreateEventStreamingRequestSink) *CreateEventStreamingRequest {
	s.Sink = v
	return s
}

func (s *CreateEventStreamingRequest) SetSource(v *CreateEventStreamingRequestSource) *CreateEventStreamingRequest {
	s.Source = v
	return s
}

func (s *CreateEventStreamingRequest) SetTags(v []*CreateEventStreamingRequestTags) *CreateEventStreamingRequest {
	s.Tags = v
	return s
}

func (s *CreateEventStreamingRequest) SetTransforms(v []*CreateEventStreamingRequestTransforms) *CreateEventStreamingRequest {
	s.Transforms = v
	return s
}

func (s *CreateEventStreamingRequest) Validate() error {
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type CreateEventStreamingRequestRunOptions struct {
	// The batch window.
	BatchWindow    *CreateEventStreamingRequestRunOptionsBatchWindow    `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	BusinessOption *CreateEventStreamingRequestRunOptionsBusinessOption `json:"BusinessOption,omitempty" xml:"BusinessOption,omitempty" type:"Struct"`
	// Specifies whether to enable dead-letter queues. By default, dead-letter queues are disabled. Messages that fail to be pushed after the allowed retries as specified by the retry policy are discarded.
	DeadLetterQueue *CreateEventStreamingRequestRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The exception tolerance policy. Valid values:
	//
	// 	- NONE: does not tolerate exceptions.
	//
	// 	- ALL: tolerates all exceptions.
	//
	// example:
	//
	// ALL
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The maximum number of concurrent tasks.
	//
	// example:
	//
	// 2
	MaximumTasks *int64 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The retry policy that you want to use if events fail to be pushed.
	RetryStrategy *CreateEventStreamingRequestRunOptionsRetryStrategy `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
	Throttling    *int32                                              `json:"Throttling,omitempty" xml:"Throttling,omitempty"`
}

func (s CreateEventStreamingRequestRunOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestRunOptions) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestRunOptions) GetBatchWindow() *CreateEventStreamingRequestRunOptionsBatchWindow {
	return s.BatchWindow
}

func (s *CreateEventStreamingRequestRunOptions) GetBusinessOption() *CreateEventStreamingRequestRunOptionsBusinessOption {
	return s.BusinessOption
}

func (s *CreateEventStreamingRequestRunOptions) GetDeadLetterQueue() *CreateEventStreamingRequestRunOptionsDeadLetterQueue {
	return s.DeadLetterQueue
}

func (s *CreateEventStreamingRequestRunOptions) GetErrorsTolerance() *string {
	return s.ErrorsTolerance
}

func (s *CreateEventStreamingRequestRunOptions) GetMaximumTasks() *int64 {
	return s.MaximumTasks
}

func (s *CreateEventStreamingRequestRunOptions) GetRetryStrategy() *CreateEventStreamingRequestRunOptionsRetryStrategy {
	return s.RetryStrategy
}

func (s *CreateEventStreamingRequestRunOptions) GetThrottling() *int32 {
	return s.Throttling
}

func (s *CreateEventStreamingRequestRunOptions) SetBatchWindow(v *CreateEventStreamingRequestRunOptionsBatchWindow) *CreateEventStreamingRequestRunOptions {
	s.BatchWindow = v
	return s
}

func (s *CreateEventStreamingRequestRunOptions) SetBusinessOption(v *CreateEventStreamingRequestRunOptionsBusinessOption) *CreateEventStreamingRequestRunOptions {
	s.BusinessOption = v
	return s
}

func (s *CreateEventStreamingRequestRunOptions) SetDeadLetterQueue(v *CreateEventStreamingRequestRunOptionsDeadLetterQueue) *CreateEventStreamingRequestRunOptions {
	s.DeadLetterQueue = v
	return s
}

func (s *CreateEventStreamingRequestRunOptions) SetErrorsTolerance(v string) *CreateEventStreamingRequestRunOptions {
	s.ErrorsTolerance = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptions) SetMaximumTasks(v int64) *CreateEventStreamingRequestRunOptions {
	s.MaximumTasks = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptions) SetRetryStrategy(v *CreateEventStreamingRequestRunOptionsRetryStrategy) *CreateEventStreamingRequestRunOptions {
	s.RetryStrategy = v
	return s
}

func (s *CreateEventStreamingRequestRunOptions) SetThrottling(v int32) *CreateEventStreamingRequestRunOptions {
	s.Throttling = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptions) Validate() error {
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

type CreateEventStreamingRequestRunOptionsBatchWindow struct {
	// The maximum number of events that are allowed in the batch window. When this threshold is reached, data in the window is pushed to the downstream service. If multiple batch windows exist, data is pushed if the triggering conditions are met in one of the windows.
	//
	// example:
	//
	// 100
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	// The maximum period of time during which events are allowed in the batch window. Unit: seconds. When this threshold is reached, data in the window is pushed to the downstream service. If multiple batch windows exist, data is pushed if the triggering conditions are met in one of the windows.
	//
	// example:
	//
	// 10
	TimeBasedWindow *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
}

func (s CreateEventStreamingRequestRunOptionsBatchWindow) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestRunOptionsBatchWindow) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestRunOptionsBatchWindow) GetCountBasedWindow() *int32 {
	return s.CountBasedWindow
}

func (s *CreateEventStreamingRequestRunOptionsBatchWindow) GetTimeBasedWindow() *int32 {
	return s.TimeBasedWindow
}

func (s *CreateEventStreamingRequestRunOptionsBatchWindow) SetCountBasedWindow(v int32) *CreateEventStreamingRequestRunOptionsBatchWindow {
	s.CountBasedWindow = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsBatchWindow) SetTimeBasedWindow(v int32) *CreateEventStreamingRequestRunOptionsBatchWindow {
	s.TimeBasedWindow = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsBatchWindow) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestRunOptionsBusinessOption struct {
	BusinessMode         *string `json:"BusinessMode,omitempty" xml:"BusinessMode,omitempty"`
	MaxCapacityUnitCount *int64  `json:"MaxCapacityUnitCount,omitempty" xml:"MaxCapacityUnitCount,omitempty"`
	MinCapacityUnitCount *int64  `json:"MinCapacityUnitCount,omitempty" xml:"MinCapacityUnitCount,omitempty"`
}

func (s CreateEventStreamingRequestRunOptionsBusinessOption) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestRunOptionsBusinessOption) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestRunOptionsBusinessOption) GetBusinessMode() *string {
	return s.BusinessMode
}

func (s *CreateEventStreamingRequestRunOptionsBusinessOption) GetMaxCapacityUnitCount() *int64 {
	return s.MaxCapacityUnitCount
}

func (s *CreateEventStreamingRequestRunOptionsBusinessOption) GetMinCapacityUnitCount() *int64 {
	return s.MinCapacityUnitCount
}

func (s *CreateEventStreamingRequestRunOptionsBusinessOption) SetBusinessMode(v string) *CreateEventStreamingRequestRunOptionsBusinessOption {
	s.BusinessMode = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsBusinessOption) SetMaxCapacityUnitCount(v int64) *CreateEventStreamingRequestRunOptionsBusinessOption {
	s.MaxCapacityUnitCount = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsBusinessOption) SetMinCapacityUnitCount(v int64) *CreateEventStreamingRequestRunOptionsBusinessOption {
	s.MinCapacityUnitCount = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsBusinessOption) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestRunOptionsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue.
	//
	// example:
	//
	// acs:ram::1317334647812936:role/rdstoecsassumekms
	Arn             *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventStreamingRequestRunOptionsDeadLetterQueue) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestRunOptionsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) GetArn() *string {
	return s.Arn
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) GetNetwork() *string {
	return s.Network
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) SetArn(v string) *CreateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.Arn = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) SetNetwork(v string) *CreateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.Network = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) SetSecurityGroupId(v string) *CreateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) SetVSwitchIds(v string) *CreateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.VSwitchIds = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) SetVpcId(v string) *CreateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.VpcId = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestRunOptionsRetryStrategy struct {
	// The maximum timeout period for a retry.
	//
	// example:
	//
	// 512
	MaximumEventAgeInSeconds *int64 `json:"MaximumEventAgeInSeconds,omitempty" xml:"MaximumEventAgeInSeconds,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 2
	MaximumRetryAttempts *int64 `json:"MaximumRetryAttempts,omitempty" xml:"MaximumRetryAttempts,omitempty"`
	// The retry policy. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFFRETRY
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
}

func (s CreateEventStreamingRequestRunOptionsRetryStrategy) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestRunOptionsRetryStrategy) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestRunOptionsRetryStrategy) GetMaximumEventAgeInSeconds() *int64 {
	return s.MaximumEventAgeInSeconds
}

func (s *CreateEventStreamingRequestRunOptionsRetryStrategy) GetMaximumRetryAttempts() *int64 {
	return s.MaximumRetryAttempts
}

func (s *CreateEventStreamingRequestRunOptionsRetryStrategy) GetPushRetryStrategy() *string {
	return s.PushRetryStrategy
}

func (s *CreateEventStreamingRequestRunOptionsRetryStrategy) SetMaximumEventAgeInSeconds(v int64) *CreateEventStreamingRequestRunOptionsRetryStrategy {
	s.MaximumEventAgeInSeconds = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsRetryStrategy) SetMaximumRetryAttempts(v int64) *CreateEventStreamingRequestRunOptionsRetryStrategy {
	s.MaximumRetryAttempts = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsRetryStrategy) SetPushRetryStrategy(v string) *CreateEventStreamingRequestRunOptionsRetryStrategy {
	s.PushRetryStrategy = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsRetryStrategy) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSink struct {
	SinkApacheKafkaParameters              *CreateEventStreamingRequestSinkSinkApacheKafkaParameters              `json:"SinkApacheKafkaParameters,omitempty" xml:"SinkApacheKafkaParameters,omitempty" type:"Struct"`
	SinkApacheRocketMQCheckpointParameters *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters `json:"SinkApacheRocketMQCheckpointParameters,omitempty" xml:"SinkApacheRocketMQCheckpointParameters,omitempty" type:"Struct"`
	SinkApiDestinationParameters           *SinkApiDestinationParameters                                          `json:"SinkApiDestinationParameters,omitempty" xml:"SinkApiDestinationParameters,omitempty"`
	SinkBaiLianParameters                  *SinkBaiLianParameters                                                 `json:"SinkBaiLianParameters,omitempty" xml:"SinkBaiLianParameters,omitempty"`
	SinkCustomizedKafkaConnectorParameters *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters `json:"SinkCustomizedKafkaConnectorParameters,omitempty" xml:"SinkCustomizedKafkaConnectorParameters,omitempty" type:"Struct"`
	SinkCustomizedKafkaParameters          *CreateEventStreamingRequestSinkSinkCustomizedKafkaParameters          `json:"SinkCustomizedKafkaParameters,omitempty" xml:"SinkCustomizedKafkaParameters,omitempty" type:"Struct"`
	SinkDashVectorParameters               *CreateEventStreamingRequestSinkSinkDashVectorParameters               `json:"SinkDashVectorParameters,omitempty" xml:"SinkDashVectorParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify DataHub as the event target.
	SinkDataHubParameters          *CreateEventStreamingRequestSinkSinkDataHubParameters `json:"SinkDataHubParameters,omitempty" xml:"SinkDataHubParameters,omitempty" type:"Struct"`
	SinkDataWorksTriggerParameters *SinkDataWorksTriggerParameters                       `json:"SinkDataWorksTriggerParameters,omitempty" xml:"SinkDataWorksTriggerParameters,omitempty"`
	SinkDorisParameters            *CreateEventStreamingRequestSinkSinkDorisParameters   `json:"SinkDorisParameters,omitempty" xml:"SinkDorisParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Function Compute as the event target.
	SinkFcParameters *CreateEventStreamingRequestSinkSinkFcParameters `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify CloudFlow as the event target.
	SinkFnfParameters   *CreateEventStreamingRequestSinkSinkFnfParameters `json:"SinkFnfParameters,omitempty" xml:"SinkFnfParameters,omitempty" type:"Struct"`
	SinkHttpsParameters *SinkHttpsParameters                              `json:"SinkHttpsParameters,omitempty" xml:"SinkHttpsParameters,omitempty"`
	// The parameters that are configured if you specify ApsaraMQ for Kafka as the event target.
	SinkKafkaParameters *CreateEventStreamingRequestSinkSinkKafkaParameters `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify MNS as the event target.
	SinkMNSParameters                *CreateEventStreamingRequestSinkSinkMNSParameters                `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
	SinkOSSParameters                *SinkOSSParameters                                               `json:"SinkOSSParameters,omitempty" xml:"SinkOSSParameters,omitempty"`
	SinkOpenSourceRabbitMQParameters *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters `json:"SinkOpenSourceRabbitMQParameters,omitempty" xml:"SinkOpenSourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Managed Service for Prometheus as the event target.
	SinkPrometheusParameters      *CreateEventStreamingRequestSinkSinkPrometheusParameters `json:"SinkPrometheusParameters,omitempty" xml:"SinkPrometheusParameters,omitempty" type:"Struct"`
	SinkRabbitMQMetaParameters    *SinkRabbitMQMetaParameters                              `json:"SinkRabbitMQMetaParameters,omitempty" xml:"SinkRabbitMQMetaParameters,omitempty"`
	SinkRabbitMQMsgSyncParameters *SinkRabbitMQMsgSyncParameters                           `json:"SinkRabbitMQMsgSyncParameters,omitempty" xml:"SinkRabbitMQMsgSyncParameters,omitempty"`
	// The parameters that are configured if you specify ApsaraMQ for RabbitMQ as the event target.
	SinkRabbitMQParameters           *CreateEventStreamingRequestSinkSinkRabbitMQParameters           `json:"SinkRabbitMQParameters,omitempty" xml:"SinkRabbitMQParameters,omitempty" type:"Struct"`
	SinkRocketMQCheckpointParameters *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters `json:"SinkRocketMQCheckpointParameters,omitempty" xml:"SinkRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for RocketMQ as the event target.
	SinkRocketMQParameters *CreateEventStreamingRequestSinkSinkRocketMQParameters `json:"SinkRocketMQParameters,omitempty" xml:"SinkRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Simple Log Service as the event target.
	SinkSLSParameters *CreateEventStreamingRequestSinkSinkSLSParameters `json:"SinkSLSParameters,omitempty" xml:"SinkSLSParameters,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSink) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSink) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSink) GetSinkApacheKafkaParameters() *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	return s.SinkApacheKafkaParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkApacheRocketMQCheckpointParameters() *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	return s.SinkApacheRocketMQCheckpointParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkApiDestinationParameters() *SinkApiDestinationParameters {
	return s.SinkApiDestinationParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkBaiLianParameters() *SinkBaiLianParameters {
	return s.SinkBaiLianParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkCustomizedKafkaConnectorParameters() *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters {
	return s.SinkCustomizedKafkaConnectorParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkCustomizedKafkaParameters() *CreateEventStreamingRequestSinkSinkCustomizedKafkaParameters {
	return s.SinkCustomizedKafkaParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkDashVectorParameters() *CreateEventStreamingRequestSinkSinkDashVectorParameters {
	return s.SinkDashVectorParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkDataHubParameters() *CreateEventStreamingRequestSinkSinkDataHubParameters {
	return s.SinkDataHubParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkDataWorksTriggerParameters() *SinkDataWorksTriggerParameters {
	return s.SinkDataWorksTriggerParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkDorisParameters() *CreateEventStreamingRequestSinkSinkDorisParameters {
	return s.SinkDorisParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkFcParameters() *CreateEventStreamingRequestSinkSinkFcParameters {
	return s.SinkFcParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkFnfParameters() *CreateEventStreamingRequestSinkSinkFnfParameters {
	return s.SinkFnfParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkHttpsParameters() *SinkHttpsParameters {
	return s.SinkHttpsParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkKafkaParameters() *CreateEventStreamingRequestSinkSinkKafkaParameters {
	return s.SinkKafkaParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkMNSParameters() *CreateEventStreamingRequestSinkSinkMNSParameters {
	return s.SinkMNSParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkOSSParameters() *SinkOSSParameters {
	return s.SinkOSSParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkOpenSourceRabbitMQParameters() *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	return s.SinkOpenSourceRabbitMQParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkPrometheusParameters() *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	return s.SinkPrometheusParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkRabbitMQMetaParameters() *SinkRabbitMQMetaParameters {
	return s.SinkRabbitMQMetaParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkRabbitMQMsgSyncParameters() *SinkRabbitMQMsgSyncParameters {
	return s.SinkRabbitMQMsgSyncParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkRabbitMQParameters() *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	return s.SinkRabbitMQParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkRocketMQCheckpointParameters() *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	return s.SinkRocketMQCheckpointParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkRocketMQParameters() *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	return s.SinkRocketMQParameters
}

func (s *CreateEventStreamingRequestSink) GetSinkSLSParameters() *CreateEventStreamingRequestSinkSinkSLSParameters {
	return s.SinkSLSParameters
}

func (s *CreateEventStreamingRequestSink) SetSinkApacheKafkaParameters(v *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) *CreateEventStreamingRequestSink {
	s.SinkApacheKafkaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkApacheRocketMQCheckpointParameters(v *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) *CreateEventStreamingRequestSink {
	s.SinkApacheRocketMQCheckpointParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkApiDestinationParameters(v *SinkApiDestinationParameters) *CreateEventStreamingRequestSink {
	s.SinkApiDestinationParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkBaiLianParameters(v *SinkBaiLianParameters) *CreateEventStreamingRequestSink {
	s.SinkBaiLianParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkCustomizedKafkaConnectorParameters(v *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) *CreateEventStreamingRequestSink {
	s.SinkCustomizedKafkaConnectorParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkCustomizedKafkaParameters(v *CreateEventStreamingRequestSinkSinkCustomizedKafkaParameters) *CreateEventStreamingRequestSink {
	s.SinkCustomizedKafkaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkDashVectorParameters(v *CreateEventStreamingRequestSinkSinkDashVectorParameters) *CreateEventStreamingRequestSink {
	s.SinkDashVectorParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkDataHubParameters(v *CreateEventStreamingRequestSinkSinkDataHubParameters) *CreateEventStreamingRequestSink {
	s.SinkDataHubParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkDataWorksTriggerParameters(v *SinkDataWorksTriggerParameters) *CreateEventStreamingRequestSink {
	s.SinkDataWorksTriggerParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkDorisParameters(v *CreateEventStreamingRequestSinkSinkDorisParameters) *CreateEventStreamingRequestSink {
	s.SinkDorisParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkFcParameters(v *CreateEventStreamingRequestSinkSinkFcParameters) *CreateEventStreamingRequestSink {
	s.SinkFcParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkFnfParameters(v *CreateEventStreamingRequestSinkSinkFnfParameters) *CreateEventStreamingRequestSink {
	s.SinkFnfParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkHttpsParameters(v *SinkHttpsParameters) *CreateEventStreamingRequestSink {
	s.SinkHttpsParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkKafkaParameters(v *CreateEventStreamingRequestSinkSinkKafkaParameters) *CreateEventStreamingRequestSink {
	s.SinkKafkaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkMNSParameters(v *CreateEventStreamingRequestSinkSinkMNSParameters) *CreateEventStreamingRequestSink {
	s.SinkMNSParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkOSSParameters(v *SinkOSSParameters) *CreateEventStreamingRequestSink {
	s.SinkOSSParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkOpenSourceRabbitMQParameters(v *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) *CreateEventStreamingRequestSink {
	s.SinkOpenSourceRabbitMQParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkPrometheusParameters(v *CreateEventStreamingRequestSinkSinkPrometheusParameters) *CreateEventStreamingRequestSink {
	s.SinkPrometheusParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkRabbitMQMetaParameters(v *SinkRabbitMQMetaParameters) *CreateEventStreamingRequestSink {
	s.SinkRabbitMQMetaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkRabbitMQMsgSyncParameters(v *SinkRabbitMQMsgSyncParameters) *CreateEventStreamingRequestSink {
	s.SinkRabbitMQMsgSyncParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkRabbitMQParameters(v *CreateEventStreamingRequestSinkSinkRabbitMQParameters) *CreateEventStreamingRequestSink {
	s.SinkRabbitMQParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkRocketMQCheckpointParameters(v *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) *CreateEventStreamingRequestSink {
	s.SinkRocketMQCheckpointParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkRocketMQParameters(v *CreateEventStreamingRequestSinkSinkRocketMQParameters) *CreateEventStreamingRequestSink {
	s.SinkRocketMQParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkSLSParameters(v *CreateEventStreamingRequestSinkSinkSLSParameters) *CreateEventStreamingRequestSink {
	s.SinkSLSParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) Validate() error {
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

type CreateEventStreamingRequestSinkSinkApacheKafkaParameters struct {
	Acks                      *string                                                                  `json:"Acks,omitempty" xml:"Acks,omitempty"`
	Bootstraps                *string                                                                  `json:"Bootstraps,omitempty" xml:"Bootstraps,omitempty"`
	CompressionType           *string                                                                  `json:"CompressionType,omitempty" xml:"CompressionType,omitempty"`
	Headers                   *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders         `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	Key                       *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey             `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	NetworkType               *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType     `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	SaslMechanism             *string                                                                  `json:"SaslMechanism,omitempty" xml:"SaslMechanism,omitempty"`
	SaslPassword              *string                                                                  `json:"SaslPassword,omitempty" xml:"SaslPassword,omitempty"`
	SaslUser                  *string                                                                  `json:"SaslUser,omitempty" xml:"SaslUser,omitempty"`
	SecurityGroupId           *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	SecurityProtocol          *string                                                                  `json:"SecurityProtocol,omitempty" xml:"SecurityProtocol,omitempty"`
	SslTruststoreCertificates *string                                                                  `json:"SslTruststoreCertificates,omitempty" xml:"SslTruststoreCertificates,omitempty"`
	Topic                     *string                                                                  `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds                *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds      `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	Value                     *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue           `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
	VpcId                     *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId           `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetAcks() *string {
	return s.Acks
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetBootstraps() *string {
	return s.Bootstraps
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetCompressionType() *string {
	return s.CompressionType
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetHeaders() *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders {
	return s.Headers
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetKey() *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey {
	return s.Key
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetNetworkType() *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType {
	return s.NetworkType
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSaslMechanism() *string {
	return s.SaslMechanism
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSaslPassword() *string {
	return s.SaslPassword
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSaslUser() *string {
	return s.SaslUser
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSecurityGroupId() *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSecurityProtocol() *string {
	return s.SecurityProtocol
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetSslTruststoreCertificates() *string {
	return s.SslTruststoreCertificates
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetVSwitchIds() *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetValue() *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) GetVpcId() *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId {
	return s.VpcId
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetAcks(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Acks = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetBootstraps(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Bootstraps = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetCompressionType(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.CompressionType = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetHeaders(v *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Headers = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetKey(v *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Key = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetNetworkType(v *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.NetworkType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSaslMechanism(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SaslMechanism = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSaslPassword(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SaslPassword = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSaslUser(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SaslUser = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSecurityGroupId(v *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SecurityGroupId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSecurityProtocol(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SecurityProtocol = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetSslTruststoreCertificates(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.SslTruststoreCertificates = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetTopic(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetVSwitchIds(v *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.VSwitchIds = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetValue(v *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.Value = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) SetVpcId(v *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) *CreateEventStreamingRequestSinkSinkApacheKafkaParameters {
	s.VpcId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) SetForm(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) SetValue(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey) SetForm(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey) SetValue(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersKey) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) SetForm(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) SetValue(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersNetworkType) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) SetForm(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) SetValue(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) SetForm(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) SetValue(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue) SetForm(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue) SetValue(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersValue) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) SetForm(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) SetValue(v string) *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheKafkaParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters struct {
	ConsumeTimestamp *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp `json:"ConsumeTimestamp,omitempty" xml:"ConsumeTimestamp,omitempty" type:"Struct"`
	Group            *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup            `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	InstanceEndpoint *string                                                                                `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	InstancePassword *string                                                                                `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	InstanceUsername *string                                                                                `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	NetworkType      *string                                                                                `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	SecurityGroupId  *string                                                                                `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Topic            *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic            `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	VSwitchId        *string                                                                                `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId            *string                                                                                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetConsumeTimestamp() *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	return s.ConsumeTimestamp
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetGroup() *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup {
	return s.Group
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetTopic() *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic {
	return s.Topic
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetConsumeTimestamp(v *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.ConsumeTimestamp = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetGroup(v *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.Group = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetInstanceEndpoint(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetInstancePassword(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.InstancePassword = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetInstanceUsername(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.InstanceUsername = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetNetworkType(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.NetworkType = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetSecurityGroupId(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetTopic(v *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.Topic = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetVSwitchId(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.VSwitchId = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) SetVpcId(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	s.VpcId = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetForm(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetValue(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) SetForm(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) SetValue(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) SetForm(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) SetValue(v string) *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters struct {
	// example:
	//
	// "https://examplebucket.oss-cn-hangzhou.aliyuncs.com/testDoc/Old_Homebrew/2024-06-26%2022%3A34%3A08/opt/homebrew/homebrew/Library/Homebrew/test/support/fixtures/cask/AppWithBinary.zip?OSSAccessKeyId=ri&Expires=1725539627&Signature=rb8q3OpV2i3gZJ"
	ConnectorPackageUrl *string                                                                                   `json:"ConnectorPackageUrl,omitempty" xml:"ConnectorPackageUrl,omitempty"`
	ConnectorParameters *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters `json:"ConnectorParameters,omitempty" xml:"ConnectorParameters,omitempty" type:"Struct"`
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

func (s CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) GetConnectorPackageUrl() *string {
	return s.ConnectorPackageUrl
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) GetConnectorParameters() *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	return s.ConnectorParameters
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) GetWorkerParameters() map[string]interface{} {
	return s.WorkerParameters
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) SetConnectorPackageUrl(v string) *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters {
	s.ConnectorPackageUrl = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) SetConnectorParameters(v *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters {
	s.ConnectorParameters = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) SetWorkerParameters(v map[string]interface{}) *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters {
	s.WorkerParameters = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters) Validate() error {
	if s.ConnectorParameters != nil {
		if err := s.ConnectorParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters struct {
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
	// example:
	//
	// mongo-sink
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GetName() *string {
	return s.Name
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) SetConfig(v map[string]interface{}) *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	s.Config = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) SetName(v string) *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	s.Name = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkCustomizedKafkaParameters struct {
	// example:
	//
	// 90be1f96-4229-4535-bb76-34b4f6fb2b71
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkCustomizedKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkCustomizedKafkaParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaParameters) SetInstanceId(v string) *CreateEventStreamingRequestSinkSinkCustomizedKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkCustomizedKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDashVectorParameters struct {
	// example:
	//
	// Q34nExQH7sQ****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// collection1
	Collection                 *string                                                                              `json:"Collection,omitempty" xml:"Collection,omitempty"`
	DashVectorSchemaParameters []*CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters `json:"DashVectorSchemaParameters,omitempty" xml:"DashVectorSchemaParameters,omitempty" type:"Repeated"`
	// example:
	//
	// vrs-cn-lbj3ru1***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// PublicNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// Upsert
	Operation    *string                                                              `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Partition    *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition    `json:"Partition,omitempty" xml:"Partition,omitempty" type:"Struct"`
	PrimaryKeyId *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId `json:"PrimaryKeyId,omitempty" xml:"PrimaryKeyId,omitempty" type:"Struct"`
	Vector       *CreateEventStreamingRequestSinkSinkDashVectorParametersVector       `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) GetCollection() *string {
	return s.Collection
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) GetDashVectorSchemaParameters() []*CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
	return s.DashVectorSchemaParameters
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) GetNetwork() *string {
	return s.Network
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) GetOperation() *string {
	return s.Operation
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) GetPartition() *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition {
	return s.Partition
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) GetPrimaryKeyId() *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId {
	return s.PrimaryKeyId
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) GetVector() *CreateEventStreamingRequestSinkSinkDashVectorParametersVector {
	return s.Vector
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) SetApiKey(v string) *CreateEventStreamingRequestSinkSinkDashVectorParameters {
	s.ApiKey = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) SetCollection(v string) *CreateEventStreamingRequestSinkSinkDashVectorParameters {
	s.Collection = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) SetDashVectorSchemaParameters(v []*CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) *CreateEventStreamingRequestSinkSinkDashVectorParameters {
	s.DashVectorSchemaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) SetInstanceId(v string) *CreateEventStreamingRequestSinkSinkDashVectorParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) SetNetwork(v string) *CreateEventStreamingRequestSinkSinkDashVectorParameters {
	s.Network = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) SetOperation(v string) *CreateEventStreamingRequestSinkSinkDashVectorParameters {
	s.Operation = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) SetPartition(v *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition) *CreateEventStreamingRequestSinkSinkDashVectorParameters {
	s.Partition = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) SetPrimaryKeyId(v *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) *CreateEventStreamingRequestSinkSinkDashVectorParameters {
	s.PrimaryKeyId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) SetVector(v *CreateEventStreamingRequestSinkSinkDashVectorParametersVector) *CreateEventStreamingRequestSinkSinkDashVectorParameters {
	s.Vector = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters struct {
	Name  *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName  `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	Type  *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType  `json:"Type,omitempty" xml:"Type,omitempty" type:"Struct"`
	Value *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GetName() *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	return s.Name
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GetType() *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	return s.Type
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GetValue() *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) SetName(v *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Name = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) SetType(v *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Type = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) SetValue(v *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Value = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// content
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) SetForm(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) SetValue(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// STRING
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) SetForm(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) SetValue(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersType) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue struct {
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// ${content}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// $.data.content
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) SetForm(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) SetValue(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParametersValue) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDashVectorParametersPartition struct {
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// ${partition}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// default
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersPartition) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersPartition) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition) SetForm(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition) SetValue(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPartition) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId struct {
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// ${ID}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// $.data.requestId
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) SetForm(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) SetValue(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDashVectorParametersVector struct {
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// ${vector}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// $.data.messageBody
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersVector) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDashVectorParametersVector) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersVector) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersVector) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersVector) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersVector) SetForm(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersVector {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersVector) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersVector {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersVector) SetValue(v string) *CreateEventStreamingRequestSinkSinkDashVectorParametersVector {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDashVectorParametersVector) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDataHubParameters struct {
	// The BLOB topic.
	Body *CreateEventStreamingRequestSinkSinkDataHubParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The name of the DataHub project.
	Project *CreateEventStreamingRequestSinkSinkDataHubParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The role name.
	RoleName *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The name of the DataHub topic.
	Topic *CreateEventStreamingRequestSinkSinkDataHubParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The TUBLE topic.
	TopicSchema *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema `json:"TopicSchema,omitempty" xml:"TopicSchema,omitempty" type:"Struct"`
	// The topic type. Valid values:
	//
	// 	- TUPLE
	//
	// 	- BLOB
	TopicType *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType `json:"TopicType,omitempty" xml:"TopicType,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) GetBody() *CreateEventStreamingRequestSinkSinkDataHubParametersBody {
	return s.Body
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) GetProject() *CreateEventStreamingRequestSinkSinkDataHubParametersProject {
	return s.Project
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) GetRoleName() *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName {
	return s.RoleName
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) GetTopic() *CreateEventStreamingRequestSinkSinkDataHubParametersTopic {
	return s.Topic
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) GetTopicSchema() *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema {
	return s.TopicSchema
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) GetTopicType() *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType {
	return s.TopicType
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetBody(v *CreateEventStreamingRequestSinkSinkDataHubParametersBody) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetProject(v *CreateEventStreamingRequestSinkSinkDataHubParametersProject) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.Project = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetRoleName(v *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.RoleName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetTopic(v *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.Topic = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetTopicSchema(v *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.TopicSchema = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetTopicType(v *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.TopicType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
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

type CreateEventStreamingRequestSinkSinkDataHubParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// ORIGINAL
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The BLOB topic.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersBody) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersBody) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersBody {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersBody) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDataHubParametersProject struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the DataHub project.
	//
	// example:
	//
	// demo-project
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersProject) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersProject) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersProject) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersProject) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersProject) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersProject) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersProject {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersProject) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersProject {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersProject) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersProject {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersProject) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDataHubParametersRoleName struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name.
	//
	// example:
	//
	// test-role
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDataHubParametersTopic struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the DataHub topic.
	//
	// example:
	//
	// demo-topic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopic) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopic {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopic {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopic {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// {"k1":"${k1}","k2":"${k2}"}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The TUBLE topic.
	//
	// example:
	//
	// {"k1":"value1","k2":"value2"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDataHubParametersTopicType struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic type. Valid values:
	//
	// 	- TUPLE
	//
	// 	- BLOB
	//
	// example:
	//
	// TUPLE
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParameters struct {
	BeHttpEndpoint  *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint  `json:"BeHttpEndpoint,omitempty" xml:"BeHttpEndpoint,omitempty" type:"Struct"`
	Body            *CreateEventStreamingRequestSinkSinkDorisParametersBody            `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Database        *CreateEventStreamingRequestSinkSinkDorisParametersDatabase        `json:"Database,omitempty" xml:"Database,omitempty" type:"Struct"`
	FeHttpEndpoint  *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint  `json:"FeHttpEndpoint,omitempty" xml:"FeHttpEndpoint,omitempty" type:"Struct"`
	NetworkType     *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType     `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	Password        *CreateEventStreamingRequestSinkSinkDorisParametersPassword        `json:"Password,omitempty" xml:"Password,omitempty" type:"Struct"`
	QueryEndpoint   *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint   `json:"QueryEndpoint,omitempty" xml:"QueryEndpoint,omitempty" type:"Struct"`
	SecurityGroupId *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	Table           *CreateEventStreamingRequestSinkSinkDorisParametersTable           `json:"Table,omitempty" xml:"Table,omitempty" type:"Struct"`
	UserName        *CreateEventStreamingRequestSinkSinkDorisParametersUserName        `json:"UserName,omitempty" xml:"UserName,omitempty" type:"Struct"`
	VSwitchIds      *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds      `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	VpcId           *CreateEventStreamingRequestSinkSinkDorisParametersVpcId           `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetBeHttpEndpoint() *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint {
	return s.BeHttpEndpoint
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetBody() *CreateEventStreamingRequestSinkSinkDorisParametersBody {
	return s.Body
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetDatabase() *CreateEventStreamingRequestSinkSinkDorisParametersDatabase {
	return s.Database
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetFeHttpEndpoint() *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint {
	return s.FeHttpEndpoint
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetNetworkType() *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType {
	return s.NetworkType
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetPassword() *CreateEventStreamingRequestSinkSinkDorisParametersPassword {
	return s.Password
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetQueryEndpoint() *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint {
	return s.QueryEndpoint
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetSecurityGroupId() *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetTable() *CreateEventStreamingRequestSinkSinkDorisParametersTable {
	return s.Table
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetUserName() *CreateEventStreamingRequestSinkSinkDorisParametersUserName {
	return s.UserName
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetVSwitchIds() *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) GetVpcId() *CreateEventStreamingRequestSinkSinkDorisParametersVpcId {
	return s.VpcId
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetBeHttpEndpoint(v *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.BeHttpEndpoint = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetBody(v *CreateEventStreamingRequestSinkSinkDorisParametersBody) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetDatabase(v *CreateEventStreamingRequestSinkSinkDorisParametersDatabase) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.Database = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetFeHttpEndpoint(v *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.FeHttpEndpoint = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetNetworkType(v *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.NetworkType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetPassword(v *CreateEventStreamingRequestSinkSinkDorisParametersPassword) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.Password = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetQueryEndpoint(v *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.QueryEndpoint = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetSecurityGroupId(v *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.SecurityGroupId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetTable(v *CreateEventStreamingRequestSinkSinkDorisParametersTable) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.Table = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetUserName(v *CreateEventStreamingRequestSinkSinkDorisParametersUserName) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.UserName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetVSwitchIds(v *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.VSwitchIds = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) SetVpcId(v *CreateEventStreamingRequestSinkSinkDorisParametersVpcId) *CreateEventStreamingRequestSinkSinkDorisParameters {
	s.VpcId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBeHttpEndpoint) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParametersBody struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBody) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBody) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersBody {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersBody) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParametersDatabase struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersDatabase) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersDatabase) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersDatabase) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersDatabase) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersDatabase) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersDatabase) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersDatabase {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersDatabase) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersDatabase {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersDatabase) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersDatabase {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersDatabase) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersFeHttpEndpoint) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParametersNetworkType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersNetworkType) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersNetworkType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersNetworkType) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParametersPassword struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersPassword) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersPassword) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersPassword) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersPassword) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersPassword) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersPassword) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersPassword {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersPassword) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersPassword {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersPassword) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersPassword {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersPassword) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersQueryEndpoint) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParametersTable struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersTable) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersTable) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersTable) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersTable) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersTable) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersTable) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersTable {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersTable) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersTable {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersTable) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersTable {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersTable) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParametersUserName struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersUserName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersUserName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersUserName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersUserName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersUserName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersUserName) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersUserName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersUserName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersUserName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersUserName) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersUserName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersUserName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkDorisParametersVpcId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDorisParametersVpcId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVpcId) SetForm(v string) *CreateEventStreamingRequestSinkSinkDorisParametersVpcId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVpcId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDorisParametersVpcId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVpcId) SetValue(v string) *CreateEventStreamingRequestSinkSinkDorisParametersVpcId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDorisParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkFcParameters struct {
	// The message body that you want to deliver to Function Compute.
	Body *CreateEventStreamingRequestSinkSinkFcParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The delivery concurrency. Minimum value: 1.
	Concurrency *CreateEventStreamingRequestSinkSinkFcParametersConcurrency `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	DataFormat  *CreateEventStreamingRequestSinkSinkFcParametersDataFormat  `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" type:"Struct"`
	// The function name.
	FunctionName *CreateEventStreamingRequestSinkSinkFcParametersFunctionName `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	// The invocation mode. Valid values: Sync and Async.
	InvocationType *CreateEventStreamingRequestSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	// The service version.
	Qualifier *CreateEventStreamingRequestSinkSinkFcParametersQualifier `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	// The service name.
	ServiceName *CreateEventStreamingRequestSinkSinkFcParametersServiceName `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkFcParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) GetBody() *CreateEventStreamingRequestSinkSinkFcParametersBody {
	return s.Body
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) GetConcurrency() *CreateEventStreamingRequestSinkSinkFcParametersConcurrency {
	return s.Concurrency
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) GetDataFormat() *CreateEventStreamingRequestSinkSinkFcParametersDataFormat {
	return s.DataFormat
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) GetFunctionName() *CreateEventStreamingRequestSinkSinkFcParametersFunctionName {
	return s.FunctionName
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) GetInvocationType() *CreateEventStreamingRequestSinkSinkFcParametersInvocationType {
	return s.InvocationType
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) GetQualifier() *CreateEventStreamingRequestSinkSinkFcParametersQualifier {
	return s.Qualifier
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) GetServiceName() *CreateEventStreamingRequestSinkSinkFcParametersServiceName {
	return s.ServiceName
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetBody(v *CreateEventStreamingRequestSinkSinkFcParametersBody) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetConcurrency(v *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.Concurrency = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetDataFormat(v *CreateEventStreamingRequestSinkSinkFcParametersDataFormat) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.DataFormat = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetFunctionName(v *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.FunctionName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetInvocationType(v *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.InvocationType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetQualifier(v *CreateEventStreamingRequestSinkSinkFcParametersQualifier) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.Qualifier = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetServiceName(v *CreateEventStreamingRequestSinkSinkFcParametersServiceName) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.ServiceName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkFcParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
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

func (s CreateEventStreamingRequestSinkSinkFcParametersBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersBody) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersBody) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersBody {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersBody) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkFcParametersConcurrency struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The delivery concurrency. Minimum value: 1.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersConcurrency) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersConcurrency) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkFcParametersDataFormat struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersDataFormat) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersDataFormat) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersDataFormat) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersDataFormat) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersDataFormat) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersDataFormat) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersDataFormat {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersDataFormat) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersDataFormat {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersDataFormat) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersDataFormat {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersDataFormat) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkFcParametersFunctionName struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The function name.
	//
	// example:
	//
	// myFunction
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersFunctionName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersFunctionName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkFcParametersInvocationType struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The invocation mode. Valid values: Sync and Async.
	//
	// example:
	//
	// Async
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersInvocationType) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersInvocationType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkFcParametersQualifier struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The service version.
	//
	// example:
	//
	// LATEST
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersQualifier) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersQualifier) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersQualifier) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersQualifier) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersQualifier) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersQualifier) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersQualifier) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersQualifier) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersQualifier) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkFcParametersServiceName struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The service name.
	//
	// example:
	//
	// myService
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersServiceName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersServiceName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersServiceName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersServiceName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersServiceName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersServiceName) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersServiceName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersServiceName) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersServiceName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkFnfParameters struct {
	// The execution name.
	ExecutionName *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty" type:"Struct"`
	// The flow name.
	FlowName *CreateEventStreamingRequestSinkSinkFnfParametersFlowName `json:"FlowName,omitempty" xml:"FlowName,omitempty" type:"Struct"`
	// The input information of the execution.
	Input *CreateEventStreamingRequestSinkSinkFnfParametersInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The role name.
	RoleName *CreateEventStreamingRequestSinkSinkFnfParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkFnfParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFnfParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) GetExecutionName() *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	return s.ExecutionName
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) GetFlowName() *CreateEventStreamingRequestSinkSinkFnfParametersFlowName {
	return s.FlowName
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) GetInput() *CreateEventStreamingRequestSinkSinkFnfParametersInput {
	return s.Input
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) GetRoleName() *CreateEventStreamingRequestSinkSinkFnfParametersRoleName {
	return s.RoleName
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) SetExecutionName(v *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) *CreateEventStreamingRequestSinkSinkFnfParameters {
	s.ExecutionName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) SetFlowName(v *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) *CreateEventStreamingRequestSinkSinkFnfParameters {
	s.FlowName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) SetInput(v *CreateEventStreamingRequestSinkSinkFnfParametersInput) *CreateEventStreamingRequestSinkSinkFnfParameters {
	s.Input = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) SetRoleName(v *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) *CreateEventStreamingRequestSinkSinkFnfParameters {
	s.RoleName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkFnfParametersExecutionName struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The execution name.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetForm(v string) *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetValue(v string) *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkFnfParametersFlowName struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The flow name.
	//
	// example:
	//
	// test-streaming-fnf
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersFlowName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersFlowName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) SetForm(v string) *CreateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) SetValue(v string) *CreateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkFnfParametersInput struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The input information of the execution.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersInput) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersInput) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersInput) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersInput) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersInput) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersInput) SetForm(v string) *CreateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersInput) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersInput) SetValue(v string) *CreateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersInput) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkFnfParametersRoleName struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name.
	//
	// example:
	//
	// Al****FNF-x****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersRoleName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) SetForm(v string) *CreateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) SetValue(v string) *CreateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkKafkaParameters struct {
	// The acknowledgment (ACK) mode.
	//
	// 	- If you set this parameter to 0, no response is returned from the broker. In this mode, the performance is high, but the risk of data loss is also high.
	//
	// 	- If you set this parameter to 1, a response is returned when data is written to the leader. In this mode, the performance and the risk of data loss are moderate. Data loss may occur if a failure occurs on the leader.
	//
	// 	- If you set this parameter to all, a response is returned when data is written to the leader and synchronized to the followers. In this mode, the performance is low, but the risk of data loss is also low. Data loss occurs if the leader and the followers fail at the same time.
	Acks            *CreateEventStreamingRequestSinkSinkKafkaParametersAcks    `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	CompressionType *string                                                    `json:"CompressionType,omitempty" xml:"CompressionType,omitempty"`
	Headers         *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The ID of the ApsaraMQ for Kafka instance.
	InstanceId *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message key.
	Key *CreateEventStreamingRequestSinkSinkKafkaParametersKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	// The name of the topic on the ApsaraMQ for Kafka instance.
	Topic *CreateEventStreamingRequestSinkSinkKafkaParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The message body.
	Value *CreateEventStreamingRequestSinkSinkKafkaParametersValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) GetAcks() *CreateEventStreamingRequestSinkSinkKafkaParametersAcks {
	return s.Acks
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) GetCompressionType() *string {
	return s.CompressionType
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) GetHeaders() *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders {
	return s.Headers
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) GetInstanceId() *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) GetKey() *CreateEventStreamingRequestSinkSinkKafkaParametersKey {
	return s.Key
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) GetTopic() *CreateEventStreamingRequestSinkSinkKafkaParametersTopic {
	return s.Topic
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) GetValue() *CreateEventStreamingRequestSinkSinkKafkaParametersValue {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetAcks(v *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.Acks = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetCompressionType(v string) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.CompressionType = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetHeaders(v *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.Headers = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetInstanceId(v *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.InstanceId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetKey(v *CreateEventStreamingRequestSinkSinkKafkaParametersKey) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.Key = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetTopic(v *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.Topic = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetValue(v *CreateEventStreamingRequestSinkSinkKafkaParametersValue) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.Value = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) Validate() error {
	if s.Acks != nil {
		if err := s.Acks.Validate(); err != nil {
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

type CreateEventStreamingRequestSinkSinkKafkaParametersAcks struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ACK mode.
	//
	// 	- If you set this parameter to 0, no response is returned from the broker. In this mode, the performance is high, but the risk of data loss is also high.
	//
	// 	- If you set this parameter to 1, a response is returned when data is written to the leader. In this mode, the performance and the risk of data loss are moderate. Data loss may occur if a failure occurs on the leader.
	//
	// 	- If you set this parameter to all, a response is returned when data is written to the leader and synchronized to the followers. In this mode, the performance is low, but the risk of data loss is also low. Data loss occurs if the leader and the followers fail at the same time.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersAcks) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersAcks) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) SetForm(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) SetValue(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkKafkaParametersHeaders struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersHeaders) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders) SetForm(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders) SetValue(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the ApsaraMQ for Kafka instance.
	//
	// example:
	//
	// Defaut_1283278472_sadkj
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetForm(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetValue(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkKafkaParametersKey struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The message key.
	//
	// example:
	//
	// key
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersKey) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersKey) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersKey) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersKey) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersKey) SetForm(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersKey) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersKey) SetValue(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersKey) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkKafkaParametersTopic struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic on the ApsaraMQ for Kafka instance.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersTopic) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) SetForm(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) SetValue(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkKafkaParametersValue struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
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

func (s CreateEventStreamingRequestSinkSinkKafkaParametersValue) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersValue) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersValue) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersValue) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersValue) SetForm(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersValue) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersValue) SetValue(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersValue) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkMNSParameters struct {
	// The message content.
	Body *CreateEventStreamingRequestSinkSinkMNSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// Specifies whether to enable Base64 encoding.
	IsBase64Encode *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	// The name of the MNS queue.
	QueueName *CreateEventStreamingRequestSinkSinkMNSParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkMNSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkMNSParameters) GetBody() *CreateEventStreamingRequestSinkSinkMNSParametersBody {
	return s.Body
}

func (s *CreateEventStreamingRequestSinkSinkMNSParameters) GetIsBase64Encode() *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	return s.IsBase64Encode
}

func (s *CreateEventStreamingRequestSinkSinkMNSParameters) GetQueueName() *CreateEventStreamingRequestSinkSinkMNSParametersQueueName {
	return s.QueueName
}

func (s *CreateEventStreamingRequestSinkSinkMNSParameters) SetBody(v *CreateEventStreamingRequestSinkSinkMNSParametersBody) *CreateEventStreamingRequestSinkSinkMNSParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParameters) SetIsBase64Encode(v *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) *CreateEventStreamingRequestSinkSinkMNSParameters {
	s.IsBase64Encode = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParameters) SetQueueName(v *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) *CreateEventStreamingRequestSinkSinkMNSParameters {
	s.QueueName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkMNSParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
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
	//   "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersBody) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersBody) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersBody) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Specifies whether to enable Base64 encoding.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetForm(v string) *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetValue(v string) *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkMNSParametersQueueName struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the MNS queue.
	//
	// example:
	//
	// MyQueue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersQueueName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersQueueName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) SetForm(v string) *CreateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) SetValue(v string) *CreateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters struct {
	AuthType        *string                                                                    `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Body            *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody       `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Endpoint        *string                                                                    `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Exchange        *string                                                                    `json:"Exchange,omitempty" xml:"Exchange,omitempty"`
	MessageId       *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId  `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	NetworkType     *string                                                                    `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Password        *string                                                                    `json:"Password,omitempty" xml:"Password,omitempty"`
	Properties      *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	QueueName       *string                                                                    `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RoutingKey      *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	SecurityGroupId *string                                                                    `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	TargetType      *string                                                                    `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	Username        *string                                                                    `json:"Username,omitempty" xml:"Username,omitempty"`
	VSwitchIds      *string                                                                    `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VirtualHostName *string                                                                    `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	VpcId           *string                                                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetBody() *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody {
	return s.Body
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetExchange() *string {
	return s.Exchange
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetMessageId() *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId {
	return s.MessageId
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetPassword() *string {
	return s.Password
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetProperties() *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties {
	return s.Properties
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetRoutingKey() *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	return s.RoutingKey
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetUsername() *string {
	return s.Username
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetAuthType(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.AuthType = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetBody(v *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetEndpoint(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Endpoint = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetExchange(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Exchange = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetMessageId(v *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetNetworkType(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.NetworkType = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetPassword(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Password = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetProperties(v *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetQueueName(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetRoutingKey(v *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetSecurityGroupId(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetTargetType(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.TargetType = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetUsername(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.Username = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetVSwitchIds(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetVirtualHostName(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) SetVpcId(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	s.VpcId = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) SetForm(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) SetValue(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) SetForm(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) SetValue(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetForm(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetValue(v string) *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkPrometheusParameters struct {
	// The authentication method.
	AuthorizationType *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty" type:"Struct"`
	// The metric data.
	Data             *CreateEventStreamingRequestSinkSinkPrometheusParametersData             `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HeaderParameters *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Struct"`
	// The network type.
	NetworkType *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	// The password.
	Password *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword `json:"Password,omitempty" xml:"Password,omitempty" type:"Struct"`
	// The ID of the security group to which the Managed Service for Prometheus instance belongs.
	SecurityGroupId *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	// The URL for the remote write configuration item of the Managed Service for Prometheus instance.
	URL *CreateEventStreamingRequestSinkSinkPrometheusParametersURL `json:"URL,omitempty" xml:"URL,omitempty" type:"Struct"`
	// The username.
	Username *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername `json:"Username,omitempty" xml:"Username,omitempty" type:"Struct"`
	// The ID of the vSwitch with which the Managed Service for Prometheus instance is associated.
	VSwitchId *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Struct"`
	// The ID of the VPC to which the Managed Service for Prometheus instance belongs.
	VpcId *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) GetAuthorizationType() *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	return s.AuthorizationType
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) GetData() *CreateEventStreamingRequestSinkSinkPrometheusParametersData {
	return s.Data
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) GetHeaderParameters() *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters {
	return s.HeaderParameters
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) GetNetworkType() *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	return s.NetworkType
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) GetPassword() *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	return s.Password
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) GetSecurityGroupId() *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) GetURL() *CreateEventStreamingRequestSinkSinkPrometheusParametersURL {
	return s.URL
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) GetUsername() *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	return s.Username
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) GetVSwitchId() *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	return s.VSwitchId
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) GetVpcId() *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	return s.VpcId
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetAuthorizationType(v *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.AuthorizationType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetData(v *CreateEventStreamingRequestSinkSinkPrometheusParametersData) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Data = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetHeaderParameters(v *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.HeaderParameters = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetNetworkType(v *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.NetworkType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetPassword(v *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Password = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetSecurityGroupId(v *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.SecurityGroupId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetURL(v *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.URL = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetUsername(v *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Username = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetVSwitchId(v *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.VSwitchId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetVpcId(v *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.VpcId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The authentication method.
	//
	// example:
	//
	// BASIC_AUTH
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersData struct {
	// The method that you want to use to transform events. Default value: JSAONPATH.
	//
	// example:
	//
	// JSAONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The metric data.
	//
	// example:
	//
	// $.data
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersData) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersData) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersData) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersData) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersData) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersData) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersData) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersData) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersData) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters struct {
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// {
	//
	//     "user_name":"${name}"
	//
	// }
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The network type.
	//
	// example:
	//
	// PrivateNetwork
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersPassword struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
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
	// *****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the security group to which the Managed Service for Prometheus instance belongs.
	//
	// example:
	//
	// sg-mw43*****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersURL struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The URL for the remote write configuration item of the Managed Service for Prometheus instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersURL) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersURL) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersUsername struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
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
	// ****admin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the vSwitch with which the Managed Service for Prometheus instance is associated.
	//
	// example:
	//
	// vsw-dwaafds****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the VPC to which the Managed Service for Prometheus instance belongs.
	//
	// example:
	//
	// i-2ze7u5i17mbqtx1p****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRabbitMQParameters struct {
	// The message content.
	Body *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The exchange mode. This parameter is required only if you set TargetType to Exchange.
	Exchange *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
	InstanceId *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message ID.
	MessageId *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// The properties that you want to use to filter messages.
	Properties *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The queue mode. This parameter is required only if you set TargetType to Queue.
	QueueName *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	// The rule that you want to use to route messages. This parameter is required only if you set TargetType to Exchange.
	RoutingKey *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// The type of the resource to which you want to deliver events.
	TargetType *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
	// The name of the vhost to which the ApsaraMQ for RabbitMQ instance belongs.
	VirtualHostName *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) GetBody() *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	return s.Body
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) GetExchange() *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	return s.Exchange
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) GetInstanceId() *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) GetMessageId() *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	return s.MessageId
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) GetProperties() *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	return s.Properties
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) GetQueueName() *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	return s.QueueName
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) GetRoutingKey() *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	return s.RoutingKey
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) GetTargetType() *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	return s.TargetType
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) GetVirtualHostName() *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	return s.VirtualHostName
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetBody(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetExchange(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Exchange = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetInstanceId(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.InstanceId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetMessageId(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetProperties(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetQueueName(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.QueueName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetRoutingKey(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetTargetType(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.TargetType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetVirtualHostName(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.VirtualHostName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkRabbitMQParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
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
	//   "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the exchange on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// a_exchange
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// a5ff91ad4f3f24947887fe184fc20d07
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
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
	//   "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
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
	//   "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the queue on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// MyQueue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The rule that you want to use to route messages.
	//
	// example:
	//
	// housekeeping
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the resource to which you want to deliver events. Valid values:
	//
	// 	- Exchange
	//
	// 	- Queue
	//
	// example:
	//
	// Exchange/Queue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the vhost to which the ApsaraMQ for RabbitMQ instance belongs.
	//
	// example:
	//
	// rabbit-host
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters struct {
	ConsumeTimestamp *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp `json:"ConsumeTimestamp,omitempty" xml:"ConsumeTimestamp,omitempty" type:"Struct"`
	Group            *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup            `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	InstanceId       *string                                                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType     *string                                                                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Topic            *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic            `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GetConsumeTimestamp() *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	return s.ConsumeTimestamp
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GetGroup() *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup {
	return s.Group
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) GetTopic() *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic {
	return s.Topic
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) SetConsumeTimestamp(v *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	s.ConsumeTimestamp = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) SetGroup(v *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	s.Group = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) SetInstanceId(v string) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) SetInstanceType(v string) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	s.InstanceType = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) SetTopic(v *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters {
	s.Topic = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQCheckpointParametersTopic) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParameters struct {
	// The message content.
	Body              *CreateEventStreamingRequestSinkSinkRocketMQParametersBody              `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	DeliveryOrderType *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType `json:"DeliveryOrderType,omitempty" xml:"DeliveryOrderType,omitempty" type:"Struct"`
	// The endpoint that you want to use to access the ApsaraMQ for RocketMQ instance.
	InstanceEndpoint *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty" type:"Struct"`
	// The ID of the ApsaraMQ for RocketMQ instance.
	InstanceId *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The password that you want to use to access the ApsaraMQ for RocketMQ instance.
	InstancePassword *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty" type:"Struct"`
	// The type of the ApsaraMQ for RocketMQ instance.
	InstanceType *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Struct"`
	// The username that you want to use to access the ApsaraMQ for RocketMQ instance.
	InstanceUsername *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty" type:"Struct"`
	// The keys that you want to use to filter messages.
	Keys *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	// The network type.
	//
	// 	- PublicNetwork
	//
	// 	- PrivateNetwork
	Network *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The properties that you want to use to filter messages.
	Properties *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The ID of the security group to which the ApsaraMQ for RocketMQ instance belongs.
	SecurityGroupId *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	ShardingKey     *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey     `json:"ShardingKey,omitempty" xml:"ShardingKey,omitempty" type:"Struct"`
	// The tags that you want to use to filter messages.
	Tags *CreateEventStreamingRequestSinkSinkRocketMQParametersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The name of the topic on the ApsaraMQ for RocketMQ instance.
	Topic *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The ID of the vSwitch with which the ApsaraMQ for RocketMQ instance is associated.
	VSwitchIds *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	// The ID of the VPC to which the ApsaraMQ for RocketMQ instance belongs.
	VpcId *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetBody() *CreateEventStreamingRequestSinkSinkRocketMQParametersBody {
	return s.Body
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetDeliveryOrderType() *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType {
	return s.DeliveryOrderType
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetInstanceEndpoint() *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint {
	return s.InstanceEndpoint
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetInstanceId() *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetInstancePassword() *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword {
	return s.InstancePassword
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetInstanceType() *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType {
	return s.InstanceType
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetInstanceUsername() *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername {
	return s.InstanceUsername
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetKeys() *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	return s.Keys
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetNetwork() *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork {
	return s.Network
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetProperties() *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	return s.Properties
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetSecurityGroupId() *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetShardingKey() *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey {
	return s.ShardingKey
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetTags() *CreateEventStreamingRequestSinkSinkRocketMQParametersTags {
	return s.Tags
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetTopic() *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	return s.Topic
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetVSwitchIds() *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) GetVpcId() *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId {
	return s.VpcId
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetBody(v *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetDeliveryOrderType(v *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.DeliveryOrderType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceEndpoint(v *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceEndpoint = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceId(v *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetInstancePassword(v *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstancePassword = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceType(v *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceUsername(v *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceUsername = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetKeys(v *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Keys = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetNetwork(v *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Network = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetProperties(v *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Properties = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetSecurityGroupId(v *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.SecurityGroupId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetShardingKey(v *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.ShardingKey = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetTags(v *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Tags = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetTopic(v *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Topic = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetVSwitchIds(v *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.VSwitchIds = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetVpcId(v *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.VpcId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkRocketMQParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
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
	//   "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The endpoint that you want to use to access the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// vbr-8vbsvkkbpf3vb0zefs7ex
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// MQ_INST_164901546557****_BAAN****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The password that you want to use to access the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// admin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// 2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The username that you want to use to access the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// admin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersKeys struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
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
	//   "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The network type.
	//
	// 	- PublicNetwork
	//
	// 	- PrivateNetwork
	//
	// example:
	//
	// PublicNetwork
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersProperties struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
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
	//   "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the security group to which the ApsaraMQ for RocketMQ instance belongs.
	//
	// example:
	//
	// b4bf375515f6440f942e3a20c33d5b9c
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersShardingKey) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersTags struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
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
	//   "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersTags) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersTags) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersTopic struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic on the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// Mytopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the vSwitch with which the ApsaraMQ for RocketMQ instance is associated.
	//
	// example:
	//
	// vbr-8vb835n3zf9shwlvbwlmp
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the VPC to which the ApsaraMQ for RocketMQ instance belongs.
	//
	// example:
	//
	// vbr-8vb835n3zf9shwlvbwlmp
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkSLSParameters struct {
	// The message body that you want to deliver to Simple Log Service.
	Body          *CreateEventStreamingRequestSinkSinkSLSParametersBody          `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	ContentSchema *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema `json:"ContentSchema,omitempty" xml:"ContentSchema,omitempty" type:"Struct"`
	ContentType   *CreateEventStreamingRequestSinkSinkSLSParametersContentType   `json:"ContentType,omitempty" xml:"ContentType,omitempty" type:"Struct"`
	// The Simple Log Service Logstore.
	LogStore *CreateEventStreamingRequestSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	// The Simple Log Service project.
	Project *CreateEventStreamingRequestSinkSinkSLSParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Simple Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	RoleName *CreateEventStreamingRequestSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The topic that you want to use to store logs. This parameter corresponds to the reserved field **topic*	- in Simple Log Service.
	Topic *CreateEventStreamingRequestSinkSinkSLSParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) GetBody() *CreateEventStreamingRequestSinkSinkSLSParametersBody {
	return s.Body
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) GetContentSchema() *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema {
	return s.ContentSchema
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) GetContentType() *CreateEventStreamingRequestSinkSinkSLSParametersContentType {
	return s.ContentType
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) GetLogStore() *CreateEventStreamingRequestSinkSinkSLSParametersLogStore {
	return s.LogStore
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) GetProject() *CreateEventStreamingRequestSinkSinkSLSParametersProject {
	return s.Project
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) GetRoleName() *CreateEventStreamingRequestSinkSinkSLSParametersRoleName {
	return s.RoleName
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) GetTopic() *CreateEventStreamingRequestSinkSinkSLSParametersTopic {
	return s.Topic
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetBody(v *CreateEventStreamingRequestSinkSinkSLSParametersBody) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetContentSchema(v *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.ContentSchema = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetContentType(v *CreateEventStreamingRequestSinkSinkSLSParametersContentType) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.ContentType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetLogStore(v *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.LogStore = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetProject(v *CreateEventStreamingRequestSinkSinkSLSParametersProject) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.Project = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetRoleName(v *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.RoleName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetTopic(v *CreateEventStreamingRequestSinkSinkSLSParametersTopic) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.Topic = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) Validate() error {
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

type CreateEventStreamingRequestSinkSinkSLSParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
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

func (s CreateEventStreamingRequestSinkSinkSLSParametersBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersBody) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersBody) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersBody) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkSLSParametersContentSchema struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// {"Key_1":{"form":"CONSTANT","value":"demoKey"},"Value_1":{"form":"JSONPATH","value":"$.data.value"}}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersContentSchema) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersContentSchema) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentSchema) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkSLSParametersContentType struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// JSON
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersContentType) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersContentType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentType) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentType) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentType) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentType) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersContentType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersContentType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentType) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersContentType {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersContentType) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkSLSParametersLogStore struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Simple Log Service Logstore.
	//
	// example:
	//
	// test-logstore
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersLogStore) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersLogStore) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkSLSParametersProject struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Simple Log Service project.
	//
	// example:
	//
	// test-project
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersProject) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersProject) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersProject) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersProject) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersProject) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersProject) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersProject) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersProject) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersProject) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkSLSParametersRoleName struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Simple Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	//
	// example:
	//
	// testRole
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersRoleName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSinkSinkSLSParametersTopic struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic that you want to use to store logs. This parameter corresponds to the reserved field **topic*	- in Simple Log Service.
	//
	// example:
	//
	// testTopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersTopic) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersTopic) GetForm() *string {
	return s.Form
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersTopic) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersTopic) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersTopic) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersTopic) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersTopic) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSource struct {
	SourceApacheKafkaParameters              *CreateEventStreamingRequestSourceSourceApacheKafkaParameters              `json:"SourceApacheKafkaParameters,omitempty" xml:"SourceApacheKafkaParameters,omitempty" type:"Struct"`
	SourceApacheRocketMQCheckpointParameters *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters `json:"SourceApacheRocketMQCheckpointParameters,omitempty" xml:"SourceApacheRocketMQCheckpointParameters,omitempty" type:"Struct"`
	SourceCustomizedKafkaConnectorParameters *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters `json:"SourceCustomizedKafkaConnectorParameters,omitempty" xml:"SourceCustomizedKafkaConnectorParameters,omitempty" type:"Struct"`
	SourceCustomizedKafkaParameters          *CreateEventStreamingRequestSourceSourceCustomizedKafkaParameters          `json:"SourceCustomizedKafkaParameters,omitempty" xml:"SourceCustomizedKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Data Transmission Service (DTS) as the event source.
	SourceDTSParameters      *CreateEventStreamingRequestSourceSourceDTSParameters      `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	SourceEventBusParameters *CreateEventStreamingRequestSourceSourceEventBusParameters `json:"SourceEventBusParameters,omitempty" xml:"SourceEventBusParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for Kafka as the event source.
	SourceKafkaParameters *CreateEventStreamingRequestSourceSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Service (MNS) as the event source.
	SourceMNSParameters *CreateEventStreamingRequestSourceSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for MQTT as the event source.
	SourceMQTTParameters               *CreateEventStreamingRequestSourceSourceMQTTParameters               `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	SourceMySQLParameters              *SourceMySQLParameters                                               `json:"SourceMySQLParameters,omitempty" xml:"SourceMySQLParameters,omitempty"`
	SourceOSSParameters                *CreateEventStreamingRequestSourceSourceOSSParameters                `json:"SourceOSSParameters,omitempty" xml:"SourceOSSParameters,omitempty" type:"Struct"`
	SourceOpenSourceRabbitMQParameters *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters `json:"SourceOpenSourceRabbitMQParameters,omitempty" xml:"SourceOpenSourceRabbitMQParameters,omitempty" type:"Struct"`
	SourcePostgreSQLParameters         *SourcePostgreSQLParameters                                          `json:"SourcePostgreSQLParameters,omitempty" xml:"SourcePostgreSQLParameters,omitempty"`
	// Parameters that are configured if you specify Managed Service for Prometheus as the event source.
	SourcePrometheusParameters      *CreateEventStreamingRequestSourceSourcePrometheusParameters `json:"SourcePrometheusParameters,omitempty" xml:"SourcePrometheusParameters,omitempty" type:"Struct"`
	SourceRabbitMQMetaParameters    *SourceRabbitMQMetaParameters                                `json:"SourceRabbitMQMetaParameters,omitempty" xml:"SourceRabbitMQMetaParameters,omitempty"`
	SourceRabbitMQMsgSyncParameters *SourceRabbitMQMsgSyncParameters                             `json:"SourceRabbitMQMsgSyncParameters,omitempty" xml:"SourceRabbitMQMsgSyncParameters,omitempty"`
	// The parameters that are configured if you specify ApsaraMQ for RabbitMQ as the event source.
	SourceRabbitMQParameters           *CreateEventStreamingRequestSourceSourceRabbitMQParameters           `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	SourceRocketMQCheckpointParameters *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters `json:"SourceRocketMQCheckpointParameters,omitempty" xml:"SourceRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for RocketMQ as the event source.
	SourceRocketMQParameters *CreateEventStreamingRequestSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Simple Log Service as the event source.
	SourceSLSParameters *CreateEventStreamingRequestSourceSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSource) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSource) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSource) GetSourceApacheKafkaParameters() *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	return s.SourceApacheKafkaParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceApacheRocketMQCheckpointParameters() *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	return s.SourceApacheRocketMQCheckpointParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceCustomizedKafkaConnectorParameters() *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters {
	return s.SourceCustomizedKafkaConnectorParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceCustomizedKafkaParameters() *CreateEventStreamingRequestSourceSourceCustomizedKafkaParameters {
	return s.SourceCustomizedKafkaParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceDTSParameters() *CreateEventStreamingRequestSourceSourceDTSParameters {
	return s.SourceDTSParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceEventBusParameters() *CreateEventStreamingRequestSourceSourceEventBusParameters {
	return s.SourceEventBusParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceKafkaParameters() *CreateEventStreamingRequestSourceSourceKafkaParameters {
	return s.SourceKafkaParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceMNSParameters() *CreateEventStreamingRequestSourceSourceMNSParameters {
	return s.SourceMNSParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceMQTTParameters() *CreateEventStreamingRequestSourceSourceMQTTParameters {
	return s.SourceMQTTParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceMySQLParameters() *SourceMySQLParameters {
	return s.SourceMySQLParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceOSSParameters() *CreateEventStreamingRequestSourceSourceOSSParameters {
	return s.SourceOSSParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceOpenSourceRabbitMQParameters() *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	return s.SourceOpenSourceRabbitMQParameters
}

func (s *CreateEventStreamingRequestSource) GetSourcePostgreSQLParameters() *SourcePostgreSQLParameters {
	return s.SourcePostgreSQLParameters
}

func (s *CreateEventStreamingRequestSource) GetSourcePrometheusParameters() *CreateEventStreamingRequestSourceSourcePrometheusParameters {
	return s.SourcePrometheusParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceRabbitMQMetaParameters() *SourceRabbitMQMetaParameters {
	return s.SourceRabbitMQMetaParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceRabbitMQMsgSyncParameters() *SourceRabbitMQMsgSyncParameters {
	return s.SourceRabbitMQMsgSyncParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceRabbitMQParameters() *CreateEventStreamingRequestSourceSourceRabbitMQParameters {
	return s.SourceRabbitMQParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceRocketMQCheckpointParameters() *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters {
	return s.SourceRocketMQCheckpointParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceRocketMQParameters() *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	return s.SourceRocketMQParameters
}

func (s *CreateEventStreamingRequestSource) GetSourceSLSParameters() *CreateEventStreamingRequestSourceSourceSLSParameters {
	return s.SourceSLSParameters
}

func (s *CreateEventStreamingRequestSource) SetSourceApacheKafkaParameters(v *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) *CreateEventStreamingRequestSource {
	s.SourceApacheKafkaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceApacheRocketMQCheckpointParameters(v *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) *CreateEventStreamingRequestSource {
	s.SourceApacheRocketMQCheckpointParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceCustomizedKafkaConnectorParameters(v *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) *CreateEventStreamingRequestSource {
	s.SourceCustomizedKafkaConnectorParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceCustomizedKafkaParameters(v *CreateEventStreamingRequestSourceSourceCustomizedKafkaParameters) *CreateEventStreamingRequestSource {
	s.SourceCustomizedKafkaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceDTSParameters(v *CreateEventStreamingRequestSourceSourceDTSParameters) *CreateEventStreamingRequestSource {
	s.SourceDTSParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceEventBusParameters(v *CreateEventStreamingRequestSourceSourceEventBusParameters) *CreateEventStreamingRequestSource {
	s.SourceEventBusParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceKafkaParameters(v *CreateEventStreamingRequestSourceSourceKafkaParameters) *CreateEventStreamingRequestSource {
	s.SourceKafkaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceMNSParameters(v *CreateEventStreamingRequestSourceSourceMNSParameters) *CreateEventStreamingRequestSource {
	s.SourceMNSParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceMQTTParameters(v *CreateEventStreamingRequestSourceSourceMQTTParameters) *CreateEventStreamingRequestSource {
	s.SourceMQTTParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceMySQLParameters(v *SourceMySQLParameters) *CreateEventStreamingRequestSource {
	s.SourceMySQLParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceOSSParameters(v *CreateEventStreamingRequestSourceSourceOSSParameters) *CreateEventStreamingRequestSource {
	s.SourceOSSParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceOpenSourceRabbitMQParameters(v *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) *CreateEventStreamingRequestSource {
	s.SourceOpenSourceRabbitMQParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourcePostgreSQLParameters(v *SourcePostgreSQLParameters) *CreateEventStreamingRequestSource {
	s.SourcePostgreSQLParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourcePrometheusParameters(v *CreateEventStreamingRequestSourceSourcePrometheusParameters) *CreateEventStreamingRequestSource {
	s.SourcePrometheusParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceRabbitMQMetaParameters(v *SourceRabbitMQMetaParameters) *CreateEventStreamingRequestSource {
	s.SourceRabbitMQMetaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceRabbitMQMsgSyncParameters(v *SourceRabbitMQMsgSyncParameters) *CreateEventStreamingRequestSource {
	s.SourceRabbitMQMsgSyncParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceRabbitMQParameters(v *CreateEventStreamingRequestSourceSourceRabbitMQParameters) *CreateEventStreamingRequestSource {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceRocketMQCheckpointParameters(v *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) *CreateEventStreamingRequestSource {
	s.SourceRocketMQCheckpointParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceRocketMQParameters(v *CreateEventStreamingRequestSourceSourceRocketMQParameters) *CreateEventStreamingRequestSource {
	s.SourceRocketMQParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceSLSParameters(v *CreateEventStreamingRequestSourceSourceSLSParameters) *CreateEventStreamingRequestSource {
	s.SourceSLSParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) Validate() error {
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

type CreateEventStreamingRequestSourceSourceApacheKafkaParameters struct {
	Bootstraps                *string `json:"Bootstraps,omitempty" xml:"Bootstraps,omitempty"`
	ConsumerGroup             *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	NetworkType               *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OffsetReset               *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	SaslMechanism             *string `json:"SaslMechanism,omitempty" xml:"SaslMechanism,omitempty"`
	SaslPassword              *string `json:"SaslPassword,omitempty" xml:"SaslPassword,omitempty"`
	SaslUser                  *string `json:"SaslUser,omitempty" xml:"SaslUser,omitempty"`
	SecurityGroupId           *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityProtocol          *string `json:"SecurityProtocol,omitempty" xml:"SecurityProtocol,omitempty"`
	SslTruststoreCertificates *string `json:"SslTruststoreCertificates,omitempty" xml:"SslTruststoreCertificates,omitempty"`
	Topic                     *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds                *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	ValueDataType             *string `json:"ValueDataType,omitempty" xml:"ValueDataType,omitempty"`
	VpcId                     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceApacheKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetBootstraps() *string {
	return s.Bootstraps
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSaslMechanism() *string {
	return s.SaslMechanism
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSaslPassword() *string {
	return s.SaslPassword
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSaslUser() *string {
	return s.SaslUser
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSecurityProtocol() *string {
	return s.SecurityProtocol
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetSslTruststoreCertificates() *string {
	return s.SslTruststoreCertificates
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetValueDataType() *string {
	return s.ValueDataType
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetBootstraps(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.Bootstraps = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetConsumerGroup(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetNetworkType(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.NetworkType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetOffsetReset(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSaslMechanism(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SaslMechanism = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSaslPassword(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SaslPassword = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSaslUser(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SaslUser = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSecurityGroupId(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSecurityProtocol(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SecurityProtocol = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetSslTruststoreCertificates(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.SslTruststoreCertificates = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetTopic(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetVSwitchIds(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetValueDataType(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.ValueDataType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) SetVpcId(v string) *CreateEventStreamingRequestSourceSourceApacheKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters struct {
	InstanceEndpoint *string   `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	InstancePassword *string   `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	InstanceUsername *string   `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	NetworkType      *string   `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId  *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Topics           []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
	VSwitchId        *string   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId            *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetTopics() []*string {
	return s.Topics
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetInstanceEndpoint(v string) *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetInstancePassword(v string) *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.InstancePassword = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetInstanceUsername(v string) *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.InstanceUsername = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetNetworkType(v string) *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.NetworkType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetSecurityGroupId(v string) *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetTopics(v []*string) *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.Topics = v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetVSwitchId(v string) *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.VSwitchId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) SetVpcId(v string) *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters {
	s.VpcId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters struct {
	// example:
	//
	// "https://examplebucket.oss-cn-hangzhou.aliyuncs.com/testDoc/Old_Homebrew/2024-06-26%2022%3A34%3A08/opt/homebrew/homebrew/Library/Homebrew/test/support/fixtures/cask/AppWithBinary.zip?OSSAccessKeyId=ri&Expires=1725539627&Signature=rb8q3OpV2i3gZJ"
	ConnectorPackageUrl *string                                                                                       `json:"ConnectorPackageUrl,omitempty" xml:"ConnectorPackageUrl,omitempty"`
	ConnectorParameters *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters `json:"ConnectorParameters,omitempty" xml:"ConnectorParameters,omitempty" type:"Struct"`
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

func (s CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) GetConnectorPackageUrl() *string {
	return s.ConnectorPackageUrl
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) GetConnectorParameters() *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	return s.ConnectorParameters
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) GetWorkerParameters() map[string]interface{} {
	return s.WorkerParameters
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) SetConnectorPackageUrl(v string) *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters {
	s.ConnectorPackageUrl = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) SetConnectorParameters(v *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters {
	s.ConnectorParameters = v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) SetWorkerParameters(v map[string]interface{}) *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters {
	s.WorkerParameters = v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters) Validate() error {
	if s.ConnectorParameters != nil {
		if err := s.ConnectorParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters struct {
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
	// example:
	//
	// test-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GetName() *string {
	return s.Name
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) SetConfig(v map[string]interface{}) *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	s.Config = v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) SetName(v string) *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	s.Name = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceCustomizedKafkaParameters struct {
	// example:
	//
	// r-8vb64581862c****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceCustomizedKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceCustomizedKafkaParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaParameters) SetInstanceId(v string) *CreateEventStreamingRequestSourceSourceCustomizedKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceCustomizedKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceDTSParameters struct {
	// The URL and port number of the change tracking instance.
	BrokerUrl *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	// The UNIX timestamp that is generated when the SDK client consumes the first data record.
	//
	// example:
	//
	// 1620962769
	InitCheckPoint *int64 `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	// The consumer group password.
	//
	// example:
	//
	// admin
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The consumer group ID.
	//
	// example:
	//
	// HD3
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The task ID.
	//
	// example:
	//
	// f86e5814-b223-482c-b768-3b873297dade
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the tracked topic of the change tracking instance.
	//
	// example:
	//
	// LTC_CACHE_PRD
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The consumer group username.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceDTSParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceDTSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) GetBrokerUrl() *string {
	return s.BrokerUrl
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) GetInitCheckPoint() *int64 {
	return s.InitCheckPoint
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) GetPassword() *string {
	return s.Password
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) GetSid() *string {
	return s.Sid
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) GetTopic() *string {
	return s.Topic
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) GetUsername() *string {
	return s.Username
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetBrokerUrl(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.BrokerUrl = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetInitCheckPoint(v int64) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.InitCheckPoint = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetPassword(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.Password = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetSid(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.Sid = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetTaskId(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.TaskId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetTopic(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetUsername(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.Username = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceEventBusParameters struct {
	EventBusName  *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventRuleName *string `json:"EventRuleName,omitempty" xml:"EventRuleName,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceEventBusParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceEventBusParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceEventBusParameters) GetEventBusName() *string {
	return s.EventBusName
}

func (s *CreateEventStreamingRequestSourceSourceEventBusParameters) GetEventRuleName() *string {
	return s.EventRuleName
}

func (s *CreateEventStreamingRequestSourceSourceEventBusParameters) SetEventBusName(v string) *CreateEventStreamingRequestSourceSourceEventBusParameters {
	s.EventBusName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceEventBusParameters) SetEventRuleName(v string) *CreateEventStreamingRequestSourceSourceEventBusParameters {
	s.EventRuleName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceEventBusParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceKafkaParameters struct {
	// The group ID of the consumer that subscribes to the topic.
	//
	// example:
	//
	// DEFAULT_GROUP
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The ID of the ApsaraMQ for Kafka instance.
	//
	// example:
	//
	// r-8vb64581862cd814
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type. Default value: Default. The value PublicNetwork specifies a virtual private cloud (VPC).
	//
	// example:
	//
	// Default
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The offset from which messages are consumed.
	//
	// example:
	//
	// latest
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The ID of the region where the ApsaraMQ for Kafka instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group to which the ApsaraMQ for Kafka instance belongs.
	//
	// example:
	//
	// sg-bp1iv19sp1msc7zot4wr
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the topic on the ApsaraMQ for Kafka instance.
	//
	// example:
	//
	// popvip_center_robot_order
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The ID of the vSwitch with which the ApsaraMQ for Kafka instance is associated.
	//
	// example:
	//
	// vsw-bp179l3llg3jjxwrq72hh
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// example:
	//
	// JSON
	ValueDataType *string `json:"ValueDataType,omitempty" xml:"ValueDataType,omitempty"`
	// The ID of the VPC to which the ApsaraMQ for Kafka instance belongs.
	//
	// example:
	//
	// vpc-8vblalsi0vbhizr77cbhu
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) GetNetwork() *string {
	return s.Network
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) GetValueDataType() *string {
	return s.ValueDataType
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetConsumerGroup(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetInstanceId(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetNetwork(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetOffsetReset(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetSecurityGroupId(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetTopic(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetVSwitchIds(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetValueDataType(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.ValueDataType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetVpcId(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceMNSParameters struct {
	// Specifies whether to enable Base64 encoding. Default value: true.
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
	// The ID of the region where the MNS queue resides.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceMNSParameters) GetIsBase64Decode() *bool {
	return s.IsBase64Decode
}

func (s *CreateEventStreamingRequestSourceSourceMNSParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateEventStreamingRequestSourceSourceMNSParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventStreamingRequestSourceSourceMNSParameters) SetIsBase64Decode(v bool) *CreateEventStreamingRequestSourceSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMNSParameters) SetQueueName(v string) *CreateEventStreamingRequestSourceSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMNSParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceMNSParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMNSParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceMQTTParameters struct {
	// example:
	//
	// JSON
	BodyDataType *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance.
	//
	// example:
	//
	// r-bp1b5ncun5lqerzg4r
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the ApsaraMQ for MQTT instance resides.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic from which messages are sent.
	//
	// example:
	//
	// CANAL_VICUTU_UAT
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceMQTTParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceMQTTParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) GetTopic() *string {
	return s.Topic
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) SetBodyDataType(v string) *CreateEventStreamingRequestSourceSourceMQTTParameters {
	s.BodyDataType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) SetInstanceId(v string) *CreateEventStreamingRequestSourceSourceMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceMQTTParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) SetTopic(v string) *CreateEventStreamingRequestSourceSourceMQTTParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceOSSParameters struct {
	// example:
	//
	// bucket_abc
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// \\n
	Delimiter *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	// example:
	//
	// TextLoader
	LoadFormat *string `json:"LoadFormat,omitempty" xml:"LoadFormat,omitempty"`
	// example:
	//
	// single
	LoadMode *string `json:"LoadMode,omitempty" xml:"LoadMode,omitempty"`
	// example:
	//
	// fun/document/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// example:
	//
	// eventbridge_oss_role
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceOSSParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceOSSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) GetDelimiter() *string {
	return s.Delimiter
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) GetLoadFormat() *string {
	return s.LoadFormat
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) GetLoadMode() *string {
	return s.LoadMode
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) GetPrefix() *string {
	return s.Prefix
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) SetBucketName(v string) *CreateEventStreamingRequestSourceSourceOSSParameters {
	s.BucketName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) SetDelimiter(v string) *CreateEventStreamingRequestSourceSourceOSSParameters {
	s.Delimiter = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) SetLoadFormat(v string) *CreateEventStreamingRequestSourceSourceOSSParameters {
	s.LoadFormat = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) SetLoadMode(v string) *CreateEventStreamingRequestSourceSourceOSSParameters {
	s.LoadMode = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) SetPrefix(v string) *CreateEventStreamingRequestSourceSourceOSSParameters {
	s.Prefix = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) SetRoleName(v string) *CreateEventStreamingRequestSourceSourceOSSParameters {
	s.RoleName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOSSParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters struct {
	AuthType        *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	BodyDataType    *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	Endpoint        *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	QueueName       *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetPassword() *string {
	return s.Password
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetUsername() *string {
	return s.Username
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetAuthType(v string) *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.AuthType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetBodyDataType(v string) *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.BodyDataType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetEndpoint(v string) *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.Endpoint = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetNetworkType(v string) *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.NetworkType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetPassword(v string) *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.Password = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetQueueName(v string) *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetSecurityGroupId(v string) *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetUsername(v string) *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.Username = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetVSwitchIds(v string) *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetVirtualHostName(v string) *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) SetVpcId(v string) *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters {
	s.VpcId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourcePrometheusParameters struct {
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
	// json
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	ExternalLabels *string `json:"ExternalLabels,omitempty" xml:"ExternalLabels,omitempty"`
	// The labels.
	//
	// example:
	//
	// __name__=.*
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourcePrometheusParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourcePrometheusParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) GetDataType() *string {
	return s.DataType
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) GetExternalLabels() *string {
	return s.ExternalLabels
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) GetLabels() *string {
	return s.Labels
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) SetClusterId(v string) *CreateEventStreamingRequestSourceSourcePrometheusParameters {
	s.ClusterId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) SetDataType(v string) *CreateEventStreamingRequestSourceSourcePrometheusParameters {
	s.DataType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) SetExternalLabels(v string) *CreateEventStreamingRequestSourceSourcePrometheusParameters {
	s.ExternalLabels = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) SetLabels(v string) *CreateEventStreamingRequestSourceSourcePrometheusParameters {
	s.Labels = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourcePrometheusParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) SetRoleName(v string) *CreateEventStreamingRequestSourceSourcePrometheusParameters {
	s.RoleName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceRabbitMQParameters struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// gtm-cn-k2c2yfgzt02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// demo
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the ApsaraMQ for RabbitMQ instance resides. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/62010.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the vhost to which the ApsaraMQ for RabbitMQ instance belongs.
	//
	// example:
	//
	// eb-connect
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) SetInstanceId(v string) *CreateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) SetQueueName(v string) *CreateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) SetVirtualHostName(v string) *CreateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters struct {
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topics       []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
}

func (s CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) GetTopics() []*string {
	return s.Topics
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) SetInstanceId(v string) *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) SetInstanceType(v string) *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters {
	s.InstanceType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) SetTopics(v []*string) *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters {
	s.Topics = v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQCheckpointParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceRocketMQParameters struct {
	// The authentication method.
	//
	// example:
	//
	// ACL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// JSON
	BodyDataType *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	// The SQL statement that you want to use to filter messages.
	//
	// example:
	//
	// index > 10
	FilterSql *string `json:"FilterSql,omitempty" xml:"FilterSql,omitempty"`
	// The method that you want to use to filter messages.
	//
	// example:
	//
	// Tag
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// The ID of the consumer group on the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// GID_group1
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// The endpoint that you want to use to access the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// registry-vpc.cn-zhangjiakou.aliyuncs.com
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// default_C56C360261515
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type. Valid values:
	//
	// 	- PublicNetwork
	//
	// 	- PrivateNetwork
	//
	// example:
	//
	// PublicNetwork
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The password that you want to use to access the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// admin
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The ID of the security group to which the ApsaraMQ for RocketMQ instance belongs.
	//
	// example:
	//
	// sg-m5edtu24f123456789
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// The type of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// 2
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The username that you want to use to access the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// admin
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The ID of the vSwitch with which the ApsaraMQ for RocketMQ instance is associated.
	//
	// example:
	//
	// vsw-m5ev8asdc6h123456789
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The ID of the VPC to which the ApsaraMQ for RocketMQ instance belongs.
	//
	// example:
	//
	// vpc-m5e3sv4b123456789
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The network type. Valid values: PublicNetwork and PrivateNetwork.
	//
	// example:
	//
	// PrivateNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The offset from which messages are consumed. Valid values:
	//
	// 	- CONSUME_FROM_LAST_OFFSET: Messages are consumed from the latest offset.
	//
	// 	- CONSUME_FROM_FIRST_OFFSET: Messages are consumed from the earliest offset.
	//
	// 	- CONSUME_FROM_TIMESTAMP: Messages are consumed from the offset at the specified point in time.
	//
	// Default value: CONSUME_FROM_LAST_OFFSET.
	//
	// example:
	//
	// CONSUMEFROMLAST_OFFSET
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the region where the ApsaraMQ for RocketMQ instance resides.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID of the cross-border task.
	//
	// example:
	//
	// sg-m5edtu24f123456789
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The tag that you want to use to filter messages.
	//
	// example:
	//
	// test
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp that indicates the time from which messages are consumed. This parameter is valid only if you set Offset to CONSUME_FROM_TIMESTAMP.
	//
	// example:
	//
	// 1670656652009
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The name of the topic on the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// Topic_publicRule_api_1667273421288
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The vSwitch ID of the cross-border task.
	//
	// example:
	//
	// vsw-m5ev8asdc6h123456789
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The VPC ID of the cross-border task.
	//
	// example:
	//
	// vpc-m5e3sv4b123456789
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetFilterSql() *string {
	return s.FilterSql
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetFilterType() *string {
	return s.FilterType
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetGroupID() *string {
	return s.GroupID
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceNetwork() *string {
	return s.InstanceNetwork
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceSecurityGroupId() *string {
	return s.InstanceSecurityGroupId
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceVSwitchIds() *string {
	return s.InstanceVSwitchIds
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetInstanceVpcId() *string {
	return s.InstanceVpcId
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetNetwork() *string {
	return s.Network
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetOffset() *string {
	return s.Offset
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetTag() *string {
	return s.Tag
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetTopic() *string {
	return s.Topic
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetAuthType(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetBodyDataType(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.BodyDataType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetFilterSql(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.FilterSql = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetFilterType(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.FilterType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetGroupID(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceEndpoint(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceNetwork(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstancePassword(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceType(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceUsername(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceVpcId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetNetwork(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Network = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetOffset(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetSecurityGroupId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetTag(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetTimestamp(v int64) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetTopic(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetVSwitchIds(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetVpcId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.VpcId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestSourceSourceSLSParameters struct {
	// The consumer offset. The value begin specifies the earliest offset. The value end specifies the latest offset. You can also specify a time in seconds to start consumption.
	//
	// example:
	//
	// end
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The Simple Log Service Logstore.
	//
	// example:
	//
	// sas-log
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Simple Log Service project.
	//
	// example:
	//
	// test
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Simple Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console.
	//
	// example:
	//
	// testRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceSLSParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) GetConsumePosition() *string {
	return s.ConsumePosition
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) GetLogStore() *string {
	return s.LogStore
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) GetProject() *string {
	return s.Project
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) SetConsumePosition(v string) *CreateEventStreamingRequestSourceSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) SetLogStore(v string) *CreateEventStreamingRequestSourceSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) SetProject(v string) *CreateEventStreamingRequestSourceSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) SetRoleName(v string) *CreateEventStreamingRequestSourceSourceSLSParameters {
	s.RoleName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestTags) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateEventStreamingRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingRequestTags) SetKey(v string) *CreateEventStreamingRequestTags {
	s.Key = &v
	return s
}

func (s *CreateEventStreamingRequestTags) SetValue(v string) *CreateEventStreamingRequestTags {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingRequestTags) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingRequestTransforms struct {
	// example:
	//
	// acs:fc:cn-hangzhou:*****:services/demo-service.LATEST/functions/demo-func
	Arn                             *string                          `json:"Arn,omitempty" xml:"Arn,omitempty"`
	BaiLianAgentTransformParameters *BaiLianAgentTransformParameters `json:"BaiLianAgentTransformParameters,omitempty" xml:"BaiLianAgentTransformParameters,omitempty"`
	DashScopeTransformParameters    *DashScopeTransformParameters    `json:"DashScopeTransformParameters,omitempty" xml:"DashScopeTransformParameters,omitempty"`
	EmbeddingTransformParameters    *EmbeddingTransformParameters    `json:"EmbeddingTransformParameters,omitempty" xml:"EmbeddingTransformParameters,omitempty"`
}

func (s CreateEventStreamingRequestTransforms) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingRequestTransforms) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestTransforms) GetArn() *string {
	return s.Arn
}

func (s *CreateEventStreamingRequestTransforms) GetBaiLianAgentTransformParameters() *BaiLianAgentTransformParameters {
	return s.BaiLianAgentTransformParameters
}

func (s *CreateEventStreamingRequestTransforms) GetDashScopeTransformParameters() *DashScopeTransformParameters {
	return s.DashScopeTransformParameters
}

func (s *CreateEventStreamingRequestTransforms) GetEmbeddingTransformParameters() *EmbeddingTransformParameters {
	return s.EmbeddingTransformParameters
}

func (s *CreateEventStreamingRequestTransforms) SetArn(v string) *CreateEventStreamingRequestTransforms {
	s.Arn = &v
	return s
}

func (s *CreateEventStreamingRequestTransforms) SetBaiLianAgentTransformParameters(v *BaiLianAgentTransformParameters) *CreateEventStreamingRequestTransforms {
	s.BaiLianAgentTransformParameters = v
	return s
}

func (s *CreateEventStreamingRequestTransforms) SetDashScopeTransformParameters(v *DashScopeTransformParameters) *CreateEventStreamingRequestTransforms {
	s.DashScopeTransformParameters = v
	return s
}

func (s *CreateEventStreamingRequestTransforms) SetEmbeddingTransformParameters(v *EmbeddingTransformParameters) *CreateEventStreamingRequestTransforms {
	s.EmbeddingTransformParameters = v
	return s
}

func (s *CreateEventStreamingRequestTransforms) Validate() error {
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
	if s.EmbeddingTransformParameters != nil {
		if err := s.EmbeddingTransformParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}
