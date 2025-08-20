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
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	//
	// example:
	//
	// amv-bp1hx5n1o8f61****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-11T15:01Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The queried performance metrics.
	Performances []*DescribeDBClusterPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BD8C3096-8BC6-51DF-A4AB-BACD9DC10435
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-10T23:56Z
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
	return dara.Validate(s)
}

type DescribeDBClusterPerformanceResponseBodyPerformances struct {
	// The name of the performance metric.
	//
	// example:
	//
	// AnalyticDB_CPU_Usage_Percentage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The queried performance metric data.
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

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetSeries(v []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Series = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetUnit(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterPerformanceResponseBodyPerformancesSeries struct {
	// The name of the performance metric value. For more information about the performance metrics, see [Metric overview](https://help.aliyun.com/document_detail/2863211.html).
	//
	// example:
	//
	// AnalyticDB_Storage_CPU_Avg_Usage_Percentage
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tags that are added to the cluster.
	//
	// example:
	//
	// {instance_name: "am-***"}
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TranslateKey *string `json:"TranslateKey,omitempty" xml:"TranslateKey,omitempty"`
	// The values of the performance metric at different points in time.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) GetTags() *string {
	return s.Tags
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) GetTranslateKey() *string {
	return s.TranslateKey
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) GetValues() []*string {
	return s.Values
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetTags(v string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Tags = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetTranslateKey(v string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.TranslateKey = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetValues(v []*string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Values = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) Validate() error {
	return dara.Validate(s)
}
