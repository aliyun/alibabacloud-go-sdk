// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBProxyPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBProxyPerformanceResponseBody
	GetDBClusterId() *string
	SetDBType(v string) *DescribeDBProxyPerformanceResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBProxyPerformanceResponseBody
	GetDBVersion() *string
	SetEndTime(v string) *DescribeDBProxyPerformanceResponseBody
	GetEndTime() *string
	SetPerformanceKeys(v *DescribeDBProxyPerformanceResponseBodyPerformanceKeys) *DescribeDBProxyPerformanceResponseBody
	GetPerformanceKeys() *DescribeDBProxyPerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeDBProxyPerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDBProxyPerformanceResponseBody
	GetStartTime() *string
}

type DescribeDBProxyPerformanceResponseBody struct {
	// The ID of the cluster.
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
	// The end time of the query. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-09-23T01:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Details about the performance metrics.
	PerformanceKeys *DescribeDBProxyPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 35D3E3DA-4650-407A-BFF5-59BFF1******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-09-23T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBProxyPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBProxyPerformanceResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBProxyPerformanceResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBProxyPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBProxyPerformanceResponseBody) GetPerformanceKeys() *DescribeDBProxyPerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeDBProxyPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBProxyPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBProxyPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBProxyPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) SetDBType(v string) *DescribeDBProxyPerformanceResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) SetDBVersion(v string) *DescribeDBProxyPerformanceResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) SetEndTime(v string) *DescribeDBProxyPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) SetPerformanceKeys(v *DescribeDBProxyPerformanceResponseBodyPerformanceKeys) *DescribeDBProxyPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) SetRequestId(v string) *DescribeDBProxyPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) SetStartTime(v string) *DescribeDBProxyPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBProxyPerformanceResponseBodyPerformanceKeys struct {
	PerformanceItem []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem `json:"PerformanceItem,omitempty" xml:"PerformanceItem,omitempty" type:"Repeated"`
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeys) GetPerformanceItem() []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem {
	return s.PerformanceItem
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeys) SetPerformanceItem(v []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) *DescribeDBProxyPerformanceResponseBodyPerformanceKeys {
	s.PerformanceItem = v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeys) Validate() error {
	return dara.Validate(s)
}

type DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem struct {
	// The ID of the node.
	//
	// example:
	//
	// pi-*************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The performance metric.
	//
	// example:
	//
	// PolarProxy_CpuUsage
	Measurement *string `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	// The name of the performance metric.
	//
	// example:
	//
	// service_connections_ps
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The list of the performance metrics.
	Points *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) GetMeasurement() *string {
	return s.Measurement
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) GetPoints() *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	return s.Points
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) SetDBNodeId(v string) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMetricName(v string) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) SetPoints(v *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItem) Validate() error {
	return dara.Validate(s)
}

type DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue,omitempty" xml:"PerformanceItemValue,omitempty" type:"Repeated"`
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GetPerformanceItemValue() []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	return s.PerformanceItemValue
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) Validate() error {
	return dara.Validate(s)
}

type DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	// The time when the metric value was collected. This value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1600822800000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetValue() *string {
	return s.Value
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) Validate() error {
	return dara.Validate(s)
}
