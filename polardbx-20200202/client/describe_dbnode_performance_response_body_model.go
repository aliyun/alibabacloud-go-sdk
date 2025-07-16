// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBNodePerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBNodePerformanceResponseBody
	GetDBInstanceName() *string
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
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 2021-10-20T02:00Z
	EndTime         *string                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PerformanceKeys *DescribeDBNodePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// example:
	//
	// EFB5E10B-5268-170F-A378-9AF86CCEACC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-10-20T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBNodePerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
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

func (s *DescribeDBNodePerformanceResponseBody) SetDBInstanceName(v string) *DescribeDBNodePerformanceResponseBody {
	s.DBInstanceName = &v
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem struct {
	// example:
	//
	// pxc-i-********
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// example:
	//
	// Logic_TPS
	Measurement *string `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	// example:
	//
	// logic_tps
	MetricName *string                                                                    `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Points     *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) GetDBNodeId() *string {
	return s.DBNodeId
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

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) SetDBNodeId(v string) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.DBNodeId = &v
	return s
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	// example:
	//
	// 1600822800000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 42.38
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
