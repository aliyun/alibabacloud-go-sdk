// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventStreamingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEventStreamingsResponseBody
	GetCode() *string
	SetData(v *ListEventStreamingsResponseBodyData) *ListEventStreamingsResponseBody
	GetData() *ListEventStreamingsResponseBodyData
	SetMessage(v string) *ListEventStreamingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEventStreamingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEventStreamingsResponseBody
	GetSuccess() *bool
}

type ListEventStreamingsResponseBody struct {
	// The response code. Valid values:
	//
	// Success: The request is successful.
	//
	// Other codes: The request failed. For more information about error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListEventStreamingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// The event streaming [xxxx] not existed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 283FF852-C4B8-58C9-9777-F88A5A16A79F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. The value true indicates that the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEventStreamingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEventStreamingsResponseBody) GetData() *ListEventStreamingsResponseBodyData {
	return s.Data
}

func (s *ListEventStreamingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEventStreamingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventStreamingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEventStreamingsResponseBody) SetCode(v string) *ListEventStreamingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventStreamingsResponseBody) SetData(v *ListEventStreamingsResponseBodyData) *ListEventStreamingsResponseBody {
	s.Data = v
	return s
}

func (s *ListEventStreamingsResponseBody) SetMessage(v string) *ListEventStreamingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEventStreamingsResponseBody) SetRequestId(v string) *ListEventStreamingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventStreamingsResponseBody) SetSuccess(v bool) *ListEventStreamingsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEventStreamingsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyData struct {
	// The event streams.
	EventStreamings []*ListEventStreamingsResponseBodyDataEventStreamings `json:"EventStreamings,omitempty" xml:"EventStreamings,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists. You must specify the pagination token in the next request.
	//
	// example:
	//
	// 177
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEventStreamingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyData) GetEventStreamings() []*ListEventStreamingsResponseBodyDataEventStreamings {
	return s.EventStreamings
}

func (s *ListEventStreamingsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventStreamingsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListEventStreamingsResponseBodyData) SetEventStreamings(v []*ListEventStreamingsResponseBodyDataEventStreamings) *ListEventStreamingsResponseBodyData {
	s.EventStreamings = v
	return s
}

func (s *ListEventStreamingsResponseBodyData) SetNextToken(v string) *ListEventStreamingsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListEventStreamingsResponseBodyData) SetTotal(v int32) *ListEventStreamingsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListEventStreamingsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamings struct {
	// The description of the event stream.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event stream.
	//
	// example:
	//
	// name
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	// The rule that is used to filter events. If you leave this parameter empty, all events are matched.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The parameters that are returned for the runtime environment.
	RunOptions *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	// The event target.
	Sink *ListEventStreamingsResponseBodyDataEventStreamingsSink `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	// The event provider, which is also known as the event source.
	Source *ListEventStreamingsResponseBodyDataEventStreamingsSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// The status of the event stream that is returned.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The transformation-related configurations.
	Transforms []*ListEventStreamingsResponseBodyDataEventStreamingsTransforms `json:"Transforms,omitempty" xml:"Transforms,omitempty" type:"Repeated"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamings) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamings) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) GetDescription() *string {
	return s.Description
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) GetEventStreamingName() *string {
	return s.EventStreamingName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) GetFilterPattern() *string {
	return s.FilterPattern
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) GetRunOptions() *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	return s.RunOptions
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) GetSink() *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	return s.Sink
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) GetSource() *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	return s.Source
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) GetStatus() *string {
	return s.Status
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) GetTransforms() []*ListEventStreamingsResponseBodyDataEventStreamingsTransforms {
	return s.Transforms
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetDescription(v string) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.Description = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetEventStreamingName(v string) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.EventStreamingName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetFilterPattern(v string) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.FilterPattern = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetRunOptions(v *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.RunOptions = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetSink(v *ListEventStreamingsResponseBodyDataEventStreamingsSink) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.Sink = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetSource(v *ListEventStreamingsResponseBodyDataEventStreamingsSource) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.Source = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetStatus(v string) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.Status = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetTransforms(v []*ListEventStreamingsResponseBodyDataEventStreamingsTransforms) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.Transforms = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsRunOptions struct {
	// The batch window.
	BatchWindow    *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow    `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	BusinessOption *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption `json:"BusinessOption,omitempty" xml:"BusinessOption,omitempty" type:"Struct"`
	// Indicates whether dead-letter queues are enabled. By default, dead-letter queues are disabled. Events that fail to be pushed are discarded after the maximum number of retries that is specified by the retry policy is reached.
	DeadLetterQueue *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
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
	MaximumTasks *int32 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The retry policy that is used if events fail to be pushed.
	RetryStrategy *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
	Throttling    *int32                                                                     `json:"Throttling,omitempty" xml:"Throttling,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) GetBatchWindow() *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow {
	return s.BatchWindow
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) GetBusinessOption() *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption {
	return s.BusinessOption
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) GetDeadLetterQueue() *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue {
	return s.DeadLetterQueue
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) GetErrorsTolerance() *string {
	return s.ErrorsTolerance
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) GetMaximumTasks() *int32 {
	return s.MaximumTasks
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) GetRetryStrategy() *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy {
	return s.RetryStrategy
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) GetThrottling() *int32 {
	return s.Throttling
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetBatchWindow(v *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.BatchWindow = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetBusinessOption(v *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.BusinessOption = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetDeadLetterQueue(v *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.DeadLetterQueue = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetErrorsTolerance(v string) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.ErrorsTolerance = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetMaximumTasks(v int32) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.MaximumTasks = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetRetryStrategy(v *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.RetryStrategy = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetThrottling(v int32) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.Throttling = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow struct {
	// The maximum number of events that are allowed in the batch window. When this threshold is reached, data in the window is pushed to the downstream service. When multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	//
	// example:
	//
	// 100
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	// The maximum period of time during which events are allowed in the batch window. Unit: seconds. When this threshold is reached, data in the window is pushed to the downstream service. When multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	//
	// example:
	//
	// 10
	TimeBasedWindow *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) GetCountBasedWindow() *int32 {
	return s.CountBasedWindow
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) GetTimeBasedWindow() *int32 {
	return s.TimeBasedWindow
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) SetCountBasedWindow(v int32) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow {
	s.CountBasedWindow = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) SetTimeBasedWindow(v int32) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow {
	s.TimeBasedWindow = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption struct {
	BusinessMode         *string `json:"BusinessMode,omitempty" xml:"BusinessMode,omitempty"`
	MaxCapacityUnitCount *int64  `json:"MaxCapacityUnitCount,omitempty" xml:"MaxCapacityUnitCount,omitempty"`
	MinCapacityUnitCount *int64  `json:"MinCapacityUnitCount,omitempty" xml:"MinCapacityUnitCount,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption) GetBusinessMode() *string {
	return s.BusinessMode
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption) GetMaxCapacityUnitCount() *int64 {
	return s.MaxCapacityUnitCount
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption) GetMinCapacityUnitCount() *int64 {
	return s.MinCapacityUnitCount
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption) SetBusinessMode(v string) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption {
	s.BusinessMode = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption) SetMaxCapacityUnitCount(v int64) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption {
	s.MaxCapacityUnitCount = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption) SetMinCapacityUnitCount(v int64) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption {
	s.MinCapacityUnitCount = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBusinessOption) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue struct {
	// The ARN of the dead-letter queue.
	//
	// example:
	//
	// acs:ram::1597871211794192:role/aliyunsaedefaultrole
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) GetArn() *string {
	return s.Arn
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) GetNetwork() *string {
	return s.Network
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) SetArn(v string) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue {
	s.Arn = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) SetNetwork(v string) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue {
	s.Network = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) SetSecurityGroupId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue {
	s.SecurityGroupId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) SetVSwitchIds(v string) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue {
	s.VSwitchIds = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) SetVpcId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue {
	s.VpcId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy struct {
	// The retry policy. Valid values: BACKOFF_RETRY and EXPONENTIAL_DECAY_RETRY.
	//
	// example:
	//
	// EXPONENTIALDECAY_RETRY
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) GetPushRetryStrategy() *string {
	return s.PushRetryStrategy
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) SetPushRetryStrategy(v string) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy {
	s.PushRetryStrategy = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSink struct {
	SinkApacheKafkaParameters              *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters              `json:"SinkApacheKafkaParameters,omitempty" xml:"SinkApacheKafkaParameters,omitempty" type:"Struct"`
	SinkApacheRocketMQCheckpointParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters `json:"SinkApacheRocketMQCheckpointParameters,omitempty" xml:"SinkApacheRocketMQCheckpointParameters,omitempty" type:"Struct"`
	SinkBaiLianParameters                  *SinkBaiLianParameters                                                                        `json:"SinkBaiLianParameters,omitempty" xml:"SinkBaiLianParameters,omitempty"`
	SinkCustomizedKafkaConnectorParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters `json:"SinkCustomizedKafkaConnectorParameters,omitempty" xml:"SinkCustomizedKafkaConnectorParameters,omitempty" type:"Struct"`
	SinkCustomizedKafkaParameters          *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaParameters          `json:"SinkCustomizedKafkaParameters,omitempty" xml:"SinkCustomizedKafkaParameters,omitempty" type:"Struct"`
	SinkDashVectorParameters               *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters               `json:"SinkDashVectorParameters,omitempty" xml:"SinkDashVectorParameters,omitempty" type:"Struct"`
	SinkDataHubParameters                  *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters                  `json:"SinkDataHubParameters,omitempty" xml:"SinkDataHubParameters,omitempty" type:"Struct"`
	SinkDataWorksTriggerParameters         *SinkDataWorksTriggerParameters                                                               `json:"SinkDataWorksTriggerParameters,omitempty" xml:"SinkDataWorksTriggerParameters,omitempty"`
	SinkDorisParameters                    *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters                    `json:"SinkDorisParameters,omitempty" xml:"SinkDorisParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Function Compute is specified as the event target.
	SinkFcParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	// The parameters that are returned if CloudFlow is specified as the event target.
	SinkFnfParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters `json:"SinkFnfParameters,omitempty" xml:"SinkFnfParameters,omitempty" type:"Struct"`
	// The parameters that are returned if ApsaraMQ for Kafka is specified as the event target.
	SinkKafkaParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are returned if MNS is specified as the event target.
	SinkMNSParameters                *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters                `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
	SinkOpenSourceRabbitMQParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters `json:"SinkOpenSourceRabbitMQParameters,omitempty" xml:"SinkOpenSourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if ApsaraMQ for RabbitMQ is specified as the event target.
	SinkRabbitMQParameters           *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters           `json:"SinkRabbitMQParameters,omitempty" xml:"SinkRabbitMQParameters,omitempty" type:"Struct"`
	SinkRocketMQCheckpointParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters `json:"SinkRocketMQCheckpointParameters,omitempty" xml:"SinkRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The parameters that are returned if ApsaraMQ for RocketMQ is specified as the event target.
	SinkRocketMQParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters `json:"SinkRocketMQParameters,omitempty" xml:"SinkRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Simple Log Service is specified as the event target.
	SinkSLSParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters `json:"SinkSLSParameters,omitempty" xml:"SinkSLSParameters,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSink) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSink) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkApacheKafkaParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	return s.SinkApacheKafkaParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkApacheRocketMQCheckpointParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters {
	return s.SinkApacheRocketMQCheckpointParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkBaiLianParameters() *SinkBaiLianParameters {
	return s.SinkBaiLianParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkCustomizedKafkaConnectorParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters {
	return s.SinkCustomizedKafkaConnectorParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkCustomizedKafkaParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaParameters {
	return s.SinkCustomizedKafkaParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkDashVectorParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters {
	return s.SinkDashVectorParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkDataHubParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters {
	return s.SinkDataHubParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkDataWorksTriggerParameters() *SinkDataWorksTriggerParameters {
	return s.SinkDataWorksTriggerParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkDorisParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	return s.SinkDorisParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkFcParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	return s.SinkFcParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkFnfParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters {
	return s.SinkFnfParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkKafkaParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters {
	return s.SinkKafkaParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkMNSParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters {
	return s.SinkMNSParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkOpenSourceRabbitMQParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	return s.SinkOpenSourceRabbitMQParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkRabbitMQParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	return s.SinkRabbitMQParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkRocketMQCheckpointParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters {
	return s.SinkRocketMQCheckpointParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkRocketMQParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	return s.SinkRocketMQParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) GetSinkSLSParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	return s.SinkSLSParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkApacheKafkaParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkApacheKafkaParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkApacheRocketMQCheckpointParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkApacheRocketMQCheckpointParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkBaiLianParameters(v *SinkBaiLianParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkBaiLianParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkCustomizedKafkaConnectorParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkCustomizedKafkaConnectorParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkCustomizedKafkaParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkCustomizedKafkaParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkDashVectorParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkDashVectorParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkDataHubParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkDataHubParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkDataWorksTriggerParameters(v *SinkDataWorksTriggerParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkDataWorksTriggerParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkDorisParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkDorisParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkFcParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkFcParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkFnfParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkFnfParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkKafkaParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkKafkaParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkMNSParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkMNSParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkOpenSourceRabbitMQParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkOpenSourceRabbitMQParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkRabbitMQParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkRabbitMQParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkRocketMQCheckpointParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkRocketMQCheckpointParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkRocketMQParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkRocketMQParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkSLSParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkSLSParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters struct {
	Acks             *string                                                                                         `json:"Acks,omitempty" xml:"Acks,omitempty"`
	Bootstraps       *string                                                                                         `json:"Bootstraps,omitempty" xml:"Bootstraps,omitempty"`
	Key              *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey             `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	NetworkType      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType     `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	SaslMechanism    *string                                                                                         `json:"SaslMechanism,omitempty" xml:"SaslMechanism,omitempty"`
	SaslPassword     *string                                                                                         `json:"SaslPassword,omitempty" xml:"SaslPassword,omitempty"`
	SaslUser         *string                                                                                         `json:"SaslUser,omitempty" xml:"SaslUser,omitempty"`
	SecurityGroupId  *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	SecurityProtocol *string                                                                                         `json:"SecurityProtocol,omitempty" xml:"SecurityProtocol,omitempty"`
	Topic            *string                                                                                         `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds       *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds      `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	Value            *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue           `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
	VpcId            *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId           `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetAcks() *string {
	return s.Acks
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetBootstraps() *string {
	return s.Bootstraps
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetKey() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey {
	return s.Key
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetNetworkType() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType {
	return s.NetworkType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetSaslMechanism() *string {
	return s.SaslMechanism
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetSaslPassword() *string {
	return s.SaslPassword
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetSaslUser() *string {
	return s.SaslUser
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetSecurityGroupId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetSecurityProtocol() *string {
	return s.SecurityProtocol
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetVSwitchIds() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetValue() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) GetVpcId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId {
	return s.VpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetAcks(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.Acks = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetBootstraps(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.Bootstraps = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetKey(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.Key = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetNetworkType(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.NetworkType = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetSaslMechanism(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.SaslMechanism = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetSaslPassword(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.SaslPassword = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetSaslUser(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.SaslUser = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetSecurityGroupId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.SecurityGroupId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetSecurityProtocol(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.SecurityProtocol = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetTopic(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.Topic = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetVSwitchIds(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.VSwitchIds = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetValue(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.Value = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) SetVpcId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters {
	s.VpcId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersKey) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersNetworkType) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersValue) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheKafkaParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters struct {
	ConsumeTimestamp *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp `json:"ConsumeTimestamp,omitempty" xml:"ConsumeTimestamp,omitempty" type:"Struct"`
	Group            *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup            `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
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
	// sg-2vcgdxz7o1n9zapp****
	SecurityGroupId *string                                                                                            `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Topic           *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// example:
	//
	// vsw-wz9qqeovkwjxlu9uc****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-2zehizpoendb3****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) GetConsumeTimestamp() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	return s.ConsumeTimestamp
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) GetGroup() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup {
	return s.Group
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) GetTopic() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) SetConsumeTimestamp(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters {
	s.ConsumeTimestamp = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) SetGroup(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters {
	s.Group = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) SetInstanceEndpoint(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) SetInstancePassword(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters {
	s.InstancePassword = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) SetInstanceUsername(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters {
	s.InstanceUsername = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) SetNetworkType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters {
	s.NetworkType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) SetSecurityGroupId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) SetTopic(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters {
	s.Topic = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) SetVSwitchId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters {
	s.VSwitchId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) SetVpcId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters {
	s.VpcId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// 1636597951964
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersConsumeTimestamp) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// GID_******
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersGroup) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// myTopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkApacheRocketMQCheckpointParametersTopic) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters struct {
	// example:
	//
	// "https://examplebucket.oss-cn-hangzhou.aliyuncs.com/testDoc/Old_Homebrew/2024-06-26%2022%3A34%3A08/opt/homebrew/homebrew/Library/Homebrew/test/support/fixtures/cask/AppWithBinary.zip?OSSAccessKeyId=ri&Expires=1725539627&Signature=rb8q3OpV2i3gZJ"
	ConnectorPackageUrl *string                                                                                                          `json:"ConnectorPackageUrl,omitempty" xml:"ConnectorPackageUrl,omitempty"`
	ConnectorParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters `json:"ConnectorParameters,omitempty" xml:"ConnectorParameters,omitempty" type:"Struct"`
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters) GetConnectorPackageUrl() *string {
	return s.ConnectorPackageUrl
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters) GetConnectorParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	return s.ConnectorParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters) GetWorkerParameters() map[string]interface{} {
	return s.WorkerParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters) SetConnectorPackageUrl(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters {
	s.ConnectorPackageUrl = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters) SetConnectorParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters {
	s.ConnectorParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters) SetWorkerParameters(v map[string]interface{}) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters {
	s.WorkerParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) GetName() *string {
	return s.Name
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) SetConfig(v map[string]interface{}) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	s.Config = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) SetName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters {
	s.Name = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaConnectorParametersConnectorParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaParameters struct {
	// example:
	//
	// 90be1f96-4229-4535-bb76-34b4f6fb2b71
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkCustomizedKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters struct {
	// example:
	//
	// Q34nExQH7sQ****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// collection1
	Collection                 *string                                                                                                   `json:"Collection,omitempty" xml:"Collection,omitempty"`
	DashVectorSchemaParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters `json:"DashVectorSchemaParameters,omitempty" xml:"DashVectorSchemaParameters,omitempty" type:"Struct"`
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
	Operation    *string                                                                                     `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Partition    *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition    `json:"Partition,omitempty" xml:"Partition,omitempty" type:"Struct"`
	PrimaryKeyId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId `json:"PrimaryKeyId,omitempty" xml:"PrimaryKeyId,omitempty" type:"Struct"`
	Vector       *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector       `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) GetCollection() *string {
	return s.Collection
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) GetDashVectorSchemaParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters {
	return s.DashVectorSchemaParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) GetNetwork() *string {
	return s.Network
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) GetOperation() *string {
	return s.Operation
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) GetPartition() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition {
	return s.Partition
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) GetPrimaryKeyId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId {
	return s.PrimaryKeyId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) GetVector() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector {
	return s.Vector
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) SetApiKey(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters {
	s.ApiKey = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) SetCollection(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters {
	s.Collection = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) SetDashVectorSchemaParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters {
	s.DashVectorSchemaParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) SetNetwork(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters {
	s.Network = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) SetOperation(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters {
	s.Operation = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) SetPartition(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters {
	s.Partition = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) SetPrimaryKeyId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters {
	s.PrimaryKeyId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) SetVector(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters {
	s.Vector = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters struct {
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersDashVectorSchemaParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPartition) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersPrimaryKeyId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDashVectorParametersVector) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters struct {
	Body        *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody        `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Project     *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject     `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	RoleName    *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName    `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	Topic       *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic       `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	TopicSchema *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema `json:"TopicSchema,omitempty" xml:"TopicSchema,omitempty" type:"Struct"`
	TopicType   *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType   `json:"TopicType,omitempty" xml:"TopicType,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) GetBody() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody {
	return s.Body
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) GetProject() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject {
	return s.Project
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) GetRoleName() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName {
	return s.RoleName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) GetTopic() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) GetTopicSchema() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema {
	return s.TopicSchema
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) GetTopicType() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType {
	return s.TopicType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) SetProject(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters {
	s.Project = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) SetRoleName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters {
	s.RoleName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) SetTopic(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters {
	s.Topic = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) SetTopicSchema(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters {
	s.TopicSchema = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) SetTopicType(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters {
	s.TopicType = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody struct {
	// example:
	//
	// ORIGINAL
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersBody) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersProject) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// testRole
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopic) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// {"k1":"value1","k2":"value2"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicSchema) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDataHubParametersTopicType) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters struct {
	BeHttpEndpoint  *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint  `json:"BeHttpEndpoint,omitempty" xml:"BeHttpEndpoint,omitempty" type:"Struct"`
	Body            *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody            `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Database        *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase        `json:"Database,omitempty" xml:"Database,omitempty" type:"Struct"`
	FeHttpEndpoint  *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint  `json:"FeHttpEndpoint,omitempty" xml:"FeHttpEndpoint,omitempty" type:"Struct"`
	NetworkType     *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType     `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	Password        *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword        `json:"Password,omitempty" xml:"Password,omitempty" type:"Struct"`
	QueryEndpoint   *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint   `json:"QueryEndpoint,omitempty" xml:"QueryEndpoint,omitempty" type:"Struct"`
	SecurityGroupId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	Table           *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable           `json:"Table,omitempty" xml:"Table,omitempty" type:"Struct"`
	UserName        *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName        `json:"UserName,omitempty" xml:"UserName,omitempty" type:"Struct"`
	VSwitchIds      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds      `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	VpcId           *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId           `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetBeHttpEndpoint() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint {
	return s.BeHttpEndpoint
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetBody() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody {
	return s.Body
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetDatabase() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase {
	return s.Database
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetFeHttpEndpoint() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint {
	return s.FeHttpEndpoint
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetNetworkType() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType {
	return s.NetworkType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetPassword() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword {
	return s.Password
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetQueryEndpoint() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint {
	return s.QueryEndpoint
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetSecurityGroupId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetTable() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable {
	return s.Table
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetUserName() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName {
	return s.UserName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetVSwitchIds() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) GetVpcId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId {
	return s.VpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetBeHttpEndpoint(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.BeHttpEndpoint = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetDatabase(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.Database = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetFeHttpEndpoint(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.FeHttpEndpoint = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetNetworkType(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.NetworkType = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetPassword(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.Password = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetQueryEndpoint(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.QueryEndpoint = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetSecurityGroupId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.SecurityGroupId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetTable(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.Table = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetUserName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.UserName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetVSwitchIds(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.VSwitchIds = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) SetVpcId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters {
	s.VpcId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBeHttpEndpoint) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersBody) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersDatabase) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersFeHttpEndpoint) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersNetworkType) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersPassword) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersQueryEndpoint) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersTable) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersUserName) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkDorisParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters struct {
	// The message body that is delivered to Function Compute.
	Body *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The delivery concurrency. Minimum value: 1.
	Concurrency *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	DataFormat  *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat  `json:"DataFormat,omitempty" xml:"DataFormat,omitempty" type:"Struct"`
	// The function name.
	FunctionName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	// The invocation mode. Valid values:
	//
	// 	- Sync
	//
	// 	- Async
	InvocationType *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	// The alias of the service to which the function belongs.
	Qualifier *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	// The service name.
	ServiceName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) GetBody() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody {
	return s.Body
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) GetConcurrency() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency {
	return s.Concurrency
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) GetDataFormat() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat {
	return s.DataFormat
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) GetFunctionName() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName {
	return s.FunctionName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) GetInvocationType() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType {
	return s.InvocationType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) GetQualifier() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier {
	return s.Qualifier
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) GetServiceName() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName {
	return s.ServiceName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetConcurrency(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.Concurrency = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetDataFormat(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.DataFormat = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetFunctionName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.FunctionName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetInvocationType(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.InvocationType = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetQualifier(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.Qualifier = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetServiceName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.ServiceName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody struct {
	// The method that is used to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersDataFormat) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The invocation mode. Valid values:
	//
	// 	- Sync
	//
	// 	- Async
	//
	// example:
	//
	// Async
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters struct {
	// The execution name.
	ExecutionName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty" type:"Struct"`
	// The flow name.
	FlowName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName `json:"FlowName,omitempty" xml:"FlowName,omitempty" type:"Struct"`
	// The input information of the execution.
	Input *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The role name.
	RoleName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) GetExecutionName() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName {
	return s.ExecutionName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) GetFlowName() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName {
	return s.FlowName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) GetInput() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput {
	return s.Input
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) GetRoleName() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName {
	return s.RoleName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) SetExecutionName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters {
	s.ExecutionName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) SetFlowName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters {
	s.FlowName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) SetInput(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters {
	s.Input = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) SetRoleName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters {
	s.RoleName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	//
	// example:
	//
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The input information of the execution.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters struct {
	// The acknowledgment (ACK) mode.
	//
	// 	- If you set this parameter to 0, no response is returned from the broker. In this mode, the performance is high, but the risk of data loss is also high.
	//
	// 	- If you set this parameter to 1, a response is returned when data is written to the leader. In this mode, the performance and the risk of data loss are moderate. Data loss may occur if a failure occurs on the leader.
	//
	// 	- If you set this parameter to all, a response is returned when data is written to the leader and synchronized to the followers. In this mode, the performance is low, but the risk of data loss is also low. Data loss occurs if the leader and the followers fail at the same time.
	Acks *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	// The ID of the ApsaraMQ for Kafka instance.
	InstanceId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message key.
	Key *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	// The name of the topic on the ApsaraMQ for Kafka instance.
	Topic *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The message body.
	Value *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) GetAcks() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks {
	return s.Acks
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) GetInstanceId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) GetKey() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey {
	return s.Key
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) GetTopic() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) GetValue() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) SetAcks(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters {
	s.Acks = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) SetInstanceId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters {
	s.InstanceId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) SetKey(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters {
	s.Key = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) SetTopic(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters {
	s.Topic = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) SetValue(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters {
	s.Value = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic struct {
	// The method that is used to transform events. Default value: CONSTANT.
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
	// topic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue struct {
	// The method that is used to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters struct {
	// The message content.
	Body *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// Indicates whether Base64 encoding is enabled.
	IsBase64Encode *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	// The name of the MNS queue.
	QueueName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) GetBody() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody {
	return s.Body
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) GetIsBase64Encode() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode {
	return s.IsBase64Encode
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) GetQueueName() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName {
	return s.QueueName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) SetIsBase64Encode(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters {
	s.IsBase64Encode = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) SetQueueName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters {
	s.QueueName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody struct {
	// The method that is used to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Indicates that Base64 encoding is enabled.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the MNS queue.
	//
	// example:
	//
	// MyQueue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters struct {
	// example:
	//
	// ACL
	AuthType *string                                                                                     `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Body     *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// example:
	//
	// 192.168.1.1:9876
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// my_exchange
	Exchange  *string                                                                                          `json:"Exchange,omitempty" xml:"Exchange,omitempty"`
	MessageId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// ****
	Password   *string                                                                                           `json:"Password,omitempty" xml:"Password,omitempty"`
	Properties *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// example:
	//
	// my_queue
	QueueName  *string                                                                                           `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RoutingKey *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// example:
	//
	// sg-2vcgdxz7o1n9zapp****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// exchange
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// example:
	//
	// vsw-wz9qqeovkwjxlu9uc****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// example:
	//
	// Vhost1
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-2zehizpoendb3****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetBody() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody {
	return s.Body
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetExchange() *string {
	return s.Exchange
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetMessageId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId {
	return s.MessageId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetPassword() *string {
	return s.Password
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetProperties() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties {
	return s.Properties
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetRoutingKey() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	return s.RoutingKey
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetTargetType() *string {
	return s.TargetType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetUsername() *string {
	return s.Username
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetAuthType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.AuthType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetEndpoint(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.Endpoint = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetExchange(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.Exchange = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetMessageId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetNetworkType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.NetworkType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetPassword(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.Password = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetProperties(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetQueueName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetRoutingKey(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetSecurityGroupId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetTargetType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.TargetType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetUsername(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.Username = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetVSwitchIds(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetVirtualHostName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) SetVpcId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters {
	s.VpcId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody struct {
	// example:
	//
	// TEMPLATE
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId struct {
	// example:
	//
	// TEMPLATE
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersMessageId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties struct {
	// example:
	//
	// TEMPLATE
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey struct {
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// housekeeping
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkOpenSourceRabbitMQParametersRoutingKey) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters struct {
	// The message content.
	Body *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The exchange mode. This parameter is required only if TargetType is set to Exchange.
	Exchange *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
	InstanceId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message ID.
	MessageId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// The properties that are used to filter messages.
	Properties *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The queue mode. This parameter is required only if TargetType is set to Queue.
	QueueName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	// The rule that is used to route messages. This parameter is required only if TargetType is set to Exchange.
	RoutingKey *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// The type of the resource to which events are delivered.
	TargetType *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
	// The name of the vhost to which the ApsaraMQ for RabbitMQ instance belongs.
	VirtualHostName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) GetBody() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody {
	return s.Body
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) GetExchange() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange {
	return s.Exchange
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) GetInstanceId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) GetMessageId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId {
	return s.MessageId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) GetProperties() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties {
	return s.Properties
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) GetQueueName() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName {
	return s.QueueName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) GetRoutingKey() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey {
	return s.RoutingKey
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) GetTargetType() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType {
	return s.TargetType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) GetVirtualHostName() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName {
	return s.VirtualHostName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetExchange(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.Exchange = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetInstanceId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.InstanceId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetMessageId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetProperties(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetQueueName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.QueueName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetRoutingKey(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetTargetType(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.TargetType = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetVirtualHostName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.VirtualHostName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody struct {
	// The method that is used to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId struct {
	// The method that is used to transform events. Default value: CONSTANT.
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
	// si-296cd57939a1421b94ec
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId struct {
	// The method that is used to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties struct {
	// The method that is used to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The rule that is used to route messages.
	//
	// example:
	//
	// housekeeping
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the resource to which events are delivered. Valid values: Exchange and Queue.
	//
	// example:
	//
	// Queue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the vhost to which the ApsaraMQ for RabbitMQ instance belongs.
	//
	// example:
	//
	// rabbit-host
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters struct {
	ConsumeTimestamp *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp `json:"ConsumeTimestamp,omitempty" xml:"ConsumeTimestamp,omitempty" type:"Struct"`
	Group            *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup            `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// example:
	//
	// MQ_INST_1825725063814405_BZ******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Cloud_5
	InstanceType *string                                                                                      `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Topic        *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) GetConsumeTimestamp() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	return s.ConsumeTimestamp
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) GetGroup() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup {
	return s.Group
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) GetTopic() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) SetConsumeTimestamp(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters {
	s.ConsumeTimestamp = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) SetGroup(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters {
	s.Group = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) SetInstanceType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters {
	s.InstanceType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) SetTopic(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters {
	s.Topic = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// 1636597951964
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersConsumeTimestamp) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// GID_******
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersGroup) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// myTopic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQCheckpointParametersTopic) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters struct {
	// The message content.
	Body              *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody              `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	DeliveryOrderType *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType `json:"DeliveryOrderType,omitempty" xml:"DeliveryOrderType,omitempty" type:"Struct"`
	InstanceEndpoint  *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint  `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty" type:"Struct"`
	// The ID of the ApsaraMQ for RocketMQ instance.
	InstanceId       *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	InstancePassword *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty" type:"Struct"`
	InstanceType     *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType     `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Struct"`
	InstanceUsername *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty" type:"Struct"`
	// The keys that are used to filter messages.
	Keys    *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys    `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	Network *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The properties that are used to filter messages.
	Properties      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties      `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	SecurityGroupId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	ShardingKey     *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey     `json:"ShardingKey,omitempty" xml:"ShardingKey,omitempty" type:"Struct"`
	// The tags that are used to filter messages.
	Tags *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The topic on the ApsaraMQ for RocketMQ instance.
	Topic      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic      `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	VSwitchIds *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	VpcId      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId      `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetBody() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody {
	return s.Body
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetDeliveryOrderType() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType {
	return s.DeliveryOrderType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetInstanceEndpoint() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint {
	return s.InstanceEndpoint
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetInstanceId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetInstancePassword() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword {
	return s.InstancePassword
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetInstanceType() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType {
	return s.InstanceType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetInstanceUsername() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername {
	return s.InstanceUsername
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetKeys() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys {
	return s.Keys
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetNetwork() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork {
	return s.Network
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetProperties() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties {
	return s.Properties
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetSecurityGroupId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId {
	return s.SecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetShardingKey() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey {
	return s.ShardingKey
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetTags() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags {
	return s.Tags
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetTopic() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetVSwitchIds() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds {
	return s.VSwitchIds
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GetVpcId() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId {
	return s.VpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetDeliveryOrderType(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.DeliveryOrderType = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetInstanceEndpoint(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.InstanceEndpoint = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetInstanceId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.InstanceId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetInstancePassword(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.InstancePassword = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetInstanceType(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.InstanceType = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetInstanceUsername(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.InstanceUsername = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetKeys(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.Keys = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetNetwork(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.Network = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetProperties(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.Properties = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetSecurityGroupId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.SecurityGroupId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetShardingKey(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.ShardingKey = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetTags(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.Tags = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetTopic(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.Topic = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetVSwitchIds(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.VSwitchIds = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetVpcId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.VpcId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody struct {
	// The method that is used to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// Orderly
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersDeliveryOrderType) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceEndpoint) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword struct {
	// example:
	//
	// CONSTANT
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// admin******
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstancePassword) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceType) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceUsername) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys struct {
	// The method that is used to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersNetwork) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties struct {
	// The method that is used to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersSecurityGroupId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersShardingKey) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags struct {
	// The method that is used to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic on the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// topic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersVpcId) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters struct {
	// The message body that is sent to Simple Log Service.
	Body          *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody          `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	ContentSchema *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema `json:"ContentSchema,omitempty" xml:"ContentSchema,omitempty" type:"Struct"`
	ContentType   *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType   `json:"ContentType,omitempty" xml:"ContentType,omitempty" type:"Struct"`
	// The Simple Log Service Logstore.
	LogStore *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	// The Simple Log Service project.
	Project *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	RoleName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The name of the topic in which logs are stored. The topic corresponds to the topic reserved field in Simple Log Service.
	Topic *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) GetBody() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody {
	return s.Body
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) GetContentSchema() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema {
	return s.ContentSchema
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) GetContentType() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType {
	return s.ContentType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) GetLogStore() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore {
	return s.LogStore
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) GetProject() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject {
	return s.Project
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) GetRoleName() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName {
	return s.RoleName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) GetTopic() *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetContentSchema(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.ContentSchema = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetContentType(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.ContentType = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetLogStore(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.LogStore = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetProject(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.Project = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetRoleName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.RoleName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetTopic(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.Topic = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody struct {
	// The method that is used to transform events.
	//
	// example:
	//
	// TEMPLATE
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentSchema) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersContentType) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject struct {
	// The method that is used to transform events. Default value: CONSTANT.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	//
	// example:
	//
	// test-role
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic struct {
	// The method that is used to transform events. Default value: CONSTANT.
	//
	// example:
	//
	// CONSTANT
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic in which logs are stored. The topic corresponds to the topic reserved field in Simple Log Service.
	//
	// example:
	//
	// topic
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) GetForm() *string {
	return s.Form
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) GetTemplate() *string {
	return s.Template
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSource struct {
	SourceApacheKafkaParameters              *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters              `json:"SourceApacheKafkaParameters,omitempty" xml:"SourceApacheKafkaParameters,omitempty" type:"Struct"`
	SourceApacheRocketMQCheckpointParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters `json:"SourceApacheRocketMQCheckpointParameters,omitempty" xml:"SourceApacheRocketMQCheckpointParameters,omitempty" type:"Struct"`
	SourceCustomizedKafkaConnectorParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters `json:"SourceCustomizedKafkaConnectorParameters,omitempty" xml:"SourceCustomizedKafkaConnectorParameters,omitempty" type:"Struct"`
	SourceCustomizedKafkaParameters          *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaParameters          `json:"SourceCustomizedKafkaParameters,omitempty" xml:"SourceCustomizedKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Data Transmission Service (DTS) is specified as the event source.
	SourceDTSParameters      *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters      `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	SourceEventBusParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters `json:"SourceEventBusParameters,omitempty" xml:"SourceEventBusParameters,omitempty" type:"Struct"`
	// The parameters that are returned if ApsaraMQ for Kafka is specified as the event source.
	SourceKafkaParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue (MNS) is specified as the event source.
	SourceMNSParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters that are returned if ApsaraMQ for MQTT is specified as the event source.
	SourceMQTTParameters               *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters               `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	SourceMySQLParameters              *SourceMySQLParameters                                                                      `json:"SourceMySQLParameters,omitempty" xml:"SourceMySQLParameters,omitempty"`
	SourceOSSParameters                *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters                `json:"SourceOSSParameters,omitempty" xml:"SourceOSSParameters,omitempty" type:"Struct"`
	SourceOpenSourceRabbitMQParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters `json:"SourceOpenSourceRabbitMQParameters,omitempty" xml:"SourceOpenSourceRabbitMQParameters,omitempty" type:"Struct"`
	SourcePrometheusParameters         *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters         `json:"SourcePrometheusParameters,omitempty" xml:"SourcePrometheusParameters,omitempty" type:"Struct"`
	// The parameters that are returned if ApsaraMQ for RabbitMQ is specified as the event source.
	SourceRabbitMQParameters           *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters           `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	SourceRocketMQCheckpointParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters `json:"SourceRocketMQCheckpointParameters,omitempty" xml:"SourceRocketMQCheckpointParameters,omitempty" type:"Struct"`
	// The parameters that are returned if ApsaraMQ for RocketMQ is specified as the event source.
	SourceRocketMQParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Simple Log Service is specified as the event source.
	SourceSLSParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSource) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSource) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceApacheKafkaParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	return s.SourceApacheKafkaParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceApacheRocketMQCheckpointParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters {
	return s.SourceApacheRocketMQCheckpointParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceCustomizedKafkaConnectorParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters {
	return s.SourceCustomizedKafkaConnectorParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceCustomizedKafkaParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaParameters {
	return s.SourceCustomizedKafkaParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceDTSParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	return s.SourceDTSParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceEventBusParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters {
	return s.SourceEventBusParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceKafkaParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	return s.SourceKafkaParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceMNSParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters {
	return s.SourceMNSParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceMQTTParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters {
	return s.SourceMQTTParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceMySQLParameters() *SourceMySQLParameters {
	return s.SourceMySQLParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceOSSParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters {
	return s.SourceOSSParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceOpenSourceRabbitMQParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	return s.SourceOpenSourceRabbitMQParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourcePrometheusParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters {
	return s.SourcePrometheusParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceRabbitMQParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters {
	return s.SourceRabbitMQParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceRocketMQCheckpointParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters {
	return s.SourceRocketMQCheckpointParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceRocketMQParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	return s.SourceRocketMQParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) GetSourceSLSParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters {
	return s.SourceSLSParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceApacheKafkaParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceApacheKafkaParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceApacheRocketMQCheckpointParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceApacheRocketMQCheckpointParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceCustomizedKafkaConnectorParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceCustomizedKafkaConnectorParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceCustomizedKafkaParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceCustomizedKafkaParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceDTSParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceDTSParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceEventBusParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceEventBusParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceKafkaParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceKafkaParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceMNSParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceMNSParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceMQTTParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceMQTTParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceMySQLParameters(v *SourceMySQLParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceMySQLParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceOSSParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceOSSParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceOpenSourceRabbitMQParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceOpenSourceRabbitMQParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourcePrometheusParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourcePrometheusParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceRabbitMQParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceRocketMQCheckpointParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceRocketMQCheckpointParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceRocketMQParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceRocketMQParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceSLSParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceSLSParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetBootstraps() *string {
	return s.Bootstraps
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetSaslMechanism() *string {
	return s.SaslMechanism
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetSaslPassword() *string {
	return s.SaslPassword
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetSaslUser() *string {
	return s.SaslUser
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetSecurityProtocol() *string {
	return s.SecurityProtocol
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetValueDataType() *string {
	return s.ValueDataType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetBootstraps(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.Bootstraps = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetConsumerGroup(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetNetworkType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.NetworkType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetOffsetReset(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetSaslMechanism(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.SaslMechanism = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetSaslPassword(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.SaslPassword = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetSaslUser(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.SaslUser = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetSecurityGroupId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetSecurityProtocol(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.SecurityProtocol = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetTopic(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.Topic = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetVSwitchIds(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetValueDataType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.ValueDataType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) SetVpcId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) GetTopics() []*string {
	return s.Topics
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) SetInstanceEndpoint(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) SetInstancePassword(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters {
	s.InstancePassword = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) SetInstanceUsername(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters {
	s.InstanceUsername = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) SetNetworkType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters {
	s.NetworkType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) SetSecurityGroupId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) SetTopics(v []*string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters {
	s.Topics = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) SetVSwitchId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters {
	s.VSwitchId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) SetVpcId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters {
	s.VpcId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceApacheRocketMQCheckpointParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters struct {
	// example:
	//
	// "https://examplebucket.oss-cn-hangzhou.aliyuncs.com/testDoc/Old_Homebrew/2024-06-26%2022%3A34%3A08/opt/homebrew/homebrew/Library/Homebrew/test/support/fixtures/cask/AppWithBinary.zip?OSSAccessKeyId=ri&Expires=1725539627&Signature=rb8q3OpV2i3gZJ"
	ConnectorPackageUrl *string                                                                                                              `json:"ConnectorPackageUrl,omitempty" xml:"ConnectorPackageUrl,omitempty"`
	ConnectorParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters `json:"ConnectorParameters,omitempty" xml:"ConnectorParameters,omitempty" type:"Struct"`
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters) GetConnectorPackageUrl() *string {
	return s.ConnectorPackageUrl
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters) GetConnectorParameters() *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	return s.ConnectorParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters) GetWorkerParameters() map[string]interface{} {
	return s.WorkerParameters
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters) SetConnectorPackageUrl(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters {
	s.ConnectorPackageUrl = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters) SetConnectorParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters {
	s.ConnectorParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters) SetWorkerParameters(v map[string]interface{}) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters {
	s.WorkerParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) GetName() *string {
	return s.Name
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) SetConfig(v map[string]interface{}) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	s.Config = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) SetName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters {
	s.Name = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaConnectorParametersConnectorParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaParameters struct {
	// example:
	//
	// gtm-cn-k2c2yfg****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceCustomizedKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters struct {
	// The URL and port number of the change tracking instance.
	BrokerUrl *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	// The UNIX timestamp that is generated when the SDK client consumes the first data record.
	//
	// example:
	//
	// 1620962769
	InitCheckPoint *string `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	// The consumer group password.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The consumer group ID.
	//
	// example:
	//
	// HG9
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1611b337285f44e2936a2c4170bbbb7f
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the tracked topic of the change tracking instance.
	//
	// example:
	//
	// TP_TEST_UNDERWRITE_ISSUE
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The consumer group username.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) GetBrokerUrl() *string {
	return s.BrokerUrl
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) GetInitCheckPoint() *string {
	return s.InitCheckPoint
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) GetPassword() *string {
	return s.Password
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) GetSid() *string {
	return s.Sid
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) GetTaskId() *string {
	return s.TaskId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) GetTopic() *string {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) GetUsername() *string {
	return s.Username
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetBrokerUrl(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.BrokerUrl = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetInitCheckPoint(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.InitCheckPoint = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetPassword(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.Password = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetSid(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.Sid = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetTaskId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.TaskId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetTopic(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.Topic = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetUsername(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.Username = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters struct {
	// example:
	//
	// my-event-bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// example:
	//
	// my-event-rule
	EventRuleName *string `json:"EventRuleName,omitempty" xml:"EventRuleName,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters) GetEventBusName() *string {
	return s.EventBusName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters) GetEventRuleName() *string {
	return s.EventRuleName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters) SetEventBusName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters {
	s.EventBusName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters) SetEventRuleName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters {
	s.EventRuleName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceEventBusParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters struct {
	// The group ID of the consumer that subscribes to the topic.
	//
	// example:
	//
	// GID_TEST
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The ID of the ApsaraMQ for Kafka instance.
	//
	// example:
	//
	// bastionhost-cn-i7m2gwt7z1n
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type. Default value: Default. The value PublicNetwork indicates a VPC.
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
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group to which the ApsaraMQ for Kafka instance belongs.
	//
	// example:
	//
	// sg-5ud5f3p0rqqis69tpp8eho7cp
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the topic on the ApsaraMQ for Kafka instance.
	//
	// example:
	//
	// topic_empower_1642473600414
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The ID of the vSwitch with which the ApsaraMQ for Kafka instance is associated.
	//
	// example:
	//
	// vsw-bp1rmi8rind7eo50cbied
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// example:
	//
	// JSON
	ValueDataType *string `json:"ValueDataType,omitempty" xml:"ValueDataType,omitempty"`
	// The ID of the VPC to which the ApsaraMQ for Kafka instance belongs.
	//
	// example:
	//
	// vpc-wz9ki1qdlx3cx5cbbhowf
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GetNetwork() *string {
	return s.Network
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GetValueDataType() *string {
	return s.ValueDataType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetConsumerGroup(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetNetwork(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetOffsetReset(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetSecurityGroupId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetTopic(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetVSwitchIds(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetValueDataType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.ValueDataType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetVpcId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters struct {
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
	// work4
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the MNS queue resides.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) GetIsBase64Decode() *bool {
	return s.IsBase64Decode
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) SetIsBase64Decode(v bool) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) SetQueueName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters struct {
	// example:
	//
	// JSON
	BodyDataType *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance.
	//
	// example:
	//
	// bastionhost-cn-zvp27kcha1r
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the ApsaraMQ for MQTT instance resides.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic on the ApsaraMQ for MQTT instance.
	//
	// example:
	//
	// migration_instance
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) GetTopic() *string {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) SetBodyDataType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters {
	s.BodyDataType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) SetTopic(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters {
	s.Topic = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) GetBucketName() *string {
	return s.BucketName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) GetLoadFormat() *string {
	return s.LoadFormat
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) GetLoadMode() *string {
	return s.LoadMode
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) GetPrefix() *string {
	return s.Prefix
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) SetBucketName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters {
	s.BucketName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) SetDelimiter(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters {
	s.Delimiter = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) SetLoadFormat(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters {
	s.LoadFormat = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) SetLoadMode(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters {
	s.LoadMode = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) SetPrefix(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters {
	s.Prefix = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) SetRoleName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters {
	s.RoleName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOSSParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GetPassword() *string {
	return s.Password
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GetUsername() *string {
	return s.Username
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) SetAuthType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	s.AuthType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) SetBodyDataType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	s.BodyDataType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) SetEndpoint(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	s.Endpoint = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) SetNetworkType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	s.NetworkType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) SetPassword(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	s.Password = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) SetQueueName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) SetSecurityGroupId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) SetUsername(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	s.Username = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) SetVSwitchIds(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) SetVirtualHostName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) SetVpcId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters {
	s.VpcId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceOpenSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters struct {
	// example:
	//
	// c83555068b6******ad213f565f209
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// JSON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// {"source":"prometheus-prod"}
	ExternalLabels *string `json:"ExternalLabels,omitempty" xml:"ExternalLabels,omitempty"`
	// example:
	//
	// __name__=.*
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// testRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) GetDataType() *string {
	return s.DataType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) GetExternalLabels() *string {
	return s.ExternalLabels
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) GetLabels() *string {
	return s.Labels
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) SetClusterId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters {
	s.ClusterId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) SetDataType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters {
	s.DataType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) SetExternalLabels(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters {
	s.ExternalLabels = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) SetLabels(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters {
	s.Labels = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) SetRoleName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters {
	s.RoleName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourcePrometheusParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// dbaudit-cn-7mz2hqolc06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// liuyang
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the ApsaraMQ for RabbitMQ instance resides.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the vhost to which the ApsaraMQ for RabbitMQ instance belongs.
	//
	// example:
	//
	// eb-connect
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) SetQueueName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) SetVirtualHostName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters struct {
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) GetTopics() []*string {
	return s.Topics
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) SetInstanceType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters {
	s.InstanceType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) SetTopics(v []*string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters {
	s.Topics = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQCheckpointParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters struct {
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
	// GID_group1
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// The endpoint that is used to access the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance
	//
	// example:
	//
	// i-f8zbher64dlm58plyfte
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type of the ApsaraMQ for RocketMQ instance. Valid values:
	//
	// 	- PublicNetwork
	//
	// 	- PrivateNetwork
	//
	// example:
	//
	// PublicNetwork
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The password that is used to access the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// 123
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
	// The username that is used to access the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// admin**
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The ID of the vSwitch with which the ApsaraMQ for RocketMQ instance is associated.
	//
	// example:
	//
	// vsw-m5ev8asdc6h123456****
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the ApsaraMQ for RocketMQ instance belongs.
	//
	// example:
	//
	// vpc-bp1a4gmlk31hy***l3ss
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// example:
	//
	// PublicNetwork
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The offset from which messages are consumed. Valid values: CONSUMEFROMLASTOFFSET: Messages are consumed from the latest offset. CONSUMEFROMFIRSTOFFSET: Messages are consumed from the earliest offset. CONSUME_FROM_TIMESTAMP: Messages are consumed from the offset at the specified point in time.
	//
	// example:
	//
	// CONSUMEFROMTIMESTAMP
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the region where the ApsaraMQ for RocketMQ instance resides.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-m5edtu24f12345****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The tag that is used to filter messages.
	//
	// example:
	//
	// v1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp that indicates the time from which messages are consumed. This parameter is valid only if Offset is set to CONSUMEFROMTIMESTAMP.
	//
	// example:
	//
	// 1670742074043
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The topic from which messages are sent.
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

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetBodyDataType() *string {
	return s.BodyDataType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetFilterSql() *string {
	return s.FilterSql
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetFilterType() *string {
	return s.FilterType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetGroupID() *string {
	return s.GroupID
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetInstanceNetwork() *string {
	return s.InstanceNetwork
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetInstanceSecurityGroupId() *string {
	return s.InstanceSecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetInstanceVSwitchIds() *string {
	return s.InstanceVSwitchIds
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetInstanceVpcId() *string {
	return s.InstanceVpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetNetwork() *string {
	return s.Network
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetOffset() *string {
	return s.Offset
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetTag() *string {
	return s.Tag
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetTopic() *string {
	return s.Topic
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetAuthType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetBodyDataType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.BodyDataType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetFilterSql(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.FilterSql = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetFilterType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.FilterType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetGroupID(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceEndpoint(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceNetwork(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstancePassword(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceUsername(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceVpcId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetNetwork(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.Network = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetOffset(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetSecurityGroupId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetTag(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetTimestamp(v int64) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetTopic(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.Topic = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetVSwitchIds(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetVpcId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.VpcId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters struct {
	// The consumer offset. The value begin indicates the earliest offset. The value end indicates the latest offset. You can also specify a time in seconds to start message consumption.
	//
	// example:
	//
	// begin
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The group ID of the consumer that subscribes to the topic.
	//
	// example:
	//
	// go-dts-shelf-group
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The Simple Log Service Logstore.
	//
	// example:
	//
	// waf-logstore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Simple Log Service project.
	//
	// example:
	//
	// dmmzk
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Simple Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console.
	//
	// example:
	//
	// testRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) GetConsumePosition() *string {
	return s.ConsumePosition
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) GetLogStore() *string {
	return s.LogStore
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) GetProject() *string {
	return s.Project
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) SetConsumePosition(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) SetConsumerGroup(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) SetLogStore(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) SetProject(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) SetRoleName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters {
	s.RoleName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) Validate() error {
	return dara.Validate(s)
}

type ListEventStreamingsResponseBodyDataEventStreamingsTransforms struct {
	// The Alibaba Cloud Resource Name (ARN) of the cloud service, such as the ARN of a Function Compute function.
	//
	// example:
	//
	// acs:fc:cn-hangzhou:*****:services/demo-service.LATEST/functions/demo-func
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsTransforms) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsTransforms) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsTransforms) GetArn() *string {
	return s.Arn
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsTransforms) SetArn(v string) *ListEventStreamingsResponseBodyDataEventStreamingsTransforms {
	s.Arn = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsTransforms) Validate() error {
	return dara.Validate(s)
}
