// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHadoopConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigName(v string) *DescribeHadoopConfigsRequest
	GetConfigName() *string
	SetDBInstanceId(v string) *DescribeHadoopConfigsRequest
	GetDBInstanceId() *string
	SetEmrInstanceId(v string) *DescribeHadoopConfigsRequest
	GetEmrInstanceId() *string
	SetRegionId(v string) *DescribeHadoopConfigsRequest
	GetRegionId() *string
}

type DescribeHadoopConfigsRequest struct {
	// The name of the configuration file. Valid values:
	//
	// 	- hdfs-site
	//
	// 	- core-site
	//
	// 	- yarn-site
	//
	// 	- mapred-site
	//
	// 	- hive-site
	//
	// This parameter is required.
	//
	// example:
	//
	// hdfs-site
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The E-MapReduce (EMR) Hadoop cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-xxx
	EmrInstanceId *string `json:"EmrInstanceId,omitempty" xml:"EmrInstanceId,omitempty"`
	// The region ID of the instance.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeHadoopConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHadoopConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHadoopConfigsRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *DescribeHadoopConfigsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeHadoopConfigsRequest) GetEmrInstanceId() *string {
	return s.EmrInstanceId
}

func (s *DescribeHadoopConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHadoopConfigsRequest) SetConfigName(v string) *DescribeHadoopConfigsRequest {
	s.ConfigName = &v
	return s
}

func (s *DescribeHadoopConfigsRequest) SetDBInstanceId(v string) *DescribeHadoopConfigsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeHadoopConfigsRequest) SetEmrInstanceId(v string) *DescribeHadoopConfigsRequest {
	s.EmrInstanceId = &v
	return s
}

func (s *DescribeHadoopConfigsRequest) SetRegionId(v string) *DescribeHadoopConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHadoopConfigsRequest) Validate() error {
	return dara.Validate(s)
}
