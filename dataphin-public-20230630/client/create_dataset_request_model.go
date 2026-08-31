// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateDatasetRequestCreateCommand) *CreateDatasetRequest
	GetCreateCommand() *CreateDatasetRequestCreateCommand
	SetOpTenantId(v int64) *CreateDatasetRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateDatasetRequest
	GetOpUserId() *string
	SetProjectId(v string) *CreateDatasetRequest
	GetProjectId() *string
}

type CreateDatasetRequest struct {
	// The creation request.
	//
	// This parameter is required.
	CreateCommand *CreateDatasetRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
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
	// 7255013756724992
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) GetCreateCommand() *CreateDatasetRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateDatasetRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDatasetRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateDatasetRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateDatasetRequest) SetCreateCommand(v *CreateDatasetRequestCreateCommand) *CreateDatasetRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateDatasetRequest) SetOpTenantId(v int64) *CreateDatasetRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDatasetRequest) SetOpUserId(v string) *CreateDatasetRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateDatasetRequest) SetProjectId(v string) *CreateDatasetRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDatasetRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDatasetRequestCreateCommand struct {
	ApiInfo *CreateDatasetRequestCreateCommandApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Struct"`
	// The dataset content type. Valid values: GENERAL, TEXT, AUDIO, VIDEO, IMAGE, TABLE, INDEX.
	//
	// This parameter is required.
	//
	// example:
	//
	// GENERAL
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The data domain ID.
	//
	// example:
	//
	// 78201
	DataCellId *string `json:"DataCellId,omitempty" xml:"DataCellId,omitempty"`
	// The description.
	//
	// example:
	//
	// 测试数据集
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The directory (obtained from the file service by using the fileId).
	//
	// This parameter is required.
	//
	// example:
	//
	// /
	DirName *string `json:"DirName,omitempty" xml:"DirName,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 7255018404650688
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The metastore type.
	//
	// example:
	//
	// POSTGRESQL
	MetadataStorageType *string `json:"MetadataStorageType,omitempty" xml:"MetadataStorageType,omitempty"`
	// The dataset name.
	//
	// This parameter is required.
	//
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of owner IDs, separated by commas.
	//
	// example:
	//
	// 300000913
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The dataset scenarios. Valid values:
	//
	// - OFFLINE: offline. This is the default value.
	//
	// - REALTIME: real-time.
	//
	// This parameter is required.
	//
	// example:
	//
	// OFFLINE
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The storage type.
	//
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The dataset type. Valid values: FILE, TABLE, HYBRID.
	//
	// This parameter is required.
	//
	// example:
	//
	// FILE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version number. If not specified, the default version V1 is used.
	//
	// example:
	//
	// V1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version configuration.
	VersionConfig *CreateDatasetRequestCreateCommandVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Struct"`
}

func (s CreateDatasetRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommand) GetApiInfo() *CreateDatasetRequestCreateCommandApiInfo {
	return s.ApiInfo
}

func (s *CreateDatasetRequestCreateCommand) GetContentType() *string {
	return s.ContentType
}

func (s *CreateDatasetRequestCreateCommand) GetDataCellId() *string {
	return s.DataCellId
}

func (s *CreateDatasetRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateDatasetRequestCreateCommand) GetDirName() *string {
	return s.DirName
}

func (s *CreateDatasetRequestCreateCommand) GetFileId() *string {
	return s.FileId
}

func (s *CreateDatasetRequestCreateCommand) GetMetadataStorageType() *string {
	return s.MetadataStorageType
}

func (s *CreateDatasetRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateDatasetRequestCreateCommand) GetOwner() *string {
	return s.Owner
}

func (s *CreateDatasetRequestCreateCommand) GetScenario() *string {
	return s.Scenario
}

func (s *CreateDatasetRequestCreateCommand) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDatasetRequestCreateCommand) GetType() *string {
	return s.Type
}

func (s *CreateDatasetRequestCreateCommand) GetVersion() *string {
	return s.Version
}

func (s *CreateDatasetRequestCreateCommand) GetVersionConfig() *CreateDatasetRequestCreateCommandVersionConfig {
	return s.VersionConfig
}

func (s *CreateDatasetRequestCreateCommand) SetApiInfo(v *CreateDatasetRequestCreateCommandApiInfo) *CreateDatasetRequestCreateCommand {
	s.ApiInfo = v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetContentType(v string) *CreateDatasetRequestCreateCommand {
	s.ContentType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetDataCellId(v string) *CreateDatasetRequestCreateCommand {
	s.DataCellId = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetDescription(v string) *CreateDatasetRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetDirName(v string) *CreateDatasetRequestCreateCommand {
	s.DirName = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetFileId(v string) *CreateDatasetRequestCreateCommand {
	s.FileId = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetMetadataStorageType(v string) *CreateDatasetRequestCreateCommand {
	s.MetadataStorageType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetName(v string) *CreateDatasetRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetOwner(v string) *CreateDatasetRequestCreateCommand {
	s.Owner = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetScenario(v string) *CreateDatasetRequestCreateCommand {
	s.Scenario = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetStorageType(v string) *CreateDatasetRequestCreateCommand {
	s.StorageType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetType(v string) *CreateDatasetRequestCreateCommand {
	s.Type = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetVersion(v string) *CreateDatasetRequestCreateCommand {
	s.Version = &v
	return s
}

func (s *CreateDatasetRequestCreateCommand) SetVersionConfig(v *CreateDatasetRequestCreateCommandVersionConfig) *CreateDatasetRequestCreateCommand {
	s.VersionConfig = v
	return s
}

func (s *CreateDatasetRequestCreateCommand) Validate() error {
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

type CreateDatasetRequestCreateCommandApiInfo struct {
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
	// 1022
	OsProject *int32 `json:"OsProject,omitempty" xml:"OsProject,omitempty"`
	// example:
	//
	// 1
	Protocol *int32 `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 1
	RequestMethod     *int32                                                       `json:"RequestMethod,omitempty" xml:"RequestMethod,omitempty"`
	RequestParamList  []*CreateDatasetRequestCreateCommandApiInfoRequestParamList  `json:"RequestParamList,omitempty" xml:"RequestParamList,omitempty" type:"Repeated"`
	ResponseParamList []*CreateDatasetRequestCreateCommandApiInfoResponseParamList `json:"ResponseParamList,omitempty" xml:"ResponseParamList,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateDatasetRequestCreateCommandApiInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandApiInfo) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandApiInfo) GetExecTimeout() *int32 {
	return s.ExecTimeout
}

func (s *CreateDatasetRequestCreateCommandApiInfo) GetExecuteMode() *int32 {
	return s.ExecuteMode
}

func (s *CreateDatasetRequestCreateCommandApiInfo) GetOsApiGroup() *int32 {
	return s.OsApiGroup
}

func (s *CreateDatasetRequestCreateCommandApiInfo) GetOsProject() *int32 {
	return s.OsProject
}

func (s *CreateDatasetRequestCreateCommandApiInfo) GetProtocol() *int32 {
	return s.Protocol
}

func (s *CreateDatasetRequestCreateCommandApiInfo) GetRequestMethod() *int32 {
	return s.RequestMethod
}

func (s *CreateDatasetRequestCreateCommandApiInfo) GetRequestParamList() []*CreateDatasetRequestCreateCommandApiInfoRequestParamList {
	return s.RequestParamList
}

func (s *CreateDatasetRequestCreateCommandApiInfo) GetResponseParamList() []*CreateDatasetRequestCreateCommandApiInfoResponseParamList {
	return s.ResponseParamList
}

func (s *CreateDatasetRequestCreateCommandApiInfo) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateDatasetRequestCreateCommandApiInfo) SetExecTimeout(v int32) *CreateDatasetRequestCreateCommandApiInfo {
	s.ExecTimeout = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfo) SetExecuteMode(v int32) *CreateDatasetRequestCreateCommandApiInfo {
	s.ExecuteMode = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfo) SetOsApiGroup(v int32) *CreateDatasetRequestCreateCommandApiInfo {
	s.OsApiGroup = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfo) SetOsProject(v int32) *CreateDatasetRequestCreateCommandApiInfo {
	s.OsProject = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfo) SetProtocol(v int32) *CreateDatasetRequestCreateCommandApiInfo {
	s.Protocol = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfo) SetRequestMethod(v int32) *CreateDatasetRequestCreateCommandApiInfo {
	s.RequestMethod = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfo) SetRequestParamList(v []*CreateDatasetRequestCreateCommandApiInfoRequestParamList) *CreateDatasetRequestCreateCommandApiInfo {
	s.RequestParamList = v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfo) SetResponseParamList(v []*CreateDatasetRequestCreateCommandApiInfoResponseParamList) *CreateDatasetRequestCreateCommandApiInfo {
	s.ResponseParamList = v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfo) SetTimeout(v int32) *CreateDatasetRequestCreateCommandApiInfo {
	s.Timeout = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfo) Validate() error {
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

type CreateDatasetRequestCreateCommandApiInfoRequestParamList struct {
	// example:
	//
	// test
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
	// test
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
}

func (s CreateDatasetRequestCreateCommandApiInfoRequestParamList) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandApiInfoRequestParamList) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) GetDescr() *string {
	return s.Descr
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) GetIsUrl() *bool {
	return s.IsUrl
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) GetMust() *bool {
	return s.Must
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) GetParamName() *string {
	return s.ParamName
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) GetParamType() *string {
	return s.ParamType
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) GetSample() *string {
	return s.Sample
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) SetDefaultValue(v string) *CreateDatasetRequestCreateCommandApiInfoRequestParamList {
	s.DefaultValue = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) SetDescr(v string) *CreateDatasetRequestCreateCommandApiInfoRequestParamList {
	s.Descr = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) SetIsUrl(v bool) *CreateDatasetRequestCreateCommandApiInfoRequestParamList {
	s.IsUrl = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) SetMust(v bool) *CreateDatasetRequestCreateCommandApiInfoRequestParamList {
	s.Must = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) SetParamName(v string) *CreateDatasetRequestCreateCommandApiInfoRequestParamList {
	s.ParamName = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) SetParamType(v string) *CreateDatasetRequestCreateCommandApiInfoRequestParamList {
	s.ParamType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) SetSample(v string) *CreateDatasetRequestCreateCommandApiInfoRequestParamList {
	s.Sample = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoRequestParamList) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestCreateCommandApiInfoResponseParamList struct {
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
	// test
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
}

func (s CreateDatasetRequestCreateCommandApiInfoResponseParamList) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandApiInfoResponseParamList) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandApiInfoResponseParamList) GetDescr() *string {
	return s.Descr
}

func (s *CreateDatasetRequestCreateCommandApiInfoResponseParamList) GetIsUrl() *bool {
	return s.IsUrl
}

func (s *CreateDatasetRequestCreateCommandApiInfoResponseParamList) GetParamName() *string {
	return s.ParamName
}

func (s *CreateDatasetRequestCreateCommandApiInfoResponseParamList) GetParamType() *string {
	return s.ParamType
}

func (s *CreateDatasetRequestCreateCommandApiInfoResponseParamList) GetSample() *string {
	return s.Sample
}

func (s *CreateDatasetRequestCreateCommandApiInfoResponseParamList) SetDescr(v string) *CreateDatasetRequestCreateCommandApiInfoResponseParamList {
	s.Descr = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoResponseParamList) SetIsUrl(v bool) *CreateDatasetRequestCreateCommandApiInfoResponseParamList {
	s.IsUrl = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoResponseParamList) SetParamName(v string) *CreateDatasetRequestCreateCommandApiInfoResponseParamList {
	s.ParamName = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoResponseParamList) SetParamType(v string) *CreateDatasetRequestCreateCommandApiInfoResponseParamList {
	s.ParamType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoResponseParamList) SetSample(v string) *CreateDatasetRequestCreateCommandApiInfoResponseParamList {
	s.Sample = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandApiInfoResponseParamList) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestCreateCommandVersionConfig struct {
	// The file storage configuration.
	FileStorageConfig *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig `json:"FileStorageConfig,omitempty" xml:"FileStorageConfig,omitempty" type:"Struct"`
	// The metastore configuration.
	MetadataStorageConfig *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig `json:"MetadataStorageConfig,omitempty" xml:"MetadataStorageConfig,omitempty" type:"Struct"`
	// The real-time meta table configuration. Takes effect when metadataStorageType is set to STREAM_TABLE.
	RealtimeMetaTableConfig *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig `json:"RealtimeMetaTableConfig,omitempty" xml:"RealtimeMetaTableConfig,omitempty" type:"Struct"`
	// The version description.
	//
	// example:
	//
	// 测试数据集版本
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
}

func (s CreateDatasetRequestCreateCommandVersionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandVersionConfig) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandVersionConfig) GetFileStorageConfig() *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig {
	return s.FileStorageConfig
}

func (s *CreateDatasetRequestCreateCommandVersionConfig) GetMetadataStorageConfig() *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig {
	return s.MetadataStorageConfig
}

func (s *CreateDatasetRequestCreateCommandVersionConfig) GetRealtimeMetaTableConfig() *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig {
	return s.RealtimeMetaTableConfig
}

func (s *CreateDatasetRequestCreateCommandVersionConfig) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *CreateDatasetRequestCreateCommandVersionConfig) SetFileStorageConfig(v *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) *CreateDatasetRequestCreateCommandVersionConfig {
	s.FileStorageConfig = v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfig) SetMetadataStorageConfig(v *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) *CreateDatasetRequestCreateCommandVersionConfig {
	s.MetadataStorageConfig = v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfig) SetRealtimeMetaTableConfig(v *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) *CreateDatasetRequestCreateCommandVersionConfig {
	s.RealtimeMetaTableConfig = v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfig) SetVersionDescription(v string) *CreateDatasetRequestCreateCommandVersionConfig {
	s.VersionDescription = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfig) Validate() error {
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

type CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig struct {
	// The data source ID.
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
	// 测试数据源
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

func (s CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) GetDevPath() *string {
	return s.DevPath
}

func (s *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) GetProdPath() *string {
	return s.ProdPath
}

func (s *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) SetDataSourceId(v string) *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig {
	s.DataSourceId = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) SetDataSourceName(v string) *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig {
	s.DataSourceName = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) SetDevPath(v string) *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig {
	s.DevPath = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) SetMountPath(v string) *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig {
	s.MountPath = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) SetProdPath(v string) *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig {
	s.ProdPath = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigFileStorageConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig struct {
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7429133693081710272
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The data source name.
	//
	// example:
	//
	// 测试数据源
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The development database/schema.
	//
	// example:
	//
	// HTML正文提取/test423/
	DevSchema *string `json:"DevSchema,omitempty" xml:"DevSchema,omitempty"`
	// The metadata storage mode. Valid values:
	//
	// - CREATE: create a new table.
	//
	// - EXISTING: use an existing table.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREATE
	MetadataStorageMode *string `json:"MetadataStorageMode,omitempty" xml:"MetadataStorageMode,omitempty"`
	// The metastore type.
	//
	// example:
	//
	// MILVUS
	MetadataStorageType *string `json:"MetadataStorageType,omitempty" xml:"MetadataStorageType,omitempty"`
	// The production database/schema.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTML正文提取/test423/
	ProdSchema *string `json:"ProdSchema,omitempty" xml:"ProdSchema,omitempty"`
	// The table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// s_crm_all_plt_jala_shop
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The table schema.
	TableSchema *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) GetDevSchema() *string {
	return s.DevSchema
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) GetMetadataStorageMode() *string {
	return s.MetadataStorageMode
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) GetMetadataStorageType() *string {
	return s.MetadataStorageType
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) GetProdSchema() *string {
	return s.ProdSchema
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) GetTableName() *string {
	return s.TableName
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) GetTableSchema() *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchema {
	return s.TableSchema
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) SetDataSourceId(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig {
	s.DataSourceId = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) SetDataSourceName(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig {
	s.DataSourceName = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) SetDevSchema(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig {
	s.DevSchema = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) SetMetadataStorageMode(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig {
	s.MetadataStorageMode = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) SetMetadataStorageType(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig {
	s.MetadataStorageType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) SetProdSchema(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig {
	s.ProdSchema = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) SetTableName(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig {
	s.TableName = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) SetTableSchema(v *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchema) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig {
	s.TableSchema = v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfig) Validate() error {
	if s.TableSchema != nil {
		if err := s.TableSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchema struct {
	// The column list.
	Columns []*CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchema) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchema) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchema) GetColumns() []*CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	return s.Columns
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchema) SetColumns(v []*CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchema {
	s.Columns = v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchema) Validate() error {
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

type CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns struct {
	// The field comment.
	//
	// example:
	//
	// primary key
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The array element subtype. Valid only when type is set to ARRAY.
	//
	// example:
	//
	// INT64
	ElementType *string `json:"ElementType,omitempty" xml:"ElementType,omitempty"`
	// The maximum capacity of the array. Valid only when type is set to ARRAY. Default value: 4096.
	//
	// example:
	//
	// 35
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The field name.
	//
	// This parameter is required.
	//
	// example:
	//
	// id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the field is a primary key.
	//
	// example:
	//
	// true
	Pk *bool `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// The field type.
	//
	// This parameter is required.
	//
	// example:
	//
	// int8
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Specifies whether the field is a URL.
	//
	// example:
	//
	// false
	Url *bool `json:"Url,omitempty" xml:"Url,omitempty"`
	// The vector index configuration.
	VectorIndexConfig *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
}

func (s CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetElementType() *string {
	return s.ElementType
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetName() *string {
	return s.Name
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetPk() *bool {
	return s.Pk
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetType() *string {
	return s.Type
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetUrl() *bool {
	return s.Url
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetVectorIndexConfig() *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetComment(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Comment = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetElementType(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.ElementType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetMaxCapacity(v int32) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.MaxCapacity = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetName(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Name = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetPk(v bool) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Pk = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetType(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Type = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetUrl(v bool) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Url = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetVectorIndexConfig(v *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.VectorIndexConfig = v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) Validate() error {
	if s.VectorIndexConfig != nil {
		if err := s.VectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig struct {
	// The vector dimensions.
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
	// The index build parameters. Different parameters are required based on the indexType. For example, HNSW requires {M:30, efConstruction:360} and IVF_FLAT requires {nlist:128}.
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
	// The similarity type. Default value: COSINE. Valid values: COSINE, L2, IP.
	//
	// This parameter is required.
	//
	// example:
	//
	// COSINE
	SimilarityType *string `json:"SimilarityType,omitempty" xml:"SimilarityType,omitempty"`
}

func (s CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetDimension() *int64 {
	return s.Dimension
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetIndexParams() map[string]interface{} {
	return s.IndexParams
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetIndexType() *string {
	return s.IndexType
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetSimilarityType() *string {
	return s.SimilarityType
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetDimension(v int64) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.Dimension = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetEmbeddingModel(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.EmbeddingModel = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetIndexParams(v map[string]interface{}) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexParams = v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetIndexType(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetSimilarityType(v string) *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.SimilarityType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig struct {
	// The meta table data source type (only KAFKA is supported in this version).
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
	// 测试元表
	MetaTableName *string `json:"MetaTableName,omitempty" xml:"MetaTableName,omitempty"`
	// The project ID of the meta table (cross-project supported).
	//
	// This parameter is required.
	//
	// example:
	//
	// 7256391656294144
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The table schema.
	TableSchema *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) GetDatasourceType() *string {
	return s.DatasourceType
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) GetMetaTableName() *string {
	return s.MetaTableName
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) GetTableSchema() *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchema {
	return s.TableSchema
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) SetDatasourceType(v string) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig {
	s.DatasourceType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) SetMetaTableName(v string) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig {
	s.MetaTableName = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) SetProjectId(v int64) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig {
	s.ProjectId = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) SetTableSchema(v *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchema) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig {
	s.TableSchema = v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfig) Validate() error {
	if s.TableSchema != nil {
		if err := s.TableSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchema struct {
	// The column list.
	Columns []*CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchema) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchema) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchema) GetColumns() []*CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	return s.Columns
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchema) SetColumns(v []*CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchema {
	s.Columns = v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchema) Validate() error {
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

type CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns struct {
	// The field comment.
	//
	// example:
	//
	// happen time
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The array element subtype. Valid only when type is set to ARRAY.
	//
	// example:
	//
	// INT64
	ElementType *string `json:"ElementType,omitempty" xml:"ElementType,omitempty"`
	// The maximum capacity of the array. Valid only when type is set to ARRAY. Default value: 4096.
	//
	// example:
	//
	// 35
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The field name.
	//
	// This parameter is required.
	//
	// example:
	//
	// happen_time
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the field is a primary key.
	//
	// example:
	//
	// false
	Pk *bool `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// The field type.
	//
	// This parameter is required.
	//
	// example:
	//
	// date
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Specifies whether the field is a URL.
	//
	// example:
	//
	// false
	Url *bool `json:"Url,omitempty" xml:"Url,omitempty"`
	// The vector index configuration. Configure this when the field type is FLOAT_VECTOR/FLOAT16_VECTOR/BFLOAT16_VECTOR to set the dimension, index type, and similarity.
	VectorIndexConfig *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
}

func (s CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetElementType() *string {
	return s.ElementType
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetName() *string {
	return s.Name
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetPk() *bool {
	return s.Pk
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetType() *string {
	return s.Type
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetUrl() *bool {
	return s.Url
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetVectorIndexConfig() *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetComment(v string) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Comment = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetElementType(v string) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.ElementType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetMaxCapacity(v int32) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.MaxCapacity = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetName(v string) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Name = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetPk(v bool) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Pk = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetType(v string) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Type = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetUrl(v bool) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Url = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetVectorIndexConfig(v *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.VectorIndexConfig = v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) Validate() error {
	if s.VectorIndexConfig != nil {
		if err := s.VectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig struct {
	// The vector dimensions.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance:MySQL.monitor
	Dimension *int64 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The embedding model.
	//
	// This parameter is required.
	//
	// example:
	//
	// MultiModal-Embedding
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// The index build parameters. Different parameters are required based on the indexType. For example, HNSW requires {M:30, efConstruction:360} and IVF_FLAT requires {nlist:128}.
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
	// The similarity type. Default value: COSINE. Valid values: COSINE, L2, IP.
	//
	// This parameter is required.
	//
	// example:
	//
	// COSINE
	SimilarityType *string `json:"SimilarityType,omitempty" xml:"SimilarityType,omitempty"`
}

func (s CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetDimension() *int64 {
	return s.Dimension
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetIndexParams() map[string]interface{} {
	return s.IndexParams
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetIndexType() *string {
	return s.IndexType
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetSimilarityType() *string {
	return s.SimilarityType
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetDimension(v int64) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.Dimension = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetEmbeddingModel(v string) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.EmbeddingModel = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetIndexParams(v map[string]interface{}) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexParams = v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetIndexType(v string) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetSimilarityType(v string) *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.SimilarityType = &v
	return s
}

func (s *CreateDatasetRequestCreateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}
