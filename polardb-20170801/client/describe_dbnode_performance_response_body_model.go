// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBNodePerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBNodeId(v string) *DescribeDBNodePerformanceResponseBody
	GetDBNodeId() *string
	SetDBType(v string) *DescribeDBNodePerformanceResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBNodePerformanceResponseBody
	GetDBVersion() *string
	SetEndTime(v string) *DescribeDBNodePerformanceResponseBody
	GetEndTime() *string
	SetPerformanceKeys(v *DescribeDBNodePerformanceResponseBodyPerformanceKeys) *DescribeDBNodePerformanceResponseBody
	GetPerformanceKeys() *DescribeDBNodePerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeDBNodePerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDBNodePerformanceResponseBody
	GetStartTime() *string
}

type DescribeDBNodePerformanceResponseBody struct {
	// The ID of the cluster node.
	//
	// example:
	//
	// pi-*****************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-09-23T01:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The cluster performance metrics.
	PerformanceKeys *DescribeDBNodePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E2FDB684-751D-424D-98B9-704BEA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-09-23T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBNodePerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBody) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBNodePerformanceResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBNodePerformanceResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBNodePerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBNodePerformanceResponseBody) GetPerformanceKeys() *DescribeDBNodePerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeDBNodePerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBNodePerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBNodePerformanceResponseBody) SetDBNodeId(v string) *DescribeDBNodePerformanceResponseBody {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetDBType(v string) *DescribeDBNodePerformanceResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetDBVersion(v string) *DescribeDBNodePerformanceResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetEndTime(v string) *DescribeDBNodePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetPerformanceKeys(v *DescribeDBNodePerformanceResponseBodyPerformanceKeys) *DescribeDBNodePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetRequestId(v string) *DescribeDBNodePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetStartTime(v string) *DescribeDBNodePerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) Validate() error {
	if s.PerformanceKeys != nil {
		if err := s.PerformanceKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeys struct {
	PerformanceItem []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem `json:"PerformanceItem,omitempty" xml:"PerformanceItem,omitempty" type:"Repeated"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeys) GetPerformanceItem() []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	return s.PerformanceItem
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeys) SetPerformanceItem(v []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) *DescribeDBNodePerformanceResponseBodyPerformanceKeys {
	s.PerformanceItem = v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeys) Validate() error {
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

type DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem struct {
	// The performance metrics that you want to query.
	//
	// example:
	//
	// PolarDBDiskUsage
	Measurement *string `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	// The name of the performance metric.
	//
	// example:
	//
	// mean_sys_dir_size
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The performance metrics.
	Points *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) GetMeasurement() *string {
	return s.Measurement
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) GetPoints() *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	return s.Points
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) SetMetricName(v string) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) SetPoints(v *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) Validate() error {
	if s.Points != nil {
		if err := s.Points.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue,omitempty" xml:"PerformanceItemValue,omitempty" type:"Repeated"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GetPerformanceItemValue() []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	return s.PerformanceItemValue
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) Validate() error {
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

type DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	// The timestamp of the metric. This value is a UNIX timestamp. Unit: millisecond.
	//
	// example:
	//
	// 1600822800000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 9.33
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetValue() *string {
	return s.Value
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) Validate() error {
	return dara.Validate(s)
}
