// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceNamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDefenseResourceNamesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeDefenseResourceNamesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDefenseResourceNamesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDefenseResourceNamesRequest
	GetRegionId() *string
	SetResource(v string) *DescribeDefenseResourceNamesRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceNamesRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeDefenseResourceNamesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-9lb****5s03
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
	// The name of the protected object that you want to query.
	//
	// example:
	//
	// example.xxxxaliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aekzd4c****pdwy
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefenseResourceNamesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceNamesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseResourceNamesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDefenseResourceNamesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDefenseResourceNamesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseResourceNamesRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeDefenseResourceNamesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseResourceNamesRequest) SetInstanceId(v string) *DescribeDefenseResourceNamesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceNamesRequest) SetPageNumber(v int32) *DescribeDefenseResourceNamesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseResourceNamesRequest) SetPageSize(v int32) *DescribeDefenseResourceNamesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseResourceNamesRequest) SetRegionId(v string) *DescribeDefenseResourceNamesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceNamesRequest) SetResource(v string) *DescribeDefenseResourceNamesRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseResourceNamesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceNamesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourceNamesRequest) Validate() error {
	return dara.Validate(s)
}
