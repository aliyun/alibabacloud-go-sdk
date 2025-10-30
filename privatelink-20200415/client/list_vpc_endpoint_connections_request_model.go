// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStatus(v string) *ListVpcEndpointConnectionsRequest
	GetConnectionStatus() *string
	SetEndpointId(v string) *ListVpcEndpointConnectionsRequest
	GetEndpointId() *string
	SetEndpointOwnerId(v int64) *ListVpcEndpointConnectionsRequest
	GetEndpointOwnerId() *int64
	SetEniId(v string) *ListVpcEndpointConnectionsRequest
	GetEniId() *string
	SetMaxResults(v int32) *ListVpcEndpointConnectionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointConnectionsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVpcEndpointConnectionsRequest
	GetRegionId() *string
	SetReplacedResourceId(v string) *ListVpcEndpointConnectionsRequest
	GetReplacedResourceId() *string
	SetResourceGroupId(v string) *ListVpcEndpointConnectionsRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ListVpcEndpointConnectionsRequest
	GetResourceId() *string
	SetServiceId(v string) *ListVpcEndpointConnectionsRequest
	GetServiceId() *string
}

type ListVpcEndpointConnectionsRequest struct {
	// The state of the endpoint connection. Valid values:
	//
	// 	- **Pending**: The endpoint connection is being modified.
	//
	// 	- **Connecting**: The endpoint connection is being established.
	//
	// 	- **Connected**: The endpoint connection is established.
	//
	// 	- **Disconnecting**: The endpoint is being disconnected from the endpoint service.
	//
	// 	- **Disconnected**: The endpoint is disconnected from the endpoint service.
	//
	// 	- **Deleting**: The connection is being deleted.
	//
	// 	- **ServiceDeleted**: The corresponding endpoint service has been deleted.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the Alibaba Cloud account to which the endpoint belongs.
	//
	// example:
	//
	// 25346073170691****
	EndpointOwnerId *int64 `json:"EndpointOwnerId,omitempty" xml:"EndpointOwnerId,omitempty"`
	// The ID of the endpoint elastic network interface (ENI).
	//
	// example:
	//
	// eni-hp32lk0pyv6o94hs****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the value to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replaced service resource in smooth migration scenarios.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ReplacedResourceId *string `json:"ReplacedResourceId,omitempty" xml:"ReplacedResourceId,omitempty"`
	// The ID of the resource group to which the endpoint belongs.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service resource ID.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The endpoint service ID.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ListVpcEndpointConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointConnectionsRequest) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *ListVpcEndpointConnectionsRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListVpcEndpointConnectionsRequest) GetEndpointOwnerId() *int64 {
	return s.EndpointOwnerId
}

func (s *ListVpcEndpointConnectionsRequest) GetEniId() *string {
	return s.EniId
}

func (s *ListVpcEndpointConnectionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointConnectionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointConnectionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointConnectionsRequest) GetReplacedResourceId() *string {
	return s.ReplacedResourceId
}

func (s *ListVpcEndpointConnectionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcEndpointConnectionsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListVpcEndpointConnectionsRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListVpcEndpointConnectionsRequest) SetConnectionStatus(v string) *ListVpcEndpointConnectionsRequest {
	s.ConnectionStatus = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetEndpointId(v string) *ListVpcEndpointConnectionsRequest {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetEndpointOwnerId(v int64) *ListVpcEndpointConnectionsRequest {
	s.EndpointOwnerId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetEniId(v string) *ListVpcEndpointConnectionsRequest {
	s.EniId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetMaxResults(v int32) *ListVpcEndpointConnectionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetNextToken(v string) *ListVpcEndpointConnectionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetRegionId(v string) *ListVpcEndpointConnectionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetReplacedResourceId(v string) *ListVpcEndpointConnectionsRequest {
	s.ReplacedResourceId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetResourceGroupId(v string) *ListVpcEndpointConnectionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetResourceId(v string) *ListVpcEndpointConnectionsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetServiceId(v string) *ListVpcEndpointConnectionsRequest {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
