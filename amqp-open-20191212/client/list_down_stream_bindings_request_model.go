// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownStreamBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExchangeName(v string) *ListDownStreamBindingsRequest
	GetExchangeName() *string
	SetInstanceId(v string) *ListDownStreamBindingsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListDownStreamBindingsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDownStreamBindingsRequest
	GetNextToken() *string
	SetVirtualHost(v string) *ListDownStreamBindingsRequest
	GetVirtualHost() *string
}

type ListDownStreamBindingsRequest struct {
	// The name of the exchange.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance to which the exchange belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. If this is your first call, or if you have retrieved all results, leave this parameter empty.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the vhost to which the exchange belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListDownStreamBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDownStreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsRequest) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *ListDownStreamBindingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDownStreamBindingsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDownStreamBindingsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDownStreamBindingsRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *ListDownStreamBindingsRequest) SetExchangeName(v string) *ListDownStreamBindingsRequest {
	s.ExchangeName = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetInstanceId(v string) *ListDownStreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetMaxResults(v int32) *ListDownStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetNextToken(v string) *ListDownStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetVirtualHost(v string) *ListDownStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListDownStreamBindingsRequest) Validate() error {
	return dara.Validate(s)
}
