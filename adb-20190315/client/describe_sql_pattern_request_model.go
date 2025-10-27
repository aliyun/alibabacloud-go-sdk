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
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON string format. Example: `[{"Field":"Pattern","Type":"Asc"}]`. Parameters:
	//
	// 	- `Field` specifies the field by which to sort the query results. Valid values:
	//
	//     	- `Pattern`: the SQL pattern.
	//
	//     	- `AccessIP`: the IP address of the client.
	//
	//     	- `User`: the username.
	//
	//     	- `QueryCount`: the number of queries performed in association with the SQL pattern within the time range to query.
	//
	//     	- `AvgPeakMemory`: the average peak memory usage of the SQL pattern within the time range to query. Unit: KB.
	//
	//     	- `MaxPeakMemory`: the maximum peak memory usage of the SQL pattern within the time range to query. Unit: KB.
	//
	//     	- `AvgCpuTime`: the average execution duration of the SQL pattern within the time range to query. Unit: milliseconds.
	//
	//     	- `MaxCpuTime`: the maximum execution duration of the SQL pattern within the time range to query. Unit: milliseconds.
	//
	//     	- `AvgStageCount`: the average number of stages.
	//
	//     	- `MaxStageCount`: the maximum number of stages.
	//
	//     	- `AvgTaskCount`: the average number of tasks.
	//
	//     	- `MaxTaskCount`: the maximum number of tasks.
	//
	//     	- `AvgScanSize`: the average amount of data scanned based on the SQL pattern within the time range to query. Unit: KB.
	//
	//     	- `MaxScanSize`: the maximum amount of data scanned based on the SQL pattern within the time range to query. Unit: KB.
	//
	// 	- `Type` specifies the sorting order. Valid values:
	//
	//     	- `Asc`: ascending order.
	//
	//     	- `Desc`: descending order.
	//
	// >
	//
	// 	- If you do not specify this parameter, query results are sorted in ascending order of `Pattern`.
	//
	// 	- If you want to sort query results by `AccessIP`, you must set the `Type` parameter to `accessip`. If you want to sort query results by `User`, you must leave the `Type` parameter empty or set it to `user`.
	//
	// example:
	//
	// [{"Field":"Pattern","Type":"Asc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The value must be a positive integer. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keyword that is used for the query.
	//
	// > If you do not specify this parameter, all SQL patterns of the AnalyticDB for MySQL cluster within the time period specified by `StartTime` are returned.
	//
	// example:
	//
	// SELECT
	SqlPattern *string `json:"SqlPattern,omitempty" xml:"SqlPattern,omitempty"`
	// The start date to query. Specify the time in the *yyyy-MM-dd	- format. The time must be in UTC.
	//
	// > Only data within the last 30 days can be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-08-30
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The dimension by which to aggregate the SQL patterns. Valid values:
	//
	// 	- `user`: aggregates the SQL patterns by user.
	//
	// 	- `accessip`: aggregates the SQL patterns by client IP address.
	//
	// > If you do not specify this parameter, the SQL patterns are aggregated by `user`.
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
