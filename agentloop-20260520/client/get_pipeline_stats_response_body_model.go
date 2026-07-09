// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetPipelineStatsResponseBody
	GetEndTime() *int64
	SetGranularity(v string) *GetPipelineStatsResponseBody
	GetGranularity() *string
	SetPipelineName(v string) *GetPipelineStatsResponseBody
	GetPipelineName() *string
	SetRequestId(v string) *GetPipelineStatsResponseBody
	GetRequestId() *string
	SetStartTime(v int64) *GetPipelineStatsResponseBody
	GetStartTime() *int64
	SetSummary(v *GetPipelineStatsResponseBodySummary) *GetPipelineStatsResponseBody
	GetSummary() *GetPipelineStatsResponseBodySummary
	SetTimeSeries(v []*GetPipelineStatsResponseBodyTimeSeries) *GetPipelineStatsResponseBody
	GetTimeSeries() []*GetPipelineStatsResponseBodyTimeSeries
}

type GetPipelineStatsResponseBody struct {
	// example:
	//
	// 1735660800
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// Hour
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The name of the pipeline.
	//
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The request ID, which is used to locate the request during troubleshooting.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1735574400
	StartTime  *int64                                    `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Summary    *GetPipelineStatsResponseBodySummary      `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
	TimeSeries []*GetPipelineStatsResponseBodyTimeSeries `json:"timeSeries,omitempty" xml:"timeSeries,omitempty" type:"Repeated"`
}

func (s GetPipelineStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineStatsResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetPipelineStatsResponseBody) GetGranularity() *string {
	return s.Granularity
}

func (s *GetPipelineStatsResponseBody) GetPipelineName() *string {
	return s.PipelineName
}

func (s *GetPipelineStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineStatsResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetPipelineStatsResponseBody) GetSummary() *GetPipelineStatsResponseBodySummary {
	return s.Summary
}

func (s *GetPipelineStatsResponseBody) GetTimeSeries() []*GetPipelineStatsResponseBodyTimeSeries {
	return s.TimeSeries
}

func (s *GetPipelineStatsResponseBody) SetEndTime(v int64) *GetPipelineStatsResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetPipelineStatsResponseBody) SetGranularity(v string) *GetPipelineStatsResponseBody {
	s.Granularity = &v
	return s
}

func (s *GetPipelineStatsResponseBody) SetPipelineName(v string) *GetPipelineStatsResponseBody {
	s.PipelineName = &v
	return s
}

func (s *GetPipelineStatsResponseBody) SetRequestId(v string) *GetPipelineStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineStatsResponseBody) SetStartTime(v int64) *GetPipelineStatsResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetPipelineStatsResponseBody) SetSummary(v *GetPipelineStatsResponseBodySummary) *GetPipelineStatsResponseBody {
	s.Summary = v
	return s
}

func (s *GetPipelineStatsResponseBody) SetTimeSeries(v []*GetPipelineStatsResponseBodyTimeSeries) *GetPipelineStatsResponseBody {
	s.TimeSeries = v
	return s
}

func (s *GetPipelineStatsResponseBody) Validate() error {
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	if s.TimeSeries != nil {
		for _, item := range s.TimeSeries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPipelineStatsResponseBodySummary struct {
	// example:
	//
	// 2500
	AvgElapsedMs *int64 `json:"avgElapsedMs,omitempty" xml:"avgElapsedMs,omitempty"`
	// example:
	//
	// 0
	CancelledRuns *int64 `json:"cancelledRuns,omitempty" xml:"cancelledRuns,omitempty"`
	// example:
	//
	// 1735660800
	CommittedWatermark *int64 `json:"committedWatermark,omitempty" xml:"committedWatermark,omitempty"`
	// example:
	//
	// 0
	FailedRuns *int64 `json:"failedRuns,omitempty" xml:"failedRuns,omitempty"`
	// example:
	//
	// 120
	ScheduleLagSeconds *int64 `json:"scheduleLagSeconds,omitempty" xml:"scheduleLagSeconds,omitempty"`
	// example:
	//
	// 44
	SucceededRuns *int64 `json:"succeededRuns,omitempty" xml:"succeededRuns,omitempty"`
	// example:
	//
	// 1.0
	SuccessRate *float64 `json:"successRate,omitempty" xml:"successRate,omitempty"`
	// example:
	//
	// 3221225472
	TotalOutputBytes *int64 `json:"totalOutputBytes,omitempty" xml:"totalOutputBytes,omitempty"`
	// example:
	//
	// 1200000
	TotalOutputRows *int64 `json:"totalOutputRows,omitempty" xml:"totalOutputRows,omitempty"`
	// example:
	//
	// 5368709120
	TotalProcessedBytes *int64 `json:"totalProcessedBytes,omitempty" xml:"totalProcessedBytes,omitempty"`
	// example:
	//
	// 1500000
	TotalProcessedRows *int64 `json:"totalProcessedRows,omitempty" xml:"totalProcessedRows,omitempty"`
	// example:
	//
	// 44
	TotalRuns *int64 `json:"totalRuns,omitempty" xml:"totalRuns,omitempty"`
}

func (s GetPipelineStatsResponseBodySummary) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineStatsResponseBodySummary) GoString() string {
	return s.String()
}

func (s *GetPipelineStatsResponseBodySummary) GetAvgElapsedMs() *int64 {
	return s.AvgElapsedMs
}

func (s *GetPipelineStatsResponseBodySummary) GetCancelledRuns() *int64 {
	return s.CancelledRuns
}

func (s *GetPipelineStatsResponseBodySummary) GetCommittedWatermark() *int64 {
	return s.CommittedWatermark
}

func (s *GetPipelineStatsResponseBodySummary) GetFailedRuns() *int64 {
	return s.FailedRuns
}

func (s *GetPipelineStatsResponseBodySummary) GetScheduleLagSeconds() *int64 {
	return s.ScheduleLagSeconds
}

func (s *GetPipelineStatsResponseBodySummary) GetSucceededRuns() *int64 {
	return s.SucceededRuns
}

func (s *GetPipelineStatsResponseBodySummary) GetSuccessRate() *float64 {
	return s.SuccessRate
}

func (s *GetPipelineStatsResponseBodySummary) GetTotalOutputBytes() *int64 {
	return s.TotalOutputBytes
}

func (s *GetPipelineStatsResponseBodySummary) GetTotalOutputRows() *int64 {
	return s.TotalOutputRows
}

func (s *GetPipelineStatsResponseBodySummary) GetTotalProcessedBytes() *int64 {
	return s.TotalProcessedBytes
}

func (s *GetPipelineStatsResponseBodySummary) GetTotalProcessedRows() *int64 {
	return s.TotalProcessedRows
}

func (s *GetPipelineStatsResponseBodySummary) GetTotalRuns() *int64 {
	return s.TotalRuns
}

func (s *GetPipelineStatsResponseBodySummary) SetAvgElapsedMs(v int64) *GetPipelineStatsResponseBodySummary {
	s.AvgElapsedMs = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) SetCancelledRuns(v int64) *GetPipelineStatsResponseBodySummary {
	s.CancelledRuns = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) SetCommittedWatermark(v int64) *GetPipelineStatsResponseBodySummary {
	s.CommittedWatermark = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) SetFailedRuns(v int64) *GetPipelineStatsResponseBodySummary {
	s.FailedRuns = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) SetScheduleLagSeconds(v int64) *GetPipelineStatsResponseBodySummary {
	s.ScheduleLagSeconds = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) SetSucceededRuns(v int64) *GetPipelineStatsResponseBodySummary {
	s.SucceededRuns = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) SetSuccessRate(v float64) *GetPipelineStatsResponseBodySummary {
	s.SuccessRate = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) SetTotalOutputBytes(v int64) *GetPipelineStatsResponseBodySummary {
	s.TotalOutputBytes = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) SetTotalOutputRows(v int64) *GetPipelineStatsResponseBodySummary {
	s.TotalOutputRows = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) SetTotalProcessedBytes(v int64) *GetPipelineStatsResponseBodySummary {
	s.TotalProcessedBytes = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) SetTotalProcessedRows(v int64) *GetPipelineStatsResponseBodySummary {
	s.TotalProcessedRows = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) SetTotalRuns(v int64) *GetPipelineStatsResponseBodySummary {
	s.TotalRuns = &v
	return s
}

func (s *GetPipelineStatsResponseBodySummary) Validate() error {
	return dara.Validate(s)
}

type GetPipelineStatsResponseBodyTimeSeries struct {
	// example:
	//
	// 2500
	AvgElapsedMs *int64 `json:"avgElapsedMs,omitempty" xml:"avgElapsedMs,omitempty"`
	// example:
	//
	// 322122547
	OutputBytes *int64 `json:"outputBytes,omitempty" xml:"outputBytes,omitempty"`
	// example:
	//
	// 80000
	OutputRows *int64 `json:"outputRows,omitempty" xml:"outputRows,omitempty"`
	// example:
	//
	// 536870912
	ProcessedBytes *int64 `json:"processedBytes,omitempty" xml:"processedBytes,omitempty"`
	// example:
	//
	// 100000
	ProcessedRows *int64 `json:"processedRows,omitempty" xml:"processedRows,omitempty"`
	// example:
	//
	// 5
	Runs *int64 `json:"runs,omitempty" xml:"runs,omitempty"`
	// example:
	//
	// 5
	SucceededRuns *int64 `json:"succeededRuns,omitempty" xml:"succeededRuns,omitempty"`
	// example:
	//
	// 1735574400
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetPipelineStatsResponseBodyTimeSeries) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineStatsResponseBodyTimeSeries) GoString() string {
	return s.String()
}

func (s *GetPipelineStatsResponseBodyTimeSeries) GetAvgElapsedMs() *int64 {
	return s.AvgElapsedMs
}

func (s *GetPipelineStatsResponseBodyTimeSeries) GetOutputBytes() *int64 {
	return s.OutputBytes
}

func (s *GetPipelineStatsResponseBodyTimeSeries) GetOutputRows() *int64 {
	return s.OutputRows
}

func (s *GetPipelineStatsResponseBodyTimeSeries) GetProcessedBytes() *int64 {
	return s.ProcessedBytes
}

func (s *GetPipelineStatsResponseBodyTimeSeries) GetProcessedRows() *int64 {
	return s.ProcessedRows
}

func (s *GetPipelineStatsResponseBodyTimeSeries) GetRuns() *int64 {
	return s.Runs
}

func (s *GetPipelineStatsResponseBodyTimeSeries) GetSucceededRuns() *int64 {
	return s.SucceededRuns
}

func (s *GetPipelineStatsResponseBodyTimeSeries) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetPipelineStatsResponseBodyTimeSeries) SetAvgElapsedMs(v int64) *GetPipelineStatsResponseBodyTimeSeries {
	s.AvgElapsedMs = &v
	return s
}

func (s *GetPipelineStatsResponseBodyTimeSeries) SetOutputBytes(v int64) *GetPipelineStatsResponseBodyTimeSeries {
	s.OutputBytes = &v
	return s
}

func (s *GetPipelineStatsResponseBodyTimeSeries) SetOutputRows(v int64) *GetPipelineStatsResponseBodyTimeSeries {
	s.OutputRows = &v
	return s
}

func (s *GetPipelineStatsResponseBodyTimeSeries) SetProcessedBytes(v int64) *GetPipelineStatsResponseBodyTimeSeries {
	s.ProcessedBytes = &v
	return s
}

func (s *GetPipelineStatsResponseBodyTimeSeries) SetProcessedRows(v int64) *GetPipelineStatsResponseBodyTimeSeries {
	s.ProcessedRows = &v
	return s
}

func (s *GetPipelineStatsResponseBodyTimeSeries) SetRuns(v int64) *GetPipelineStatsResponseBodyTimeSeries {
	s.Runs = &v
	return s
}

func (s *GetPipelineStatsResponseBodyTimeSeries) SetSucceededRuns(v int64) *GetPipelineStatsResponseBodyTimeSeries {
	s.SucceededRuns = &v
	return s
}

func (s *GetPipelineStatsResponseBodyTimeSeries) SetTimestamp(v int64) *GetPipelineStatsResponseBodyTimeSeries {
	s.Timestamp = &v
	return s
}

func (s *GetPipelineStatsResponseBodyTimeSeries) Validate() error {
	return dara.Validate(s)
}
