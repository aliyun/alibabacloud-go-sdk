// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iItemsMetricValuesValue interface {
	dara.Model
	String() string
	GoString() string
	SetMetricCode(v string) *ItemsMetricValuesValue
	GetMetricCode() *string
	SetMetricName(v string) *ItemsMetricValuesValue
	GetMetricName() *string
	SetPrimary(v bool) *ItemsMetricValuesValue
	GetPrimary() *bool
	SetTime2(v *ItemsMetricValuesValueTime2) *ItemsMetricValuesValue
	GetTime2() *ItemsMetricValuesValueTime2
	SetAvg(v *ItemsMetricValuesValueAvg) *ItemsMetricValuesValue
	GetAvg() *ItemsMetricValuesValueAvg
	SetSum(v *ItemsMetricValuesValueSum) *ItemsMetricValuesValue
	GetSum() *ItemsMetricValuesValueSum
	SetMax(v *ItemsMetricValuesValueMax) *ItemsMetricValuesValue
	GetMax() *ItemsMetricValuesValueMax
}

type ItemsMetricValuesValue struct {
	// The primary metric code, which matches the key in `MetricValues` and the `MetricType` request parameter. Valid values:
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
	// example:
	//
	// CPU_COST
	MetricCode *string `json:"MetricCode,omitempty" xml:"MetricCode,omitempty"`
	// The primary metric name. The mapping is as follows:
	//
	// - `QUERY_COUNT`: `QueryCount`.
	//
	// - `CPU_COST`: `OperatorCost`.
	//
	// - `SHUFFLE_SIZE`: `ShuffleSize`.
	//
	// - `PEAK_MEMORY`: `PeakMemory`.
	//
	// - `SCAN_SIZE`: `ScanSize`.
	//
	// example:
	//
	// OperatorCost
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// Indicates whether this is the primary metric for the current analysis dimension. The current value is true.
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// The aggregated result for Time 2 in the NEW report. This field is returned only for NEW reports.
	Time2 *ItemsMetricValuesValueTime2 `json:"Time2,omitempty" xml:"Time2,omitempty" type:"Struct"`
	// The dual-window comparison of the average value across active query minute buckets for the CHANGED report. This field is returned only for CHANGED reports.
	Avg *ItemsMetricValuesValueAvg `json:"Avg,omitempty" xml:"Avg,omitempty" type:"Struct"`
	// The dual-window comparison of the sum of metric values across active query minute buckets for the CHANGED report. This field is returned only for CHANGED reports.
	Sum *ItemsMetricValuesValueSum `json:"Sum,omitempty" xml:"Sum,omitempty" type:"Struct"`
	// The dual-window comparison of the peak value in a single minute bucket for the CHANGED report. This field is returned only for CHANGED reports.
	Max *ItemsMetricValuesValueMax `json:"Max,omitempty" xml:"Max,omitempty" type:"Struct"`
}

func (s ItemsMetricValuesValue) String() string {
	return dara.Prettify(s)
}

func (s ItemsMetricValuesValue) GoString() string {
	return s.String()
}

func (s *ItemsMetricValuesValue) GetMetricCode() *string {
	return s.MetricCode
}

func (s *ItemsMetricValuesValue) GetMetricName() *string {
	return s.MetricName
}

func (s *ItemsMetricValuesValue) GetPrimary() *bool {
	return s.Primary
}

func (s *ItemsMetricValuesValue) GetTime2() *ItemsMetricValuesValueTime2 {
	return s.Time2
}

func (s *ItemsMetricValuesValue) GetAvg() *ItemsMetricValuesValueAvg {
	return s.Avg
}

func (s *ItemsMetricValuesValue) GetSum() *ItemsMetricValuesValueSum {
	return s.Sum
}

func (s *ItemsMetricValuesValue) GetMax() *ItemsMetricValuesValueMax {
	return s.Max
}

func (s *ItemsMetricValuesValue) SetMetricCode(v string) *ItemsMetricValuesValue {
	s.MetricCode = &v
	return s
}

func (s *ItemsMetricValuesValue) SetMetricName(v string) *ItemsMetricValuesValue {
	s.MetricName = &v
	return s
}

func (s *ItemsMetricValuesValue) SetPrimary(v bool) *ItemsMetricValuesValue {
	s.Primary = &v
	return s
}

func (s *ItemsMetricValuesValue) SetTime2(v *ItemsMetricValuesValueTime2) *ItemsMetricValuesValue {
	s.Time2 = v
	return s
}

func (s *ItemsMetricValuesValue) SetAvg(v *ItemsMetricValuesValueAvg) *ItemsMetricValuesValue {
	s.Avg = v
	return s
}

func (s *ItemsMetricValuesValue) SetSum(v *ItemsMetricValuesValueSum) *ItemsMetricValuesValue {
	s.Sum = v
	return s
}

func (s *ItemsMetricValuesValue) SetMax(v *ItemsMetricValuesValueMax) *ItemsMetricValuesValue {
	s.Max = v
	return s
}

func (s *ItemsMetricValuesValue) Validate() error {
	if s.Time2 != nil {
		if err := s.Time2.Validate(); err != nil {
			return err
		}
	}
	if s.Avg != nil {
		if err := s.Avg.Validate(); err != nil {
			return err
		}
	}
	if s.Sum != nil {
		if err := s.Sum.Validate(); err != nil {
			return err
		}
	}
	if s.Max != nil {
		if err := s.Max.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ItemsMetricValuesValueTime2 struct {
	// The sum of metric values across active query minute buckets for Time 2. The unit depends on MetricCode: count for QUERY_COUNT, seconds for CPU_COST, and GB (1 GB = 1024³ bytes) for SHUFFLE_SIZE, PEAK_MEMORY, and SCAN_SIZE.
	//
	// example:
	//
	// 180
	SumValue *float64 `json:"SumValue,omitempty" xml:"SumValue,omitempty"`
	// The display string of the total sum for Time 2, with the unit included.
	//
	// example:
	//
	// 180s
	SumDisplayValue *string `json:"SumDisplayValue,omitempty" xml:"SumDisplayValue,omitempty"`
	// The average value across active query minute buckets for Time 2, calculated as the total sum divided by the number of minute buckets that contain queries for this Pattern. The unit depends on MetricCode: count for QUERY_COUNT, seconds for CPU_COST, and GB (1 GB = 1024³ bytes) for SHUFFLE_SIZE, PEAK_MEMORY, and SCAN_SIZE.
	//
	// example:
	//
	// 3
	AvgValue *float64 `json:"AvgValue,omitempty" xml:"AvgValue,omitempty"`
	// The display string of the average value across active query minute buckets for Time 2, with the unit included.
	//
	// example:
	//
	// 3s
	AvgDisplayValue *string `json:"AvgDisplayValue,omitempty" xml:"AvgDisplayValue,omitempty"`
	// The maximum metric value in a single minute bucket for Time 2. The unit depends on MetricCode: count for QUERY_COUNT, seconds for CPU_COST, and GB (1 GB = 1024³ bytes) for SHUFFLE_SIZE, PEAK_MEMORY, and SCAN_SIZE.
	//
	// example:
	//
	// 9
	MaxValue *float64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The display string of the peak value in a single minute bucket for Time 2, with the unit included.
	//
	// example:
	//
	// 9s
	MaxDisplayValue *string `json:"MaxDisplayValue,omitempty" xml:"MaxDisplayValue,omitempty"`
	// The percentage of this Pattern\\"s Time 2 total sum relative to the total sum of all results before dimension filtering in the current report. A value of 10 indicates 10%.
	//
	// example:
	//
	// 10
	SumRatioPercent *float64 `json:"SumRatioPercent,omitempty" xml:"SumRatioPercent,omitempty"`
	// The percentage of this Pattern\\"s Time 2 average value relative to the sum of average values across all Patterns before dimension filtering in the current report. A value of 10 indicates 10%.
	//
	// example:
	//
	// 10
	AvgRatioPercent *float64 `json:"AvgRatioPercent,omitempty" xml:"AvgRatioPercent,omitempty"`
}

func (s ItemsMetricValuesValueTime2) String() string {
	return dara.Prettify(s)
}

func (s ItemsMetricValuesValueTime2) GoString() string {
	return s.String()
}

func (s *ItemsMetricValuesValueTime2) GetSumValue() *float64 {
	return s.SumValue
}

func (s *ItemsMetricValuesValueTime2) GetSumDisplayValue() *string {
	return s.SumDisplayValue
}

func (s *ItemsMetricValuesValueTime2) GetAvgValue() *float64 {
	return s.AvgValue
}

func (s *ItemsMetricValuesValueTime2) GetAvgDisplayValue() *string {
	return s.AvgDisplayValue
}

func (s *ItemsMetricValuesValueTime2) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *ItemsMetricValuesValueTime2) GetMaxDisplayValue() *string {
	return s.MaxDisplayValue
}

func (s *ItemsMetricValuesValueTime2) GetSumRatioPercent() *float64 {
	return s.SumRatioPercent
}

func (s *ItemsMetricValuesValueTime2) GetAvgRatioPercent() *float64 {
	return s.AvgRatioPercent
}

func (s *ItemsMetricValuesValueTime2) SetSumValue(v float64) *ItemsMetricValuesValueTime2 {
	s.SumValue = &v
	return s
}

func (s *ItemsMetricValuesValueTime2) SetSumDisplayValue(v string) *ItemsMetricValuesValueTime2 {
	s.SumDisplayValue = &v
	return s
}

func (s *ItemsMetricValuesValueTime2) SetAvgValue(v float64) *ItemsMetricValuesValueTime2 {
	s.AvgValue = &v
	return s
}

func (s *ItemsMetricValuesValueTime2) SetAvgDisplayValue(v string) *ItemsMetricValuesValueTime2 {
	s.AvgDisplayValue = &v
	return s
}

func (s *ItemsMetricValuesValueTime2) SetMaxValue(v float64) *ItemsMetricValuesValueTime2 {
	s.MaxValue = &v
	return s
}

func (s *ItemsMetricValuesValueTime2) SetMaxDisplayValue(v string) *ItemsMetricValuesValueTime2 {
	s.MaxDisplayValue = &v
	return s
}

func (s *ItemsMetricValuesValueTime2) SetSumRatioPercent(v float64) *ItemsMetricValuesValueTime2 {
	s.SumRatioPercent = &v
	return s
}

func (s *ItemsMetricValuesValueTime2) SetAvgRatioPercent(v float64) *ItemsMetricValuesValueTime2 {
	s.AvgRatioPercent = &v
	return s
}

func (s *ItemsMetricValuesValueTime2) Validate() error {
	return dara.Validate(s)
}

type ItemsMetricValuesValueAvg struct {
	// The change rate of the average value across active query minute buckets, calculated as (Time 2 value − Time 1 value) / Time 1 value × 100. A value of 200 indicates a 200% increase. When the Time 1 value is 0, a finite change rate cannot be calculated. This field may not be returned and must not be treated as 0%.
	//
	// example:
	//
	// 200
	ChangeRatePercent *float64 `json:"ChangeRatePercent,omitempty" xml:"ChangeRatePercent,omitempty"`
	// The average value across active query minute buckets for Time 1. The unit depends on MetricCode: count for QUERY_COUNT, seconds for CPU_COST, and GB (1 GB = 1024³ bytes) for SHUFFLE_SIZE, PEAK_MEMORY, and SCAN_SIZE.
	//
	// example:
	//
	// 1
	Time1Value *float64 `json:"Time1Value,omitempty" xml:"Time1Value,omitempty"`
	// The display string of the average value across active query minute buckets for Time 1, with the unit included.
	//
	// example:
	//
	// 1s
	Time1DisplayValue *string `json:"Time1DisplayValue,omitempty" xml:"Time1DisplayValue,omitempty"`
	// The percentage of this Pattern\\"s Time 1 average value across active query minute buckets relative to the sum of the corresponding statistics for all results before dimension filtering in the current report. A value of 10 indicates 10%.
	//
	// example:
	//
	// 10
	Time1RatioPercent *float64 `json:"Time1RatioPercent,omitempty" xml:"Time1RatioPercent,omitempty"`
	// The average value across active query minute buckets for Time 2. The unit depends on MetricCode: count for QUERY_COUNT, seconds for CPU_COST, and GB (1 GB = 1024³ bytes) for SHUFFLE_SIZE, PEAK_MEMORY, and SCAN_SIZE.
	//
	// example:
	//
	// 3
	Time2Value *float64 `json:"Time2Value,omitempty" xml:"Time2Value,omitempty"`
	// The display string of the average value across active query minute buckets for Time 2, with the unit included.
	//
	// example:
	//
	// 3s
	Time2DisplayValue *string `json:"Time2DisplayValue,omitempty" xml:"Time2DisplayValue,omitempty"`
	// The percentage of this Pattern\\"s Time 2 average value across active query minute buckets relative to the sum of the corresponding statistics for all results before dimension filtering in the current report. A value of 10 indicates 10%.
	//
	// example:
	//
	// 10
	Time2RatioPercent *float64 `json:"Time2RatioPercent,omitempty" xml:"Time2RatioPercent,omitempty"`
}

func (s ItemsMetricValuesValueAvg) String() string {
	return dara.Prettify(s)
}

func (s ItemsMetricValuesValueAvg) GoString() string {
	return s.String()
}

func (s *ItemsMetricValuesValueAvg) GetChangeRatePercent() *float64 {
	return s.ChangeRatePercent
}

func (s *ItemsMetricValuesValueAvg) GetTime1Value() *float64 {
	return s.Time1Value
}

func (s *ItemsMetricValuesValueAvg) GetTime1DisplayValue() *string {
	return s.Time1DisplayValue
}

func (s *ItemsMetricValuesValueAvg) GetTime1RatioPercent() *float64 {
	return s.Time1RatioPercent
}

func (s *ItemsMetricValuesValueAvg) GetTime2Value() *float64 {
	return s.Time2Value
}

func (s *ItemsMetricValuesValueAvg) GetTime2DisplayValue() *string {
	return s.Time2DisplayValue
}

func (s *ItemsMetricValuesValueAvg) GetTime2RatioPercent() *float64 {
	return s.Time2RatioPercent
}

func (s *ItemsMetricValuesValueAvg) SetChangeRatePercent(v float64) *ItemsMetricValuesValueAvg {
	s.ChangeRatePercent = &v
	return s
}

func (s *ItemsMetricValuesValueAvg) SetTime1Value(v float64) *ItemsMetricValuesValueAvg {
	s.Time1Value = &v
	return s
}

func (s *ItemsMetricValuesValueAvg) SetTime1DisplayValue(v string) *ItemsMetricValuesValueAvg {
	s.Time1DisplayValue = &v
	return s
}

func (s *ItemsMetricValuesValueAvg) SetTime1RatioPercent(v float64) *ItemsMetricValuesValueAvg {
	s.Time1RatioPercent = &v
	return s
}

func (s *ItemsMetricValuesValueAvg) SetTime2Value(v float64) *ItemsMetricValuesValueAvg {
	s.Time2Value = &v
	return s
}

func (s *ItemsMetricValuesValueAvg) SetTime2DisplayValue(v string) *ItemsMetricValuesValueAvg {
	s.Time2DisplayValue = &v
	return s
}

func (s *ItemsMetricValuesValueAvg) SetTime2RatioPercent(v float64) *ItemsMetricValuesValueAvg {
	s.Time2RatioPercent = &v
	return s
}

func (s *ItemsMetricValuesValueAvg) Validate() error {
	return dara.Validate(s)
}

type ItemsMetricValuesValueSum struct {
	// The change rate of the sum of metric values across active query minute buckets, calculated as (Time 2 value − Time 1 value) / Time 1 value × 100. A value of 200 indicates a 200% increase. When the Time 1 value is 0, a finite change rate cannot be calculated. This field may not be returned and must not be treated as 0%.
	//
	// example:
	//
	// 200
	ChangeRatePercent *float64 `json:"ChangeRatePercent,omitempty" xml:"ChangeRatePercent,omitempty"`
	// The sum of metric values across active query minute buckets for Time 1. The unit depends on MetricCode: count for QUERY_COUNT, seconds for CPU_COST, and GB (1 GB = 1024³ bytes) for SHUFFLE_SIZE, PEAK_MEMORY, and SCAN_SIZE.
	//
	// example:
	//
	// 60
	Time1Value *float64 `json:"Time1Value,omitempty" xml:"Time1Value,omitempty"`
	// The display string of the sum of metric values across active query minute buckets for Time 1, with the unit included.
	//
	// example:
	//
	// 60s
	Time1DisplayValue *string `json:"Time1DisplayValue,omitempty" xml:"Time1DisplayValue,omitempty"`
	// The percentage of this Pattern\\"s Time 1 sum of metric values across active query minute buckets relative to the sum of the corresponding statistics for all results before dimension filtering in the current report. A value of 10 indicates 10%.
	//
	// example:
	//
	// 10
	Time1RatioPercent *float64 `json:"Time1RatioPercent,omitempty" xml:"Time1RatioPercent,omitempty"`
	// The sum of metric values across active query minute buckets for Time 2. The unit depends on MetricCode: count for QUERY_COUNT, seconds for CPU_COST, and GB (1 GB = 1024³ bytes) for SHUFFLE_SIZE, PEAK_MEMORY, and SCAN_SIZE.
	//
	// example:
	//
	// 180
	Time2Value *float64 `json:"Time2Value,omitempty" xml:"Time2Value,omitempty"`
	// The display string of the sum of metric values across active query minute buckets for Time 2, with the unit included.
	//
	// example:
	//
	// 180s
	Time2DisplayValue *string `json:"Time2DisplayValue,omitempty" xml:"Time2DisplayValue,omitempty"`
	// The percentage of this Pattern\\"s Time 2 sum of metric values across active query minute buckets relative to the sum of the corresponding statistics for all results before dimension filtering in the current report. A value of 10 indicates 10%.
	//
	// example:
	//
	// 10
	Time2RatioPercent *float64 `json:"Time2RatioPercent,omitempty" xml:"Time2RatioPercent,omitempty"`
}

func (s ItemsMetricValuesValueSum) String() string {
	return dara.Prettify(s)
}

func (s ItemsMetricValuesValueSum) GoString() string {
	return s.String()
}

func (s *ItemsMetricValuesValueSum) GetChangeRatePercent() *float64 {
	return s.ChangeRatePercent
}

func (s *ItemsMetricValuesValueSum) GetTime1Value() *float64 {
	return s.Time1Value
}

func (s *ItemsMetricValuesValueSum) GetTime1DisplayValue() *string {
	return s.Time1DisplayValue
}

func (s *ItemsMetricValuesValueSum) GetTime1RatioPercent() *float64 {
	return s.Time1RatioPercent
}

func (s *ItemsMetricValuesValueSum) GetTime2Value() *float64 {
	return s.Time2Value
}

func (s *ItemsMetricValuesValueSum) GetTime2DisplayValue() *string {
	return s.Time2DisplayValue
}

func (s *ItemsMetricValuesValueSum) GetTime2RatioPercent() *float64 {
	return s.Time2RatioPercent
}

func (s *ItemsMetricValuesValueSum) SetChangeRatePercent(v float64) *ItemsMetricValuesValueSum {
	s.ChangeRatePercent = &v
	return s
}

func (s *ItemsMetricValuesValueSum) SetTime1Value(v float64) *ItemsMetricValuesValueSum {
	s.Time1Value = &v
	return s
}

func (s *ItemsMetricValuesValueSum) SetTime1DisplayValue(v string) *ItemsMetricValuesValueSum {
	s.Time1DisplayValue = &v
	return s
}

func (s *ItemsMetricValuesValueSum) SetTime1RatioPercent(v float64) *ItemsMetricValuesValueSum {
	s.Time1RatioPercent = &v
	return s
}

func (s *ItemsMetricValuesValueSum) SetTime2Value(v float64) *ItemsMetricValuesValueSum {
	s.Time2Value = &v
	return s
}

func (s *ItemsMetricValuesValueSum) SetTime2DisplayValue(v string) *ItemsMetricValuesValueSum {
	s.Time2DisplayValue = &v
	return s
}

func (s *ItemsMetricValuesValueSum) SetTime2RatioPercent(v float64) *ItemsMetricValuesValueSum {
	s.Time2RatioPercent = &v
	return s
}

func (s *ItemsMetricValuesValueSum) Validate() error {
	return dara.Validate(s)
}

type ItemsMetricValuesValueMax struct {
	// The change rate of the peak value in a single minute bucket, calculated as (Time 2 value − Time 1 value) / Time 1 value × 100. A value of 200 indicates a 200% increase. When the Time 1 value is 0, a finite change rate cannot be calculated. This field may not be returned and must not be treated as 0%.
	//
	// example:
	//
	// 200
	ChangeRatePercent *float64 `json:"ChangeRatePercent,omitempty" xml:"ChangeRatePercent,omitempty"`
	// The peak value in a single minute bucket for Time 1. The unit depends on MetricCode: count for QUERY_COUNT, seconds for CPU_COST, and GB (1 GB = 1024³ bytes) for SHUFFLE_SIZE, PEAK_MEMORY, and SCAN_SIZE.
	//
	// example:
	//
	// 3
	Time1Value *float64 `json:"Time1Value,omitempty" xml:"Time1Value,omitempty"`
	// The display string of the peak value in a single minute bucket for Time 1, with the unit included.
	//
	// example:
	//
	// 3s
	Time1DisplayValue *string `json:"Time1DisplayValue,omitempty" xml:"Time1DisplayValue,omitempty"`
	// The peak value in a single minute bucket for Time 2. The unit depends on MetricCode: count for QUERY_COUNT, seconds for CPU_COST, and GB (1 GB = 1024³ bytes) for SHUFFLE_SIZE, PEAK_MEMORY, and SCAN_SIZE.
	//
	// example:
	//
	// 9
	Time2Value *float64 `json:"Time2Value,omitempty" xml:"Time2Value,omitempty"`
	// The display string of the peak value in a single minute bucket for Time 2, with the unit included.
	//
	// example:
	//
	// 9s
	Time2DisplayValue *string `json:"Time2DisplayValue,omitempty" xml:"Time2DisplayValue,omitempty"`
}

func (s ItemsMetricValuesValueMax) String() string {
	return dara.Prettify(s)
}

func (s ItemsMetricValuesValueMax) GoString() string {
	return s.String()
}

func (s *ItemsMetricValuesValueMax) GetChangeRatePercent() *float64 {
	return s.ChangeRatePercent
}

func (s *ItemsMetricValuesValueMax) GetTime1Value() *float64 {
	return s.Time1Value
}

func (s *ItemsMetricValuesValueMax) GetTime1DisplayValue() *string {
	return s.Time1DisplayValue
}

func (s *ItemsMetricValuesValueMax) GetTime2Value() *float64 {
	return s.Time2Value
}

func (s *ItemsMetricValuesValueMax) GetTime2DisplayValue() *string {
	return s.Time2DisplayValue
}

func (s *ItemsMetricValuesValueMax) SetChangeRatePercent(v float64) *ItemsMetricValuesValueMax {
	s.ChangeRatePercent = &v
	return s
}

func (s *ItemsMetricValuesValueMax) SetTime1Value(v float64) *ItemsMetricValuesValueMax {
	s.Time1Value = &v
	return s
}

func (s *ItemsMetricValuesValueMax) SetTime1DisplayValue(v string) *ItemsMetricValuesValueMax {
	s.Time1DisplayValue = &v
	return s
}

func (s *ItemsMetricValuesValueMax) SetTime2Value(v float64) *ItemsMetricValuesValueMax {
	s.Time2Value = &v
	return s
}

func (s *ItemsMetricValuesValueMax) SetTime2DisplayValue(v string) *ItemsMetricValuesValueMax {
	s.Time2DisplayValue = &v
	return s
}

func (s *ItemsMetricValuesValueMax) Validate() error {
	return dara.Validate(s)
}
