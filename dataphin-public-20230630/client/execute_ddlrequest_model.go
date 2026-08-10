// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteDDLRequest interface {
  dara.Model
  String() string
  GoString() string
  SetContext(v *ExecuteDDLRequestContext) *ExecuteDDLRequest
  GetContext() *ExecuteDDLRequestContext 
  SetDDLCommand(v *ExecuteDDLRequestDDLCommand) *ExecuteDDLRequest
  GetDDLCommand() *ExecuteDDLRequestDDLCommand 
  SetOpTenantId(v int64) *ExecuteDDLRequest
  GetOpTenantId() *int64 
}

type ExecuteDDLRequest struct {
  // The request context information.
  // 
  // This parameter is required.
  Context *ExecuteDDLRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
  // The one-click table creation parameters.
  // 
  // This parameter is required.
  DDLCommand *ExecuteDDLRequestDDLCommand `json:"DDLCommand,omitempty" xml:"DDLCommand,omitempty" type:"Struct"`
  // The tenant ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 30001011
  OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ExecuteDDLRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDDLRequest) GoString() string {
  return s.String()
}

func (s *ExecuteDDLRequest) GetContext() *ExecuteDDLRequestContext  {
  return s.Context
}

func (s *ExecuteDDLRequest) GetDDLCommand() *ExecuteDDLRequestDDLCommand  {
  return s.DDLCommand
}

func (s *ExecuteDDLRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExecuteDDLRequest) SetContext(v *ExecuteDDLRequestContext) *ExecuteDDLRequest {
  s.Context = v
  return s
}

func (s *ExecuteDDLRequest) SetDDLCommand(v *ExecuteDDLRequestDDLCommand) *ExecuteDDLRequest {
  s.DDLCommand = v
  return s
}

func (s *ExecuteDDLRequest) SetOpTenantId(v int64) *ExecuteDDLRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExecuteDDLRequest) Validate() error {
  if s.Context != nil {
    if err := s.Context.Validate(); err != nil {
      return err
    }
  }
  if s.DDLCommand != nil {
    if err := s.DDLCommand.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteDDLRequestContext struct {
  // The current operating environment. Valid values:
  // 
  // - DEV: development environment.
  // 
  // - PROD: production environment.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // DEV
  Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
  // The ID of the project to which the integration pipeline task belongs.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 123
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ExecuteDDLRequestContext) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDDLRequestContext) GoString() string {
  return s.String()
}

func (s *ExecuteDDLRequestContext) GetEnv() *string  {
  return s.Env
}

func (s *ExecuteDDLRequestContext) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *ExecuteDDLRequestContext) SetEnv(v string) *ExecuteDDLRequestContext {
  s.Env = &v
  return s
}

func (s *ExecuteDDLRequestContext) SetProjectId(v int64) *ExecuteDDLRequestContext {
  s.ProjectId = &v
  return s
}

func (s *ExecuteDDLRequestContext) Validate() error {
  return dara.Validate(s)
}

type ExecuteDDLRequestDDLCommand struct {
  // The identifier of the data source, compute source, or dataset used for table creation.
  // 
  // This parameter is required.
  DatasourceId *ExecuteDDLRequestDDLCommandDatasourceId `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty" type:"Struct"`
  // The DDL statement for table creation.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // CREATE TABLE test (id bigint)
  Ddl *string `json:"Ddl,omitempty" xml:"Ddl,omitempty"`
  // Specifies whether to drop the table if it already exists.
  // 
  // example:
  // 
  // true
  DropTable *bool `json:"DropTable,omitempty" xml:"DropTable,omitempty"`
}

func (s ExecuteDDLRequestDDLCommand) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDDLRequestDDLCommand) GoString() string {
  return s.String()
}

func (s *ExecuteDDLRequestDDLCommand) GetDatasourceId() *ExecuteDDLRequestDDLCommandDatasourceId  {
  return s.DatasourceId
}

func (s *ExecuteDDLRequestDDLCommand) GetDdl() *string  {
  return s.Ddl
}

func (s *ExecuteDDLRequestDDLCommand) GetDropTable() *bool  {
  return s.DropTable
}

func (s *ExecuteDDLRequestDDLCommand) SetDatasourceId(v *ExecuteDDLRequestDDLCommandDatasourceId) *ExecuteDDLRequestDDLCommand {
  s.DatasourceId = v
  return s
}

func (s *ExecuteDDLRequestDDLCommand) SetDdl(v string) *ExecuteDDLRequestDDLCommand {
  s.Ddl = &v
  return s
}

func (s *ExecuteDDLRequestDDLCommand) SetDropTable(v bool) *ExecuteDDLRequestDDLCommand {
  s.DropTable = &v
  return s
}

func (s *ExecuteDDLRequestDDLCommand) Validate() error {
  if s.DatasourceId != nil {
    if err := s.DatasourceId.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteDDLRequestDDLCommandDatasourceId struct {
  // The catalog of the data source or compute cluster. This parameter is required only in OneCatalog scenarios.
  // 
  // example:
  // 
  // dummy_cdm_dev
  Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
  // The data source category. Valid values:
  // 
  // - DATA_SOURCE: physical data source.
  // 
  // - PROJECT_COMPUTE_SOURCE: compute source bound to a project.
  // 
  // - ONE_CATALOG: compute source or data source in multi-engine mode (OneCatalog).
  // 
  // - DATA_SET: dataset.
  // 
  // This parameter is optional. The system automatically infers the category based on other fields if this parameter is not specified.
  // 
  // example:
  // 
  // DATA_SOURCE
  DsCategory *string `json:"DsCategory,omitempty" xml:"DsCategory,omitempty"`
  // The ID of the data source, compute source, or dataset. This parameter is optional when DsCategory is set to PROJECT_COMPUTE_SOURCE.
  // 
  // example:
  // 
  // 123
  DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
  // The environment. Valid values:
  // 
  // - DEV: development environment.
  // 
  // - PROD: production environment.
  // 
  // example:
  // 
  // DEV
  Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
  // The catalog type when DsCategory is set to ONE_CATALOG. Valid values:
  // 
  // - COMPUTE_CLUSTER: compute cluster.
  // 
  // - DATA_SOURCE: physical data source.
  // 
  // example:
  // 
  // COMPUTE_CLUSTER
  OneCatalogType *string `json:"OneCatalogType,omitempty" xml:"OneCatalogType,omitempty"`
  // The ID of the project bound to the compute source. This parameter is required only when DsCategory is set to PROJECT_COMPUTE_SOURCE.
  // 
  // example:
  // 
  // 123
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
  // The dataset version. This parameter is required only when DsCategory is set to DATA_SET.
  // 
  // example:
  // 
  // 3
  Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ExecuteDDLRequestDDLCommandDatasourceId) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDDLRequestDDLCommandDatasourceId) GoString() string {
  return s.String()
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) GetCatalog() *string  {
  return s.Catalog
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) GetDsCategory() *string  {
  return s.DsCategory
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) GetDsId() *string  {
  return s.DsId
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) GetEnv() *string  {
  return s.Env
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) GetOneCatalogType() *string  {
  return s.OneCatalogType
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) GetVersion() *string  {
  return s.Version
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) SetCatalog(v string) *ExecuteDDLRequestDDLCommandDatasourceId {
  s.Catalog = &v
  return s
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) SetDsCategory(v string) *ExecuteDDLRequestDDLCommandDatasourceId {
  s.DsCategory = &v
  return s
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) SetDsId(v string) *ExecuteDDLRequestDDLCommandDatasourceId {
  s.DsId = &v
  return s
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) SetEnv(v string) *ExecuteDDLRequestDDLCommandDatasourceId {
  s.Env = &v
  return s
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) SetOneCatalogType(v string) *ExecuteDDLRequestDDLCommandDatasourceId {
  s.OneCatalogType = &v
  return s
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) SetProjectId(v int64) *ExecuteDDLRequestDDLCommandDatasourceId {
  s.ProjectId = &v
  return s
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) SetVersion(v string) *ExecuteDDLRequestDDLCommandDatasourceId {
  s.Version = &v
  return s
}

func (s *ExecuteDDLRequestDDLCommandDatasourceId) Validate() error {
  return dara.Validate(s)
}

