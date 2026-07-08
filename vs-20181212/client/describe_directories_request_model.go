// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeDirectoriesRequest
	GetGroupId() *string
	SetNoPagination(v bool) *DescribeDirectoriesRequest
	GetNoPagination() *bool
	SetOwnerId(v int64) *DescribeDirectoriesRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeDirectoriesRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeDirectoriesRequest
	GetPageSize() *int64
	SetParentId(v string) *DescribeDirectoriesRequest
	GetParentId() *string
	SetSortBy(v string) *DescribeDirectoriesRequest
	GetSortBy() *string
	SetSortDirection(v string) *DescribeDirectoriesRequest
	GetSortDirection() *string
}

type DescribeDirectoriesRequest struct {
	// ID of the group to which the directory belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Disable paging. Default is false.
	//
	// example:
	//
	// false
	NoPagination *bool  `json:"NoPagination,omitempty" xml:"NoPagination,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Page number. Default is 1.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Page size. Default is 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Query by parent directory ID.
	//
	// example:
	//
	// 399*****774-cn-qingdao
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// Sort by the specified field. Default is by ID.
	//
	// example:
	//
	// ID
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Sort order. Default is ascending. Values:
	//
	// - asc (ascending)
	//
	// - desc (descending)
	//
	// example:
	//
	// asc
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
}

func (s DescribeDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDirectoriesRequest) GetNoPagination() *bool {
	return s.NoPagination
}

func (s *DescribeDirectoriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDirectoriesRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeDirectoriesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDirectoriesRequest) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDirectoriesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeDirectoriesRequest) GetSortDirection() *string {
	return s.SortDirection
}

func (s *DescribeDirectoriesRequest) SetGroupId(v string) *DescribeDirectoriesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetNoPagination(v bool) *DescribeDirectoriesRequest {
	s.NoPagination = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetOwnerId(v int64) *DescribeDirectoriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetPageNum(v int64) *DescribeDirectoriesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetPageSize(v int64) *DescribeDirectoriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetParentId(v string) *DescribeDirectoriesRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetSortBy(v string) *DescribeDirectoriesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetSortDirection(v string) *DescribeDirectoriesRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
