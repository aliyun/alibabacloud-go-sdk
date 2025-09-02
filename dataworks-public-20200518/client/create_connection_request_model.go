// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionType(v string) *CreateConnectionRequest
	GetConnectionType() *string
	SetContent(v string) *CreateConnectionRequest
	GetContent() *string
	SetDescription(v string) *CreateConnectionRequest
	GetDescription() *string
	SetEnvType(v int32) *CreateConnectionRequest
	GetEnvType() *int32
	SetName(v string) *CreateConnectionRequest
	GetName() *string
	SetProjectId(v int64) *CreateConnectionRequest
	GetProjectId() *int64
	SetSubType(v string) *CreateConnectionRequest
	GetSubType() *string
}

type CreateConnectionRequest struct {
	// The type of the connection string. Valid values:
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
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The details of the data source. Examples of details of some common data sources:
	//
	// 	- odps
	//
	// <!---->
	//
	//     {
	//
	//       "accessId": "xssssss",
	//
	//       "accessKey": "xsaxsaxsa",
	//
	//       "authType": 2,
	//
	//       "endpoint": "http://service.odps.aliyun.com/api",
	//
	//       "project": "xsaxsax",
	//
	//       "tag": "public"
	//
	//     }
	//
	// 	- mysql
	//
	// <!---->
	//
	//     {
	//
	//       "database": "xsaxsa",
	//
	//       "instanceName": "rm-xsaxsa",
	//
	//       "password": "xsaxsa",
	//
	//       "rdsOwnerId": "xasxsa",
	//
	//       "regionId": "cn-shanghai",
	//
	//       "tag": "rds",
	//
	//       "username": "xsaxsa"
	//
	//     }
	//
	// 	- rds
	//
	// <!---->
	//
	//     {
	//
	//       "configType": 1,
	//
	//       "tag": "rds",
	//
	//       "database": "xsaxsa",
	//
	//       "username": "xsaxsa",
	//
	//       "password": "xssaxsa$32050",
	//
	//       "instanceName": "rm-xsaxs",
	//
	//       "rdsOwnerId": "11111111"
	//
	//     }
	//
	// 	- oss
	//
	// <!---->
	//
	//     {
	//
	//       "accessId": "sssssxx",
	//
	//       "accessKey": "xsaxaxsaxs",
	//
	//       "bucket": "xsa-xs-xs",
	//
	//       "endpoint": "http://oss-cn-shanghai.aliyuncs.com",
	//
	//       "tag": "public"
	//
	//     }
	//
	// 	- sqlserver
	//
	// <!---->
	//
	//     {
	//
	//       "jdbcUrl": "jdbc:sqlserver://xsaxsa-xsaxsa.database.xxx.cn:123;DatabaseName=xsxs-xsxs",
	//
	//       "password": "sdasda$fs",
	//
	//       "tag": "public",
	//
	//       "username": "sxaxacdacdd"
	//
	//     }
	//
	// 	- polardb
	//
	// <!---->
	//
	//     {
	//
	//       "clusterId": "pc-sdadsadsa",
	//
	//       "database": "dsadsadsa",
	//
	//       "ownerId": "121212122",
	//
	//       "password": "sdasdafssa",
	//
	//       "region": "cn-shanghai",
	//
	//       "tag": "polardb",
	//
	//       "username": "asdadsads"
	//
	//     }
	//
	// 	- oracle
	//
	// <!---->
	//
	//     {
	//
	//       "jdbcUrl": "jdbc:oracle:saaa:@xxxxx:1521:PROD",
	//
	//       "password": "sxasaxsa",
	//
	//       "tag": "public",
	//
	//       "username": "sasfadfa"
	//
	//     }
	//
	// 	- mongodb
	//
	// <!---->
	//
	//     {
	//
	//       "address": "[\\"xsaxxsa.mongodb.rds.aliyuncs.com:3717\\"]",
	//
	//       "database": "admin",
	//
	//       "password": "sadsda@",
	//
	//       "tag": "public",
	//
	//       "username": "dsadsadas"
	//
	//     }
	//
	// 	- emr
	//
	// <!---->
	//
	//     {
	//
	//       "accessId": "xsaxsa",
	//
	//       "emrClusterId": "C-dsads",
	//
	//       "emrResourceQueueName": "default",
	//
	//       "emrEndpoint": "emr.aliyuncs.com",
	//
	//       "accessKey": "dsadsad",
	//
	//       "emrUserId": "224833315798889783",
	//
	//       "name": "sasdsadsa",
	//
	//       "emrAccessMode": "simple",
	//
	//       "region": "cn-shanghai",
	//
	//       "authType": "2",
	//
	//       "emrProjectId": "FP-sdadsad"
	//
	//     }
	//
	// 	- postgresql
	//
	// <!---->
	//
	//     {
	//
	//       "jdbcUrl": "jdbc:postgresql://xxxx:1921/ssss",
	//
	//       "password": "sdadsads",
	//
	//       "tag": "public",
	//
	//       "username": "sdsasda"
	//
	//     }
	//
	// 	- analyticdb_for_mysql
	//
	// <!---->
	//
	//     {
	//
	//       "instanceId": "am-sadsada",
	//
	//       "database": "xsxsx",
	//
	//       "username": "xsxsa",
	//
	//       "password": "asdadsa",
	//
	//       "connectionString": "am-xssxsxs.ads.aliyuncs.com:3306"
	//
	//     }
	//
	// 	- hybriddb_for_postgresql
	//
	// <!---->
	//
	//     {
	//
	//       "connectionString": "gp-xsaxsaxa-master.gpdbmaster.rds.aliyuncs.com",
	//
	//       "database": "xsaxsaxas",
	//
	//       "password": "xsaxsaxsa@11",
	//
	//       "instanceId": "gp-xsaxsaxsa",
	//
	//       "port": "541132",
	//
	//       "ownerId": "xsaxsaxsas",
	//
	//       "username": "sadsad"
	//
	//     }
	//
	// 	- holo
	//
	// <!---->
	//
	//     {
	//
	//       "accessId": "xsaxsaxs",
	//
	//       "accessKey": "xsaxsaxsa",
	//
	//       "database": "xsaxsaxsa",
	//
	//       "instanceId": "xsaxa",
	//
	//       "tag": "aliyun"
	//
	//     }
	//
	// 	- kafka
	//
	// <!---->
	//
	//     {
	//
	//       "instanceId": "xsax-cn-xsaxsa",
	//
	//       "regionId": "cn-shanghai",
	//
	//       "tag": "aliyun",
	//
	//       "ownerId": "1212121212112"
	//
	//     }
	//
	// This parameter is required.
	//
	// example:
	//
	// {"database":"dbname","instanceName":"instancename","password":"password","rdsOwnerId":"123","username":"username"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the connection string.
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
	// The ID of the workspace with which the data source is associated. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The subtype of the connection string. This parameter is used for scenarios where a type includes subtypes. The following type and subtypes are supported:
	//
	// 	- Type: `rds`
	//
	// 	- Subtypes: `mysql`, `sqlserver`, and `postgresql`.
	//
	// example:
	//
	// mysql
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s CreateConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *CreateConnectionRequest) GetContent() *string {
	return s.Content
}

func (s *CreateConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConnectionRequest) GetEnvType() *int32 {
	return s.EnvType
}

func (s *CreateConnectionRequest) GetName() *string {
	return s.Name
}

func (s *CreateConnectionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateConnectionRequest) GetSubType() *string {
	return s.SubType
}

func (s *CreateConnectionRequest) SetConnectionType(v string) *CreateConnectionRequest {
	s.ConnectionType = &v
	return s
}

func (s *CreateConnectionRequest) SetContent(v string) *CreateConnectionRequest {
	s.Content = &v
	return s
}

func (s *CreateConnectionRequest) SetDescription(v string) *CreateConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreateConnectionRequest) SetEnvType(v int32) *CreateConnectionRequest {
	s.EnvType = &v
	return s
}

func (s *CreateConnectionRequest) SetName(v string) *CreateConnectionRequest {
	s.Name = &v
	return s
}

func (s *CreateConnectionRequest) SetProjectId(v int64) *CreateConnectionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateConnectionRequest) SetSubType(v string) *CreateConnectionRequest {
	s.SubType = &v
	return s
}

func (s *CreateConnectionRequest) Validate() error {
	return dara.Validate(s)
}
