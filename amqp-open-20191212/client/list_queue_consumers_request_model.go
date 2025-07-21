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
	// 188077086902***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of data entries to return. If you do not configure this parameter, the default value 1 is used.
	//
	// Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	QueryCount *int32 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The name of the queue for which you want to query online consumers.
	//
	// This parameter is required.
	//
	// example:
	//
	// queue-rabbit-springboot-advance5
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// The virtual host (vhost) name.
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
