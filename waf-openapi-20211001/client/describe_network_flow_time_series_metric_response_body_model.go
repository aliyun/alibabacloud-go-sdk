// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkFlowTimeSeriesMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkFlowTimeSeries(v []*DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries) *DescribeNetworkFlowTimeSeriesMetricResponseBody
	GetNetworkFlowTimeSeries() []*DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries
	SetRequestId(v string) *DescribeNetworkFlowTimeSeriesMetricResponseBody
	GetRequestId() *string
	SetTimeSeriesMetaData(v *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData) *DescribeNetworkFlowTimeSeriesMetricResponseBody
	GetTimeSeriesMetaData() *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData
}

type DescribeNetworkFlowTimeSeriesMetricResponseBody struct {
	// The time series data. Multiple data series can be returned.
	NetworkFlowTimeSeries []*DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries `json:"NetworkFlowTimeSeries,omitempty" xml:"NetworkFlowTimeSeries,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D827FCFE-90A7-4330-9326-D33C8B4*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the returned data.
	TimeSeriesMetaData *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData `json:"TimeSeriesMetaData,omitempty" xml:"TimeSeriesMetaData,omitempty" type:"Struct"`
}

func (s DescribeNetworkFlowTimeSeriesMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTimeSeriesMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBody) GetNetworkFlowTimeSeries() []*DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries {
	return s.NetworkFlowTimeSeries
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBody) GetTimeSeriesMetaData() *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData {
	return s.TimeSeriesMetaData
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBody) SetNetworkFlowTimeSeries(v []*DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries) *DescribeNetworkFlowTimeSeriesMetricResponseBody {
	s.NetworkFlowTimeSeries = v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBody) SetRequestId(v string) *DescribeNetworkFlowTimeSeriesMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBody) SetTimeSeriesMetaData(v *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData) *DescribeNetworkFlowTimeSeriesMetricResponseBody {
	s.TimeSeriesMetaData = v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBody) Validate() error {
	if s.NetworkFlowTimeSeries != nil {
		for _, item := range s.NetworkFlowTimeSeries {
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

type DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries struct {
	// The metric name. This value is the same as the Metric request parameter.
	//
	// example:
	//
	// total_requests
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The timestamps. Each value represents the start of a time interval.
	Timestamps []*string `json:"Timestamps,omitempty" xml:"Timestamps,omitempty" type:"Repeated"`
	// The metric values. Each value represents the count within the corresponding time interval.
	Values []*int64 `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries) GetMetric() *string {
	return s.Metric
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries) GetTimestamps() []*string {
	return s.Timestamps
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries) GetValues() []*int64 {
	return s.Values
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries) SetMetric(v string) *DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries {
	s.Metric = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries) SetTimestamps(v []*string) *DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries {
	s.Timestamps = v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries) SetValues(v []*int64) *DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries {
	s.Values = v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyNetworkFlowTimeSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData struct {
	// The time granularity of each data point in the returned time series. For example, "15m" indicates that each data point represents statistics for a 15-minute interval. For more information about the time granularity, see the **Time granularity of data points*	- section.
	//
	// example:
	//
	// 1m
	AggregateInterval *string `json:"AggregateInterval,omitempty" xml:"AggregateInterval,omitempty"`
	// The time range that was queried.
	DateRange *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange `json:"DateRange,omitempty" xml:"DateRange,omitempty" type:"Struct"`
	// The unit of the returned data. Default value: requests.
	//
	// example:
	//
	// requests
	Units *string `json:"Units,omitempty" xml:"Units,omitempty"`
}

func (s DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData) GetAggregateInterval() *string {
	return s.AggregateInterval
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData) GetDateRange() *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange {
	return s.DateRange
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData) GetUnits() *string {
	return s.Units
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData) SetAggregateInterval(v string) *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData {
	s.AggregateInterval = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData) SetDateRange(v *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData {
	s.DateRange = v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData) SetUnits(v string) *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData {
	s.Units = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaData) Validate() error {
	if s.DateRange != nil {
		if err := s.DateRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange struct {
	// The end of the time range that was queried. This value is a UNIX timestamp. Unit: seconds. This value is the same as the EndDate request parameter.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The start of the time range that was queried. This value is a UNIX timestamp. Unit: seconds. This value is the same as the StartDate request parameter.
	//
	// example:
	//
	// 1713888000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) SetEndDate(v int64) *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange {
	s.EndDate = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) SetStartDate(v int64) *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange {
	s.StartDate = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponseBodyTimeSeriesMetaDataDateRange) Validate() error {
	return dara.Validate(s)
}
