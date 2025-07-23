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
	// The data returned.
	Data *GetQueueAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
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
	return dara.Validate(s)
}

type GetQueueAttributesResponseBodyData struct {
	// The total number of messages that are in the Active state in the queue. The value is an approximate value. Default value: 0. We recommend that you do not use the return value and that you call CloudMonitor API operations to query the metric value.
	//
	// example:
	//
	// 20
	ActiveMessages *int64 `json:"ActiveMessages,omitempty" xml:"ActiveMessages,omitempty"`
	// The time when the queue was created.
	//
	// example:
	//
	// 1250700999
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The total number of messages that are in the Delayed state in the queue. The value is an approximate value. Default value: 0. We recommend that you do not use the return value and that you call CloudMonitor API operations to query the metric value.
	//
	// example:
	//
	// 0
	DelayMessages *int64 `json:"DelayMessages,omitempty" xml:"DelayMessages,omitempty"`
	// The period after which all messages sent to the queue are consumed. Unit: seconds.
	//
	// example:
	//
	// 30
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *GetQueueAttributesResponseBodyDataDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// The total number of messages that are in the Inactive state in the queue. The value is an approximate value. Default value: 0. We recommend that you do not use the return value and that you call CloudMonitor API operations to query the metric value.
	//
	// example:
	//
	// 0
	InactiveMessages *int64 `json:"InactiveMessages,omitempty" xml:"InactiveMessages,omitempty"`
	// The time when the queue was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1250700999
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// Indicates whether the logging feature is enabled. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	LoggingEnabled *bool `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
	// The maximum length of the message that is sent to the queue. Unit: bytes.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Unit: seconds.
	//
	// example:
	//
	// 65536
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Unit: seconds.
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
	// The tag.
	Tags                  []*GetQueueAttributesResponseBodyDataTags                `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TenantRateLimitPolicy *GetQueueAttributesResponseBodyDataTenantRateLimitPolicy `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty" type:"Struct"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
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

func (s *GetQueueAttributesResponseBodyData) GetInactiveMessages() *int64 {
	return s.InactiveMessages
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

func (s *GetQueueAttributesResponseBodyData) SetInactiveMessages(v int64) *GetQueueAttributesResponseBodyData {
	s.InactiveMessages = &v
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
	return dara.Validate(s)
}

type GetQueueAttributesResponseBodyDataDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// deadLetterTargetQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable the dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of retries.
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
	// The tag key.
	//
	// example:
	//
	// tag1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
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
	Enabled              *bool  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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
