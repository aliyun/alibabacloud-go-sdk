// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeCloudAssistantStatusRequest
	GetInstanceIds() []*string
	SetPageNumber(v int32) *DescribeCloudAssistantStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCloudAssistantStatusRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCloudAssistantStatusRequest
	GetRegionId() *string
}

type DescribeCloudAssistantStatusRequest struct {
	// The IDs of the simple application servers.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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

func (s DescribeCloudAssistantStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeCloudAssistantStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCloudAssistantStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudAssistantStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudAssistantStatusRequest) SetInstanceIds(v []*string) *DescribeCloudAssistantStatusRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetPageNumber(v int32) *DescribeCloudAssistantStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetPageSize(v int32) *DescribeCloudAssistantStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetRegionId(v string) *DescribeCloudAssistantStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) Validate() error {
	return dara.Validate(s)
}
