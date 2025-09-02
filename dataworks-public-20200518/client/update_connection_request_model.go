// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionId(v int64) *UpdateConnectionRequest
	GetConnectionId() *int64
	SetContent(v string) *UpdateConnectionRequest
	GetContent() *string
	SetDescription(v string) *UpdateConnectionRequest
	GetDescription() *string
	SetEnvType(v int32) *UpdateConnectionRequest
	GetEnvType() *int32
	SetStatus(v string) *UpdateConnectionRequest
	GetStatus() *string
}

type UpdateConnectionRequest struct {
	// The data source ID. You can call the [ListConnections](https://help.aliyun.com/document_detail/173911.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ConnectionId *int64 `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
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
	// example:
	//
	// {"database":"dbname","instanceName":"instancename","password":"password","rdsOwnerId":"123","username":"username"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment in which the data source is used. Valid values: 0 and 1. The value 0 indicates the development environment. The value 1 indicates the production environment.
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The status of the data source. Valid values: ENABLED and DISABLED. The value ENABLED indicates that the data source is in the normal state. The value DISABLED indicates that the data source is in an abnormal state.
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequest) GetConnectionId() *int64 {
	return s.ConnectionId
}

func (s *UpdateConnectionRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConnectionRequest) GetEnvType() *int32 {
	return s.EnvType
}

func (s *UpdateConnectionRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateConnectionRequest) SetConnectionId(v int64) *UpdateConnectionRequest {
	s.ConnectionId = &v
	return s
}

func (s *UpdateConnectionRequest) SetContent(v string) *UpdateConnectionRequest {
	s.Content = &v
	return s
}

func (s *UpdateConnectionRequest) SetDescription(v string) *UpdateConnectionRequest {
	s.Description = &v
	return s
}

func (s *UpdateConnectionRequest) SetEnvType(v int32) *UpdateConnectionRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateConnectionRequest) SetStatus(v string) *UpdateConnectionRequest {
	s.Status = &v
	return s
}

func (s *UpdateConnectionRequest) Validate() error {
	return dara.Validate(s)
}
