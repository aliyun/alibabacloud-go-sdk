// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeTagKeysRequest
	GetCategory() *string
	SetPageNumber(v int32) *DescribeTagKeysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagKeysRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTagKeysRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeTagKeysRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *DescribeTagKeysRequest
	GetResourceId() *string
	SetResourceType(v string) *DescribeTagKeysRequest
	GetResourceType() *string
}

type DescribeTagKeysRequest struct {
	// The type of the tag. Valid values:
	//
	// 	- **Custom**: The tag is added by a user.
	//
	// 	- **System**: The tag is added by the system.
	//
	// >  By default, if the parameter is left empty, custom tags and system tags are returned.
	//
	// example:
	//
	// Custom
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1. This parameter is used together with PageSize.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of tags to return on each page if the DTS instance has multiple tags. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the DTS instance. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking instance. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID.
	//
	// example:
	//
	// dtsl5o11f9029c****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Set the value to **ALIYUN::DTS::INSTANCE**.
	//
	// example:
	//
	// ALIYUN::DTS::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysRequest) GetCategory() *string {
	return s.Category
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

func (s *DescribeTagKeysRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeTagKeysRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagKeysRequest) SetCategory(v string) *DescribeTagKeysRequest {
	s.Category = &v
	return s
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

func (s *DescribeTagKeysRequest) SetResourceId(v string) *DescribeTagKeysRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeTagKeysRequest) SetResourceType(v string) *DescribeTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagKeysRequest) Validate() error {
	return dara.Validate(s)
}
