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
	SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody
	GetEndTime() *string
	SetPerformances(v []*DescribeDBClusterPerformanceResponseBodyPerformances) *DescribeDBClusterPerformanceResponseBody
	GetPerformances() []*DescribeDBClusterPerformanceResponseBodyPerformances
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
	// cc-bp125e3uu94wo****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2021-11-27T16:38Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The values of the queried performance metrics of the cluster.
	Performances []*DescribeDBClusterPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// FE242962-6DA3-5FC8-9691-37B62A3210F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range to query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-27T16:37Z
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

func (s *DescribeDBClusterPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBClusterPerformanceResponseBody) GetPerformances() []*DescribeDBClusterPerformanceResponseBodyPerformances {
	return s.Performances
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

func (s *DescribeDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetPerformances(v []*DescribeDBClusterPerformanceResponseBodyPerformances) *DescribeDBClusterPerformanceResponseBody {
	s.Performances = v
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
	if s.Performances != nil {
		for _, item := range s.Performances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterPerformanceResponseBodyPerformances struct {
	// The name of the performance metric.
	//
	// example:
	//
	// MEM_USAGE
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the performance metric value.
	//
	// example:
	//
	// mem_usage
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The queried performance pamaters.
	Series []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The unit of the performance metric.
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) GetName() *string {
	return s.Name
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) GetSeries() []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	return s.Series
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetKey(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetSeries(v []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Series = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetUnit(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) Validate() error {
	if s.Series != nil {
		for _, item := range s.Series {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterPerformanceResponseBodyPerformancesSeries struct {
	// The name of the list of performance metric values.
	//
	// example:
	//
	// cc-bp125e3uu94wo1s0k16****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the performance parameter. Each value of the performance parameter is collected at a point in time.
	Values []*DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) GetName() *string {
	return s.Name
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) GetValues() []*DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues {
	return s.Values
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetValues(v []*DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Values = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) Validate() error {
	if s.Values != nil {
		for _, item := range s.Values {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues struct {
	// The values of a metric.
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues) GetPoint() []*string {
	return s.Point
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues) SetPoint(v []*string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues {
	s.Point = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues) Validate() error {
	return dara.Validate(s)
}
