// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventLogsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *DescribeSecurityEventLogsShrinkRequest
	GetFilterShrink() *string
	SetInstanceId(v string) *DescribeSecurityEventLogsShrinkRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeSecurityEventLogsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSecurityEventLogsShrinkRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeSecurityEventLogsShrinkRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventLogsShrinkRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeSecurityEventLogsShrinkRequest struct {
	// The query filter conditions. Multiple filter conditions are connected by the logical AND operator.
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number to return during paging in a paged query. Default value: **1**, which indicates the first page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paging in a paged query. Maximum value: **100**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeSecurityEventLogsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventLogsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventLogsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *DescribeSecurityEventLogsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityEventLogsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSecurityEventLogsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSecurityEventLogsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityEventLogsShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSecurityEventLogsShrinkRequest) SetFilterShrink(v string) *DescribeSecurityEventLogsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *DescribeSecurityEventLogsShrinkRequest) SetInstanceId(v string) *DescribeSecurityEventLogsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityEventLogsShrinkRequest) SetPageNumber(v int64) *DescribeSecurityEventLogsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecurityEventLogsShrinkRequest) SetPageSize(v int64) *DescribeSecurityEventLogsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityEventLogsShrinkRequest) SetRegionId(v string) *DescribeSecurityEventLogsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventLogsShrinkRequest) SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventLogsShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSecurityEventLogsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
