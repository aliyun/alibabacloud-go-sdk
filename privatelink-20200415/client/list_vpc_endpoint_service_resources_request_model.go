// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServiceResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVpcEndpointServiceResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointServiceResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVpcEndpointServiceResourcesRequest
	GetRegionId() *string
	SetServiceId(v string) *ListVpcEndpointServiceResourcesRequest
	GetServiceId() *string
}

type ListVpcEndpointServiceResourcesRequest struct {
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
	// 	- If a next request is to be performed, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the service resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ListVpcEndpointServiceResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServiceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointServiceResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointServiceResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointServiceResourcesRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListVpcEndpointServiceResourcesRequest) SetMaxResults(v int32) *ListVpcEndpointServiceResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesRequest) SetNextToken(v string) *ListVpcEndpointServiceResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesRequest) SetRegionId(v string) *ListVpcEndpointServiceResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesRequest) SetServiceId(v string) *ListVpcEndpointServiceResourcesRequest {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesRequest) Validate() error {
	return dara.Validate(s)
}
