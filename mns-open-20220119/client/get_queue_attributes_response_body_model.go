// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetQueueAttributesResponseBody
	GetCode() *int64
	SetData(v *GetQueueAttributesResponseBodyData) *GetQueueAttributesResponseBody
	GetData() *GetQueueAttributesResponseBodyData
	SetMessage(v string) *GetQueueAttributesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueueAttributesResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetQueueAttributesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetQueueAttributesResponseBody
	GetSuccess() *bool
}

type GetQueueAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetQueueAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the response.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueueAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueueAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetQueueAttributesResponseBody) GetData() *GetQueueAttributesResponseBodyData {
	return s.Data
}

func (s *GetQueueAttributesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueueAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueueAttributesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetQueueAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQueueAttributesResponseBody) SetCode(v int64) *GetQueueAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetData(v *GetQueueAttributesResponseBodyData) *GetQueueAttributesResponseBody {
	s.Data = v
	return s
}

func (s *GetQueueAttributesResponseBody) SetMessage(v string) *GetQueueAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetRequestId(v string) *GetQueueAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetStatus(v string) *GetQueueAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetSuccess(v bool) *GetQueueAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueueAttributesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueueAttributesResponseBodyData struct {
	// The approximate total number of messages in the Active state in the queue.
	//
	// <warning>This field will be deprecated and defaults to 0. Use the CloudMonitor API to retrieve this metric instead.</warning>
	//
	// example:
	//
	// 0
	ActiveMessages *int64 `json:"ActiveMessages,omitempty" xml:"ActiveMessages,omitempty"`
	// The time when the queue was created.
	//
	// example:
	//
	// 1250700999
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The approximate total number of messages in the Delayed state in the queue.
	//
	// <warning>This field will be deprecated and defaults to 0. Use the CloudMonitor API to retrieve this metric instead.</warning>
	//
	// example:
	//
	// 0
	DelayMessages *int64 `json:"DelayMessages,omitempty" xml:"DelayMessages,omitempty"`
	// The delay period for all messages sent to the queue. Messages sent to the queue can be consumed only after the delay period specified by this parameter elapses. Unit: seconds.
	//
	// example:
	//
	// 30
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy         *GetQueueAttributesResponseBodyDataDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	EnableSSE         *bool                                        `json:"EnableSSE,omitempty" xml:"EnableSSE,omitempty"`
	EncryptionEnabled *bool                                        `json:"EncryptionEnabled,omitempty" xml:"EncryptionEnabled,omitempty"`
	// The approximate total number of messages in the Inactive state in the queue.
	//
	// <warning>This field will be deprecated and defaults to 0. Use the CloudMonitor API to retrieve this metric instead.</warning>
	//
	// example:
	//
	// 0
	InactiveMessages *int64  `json:"InactiveMessages,omitempty" xml:"InactiveMessages,omitempty"`
	KmsKeyId         *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The most recent time when the queue attributes were modified. The value is a UNIX timestamp representing the number of seconds elapsed since 1970-01-01 00:00:00.
	//
	// example:
	//
	// 1250700999
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// Indicates whether the log management feature is enabled.
	//
	// - True: Enabled.
	//
	// - False: Disabled.
	//
	// example:
	//
	// True
	LoggingEnabled *bool `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
	// The maximum length of the message body sent to the queue. Unit: bytes.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the period specified by this parameter elapses since the message is sent to the queue, the message is deleted regardless of whether it has been consumed. Unit: seconds.
	//
	// example:
	//
	// 65536
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum wait time for a ReceiveMessage request on the queue when the queue has no messages. Unit: seconds.
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// demo-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The type of the queue. Valid values:
	//
	//    	- normal: standard queue
	//
	//    	- fifo: FIFO queue
	//
	// example:
	//
	// normal
	QueueType    *string `json:"QueueType,omitempty" xml:"QueueType,omitempty"`
	SseAlgorithm *string `json:"SseAlgorithm,omitempty" xml:"SseAlgorithm,omitempty"`
	SseType      *string `json:"SseType,omitempty" xml:"SseType,omitempty"`
	// The list of resource tags.
	Tags []*GetQueueAttributesResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The rate limiting policy.
	TenantRateLimitPolicy *GetQueueAttributesResponseBodyDataTenantRateLimitPolicy `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty" type:"Struct"`
	// The duration for which a message stays in the Inactive state after it is consumed from the queue and changes from the Active state to the Inactive state.
	//
	// Valid values: 1 to 43200. Unit: seconds.
	//
	// Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s GetQueueAttributesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueueAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBodyData) GetActiveMessages() *int64 {
	return s.ActiveMessages
}

func (s *GetQueueAttributesResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetQueueAttributesResponseBodyData) GetDelayMessages() *int64 {
	return s.DelayMessages
}

func (s *GetQueueAttributesResponseBodyData) GetDelaySeconds() *int64 {
	return s.DelaySeconds
}

func (s *GetQueueAttributesResponseBodyData) GetDlqPolicy() *GetQueueAttributesResponseBodyDataDlqPolicy {
	return s.DlqPolicy
}

func (s *GetQueueAttributesResponseBodyData) GetEnableSSE() *bool {
	return s.EnableSSE
}

func (s *GetQueueAttributesResponseBodyData) GetEncryptionEnabled() *bool {
	return s.EncryptionEnabled
}

func (s *GetQueueAttributesResponseBodyData) GetInactiveMessages() *int64 {
	return s.InactiveMessages
}

func (s *GetQueueAttributesResponseBodyData) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *GetQueueAttributesResponseBodyData) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *GetQueueAttributesResponseBodyData) GetLoggingEnabled() *bool {
	return s.LoggingEnabled
}

func (s *GetQueueAttributesResponseBodyData) GetMaximumMessageSize() *int64 {
	return s.MaximumMessageSize
}

func (s *GetQueueAttributesResponseBodyData) GetMessageRetentionPeriod() *int64 {
	return s.MessageRetentionPeriod
}

func (s *GetQueueAttributesResponseBodyData) GetPollingWaitSeconds() *int64 {
	return s.PollingWaitSeconds
}

func (s *GetQueueAttributesResponseBodyData) GetQueueName() *string {
	return s.QueueName
}

func (s *GetQueueAttributesResponseBodyData) GetQueueType() *string {
	return s.QueueType
}

func (s *GetQueueAttributesResponseBodyData) GetSseAlgorithm() *string {
	return s.SseAlgorithm
}

func (s *GetQueueAttributesResponseBodyData) GetSseType() *string {
	return s.SseType
}

func (s *GetQueueAttributesResponseBodyData) GetTags() []*GetQueueAttributesResponseBodyDataTags {
	return s.Tags
}

func (s *GetQueueAttributesResponseBodyData) GetTenantRateLimitPolicy() *GetQueueAttributesResponseBodyDataTenantRateLimitPolicy {
	return s.TenantRateLimitPolicy
}

func (s *GetQueueAttributesResponseBodyData) GetVisibilityTimeout() *int64 {
	return s.VisibilityTimeout
}

func (s *GetQueueAttributesResponseBodyData) SetActiveMessages(v int64) *GetQueueAttributesResponseBodyData {
	s.ActiveMessages = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetCreateTime(v int64) *GetQueueAttributesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetDelayMessages(v int64) *GetQueueAttributesResponseBodyData {
	s.DelayMessages = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetDelaySeconds(v int64) *GetQueueAttributesResponseBodyData {
	s.DelaySeconds = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetDlqPolicy(v *GetQueueAttributesResponseBodyDataDlqPolicy) *GetQueueAttributesResponseBodyData {
	s.DlqPolicy = v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetEnableSSE(v bool) *GetQueueAttributesResponseBodyData {
	s.EnableSSE = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetEncryptionEnabled(v bool) *GetQueueAttributesResponseBodyData {
	s.EncryptionEnabled = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetInactiveMessages(v int64) *GetQueueAttributesResponseBodyData {
	s.InactiveMessages = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetKmsKeyId(v string) *GetQueueAttributesResponseBodyData {
	s.KmsKeyId = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetLastModifyTime(v int64) *GetQueueAttributesResponseBodyData {
	s.LastModifyTime = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetLoggingEnabled(v bool) *GetQueueAttributesResponseBodyData {
	s.LoggingEnabled = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetMaximumMessageSize(v int64) *GetQueueAttributesResponseBodyData {
	s.MaximumMessageSize = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetMessageRetentionPeriod(v int64) *GetQueueAttributesResponseBodyData {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetPollingWaitSeconds(v int64) *GetQueueAttributesResponseBodyData {
	s.PollingWaitSeconds = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetQueueName(v string) *GetQueueAttributesResponseBodyData {
	s.QueueName = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetQueueType(v string) *GetQueueAttributesResponseBodyData {
	s.QueueType = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetSseAlgorithm(v string) *GetQueueAttributesResponseBodyData {
	s.SseAlgorithm = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetSseType(v string) *GetQueueAttributesResponseBodyData {
	s.SseType = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetTags(v []*GetQueueAttributesResponseBodyDataTags) *GetQueueAttributesResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetTenantRateLimitPolicy(v *GetQueueAttributesResponseBodyDataTenantRateLimitPolicy) *GetQueueAttributesResponseBodyData {
	s.TenantRateLimitPolicy = v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetVisibilityTimeout(v int64) *GetQueueAttributesResponseBodyData {
	s.VisibilityTimeout = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) Validate() error {
	if s.DlqPolicy != nil {
		if err := s.DlqPolicy.Validate(); err != nil {
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
	if s.TenantRateLimitPolicy != nil {
		if err := s.TenantRateLimitPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueueAttributesResponseBodyDataDlqPolicy struct {
	// The target queue for dead-letter message delivery.
	//
	// example:
	//
	// deadLetterTargetQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Indicates whether dead-letter message delivery is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of times a message can be delivered.
	//
	// example:
	//
	// 3
	MaxReceiveCount *string `json:"MaxReceiveCount,omitempty" xml:"MaxReceiveCount,omitempty"`
}

func (s GetQueueAttributesResponseBodyDataDlqPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetQueueAttributesResponseBodyDataDlqPolicy) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBodyDataDlqPolicy) GetDeadLetterTargetQueue() *string {
	return s.DeadLetterTargetQueue
}

func (s *GetQueueAttributesResponseBodyDataDlqPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetQueueAttributesResponseBodyDataDlqPolicy) GetMaxReceiveCount() *string {
	return s.MaxReceiveCount
}

func (s *GetQueueAttributesResponseBodyDataDlqPolicy) SetDeadLetterTargetQueue(v string) *GetQueueAttributesResponseBodyDataDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *GetQueueAttributesResponseBodyDataDlqPolicy) SetEnabled(v bool) *GetQueueAttributesResponseBodyDataDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *GetQueueAttributesResponseBodyDataDlqPolicy) SetMaxReceiveCount(v string) *GetQueueAttributesResponseBodyDataDlqPolicy {
	s.MaxReceiveCount = &v
	return s
}

func (s *GetQueueAttributesResponseBodyDataDlqPolicy) Validate() error {
	return dara.Validate(s)
}

type GetQueueAttributesResponseBodyDataTags struct {
	// The key of the tag.
	//
	// example:
	//
	// tag1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetQueueAttributesResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetQueueAttributesResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBodyDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetQueueAttributesResponseBodyDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetQueueAttributesResponseBodyDataTags) SetTagKey(v string) *GetQueueAttributesResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *GetQueueAttributesResponseBodyDataTags) SetTagValue(v string) *GetQueueAttributesResponseBodyDataTags {
	s.TagValue = &v
	return s
}

func (s *GetQueueAttributesResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}

type GetQueueAttributesResponseBodyDataTenantRateLimitPolicy struct {
	// Specifies whether rate limiting is enabled. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of receives per second.
	//
	// example:
	//
	// 1000
	MaxReceivesPerSecond *int32 `json:"MaxReceivesPerSecond,omitempty" xml:"MaxReceivesPerSecond,omitempty"`
}

func (s GetQueueAttributesResponseBodyDataTenantRateLimitPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetQueueAttributesResponseBodyDataTenantRateLimitPolicy) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBodyDataTenantRateLimitPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetQueueAttributesResponseBodyDataTenantRateLimitPolicy) GetMaxReceivesPerSecond() *int32 {
	return s.MaxReceivesPerSecond
}

func (s *GetQueueAttributesResponseBodyDataTenantRateLimitPolicy) SetEnabled(v bool) *GetQueueAttributesResponseBodyDataTenantRateLimitPolicy {
	s.Enabled = &v
	return s
}

func (s *GetQueueAttributesResponseBodyDataTenantRateLimitPolicy) SetMaxReceivesPerSecond(v int32) *GetQueueAttributesResponseBodyDataTenantRateLimitPolicy {
	s.MaxReceivesPerSecond = &v
	return s
}

func (s *GetQueueAttributesResponseBodyDataTenantRateLimitPolicy) Validate() error {
	return dara.Validate(s)
}
