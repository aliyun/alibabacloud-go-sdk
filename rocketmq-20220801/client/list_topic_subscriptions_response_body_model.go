// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicSubscriptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTopicSubscriptionsResponseBody
	GetCode() *string
	SetData(v []*ListTopicSubscriptionsResponseBodyData) *ListTopicSubscriptionsResponseBody
	GetData() []*ListTopicSubscriptionsResponseBodyData
	SetDynamicCode(v string) *ListTopicSubscriptionsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListTopicSubscriptionsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListTopicSubscriptionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTopicSubscriptionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTopicSubscriptionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTopicSubscriptionsResponseBody
	GetSuccess() *bool
}

type ListTopicSubscriptionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data []*ListTopicSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// Topic
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
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
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 92A9BE4E-B794-50C8-979C-0456E4D32943
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTopicSubscriptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTopicSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicSubscriptionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTopicSubscriptionsResponseBody) GetData() []*ListTopicSubscriptionsResponseBodyData {
	return s.Data
}

func (s *ListTopicSubscriptionsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListTopicSubscriptionsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListTopicSubscriptionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTopicSubscriptionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTopicSubscriptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTopicSubscriptionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTopicSubscriptionsResponseBody) SetCode(v string) *ListTopicSubscriptionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetData(v []*ListTopicSubscriptionsResponseBodyData) *ListTopicSubscriptionsResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetDynamicCode(v string) *ListTopicSubscriptionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetDynamicMessage(v string) *ListTopicSubscriptionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetHttpStatusCode(v int32) *ListTopicSubscriptionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetMessage(v string) *ListTopicSubscriptionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetRequestId(v string) *ListTopicSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetSuccess(v bool) *ListTopicSubscriptionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTopicSubscriptionsResponseBodyData struct {
	// Indicates whether message consumption is consistent. Valid values:
	//
	// 	- false: Unconsumed messages exist in the consumer group.
	//
	// 	- true: No unconsumed message exists in the consumer group.
	//
	// example:
	//
	// true
	Consistency *string `json:"consistency,omitempty" xml:"consistency,omitempty"`
	// The consumer group ID.
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
	// The type of the filter expression. Valid values: SQL, TAG, and UNSPECIFIED. The value SQL indicates that messages are filtered by using SQL expressions. The value TAG indicates that messages are filtered by using tags. The value UNSPECIFIED indicates that no filter expression type is specified.
	//
	// example:
	//
	// SQL
	FilterExpressionType *string `json:"filterExpressionType,omitempty" xml:"filterExpressionType,omitempty"`
	// The consumption mode. Valid values: BROADCASTING and CLUSTERING.
	//
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The subscription status. Valid values: ONLINE and OFFLINE.
	//
	// example:
	//
	// ONLINE
	SubscriptionStatus *string `json:"subscriptionStatus,omitempty" xml:"subscriptionStatus,omitempty"`
	// The topic name.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListTopicSubscriptionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTopicSubscriptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTopicSubscriptionsResponseBodyData) GetConsistency() *string {
	return s.Consistency
}

func (s *ListTopicSubscriptionsResponseBodyData) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ListTopicSubscriptionsResponseBodyData) GetFilterExpression() *string {
	return s.FilterExpression
}

func (s *ListTopicSubscriptionsResponseBodyData) GetFilterExpressionType() *string {
	return s.FilterExpressionType
}

func (s *ListTopicSubscriptionsResponseBodyData) GetMessageModel() *string {
	return s.MessageModel
}

func (s *ListTopicSubscriptionsResponseBodyData) GetSubscriptionStatus() *string {
	return s.SubscriptionStatus
}

func (s *ListTopicSubscriptionsResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTopicSubscriptionsResponseBodyData) SetConsistency(v string) *ListTopicSubscriptionsResponseBodyData {
	s.Consistency = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetConsumerGroupId(v string) *ListTopicSubscriptionsResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetFilterExpression(v string) *ListTopicSubscriptionsResponseBodyData {
	s.FilterExpression = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetFilterExpressionType(v string) *ListTopicSubscriptionsResponseBodyData {
	s.FilterExpressionType = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetMessageModel(v string) *ListTopicSubscriptionsResponseBodyData {
	s.MessageModel = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetSubscriptionStatus(v string) *ListTopicSubscriptionsResponseBodyData {
	s.SubscriptionStatus = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetTopicName(v string) *ListTopicSubscriptionsResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
