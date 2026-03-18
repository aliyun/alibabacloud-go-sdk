// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenedAccessLogInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListOpenedAccessLogInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOpenedAccessLogInstancesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListOpenedAccessLogInstancesRequest
	GetResourceGroupId() *string
}

type ListOpenedAccessLogInstancesRequest struct {
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListOpenedAccessLogInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOpenedAccessLogInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOpenedAccessLogInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOpenedAccessLogInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListOpenedAccessLogInstancesRequest) SetPageNumber(v int32) *ListOpenedAccessLogInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOpenedAccessLogInstancesRequest) SetPageSize(v int32) *ListOpenedAccessLogInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOpenedAccessLogInstancesRequest) SetResourceGroupId(v string) *ListOpenedAccessLogInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListOpenedAccessLogInstancesRequest) Validate() error {
	return dara.Validate(s)
}
