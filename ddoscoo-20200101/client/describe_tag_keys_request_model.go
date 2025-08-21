// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeTagKeysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagKeysRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTagKeysRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeTagKeysRequest
	GetResourceGroupId() *string
	SetResourceType(v string) *DescribeTagKeysRequest
	GetResourceType() *string
}

type DescribeTagKeysRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance. Set the value to **cn-hangzhou**, which indicates an Anti-DDoS Proxy (Chinese Mainland) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the resource to which the tag belongs. Set the value to **INSTANCE**, which indicates an Anti-DDoS Pro instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTagKeysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTagKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagKeysRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTagKeysRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagKeysRequest) SetPageNumber(v int32) *DescribeTagKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagKeysRequest) SetPageSize(v int32) *DescribeTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagKeysRequest) SetRegionId(v string) *DescribeTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagKeysRequest) SetResourceGroupId(v string) *DescribeTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTagKeysRequest) SetResourceType(v string) *DescribeTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagKeysRequest) Validate() error {
	return dara.Validate(s)
}
