// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeDBClusterPerformanceResponseBody
	GetAccessDeniedDetail() *string
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
	// The information about the request denial. This parameter is returned only if Resource Access Management (RAM) permission verification failed.
	//
	// example:
	//
	// {
	//
	//   "AuthAction": "xxx",
	//
	//   "AuthPrincipalDisplayName": "sampleName",
	//
	//   "AuthPrincipalOwnerId": "111111111111111111",
	//
	//   "AuthPrincipalType": "SubUser",
	//
	//   "AuthResource": "xxx",
	//
	//   "NoPermissionType": "xxx",
	//
	//   "PolicyType": "xxx"
	//
	// }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// am-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-03T15:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The queried performance metrics.
	Performances []*DescribeDBClusterPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25B56BC7-4978-40B3-9E48-4B7067******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-03T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
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

func (s *DescribeDBClusterPerformanceResponseBody) SetAccessDeniedDetail(v string) *DescribeDBClusterPerformanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// AnalyticDB_CPU
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
	// The name of the performance metric value. For more information about the performance metrics, see [Metric overview](https://help.aliyun.com/document_detail/2863211.html).
	//
	// example:
	//
	// worker_avg_cpu_used
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tags that are added to the cluster.
	//
	// example:
	//
	// instance_name: "amv-8vbf80pjcz*****"
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The key that is used to obtain the name of the performance metric value.
	//
	// example:
	//
	// elastic_executor_avg_cpu_use
	TranslateKey *string `json:"TranslateKey,omitempty" xml:"TranslateKey,omitempty"`
	// The values of the queried performance metric.
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
