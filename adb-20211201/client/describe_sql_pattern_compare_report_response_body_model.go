// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlPatternCompareReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeSqlPatternCompareReportResponseBodyItems) *DescribeSqlPatternCompareReportResponseBody
	GetItems() []*DescribeSqlPatternCompareReportResponseBodyItems
	SetMetricType(v string) *DescribeSqlPatternCompareReportResponseBody
	GetMetricType() *string
	SetPageNumber(v int32) *DescribeSqlPatternCompareReportResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSqlPatternCompareReportResponseBody
	GetPageSize() *int32
	SetReportId(v int64) *DescribeSqlPatternCompareReportResponseBody
	GetReportId() *int64
	SetRequestId(v string) *DescribeSqlPatternCompareReportResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSqlPatternCompareReportResponseBody
	GetTotalCount() *int32
}

type DescribeSqlPatternCompareReportResponseBody struct {
	// The Pattern details on the current page. An empty array is returned if no results match the conditions.
	Items []*DescribeSqlPatternCompareReportResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The analysis metric. Valid values:
	//
	// - `QUERY_COUNT`: the number of query executions.
	//
	// - `CPU_COST`: the CPU consumption.
	//
	// - `SHUFFLE_SIZE`: the amount of shuffle data.
	//
	// - `PEAK_MEMORY`: the peak memory consumption.
	//
	// - `SCAN_SIZE`: the amount of scanned data.
	//
	// example:
	//
	// CPU_COST
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The page number of the returned page, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries returned per page for this query.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the SQL Pattern comparison report.
	//
	// example:
	//
	// 1001
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A1B2C3D-4E5F-6789-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of Patterns that match the current report, analysis dimension, and change rate filter conditions. This is not the number of entries on the current page.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSqlPatternCompareReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternCompareReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternCompareReportResponseBody) GetItems() []*DescribeSqlPatternCompareReportResponseBodyItems {
	return s.Items
}

func (s *DescribeSqlPatternCompareReportResponseBody) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeSqlPatternCompareReportResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSqlPatternCompareReportResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSqlPatternCompareReportResponseBody) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeSqlPatternCompareReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlPatternCompareReportResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSqlPatternCompareReportResponseBody) SetItems(v []*DescribeSqlPatternCompareReportResponseBodyItems) *DescribeSqlPatternCompareReportResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBody) SetMetricType(v string) *DescribeSqlPatternCompareReportResponseBody {
	s.MetricType = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBody) SetPageNumber(v int32) *DescribeSqlPatternCompareReportResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBody) SetPageSize(v int32) *DescribeSqlPatternCompareReportResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBody) SetReportId(v int64) *DescribeSqlPatternCompareReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBody) SetRequestId(v string) *DescribeSqlPatternCompareReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBody) SetTotalCount(v int32) *DescribeSqlPatternCompareReportResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSqlPatternCompareReportResponseBodyItems struct {
	// The display string of the average execution duration for Time 2, in seconds. This field is returned only for the CPU_COST dimension.
	//
	// example:
	//
	// 0.4s
	AvgExecutionTime *string `json:"AvgExecutionTime,omitempty" xml:"AvgExecutionTime,omitempty"`
	// The display string of the average planning duration for Time 2, in seconds. This field is returned only for the CPU_COST dimension.
	//
	// example:
	//
	// 0.1s
	AvgPlanningTime *string `json:"AvgPlanningTime,omitempty" xml:"AvgPlanningTime,omitempty"`
	// The display string of the average query response time for Time 2, in seconds. This field is returned for all analysis dimensions.
	//
	// example:
	//
	// 0.5s
	AvgRt *string `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	// The display string of the maximum execution duration for Time 2, in seconds. This field is returned only for the CPU_COST dimension.
	//
	// example:
	//
	// 1.8s
	MaxExecutionTime *string `json:"MaxExecutionTime,omitempty" xml:"MaxExecutionTime,omitempty"`
	// The display string of the maximum planning duration for Time 2, in seconds. This field is returned only for the CPU_COST dimension.
	//
	// example:
	//
	// 0.2s
	MaxPlanningTime *string `json:"MaxPlanningTime,omitempty" xml:"MaxPlanningTime,omitempty"`
	// The display string of the maximum query response time for Time 2, in seconds. This field is returned for all analysis dimensions.
	//
	// example:
	//
	// 2s
	MaxRt *string `json:"MaxRt,omitempty" xml:"MaxRt,omitempty"`
	// The primary metric mapping for the current analysis dimension. Valid keys:
	//
	// - `QUERY_COUNT`: the number of query executions.
	//
	// - `CPU_COST`: the CPU consumption.
	//
	// - `SHUFFLE_SIZE`: the shuffle data volume.
	//
	// - `PEAK_MEMORY`: the peak memory consumption.
	//
	// - `SCAN_SIZE`: the scan data volume.
	//
	// > Each result contains only one key that matches the `MetricType` request parameter.
	MetricValues map[string]*ItemsMetricValuesValue `json:"MetricValues,omitempty" xml:"MetricValues,omitempty"`
	// The parameterized SQL Pattern text. This field is empty or not returned when IncludePattern is set to false. When the text is unavailable, a prompt containing a hash identifier may be returned.
	//
	// example:
	//
	// SELECT 	- FROM orders WHERE order_id = ?
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The number of query executions for Time 2, in count. This field is returned for the CPU_COST, SHUFFLE_SIZE, PEAK_MEMORY, and SCAN_SIZE dimensions.
	//
	// example:
	//
	// 120
	QueryCount *int64 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The display string of the number of query executions for Time 2. The applicable scope is the same as QueryCount.
	//
	// example:
	//
	// 120 times
	QueryCountDisplayValue *string `json:"QueryCountDisplayValue,omitempty" xml:"QueryCountDisplayValue,omitempty"`
	// The global sequence number in the current filtered and sorted results, starting from 1 and numbered continuously across pages.
	//
	// example:
	//
	// 1
	Rank *int64 `json:"Rank,omitempty" xml:"Rank,omitempty"`
	// The change level for the current analysis dimension. Valid values:
	//
	// - `NEW`: A new Pattern. Returned only for NEW reports.
	//
	// - `SLIGHT`: A slight change. The average change rate is in the range of (0%, 20%].
	//
	// - `MODERATE`: A moderate change. The average change rate is in the range of (20%, 50%].
	//
	// - `HIGH`: A high change. The average change rate is in the range of (50%, 100%].
	//
	// - `SEVERE`: A severe change. The average change rate is greater than 100%, or the change represents zero-baseline growth.
	//
	// > The change level only indicates the magnitude of metric growth and cannot be used alone to determine the cause of a fault.
	//
	// example:
	//
	// SEVERE
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The hash identifier of the SQL Pattern, returned as a string. Store and pass this value as a string to avoid precision loss caused by numeric conversion.
	//
	// example:
	//
	// 1234567890123456789
	SqlPatternHash *string `json:"SqlPatternHash,omitempty" xml:"SqlPatternHash,omitempty"`
	// The display string of the total query duration for Time 2, in seconds. This field is returned only for the QUERY_COUNT dimension.
	//
	// example:
	//
	// 60s
	TotalQueryTime *string `json:"TotalQueryTime,omitempty" xml:"TotalQueryTime,omitempty"`
	// The display string of the total scan duration for Time 2, in seconds. This field is returned only for the SCAN_SIZE dimension.
	//
	// example:
	//
	// 12s
	TotalScanCost *string `json:"TotalScanCost,omitempty" xml:"TotalScanCost,omitempty"`
}

func (s DescribeSqlPatternCompareReportResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternCompareReportResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetAvgExecutionTime() *string {
	return s.AvgExecutionTime
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetAvgPlanningTime() *string {
	return s.AvgPlanningTime
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetAvgRt() *string {
	return s.AvgRt
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetMaxExecutionTime() *string {
	return s.MaxExecutionTime
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetMaxPlanningTime() *string {
	return s.MaxPlanningTime
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetMaxRt() *string {
	return s.MaxRt
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetMetricValues() map[string]*ItemsMetricValuesValue {
	return s.MetricValues
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetPattern() *string {
	return s.Pattern
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetQueryCount() *int64 {
	return s.QueryCount
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetQueryCountDisplayValue() *string {
	return s.QueryCountDisplayValue
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetRank() *int64 {
	return s.Rank
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetSqlPatternHash() *string {
	return s.SqlPatternHash
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetTotalQueryTime() *string {
	return s.TotalQueryTime
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) GetTotalScanCost() *string {
	return s.TotalScanCost
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetAvgExecutionTime(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.AvgExecutionTime = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetAvgPlanningTime(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.AvgPlanningTime = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetAvgRt(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.AvgRt = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetMaxExecutionTime(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.MaxExecutionTime = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetMaxPlanningTime(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.MaxPlanningTime = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetMaxRt(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.MaxRt = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetMetricValues(v map[string]*ItemsMetricValuesValue) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.MetricValues = v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetPattern(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.Pattern = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetQueryCount(v int64) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.QueryCount = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetQueryCountDisplayValue(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.QueryCountDisplayValue = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetRank(v int64) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.Rank = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetRiskLevel(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.RiskLevel = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetSqlPatternHash(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.SqlPatternHash = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetTotalQueryTime(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.TotalQueryTime = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) SetTotalScanCost(v string) *DescribeSqlPatternCompareReportResponseBodyItems {
	s.TotalScanCost = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
