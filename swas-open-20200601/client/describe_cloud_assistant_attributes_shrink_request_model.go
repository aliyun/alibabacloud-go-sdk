// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantAttributesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *DescribeCloudAssistantAttributesShrinkRequest
	GetInstanceIdsShrink() *string
	SetPageNumber(v int32) *DescribeCloudAssistantAttributesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCloudAssistantAttributesShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCloudAssistantAttributesShrinkRequest
	GetRegionId() *string
}

type DescribeCloudAssistantAttributesShrinkRequest struct {
	// The instance IDs of the simple application server instances.
	//
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The page number.
	//
	// Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers.
	//
	// You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query the list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudAssistantAttributesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantAttributesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) SetInstanceIdsShrink(v string) *DescribeCloudAssistantAttributesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) SetPageNumber(v int32) *DescribeCloudAssistantAttributesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) SetPageSize(v int32) *DescribeCloudAssistantAttributesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) SetRegionId(v string) *DescribeCloudAssistantAttributesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
