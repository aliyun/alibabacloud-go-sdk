// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentPlatformsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGbId(v string) *DescribeParentPlatformsRequest
	GetGbId() *string
	SetOwnerId(v int64) *DescribeParentPlatformsRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeParentPlatformsRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeParentPlatformsRequest
	GetPageSize() *int64
	SetSortBy(v string) *DescribeParentPlatformsRequest
	GetSortBy() *string
	SetSortDirection(v string) *DescribeParentPlatformsRequest
	GetSortDirection() *string
	SetStatus(v string) *DescribeParentPlatformsRequest
	GetStatus() *string
}

type DescribeParentPlatformsRequest struct {
	// example:
	//
	// 31000*****2170123451
	GbId    *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
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
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeParentPlatformsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentPlatformsRequest) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformsRequest) GetGbId() *string {
	return s.GbId
}

func (s *DescribeParentPlatformsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParentPlatformsRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeParentPlatformsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeParentPlatformsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeParentPlatformsRequest) GetSortDirection() *string {
	return s.SortDirection
}

func (s *DescribeParentPlatformsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeParentPlatformsRequest) SetGbId(v string) *DescribeParentPlatformsRequest {
	s.GbId = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetOwnerId(v int64) *DescribeParentPlatformsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetPageNum(v int64) *DescribeParentPlatformsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetPageSize(v int64) *DescribeParentPlatformsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetSortBy(v string) *DescribeParentPlatformsRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetSortDirection(v string) *DescribeParentPlatformsRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetStatus(v string) *DescribeParentPlatformsRequest {
	s.Status = &v
	return s
}

func (s *DescribeParentPlatformsRequest) Validate() error {
	return dara.Validate(s)
}
