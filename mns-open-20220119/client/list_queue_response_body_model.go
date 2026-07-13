// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListQueueResponseBody
	GetCode() *int64
	SetData(v *ListQueueResponseBodyData) *ListQueueResponseBody
	GetData() *ListQueueResponseBodyData
	SetMessage(v string) *ListQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListQueueResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListQueueResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ListQueueResponseBody
	GetSuccess() *bool
}

type ListQueueResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *ListQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueueResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListQueueResponseBody) GetData() *ListQueueResponseBodyData {
	return s.Data
}

func (s *ListQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueueResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQueueResponseBody) SetCode(v int64) *ListQueueResponseBody {
	s.Code = &v
	return s
}

func (s *ListQueueResponseBody) SetData(v *ListQueueResponseBodyData) *ListQueueResponseBody {
	s.Data = v
	return s
}

func (s *ListQueueResponseBody) SetMessage(v string) *ListQueueResponseBody {
	s.Message = &v
	return s
}

func (s *ListQueueResponseBody) SetRequestId(v string) *ListQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueueResponseBody) SetStatus(v string) *ListQueueResponseBody {
	s.Status = &v
	return s
}

func (s *ListQueueResponseBody) SetSuccess(v bool) *ListQueueResponseBody {
	s.Success = &v
	return s
}

func (s *ListQueueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQueueResponseBodyData struct {
	// The results returned on the current page.
	PageData []*ListQueueResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// The page number of the returned results.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 3
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 130
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyData) GetPageData() []*ListQueueResponseBodyDataPageData {
	return s.PageData
}

func (s *ListQueueResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListQueueResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListQueueResponseBodyData) GetPages() *int64 {
	return s.Pages
}

func (s *ListQueueResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *ListQueueResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListQueueResponseBodyData) SetPageData(v []*ListQueueResponseBodyDataPageData) *ListQueueResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListQueueResponseBodyData) SetPageNum(v int64) *ListQueueResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListQueueResponseBodyData) SetPageSize(v int64) *ListQueueResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListQueueResponseBodyData) SetPages(v int64) *ListQueueResponseBodyData {
	s.Pages = &v
	return s
}

func (s *ListQueueResponseBodyData) SetSize(v int64) *ListQueueResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListQueueResponseBodyData) SetTotal(v int64) *ListQueueResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListQueueResponseBodyData) Validate() error {
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQueueResponseBodyDataPageData struct {
	// The approximate total number of messages in the Active state in this queue.
	//
	// This field will default to 0 in the future and is not recommended. Use CloudMonitor API to retrieve this metric instead.
	//
	// example:
	//
	// 0
	ActiveMessages *int64 `json:"ActiveMessages,omitempty" xml:"ActiveMessages,omitempty"`
	// The time when the queue was created. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1250700999
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The approximate total number of messages in the Delayed state in this queue.
	//
	// This field will default to 0 in the future and is not recommended. Use CloudMonitor API to retrieve this metric instead.
	//
	// example:
	//
	// 0
	DelayMessages *int64 `json:"DelayMessages,omitempty" xml:"DelayMessages,omitempty"`
	// The delay period after which all messages sent to this queue become consumable. Unit: seconds.
	//
	// example:
	//
	// 30
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy         *ListQueueResponseBodyDataPageDataDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	EnableSSE         *bool                                       `json:"EnableSSE,omitempty" xml:"EnableSSE,omitempty"`
	EncryptionEnabled *bool                                       `json:"EncryptionEnabled,omitempty" xml:"EncryptionEnabled,omitempty"`
	// The approximate total number of messages in the Inactive state in this queue.
	//
	// This field will default to 0 in the future and is not recommended. Use CloudMonitor API to retrieve this metric instead.
	//
	// example:
	//
	// 0
	InactiveMessages *int64  `json:"InactiveMessages,omitempty" xml:"InactiveMessages,omitempty"`
	KmsKeyId         *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The most recent time when the queue attributes were modified. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
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
	// The maximum size of a message body that can be sent to this queue. Unit: bytes.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum period for which a message can be retained in this queue. After the specified period elapses since a message is sent to the queue, the message is deleted regardless of whether it has been consumed. Unit: seconds.
	//
	// example:
	//
	// 65536
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum wait time for a ReceiveMessage request when the queue is empty. Unit: seconds.
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
	Tags []*ListQueueResponseBodyDataPageDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The duration for which a message stays in the Inactive state after it is consumed from the queue.
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

func (s ListQueueResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListQueueResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyDataPageData) GetActiveMessages() *int64 {
	return s.ActiveMessages
}

func (s *ListQueueResponseBodyDataPageData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListQueueResponseBodyDataPageData) GetDelayMessages() *int64 {
	return s.DelayMessages
}

func (s *ListQueueResponseBodyDataPageData) GetDelaySeconds() *int64 {
	return s.DelaySeconds
}

func (s *ListQueueResponseBodyDataPageData) GetDlqPolicy() *ListQueueResponseBodyDataPageDataDlqPolicy {
	return s.DlqPolicy
}

func (s *ListQueueResponseBodyDataPageData) GetEnableSSE() *bool {
	return s.EnableSSE
}

func (s *ListQueueResponseBodyDataPageData) GetEncryptionEnabled() *bool {
	return s.EncryptionEnabled
}

func (s *ListQueueResponseBodyDataPageData) GetInactiveMessages() *int64 {
	return s.InactiveMessages
}

func (s *ListQueueResponseBodyDataPageData) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *ListQueueResponseBodyDataPageData) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *ListQueueResponseBodyDataPageData) GetLoggingEnabled() *bool {
	return s.LoggingEnabled
}

func (s *ListQueueResponseBodyDataPageData) GetMaximumMessageSize() *int64 {
	return s.MaximumMessageSize
}

func (s *ListQueueResponseBodyDataPageData) GetMessageRetentionPeriod() *int64 {
	return s.MessageRetentionPeriod
}

func (s *ListQueueResponseBodyDataPageData) GetPollingWaitSeconds() *int64 {
	return s.PollingWaitSeconds
}

func (s *ListQueueResponseBodyDataPageData) GetQueueName() *string {
	return s.QueueName
}

func (s *ListQueueResponseBodyDataPageData) GetQueueType() *string {
	return s.QueueType
}

func (s *ListQueueResponseBodyDataPageData) GetSseAlgorithm() *string {
	return s.SseAlgorithm
}

func (s *ListQueueResponseBodyDataPageData) GetSseType() *string {
	return s.SseType
}

func (s *ListQueueResponseBodyDataPageData) GetTags() []*ListQueueResponseBodyDataPageDataTags {
	return s.Tags
}

func (s *ListQueueResponseBodyDataPageData) GetVisibilityTimeout() *int64 {
	return s.VisibilityTimeout
}

func (s *ListQueueResponseBodyDataPageData) SetActiveMessages(v int64) *ListQueueResponseBodyDataPageData {
	s.ActiveMessages = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetCreateTime(v int64) *ListQueueResponseBodyDataPageData {
	s.CreateTime = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetDelayMessages(v int64) *ListQueueResponseBodyDataPageData {
	s.DelayMessages = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetDelaySeconds(v int64) *ListQueueResponseBodyDataPageData {
	s.DelaySeconds = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetDlqPolicy(v *ListQueueResponseBodyDataPageDataDlqPolicy) *ListQueueResponseBodyDataPageData {
	s.DlqPolicy = v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetEnableSSE(v bool) *ListQueueResponseBodyDataPageData {
	s.EnableSSE = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetEncryptionEnabled(v bool) *ListQueueResponseBodyDataPageData {
	s.EncryptionEnabled = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetInactiveMessages(v int64) *ListQueueResponseBodyDataPageData {
	s.InactiveMessages = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetKmsKeyId(v string) *ListQueueResponseBodyDataPageData {
	s.KmsKeyId = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetLastModifyTime(v int64) *ListQueueResponseBodyDataPageData {
	s.LastModifyTime = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetLoggingEnabled(v bool) *ListQueueResponseBodyDataPageData {
	s.LoggingEnabled = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetMaximumMessageSize(v int64) *ListQueueResponseBodyDataPageData {
	s.MaximumMessageSize = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetMessageRetentionPeriod(v int64) *ListQueueResponseBodyDataPageData {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetPollingWaitSeconds(v int64) *ListQueueResponseBodyDataPageData {
	s.PollingWaitSeconds = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetQueueName(v string) *ListQueueResponseBodyDataPageData {
	s.QueueName = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetQueueType(v string) *ListQueueResponseBodyDataPageData {
	s.QueueType = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetSseAlgorithm(v string) *ListQueueResponseBodyDataPageData {
	s.SseAlgorithm = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetSseType(v string) *ListQueueResponseBodyDataPageData {
	s.SseType = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetTags(v []*ListQueueResponseBodyDataPageDataTags) *ListQueueResponseBodyDataPageData {
	s.Tags = v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetVisibilityTimeout(v int64) *ListQueueResponseBodyDataPageData {
	s.VisibilityTimeout = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) Validate() error {
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
	return nil
}

type ListQueueResponseBodyDataPageDataDlqPolicy struct {
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

func (s ListQueueResponseBodyDataPageDataDlqPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListQueueResponseBodyDataPageDataDlqPolicy) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyDataPageDataDlqPolicy) GetDeadLetterTargetQueue() *string {
	return s.DeadLetterTargetQueue
}

func (s *ListQueueResponseBodyDataPageDataDlqPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListQueueResponseBodyDataPageDataDlqPolicy) GetMaxReceiveCount() *string {
	return s.MaxReceiveCount
}

func (s *ListQueueResponseBodyDataPageDataDlqPolicy) SetDeadLetterTargetQueue(v string) *ListQueueResponseBodyDataPageDataDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *ListQueueResponseBodyDataPageDataDlqPolicy) SetEnabled(v bool) *ListQueueResponseBodyDataPageDataDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *ListQueueResponseBodyDataPageDataDlqPolicy) SetMaxReceiveCount(v string) *ListQueueResponseBodyDataPageDataDlqPolicy {
	s.MaxReceiveCount = &v
	return s
}

func (s *ListQueueResponseBodyDataPageDataDlqPolicy) Validate() error {
	return dara.Validate(s)
}

type ListQueueResponseBodyDataPageDataTags struct {
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

func (s ListQueueResponseBodyDataPageDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListQueueResponseBodyDataPageDataTags) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyDataPageDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListQueueResponseBodyDataPageDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListQueueResponseBodyDataPageDataTags) SetTagKey(v string) *ListQueueResponseBodyDataPageDataTags {
	s.TagKey = &v
	return s
}

func (s *ListQueueResponseBodyDataPageDataTags) SetTagValue(v string) *ListQueueResponseBodyDataPageDataTags {
	s.TagValue = &v
	return s
}

func (s *ListQueueResponseBodyDataPageDataTags) Validate() error {
	return dara.Validate(s)
}
