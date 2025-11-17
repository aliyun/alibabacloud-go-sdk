// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryComponentPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryComponentPerformanceResponseBody
	GetRequestId() *string
	SetResult(v []*QueryComponentPerformanceResponseBodyResult) *QueryComponentPerformanceResponseBody
	GetResult() []*QueryComponentPerformanceResponseBodyResult
	SetSuccess(v bool) *QueryComponentPerformanceResponseBody
	GetSuccess() *bool
}

type QueryComponentPerformanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BCE45E6D-9304-4F94-86BB-5A772B1615FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result returned.
	Result []*QueryComponentPerformanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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

func (s QueryComponentPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryComponentPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryComponentPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryComponentPerformanceResponseBody) GetResult() []*QueryComponentPerformanceResponseBodyResult {
	return s.Result
}

func (s *QueryComponentPerformanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryComponentPerformanceResponseBody) SetRequestId(v string) *QueryComponentPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryComponentPerformanceResponseBody) SetResult(v []*QueryComponentPerformanceResponseBodyResult) *QueryComponentPerformanceResponseBody {
	s.Result = v
	return s
}

func (s *QueryComponentPerformanceResponseBody) SetSuccess(v bool) *QueryComponentPerformanceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryComponentPerformanceResponseBody) Validate() error {
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

type QueryComponentPerformanceResponseBodyResult struct {
	// The average duration of cache hits.
	//
	// example:
	//
	// 0.3
	CacheCostTimeAvg *float64 `json:"CacheCostTimeAvg,omitempty" xml:"CacheCostTimeAvg,omitempty"`
	// The number of cache hits.
	//
	// example:
	//
	// 3
	CacheQueryCount *int32 `json:"CacheQueryCount,omitempty" xml:"CacheQueryCount,omitempty"`
	// The component ID.
	//
	// example:
	//
	// 0696083a-ca72-4d89-8e7a-c017910e0***
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The name of the add-on.
	//
	// example:
	//
	// test
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The average query duration associated with the SQL pattern.
	//
	// example:
	//
	// 0.3
	CostTimeAvg *float64 `json:"CostTimeAvg,omitempty" xml:"CostTimeAvg,omitempty"`
	// The number of queries.
	//
	// example:
	//
	// 5
	QueryCount *int32 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The average number of queries.
	//
	// example:
	//
	// 0.3
	QueryCountAvg *float64 `json:"QueryCountAvg,omitempty" xml:"QueryCountAvg,omitempty"`
	// The query exceeds the 5S number of queries.
	//
	// example:
	//
	// 5
	QueryOverFivePercentNum *float64 `json:"QueryOverFivePercentNum,omitempty" xml:"QueryOverFivePercentNum,omitempty"`
	// Query the proportion of more than 5S.
	//
	// example:
	//
	// 0.3
	QueryOverFiveSecPercent *string `json:"QueryOverFiveSecPercent,omitempty" xml:"QueryOverFiveSecPercent,omitempty"`
	// The percentage of queries that exceed 10s.
	//
	// example:
	//
	// 0.3
	QueryOverTenSecPercent *string `json:"QueryOverTenSecPercent,omitempty" xml:"QueryOverTenSecPercent,omitempty"`
	// The percentage of queries that exceed 10s.
	//
	// example:
	//
	// 0.3
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
	// 0.3
	QueryTimeoutCountPercent *float64 `json:"QueryTimeoutCountPercent,omitempty" xml:"QueryTimeoutCountPercent,omitempty"`
	// The average time consumed by the Quick engine query.
	//
	// example:
	//
	// 0.3
	QuickIndexCostTimeAvg *float64 `json:"QuickIndexCostTimeAvg,omitempty" xml:"QuickIndexCostTimeAvg,omitempty"`
	// The number of times that the Quick engine is hit.
	//
	// example:
	//
	// 3
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
	// 2
	RepeatQueryPercentNum *float64 `json:"RepeatQueryPercentNum,omitempty" xml:"RepeatQueryPercentNum,omitempty"`
	// The number of times the query is repeated.
	//
	// example:
	//
	// 5
	RepeatSqlQueryCount *int32 `json:"RepeatSqlQueryCount,omitempty" xml:"RepeatSqlQueryCount,omitempty"`
	// The proportion of duplicate queries.
	//
	// example:
	//
	// 0.3
	RepeatSqlQueryPercent *string `json:"RepeatSqlQueryPercent,omitempty" xml:"RepeatSqlQueryPercent,omitempty"`
	// The ID of the work.
	//
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The name of the report.
	//
	// example:
	//
	// ClusterRiskReport
	ReportName *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	// The format of the report.
	//
	// example:
	//
	// report
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The unique ID of the space.
	//
	// example:
	//
	// 89713491-cb4f-4579-b889-e82c35f1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// test
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryComponentPerformanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryComponentPerformanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryComponentPerformanceResponseBodyResult) GetCacheCostTimeAvg() *float64 {
	return s.CacheCostTimeAvg
}

func (s *QueryComponentPerformanceResponseBodyResult) GetCacheQueryCount() *int32 {
	return s.CacheQueryCount
}

func (s *QueryComponentPerformanceResponseBodyResult) GetComponentId() *string {
	return s.ComponentId
}

func (s *QueryComponentPerformanceResponseBodyResult) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryComponentPerformanceResponseBodyResult) GetCostTimeAvg() *float64 {
	return s.CostTimeAvg
}

func (s *QueryComponentPerformanceResponseBodyResult) GetQueryCount() *int32 {
	return s.QueryCount
}

func (s *QueryComponentPerformanceResponseBodyResult) GetQueryCountAvg() *float64 {
	return s.QueryCountAvg
}

func (s *QueryComponentPerformanceResponseBodyResult) GetQueryOverFivePercentNum() *float64 {
	return s.QueryOverFivePercentNum
}

func (s *QueryComponentPerformanceResponseBodyResult) GetQueryOverFiveSecPercent() *string {
	return s.QueryOverFiveSecPercent
}

func (s *QueryComponentPerformanceResponseBodyResult) GetQueryOverTenSecPercent() *string {
	return s.QueryOverTenSecPercent
}

func (s *QueryComponentPerformanceResponseBodyResult) GetQueryOverTenSecPercentNum() *float64 {
	return s.QueryOverTenSecPercentNum
}

func (s *QueryComponentPerformanceResponseBodyResult) GetQueryTimeoutCount() *int32 {
	return s.QueryTimeoutCount
}

func (s *QueryComponentPerformanceResponseBodyResult) GetQueryTimeoutCountPercent() *float64 {
	return s.QueryTimeoutCountPercent
}

func (s *QueryComponentPerformanceResponseBodyResult) GetQuickIndexCostTimeAvg() *float64 {
	return s.QuickIndexCostTimeAvg
}

func (s *QueryComponentPerformanceResponseBodyResult) GetQuickIndexQueryCount() *int32 {
	return s.QuickIndexQueryCount
}

func (s *QueryComponentPerformanceResponseBodyResult) GetRepeatQueryPercent() *string {
	return s.RepeatQueryPercent
}

func (s *QueryComponentPerformanceResponseBodyResult) GetRepeatQueryPercentNum() *float64 {
	return s.RepeatQueryPercentNum
}

func (s *QueryComponentPerformanceResponseBodyResult) GetRepeatSqlQueryCount() *int32 {
	return s.RepeatSqlQueryCount
}

func (s *QueryComponentPerformanceResponseBodyResult) GetRepeatSqlQueryPercent() *string {
	return s.RepeatSqlQueryPercent
}

func (s *QueryComponentPerformanceResponseBodyResult) GetReportId() *string {
	return s.ReportId
}

func (s *QueryComponentPerformanceResponseBodyResult) GetReportName() *string {
	return s.ReportName
}

func (s *QueryComponentPerformanceResponseBodyResult) GetReportType() *string {
	return s.ReportType
}

func (s *QueryComponentPerformanceResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryComponentPerformanceResponseBodyResult) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryComponentPerformanceResponseBodyResult) SetCacheCostTimeAvg(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.CacheCostTimeAvg = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetCacheQueryCount(v int32) *QueryComponentPerformanceResponseBodyResult {
	s.CacheQueryCount = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetComponentId(v string) *QueryComponentPerformanceResponseBodyResult {
	s.ComponentId = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetComponentName(v string) *QueryComponentPerformanceResponseBodyResult {
	s.ComponentName = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetCostTimeAvg(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.CostTimeAvg = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryCount(v int32) *QueryComponentPerformanceResponseBodyResult {
	s.QueryCount = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryCountAvg(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.QueryCountAvg = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryOverFivePercentNum(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.QueryOverFivePercentNum = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryOverFiveSecPercent(v string) *QueryComponentPerformanceResponseBodyResult {
	s.QueryOverFiveSecPercent = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryOverTenSecPercent(v string) *QueryComponentPerformanceResponseBodyResult {
	s.QueryOverTenSecPercent = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryOverTenSecPercentNum(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.QueryOverTenSecPercentNum = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryTimeoutCount(v int32) *QueryComponentPerformanceResponseBodyResult {
	s.QueryTimeoutCount = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryTimeoutCountPercent(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.QueryTimeoutCountPercent = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQuickIndexCostTimeAvg(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.QuickIndexCostTimeAvg = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQuickIndexQueryCount(v int32) *QueryComponentPerformanceResponseBodyResult {
	s.QuickIndexQueryCount = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetRepeatQueryPercent(v string) *QueryComponentPerformanceResponseBodyResult {
	s.RepeatQueryPercent = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetRepeatQueryPercentNum(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.RepeatQueryPercentNum = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetRepeatSqlQueryCount(v int32) *QueryComponentPerformanceResponseBodyResult {
	s.RepeatSqlQueryCount = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetRepeatSqlQueryPercent(v string) *QueryComponentPerformanceResponseBodyResult {
	s.RepeatSqlQueryPercent = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetReportId(v string) *QueryComponentPerformanceResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetReportName(v string) *QueryComponentPerformanceResponseBodyResult {
	s.ReportName = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetReportType(v string) *QueryComponentPerformanceResponseBodyResult {
	s.ReportType = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetWorkspaceId(v string) *QueryComponentPerformanceResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetWorkspaceName(v string) *QueryComponentPerformanceResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
