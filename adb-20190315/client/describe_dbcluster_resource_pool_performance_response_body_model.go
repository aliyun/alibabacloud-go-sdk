// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterResourcePoolPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody
	GetEndTime() *string
	SetPerformances(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) *DescribeDBClusterResourcePoolPerformanceResponseBody
	GetPerformances() []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances
	SetRequestId(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody
	GetStartTime() *string
}

type DescribeDBClusterResourcePoolPerformanceResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// am-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range for monitoring the resource group. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-10T07:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The queried monitoring information about the metrics.
	Performances []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C7EDB8E4-9769-4233-88C7-DCA4C9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range for monitoring the resource group. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-10T07:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) GetPerformances() []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	return s.Performances
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetPerformances(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) Validate() error {
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

type DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances struct {
	// The metric of the resource group.
	//
	// example:
	//
	// AnalyticDB_RP_CPU
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The queried monitoring information about the resource groups.
	ResourcePoolPerformances []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances `json:"ResourcePoolPerformances,omitempty" xml:"ResourcePoolPerformances,omitempty" type:"Repeated"`
	// The unit of the metric value.
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) GetResourcePoolPerformances() []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances {
	return s.ResourcePoolPerformances
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) SetKey(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) SetResourcePoolPerformances(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	s.ResourcePoolPerformances = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) SetUnit(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) Validate() error {
	if s.ResourcePoolPerformances != nil {
		for _, item := range s.ResourcePoolPerformances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances struct {
	// The name of the resource group.
	//
	// example:
	//
	// test_pool
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	// The sequential monitoring information about the resource groups.
	ResourcePoolSeries []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries `json:"ResourcePoolSeries,omitempty" xml:"ResourcePoolSeries,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) GetResourcePoolName() *string {
	return s.ResourcePoolName
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) GetResourcePoolSeries() []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries {
	return s.ResourcePoolSeries
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) SetResourcePoolName(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) SetResourcePoolSeries(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances {
	s.ResourcePoolSeries = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) Validate() error {
	if s.ResourcePoolSeries != nil {
		for _, item := range s.ResourcePoolSeries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries struct {
	// The name of the metric.
	//
	// example:
	//
	// cpu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the metric.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) GetName() *string {
	return s.Name
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) GetValues() []*string {
	return s.Values
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) SetName(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) SetValues(v []*string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries {
	s.Values = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) Validate() error {
	return dara.Validate(s)
}
