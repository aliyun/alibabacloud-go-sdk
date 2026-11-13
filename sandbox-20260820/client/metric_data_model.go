// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricData interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *MetricData
	GetEndTime() *string
	SetGroupBy(v []*string) *MetricData
	GetGroupBy() []*string
	SetPeriod(v int32) *MetricData
	GetPeriod() *int32
	SetResults(v []*MetricDataResults) *MetricData
	GetResults() []*MetricDataResults
	SetStartTime(v string) *MetricData
	GetStartTime() *string
	SetTime(v string) *MetricData
	GetTime() *string
}

type MetricData struct {
	// The end time of the time series query, exclusive.
	//
	// example:
	//
	// 2026-11-05T00:02:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The grouping keys used for this query. An empty array is returned if no grouping is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["teamID"]
	GroupBy []*string `json:"groupBy" xml:"groupBy" type:"Repeated"`
	// The data point interval used for this query, in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
	// The result for each requested metric.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"metricName":"cpuUsedCores","unit":"Core","series":[{"dimensions":{"teamID":"70d1c834-0333-58d8-97ac-5336eb9179eb"},"points":[{"timestamp":"2026-11-05T00:00:00Z","value":0.25},{"timestamp":"2026-11-05T00:01:00Z","value":0.5}]}]}]
	Results []*MetricDataResults `json:"results" xml:"results" type:"Repeated"`
	// The start time of the time series query, inclusive.
	//
	// example:
	//
	// 2026-11-05T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The time of the query. Returned only for single-point queries.
	//
	// example:
	//
	// 2026-11-05T00:01:00Z
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s MetricData) String() string {
	return dara.Prettify(s)
}

func (s MetricData) GoString() string {
	return s.String()
}

func (s *MetricData) GetEndTime() *string {
	return s.EndTime
}

func (s *MetricData) GetGroupBy() []*string {
	return s.GroupBy
}

func (s *MetricData) GetPeriod() *int32 {
	return s.Period
}

func (s *MetricData) GetResults() []*MetricDataResults {
	return s.Results
}

func (s *MetricData) GetStartTime() *string {
	return s.StartTime
}

func (s *MetricData) GetTime() *string {
	return s.Time
}

func (s *MetricData) SetEndTime(v string) *MetricData {
	s.EndTime = &v
	return s
}

func (s *MetricData) SetGroupBy(v []*string) *MetricData {
	s.GroupBy = v
	return s
}

func (s *MetricData) SetPeriod(v int32) *MetricData {
	s.Period = &v
	return s
}

func (s *MetricData) SetResults(v []*MetricDataResults) *MetricData {
	s.Results = v
	return s
}

func (s *MetricData) SetStartTime(v string) *MetricData {
	s.StartTime = &v
	return s
}

func (s *MetricData) SetTime(v string) *MetricData {
	s.Time = &v
	return s
}

func (s *MetricData) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MetricDataResults struct {
	// The name of the metric in the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpuUsedCores
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The time series for the metric. An empty array is returned if no data points exist.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"dimensions":{"teamID":"70d1c834-0333-58d8-97ac-5336eb9179eb"},"points":[{"timestamp":"2026-11-05T00:00:00Z","value":0.25},{"timestamp":"2026-11-05T00:01:00Z","value":0.5}]}]
	Series []*MetricDataResultsSeries `json:"series" xml:"series" type:"Repeated"`
	// The unit of the metric value, such as Core for CPU cores.
	//
	// This parameter is required.
	//
	// example:
	//
	// Core
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s MetricDataResults) String() string {
	return dara.Prettify(s)
}

func (s MetricDataResults) GoString() string {
	return s.String()
}

func (s *MetricDataResults) GetMetricName() *string {
	return s.MetricName
}

func (s *MetricDataResults) GetSeries() []*MetricDataResultsSeries {
	return s.Series
}

func (s *MetricDataResults) GetUnit() *string {
	return s.Unit
}

func (s *MetricDataResults) SetMetricName(v string) *MetricDataResults {
	s.MetricName = &v
	return s
}

func (s *MetricDataResults) SetSeries(v []*MetricDataResultsSeries) *MetricDataResults {
	s.Series = v
	return s
}

func (s *MetricDataResults) SetUnit(v string) *MetricDataResults {
	s.Unit = &v
	return s
}

func (s *MetricDataResults) Validate() error {
	if s.Series != nil {
		for _, item := range s.Series {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MetricDataResultsSeries struct {
	// The team, template, or other grouping values for this series.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"teamID":"70d1c834-0333-58d8-97ac-5336eb9179eb"}
	Dimensions map[string]*string `json:"dimensions" xml:"dimensions"`
	// The data points in chronological order. Only times with metric values are listed.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"timestamp":"2026-11-05T00:00:00Z","value":0.25},{"timestamp":"2026-11-05T00:01:00Z","value":0.5}]
	Points []*MetricDataResultsSeriesPoints `json:"points" xml:"points" type:"Repeated"`
}

func (s MetricDataResultsSeries) String() string {
	return dara.Prettify(s)
}

func (s MetricDataResultsSeries) GoString() string {
	return s.String()
}

func (s *MetricDataResultsSeries) GetDimensions() map[string]*string {
	return s.Dimensions
}

func (s *MetricDataResultsSeries) GetPoints() []*MetricDataResultsSeriesPoints {
	return s.Points
}

func (s *MetricDataResultsSeries) SetDimensions(v map[string]*string) *MetricDataResultsSeries {
	s.Dimensions = v
	return s
}

func (s *MetricDataResultsSeries) SetPoints(v []*MetricDataResultsSeriesPoints) *MetricDataResultsSeries {
	s.Points = v
	return s
}

func (s *MetricDataResultsSeries) Validate() error {
	if s.Points != nil {
		for _, item := range s.Points {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MetricDataResultsSeriesPoints struct {
	// The time corresponding to the metric value, in UTC seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-11-05T00:00:00Z
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// The metric value at this time. The unit is specified in unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.25
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MetricDataResultsSeriesPoints) String() string {
	return dara.Prettify(s)
}

func (s MetricDataResultsSeriesPoints) GoString() string {
	return s.String()
}

func (s *MetricDataResultsSeriesPoints) GetTimestamp() *string {
	return s.Timestamp
}

func (s *MetricDataResultsSeriesPoints) GetValue() *float64 {
	return s.Value
}

func (s *MetricDataResultsSeriesPoints) SetTimestamp(v string) *MetricDataResultsSeriesPoints {
	s.Timestamp = &v
	return s
}

func (s *MetricDataResultsSeriesPoints) SetValue(v float64) *MetricDataResultsSeriesPoints {
	s.Value = &v
	return s
}

func (s *MetricDataResultsSeriesPoints) Validate() error {
	return dara.Validate(s)
}
