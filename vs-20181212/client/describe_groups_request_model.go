// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeGroupsRequest
	GetId() *string
	SetInProtocol(v string) *DescribeGroupsRequest
	GetInProtocol() *string
	SetIncludeStats(v bool) *DescribeGroupsRequest
	GetIncludeStats() *bool
	SetName(v string) *DescribeGroupsRequest
	GetName() *string
	SetOwnerId(v int64) *DescribeGroupsRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeGroupsRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeGroupsRequest
	GetPageSize() *int64
	SetRegion(v string) *DescribeGroupsRequest
	GetRegion() *string
	SetSortBy(v string) *DescribeGroupsRequest
	GetSortBy() *string
	SetSortDirection(v string) *DescribeGroupsRequest
	GetSortDirection() *string
	SetStatus(v string) *DescribeGroupsRequest
	GetStatus() *string
}

type DescribeGroupsRequest struct {
	// example:
	//
	// 33763950877224964-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// rtmp
	InProtocol *string `json:"InProtocol,omitempty" xml:"InProtocol,omitempty"`
	// example:
	//
	// false
	IncludeStats *bool   `json:"IncludeStats,omitempty" xml:"IncludeStats,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// asc
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupsRequest) GetId() *string {
	return s.Id
}

func (s *DescribeGroupsRequest) GetInProtocol() *string {
	return s.InProtocol
}

func (s *DescribeGroupsRequest) GetIncludeStats() *bool {
	return s.IncludeStats
}

func (s *DescribeGroupsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGroupsRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeGroupsRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeGroupsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeGroupsRequest) GetSortDirection() *string {
	return s.SortDirection
}

func (s *DescribeGroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeGroupsRequest) SetId(v string) *DescribeGroupsRequest {
	s.Id = &v
	return s
}

func (s *DescribeGroupsRequest) SetInProtocol(v string) *DescribeGroupsRequest {
	s.InProtocol = &v
	return s
}

func (s *DescribeGroupsRequest) SetIncludeStats(v bool) *DescribeGroupsRequest {
	s.IncludeStats = &v
	return s
}

func (s *DescribeGroupsRequest) SetName(v string) *DescribeGroupsRequest {
	s.Name = &v
	return s
}

func (s *DescribeGroupsRequest) SetOwnerId(v int64) *DescribeGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGroupsRequest) SetPageNum(v int64) *DescribeGroupsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeGroupsRequest) SetPageSize(v int64) *DescribeGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupsRequest) SetRegion(v string) *DescribeGroupsRequest {
	s.Region = &v
	return s
}

func (s *DescribeGroupsRequest) SetSortBy(v string) *DescribeGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeGroupsRequest) SetSortDirection(v string) *DescribeGroupsRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeGroupsRequest) SetStatus(v string) *DescribeGroupsRequest {
	s.Status = &v
	return s
}

func (s *DescribeGroupsRequest) Validate() error {
	return dara.Validate(s)
}
