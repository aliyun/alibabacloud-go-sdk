// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DescribeAIDBClusterPerformanceResponseBody
	GetApiKey() *string
	SetDBClusterId(v string) *DescribeAIDBClusterPerformanceResponseBody
	GetDBClusterId() *string
	SetDBType(v string) *DescribeAIDBClusterPerformanceResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeAIDBClusterPerformanceResponseBody
	GetDBVersion() *string
	SetEndTime(v string) *DescribeAIDBClusterPerformanceResponseBody
	GetEndTime() *string
	SetInterval(v string) *DescribeAIDBClusterPerformanceResponseBody
	GetInterval() *string
	SetPerformanceKeys(v []*DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) *DescribeAIDBClusterPerformanceResponseBody
	GetPerformanceKeys() []*DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeAIDBClusterPerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeAIDBClusterPerformanceResponseBody
	GetStartTime() *string
}

type DescribeAIDBClusterPerformanceResponseBody struct {
	// The API key for the model service.
	//
	// example:
	//
	// xxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The ID of the database cluster.
	//
	// example:
	//
	// pc-a************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The type of the database engine. Only **polardb_ai*	- is supported.
	//
	// example:
	//
	// polardb_ai
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version number of the database AI engine.
	//
	// Example: 3.0
	//
	// example:
	//
	// 3.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The end of the time range that was queried. The time is in the `YYYY-MM-DDThh:mmZ` format and is in UTC.
	//
	// example:
	//
	// 2022-11-16T16:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the performance data. Valid values:
	//
	// - 60
	//
	// - 3600
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The details of the instance performance parameters.
	PerformanceKeys []*DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range that was queried. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is in UTC.
	//
	// example:
	//
	// 2022-11-15T16:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAIDBClusterPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterPerformanceResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeAIDBClusterPerformanceResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterPerformanceResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeAIDBClusterPerformanceResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeAIDBClusterPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAIDBClusterPerformanceResponseBody) GetInterval() *string {
	return s.Interval
}

func (s *DescribeAIDBClusterPerformanceResponseBody) GetPerformanceKeys() []*DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeAIDBClusterPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIDBClusterPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAIDBClusterPerformanceResponseBody) SetApiKey(v string) *DescribeAIDBClusterPerformanceResponseBody {
	s.ApiKey = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBody) SetDBClusterId(v string) *DescribeAIDBClusterPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBody) SetDBType(v string) *DescribeAIDBClusterPerformanceResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBody) SetDBVersion(v string) *DescribeAIDBClusterPerformanceResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeAIDBClusterPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBody) SetInterval(v string) *DescribeAIDBClusterPerformanceResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBody) SetPerformanceKeys(v []*DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) *DescribeAIDBClusterPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBody) SetRequestId(v string) *DescribeAIDBClusterPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBody) SetStartTime(v string) *DescribeAIDBClusterPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBody) Validate() error {
	if s.PerformanceKeys != nil {
		for _, item := range s.PerformanceKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys struct {
	// The ID of the cluster node.
	//
	// example:
	//
	// pi-*************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The performance metric.
	//
	// example:
	//
	// PolarDBAIModelCall
	Measurement *string `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	// The name of the specific performance metric.
	//
	// example:
	//
	// model_input_amount
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The array of performance data.
	Points []*DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) GetMeasurement() *string {
	return s.Measurement
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) GetPoints() []*DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints {
	return s.Points
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) SetDBNodeId(v string) *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys {
	s.DBNodeId = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) SetMeasurement(v string) *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Measurement = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) SetMetricName(v string) *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys {
	s.MetricName = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) SetPoints(v []*DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints) *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Points = v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeys) Validate() error {
	if s.Points != nil {
		for _, item := range s.Points {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints struct {
	// The UNIX timestamp that indicates when the metric was collected. Unit: milliseconds.
	//
	// example:
	//
	// 1724206183
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 42.38
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints) GetValue() *string {
	return s.Value
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints) SetTimestamp(v int64) *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints {
	s.Timestamp = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints) SetValue(v string) *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints {
	s.Value = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponseBodyPerformanceKeysPoints) Validate() error {
	return dara.Validate(s)
}
