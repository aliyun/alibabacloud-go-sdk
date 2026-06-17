// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTaskMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAIDBClusterTaskMetricsRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeAIDBClusterTaskMetricsRequest
	GetEndTime() *string
	SetMetricType(v string) *DescribeAIDBClusterTaskMetricsRequest
	GetMetricType() *string
	SetPageNumber(v int64) *DescribeAIDBClusterTaskMetricsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAIDBClusterTaskMetricsRequest
	GetPageSize() *int64
	SetRelativeDBClusterId(v string) *DescribeAIDBClusterTaskMetricsRequest
	GetRelativeDBClusterId() *string
	SetReverse(v bool) *DescribeAIDBClusterTaskMetricsRequest
	GetReverse() *bool
	SetStartTime(v string) *DescribeAIDBClusterTaskMetricsRequest
	GetStartTime() *string
}

type DescribeAIDBClusterTaskMetricsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pm-2zejpr***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-01-15T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metric type. Valid values:
	//
	// - `all`
	//
	// - `train`
	//
	// - `eval`
	//
	// > The default value is *all*.
	//
	// example:
	//
	// all
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The page number of the query result.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records to return on each page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// The default value is 100.
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the associated PolarDB instance.
	//
	// example:
	//
	// pc-2zejpr***
	RelativeDBClusterId *string `json:"RelativeDBClusterId,omitempty" xml:"RelativeDBClusterId,omitempty"`
	// Specifies whether to sort the results in reverse order. The default value is *false*.
	//
	// example:
	//
	// false
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The start time of the query. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-01-15T14:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAIDBClusterTaskMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskMetricsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterTaskMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAIDBClusterTaskMetricsRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeAIDBClusterTaskMetricsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAIDBClusterTaskMetricsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAIDBClusterTaskMetricsRequest) GetRelativeDBClusterId() *string {
	return s.RelativeDBClusterId
}

func (s *DescribeAIDBClusterTaskMetricsRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *DescribeAIDBClusterTaskMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAIDBClusterTaskMetricsRequest) SetDBClusterId(v string) *DescribeAIDBClusterTaskMetricsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsRequest) SetEndTime(v string) *DescribeAIDBClusterTaskMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsRequest) SetMetricType(v string) *DescribeAIDBClusterTaskMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsRequest) SetPageNumber(v int64) *DescribeAIDBClusterTaskMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsRequest) SetPageSize(v int64) *DescribeAIDBClusterTaskMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsRequest) SetRelativeDBClusterId(v string) *DescribeAIDBClusterTaskMetricsRequest {
	s.RelativeDBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsRequest) SetReverse(v bool) *DescribeAIDBClusterTaskMetricsRequest {
	s.Reverse = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsRequest) SetStartTime(v string) *DescribeAIDBClusterTaskMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsRequest) Validate() error {
	return dara.Validate(s)
}
