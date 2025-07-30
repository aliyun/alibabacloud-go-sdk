// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConfigChangeLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *DescribeDBClusterConfigChangeLogsRequest
	GetConfigKey() *string
	SetDBClusterId(v string) *DescribeDBClusterConfigChangeLogsRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *DescribeDBClusterConfigChangeLogsRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDBClusterConfigChangeLogsRequest
	GetEndTime() *string
	SetRegionId(v string) *DescribeDBClusterConfigChangeLogsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDBClusterConfigChangeLogsRequest
	GetStartTime() *string
}

type DescribeDBClusterConfigChangeLogsRequest struct {
	// The configuration file that you want to modify. For a compute cluster, set the value to be.conf. For a frontend (FE) cluster, set the value to fe.conf.
	//
	// example:
	//
	// be.conf
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8yvv09-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-jia3ma3b003
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Format: yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-08T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Format: yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-25T09:48:23Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterConfigChangeLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigChangeLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigChangeLogsRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DescribeDBClusterConfigChangeLogsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterConfigChangeLogsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBClusterConfigChangeLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBClusterConfigChangeLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterConfigChangeLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetConfigKey(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.ConfigKey = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetDBClusterId(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetDBInstanceId(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetEndTime(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetRegionId(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsRequest) SetStartTime(v string) *DescribeDBClusterConfigChangeLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsRequest) Validate() error {
	return dara.Validate(s)
}
