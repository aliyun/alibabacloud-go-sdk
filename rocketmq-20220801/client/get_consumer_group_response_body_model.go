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
	// The error code.
	//
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetConsumerGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
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
	// The error message.
	//
	// example:
	//
	// Parameter consumerGroupId is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// C7F94090-3358-506A-97DC-34BC803C****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
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
	return dara.Validate(s)
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
	// The time when the consumer group was created.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The message delivery method of the consumer group.
	//
	// Valid values:
	//
	// 	- Concurrently: concurrent delivery
	//
	// 	- Orderly: ordered delivery
	//
	// example:
	//
	// Concurrently
	DeliveryOrderType *string `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Maximum received message tps
	//
	// example:
	//
	// 1000
	MaxReceiveTps *int64 `json:"maxReceiveTps,omitempty" xml:"maxReceiveTps,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The remarks on the consumer group.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The status of the consumer group.
	//
	// Valid values:
	//
	// 	- RUNNING
	//
	// 	- CREATING
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the consumer group was last updated.
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

func (s *GetConsumerGroupResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConsumerGroupResponseBodyData) GetMaxReceiveTps() *int64 {
	return s.MaxReceiveTps
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

func (s *GetConsumerGroupResponseBodyData) SetInstanceId(v string) *GetConsumerGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetMaxReceiveTps(v int64) *GetConsumerGroupResponseBodyData {
	s.MaxReceiveTps = &v
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

func (s *GetConsumerGroupResponseBodyData) SetUpdateTime(v string) *GetConsumerGroupResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetConsumerGroupResponseBodyDataConsumeRetryPolicy struct {
	// The dead-letter topic.
	//
	// If a consumer still fails to consume a message after the message is retried for a specified number of times, the message is delivered to a dead-letter topic for subsequent business recovery or troubleshooting. For more information, see [Consumption retry and dead-letter messages](https://help.aliyun.com/document_detail/440356.html).
	//
	// example:
	//
	// DLQ_mqtest
	DeadLetterTargetTopic *string `json:"deadLetterTargetTopic,omitempty" xml:"deadLetterTargetTopic,omitempty"`
	// Fixed interval retry time,Value range, unit: seconds
	//
	//  - Concurrently:10-1800
	//
	//  - Orderly:1-600
	//
	// example:
	//
	// 20
	FixedIntervalRetryTime *int32 `json:"fixedIntervalRetryTime,omitempty" xml:"fixedIntervalRetryTime,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 16
	MaxRetryTimes *int32 `json:"maxRetryTimes,omitempty" xml:"maxRetryTimes,omitempty"`
	// The retry policy.
	//
	// Valid values:
	//
	// 	- FixedRetryPolicy: fixed-interval retry
	//
	// 	- DefaultRetryPolicy: exponential backoff retry
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
