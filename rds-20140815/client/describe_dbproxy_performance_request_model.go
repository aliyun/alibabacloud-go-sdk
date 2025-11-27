// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBProxyPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBProxyPerformanceRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *DescribeDBProxyPerformanceRequest
	GetDBProxyEngineType() *string
	SetDBProxyInstanceType(v string) *DescribeDBProxyPerformanceRequest
	GetDBProxyInstanceType() *string
	SetDimension(v string) *DescribeDBProxyPerformanceRequest
	GetDimension() *string
	SetEndTime(v string) *DescribeDBProxyPerformanceRequest
	GetEndTime() *string
	SetMetricsName(v string) *DescribeDBProxyPerformanceRequest
	GetMetricsName() *string
	SetOwnerId(v int64) *DescribeDBProxyPerformanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBProxyPerformanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBProxyPerformanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBProxyPerformanceRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeDBProxyPerformanceRequest
	GetStartTime() *string
}

type DescribeDBProxyPerformanceRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n3axxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The type of the database proxy instance. Valid values:
	//
	// 	- common: the general-purpose database proxy
	//
	// 	- exclusive: the dedicated database proxy
	//
	// example:
	//
	// DedicatedProxy
	DBProxyInstanceType *string `json:"DBProxyInstanceType,omitempty" xml:"DBProxyInstanceType,omitempty"`
	// Dimension.
	//
	// example:
	//
	// service
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-21T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance metrics that you want to query.
	//
	// If the instance runs MySQL, you can query only the **Maxscale_CpuUsage*	- performance metric, which indicates the CPU utilization of the instance.
	//
	// If the instance runs PostgreSQL, you can query the following performance metrics:
	//
	// 	- **Maxscale_TotalConns**: the number of connections per second
	//
	// 	- **Maxscale_CurrentConns**: the number of connections that are established
	//
	// 	- **Maxscale_DownFlows**: outbound traffic
	//
	// 	- **Maxscale_UpFlows**: inbound traffic
	//
	// 	- **Maxscale_QPS**: QPS
	//
	// 	- **Maxscale_MemUsage**: memory usage
	//
	// 	- **Maxscale_CpuUsage**: CPU utilization
	//
	// If you want to query more than one performance metric, separate the performance metrics with commas (,). You can specify up to six performance metrics in a single request.
	//
	// This parameter is required.
	//
	// example:
	//
	// Maxscale_CpuUsage
	MetricsName *string `json:"MetricsName,omitempty" xml:"MetricsName,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-19T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBProxyPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBProxyPerformanceRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *DescribeDBProxyPerformanceRequest) GetDBProxyInstanceType() *string {
	return s.DBProxyInstanceType
}

func (s *DescribeDBProxyPerformanceRequest) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeDBProxyPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBProxyPerformanceRequest) GetMetricsName() *string {
	return s.MetricsName
}

func (s *DescribeDBProxyPerformanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBProxyPerformanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBProxyPerformanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBProxyPerformanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBProxyPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBProxyPerformanceRequest) SetDBInstanceId(v string) *DescribeDBProxyPerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetDBProxyEngineType(v string) *DescribeDBProxyPerformanceRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetDBProxyInstanceType(v string) *DescribeDBProxyPerformanceRequest {
	s.DBProxyInstanceType = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetDimension(v string) *DescribeDBProxyPerformanceRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetEndTime(v string) *DescribeDBProxyPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetMetricsName(v string) *DescribeDBProxyPerformanceRequest {
	s.MetricsName = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetOwnerId(v int64) *DescribeDBProxyPerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetRegionId(v string) *DescribeDBProxyPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBProxyPerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBProxyPerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetStartTime(v string) *DescribeDBProxyPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
