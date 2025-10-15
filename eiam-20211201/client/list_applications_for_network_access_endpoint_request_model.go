// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForNetworkAccessEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListApplicationsForNetworkAccessEndpointRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListApplicationsForNetworkAccessEndpointRequest
	GetMaxResults() *int64
	SetNetworkAccessEndpointId(v string) *ListApplicationsForNetworkAccessEndpointRequest
	GetNetworkAccessEndpointId() *string
	SetNextToken(v string) *ListApplicationsForNetworkAccessEndpointRequest
	GetNextToken() *string
}

type ListApplicationsForNetworkAccessEndpointRequest struct {
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
	// nae_mxzj4c44ctyectkggtg4mxxxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListApplicationsForNetworkAccessEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForNetworkAccessEndpointRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForNetworkAccessEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsForNetworkAccessEndpointRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListApplicationsForNetworkAccessEndpointRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *ListApplicationsForNetworkAccessEndpointRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsForNetworkAccessEndpointRequest) SetInstanceId(v string) *ListApplicationsForNetworkAccessEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointRequest) SetMaxResults(v int64) *ListApplicationsForNetworkAccessEndpointRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointRequest) SetNetworkAccessEndpointId(v string) *ListApplicationsForNetworkAccessEndpointRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointRequest) SetNextToken(v string) *ListApplicationsForNetworkAccessEndpointRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointRequest) Validate() error {
	return dara.Validate(s)
}
