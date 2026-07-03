// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventDisposeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeEventDisposeRequest
	GetCurrentPage() *int32
	SetIncidentUuid(v string) *DescribeEventDisposeRequest
	GetIncidentUuid() *string
	SetPageSize(v int32) *DescribeEventDisposeRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeEventDisposeRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeEventDisposeRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeEventDisposeRequest
	GetRoleType() *int32
}

type DescribeEventDisposeRequest struct {
	// The number of the page to return. The value must be greater than or equal to 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The globally unique UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The number of entries to return on each page. The maximum value is 500.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the Data Management center of Threat Analysis is located. Select the region based on the region where your assets are deployed. Valid values:
	//
	// - cn-hangzhou: Your assets are deployed in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are deployed in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can use this parameter to switch to the perspective of a member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts that belong to the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeEventDisposeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDisposeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeEventDisposeRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeEventDisposeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEventDisposeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventDisposeRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeEventDisposeRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeEventDisposeRequest) SetCurrentPage(v int32) *DescribeEventDisposeRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetIncidentUuid(v string) *DescribeEventDisposeRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetPageSize(v int32) *DescribeEventDisposeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetRegionId(v string) *DescribeEventDisposeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetRoleFor(v int64) *DescribeEventDisposeRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetRoleType(v int32) *DescribeEventDisposeRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeEventDisposeRequest) Validate() error {
	return dara.Validate(s)
}
