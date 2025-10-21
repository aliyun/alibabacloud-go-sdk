// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeProcessListRequest
	GetDBInstanceId() *string
	SetInitialQueryId(v string) *DescribeProcessListRequest
	GetInitialQueryId() *string
	SetInitialUser(v string) *DescribeProcessListRequest
	GetInitialUser() *string
	SetKeyword(v string) *DescribeProcessListRequest
	GetKeyword() *string
	SetPageNumber(v int32) *DescribeProcessListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeProcessListRequest
	GetPageSize() *int32
	SetQueryDurationMs(v string) *DescribeProcessListRequest
	GetQueryDurationMs() *string
	SetQueryOrder(v int64) *DescribeProcessListRequest
	GetQueryOrder() *int64
	SetRegionId(v string) *DescribeProcessListRequest
	GetRegionId() *string
}

type DescribeProcessListRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 1
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The user who executes the query statement.
	//
	// example:
	//
	// testuser
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The keyword of the query statement.
	//
	// example:
	//
	// SELECT
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The execution duration of slow SQL queries. Minimum value: 1000. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	QueryDurationMs *string `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// Specifies the columns by which the query results are sorted in descending order.
	//
	// 	- 0: The query results are sorted by the query_duration_ms column.
	//
	// 	- 1: The query results are sorted by the query_duration_ms and query_start_time columns.
	//
	// 	- 2: The query results are sorted by the query_duration_ms, query_start_time, and user columns.
	//
	// example:
	//
	// id
	QueryOrder *int64 `json:"QueryOrder,omitempty" xml:"QueryOrder,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeProcessListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeProcessListRequest) GetInitialQueryId() *string {
	return s.InitialQueryId
}

func (s *DescribeProcessListRequest) GetInitialUser() *string {
	return s.InitialUser
}

func (s *DescribeProcessListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeProcessListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeProcessListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeProcessListRequest) GetQueryDurationMs() *string {
	return s.QueryDurationMs
}

func (s *DescribeProcessListRequest) GetQueryOrder() *int64 {
	return s.QueryOrder
}

func (s *DescribeProcessListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeProcessListRequest) SetDBInstanceId(v string) *DescribeProcessListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeProcessListRequest) SetInitialQueryId(v string) *DescribeProcessListRequest {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeProcessListRequest) SetInitialUser(v string) *DescribeProcessListRequest {
	s.InitialUser = &v
	return s
}

func (s *DescribeProcessListRequest) SetKeyword(v string) *DescribeProcessListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageNumber(v int32) *DescribeProcessListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageSize(v int32) *DescribeProcessListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessListRequest) SetQueryDurationMs(v string) *DescribeProcessListRequest {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeProcessListRequest) SetQueryOrder(v int64) *DescribeProcessListRequest {
	s.QueryOrder = &v
	return s
}

func (s *DescribeProcessListRequest) SetRegionId(v string) *DescribeProcessListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProcessListRequest) Validate() error {
	return dara.Validate(s)
}
