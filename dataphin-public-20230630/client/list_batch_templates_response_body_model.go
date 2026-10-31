// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBatchTemplatesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListBatchTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBatchTemplatesResponseBody
	GetMessage() *string
	SetPageResult(v *ListBatchTemplatesResponseBodyPageResult) *ListBatchTemplatesResponseBody
	GetPageResult() *ListBatchTemplatesResponseBodyPageResult
	SetRequestId(v string) *ListBatchTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBatchTemplatesResponseBody
	GetSuccess() *bool
}

type ListBatchTemplatesResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paging query result.
	PageResult *ListBatchTemplatesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBatchTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBatchTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBatchTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBatchTemplatesResponseBody) GetPageResult() *ListBatchTemplatesResponseBodyPageResult {
	return s.PageResult
}

func (s *ListBatchTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBatchTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBatchTemplatesResponseBody) SetCode(v string) *ListBatchTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListBatchTemplatesResponseBody) SetHttpStatusCode(v int32) *ListBatchTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBatchTemplatesResponseBody) SetMessage(v string) *ListBatchTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListBatchTemplatesResponseBody) SetPageResult(v *ListBatchTemplatesResponseBodyPageResult) *ListBatchTemplatesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListBatchTemplatesResponseBody) SetRequestId(v string) *ListBatchTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBatchTemplatesResponseBody) SetSuccess(v bool) *ListBatchTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListBatchTemplatesResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBatchTemplatesResponseBodyPageResult struct {
	// The paginated template record list.
	TemplateList []*ListBatchTemplatesResponseBodyPageResultTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBatchTemplatesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesResponseBodyPageResult) GetTemplateList() []*ListBatchTemplatesResponseBodyPageResultTemplateList {
	return s.TemplateList
}

func (s *ListBatchTemplatesResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBatchTemplatesResponseBodyPageResult) SetTemplateList(v []*ListBatchTemplatesResponseBodyPageResultTemplateList) *ListBatchTemplatesResponseBodyPageResult {
	s.TemplateList = v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResult) SetTotalCount(v int32) *ListBatchTemplatesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResult) Validate() error {
	if s.TemplateList != nil {
		for _, item := range s.TemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBatchTemplatesResponseBodyPageResultTemplateList struct {
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
	// The latest version.
	//
	// example:
	//
	// 1
	CurrentVersion *int64 `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The datasource config of the database SQL template.
	DataSourceConfig *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig `json:"DataSourceConfig,omitempty" xml:"DataSourceConfig,omitempty" type:"Struct"`
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
	// The ID of the template. This ID is the same as the menu tree node ID.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The user who last updated the template.
	//
	// example:
	//
	// Li Si
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The ID of the user who last updated the template.
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
	// The type of the node. Valid values: 10 (Shell) and 21 (Python).
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
	Params []*ListBatchTemplatesResponseBodyPageResultTemplateListParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The ID of the project.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of dependent Python third-party packages.
	PythonModuleList []*ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList `json:"PythonModuleList,omitempty" xml:"PythonModuleList,omitempty" type:"Repeated"`
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
}

func (s ListBatchTemplatesResponseBodyPageResultTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesResponseBodyPageResultTemplateList) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetComment() *string {
	return s.Comment
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetContent() *string {
	return s.Content
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetCurrentVersion() *int64 {
	return s.CurrentVersion
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetDataSourceConfig() *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig {
	return s.DataSourceConfig
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetDescription() *string {
	return s.Description
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetDevResourceGroupId() *string {
	return s.DevResourceGroupId
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetDevResourceGroupName() *string {
	return s.DevResourceGroupName
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetEngine() *string {
	return s.Engine
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetId() *int64 {
	return s.Id
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetModifier() *string {
	return s.Modifier
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetModifierId() *string {
	return s.ModifierId
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetName() *string {
	return s.Name
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetOperatorType() *int32 {
	return s.OperatorType
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetOwner() *string {
	return s.Owner
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetParams() []*ListBatchTemplatesResponseBodyPageResultTemplateListParams {
	return s.Params
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetPythonModuleList() []*ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList {
	return s.PythonModuleList
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetShareRuntimeResource() *bool {
	return s.ShareRuntimeResource
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetStatus() *string {
	return s.Status
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetComment(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Comment = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetContent(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Content = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetCurrentVersion(v int64) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.CurrentVersion = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetDataSourceConfig(v *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.DataSourceConfig = v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetDescription(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Description = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetDevResourceGroupId(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.DevResourceGroupId = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetDevResourceGroupName(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.DevResourceGroupName = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetEngine(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Engine = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetGmtCreate(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.GmtCreate = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetGmtModified(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.GmtModified = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetId(v int64) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Id = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetModifier(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Modifier = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetModifierId(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.ModifierId = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetName(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Name = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetOperatorType(v int32) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.OperatorType = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetOwner(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Owner = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetOwnerId(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.OwnerId = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetParams(v []*ListBatchTemplatesResponseBodyPageResultTemplateListParams) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Params = v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetProjectId(v int64) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.ProjectId = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetPythonModuleList(v []*ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.PythonModuleList = v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetResourceGroupId(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetResourceGroupName(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.ResourceGroupName = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetShareRuntimeResource(v bool) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.ShareRuntimeResource = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetStatus(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Status = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) Validate() error {
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

type ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig struct {
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

func (s ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) GetDsId() *string {
	return s.DsId
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) GetSchema() *string {
	return s.Schema
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) GetType() *string {
	return s.Type
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) SetCatalog(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig {
	s.Catalog = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) SetDsId(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig {
	s.DsId = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) SetSchema(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig {
	s.Schema = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) SetType(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig {
	s.Type = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListDataSourceConfig) Validate() error {
	return dara.Validate(s)
}

type ListBatchTemplatesResponseBodyPageResultTemplateListParams struct {
	// The parameter description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether encryption is enabled for the parameter.
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

func (s ListBatchTemplatesResponseBodyPageResultTemplateListParams) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesResponseBodyPageResultTemplateListParams) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListParams) GetDescription() *string {
	return s.Description
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListParams) GetEncryptEnabled() *bool {
	return s.EncryptEnabled
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListParams) GetKey() *string {
	return s.Key
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListParams) GetValue() *string {
	return s.Value
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListParams) SetDescription(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListParams {
	s.Description = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListParams) SetEncryptEnabled(v bool) *ListBatchTemplatesResponseBodyPageResultTemplateListParams {
	s.EncryptEnabled = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListParams) SetKey(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListParams {
	s.Key = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListParams) SetValue(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListParams {
	s.Value = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListParams) Validate() error {
	return dara.Validate(s)
}

type ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList struct {
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
	// The status of the third-party package. Valid values: INSTALL_SUCCESS (installed successfully), INSTALL_FAILED (installation failed), INSTALLING (installing), UPGRADE_SUCCESS (upgraded successfully), UPGRADE_FAILED (upgrade failed), and UPGRADING (upgrading).
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

func (s ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) GetId() *int64 {
	return s.Id
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) GetInstallTime() *string {
	return s.InstallTime
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) GetName() *string {
	return s.Name
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) GetPythonVersion() *string {
	return s.PythonVersion
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) GetStatus() *string {
	return s.Status
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) GetVersion() *string {
	return s.Version
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) SetId(v int64) *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList {
	s.Id = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) SetInstallTime(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList {
	s.InstallTime = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) SetName(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList {
	s.Name = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) SetPythonVersion(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList {
	s.PythonVersion = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) SetStatus(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList {
	s.Status = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) SetVersion(v string) *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList {
	s.Version = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateListPythonModuleList) Validate() error {
	return dara.Validate(s)
}
