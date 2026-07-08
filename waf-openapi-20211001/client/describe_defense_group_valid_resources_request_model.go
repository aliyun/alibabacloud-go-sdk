// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseGroupValidResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DescribeDefenseGroupValidResourcesRequest
	GetGroupName() *string
	SetInstanceId(v string) *DescribeDefenseGroupValidResourcesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeDefenseGroupValidResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDefenseGroupValidResourcesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDefenseGroupValidResourcesRequest
	GetRegionId() *string
	SetResource(v string) *DescribeDefenseGroupValidResourcesRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseGroupValidResourcesRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeDefenseGroupValidResourcesRequest struct {
	// The name of the defense group to query.
	//
	// > Leave this parameter empty to query for protected objects that can be associated with a new defense group.
	//
	// example:
	//
	// group221
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-zxu****jc01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object that you want to use as a filter for the query results.
	//
	// example:
	//
	// example.aliyun-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek2lrm****6pnq
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefenseGroupValidResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseGroupValidResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseGroupValidResourcesRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDefenseGroupValidResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseGroupValidResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDefenseGroupValidResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDefenseGroupValidResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseGroupValidResourcesRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeDefenseGroupValidResourcesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseGroupValidResourcesRequest) SetGroupName(v string) *DescribeDefenseGroupValidResourcesRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDefenseGroupValidResourcesRequest) SetInstanceId(v string) *DescribeDefenseGroupValidResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseGroupValidResourcesRequest) SetPageNumber(v int32) *DescribeDefenseGroupValidResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseGroupValidResourcesRequest) SetPageSize(v int32) *DescribeDefenseGroupValidResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseGroupValidResourcesRequest) SetRegionId(v string) *DescribeDefenseGroupValidResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseGroupValidResourcesRequest) SetResource(v string) *DescribeDefenseGroupValidResourcesRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseGroupValidResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseGroupValidResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseGroupValidResourcesRequest) Validate() error {
	return dara.Validate(s)
}
