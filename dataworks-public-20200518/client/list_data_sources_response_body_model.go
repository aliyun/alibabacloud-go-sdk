// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataSourcesResponseBodyData) *ListDataSourcesResponseBody
	GetData() *ListDataSourcesResponseBodyData
	SetHttpStatusCode(v int32) *ListDataSourcesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListDataSourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataSourcesResponseBody
	GetSuccess() *bool
}

type ListDataSourcesResponseBody struct {
	// The query result returned.
	Data *ListDataSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) GetData() *ListDataSourcesResponseBodyData {
	return s.Data
}

func (s *ListDataSourcesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataSourcesResponseBody) SetData(v *ListDataSourcesResponseBodyData) *ListDataSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListDataSourcesResponseBody) SetHttpStatusCode(v int32) *ListDataSourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetSuccess(v bool) *ListDataSourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataSourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataSourcesResponseBodyData struct {
	// The data sources.
	DataSources []*ListDataSourcesResponseBodyDataDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of data sources.
	//
	// example:
	//
	// 233
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyData) GetDataSources() []*ListDataSourcesResponseBodyDataDataSources {
	return s.DataSources
}

func (s *ListDataSourcesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataSourcesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSourcesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataSourcesResponseBodyData) SetDataSources(v []*ListDataSourcesResponseBodyDataDataSources) *ListDataSourcesResponseBodyData {
	s.DataSources = v
	return s
}

func (s *ListDataSourcesResponseBodyData) SetPageNumber(v int32) *ListDataSourcesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesResponseBodyData) SetPageSize(v int32) *ListDataSourcesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesResponseBodyData) SetTotalCount(v int32) *ListDataSourcesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDataSourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDataSourcesResponseBodyDataDataSources struct {
	// The ID of the compute engine with which the data source is associated.
	//
	// example:
	//
	// 123
	BindingCalcEngineId *int64 `json:"BindingCalcEngineId,omitempty" xml:"BindingCalcEngineId,omitempty"`
	// The status of the data source. Valid values:
	//
	// 	- 1: The data source is accessible.
	//
	// 	- 2: The data source is inaccessible.
	//
	// example:
	//
	// 1
	ConnectStatus *int32 `json:"ConnectStatus,omitempty" xml:"ConnectStatus,omitempty"`
	// The data connection string. The value of this parameter is in the JSON format. Examples of connection strings of common data sources:
	//
	// 	- MaxCompute
	//
	//         {
	//
	//           "pubEndpoint": "http://service.cn.maxcompute.aliyun.com/api",
	//
	//           "accessId": "xxxxxxx",
	//
	//           "securityToken": null,
	//
	//           "endpoint": "http://service.cn.maxcompute.aliyun-inc.com/api",
	//
	//           "accessKey": "***",
	//
	//           "name": "PRE_PROJECT_A_engine",
	//
	//           "project": "PRE_PROJECT_A",
	//
	//           "vpcEndpoint": "http://service.cn.maxcompute.aliyun-inc.com/api",
	//
	//           "region": "cn-shanghai",
	//
	//           "authType": "2"
	//
	//         }
	//
	// 	- mysql
	//
	//         {
	//
	//           "configType": "1",
	//
	//           "database": "mysql_d111b",
	//
	//           "instanceName": "rm-xxxxxx",
	//
	//           "password": "***",
	//
	//           "rdsOwnerId": "12133xxxxxx",
	//
	//           "tag": "rds",
	//
	//           "username": "mysql_db111"
	//
	//         }
	//
	// 	- sqlserver
	//
	//         {
	//
	//           "configType": "1",
	//
	//           "jdbcUrl": "jdbc:sqlserver://rm-xxxxx.sqlserver.rds.aliyuncs.com:1433;DatabaseName=sqlserver_db1",
	//
	//           "password": "***",
	//
	//           "tag": "public",
	//
	//           "username": "sqlserver_db111"
	//
	//         }
	//
	// 	- oss
	//
	//         {
	//
	//           "accessId": "***********",
	//
	//           "accessKey": "***********",
	//
	//           "bucket": "bigxxx1223",
	//
	//           "configType": "1",
	//
	//           "endpoint": "http://oss-cn-hangzhou.aliyuncs.com",
	//
	//           "tag": "public"
	//
	//         }
	//
	// 	- postgresql
	//
	//         {
	//
	//           "configType": "1",
	//
	//           "database": "cdp_xxx",
	//
	//           "instanceName": "rm-xxxx",
	//
	//           "password": "***",
	//
	//           "rdsOwnerId": "121xxxxx",
	//
	//           "tag": "rds",
	//
	//           "username": "cdp_xxx"
	//
	//         }
	//
	// 	- ads
	//
	//         {
	//
	//           "configType": "1",
	//
	//           "password": "***",
	//
	//           "schema": "ads_demo",
	//
	//           "tag": "public",
	//
	//           "url": "ads-xxx-xxxx.cn-hangzhou-1.ads.aliyuncs.com:3029",
	//
	//           "username": "lslslsls"
	//
	//         }
	//
	// example:
	//
	// {"pubEndpoint":"http://service.cn.maxcompute.aliyun.com/api","accessId":"TMP.3KecGjvzy3i8MYfn2BGHgF7EHGyBFZcHm7GgngrABVRyvvKQrfF5kskR36xP361C3dqwbGo7SGYptAeGyiTwHXqLaBUvYC","securityToken":null,"endpoint":"http://service.cn.maxcompute.aliyun-inc.com/api","accessKey":"***","name":"PRE_PROJECT_A_engine","project":"PRE_PROJECT_A","vpcEndpoint":"http://service.cn.maxcompute.aliyun-inc.com/api","region":"cn-shanghai","authType":"2"}
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
	// example:
	//
	// rds
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
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
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The ID of the workspace to which the data source belongs.
	//
	// example:
	//
	// 123
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sequence number of the data source. Data sources are sorted in descending order based on the value of this parameter.
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
	// 	- 1: The data source is accessible.
	//
	// 	- 2: The data source is inaccessible.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subtype of the data source. This parameter takes effect only when the DataSourceType parameter is set to rds.
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

func (s ListDataSourcesResponseBodyDataDataSources) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyDataDataSources) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetBindingCalcEngineId() *int64 {
	return s.BindingCalcEngineId
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetConnectStatus() *int32 {
	return s.ConnectStatus
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetContent() *string {
	return s.Content
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetDefaultEngine() *bool {
	return s.DefaultEngine
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetDescription() *string {
	return s.Description
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetEnvType() *int32 {
	return s.EnvType
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetId() *int64 {
	return s.Id
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetName() *string {
	return s.Name
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetOperator() *string {
	return s.Operator
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetShared() *bool {
	return s.Shared
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetStatus() *int32 {
	return s.Status
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetSubType() *string {
	return s.SubType
}

func (s *ListDataSourcesResponseBodyDataDataSources) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetBindingCalcEngineId(v int64) *ListDataSourcesResponseBodyDataDataSources {
	s.BindingCalcEngineId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetConnectStatus(v int32) *ListDataSourcesResponseBodyDataDataSources {
	s.ConnectStatus = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetContent(v string) *ListDataSourcesResponseBodyDataDataSources {
	s.Content = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetDataSourceType(v string) *ListDataSourcesResponseBodyDataDataSources {
	s.DataSourceType = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetDefaultEngine(v bool) *ListDataSourcesResponseBodyDataDataSources {
	s.DefaultEngine = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetDescription(v string) *ListDataSourcesResponseBodyDataDataSources {
	s.Description = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetEnvType(v int32) *ListDataSourcesResponseBodyDataDataSources {
	s.EnvType = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetGmtCreate(v string) *ListDataSourcesResponseBodyDataDataSources {
	s.GmtCreate = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetGmtModified(v string) *ListDataSourcesResponseBodyDataDataSources {
	s.GmtModified = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetId(v int64) *ListDataSourcesResponseBodyDataDataSources {
	s.Id = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetName(v string) *ListDataSourcesResponseBodyDataDataSources {
	s.Name = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetOperator(v string) *ListDataSourcesResponseBodyDataDataSources {
	s.Operator = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetProjectId(v int32) *ListDataSourcesResponseBodyDataDataSources {
	s.ProjectId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetSequence(v int32) *ListDataSourcesResponseBodyDataDataSources {
	s.Sequence = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetShared(v bool) *ListDataSourcesResponseBodyDataDataSources {
	s.Shared = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetStatus(v int32) *ListDataSourcesResponseBodyDataDataSources {
	s.Status = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetSubType(v string) *ListDataSourcesResponseBodyDataDataSources {
	s.SubType = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) SetTenantId(v int64) *ListDataSourcesResponseBodyDataDataSources {
	s.TenantId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataDataSources) Validate() error {
	return dara.Validate(s)
}
