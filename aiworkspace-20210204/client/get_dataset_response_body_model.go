// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetDatasetResponseBody
	GetAccessibility() *string
	SetAccessibleRoleIdList(v []*string) *GetDatasetResponseBody
	GetAccessibleRoleIdList() []*string
	SetDataSourceType(v string) *GetDatasetResponseBody
	GetDataSourceType() *string
	SetDataType(v string) *GetDatasetResponseBody
	GetDataType() *string
	SetDatasetId(v string) *GetDatasetResponseBody
	GetDatasetId() *string
	SetDescription(v string) *GetDatasetResponseBody
	GetDescription() *string
	SetEdition(v string) *GetDatasetResponseBody
	GetEdition() *string
	SetGmtCreateTime(v string) *GetDatasetResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetDatasetResponseBody
	GetGmtModifiedTime() *string
	SetImportInfo(v string) *GetDatasetResponseBody
	GetImportInfo() *string
	SetIsShared(v bool) *GetDatasetResponseBody
	GetIsShared() *bool
	SetLabels(v []*Label) *GetDatasetResponseBody
	GetLabels() []*Label
	SetLatestVersion(v *DatasetVersion) *GetDatasetResponseBody
	GetLatestVersion() *DatasetVersion
	SetMountAccess(v string) *GetDatasetResponseBody
	GetMountAccess() *string
	SetMountAccessReadWriteRoleIdList(v []*string) *GetDatasetResponseBody
	GetMountAccessReadWriteRoleIdList() []*string
	SetName(v string) *GetDatasetResponseBody
	GetName() *string
	SetOptions(v string) *GetDatasetResponseBody
	GetOptions() *string
	SetOwnerId(v string) *GetDatasetResponseBody
	GetOwnerId() *string
	SetProperty(v string) *GetDatasetResponseBody
	GetProperty() *string
	SetProvider(v string) *GetDatasetResponseBody
	GetProvider() *string
	SetProviderType(v string) *GetDatasetResponseBody
	GetProviderType() *string
	SetRequestId(v string) *GetDatasetResponseBody
	GetRequestId() *string
	SetSharedFrom(v *DatasetShareRelationship) *GetDatasetResponseBody
	GetSharedFrom() *DatasetShareRelationship
	SetSharingConfig(v *GetDatasetResponseBodySharingConfig) *GetDatasetResponseBody
	GetSharingConfig() *GetDatasetResponseBodySharingConfig
	SetSourceDatasetId(v string) *GetDatasetResponseBody
	GetSourceDatasetId() *string
	SetSourceDatasetVersion(v string) *GetDatasetResponseBody
	GetSourceDatasetVersion() *string
	SetSourceId(v string) *GetDatasetResponseBody
	GetSourceId() *string
	SetSourceType(v string) *GetDatasetResponseBody
	GetSourceType() *string
	SetTagTemplateType(v string) *GetDatasetResponseBody
	GetTagTemplateType() *string
	SetUri(v string) *GetDatasetResponseBody
	GetUri() *string
	SetUserId(v string) *GetDatasetResponseBody
	GetUserId() *string
	SetWorkspaceId(v string) *GetDatasetResponseBody
	GetWorkspaceId() *string
}

type GetDatasetResponseBody struct {
	// The workspace visibility. Valid values:
	//
	// - PRIVATE: Only the dataset owner and administrators in the workspace can access the dataset.
	//
	// - PUBLIC: All members in the workspace can access the dataset.
	//
	// - ROLE_PUBLIC: Only specified workspace roles can access the dataset. For the role list, see AccessibleRoleIdList. The dataset owner and administrators always have access under this condition.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The list of workspace role names that can access the dataset. This field takes effect when Accessibility is ROLE_PUBLIC. IDs starting with PAI are basic role IDs, and IDs starting with role- are custom role IDs.
	AccessibleRoleIdList []*string `json:"AccessibleRoleIdList,omitempty" xml:"AccessibleRoleIdList,omitempty" type:"Repeated"`
	// The data source type. Valid values:
	//
	// - OSS: Alibaba Cloud Object Storage Service (OSS).
	//
	// - NAS: Alibaba Cloud Apsara File Storage NAS (NAS).
	//
	// example:
	//
	// NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The data type of the dataset. Valid values:
	//
	// - COMMON: common.
	//
	// - PIC: image.
	//
	// - TEXT: text.
	//
	// - VIDEO: video.
	//
	// - AUDIO: audio.
	//
	// example:
	//
	// COMMON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// d-rbvg5wz****c9ks92
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The description.
	//
	// example:
	//
	// Data for labeling
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The dataset type. Valid values:
	//
	// - BASIC: Basic. Does not support dataset file metadata management.
	//
	//
	//
	// - ADVANCED: Advanced. Only supported for OSS type. Each version supports metadata management for up to 1 million files.
	//
	// example:
	//
	// BASIC
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The storage import configuration of the dataset. OSS, NAS, and CPFS are supported.
	//
	// <details>
	//
	// <summary>OSS</summary>
	//
	// {<BR>
	//
	// "region": "${region}",//Region ID<BR>
	//
	// "bucket": "${bucket}",//Bucket name<BR>
	//
	// "path": "${path}" //File path<BR>
	//
	// }<BR>
	//
	// </details>
	//
	// <details>
	//
	// <summary>NAS</summary>
	//
	// {<BR>
	//
	// "region": "${region}",//Region ID<BR>
	//
	// "fileSystemId": "${file_system_id}", //File system ID<BR>
	//
	// "path": "${path}", //File system path<BR>
	//
	// "mountTarget": "${mount_target}" //File system mount target<BR>
	//
	// }<BR>
	//
	// </details>
	//
	// <details>
	//
	// <summary>CPFS</summary>
	//
	// {<BR>
	//
	// "region": "${region}",//Region ID<BR>
	//
	// "fileSystemId": "${file_system_id}", //File system ID<BR>
	//
	// "protocolServiceId":"${protocol_service_id}", //File system protocol service<BR>
	//
	// "exportId": "${export_id}", //File system export directory<BR>
	//
	// "path": "${path}", //File system path<BR>
	//
	// }<BR>
	//
	// </details>
	//
	// <details>
	//
	// <summary>Lingjun CPFS</summary>
	//
	// {<BR>
	//
	// "region": "${region}",//Region ID<BR>
	//
	// "fileSystemId": "${file_system_id}", //File system ID<BR>
	//
	// "path": "${path}", //File system path<BR>
	//
	// "mountTarget": "${mount_target}" //File system mount target, specific to Lingjun edition<BR>
	//
	// "isVpcMount": boolean, //Whether it is a VPC mount target, specific to Lingjun edition<BR>
	//
	// }<BR>
	//
	// </details>
	//
	// example:
	//
	// {
	//
	//     "region": "cn-wulanchabu",
	//
	//     "fileSystemId": "bmcpfs-xxxxxxxxxxx",
	//
	//     "path": "/mnt",
	//
	//     "mountTarget": "cpfs-xxxxxxxxxxxx-vpc-gacs9f.cn-wulanchabu.cpfs.aliyuncs.com",
	//
	//     "isVpcMount": true
	//
	// }
	ImportInfo *string `json:"ImportInfo,omitempty" xml:"ImportInfo,omitempty"`
	// Indicates whether the dataset is a shared dataset.
	//
	// example:
	//
	// false
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The list of labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The latest version of the dataset.
	LatestVersion *DatasetVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The permission when the dataset is mounted. Valid values:
	//
	// - RO: read-only mount.
	//
	// - RW: read-write mount.
	//
	// example:
	//
	// RW
	MountAccess *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// The list of workspace role names that have read and write permission on the dataset. IDs starting with PAI are basic role IDs, and IDs starting with role- are custom role IDs. If the list contains "*", all roles have read and write permission.
	MountAccessReadWriteRoleIdList []*string `json:"MountAccessReadWriteRoleIdList,omitempty" xml:"MountAccessReadWriteRoleIdList,omitempty" type:"Repeated"`
	// The dataset name.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The extension field of the initial version v1, in JsonString format.
	//
	// When DLC uses the dataset, you can specify the default mount path of the dataset by configuring the mountPath field.
	//
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 1631044****3440
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The property of the initial dataset version v1. Valid values:
	//
	// - FILE: file.
	//
	// - DIRECTORY: folder.
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The dataset provider. If the value is "pai", the dataset is a PAI platform public dataset.
	//
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The data source provider type of the dataset. Valid values:
	//
	// - Ecs (default)
	//
	// - Lingjun
	//
	// example:
	//
	// Ecs
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source relationship of the shared dataset. This field is valid only when IsShared is true.
	SharedFrom *DatasetShareRelationship `json:"SharedFrom,omitempty" xml:"SharedFrom,omitempty"`
	// The sharing configuration of the current dataset.
	SharingConfig *GetDatasetResponseBodySharingConfig `json:"SharingConfig,omitempty" xml:"SharingConfig,omitempty" type:"Struct"`
	// The source dataset ID of the iTag labeling dataset.
	//
	// example:
	//
	// d-rcdg3wxxxxxhc5jk87
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	// The source dataset version of the labeling dataset.
	//
	// example:
	//
	// v2
	SourceDatasetVersion *string `json:"SourceDatasetVersion,omitempty" xml:"SourceDatasetVersion,omitempty"`
	// The source ID of the initial version v1. Valid values:
	//
	// - If SourceType is USER, SourceId can be customized.
	//
	// - If SourceType is ITAG, which indicates a dataset generated from iTAG labeling results, SourceId is the iTAG task ID.
	//
	// - If SourceType is PAI_PUBLIC_DATASET, which indicates a dataset created from a PAI public dataset, SourceId is empty by default.
	//
	// example:
	//
	// jdnhf***fnrimv
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the initial version v1.
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The labeling template of the iTag labeling dataset.
	//
	// example:
	//
	// TextClassification
	TagTemplateType *string `json:"TagTemplateType,omitempty" xml:"TagTemplateType,omitempty"`
	// The URI of the initial version v1. Example formats:
	//
	// - If the data source type is OSS: `oss://bucket.endpoint/object`.
	//
	// - If the data source type is NAS:
	//
	// General-purpose NAS format: `nas://<nasfisid>.region/subpath/to/dir/`.
	//
	// CPFS 1.0: `nas://<cpfs-fsid>.region/subpath/to/dir/`.
	//
	// CPFS 2.0: `nas://<cpfs-fsid>.region/<protocolserviceid>/`.
	//
	// CPFS 1.0 and CPFS 2.0 are distinguished by the format of the fsid:
	//
	// CPFS 1.0 format: cpfs-<8 ASCII characters>.
	//
	// CPFS 2.0 format: cpfs-<16 ASCII characters>.
	//
	// example:
	//
	// nas://09f****f2.cn-hangzhou/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The ID of the user to whom the dataset belongs.
	//
	// example:
	//
	// 2485765****023475
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the workspace to which the dataset belongs.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetDatasetResponseBody) GetAccessibleRoleIdList() []*string {
	return s.AccessibleRoleIdList
}

func (s *GetDatasetResponseBody) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetDatasetResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetDatasetResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetDatasetResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetDatasetResponseBody) GetEdition() *string {
	return s.Edition
}

func (s *GetDatasetResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetDatasetResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetDatasetResponseBody) GetImportInfo() *string {
	return s.ImportInfo
}

func (s *GetDatasetResponseBody) GetIsShared() *bool {
	return s.IsShared
}

func (s *GetDatasetResponseBody) GetLabels() []*Label {
	return s.Labels
}

func (s *GetDatasetResponseBody) GetLatestVersion() *DatasetVersion {
	return s.LatestVersion
}

func (s *GetDatasetResponseBody) GetMountAccess() *string {
	return s.MountAccess
}

func (s *GetDatasetResponseBody) GetMountAccessReadWriteRoleIdList() []*string {
	return s.MountAccessReadWriteRoleIdList
}

func (s *GetDatasetResponseBody) GetName() *string {
	return s.Name
}

func (s *GetDatasetResponseBody) GetOptions() *string {
	return s.Options
}

func (s *GetDatasetResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetDatasetResponseBody) GetProperty() *string {
	return s.Property
}

func (s *GetDatasetResponseBody) GetProvider() *string {
	return s.Provider
}

func (s *GetDatasetResponseBody) GetProviderType() *string {
	return s.ProviderType
}

func (s *GetDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetResponseBody) GetSharedFrom() *DatasetShareRelationship {
	return s.SharedFrom
}

func (s *GetDatasetResponseBody) GetSharingConfig() *GetDatasetResponseBodySharingConfig {
	return s.SharingConfig
}

func (s *GetDatasetResponseBody) GetSourceDatasetId() *string {
	return s.SourceDatasetId
}

func (s *GetDatasetResponseBody) GetSourceDatasetVersion() *string {
	return s.SourceDatasetVersion
}

func (s *GetDatasetResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *GetDatasetResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetDatasetResponseBody) GetTagTemplateType() *string {
	return s.TagTemplateType
}

func (s *GetDatasetResponseBody) GetUri() *string {
	return s.Uri
}

func (s *GetDatasetResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetDatasetResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDatasetResponseBody) SetAccessibility(v string) *GetDatasetResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetDatasetResponseBody) SetAccessibleRoleIdList(v []*string) *GetDatasetResponseBody {
	s.AccessibleRoleIdList = v
	return s
}

func (s *GetDatasetResponseBody) SetDataSourceType(v string) *GetDatasetResponseBody {
	s.DataSourceType = &v
	return s
}

func (s *GetDatasetResponseBody) SetDataType(v string) *GetDatasetResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetId(v string) *GetDatasetResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetResponseBody) SetDescription(v string) *GetDatasetResponseBody {
	s.Description = &v
	return s
}

func (s *GetDatasetResponseBody) SetEdition(v string) *GetDatasetResponseBody {
	s.Edition = &v
	return s
}

func (s *GetDatasetResponseBody) SetGmtCreateTime(v string) *GetDatasetResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetDatasetResponseBody) SetGmtModifiedTime(v string) *GetDatasetResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetDatasetResponseBody) SetImportInfo(v string) *GetDatasetResponseBody {
	s.ImportInfo = &v
	return s
}

func (s *GetDatasetResponseBody) SetIsShared(v bool) *GetDatasetResponseBody {
	s.IsShared = &v
	return s
}

func (s *GetDatasetResponseBody) SetLabels(v []*Label) *GetDatasetResponseBody {
	s.Labels = v
	return s
}

func (s *GetDatasetResponseBody) SetLatestVersion(v *DatasetVersion) *GetDatasetResponseBody {
	s.LatestVersion = v
	return s
}

func (s *GetDatasetResponseBody) SetMountAccess(v string) *GetDatasetResponseBody {
	s.MountAccess = &v
	return s
}

func (s *GetDatasetResponseBody) SetMountAccessReadWriteRoleIdList(v []*string) *GetDatasetResponseBody {
	s.MountAccessReadWriteRoleIdList = v
	return s
}

func (s *GetDatasetResponseBody) SetName(v string) *GetDatasetResponseBody {
	s.Name = &v
	return s
}

func (s *GetDatasetResponseBody) SetOptions(v string) *GetDatasetResponseBody {
	s.Options = &v
	return s
}

func (s *GetDatasetResponseBody) SetOwnerId(v string) *GetDatasetResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetDatasetResponseBody) SetProperty(v string) *GetDatasetResponseBody {
	s.Property = &v
	return s
}

func (s *GetDatasetResponseBody) SetProvider(v string) *GetDatasetResponseBody {
	s.Provider = &v
	return s
}

func (s *GetDatasetResponseBody) SetProviderType(v string) *GetDatasetResponseBody {
	s.ProviderType = &v
	return s
}

func (s *GetDatasetResponseBody) SetRequestId(v string) *GetDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetResponseBody) SetSharedFrom(v *DatasetShareRelationship) *GetDatasetResponseBody {
	s.SharedFrom = v
	return s
}

func (s *GetDatasetResponseBody) SetSharingConfig(v *GetDatasetResponseBodySharingConfig) *GetDatasetResponseBody {
	s.SharingConfig = v
	return s
}

func (s *GetDatasetResponseBody) SetSourceDatasetId(v string) *GetDatasetResponseBody {
	s.SourceDatasetId = &v
	return s
}

func (s *GetDatasetResponseBody) SetSourceDatasetVersion(v string) *GetDatasetResponseBody {
	s.SourceDatasetVersion = &v
	return s
}

func (s *GetDatasetResponseBody) SetSourceId(v string) *GetDatasetResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetDatasetResponseBody) SetSourceType(v string) *GetDatasetResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetDatasetResponseBody) SetTagTemplateType(v string) *GetDatasetResponseBody {
	s.TagTemplateType = &v
	return s
}

func (s *GetDatasetResponseBody) SetUri(v string) *GetDatasetResponseBody {
	s.Uri = &v
	return s
}

func (s *GetDatasetResponseBody) SetUserId(v string) *GetDatasetResponseBody {
	s.UserId = &v
	return s
}

func (s *GetDatasetResponseBody) SetWorkspaceId(v string) *GetDatasetResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetDatasetResponseBody) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LatestVersion != nil {
		if err := s.LatestVersion.Validate(); err != nil {
			return err
		}
	}
	if s.SharedFrom != nil {
		if err := s.SharedFrom.Validate(); err != nil {
			return err
		}
	}
	if s.SharingConfig != nil {
		if err := s.SharingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetResponseBodySharingConfig struct {
	// The list of sharing configuration relationships.
	SharedTo []*DatasetShareRelationship `json:"SharedTo,omitempty" xml:"SharedTo,omitempty" type:"Repeated"`
}

func (s GetDatasetResponseBodySharingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodySharingConfig) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodySharingConfig) GetSharedTo() []*DatasetShareRelationship {
	return s.SharedTo
}

func (s *GetDatasetResponseBodySharingConfig) SetSharedTo(v []*DatasetShareRelationship) *GetDatasetResponseBodySharingConfig {
	s.SharedTo = v
	return s
}

func (s *GetDatasetResponseBodySharingConfig) Validate() error {
	if s.SharedTo != nil {
		for _, item := range s.SharedTo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
