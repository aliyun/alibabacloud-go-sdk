// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCapacityPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComplexQueryAvailable(v bool) *CapacityPlanRequest
	GetComplexQueryAvailable() *bool
	SetDataInfo(v []*CapacityPlanRequestDataInfo) *CapacityPlanRequest
	GetDataInfo() []*CapacityPlanRequestDataInfo
	SetMetric(v []*CapacityPlanRequestMetric) *CapacityPlanRequest
	GetMetric() []*CapacityPlanRequestMetric
	SetUsageScenario(v string) *CapacityPlanRequest
	GetUsageScenario() *string
}

type CapacityPlanRequest struct {
	// Specifies whether complex aggregate query is required. Valid values:
	//
	// - true: Required.
	//
	// - false (default): Not required.
	//
	// example:
	//
	// true
	ComplexQueryAvailable *bool `json:"complexQueryAvailable,omitempty" xml:"complexQueryAvailable,omitempty"`
	// The disk usage information.
	DataInfo []*CapacityPlanRequestDataInfo `json:"dataInfo,omitempty" xml:"dataInfo,omitempty" type:"Repeated"`
	// The metric information, including disk usage, search and write operations, and aggregation requests.
	Metric []*CapacityPlanRequestMetric `json:"metric,omitempty" xml:"metric,omitempty" type:"Repeated"`
	// Scenarios. Valid values:
	//
	// - general: general-purpose scenario
	//
	// - analysisVisualization: data analytics scenario
	//
	// - dbAcceleration: database acceleration scenario
	//
	// - search: search scenario
	//
	// - log: log scenario.
	//
	// example:
	//
	// general
	UsageScenario *string `json:"usageScenario,omitempty" xml:"usageScenario,omitempty"`
}

func (s CapacityPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CapacityPlanRequest) GoString() string {
	return s.String()
}

func (s *CapacityPlanRequest) GetComplexQueryAvailable() *bool {
	return s.ComplexQueryAvailable
}

func (s *CapacityPlanRequest) GetDataInfo() []*CapacityPlanRequestDataInfo {
	return s.DataInfo
}

func (s *CapacityPlanRequest) GetMetric() []*CapacityPlanRequestMetric {
	return s.Metric
}

func (s *CapacityPlanRequest) GetUsageScenario() *string {
	return s.UsageScenario
}

func (s *CapacityPlanRequest) SetComplexQueryAvailable(v bool) *CapacityPlanRequest {
	s.ComplexQueryAvailable = &v
	return s
}

func (s *CapacityPlanRequest) SetDataInfo(v []*CapacityPlanRequestDataInfo) *CapacityPlanRequest {
	s.DataInfo = v
	return s
}

func (s *CapacityPlanRequest) SetMetric(v []*CapacityPlanRequestMetric) *CapacityPlanRequest {
	s.Metric = v
	return s
}

func (s *CapacityPlanRequest) SetUsageScenario(v string) *CapacityPlanRequest {
	s.UsageScenario = &v
	return s
}

func (s *CapacityPlanRequest) Validate() error {
	if s.DataInfo != nil {
		for _, item := range s.DataInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Metric != nil {
		for _, item := range s.Metric {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CapacityPlanRequestDataInfo struct {
	// The disk data metric code. Valid values:
	//
	// - totalRawData: source data information
	//
	// - document: data document information, estimated number of documents
	//
	// - dailyIncrement: daily data growth
	//
	// - dailyIncrement: daily incremental documents
	//
	// - retentionTime: data retention period
	//
	// - replica: replica settings.
	//
	// example:
	//
	// totalRawData
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The metric value of disk usage.
	//
	// example:
	//
	// 100
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The total number of data entries.
	//
	// example:
	//
	// 10000
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The disk data type. Valid values:
	//
	// - hot: hot data
	//
	// - warm: warm data.
	//
	// example:
	//
	// hot
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The data unit or time unit. Valid values:
	//
	// - Data units: MiB, GiB, TB, PB
	//
	// - Time units: DAYS, WEEKS, MONTHS, YEARS.
	//
	// example:
	//
	// MiB
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s CapacityPlanRequestDataInfo) String() string {
	return dara.Prettify(s)
}

func (s CapacityPlanRequestDataInfo) GoString() string {
	return s.String()
}

func (s *CapacityPlanRequestDataInfo) GetCode() *string {
	return s.Code
}

func (s *CapacityPlanRequestDataInfo) GetSize() *int64 {
	return s.Size
}

func (s *CapacityPlanRequestDataInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *CapacityPlanRequestDataInfo) GetType() *string {
	return s.Type
}

func (s *CapacityPlanRequestDataInfo) GetUnit() *string {
	return s.Unit
}

func (s *CapacityPlanRequestDataInfo) SetCode(v string) *CapacityPlanRequestDataInfo {
	s.Code = &v
	return s
}

func (s *CapacityPlanRequestDataInfo) SetSize(v int64) *CapacityPlanRequestDataInfo {
	s.Size = &v
	return s
}

func (s *CapacityPlanRequestDataInfo) SetTotalCount(v int32) *CapacityPlanRequestDataInfo {
	s.TotalCount = &v
	return s
}

func (s *CapacityPlanRequestDataInfo) SetType(v string) *CapacityPlanRequestDataInfo {
	s.Type = &v
	return s
}

func (s *CapacityPlanRequestDataInfo) SetUnit(v string) *CapacityPlanRequestDataInfo {
	s.Unit = &v
	return s
}

func (s *CapacityPlanRequestDataInfo) Validate() error {
	return dara.Validate(s)
}

type CapacityPlanRequestMetric struct {
	// The average QPS.
	//
	// example:
	//
	// 30
	AverageQps *int32 `json:"averageQps,omitempty" xml:"averageQps,omitempty"`
	// The search or write metric code. Valid values:
	//
	// - write: write
	//
	// - search: search.
	//
	// example:
	//
	// write
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The number of concurrent connections.
	//
	// example:
	//
	// 2
	Concurrent *int64 `json:"concurrent,omitempty" xml:"concurrent,omitempty"`
	// The peak QPS.
	//
	// example:
	//
	// 30
	PeakQps *int32 `json:"peakQps,omitempty" xml:"peakQps,omitempty"`
	// The expected average response time, in milliseconds.
	//
	// example:
	//
	// 100
	ResponseTime *int32 `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
	// The throughput, in MB/s.
	//
	// example:
	//
	// 100
	Throughput *int64 `json:"throughput,omitempty" xml:"throughput,omitempty"`
	// The search or write peak type. Valid values:
	//
	// - common: normal
	//
	// - peak: peak.
	//
	// example:
	//
	// common
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CapacityPlanRequestMetric) String() string {
	return dara.Prettify(s)
}

func (s CapacityPlanRequestMetric) GoString() string {
	return s.String()
}

func (s *CapacityPlanRequestMetric) GetAverageQps() *int32 {
	return s.AverageQps
}

func (s *CapacityPlanRequestMetric) GetCode() *string {
	return s.Code
}

func (s *CapacityPlanRequestMetric) GetConcurrent() *int64 {
	return s.Concurrent
}

func (s *CapacityPlanRequestMetric) GetPeakQps() *int32 {
	return s.PeakQps
}

func (s *CapacityPlanRequestMetric) GetResponseTime() *int32 {
	return s.ResponseTime
}

func (s *CapacityPlanRequestMetric) GetThroughput() *int64 {
	return s.Throughput
}

func (s *CapacityPlanRequestMetric) GetType() *string {
	return s.Type
}

func (s *CapacityPlanRequestMetric) SetAverageQps(v int32) *CapacityPlanRequestMetric {
	s.AverageQps = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetCode(v string) *CapacityPlanRequestMetric {
	s.Code = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetConcurrent(v int64) *CapacityPlanRequestMetric {
	s.Concurrent = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetPeakQps(v int32) *CapacityPlanRequestMetric {
	s.PeakQps = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetResponseTime(v int32) *CapacityPlanRequestMetric {
	s.ResponseTime = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetThroughput(v int64) *CapacityPlanRequestMetric {
	s.Throughput = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetType(v string) *CapacityPlanRequestMetric {
	s.Type = &v
	return s
}

func (s *CapacityPlanRequestMetric) Validate() error {
	return dara.Validate(s)
}
