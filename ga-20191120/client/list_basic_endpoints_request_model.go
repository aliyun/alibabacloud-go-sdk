// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListBasicEndpointsRequest
	GetClientToken() *string
	SetEndpointGroupId(v string) *ListBasicEndpointsRequest
	GetEndpointGroupId() *string
	SetEndpointId(v string) *ListBasicEndpointsRequest
	GetEndpointId() *string
	SetEndpointType(v string) *ListBasicEndpointsRequest
	GetEndpointType() *string
	SetMaxResults(v int32) *ListBasicEndpointsRequest
	GetMaxResults() *int32
	SetName(v string) *ListBasicEndpointsRequest
	GetName() *string
	SetNextToken(v string) *ListBasicEndpointsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListBasicEndpointsRequest
	GetRegionId() *string
}

type ListBasicEndpointsRequest struct {
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
	// The ID of the endpoint group to which the endpoint that you want to query belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the endpoint that you want to query.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The type of endpoint that you want to query. Valid values:
	//
	// 	- **ENI**: elastic network interface (ENI).
	//
	// 	- **SLB**: Classic Load Balancer (CLB) instance.
	//
	// 	- **ECS**: Elastic Compute Service (ECS) instance.
	//
	// 	- **NLB**: Network Load Balancer (NLB) instance.
	//
	// example:
	//
	// ENI
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the endpoint that you want to query.
	//
	// The name must be 1 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// ep01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s ListBasicEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBasicEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListBasicEndpointsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListBasicEndpointsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListBasicEndpointsRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListBasicEndpointsRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListBasicEndpointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBasicEndpointsRequest) GetName() *string {
	return s.Name
}

func (s *ListBasicEndpointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBasicEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBasicEndpointsRequest) SetClientToken(v string) *ListBasicEndpointsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetEndpointGroupId(v string) *ListBasicEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetEndpointId(v string) *ListBasicEndpointsRequest {
	s.EndpointId = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetEndpointType(v string) *ListBasicEndpointsRequest {
	s.EndpointType = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetMaxResults(v int32) *ListBasicEndpointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetName(v string) *ListBasicEndpointsRequest {
	s.Name = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetNextToken(v string) *ListBasicEndpointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetRegionId(v string) *ListBasicEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *ListBasicEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
