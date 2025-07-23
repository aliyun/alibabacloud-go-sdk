// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetSubscriptionAttributesResponseBody
	GetCode() *int64
	SetData(v *GetSubscriptionAttributesResponseBodyData) *GetSubscriptionAttributesResponseBody
	GetData() *GetSubscriptionAttributesResponseBodyData
	SetMessage(v string) *GetSubscriptionAttributesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSubscriptionAttributesResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetSubscriptionAttributesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetSubscriptionAttributesResponseBody
	GetSuccess() *bool
}

type GetSubscriptionAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetSubscriptionAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetSubscriptionAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetSubscriptionAttributesResponseBody) GetData() *GetSubscriptionAttributesResponseBodyData {
	return s.Data
}

func (s *GetSubscriptionAttributesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubscriptionAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubscriptionAttributesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSubscriptionAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSubscriptionAttributesResponseBody) SetCode(v int64) *GetSubscriptionAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetData(v *GetSubscriptionAttributesResponseBodyData) *GetSubscriptionAttributesResponseBody {
	s.Data = v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetMessage(v string) *GetSubscriptionAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetRequestId(v string) *GetSubscriptionAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetStatus(v string) *GetSubscriptionAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetSuccess(v bool) *GetSubscriptionAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSubscriptionAttributesResponseBodyData struct {
	// The time when the subscription was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554806
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *GetSubscriptionAttributesResponseBodyDataDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
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
	// 1449554962
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
	SubscriptionName      *string                                                         `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	TenantRateLimitPolicy *GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty" type:"Struct"`
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

func (s GetSubscriptionAttributesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetSubscriptionAttributesResponseBodyData) GetDlqPolicy() *GetSubscriptionAttributesResponseBodyDataDlqPolicy {
	return s.DlqPolicy
}

func (s *GetSubscriptionAttributesResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetSubscriptionAttributesResponseBodyData) GetFilterTag() *string {
	return s.FilterTag
}

func (s *GetSubscriptionAttributesResponseBodyData) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *GetSubscriptionAttributesResponseBodyData) GetNotifyContentFormat() *string {
	return s.NotifyContentFormat
}

func (s *GetSubscriptionAttributesResponseBodyData) GetNotifyStrategy() *string {
	return s.NotifyStrategy
}

func (s *GetSubscriptionAttributesResponseBodyData) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *GetSubscriptionAttributesResponseBodyData) GetTenantRateLimitPolicy() *GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy {
	return s.TenantRateLimitPolicy
}

func (s *GetSubscriptionAttributesResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *GetSubscriptionAttributesResponseBodyData) GetTopicOwner() *string {
	return s.TopicOwner
}

func (s *GetSubscriptionAttributesResponseBodyData) SetCreateTime(v int64) *GetSubscriptionAttributesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetDlqPolicy(v *GetSubscriptionAttributesResponseBodyDataDlqPolicy) *GetSubscriptionAttributesResponseBodyData {
	s.DlqPolicy = v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetEndpoint(v string) *GetSubscriptionAttributesResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetFilterTag(v string) *GetSubscriptionAttributesResponseBodyData {
	s.FilterTag = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetLastModifyTime(v int64) *GetSubscriptionAttributesResponseBodyData {
	s.LastModifyTime = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetNotifyContentFormat(v string) *GetSubscriptionAttributesResponseBodyData {
	s.NotifyContentFormat = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetNotifyStrategy(v string) *GetSubscriptionAttributesResponseBodyData {
	s.NotifyStrategy = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetSubscriptionName(v string) *GetSubscriptionAttributesResponseBodyData {
	s.SubscriptionName = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetTenantRateLimitPolicy(v *GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy) *GetSubscriptionAttributesResponseBodyData {
	s.TenantRateLimitPolicy = v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetTopicName(v string) *GetSubscriptionAttributesResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetTopicOwner(v string) *GetSubscriptionAttributesResponseBodyData {
	s.TopicOwner = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetSubscriptionAttributesResponseBodyDataDlqPolicy struct {
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
}

func (s GetSubscriptionAttributesResponseBodyDataDlqPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionAttributesResponseBodyDataDlqPolicy) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponseBodyDataDlqPolicy) GetDeadLetterTargetQueue() *string {
	return s.DeadLetterTargetQueue
}

func (s *GetSubscriptionAttributesResponseBodyDataDlqPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetSubscriptionAttributesResponseBodyDataDlqPolicy) SetDeadLetterTargetQueue(v string) *GetSubscriptionAttributesResponseBodyDataDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyDataDlqPolicy) SetEnabled(v bool) *GetSubscriptionAttributesResponseBodyDataDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyDataDlqPolicy) Validate() error {
	return dara.Validate(s)
}

type GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy struct {
	Enabled              *bool  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	MaxReceivesPerSecond *int32 `json:"MaxReceivesPerSecond,omitempty" xml:"MaxReceivesPerSecond,omitempty"`
}

func (s GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy) GetMaxReceivesPerSecond() *int32 {
	return s.MaxReceivesPerSecond
}

func (s *GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy) SetEnabled(v bool) *GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy {
	s.Enabled = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy) SetMaxReceivesPerSecond(v int32) *GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy {
	s.MaxReceivesPerSecond = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyDataTenantRateLimitPolicy) Validate() error {
	return dara.Validate(s)
}
