// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationComponentPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationComponentPerformanceResponseBody
	GetApplicationId() *string
	SetApplicationType(v string) *DescribeApplicationComponentPerformanceResponseBody
	GetApplicationType() *string
	SetApplicationVersion(v string) *DescribeApplicationComponentPerformanceResponseBody
	GetApplicationVersion() *string
	SetComponentId(v string) *DescribeApplicationComponentPerformanceResponseBody
	GetComponentId() *string
	SetComponentType(v string) *DescribeApplicationComponentPerformanceResponseBody
	GetComponentType() *string
	SetEndTime(v string) *DescribeApplicationComponentPerformanceResponseBody
	GetEndTime() *string
	SetPerformanceKeys(v *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys) *DescribeApplicationComponentPerformanceResponseBody
	GetPerformanceKeys() *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeApplicationComponentPerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeApplicationComponentPerformanceResponseBody
	GetStartTime() *string
}

type DescribeApplicationComponentPerformanceResponseBody struct {
	// The ID of the application.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The type of the application.
	//
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The version of the application.
	//
	// example:
	//
	// 1.0.0
	ApplicationVersion *string `json:"ApplicationVersion,omitempty" xml:"ApplicationVersion,omitempty"`
	// The ID of the application component.
	//
	// example:
	//
	// pac-*******************
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The type of the application component.
	//
	// example:
	//
	// backend
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The end time of the query. The time is in UTC.
	//
	// example:
	//
	// 2020-09-23T02:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// []
	PerformanceKeys *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time is in UTC.
	//
	// example:
	//
	// 2020-09-23T01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApplicationComponentPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationComponentPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationComponentPerformanceResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationComponentPerformanceResponseBody) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribeApplicationComponentPerformanceResponseBody) GetApplicationVersion() *string {
	return s.ApplicationVersion
}

func (s *DescribeApplicationComponentPerformanceResponseBody) GetComponentId() *string {
	return s.ComponentId
}

func (s *DescribeApplicationComponentPerformanceResponseBody) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeApplicationComponentPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApplicationComponentPerformanceResponseBody) GetPerformanceKeys() *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeApplicationComponentPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationComponentPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApplicationComponentPerformanceResponseBody) SetApplicationId(v string) *DescribeApplicationComponentPerformanceResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBody) SetApplicationType(v string) *DescribeApplicationComponentPerformanceResponseBody {
	s.ApplicationType = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBody) SetApplicationVersion(v string) *DescribeApplicationComponentPerformanceResponseBody {
	s.ApplicationVersion = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBody) SetComponentId(v string) *DescribeApplicationComponentPerformanceResponseBody {
	s.ComponentId = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBody) SetComponentType(v string) *DescribeApplicationComponentPerformanceResponseBody {
	s.ComponentType = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBody) SetEndTime(v string) *DescribeApplicationComponentPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBody) SetPerformanceKeys(v *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys) *DescribeApplicationComponentPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBody) SetRequestId(v string) *DescribeApplicationComponentPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBody) SetStartTime(v string) *DescribeApplicationComponentPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBody) Validate() error {
	if s.PerformanceKeys != nil {
		if err := s.PerformanceKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys struct {
	PerformanceItem []*DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem `json:"PerformanceItem,omitempty" xml:"PerformanceItem,omitempty" type:"Repeated"`
}

func (s DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys) GetPerformanceItem() []*DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem {
	return s.PerformanceItem
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys) SetPerformanceItem(v []*DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem) *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys {
	s.PerformanceItem = v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeys) Validate() error {
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

type DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem struct {
	Measurement *string                                                                                  `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	MetricName  *string                                                                                  `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Points      *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
}

func (s DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem) GetMeasurement() *string {
	return s.Measurement
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem) GetPoints() *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	return s.Points
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMetricName(v string) *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem) SetPoints(v *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItem) Validate() error {
	if s.Points != nil {
		if err := s.Points.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue,omitempty" xml:"PerformanceItemValue,omitempty" type:"Repeated"`
}

func (s DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GetPerformanceItemValue() []*DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	return s.PerformanceItemValue
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) Validate() error {
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

type DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetValue() *string {
	return s.Value
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) Validate() error {
	return dara.Validate(s)
}
