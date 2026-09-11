// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeTagValuesRequest
	GetCategory() *string
	SetKey(v string) *DescribeTagValuesRequest
	GetKey() *string
	SetPageNumber(v int32) *DescribeTagValuesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagValuesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTagValuesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeTagValuesRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *DescribeTagValuesRequest
	GetResourceId() *string
	SetResourceType(v string) *DescribeTagValuesRequest
	GetResourceType() *string
}

type DescribeTagValuesRequest struct {
	// The type of the tag key. Valid values:
	//
	// - **Custom**: a user-added tag key.
	//
	// - **System**: a system-created tag key.
	//
	// > If this parameter is left empty, all tag keys are returned by default.
	//
	// example:
	//
	// Custom
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The key of the tag.
	//
	// > This parameter is required.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The page number. Specifies the page to return when a tag key has multiple tag values. The value must be a positive integer that does not exceed the maximum value of the Integer data type. This parameter is typically used together with PageSize. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of tag values to return per page when a tag key has multiple tag values. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking instance. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID.
	//
	// > If this parameter is left empty, all tag values of the specified tag key for the current account are returned.
	//
	// example:
	//
	// dtsl5o11f9029c****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type defined by the system. Set the value to **ALIYUN::DTS::INSTANCE**.
	//
	// example:
	//
	// ALIYUN::DTS::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagValuesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagValuesRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeTagValuesRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeTagValuesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTagValuesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTagValuesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagValuesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTagValuesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeTagValuesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagValuesRequest) SetCategory(v string) *DescribeTagValuesRequest {
	s.Category = &v
	return s
}

func (s *DescribeTagValuesRequest) SetKey(v string) *DescribeTagValuesRequest {
	s.Key = &v
	return s
}

func (s *DescribeTagValuesRequest) SetPageNumber(v int32) *DescribeTagValuesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagValuesRequest) SetPageSize(v int32) *DescribeTagValuesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagValuesRequest) SetRegionId(v string) *DescribeTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagValuesRequest) SetResourceGroupId(v string) *DescribeTagValuesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTagValuesRequest) SetResourceId(v string) *DescribeTagValuesRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeTagValuesRequest) SetResourceType(v string) *DescribeTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagValuesRequest) Validate() error {
	return dara.Validate(s)
}
