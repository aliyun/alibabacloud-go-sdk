// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProvidersForNetworkAccessEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListIdentityProvidersForNetworkAccessEndpointRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListIdentityProvidersForNetworkAccessEndpointRequest
	GetMaxResults() *int64
	SetNetworkAccessEndpointId(v string) *ListIdentityProvidersForNetworkAccessEndpointRequest
	GetNetworkAccessEndpointId() *string
	SetNextToken(v string) *ListIdentityProvidersForNetworkAccessEndpointRequest
	GetNextToken() *string
}

type ListIdentityProvidersForNetworkAccessEndpointRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询时每页行数。默认值为20，最大值为100。
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 网络端点ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// nae_ms5ewjcjzed3ysaau5t5kxxxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListIdentityProvidersForNetworkAccessEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersForNetworkAccessEndpointRequest) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersForNetworkAccessEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIdentityProvidersForNetworkAccessEndpointRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListIdentityProvidersForNetworkAccessEndpointRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *ListIdentityProvidersForNetworkAccessEndpointRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIdentityProvidersForNetworkAccessEndpointRequest) SetInstanceId(v string) *ListIdentityProvidersForNetworkAccessEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointRequest) SetMaxResults(v int64) *ListIdentityProvidersForNetworkAccessEndpointRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointRequest) SetNetworkAccessEndpointId(v string) *ListIdentityProvidersForNetworkAccessEndpointRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointRequest) SetNextToken(v string) *ListIdentityProvidersForNetworkAccessEndpointRequest {
	s.NextToken = &v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointRequest) Validate() error {
	return dara.Validate(s)
}
