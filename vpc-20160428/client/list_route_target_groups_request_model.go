// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteTargetGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListRouteTargetGroupsRequest
	GetClientToken() *string
	SetMaxResults(v int32) *ListRouteTargetGroupsRequest
	GetMaxResults() *int32
	SetMemberId(v string) *ListRouteTargetGroupsRequest
	GetMemberId() *string
	SetNextToken(v string) *ListRouteTargetGroupsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListRouteTargetGroupsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListRouteTargetGroupsRequest
	GetResourceGroupId() *string
	SetRouteTargetGroupIds(v []*string) *ListRouteTargetGroupsRequest
	GetRouteTargetGroupIds() []*string
	SetTag(v []*ListRouteTargetGroupsRequestTag) *ListRouteTargetGroupsRequest
	GetTag() []*ListRouteTargetGroupsRequestTag
	SetVpcId(v string) *ListRouteTargetGroupsRequest
	GetVpcId() *string
}

type ListRouteTargetGroupsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters. If you do not specify this parameter, the system uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **50**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of a route target group member instance. This parameter filters route target groups that contain the specified member instance ID.
	//
	// example:
	//
	// ep-xxxx
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. If not all results are returned in a query, a value is returned for NextToken. In the next query, set NextToken to this value.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the VPC to which the route target group belongs. You can call the DescribeRegions operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. For more information about resource groups, see What is a resource group.
	//
	// example:
	//
	// rg-acfmxazdjdhd****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of route target group IDs.
	//
	// You can specify up to 50 instance IDs in a single query.
	RouteTargetGroupIds []*string `json:"RouteTargetGroupIds,omitempty" xml:"RouteTargetGroupIds,omitempty" type:"Repeated"`
	// The list of tags.
	Tag []*ListRouteTargetGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC to which the route target group belongs.
	//
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListRouteTargetGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRouteTargetGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListRouteTargetGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListRouteTargetGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRouteTargetGroupsRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *ListRouteTargetGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRouteTargetGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRouteTargetGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListRouteTargetGroupsRequest) GetRouteTargetGroupIds() []*string {
	return s.RouteTargetGroupIds
}

func (s *ListRouteTargetGroupsRequest) GetTag() []*ListRouteTargetGroupsRequestTag {
	return s.Tag
}

func (s *ListRouteTargetGroupsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListRouteTargetGroupsRequest) SetClientToken(v string) *ListRouteTargetGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListRouteTargetGroupsRequest) SetMaxResults(v int32) *ListRouteTargetGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRouteTargetGroupsRequest) SetMemberId(v string) *ListRouteTargetGroupsRequest {
	s.MemberId = &v
	return s
}

func (s *ListRouteTargetGroupsRequest) SetNextToken(v string) *ListRouteTargetGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRouteTargetGroupsRequest) SetRegionId(v string) *ListRouteTargetGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListRouteTargetGroupsRequest) SetResourceGroupId(v string) *ListRouteTargetGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListRouteTargetGroupsRequest) SetRouteTargetGroupIds(v []*string) *ListRouteTargetGroupsRequest {
	s.RouteTargetGroupIds = v
	return s
}

func (s *ListRouteTargetGroupsRequest) SetTag(v []*ListRouteTargetGroupsRequestTag) *ListRouteTargetGroupsRequest {
	s.Tag = v
	return s
}

func (s *ListRouteTargetGroupsRequest) SetVpcId(v string) *ListRouteTargetGroupsRequest {
	s.VpcId = &v
	return s
}

func (s *ListRouteTargetGroupsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRouteTargetGroupsRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRouteTargetGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListRouteTargetGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *ListRouteTargetGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListRouteTargetGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListRouteTargetGroupsRequestTag) SetKey(v string) *ListRouteTargetGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *ListRouteTargetGroupsRequestTag) SetValue(v string) *ListRouteTargetGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *ListRouteTargetGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
