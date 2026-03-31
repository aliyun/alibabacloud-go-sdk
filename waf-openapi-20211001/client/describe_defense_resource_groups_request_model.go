// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupNameLike(v string) *DescribeDefenseResourceGroupsRequest
	GetGroupNameLike() *string
	SetGroupNames(v string) *DescribeDefenseResourceGroupsRequest
	GetGroupNames() *string
	SetInstanceId(v string) *DescribeDefenseResourceGroupsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeDefenseResourceGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDefenseResourceGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDefenseResourceGroupsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceGroupsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeDefenseResourceGroupsRequest struct {
	// The name of the protected object group that you want to query. Fuzzy queries are supported.
	//
	// example:
	//
	// demoGroupName
	GroupNameLike *string `json:"GroupNameLike,omitempty" xml:"GroupNameLike,omitempty"`
	// The names of the protected object groups that you want to query. Separate multiple names with commas (,).
	//
	// example:
	//
	// groupName1,groupName2
	GroupNames *string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-wwo36****0i
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfmxc7lf****eq
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefenseResourceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupsRequest) GetGroupNameLike() *string {
	return s.GroupNameLike
}

func (s *DescribeDefenseResourceGroupsRequest) GetGroupNames() *string {
	return s.GroupNames
}

func (s *DescribeDefenseResourceGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseResourceGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDefenseResourceGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDefenseResourceGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseResourceGroupsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseResourceGroupsRequest) SetGroupNameLike(v string) *DescribeDefenseResourceGroupsRequest {
	s.GroupNameLike = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetGroupNames(v string) *DescribeDefenseResourceGroupsRequest {
	s.GroupNames = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetInstanceId(v string) *DescribeDefenseResourceGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetPageNumber(v int32) *DescribeDefenseResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetPageSize(v int32) *DescribeDefenseResourceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetRegionId(v string) *DescribeDefenseResourceGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceGroupsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
