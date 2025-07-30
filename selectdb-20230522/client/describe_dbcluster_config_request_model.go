// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *DescribeDBClusterConfigRequest
	GetConfigKey() *string
	SetDBClusterId(v string) *DescribeDBClusterConfigRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *DescribeDBClusterConfigRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeDBClusterConfigRequest
	GetRegionId() *string
}

type DescribeDBClusterConfigRequest struct {
	// The configuration file to be modified.
	//
	// 	- For a compute cluster, set the value to be.conf.
	//
	// 	- For a frontend (FE) cluster, set the value to fe.conf.
	//
	// This parameter is required.
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
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DescribeDBClusterConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBClusterConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterConfigRequest) SetConfigKey(v string) *DescribeDBClusterConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetDBClusterId(v string) *DescribeDBClusterConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetDBInstanceId(v string) *DescribeDBClusterConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetRegionId(v string) *DescribeDBClusterConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) Validate() error {
	return dara.Validate(s)
}
