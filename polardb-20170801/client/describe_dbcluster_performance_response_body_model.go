// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody
	GetDBClusterId() *string
	SetDBType(v string) *DescribeDBClusterPerformanceResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBClusterPerformanceResponseBody
	GetDBVersion() *string
	SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody
	GetEndTime() *string
	SetPerformanceKeys(v *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) *DescribeDBClusterPerformanceResponseBody
	GetPerformanceKeys() *DescribeDBClusterPerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeDBClusterPerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDBClusterPerformanceResponseBody
	GetStartTime() *string
}

type DescribeDBClusterPerformanceResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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
	PerformanceKeys *DescribeDBClusterPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 35D3E3DA-4650-407A-BFF5-59BFF1******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-09-23T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterPerformanceResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClusterPerformanceResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClusterPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBClusterPerformanceResponseBody) GetPerformanceKeys() *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeDBClusterPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBType(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBVersion(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetPerformanceKeys(v *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) *DescribeDBClusterPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) Validate() error {
	if s.PerformanceKeys != nil {
		if err := s.PerformanceKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeys struct {
	PerformanceItem []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem `json:"PerformanceItem,omitempty" xml:"PerformanceItem,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) GetPerformanceItem() []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem {
	return s.PerformanceItem
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetPerformanceItem(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.PerformanceItem = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) Validate() error {
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

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem struct {
	// The ID of the cluster node.
	//
	// >  The value of this parameter is not returned if the `Key` parameter is set to `PolarDBDiskUsage`.
	//
	// example:
	//
	// pi-*************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The performance metrics that are returned.
	//
	// example:
	//
	// PolarDBDiskUsage
	Measurement *string `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	// The name of the performance metric.
	//
	// example:
	//
	// mean_data_size
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The performance metrics.
	Points *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) GetMeasurement() *string {
	return s.Measurement
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) GetPoints() *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	return s.Points
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) SetDBNodeId(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMetricName(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) SetPoints(v *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) Validate() error {
	if s.Points != nil {
		if err := s.Points.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue,omitempty" xml:"PerformanceItemValue,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GetPerformanceItemValue() []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	return s.PerformanceItemValue
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) Validate() error {
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

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
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
	// 42.38
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) Validate() error {
	return dara.Validate(s)
}
