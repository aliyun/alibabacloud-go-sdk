// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNeedContainResourceGroupWithOfficeSite(v int64) *DescribeResourceGroupsRequest
	GetNeedContainResourceGroupWithOfficeSite() *int64
	SetPageNumber(v int32) *DescribeResourceGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeResourceGroupsRequest
	GetPageSize() *int32
	SetPlatform(v string) *DescribeResourceGroupsRequest
	GetPlatform() *string
	SetResourceGroupIds(v []*string) *DescribeResourceGroupsRequest
	GetResourceGroupIds() []*string
	SetResourceGroupName(v string) *DescribeResourceGroupsRequest
	GetResourceGroupName() *string
}

type DescribeResourceGroupsRequest struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 0
	NeedContainResourceGroupWithOfficeSite *int64 `json:"NeedContainResourceGroupWithOfficeSite,omitempty" xml:"NeedContainResourceGroupWithOfficeSite,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// >  Set the value to AliyunConsole.
	//
	// 	- This parameter is not publicly available on other platforms.
	//
	// example:
	//
	// AliyunConsole
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The IDs of the resource groups that you want to query.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// The name of the resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s DescribeResourceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsRequest) GetNeedContainResourceGroupWithOfficeSite() *int64 {
	return s.NeedContainResourceGroupWithOfficeSite
}

func (s *DescribeResourceGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeResourceGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeResourceGroupsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeResourceGroupsRequest) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *DescribeResourceGroupsRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeResourceGroupsRequest) SetNeedContainResourceGroupWithOfficeSite(v int64) *DescribeResourceGroupsRequest {
	s.NeedContainResourceGroupWithOfficeSite = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetPageNumber(v int32) *DescribeResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetPageSize(v int32) *DescribeResourceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetPlatform(v string) *DescribeResourceGroupsRequest {
	s.Platform = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetResourceGroupIds(v []*string) *DescribeResourceGroupsRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *DescribeResourceGroupsRequest) SetResourceGroupName(v string) *DescribeResourceGroupsRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeResourceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
