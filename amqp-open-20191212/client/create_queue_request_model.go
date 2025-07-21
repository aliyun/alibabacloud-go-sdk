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
	// 	- true: The queue is automatically deleted. After the last consumer unsubscribes from the queue, the queue is automatically deleted.
	//
	// 	- false: The queue is not automatically deleted.
	//
	// example:
	//
	// false
	AutoDeleteState *bool `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	// The validity period after which the queue is automatically deleted. If the queue is not accessed within the specified period of time, the queue is automatically deleted.
	//
	// Unit: milliseconds.
	//
	// >  You can use the feature that corresponds to this parameter only after you enable the feature. To enable the feature, [submit a ticket](https://ticket-intl.console.aliyun.com/#/ticket/createIndex).
	//
	// example:
	//
	// 10000
	AutoExpireState *int64 `json:"AutoExpireState,omitempty" xml:"AutoExpireState,omitempty"`
	// The dead-letter exchange. A dead-letter exchange is used to receive rejected messages.
	//
	// If a consumer rejects a message that cannot be redelivered, ApsaraMQ for RabbitMQ routes the message to the specified dead-letter exchange. Then, the dead-letter exchange routes the message to the queue that is bound to the dead-letter exchange for storage.
	//
	// example:
	//
	// DLExchange
	DeadLetterExchange *string `json:"DeadLetterExchange,omitempty" xml:"DeadLetterExchange,omitempty"`
	// The dead-letter routing key. The key must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	// example:
	//
	// test.dl
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" xml:"DeadLetterRoutingKey,omitempty"`
	// Specifies whether the exchange is an exclusive exchange. Valid values:
	//
	// 	- true: The exchange is an exclusive exchange. Only the connection that declares the exclusive exchange can use the exclusive exchange. After the connection is closed, the exclusive exchange is automatically deleted.
	//
	// 	- false: The exchange is not an exclusive exchange.
	//
	// example:
	//
	// false
	ExclusiveState *bool `json:"ExclusiveState,omitempty" xml:"ExclusiveState,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance on which you want to create a queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is unavailable in the current version of ApsaraMQ for RabbitMQ.
	//
	// The maximum number of messages that can be stored in the queue. If this threshold is exceeded, the earliest stored messages in the queue are deleted.
	//
	// example:
	//
	// 1000
	MaxLength *int64 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// Queue priorities are not supported. The value does not affect the call or return results.
	//
	// example:
	//
	// 10
	MaximumPriority *int32 `json:"MaximumPriority,omitempty" xml:"MaximumPriority,omitempty"`
	// The message time to live (TTL) of the queue.
	//
	// 	- If the retention period of a message in the queue exceeds the message TTL of the queue, the message expires.
	//
	// 	- The message TTL must be set to a non-negative integer. The maximum message TTL is one day. Unit: milliseconds. For example, if the message TTL is 1,000 milliseconds, the message can be retained for up to 1 second in the queue.
	//
	// example:
	//
	// 1000
	MessageTTL *int64 `json:"MessageTTL,omitempty" xml:"MessageTTL,omitempty"`
	// The name of the queue that you want to create.
	//
	// 	- The name must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	// 	- After the queue is created, you cannot change the name of the queue. If you want to change the name of the queue, delete the queue and create another queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoQueue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The name of the vhost to which the queue that you want to create belongs. The name must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
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
