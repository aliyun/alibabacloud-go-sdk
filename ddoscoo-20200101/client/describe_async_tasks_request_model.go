// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAsyncTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeAsyncTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAsyncTasksRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeAsyncTasksRequest
	GetResourceGroupId() *string
}

type DescribeAsyncTasksRequest struct {
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeAsyncTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAsyncTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAsyncTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAsyncTasksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAsyncTasksRequest) SetPageNumber(v int32) *DescribeAsyncTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAsyncTasksRequest) SetPageSize(v int32) *DescribeAsyncTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAsyncTasksRequest) SetResourceGroupId(v string) *DescribeAsyncTasksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAsyncTasksRequest) Validate() error {
	return dara.Validate(s)
}
