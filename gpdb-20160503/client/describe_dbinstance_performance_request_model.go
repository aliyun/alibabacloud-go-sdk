// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDBInstancePerformanceRequest
	GetEndTime() *string
	SetKey(v string) *DescribeDBInstancePerformanceRequest
	GetKey() *string
	SetResourceGroupId(v string) *DescribeDBInstancePerformanceRequest
	GetResourceGroupId() *string
	SetStartTime(v string) *DescribeDBInstancePerformanceRequest
	GetStartTime() *string
}

type DescribeDBInstancePerformanceRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-07-09T03:47Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance metric. Separate multiple values with commas (,). For more information, see [Performance parameters](https://help.aliyun.com/document_detail/86943.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// CpuUsage,MemoryUsage,Gpdb_SpaceUsage,Gpdb_IOPS,Gpdb_session
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is no longer used.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-07-08T03:47Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancePerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstancePerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancePerformanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancePerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstancePerformanceRequest) SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetEndTime(v string) *DescribeDBInstancePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetKey(v string) *DescribeDBInstancePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetResourceGroupId(v string) *DescribeDBInstancePerformanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetStartTime(v string) *DescribeDBInstancePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) Validate() error {
	return dara.Validate(s)
}
