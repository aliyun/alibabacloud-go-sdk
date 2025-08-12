// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogMonitorListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DescribeLogMonitorListRequest
	GetGroupId() *int64
	SetPageNumber(v int32) *DescribeLogMonitorListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLogMonitorListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLogMonitorListRequest
	GetRegionId() *string
	SetSearchValue(v string) *DescribeLogMonitorListRequest
	GetSearchValue() *string
}

type DescribeLogMonitorListRequest struct {
	// The ID of the application group.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keyword that is used to search for log monitoring metrics. Fuzzy match is supported.
	//
	// example:
	//
	// test
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
}

func (s DescribeLogMonitorListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMonitorListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorListRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeLogMonitorListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLogMonitorListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLogMonitorListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogMonitorListRequest) GetSearchValue() *string {
	return s.SearchValue
}

func (s *DescribeLogMonitorListRequest) SetGroupId(v int64) *DescribeLogMonitorListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeLogMonitorListRequest) SetPageNumber(v int32) *DescribeLogMonitorListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogMonitorListRequest) SetPageSize(v int32) *DescribeLogMonitorListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogMonitorListRequest) SetRegionId(v string) *DescribeLogMonitorListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogMonitorListRequest) SetSearchValue(v string) *DescribeLogMonitorListRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeLogMonitorListRequest) Validate() error {
	return dara.Validate(s)
}
