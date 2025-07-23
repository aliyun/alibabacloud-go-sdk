// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionByTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListSubscriptionByTopicResponseBody
	GetCode() *int64
	SetData(v *ListSubscriptionByTopicResponseBodyData) *ListSubscriptionByTopicResponseBody
	GetData() *ListSubscriptionByTopicResponseBodyData
	SetMessage(v string) *ListSubscriptionByTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSubscriptionByTopicResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListSubscriptionByTopicResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ListSubscriptionByTopicResponseBody
	GetSuccess() *bool
}

type ListSubscriptionByTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListSubscriptionByTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListSubscriptionByTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListSubscriptionByTopicResponseBody) GetData() *ListSubscriptionByTopicResponseBodyData {
	return s.Data
}

func (s *ListSubscriptionByTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSubscriptionByTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubscriptionByTopicResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListSubscriptionByTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSubscriptionByTopicResponseBody) SetCode(v int64) *ListSubscriptionByTopicResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetData(v *ListSubscriptionByTopicResponseBodyData) *ListSubscriptionByTopicResponseBody {
	s.Data = v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetMessage(v string) *ListSubscriptionByTopicResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetRequestId(v string) *ListSubscriptionByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetStatus(v string) *ListSubscriptionByTopicResponseBody {
	s.Status = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetSuccess(v bool) *ListSubscriptionByTopicResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSubscriptionByTopicResponseBodyData struct {
	// The data returned on the current page.
	PageData []*ListSubscriptionByTopicResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 3
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 130
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSubscriptionByTopicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionByTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponseBodyData) GetPageData() []*ListSubscriptionByTopicResponseBodyDataPageData {
	return s.PageData
}

func (s *ListSubscriptionByTopicResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListSubscriptionByTopicResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSubscriptionByTopicResponseBodyData) GetPages() *int64 {
	return s.Pages
}

func (s *ListSubscriptionByTopicResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *ListSubscriptionByTopicResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPageData(v []*ListSubscriptionByTopicResponseBodyDataPageData) *ListSubscriptionByTopicResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPageNum(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPageSize(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPages(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.Pages = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetSize(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetTotal(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListSubscriptionByTopicResponseBodyDataPageData struct {
	// The time when the subscription was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554806
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// The endpoint to which the messages are pushed.
	//
	// example:
	//
	// http://example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The tag that is used to filter messages. Only the messages that are attached with the specified tag can be pushed.
	//
	// example:
	//
	// important
	FilterTag *string `json:"FilterTag,omitempty" xml:"FilterTag,omitempty"`
	// The time when the subscription was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554806
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The content format of the messages that are pushed to the endpoint. Valid values:
	//
	// 	- XML
	//
	// 	- JSON
	//
	// 	- SIMPLIFIED
	//
	// example:
	//
	// XML
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" xml:"NotifyContentFormat,omitempty"`
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The name of the subscription.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// MyTopic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The Alibaba Cloud account ID of the topic owner.
	//
	// example:
	//
	// 123456789098****
	TopicOwner *string `json:"TopicOwner,omitempty" xml:"TopicOwner,omitempty"`
}

func (s ListSubscriptionByTopicResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionByTopicResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) GetDlqPolicy() *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy {
	return s.DlqPolicy
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) GetFilterTag() *string {
	return s.FilterTag
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) GetNotifyContentFormat() *string {
	return s.NotifyContentFormat
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) GetNotifyStrategy() *string {
	return s.NotifyStrategy
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) GetTopicName() *string {
	return s.TopicName
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) GetTopicOwner() *string {
	return s.TopicOwner
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetCreateTime(v int64) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.CreateTime = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetDlqPolicy(v *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.DlqPolicy = v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetEndpoint(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.Endpoint = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetFilterTag(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.FilterTag = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetLastModifyTime(v int64) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.LastModifyTime = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetNotifyContentFormat(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.NotifyContentFormat = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetNotifyStrategy(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.NotifyStrategy = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetSubscriptionName(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.SubscriptionName = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetTopicName(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.TopicName = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetTopicOwner(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.TopicOwner = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}

type ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// dead-letter-queue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable the dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) GetDeadLetterTargetQueue() *string {
	return s.DeadLetterTargetQueue
}

func (s *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) SetDeadLetterTargetQueue(v string) *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) SetEnabled(v bool) *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageDataDlqPolicy) Validate() error {
	return dara.Validate(s)
}
