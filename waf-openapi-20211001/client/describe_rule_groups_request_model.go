// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeRuleGroupsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeRuleGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRuleGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRuleGroupsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeRuleGroupsRequest
	GetResourceManagerResourceGroupId() *string
	SetSearchType(v string) *DescribeRuleGroupsRequest
	GetSearchType() *string
	SetSearchValue(v string) *DescribeRuleGroupsRequest
	GetSearchValue() *string
}

type DescribeRuleGroupsRequest struct {
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the results to return. The default value is **1**, which returns the first page of results.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The default value is **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The query type. Valid values:
	//
	// - **id**: Queries by rule group ID.
	//
	// - **name**: Queries by rule group name. Both Chinese and English names are supported.
	//
	// example:
	//
	// name
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	// The content to query.
	//
	// example:
	//
	// test
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
}

func (s DescribeRuleGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRuleGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRuleGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRuleGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRuleGroupsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeRuleGroupsRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *DescribeRuleGroupsRequest) GetSearchValue() *string {
	return s.SearchValue
}

func (s *DescribeRuleGroupsRequest) SetInstanceId(v string) *DescribeRuleGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetPageNumber(v int32) *DescribeRuleGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetPageSize(v int32) *DescribeRuleGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetRegionId(v string) *DescribeRuleGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleGroupsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetSearchType(v string) *DescribeRuleGroupsRequest {
	s.SearchType = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetSearchValue(v string) *DescribeRuleGroupsRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeRuleGroupsRequest) Validate() error {
	return dara.Validate(s)
}
