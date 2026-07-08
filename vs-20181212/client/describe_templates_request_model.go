// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeTemplatesRequest
	GetId() *string
	SetInstanceId(v string) *DescribeTemplatesRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DescribeTemplatesRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeTemplatesRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeTemplatesRequest
	GetPageSize() *int64
	SetSortBy(v string) *DescribeTemplatesRequest
	GetSortBy() *string
	SetSortDirection(v string) *DescribeTemplatesRequest
	GetSortDirection() *string
	SetType(v string) *DescribeTemplatesRequest
	GetType() *string
}

type DescribeTemplatesRequest struct {
	// Template ID.
	//
	// example:
	//
	// 323434****83423432
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Filter templates by the bound instance ID.
	//
	// example:
	//
	// 323*****998-cn-qingdao
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Page number. Default: 1.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of entries per page. Default: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Sort results by the specified field. Default: sort by ID.
	//
	// example:
	//
	// id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Sort order. Default: ascending. Valid values:
	//
	// - asc (ascending)
	//
	// - desc (descending)
	//
	// example:
	//
	// asc
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	// Filter templates by template type.
	//
	// example:
	//
	// record
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesRequest) GetId() *string {
	return s.Id
}

func (s *DescribeTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTemplatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTemplatesRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeTemplatesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeTemplatesRequest) GetSortDirection() *string {
	return s.SortDirection
}

func (s *DescribeTemplatesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeTemplatesRequest) SetId(v string) *DescribeTemplatesRequest {
	s.Id = &v
	return s
}

func (s *DescribeTemplatesRequest) SetInstanceId(v string) *DescribeTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTemplatesRequest) SetOwnerId(v int64) *DescribeTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTemplatesRequest) SetPageNum(v int64) *DescribeTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeTemplatesRequest) SetPageSize(v int64) *DescribeTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplatesRequest) SetSortBy(v string) *DescribeTemplatesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeTemplatesRequest) SetSortDirection(v string) *DescribeTemplatesRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeTemplatesRequest) SetType(v string) *DescribeTemplatesRequest {
	s.Type = &v
	return s
}

func (s *DescribeTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
