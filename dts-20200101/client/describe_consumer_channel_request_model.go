// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumerChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *DescribeConsumerChannelRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *DescribeConsumerChannelRequest
	GetDtsJobId() *string
	SetPageNumber(v int32) *DescribeConsumerChannelRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeConsumerChannelRequest
	GetPageSize() *int32
	SetParentChannelId(v string) *DescribeConsumerChannelRequest
	GetParentChannelId() *string
	SetRegionId(v string) *DescribeConsumerChannelRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeConsumerChannelRequest
	GetResourceGroupId() *string
}

type DescribeConsumerChannelRequest struct {
	// The ID of the change tracking instance. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID.
	//
	// >  You must specify at least one of the **DtsInstanceId*	- and **DtsJobId*	- parameters.
	//
	// example:
	//
	// dtsboss6pn1w******
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the change tracking task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// >  You must specify at least one of the **DtsInstanceId*	- and **DtsJobId*	- parameters.
	//
	// example:
	//
	// boss6pn1w******
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than **0*	- and does not exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The parent task ID of the distributed task.
	//
	// example:
	//
	// dtsan5114c52******
	ParentChannelId *string `json:"ParentChannelId,omitempty" xml:"ParentChannelId,omitempty"`
	// The ID of the region in which the change tracking instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
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
}

func (s DescribeConsumerChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerChannelRequest) GoString() string {
	return s.String()
}

func (s *DescribeConsumerChannelRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DescribeConsumerChannelRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeConsumerChannelRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeConsumerChannelRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeConsumerChannelRequest) GetParentChannelId() *string {
	return s.ParentChannelId
}

func (s *DescribeConsumerChannelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeConsumerChannelRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeConsumerChannelRequest) SetDtsInstanceId(v string) *DescribeConsumerChannelRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeConsumerChannelRequest) SetDtsJobId(v string) *DescribeConsumerChannelRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeConsumerChannelRequest) SetPageNumber(v int32) *DescribeConsumerChannelRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeConsumerChannelRequest) SetPageSize(v int32) *DescribeConsumerChannelRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConsumerChannelRequest) SetParentChannelId(v string) *DescribeConsumerChannelRequest {
	s.ParentChannelId = &v
	return s
}

func (s *DescribeConsumerChannelRequest) SetRegionId(v string) *DescribeConsumerChannelRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeConsumerChannelRequest) SetResourceGroupId(v string) *DescribeConsumerChannelRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeConsumerChannelRequest) Validate() error {
	return dara.Validate(s)
}
