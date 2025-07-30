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
	// 	- **Custom**: The tag key is created by users.
	//
	// 	- **System**: The tag key is created by the system.
	//
	// >  By default, if the parameter is left empty, both custom tag keys and system tag keys are returned.
	//
	// example:
	//
	// Custom
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The tag key.
	//
	// >  This parameter is required.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The page number of the start page to return for a tag key that has multiple values. The valid value ranges from 1 to the maximum value of the INTEGER data type. This parameter is often used with the PageSize parameter. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of tag values to return each time for a tag key that has multiple values. Default value: 20.
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
	// >  If this parameter is left empty, the values of all tag keys of the current user are returned.
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
