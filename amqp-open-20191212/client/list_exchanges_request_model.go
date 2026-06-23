// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListExchangesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListExchangesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListExchangesRequest
	GetNextToken() *string
	SetVirtualHost(v string) *ListExchangesRequest
	GetVirtualHost() *string
}

type ListExchangesRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-7pp2mwbc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of results to return. Valid values: **1 to 100**.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for retrieving the next page of results. Set this parameter to the token value returned from the previous call.
	//
	// - Omit this parameter on your first call.
	//
	// - If a subsequent call is required, set this parameter to the `NextToken` value returned from the previous call.
	//
	// example:
	//
	// AAAANDQBYW1xcC1jbi03cHAybXdiY3AwMGEBdmhvc3QBAXNkZndhYWJhATE2NDkzMTM4OTU5NDIB4o3z1pPwWzk4aYuiRffi8R6-****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the vhost.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListExchangesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExchangesRequest) GoString() string {
	return s.String()
}

func (s *ListExchangesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListExchangesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExchangesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExchangesRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *ListExchangesRequest) SetInstanceId(v string) *ListExchangesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExchangesRequest) SetMaxResults(v int32) *ListExchangesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExchangesRequest) SetNextToken(v string) *ListExchangesRequest {
	s.NextToken = &v
	return s
}

func (s *ListExchangesRequest) SetVirtualHost(v string) *ListExchangesRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListExchangesRequest) Validate() error {
	return dara.Validate(s)
}
