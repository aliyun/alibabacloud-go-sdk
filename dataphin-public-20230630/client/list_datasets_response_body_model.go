// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDatasetsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListDatasetsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDatasetsResponseBody
	GetMessage() *string
	SetPageResult(v *ListDatasetsResponseBodyPageResult) *ListDatasetsResponseBody
	GetPageResult() *ListDatasetsResponseBodyPageResult
	SetRequestId(v string) *ListDatasetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDatasetsResponseBody
	GetSuccess() *bool
}

type ListDatasetsResponseBody struct {
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
	// The paged result.
	PageResult *ListDatasetsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
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

func (s ListDatasetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDatasetsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDatasetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDatasetsResponseBody) GetPageResult() *ListDatasetsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListDatasetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDatasetsResponseBody) SetCode(v string) *ListDatasetsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDatasetsResponseBody) SetHttpStatusCode(v int32) *ListDatasetsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDatasetsResponseBody) SetMessage(v string) *ListDatasetsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDatasetsResponseBody) SetPageResult(v *ListDatasetsResponseBodyPageResult) *ListDatasetsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListDatasetsResponseBody) SetRequestId(v string) *ListDatasetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetsResponseBody) SetSuccess(v bool) *ListDatasetsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDatasetsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDatasetsResponseBodyPageResult struct {
	// The total count.
	//
	// example:
	//
	// 12
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The object.
	ResultData []*ListDatasetsResponseBodyPageResultResultData `json:"ResultData,omitempty" xml:"ResultData,omitempty" type:"Repeated"`
}

func (s ListDatasetsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResult) GetCount() *int32 {
	return s.Count
}

func (s *ListDatasetsResponseBodyPageResult) GetResultData() []*ListDatasetsResponseBodyPageResultResultData {
	return s.ResultData
}

func (s *ListDatasetsResponseBodyPageResult) SetCount(v int32) *ListDatasetsResponseBodyPageResult {
	s.Count = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResult) SetResultData(v []*ListDatasetsResponseBodyPageResultResultData) *ListDatasetsResponseBodyPageResult {
	s.ResultData = v
	return s
}

func (s *ListDatasetsResponseBodyPageResult) Validate() error {
	if s.ResultData != nil {
		for _, item := range s.ResultData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDatasetsResponseBodyPageResultResultData struct {
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
	// The data domain ID.
	//
	// example:
	//
	// 74280
	DataCellId *string `json:"DataCellId,omitempty" xml:"DataCellId,omitempty"`
	// **The data domain name.**
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
	// The directory (retrieved from the file service by fileId).
	//
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 7285340579984832
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-05-18 17:07:54.237771
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-05-18 17:07:54.237771
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The dataset ID (business primary key).
	//
	// example:
	//
	// 7128268454335680
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The development owner ID.
	//
	// example:
	//
	// 300001391
	LockOwner *string `json:"LockOwner,omitempty" xml:"LockOwner,omitempty"`
	// The name of the development owner (interface Displayed Fields).
	//
	// example:
	//
	// 张三
	LockOwnerName *string `json:"LockOwnerName,omitempty" xml:"LockOwnerName,omitempty"`
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
	// The owner list.
	OwnerList []*ListDatasetsResponseBodyPageResultResultDataOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// The project ID.
	//
	// example:
	//
	// 7290731813137728
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name.
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
	// **The storage type.**
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
	// **The dataset type.**
	//
	// example:
	//
	// HYBRID
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version list.
	VersionList []*ListDatasetsResponseBodyPageResultResultDataVersionList `json:"VersionList,omitempty" xml:"VersionList,omitempty" type:"Repeated"`
}

func (s ListDatasetsResponseBodyPageResultResultData) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultData) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetContentType() *string {
	return s.ContentType
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetCreator() *string {
	return s.Creator
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetDataCellId() *string {
	return s.DataCellId
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetDataCellName() *string {
	return s.DataCellName
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetDescription() *string {
	return s.Description
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetDirectory() *string {
	return s.Directory
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetFileId() *int64 {
	return s.FileId
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetId() *int64 {
	return s.Id
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetLockOwner() *string {
	return s.LockOwner
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetLockOwnerName() *string {
	return s.LockOwnerName
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetMetadataStorageType() *string {
	return s.MetadataStorageType
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetName() *string {
	return s.Name
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetOwnerList() []*ListDatasetsResponseBodyPageResultResultDataOwnerList {
	return s.OwnerList
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetScenario() *string {
	return s.Scenario
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetStorageType() *string {
	return s.StorageType
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetType() *string {
	return s.Type
}

func (s *ListDatasetsResponseBodyPageResultResultData) GetVersionList() []*ListDatasetsResponseBodyPageResultResultDataVersionList {
	return s.VersionList
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetContentType(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.ContentType = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetCreator(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.Creator = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetCreatorName(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.CreatorName = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetDataCellId(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.DataCellId = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetDataCellName(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.DataCellName = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetDescription(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.Description = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetDirectory(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.Directory = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetFileId(v int64) *ListDatasetsResponseBodyPageResultResultData {
	s.FileId = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetGmtCreate(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetGmtModified(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.GmtModified = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetId(v int64) *ListDatasetsResponseBodyPageResultResultData {
	s.Id = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetLockOwner(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.LockOwner = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetLockOwnerName(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.LockOwnerName = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetMetadataStorageType(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.MetadataStorageType = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetName(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.Name = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetOwnerList(v []*ListDatasetsResponseBodyPageResultResultDataOwnerList) *ListDatasetsResponseBodyPageResultResultData {
	s.OwnerList = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetProjectId(v int64) *ListDatasetsResponseBodyPageResultResultData {
	s.ProjectId = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetProjectName(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.ProjectName = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetScenario(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.Scenario = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetStorageType(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.StorageType = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetTenantId(v int64) *ListDatasetsResponseBodyPageResultResultData {
	s.TenantId = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetType(v string) *ListDatasetsResponseBodyPageResultResultData {
	s.Type = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) SetVersionList(v []*ListDatasetsResponseBodyPageResultResultDataVersionList) *ListDatasetsResponseBodyPageResultResultData {
	s.VersionList = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultData) Validate() error {
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

type ListDatasetsResponseBodyPageResultResultDataOwnerList struct {
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

func (s ListDatasetsResponseBodyPageResultResultDataOwnerList) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataOwnerList) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataOwnerList) GetUserId() *string {
	return s.UserId
}

func (s *ListDatasetsResponseBodyPageResultResultDataOwnerList) GetUserName() *string {
	return s.UserName
}

func (s *ListDatasetsResponseBodyPageResultResultDataOwnerList) SetUserId(v string) *ListDatasetsResponseBodyPageResultResultDataOwnerList {
	s.UserId = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataOwnerList) SetUserName(v string) *ListDatasetsResponseBodyPageResultResultDataOwnerList {
	s.UserName = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataOwnerList) Validate() error {
	return dara.Validate(s)
}

type ListDatasetsResponseBodyPageResultResultDataVersionList struct {
	// The creator ID.
	//
	// example:
	//
	// 300001391
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The dataset version configuration.
	DataVersionConfig *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig `json:"DataVersionConfig,omitempty" xml:"DataVersionConfig,omitempty" type:"Struct"`
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
	// 2026-04-28 10:03:49
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

func (s ListDatasetsResponseBodyPageResultResultDataVersionList) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionList) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) GetCreator() *string {
	return s.Creator
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) GetDataVersionConfig() *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig {
	return s.DataVersionConfig
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) GetId() *int64 {
	return s.Id
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) GetVersion() *string {
	return s.Version
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) SetCreator(v string) *ListDatasetsResponseBodyPageResultResultDataVersionList {
	s.Creator = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) SetDataVersionConfig(v *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) *ListDatasetsResponseBodyPageResultResultDataVersionList {
	s.DataVersionConfig = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) SetDatasetId(v int64) *ListDatasetsResponseBodyPageResultResultDataVersionList {
	s.DatasetId = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) SetGmtCreate(v string) *ListDatasetsResponseBodyPageResultResultDataVersionList {
	s.GmtCreate = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) SetGmtModified(v string) *ListDatasetsResponseBodyPageResultResultDataVersionList {
	s.GmtModified = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) SetId(v int64) *ListDatasetsResponseBodyPageResultResultDataVersionList {
	s.Id = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) SetVersion(v string) *ListDatasetsResponseBodyPageResultResultDataVersionList {
	s.Version = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionList) Validate() error {
	if s.DataVersionConfig != nil {
		if err := s.DataVersionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig struct {
	// The file storage configuration.
	FileStorageConfig *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig `json:"FileStorageConfig,omitempty" xml:"FileStorageConfig,omitempty" type:"Struct"`
	// The metastore configuration.
	MetadataStorageConfig *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig `json:"MetadataStorageConfig,omitempty" xml:"MetadataStorageConfig,omitempty" type:"Struct"`
	// The real-time meta table configuration (takes effect only when `metadataStorageType=REALTIME_META_TABLE`).
	RealtimeMetaTableConfig *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig `json:"RealtimeMetaTableConfig,omitempty" xml:"RealtimeMetaTableConfig,omitempty" type:"Struct"`
	// The version description.
	//
	// example:
	//
	// 测试数据集版本
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) GetFileStorageConfig() *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig {
	return s.FileStorageConfig
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) GetMetadataStorageConfig() *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig {
	return s.MetadataStorageConfig
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) GetRealtimeMetaTableConfig() *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig {
	return s.RealtimeMetaTableConfig
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) SetFileStorageConfig(v *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig {
	s.FileStorageConfig = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) SetMetadataStorageConfig(v *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig {
	s.MetadataStorageConfig = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) SetRealtimeMetaTableConfig(v *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig {
	s.RealtimeMetaTableConfig = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) SetVersionDescription(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig {
	s.VersionDescription = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfig) Validate() error {
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

type ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig struct {
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

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) GetDevPath() *string {
	return s.DevPath
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) GetMountPath() *string {
	return s.MountPath
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) GetProdPath() *string {
	return s.ProdPath
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) SetDataSourceId(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig {
	s.DataSourceId = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) SetDataSourceName(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig {
	s.DataSourceName = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) SetDevPath(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig {
	s.DevPath = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) SetMountPath(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig {
	s.MountPath = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) SetProdPath(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig {
	s.ProdPath = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigFileStorageConfig) Validate() error {
	return dara.Validate(s)
}

type ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig struct {
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
	TableSchema *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) GetDevSchema() *string {
	return s.DevSchema
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) GetMetadataStorageMode() *string {
	return s.MetadataStorageMode
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) GetMetadataStorageType() *string {
	return s.MetadataStorageType
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) GetProdSchema() *string {
	return s.ProdSchema
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) GetTableName() *string {
	return s.TableName
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) GetTableSchema() *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchema {
	return s.TableSchema
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) SetDataSourceId(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig {
	s.DataSourceId = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) SetDataSourceName(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig {
	s.DataSourceName = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) SetDevSchema(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig {
	s.DevSchema = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) SetMetadataStorageMode(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig {
	s.MetadataStorageMode = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) SetMetadataStorageType(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig {
	s.MetadataStorageType = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) SetProdSchema(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig {
	s.ProdSchema = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) SetTableName(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig {
	s.TableName = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) SetTableSchema(v *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchema) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig {
	s.TableSchema = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfig) Validate() error {
	if s.TableSchema != nil {
		if err := s.TableSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchema struct {
	// The field list.
	Columns []*ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchema) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchema) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchema) GetColumns() []*ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	return s.Columns
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchema) SetColumns(v []*ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchema {
	s.Columns = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchema) Validate() error {
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

type ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns struct {
	// The field description.
	//
	// example:
	//
	// 主键
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The array element subtype. This parameter takes effect only when type is set to ARRAY.
	//
	// example:
	//
	// INT64
	ElementType *string `json:"ElementType,omitempty" xml:"ElementType,omitempty"`
	// The maximum array capacity. This parameter takes effect only when type is set to ARRAY. Default value: 4096.
	//
	// example:
	//
	// 20
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The field name.
	//
	// This parameter is required.
	//
	// example:
	//
	// id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the field is a primary key.
	//
	// This parameter is required.
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
	// Indicates whether the field is a URL.
	//
	// example:
	//
	// false
	Url *bool `json:"Url,omitempty" xml:"Url,omitempty"`
	// The vector index configuration.
	VectorIndexConfig *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetElementType() *string {
	return s.ElementType
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetName() *string {
	return s.Name
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetPk() *bool {
	return s.Pk
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetType() *string {
	return s.Type
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetUrl() *bool {
	return s.Url
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) GetVectorIndexConfig() *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetComment(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Comment = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetElementType(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.ElementType = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetMaxCapacity(v int32) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.MaxCapacity = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetName(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Name = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetPk(v bool) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Pk = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetType(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Type = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetUrl(v bool) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Url = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) SetVectorIndexConfig(v *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.VectorIndexConfig = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumns) Validate() error {
	if s.VectorIndexConfig != nil {
		if err := s.VectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig struct {
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
	// The index build parameters.
	IndexParams map[string]interface{} `json:"IndexParams,omitempty" xml:"IndexParams,omitempty"`
	// The index type. PG supports IVFFlat and HNSW. Milvus supports all index types.
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

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetDimension() *int64 {
	return s.Dimension
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetIndexParams() map[string]interface{} {
	return s.IndexParams
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetIndexType() *string {
	return s.IndexType
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetSimilarityType() *string {
	return s.SimilarityType
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetDimension(v int64) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.Dimension = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetEmbeddingModel(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.EmbeddingModel = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetIndexParams(v map[string]interface{}) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexParams = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetIndexType(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexType = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetSimilarityType(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.SimilarityType = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}

type ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig struct {
	// The meta table data source type (only KAFKA is available in this release).
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
	// The project ID of the meta table (cross-project access is supported).
	//
	// This parameter is required.
	//
	// example:
	//
	// 7128268454335680
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The table schema configuration (reuses `MetadataStorageConfigDTO.TableSchemaDTO`).
	TableSchema *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) GetDatasourceType() *string {
	return s.DatasourceType
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) GetMetaTableName() *string {
	return s.MetaTableName
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) GetTableSchema() *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema {
	return s.TableSchema
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) SetDatasourceType(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig {
	s.DatasourceType = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) SetMetaTableName(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig {
	s.MetaTableName = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) SetProjectId(v int64) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig {
	s.ProjectId = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) SetTableSchema(v *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig {
	s.TableSchema = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfig) Validate() error {
	if s.TableSchema != nil {
		if err := s.TableSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema struct {
	// The field list.
	Columns []*ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) GetColumns() []*ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	return s.Columns
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) SetColumns(v []*ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema {
	s.Columns = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchema) Validate() error {
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

type ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns struct {
	// **The field description.**
	//
	// example:
	//
	// 录入时间
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// **The array element subtype. This parameter takes effect only when type is set to ARRAY.**
	//
	// example:
	//
	// INT64
	ElementType *string `json:"ElementType,omitempty" xml:"ElementType,omitempty"`
	// **The maximum array capacity. This parameter takes effect only when type is set to ARRAY. Default value: 4096.**
	//
	// example:
	//
	// 35
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
	// This parameter is required.
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
	// The vector index configuration.
	VectorIndexConfig *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetElementType() *string {
	return s.ElementType
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetName() *string {
	return s.Name
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetPk() *bool {
	return s.Pk
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetType() *string {
	return s.Type
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetUrl() *bool {
	return s.Url
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetVectorIndexConfig() *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetComment(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Comment = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetElementType(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.ElementType = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetMaxCapacity(v int32) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.MaxCapacity = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetName(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Name = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetPk(v bool) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Pk = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetType(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Type = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetUrl(v bool) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Url = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetVectorIndexConfig(v *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.VectorIndexConfig = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumns) Validate() error {
	if s.VectorIndexConfig != nil {
		if err := s.VectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig struct {
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
	// The index type. PG supports IVFFlat and HNSW. Milvus supports all index types.
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

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetDimension() *int64 {
	return s.Dimension
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetIndexParams() map[string]interface{} {
	return s.IndexParams
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetIndexType() *string {
	return s.IndexType
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetSimilarityType() *string {
	return s.SimilarityType
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetDimension(v int64) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.Dimension = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetEmbeddingModel(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.EmbeddingModel = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetIndexParams(v map[string]interface{}) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexParams = v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetIndexType(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexType = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetSimilarityType(v string) *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.SimilarityType = &v
	return s
}

func (s *ListDatasetsResponseBodyPageResultResultDataVersionListDataVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}
