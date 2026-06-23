// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListBindingsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListBindingsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBindingsRequest
	GetNextToken() *string
	SetVirtualHost(v string) *ListBindingsRequest
	GetVirtualHost() *string
}

type ListBindingsRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return. Valid values:
	//
	// **1 to 100**
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current query. To retrieve the next page of results, pass the token that was returned in the previous call. For the first call or if the last page is returned, the value of this parameter is an empty string.
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
	// Test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListBindingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBindingsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBindingsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBindingsRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *ListBindingsRequest) SetInstanceId(v string) *ListBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBindingsRequest) SetMaxResults(v int32) *ListBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBindingsRequest) SetNextToken(v string) *ListBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBindingsRequest) SetVirtualHost(v string) *ListBindingsRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListBindingsRequest) Validate() error {
	return dara.Validate(s)
}
