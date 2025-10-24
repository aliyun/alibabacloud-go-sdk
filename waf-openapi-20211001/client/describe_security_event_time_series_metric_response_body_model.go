// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventTimeSeriesMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSecurityEventTimeSeriesMetricResponseBody
	GetRequestId() *string
	SetSecurityEventTimeSeries(v []*DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries) *DescribeSecurityEventTimeSeriesMetricResponseBody
	GetSecurityEventTimeSeries() []*DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries
	SetTimeSeriesMetaData(v *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData) *DescribeSecurityEventTimeSeriesMetricResponseBody
	GetTimeSeriesMetaData() *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData
}

type DescribeSecurityEventTimeSeriesMetricResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D827FCFE-90A7-4330-9326-*****4C7726
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time series data returned. This operation can return time series for multiple metrics.
	SecurityEventTimeSeries []*DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries `json:"SecurityEventTimeSeries,omitempty" xml:"SecurityEventTimeSeries,omitempty" type:"Repeated"`
	// The metadata of the time series data returned.
	TimeSeriesMetaData *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData `json:"TimeSeriesMetaData,omitempty" xml:"TimeSeriesMetaData,omitempty" type:"Struct"`
}

func (s DescribeSecurityEventTimeSeriesMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTimeSeriesMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBody) GetSecurityEventTimeSeries() []*DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries {
	return s.SecurityEventTimeSeries
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBody) GetTimeSeriesMetaData() *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData {
	return s.TimeSeriesMetaData
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBody) SetRequestId(v string) *DescribeSecurityEventTimeSeriesMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBody) SetSecurityEventTimeSeries(v []*DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries) *DescribeSecurityEventTimeSeriesMetricResponseBody {
	s.SecurityEventTimeSeries = v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBody) SetTimeSeriesMetaData(v *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData) *DescribeSecurityEventTimeSeriesMetricResponseBody {
	s.TimeSeriesMetaData = v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBody) Validate() error {
	if s.SecurityEventTimeSeries != nil {
		for _, item := range s.SecurityEventTimeSeries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TimeSeriesMetaData != nil {
		if err := s.TimeSeriesMetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries struct {
	// The metric. This value is the same as the value of Metric in the request parameters.
	//
	// example:
	//
	// monitored_requests
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The time points. Each point represents a time range.
	Timestamps []*string `json:"Timestamps,omitempty" xml:"Timestamps,omitempty" type:"Repeated"`
	// The data points. Each point represents a count for a time range.
	Values []*int64 `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries) GetMetric() *string {
	return s.Metric
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries) GetTimestamps() []*string {
	return s.Timestamps
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries) GetValues() []*int64 {
	return s.Values
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries) SetMetric(v string) *DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries {
	s.Metric = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries) SetTimestamps(v []*string) *DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries {
	s.Timestamps = v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries) SetValues(v []*int64) *DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries {
	s.Values = v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodySecurityEventTimeSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData struct {
	// The time granularity of data points in the time series data returned. For example, a value of 15m indicates that data points are collected at 15-minute intervals. For more information about time granularities, see the **Time granularities of data points in time series*	- section below.
	//
	// example:
	//
	// 1m
	AggregateInterval *string `json:"AggregateInterval,omitempty" xml:"AggregateInterval,omitempty"`
	// The time range that is used for the query.
	DateRange *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange `json:"DateRange,omitempty" xml:"DateRange,omitempty" type:"Struct"`
	// The unit of the statistics returned. It is fixed as requests.
	//
	// example:
	//
	// requests
	Units *string `json:"Units,omitempty" xml:"Units,omitempty"`
}

func (s DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData) GetAggregateInterval() *string {
	return s.AggregateInterval
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData) GetDateRange() *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange {
	return s.DateRange
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData) GetUnits() *string {
	return s.Units
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData) SetAggregateInterval(v string) *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData {
	s.AggregateInterval = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData) SetDateRange(v *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData {
	s.DateRange = v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData) SetUnits(v string) *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData {
	s.Units = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaData) Validate() error {
	if s.DateRange != nil {
		if err := s.DateRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange struct {
	// The end of the time range to query. The value is a Unix timestamp. Unit: seconds. This value is the same as the value of EndDate in the request parameters.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The beginning of the time range to query. The value is a Unix timestamp. Unit: seconds. This value is the same as the value of StartDate in the request parameters.
	//
	// example:
	//
	// 1713888000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) SetEndDate(v int64) *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange {
	s.EndDate = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) SetStartDate(v int64) *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange {
	s.StartDate = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) Validate() error {
	return dara.Validate(s)
}
