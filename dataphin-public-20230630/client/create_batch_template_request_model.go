// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateBatchTemplateRequestCreateCommand) *CreateBatchTemplateRequest
	GetCreateCommand() *CreateBatchTemplateRequestCreateCommand
	SetOpTenantId(v int64) *CreateBatchTemplateRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateBatchTemplateRequest
	GetOpUserId() *string
	SetProjectId(v int64) *CreateBatchTemplateRequest
	GetProjectId() *int64
}

type CreateBatchTemplateRequest struct {
	// The command to create a template.
	//
	// This parameter is required.
	CreateCommand *CreateBatchTemplateRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateBatchTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTemplateRequest) GetCreateCommand() *CreateBatchTemplateRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateBatchTemplateRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBatchTemplateRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateBatchTemplateRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateBatchTemplateRequest) SetCreateCommand(v *CreateBatchTemplateRequestCreateCommand) *CreateBatchTemplateRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateBatchTemplateRequest) SetOpTenantId(v int64) *CreateBatchTemplateRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBatchTemplateRequest) SetOpUserId(v string) *CreateBatchTemplateRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateBatchTemplateRequest) SetProjectId(v int64) *CreateBatchTemplateRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateBatchTemplateRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBatchTemplateRequestCreateCommand struct {
	// The content of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// echo 123
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The datasource config for database SQL templates, which takes effect only for database SQL templates.
	DataSourceConfig *CreateBatchTemplateRequestCreateCommandDataSourceConfig `json:"DataSourceConfig,omitempty" xml:"DataSourceConfig,omitempty" type:"Struct"`
	// The description of the template.
	//
	// example:
	//
	// Shell template for data processing
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the resource group in the development environment. By default, the default resource group of the project is used.
	//
	// example:
	//
	// default
	DevResourceGroupId *string `json:"DevResourceGroupId,omitempty" xml:"DevResourceGroupId,omitempty"`
	// The directory of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// /Data Development/Templates
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The compute engine version. Valid values: PYTHON_2_7, PYTHON_3_7, and PYTHON_3_11.
	//
	// example:
	//
	// PYTHON_3_7
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Data processing template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node type of the template. Valid values: 10: Shell, 21: Python, 998: Database SQL, 5: MaxCompute SQL, 6: MapReduce on MaxCompute, 8: Spark on MaxCompute, 28: Hologres SQL, 1: Hive SQL, 2: MapReduce on Yarn, 18: Spark on Yarn, 78: Impala SQL, 80: Spark SQL, 10000: Inceptor SQL, 51: AnalyticDB for PostgreSQL SQL, 81: GaussDB SQL, 82: Databricks SQL, 76: ArgoDB SQL, 79: StarRocks SQL, 202: SelectDB SQL, 201: Doris SQL, and 203: OushuDB SQL.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	OperatorType *int32 `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// The list of template parameters.
	Params []*CreateBatchTemplateRequestCreateCommandParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The list of IDs of dependent third-party Python packages.
	PythonModuleIds []*int64 `json:"PythonModuleIds,omitempty" xml:"PythonModuleIds,omitempty" type:"Repeated"`
	// The ID of the resource group in the production environment. By default, the default resource group of the project is used.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to share runtime resources.
	ShareRuntimeResource *bool `json:"ShareRuntimeResource,omitempty" xml:"ShareRuntimeResource,omitempty"`
}

func (s CreateBatchTemplateRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTemplateRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateBatchTemplateRequestCreateCommand) GetContent() *string {
	return s.Content
}

func (s *CreateBatchTemplateRequestCreateCommand) GetDataSourceConfig() *CreateBatchTemplateRequestCreateCommandDataSourceConfig {
	return s.DataSourceConfig
}

func (s *CreateBatchTemplateRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateBatchTemplateRequestCreateCommand) GetDevResourceGroupId() *string {
	return s.DevResourceGroupId
}

func (s *CreateBatchTemplateRequestCreateCommand) GetDirectory() *string {
	return s.Directory
}

func (s *CreateBatchTemplateRequestCreateCommand) GetEngine() *string {
	return s.Engine
}

func (s *CreateBatchTemplateRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateBatchTemplateRequestCreateCommand) GetOperatorType() *int32 {
	return s.OperatorType
}

func (s *CreateBatchTemplateRequestCreateCommand) GetParams() []*CreateBatchTemplateRequestCreateCommandParams {
	return s.Params
}

func (s *CreateBatchTemplateRequestCreateCommand) GetPythonModuleIds() []*int64 {
	return s.PythonModuleIds
}

func (s *CreateBatchTemplateRequestCreateCommand) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateBatchTemplateRequestCreateCommand) GetShareRuntimeResource() *bool {
	return s.ShareRuntimeResource
}

func (s *CreateBatchTemplateRequestCreateCommand) SetContent(v string) *CreateBatchTemplateRequestCreateCommand {
	s.Content = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) SetDataSourceConfig(v *CreateBatchTemplateRequestCreateCommandDataSourceConfig) *CreateBatchTemplateRequestCreateCommand {
	s.DataSourceConfig = v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) SetDescription(v string) *CreateBatchTemplateRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) SetDevResourceGroupId(v string) *CreateBatchTemplateRequestCreateCommand {
	s.DevResourceGroupId = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) SetDirectory(v string) *CreateBatchTemplateRequestCreateCommand {
	s.Directory = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) SetEngine(v string) *CreateBatchTemplateRequestCreateCommand {
	s.Engine = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) SetName(v string) *CreateBatchTemplateRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) SetOperatorType(v int32) *CreateBatchTemplateRequestCreateCommand {
	s.OperatorType = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) SetParams(v []*CreateBatchTemplateRequestCreateCommandParams) *CreateBatchTemplateRequestCreateCommand {
	s.Params = v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) SetPythonModuleIds(v []*int64) *CreateBatchTemplateRequestCreateCommand {
	s.PythonModuleIds = v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) SetResourceGroupId(v string) *CreateBatchTemplateRequestCreateCommand {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) SetShareRuntimeResource(v bool) *CreateBatchTemplateRequestCreateCommand {
	s.ShareRuntimeResource = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommand) Validate() error {
	if s.DataSourceConfig != nil {
		if err := s.DataSourceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Params != nil {
		for _, item := range s.Params {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateBatchTemplateRequestCreateCommandDataSourceConfig struct {
	// The catalog of the data source for the database SQL template.
	//
	// example:
	//
	// default
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The ID of the data source for the database SQL template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// The schema of the data source for the database SQL template.
	//
	// example:
	//
	// default
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The type of the data source for the database SQL template. Valid values: MYSQL, ORACLE, ADB_FOR_PG, OUSHUDB, PRESTO, POSTGRE_SQL, STARROCKS, CLICKHOUSE, ANALYTICDB, ADB_FOR_MYSQL_V3, GAUSSDB, SELECTDB, DORIS, SQL_SERVER, HOLOGRES, OPENGAUSS, DAMENG, OCEANBASE, TRINO, POLARDB_X2, LINDORM, EMR_SPARK_SERVERLESS, TDSQL_FOR_PG, TDSQL_FOR_MYSQL, GBASE_8A, and TDENGINE.
	//
	// This parameter is required.
	//
	// example:
	//
	// MYSQL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateBatchTemplateRequestCreateCommandDataSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTemplateRequestCreateCommandDataSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateBatchTemplateRequestCreateCommandDataSourceConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *CreateBatchTemplateRequestCreateCommandDataSourceConfig) GetDsId() *string {
	return s.DsId
}

func (s *CreateBatchTemplateRequestCreateCommandDataSourceConfig) GetSchema() *string {
	return s.Schema
}

func (s *CreateBatchTemplateRequestCreateCommandDataSourceConfig) GetType() *string {
	return s.Type
}

func (s *CreateBatchTemplateRequestCreateCommandDataSourceConfig) SetCatalog(v string) *CreateBatchTemplateRequestCreateCommandDataSourceConfig {
	s.Catalog = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommandDataSourceConfig) SetDsId(v string) *CreateBatchTemplateRequestCreateCommandDataSourceConfig {
	s.DsId = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommandDataSourceConfig) SetSchema(v string) *CreateBatchTemplateRequestCreateCommandDataSourceConfig {
	s.Schema = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommandDataSourceConfig) SetType(v string) *CreateBatchTemplateRequestCreateCommandDataSourceConfig {
	s.Type = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommandDataSourceConfig) Validate() error {
	return dara.Validate(s)
}

type CreateBatchTemplateRequestCreateCommandParams struct {
	// The parameter description.
	//
	// example:
	//
	// Input date
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether encryption is enabled for the parameter.
	EncryptEnabled *bool `json:"EncryptEnabled,omitempty" xml:"EncryptEnabled,omitempty"`
	// The key of the parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// param1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The default value of the parameter.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateBatchTemplateRequestCreateCommandParams) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTemplateRequestCreateCommandParams) GoString() string {
	return s.String()
}

func (s *CreateBatchTemplateRequestCreateCommandParams) GetDescription() *string {
	return s.Description
}

func (s *CreateBatchTemplateRequestCreateCommandParams) GetEncryptEnabled() *bool {
	return s.EncryptEnabled
}

func (s *CreateBatchTemplateRequestCreateCommandParams) GetKey() *string {
	return s.Key
}

func (s *CreateBatchTemplateRequestCreateCommandParams) GetValue() *string {
	return s.Value
}

func (s *CreateBatchTemplateRequestCreateCommandParams) SetDescription(v string) *CreateBatchTemplateRequestCreateCommandParams {
	s.Description = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommandParams) SetEncryptEnabled(v bool) *CreateBatchTemplateRequestCreateCommandParams {
	s.EncryptEnabled = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommandParams) SetKey(v string) *CreateBatchTemplateRequestCreateCommandParams {
	s.Key = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommandParams) SetValue(v string) *CreateBatchTemplateRequestCreateCommandParams {
	s.Value = &v
	return s
}

func (s *CreateBatchTemplateRequestCreateCommandParams) Validate() error {
	return dara.Validate(s)
}
