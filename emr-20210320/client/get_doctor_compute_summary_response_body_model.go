// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorComputeSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorComputeSummaryResponseBodyData) *GetDoctorComputeSummaryResponseBody
	GetData() *GetDoctorComputeSummaryResponseBodyData
	SetRequestId(v string) *GetDoctorComputeSummaryResponseBody
	GetRequestId() *string
}

type GetDoctorComputeSummaryResponseBody struct {
	// The details of resource usage.
	Data *GetDoctorComputeSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorComputeSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBody) GetData() *GetDoctorComputeSummaryResponseBodyData {
	return s.Data
}

func (s *GetDoctorComputeSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorComputeSummaryResponseBody) SetData(v *GetDoctorComputeSummaryResponseBodyData) *GetDoctorComputeSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorComputeSummaryResponseBody) SetRequestId(v string) *GetDoctorComputeSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDoctorComputeSummaryResponseBodyData struct {
	// The resource analysis information.
	Analysis *GetDoctorComputeSummaryResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The metrics.
	Metrics *GetDoctorComputeSummaryResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
}

func (s GetDoctorComputeSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBodyData) GetAnalysis() *GetDoctorComputeSummaryResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *GetDoctorComputeSummaryResponseBodyData) GetMetrics() *GetDoctorComputeSummaryResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorComputeSummaryResponseBodyData) SetAnalysis(v *GetDoctorComputeSummaryResponseBodyDataAnalysis) *GetDoctorComputeSummaryResponseBodyData {
	s.Analysis = v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyData) SetMetrics(v *GetDoctorComputeSummaryResponseBodyDataMetrics) *GetDoctorComputeSummaryResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyData) Validate() error {
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

type GetDoctorComputeSummaryResponseBodyDataAnalysis struct {
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
	// 234
	NeedAttentionJobCount *int64 `json:"NeedAttentionJobCount,omitempty" xml:"NeedAttentionJobCount,omitempty"`
	// The score for jobs.
	//
	// example:
	//
	// 73
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The day-to-day growth rate of the score for jobs.
	//
	// example:
	//
	// 0.02
	ScoreDayGrowthRatio *float32 `json:"ScoreDayGrowthRatio,omitempty" xml:"ScoreDayGrowthRatio,omitempty"`
	// The total number of sub-healthy jobs.
	//
	// example:
	//
	// 1123
	SubHealthyJobCount *int64 `json:"SubHealthyJobCount,omitempty" xml:"SubHealthyJobCount,omitempty"`
	// The total number of unhealthy jobs.
	//
	// example:
	//
	// 23
	UnhealthyJobCount *int64 `json:"UnhealthyJobCount,omitempty" xml:"UnhealthyJobCount,omitempty"`
}

func (s GetDoctorComputeSummaryResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) GetHealthyJobCount() *int64 {
	return s.HealthyJobCount
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) GetNeedAttentionJobCount() *int64 {
	return s.NeedAttentionJobCount
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) GetScore() *int32 {
	return s.Score
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) GetScoreDayGrowthRatio() *float32 {
	return s.ScoreDayGrowthRatio
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) GetSubHealthyJobCount() *int64 {
	return s.SubHealthyJobCount
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) GetUnhealthyJobCount() *int64 {
	return s.UnhealthyJobCount
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) SetHealthyJobCount(v int64) *GetDoctorComputeSummaryResponseBodyDataAnalysis {
	s.HealthyJobCount = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) SetNeedAttentionJobCount(v int64) *GetDoctorComputeSummaryResponseBodyDataAnalysis {
	s.NeedAttentionJobCount = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) SetScore(v int32) *GetDoctorComputeSummaryResponseBodyDataAnalysis {
	s.Score = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) SetScoreDayGrowthRatio(v float32) *GetDoctorComputeSummaryResponseBodyDataAnalysis {
	s.ScoreDayGrowthRatio = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) SetSubHealthyJobCount(v int64) *GetDoctorComputeSummaryResponseBodyDataAnalysis {
	s.SubHealthyJobCount = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) SetUnhealthyJobCount(v int64) *GetDoctorComputeSummaryResponseBodyDataAnalysis {
	s.UnhealthyJobCount = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type GetDoctorComputeSummaryResponseBodyDataMetrics struct {
	// The total memory consumption over time in seconds.
	MemSeconds *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds `json:"MemSeconds,omitempty" xml:"MemSeconds,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total memory consumption over time in seconds.
	MemSecondsDayGrowthRatio *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio `json:"MemSecondsDayGrowthRatio,omitempty" xml:"MemSecondsDayGrowthRatio,omitempty" type:"Struct"`
	// The average memory usage.
	MemUtilization *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization `json:"MemUtilization,omitempty" xml:"MemUtilization,omitempty" type:"Struct"`
	// The total amount of data read from the file system.
	ReadSize *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize `json:"ReadSize,omitempty" xml:"ReadSize,omitempty" type:"Struct"`
	// The total CPU consumption over time in seconds.
	VcoreSeconds *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds `json:"VcoreSeconds,omitempty" xml:"VcoreSeconds,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total CPU consumption over time in seconds.
	VcoreSecondsDayGrowthRatio *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio `json:"VcoreSecondsDayGrowthRatio,omitempty" xml:"VcoreSecondsDayGrowthRatio,omitempty" type:"Struct"`
	// The average CPU utilization. The meaning is the same as the %CPU parameter in the output of the top command in Linux.
	VcoreUtilization *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization `json:"VcoreUtilization,omitempty" xml:"VcoreUtilization,omitempty" type:"Struct"`
	// The total amount of data written to the file system.
	WriteSize *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize `json:"WriteSize,omitempty" xml:"WriteSize,omitempty" type:"Struct"`
}

func (s GetDoctorComputeSummaryResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) GetMemSeconds() *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds {
	return s.MemSeconds
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) GetMemSecondsDayGrowthRatio() *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio {
	return s.MemSecondsDayGrowthRatio
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) GetMemUtilization() *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization {
	return s.MemUtilization
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) GetReadSize() *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize {
	return s.ReadSize
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) GetVcoreSeconds() *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds {
	return s.VcoreSeconds
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) GetVcoreSecondsDayGrowthRatio() *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio {
	return s.VcoreSecondsDayGrowthRatio
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) GetVcoreUtilization() *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization {
	return s.VcoreUtilization
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) GetWriteSize() *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize {
	return s.WriteSize
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) SetMemSeconds(v *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) *GetDoctorComputeSummaryResponseBodyDataMetrics {
	s.MemSeconds = v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) SetMemSecondsDayGrowthRatio(v *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) *GetDoctorComputeSummaryResponseBodyDataMetrics {
	s.MemSecondsDayGrowthRatio = v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) SetMemUtilization(v *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) *GetDoctorComputeSummaryResponseBodyDataMetrics {
	s.MemUtilization = v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) SetReadSize(v *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) *GetDoctorComputeSummaryResponseBodyDataMetrics {
	s.ReadSize = v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) SetVcoreSeconds(v *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) *GetDoctorComputeSummaryResponseBodyDataMetrics {
	s.VcoreSeconds = v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) SetVcoreSecondsDayGrowthRatio(v *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) *GetDoctorComputeSummaryResponseBodyDataMetrics {
	s.VcoreSecondsDayGrowthRatio = v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) SetVcoreUtilization(v *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) *GetDoctorComputeSummaryResponseBodyDataMetrics {
	s.VcoreUtilization = v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) SetWriteSize(v *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) *GetDoctorComputeSummaryResponseBodyDataMetrics {
	s.WriteSize = v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetrics) Validate() error {
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

type GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds struct {
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

func (s GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) GetName() *string {
	return s.Name
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) SetDescription(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds {
	s.Description = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) SetName(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds {
	s.Name = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) SetUnit(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds {
	s.Unit = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) SetValue(v int64) *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds {
	s.Value = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSeconds) Validate() error {
	return dara.Validate(s)
}

type GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio struct {
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

func (s GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) SetDescription(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) SetName(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) SetUnit(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) SetValue(v float32) *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemSecondsDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.82
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) GetName() *string {
	return s.Name
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) SetDescription(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization {
	s.Description = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) SetName(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization {
	s.Name = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) SetUnit(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization {
	s.Unit = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) SetValue(v float32) *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization {
	s.Value = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsMemUtilization) Validate() error {
	return dara.Validate(s)
}

type GetDoctorComputeSummaryResponseBodyDataMetricsReadSize struct {
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

func (s GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) SetDescription(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize {
	s.Description = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) SetName(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize {
	s.Name = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) SetUnit(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) SetValue(v int64) *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize {
	s.Value = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsReadSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds struct {
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

func (s GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) GetName() *string {
	return s.Name
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) SetDescription(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds {
	s.Description = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) SetName(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds {
	s.Name = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) SetUnit(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds {
	s.Unit = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) SetValue(v int64) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds {
	s.Value = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSeconds) Validate() error {
	return dara.Validate(s)
}

type GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio struct {
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

func (s GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) SetDescription(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) SetName(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) SetUnit(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) SetValue(v float32) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreSecondsDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization struct {
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

func (s GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) GetName() *string {
	return s.Name
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) SetDescription(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization {
	s.Description = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) SetName(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization {
	s.Name = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) SetUnit(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization {
	s.Unit = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) SetValue(v float32) *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization {
	s.Value = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsVcoreUtilization) Validate() error {
	return dara.Validate(s)
}

type GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize struct {
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

func (s GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) SetDescription(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize {
	s.Description = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) SetName(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize {
	s.Name = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) SetUnit(v string) *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) SetValue(v int64) *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize {
	s.Value = &v
	return s
}

func (s *GetDoctorComputeSummaryResponseBodyDataMetricsWriteSize) Validate() error {
	return dara.Validate(s)
}
