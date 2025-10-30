// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *ListVpcEndpointZonesRequest
	GetEndpointId() *string
	SetMaxResults(v int32) *ListVpcEndpointZonesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointZonesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVpcEndpointZonesRequest
	GetRegionId() *string
}

type ListVpcEndpointZonesRequest struct {
	// The ID of the endpoint for which you want to query zones.
	//
	// After you specify an endpoint ID, the system queries the zones of the specified endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
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
	// 	- If a next request is to be performed, set the parameter to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListVpcEndpointZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointZonesRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointZonesRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListVpcEndpointZonesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointZonesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointZonesRequest) SetEndpointId(v string) *ListVpcEndpointZonesRequest {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointZonesRequest) SetMaxResults(v int32) *ListVpcEndpointZonesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointZonesRequest) SetNextToken(v string) *ListVpcEndpointZonesRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointZonesRequest) SetRegionId(v string) *ListVpcEndpointZonesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointZonesRequest) Validate() error {
	return dara.Validate(s)
}
