// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamResourceCidrsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpamPoolId(v string) *ListIpamResourceCidrsRequest
	GetIpamPoolId() *string
	SetIpamScopeId(v string) *ListIpamResourceCidrsRequest
	GetIpamScopeId() *string
	SetMaxResults(v int32) *ListIpamResourceCidrsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamResourceCidrsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListIpamResourceCidrsRequest
	GetRegionId() *string
	SetResourceId(v string) *ListIpamResourceCidrsRequest
	GetResourceId() *string
	SetResourceOwnerId(v int64) *ListIpamResourceCidrsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *ListIpamResourceCidrsRequest
	GetResourceType() *string
	SetVpcId(v string) *ListIpamResourceCidrsRequest
	GetVpcId() *string
}

type ListIpamResourceCidrsRequest struct {
	// The ID of the IPAM pool.
	//
	// >  You must specify at least one of **IpamScopeId*	- and **IpamPoolId**.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The ID of the IPAM scope.
	//
	// >  You must specify at least one of **IpamScopeId*	- and **IpamPoolId**.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the IPAM instance is hosted.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of resource. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **VSwitch**
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1fjfnrg3av6zb9e****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListIpamResourceCidrsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceCidrsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsRequest) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *ListIpamResourceCidrsRequest) GetIpamScopeId() *string {
	return s.IpamScopeId
}

func (s *ListIpamResourceCidrsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamResourceCidrsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamResourceCidrsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamResourceCidrsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListIpamResourceCidrsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListIpamResourceCidrsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListIpamResourceCidrsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListIpamResourceCidrsRequest) SetIpamPoolId(v string) *ListIpamResourceCidrsRequest {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetIpamScopeId(v string) *ListIpamResourceCidrsRequest {
	s.IpamScopeId = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetMaxResults(v int32) *ListIpamResourceCidrsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetNextToken(v string) *ListIpamResourceCidrsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetRegionId(v string) *ListIpamResourceCidrsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetResourceId(v string) *ListIpamResourceCidrsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetResourceOwnerId(v int64) *ListIpamResourceCidrsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetResourceType(v string) *ListIpamResourceCidrsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetVpcId(v string) *ListIpamResourceCidrsRequest {
	s.VpcId = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) Validate() error {
	return dara.Validate(s)
}
