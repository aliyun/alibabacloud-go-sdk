// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *DescribeCloudAssistantStatusShrinkRequest
	GetInstanceIdsShrink() *string
	SetPageNumber(v int32) *DescribeCloudAssistantStatusShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCloudAssistantStatusShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCloudAssistantStatusShrinkRequest
	GetRegionId() *string
}

type DescribeCloudAssistantStatusShrinkRequest struct {
	// The IDs of the simple application servers.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudAssistantStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DescribeCloudAssistantStatusShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCloudAssistantStatusShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudAssistantStatusShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetInstanceIdsShrink(v string) *DescribeCloudAssistantStatusShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetPageNumber(v int32) *DescribeCloudAssistantStatusShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetPageSize(v int32) *DescribeCloudAssistantStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetRegionId(v string) *DescribeCloudAssistantStatusShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudAssistantStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
