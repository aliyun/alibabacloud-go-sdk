// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHadoopDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyHadoopDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceDescription(v string) *ModifyHadoopDataSourceRequest
	GetDataSourceDescription() *string
	SetDataSourceId(v string) *ModifyHadoopDataSourceRequest
	GetDataSourceId() *string
	SetDataSourceType(v string) *ModifyHadoopDataSourceRequest
	GetDataSourceType() *string
	SetEmrInstanceId(v string) *ModifyHadoopDataSourceRequest
	GetEmrInstanceId() *string
	SetHDFSConf(v string) *ModifyHadoopDataSourceRequest
	GetHDFSConf() *string
	SetHadoopCoreConf(v string) *ModifyHadoopDataSourceRequest
	GetHadoopCoreConf() *string
	SetHadoopCreateType(v string) *ModifyHadoopDataSourceRequest
	GetHadoopCreateType() *string
	SetHadoopHostsAddress(v string) *ModifyHadoopDataSourceRequest
	GetHadoopHostsAddress() *string
	SetHiveConf(v string) *ModifyHadoopDataSourceRequest
	GetHiveConf() *string
	SetMapReduceConf(v string) *ModifyHadoopDataSourceRequest
	GetMapReduceConf() *string
	SetRegionId(v string) *ModifyHadoopDataSourceRequest
	GetRegionId() *string
	SetYarnConf(v string) *ModifyHadoopDataSourceRequest
	GetYarnConf() *string
}

type ModifyHadoopDataSourceRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Service description, with a maximum length of 256 characters.
	//
	// example:
	//
	// pxf for hdfs data source
	DataSourceDescription *string `json:"DataSourceDescription,omitempty" xml:"DataSourceDescription,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- mysql
	//
	// - postgresql
	//
	// 	- hdfs
	//
	// - hive
	//
	// example:
	//
	// mysql
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// When HadoopCreateType is Emr, the value of this field is the EMR instance ID.
	//
	// example:
	//
	// c-1234567
	EmrInstanceId *string `json:"EmrInstanceId,omitempty" xml:"EmrInstanceId,omitempty"`
	// The content of the Hadoop hdfs-site.xml file. This parameter must be specified when DataSourceType is set to HDFS.
	//
	// example:
	//
	// xxxxxx
	HDFSConf *string `json:"HDFSConf,omitempty" xml:"HDFSConf,omitempty"`
	// The content of the Hadoop core-site.xml file.
	//
	// example:
	//
	// xxxxxx
	HadoopCoreConf *string `json:"HadoopCoreConf,omitempty" xml:"HadoopCoreConf,omitempty"`
	// The type of the external service. Valid values:
	//
	// 	- emr: E-MapReduce (EMR) Hadoop cluster.
	//
	// 	- selfCreate: self-managed Hadoop cluster.
	//
	// example:
	//
	// emr
	HadoopCreateType *string `json:"HadoopCreateType,omitempty" xml:"HadoopCreateType,omitempty"`
	// The IP address and hostname of the Hadoop cluster (data source) in the /etc/hosts file.
	//
	// example:
	//
	// 127.0.0.1 localhost
	HadoopHostsAddress *string `json:"HadoopHostsAddress,omitempty" xml:"HadoopHostsAddress,omitempty"`
	// The content of the Hadoop hive-site.xml file. This parameter must be specified when DataSourceType is set to Hive.
	//
	// example:
	//
	// xxxxxx
	HiveConf *string `json:"HiveConf,omitempty" xml:"HiveConf,omitempty"`
	// The content of the Hadoop mapred-site.xml file. This parameter must be specified when DataSourceType is set to HDFS.
	//
	// example:
	//
	// <?xml version="1.0" ?>
	//
	// <!-- Created at 2023-08-15 13:53:28.962 -->
	//
	// <configuration>
	//
	//     <property>
	//
	//         <name>mapreduce.map.speculative</name>
	//
	//         <value>true</value>
	//
	//     </property>
	//
	//     <property>
	//
	//         <name>mapreduce.jobhistory.keytab</name>
	//
	//         <value></value>
	//
	//     </property>
	//
	// </configuration>
	MapReduceConf *string `json:"MapReduceConf,omitempty" xml:"MapReduceConf,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The content of the Hadoop yarn-site.xml file. This parameter must be specified when DataSourceType is set to HDFS.
	//
	// example:
	//
	// xxxxxx
	YarnConf *string `json:"YarnConf,omitempty" xml:"YarnConf,omitempty"`
}

func (s ModifyHadoopDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHadoopDataSourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyHadoopDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyHadoopDataSourceRequest) GetDataSourceDescription() *string {
	return s.DataSourceDescription
}

func (s *ModifyHadoopDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ModifyHadoopDataSourceRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ModifyHadoopDataSourceRequest) GetEmrInstanceId() *string {
	return s.EmrInstanceId
}

func (s *ModifyHadoopDataSourceRequest) GetHDFSConf() *string {
	return s.HDFSConf
}

func (s *ModifyHadoopDataSourceRequest) GetHadoopCoreConf() *string {
	return s.HadoopCoreConf
}

func (s *ModifyHadoopDataSourceRequest) GetHadoopCreateType() *string {
	return s.HadoopCreateType
}

func (s *ModifyHadoopDataSourceRequest) GetHadoopHostsAddress() *string {
	return s.HadoopHostsAddress
}

func (s *ModifyHadoopDataSourceRequest) GetHiveConf() *string {
	return s.HiveConf
}

func (s *ModifyHadoopDataSourceRequest) GetMapReduceConf() *string {
	return s.MapReduceConf
}

func (s *ModifyHadoopDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHadoopDataSourceRequest) GetYarnConf() *string {
	return s.YarnConf
}

func (s *ModifyHadoopDataSourceRequest) SetDBInstanceId(v string) *ModifyHadoopDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetDataSourceDescription(v string) *ModifyHadoopDataSourceRequest {
	s.DataSourceDescription = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetDataSourceId(v string) *ModifyHadoopDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetDataSourceType(v string) *ModifyHadoopDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetEmrInstanceId(v string) *ModifyHadoopDataSourceRequest {
	s.EmrInstanceId = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetHDFSConf(v string) *ModifyHadoopDataSourceRequest {
	s.HDFSConf = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetHadoopCoreConf(v string) *ModifyHadoopDataSourceRequest {
	s.HadoopCoreConf = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetHadoopCreateType(v string) *ModifyHadoopDataSourceRequest {
	s.HadoopCreateType = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetHadoopHostsAddress(v string) *ModifyHadoopDataSourceRequest {
	s.HadoopHostsAddress = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetHiveConf(v string) *ModifyHadoopDataSourceRequest {
	s.HiveConf = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetMapReduceConf(v string) *ModifyHadoopDataSourceRequest {
	s.MapReduceConf = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetRegionId(v string) *ModifyHadoopDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) SetYarnConf(v string) *ModifyHadoopDataSourceRequest {
	s.YarnConf = &v
	return s
}

func (s *ModifyHadoopDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
