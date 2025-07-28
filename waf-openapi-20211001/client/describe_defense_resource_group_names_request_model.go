// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceGroupNamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupNameLike(v string) *DescribeDefenseResourceGroupNamesRequest
	GetGroupNameLike() *string
	SetInstanceId(v string) *DescribeDefenseResourceGroupNamesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeDefenseResourceGroupNamesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDefenseResourceGroupNamesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDefenseResourceGroupNamesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceGroupNamesRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeDefenseResourceGroupNamesRequest struct {
	// The name of the protected object group. Fuzzy queries are supported.
	//
	// example:
	//
	// example-group
	GroupNameLike *string `json:"GroupNameLike,omitempty" xml:"GroupNameLike,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-lbj****cc03
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
	// 10
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
	// rg-aekzwwk****cv5i
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefenseResourceGroupNamesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceGroupNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupNamesRequest) GetGroupNameLike() *string {
	return s.GroupNameLike
}

func (s *DescribeDefenseResourceGroupNamesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseResourceGroupNamesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDefenseResourceGroupNamesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDefenseResourceGroupNamesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseResourceGroupNamesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetGroupNameLike(v string) *DescribeDefenseResourceGroupNamesRequest {
	s.GroupNameLike = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetInstanceId(v string) *DescribeDefenseResourceGroupNamesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetPageNumber(v int32) *DescribeDefenseResourceGroupNamesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetPageSize(v int32) *DescribeDefenseResourceGroupNamesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetRegionId(v string) *DescribeDefenseResourceGroupNamesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceGroupNamesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesRequest) Validate() error {
	return dara.Validate(s)
}
