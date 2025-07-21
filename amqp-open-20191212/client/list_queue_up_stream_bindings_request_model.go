// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueUpStreamBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListQueueUpStreamBindingsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListQueueUpStreamBindingsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListQueueUpStreamBindingsRequest
	GetNextToken() *string
	SetQueueName(v string) *ListQueueUpStreamBindingsRequest
	GetQueueName() *string
	SetVirtualHost(v string) *ListQueueUpStreamBindingsRequest
	GetVirtualHost() *string
}

type ListQueueUpStreamBindingsRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The queue name.
	//
	// This parameter is required.
	//
	// example:
	//
	// QueueTest
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The virtual host (vhost) name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListQueueUpStreamBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueueUpStreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListQueueUpStreamBindingsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQueueUpStreamBindingsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQueueUpStreamBindingsRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *ListQueueUpStreamBindingsRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *ListQueueUpStreamBindingsRequest) SetInstanceId(v string) *ListQueueUpStreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetMaxResults(v int32) *ListQueueUpStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetNextToken(v string) *ListQueueUpStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetQueueName(v string) *ListQueueUpStreamBindingsRequest {
	s.QueueName = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetVirtualHost(v string) *ListQueueUpStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) Validate() error {
	return dara.Validate(s)
}
