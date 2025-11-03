// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventStreamingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEventStreamingResponseBody
	GetCode() *string
	SetData(v *GetEventStreamingResponseBodyData) *GetEventStreamingResponseBody
	GetData() *GetEventStreamingResponseBodyData
	SetMessage(v string) *GetEventStreamingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEventStreamingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEventStreamingResponseBody
	GetSuccess() *bool
}

type GetEventStreamingResponseBody struct {
	// The response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For a list of error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetEventStreamingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// The event streaming [xxxx] not existed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7892F480-58C9-5067-AB35-8A7BEF0F726A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. The value true indicates that the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEventStreamingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEventStreamingResponseBody) GetData() *GetEventStreamingResponseBodyData {
	return s.Data
}

func (s *GetEventStreamingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEventStreamingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEventStreamingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEventStreamingResponseBody) SetCode(v string) *GetEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *GetEventStreamingResponseBody) SetData(v *GetEventStreamingResponseBodyData) *GetEventStreamingResponseBody {
	s.Data = v
	return s
}

func (s *GetEventStreamingResponseBody) SetMessage(v string) *GetEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *GetEventStreamingResponseBody) SetRequestId(v string) *GetEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventStreamingResponseBody) SetSuccess(v bool) *GetEventStreamingResponseBody {
	s.Success = &v
	return s
}

func (s *GetEventStreamingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEventStreamingResponseBodyData struct {
	// The description of the event stream that is returned.
	//
	// example:
	//
	// RocketMQ-to-RocketMQ
	Description    *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	DetailedStatus *GetEventStreamingResponseBodyDataDetailedStatus `json:"DetailedStatus,omitempty" xml:"DetailedStatus,omitempty" type:"Struct"`
	// The name of the event stream that is returned.
	//
	// example:
	//
	// rocketmq-sync
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	// The rule that is used to filter events. If you leave this parameter empty, all events are matched.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The runtime environment-related configurations.
	RunOptions *GetEventStreamingResponseBodyDataRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	// The event target.
	Sink *GetEventStreamingResponseBodyDataSink `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	// The event source.
	Source *GetEventStreamingResponseBodyDataSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// The status of the event stream that is returned.
	//
	// example:
	//
	// RUNNING
	Status     *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Transforms []*GetEventStreamingResponseBodyDataTransforms `json:"Transforms,omitempty" xml:"Transforms,omitempty" type:"Repeated"`
}

func (s GetEventStreamingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetEventStreamingResponseBodyData) GetDetailedStatus() *GetEventStreamingResponseBodyDataDetailedStatus {
	return s.DetailedStatus
}

func (s *GetEventStreamingResponseBodyData) GetEventStreamingName() *string {
	return s.EventStreamingName
}

func (s *GetEventStreamingResponseBodyData) GetFilterPattern() *string {
	return s.FilterPattern
}

func (s *GetEventStreamingResponseBodyData) GetRunOptions() *GetEventStreamingResponseBodyDataRunOptions {
	return s.RunOptions
}

func (s *GetEventStreamingResponseBodyData) GetSink() *GetEventStreamingResponseBodyDataSink {
	return s.Sink
}

func (s *GetEventStreamingResponseBodyData) GetSource() *GetEventStreamingResponseBodyDataSource {
	return s.Source
}

func (s *GetEventStreamingResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetEventStreamingResponseBodyData) GetTransforms() []*GetEventStreamingResponseBodyDataTransforms {
	return s.Transforms
}

func (s *GetEventStreamingResponseBodyData) SetDescription(v string) *GetEventStreamingResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetDetailedStatus(v *GetEventStreamingResponseBodyDataDetailedStatus) *GetEventStreamingResponseBodyData {
	s.DetailedStatus = v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetEventStreamingName(v string) *GetEventStreamingResponseBodyData {
	s.EventStreamingName = &v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetFilterPattern(v string) *GetEventStreamingResponseBodyData {
	s.FilterPattern = &v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetRunOptions(v *GetEventStreamingResponseBodyDataRunOptions) *GetEventStreamingResponseBodyData {
	s.RunOptions = v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetSink(v *GetEventStreamingResponseBodyDataSink) *GetEventStreamingResponseBodyData {
	s.Sink = v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetSource(v *GetEventStreamingResponseBodyDataSource) *GetEventStreamingResponseBodyData {
	s.Source = v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetStatus(v string) *GetEventStreamingResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetTransforms(v []*GetEventStreamingResponseBodyDataTransforms) *GetEventStreamingResponseBodyData {
	s.Transforms = v
	return s
}

func (s *GetEventStreamingResponseBodyData) Validate() error {
	if s.DetailedStatus != nil {
		if err := s.DetailedStatus.Validate(); err != nil {
			return err
		}
	}
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

type GetEventStreamingResponseBodyDataDetailedStatus struct {
	// example:
	//
	// 3
	DelayTime *int64 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// example:
	//
	// 0
	DiffOffset *int64 `json:"DiffOffset,omitempty" xml:"DiffOffset,omitempty"`
	// example:
	//
	// {
	//
	//         "test": "test",
	//
	//         "test2": 1
	//
	//       }
	Extensions map[string]interface{} `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	// example:
	//
	// 5
	TPS *float64 `json:"TPS,omitempty" xml:"TPS,omitempty"`
}

func (s GetEventStreamingResponseBodyDataDetailedStatus) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataDetailedStatus) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataDetailedStatus) GetDelayTime() *int64 {
	return s.DelayTime
}

func (s *GetEventStreamingResponseBodyDataDetailedStatus) GetDiffOffset() *int64 {
	return s.DiffOffset
}

func (s *GetEventStreamingResponseBodyDataDetailedStatus) GetExtensions() map[string]interface{} {
	return s.Extensions
}

func (s *GetEventStreamingResponseBodyDataDetailedStatus) GetTPS() *float64 {
	return s.TPS
}

func (s *GetEventStreamingResponseBodyDataDetailedStatus) SetDelayTime(v int64) *GetEventStreamingResponseBodyDataDetailedStatus {
	s.DelayTime = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataDetailedStatus) SetDiffOffset(v int64) *GetEventStreamingResponseBodyDataDetailedStatus {
	s.DiffOffset = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataDetailedStatus) SetExtensions(v map[string]interface{}) *GetEventStreamingResponseBodyDataDetailedStatus {
	s.Extensions = v
	return s
}

func (s *GetEventStreamingResponseBodyDataDetailedStatus) SetTPS(v float64) *GetEventStreamingResponseBodyDataDetailedStatus {
	s.TPS = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataDetailedStatus) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataRunOptions struct {
	// The batch window.
	BatchWindow    *GetEventStreamingResponseBodyDataRunOptionsBatchWindow    `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	BusinessOption *GetEventStreamingResponseBodyDataRunOptionsBusinessOption `json:"BusinessOption,omitempty" xml:"BusinessOption,omitempty" type:"Struct"`
	// Indicates whether dead-letter queues are enabled. By default, dead-letter queues are disabled. Messages that fail to be pushed after allowed retries as specified by the retry policy are discarded.
	DeadLetterQueue *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The fault tolerance policy. The value NONE specifies that faults are not tolerated, and the value All specifies that all faults are tolerated.
	//
	// example:
	//
	// ALL
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The concurrency level.
	//
	// example:
	//
	// 2
	MaximumTasks *int32 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The information about the retry policy that is used if the event fails to be pushed.
	RetryStrategy *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
	Throttling    *int32                                                    `json:"Throttling,omitempty" xml:"Throttling,omitempty"`
}

func (s GetEventStreamingResponseBodyDataRunOptions) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataRunOptions) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataRunOptions) GetBatchWindow() *GetEventStreamingResponseBodyDataRunOptionsBatchWindow {
	return s.BatchWindow
}

func (s *GetEventStreamingResponseBodyDataRunOptions) GetBusinessOption() *GetEventStreamingResponseBodyDataRunOptionsBusinessOption {
	return s.BusinessOption
}

func (s *GetEventStreamingResponseBodyDataRunOptions) GetDeadLetterQueue() *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue {
	return s.DeadLetterQueue
}

func (s *GetEventStreamingResponseBodyDataRunOptions) GetErrorsTolerance() *string {
	return s.ErrorsTolerance
}

func (s *GetEventStreamingResponseBodyDataRunOptions) GetMaximumTasks() *int32 {
	return s.MaximumTasks
}

func (s *GetEventStreamingResponseBodyDataRunOptions) GetRetryStrategy() *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy {
	return s.RetryStrategy
}

func (s *GetEventStreamingResponseBodyDataRunOptions) GetThrottling() *int32 {
	return s.Throttling
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetBatchWindow(v *GetEventStreamingResponseBodyDataRunOptionsBatchWindow) *GetEventStreamingResponseBodyDataRunOptions {
	s.BatchWindow = v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetBusinessOption(v *GetEventStreamingResponseBodyDataRunOptionsBusinessOption) *GetEventStreamingResponseBodyDataRunOptions {
	s.BusinessOption = v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetDeadLetterQueue(v *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) *GetEventStreamingResponseBodyDataRunOptions {
	s.DeadLetterQueue = v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetErrorsTolerance(v string) *GetEventStreamingResponseBodyDataRunOptions {
	s.ErrorsTolerance = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetMaximumTasks(v int32) *GetEventStreamingResponseBodyDataRunOptions {
	s.MaximumTasks = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetRetryStrategy(v *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) *GetEventStreamingResponseBodyDataRunOptions {
	s.RetryStrategy = v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetThrottling(v int32) *GetEventStreamingResponseBodyDataRunOptions {
	s.Throttling = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptions) Validate() error {
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

type GetEventStreamingResponseBodyDataRunOptionsBatchWindow struct {
	// The maximum number of events that are allowed in the batch window. If this threshold is reached, data in the window is pushed downstream. When multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	//
	// example:
	//
	// 100
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	// The maximum period of time during which events are allowed in the batch window. Unit: seconds. If this threshold is reached, data in the window is pushed downstream. When multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	//
	// example:
	//
	// 10
	TimeBasedWindow *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
}

func (s GetEventStreamingResponseBodyDataRunOptionsBatchWindow) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataRunOptionsBatchWindow) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBatchWindow) GetCountBasedWindow() *int32 {
	return s.CountBasedWindow
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBatchWindow) GetTimeBasedWindow() *int32 {
	return s.TimeBasedWindow
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBatchWindow) SetCountBasedWindow(v int32) *GetEventStreamingResponseBodyDataRunOptionsBatchWindow {
	s.CountBasedWindow = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBatchWindow) SetTimeBasedWindow(v int32) *GetEventStreamingResponseBodyDataRunOptionsBatchWindow {
	s.TimeBasedWindow = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBatchWindow) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataRunOptionsBusinessOption struct {
	BusinessMode         *string `json:"BusinessMode,omitempty" xml:"BusinessMode,omitempty"`
	MaxCapacityUnitCount *int64  `json:"MaxCapacityUnitCount,omitempty" xml:"MaxCapacityUnitCount,omitempty"`
	MinCapacityUnitCount *int64  `json:"MinCapacityUnitCount,omitempty" xml:"MinCapacityUnitCount,omitempty"`
}

func (s GetEventStreamingResponseBodyDataRunOptionsBusinessOption) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataRunOptionsBusinessOption) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBusinessOption) GetBusinessMode() *string {
	return s.BusinessMode
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBusinessOption) GetMaxCapacityUnitCount() *int64 {
	return s.MaxCapacityUnitCount
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBusinessOption) GetMinCapacityUnitCount() *int64 {
	return s.MinCapacityUnitCount
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBusinessOption) SetBusinessMode(v string) *GetEventStreamingResponseBodyDataRunOptionsBusinessOption {
	s.BusinessMode = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBusinessOption) SetMaxCapacityUnitCount(v int64) *GetEventStreamingResponseBodyDataRunOptionsBusinessOption {
	s.MaxCapacityUnitCount = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBusinessOption) SetMinCapacityUnitCount(v int64) *GetEventStreamingResponseBodyDataRunOptionsBusinessOption {
	s.MinCapacityUnitCount = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBusinessOption) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue.
	//
	// example:
	//
	// acs:ram::1550203943326350:role/edskmstoecs
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// example:
	//
	// PrivateNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// sg-2vcgdxz7o1n9zapp****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// vsw-m5ev8asdc6h12345****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// example:
	//
	// vpc-2zehizpoendb3****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) GetArn() *string {
	return s.Arn
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) GetNetwork() *string {
	return s.Network
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) GetVpcId() *string {
	return s.VpcId
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) SetArn(v string) *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue {
	s.Arn = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) SetNetwork(v string) *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue {
	s.Network = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) SetSecurityGroupId(v string) *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue {
	s.SecurityGroupId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) SetVSwitchIds(v string) *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue {
	s.VSwitchIds = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) SetVpcId(v string) *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue {
	s.VpcId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataRunOptionsRetryStrategy struct {
	// The maximum period of time during which retries are performed.
	//
	// example:
	//
	// 512
	MaximumEventAgeInSeconds *float32 `json:"MaximumEventAgeInSeconds,omitempty" xml:"MaximumEventAgeInSeconds,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 2
	MaximumRetryAttempts *float32 `json:"MaximumRetryAttempts,omitempty" xml:"MaximumRetryAttempts,omitempty"`
	// The retry policy. Valid values: BACKOFFRETRY and EXPONENTIALDECAY_RETRY.
	//
	// example:
	//
	// BACKOFFRETRY
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
}

func (s GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) GetMaximumEventAgeInSeconds() *float32 {
	return s.MaximumEventAgeInSeconds
}

func (s *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) GetMaximumRetryAttempts() *float32 {
	return s.MaximumRetryAttempts
}

func (s *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) GetPushRetryStrategy() *string {
	return s.PushRetryStrategy
}

func (s *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) SetMaximumEventAgeInSeconds(v float32) *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy {
	s.MaximumEventAgeInSeconds = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) SetMaximumRetryAttempts(v float32) *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy {
	s.MaximumRetryAttempts = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) SetPushRetryStrategy(v string) *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy {
	s.PushRetryStrategy = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSink struct {
	SinkApacheKafkaParameters *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters `json:"SinkApacheKafkaParameters,omitempty" xml:"SinkApacheKafkaParameters,omitempty" type:"Struct"`
	// Sink Apache RocketMQ Checkpoint Parameters
	SinkApacheRocketMQCheckpointParameters *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters `json:"SinkApacheRocketMQCheckpointParameters,omitempty" xml:"SinkApacheRocketMQCheckpointParameters,omitempty" type:"Struct"`
	SinkApiDestinationParameters           *SinkApiDestinationParameters                                                `json:"SinkApiDestinationParameters,omitempty" xml:"SinkApiDestinationParameters,omitempty"`
	// Sink BaiLian Parameters
	SinkBaiLianParameters                  *SinkBaiLianParameters                                                       `json:"SinkBaiLianParameters,omitempty" xml:"SinkBaiLianParameters,omitempty"`
	SinkCustomizedKafkaConnectorParameters *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters `json:"SinkCustomizedKafkaConnectorParameters,omitempty" xml:"SinkCustomizedKafkaConnectorParameters,omitempty" type:"Struct"`
	SinkCustomizedKafkaParameters          *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaParameters          `json:"SinkCustomizedKafkaParameters,omitempty" xml:"SinkCustomizedKafkaParameters,omitempty" type:"Struct"`
	SinkDashVectorParameters               *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters               `json:"SinkDashVectorParameters,omitempty" xml:"SinkDashVectorParameters,omitempty" type:"Struct"`
	SinkDataHubParameters                  *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters                  `json:"SinkDataHubParameters,omitempty" xml:"SinkDataHubParameters,omitempty" type:"Struct"`
	SinkDataWorksTriggerParameters         *SinkDataWorksTriggerParameters                                              `json:"SinkDataWorksTriggerParameters,omitempty" xml:"SinkDataWorksTriggerParameters,omitempty"`
	SinkDorisParameters                    *GetEventStreamingResponseBodyDataSinkSinkDorisParameters                    `json:"SinkDorisParameters,omitempty" xml:"SinkDorisParameters,omitempty" type:"Struct"`
	// The parameters that are returned if the event target is Function Compute.
	SinkFcParameters *GetEventStreamingResponseBodyDataSinkSinkFcParameters `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	// The Sink Fnf parameters.
	SinkFnfParameters   *GetEventStreamingResponseBodyDataSinkSinkFnfParameters `json:"SinkFnfParameters,omitempty" xml:"SinkFnfParameters,omitempty" type:"Struct"`
	SinkHttpsParameters *SinkHttpsParameters                                    `json:"SinkHttpsParameters,omitempty" xml:"SinkHttpsParameters,omitempty"`
	// The parameters that are returned if the event target is Message Queue for Apache Kafka.
	SinkKafkaParameters *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are returned if the event target is Message Service (MNS).
	SinkMNSParameters *GetEventStreamingResponseBodyDataSinkSinkMNSParameters `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
	// Sink Open Source RabbitMQ Parameters
	SinkOpenSourceRabbitMQParameters *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters `json:"SinkOpenSourceRabbitMQParameters,omitempty" xml:"SinkOpenSourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if the event target is Message Queue for RabbitMQ.
	SinkRabbitMQParameters *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters `json:"SinkRabbitMQParameters,omitempty" xml:"SinkRabbitMQParameters,omitempty" type:"Struct"`
	// Sink RocketMQ Checkpoint Parameters
	SinkRocketMQCheckpointParameters *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters `json:"SinkRocketMQCheckpointParameters,omitempty" xml:"SinkRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The parameters that are returned if ApsaraMQ for RocketMQ is specified as the event target.
	SinkRocketMQParameters *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters `json:"SinkRocketMQParameters,omitempty" xml:"SinkRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Simple Log Service is specified as the event target.
	SinkSLSParameters *GetEventStreamingResponseBodyDataSinkSinkSLSParameters `json:"SinkSLSParameters,omitempty" xml:"SinkSLSParameters,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSink) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSink) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkApacheKafkaParameters() *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	return s.SinkApacheKafkaParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkApacheRocketMQCheckpointParameters() *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters {
	return s.SinkApacheRocketMQCheckpointParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkApiDestinationParameters() *SinkApiDestinationParameters {
	return s.SinkApiDestinationParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkBaiLianParameters() *SinkBaiLianParameters {
	return s.SinkBaiLianParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkCustomizedKafkaConnectorParameters() *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters {
	return s.SinkCustomizedKafkaConnectorParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkCustomizedKafkaParameters() *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaParameters {
	return s.SinkCustomizedKafkaParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkDashVectorParameters() *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters {
	return s.SinkDashVectorParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkDataHubParameters() *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters {
	return s.SinkDataHubParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkDataWorksTriggerParameters() *SinkDataWorksTriggerParameters {
	return s.SinkDataWorksTriggerParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkDorisParameters() *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	return s.SinkDorisParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkFcParameters() *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	return s.SinkFcParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkFnfParameters() *GetEventStreamingResponseBodyDataSinkSinkFnfParameters {
	return s.SinkFnfParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkHttpsParameters() *SinkHttpsParameters {
	return s.SinkHttpsParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkKafkaParameters() *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	return s.SinkKafkaParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkMNSParameters() *GetEventStreamingResponseBodyDataSinkSinkMNSParameters {
	return s.SinkMNSParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkOpenSourceRabbitMQParameters() *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	return s.SinkOpenSourceRabbitMQParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkRabbitMQParameters() *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	return s.SinkRabbitMQParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkRocketMQCheckpointParameters() *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters {
	return s.SinkRocketMQCheckpointParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkRocketMQParameters() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	return s.SinkRocketMQParameters
}

func (s *GetEventStreamingResponseBodyDataSink) GetSinkSLSParameters() *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	return s.SinkSLSParameters
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkApacheKafkaParameters(v *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkApacheKafkaParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkApacheRocketMQCheckpointParameters(v *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkApacheRocketMQCheckpointParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkApiDestinationParameters(v *SinkApiDestinationParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkApiDestinationParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkBaiLianParameters(v *SinkBaiLianParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkBaiLianParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkCustomizedKafkaConnectorParameters(v *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkCustomizedKafkaConnectorParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkCustomizedKafkaParameters(v *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkCustomizedKafkaParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkDashVectorParameters(v *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkDashVectorParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkDataHubParameters(v *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkDataHubParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkDataWorksTriggerParameters(v *SinkDataWorksTriggerParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkDataWorksTriggerParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkDorisParameters(v *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkDorisParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkFcParameters(v *GetEventStreamingResponseBodyDataSinkSinkFcParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkFcParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkFnfParameters(v *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkFnfParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkHttpsParameters(v *SinkHttpsParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkHttpsParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkKafkaParameters(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkKafkaParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkMNSParameters(v *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkMNSParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkOpenSourceRabbitMQParameters(v *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkOpenSourceRabbitMQParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkRabbitMQParameters(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkRabbitMQParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkRocketMQCheckpointParameters(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkRocketMQCheckpointParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkRocketMQParameters(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkRocketMQParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkSLSParameters(v *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkSLSParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) Validate() error {
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
	if s.SinkOpenSourceRabbitMQParameters != nil {
		if err := s.SinkOpenSourceRabbitMQParameters.Validate(); err != nil {
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

type GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters struct {
	Acks             *string                                                                        `json:"Acks,omitempty" xml:"Acks,omitempty"`
	Bootstraps       *string                                                                        `json:"Bootstraps,omitempty" xml:"Bootstraps,omitempty"`
	Headers          *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders         `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	Key              *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey             `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	NetworkType      *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType     `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	SaslMechanism    *string                                                                        `json:"SaslMechanism,omitempty" xml:"SaslMechanism,omitempty"`
	SaslPassword     *string                                                                        `json:"SaslPassword,omitempty" xml:"SaslPassword,omitempty"`
	SaslUser         *string                                                                        `json:"SaslUser,omitempty" xml:"SaslUser,omitempty"`
	SecurityGroupId  *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	SecurityProtocol *string                                                                        `json:"SecurityProtocol,omitempty" xml:"SecurityProtocol,omitempty"`
	Topic            *string                                                                        `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds       *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds      `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	Value            *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue           `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
	VpcId            *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId           `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetAcks() *string {
	return s.Acks
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetBootstraps() *string {
	return s.Bootstraps
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetHeaders() *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders {
	return s.Headers
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetKey() *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey {
	return s.Key
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetNetworkType() *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType {
	return s.NetworkType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetSaslMechanism() *string {
	return s.SaslMechanism
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetSaslPassword() *string {
	return s.SaslPassword
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetSaslUser() *string {
	return s.SaslUser
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetSecurityGroupId() *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetSecurityProtocol() *string {
	return s.SecurityProtocol
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetVSwitchIds() *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetValue() *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) GetVpcId() *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId {
	return s.VpcId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetAcks(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.Acks = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetBootstraps(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.Bootstraps = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetHeaders(v *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.Headers = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetKey(v *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.Key = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetNetworkType(v *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.NetworkType = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetSaslMechanism(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.SaslMechanism = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetSaslPassword(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.SaslPassword = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetSaslUser(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.SaslUser = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetSecurityGroupId(v *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.SecurityGroupId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetSecurityProtocol(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.SecurityProtocol = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetTopic(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.Topic = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetVSwitchIds(v *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.VSwitchIds = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetValue(v *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.Value = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) SetVpcId(v *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters {
	s.VpcId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersHeaders) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersKey) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersNetworkType) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersValue) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheKafkaParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters struct {
	ConsumeTimestamp *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp `json:"ConsumeTimestamp,omitempty" xml:"ConsumeTimestamp,omitempty" type:"Struct"`
	Group            *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup            `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
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
	// sg-2ze5bmpw6adn0q******
	SecurityGroupId *string                                                                           `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Topic           *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// example:
	//
	// vsw-uf62oqt1twuevrt******
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-2zeccak5pb0j3ay******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) GetConsumeTimestamp() *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	return s.ConsumeTimestamp
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) GetGroup() *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup {
	return s.Group
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) GetTopic() *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) SetConsumeTimestamp(v *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters {
	s.ConsumeTimestamp = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) SetGroup(v *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters {
	s.Group = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) SetInstanceEndpoint(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) SetInstancePassword(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters {
	s.InstancePassword = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) SetInstanceUsername(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters {
	s.InstanceUsername = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) SetNetworkType(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters {
	s.NetworkType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) SetSecurityGroupId(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) SetTopic(v *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters {
	s.Topic = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) SetVSwitchId(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters {
	s.VSwitchId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) SetVpcId(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters {
	s.VpcId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// 1570761026400
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Group ID
	//
	// example:
	//
	// GID_EVENTBRIDGE_1736234******
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersGroup) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// Mytopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkApacheRocketMQCheckpointParametersTopic) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters struct {
	// example:
	//
	// "https://examplebucket.oss-cn-hangzhou.aliyuncs.com/testDoc/Old_Homebrew/2024-06-26%2022%3A34%3A08/opt/homebrew/homebrew/Library/Homebrew/test/support/fixtures/cask/AppWithBinary.zip?OSSAccessKeyId=ri&Expires=1725539627&Signature=rb8q3OpV******"
	ConnectorPackageUrl *string                                                                                         `json:"ConnectorPackageUrl,omitempty" xml:"ConnectorPackageUrl,omitempty"`
	ConnectorParameters *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters `json:"ConnectorParameters,omitempty" xml:"ConnectorParameters,omitempty" type:"Struct"`
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

func (s GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters) GetConnectorPackageUrl() *string {
	return s.ConnectorPackageUrl
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters) GetConnectorParameters() *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	return s.ConnectorParameters
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters) GetWorkerParameters() map[string]interface{} {
	return s.WorkerParameters
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters) SetConnectorPackageUrl(v string) *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters {
	s.ConnectorPackageUrl = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters) SetConnectorParameters(v *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters {
	s.ConnectorParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters) SetWorkerParameters(v map[string]interface{}) *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters {
	s.WorkerParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParameters) Validate() error {
	if s.ConnectorParameters != nil {
		if err := s.ConnectorParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters struct {
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

func (s GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GetName() *string {
	return s.Name
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) SetConfig(v map[string]interface{}) *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	s.Config = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) SetName(v string) *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	s.Name = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaParameters struct {
	// example:
	//
	// 90be1f96-4229-4535-bb76-34b4f6fb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkCustomizedKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters struct {
	// example:
	//
	// Q34nExQH7sQ****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// collection1
	Collection                 *string                                                                                    `json:"Collection,omitempty" xml:"Collection,omitempty"`
	DashVectorSchemaParameters []*GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters `json:"DashVectorSchemaParameters,omitempty" xml:"DashVectorSchemaParameters,omitempty" type:"Repeated"`
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
	Operation    *string                                                                    `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Partition    *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition    `json:"Partition,omitempty" xml:"Partition,omitempty" type:"Struct"`
	PrimaryKeyId *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId `json:"PrimaryKeyId,omitempty" xml:"PrimaryKeyId,omitempty" type:"Struct"`
	Vector       *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector       `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) GetApiKey() *string {
	return s.ApiKey
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) GetCollection() *string {
	return s.Collection
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) GetDashVectorSchemaParameters() []*GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters {
	return s.DashVectorSchemaParameters
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) GetNetwork() *string {
	return s.Network
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) GetOperation() *string {
	return s.Operation
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) GetPartition() *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition {
	return s.Partition
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) GetPrimaryKeyId() *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId {
	return s.PrimaryKeyId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) GetVector() *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector {
	return s.Vector
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) SetApiKey(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters {
	s.ApiKey = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) SetCollection(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters {
	s.Collection = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) SetDashVectorSchemaParameters(v []*GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters {
	s.DashVectorSchemaParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) SetNetwork(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters {
	s.Network = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) SetOperation(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters {
	s.Operation = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) SetPartition(v *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters {
	s.Partition = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) SetPrimaryKeyId(v *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters {
	s.PrimaryKeyId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) SetVector(v *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters {
	s.Vector = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters struct {
	Name  *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName  `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	Type  *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType  `json:"Type,omitempty" xml:"Type,omitempty" type:"Struct"`
	Value *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters) GetName() *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	return s.Name
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters) GetType() *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	return s.Type
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters) GetValue() *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters) SetName(v *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Name = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters) SetType(v *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Type = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters) SetValue(v *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Value = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersType) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersDashVectorSchemaParametersValue) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// default
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPartition) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId struct {
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

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersPrimaryKeyId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector struct {
	// example:
	//
	// JSONPATH
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// $.data.messageBody
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDashVectorParametersVector) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDataHubParameters struct {
	Body        *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody        `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Project     *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject     `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	RoleName    *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName    `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	Topic       *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic       `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	TopicSchema *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema `json:"TopicSchema,omitempty" xml:"TopicSchema,omitempty" type:"Struct"`
	TopicType   *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType   `json:"TopicType,omitempty" xml:"TopicType,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) GetBody() *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody {
	return s.Body
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) GetProject() *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject {
	return s.Project
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) GetRoleName() *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName {
	return s.RoleName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) GetTopic() *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) GetTopicSchema() *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema {
	return s.TopicSchema
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) GetTopicType() *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType {
	return s.TopicType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody) *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) SetProject(v *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject) *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters {
	s.Project = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) SetRoleName(v *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName) *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters {
	s.RoleName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) SetTopic(v *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic) *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters {
	s.Topic = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) SetTopicSchema(v *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema) *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters {
	s.TopicSchema = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) SetTopicType(v *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType) *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters {
	s.TopicType = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody struct {
	// example:
	//
	// ORIGINAL
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersBody) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// demo-project
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersProject) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// test-role
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// demo-topic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopic) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema struct {
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// {"k1":"${k1}","k2":"${k2}"}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// {"k1":"value1","k2":"value2"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicSchema) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// TUPLE
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDataHubParametersTopicType) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParameters struct {
	BeHttpEndpoint  *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint  `json:"BeHttpEndpoint,omitempty" xml:"BeHttpEndpoint,omitempty" type:"Struct"`
	Body            *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody            `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Database        *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase        `json:"Database,omitempty" xml:"Database,omitempty" type:"Struct"`
	FeHttpEndpoint  *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint  `json:"FeHttpEndpoint,omitempty" xml:"FeHttpEndpoint,omitempty" type:"Struct"`
	NetworkType     *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType     `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	Password        *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword        `json:"Password,omitempty" xml:"Password,omitempty" type:"Struct"`
	QueryEndpoint   *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint   `json:"QueryEndpoint,omitempty" xml:"QueryEndpoint,omitempty" type:"Struct"`
	SecurityGroupId *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	Table           *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable           `json:"Table,omitempty" xml:"Table,omitempty" type:"Struct"`
	UserName        *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName        `json:"UserName,omitempty" xml:"UserName,omitempty" type:"Struct"`
	VSwitchIds      *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds      `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	VpcId           *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId           `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetBeHttpEndpoint() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint {
	return s.BeHttpEndpoint
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetBody() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody {
	return s.Body
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetDatabase() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase {
	return s.Database
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetFeHttpEndpoint() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint {
	return s.FeHttpEndpoint
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetNetworkType() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType {
	return s.NetworkType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetPassword() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword {
	return s.Password
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetQueryEndpoint() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint {
	return s.QueryEndpoint
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetSecurityGroupId() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetTable() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable {
	return s.Table
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetUserName() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName {
	return s.UserName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetVSwitchIds() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) GetVpcId() *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId {
	return s.VpcId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetBeHttpEndpoint(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.BeHttpEndpoint = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetDatabase(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.Database = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetFeHttpEndpoint(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.FeHttpEndpoint = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetNetworkType(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.NetworkType = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetPassword(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.Password = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetQueryEndpoint(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.QueryEndpoint = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetSecurityGroupId(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.SecurityGroupId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetTable(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.Table = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetUserName(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.UserName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetVSwitchIds(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.VSwitchIds = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) SetVpcId(v *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId) *GetEventStreamingResponseBodyDataSinkSinkDorisParameters {
	s.VpcId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBeHttpEndpoint) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersBody) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersDatabase) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersFeHttpEndpoint) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersNetworkType) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersPassword) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersQueryEndpoint) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersTable) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersUserName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkDorisParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkFcParameters struct {
	// The message body that is sent to the function.
	Body *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The delivery concurrency. Minimum value: 1.
	Concurrency *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	DataFormat  *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat  `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" type:"Struct"`
	// The function name.
	FunctionName *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	// The invocation type. Valid values: Sync: synchronous Async: asynchronous
	InvocationType *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	// The alias of the service to which the function belongs.
	Qualifier *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	// The service name.
	ServiceName *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) GetBody() *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody {
	return s.Body
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) GetConcurrency() *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency {
	return s.Concurrency
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) GetDataFormat() *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat {
	return s.DataFormat
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) GetFunctionName() *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName {
	return s.FunctionName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) GetInvocationType() *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType {
	return s.InvocationType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) GetQualifier() *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier {
	return s.Qualifier
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) GetServiceName() *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName {
	return s.ServiceName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetConcurrency(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.Concurrency = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetDataFormat(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.DataFormat = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetFunctionName(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.FunctionName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetInvocationType(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.InvocationType = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetQualifier(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.Qualifier = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetServiceName(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.ServiceName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkFcParametersBody struct {
	// The method that is used to transform the event.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
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

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
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

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat struct {
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// $.data.key
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersDataFormat) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The function name.
	//
	// example:
	//
	// mFunction
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The invocation type.
	//
	// example:
	//
	// Async
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The alias of the service to which the function belongs.
	//
	// example:
	//
	// LATEST
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The name of the service.
	//
	// example:
	//
	// myService
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkFnfParameters struct {
	// The execution name.
	ExecutionName *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty" type:"Struct"`
	// The flow name.
	FlowName *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName `json:"FlowName,omitempty" xml:"FlowName,omitempty" type:"Struct"`
	// The execution input information.
	Input *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The role name.
	RoleName *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) GetExecutionName() *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName {
	return s.ExecutionName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) GetFlowName() *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName {
	return s.FlowName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) GetInput() *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput {
	return s.Input
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) GetRoleName() *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName {
	return s.RoleName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) SetExecutionName(v *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) *GetEventStreamingResponseBodyDataSinkSinkFnfParameters {
	s.ExecutionName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) SetFlowName(v *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) *GetEventStreamingResponseBodyDataSinkSinkFnfParameters {
	s.FlowName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) SetInput(v *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) *GetEventStreamingResponseBodyDataSinkSinkFnfParameters {
	s.Input = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) SetRoleName(v *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) *GetEventStreamingResponseBodyDataSinkSinkFnfParameters {
	s.RoleName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The execution name.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The flow name.
	//
	// example:
	//
	// test-streaming-fnf
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The execution input information.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role configuration.
	//
	// example:
	//
	// Al****FNF-x****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParameters struct {
	// The acknowledgment information.
	Acks    *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks    `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	Headers *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The target service type is Message Queue for Apache Kafka.
	InstanceId *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message key.
	Key *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	// The topic name.
	Topic *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The message content.
	Value *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) GetAcks() *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks {
	return s.Acks
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) GetHeaders() *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders {
	return s.Headers
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) GetInstanceId() *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) GetKey() *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey {
	return s.Key
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) GetTopic() *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) GetValue() *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) SetAcks(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	s.Acks = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) SetHeaders(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders) *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	s.Headers = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) SetInstanceId(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	s.InstanceId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) SetKey(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	s.Key = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) SetTopic(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	s.Topic = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) SetValue(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	s.Value = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The acknowledgment information.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders struct {
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersHeaders) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The instance ID.
	//
	// example:
	//
	// Defaut_1283278472_sadkj
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The message key.
	//
	// example:
	//
	// key
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The topic name.
	//
	// example:
	//
	// topic-test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue struct {
	// The method that is used to transform the event.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
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

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkMNSParameters struct {
	// The message content.
	Body *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// Indicates whether Base64 encoding is enabled.
	IsBase64Encode *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	// The target service type is MNS.
	QueueName *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) GetBody() *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody {
	return s.Body
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) GetIsBase64Encode() *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode {
	return s.IsBase64Encode
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) GetQueueName() *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName {
	return s.QueueName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) *GetEventStreamingResponseBodyDataSinkSinkMNSParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) SetIsBase64Encode(v *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) *GetEventStreamingResponseBodyDataSinkSinkMNSParameters {
	s.IsBase64Encode = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) SetQueueName(v *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) *GetEventStreamingResponseBodyDataSinkSinkMNSParameters {
	s.QueueName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody struct {
	// The method that is used to transform the event.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
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

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Specifies that Base64 encoding is enabled.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the MNS queue.
	//
	// example:
	//
	// MyQueue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters struct {
	// example:
	//
	// ACL
	AuthType *string                                                                    `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Body     *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// example:
	//
	// 192.168.1.1:9876
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// my-exchange
	Exchange  *string                                                                         `json:"Exchange,omitempty" xml:"Exchange,omitempty"`
	MessageId *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// example:
	//
	// PublicNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// ****
	Password   *string                                                                          `json:"Password,omitempty" xml:"Password,omitempty"`
	Properties *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// example:
	//
	// my-queue
	QueueName  *string                                                                          `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RoutingKey *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// example:
	//
	// sg-uf6of9452b2pba82c ****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// Exchange
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// example:
	//
	// vsw-uf6of9452b2pba82c ****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// example:
	//
	// vhost1
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	// example:
	//
	// vpc-uf6of9452b2pba82c ****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetBody() *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody {
	return s.Body
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetExchange() *string {
	return s.Exchange
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetMessageId() *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId {
	return s.MessageId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetPassword() *string {
	return s.Password
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetProperties() *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties {
	return s.Properties
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetRoutingKey() *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	return s.RoutingKey
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetTargetType() *string {
	return s.TargetType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetUsername() *string {
	return s.Username
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetAuthType(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.AuthType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetEndpoint(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.Endpoint = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetExchange(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.Exchange = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetMessageId(v *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetNetworkType(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.NetworkType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetPassword(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.Password = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetProperties(v *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetQueueName(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetRoutingKey(v *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetSecurityGroupId(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetTargetType(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.TargetType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetUsername(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.Username = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetVSwitchIds(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetVirtualHostName(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) SetVpcId(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters {
	s.VpcId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody struct {
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// {"key": "value"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// 12345
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersMessageId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties struct {
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// {"env": "prod"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey struct {
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// {"Form": "CONSTANT", "Value": "my-routing-key"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkOpenSourceRabbitMQParametersRoutingKey) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters struct {
	// The message content.
	Body *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The exchange mode. This parameter is available only if TargetType is set to Exchange.
	Exchange *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	// The target service type is Message Queue for RabbitMQ instance.
	InstanceId *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message ID.
	MessageId *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// The tags that are used to filter messages.
	Properties *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The queue mode. This parameter is available only if TargetType is set to Queue.
	QueueName *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	// The routing rule for the message. This parameter is available only if TargetType is set to Exchange.
	RoutingKey *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// The target type.
	TargetType *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
	// The name of the vhost of the Message Queue for RabbitMQ instance.
	VirtualHostName *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) GetBody() *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody {
	return s.Body
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) GetExchange() *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange {
	return s.Exchange
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) GetInstanceId() *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) GetMessageId() *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId {
	return s.MessageId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) GetProperties() *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties {
	return s.Properties
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) GetQueueName() *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName {
	return s.QueueName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) GetRoutingKey() *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey {
	return s.RoutingKey
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) GetTargetType() *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType {
	return s.TargetType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) GetVirtualHostName() *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName {
	return s.VirtualHostName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetExchange(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.Exchange = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetInstanceId(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.InstanceId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetMessageId(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetProperties(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetQueueName(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.QueueName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetRoutingKey(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetTargetType(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.TargetType = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetVirtualHostName(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.VirtualHostName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody struct {
	// The method that is used to transform the event.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
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

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The name of the exchange in the Message Queue for RabbitMQ instance.
	//
	// example:
	//
	// a_exchange
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The ID of the Message Queue for RabbitMQ instance.
	//
	// example:
	//
	// e5c9b727-e06c-4d7e-84d5-f8ce644e00bf
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId struct {
	// The method that is used to transform the event.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
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

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties struct {
	// The method that is used to transform the event.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
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

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The name of the queue in the Message Queue for RabbitMQ instance.
	//
	// example:
	//
	// MyQueue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The routing rule for the message.
	//
	// example:
	//
	// housekeeping
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The type of the resource to which the event is delivered. Valid values: Exchange: exchanges. Queue: queues.
	//
	// example:
	//
	// Exchange/Queue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The vhost name of the Message Queue for RabbitMQ instance.
	//
	// example:
	//
	// rabbit-host
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters struct {
	ConsumeTimestamp *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp `json:"ConsumeTimestamp,omitempty" xml:"ConsumeTimestamp,omitempty" type:"Struct"`
	Group            *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup            `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// example:
	//
	// MQ_INST_164901546557****_BAAN****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Cloud_4
	InstanceType *string                                                                     `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Topic        *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) GetConsumeTimestamp() *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	return s.ConsumeTimestamp
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) GetGroup() *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup {
	return s.Group
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) GetTopic() *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) SetConsumeTimestamp(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters {
	s.ConsumeTimestamp = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) SetGroup(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters {
	s.Group = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) SetInstanceType(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters {
	s.InstanceType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) SetTopic(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters {
	s.Topic = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// 1570761026400
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersConsumeTimestamp) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Group ID
	//
	// example:
	//
	// GID_EVENTBRIDGE_1736234******
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersGroup) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// Mytopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQCheckpointParametersTopic) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters struct {
	// The message content.
	Body              *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody              `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	DeliveryOrderType *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType `json:"DeliveryOrderType,omitempty" xml:"DeliveryOrderType,omitempty" type:"Struct"`
	InstanceEndpoint  *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint  `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty" type:"Struct"`
	// The target service type is Message Queue for Apache RocketMQ.
	InstanceId       *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	InstancePassword *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty" type:"Struct"`
	InstanceType     *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType     `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Struct"`
	InstanceUsername *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty" type:"Struct"`
	// The tags that are used to filter messages.
	Keys    *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys    `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	Network *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The tags that are used to filter messages.
	Properties      *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties      `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	SecurityGroupId *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	ShardingKey     *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey     `json:"ShardingKey,omitempty" xml:"ShardingKey,omitempty" type:"Struct"`
	// The tags that are used to filter messages.
	Tags *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The name of the topic in the Message Queue for Apache RocketMQ instance.
	Topic      *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic      `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	VSwitchIds *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	VpcId      *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId      `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetBody() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody {
	return s.Body
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetDeliveryOrderType() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType {
	return s.DeliveryOrderType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetInstanceEndpoint() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint {
	return s.InstanceEndpoint
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetInstanceId() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetInstancePassword() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword {
	return s.InstancePassword
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetInstanceType() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType {
	return s.InstanceType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetInstanceUsername() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername {
	return s.InstanceUsername
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetKeys() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys {
	return s.Keys
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetNetwork() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork {
	return s.Network
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetProperties() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties {
	return s.Properties
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetSecurityGroupId() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetShardingKey() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey {
	return s.ShardingKey
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetTags() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags {
	return s.Tags
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetTopic() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetVSwitchIds() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GetVpcId() *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId {
	return s.VpcId
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetDeliveryOrderType(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.DeliveryOrderType = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetInstanceEndpoint(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.InstanceEndpoint = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetInstanceId(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.InstanceId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetInstancePassword(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.InstancePassword = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetInstanceType(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.InstanceType = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetInstanceUsername(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.InstanceUsername = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetKeys(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.Keys = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetNetwork(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.Network = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetProperties(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.Properties = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetSecurityGroupId(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.SecurityGroupId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetShardingKey(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.ShardingKey = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetTags(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.Tags = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetTopic(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.Topic = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetVSwitchIds(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.VSwitchIds = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetVpcId(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.VpcId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody struct {
	// The method that is used to transform the event.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
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

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// Concurrently
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersDeliveryOrderType) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// vbr-8vbsvkkbpf3vb0zef****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceEndpoint) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The ID of the Message Queue for Apache RocketMQ instance.
	//
	// example:
	//
	// MQ_INST_164901546557****_BAAN****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// admin****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstancePassword) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// Cloud_4
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceType) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// admin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceUsername) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys struct {
	// The method that is used to transform the event.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
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

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// PublicNetwork
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersNetwork) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties struct {
	// The method that is used to transform the event.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
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

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// b4bf375515f6440f942e3a20c33d****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// order_id
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersShardingKey) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags struct {
	// The method that is used to transform the event.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
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

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic struct {
	// The method that is used to transform the event. Default value: CONSTANT.
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
	// The name of the topic in the Message Queue for Apache RocketMQ instance.
	//
	// example:
	//
	// Mytopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// vbr-8vb835n3zf9shwl****mp
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// vbr-8vb835n3zf9shwlvb****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParameters struct {
	// The message content.
	Body          *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody          `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	ContentSchema *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema `json:"ContentSchema,omitempty" xml:"ContentSchema,omitempty" type:"Struct"`
	ContentType   *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType   `json:"ContentType,omitempty" xml:"ContentType,omitempty" type:"Struct"`
	// The Simple Log Service Logstore.
	LogStore *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	// The Simple Log Service project.
	Project *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Simple Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console.
	RoleName *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The name of the topic in which logs are stored. The topic corresponds to the topic reserved field in Simple Log Service.
	Topic *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) GetBody() *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody {
	return s.Body
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) GetContentSchema() *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema {
	return s.ContentSchema
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) GetContentType() *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType {
	return s.ContentType
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) GetLogStore() *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore {
	return s.LogStore
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) GetProject() *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject {
	return s.Project
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) GetRoleName() *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName {
	return s.RoleName
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) GetTopic() *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetContentSchema(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.ContentSchema = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetContentType(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.ContentType = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetLogStore(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.LogStore = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetProject(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.Project = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetRoleName(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.RoleName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetTopic(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.Topic = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) Validate() error {
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

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody struct {
	// The method that is used to transform the event.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
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

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema struct {
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

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentSchema) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType struct {
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

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersContentType) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Log Service Logstore.
	//
	// example:
	//
	// test-logstore
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Log Service project.
	//
	// example:
	//
	// test-project
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	//
	// example:
	//
	// testRole
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic in which logs are stored. The topic corresponds to the topic reserved field in Log Service.
	//
	// example:
	//
	// testTopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) GetForm() *string {
	return s.Form
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) GetValue() *string {
	return s.Value
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic {
	s.Value = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSource struct {
	SourceApacheKafkaParameters *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters `json:"SourceApacheKafkaParameters,omitempty" xml:"SourceApacheKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Apache RocketMQ (Offset Data) is specified as the event source.
	SourceApacheRocketMQCheckpointParameters *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters `json:"SourceApacheRocketMQCheckpointParameters,omitempty" xml:"SourceApacheRocketMQCheckpointParameters,omitempty" type:"Struct"`
	SourceCustomizedKafkaConnectorParameters *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters `json:"SourceCustomizedKafkaConnectorParameters,omitempty" xml:"SourceCustomizedKafkaConnectorParameters,omitempty" type:"Struct"`
	SourceCustomizedKafkaParameters          *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaParameters          `json:"SourceCustomizedKafkaParameters,omitempty" xml:"SourceCustomizedKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are returned if the event source is Data Transmission Service (DTS).
	SourceDTSParameters      *GetEventStreamingResponseBodyDataSourceSourceDTSParameters      `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	SourceEventBusParameters *GetEventStreamingResponseBodyDataSourceSourceEventBusParameters `json:"SourceEventBusParameters,omitempty" xml:"SourceEventBusParameters,omitempty" type:"Struct"`
	// The parameters that are returned if ApsaraMQ for Kafka is specified as the event source.
	SourceKafkaParameters *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// Source MNS Parameters
	SourceMNSParameters *GetEventStreamingResponseBodyDataSourceSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters that are returned if ApsaraMQ for MQTT is specified as the event source.
	SourceMQTTParameters               *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters               `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	SourceMySQLParameters              *SourceMySQLParameters                                                     `json:"SourceMySQLParameters,omitempty" xml:"SourceMySQLParameters,omitempty"`
	SourceOSSParameters                *GetEventStreamingResponseBodyDataSourceSourceOSSParameters                `json:"SourceOSSParameters,omitempty" xml:"SourceOSSParameters,omitempty" type:"Struct"`
	SourceOpenSourceRabbitMQParameters *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters `json:"SourceOpenSourceRabbitMQParameters,omitempty" xml:"SourceOpenSourceRabbitMQParameters,omitempty" type:"Struct"`
	SourcePrometheusParameters         *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters         `json:"SourcePrometheusParameters,omitempty" xml:"SourcePrometheusParameters,omitempty" type:"Struct"`
	// Source RabbitMQ Parameters
	SourceRabbitMQParameters           *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters           `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	SourceRocketMQCheckpointParameters *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters `json:"SourceRocketMQCheckpointParameters,omitempty" xml:"SourceRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The parameters that are returned if ApsaraMQ for RocketMQ is specified as the event source.
	SourceRocketMQParameters *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if the event provider is Simple Log Service.
	SourceSLSParameters *GetEventStreamingResponseBodyDataSourceSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSource) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceApacheKafkaParameters() *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	return s.SourceApacheKafkaParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceApacheRocketMQCheckpointParameters() *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters {
	return s.SourceApacheRocketMQCheckpointParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceCustomizedKafkaConnectorParameters() *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters {
	return s.SourceCustomizedKafkaConnectorParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceCustomizedKafkaParameters() *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaParameters {
	return s.SourceCustomizedKafkaParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceDTSParameters() *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	return s.SourceDTSParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceEventBusParameters() *GetEventStreamingResponseBodyDataSourceSourceEventBusParameters {
	return s.SourceEventBusParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceKafkaParameters() *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	return s.SourceKafkaParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceMNSParameters() *GetEventStreamingResponseBodyDataSourceSourceMNSParameters {
	return s.SourceMNSParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceMQTTParameters() *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters {
	return s.SourceMQTTParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceMySQLParameters() *SourceMySQLParameters {
	return s.SourceMySQLParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceOSSParameters() *GetEventStreamingResponseBodyDataSourceSourceOSSParameters {
	return s.SourceOSSParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceOpenSourceRabbitMQParameters() *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	return s.SourceOpenSourceRabbitMQParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourcePrometheusParameters() *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters {
	return s.SourcePrometheusParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceRabbitMQParameters() *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters {
	return s.SourceRabbitMQParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceRocketMQCheckpointParameters() *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters {
	return s.SourceRocketMQCheckpointParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceRocketMQParameters() *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	return s.SourceRocketMQParameters
}

func (s *GetEventStreamingResponseBodyDataSource) GetSourceSLSParameters() *GetEventStreamingResponseBodyDataSourceSourceSLSParameters {
	return s.SourceSLSParameters
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceApacheKafkaParameters(v *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceApacheKafkaParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceApacheRocketMQCheckpointParameters(v *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceApacheRocketMQCheckpointParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceCustomizedKafkaConnectorParameters(v *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceCustomizedKafkaConnectorParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceCustomizedKafkaParameters(v *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceCustomizedKafkaParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceDTSParameters(v *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceDTSParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceEventBusParameters(v *GetEventStreamingResponseBodyDataSourceSourceEventBusParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceEventBusParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceKafkaParameters(v *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceKafkaParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceMNSParameters(v *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceMNSParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceMQTTParameters(v *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceMQTTParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceMySQLParameters(v *SourceMySQLParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceMySQLParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceOSSParameters(v *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceOSSParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceOpenSourceRabbitMQParameters(v *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceOpenSourceRabbitMQParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourcePrometheusParameters(v *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourcePrometheusParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceRabbitMQParameters(v *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceRocketMQCheckpointParameters(v *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceRocketMQCheckpointParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceRocketMQParameters(v *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceRocketMQParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceSLSParameters(v *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceSLSParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) Validate() error {
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
	if s.SourcePrometheusParameters != nil {
		if err := s.SourcePrometheusParameters.Validate(); err != nil {
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

type GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters struct {
	Bootstraps       *string `json:"Bootstraps,omitempty" xml:"Bootstraps,omitempty"`
	ConsumerGroup    *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OffsetReset      *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	SaslMechanism    *string `json:"SaslMechanism,omitempty" xml:"SaslMechanism,omitempty"`
	SaslPassword     *string `json:"SaslPassword,omitempty" xml:"SaslPassword,omitempty"`
	SaslUser         *string `json:"SaslUser,omitempty" xml:"SaslUser,omitempty"`
	SecurityGroupId  *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityProtocol *string `json:"SecurityProtocol,omitempty" xml:"SecurityProtocol,omitempty"`
	Topic            *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds       *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	ValueDataType    *string `json:"ValueDataType,omitempty" xml:"ValueDataType,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetBootstraps() *string {
	return s.Bootstraps
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetSaslMechanism() *string {
	return s.SaslMechanism
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetSaslPassword() *string {
	return s.SaslPassword
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetSaslUser() *string {
	return s.SaslUser
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetSecurityProtocol() *string {
	return s.SecurityProtocol
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetValueDataType() *string {
	return s.ValueDataType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetBootstraps(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.Bootstraps = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetConsumerGroup(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetNetworkType(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.NetworkType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetOffsetReset(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetSaslMechanism(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.SaslMechanism = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetSaslPassword(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.SaslPassword = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetSaslUser(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.SaslUser = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetSecurityGroupId(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetSecurityProtocol(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.SecurityProtocol = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetTopic(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.Topic = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetVSwitchIds(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetValueDataType(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.ValueDataType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) SetVpcId(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters struct {
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the Apache RocketMQ instance.
	//
	// example:
	//
	// MQ_INST_164901546557****_BAAN****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 123456
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

func (s GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) GetTopics() []*string {
	return s.Topics
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) SetInstanceEndpoint(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) SetInstancePassword(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters {
	s.InstancePassword = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) SetInstanceUsername(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters {
	s.InstanceUsername = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) SetNetworkType(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters {
	s.NetworkType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) SetSecurityGroupId(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) SetTopics(v []*string) *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters {
	s.Topics = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) SetVSwitchId(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters {
	s.VSwitchId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) SetVpcId(v string) *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters {
	s.VpcId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceApacheRocketMQCheckpointParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters struct {
	// example:
	//
	// "https://examplebucket.oss-cn-hangzhou.aliyuncs.com/testDoc/Old_Homebrew/2024-06-26%2022%3A34%3A08/opt/homebrew/homebrew/Library/Homebrew/test/support/fixtures/cask/AppWithBinary.zip?OSSAccessKeyId=ri&Expires=1725539627&Signature=rb8q3OpV2i3gZJ"
	ConnectorPackageUrl *string                                                                                             `json:"ConnectorPackageUrl,omitempty" xml:"ConnectorPackageUrl,omitempty"`
	ConnectorParameters *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters `json:"ConnectorParameters,omitempty" xml:"ConnectorParameters,omitempty" type:"Struct"`
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

func (s GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters) GetConnectorPackageUrl() *string {
	return s.ConnectorPackageUrl
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters) GetConnectorParameters() *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	return s.ConnectorParameters
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters) GetWorkerParameters() map[string]interface{} {
	return s.WorkerParameters
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters) SetConnectorPackageUrl(v string) *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters {
	s.ConnectorPackageUrl = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters) SetConnectorParameters(v *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters {
	s.ConnectorParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters) SetWorkerParameters(v map[string]interface{}) *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters {
	s.WorkerParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParameters) Validate() error {
	if s.ConnectorParameters != nil {
		if err := s.ConnectorParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters struct {
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

func (s GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GetName() *string {
	return s.Name
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) SetConfig(v map[string]interface{}) *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	s.Config = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) SetName(v string) *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	s.Name = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaParameters struct {
	// example:
	//
	// r-8vb64581862c****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceCustomizedKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceDTSParameters struct {
	// The URL and port number of the data subscription channel.
	BrokerUrl *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	// The consumer offset. A consumer offset is a timestamp that indicates when the SDK client consumes the first data record. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1620962769
	InitCheckPoint *string `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	// The password of the consumer group.
	//
	// example:
	//
	// admin
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the consumer group.
	//
	// example:
	//
	// HD1
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The task ID.
	//
	// example:
	//
	// f86e5814-b223-482c-b768-3b873297dade
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The topic to which you want to subscribe by using the data subscription channel.
	//
	// example:
	//
	// TP_TEST_UNDERWRITE_ISSUE
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The account of the consumer group.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceDTSParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceDTSParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) GetBrokerUrl() *string {
	return s.BrokerUrl
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) GetInitCheckPoint() *string {
	return s.InitCheckPoint
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) GetPassword() *string {
	return s.Password
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) GetSid() *string {
	return s.Sid
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) GetTaskId() *string {
	return s.TaskId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) GetTopic() *string {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) GetUsername() *string {
	return s.Username
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetBrokerUrl(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.BrokerUrl = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetInitCheckPoint(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.InitCheckPoint = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetPassword(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.Password = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetSid(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.Sid = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetTaskId(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.TaskId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetTopic(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.Topic = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetUsername(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.Username = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceEventBusParameters struct {
	// example:
	//
	// demo
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// example:
	//
	// test
	EventRuleName *string `json:"EventRuleName,omitempty" xml:"EventRuleName,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceEventBusParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceEventBusParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceEventBusParameters) GetEventBusName() *string {
	return s.EventBusName
}

func (s *GetEventStreamingResponseBodyDataSourceSourceEventBusParameters) GetEventRuleName() *string {
	return s.EventRuleName
}

func (s *GetEventStreamingResponseBodyDataSourceSourceEventBusParameters) SetEventBusName(v string) *GetEventStreamingResponseBodyDataSourceSourceEventBusParameters {
	s.EventBusName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceEventBusParameters) SetEventRuleName(v string) *GetEventStreamingResponseBodyDataSourceSourceEventBusParameters {
	s.EventRuleName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceEventBusParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceKafkaParameters struct {
	// The ID of the consumer group that subscribes to the topic.
	//
	// example:
	//
	// GID_TEST
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1fbtrnqmjvgq66ajdw
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network. Default value: Default. The value PublicNetwork specifies a virtual private cloud (VPC).
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
	// The region ID of the Message Queue for Apache Kafka instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-2vcgdxz7o1n9zappuimt
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// topic_empower_1642473600414
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-wz9qqeovkwjxlu9uc8rst
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The encoding or decoding format. Valid values: Json, Text, and Binary. The value Json indicates that bytes are decoded into UTF-8 strings and then parsed into JSON format. The value Text indicates that bytes are decoded into UTF-8 strings and then put into the payload. The value Binary indicates that bytes are encoded into Base64 strings and put into the payload.
	//
	// example:
	//
	// Text
	ValueDataType *string `json:"ValueDataType,omitempty" xml:"ValueDataType,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2zehizpoendb3nwwu9w5o
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GetNetwork() *string {
	return s.Network
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GetValueDataType() *string {
	return s.ValueDataType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetConsumerGroup(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetNetwork(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetOffsetReset(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetSecurityGroupId(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetTopic(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetVSwitchIds(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetValueDataType(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.ValueDataType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetVpcId(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceMNSParameters struct {
	// Indicates whether Base64 encoding is enabled.
	//
	// example:
	//
	// true
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the MNS queue.
	//
	// example:
	//
	// demo
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region ID of the MNS queue.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) GetIsBase64Decode() *bool {
	return s.IsBase64Decode
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) SetIsBase64Decode(v bool) *GetEventStreamingResponseBodyDataSourceSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) SetQueueName(v string) *GetEventStreamingResponseBodyDataSourceSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceMNSParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceMQTTParameters struct {
	// example:
	//
	// JSON
	BodyDataType *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-2ze06wqdwk0uq14krrzv
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Message Queue for MQTT instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic in the Message Queue for MQTT instance.
	//
	// example:
	//
	// TOPIC-cainiao-pcs-wms-instock-noPrealertPrintLabel
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) GetTopic() *string {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) SetBodyDataType(v string) *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters {
	s.BodyDataType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) SetTopic(v string) *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters {
	s.Topic = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceOSSParameters struct {
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

func (s GetEventStreamingResponseBodyDataSourceSourceOSSParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceOSSParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) GetBucketName() *string {
	return s.BucketName
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) GetDelimiter() *string {
	return s.Delimiter
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) GetLoadFormat() *string {
	return s.LoadFormat
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) GetLoadMode() *string {
	return s.LoadMode
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) GetPrefix() *string {
	return s.Prefix
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) SetBucketName(v string) *GetEventStreamingResponseBodyDataSourceSourceOSSParameters {
	s.BucketName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) SetDelimiter(v string) *GetEventStreamingResponseBodyDataSourceSourceOSSParameters {
	s.Delimiter = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) SetLoadFormat(v string) *GetEventStreamingResponseBodyDataSourceSourceOSSParameters {
	s.LoadFormat = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) SetLoadMode(v string) *GetEventStreamingResponseBodyDataSourceSourceOSSParameters {
	s.LoadMode = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) SetPrefix(v string) *GetEventStreamingResponseBodyDataSourceSourceOSSParameters {
	s.Prefix = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) SetRoleName(v string) *GetEventStreamingResponseBodyDataSourceSourceOSSParameters {
	s.RoleName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOSSParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters struct {
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
	// 192.168.1.100:5672
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// demo
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// example:
	//
	// sg-2ze65razphjfz3******
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// example:
	//
	// vsw-uf65zhil5oukof5******
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// example:
	//
	// eb-connect
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	// example:
	//
	// vpc-bp1vllc1lnw1v657******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GetPassword() *string {
	return s.Password
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GetUsername() *string {
	return s.Username
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) SetAuthType(v string) *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	s.AuthType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) SetBodyDataType(v string) *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	s.BodyDataType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) SetEndpoint(v string) *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	s.Endpoint = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) SetNetworkType(v string) *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	s.NetworkType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) SetPassword(v string) *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	s.Password = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) SetQueueName(v string) *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) SetSecurityGroupId(v string) *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) SetUsername(v string) *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	s.Username = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) SetVSwitchIds(v string) *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) SetVirtualHostName(v string) *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) SetVpcId(v string) *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters {
	s.VpcId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceOpenSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters struct {
	// example:
	//
	// c83555068b6******ad213f565f209
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// json
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// {"env":"test"}
	ExternalLabels *string `json:"ExternalLabels,omitempty" xml:"ExternalLabels,omitempty"`
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

func (s GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) GetDataType() *string {
	return s.DataType
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) GetExternalLabels() *string {
	return s.ExternalLabels
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) GetLabels() *string {
	return s.Labels
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) SetClusterId(v string) *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters {
	s.ClusterId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) SetDataType(v string) *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters {
	s.DataType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) SetExternalLabels(v string) *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters {
	s.ExternalLabels = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) SetLabels(v string) *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters {
	s.Labels = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) SetRoleName(v string) *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters {
	s.RoleName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters struct {
	// The ID of the Message Queue for RabbitMQ instance.
	//
	// example:
	//
	// i-f8z9a9mcgwri1c1idd0z
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue in the Message Queue for RabbitMQ instance.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region ID of the Message Queue for RabbitMQ instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vhost name of the Message Queue for RabbitMQ instance.
	//
	// example:
	//
	// eb-connect
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) SetQueueName(v string) *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) SetVirtualHostName(v string) *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters struct {
	// example:
	//
	// rmp-cn-jte3w******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Cloud_4
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topics   []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) GetTopics() []*string {
	return s.Topics
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) SetInstanceType(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters {
	s.InstanceType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) SetTopics(v []*string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters {
	s.Topics = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQCheckpointParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters struct {
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
	// The ID of the consumer group in the Message Queue for Apache RocketMQ instance.
	//
	// example:
	//
	// GID_group1
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// example:
	//
	// reg****-vpc.cn-zhangjiakou.aliyuncs.com
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance.
	//
	// example:
	//
	// i-f8zbher64dlm58plyfte
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// PublicNetwork
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// example:
	//
	// xxxa
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// example:
	//
	// sg-m5edtu24f12345****
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// example:
	//
	// 2
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// xxxa
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// example:
	//
	// vsw-m5ev8asdc6h12****
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// example:
	//
	// vpc-m5e3sv4b12345****
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// example:
	//
	// PublicNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The consumer offset of messages. Valid values: CONSUME_FROM_LAST_OFFSET: Start consumption from the latest offset. CONSUME_FROM_FIRST_OFFSET: Start consumption from the earliest offset. CONSUME_FROM_TIMESTAMP: Start consumption from the offset at the specified point in time.
	//
	// example:
	//
	// CONSUMEFROMLASTOFFSET
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The region ID of the Message Queue for Apache RocketMQ instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-m5edtu24f12345****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The tags that are used to filter messages.
	//
	// example:
	//
	// v1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp of the offset from which consumption starts. This parameter is valid only if you set the Offset parameter to CONSUME_FROM_TIMESTAMP.
	//
	// example:
	//
	// 1636597951964
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The topic to which the message belongs.
	//
	// example:
	//
	// topic_add_anima
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

func (s GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetFilterSql() *string {
	return s.FilterSql
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetFilterType() *string {
	return s.FilterType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetGroupID() *string {
	return s.GroupID
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetInstanceNetwork() *string {
	return s.InstanceNetwork
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetInstanceSecurityGroupId() *string {
	return s.InstanceSecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetInstanceVSwitchIds() *string {
	return s.InstanceVSwitchIds
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetInstanceVpcId() *string {
	return s.InstanceVpcId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetNetwork() *string {
	return s.Network
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetOffset() *string {
	return s.Offset
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetTag() *string {
	return s.Tag
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetTopic() *string {
	return s.Topic
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetAuthType(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetBodyDataType(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.BodyDataType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetFilterSql(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.FilterSql = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetFilterType(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.FilterType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetGroupID(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceEndpoint(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceNetwork(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstancePassword(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceType(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceUsername(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceVpcId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetNetwork(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.Network = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetOffset(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetSecurityGroupId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetTag(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetTimestamp(v int64) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetTopic(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.Topic = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetVSwitchIds(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetVpcId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.VpcId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataSourceSourceSLSParameters struct {
	// The starting consumer offset. The value begin indicates the earliest offset, and the value end indicates the latest offset. You can also specify a time in seconds to start consumption.
	//
	// example:
	//
	// begin
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The consumer group.
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The Log Service Logstore.
	//
	// example:
	//
	// waf-logstore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Log Service project.
	//
	// example:
	//
	// dmmzk
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	//
	// example:
	//
	// testRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceSLSParameters) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) GetConsumePosition() *string {
	return s.ConsumePosition
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) GetLogStore() *string {
	return s.LogStore
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) GetProject() *string {
	return s.Project
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) SetConsumePosition(v string) *GetEventStreamingResponseBodyDataSourceSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) SetConsumerGroup(v string) *GetEventStreamingResponseBodyDataSourceSourceSLSParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) SetLogStore(v string) *GetEventStreamingResponseBodyDataSourceSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) SetProject(v string) *GetEventStreamingResponseBodyDataSourceSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) SetRoleName(v string) *GetEventStreamingResponseBodyDataSourceSourceSLSParameters {
	s.RoleName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) Validate() error {
	return dara.Validate(s)
}

type GetEventStreamingResponseBodyDataTransforms struct {
	// The Alibaba Cloud Resource Name (ARN) of the cloud service, such as the ARN of a Function Compute function.
	//
	// example:
	//
	// acs:fc:cn-hangzhou:*****:services/demo-service.LATEST/functions/demo-func
	Arn                             *string                          `json:"Arn,omitempty" xml:"Arn,omitempty"`
	BaiLianAgentTransformParameters *BaiLianAgentTransformParameters `json:"BaiLianAgentTransformParameters,omitempty" xml:"BaiLianAgentTransformParameters,omitempty"`
	DashScopeTransformParameters    *DashScopeTransformParameters    `json:"DashScopeTransformParameters,omitempty" xml:"DashScopeTransformParameters,omitempty"`
}

func (s GetEventStreamingResponseBodyDataTransforms) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataTransforms) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataTransforms) GetArn() *string {
	return s.Arn
}

func (s *GetEventStreamingResponseBodyDataTransforms) GetBaiLianAgentTransformParameters() *BaiLianAgentTransformParameters {
	return s.BaiLianAgentTransformParameters
}

func (s *GetEventStreamingResponseBodyDataTransforms) GetDashScopeTransformParameters() *DashScopeTransformParameters {
	return s.DashScopeTransformParameters
}

func (s *GetEventStreamingResponseBodyDataTransforms) SetArn(v string) *GetEventStreamingResponseBodyDataTransforms {
	s.Arn = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataTransforms) SetBaiLianAgentTransformParameters(v *BaiLianAgentTransformParameters) *GetEventStreamingResponseBodyDataTransforms {
	s.BaiLianAgentTransformParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataTransforms) SetDashScopeTransformParameters(v *DashScopeTransformParameters) *GetEventStreamingResponseBodyDataTransforms {
	s.DashScopeTransformParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataTransforms) Validate() error {
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
