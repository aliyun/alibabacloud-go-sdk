// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribePolarFsPerformanceResponseBody
	GetEndTime() *string
	SetPerformanceKeys(v *DescribePolarFsPerformanceResponseBodyPerformanceKeys) *DescribePolarFsPerformanceResponseBody
	GetPerformanceKeys() *DescribePolarFsPerformanceResponseBodyPerformanceKeys
	SetPolarFsInstanceId(v string) *DescribePolarFsPerformanceResponseBody
	GetPolarFsInstanceId() *string
	SetRequestId(v string) *DescribePolarFsPerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribePolarFsPerformanceResponseBody
	GetStartTime() *string
}

type DescribePolarFsPerformanceResponseBody struct {
	// The end time of the query. The time is in UTC.
	//
	// example:
	//
	// 2024-10-07T03:04:00Z
	EndTime         *string                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PerformanceKeys *DescribePolarFsPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// The PolarFS instance ID.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time is in UTC.
	//
	// example:
	//
	// 2024-10-07T02:04:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePolarFsPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarFsPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePolarFsPerformanceResponseBody) GetPerformanceKeys() *DescribePolarFsPerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribePolarFsPerformanceResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarFsPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePolarFsPerformanceResponseBody) SetEndTime(v string) *DescribePolarFsPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribePolarFsPerformanceResponseBody) SetPerformanceKeys(v *DescribePolarFsPerformanceResponseBodyPerformanceKeys) *DescribePolarFsPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribePolarFsPerformanceResponseBody) SetPolarFsInstanceId(v string) *DescribePolarFsPerformanceResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsPerformanceResponseBody) SetRequestId(v string) *DescribePolarFsPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarFsPerformanceResponseBody) SetStartTime(v string) *DescribePolarFsPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribePolarFsPerformanceResponseBody) Validate() error {
	if s.PerformanceKeys != nil {
		if err := s.PerformanceKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarFsPerformanceResponseBodyPerformanceKeys struct {
	PerformanceItem []*DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem `json:"PerformanceItem,omitempty" xml:"PerformanceItem,omitempty" type:"Repeated"`
}

func (s DescribePolarFsPerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeys) GetPerformanceItem() []*DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem {
	return s.PerformanceItem
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeys) SetPerformanceItem(v []*DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem) *DescribePolarFsPerformanceResponseBodyPerformanceKeys {
	s.PerformanceItem = v
	return s
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeys) Validate() error {
	if s.PerformanceItem != nil {
		for _, item := range s.PerformanceItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem struct {
	Measurement *string                                                                     `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	MetricName  *string                                                                     `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Points      *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
}

func (s DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem) GetMeasurement() *string {
	return s.Measurement
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem) GetPoints() *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	return s.Points
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMetricName(v string) *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem) SetPoints(v *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItem) Validate() error {
	if s.Points != nil {
		if err := s.Points.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue,omitempty" xml:"PerformanceItemValue,omitempty" type:"Repeated"`
}

func (s DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GetPerformanceItemValue() []*DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	return s.PerformanceItemValue
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) Validate() error {
	if s.PerformanceItemValue != nil {
		for _, item := range s.PerformanceItemValue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetValue() *string {
	return s.Value
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

func (s *DescribePolarFsPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) Validate() error {
	return dara.Validate(s)
}
