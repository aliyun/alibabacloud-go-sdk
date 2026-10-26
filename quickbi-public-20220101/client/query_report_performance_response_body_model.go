// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReportPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryReportPerformanceResponseBody
	GetRequestId() *string
	SetResult(v []*QueryReportPerformanceResponseBodyResult) *QueryReportPerformanceResponseBody
	GetResult() []*QueryReportPerformanceResponseBodyResult
	SetSuccess(v bool) *QueryReportPerformanceResponseBody
	GetSuccess() *bool
}

type QueryReportPerformanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1FC71085-D5FD-08E0-813A-4D4BD1031BC5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of report performance logs, which includes report information and performance statistics such as query count, query duration, and cache hits.
	Result []*QueryReportPerformanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryReportPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryReportPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReportPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryReportPerformanceResponseBody) GetResult() []*QueryReportPerformanceResponseBodyResult {
	return s.Result
}

func (s *QueryReportPerformanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryReportPerformanceResponseBody) SetRequestId(v string) *QueryReportPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReportPerformanceResponseBody) SetResult(v []*QueryReportPerformanceResponseBodyResult) *QueryReportPerformanceResponseBody {
	s.Result = v
	return s
}

func (s *QueryReportPerformanceResponseBody) SetSuccess(v bool) *QueryReportPerformanceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryReportPerformanceResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryReportPerformanceResponseBodyResult struct {
	// The average time consumed for cache hits.
	//
	// example:
	//
	// 2.2
	CacheCostTimeAvg *float64 `json:"CacheCostTimeAvg,omitempty" xml:"CacheCostTimeAvg,omitempty"`
	// The number of cache hits.
	//
	// example:
	//
	// 1
	CacheQueryCount *int32 `json:"CacheQueryCount,omitempty" xml:"CacheQueryCount,omitempty"`
	// The number of chart queries.
	//
	// example:
	//
	// 1
	ComponentQueryCount *int32 `json:"ComponentQueryCount,omitempty" xml:"ComponentQueryCount,omitempty"`
	// The average number of chart queries.
	//
	// example:
	//
	// 2.0
	ComponentQueryCountAvg *float64 `json:"ComponentQueryCountAvg,omitempty" xml:"ComponentQueryCountAvg,omitempty"`
	// The average query time.
	//
	// example:
	//
	// 0.2
	CostTimeAvg *float64 `json:"CostTimeAvg,omitempty" xml:"CostTimeAvg,omitempty"`
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
	// 3.3
	QueryCountAvg *float64 `json:"QueryCountAvg,omitempty" xml:"QueryCountAvg,omitempty"`
	// The proportion of queries that take more than 5 seconds.
	//
	// example:
	//
	// 0.5
	QueryOverFivePercentNum *float64 `json:"QueryOverFivePercentNum,omitempty" xml:"QueryOverFivePercentNum,omitempty"`
	// The percentage of queries that take more than 5 seconds.
	//
	// example:
	//
	// 0.5
	QueryOverFiveSecPercent *string `json:"QueryOverFiveSecPercent,omitempty" xml:"QueryOverFiveSecPercent,omitempty"`
	// The percentage of queries that take more than 10 seconds.
	//
	// example:
	//
	// 0.5
	QueryOverTenSecPercent *string `json:"QueryOverTenSecPercent,omitempty" xml:"QueryOverTenSecPercent,omitempty"`
	// The number of queries that take more than 10 seconds.
	//
	// example:
	//
	// 0.5
	QueryOverTenSecPercentNum *float64 `json:"QueryOverTenSecPercentNum,omitempty" xml:"QueryOverTenSecPercentNum,omitempty"`
	// The number of chart query timeouts.
	//
	// example:
	//
	// 8
	QueryTimeoutCount *int32 `json:"QueryTimeoutCount,omitempty" xml:"QueryTimeoutCount,omitempty"`
	// The proportion of chart query timeouts.
	//
	// example:
	//
	// 0.5
	QueryTimeoutCountPercent *float64 `json:"QueryTimeoutCountPercent,omitempty" xml:"QueryTimeoutCountPercent,omitempty"`
	// The average query time for Quick engine hits.
	//
	// example:
	//
	// 10
	QuickIndexCostTimeAvg *float64 `json:"QuickIndexCostTimeAvg,omitempty" xml:"QuickIndexCostTimeAvg,omitempty"`
	// The number of Quick engine hits.
	//
	// example:
	//
	// 5
	QuickIndexQueryCount *int32 `json:"QuickIndexQueryCount,omitempty" xml:"QuickIndexQueryCount,omitempty"`
	// The proportion of repeated queries.
	//
	// example:
	//
	// 0.8
	RepeatQueryPercent *string `json:"RepeatQueryPercent,omitempty" xml:"RepeatQueryPercent,omitempty"`
	// The number of repeated queries.
	//
	// example:
	//
	// 3
	RepeatQueryPercentNum *float64 `json:"RepeatQueryPercentNum,omitempty" xml:"RepeatQueryPercentNum,omitempty"`
	// The number of repeated SQL queries.
	//
	// example:
	//
	// 1
	RepeatSqlQueryCount *int32 `json:"RepeatSqlQueryCount,omitempty" xml:"RepeatSqlQueryCount,omitempty"`
	// The proportion of repeated queries.
	//
	// example:
	//
	// 0.7
	RepeatSqlQueryPercent *string `json:"RepeatSqlQueryPercent,omitempty" xml:"RepeatSqlQueryPercent,omitempty"`
	// The report ID.
	//
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The report name.
	//
	// example:
	//
	// ClusterAddonUpgradeReport
	ReportName *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	// The report type.
	//
	// example:
	//
	// report
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The ID of the workspace to which the report belongs.
	//
	// example:
	//
	// ab46ed33-6278-4ef7-8013-8c1335f266ee
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// Custom template
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryReportPerformanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryReportPerformanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryReportPerformanceResponseBodyResult) GetCacheCostTimeAvg() *float64 {
	return s.CacheCostTimeAvg
}

func (s *QueryReportPerformanceResponseBodyResult) GetCacheQueryCount() *int32 {
	return s.CacheQueryCount
}

func (s *QueryReportPerformanceResponseBodyResult) GetComponentQueryCount() *int32 {
	return s.ComponentQueryCount
}

func (s *QueryReportPerformanceResponseBodyResult) GetComponentQueryCountAvg() *float64 {
	return s.ComponentQueryCountAvg
}

func (s *QueryReportPerformanceResponseBodyResult) GetCostTimeAvg() *float64 {
	return s.CostTimeAvg
}

func (s *QueryReportPerformanceResponseBodyResult) GetQueryCount() *int32 {
	return s.QueryCount
}

func (s *QueryReportPerformanceResponseBodyResult) GetQueryCountAvg() *float64 {
	return s.QueryCountAvg
}

func (s *QueryReportPerformanceResponseBodyResult) GetQueryOverFivePercentNum() *float64 {
	return s.QueryOverFivePercentNum
}

func (s *QueryReportPerformanceResponseBodyResult) GetQueryOverFiveSecPercent() *string {
	return s.QueryOverFiveSecPercent
}

func (s *QueryReportPerformanceResponseBodyResult) GetQueryOverTenSecPercent() *string {
	return s.QueryOverTenSecPercent
}

func (s *QueryReportPerformanceResponseBodyResult) GetQueryOverTenSecPercentNum() *float64 {
	return s.QueryOverTenSecPercentNum
}

func (s *QueryReportPerformanceResponseBodyResult) GetQueryTimeoutCount() *int32 {
	return s.QueryTimeoutCount
}

func (s *QueryReportPerformanceResponseBodyResult) GetQueryTimeoutCountPercent() *float64 {
	return s.QueryTimeoutCountPercent
}

func (s *QueryReportPerformanceResponseBodyResult) GetQuickIndexCostTimeAvg() *float64 {
	return s.QuickIndexCostTimeAvg
}

func (s *QueryReportPerformanceResponseBodyResult) GetQuickIndexQueryCount() *int32 {
	return s.QuickIndexQueryCount
}

func (s *QueryReportPerformanceResponseBodyResult) GetRepeatQueryPercent() *string {
	return s.RepeatQueryPercent
}

func (s *QueryReportPerformanceResponseBodyResult) GetRepeatQueryPercentNum() *float64 {
	return s.RepeatQueryPercentNum
}

func (s *QueryReportPerformanceResponseBodyResult) GetRepeatSqlQueryCount() *int32 {
	return s.RepeatSqlQueryCount
}

func (s *QueryReportPerformanceResponseBodyResult) GetRepeatSqlQueryPercent() *string {
	return s.RepeatSqlQueryPercent
}

func (s *QueryReportPerformanceResponseBodyResult) GetReportId() *string {
	return s.ReportId
}

func (s *QueryReportPerformanceResponseBodyResult) GetReportName() *string {
	return s.ReportName
}

func (s *QueryReportPerformanceResponseBodyResult) GetReportType() *string {
	return s.ReportType
}

func (s *QueryReportPerformanceResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryReportPerformanceResponseBodyResult) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryReportPerformanceResponseBodyResult) SetCacheCostTimeAvg(v float64) *QueryReportPerformanceResponseBodyResult {
	s.CacheCostTimeAvg = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetCacheQueryCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.CacheQueryCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetComponentQueryCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.ComponentQueryCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetComponentQueryCountAvg(v float64) *QueryReportPerformanceResponseBodyResult {
	s.ComponentQueryCountAvg = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetCostTimeAvg(v float64) *QueryReportPerformanceResponseBodyResult {
	s.CostTimeAvg = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.QueryCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryCountAvg(v float64) *QueryReportPerformanceResponseBodyResult {
	s.QueryCountAvg = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryOverFivePercentNum(v float64) *QueryReportPerformanceResponseBodyResult {
	s.QueryOverFivePercentNum = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryOverFiveSecPercent(v string) *QueryReportPerformanceResponseBodyResult {
	s.QueryOverFiveSecPercent = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryOverTenSecPercent(v string) *QueryReportPerformanceResponseBodyResult {
	s.QueryOverTenSecPercent = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryOverTenSecPercentNum(v float64) *QueryReportPerformanceResponseBodyResult {
	s.QueryOverTenSecPercentNum = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryTimeoutCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.QueryTimeoutCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryTimeoutCountPercent(v float64) *QueryReportPerformanceResponseBodyResult {
	s.QueryTimeoutCountPercent = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQuickIndexCostTimeAvg(v float64) *QueryReportPerformanceResponseBodyResult {
	s.QuickIndexCostTimeAvg = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQuickIndexQueryCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.QuickIndexQueryCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetRepeatQueryPercent(v string) *QueryReportPerformanceResponseBodyResult {
	s.RepeatQueryPercent = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetRepeatQueryPercentNum(v float64) *QueryReportPerformanceResponseBodyResult {
	s.RepeatQueryPercentNum = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetRepeatSqlQueryCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.RepeatSqlQueryCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetRepeatSqlQueryPercent(v string) *QueryReportPerformanceResponseBodyResult {
	s.RepeatSqlQueryPercent = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetReportId(v string) *QueryReportPerformanceResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetReportName(v string) *QueryReportPerformanceResponseBodyResult {
	s.ReportName = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetReportType(v string) *QueryReportPerformanceResponseBodyResult {
	s.ReportType = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetWorkspaceId(v string) *QueryReportPerformanceResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetWorkspaceName(v string) *QueryReportPerformanceResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
