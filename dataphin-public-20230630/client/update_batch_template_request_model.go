// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBatchTemplateRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateBatchTemplateRequest
	GetOpUserId() *string
	SetProjectId(v int64) *UpdateBatchTemplateRequest
	GetProjectId() *int64
	SetUpdateCommand(v *UpdateBatchTemplateRequestUpdateCommand) *UpdateBatchTemplateRequest
	GetUpdateCommand() *UpdateBatchTemplateRequestUpdateCommand
}

type UpdateBatchTemplateRequest struct {
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
	// The template update command.
	//
	// This parameter is required.
	UpdateCommand *UpdateBatchTemplateRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateBatchTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateBatchTemplateRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBatchTemplateRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateBatchTemplateRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateBatchTemplateRequest) GetUpdateCommand() *UpdateBatchTemplateRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateBatchTemplateRequest) SetOpTenantId(v int64) *UpdateBatchTemplateRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBatchTemplateRequest) SetOpUserId(v string) *UpdateBatchTemplateRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateBatchTemplateRequest) SetProjectId(v int64) *UpdateBatchTemplateRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateBatchTemplateRequest) SetUpdateCommand(v *UpdateBatchTemplateRequestUpdateCommand) *UpdateBatchTemplateRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateBatchTemplateRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBatchTemplateRequestUpdateCommand struct {
	// The template content.
	//
	// This parameter is required.
	//
	// example:
	//
	// echo 123
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The datasource config for database SQL templates. This parameter takes effect only for database SQL templates.
	DataSourceConfig *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig `json:"DataSourceConfig,omitempty" xml:"DataSourceConfig,omitempty" type:"Struct"`
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
	// The folder of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// /Data development/Templates
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Data processing template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of template parameters.
	Params []*UpdateBatchTemplateRequestUpdateCommandParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
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
	// The ID of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateBatchTemplateRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTemplateRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateBatchTemplateRequestUpdateCommand) GetContent() *string {
	return s.Content
}

func (s *UpdateBatchTemplateRequestUpdateCommand) GetDataSourceConfig() *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig {
	return s.DataSourceConfig
}

func (s *UpdateBatchTemplateRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateBatchTemplateRequestUpdateCommand) GetDevResourceGroupId() *string {
	return s.DevResourceGroupId
}

func (s *UpdateBatchTemplateRequestUpdateCommand) GetDirectory() *string {
	return s.Directory
}

func (s *UpdateBatchTemplateRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateBatchTemplateRequestUpdateCommand) GetParams() []*UpdateBatchTemplateRequestUpdateCommandParams {
	return s.Params
}

func (s *UpdateBatchTemplateRequestUpdateCommand) GetPythonModuleIds() []*int64 {
	return s.PythonModuleIds
}

func (s *UpdateBatchTemplateRequestUpdateCommand) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateBatchTemplateRequestUpdateCommand) GetShareRuntimeResource() *bool {
	return s.ShareRuntimeResource
}

func (s *UpdateBatchTemplateRequestUpdateCommand) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *UpdateBatchTemplateRequestUpdateCommand) SetContent(v string) *UpdateBatchTemplateRequestUpdateCommand {
	s.Content = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommand) SetDataSourceConfig(v *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) *UpdateBatchTemplateRequestUpdateCommand {
	s.DataSourceConfig = v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommand) SetDescription(v string) *UpdateBatchTemplateRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommand) SetDevResourceGroupId(v string) *UpdateBatchTemplateRequestUpdateCommand {
	s.DevResourceGroupId = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommand) SetDirectory(v string) *UpdateBatchTemplateRequestUpdateCommand {
	s.Directory = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommand) SetName(v string) *UpdateBatchTemplateRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommand) SetParams(v []*UpdateBatchTemplateRequestUpdateCommandParams) *UpdateBatchTemplateRequestUpdateCommand {
	s.Params = v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommand) SetPythonModuleIds(v []*int64) *UpdateBatchTemplateRequestUpdateCommand {
	s.PythonModuleIds = v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommand) SetResourceGroupId(v string) *UpdateBatchTemplateRequestUpdateCommand {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommand) SetShareRuntimeResource(v bool) *UpdateBatchTemplateRequestUpdateCommand {
	s.ShareRuntimeResource = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommand) SetTemplateId(v int64) *UpdateBatchTemplateRequestUpdateCommand {
	s.TemplateId = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommand) Validate() error {
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

type UpdateBatchTemplateRequestUpdateCommandDataSourceConfig struct {
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

func (s UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) GoString() string {
	return s.String()
}

func (s *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) GetDsId() *string {
	return s.DsId
}

func (s *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) GetSchema() *string {
	return s.Schema
}

func (s *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) GetType() *string {
	return s.Type
}

func (s *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) SetCatalog(v string) *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig {
	s.Catalog = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) SetDsId(v string) *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig {
	s.DsId = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) SetSchema(v string) *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig {
	s.Schema = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) SetType(v string) *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig {
	s.Type = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommandDataSourceConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateBatchTemplateRequestUpdateCommandParams struct {
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

func (s UpdateBatchTemplateRequestUpdateCommandParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTemplateRequestUpdateCommandParams) GoString() string {
	return s.String()
}

func (s *UpdateBatchTemplateRequestUpdateCommandParams) GetDescription() *string {
	return s.Description
}

func (s *UpdateBatchTemplateRequestUpdateCommandParams) GetEncryptEnabled() *bool {
	return s.EncryptEnabled
}

func (s *UpdateBatchTemplateRequestUpdateCommandParams) GetKey() *string {
	return s.Key
}

func (s *UpdateBatchTemplateRequestUpdateCommandParams) GetValue() *string {
	return s.Value
}

func (s *UpdateBatchTemplateRequestUpdateCommandParams) SetDescription(v string) *UpdateBatchTemplateRequestUpdateCommandParams {
	s.Description = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommandParams) SetEncryptEnabled(v bool) *UpdateBatchTemplateRequestUpdateCommandParams {
	s.EncryptEnabled = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommandParams) SetKey(v string) *UpdateBatchTemplateRequestUpdateCommandParams {
	s.Key = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommandParams) SetValue(v string) *UpdateBatchTemplateRequestUpdateCommandParams {
	s.Value = &v
	return s
}

func (s *UpdateBatchTemplateRequestUpdateCommandParams) Validate() error {
	return dara.Validate(s)
}
