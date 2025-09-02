// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateDataSourceRequest
	GetContent() *string
	SetDataSourceId(v int64) *UpdateDataSourceRequest
	GetDataSourceId() *int64
	SetDescription(v string) *UpdateDataSourceRequest
	GetDescription() *string
	SetEnvType(v int32) *UpdateDataSourceRequest
	GetEnvType() *int32
	SetStatus(v string) *UpdateDataSourceRequest
	GetStatus() *string
}

type UpdateDataSourceRequest struct {
	// The details about the data source. You are not allowed to change the type of the data source. For example, you are not allowed to change the data source type from MaxCompute to MySQL. Examples of details of some common data sources:
	//
	// 	- odps
	//
	//         {
	//
	//           "accessId": "*****",
	//
	//           "accessKey": "*****",
	//
	//           "authType": 2,
	//
	//           "endpoint": "http://service.odps.aliyun.com/api",
	//
	//           "project": "xsaxsax",
	//
	//           "tag": "public"
	//
	//         }
	//
	// 	- mysql
	//
	//         {
	//
	//           "database": "xsaxsa",
	//
	//           "instanceName": "rm-xsaxsa",
	//
	//           "password": "xsaxsa",
	//
	//           "rdsOwnerId": "xasxsa",
	//
	//           "regionId": "cn-shanghai",
	//
	//           "tag": "rds",
	//
	//           "username": "xsaxsa"
	//
	//         }
	//
	// 	- rds
	//
	//         {
	//
	//           "configType": 1,
	//
	//           "tag": "rds",
	//
	//           "database": "xsaxsa",
	//
	//           "username": "xsaxsa",
	//
	//           "password": "xssaxsa$32050",
	//
	//           "instanceName": "rm-xsaxs",
	//
	//           "rdsOwnerId": "11111111"
	//
	//         }
	//
	// 	- oss
	//
	//         {
	//
	//           "accessId": "*****",
	//
	//           "accessKey": "*****",
	//
	//           "bucket": "xsa-xs-xs",
	//
	//           "endpoint": "http://oss-cn-shanghai.aliyuncs.com",
	//
	//           "tag": "public"
	//
	//         }
	//
	// 	- sqlserver
	//
	//         {
	//
	//           "jdbcUrl": "jdbc:sqlserver://xsaxsa-xsaxsa.database.xxx.cn:123;DatabaseName=xsxs-xsxs",
	//
	//           "password": "sdasda$fs",
	//
	//           "tag": "public",
	//
	//           "username": "sxaxacdacdd"
	//
	//         }
	//
	// 	- polardb
	//
	//         {
	//
	//           "clusterId": "pc-sdadsadsa",
	//
	//           "database": "dsadsadsa",
	//
	//           "ownerId": "121212122",
	//
	//           "password": "sdasdafssa",
	//
	//           "region": "cn-shanghai",
	//
	//           "tag": "polardb",
	//
	//           "username": "asdadsads"
	//
	//         }
	//
	// 	- oracle
	//
	//         {
	//
	//           "jdbcUrl": "jdbc:oracle:saaa:@xxxxx:1521:PROD",
	//
	//           "password": "sxasaxsa",
	//
	//           "tag": "public",
	//
	//           "username": "sasfadfa"
	//
	//         }
	//
	// 	- mongodb
	//
	//         {
	//
	//           "address": "[\\"xsaxxsa.mongodb.rds.aliyuncs.com:3717\\"]",
	//
	//           "database": "admin",
	//
	//           "password": "sadsda@",
	//
	//           "tag": "public",
	//
	//           "username": "dsadsadas"
	//
	//         }
	//
	// 	- emr
	//
	//         {
	//
	//           "accessId": "*****",
	//
	//           "emrClusterId": "C-dsads",
	//
	//           "emrResourceQueueName": "default",
	//
	//           "emrEndpoint": "emr.aliyuncs.com",
	//
	//           "accessKey": "*****",
	//
	//           "emrUserId": "224833315798889783",
	//
	//           "name": "sasdsadsa",
	//
	//           "emrAccessMode": "simple",
	//
	//           "region": "cn-shanghai",
	//
	//           "authType": "2",
	//
	//           "emrProjectId": "FP-sdadsad"
	//
	//         }
	//
	// 	- postgresql
	//
	//         {
	//
	//           "jdbcUrl": "jdbc:postgresql://xxxx:1921/ssss",
	//
	//           "password": "sdadsads",
	//
	//           "tag": "public",
	//
	//           "username": "sdsasda"
	//
	//         }
	//
	// 	- analyticdb_for_mysql
	//
	//         {
	//
	//           "instanceId": "am-sadsada",
	//
	//           "database": "xsxsx",
	//
	//           "username": "xsxsa",
	//
	//           "password": "asdadsa",
	//
	//           "connectionString": "am-xssxsxs.ads.aliyuncs.com:3306"
	//
	//         }
	//
	// 	- hybriddb_for_postgresql
	//
	//         {
	//
	//           "connectionString": "gp-xsaxsaxa-master.gpdbmaster.rds.aliyuncs.com",
	//
	//           "database": "xsaxsaxas",
	//
	//           "password": "xsaxsaxsa@11",
	//
	//           "instanceId": "gp-xsaxsaxsa",
	//
	//           "port": "541132",
	//
	//           "ownerId": "xsaxsaxsas",
	//
	//           "username": "sadsad"
	//
	//         }
	//
	// 	- holo
	//
	//         {
	//
	//           "accessId": "*****",
	//
	//           "accessKey": "*****",
	//
	//           "database": "xsaxsaxsa",
	//
	//           "instanceId": "xsaxa",
	//
	//           "tag": "aliyun"
	//
	//         }
	//
	// 	- kafka
	//
	//         {
	//
	//           "instanceId": "xsax-cn-xsaxsa",
	//
	//           "regionId": "cn-shanghai",
	//
	//           "tag": "aliyun",
	//
	//           "ownerId": "1212121212112"
	//
	//         }
	//
	// example:
	//
	// {"accessId":"xssssss","accessKey":"xsaxsaxsa","authType":2,"endpoint":"http://service.odps.aliyun.com/api","project":"xsaxsax","tag":"public"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the data source. You can call the [ListDataSources](https://help.aliyun.com/document_detail/2780072.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment in which the data source resides. Valid values:
	//
	// 	- 0: development environment
	//
	// 	- 1: production environment
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The status of the data source. This parameter is deprecated. Do not use this parameter.
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateDataSourceRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *UpdateDataSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataSourceRequest) GetEnvType() *int32 {
	return s.EnvType
}

func (s *UpdateDataSourceRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateDataSourceRequest) SetContent(v string) *UpdateDataSourceRequest {
	s.Content = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDataSourceId(v int64) *UpdateDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDescription(v string) *UpdateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataSourceRequest) SetEnvType(v int32) *UpdateDataSourceRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateDataSourceRequest) SetStatus(v string) *UpdateDataSourceRequest {
	s.Status = &v
	return s
}

func (s *UpdateDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
