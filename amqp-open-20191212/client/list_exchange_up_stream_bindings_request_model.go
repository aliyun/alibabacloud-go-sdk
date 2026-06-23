// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeUpStreamBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExchangeName(v string) *ListExchangeUpStreamBindingsRequest
	GetExchangeName() *string
	SetInstanceId(v string) *ListExchangeUpStreamBindingsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListExchangeUpStreamBindingsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListExchangeUpStreamBindingsRequest
	GetNextToken() *string
	SetVirtualHost(v string) *ListExchangeUpStreamBindingsRequest
	GetVirtualHost() *string
}

type ListExchangeUpStreamBindingsRequest struct {
	// The exchange name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
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
	// The token that marks the end of the current query. Pass this token in the next call to query the next page. This parameter is an empty string for the first query or when the last page is returned.
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

func (s ListExchangeUpStreamBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeUpStreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsRequest) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *ListExchangeUpStreamBindingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListExchangeUpStreamBindingsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExchangeUpStreamBindingsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExchangeUpStreamBindingsRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *ListExchangeUpStreamBindingsRequest) SetExchangeName(v string) *ListExchangeUpStreamBindingsRequest {
	s.ExchangeName = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetInstanceId(v string) *ListExchangeUpStreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetMaxResults(v int32) *ListExchangeUpStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetNextToken(v string) *ListExchangeUpStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetVirtualHost(v string) *ListExchangeUpStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) Validate() error {
	return dara.Validate(s)
}
