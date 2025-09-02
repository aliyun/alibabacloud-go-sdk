// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDataSourcesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExportDataSourcesResponseBodyData) *ExportDataSourcesResponseBody
  GetData() *ExportDataSourcesResponseBodyData 
  SetHttpStatusCode(v int32) *ExportDataSourcesResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExportDataSourcesResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportDataSourcesResponseBody
  GetSuccess() *bool 
}

type ExportDataSourcesResponseBody struct {
  // The information about the exported data sources.
  Data *ExportDataSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The HTTP status code returned. Valid values:
  // 
  // 	- 200: The request was successful.
  // 
  // 	- Other values: The request failed. You can troubleshoot issues based on the HTTP status code returned.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The request ID. You can locate logs and troubleshoot issues based on the ID.
  // 
  // example:
  // 
  // 0bc14115159376359****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportDataSourcesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportDataSourcesResponseBody) GoString() string {
  return s.String()
}

func (s *ExportDataSourcesResponseBody) GetData() *ExportDataSourcesResponseBodyData  {
  return s.Data
}

func (s *ExportDataSourcesResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportDataSourcesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportDataSourcesResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportDataSourcesResponseBody) SetData(v *ExportDataSourcesResponseBodyData) *ExportDataSourcesResponseBody {
  s.Data = v
  return s
}

func (s *ExportDataSourcesResponseBody) SetHttpStatusCode(v int32) *ExportDataSourcesResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportDataSourcesResponseBody) SetRequestId(v string) *ExportDataSourcesResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportDataSourcesResponseBody) SetSuccess(v bool) *ExportDataSourcesResponseBody {
  s.Success = &v
  return s
}

func (s *ExportDataSourcesResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExportDataSourcesResponseBodyData struct {
  // The details of the exported data sources. The value is an array. The following parameters are the elements in the array. The sample values of these parameters show the details of a sample data source.
  DataSources []*ExportDataSourcesResponseBodyDataDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
  // The page number. Pages start from page 1.
  // 
  // example:
  // 
  // 10
  PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
  // The number of entries per page.
  // 
  // example:
  // 
  // 1
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  // The total number of entries returned.
  // 
  // example:
  // 
  // 100
  TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ExportDataSourcesResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExportDataSourcesResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExportDataSourcesResponseBodyData) GetDataSources() []*ExportDataSourcesResponseBodyDataDataSources  {
  return s.DataSources
}

func (s *ExportDataSourcesResponseBodyData) GetPageNumber() *int32  {
  return s.PageNumber
}

func (s *ExportDataSourcesResponseBodyData) GetPageSize() *int32  {
  return s.PageSize
}

func (s *ExportDataSourcesResponseBodyData) GetTotalCount() *int32  {
  return s.TotalCount
}

func (s *ExportDataSourcesResponseBodyData) SetDataSources(v []*ExportDataSourcesResponseBodyDataDataSources) *ExportDataSourcesResponseBodyData {
  s.DataSources = v
  return s
}

func (s *ExportDataSourcesResponseBodyData) SetPageNumber(v int32) *ExportDataSourcesResponseBodyData {
  s.PageNumber = &v
  return s
}

func (s *ExportDataSourcesResponseBodyData) SetPageSize(v int32) *ExportDataSourcesResponseBodyData {
  s.PageSize = &v
  return s
}

func (s *ExportDataSourcesResponseBodyData) SetTotalCount(v int32) *ExportDataSourcesResponseBodyData {
  s.TotalCount = &v
  return s
}

func (s *ExportDataSourcesResponseBodyData) Validate() error {
  return dara.Validate(s)
}

type ExportDataSourcesResponseBodyDataDataSources struct {
  // The ID of the compute engine that is added as the data source.
  // 
  // example:
  // 
  // 123
  BindingCalcEngineId *int32 `json:"BindingCalcEngineId,omitempty" xml:"BindingCalcEngineId,omitempty"`
  // Indicates whether the data source is connected to an exclusive resource group. Valid values:
  // 
  // 	- 1: The data source is connected to at least one exclusive resource group.
  // 
  // 	- 0: The data source is not connected to any exclusive resource group.
  // 
  // example:
  // 
  // 1
  ConnectStatus *int32 `json:"ConnectStatus,omitempty" xml:"ConnectStatus,omitempty"`
  // The configuration of the data source.
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
  // Indicates whether the compute engine that is added as the data source is the default compute engine. Valid values:
  // 
  // 	- true: The compute engine is the default compute engine.
  // 
  // 	- false: The compute engine is not the default compute engine.
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
  // The ID of the user who exported the data source.
  // 
  // example:
  // 
  // 193543050****
  Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
  // The ID of the DataWorks workspace to which the data source belongs.
  // 
  // example:
  // 
  // 123
  ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
  // The sequence number of the data source.
  // 
  // example:
  // 
  // 300
  Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
  // Indicates whether the data source can be shared. Valid values:
  // 
  // 	- true: The data source can be shared.
  // 
  // 	- false: The data source cannot be shared.
  // 
  // example:
  // 
  // false
  Shared *bool `json:"Shared,omitempty" xml:"Shared,omitempty"`
  // Indicates whether the data source is available. Valid values:
  // 
  // 	- 1: The data source is available.
  // 
  // 	- 0: The data source is unavailable.
  // 
  // example:
  // 
  // 1
  Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
  // The subtype of the data source. This parameter takes effect only when the DataSourceType parameter is set to rds.
  // 
  // If the value of the DataSourceType parameter is rds, the value of this parameter can be mysql, sqlserver, or postgresql.
  // 
  // example:
  // 
  // mysql
  SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
  // The ID of the Alibaba Cloud account to which the data source belongs.
  // 
  // example:
  // 
  // 1234567
  TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ExportDataSourcesResponseBodyDataDataSources) String() string {
  return dara.Prettify(s)
}

func (s ExportDataSourcesResponseBodyDataDataSources) GoString() string {
  return s.String()
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetBindingCalcEngineId() *int32  {
  return s.BindingCalcEngineId
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetConnectStatus() *int32  {
  return s.ConnectStatus
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetContent() *string  {
  return s.Content
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetDataSourceType() *string  {
  return s.DataSourceType
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetDefaultEngine() *bool  {
  return s.DefaultEngine
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetDescription() *string  {
  return s.Description
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetEnvType() *int32  {
  return s.EnvType
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetGmtCreate() *string  {
  return s.GmtCreate
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetGmtModified() *string  {
  return s.GmtModified
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetId() *int32  {
  return s.Id
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetName() *string  {
  return s.Name
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetOperator() *string  {
  return s.Operator
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetProjectId() *int32  {
  return s.ProjectId
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetSequence() *int32  {
  return s.Sequence
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetShared() *bool  {
  return s.Shared
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetStatus() *int32  {
  return s.Status
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetSubType() *string  {
  return s.SubType
}

func (s *ExportDataSourcesResponseBodyDataDataSources) GetTenantId() *int64  {
  return s.TenantId
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetBindingCalcEngineId(v int32) *ExportDataSourcesResponseBodyDataDataSources {
  s.BindingCalcEngineId = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetConnectStatus(v int32) *ExportDataSourcesResponseBodyDataDataSources {
  s.ConnectStatus = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetContent(v string) *ExportDataSourcesResponseBodyDataDataSources {
  s.Content = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetDataSourceType(v string) *ExportDataSourcesResponseBodyDataDataSources {
  s.DataSourceType = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetDefaultEngine(v bool) *ExportDataSourcesResponseBodyDataDataSources {
  s.DefaultEngine = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetDescription(v string) *ExportDataSourcesResponseBodyDataDataSources {
  s.Description = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetEnvType(v int32) *ExportDataSourcesResponseBodyDataDataSources {
  s.EnvType = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetGmtCreate(v string) *ExportDataSourcesResponseBodyDataDataSources {
  s.GmtCreate = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetGmtModified(v string) *ExportDataSourcesResponseBodyDataDataSources {
  s.GmtModified = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetId(v int32) *ExportDataSourcesResponseBodyDataDataSources {
  s.Id = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetName(v string) *ExportDataSourcesResponseBodyDataDataSources {
  s.Name = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetOperator(v string) *ExportDataSourcesResponseBodyDataDataSources {
  s.Operator = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetProjectId(v int32) *ExportDataSourcesResponseBodyDataDataSources {
  s.ProjectId = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetSequence(v int32) *ExportDataSourcesResponseBodyDataDataSources {
  s.Sequence = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetShared(v bool) *ExportDataSourcesResponseBodyDataDataSources {
  s.Shared = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetStatus(v int32) *ExportDataSourcesResponseBodyDataDataSources {
  s.Status = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetSubType(v string) *ExportDataSourcesResponseBodyDataDataSources {
  s.SubType = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) SetTenantId(v int64) *ExportDataSourcesResponseBodyDataDataSources {
  s.TenantId = &v
  return s
}

func (s *ExportDataSourcesResponseBodyDataDataSources) Validate() error {
  return dara.Validate(s)
}

