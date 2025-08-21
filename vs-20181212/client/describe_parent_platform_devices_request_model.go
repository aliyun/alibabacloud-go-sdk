// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentPlatformDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeParentPlatformDevicesRequest
	GetId() *string
	SetOwnerId(v int64) *DescribeParentPlatformDevicesRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeParentPlatformDevicesRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeParentPlatformDevicesRequest
	GetPageSize() *int64
	SetSortBy(v string) *DescribeParentPlatformDevicesRequest
	GetSortBy() *string
	SetSortDirection(v string) *DescribeParentPlatformDevicesRequest
	GetSortDirection() *string
}

type DescribeParentPlatformDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 359*****374-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// asc
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
}

func (s DescribeParentPlatformDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentPlatformDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformDevicesRequest) GetId() *string {
	return s.Id
}

func (s *DescribeParentPlatformDevicesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParentPlatformDevicesRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeParentPlatformDevicesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeParentPlatformDevicesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeParentPlatformDevicesRequest) GetSortDirection() *string {
	return s.SortDirection
}

func (s *DescribeParentPlatformDevicesRequest) SetId(v string) *DescribeParentPlatformDevicesRequest {
	s.Id = &v
	return s
}

func (s *DescribeParentPlatformDevicesRequest) SetOwnerId(v int64) *DescribeParentPlatformDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParentPlatformDevicesRequest) SetPageNum(v int64) *DescribeParentPlatformDevicesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeParentPlatformDevicesRequest) SetPageSize(v int64) *DescribeParentPlatformDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeParentPlatformDevicesRequest) SetSortBy(v string) *DescribeParentPlatformDevicesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeParentPlatformDevicesRequest) SetSortDirection(v string) *DescribeParentPlatformDevicesRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeParentPlatformDevicesRequest) Validate() error {
	return dara.Validate(s)
}
