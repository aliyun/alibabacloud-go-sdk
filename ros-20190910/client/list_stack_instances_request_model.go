// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListStackInstancesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListStackInstancesRequest
	GetPageSize() *int64
	SetRegionId(v string) *ListStackInstancesRequest
	GetRegionId() *string
	SetStackGroupName(v string) *ListStackInstancesRequest
	GetStackGroupName() *string
	SetStackInstanceAccountId(v string) *ListStackInstancesRequest
	GetStackInstanceAccountId() *string
	SetStackInstanceRegionId(v string) *ListStackInstancesRequest
	GetStackInstanceRegionId() *string
}

type ListStackInstancesRequest struct {
	// The number of the page to return.
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// 	- Valid values: 1 to 50.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group. The name must be unique within a region.\\
	//
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The ID of the destination account to which the stack belongs.
	//
	// 	- If the stack group is granted self-managed permissions, the stack belongs to an Alibaba Cloud account.
	//
	// 	- If the stack group is granted service-managed permissions, the stack belongs to a member in a resource directory.
	//
	// > For more information about the destination account, see [Overview](https://help.aliyun.com/document_detail/154578.html).
	//
	// example:
	//
	// 156552876021****
	StackInstanceAccountId *string `json:"StackInstanceAccountId,omitempty" xml:"StackInstanceAccountId,omitempty"`
	// The region ID of the stack.
	//
	// example:
	//
	// cn-beijing
	StackInstanceRegionId *string `json:"StackInstanceRegionId,omitempty" xml:"StackInstanceRegionId,omitempty"`
}

func (s ListStackInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStackInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListStackInstancesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStackInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStackInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStackInstancesRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *ListStackInstancesRequest) GetStackInstanceAccountId() *string {
	return s.StackInstanceAccountId
}

func (s *ListStackInstancesRequest) GetStackInstanceRegionId() *string {
	return s.StackInstanceRegionId
}

func (s *ListStackInstancesRequest) SetPageNumber(v int64) *ListStackInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStackInstancesRequest) SetPageSize(v int64) *ListStackInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackInstancesRequest) SetRegionId(v string) *ListStackInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackInstancesRequest) SetStackGroupName(v string) *ListStackInstancesRequest {
	s.StackGroupName = &v
	return s
}

func (s *ListStackInstancesRequest) SetStackInstanceAccountId(v string) *ListStackInstancesRequest {
	s.StackInstanceAccountId = &v
	return s
}

func (s *ListStackInstancesRequest) SetStackInstanceRegionId(v string) *ListStackInstancesRequest {
	s.StackInstanceRegionId = &v
	return s
}

func (s *ListStackInstancesRequest) Validate() error {
	return dara.Validate(s)
}
