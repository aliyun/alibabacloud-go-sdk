// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCubePerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryCubePerformanceResponseBody
	GetRequestId() *string
	SetResult(v []*QueryCubePerformanceResponseBodyResult) *QueryCubePerformanceResponseBody
	GetResult() []*QueryCubePerformanceResponseBodyResult
	SetSuccess(v bool) *QueryCubePerformanceResponseBody
	GetSuccess() *bool
}

type QueryCubePerformanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 685072a0-1fd5-40ef-ae6b-cf94e79e718f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Array of report objects
	Result []*QueryCubePerformanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCubePerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCubePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCubePerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCubePerformanceResponseBody) GetResult() []*QueryCubePerformanceResponseBodyResult {
	return s.Result
}

func (s *QueryCubePerformanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCubePerformanceResponseBody) SetRequestId(v string) *QueryCubePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCubePerformanceResponseBody) SetResult(v []*QueryCubePerformanceResponseBodyResult) *QueryCubePerformanceResponseBody {
	s.Result = v
	return s
}

func (s *QueryCubePerformanceResponseBody) SetSuccess(v bool) *QueryCubePerformanceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCubePerformanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryCubePerformanceResponseBodyResult struct {
	// The average duration of cache hits.
	//
	// example:
	//
	// 1
	CacheCostTimeAvg *float64 `json:"CacheCostTimeAvg,omitempty" xml:"CacheCostTimeAvg,omitempty"`
	// The number of cache hits.
	//
	// example:
	//
	// 1
	CacheQueryCount *int32 `json:"CacheQueryCount,omitempty" xml:"CacheQueryCount,omitempty"`
	// The average query duration associated with the SQL pattern.
	//
	// example:
	//
	// 1
	CostTimeAvg *float64 `json:"CostTimeAvg,omitempty" xml:"CostTimeAvg,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// test
	CubeName *string `json:"CubeName,omitempty" xml:"CubeName,omitempty"`
	// The number of queries.
	//
	// example:
	//
	// 50
	QueryCount *int32 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The average number of queries.
	//
	// example:
	//
	// 1
	QueryCountAvg *float64 `json:"QueryCountAvg,omitempty" xml:"QueryCountAvg,omitempty"`
	// The percentage of the number of queries that exceed the 5S.
	//
	// example:
	//
	// 1.0
	QueryOverFivePercentNum *float64 `json:"QueryOverFivePercentNum,omitempty" xml:"QueryOverFivePercentNum,omitempty"`
	// Query the proportion of more than 5S.
	//
	// example:
	//
	// 1.0
	QueryOverFiveSecPercent *string `json:"QueryOverFiveSecPercent,omitempty" xml:"QueryOverFiveSecPercent,omitempty"`
	// The percentage of queries that exceed 10s.
	//
	// example:
	//
	// 1.0
	QueryOverTenSecPercent *string `json:"QueryOverTenSecPercent,omitempty" xml:"QueryOverTenSecPercent,omitempty"`
	// The percentage of queries that exceed 10s.
	//
	// example:
	//
	// 1.0
	QueryOverTenSecPercentNum *float64 `json:"QueryOverTenSecPercentNum,omitempty" xml:"QueryOverTenSecPercentNum,omitempty"`
	// The number of times that the chart query times out.
	//
	// example:
	//
	// 1
	QueryTimeoutCount *int32 `json:"QueryTimeoutCount,omitempty" xml:"QueryTimeoutCount,omitempty"`
	// The percentage of timeout times for chart queries.
	//
	// example:
	//
	// 1
	QueryTimeoutCountPercent *float64 `json:"QueryTimeoutCountPercent,omitempty" xml:"QueryTimeoutCountPercent,omitempty"`
	// The average time consumed by the Quick engine query.
	//
	// example:
	//
	// 1
	QuickIndexCostTimeAvg *float64 `json:"QuickIndexCostTimeAvg,omitempty" xml:"QuickIndexCostTimeAvg,omitempty"`
	// The number of times that the Quick engine is hit.
	//
	// example:
	//
	// 1
	QuickIndexQueryCount *int32 `json:"QuickIndexQueryCount,omitempty" xml:"QuickIndexQueryCount,omitempty"`
	// The proportion of duplicate queries.
	//
	// example:
	//
	// 0.3
	RepeatQueryPercent *string `json:"RepeatQueryPercent,omitempty" xml:"RepeatQueryPercent,omitempty"`
	// The number of duplicate queries.
	//
	// example:
	//
	// 1
	RepeatQueryPercentNum *float64 `json:"RepeatQueryPercentNum,omitempty" xml:"RepeatQueryPercentNum,omitempty"`
	// The number of times the query is repeated.
	//
	// example:
	//
	// 1
	RepeatSqlQueryCount *int32 `json:"RepeatSqlQueryCount,omitempty" xml:"RepeatSqlQueryCount,omitempty"`
	// The proportion of duplicate queries.
	//
	// example:
	//
	// 1
	RepeatSqlQueryPercent *string `json:"RepeatSqlQueryPercent,omitempty" xml:"RepeatSqlQueryPercent,omitempty"`
	// The ID of the workspace to which the work belongs.
	//
	// example:
	//
	// 87c6b145-090c-43e1-9426-8f93be23****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// taascontainerprod
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryCubePerformanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryCubePerformanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCubePerformanceResponseBodyResult) GetCacheCostTimeAvg() *float64 {
	return s.CacheCostTimeAvg
}

func (s *QueryCubePerformanceResponseBodyResult) GetCacheQueryCount() *int32 {
	return s.CacheQueryCount
}

func (s *QueryCubePerformanceResponseBodyResult) GetCostTimeAvg() *float64 {
	return s.CostTimeAvg
}

func (s *QueryCubePerformanceResponseBodyResult) GetCubeId() *string {
	return s.CubeId
}

func (s *QueryCubePerformanceResponseBodyResult) GetCubeName() *string {
	return s.CubeName
}

func (s *QueryCubePerformanceResponseBodyResult) GetQueryCount() *int32 {
	return s.QueryCount
}

func (s *QueryCubePerformanceResponseBodyResult) GetQueryCountAvg() *float64 {
	return s.QueryCountAvg
}

func (s *QueryCubePerformanceResponseBodyResult) GetQueryOverFivePercentNum() *float64 {
	return s.QueryOverFivePercentNum
}

func (s *QueryCubePerformanceResponseBodyResult) GetQueryOverFiveSecPercent() *string {
	return s.QueryOverFiveSecPercent
}

func (s *QueryCubePerformanceResponseBodyResult) GetQueryOverTenSecPercent() *string {
	return s.QueryOverTenSecPercent
}

func (s *QueryCubePerformanceResponseBodyResult) GetQueryOverTenSecPercentNum() *float64 {
	return s.QueryOverTenSecPercentNum
}

func (s *QueryCubePerformanceResponseBodyResult) GetQueryTimeoutCount() *int32 {
	return s.QueryTimeoutCount
}

func (s *QueryCubePerformanceResponseBodyResult) GetQueryTimeoutCountPercent() *float64 {
	return s.QueryTimeoutCountPercent
}

func (s *QueryCubePerformanceResponseBodyResult) GetQuickIndexCostTimeAvg() *float64 {
	return s.QuickIndexCostTimeAvg
}

func (s *QueryCubePerformanceResponseBodyResult) GetQuickIndexQueryCount() *int32 {
	return s.QuickIndexQueryCount
}

func (s *QueryCubePerformanceResponseBodyResult) GetRepeatQueryPercent() *string {
	return s.RepeatQueryPercent
}

func (s *QueryCubePerformanceResponseBodyResult) GetRepeatQueryPercentNum() *float64 {
	return s.RepeatQueryPercentNum
}

func (s *QueryCubePerformanceResponseBodyResult) GetRepeatSqlQueryCount() *int32 {
	return s.RepeatSqlQueryCount
}

func (s *QueryCubePerformanceResponseBodyResult) GetRepeatSqlQueryPercent() *string {
	return s.RepeatSqlQueryPercent
}

func (s *QueryCubePerformanceResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryCubePerformanceResponseBodyResult) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryCubePerformanceResponseBodyResult) SetCacheCostTimeAvg(v float64) *QueryCubePerformanceResponseBodyResult {
	s.CacheCostTimeAvg = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetCacheQueryCount(v int32) *QueryCubePerformanceResponseBodyResult {
	s.CacheQueryCount = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetCostTimeAvg(v float64) *QueryCubePerformanceResponseBodyResult {
	s.CostTimeAvg = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetCubeId(v string) *QueryCubePerformanceResponseBodyResult {
	s.CubeId = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetCubeName(v string) *QueryCubePerformanceResponseBodyResult {
	s.CubeName = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryCount(v int32) *QueryCubePerformanceResponseBodyResult {
	s.QueryCount = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryCountAvg(v float64) *QueryCubePerformanceResponseBodyResult {
	s.QueryCountAvg = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryOverFivePercentNum(v float64) *QueryCubePerformanceResponseBodyResult {
	s.QueryOverFivePercentNum = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryOverFiveSecPercent(v string) *QueryCubePerformanceResponseBodyResult {
	s.QueryOverFiveSecPercent = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryOverTenSecPercent(v string) *QueryCubePerformanceResponseBodyResult {
	s.QueryOverTenSecPercent = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryOverTenSecPercentNum(v float64) *QueryCubePerformanceResponseBodyResult {
	s.QueryOverTenSecPercentNum = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryTimeoutCount(v int32) *QueryCubePerformanceResponseBodyResult {
	s.QueryTimeoutCount = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryTimeoutCountPercent(v float64) *QueryCubePerformanceResponseBodyResult {
	s.QueryTimeoutCountPercent = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQuickIndexCostTimeAvg(v float64) *QueryCubePerformanceResponseBodyResult {
	s.QuickIndexCostTimeAvg = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQuickIndexQueryCount(v int32) *QueryCubePerformanceResponseBodyResult {
	s.QuickIndexQueryCount = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetRepeatQueryPercent(v string) *QueryCubePerformanceResponseBodyResult {
	s.RepeatQueryPercent = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetRepeatQueryPercentNum(v float64) *QueryCubePerformanceResponseBodyResult {
	s.RepeatQueryPercentNum = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetRepeatSqlQueryCount(v int32) *QueryCubePerformanceResponseBodyResult {
	s.RepeatSqlQueryCount = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetRepeatSqlQueryPercent(v string) *QueryCubePerformanceResponseBodyResult {
	s.RepeatSqlQueryPercent = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetWorkspaceId(v string) *QueryCubePerformanceResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetWorkspaceName(v string) *QueryCubePerformanceResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
