// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDatasetRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateDatasetRequest
	GetOpUserId() *string
	SetProjectId(v string) *UpdateDatasetRequest
	GetProjectId() *string
	SetUpdateCommand(v *UpdateDatasetRequestUpdateCommand) *UpdateDatasetRequest
	GetUpdateCommand() *UpdateDatasetRequestUpdateCommand
}

type UpdateDatasetRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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
	// 7273382541481536
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The update request struct.
	//
	// This parameter is required.
	UpdateCommand *UpdateDatasetRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDatasetRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateDatasetRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateDatasetRequest) GetUpdateCommand() *UpdateDatasetRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateDatasetRequest) SetOpTenantId(v int64) *UpdateDatasetRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDatasetRequest) SetOpUserId(v string) *UpdateDatasetRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateDatasetRequest) SetProjectId(v string) *UpdateDatasetRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDatasetRequest) SetUpdateCommand(v *UpdateDatasetRequestUpdateCommand) *UpdateDatasetRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateDatasetRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommand struct {
	ApiInfo *UpdateDatasetRequestUpdateCommandApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Struct"`
	// **The content type.**
	//
	// example:
	//
	// GENERAL
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The subject area ID.
	//
	// example:
	//
	// 78201
	DataCellId *string `json:"DataCellId,omitempty" xml:"DataCellId,omitempty"`
	// **The description.**
	//
	// example:
	//
	// Test dataset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file ID (the file ID at creation time).
	//
	// This parameter is required.
	//
	// example:
	//
	// 7261110566632832
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The dataset ID (business primary key).
	//
	// This parameter is required.
	//
	// example:
	//
	// 7280832407583104
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// **The metastore type.**
	//
	// example:
	//
	// POSTGRESQL
	MetadataStorageType *string `json:"MetadataStorageType,omitempty" xml:"MetadataStorageType,omitempty"`
	// The dataset name.
	//
	// example:
	//
	// audio_dataset
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of owner IDs, separated by commas.
	//
	// example:
	//
	// 300001391
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// **Scenarios:*	- `OFFLINE` (offline, default) / `REALTIME` (real-time).
	//
	// example:
	//
	// OFFLINE
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// **The storage type.**
	//
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The dataset type.
	//
	// example:
	//
	// HYBRID
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version.
	//
	// example:
	//
	// V1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The dataset version configuration.
	VersionConfig *UpdateDatasetRequestUpdateCommandVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommand) GetApiInfo() *UpdateDatasetRequestUpdateCommandApiInfo {
	return s.ApiInfo
}

func (s *UpdateDatasetRequestUpdateCommand) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateDatasetRequestUpdateCommand) GetDataCellId() *string {
	return s.DataCellId
}

func (s *UpdateDatasetRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateDatasetRequestUpdateCommand) GetFileId() *string {
	return s.FileId
}

func (s *UpdateDatasetRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateDatasetRequestUpdateCommand) GetMetadataStorageType() *string {
	return s.MetadataStorageType
}

func (s *UpdateDatasetRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateDatasetRequestUpdateCommand) GetOwner() *string {
	return s.Owner
}

func (s *UpdateDatasetRequestUpdateCommand) GetScenario() *string {
	return s.Scenario
}

func (s *UpdateDatasetRequestUpdateCommand) GetStorageType() *string {
	return s.StorageType
}

func (s *UpdateDatasetRequestUpdateCommand) GetType() *string {
	return s.Type
}

func (s *UpdateDatasetRequestUpdateCommand) GetVersion() *string {
	return s.Version
}

func (s *UpdateDatasetRequestUpdateCommand) GetVersionConfig() *UpdateDatasetRequestUpdateCommandVersionConfig {
	return s.VersionConfig
}

func (s *UpdateDatasetRequestUpdateCommand) SetApiInfo(v *UpdateDatasetRequestUpdateCommandApiInfo) *UpdateDatasetRequestUpdateCommand {
	s.ApiInfo = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetContentType(v string) *UpdateDatasetRequestUpdateCommand {
	s.ContentType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetDataCellId(v string) *UpdateDatasetRequestUpdateCommand {
	s.DataCellId = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetDescription(v string) *UpdateDatasetRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetFileId(v string) *UpdateDatasetRequestUpdateCommand {
	s.FileId = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetId(v int64) *UpdateDatasetRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetMetadataStorageType(v string) *UpdateDatasetRequestUpdateCommand {
	s.MetadataStorageType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetName(v string) *UpdateDatasetRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetOwner(v string) *UpdateDatasetRequestUpdateCommand {
	s.Owner = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetScenario(v string) *UpdateDatasetRequestUpdateCommand {
	s.Scenario = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetStorageType(v string) *UpdateDatasetRequestUpdateCommand {
	s.StorageType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetType(v string) *UpdateDatasetRequestUpdateCommand {
	s.Type = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetVersion(v string) *UpdateDatasetRequestUpdateCommand {
	s.Version = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetVersionConfig(v *UpdateDatasetRequestUpdateCommandVersionConfig) *UpdateDatasetRequestUpdateCommand {
	s.VersionConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) Validate() error {
	if s.ApiInfo != nil {
		if err := s.ApiInfo.Validate(); err != nil {
			return err
		}
	}
	if s.VersionConfig != nil {
		if err := s.VersionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandApiInfo struct {
	// example:
	//
	// 60
	ExecTimeout *int32 `json:"ExecTimeout,omitempty" xml:"ExecTimeout,omitempty"`
	// example:
	//
	// 1
	ExecuteMode *int32 `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// example:
	//
	// 1011
	OsApiGroup *int32 `json:"OsApiGroup,omitempty" xml:"OsApiGroup,omitempty"`
	// example:
	//
	// 1012
	OsProject *int32 `json:"OsProject,omitempty" xml:"OsProject,omitempty"`
	// example:
	//
	// 1
	Protocol *int32 `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 1
	RequestMethod     *int32                                                       `json:"RequestMethod,omitempty" xml:"RequestMethod,omitempty"`
	RequestParamList  []*UpdateDatasetRequestUpdateCommandApiInfoRequestParamList  `json:"RequestParamList,omitempty" xml:"RequestParamList,omitempty" type:"Repeated"`
	ResponseParamList []*UpdateDatasetRequestUpdateCommandApiInfoResponseParamList `json:"ResponseParamList,omitempty" xml:"ResponseParamList,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateDatasetRequestUpdateCommandApiInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandApiInfo) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) GetExecTimeout() *int32 {
	return s.ExecTimeout
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) GetExecuteMode() *int32 {
	return s.ExecuteMode
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) GetOsApiGroup() *int32 {
	return s.OsApiGroup
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) GetOsProject() *int32 {
	return s.OsProject
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) GetProtocol() *int32 {
	return s.Protocol
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) GetRequestMethod() *int32 {
	return s.RequestMethod
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) GetRequestParamList() []*UpdateDatasetRequestUpdateCommandApiInfoRequestParamList {
	return s.RequestParamList
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) GetResponseParamList() []*UpdateDatasetRequestUpdateCommandApiInfoResponseParamList {
	return s.ResponseParamList
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) SetExecTimeout(v int32) *UpdateDatasetRequestUpdateCommandApiInfo {
	s.ExecTimeout = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) SetExecuteMode(v int32) *UpdateDatasetRequestUpdateCommandApiInfo {
	s.ExecuteMode = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) SetOsApiGroup(v int32) *UpdateDatasetRequestUpdateCommandApiInfo {
	s.OsApiGroup = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) SetOsProject(v int32) *UpdateDatasetRequestUpdateCommandApiInfo {
	s.OsProject = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) SetProtocol(v int32) *UpdateDatasetRequestUpdateCommandApiInfo {
	s.Protocol = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) SetRequestMethod(v int32) *UpdateDatasetRequestUpdateCommandApiInfo {
	s.RequestMethod = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) SetRequestParamList(v []*UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) *UpdateDatasetRequestUpdateCommandApiInfo {
	s.RequestParamList = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) SetResponseParamList(v []*UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) *UpdateDatasetRequestUpdateCommandApiInfo {
	s.ResponseParamList = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) SetTimeout(v int32) *UpdateDatasetRequestUpdateCommandApiInfo {
	s.Timeout = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfo) Validate() error {
	if s.RequestParamList != nil {
		for _, item := range s.RequestParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResponseParamList != nil {
		for _, item := range s.ResponseParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandApiInfoRequestParamList struct {
	// example:
	//
	// 1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// test
	Descr *string `json:"Descr,omitempty" xml:"Descr,omitempty"`
	IsUrl *bool   `json:"IsUrl,omitempty" xml:"IsUrl,omitempty"`
	Must  *bool   `json:"Must,omitempty" xml:"Must,omitempty"`
	// example:
	//
	// col01
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// example:
	//
	// int
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// 1
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
}

func (s UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) GetDescr() *string {
	return s.Descr
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) GetIsUrl() *bool {
	return s.IsUrl
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) GetMust() *bool {
	return s.Must
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) GetParamName() *string {
	return s.ParamName
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) GetParamType() *string {
	return s.ParamType
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) GetSample() *string {
	return s.Sample
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) SetDefaultValue(v string) *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList {
	s.DefaultValue = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) SetDescr(v string) *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList {
	s.Descr = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) SetIsUrl(v bool) *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList {
	s.IsUrl = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) SetMust(v bool) *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList {
	s.Must = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) SetParamName(v string) *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList {
	s.ParamName = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) SetParamType(v string) *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList {
	s.ParamType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) SetSample(v string) *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList {
	s.Sample = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoRequestParamList) Validate() error {
	return dara.Validate(s)
}

type UpdateDatasetRequestUpdateCommandApiInfoResponseParamList struct {
	// example:
	//
	// test
	Descr *string `json:"Descr,omitempty" xml:"Descr,omitempty"`
	IsUrl *bool   `json:"IsUrl,omitempty" xml:"IsUrl,omitempty"`
	// example:
	//
	// col01
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// example:
	//
	// int
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// 1
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
}

func (s UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) GetDescr() *string {
	return s.Descr
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) GetIsUrl() *bool {
	return s.IsUrl
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) GetParamName() *string {
	return s.ParamName
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) GetParamType() *string {
	return s.ParamType
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) GetSample() *string {
	return s.Sample
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) SetDescr(v string) *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList {
	s.Descr = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) SetIsUrl(v bool) *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList {
	s.IsUrl = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) SetParamName(v string) *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList {
	s.ParamName = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) SetParamType(v string) *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList {
	s.ParamType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) SetSample(v string) *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList {
	s.Sample = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandApiInfoResponseParamList) Validate() error {
	return dara.Validate(s)
}

type UpdateDatasetRequestUpdateCommandVersionConfig struct {
	// The file storage configuration.
	FileStorageConfig *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig `json:"FileStorageConfig,omitempty" xml:"FileStorageConfig,omitempty" type:"Struct"`
	// The metastore configuration.
	MetadataStorageConfig *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig `json:"MetadataStorageConfig,omitempty" xml:"MetadataStorageConfig,omitempty" type:"Struct"`
	// The real-time meta table configuration. Takes effect when metadataStorageType is set to STREAM_TABLE.
	RealtimeMetaTableConfig *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig `json:"RealtimeMetaTableConfig,omitempty" xml:"RealtimeMetaTableConfig,omitempty" type:"Struct"`
	// **Version description**
	//
	// example:
	//
	// Test dataset version.
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) GetFileStorageConfig() *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	return s.FileStorageConfig
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) GetMetadataStorageConfig() *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	return s.MetadataStorageConfig
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) GetRealtimeMetaTableConfig() *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig {
	return s.RealtimeMetaTableConfig
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) SetFileStorageConfig(v *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) *UpdateDatasetRequestUpdateCommandVersionConfig {
	s.FileStorageConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) SetMetadataStorageConfig(v *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) *UpdateDatasetRequestUpdateCommandVersionConfig {
	s.MetadataStorageConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) SetRealtimeMetaTableConfig(v *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) *UpdateDatasetRequestUpdateCommandVersionConfig {
	s.RealtimeMetaTableConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) SetVersionDescription(v string) *UpdateDatasetRequestUpdateCommandVersionConfig {
	s.VersionDescription = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) Validate() error {
	if s.FileStorageConfig != nil {
		if err := s.FileStorageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MetadataStorageConfig != nil {
		if err := s.MetadataStorageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RealtimeMetaTableConfig != nil {
		if err := s.RealtimeMetaTableConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig struct {
	// **The data source ID.**
	//
	// This parameter is required.
	//
	// example:
	//
	// 7445343860022804608
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The data source name.
	//
	// example:
	//
	// Test data source.
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The development path (not required for basic projects).
	//
	// example:
	//
	// HTML正文提取/test423/
	DevPath *string `json:"DevPath,omitempty" xml:"DevPath,omitempty"`
	// The mount path.
	//
	// This parameter is required.
	//
	// example:
	//
	// /var/run/openresty/cache/corp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The production path.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTML正文提取/test423/
	ProdPath *string `json:"ProdPath,omitempty" xml:"ProdPath,omitempty"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GetDevPath() *string {
	return s.DevPath
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GetProdPath() *string {
	return s.ProdPath
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) SetDataSourceId(v string) *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	s.DataSourceId = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) SetDataSourceName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	s.DataSourceName = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) SetDevPath(v string) *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	s.DevPath = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) SetMountPath(v string) *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	s.MountPath = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) SetProdPath(v string) *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	s.ProdPath = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig struct {
	// **The data source ID.**
	//
	// This parameter is required.
	//
	// example:
	//
	// 7429133693081710272
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// **The data source name.**
	//
	// example:
	//
	// Test data source.
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// **The development database/schema.**
	//
	// example:
	//
	// HTML正文提取/test423/
	DevSchema *string `json:"DevSchema,omitempty" xml:"DevSchema,omitempty"`
	// Specifies whether to store metadata in a new table or an existing table.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREATE
	MetadataStorageMode *string `json:"MetadataStorageMode,omitempty" xml:"MetadataStorageMode,omitempty"`
	// **The metastore type.**
	//
	// example:
	//
	// MILVUS
	MetadataStorageType *string `json:"MetadataStorageType,omitempty" xml:"MetadataStorageType,omitempty"`
	// **The production database/schema.**
	//
	// This parameter is required.
	//
	// example:
	//
	// HTML正文提取/test423/
	ProdSchema *string `json:"ProdSchema,omitempty" xml:"ProdSchema,omitempty"`
	// **The table name.**
	//
	// This parameter is required.
	//
	// example:
	//
	// s_crm_all_plt_jala_shop
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The table schema.
	TableSchema *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetDevSchema() *string {
	return s.DevSchema
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetMetadataStorageMode() *string {
	return s.MetadataStorageMode
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetMetadataStorageType() *string {
	return s.MetadataStorageType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetProdSchema() *string {
	return s.ProdSchema
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetTableName() *string {
	return s.TableName
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetTableSchema() *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema {
	return s.TableSchema
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetDataSourceId(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.DataSourceId = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetDataSourceName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.DataSourceName = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetDevSchema(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.DevSchema = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetMetadataStorageMode(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.MetadataStorageMode = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetMetadataStorageType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.MetadataStorageType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetProdSchema(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.ProdSchema = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetTableName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.TableName = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetTableSchema(v *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.TableSchema = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) Validate() error {
	if s.TableSchema != nil {
		if err := s.TableSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema struct {
	// The field list.
	Columns []*UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) GetColumns() []*UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	return s.Columns
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) SetColumns(v []*UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema {
	s.Columns = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns struct {
	// The field description.
	//
	// example:
	//
	// primary key
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// **The array element subtype. Valid only when type is set to ARRAY.**
	//
	// example:
	//
	// INT64
	ElementType *string `json:"ElementType,omitempty" xml:"ElementType,omitempty"`
	// The maximum capacity of the array. Valid only when type is set to ARRAY. Default value: 4096.
	//
	// example:
	//
	// 250
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// **The field name.**
	//
	// This parameter is required.
	//
	// example:
	//
	// id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the field is a primary key.
	//
	// example:
	//
	// true
	Pk *bool `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// **The field type.**
	//
	// This parameter is required.
	//
	// example:
	//
	// int8
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Indicates whether the field is a URL.
	//
	// example:
	//
	// false
	Url *bool `json:"Url,omitempty" xml:"Url,omitempty"`
	// The vector index configuration.
	VectorIndexConfig *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetElementType() *string {
	return s.ElementType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetName() *string {
	return s.Name
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetPk() *bool {
	return s.Pk
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetType() *string {
	return s.Type
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetUrl() *bool {
	return s.Url
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetVectorIndexConfig() *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetComment(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Comment = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetElementType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.ElementType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetMaxCapacity(v int32) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.MaxCapacity = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Name = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetPk(v bool) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Pk = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Type = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetUrl(v bool) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Url = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetVectorIndexConfig(v *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.VectorIndexConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) Validate() error {
	if s.VectorIndexConfig != nil {
		if err := s.VectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig struct {
	// The embedding dimension.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance:mongodb
	Dimension *int64 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The embedding model.
	//
	// This parameter is required.
	//
	// example:
	//
	// MultiModal-Embedding
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// The index build parameters. Different parameters are required based on the indexType. For example, HNSW requires {M:30, efConstruction:360}, and IVF_FLAT requires {nlist:128}.
	//
	// example:
	//
	// {M:30, efConstruction:360}
	IndexParams map[string]interface{} `json:"IndexParams,omitempty" xml:"IndexParams,omitempty"`
	// The index type. PostgreSQL supports IVFFlat and HNSW. Milvus supports all types.
	//
	// This parameter is required.
	//
	// example:
	//
	// AUTOINDEX
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// The similarity type. Default value: COSINE. Valid values: COSINE, L2, and IP.
	//
	// This parameter is required.
	//
	// example:
	//
	// COSINE
	SimilarityType *string `json:"SimilarityType,omitempty" xml:"SimilarityType,omitempty"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetDimension() *int64 {
	return s.Dimension
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetIndexParams() map[string]interface{} {
	return s.IndexParams
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetIndexType() *string {
	return s.IndexType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetSimilarityType() *string {
	return s.SimilarityType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetDimension(v int64) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.Dimension = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetEmbeddingModel(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.EmbeddingModel = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetIndexParams(v map[string]interface{}) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexParams = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetIndexType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetSimilarityType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.SimilarityType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig struct {
	// The meta table data source type (only KAFKA is supported in the current release).
	//
	// This parameter is required.
	//
	// example:
	//
	// KAFKA
	DatasourceType *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	// The meta table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test meta table.
	MetaTableName *string `json:"MetaTableName,omitempty" xml:"MetaTableName,omitempty"`
	// The project ID to which the meta table belongs (cross-project access is supported).
	//
	// This parameter is required.
	//
	// example:
	//
	// 7255013756724992
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The table schema.
	TableSchema *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) GetDatasourceType() *string {
	return s.DatasourceType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) GetMetaTableName() *string {
	return s.MetaTableName
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) GetTableSchema() *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema {
	return s.TableSchema
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) SetDatasourceType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig {
	s.DatasourceType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) SetMetaTableName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig {
	s.MetaTableName = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) SetProjectId(v int64) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig {
	s.ProjectId = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) SetTableSchema(v *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig {
	s.TableSchema = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) Validate() error {
	if s.TableSchema != nil {
		if err := s.TableSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema struct {
	// **The field list.**
	Columns []*UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) GetColumns() []*UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	return s.Columns
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) SetColumns(v []*UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema {
	s.Columns = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns struct {
	// The field description.
	//
	// example:
	//
	// happen time
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// **The array element subtype. Valid only when type is set to ARRAY.**
	//
	// example:
	//
	// INT64
	ElementType *string `json:"ElementType,omitempty" xml:"ElementType,omitempty"`
	// **The maximum capacity of the array. This parameter is valid only when type is set to ARRAY. Default value: 4096.**
	//
	// example:
	//
	// 10
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// **The field name.**
	//
	// This parameter is required.
	//
	// example:
	//
	// happen_time
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the field is a primary key.
	//
	// example:
	//
	// false
	Pk *bool `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// **The field type.**
	//
	// This parameter is required.
	//
	// example:
	//
	// date
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Indicates whether the field is a URL.
	//
	// example:
	//
	// false
	Url *bool `json:"Url,omitempty" xml:"Url,omitempty"`
	// The vector index configuration. Configure this parameter when the field type is FLOAT_VECTOR, FLOAT16_VECTOR, or BFLOAT16_VECTOR. This parameter is used to specify the dimensions, index type, and similarity metric.
	VectorIndexConfig *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetElementType() *string {
	return s.ElementType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetName() *string {
	return s.Name
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetPk() *bool {
	return s.Pk
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetType() *string {
	return s.Type
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetUrl() *bool {
	return s.Url
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetVectorIndexConfig() *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetComment(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Comment = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetElementType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.ElementType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetMaxCapacity(v int32) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.MaxCapacity = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Name = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetPk(v bool) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Pk = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Type = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetUrl(v bool) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Url = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetVectorIndexConfig(v *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.VectorIndexConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) Validate() error {
	if s.VectorIndexConfig != nil {
		if err := s.VectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig struct {
	// The embedding dimension.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	Dimension *int64 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The embedding model.
	//
	// This parameter is required.
	//
	// example:
	//
	// multimodal-embedding-v1
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// The index build parameters. Different parameters are required based on the indexType. For example, HNSW requires {M:30, efConstruction:360}, and IVF_FLAT requires {nlist:128}.
	//
	// example:
	//
	// {M:30, efConstruction:360}
	IndexParams map[string]interface{} `json:"IndexParams,omitempty" xml:"IndexParams,omitempty"`
	// The index type. PostgreSQL supports IVFFlat and HNSW. Milvus supports all types.
	//
	// This parameter is required.
	//
	// example:
	//
	// AUTOINDEX
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// The similarity type. Default value: COSINE. Valid values: COSINE, L2, and IP.
	//
	// This parameter is required.
	//
	// example:
	//
	// COSINE
	SimilarityType *string `json:"SimilarityType,omitempty" xml:"SimilarityType,omitempty"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetDimension() *int64 {
	return s.Dimension
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetIndexParams() map[string]interface{} {
	return s.IndexParams
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetIndexType() *string {
	return s.IndexType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetSimilarityType() *string {
	return s.SimilarityType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetDimension(v int64) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.Dimension = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetEmbeddingModel(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.EmbeddingModel = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetIndexParams(v map[string]interface{}) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexParams = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetIndexType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetSimilarityType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.SimilarityType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}
