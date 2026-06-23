// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueConsumersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListQueueConsumersRequest
	GetInstanceId() *string
	SetNextToken(v string) *ListQueueConsumersRequest
	GetNextToken() *string
	SetQueryCount(v int32) *ListQueueConsumersRequest
	GetQueryCount() *int32
	SetQueue(v string) *ListQueueConsumersRequest
	GetQueue() *string
	SetVirtualHost(v string) *ListQueueConsumersRequest
	GetVirtualHost() *string
}

type ListQueueConsumersRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-5yd3aw******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The token that marks the end of the current page of results. To retrieve the next page, include this token in the next request. If this is your first request or the last page is returned, the value is an empty string.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to return. If you do not set this parameter, the default value is 1.
	//
	// Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	QueryCount *int32 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The queue name.
	//
	// This parameter is required.
	//
	// example:
	//
	// queue-rabbit-springboot-advance5
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// The vhost name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListQueueConsumersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueueConsumersRequest) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListQueueConsumersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQueueConsumersRequest) GetQueryCount() *int32 {
	return s.QueryCount
}

func (s *ListQueueConsumersRequest) GetQueue() *string {
	return s.Queue
}

func (s *ListQueueConsumersRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *ListQueueConsumersRequest) SetInstanceId(v string) *ListQueueConsumersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQueueConsumersRequest) SetNextToken(v string) *ListQueueConsumersRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueueConsumersRequest) SetQueryCount(v int32) *ListQueueConsumersRequest {
	s.QueryCount = &v
	return s
}

func (s *ListQueueConsumersRequest) SetQueue(v string) *ListQueueConsumersRequest {
	s.Queue = &v
	return s
}

func (s *ListQueueConsumersRequest) SetVirtualHost(v string) *ListQueueConsumersRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListQueueConsumersRequest) Validate() error {
	return dara.Validate(s)
}
