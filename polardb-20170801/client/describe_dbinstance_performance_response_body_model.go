// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstancePerformanceResponseBody
	GetDBInstanceId() *string
	SetDBType(v string) *DescribeDBInstancePerformanceResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBInstancePerformanceResponseBody
	GetDBVersion() *string
	SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody
	GetEndTime() *string
	SetEngine(v string) *DescribeDBInstancePerformanceResponseBody
	GetEngine() *string
	SetPerformanceKeys(v *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) *DescribeDBInstancePerformanceResponseBody
	GetPerformanceKeys() *DescribeDBInstancePerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeDBInstancePerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDBInstancePerformanceResponseBody
	GetStartTime() *string
}

type DescribeDBInstancePerformanceResponseBody struct {
	// example:
	//
	// pi-*************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// example:
	//
	// 2020-09-23T01:01Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// POLARDB
	Engine          *string                                                   `json:"Engine,omitempty" xml:"Engine,omitempty"`
	PerformanceKeys *DescribeDBInstancePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// example:
	//
	// F2A9EFA7-915F-4572-8299-85A307******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2020-09-23T01:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancePerformanceResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBInstancePerformanceResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBInstancePerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstancePerformanceResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancePerformanceResponseBody) GetPerformanceKeys() *DescribeDBInstancePerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeDBInstancePerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancePerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstancePerformanceResponseBody) SetDBInstanceId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetDBType(v string) *DescribeDBInstancePerformanceResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetDBVersion(v string) *DescribeDBInstancePerformanceResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEngine(v string) *DescribeDBInstancePerformanceResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetPerformanceKeys(v *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) *DescribeDBInstancePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetRequestId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetStartTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeys struct {
	PerformanceItem []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem `json:"PerformanceItem,omitempty" xml:"PerformanceItem,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) GetPerformanceItem() []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem {
	return s.PerformanceItem
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) SetPerformanceItem(v []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem) *DescribeDBInstancePerformanceResponseBodyPerformanceKeys {
	s.PerformanceItem = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem struct {
	// example:
	//
	// PolarDBDiskUsage
	Measurement *string `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	// example:
	//
	// mean_data_size
	MetricName *string                                                                        `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Points     *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem) GetMeasurement() *string {
	return s.Measurement
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem) GetPoints() *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	return s.Points
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem) SetMetricName(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem) SetPoints(v *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItem) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue,omitempty" xml:"PerformanceItemValue,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GetPerformanceItemValue() []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	return s.PerformanceItemValue
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	// example:
	//
	// 1737424822
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 12.33
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) Validate() error {
	return dara.Validate(s)
}
