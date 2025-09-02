// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateDataSourceRequest
	GetContent() *string
	SetDataSourceType(v string) *CreateDataSourceRequest
	GetDataSourceType() *string
	SetDescription(v string) *CreateDataSourceRequest
	GetDescription() *string
	SetEnvType(v int32) *CreateDataSourceRequest
	GetEnvType() *int32
	SetName(v string) *CreateDataSourceRequest
	GetName() *string
	SetProjectId(v int64) *CreateDataSourceRequest
	GetProjectId() *int64
	SetSubType(v string) *CreateDataSourceRequest
	GetSubType() *string
}

type CreateDataSourceRequest struct {
	// The details of the data source. Examples of details of some common data sources:
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
	// 	- redis
	//
	//         {
	//
	//         "password": "xxxxxx",
	//
	//          "address":"[{\\"host\\":\\"xxxxxxx.redis.rds.aliyuncs.com\\",\\"port\\":6379}]",
	//
	//         "tag": "public"
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
	// This parameter is required.
	//
	// example:
	//
	// {"database":"dbname","instanceName":"instancename","password":"password","rdsOwnerId":"123","username":"username"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- odps
	//
	// 	- mysql
	//
	// 	- rds
	//
	// 	- oss
	//
	// 	- sqlserver
	//
	// 	- polardb
	//
	// 	- oracle
	//
	// 	- mongodb
	//
	// 	- emr
	//
	// 	- postgresql
	//
	// 	- analyticdb_for_mysql
	//
	// 	- hybriddb_for_postgresql
	//
	// 	- holo
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment in which the data source is used. Valid values: 0 and 1. The value 0 indicates the development environment. The value 1 indicates the production environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace to which the data source belongs. You can call the [ListProjects](https://help.aliyun.com/document_detail/2780068.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The subtype of the data source.
	//
	// 	- This parameter takes effect only if you set the DataSourceType parameter to rds.
	//
	// 	- If the DataSourceType parameter is set to rds, this parameter can be set to mysql, sqlserver, or postgresql.
	//
	// example:
	//
	// mysql
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s CreateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) GetContent() *string {
	return s.Content
}

func (s *CreateDataSourceRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CreateDataSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataSourceRequest) GetEnvType() *int32 {
	return s.EnvType
}

func (s *CreateDataSourceRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataSourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataSourceRequest) GetSubType() *string {
	return s.SubType
}

func (s *CreateDataSourceRequest) SetContent(v string) *CreateDataSourceRequest {
	s.Content = &v
	return s
}

func (s *CreateDataSourceRequest) SetDataSourceType(v string) *CreateDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateDataSourceRequest) SetDescription(v string) *CreateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateDataSourceRequest) SetEnvType(v int32) *CreateDataSourceRequest {
	s.EnvType = &v
	return s
}

func (s *CreateDataSourceRequest) SetName(v string) *CreateDataSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequest) SetProjectId(v int64) *CreateDataSourceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataSourceRequest) SetSubType(v string) *CreateDataSourceRequest {
	s.SubType = &v
	return s
}

func (s *CreateDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
