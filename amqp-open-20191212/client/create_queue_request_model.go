// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDeleteState(v bool) *CreateQueueRequest
	GetAutoDeleteState() *bool
	SetAutoExpireState(v int64) *CreateQueueRequest
	GetAutoExpireState() *int64
	SetDeadLetterExchange(v string) *CreateQueueRequest
	GetDeadLetterExchange() *string
	SetDeadLetterRoutingKey(v string) *CreateQueueRequest
	GetDeadLetterRoutingKey() *string
	SetExclusiveState(v bool) *CreateQueueRequest
	GetExclusiveState() *bool
	SetInstanceId(v string) *CreateQueueRequest
	GetInstanceId() *string
	SetMaxLength(v int64) *CreateQueueRequest
	GetMaxLength() *int64
	SetMaximumPriority(v int32) *CreateQueueRequest
	GetMaximumPriority() *int32
	SetMessageTTL(v int64) *CreateQueueRequest
	GetMessageTTL() *int64
	SetQueueName(v string) *CreateQueueRequest
	GetQueueName() *string
	SetVirtualHost(v string) *CreateQueueRequest
	GetVirtualHost() *string
}

type CreateQueueRequest struct {
	// Specifies whether to automatically delete the queue. Valid values:
	//
	// - true: The queue is automatically deleted after the last consumer unsubscribes from it.
	//
	// - false: The queue is not automatically deleted.
	//
	// example:
	//
	// false
	AutoDeleteState *bool `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	// The auto-expiration time for the queue. The queue is automatically deleted if it is not accessed within the specified time period.
	//
	// Unit: milliseconds.
	//
	// > This feature must be enabled before you can use this parameter. To enable the feature, <props="china">[submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex)<props="intl">[submit a ticket](https://ticket-intl.console.aliyun.com/#/ticket/createIndex).
	//
	// example:
	//
	// 10000
	AutoExpireState *int64 `json:"AutoExpireState,omitempty" xml:"AutoExpireState,omitempty"`
	// The dead-letter exchange. This type of exchange is used to receive rejected messages.
	//
	// If a consumer rejects a message and the message is not requeued, ApsaraMQ for RabbitMQ routes the message to the specified dead-letter exchange. The dead-letter exchange then routes the message to a bound queue for storage.
	//
	// example:
	//
	// DLExchange
	DeadLetterExchange *string `json:"DeadLetterExchange,omitempty" xml:"DeadLetterExchange,omitempty"`
	// The dead-letter routing key.
	//
	// The key can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@). The key must be 1 to 255 characters in length.
	//
	// example:
	//
	// test.dl
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" xml:"DeadLetterRoutingKey,omitempty"`
	// Specifies whether the queue is an exclusive queue. Valid values:
	//
	// - true: The queue is an exclusive queue. An exclusive queue can be used only by the connection that declares it. The queue is automatically deleted after the connection is closed.
	//
	// - false: The queue is not an exclusive queue.
	//
	// example:
	//
	// false
	ExclusiveState *bool `json:"ExclusiveState,omitempty" xml:"ExclusiveState,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance to which the queue belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is not supported in the current version.
	//
	// The maximum number of messages that can be stored in the queue. If this limit is exceeded, the earliest messages in the queue are deleted.
	//
	// example:
	//
	// 1000
	MaxLength *int64 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// The priority of the queue. The recommended value is an integer from 1 to 10.
	//
	// > This parameter is used for message priority. It is supported only by dedicated instances and can be used only after the message priority feature is enabled. To enable the feature, <props="china">[submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex)<props="intl">[submit a ticket](https://ticket-intl.console.aliyun.com/#/ticket/createIndex).
	//
	// example:
	//
	// 10
	MaximumPriority *int32 `json:"MaximumPriority,omitempty" xml:"MaximumPriority,omitempty"`
	// The time to live (TTL) of a message in the queue.
	//
	// - A message expires if its retention time in the queue exceeds the configured TTL.
	//
	// - The message TTL must be a non-negative integer. The maximum value is 1 day. The unit is milliseconds. For example, if the value is 1000, the message can be stored in the queue for a maximum of 1 second.
	//
	// example:
	//
	// 1000
	MessageTTL *int64 `json:"MessageTTL,omitempty" xml:"MessageTTL,omitempty"`
	// The name of the queue to create.
	//
	// - The queue name can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@). The name must be 1 to 255 characters in length.
	//
	// - After a queue is created, its name cannot be changed. To change the name, delete the queue and create a new one.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoQueue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The name of the vhost to which the queue belongs.
	//
	// The name can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@). The name must be 1 to 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s CreateQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueRequest) GetAutoDeleteState() *bool {
	return s.AutoDeleteState
}

func (s *CreateQueueRequest) GetAutoExpireState() *int64 {
	return s.AutoExpireState
}

func (s *CreateQueueRequest) GetDeadLetterExchange() *string {
	return s.DeadLetterExchange
}

func (s *CreateQueueRequest) GetDeadLetterRoutingKey() *string {
	return s.DeadLetterRoutingKey
}

func (s *CreateQueueRequest) GetExclusiveState() *bool {
	return s.ExclusiveState
}

func (s *CreateQueueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateQueueRequest) GetMaxLength() *int64 {
	return s.MaxLength
}

func (s *CreateQueueRequest) GetMaximumPriority() *int32 {
	return s.MaximumPriority
}

func (s *CreateQueueRequest) GetMessageTTL() *int64 {
	return s.MessageTTL
}

func (s *CreateQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateQueueRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *CreateQueueRequest) SetAutoDeleteState(v bool) *CreateQueueRequest {
	s.AutoDeleteState = &v
	return s
}

func (s *CreateQueueRequest) SetAutoExpireState(v int64) *CreateQueueRequest {
	s.AutoExpireState = &v
	return s
}

func (s *CreateQueueRequest) SetDeadLetterExchange(v string) *CreateQueueRequest {
	s.DeadLetterExchange = &v
	return s
}

func (s *CreateQueueRequest) SetDeadLetterRoutingKey(v string) *CreateQueueRequest {
	s.DeadLetterRoutingKey = &v
	return s
}

func (s *CreateQueueRequest) SetExclusiveState(v bool) *CreateQueueRequest {
	s.ExclusiveState = &v
	return s
}

func (s *CreateQueueRequest) SetInstanceId(v string) *CreateQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateQueueRequest) SetMaxLength(v int64) *CreateQueueRequest {
	s.MaxLength = &v
	return s
}

func (s *CreateQueueRequest) SetMaximumPriority(v int32) *CreateQueueRequest {
	s.MaximumPriority = &v
	return s
}

func (s *CreateQueueRequest) SetMessageTTL(v int64) *CreateQueueRequest {
	s.MessageTTL = &v
	return s
}

func (s *CreateQueueRequest) SetQueueName(v string) *CreateQueueRequest {
	s.QueueName = &v
	return s
}

func (s *CreateQueueRequest) SetVirtualHost(v string) *CreateQueueRequest {
	s.VirtualHost = &v
	return s
}

func (s *CreateQueueRequest) Validate() error {
	return dara.Validate(s)
}
