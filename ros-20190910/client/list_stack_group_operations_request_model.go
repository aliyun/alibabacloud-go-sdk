// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackGroupOperationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListStackGroupOperationsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListStackGroupOperationsRequest
	GetPageSize() *int64
	SetRegionId(v string) *ListStackGroupOperationsRequest
	GetRegionId() *string
	SetStackGroupName(v string) *ListStackGroupOperationsRequest
	GetStackGroupName() *string
}

type ListStackGroupOperationsRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group. The name must be unique within a region.
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s ListStackGroupOperationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupOperationsRequest) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStackGroupOperationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStackGroupOperationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStackGroupOperationsRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *ListStackGroupOperationsRequest) SetPageNumber(v int64) *ListStackGroupOperationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupOperationsRequest) SetPageSize(v int64) *ListStackGroupOperationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationsRequest) SetRegionId(v string) *ListStackGroupOperationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackGroupOperationsRequest) SetStackGroupName(v string) *ListStackGroupOperationsRequest {
	s.StackGroupName = &v
	return s
}

func (s *ListStackGroupOperationsRequest) Validate() error {
	return dara.Validate(s)
}
