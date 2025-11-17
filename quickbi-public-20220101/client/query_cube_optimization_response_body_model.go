// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCubeOptimizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryCubeOptimizationResponseBody
	GetRequestId() *string
	SetResult(v []*QueryCubeOptimizationResponseBodyResult) *QueryCubeOptimizationResponseBody
	GetResult() []*QueryCubeOptimizationResponseBodyResult
	SetSuccess(v bool) *QueryCubeOptimizationResponseBody
	GetSuccess() *bool
}

type QueryCubeOptimizationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The query results are returned.
	Result []*QueryCubeOptimizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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

func (s QueryCubeOptimizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCubeOptimizationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCubeOptimizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCubeOptimizationResponseBody) GetResult() []*QueryCubeOptimizationResponseBodyResult {
	return s.Result
}

func (s *QueryCubeOptimizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCubeOptimizationResponseBody) SetRequestId(v string) *QueryCubeOptimizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCubeOptimizationResponseBody) SetResult(v []*QueryCubeOptimizationResponseBodyResult) *QueryCubeOptimizationResponseBody {
	s.Result = v
	return s
}

func (s *QueryCubeOptimizationResponseBody) SetSuccess(v bool) *QueryCubeOptimizationResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCubeOptimizationResponseBody) Validate() error {
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

type QueryCubeOptimizationResponseBodyResult struct {
	// The type of the suggestion. Valid values:
	//
	// 	- **OPEN_CACHE**: Open cache.
	//
	// 	- **OPEN_QUICK_ENGINE**: Open FAST Cache.
	//
	// 	- **INCREASE_CACHE_TIME**: Increase the cache time.
	//
	// example:
	//
	// OPENQUICKENGINE
	AdviceType *string `json:"AdviceType,omitempty" xml:"AdviceType,omitempty"`
	// The diagnostic information about the dataset.
	CubePerformanceDiagnoseModel *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel `json:"CubePerformanceDiagnoseModel,omitempty" xml:"CubePerformanceDiagnoseModel,omitempty" type:"Struct"`
}

func (s QueryCubeOptimizationResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryCubeOptimizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCubeOptimizationResponseBodyResult) GetAdviceType() *string {
	return s.AdviceType
}

func (s *QueryCubeOptimizationResponseBodyResult) GetCubePerformanceDiagnoseModel() *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	return s.CubePerformanceDiagnoseModel
}

func (s *QueryCubeOptimizationResponseBodyResult) SetAdviceType(v string) *QueryCubeOptimizationResponseBodyResult {
	s.AdviceType = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResult) SetCubePerformanceDiagnoseModel(v *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) *QueryCubeOptimizationResponseBodyResult {
	s.CubePerformanceDiagnoseModel = v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResult) Validate() error {
	if s.CubePerformanceDiagnoseModel != nil {
		if err := s.CubePerformanceDiagnoseModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel struct {
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
	// 2
	CacheQueryCount *int32 `json:"CacheQueryCount,omitempty" xml:"CacheQueryCount,omitempty"`
	// The average query duration associated with the SQL pattern.
	//
	// example:
	//
	// 1.0
	CostTimeAvg *float64 `json:"CostTimeAvg,omitempty" xml:"CostTimeAvg,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// 3e45b61a-9ba8-4c7c-8248-8dbe69945636
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
	// 2
	QueryCountAvg *float64 `json:"QueryCountAvg,omitempty" xml:"QueryCountAvg,omitempty"`
	// The percentage of the number of queries that exceed the 5S.
	//
	// example:
	//
	// 0.1
	QueryOverFivePercentNum *float64 `json:"QueryOverFivePercentNum,omitempty" xml:"QueryOverFivePercentNum,omitempty"`
	// Query the proportion of more than 5S.
	//
	// example:
	//
	// 0.5
	QueryOverFiveSecPercent *string `json:"QueryOverFiveSecPercent,omitempty" xml:"QueryOverFiveSecPercent,omitempty"`
	// The percentage of queries that exceed 10s.
	//
	// example:
	//
	// 0.1
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
	// 1
	QuickIndexCostTimeAvg *float64 `json:"QuickIndexCostTimeAvg,omitempty" xml:"QuickIndexCostTimeAvg,omitempty"`
	// The number of times that the Quick engine is hit.
	//
	// example:
	//
	// 2
	QuickIndexQueryCount *int32 `json:"QuickIndexQueryCount,omitempty" xml:"QuickIndexQueryCount,omitempty"`
	// The proportion of duplicate queries.
	//
	// example:
	//
	// 0.1
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
	// 2
	RepeatSqlQueryCount *int32 `json:"RepeatSqlQueryCount,omitempty" xml:"RepeatSqlQueryCount,omitempty"`
	// The proportion of duplicate queries.
	//
	// example:
	//
	// 0.3
	RepeatSqlQueryPercent *string `json:"RepeatSqlQueryPercent,omitempty" xml:"RepeatSqlQueryPercent,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 6ea74bff-c818-4188-b462-dbb45a24dbac
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// eco0sh0prods
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) String() string {
	return dara.Prettify(s)
}

func (s QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GoString() string {
	return s.String()
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetCacheCostTimeAvg() *float64 {
	return s.CacheCostTimeAvg
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetCacheQueryCount() *int32 {
	return s.CacheQueryCount
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetCostTimeAvg() *float64 {
	return s.CostTimeAvg
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetCubeId() *string {
	return s.CubeId
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetCubeName() *string {
	return s.CubeName
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetQueryCount() *int32 {
	return s.QueryCount
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetQueryCountAvg() *float64 {
	return s.QueryCountAvg
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetQueryOverFivePercentNum() *float64 {
	return s.QueryOverFivePercentNum
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetQueryOverFiveSecPercent() *string {
	return s.QueryOverFiveSecPercent
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetQueryOverTenSecPercent() *string {
	return s.QueryOverTenSecPercent
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetQueryOverTenSecPercentNum() *float64 {
	return s.QueryOverTenSecPercentNum
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetQueryTimeoutCount() *int32 {
	return s.QueryTimeoutCount
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetQueryTimeoutCountPercent() *float64 {
	return s.QueryTimeoutCountPercent
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetQuickIndexCostTimeAvg() *float64 {
	return s.QuickIndexCostTimeAvg
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetQuickIndexQueryCount() *int32 {
	return s.QuickIndexQueryCount
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetRepeatQueryPercent() *string {
	return s.RepeatQueryPercent
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetRepeatQueryPercentNum() *float64 {
	return s.RepeatQueryPercentNum
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetRepeatSqlQueryCount() *int32 {
	return s.RepeatSqlQueryCount
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetRepeatSqlQueryPercent() *string {
	return s.RepeatSqlQueryPercent
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetCacheCostTimeAvg(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.CacheCostTimeAvg = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetCacheQueryCount(v int32) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.CacheQueryCount = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetCostTimeAvg(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.CostTimeAvg = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetCubeId(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.CubeId = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetCubeName(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.CubeName = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryCount(v int32) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryCount = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryCountAvg(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryCountAvg = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryOverFivePercentNum(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryOverFivePercentNum = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryOverFiveSecPercent(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryOverFiveSecPercent = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryOverTenSecPercent(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryOverTenSecPercent = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryOverTenSecPercentNum(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryOverTenSecPercentNum = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryTimeoutCount(v int32) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryTimeoutCount = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryTimeoutCountPercent(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryTimeoutCountPercent = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQuickIndexCostTimeAvg(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QuickIndexCostTimeAvg = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQuickIndexQueryCount(v int32) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QuickIndexQueryCount = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetRepeatQueryPercent(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.RepeatQueryPercent = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetRepeatQueryPercentNum(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.RepeatQueryPercentNum = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetRepeatSqlQueryCount(v int32) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.RepeatSqlQueryCount = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetRepeatSqlQueryPercent(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.RepeatSqlQueryPercent = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetWorkspaceId(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.WorkspaceId = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetWorkspaceName(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.WorkspaceName = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) Validate() error {
	return dara.Validate(s)
}
