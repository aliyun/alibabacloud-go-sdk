// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResourceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeJobResourceUsageRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeJobResourceUsageRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeJobResourceUsageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeJobResourceUsageRequest
	GetPageSize() *int32
	SetSparkAppName(v string) *DescribeJobResourceUsageRequest
	GetSparkAppName() *string
	SetStartTime(v string) *DescribeJobResourceUsageRequest
	GetStartTime() *string
}

type DescribeJobResourceUsageRequest struct {
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6g8w25jacm7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The end time must be later than the start time. Format: <i>yyyy-MM-ddTHH:mm:ssZ</i> (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-03-17T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. The value must be a positive integer. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - 30
	//
	// - 50
	//
	// - 100
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SparkAppName *string `json:"SparkAppName,omitempty" xml:"SparkAppName,omitempty"`
	// The start time of the query. Format: <i>yyyy-MM-ddTHH:mm:ssZ</i> (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-04T03:45:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeJobResourceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeJobResourceUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeJobResourceUsageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeJobResourceUsageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeJobResourceUsageRequest) GetSparkAppName() *string {
	return s.SparkAppName
}

func (s *DescribeJobResourceUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeJobResourceUsageRequest) SetDBClusterId(v string) *DescribeJobResourceUsageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeJobResourceUsageRequest) SetEndTime(v string) *DescribeJobResourceUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeJobResourceUsageRequest) SetPageNumber(v int32) *DescribeJobResourceUsageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeJobResourceUsageRequest) SetPageSize(v int32) *DescribeJobResourceUsageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeJobResourceUsageRequest) SetSparkAppName(v string) *DescribeJobResourceUsageRequest {
	s.SparkAppName = &v
	return s
}

func (s *DescribeJobResourceUsageRequest) SetStartTime(v string) *DescribeJobResourceUsageRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeJobResourceUsageRequest) Validate() error {
	return dara.Validate(s)
}
