// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDatasetResponseBody
	GetCode() *string
	SetDatasetDTO(v *GetDatasetResponseBodyDatasetDTO) *GetDatasetResponseBody
	GetDatasetDTO() *GetDatasetResponseBodyDatasetDTO
	SetHttpStatusCode(v int32) *GetDatasetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDatasetResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDatasetResponseBody
	GetSuccess() *bool
}

type GetDatasetResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dataset object.
	DatasetDTO *GetDatasetResponseBodyDatasetDTO `json:"DatasetDTO,omitempty" xml:"DatasetDTO,omitempty" type:"Struct"`
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDatasetResponseBody) GetDatasetDTO() *GetDatasetResponseBodyDatasetDTO {
	return s.DatasetDTO
}

func (s *GetDatasetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDatasetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDatasetResponseBody) SetCode(v string) *GetDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetDTO(v *GetDatasetResponseBodyDatasetDTO) *GetDatasetResponseBody {
	s.DatasetDTO = v
	return s
}

func (s *GetDatasetResponseBody) SetHttpStatusCode(v int32) *GetDatasetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDatasetResponseBody) SetMessage(v string) *GetDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *GetDatasetResponseBody) SetRequestId(v string) *GetDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetResponseBody) SetSuccess(v bool) *GetDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *GetDatasetResponseBody) Validate() error {
	if s.DatasetDTO != nil {
		if err := s.DatasetDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetResponseBodyDatasetDTO struct {
	// The content type.
	//
	// example:
	//
	// GENERAL
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The creator ID.
	//
	// example:
	//
	// 300001391
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The creator name.
	//
	// example:
	//
	// 张三
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The subject area ID.
	//
	// example:
	//
	// 78201
	DataCellId *string `json:"DataCellId,omitempty" xml:"DataCellId,omitempty"`
	// The subject area name.
	//
	// example:
	//
	// 离线数据主题域
	DataCellName *string `json:"DataCellName,omitempty" xml:"DataCellName,omitempty"`
	// The description.
	//
	// example:
	//
	// 测试数据集
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The directory (retrieved from the file service by using the fileId).
	//
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 7255018404650688
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-04-28 10:03:49
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-04-28 10:03:49
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The dataset ID (business primary key).
	//
	// example:
	//
	// 7255018404425088
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the development owner.
	//
	// example:
	//
	// 300001391
	LockOwner *string `json:"LockOwner,omitempty" xml:"LockOwner,omitempty"`
	// The display name of the development owner on the interface.
	//
	// example:
	//
	// 张三
	LockOwnerName *string `json:"LockOwnerName,omitempty" xml:"LockOwnerName,omitempty"`
	// The metastore type.
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
	// The list of owners.
	OwnerList []*GetDatasetResponseBodyDatasetDTOOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// The ID of the project to which the dataset belongs.
	//
	// example:
	//
	// 7255013756724992
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the project to which the dataset belongs.
	//
	// example:
	//
	// ds_tm
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The dataset scenarios. Valid values:
	//
	// - OFFLINE: offline (default).
	//
	// - REALTIME: real-time.
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
	// The tenant ID.
	//
	// example:
	//
	// 300001413
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The dataset type.
	//
	// example:
	//
	// HYBRID
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The list of versions.
	VersionList []*GetDatasetResponseBodyDatasetDTOVersionList `json:"VersionList,omitempty" xml:"VersionList,omitempty" type:"Repeated"`
}

func (s GetDatasetResponseBodyDatasetDTO) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTO) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTO) GetContentType() *string {
	return s.ContentType
}

func (s *GetDatasetResponseBodyDatasetDTO) GetCreator() *string {
	return s.Creator
}

func (s *GetDatasetResponseBodyDatasetDTO) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetDatasetResponseBodyDatasetDTO) GetDataCellId() *string {
	return s.DataCellId
}

func (s *GetDatasetResponseBodyDatasetDTO) GetDataCellName() *string {
	return s.DataCellName
}

func (s *GetDatasetResponseBodyDatasetDTO) GetDescription() *string {
	return s.Description
}

func (s *GetDatasetResponseBodyDatasetDTO) GetDirectory() *string {
	return s.Directory
}

func (s *GetDatasetResponseBodyDatasetDTO) GetFileId() *int64 {
	return s.FileId
}

func (s *GetDatasetResponseBodyDatasetDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDatasetResponseBodyDatasetDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDatasetResponseBodyDatasetDTO) GetId() *int64 {
	return s.Id
}

func (s *GetDatasetResponseBodyDatasetDTO) GetLockOwner() *string {
	return s.LockOwner
}

func (s *GetDatasetResponseBodyDatasetDTO) GetLockOwnerName() *string {
	return s.LockOwnerName
}

func (s *GetDatasetResponseBodyDatasetDTO) GetMetadataStorageType() *string {
	return s.MetadataStorageType
}

func (s *GetDatasetResponseBodyDatasetDTO) GetName() *string {
	return s.Name
}

func (s *GetDatasetResponseBodyDatasetDTO) GetOwnerList() []*GetDatasetResponseBodyDatasetDTOOwnerList {
	return s.OwnerList
}

func (s *GetDatasetResponseBodyDatasetDTO) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDatasetResponseBodyDatasetDTO) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetDatasetResponseBodyDatasetDTO) GetScenario() *string {
	return s.Scenario
}

func (s *GetDatasetResponseBodyDatasetDTO) GetStorageType() *string {
	return s.StorageType
}

func (s *GetDatasetResponseBodyDatasetDTO) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetDatasetResponseBodyDatasetDTO) GetType() *string {
	return s.Type
}

func (s *GetDatasetResponseBodyDatasetDTO) GetVersionList() []*GetDatasetResponseBodyDatasetDTOVersionList {
	return s.VersionList
}

func (s *GetDatasetResponseBodyDatasetDTO) SetContentType(v string) *GetDatasetResponseBodyDatasetDTO {
	s.ContentType = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetCreator(v string) *GetDatasetResponseBodyDatasetDTO {
	s.Creator = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetCreatorName(v string) *GetDatasetResponseBodyDatasetDTO {
	s.CreatorName = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetDataCellId(v string) *GetDatasetResponseBodyDatasetDTO {
	s.DataCellId = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetDataCellName(v string) *GetDatasetResponseBodyDatasetDTO {
	s.DataCellName = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetDescription(v string) *GetDatasetResponseBodyDatasetDTO {
	s.Description = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetDirectory(v string) *GetDatasetResponseBodyDatasetDTO {
	s.Directory = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetFileId(v int64) *GetDatasetResponseBodyDatasetDTO {
	s.FileId = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetGmtCreate(v string) *GetDatasetResponseBodyDatasetDTO {
	s.GmtCreate = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetGmtModified(v string) *GetDatasetResponseBodyDatasetDTO {
	s.GmtModified = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetId(v int64) *GetDatasetResponseBodyDatasetDTO {
	s.Id = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetLockOwner(v string) *GetDatasetResponseBodyDatasetDTO {
	s.LockOwner = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetLockOwnerName(v string) *GetDatasetResponseBodyDatasetDTO {
	s.LockOwnerName = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetMetadataStorageType(v string) *GetDatasetResponseBodyDatasetDTO {
	s.MetadataStorageType = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetName(v string) *GetDatasetResponseBodyDatasetDTO {
	s.Name = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetOwnerList(v []*GetDatasetResponseBodyDatasetDTOOwnerList) *GetDatasetResponseBodyDatasetDTO {
	s.OwnerList = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetProjectId(v int64) *GetDatasetResponseBodyDatasetDTO {
	s.ProjectId = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetProjectName(v string) *GetDatasetResponseBodyDatasetDTO {
	s.ProjectName = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetScenario(v string) *GetDatasetResponseBodyDatasetDTO {
	s.Scenario = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetStorageType(v string) *GetDatasetResponseBodyDatasetDTO {
	s.StorageType = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetTenantId(v int64) *GetDatasetResponseBodyDatasetDTO {
	s.TenantId = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetType(v string) *GetDatasetResponseBodyDatasetDTO {
	s.Type = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) SetVersionList(v []*GetDatasetResponseBodyDatasetDTOVersionList) *GetDatasetResponseBodyDatasetDTO {
	s.VersionList = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTO) Validate() error {
	if s.OwnerList != nil {
		for _, item := range s.OwnerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VersionList != nil {
		for _, item := range s.VersionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDatasetResponseBodyDatasetDTOOwnerList struct {
	// The user ID.
	//
	// example:
	//
	// 300001391
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// 张三
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetDatasetResponseBodyDatasetDTOOwnerList) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOOwnerList) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOOwnerList) GetUserId() *string {
	return s.UserId
}

func (s *GetDatasetResponseBodyDatasetDTOOwnerList) GetUserName() *string {
	return s.UserName
}

func (s *GetDatasetResponseBodyDatasetDTOOwnerList) SetUserId(v string) *GetDatasetResponseBodyDatasetDTOOwnerList {
	s.UserId = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOOwnerList) SetUserName(v string) *GetDatasetResponseBodyDatasetDTOOwnerList {
	s.UserName = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOOwnerList) Validate() error {
	return dara.Validate(s)
}

type GetDatasetResponseBodyDatasetDTOVersionList struct {
	// The creator ID.
	//
	// example:
	//
	// ***
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The dataset version configuration.
	DataVersionConfig *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig `json:"DataVersionConfig,omitempty" xml:"DataVersionConfig,omitempty" type:"Struct"`
	// The dataset ID.
	//
	// example:
	//
	// 7280832407583104
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-11-13T02:11:56Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1749450490000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The version ID.
	//
	// example:
	//
	// 7280840713415040
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The version number.
	//
	// example:
	//
	// V4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetDatasetResponseBodyDatasetDTOVersionList) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOVersionList) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) GetCreator() *string {
	return s.Creator
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) GetDataVersionConfig() *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig {
	return s.DataVersionConfig
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) GetId() *int64 {
	return s.Id
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) GetVersion() *string {
	return s.Version
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) SetCreator(v string) *GetDatasetResponseBodyDatasetDTOVersionList {
	s.Creator = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) SetDataVersionConfig(v *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) *GetDatasetResponseBodyDatasetDTOVersionList {
	s.DataVersionConfig = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) SetDatasetId(v int64) *GetDatasetResponseBodyDatasetDTOVersionList {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) SetGmtCreate(v string) *GetDatasetResponseBodyDatasetDTOVersionList {
	s.GmtCreate = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) SetGmtModified(v string) *GetDatasetResponseBodyDatasetDTOVersionList {
	s.GmtModified = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) SetId(v int64) *GetDatasetResponseBodyDatasetDTOVersionList {
	s.Id = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) SetVersion(v string) *GetDatasetResponseBodyDatasetDTOVersionList {
	s.Version = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionList) Validate() error {
	if s.DataVersionConfig != nil {
		if err := s.DataVersionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig struct {
	// The file storage configuration.
	FileStorageConfig *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig `json:"FileStorageConfig,omitempty" xml:"FileStorageConfig,omitempty" type:"Struct"`
	// The metastore configuration.
	MetadataStorageConfig *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig `json:"MetadataStorageConfig,omitempty" xml:"MetadataStorageConfig,omitempty" type:"Struct"`
	// The real-time meta table configuration. This parameter takes effect only when metadataStorageType is set to REALTIME_META_TABLE.
	RealtimeMetaTableConfig *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig `json:"RealtimeMetaTableConfig,omitempty" xml:"RealtimeMetaTableConfig,omitempty" type:"Struct"`
	// The version description.
	//
	// example:
	//
	// 测试数据集版本
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) GetFileStorageConfig() *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig {
	return s.FileStorageConfig
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) GetMetadataStorageConfig() *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig {
	return s.MetadataStorageConfig
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) GetRealtimeMetaTableConfig() *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig {
	return s.RealtimeMetaTableConfig
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) SetFileStorageConfig(v *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig {
	s.FileStorageConfig = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) SetMetadataStorageConfig(v *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig {
	s.MetadataStorageConfig = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) SetRealtimeMetaTableConfig(v *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig {
	s.RealtimeMetaTableConfig = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) SetVersionDescription(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig {
	s.VersionDescription = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfig) Validate() error {
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

type GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig struct {
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

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) GetDevPath() *string {
	return s.DevPath
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) GetMountPath() *string {
	return s.MountPath
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) GetProdPath() *string {
	return s.ProdPath
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) SetDataSourceId(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig {
	s.DataSourceId = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) SetDataSourceName(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig {
	s.DataSourceName = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) SetDevPath(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig {
	s.DevPath = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) SetMountPath(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig {
	s.MountPath = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) SetProdPath(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig {
	s.ProdPath = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigFileStorageConfig) Validate() error {
	return dara.Validate(s)
}

type GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig struct {
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
	// The storage destination (new table or existing table).
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
	// The table schema configuration.
	TableSchema *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) GetDevSchema() *string {
	return s.DevSchema
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) GetMetadataStorageMode() *string {
	return s.MetadataStorageMode
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) GetMetadataStorageType() *string {
	return s.MetadataStorageType
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) GetProdSchema() *string {
	return s.ProdSchema
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) GetTableName() *string {
	return s.TableName
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) GetTableSchema() *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchema {
	return s.TableSchema
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) SetDataSourceId(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig {
	s.DataSourceId = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) SetDataSourceName(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig {
	s.DataSourceName = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) SetDevSchema(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig {
	s.DevSchema = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) SetMetadataStorageMode(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig {
	s.MetadataStorageMode = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) SetMetadataStorageType(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig {
	s.MetadataStorageType = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) SetProdSchema(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig {
	s.ProdSchema = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) SetTableName(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig {
	s.TableName = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) SetTableSchema(v *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchema) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig {
	s.TableSchema = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfig) Validate() error {
	if s.TableSchema != nil {
		if err := s.TableSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchema struct {
	// The list of columns.
	Columns []*GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchema) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchema) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchema) GetColumns() []*GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	return s.Columns
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchema) SetColumns(v []*GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchema {
	s.Columns = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchema) Validate() error {
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

type GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns struct {
	// The column comment.
	//
	// example:
	//
	// 主键
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The array element subtype. This parameter is valid only when type is set to ARRAY.
	//
	// example:
	//
	// INT64
	ElementType *string `json:"ElementType,omitempty" xml:"ElementType,omitempty"`
	// The maximum array capacity. This parameter is valid only when type is set to ARRAY. Default value: 4096.
	//
	// example:
	//
	// 35
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The column name.
	//
	// This parameter is required.
	//
	// example:
	//
	// id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the column is a primary key.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Pk *bool `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// The column type.
	//
	// This parameter is required.
	//
	// example:
	//
	// int8
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Indicates whether the column is a URL.
	//
	// example:
	//
	// false
	Url *bool `json:"Url,omitempty" xml:"Url,omitempty"`
	// The vector index configuration.
	VectorIndexConfig *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetElementType() *string {
	return s.ElementType
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetName() *string {
	return s.Name
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetPk() *bool {
	return s.Pk
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetType() *string {
	return s.Type
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetUrl() *bool {
	return s.Url
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetVectorIndexConfig() *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetComment(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Comment = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetElementType(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.ElementType = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetMaxCapacity(v int32) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.MaxCapacity = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetName(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Name = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetPk(v bool) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Pk = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetType(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Type = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetUrl(v bool) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Url = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetVectorIndexConfig(v *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.VectorIndexConfig = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) Validate() error {
	if s.VectorIndexConfig != nil {
		if err := s.VectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig struct {
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
	// The index build parameters. Different parameters are required depending on the index type. For example, HNSW requires {M:30, efConstruction:360}, and IVF_FLAT requires {nlist:128}.
	//
	// example:
	//
	// {M:30, efConstruction:360}
	IndexParams map[string]interface{} `json:"IndexParams,omitempty" xml:"IndexParams,omitempty"`
	// The index type. PG supports IVFFlat and HNSW. Milvus supports all types.
	//
	// This parameter is required.
	//
	// example:
	//
	// AUTOINDEX
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// The similarity type. Default value: COSINE. COSINE corresponds to vector_cosine_ops, L2 corresponds to vector_l2_ops, and IP corresponds to vector_ip_ops.
	//
	// This parameter is required.
	//
	// example:
	//
	// COSINE
	SimilarityType *string `json:"SimilarityType,omitempty" xml:"SimilarityType,omitempty"`
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetDimension() *int64 {
	return s.Dimension
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetIndexParams() map[string]interface{} {
	return s.IndexParams
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetIndexType() *string {
	return s.IndexType
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetSimilarityType() *string {
	return s.SimilarityType
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetDimension(v int64) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.Dimension = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetEmbeddingModel(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.EmbeddingModel = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetIndexParams(v map[string]interface{}) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexParams = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetIndexType(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexType = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetSimilarityType(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.SimilarityType = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}

type GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig struct {
	// The data source type of the meta table. Only KAFKA is supported in the current release.
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
	// The project ID to which the meta table belongs. Cross-project references are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7255013756724992
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The table schema configuration (reuses MetadataStorageConfigDTO.TableSchemaDTO).
	TableSchema *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) GetDatasourceType() *string {
	return s.DatasourceType
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) GetMetaTableName() *string {
	return s.MetaTableName
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) GetTableSchema() *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema {
	return s.TableSchema
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) SetDatasourceType(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig {
	s.DatasourceType = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) SetMetaTableName(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig {
	s.MetaTableName = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) SetProjectId(v int64) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig {
	s.ProjectId = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) SetTableSchema(v *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig {
	s.TableSchema = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfig) Validate() error {
	if s.TableSchema != nil {
		if err := s.TableSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema struct {
	// The list of columns.
	Columns []*GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) GetColumns() []*GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	return s.Columns
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) SetColumns(v []*GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema {
	s.Columns = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) Validate() error {
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

type GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns struct {
	// The column comment.
	//
	// example:
	//
	// 录入时间
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The array element subtype.
	//
	// example:
	//
	// INT64
	ElementType *string `json:"ElementType,omitempty" xml:"ElementType,omitempty"`
	// The maximum array capacity.
	//
	// example:
	//
	// 5
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The column name.
	//
	// This parameter is required.
	//
	// example:
	//
	// happen_time
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the column is a primary key.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Pk *bool `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// The column type.
	//
	// This parameter is required.
	//
	// example:
	//
	// date
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Indicates whether the column is a URL.
	//
	// example:
	//
	// false
	Url *bool `json:"Url,omitempty" xml:"Url,omitempty"`
	// The vector index configuration.
	VectorIndexConfig *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetElementType() *string {
	return s.ElementType
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetName() *string {
	return s.Name
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetPk() *bool {
	return s.Pk
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetType() *string {
	return s.Type
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetUrl() *bool {
	return s.Url
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetVectorIndexConfig() *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetComment(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Comment = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetElementType(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.ElementType = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetMaxCapacity(v int32) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.MaxCapacity = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetName(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Name = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetPk(v bool) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Pk = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetType(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Type = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetUrl(v bool) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Url = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetVectorIndexConfig(v *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.VectorIndexConfig = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) Validate() error {
	if s.VectorIndexConfig != nil {
		if err := s.VectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig struct {
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
	// The index build parameters.
	IndexParams map[string]interface{} `json:"IndexParams,omitempty" xml:"IndexParams,omitempty"`
	// The index type.
	//
	// This parameter is required.
	//
	// example:
	//
	// AUTOINDEX
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// The similarity type.
	//
	// This parameter is required.
	//
	// example:
	//
	// COSINE
	SimilarityType *string `json:"SimilarityType,omitempty" xml:"SimilarityType,omitempty"`
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetDimension() *int64 {
	return s.Dimension
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetIndexParams() map[string]interface{} {
	return s.IndexParams
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetIndexType() *string {
	return s.IndexType
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetSimilarityType() *string {
	return s.SimilarityType
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetDimension(v int64) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.Dimension = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetEmbeddingModel(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.EmbeddingModel = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetIndexParams(v map[string]interface{}) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexParams = v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetIndexType(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexType = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetSimilarityType(v string) *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.SimilarityType = &v
	return s
}

func (s *GetDatasetResponseBodyDatasetDTOVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}
