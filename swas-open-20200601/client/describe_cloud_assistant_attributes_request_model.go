// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeCloudAssistantAttributesRequest
	GetInstanceIds() []*string
	SetPageNumber(v int32) *DescribeCloudAssistantAttributesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCloudAssistantAttributesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCloudAssistantAttributesRequest
	GetRegionId() *string
}

type DescribeCloudAssistantAttributesRequest struct {
	// The instance IDs of the simple application server instances.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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

func (s DescribeCloudAssistantAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantAttributesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeCloudAssistantAttributesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCloudAssistantAttributesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudAssistantAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudAssistantAttributesRequest) SetInstanceIds(v []*string) *DescribeCloudAssistantAttributesRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeCloudAssistantAttributesRequest) SetPageNumber(v int32) *DescribeCloudAssistantAttributesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantAttributesRequest) SetPageSize(v int32) *DescribeCloudAssistantAttributesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantAttributesRequest) SetRegionId(v string) *DescribeCloudAssistantAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudAssistantAttributesRequest) Validate() error {
	return dara.Validate(s)
}
