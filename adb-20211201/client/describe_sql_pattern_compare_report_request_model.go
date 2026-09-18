// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlPatternCompareReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeRate(v string) *DescribeSqlPatternCompareReportRequest
	GetChangeRate() *string
	SetDBClusterId(v string) *DescribeSqlPatternCompareReportRequest
	GetDBClusterId() *string
	SetIncludePattern(v bool) *DescribeSqlPatternCompareReportRequest
	GetIncludePattern() *bool
	SetMetricType(v string) *DescribeSqlPatternCompareReportRequest
	GetMetricType() *string
	SetOrder(v string) *DescribeSqlPatternCompareReportRequest
	GetOrder() *string
	SetPageNumber(v int32) *DescribeSqlPatternCompareReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSqlPatternCompareReportRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSqlPatternCompareReportRequest
	GetRegionId() *string
	SetReportId(v int64) *DescribeSqlPatternCompareReportRequest
	GetReportId() *int64
}

type DescribeSqlPatternCompareReportRequest struct {
	// The average change rate filter range for CHANGED reports. The format is `left~right`, where values are expressed as percentages and the interval is left-exclusive and right-inclusive. Examples:
	//
	// - `100~500`: greater than 100% and less than or equal to 500%.
	//
	// - `100~`: greater than 100% with no upper limit.
	//
	// > - The left boundary is required and must be no less than 0. The right boundary must be no less than the left boundary.
	//
	// > - This parameter is ignored for NEW reports.
	//
	// > - When the time window 1 metric value is 0 and the time window 2 value is greater than 0, the Pattern is classified as zero-baseline growth and is categorized as `SEVERE` (significant change). To exclude such Patterns, set an upper limit for the change rate.
	//
	// example:
	//
	// 100~500
	ChangeRate *string `json:"ChangeRate,omitempty" xml:"ChangeRate,omitempty"`
	// The ID of the AnalyticDB for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-2ze1234567890****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to return the parameterized SQL Pattern text. Valid values:
	//
	// - `true`: Returns the Pattern text.
	//
	// - `false`: Does not return the Pattern text, which reduces the response size.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	IncludePattern *bool `json:"IncludePattern,omitempty" xml:"IncludePattern,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// CPU_COST
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// Sorts the query results by a specified field. The value is a JSON array string, such as `[{"Field":"Time2SumValue","Type":"Desc"}]`. The array can contain only one object. Parameters:
	//
	// - `Field`: the sort field. This parameter is case-sensitive. Valid values:
	//
	//     - NEW report: `Time2SumValue`, `Time2AvgValue`, `Time2MaxValue`.
	//
	//     - CHANGED report: `AvgChangeRatePercent`, `AvgTime1Value`, `AvgTime2Value`, `SumChangeRatePercent`, `SumTime1Value`, `SumTime2Value`, `MaxChangeRatePercent`, `MaxTime1Value`, `MaxTime2Value`.
	//
	//     - All report types and analysis metrics: `AvgRt`, `MaxRt`.
	//
	//     - `QUERY_COUNT`: `TotalQueryTime`.
	//
	//     - `CPU_COST`: `QueryCount`, `AvgPlanningTime`, `MaxPlanningTime`, `AvgExecutionTime`, `MaxExecutionTime`.
	//
	//     - `SHUFFLE_SIZE`, `PEAK_MEMORY`: `QueryCount`.
	//
	//     - `SCAN_SIZE`: `QueryCount`, `TotalScanCost`.
	//
	// - `Type`: the sort order. This parameter is case-insensitive. Valid values:
	//
	//     - `Asc`: ascending order.
	//
	//     - `Desc`: descending order.
	//
	// > - NEW reports are sorted by `Time2SumValue` in descending order by default.
	//
	// > - CHANGED reports are sorted by `AvgChangeRatePercent` in descending order by default.
	//
	// > - The value of `Field` must be applicable to the current report type and `MetricType`.
	//
	// example:
	//
	// [{"Field":"AvgChangeRatePercent","Type":"Desc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the SQL Pattern comparison report.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DescribeSqlPatternCompareReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternCompareReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternCompareReportRequest) GetChangeRate() *string {
	return s.ChangeRate
}

func (s *DescribeSqlPatternCompareReportRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSqlPatternCompareReportRequest) GetIncludePattern() *bool {
	return s.IncludePattern
}

func (s *DescribeSqlPatternCompareReportRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeSqlPatternCompareReportRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeSqlPatternCompareReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSqlPatternCompareReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSqlPatternCompareReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSqlPatternCompareReportRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeSqlPatternCompareReportRequest) SetChangeRate(v string) *DescribeSqlPatternCompareReportRequest {
	s.ChangeRate = &v
	return s
}

func (s *DescribeSqlPatternCompareReportRequest) SetDBClusterId(v string) *DescribeSqlPatternCompareReportRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSqlPatternCompareReportRequest) SetIncludePattern(v bool) *DescribeSqlPatternCompareReportRequest {
	s.IncludePattern = &v
	return s
}

func (s *DescribeSqlPatternCompareReportRequest) SetMetricType(v string) *DescribeSqlPatternCompareReportRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeSqlPatternCompareReportRequest) SetOrder(v string) *DescribeSqlPatternCompareReportRequest {
	s.Order = &v
	return s
}

func (s *DescribeSqlPatternCompareReportRequest) SetPageNumber(v int32) *DescribeSqlPatternCompareReportRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternCompareReportRequest) SetPageSize(v int32) *DescribeSqlPatternCompareReportRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternCompareReportRequest) SetRegionId(v string) *DescribeSqlPatternCompareReportRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSqlPatternCompareReportRequest) SetReportId(v int64) *DescribeSqlPatternCompareReportRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeSqlPatternCompareReportRequest) Validate() error {
	return dara.Validate(s)
}
