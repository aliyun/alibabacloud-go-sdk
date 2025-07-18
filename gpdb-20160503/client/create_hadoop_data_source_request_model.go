// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHadoopDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateHadoopDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceDescription(v string) *CreateHadoopDataSourceRequest
	GetDataSourceDescription() *string
	SetDataSourceName(v string) *CreateHadoopDataSourceRequest
	GetDataSourceName() *string
	SetDataSourceType(v string) *CreateHadoopDataSourceRequest
	GetDataSourceType() *string
	SetEmrInstanceId(v string) *CreateHadoopDataSourceRequest
	GetEmrInstanceId() *string
	SetHDFSConf(v string) *CreateHadoopDataSourceRequest
	GetHDFSConf() *string
	SetHadoopCoreConf(v string) *CreateHadoopDataSourceRequest
	GetHadoopCoreConf() *string
	SetHadoopCreateType(v string) *CreateHadoopDataSourceRequest
	GetHadoopCreateType() *string
	SetHadoopHostsAddress(v string) *CreateHadoopDataSourceRequest
	GetHadoopHostsAddress() *string
	SetHiveConf(v string) *CreateHadoopDataSourceRequest
	GetHiveConf() *string
	SetMapReduceConf(v string) *CreateHadoopDataSourceRequest
	GetMapReduceConf() *string
	SetRegionId(v string) *CreateHadoopDataSourceRequest
	GetRegionId() *string
	SetYarnConf(v string) *CreateHadoopDataSourceRequest
	GetYarnConf() *string
}

type CreateHadoopDataSourceRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Service description.
	//
	// example:
	//
	// pxf for hdfs data source
	DataSourceDescription *string `json:"DataSourceDescription,omitempty" xml:"DataSourceDescription,omitempty"`
	// Service name.
	//
	// example:
	//
	// hdfs_pxf
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// Type of Hadoop external table to be enabled, with values:
	//
	// - HDFS
	//
	// - Hive
	//
	// example:
	//
	// HDFS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// When HadoopCreateType=Emr, this field should contain the EMR instance ID.
	//
	// example:
	//
	// c-1234567
	EmrInstanceId *string `json:"EmrInstanceId,omitempty" xml:"EmrInstanceId,omitempty"`
	// Content string of the Hadoop hdfs-site.xml file. This field is required when enabling an HDFS external table.
	//
	// example:
	//
	// <?xml version="1.0" ?>
	//
	// <!-- Created at 2023-08-15 13:52:43.945 -->
	//
	// <configuration>
	//
	//     <property>
	//
	//         <name>dfs.datanode.cache.revocation.timeout.ms</name>
	//
	//         <value>900000</value>
	//
	//     </property>
	//
	//     <property>
	//
	//         <name>dfs.namenode.resource.check.interval</name>
	//
	//         <value>5000</value>
	//
	//     </property>
	//
	// </configuration>
	HDFSConf *string `json:"HDFSConf,omitempty" xml:"HDFSConf,omitempty"`
	// Content string of the Hadoop core-site.xml file.
	//
	// example:
	//
	// <?xml version="1.0" ?>
	//
	// <!-- Created at 2023-08-15 13:52:39.527 -->
	//
	// <configuration>
	//
	//     <property>
	//
	//         <name>hadoop.http.authentication.kerberos.keytab</name>
	//
	//         <value>/etc/emr/hadoop-conf/http.keytab</value>
	//
	//     </property>
	//
	//     <property>
	//
	//         <name>fs.oss.idle.timeout.millisecond</name>
	//
	//         <value>30000</value>
	//
	//     </property>
	//
	//     <property>
	//
	//         <name>fs.oss.download.thread.concurrency</name>
	//
	//         <value>32</value>
	//
	//     </property>
	//
	// </configuration>
	HadoopCoreConf *string `json:"HadoopCoreConf,omitempty" xml:"HadoopCoreConf,omitempty"`
	// External service type:
	//
	// - emr
	//
	// - hadoop: Self-built Hadoop
	//
	// example:
	//
	// emr
	HadoopCreateType *string `json:"HadoopCreateType,omitempty" xml:"HadoopCreateType,omitempty"`
	// Address and hostname of the Hadoop cluster\\"s source node in the /etc/hosts file.
	//
	// example:
	//
	// 192.168.220.128 master-1-1.c-xxx.cn-shanghai.emr.aliyuncs.com
	//
	// 192.168.220.129 core-1-1.c-xxx.cn-shanghai.emr.aliyuncs.com
	//
	// 192.168.220.130 core-1-2.c-xxx.cn-shanghai.emr.aliyuncs.com
	HadoopHostsAddress *string `json:"HadoopHostsAddress,omitempty" xml:"HadoopHostsAddress,omitempty"`
	// Content string of the Hadoop hive-site.xml file. This field is required when enabling a HIVE external table.
	//
	// example:
	//
	// <?xml version="1.0" ?>
	//
	// <!-- Created at 2023-08-15 13:52:50.646 -->
	//
	// <configuration>
	//
	//     <property>
	//
	//         <name>hive.exec.reducers.bytes.per.reducer</name>
	//
	//         <value>256000000</value>
	//
	//     </property>
	//
	//     <property>
	//
	//         <name>hive.stats.column.autogather</name>
	//
	//         <value>false</value>
	//
	//     </property>
	//
	// </configuration>
	HiveConf *string `json:"HiveConf,omitempty" xml:"HiveConf,omitempty"`
	// Content string of the Hadoop mapred-site.xml file. This field is required when enabling an HDFS external table.
	//
	// example:
	//
	// xxxxxx
	MapReduceConf *string `json:"MapReduceConf,omitempty" xml:"MapReduceConf,omitempty"`
	// Region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) interface to view available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Content string of the Hadoop yarn-site.xml file. This field is required when enabling an HDFS external table.
	//
	// example:
	//
	// <?xml version="1.0" ?>
	//
	// <!-- Created at 2023-08-15 13:53:29.021 -->
	//
	// <configuration>
	//
	//     <property>
	//
	//         <name>yarn.nodemanager.linux-container-executor.nonsecure-mode.local-user</name>
	//
	//         <value>hadoop</value>
	//
	//     </property>
	//
	//     <property>
	//
	//         <name>yarn.scheduler.fair.dynamic.max.assign</name>
	//
	//         <value>true</value>
	//
	//     </property>
	//
	// </configuration>
	YarnConf *string `json:"YarnConf,omitempty" xml:"YarnConf,omitempty"`
}

func (s CreateHadoopDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHadoopDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateHadoopDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateHadoopDataSourceRequest) GetDataSourceDescription() *string {
	return s.DataSourceDescription
}

func (s *CreateHadoopDataSourceRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *CreateHadoopDataSourceRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CreateHadoopDataSourceRequest) GetEmrInstanceId() *string {
	return s.EmrInstanceId
}

func (s *CreateHadoopDataSourceRequest) GetHDFSConf() *string {
	return s.HDFSConf
}

func (s *CreateHadoopDataSourceRequest) GetHadoopCoreConf() *string {
	return s.HadoopCoreConf
}

func (s *CreateHadoopDataSourceRequest) GetHadoopCreateType() *string {
	return s.HadoopCreateType
}

func (s *CreateHadoopDataSourceRequest) GetHadoopHostsAddress() *string {
	return s.HadoopHostsAddress
}

func (s *CreateHadoopDataSourceRequest) GetHiveConf() *string {
	return s.HiveConf
}

func (s *CreateHadoopDataSourceRequest) GetMapReduceConf() *string {
	return s.MapReduceConf
}

func (s *CreateHadoopDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHadoopDataSourceRequest) GetYarnConf() *string {
	return s.YarnConf
}

func (s *CreateHadoopDataSourceRequest) SetDBInstanceId(v string) *CreateHadoopDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetDataSourceDescription(v string) *CreateHadoopDataSourceRequest {
	s.DataSourceDescription = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetDataSourceName(v string) *CreateHadoopDataSourceRequest {
	s.DataSourceName = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetDataSourceType(v string) *CreateHadoopDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetEmrInstanceId(v string) *CreateHadoopDataSourceRequest {
	s.EmrInstanceId = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetHDFSConf(v string) *CreateHadoopDataSourceRequest {
	s.HDFSConf = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetHadoopCoreConf(v string) *CreateHadoopDataSourceRequest {
	s.HadoopCoreConf = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetHadoopCreateType(v string) *CreateHadoopDataSourceRequest {
	s.HadoopCreateType = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetHadoopHostsAddress(v string) *CreateHadoopDataSourceRequest {
	s.HadoopHostsAddress = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetHiveConf(v string) *CreateHadoopDataSourceRequest {
	s.HiveConf = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetMapReduceConf(v string) *CreateHadoopDataSourceRequest {
	s.MapReduceConf = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetRegionId(v string) *CreateHadoopDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) SetYarnConf(v string) *CreateHadoopDataSourceRequest {
	s.YarnConf = &v
	return s
}

func (s *CreateHadoopDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
