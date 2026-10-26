// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConsumerGroupResponseBody
	GetCode() *string
	SetData(v *GetConsumerGroupResponseBodyData) *GetConsumerGroupResponseBody
	GetData() *GetConsumerGroupResponseBodyData
	SetDynamicCode(v string) *GetConsumerGroupResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetConsumerGroupResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetConsumerGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetConsumerGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConsumerGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConsumerGroupResponseBody
	GetSuccess() *bool
}

type GetConsumerGroupResponseBody struct {
	// The fault code.
	//
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	Data *GetConsumerGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic fault code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic fault message.
	//
	// example:
	//
	// consumerGroupId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The fault message.
	//
	// example:
	//
	// Parameter consumerGroupId is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. Each request has a unique ID, which can be used to troubleshoot and locate faults.
	//
	// example:
	//
	// C7F94090-3358-506A-97DC-34BC803C****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the execute operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConsumerGroupResponseBody) GetData() *GetConsumerGroupResponseBodyData {
	return s.Data
}

func (s *GetConsumerGroupResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetConsumerGroupResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetConsumerGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetConsumerGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsumerGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConsumerGroupResponseBody) SetCode(v string) *GetConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetData(v *GetConsumerGroupResponseBodyData) *GetConsumerGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumerGroupResponseBody) SetDynamicCode(v string) *GetConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetDynamicMessage(v string) *GetConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetHttpStatusCode(v int32) *GetConsumerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetMessage(v string) *GetConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetRequestId(v string) *GetConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetSuccess(v bool) *GetConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetConsumerGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConsumerGroupResponseBodyData struct {
	// The consumption retry policy of the consumer group. For more information, see [Consumption retry](https://help.aliyun.com/document_detail/440356.html).
	ConsumeRetryPolicy *GetConsumerGroupResponseBodyDataConsumeRetryPolicy `json:"consumeRetryPolicy,omitempty" xml:"consumeRetryPolicy,omitempty" type:"Struct"`
	// The ID of the consumer group.
	//
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The time when the consumer group was created, in the format of YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The delivery order of the consumer group.
	//
	// example:
	//
	// Concurrently
	DeliveryOrderType *string `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	// Indicates whether exclusive consumption is enabled.
	//
	// example:
	//
	// false
	Exclusive *bool `json:"exclusive,omitempty" xml:"exclusive,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The maximum TPS for consuming messages.
	//
	// example:
	//
	// 2000
	MaxReceiveTps *int64 `json:"maxReceiveTps,omitempty" xml:"maxReceiveTps,omitempty"`
	// The consumption pattern. Valid values:
	//
	//  - CLUSTERING: clustering consumption.
	//
	//  - LITE_SELECTIVE: LITE_SELECTIVE.
	//
	// example:
	//
	// LITE_SELECTIVE
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The remarks of the consumer group.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The status of the consumer group.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The topic name for the lite type.
	//
	// example:
	//
	// test1
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// The time when the consumer group was last updated, in the format of YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetConsumerGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponseBodyData) GetConsumeRetryPolicy() *GetConsumerGroupResponseBodyDataConsumeRetryPolicy {
	return s.ConsumeRetryPolicy
}

func (s *GetConsumerGroupResponseBodyData) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *GetConsumerGroupResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetConsumerGroupResponseBodyData) GetDeliveryOrderType() *string {
	return s.DeliveryOrderType
}

func (s *GetConsumerGroupResponseBodyData) GetExclusive() *bool {
	return s.Exclusive
}

func (s *GetConsumerGroupResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConsumerGroupResponseBodyData) GetMaxReceiveTps() *int64 {
	return s.MaxReceiveTps
}

func (s *GetConsumerGroupResponseBodyData) GetMessageModel() *string {
	return s.MessageModel
}

func (s *GetConsumerGroupResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetConsumerGroupResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *GetConsumerGroupResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetConsumerGroupResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *GetConsumerGroupResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetConsumerGroupResponseBodyData) SetConsumeRetryPolicy(v *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) *GetConsumerGroupResponseBodyData {
	s.ConsumeRetryPolicy = v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetConsumerGroupId(v string) *GetConsumerGroupResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetCreateTime(v string) *GetConsumerGroupResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetDeliveryOrderType(v string) *GetConsumerGroupResponseBodyData {
	s.DeliveryOrderType = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetExclusive(v bool) *GetConsumerGroupResponseBodyData {
	s.Exclusive = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetInstanceId(v string) *GetConsumerGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetMaxReceiveTps(v int64) *GetConsumerGroupResponseBodyData {
	s.MaxReceiveTps = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetMessageModel(v string) *GetConsumerGroupResponseBodyData {
	s.MessageModel = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetRegionId(v string) *GetConsumerGroupResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetRemark(v string) *GetConsumerGroupResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetStatus(v string) *GetConsumerGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetTopicName(v string) *GetConsumerGroupResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetUpdateTime(v string) *GetConsumerGroupResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) Validate() error {
	if s.ConsumeRetryPolicy != nil {
		if err := s.ConsumeRetryPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConsumerGroupResponseBodyDataConsumeRetryPolicy struct {
	// The dead-letter topic.
	//
	// If a consumer fails to consume a message due to an exception and the message still cannot be consumed after retries, the message is delivered to the dead-letter topic for subsequent business recovery or review. For more information, see [Consumption retry and dead-letter messages](https://help.aliyun.com/document_detail/440356.html).
	//
	// example:
	//
	// DLQ_mqtest
	DeadLetterTargetTopic *string `json:"deadLetterTargetTopic,omitempty" xml:"deadLetterTargetTopic,omitempty"`
	// The fixed retry interval. Unit: seconds.
	//
	//  - Concurrently (concurrent delivery): 10-1800
	//
	//  - Orderly (ordered delivery): 1-600
	//
	// example:
	//
	// 20
	FixedIntervalRetryTime *int32 `json:"fixedIntervalRetryTime,omitempty" xml:"fixedIntervalRetryTime,omitempty"`
	// The maximum number of retry attempts.
	//
	// example:
	//
	// 16
	MaxRetryTimes *int32 `json:"maxRetryTimes,omitempty" xml:"maxRetryTimes,omitempty"`
	// The retry policy type.
	//
	// example:
	//
	// DefaultRetryPolicy
	RetryPolicy *string `json:"retryPolicy,omitempty" xml:"retryPolicy,omitempty"`
}

func (s GetConsumerGroupResponseBodyDataConsumeRetryPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupResponseBodyDataConsumeRetryPolicy) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) GetDeadLetterTargetTopic() *string {
	return s.DeadLetterTargetTopic
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) GetFixedIntervalRetryTime() *int32 {
	return s.FixedIntervalRetryTime
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) GetMaxRetryTimes() *int32 {
	return s.MaxRetryTimes
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) GetRetryPolicy() *string {
	return s.RetryPolicy
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) SetDeadLetterTargetTopic(v string) *GetConsumerGroupResponseBodyDataConsumeRetryPolicy {
	s.DeadLetterTargetTopic = &v
	return s
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) SetFixedIntervalRetryTime(v int32) *GetConsumerGroupResponseBodyDataConsumeRetryPolicy {
	s.FixedIntervalRetryTime = &v
	return s
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) SetMaxRetryTimes(v int32) *GetConsumerGroupResponseBodyDataConsumeRetryPolicy {
	s.MaxRetryTimes = &v
	return s
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) SetRetryPolicy(v string) *GetConsumerGroupResponseBodyDataConsumeRetryPolicy {
	s.RetryPolicy = &v
	return s
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) Validate() error {
	return dara.Validate(s)
}
