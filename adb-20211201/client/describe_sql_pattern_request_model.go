// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlPatternRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSqlPatternRequest
	GetDBClusterId() *string
	SetOrder(v string) *DescribeSqlPatternRequest
	GetOrder() *string
	SetPageNumber(v int32) *DescribeSqlPatternRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSqlPatternRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSqlPatternRequest
	GetRegionId() *string
	SetSqlPattern(v string) *DescribeSqlPatternRequest
	GetSqlPattern() *string
	SetStartTime(v string) *DescribeSqlPatternRequest
	GetStartTime() *string
	SetType(v string) *DescribeSqlPatternRequest
	GetType() *string
}

type DescribeSqlPatternRequest struct {
	// The cluster ID.
	//
	// > You can call the DescribeDBClusters operation to query the IDs of all clusters in a specified region.
	//
	// [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1ej1nq9n6****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies the sorting of query results by a specified field. The format is a JSON character string, such as `[{"Field":"Pattern","Type":"Asc"}]`. Parameters:
	//
	// - Field: the sorting field. Valid values:
	//
	//     - Pattern: SQL pattern.
	//
	//     - AccessIP: client IP address.
	//
	//     - User: username.
	//
	//     - QueryCount: total number of queries.
	//
	//     - AvgPeakMemory: average peak memory. Unit: KB.
	//
	//     - MaxPeakMemory: Peak Volume memory. Unit: KB.
	//
	//     - AvgCpuTime: average computing duration. Unit: milliseconds.
	//
	//     - MaxCpuTime: maximum computing duration. Unit: milliseconds.
	//
	//     - AvgStageCount: average number of stages.
	//
	//     - MaxStageCount: maximum number of stages.
	//
	//     - AvgTaskCount: average number of tasks.
	//
	//     - MaxTaskCount: maximum number of tasks.
	//
	//     - AvgScanSize: average scan size. Unit: KB.
	//
	//     - MaxScanSize: maximum scan size. Unit: KB.
	//
	// - Type: the sorting method. Valid values:
	//
	//     - Asc: ascending sorting.
	//
	//     - Desc: descending sorting.
	//
	// > - If no Settings are specified for this parameter, the results are sorted in ascending order by the Pattern field by default.
	//
	// - To perform sorting by the AccessIP field, you must set the Type parameter to accessip. To perform sorting by the User field, you must leave the Type parameter empty or set it to user.
	//
	// example:
	//
	// [{"Field":"Pattern","Type":"Asc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Valid values: integers greater than 0 and no greater than the maximum value of the Integer type.
	//
	// Default value: 1.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **10*	- (default)
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keyword that the SQL pattern must contain.
	//
	// > If this parameter is not specified, the details of all SQL patterns for the cluster on the date specified by the StartTime parameter are returned by default.
	//
	// example:
	//
	// SELECT
	SqlPattern *string `json:"SqlPattern,omitempty" xml:"SqlPattern,omitempty"`
	// The start date of the query. The time is in UTC and formatted as yyyy-MM-dd.
	//
	// > You can only query data within the last 30 days.
	//
	// example:
	//
	// 2022-08-30
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Aggregates SQL patterns by a specified dimension. Valid values:
	//
	//  - user: aggregates by user.
	//
	//  - accessip: aggregates by access IP address.
	//
	// > If this parameter is not specified, SQL patterns are aggregated by user by default.
	//
	// example:
	//
	// user
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSqlPatternRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSqlPatternRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeSqlPatternRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSqlPatternRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSqlPatternRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSqlPatternRequest) GetSqlPattern() *string {
	return s.SqlPattern
}

func (s *DescribeSqlPatternRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSqlPatternRequest) GetType() *string {
	return s.Type
}

func (s *DescribeSqlPatternRequest) SetDBClusterId(v string) *DescribeSqlPatternRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetOrder(v string) *DescribeSqlPatternRequest {
	s.Order = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetPageNumber(v int32) *DescribeSqlPatternRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetPageSize(v int32) *DescribeSqlPatternRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetRegionId(v string) *DescribeSqlPatternRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetSqlPattern(v string) *DescribeSqlPatternRequest {
	s.SqlPattern = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetStartTime(v string) *DescribeSqlPatternRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetType(v string) *DescribeSqlPatternRequest {
	s.Type = &v
	return s
}

func (s *DescribeSqlPatternRequest) Validate() error {
	return dara.Validate(s)
}
