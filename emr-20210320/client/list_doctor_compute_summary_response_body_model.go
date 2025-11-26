// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorComputeSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDoctorComputeSummaryResponseBodyData) *ListDoctorComputeSummaryResponseBody
	GetData() []*ListDoctorComputeSummaryResponseBodyData
	SetMaxResults(v int32) *ListDoctorComputeSummaryResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorComputeSummaryResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDoctorComputeSummaryResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDoctorComputeSummaryResponseBody
	GetTotalCount() *int32
}

type ListDoctorComputeSummaryResponseBody struct {
	// The details of resource usage.
	Data []*ListDoctorComputeSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of entries that are returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDoctorComputeSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBody) GetData() []*ListDoctorComputeSummaryResponseBodyData {
	return s.Data
}

func (s *ListDoctorComputeSummaryResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorComputeSummaryResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorComputeSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoctorComputeSummaryResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoctorComputeSummaryResponseBody) SetData(v []*ListDoctorComputeSummaryResponseBodyData) *ListDoctorComputeSummaryResponseBody {
	s.Data = v
	return s
}

func (s *ListDoctorComputeSummaryResponseBody) SetMaxResults(v int32) *ListDoctorComputeSummaryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBody) SetNextToken(v string) *ListDoctorComputeSummaryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBody) SetRequestId(v string) *ListDoctorComputeSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBody) SetTotalCount(v int32) *ListDoctorComputeSummaryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDoctorComputeSummaryResponseBodyData struct {
	// The resource analysis results.
	Analysis *ListDoctorComputeSummaryResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The name of the resource whose details are obtained based on the value of ComponentTypes. For example, if the value of ComponentTypes is Queue, the value of this parameter is a queue, such as DW.
	//
	// example:
	//
	// DW
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The metric information.
	Metrics *ListDoctorComputeSummaryResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
}

func (s ListDoctorComputeSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBodyData) GetAnalysis() *ListDoctorComputeSummaryResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *ListDoctorComputeSummaryResponseBodyData) GetComponentName() *string {
	return s.ComponentName
}

func (s *ListDoctorComputeSummaryResponseBodyData) GetMetrics() *ListDoctorComputeSummaryResponseBodyDataMetrics {
	return s.Metrics
}

func (s *ListDoctorComputeSummaryResponseBodyData) SetAnalysis(v *ListDoctorComputeSummaryResponseBodyDataAnalysis) *ListDoctorComputeSummaryResponseBodyData {
	s.Analysis = v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyData) SetComponentName(v string) *ListDoctorComputeSummaryResponseBodyData {
	s.ComponentName = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyData) SetMetrics(v *ListDoctorComputeSummaryResponseBodyDataMetrics) *ListDoctorComputeSummaryResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyData) Validate() error {
	if s.Analysis != nil {
		if err := s.Analysis.Validate(); err != nil {
			return err
		}
	}
	if s.Metrics != nil {
		if err := s.Metrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDoctorComputeSummaryResponseBodyDataAnalysis struct {
	// The total number of healthy jobs.
	//
	// example:
	//
	// 3
	HealthyJobCount *int64 `json:"HealthyJobCount,omitempty" xml:"HealthyJobCount,omitempty"`
	// The total number of jobs that require attention.
	//
	// example:
	//
	// 23
	NeedAttentionJobCount *int64 `json:"NeedAttentionJobCount,omitempty" xml:"NeedAttentionJobCount,omitempty"`
	// The score for jobs.
	//
	// example:
	//
	// 56
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The day-to-day growth rate of the score for jobs.
	//
	// example:
	//
	// 0.03
	ScoreDayGrowthRatio *float32 `json:"ScoreDayGrowthRatio,omitempty" xml:"ScoreDayGrowthRatio,omitempty"`
	// The total number of sub-healthy jobs.
	//
	// example:
	//
	// 13
	SubHealthyJobCount *int64 `json:"SubHealthyJobCount,omitempty" xml:"SubHealthyJobCount,omitempty"`
	// The total number of unhealthy jobs.
	//
	// example:
	//
	// 123
	UnhealthyJobCount *int64 `json:"UnhealthyJobCount,omitempty" xml:"UnhealthyJobCount,omitempty"`
}

func (s ListDoctorComputeSummaryResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) GetHealthyJobCount() *int64 {
	return s.HealthyJobCount
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) GetNeedAttentionJobCount() *int64 {
	return s.NeedAttentionJobCount
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) GetScore() *int32 {
	return s.Score
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) GetScoreDayGrowthRatio() *float32 {
	return s.ScoreDayGrowthRatio
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) GetSubHealthyJobCount() *int64 {
	return s.SubHealthyJobCount
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) GetUnhealthyJobCount() *int64 {
	return s.UnhealthyJobCount
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) SetHealthyJobCount(v int64) *ListDoctorComputeSummaryResponseBodyDataAnalysis {
	s.HealthyJobCount = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) SetNeedAttentionJobCount(v int64) *ListDoctorComputeSummaryResponseBodyDataAnalysis {
	s.NeedAttentionJobCount = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) SetScore(v int32) *ListDoctorComputeSummaryResponseBodyDataAnalysis {
	s.Score = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) SetScoreDayGrowthRatio(v float32) *ListDoctorComputeSummaryResponseBodyDataAnalysis {
	s.ScoreDayGrowthRatio = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) SetSubHealthyJobCount(v int64) *ListDoctorComputeSummaryResponseBodyDataAnalysis {
	s.SubHealthyJobCount = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) SetUnhealthyJobCount(v int64) *ListDoctorComputeSummaryResponseBodyDataAnalysis {
	s.UnhealthyJobCount = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type ListDoctorComputeSummaryResponseBodyDataMetrics struct {
	// The total memory consumption over time in seconds.
	MemSeconds *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds `json:"MemSeconds,omitempty" xml:"MemSeconds,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total memory consumption over time in seconds.
	MemSecondsDayGrowthRatio *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio `json:"MemSecondsDayGrowthRatio,omitempty" xml:"MemSecondsDayGrowthRatio,omitempty" type:"Struct"`
	// The average memory usage.
	MemUtilization *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization `json:"MemUtilization,omitempty" xml:"MemUtilization,omitempty" type:"Struct"`
	// The total amount of data read from the file system.
	ReadSize *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize `json:"ReadSize,omitempty" xml:"ReadSize,omitempty" type:"Struct"`
	// The total CPU consumption over time in seconds.
	VcoreSeconds *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds `json:"VcoreSeconds,omitempty" xml:"VcoreSeconds,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total CPU consumption over time in seconds.
	VcoreSecondsDayGrowthRatio *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio `json:"VcoreSecondsDayGrowthRatio,omitempty" xml:"VcoreSecondsDayGrowthRatio,omitempty" type:"Struct"`
	// The average CPU utilization. The meaning is the same as the %CPU parameter in the output of the top command in Linux.
	VcoreUtilization *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization `json:"VcoreUtilization,omitempty" xml:"VcoreUtilization,omitempty" type:"Struct"`
	// The total amount of data written to the file system.
	WriteSize *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize `json:"WriteSize,omitempty" xml:"WriteSize,omitempty" type:"Struct"`
}

func (s ListDoctorComputeSummaryResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) GetMemSeconds() *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds {
	return s.MemSeconds
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) GetMemSecondsDayGrowthRatio() *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio {
	return s.MemSecondsDayGrowthRatio
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) GetMemUtilization() *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization {
	return s.MemUtilization
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) GetReadSize() *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize {
	return s.ReadSize
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) GetVcoreSeconds() *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds {
	return s.VcoreSeconds
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) GetVcoreSecondsDayGrowthRatio() *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio {
	return s.VcoreSecondsDayGrowthRatio
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) GetVcoreUtilization() *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization {
	return s.VcoreUtilization
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) GetWriteSize() *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize {
	return s.WriteSize
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) SetMemSeconds(v *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) *ListDoctorComputeSummaryResponseBodyDataMetrics {
	s.MemSeconds = v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) SetMemSecondsDayGrowthRatio(v *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) *ListDoctorComputeSummaryResponseBodyDataMetrics {
	s.MemSecondsDayGrowthRatio = v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) SetMemUtilization(v *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) *ListDoctorComputeSummaryResponseBodyDataMetrics {
	s.MemUtilization = v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) SetReadSize(v *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) *ListDoctorComputeSummaryResponseBodyDataMetrics {
	s.ReadSize = v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) SetVcoreSeconds(v *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) *ListDoctorComputeSummaryResponseBodyDataMetrics {
	s.VcoreSeconds = v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) SetVcoreSecondsDayGrowthRatio(v *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) *ListDoctorComputeSummaryResponseBodyDataMetrics {
	s.VcoreSecondsDayGrowthRatio = v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) SetVcoreUtilization(v *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) *ListDoctorComputeSummaryResponseBodyDataMetrics {
	s.VcoreUtilization = v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) SetWriteSize(v *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) *ListDoctorComputeSummaryResponseBodyDataMetrics {
	s.WriteSize = v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetrics) Validate() error {
	if s.MemSeconds != nil {
		if err := s.MemSeconds.Validate(); err != nil {
			return err
		}
	}
	if s.MemSecondsDayGrowthRatio != nil {
		if err := s.MemSecondsDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.MemUtilization != nil {
		if err := s.MemUtilization.Validate(); err != nil {
			return err
		}
	}
	if s.ReadSize != nil {
		if err := s.ReadSize.Validate(); err != nil {
			return err
		}
	}
	if s.VcoreSeconds != nil {
		if err := s.VcoreSeconds.Validate(); err != nil {
			return err
		}
	}
	if s.VcoreSecondsDayGrowthRatio != nil {
		if err := s.VcoreSecondsDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.VcoreUtilization != nil {
		if err := s.VcoreUtilization.Validate(); err != nil {
			return err
		}
	}
	if s.WriteSize != nil {
		if err := s.WriteSize.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds struct {
	// The description of the metric.
	//
	// example:
	//
	// Total memory usage over time in seconds
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// memSeconds
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// MB 	- Sec
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 12312312
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) GetName() *string {
	return s.Name
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) SetDescription(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds {
	s.Description = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) SetName(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds {
	s.Name = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) SetUnit(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds {
	s.Unit = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) SetValue(v int64) *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds {
	s.Value = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) Validate() error {
	return dara.Validate(s)
}

type ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Growth ratio of memory usage in seconds per day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// memSecondsDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.36
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) SetDescription(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) SetName(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) SetUnit(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) SetValue(v float32) *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of used memory to total available memory
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// memUtilization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.82
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) GetName() *string {
	return s.Name
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) SetDescription(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization {
	s.Description = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) SetName(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization {
	s.Name = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) SetUnit(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization {
	s.Unit = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) SetValue(v float32) *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization {
	s.Value = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) Validate() error {
	return dara.Validate(s)
}

type ListDoctorComputeSummaryResponseBodyDataMetricsReadSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Total cumulative size of data read in megabytes (MB)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// readSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 504888659968
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) SetDescription(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize {
	s.Description = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) SetName(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize {
	s.Name = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) SetUnit(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) SetValue(v int64) *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize {
	s.Value = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsReadSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds struct {
	// The description of the metric.
	//
	// example:
	//
	// Total vcore usage over time in seconds
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// vcoreSeconds
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// VCores 	- Sec
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1231412
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) GetName() *string {
	return s.Name
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) SetDescription(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds {
	s.Description = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) SetName(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds {
	s.Name = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) SetUnit(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds {
	s.Unit = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) SetValue(v int64) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds {
	s.Value = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) Validate() error {
	return dara.Validate(s)
}

type ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Growth ratio of virtual core usage in seconds per day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// vcoreSecondsDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.22
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) SetDescription(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) SetName(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) SetUnit(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) SetValue(v float32) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of used vcore to total available cores
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// vcoreUtilization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 32.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) GetName() *string {
	return s.Name
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) SetDescription(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization {
	s.Description = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) SetName(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization {
	s.Name = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) SetUnit(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization {
	s.Unit = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) SetValue(v float32) *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization {
	s.Value = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) Validate() error {
	return dara.Validate(s)
}

type ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Total cumulative size of data written in megabytes (MB)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// writeSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 6294093393920
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) SetDescription(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize {
	s.Description = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) SetName(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize {
	s.Name = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) SetUnit(v string) *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) SetValue(v int64) *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize {
	s.Value = &v
	return s
}

func (s *ListDoctorComputeSummaryResponseBodyDataMetricsWriteSize) Validate() error {
	return dara.Validate(s)
}
