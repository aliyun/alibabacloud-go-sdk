// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListConnectionsResponseBodyData) *ListConnectionsResponseBody
	GetData() *ListConnectionsResponseBodyData
	SetHttpStatusCode(v int32) *ListConnectionsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListConnectionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListConnectionsResponseBody
	GetSuccess() *bool
}

type ListConnectionsResponseBody struct {
	// The query results for data sources that are returned on multiple pages.
	Data *ListConnectionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc14115159376359****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBody) GetData() *ListConnectionsResponseBodyData {
	return s.Data
}

func (s *ListConnectionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConnectionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListConnectionsResponseBody) SetData(v *ListConnectionsResponseBodyData) *ListConnectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListConnectionsResponseBody) SetHttpStatusCode(v int32) *ListConnectionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConnectionsResponseBody) SetRequestId(v string) *ListConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectionsResponseBody) SetSuccess(v bool) *ListConnectionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListConnectionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyData struct {
	// The data sources.
	Connections []*ListConnectionsResponseBodyDataConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of data sources returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConnectionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyData) GetConnections() []*ListConnectionsResponseBodyDataConnections {
	return s.Connections
}

func (s *ListConnectionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConnectionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConnectionsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListConnectionsResponseBodyData) SetConnections(v []*ListConnectionsResponseBodyDataConnections) *ListConnectionsResponseBodyData {
	s.Connections = v
	return s
}

func (s *ListConnectionsResponseBodyData) SetPageNumber(v int32) *ListConnectionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListConnectionsResponseBodyData) SetPageSize(v int32) *ListConnectionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListConnectionsResponseBodyData) SetTotalCount(v int32) *ListConnectionsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListConnectionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnections struct {
	// The ID of the compute engine with which the data source is associated.
	//
	// example:
	//
	// 123
	BindingCalcEngineId *int32 `json:"BindingCalcEngineId,omitempty" xml:"BindingCalcEngineId,omitempty"`
	// The status of the data source. Valid values:
	//
	// 	- 1: The data source is normal.
	//
	// 	- 2: The data source is disabled.
	//
	// example:
	//
	// 1
	ConnectStatus *int32 `json:"ConnectStatus,omitempty" xml:"ConnectStatus,omitempty"`
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
	// example:
	//
	// mysql
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
	// example:
	//
	// {\\"database\\":\\"xxx\\",\\"instanceName\\":\\"xxx\\",\\"password\\":\\"xxx\\",\\"rdsOwnerId\\":\\"xxx\\",\\"tag\\":\\"rds\\",\\"username\\":\\"xxx\\"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates whether the compute engine that is associated with the data source is the default compute engine used by data sources of the same type.
	//
	// example:
	//
	// false
	DefaultEngine *bool `json:"DefaultEngine,omitempty" xml:"DefaultEngine,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// a connection
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment in which the data source is used. Valid values:
	//
	// 	- 0: development environment
	//
	// 	- 1: production environment
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The time when the data source was created. Example: Mar 17, 2021 4:09:32 PM.
	//
	// example:
	//
	// Mar 17, 2021 4:09:32 PM
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the data source was last modified. Example: Mar 17, 2021 4:09:32 PM.
	//
	// example:
	//
	// Mar 17, 2021 4:09:32 PM
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account that is used to last modify the data source.
	//
	// example:
	//
	// 193543050****
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The ID of the workspace with which the data source is associated.
	//
	// example:
	//
	// 123
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The field that is used to sort data sources. Data sources are sorted in descending order based on the value of this parameter.
	//
	// example:
	//
	// 300
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Indicates whether the data source is a shared data source.
	//
	// example:
	//
	// false
	Shared *bool `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// The status of the data source. Valid values:
	//
	// 	- 1: The data source is normal.
	//
	// 	- 2: The data source is disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subtype of the data source. This parameter is used in scenarios where a type includes subtypes. The following type and subtypes are supported:
	//
	// 	- Type: `rds`
	//
	// 	- Subtypes: `mysql`, `sqlserver`, and `postgresql`.
	//
	// example:
	//
	// mysql
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnections) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnections) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnections) GetBindingCalcEngineId() *int32 {
	return s.BindingCalcEngineId
}

func (s *ListConnectionsResponseBodyDataConnections) GetConnectStatus() *int32 {
	return s.ConnectStatus
}

func (s *ListConnectionsResponseBodyDataConnections) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *ListConnectionsResponseBodyDataConnections) GetContent() *string {
	return s.Content
}

func (s *ListConnectionsResponseBodyDataConnections) GetDefaultEngine() *bool {
	return s.DefaultEngine
}

func (s *ListConnectionsResponseBodyDataConnections) GetDescription() *string {
	return s.Description
}

func (s *ListConnectionsResponseBodyDataConnections) GetEnvType() *int32 {
	return s.EnvType
}

func (s *ListConnectionsResponseBodyDataConnections) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListConnectionsResponseBodyDataConnections) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListConnectionsResponseBodyDataConnections) GetId() *int32 {
	return s.Id
}

func (s *ListConnectionsResponseBodyDataConnections) GetName() *string {
	return s.Name
}

func (s *ListConnectionsResponseBodyDataConnections) GetOperator() *string {
	return s.Operator
}

func (s *ListConnectionsResponseBodyDataConnections) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListConnectionsResponseBodyDataConnections) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListConnectionsResponseBodyDataConnections) GetShared() *bool {
	return s.Shared
}

func (s *ListConnectionsResponseBodyDataConnections) GetStatus() *int32 {
	return s.Status
}

func (s *ListConnectionsResponseBodyDataConnections) GetSubType() *string {
	return s.SubType
}

func (s *ListConnectionsResponseBodyDataConnections) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListConnectionsResponseBodyDataConnections) SetBindingCalcEngineId(v int32) *ListConnectionsResponseBodyDataConnections {
	s.BindingCalcEngineId = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetConnectStatus(v int32) *ListConnectionsResponseBodyDataConnections {
	s.ConnectStatus = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetConnectionType(v string) *ListConnectionsResponseBodyDataConnections {
	s.ConnectionType = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetContent(v string) *ListConnectionsResponseBodyDataConnections {
	s.Content = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetDefaultEngine(v bool) *ListConnectionsResponseBodyDataConnections {
	s.DefaultEngine = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetDescription(v string) *ListConnectionsResponseBodyDataConnections {
	s.Description = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetEnvType(v int32) *ListConnectionsResponseBodyDataConnections {
	s.EnvType = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetGmtCreate(v string) *ListConnectionsResponseBodyDataConnections {
	s.GmtCreate = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetGmtModified(v string) *ListConnectionsResponseBodyDataConnections {
	s.GmtModified = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetId(v int32) *ListConnectionsResponseBodyDataConnections {
	s.Id = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetName(v string) *ListConnectionsResponseBodyDataConnections {
	s.Name = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetOperator(v string) *ListConnectionsResponseBodyDataConnections {
	s.Operator = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetProjectId(v int32) *ListConnectionsResponseBodyDataConnections {
	s.ProjectId = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetSequence(v int32) *ListConnectionsResponseBodyDataConnections {
	s.Sequence = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetShared(v bool) *ListConnectionsResponseBodyDataConnections {
	s.Shared = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetStatus(v int32) *ListConnectionsResponseBodyDataConnections {
	s.Status = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetSubType(v string) *ListConnectionsResponseBodyDataConnections {
	s.SubType = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetTenantId(v int64) *ListConnectionsResponseBodyDataConnections {
	s.TenantId = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) Validate() error {
	return dara.Validate(s)
}
