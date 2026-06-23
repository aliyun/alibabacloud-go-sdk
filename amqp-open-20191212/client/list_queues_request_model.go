// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListQueuesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListQueuesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListQueuesRequest
	GetNextToken() *string
	SetVirtualHost(v string) *ListQueuesRequest
	GetVirtualHost() *string
}

type ListQueuesRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-5yd3aw******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of results to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current query. Pass this token in the next call to retrieve the next page of results. The value is an empty string for the first call or when the last page is returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The vhost name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListQueuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueuesRequest) GoString() string {
	return s.String()
}

func (s *ListQueuesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListQueuesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQueuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQueuesRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *ListQueuesRequest) SetInstanceId(v string) *ListQueuesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQueuesRequest) SetMaxResults(v int32) *ListQueuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQueuesRequest) SetNextToken(v string) *ListQueuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueuesRequest) SetVirtualHost(v string) *ListQueuesRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListQueuesRequest) Validate() error {
	return dara.Validate(s)
}
