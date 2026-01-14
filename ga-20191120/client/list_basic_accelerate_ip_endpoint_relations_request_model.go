// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicAccelerateIpEndpointRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpId(v string) *ListBasicAccelerateIpEndpointRelationsRequest
	GetAccelerateIpId() *string
	SetAcceleratorId(v string) *ListBasicAccelerateIpEndpointRelationsRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *ListBasicAccelerateIpEndpointRelationsRequest
	GetClientToken() *string
	SetEndpointId(v string) *ListBasicAccelerateIpEndpointRelationsRequest
	GetEndpointId() *string
	SetMaxResults(v int32) *ListBasicAccelerateIpEndpointRelationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBasicAccelerateIpEndpointRelationsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListBasicAccelerateIpEndpointRelationsRequest
	GetRegionId() *string
}

type ListBasicAccelerateIpEndpointRelationsRequest struct {
	// The ID of the accelerated IP address.
	//
	// example:
	//
	// gaip-bp1****
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	// The ID of the basic GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If this is your first query or no next query is to be sent, ignore this parameter.
	//
	// 	- If a next query is to be sent, set the value to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBasicAccelerateIpEndpointRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAccelerateIpEndpointRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetAccelerateIpId(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetAcceleratorId(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetClientToken(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetEndpointId(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.EndpointId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetMaxResults(v int32) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetNextToken(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetRegionId(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) Validate() error {
	return dara.Validate(s)
}
