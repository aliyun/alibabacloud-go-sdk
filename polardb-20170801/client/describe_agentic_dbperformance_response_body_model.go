// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAgenticDBPerformanceResponseBody
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeAgenticDBPerformanceResponseBody
	GetEndTime() *string
	SetPerformanceKeys(v *DescribeAgenticDBPerformanceResponseBodyPerformanceKeys) *DescribeAgenticDBPerformanceResponseBody
	GetPerformanceKeys() *DescribeAgenticDBPerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeAgenticDBPerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeAgenticDBPerformanceResponseBody
	GetStartTime() *string
}

type DescribeAgenticDBPerformanceResponseBody struct {
	// The AgenticDB cluster ID.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. Specify the time in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2020-09-23T01:01:00Z
	EndTime         *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PerformanceKeys *DescribeAgenticDBPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2F029645-FED9-4FE8-A6D3-488954******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. Specify the time in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2020-09-23T01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAgenticDBPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBPerformanceResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAgenticDBPerformanceResponseBody) GetPerformanceKeys() *DescribeAgenticDBPerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeAgenticDBPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAgenticDBPerformanceResponseBody) SetDBClusterId(v string) *DescribeAgenticDBPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBody) SetEndTime(v string) *DescribeAgenticDBPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBody) SetPerformanceKeys(v *DescribeAgenticDBPerformanceResponseBodyPerformanceKeys) *DescribeAgenticDBPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBody) SetRequestId(v string) *DescribeAgenticDBPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBody) SetStartTime(v string) *DescribeAgenticDBPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBody) Validate() error {
	if s.PerformanceKeys != nil {
		if err := s.PerformanceKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAgenticDBPerformanceResponseBodyPerformanceKeys struct {
	PerformanceKeys []*DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys `json:"performanceKeys,omitempty" xml:"performanceKeys,omitempty" type:"Repeated"`
}

func (s DescribeAgenticDBPerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeys) GetPerformanceKeys() []*DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeys) SetPerformanceKeys(v []*DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) *DescribeAgenticDBPerformanceResponseBodyPerformanceKeys {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeys) Validate() error {
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

type DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys struct {
	BranchId    *string                                                                       `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	EndpointId  *string                                                                       `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	Measurement *string                                                                       `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	MetricName  *string                                                                       `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Points      *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
}

func (s DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) GetMeasurement() *string {
	return s.Measurement
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) GetPoints() *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPoints {
	return s.Points
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) SetBranchId(v string) *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys {
	s.BranchId = &v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) SetEndpointId(v string) *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys {
	s.EndpointId = &v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) SetMeasurement(v string) *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys {
	s.Measurement = &v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) SetMetricName(v string) *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys {
	s.MetricName = &v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) SetPoints(v *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPoints) *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys {
	s.Points = v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeys) Validate() error {
	if s.Points != nil {
		if err := s.Points.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPoints struct {
	Points []*DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints `json:"points,omitempty" xml:"points,omitempty" type:"Repeated"`
}

func (s DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPoints) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPoints) GetPoints() []*DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints {
	return s.Points
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPoints) SetPoints(v []*DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints) *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPoints {
	s.Points = v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPoints) Validate() error {
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

type DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints struct {
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints) GetValue() *string {
	return s.Value
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints) SetTimestamp(v int64) *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints {
	s.Timestamp = &v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints) SetValue(v string) *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints {
	s.Value = &v
	return s
}

func (s *DescribeAgenticDBPerformanceResponseBodyPerformanceKeysPerformanceKeysPointsPoints) Validate() error {
	return dara.Validate(s)
}
