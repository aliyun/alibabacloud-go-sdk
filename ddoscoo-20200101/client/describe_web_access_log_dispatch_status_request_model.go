// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessLogDispatchStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeWebAccessLogDispatchStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeWebAccessLogDispatchStatusRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeWebAccessLogDispatchStatusRequest
	GetResourceGroupId() *string
}

type DescribeWebAccessLogDispatchStatusRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
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

func (s DescribeWebAccessLogDispatchStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessLogDispatchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogDispatchStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeWebAccessLogDispatchStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebAccessLogDispatchStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWebAccessLogDispatchStatusRequest) SetPageNumber(v int32) *DescribeWebAccessLogDispatchStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusRequest) SetPageSize(v int32) *DescribeWebAccessLogDispatchStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusRequest) SetResourceGroupId(v string) *DescribeWebAccessLogDispatchStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusRequest) Validate() error {
	return dara.Validate(s)
}
