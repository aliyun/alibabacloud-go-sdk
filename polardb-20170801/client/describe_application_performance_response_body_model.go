// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationPerformanceResponseBody
	GetApplicationId() *string
	SetApplicationType(v string) *DescribeApplicationPerformanceResponseBody
	GetApplicationType() *string
	SetEndTime(v string) *DescribeApplicationPerformanceResponseBody
	GetEndTime() *string
	SetInterval(v string) *DescribeApplicationPerformanceResponseBody
	GetInterval() *string
	SetPerformanceKeys(v *DescribeApplicationPerformanceResponseBodyPerformanceKeys) *DescribeApplicationPerformanceResponseBody
	GetPerformanceKeys() *DescribeApplicationPerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeApplicationPerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeApplicationPerformanceResponseBody
	GetStartTime() *string
}

type DescribeApplicationPerformanceResponseBody struct {
	// The application cluster ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type.
	//
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The end time of the query. The time is in UTC and follows the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2025-05-23T02:21:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The performance metrics.
	PerformanceKeys *DescribeApplicationPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time is in UTC and follows the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2022-11-15T16:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApplicationPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPerformanceResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationPerformanceResponseBody) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribeApplicationPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApplicationPerformanceResponseBody) GetInterval() *string {
	return s.Interval
}

func (s *DescribeApplicationPerformanceResponseBody) GetPerformanceKeys() *DescribeApplicationPerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeApplicationPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApplicationPerformanceResponseBody) SetApplicationId(v string) *DescribeApplicationPerformanceResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBody) SetApplicationType(v string) *DescribeApplicationPerformanceResponseBody {
	s.ApplicationType = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBody) SetEndTime(v string) *DescribeApplicationPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBody) SetInterval(v string) *DescribeApplicationPerformanceResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBody) SetPerformanceKeys(v *DescribeApplicationPerformanceResponseBodyPerformanceKeys) *DescribeApplicationPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeApplicationPerformanceResponseBody) SetRequestId(v string) *DescribeApplicationPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBody) SetStartTime(v string) *DescribeApplicationPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBody) Validate() error {
	if s.PerformanceKeys != nil {
		if err := s.PerformanceKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationPerformanceResponseBodyPerformanceKeys struct {
	PerformanceItem []*DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem `json:"PerformanceItem,omitempty" xml:"PerformanceItem,omitempty" type:"Repeated"`
}

func (s DescribeApplicationPerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeys) GetPerformanceItem() []*DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem {
	return s.PerformanceItem
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeys) SetPerformanceItem(v []*DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) *DescribeApplicationPerformanceResponseBodyPerformanceKeys {
	s.PerformanceItem = v
	return s
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeys) Validate() error {
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

type DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem struct {
	ComponentId   *string                                                                         `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	ComponentType *string                                                                         `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	Measurement   *string                                                                         `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	MetricName    *string                                                                         `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Points        *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
}

func (s DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) GetComponentId() *string {
	return s.ComponentId
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) GetMeasurement() *string {
	return s.Measurement
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) GetPoints() *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	return s.Points
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) SetComponentId(v string) *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.ComponentId = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) SetComponentType(v string) *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.ComponentType = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMetricName(v string) *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) SetPoints(v *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItem) Validate() error {
	if s.Points != nil {
		if err := s.Points.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue,omitempty" xml:"PerformanceItemValue,omitempty" type:"Repeated"`
}

func (s DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GetPerformanceItemValue() []*DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	return s.PerformanceItemValue
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) Validate() error {
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

type DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	Step      *int64  `json:"Step,omitempty" xml:"Step,omitempty"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetStep() *int64 {
	return s.Step
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetValue() *string {
	return s.Value
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetStep(v int64) *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Step = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

func (s *DescribeApplicationPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) Validate() error {
	return dara.Validate(s)
}
