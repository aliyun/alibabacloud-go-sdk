// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDataSourcesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDataSourceType(v string) *ExportDataSourcesRequest
  GetDataSourceType() *string 
  SetEnvType(v int32) *ExportDataSourcesRequest
  GetEnvType() *int32 
  SetName(v string) *ExportDataSourcesRequest
  GetName() *string 
  SetPageNumber(v int32) *ExportDataSourcesRequest
  GetPageNumber() *int32 
  SetPageSize(v int32) *ExportDataSourcesRequest
  GetPageSize() *int32 
  SetProjectId(v int64) *ExportDataSourcesRequest
  GetProjectId() *int64 
  SetSubType(v string) *ExportDataSourcesRequest
  GetSubType() *string 
}

type ExportDataSourcesRequest struct {
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
  // The keyword contained in the names of the data sources that you want to export. You can specify only one keyword. For example, if you set this parameter to test, you can call the ExportDataSources operation to export all data sources whose names contain test in the workspace.
  // 
  // example:
  // 
  // test
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The page number. Pages start from page 1.
  // 
  // example:
  // 
  // 1
  PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
  // The number of entries per page. Default value: 10. Maximum value: 100.
  // 
  // example:
  // 
  // 10
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  // The ID of the DataWorks workspace to which the data sources belong. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to query the ID of the workspace.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 10000
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
  // The subtype of the data source. This parameter takes effect only when the DataSourceType parameter is set to rds.
  // 
  // If the value of the DataSourceType parameter is rds, the value of this parameter can be mysql, sqlserver, or postgresql.
  // 
  // example:
  // 
  // mysql
  SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s ExportDataSourcesRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportDataSourcesRequest) GoString() string {
  return s.String()
}

func (s *ExportDataSourcesRequest) GetDataSourceType() *string  {
  return s.DataSourceType
}

func (s *ExportDataSourcesRequest) GetEnvType() *int32  {
  return s.EnvType
}

func (s *ExportDataSourcesRequest) GetName() *string  {
  return s.Name
}

func (s *ExportDataSourcesRequest) GetPageNumber() *int32  {
  return s.PageNumber
}

func (s *ExportDataSourcesRequest) GetPageSize() *int32  {
  return s.PageSize
}

func (s *ExportDataSourcesRequest) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *ExportDataSourcesRequest) GetSubType() *string  {
  return s.SubType
}

func (s *ExportDataSourcesRequest) SetDataSourceType(v string) *ExportDataSourcesRequest {
  s.DataSourceType = &v
  return s
}

func (s *ExportDataSourcesRequest) SetEnvType(v int32) *ExportDataSourcesRequest {
  s.EnvType = &v
  return s
}

func (s *ExportDataSourcesRequest) SetName(v string) *ExportDataSourcesRequest {
  s.Name = &v
  return s
}

func (s *ExportDataSourcesRequest) SetPageNumber(v int32) *ExportDataSourcesRequest {
  s.PageNumber = &v
  return s
}

func (s *ExportDataSourcesRequest) SetPageSize(v int32) *ExportDataSourcesRequest {
  s.PageSize = &v
  return s
}

func (s *ExportDataSourcesRequest) SetProjectId(v int64) *ExportDataSourcesRequest {
  s.ProjectId = &v
  return s
}

func (s *ExportDataSourcesRequest) SetSubType(v string) *ExportDataSourcesRequest {
  s.SubType = &v
  return s
}

func (s *ExportDataSourcesRequest) Validate() error {
  return dara.Validate(s)
}

