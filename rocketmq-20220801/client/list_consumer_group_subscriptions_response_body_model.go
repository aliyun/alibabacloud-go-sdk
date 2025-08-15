// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupSubscriptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsumerGroupSubscriptionsResponseBody
	GetCode() *string
	SetData(v []*ListConsumerGroupSubscriptionsResponseBodyData) *ListConsumerGroupSubscriptionsResponseBody
	GetData() []*ListConsumerGroupSubscriptionsResponseBodyData
	SetDynamicCode(v string) *ListConsumerGroupSubscriptionsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListConsumerGroupSubscriptionsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListConsumerGroupSubscriptionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListConsumerGroupSubscriptionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConsumerGroupSubscriptionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListConsumerGroupSubscriptionsResponseBody
	GetSuccess() *bool
}

type ListConsumerGroupSubscriptionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data []*ListConsumerGroupSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5F4D9D5F-625B-59FF-BD4F-DA8284575DB4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListConsumerGroupSubscriptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupSubscriptionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConsumerGroupSubscriptionsResponseBody) GetData() []*ListConsumerGroupSubscriptionsResponseBodyData {
	return s.Data
}

func (s *ListConsumerGroupSubscriptionsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListConsumerGroupSubscriptionsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListConsumerGroupSubscriptionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListConsumerGroupSubscriptionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConsumerGroupSubscriptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumerGroupSubscriptionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetCode(v string) *ListConsumerGroupSubscriptionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetData(v []*ListConsumerGroupSubscriptionsResponseBodyData) *ListConsumerGroupSubscriptionsResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetDynamicCode(v string) *ListConsumerGroupSubscriptionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetDynamicMessage(v string) *ListConsumerGroupSubscriptionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetHttpStatusCode(v int32) *ListConsumerGroupSubscriptionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetMessage(v string) *ListConsumerGroupSubscriptionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetRequestId(v string) *ListConsumerGroupSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetSuccess(v bool) *ListConsumerGroupSubscriptionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListConsumerGroupSubscriptionsResponseBodyData struct {
	// Indicates whether message consumption is consistent. Valid values:
	//
	// 	- false: Unconsumed messages exist in the consumer group.
	//
	// 	- true: No unconsumed message exists in the consumer group.
	//
	// example:
	//
	// true
	Consistency *bool `json:"consistency,omitempty" xml:"consistency,omitempty"`
	// The ID of the consumer group.
	//
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The filter expression.
	//
	// example:
	//
	// *
	FilterExpression *string `json:"filterExpression,omitempty" xml:"filterExpression,omitempty"`
	// The type of the filter expression. Valid values:
	//
	// 	- SQL: filters messages by using SQL expressions.
	//
	// 	- TAG: filters messages by using tags.
	//
	// example:
	//
	// SQL
	FilterExpressionType *string `json:"filterExpressionType,omitempty" xml:"filterExpressionType,omitempty"`
	// The consumption mode of the consumer group. Valid values:
	//
	// 	- BROADCASTING: broadcasting consumption
	//
	// 	- CLUSTERING: clustering consumption
	//
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The subscription status. Valid values:
	//
	// 	- ONLINE: The consumer group is online. If the consumer group contains multiple consumers, this value is returned as long as one of the consumers is online.
	//
	// 	- OFFLINE: The consumer group is offline. If the consumer group contains multiple consumers, this value is returned only if all consumers are offline.
	//
	// example:
	//
	// ONLINE
	SubscriptionStatus *string `json:"subscriptionStatus,omitempty" xml:"subscriptionStatus,omitempty"`
	// Indicates whether the topic is created.
	//
	// example:
	//
	// true
	TopicCreated *bool `json:"topicCreated,omitempty" xml:"topicCreated,omitempty"`
	// The topic to which the consumer group subscribes.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListConsumerGroupSubscriptionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupSubscriptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) GetConsistency() *bool {
	return s.Consistency
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) GetFilterExpression() *string {
	return s.FilterExpression
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) GetFilterExpressionType() *string {
	return s.FilterExpressionType
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) GetMessageModel() *string {
	return s.MessageModel
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) GetSubscriptionStatus() *string {
	return s.SubscriptionStatus
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) GetTopicCreated() *bool {
	return s.TopicCreated
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetConsistency(v bool) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.Consistency = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetConsumerGroupId(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetFilterExpression(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.FilterExpression = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetFilterExpressionType(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.FilterExpressionType = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetMessageModel(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.MessageModel = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetSubscriptionStatus(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.SubscriptionStatus = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetTopicCreated(v bool) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.TopicCreated = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetTopicName(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
