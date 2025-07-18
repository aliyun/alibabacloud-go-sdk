// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHadoopDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeHadoopDataSourceResponseBody
	GetCreateTime() *string
	SetDataSourceDescription(v string) *DescribeHadoopDataSourceResponseBody
	GetDataSourceDescription() *string
	SetDataSourceDir(v string) *DescribeHadoopDataSourceResponseBody
	GetDataSourceDir() *string
	SetDataSourceId(v string) *DescribeHadoopDataSourceResponseBody
	GetDataSourceId() *string
	SetDataSourceName(v string) *DescribeHadoopDataSourceResponseBody
	GetDataSourceName() *string
	SetDataSourceStatus(v string) *DescribeHadoopDataSourceResponseBody
	GetDataSourceStatus() *string
	SetDataSourceType(v string) *DescribeHadoopDataSourceResponseBody
	GetDataSourceType() *string
	SetEmrInstanceId(v string) *DescribeHadoopDataSourceResponseBody
	GetEmrInstanceId() *string
	SetExternalDataServiceId(v string) *DescribeHadoopDataSourceResponseBody
	GetExternalDataServiceId() *string
	SetHDFSConf(v string) *DescribeHadoopDataSourceResponseBody
	GetHDFSConf() *string
	SetHadoopCoreConf(v string) *DescribeHadoopDataSourceResponseBody
	GetHadoopCoreConf() *string
	SetHadoopCreateType(v string) *DescribeHadoopDataSourceResponseBody
	GetHadoopCreateType() *string
	SetHadoopHostsAddress(v string) *DescribeHadoopDataSourceResponseBody
	GetHadoopHostsAddress() *string
	SetHiveConf(v string) *DescribeHadoopDataSourceResponseBody
	GetHiveConf() *string
	SetMapReduceConf(v string) *DescribeHadoopDataSourceResponseBody
	GetMapReduceConf() *string
	SetModifyTime(v string) *DescribeHadoopDataSourceResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *DescribeHadoopDataSourceResponseBody
	GetRequestId() *string
	SetStatusMessage(v string) *DescribeHadoopDataSourceResponseBody
	GetStatusMessage() *string
	SetYarnConf(v string) *DescribeHadoopDataSourceResponseBody
	GetYarnConf() *string
}

type DescribeHadoopDataSourceResponseBody struct {
	// The time when the service was created.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the service. The description can be up to 256 characters in length.
	//
	// example:
	//
	// pxf for hdfs data source
	DataSourceDescription *string `json:"DataSourceDescription,omitempty" xml:"DataSourceDescription,omitempty"`
	// The service directory in which Hadoop-related configuration files are stored.
	//
	// example:
	//
	// HadoopDir
	DataSourceDir *string `json:"DataSourceDir,omitempty" xml:"DataSourceDir,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// hdfs_pxf
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- Init
	//
	// 	- Running
	//
	// 	- Exception
	//
	// example:
	//
	// Running
	DataSourceStatus *string `json:"DataSourceStatus,omitempty" xml:"DataSourceStatus,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// hive
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The E-MapReduce (EMR) Hadoop cluster ID.
	//
	// example:
	//
	// c-1234567
	EmrInstanceId *string `json:"EmrInstanceId,omitempty" xml:"EmrInstanceId,omitempty"`
	// The ID of the external data service.
	//
	// example:
	//
	// 2988
	ExternalDataServiceId *string `json:"ExternalDataServiceId,omitempty" xml:"ExternalDataServiceId,omitempty"`
	// The content of the Hadoop hdfs-site.xml file.
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
	// 	- emr
	//
	// 	- selfCreate
	//
	// example:
	//
	// HDFS
	HadoopCreateType *string `json:"HadoopCreateType,omitempty" xml:"HadoopCreateType,omitempty"`
	// The IP address and hostname of the Hadoop cluster (data source) in the /etc/hosts file.
	//
	// example:
	//
	// 127.0.0.1 localhost
	HadoopHostsAddress *string `json:"HadoopHostsAddress,omitempty" xml:"HadoopHostsAddress,omitempty"`
	// The content of the Hadoop hive-site.xml file.
	//
	// example:
	//
	// xxxxxx
	HiveConf *string `json:"HiveConf,omitempty" xml:"HiveConf,omitempty"`
	// The content of the Hadoop mapred-site.xml file.
	//
	// example:
	//
	// xxxxxx
	MapReduceConf *string `json:"MapReduceConf,omitempty" xml:"MapReduceConf,omitempty"`
	// The time when the data source was last modified.
	//
	// example:
	//
	// 2024-08-23T02:11:47Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the service status. For example, if the service is in the exception state, the cause of the exception is displayed. If the service is in the running state, this parameter is left empty.
	//
	// example:
	//
	// ""
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The content of the Hadoop yarn-site.xml file.
	//
	// example:
	//
	// xxxxxx
	YarnConf *string `json:"YarnConf,omitempty" xml:"YarnConf,omitempty"`
}

func (s DescribeHadoopDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHadoopDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHadoopDataSourceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeHadoopDataSourceResponseBody) GetDataSourceDescription() *string {
	return s.DataSourceDescription
}

func (s *DescribeHadoopDataSourceResponseBody) GetDataSourceDir() *string {
	return s.DataSourceDir
}

func (s *DescribeHadoopDataSourceResponseBody) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribeHadoopDataSourceResponseBody) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *DescribeHadoopDataSourceResponseBody) GetDataSourceStatus() *string {
	return s.DataSourceStatus
}

func (s *DescribeHadoopDataSourceResponseBody) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *DescribeHadoopDataSourceResponseBody) GetEmrInstanceId() *string {
	return s.EmrInstanceId
}

func (s *DescribeHadoopDataSourceResponseBody) GetExternalDataServiceId() *string {
	return s.ExternalDataServiceId
}

func (s *DescribeHadoopDataSourceResponseBody) GetHDFSConf() *string {
	return s.HDFSConf
}

func (s *DescribeHadoopDataSourceResponseBody) GetHadoopCoreConf() *string {
	return s.HadoopCoreConf
}

func (s *DescribeHadoopDataSourceResponseBody) GetHadoopCreateType() *string {
	return s.HadoopCreateType
}

func (s *DescribeHadoopDataSourceResponseBody) GetHadoopHostsAddress() *string {
	return s.HadoopHostsAddress
}

func (s *DescribeHadoopDataSourceResponseBody) GetHiveConf() *string {
	return s.HiveConf
}

func (s *DescribeHadoopDataSourceResponseBody) GetMapReduceConf() *string {
	return s.MapReduceConf
}

func (s *DescribeHadoopDataSourceResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeHadoopDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHadoopDataSourceResponseBody) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *DescribeHadoopDataSourceResponseBody) GetYarnConf() *string {
	return s.YarnConf
}

func (s *DescribeHadoopDataSourceResponseBody) SetCreateTime(v string) *DescribeHadoopDataSourceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetDataSourceDescription(v string) *DescribeHadoopDataSourceResponseBody {
	s.DataSourceDescription = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetDataSourceDir(v string) *DescribeHadoopDataSourceResponseBody {
	s.DataSourceDir = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetDataSourceId(v string) *DescribeHadoopDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetDataSourceName(v string) *DescribeHadoopDataSourceResponseBody {
	s.DataSourceName = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetDataSourceStatus(v string) *DescribeHadoopDataSourceResponseBody {
	s.DataSourceStatus = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetDataSourceType(v string) *DescribeHadoopDataSourceResponseBody {
	s.DataSourceType = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetEmrInstanceId(v string) *DescribeHadoopDataSourceResponseBody {
	s.EmrInstanceId = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetExternalDataServiceId(v string) *DescribeHadoopDataSourceResponseBody {
	s.ExternalDataServiceId = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetHDFSConf(v string) *DescribeHadoopDataSourceResponseBody {
	s.HDFSConf = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetHadoopCoreConf(v string) *DescribeHadoopDataSourceResponseBody {
	s.HadoopCoreConf = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetHadoopCreateType(v string) *DescribeHadoopDataSourceResponseBody {
	s.HadoopCreateType = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetHadoopHostsAddress(v string) *DescribeHadoopDataSourceResponseBody {
	s.HadoopHostsAddress = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetHiveConf(v string) *DescribeHadoopDataSourceResponseBody {
	s.HiveConf = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetMapReduceConf(v string) *DescribeHadoopDataSourceResponseBody {
	s.MapReduceConf = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetModifyTime(v string) *DescribeHadoopDataSourceResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetRequestId(v string) *DescribeHadoopDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetStatusMessage(v string) *DescribeHadoopDataSourceResponseBody {
	s.StatusMessage = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) SetYarnConf(v string) *DescribeHadoopDataSourceResponseBody {
	s.YarnConf = &v
	return s
}

func (s *DescribeHadoopDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
