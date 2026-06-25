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
	// The visibility of the dataset in the workspace. Valid values:
	//
	// - `PRIVATE`: The dataset is visible only to its owner and workspace administrators.
	//
	// - `PUBLIC`: The dataset is visible to all members in the workspace.
	//
	// - `ROLE_PUBLIC`: The dataset is visible to specific workspace roles. For the list of roles, see the `AccessibleRoleIdList` parameter. The dataset owner and workspace administrators can always view the dataset.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// A list of workspace role IDs that can view the dataset. This parameter takes effect only when `Accessibility` is set to `ROLE_PUBLIC`. A role ID that starts with `PAI` is a basic role ID. A role ID that starts with `role-` is a custom role ID.
	AccessibleRoleIdList []*string `json:"AccessibleRoleIdList,omitempty" xml:"AccessibleRoleIdList,omitempty" type:"Repeated"`
	// The data source type. Valid values:
	//
	// - `OSS`: Object Storage Service (OSS).
	//
	// - `NAS`: Apsara File Storage NAS.
	//
	// example:
	//
	// NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The data type of the dataset. Valid values:
	//
	// - `COMMON`: General data
	//
	// - `PIC`: images
	//
	// - `TEXT`: text
	//
	// - `VIDEO`: videos
	//
	// - `AUDIO`: audio
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
	// The description of the dataset.
	//
	// example:
	//
	// 用于标注的数据。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The edition of the dataset. Valid values:
	//
	// - `BASIC`: The basic edition, which does not support file metadata management.
	//
	// - `ADVANCED`: The advanced edition, which is supported only for OSS datasets and allows you to manage metadata for up to 1 million files per version.
	//
	// example:
	//
	// BASIC
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The time when the dataset was created.
	//
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the dataset was last updated.
	//
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The storage import configuration of the dataset. Storage services such as OSS, NAS, and CPFS are supported.
	//
	// <details>
	//
	// <summary>
	//
	// OSS
	//
	// </summary>
	//
	// {\\
	//
	// "region": "${region}",// The region ID.\\
	//
	// "bucket": "${bucket}",// The bucket name.\\
	//
	// "path": "${path}" // The path to the file or folder.\\
	//
	// }
	//
	// </details>
	//
	// <details>
	//
	// <summary>
	//
	// NAS
	//
	// </summary>
	//
	// {\\
	//
	// "region": "${region}",// The region ID.\\
	//
	// "fileSystemId": "${file_system_id}", // The file system ID.\\
	//
	// "path": "${path}", // The path in the file system.\\
	//
	// "mountTarget": "${mount_target}" // The file system mount target.\\
	//
	// }
	//
	// </details>
	//
	// <details>
	//
	// <summary>
	//
	// CPFS
	//
	// </summary>
	//
	// {\\
	//
	// "region": "${region}",// The region ID.\\
	//
	// "fileSystemId": "${file_system_id}", // The file system ID.\\
	//
	// "protocolServiceId":"${protocol_service_id}", // The protocol service ID.\\
	//
	// "exportId": "${export_id}", // The export directory ID.\\
	//
	// "path": "${path}", // The path in the file system.\\
	//
	// }
	//
	// </details>
	//
	// <details>
	//
	// <summary>
	//
	// CPFS for Intelligent Computing
	//
	// </summary>
	//
	// {\\
	//
	// "region": "${region}",// The region ID.\\
	//
	// "fileSystemId": "${file_system_id}", // The file system ID.\\
	//
	// "path": "${path}", // The path in the file system.\\
	//
	// "mountTarget": "${mount_target}" // The file system mount target. This parameter is specific to CPFS for Intelligent Computing.\\
	//
	// "isVpcMount": boolean, // Specifies whether the mount target is a VPC mount target. Specific to CPFS for Intelligent Computing.\\
	//
	// }
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
	// The labels attached to the dataset.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The latest version of the dataset.
	LatestVersion *DatasetVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The mount permissions for the dataset. Valid values:
	//
	// - `RO`: read-only mount
	//
	// - `RW`: read and write mount
	//
	// example:
	//
	// RW
	MountAccess *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// A list of workspace role IDs granted read/write permissions for the dataset. A role ID that starts with `PAI` is a basic role ID. A role ID that starts with `role-` is a custom role ID. If the list contains `*`, all roles have read and write permissions.
	MountAccessReadWriteRoleIdList []*string `json:"MountAccessReadWriteRoleIdList,omitempty" xml:"MountAccessReadWriteRoleIdList,omitempty" type:"Repeated"`
	// The dataset name.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Extended properties for the initial dataset version (v1), in JSON string format. For example, when using the dataset in a DLC job, you can set the `mountPath` field to specify the default mount path.
	//
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The owner ID.
	//
	// example:
	//
	// 1631044****3440
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The property of the initial dataset version (v1). Valid values:
	//
	// - `FILE`: The dataset is a file.
	//
	// - `DIRECTORY`: The dataset is a folder.
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The provider of the dataset. If the value is `pai`, the dataset is a PAI public dataset.
	//
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The type of the data source provider. Valid values:
	//
	// - `ECS` (default)
	//
	// - `Lingjun`
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
	// The source from which the dataset was shared. This parameter is returned only if `IsShared` is `true`.
	SharedFrom *DatasetShareRelationship `json:"SharedFrom,omitempty" xml:"SharedFrom,omitempty"`
	// The sharing configuration for the dataset.
	SharingConfig *GetDatasetResponseBodySharingConfig `json:"SharingConfig,omitempty" xml:"SharingConfig,omitempty" type:"Struct"`
	// The ID of the source dataset for the iTAG annotation set.
	//
	// example:
	//
	// d-rcdg3wxxxxxhc5jk87
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	// The version of the source dataset for the annotation set.
	//
	// example:
	//
	// v2
	SourceDatasetVersion *string `json:"SourceDatasetVersion,omitempty" xml:"SourceDatasetVersion,omitempty"`
	// The ID of the data source for the initial version (v1). The meaning of this parameter varies based on the `SourceType` value.
	//
	// - If `SourceType` is `USER`, you can specify a custom value for `SourceId`.
	//
	// - If `SourceType` is `ITAG`, the dataset is generated from an iTAG annotation task, and `SourceId` is the task ID.
	//
	// - If `SourceType` is `PAI_PUBLIC_DATASET`, the dataset is created from a PAI public dataset. In this case, `SourceId` is empty.
	//
	// example:
	//
	// jdnhf***fnrimv
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the initial dataset version (v1).
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The annotation template of the iTAG annotation set.
	//
	// example:
	//
	// TextClassification
	TagTemplateType *string `json:"TagTemplateType,omitempty" xml:"TagTemplateType,omitempty"`
	// The URI of the initial dataset version (v1). The supported formats are as follows:
	//
	// - For an OSS data source: `oss://bucket.endpoint/object`.
	//
	// - For a NAS data source, the format varies by NAS type:
	//
	//   CPFS 1.0 and CPFS 2.0 are distinguished by the format of the file system ID ():
	//
	// example:
	//
	// nas://09f****f2.cn-hangzhou/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The user ID of the dataset owner.
	//
	// example:
	//
	// 2485765****023475
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the workspace where the dataset is located.
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
	// A list of relationships indicating to whom the dataset is shared.
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
