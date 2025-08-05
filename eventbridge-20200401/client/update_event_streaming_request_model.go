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
	// The rule that is used to filter events. If you leave this parameter empty, all events are matched.
	//
	// This parameter is required.
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
	// The parameters that are configured for the runtime environment.
	RunOptions *UpdateEventStreamingRequestRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	// The event target. You must and can specify only one event target.
	//
	// This parameter is required.
	Sink *UpdateEventStreamingRequestSink `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	// The event provider, which is also known as the event source. You must and can specify only one event source.
	//
	// This parameter is required.
	Source     *UpdateEventStreamingRequestSource       `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestRunOptions struct {
	// The batch window.
	BatchWindow    *UpdateEventStreamingRequestRunOptionsBatchWindow    `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	BusinessOption *UpdateEventStreamingRequestRunOptionsBusinessOption `json:"BusinessOption,omitempty" xml:"BusinessOption,omitempty" type:"Struct"`
	// Specifies whether to enable dead-letter queues. By default, dead-letter queues are disabled. Events that fail to be pushed are discarded after the maximum number of retries that is specified by the retry policy is reached.
	DeadLetterQueue *UpdateEventStreamingRequestRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The exception tolerance policy. Valid values: NONE and ALL.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestRunOptionsBatchWindow struct {
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
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue.
	//
	// example:
	//
	// acs:ram::1317334647812936:role/rdstoecsassumekms
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The network type of the dead-letter queue. Valid values:
	//
	// 	- PrivateNetwork
	//
	// 	- PublicNetwork
	//
	// example:
	//
	// PrivateNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-2vcgdxz7o1n9zapp****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-m5ev8asdc6h12345****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The VPC ID.
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
	// The retry policy. Valid values: BACKOFF_RETRY and EXPONENTIAL_DECAY_RETRY.
	//
	// example:
	//
	// BACKOFFRETRY
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
	// The parameters that are configured if you specify Apache RocketMQ (Offset Data) as the event target.
	SinkApacheRocketMQCheckpointParameters *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters `json:"SinkApacheRocketMQCheckpointParameters,omitempty" xml:"SinkApacheRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify BaiLian as the event target.
	SinkBaiLianParameters *SinkBaiLianParameters `json:"SinkBaiLianParameters,omitempty" xml:"SinkBaiLianParameters,omitempty"`
	// The parameters that are configured if you specify Kafka Sink Connect as the event target.
	SinkCustomizedKafkaConnectorParameters *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParameters `json:"SinkCustomizedKafkaConnectorParameters,omitempty" xml:"SinkCustomizedKafkaConnectorParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Kafka Source Connect as the event target.
	SinkCustomizedKafkaParameters *UpdateEventStreamingRequestSinkSinkCustomizedKafkaParameters `json:"SinkCustomizedKafkaParameters,omitempty" xml:"SinkCustomizedKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify DashVector as the event target.
	SinkDashVectorParameters *UpdateEventStreamingRequestSinkSinkDashVectorParameters `json:"SinkDashVectorParameters,omitempty" xml:"SinkDashVectorParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify DataHub as the event target.
	SinkDataHubParameters *UpdateEventStreamingRequestSinkSinkDataHubParameters `json:"SinkDataHubParameters,omitempty" xml:"SinkDataHubParameters,omitempty" type:"Struct"`
	// The type of the event source.
	SinkDorisParameters *UpdateEventStreamingRequestSinkSinkDorisParameters `json:"SinkDorisParameters,omitempty" xml:"SinkDorisParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Function Compute as the event target.
	SinkFcParameters *UpdateEventStreamingRequestSinkSinkFcParameters `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify CloudFlow as the event target.
	SinkFnfParameters *UpdateEventStreamingRequestSinkSinkFnfParameters `json:"SinkFnfParameters,omitempty" xml:"SinkFnfParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for Kafka as the event target.
	SinkKafkaParameters *UpdateEventStreamingRequestSinkSinkKafkaParameters `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Simple Message Queue (SMQ, formerly MNS) as the event target.
	SinkMNSParameters *UpdateEventStreamingRequestSinkSinkMNSParameters `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify open source RabbitMQ as the event target.
	SinkOpenSourceRabbitMQParameters *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters `json:"SinkOpenSourceRabbitMQParameters,omitempty" xml:"SinkOpenSourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Managed Service for Prometheus as the event target.
	SinkPrometheusParameters *UpdateEventStreamingRequestSinkSinkPrometheusParameters `json:"SinkPrometheusParameters,omitempty" xml:"SinkPrometheusParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for RabbitMQ as the event target.
	SinkRabbitMQParameters *UpdateEventStreamingRequestSinkSinkRabbitMQParameters `json:"SinkRabbitMQParameters,omitempty" xml:"SinkRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for RocketMQ (Offset Data) as the event target.
	SinkRocketMQCheckpointParameters *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParameters `json:"SinkRocketMQCheckpointParameters,omitempty" xml:"SinkRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for RocketMQ as the event target.
	SinkRocketMQParameters *UpdateEventStreamingRequestSinkSinkRocketMQParameters `json:"SinkRocketMQParameters,omitempty" xml:"SinkRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Simple Log Service as the event target.
	SinkSLSParameters *UpdateEventStreamingRequestSinkSinkSLSParameters `json:"SinkSLSParameters,omitempty" xml:"SinkSLSParameters,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSink) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSink) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSink) GetSinkApacheRocketMQCheckpointParameters() *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters {
	return s.SinkApacheRocketMQCheckpointParameters
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

func (s *UpdateEventStreamingRequestSink) GetSinkDorisParameters() *UpdateEventStreamingRequestSinkSinkDorisParameters {
	return s.SinkDorisParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkFcParameters() *UpdateEventStreamingRequestSinkSinkFcParameters {
	return s.SinkFcParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkFnfParameters() *UpdateEventStreamingRequestSinkSinkFnfParameters {
	return s.SinkFnfParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkKafkaParameters() *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	return s.SinkKafkaParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkMNSParameters() *UpdateEventStreamingRequestSinkSinkMNSParameters {
	return s.SinkMNSParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkOpenSourceRabbitMQParameters() *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParameters {
	return s.SinkOpenSourceRabbitMQParameters
}

func (s *UpdateEventStreamingRequestSink) GetSinkPrometheusParameters() *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	return s.SinkPrometheusParameters
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

func (s *UpdateEventStreamingRequestSink) SetSinkApacheRocketMQCheckpointParameters(v *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters) *UpdateEventStreamingRequestSink {
	s.SinkApacheRocketMQCheckpointParameters = v
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

func (s *UpdateEventStreamingRequestSink) SetSinkDorisParameters(v *UpdateEventStreamingRequestSinkSinkDorisParameters) *UpdateEventStreamingRequestSink {
	s.SinkDorisParameters = v
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

func (s *UpdateEventStreamingRequestSink) SetSinkKafkaParameters(v *UpdateEventStreamingRequestSinkSinkKafkaParameters) *UpdateEventStreamingRequestSink {
	s.SinkKafkaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkMNSParameters(v *UpdateEventStreamingRequestSinkSinkMNSParameters) *UpdateEventStreamingRequestSink {
	s.SinkMNSParameters = v
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParameters struct {
	// The timestamp that specifies the time from which messages are consumed.
	ConsumeTimestamp *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp `json:"ConsumeTimestamp,omitempty" xml:"ConsumeTimestamp,omitempty" type:"Struct"`
	// The ID of the consumer group.
	Group *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The endpoint that is used to access the Apache RocketMQ instance.
	//
	// example:
	//
	// 192.168.1.1:9876
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The password that is used to access the Apache RocketMQ instance.
	//
	// example:
	//
	// ****
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The username that is used to access the Apache RocketMQ instance.
	//
	// example:
	//
	// admin
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The network type.
	//
	// 	- PublicNetwork
	//
	// 	- PrivateNetwork
	//
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-2ze5bmpw6adn0q******
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the topic on the Apache RocketMQ instance.
	Topic *UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf62oqt1twuevrt******
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The timestamp that specifies the time from which messages are consumed.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the consumer group.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic on the Apache RocketMQ instance.
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
	// The download link of the ZIP package that includes Object Storage Service (OSS) resources.
	//
	// example:
	//
	// "https://examplebucket.oss-cn-hangzhou.aliyuncs.com/testDoc/Old_Homebrew/2024-06-26%2022%3A34%3A08/opt/homebrew/homebrew/Library/Homebrew/test/support/fixtures/cask/AppWithBinary.zip?OSSAccessKeyId=ri&Expires=1725539627&Signature=rb8q3OpV2i3gZJ"
	ConnectorPackageUrl *string `json:"ConnectorPackageUrl,omitempty" xml:"ConnectorPackageUrl,omitempty"`
	// The parameters that are configured for the parsing of the .properties file in the ZIP package.
	ConnectorParameters *UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters `json:"ConnectorParameters,omitempty" xml:"ConnectorParameters,omitempty" type:"Struct"`
	// The instance configurations.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkCustomizedKafkaConnectorParametersConnectorParameters struct {
	// The connector configurations.
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
	// The ID of the ApsaraMQ for Kafka instance.
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
	// The API key that you want to create in the DashVector console.
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
	// The parameters in the Schema field of the table when data is inserted into DashVector. After the event content is transformed, the data must be in JSON format.
	DashVectorSchemaParameters *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters `json:"DashVectorSchemaParameters,omitempty" xml:"DashVectorSchemaParameters,omitempty" type:"Struct"`
	// The ID of the DashVector instance.
	//
	// example:
	//
	// vrs-cn-lbj3ru1***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type. Valid values:
	//
	// 	- PrivateNetwork
	//
	// 	- PublicNetwork
	//
	// example:
	//
	// PublicNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The type of operation that you want to perform on the DashVector database. Valid values:
	//
	// 	- Delete
	//
	// 	- Upsert
	//
	// example:
	//
	// Upsert
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The partition. Default value: default.
	Partition *UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition `json:"Partition,omitempty" xml:"Partition,omitempty" type:"Struct"`
	// The ID of the primary key that you want to use when you insert or delete records. If you do not specify this parameter, a random primary key ID is returned.
	PrimaryKeyId *UpdateEventStreamingRequestSinkSinkDashVectorParametersPrimaryKeyId `json:"PrimaryKeyId,omitempty" xml:"PrimaryKeyId,omitempty" type:"Struct"`
	// The vector that is recorded when data is inserted into DashVector.
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

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) GetDashVectorSchemaParameters() *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
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

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParameters) SetDashVectorSchemaParameters(v *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) *UpdateEventStreamingRequestSinkSinkDashVectorParameters {
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters struct {
	// The method that you want to use to transform events. Valid values:
	//
	// 	- JSONPATH
	//
	// 	- CONSTANT
	//
	// 	- TEMPLATE
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The schema template. This parameter is required only if you set Form to TEMPLATE. After the event content is transformed, the data must be an array in the JSON format. Each schema corresponds to a JSON object. The properties include only the name, type, and value fields. The value of the type field can be only of the INT, FLOAT, STRING, or BOOL type.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "name": "schema1",
	//
	//     "type": "INT",
	//
	//     "value": "${value1}"
	//
	//   },
	//
	//   {
	//
	//     "name": "schema2",
	//
	//     "type": "FLOAT",
	//
	//     "value": "${value2}"
	//
	//   }
	//
	// ]
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// 	- If you set Form to CONSTANT, specify a constant.
	//
	// 	- If you set Form to JSONPATH, specify a JSONPath rule.
	//
	// 	- If you set Form to TEMPLATE, specify variables for the template.
	//
	// >  The value of this parameter cannot exceed 10,240 characters in length.
	//
	// example:
	//
	// {
	//
	//   "value1":"v1",
	//
	//   "value2":"v2"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GetForm() *string {
	return s.Form
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GetTemplate() *string {
	return s.Template
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) GetValue() *string {
	return s.Value
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) SetForm(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) SetValue(v string) *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Value = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkDashVectorParametersDashVectorSchemaParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDashVectorParametersPartition struct {
	// The method that you want to use to transform events. Valid values:
	//
	// 	- JSONPATH
	//
	// 	- CONSTANT
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// 	- If you set Form to CONSTANT, specify a constant.
	//
	// 	- If you set Form to JSONPATH, specify a JSONPath rule.
	//
	// >  The value of this parameter cannot exceed 10,240 characters in length.
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
	// The method that you want to use to transform events. Valid values:
	//
	// 	- JSONPATH
	//
	// 	- TEMPLATE
	//
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template that you want to use to specify primary key IDs. This parameter is required only if you set Form to TEMPLATE.
	//
	// example:
	//
	// ${ID}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// If you set Form to JSONPATH, specify a JSONPath rule. If you set Form to TEMPLATE, specify variables for the template.
	//
	// >  The value of this parameter cannot exceed 10,240 characters in length.
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
	// The method that you want to use to transform events.
	//
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The JSONPath rule that you want to use to extract content.
	//
	// >  The value of this parameter cannot exceed 10,240 characters in length.
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
	// The data is of the BLOB type, and a template is defined for the record.
	Body *UpdateEventStreamingRequestSinkSinkDataHubParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The key-value pair of custom logs. This parameter takes effect only if you set ContentType to KeyValue. Each key-value pair is in the Key_n, Value_n format.
	ContentSchema *UpdateEventStreamingRequestSinkSinkDataHubParametersContentSchema `json:"ContentSchema,omitempty" xml:"ContentSchema,omitempty" type:"Struct"`
	// The data format. Valid values:
	//
	// 	- JSON
	//
	// 	- KeyValue
	ContentType *UpdateEventStreamingRequestSinkSinkDataHubParametersContentType `json:"ContentType,omitempty" xml:"ContentType,omitempty" type:"Struct"`
	// The name of the DataHub project.
	Project *UpdateEventStreamingRequestSinkSinkDataHubParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The role name.
	RoleName *UpdateEventStreamingRequestSinkSinkDataHubParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The name of the DataHub topic.
	Topic *UpdateEventStreamingRequestSinkSinkDataHubParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The data is of the TUBLE type, and a schema is defined for the DataHub topic.
	TopicSchema *UpdateEventStreamingRequestSinkSinkDataHubParametersTopicSchema `json:"TopicSchema,omitempty" xml:"TopicSchema,omitempty" type:"Struct"`
	// The data type of the DataHub topic. Valid values:
	//
	// 	- TUPLE
	//
	// 	- BLOB
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkDataHubParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// ORIGINAL
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The data is of the BLOB type, and a template is defined for the record.
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
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// {"k1":"${k1}","k2":"${k2}"}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The data is of the TUBLE type, and a schema is defined for the DataHub topic.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The data type of the DataHub topic. Valid values:
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
	return dara.Validate(s)
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

type UpdateEventStreamingRequestSinkSinkFcParameters struct {
	// The message body that you want to deliver to Function Compute.
	Body *UpdateEventStreamingRequestSinkSinkFcParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The delivery concurrency. Minimum value: 1.
	Concurrency *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	// The rule that you want to use to transform the format of event content.
	DataFormat *UpdateEventStreamingRequestSinkSinkFcParametersDataFormat `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" type:"Struct"`
	// The function name.
	FunctionName *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	// The invocation mode. Valid values: Sync and Async.
	InvocationType *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	// The alias of the service to which the function belongs.
	Qualifier *UpdateEventStreamingRequestSinkSinkFcParametersQualifier `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	// The service name.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFcParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
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
	// The method that you want to use to transform events. Valid values:
	//
	// 	- ORIGINAL: complete event
	//
	// 	- JSONPATH: partial event
	//
	// 	- CONSTANT: constant
	//
	// 	- TEMPLATE: template
	//
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// $.data.key
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The invocation mode.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The service name.
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
	// The input information of the execution.
	Input *UpdateEventStreamingRequestSinkSinkFnfParametersInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The role name.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The input information of the execution.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name.
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
	// The acknowledgment (ACK) mode.
	//
	// 	- If you set this parameter to 0, no response is returned from the broker. In this mode, the performance is high, but the risk of data loss is also high.
	//
	// 	- If you set this parameter to 1, a response is returned when data is written to the leader. In this mode, the performance and the risk of data loss are moderate. Data loss may occur if a failure occurs on the leader.
	//
	// 	- If you set this parameter to all, a response is returned when data is written to the leader and synchronized to the followers. In this mode, the performance is low, but the risk of data loss is also low. Data loss occurs if the leader and the followers fail at the same time.
	Acks *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	// The metadata added to messages in the ApsaraMQ for Kafka instance.
	Headers *UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The ID of the ApsaraMQ for Kafka instance.
	InstanceId *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message key.
	Key *UpdateEventStreamingRequestSinkSinkKafkaParametersKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	// The name of the topic on the ApsaraMQ for Kafka instance.
	Topic *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The message body.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersAcks struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
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

type UpdateEventStreamingRequestSinkSinkKafkaParametersHeaders struct {
	// The method that you want to use to transform events. Valid values:
	//
	// 	- ORIGINAL: complete event
	//
	// 	- JSONPATH: partial event
	//
	// 	- CONSTANT: constant
	//
	// 	- TEMPLATE: template
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the ApsaraMQ for Kafka instance.
	//
	// example:
	//
	// Defaut_1283278472_sadkj
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The message key.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic on the ApsaraMQ for Kafka instance.
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
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The message body.
	Body *UpdateEventStreamingRequestSinkSinkMNSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// Specifies whether to enable Base64 encoding.
	IsBase64Encode *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	// The name of the SMQ queue.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkMNSParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the SMQ queue.
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
	// 	- ACL
	//
	// 	- N/A
	//
	// example:
	//
	// ACL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The message body.
	Body *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The endpoint used to access the open source RabbitMQ instance.
	//
	// example:
	//
	// 192.168.1.1:9876
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The name of the exchange on the open source RabbitMQ instance. This parameter is valid only if you set TargetType to Exchange.
	//
	// example:
	//
	// my-exchange
	Exchange *string `json:"Exchange,omitempty" xml:"Exchange,omitempty"`
	// The message ID.
	MessageId *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// The network type. Valid values:
	//
	// 	- PrivateNetwork
	//
	// 	- PublicNetwork
	//
	// example:
	//
	// PublicNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The password that is used to access the open source RabbitMQ instance.
	//
	// example:
	//
	// ****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The attributes of the message.
	Properties *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The name of the queue on the open source RabbitMQ instance. This parameter is valid only if you set TargetType to Queue.
	//
	// example:
	//
	// my-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The routing key.
	RoutingKey *UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-uf6of9452b2pba82c ****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The type of the resource to which you want to deliver messages. Valid values:
	//
	// 	- **Exchange**: Messages are routed to the event target using an exchange.
	//
	// 	- **Queue**: Messages are delivered to a specific queue.
	//
	// example:
	//
	// Exchange
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The username that is used to access the open source RabbitMQ instance.
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
	// The name of the virtual host of the open source RabbitMQ instance.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkOpenSourceRabbitMQParametersBody struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value of the raw data.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value of the message ID.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The attribute value.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The routing key.
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
	// The authentication method.
	AuthorizationType *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty" type:"Struct"`
	// The metric data.
	Data *UpdateEventStreamingRequestSinkSinkPrometheusParametersData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The parameters that are configured for the request header.
	HeaderParameters *UpdateEventStreamingRequestSinkSinkPrometheusParametersHeaderParameters `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Struct"`
	// The network type.
	NetworkType *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	// The password.
	Password *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword `json:"Password,omitempty" xml:"Password,omitempty" type:"Struct"`
	// The ID of the security group.
	SecurityGroupId *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	// The URL for the remote write configuration item of the Managed Service for Prometheus instance.
	URL *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL `json:"URL,omitempty" xml:"URL,omitempty" type:"Struct"`
	// The username.
	Username *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername `json:"Username,omitempty" xml:"Username,omitempty" type:"Struct"`
	// The vSwitch ID.
	VSwitchId *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Struct"`
	// The VPC ID.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType struct {
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
	// The method that you want to use to transform events. Default value: JSONPATH.
	//
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The metric data.
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
	// The method that you want to use to transform events.
	//
	// **Valid values:**
	//
	// 	- JSONPATH
	//
	// 	- CONSTANT
	//
	// 	- TEMPLATE
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template that you want to use for HTTP request headers. This parameter is required only if you set Form to TEMPLATE. After the event content is transformed, the data must be in JSON format.
	//
	// example:
	//
	// {
	//
	//     "user_name":"${name}"
	//
	// }
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// 	- If you set Form to CONSTANT, specify a constant.
	//
	// 	- If you set Form to JSONPATH, specify a JSONPath rule.
	//
	// 	- If you set Form to TEMPLATE, specify variables for the template.
	//
	// Note: The value of this parameter cannot exceed 10,240 characters in length.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The network type. Valid values:
	//
	// 	- PublicNetwork
	//
	// 	- PrivateNetwork
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the security group.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
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
	// The message body.
	Body *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The exchange mode. This parameter is required only if you set TargetType to Exchange.
	Exchange *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
	InstanceId *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message ID.
	MessageId *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// The attributes that you want to use to filter messages.
	Properties *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The queue mode. This parameter is required only if you set TargetType to Queue.
	QueueName *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	// The rule that you want to use to route messages. This parameter is required only if you set TargetType to Exchange.
	RoutingKey *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// The type of the resource to which you want to deliver events.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the exchange on the ApsaraMQ for RabbitMQ instance.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
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
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the queue on the ApsaraMQ for RabbitMQ instance.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The rule that you want to use to route messages.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
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
	// The timestamp that specifies the time from which messages are consumed.
	ConsumeTimestamp *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp `json:"ConsumeTimestamp,omitempty" xml:"ConsumeTimestamp,omitempty" type:"Struct"`
	// The ID of the consumer group.
	Group *UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The ID of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// MQ_INST_164901546557****_BAAN****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type. Valid values:
	//
	// 	- Cloud_4: ApsaraMQ for RocketMQ 4.0 instance
	//
	// 	- Cloud_5: ApsaraMQ for RocketMQ 5.0 instance
	//
	// example:
	//
	// Cloud_4
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The name of the topic on the ApsaraMQ for RocketMQ instance.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQCheckpointParametersConsumeTimestamp struct {
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The timestamp that specifies the time from which messages are consumed.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the consumer group.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic on the ApsaraMQ for RocketMQ instance.
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
	// The message body.
	Body *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The type of the message delivery order. This parameter is optional. Default value: Concurrently.
	DeliveryOrderType *UpdateEventStreamingRequestSinkSinkRocketMQParametersDeliveryOrderType `json:"DeliveryOrderType,omitempty" xml:"DeliveryOrderType,omitempty" type:"Struct"`
	// The endpoint that is used to access the instance.
	InstanceEndpoint *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty" type:"Struct"`
	// The ID of the ApsaraMQ for RocketMQ instance.
	InstanceId *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The password that is used to access the instance.
	InstancePassword *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty" type:"Struct"`
	// The instance type.
	InstanceType *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Struct"`
	// The username that is used to access the instance.
	InstanceUsername *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty" type:"Struct"`
	// The keys that you want to use to filter messages.
	Keys *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	// The network type. Valid values:
	//
	// 	- PublicNetwork
	//
	// 	- PrivateNetwork
	Network *UpdateEventStreamingRequestSinkSinkRocketMQParametersNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The attributes that you want to use to filter messages.
	Properties *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The ID of the security group.
	SecurityGroupId *UpdateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	// The sharding key.
	//
	// >  If you set DeliveryOrderType to Orderly, this parameter is required. If you specify ApsaraMQ for RocketMQ as the event source, you can leave this parameter empty. In this case, the combined value of BrokerName and QueueId is used as the sharding key.
	ShardingKey *UpdateEventStreamingRequestSinkSinkRocketMQParametersShardingKey `json:"ShardingKey,omitempty" xml:"ShardingKey,omitempty" type:"Struct"`
	// The tags that you want to use to filter messages.
	Tags *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The name of the topic on the ApsaraMQ for RocketMQ instance.
	Topic *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The vSwitch ID.
	VSwitchIds *UpdateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	// The virtual private cloud (VPC) ID.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the message delivery order. Valid values:
	//
	// 	- **Orderly**
	//
	// 	- **Concurrently**
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The endpoint that is used to access the instance.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The password that is used to access the instance.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance type. Valid values:
	//
	// 	- Cloud_4 (default): ApsaraMQ for RocketMQ 4.0 instance
	//
	// 	- Cloud_5: ApsaraMQ for RocketMQ 5.0 instance
	//
	// 	- SelfBuilt: self-managed Apache RocketMQ cluster
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The username that is used to access the instance.
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
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The network type. Valid values:
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
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the security group.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value of the sharding key.
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
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic on the ApsaraMQ for RocketMQ instance.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
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
	// The message body that you want to deliver to Simple Log Service.
	Body *UpdateEventStreamingRequestSinkSinkSLSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The key-value pair of custom logs. This parameter takes effect only if you set ContentType to KeyValue. Each key-value pair is in the Key_n, Value_n format.
	ContentSchema *UpdateEventStreamingRequestSinkSinkSLSParametersContentSchema `json:"ContentSchema,omitempty" xml:"ContentSchema,omitempty" type:"Struct"`
	// The format of the Simple Log Service data. Valid values:
	//
	// 	- JSON
	//
	// 	- KeyValue
	ContentType *UpdateEventStreamingRequestSinkSinkSLSParametersContentType `json:"ContentType,omitempty" xml:"ContentType,omitempty" type:"Struct"`
	// The Simple Log Service Logstore.
	LogStore *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	// The Simple Log Service project.
	Project *UpdateEventStreamingRequestSinkSinkSLSParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Simple Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console.
	RoleName *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The topic that you want to use to store logs. This parameter corresponds to the reserved field topic in Simple Log Service.
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSinkSinkSLSParametersBody struct {
	// The method that you want to use to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
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
	// The method that you want to use to transform events.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The key-value pair of custom logs.
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
	// The method that you want to use to transform events.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want events to be transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The format of the Simple Log Service data.
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
	// The method that you want to use to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic that you want to use to store logs. This parameter corresponds to the reserved field topic in Simple Log Service.
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
	SourceApacheRocketMQCheckpointParameters *UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters `json:"SourceApacheRocketMQCheckpointParameters,omitempty" xml:"SourceApacheRocketMQCheckpointParameters,omitempty" type:"Struct"`
	SourceCustomizedKafkaConnectorParameters *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParameters `json:"SourceCustomizedKafkaConnectorParameters,omitempty" xml:"SourceCustomizedKafkaConnectorParameters,omitempty" type:"Struct"`
	SourceCustomizedKafkaParameters          *UpdateEventStreamingRequestSourceSourceCustomizedKafkaParameters          `json:"SourceCustomizedKafkaParameters,omitempty" xml:"SourceCustomizedKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Data Transmission Service (DTS) as the event source.
	SourceDTSParameters      *UpdateEventStreamingRequestSourceSourceDTSParameters      `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	SourceEventBusParameters *UpdateEventStreamingRequestSourceSourceEventBusParameters `json:"SourceEventBusParameters,omitempty" xml:"SourceEventBusParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for Kafka as the event source.
	SourceKafkaParameters *UpdateEventStreamingRequestSourceSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Service (MNS) as the event source.
	SourceMNSParameters *UpdateEventStreamingRequestSourceSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for MQTT as the event source.
	SourceMQTTParameters               *UpdateEventStreamingRequestSourceSourceMQTTParameters               `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	SourceMySQLParameters              *SourceMySQLParameters                                               `json:"SourceMySQLParameters,omitempty" xml:"SourceMySQLParameters,omitempty"`
	SourceOSSParameters                *UpdateEventStreamingRequestSourceSourceOSSParameters                `json:"SourceOSSParameters,omitempty" xml:"SourceOSSParameters,omitempty" type:"Struct"`
	SourceOpenSourceRabbitMQParameters *UpdateEventStreamingRequestSourceSourceOpenSourceRabbitMQParameters `json:"SourceOpenSourceRabbitMQParameters,omitempty" xml:"SourceOpenSourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Managed Service for Prometheus as the event source.
	SourcePrometheusParameters *UpdateEventStreamingRequestSourceSourcePrometheusParameters `json:"SourcePrometheusParameters,omitempty" xml:"SourcePrometheusParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for RabbitMQ as the event source.
	SourceRabbitMQParameters           *UpdateEventStreamingRequestSourceSourceRabbitMQParameters           `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	SourceRocketMQCheckpointParameters *UpdateEventStreamingRequestSourceSourceRocketMQCheckpointParameters `json:"SourceRocketMQCheckpointParameters,omitempty" xml:"SourceRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify ApsaraMQ for RocketMQ as the event source.
	SourceRocketMQParameters *UpdateEventStreamingRequestSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Simple Log Service as the event source.
	SourceSLSParameters *UpdateEventStreamingRequestSourceSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingRequestSource) GoString() string {
	return s.String()
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

func (s *UpdateEventStreamingRequestSource) GetSourcePrometheusParameters() *UpdateEventStreamingRequestSourceSourcePrometheusParameters {
	return s.SourcePrometheusParameters
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

func (s *UpdateEventStreamingRequestSource) SetSourcePrometheusParameters(v *UpdateEventStreamingRequestSourceSourcePrometheusParameters) *UpdateEventStreamingRequestSource {
	s.SourcePrometheusParameters = v
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceApacheRocketMQCheckpointParameters struct {
	// example:
	//
	// 192.168.1.1:9876
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// example:
	//
	// ****
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// example:
	//
	// admin
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-mw43*****
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Topics          []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
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
	// example:
	//
	// "https://examplebucket.oss-cn-hangzhou.aliyuncs.com/testDoc/Old_Homebrew/2024-06-26%2022%3A34%3A08/opt/homebrew/homebrew/Library/Homebrew/test/support/fixtures/cask/AppWithBinary.zip?OSSAccessKeyId=ri&Expires=1725539627&Signature=rb8q3OpV2i3gZJ"
	ConnectorPackageUrl *string                                                                                       `json:"ConnectorPackageUrl,omitempty" xml:"ConnectorPackageUrl,omitempty"`
	ConnectorParameters *UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters `json:"ConnectorParameters,omitempty" xml:"ConnectorParameters,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceCustomizedKafkaConnectorParametersConnectorParameters struct {
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
	// hkprdb
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
	// example:
	//
	// my-event-bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
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
	// i-8vbh4a5b9yfhgkkzm98f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network setting. Default value: Default. The value PublicNetwork specifies a virtual private cloud (VPC).
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
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group to which the ApsaraMQ for Kafka instance belongs.
	//
	// example:
	//
	// sg-uf6jcm3y5hcs7hklytxh
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the topic on the ApsaraMQ for Kafka instance.
	//
	// example:
	//
	// topic_empower_1641539400786
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The ID of the vSwitch with which the ApsaraMQ for Kafka instance is associated.
	//
	// example:
	//
	// vsw-wz9t1l1e8eu2omwjazmtm
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The encoding or decoding method. Valid values: Json, Text, and Binary. The value Json specifies that binary data is decoded into strings based on UTF-8 encoding and then parsed into the JSON format. The value Text specifies that binary data is decoded into strings based on UTF-8 encoding and then put into the payload. The value Binary specifies that binary data is encoded into strings based on Base64 encoding and then put into the payload.
	//
	// example:
	//
	// Text
	ValueDataType *string `json:"ValueDataType,omitempty" xml:"ValueDataType,omitempty"`
	// The ID of the VPC to which the ApsaraMQ for Kafka instance belongs.
	//
	// example:
	//
	// vpc-2ze6p0o345nykmekxtuop
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
	// Specifies whether to enable Base64 encoding. Default value: true.
	//
	// example:
	//
	// true
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the MNS queue.
	//
	// example:
	//
	// queue_api_1642474203601
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the MNS queue resides.
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
	// example:
	//
	// JSON
	BodyDataType *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance.
	//
	// example:
	//
	// i-bp1dsudbecqwt61jqswt
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the ApsaraMQ for MQTT instance resides.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic on the ApsaraMQ for MQTT instance.
	//
	// example:
	//
	// topic_empower_1642400400779
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) GetTopic() *string {
	return s.Topic
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetBodyDataType(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.BodyDataType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetTopic(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventStreamingRequestSourceSourceOSSParameters struct {
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
	// example:
	//
	// ACL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// Json
	BodyDataType *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	// example:
	//
	// 192.168.1.1:9876
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// ****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// demo
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// example:
	//
	// sg-m5edtu24f12345****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// example:
	//
	// vsw-m5ev8asdc6h12345****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
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
	// json
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
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
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The name of the queue on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// demo
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the ApsaraMQ for RabbitMQ instance resides.
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
	// example:
	//
	// rmq-cn-jte3w******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Cloud_5
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topics   []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
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
	// The authentication method.
	//
	// example:
	//
	// ACL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// Json
	BodyDataType *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	// example:
	//
	// index > 10
	FilterSql *string `json:"FilterSql,omitempty" xml:"FilterSql,omitempty"`
	// example:
	//
	// Tag
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// The ID of the consumer group on the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// GID_test
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// The endpoint that you want to use to access the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// reg****-vpc.cn-zhangjiakou.aliyuncs.com
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// i-f8z9a9mcgwri1c1idd0e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type of the ApsaraMQ for RocketMQ instance. Valid values:
	//
	// PublicNetwork and PrivateNetwork.
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
	// sg-m5edtu24f12345****
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
	// vsw-m5ev8asdc6h12****
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The ID of the VPC in which the ApsaraMQ for RocketMQ instance is deployed.
	//
	// example:
	//
	// vpc-m5e3sv4b12345****
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// example:
	//
	// PublicNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The offset from which messages are consumed. Valid values:
	//
	// 	- CONSUMEFROMLASTOFFSET: Messages are consumed from the latest offset.
	//
	// 	- CONSUMEFROMFIRSTOFFSET: Messages are consumed from the earliest offset.
	//
	// 	- CONSUMEFROMTIMESTAMP: Messages are consumed from the offset at the specified point in time.
	//
	// Default value: CONSUMEFROMLASTOFFSET.
	//
	// example:
	//
	// CONSUMEFROMLASTOFFSET
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the region where the ApsaraMQ for RocketMQ instance resides.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-m5edtu24f12345****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The tag that you want to use to filter messages.
	//
	// example:
	//
	// test
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp that specifies the time from which messages are consumed. This parameter is valid only if you set Offset to CONSUMEFROMTIMESTAMP.
	//
	// example:
	//
	// 1670656652009
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The name of the topic on the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// TOPIC-cainiao-pcs-order-process-inBoundConditionCheck
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// vsw-m5ev8asdc6h12345****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
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
	// The role name. If you want to authorize EventBridge to use this role to read logs in Simple Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console.
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
	// example:
	//
	// acs:fc:cn-hangzhou:*****:services/demo-service.LATEST/functions/demo-func
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
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

func (s *UpdateEventStreamingRequestTransforms) SetArn(v string) *UpdateEventStreamingRequestTransforms {
	s.Arn = &v
	return s
}

func (s *UpdateEventStreamingRequestTransforms) Validate() error {
	return dara.Validate(s)
}
