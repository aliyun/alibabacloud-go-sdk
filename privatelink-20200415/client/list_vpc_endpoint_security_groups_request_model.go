// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *ListVpcEndpointSecurityGroupsRequest
	GetEndpointId() *string
	SetMaxResults(v int32) *ListVpcEndpointSecurityGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointSecurityGroupsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVpcEndpointSecurityGroupsRequest
	GetRegionId() *string
}

type ListVpcEndpointSecurityGroupsRequest struct {
	// The ID of the endpoint that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The number of entries to return on each page. Valid values:**1*	- to **50**. Default value: **50**.
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
	// The region ID of the endpoint that you want to query.
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

func (s ListVpcEndpointSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointSecurityGroupsRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListVpcEndpointSecurityGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointSecurityGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointSecurityGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointSecurityGroupsRequest) SetEndpointId(v string) *ListVpcEndpointSecurityGroupsRequest {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsRequest) SetMaxResults(v int32) *ListVpcEndpointSecurityGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsRequest) SetNextToken(v string) *ListVpcEndpointSecurityGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsRequest) SetRegionId(v string) *ListVpcEndpointSecurityGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}
