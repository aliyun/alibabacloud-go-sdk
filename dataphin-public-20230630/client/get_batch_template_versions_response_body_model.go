// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTemplateVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBatchTemplateVersionsResponseBody
	GetCode() *string
	SetData(v *GetBatchTemplateVersionsResponseBodyData) *GetBatchTemplateVersionsResponseBody
	GetData() *GetBatchTemplateVersionsResponseBodyData
	SetHttpStatusCode(v int32) *GetBatchTemplateVersionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBatchTemplateVersionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBatchTemplateVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBatchTemplateVersionsResponseBody
	GetSuccess() *bool
}

type GetBatchTemplateVersionsResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetBatchTemplateVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBatchTemplateVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTemplateVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchTemplateVersionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBatchTemplateVersionsResponseBody) GetData() *GetBatchTemplateVersionsResponseBodyData {
	return s.Data
}

func (s *GetBatchTemplateVersionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBatchTemplateVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBatchTemplateVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchTemplateVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBatchTemplateVersionsResponseBody) SetCode(v string) *GetBatchTemplateVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) SetData(v *GetBatchTemplateVersionsResponseBodyData) *GetBatchTemplateVersionsResponseBody {
	s.Data = v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) SetHttpStatusCode(v int32) *GetBatchTemplateVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) SetMessage(v string) *GetBatchTemplateVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) SetRequestId(v string) *GetBatchTemplateVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) SetSuccess(v bool) *GetBatchTemplateVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchTemplateVersionsResponseBodyData struct {
	// The template version list.
	TemplateVersionList []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionList `json:"TemplateVersionList,omitempty" xml:"TemplateVersionList,omitempty" type:"Repeated"`
}

func (s GetBatchTemplateVersionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTemplateVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBatchTemplateVersionsResponseBodyData) GetTemplateVersionList() []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	return s.TemplateVersionList
}

func (s *GetBatchTemplateVersionsResponseBodyData) SetTemplateVersionList(v []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) *GetBatchTemplateVersionsResponseBodyData {
	s.TemplateVersionList = v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyData) Validate() error {
	if s.TemplateVersionList != nil {
		for _, item := range s.TemplateVersionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchTemplateVersionsResponseBodyDataTemplateVersionList struct {
	// The template submission comment.
	//
	// example:
	//
	// Initial submission
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The template content.
	//
	// example:
	//
	// #!/bin/bash
	//
	// echo \\"hello world\\"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The data source configuration of the database SQL template.
	DataSourceConfig *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig `json:"DataSourceConfig,omitempty" xml:"DataSourceConfig,omitempty" type:"Struct"`
	// The template description.
	//
	// example:
	//
	// Shell template for data processing
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the resource group in the development environment.
	//
	// example:
	//
	// default
	DevResourceGroupId *string `json:"DevResourceGroupId,omitempty" xml:"DevResourceGroupId,omitempty"`
	// The name of the resource group in the development environment.
	//
	// example:
	//
	// Default resource group for the project
	DevResourceGroupName *string `json:"DevResourceGroupName,omitempty" xml:"DevResourceGroupName,omitempty"`
	// The version of the compute engine. Valid values: Python 2.7 and Python 3.7.
	//
	// example:
	//
	// Python 3.7
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The time when the template was created. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2026-01-01 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the template was updated. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2026-05-28 15:30:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the template, which is the same as the menu tree node ID.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The last user who updated the template.
	//
	// example:
	//
	// Li Si
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The ID of the last user who updated the template.
	//
	// example:
	//
	// 100002
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Data processing template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node type. Valid values: 10 (Shell) and 21 (Python).
	//
	// example:
	//
	// 10
	OperatorType *int32 `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// The owner of the template.
	//
	// example:
	//
	// Zhang San
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the template owner.
	//
	// example:
	//
	// 100001
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of template parameters.
	Params []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The ID of the project.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of dependent Python third-party packages.
	PythonModuleList []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList `json:"PythonModuleList,omitempty" xml:"PythonModuleList,omitempty" type:"Repeated"`
	// The ID of the resource group in the production environment.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the resource group in the production environment.
	//
	// example:
	//
	// Default resource group for the project
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// Indicates whether runtime resources are shared.
	ShareRuntimeResource *bool `json:"ShareRuntimeResource,omitempty" xml:"ShareRuntimeResource,omitempty"`
	// The status of the template. Valid values: 0 (draft), 2 (submitted), and 100 (in development).
	//
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GoString() string {
	return s.String()
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetComment() *string {
	return s.Comment
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetContent() *string {
	return s.Content
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetDataSourceConfig() *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig {
	return s.DataSourceConfig
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetDescription() *string {
	return s.Description
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetDevResourceGroupId() *string {
	return s.DevResourceGroupId
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetDevResourceGroupName() *string {
	return s.DevResourceGroupName
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetEngine() *string {
	return s.Engine
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetId() *int64 {
	return s.Id
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetModifier() *string {
	return s.Modifier
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetName() *string {
	return s.Name
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetOperatorType() *int32 {
	return s.OperatorType
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetOwner() *string {
	return s.Owner
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetParams() []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams {
	return s.Params
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetPythonModuleList() []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList {
	return s.PythonModuleList
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetShareRuntimeResource() *bool {
	return s.ShareRuntimeResource
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetStatus() *string {
	return s.Status
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetVersion() *int64 {
	return s.Version
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetComment(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Comment = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetContent(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Content = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetDataSourceConfig(v *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.DataSourceConfig = v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetDescription(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Description = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetDevResourceGroupId(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.DevResourceGroupId = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetDevResourceGroupName(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.DevResourceGroupName = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetEngine(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Engine = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetGmtCreate(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.GmtCreate = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetGmtModified(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.GmtModified = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetId(v int64) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Id = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetModifier(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Modifier = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetModifierId(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.ModifierId = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetName(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Name = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetOperatorType(v int32) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.OperatorType = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetOwner(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Owner = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetOwnerId(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.OwnerId = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetParams(v []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Params = v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetProjectId(v int64) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetPythonModuleList(v []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.PythonModuleList = v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetResourceGroupId(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.ResourceGroupId = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetResourceGroupName(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.ResourceGroupName = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetShareRuntimeResource(v bool) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.ShareRuntimeResource = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetStatus(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Status = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetVersion(v int64) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Version = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) Validate() error {
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
	if s.PythonModuleList != nil {
		for _, item := range s.PythonModuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig struct {
	// The catalog of the data source for the database SQL template.
	//
	// example:
	//
	// default
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The ID of the data source for the database SQL template.
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
	// The data source type of the database SQL template. Valid values: MYSQL, ORACLE, ADB_FOR_PG, OUSHUDB, PRESTO, POSTGRE_SQL, STARROCKS, CLICKHOUSE, ANALYTICDB, ADB_FOR_MYSQL_V3, GAUSSDB, SELECTDB, DORIS, SQL_SERVER, HOLOGRES, OPENGAUSS, DAMENG, OCEANBASE, TRINO, POLARDB_X2, LINDORM, EMR_SPARK_SERVERLESS, TDSQL_FOR_PG, TDSQL_FOR_MYSQL, GBASE_8A, and TDENGINE.
	//
	// example:
	//
	// MYSQL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) GoString() string {
	return s.String()
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) GetDsId() *string {
	return s.DsId
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) GetSchema() *string {
	return s.Schema
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) GetType() *string {
	return s.Type
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) SetCatalog(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig {
	s.Catalog = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) SetDsId(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig {
	s.DsId = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) SetSchema(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig {
	s.Schema = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) SetType(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig {
	s.Type = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListDataSourceConfig) Validate() error {
	return dara.Validate(s)
}

type GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams struct {
	// The parameter description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether encryption is enabled for the parameter.
	EncryptEnabled *bool `json:"EncryptEnabled,omitempty" xml:"EncryptEnabled,omitempty"`
	// The key of the parameter.
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

func (s GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) GoString() string {
	return s.String()
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) GetDescription() *string {
	return s.Description
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) GetEncryptEnabled() *bool {
	return s.EncryptEnabled
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) GetKey() *string {
	return s.Key
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) GetValue() *string {
	return s.Value
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) SetDescription(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams {
	s.Description = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) SetEncryptEnabled(v bool) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams {
	s.EncryptEnabled = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) SetKey(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams {
	s.Key = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) SetValue(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams {
	s.Value = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListParams) Validate() error {
	return dara.Validate(s)
}

type GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList struct {
	// The ID of the Python third-party package.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The installation time.
	//
	// example:
	//
	// 2026-01-01 10:00:00
	InstallTime *string `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	// The name of the Python third-party package.
	//
	// example:
	//
	// numpy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Python version. Valid values: PYTHON_2_7, PYTHON_3_7, and PYTHON_3_11.
	//
	// example:
	//
	// PYTHON_2_7
	PythonVersion *string `json:"PythonVersion,omitempty" xml:"PythonVersion,omitempty"`
	// The status of the third-party package. Valid values: INSTALL_SUCCESS (installed), INSTALL_FAILED (installation failed), INSTALLING (installing), UPGRADE_SUCCESS (upgraded), UPGRADE_FAILED (upgrade failed), and UPGRADING (upgrading).
	//
	// example:
	//
	// INSTALL_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version of the third-party package.
	//
	// example:
	//
	// 1.0.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) GoString() string {
	return s.String()
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) GetId() *int64 {
	return s.Id
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) GetInstallTime() *string {
	return s.InstallTime
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) GetName() *string {
	return s.Name
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) GetPythonVersion() *string {
	return s.PythonVersion
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) GetStatus() *string {
	return s.Status
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) GetVersion() *string {
	return s.Version
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) SetId(v int64) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList {
	s.Id = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) SetInstallTime(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList {
	s.InstallTime = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) SetName(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList {
	s.Name = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) SetPythonVersion(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList {
	s.PythonVersion = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) SetStatus(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList {
	s.Status = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) SetVersion(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList {
	s.Version = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionListPythonModuleList) Validate() error {
	return dara.Validate(s)
}
