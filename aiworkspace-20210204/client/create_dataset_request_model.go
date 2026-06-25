// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateDatasetRequest
	GetAccessibility() *string
	SetAccessibleRoleIdList(v []*string) *CreateDatasetRequest
	GetAccessibleRoleIdList() []*string
	SetDataCount(v int64) *CreateDatasetRequest
	GetDataCount() *int64
	SetDataSize(v int64) *CreateDatasetRequest
	GetDataSize() *int64
	SetDataSourceType(v string) *CreateDatasetRequest
	GetDataSourceType() *string
	SetDataType(v string) *CreateDatasetRequest
	GetDataType() *string
	SetDescription(v string) *CreateDatasetRequest
	GetDescription() *string
	SetEdition(v string) *CreateDatasetRequest
	GetEdition() *string
	SetImportInfo(v string) *CreateDatasetRequest
	GetImportInfo() *string
	SetLabels(v []*Label) *CreateDatasetRequest
	GetLabels() []*Label
	SetMountAccessReadWriteRoleIdList(v []*string) *CreateDatasetRequest
	GetMountAccessReadWriteRoleIdList() []*string
	SetName(v string) *CreateDatasetRequest
	GetName() *string
	SetOptions(v string) *CreateDatasetRequest
	GetOptions() *string
	SetProperty(v string) *CreateDatasetRequest
	GetProperty() *string
	SetProvider(v string) *CreateDatasetRequest
	GetProvider() *string
	SetProviderType(v string) *CreateDatasetRequest
	GetProviderType() *string
	SetSourceDatasetId(v string) *CreateDatasetRequest
	GetSourceDatasetId() *string
	SetSourceDatasetVersion(v string) *CreateDatasetRequest
	GetSourceDatasetVersion() *string
	SetSourceId(v string) *CreateDatasetRequest
	GetSourceId() *string
	SetSourceType(v string) *CreateDatasetRequest
	GetSourceType() *string
	SetUri(v string) *CreateDatasetRequest
	GetUri() *string
	SetUserId(v string) *CreateDatasetRequest
	GetUserId() *string
	SetVersionDescription(v string) *CreateDatasetRequest
	GetVersionDescription() *string
	SetVersionLabels(v []*Label) *CreateDatasetRequest
	GetVersionLabels() []*Label
	SetWorkspaceId(v string) *CreateDatasetRequest
	GetWorkspaceId() *string
}

type CreateDatasetRequest struct {
	// The visibility of the dataset in the workspace. Valid values:
	//
	// - PRIVATE (default): The dataset is visible only to its owner and administrators in the workspace.
	//
	// - PUBLIC: The dataset is visible to all users in the workspace.
	//
	// - ROLE_PUBLIC: The dataset is visible to users with specific workspace roles. The list of roles is specified in the `AccessibleRoleIdList` parameter. The dataset owner and administrators always retain visibility.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// This parameter takes effect only when `Accessibility` is set to `ROLE_PUBLIC`. This parameter specifies a list of workspace role IDs that can view this dataset. Role IDs that start with `PAI.` are built-in roles, and role IDs that start with `role-` are custom roles.
	AccessibleRoleIdList []*string `json:"AccessibleRoleIdList,omitempty" xml:"AccessibleRoleIdList,omitempty" type:"Repeated"`
	// The number of files in the dataset.
	//
	// example:
	//
	// 500
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// The size of the dataset files, in bytes.
	//
	// example:
	//
	// 10000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The type of the data source. Valid values:
	//
	// - OSS: Object Storage Service (OSS).
	//
	// - NAS: general-purpose Apsara File Storage NAS.
	//
	// - EXTREMENAS: Extreme NAS.
	//
	// - CPFS: general-purpose Cloud Parallel File Storage (CPFS).
	//
	// - BMCPFS: AI Computing Edition of CPFS.
	//
	// - MAXCOMPUTE: MaxCompute.
	//
	// - URL: a public HTTP or HTTPS URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The data type of the dataset. The default value is `COMMON`. Valid values:
	//
	// - COMMON: common
	//
	// - PIC: image
	//
	// - TEXT: text
	//
	// - VIDEO: video
	//
	// - AUDIO: audio
	//
	// example:
	//
	// COMMON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// A custom description to distinguish the dataset from other datasets.
	//
	// example:
	//
	// This is a description of the dataset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The edition of the dataset. The default value is BASIC. Valid values:
	//
	// - BASIC: Basic. Does not support dataset file metadata management.
	//
	// - ADVANCED: Advanced. Supported only for OSS datasets. Each version supports metadata management for up to 1 million files.
	//
	// - LOGICAL: Logical. Supported only for OSS datasets. Each version supports metadata management for up to 3 million files.
	//
	// example:
	//
	// ADVANCED
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The storage import configuration of the dataset. `OSS`, `NAS`, and `CPFS` are supported.
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
	// "path": "${path}" // The file path.\\
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
	// "path": "${path}", // The file system path.\\
	//
	// "mountTarget": "${mount_target}" // The mount target of the file system.\\
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
	// "protocolServiceId":"${protocol_service_id}", // The protocol service of the file system.\\
	//
	// "exportId": "${export_id}", // The exported directory of the file system.\\
	//
	// "path": "${path}", // The file system path.\\
	//
	// }
	//
	// </details>
	//
	// <details>
	//
	// <summary>
	//
	// CPFS (AI Computing Edition)
	//
	// </summary>
	//
	// {\\
	//
	// "region": "${region}",// The region ID.\\
	//
	// "fileSystemId": "${file_system_id}", // The file system ID.\\
	//
	// "path": "${path}", // The file system path.\\
	//
	// "mountTarget": "${mount_target}", // The mount target of the file system. This parameter is specific to the AI Computing Edition.\\
	//
	// "isVpcMount": boolean, // Specifies whether the mount target is in a VPC. This parameter is specific to the AI Computing Edition.\\
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
	// A list of labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// A list of workspace role IDs that are granted read and write permissions when the dataset is mounted. Role IDs that start with `PAI.` are built-in roles, and role IDs that start with `role-` are custom roles. If the list contains an asterisk (\\*), all roles are granted read and write permissions.
	//
	// - Accounts with specified roles: `["PAI.AlgoOperator", "role-hiuwpd01ncrokkgp21"]`
	//
	// - All accounts: `["*"]`
	//
	// - Dataset creator only: `[]`
	MountAccessReadWriteRoleIdList []*string `json:"MountAccessReadWriteRoleIdList,omitempty" xml:"MountAccessReadWriteRoleIdList,omitempty" type:"Repeated"`
	// The name of the dataset. The name must meet the following requirements:
	//
	// - Starts with a lowercase letter, an uppercase letter, a number, or a Chinese character.
	//
	// - Can contain underscores (_) and hyphens (-).
	//
	// - Must be 1 to 127 characters long.
	//
	// This parameter is required.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The extended fields, which are a JSON string.
	//
	// When a Data Lake Compute (DLC) job uses the dataset, you can configure the `mountPath` field to specify the default mount path of the dataset.
	//
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The property of the dataset. Valid values:
	//
	// - FILE: A file.
	//
	// - DIRECTORY: A directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The provider of the dataset. You cannot set this parameter to `pai`.
	//
	// example:
	//
	// Github
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The type of the data source provider. Valid values:
	//
	// - Ecs (default)
	//
	// - Lingjun
	//
	// example:
	//
	// Ecs
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// The ID of the source dataset for a labeled dataset.
	//
	// example:
	//
	// d-bvfasdfxxxxj8o411
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	// The version of the source dataset for a labeled dataset.
	//
	// example:
	//
	// v2
	SourceDatasetVersion *string `json:"SourceDatasetVersion,omitempty" xml:"SourceDatasetVersion,omitempty"`
	// The ID of the data source.
	//
	// - If `SourceType` is `USER`, you can specify a custom value for `SourceId`.
	//
	// - If `SourceType` is `ITAG`, this parameter specifies the iTAG task ID from which the dataset was generated.
	//
	// - If `SourceType` is `PAI_PUBLIC_DATASET`, the dataset is from a public PAI dataset, and this parameter is empty by default.
	//
	// example:
	//
	// jdnhf***fnrimv
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source of the data. The default value is USER.
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The URI of the data. The URI format varies based on the `DataSourceType` value.
	//
	// - For an `OSS` data source: `oss://bucket.endpoint/object`
	//
	// - For a `NAS` data source:
	//
	//   For general-purpose `NAS`: `nas://<nasfisid>.region/subpath/to/dir/`.
	//
	//   For `CPFS` 1.0: `nas://<cpfs-fsid>.region/subpath/to/dir/`.
	//
	//   For `CPFS` 2.0: `nas://<cpfs-fsid>.region/<protocolserviceid>/`.
	//
	//   `CPFS` 1.0 and `CPFS` 2.0 are distinguished by the format of the file system ID (fsid). The fsid for `CPFS` 1.0 is in the `cpfs-<8-character ASCII string>` format. The fsid for `CPFS` 2.0 is in the `cpfs-<16-character ASCII string>` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// nas://09f****f2.cn-hangzhou/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The Alibaba Cloud account ID of the dataset owner. Workspace owners and administrators can create datasets for specified members of a workspace.
	//
	// example:
	//
	// 2485765****023475
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The description of the initial version of the dataset.
	//
	// example:
	//
	// This is a description of the first dataset version.
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// A list of labels for the initial version.
	VersionLabels []*Label `json:"VersionLabels,omitempty" xml:"VersionLabels,omitempty" type:"Repeated"`
	// The ID of the workspace to which the dataset belongs. For more information about how to obtain a workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// If this parameter is not specified, the default workspace is used. If the default workspace does not exist, an error is returned.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateDatasetRequest) GetAccessibleRoleIdList() []*string {
	return s.AccessibleRoleIdList
}

func (s *CreateDatasetRequest) GetDataCount() *int64 {
	return s.DataCount
}

func (s *CreateDatasetRequest) GetDataSize() *int64 {
	return s.DataSize
}

func (s *CreateDatasetRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CreateDatasetRequest) GetDataType() *string {
	return s.DataType
}

func (s *CreateDatasetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDatasetRequest) GetEdition() *string {
	return s.Edition
}

func (s *CreateDatasetRequest) GetImportInfo() *string {
	return s.ImportInfo
}

func (s *CreateDatasetRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *CreateDatasetRequest) GetMountAccessReadWriteRoleIdList() []*string {
	return s.MountAccessReadWriteRoleIdList
}

func (s *CreateDatasetRequest) GetName() *string {
	return s.Name
}

func (s *CreateDatasetRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateDatasetRequest) GetProperty() *string {
	return s.Property
}

func (s *CreateDatasetRequest) GetProvider() *string {
	return s.Provider
}

func (s *CreateDatasetRequest) GetProviderType() *string {
	return s.ProviderType
}

func (s *CreateDatasetRequest) GetSourceDatasetId() *string {
	return s.SourceDatasetId
}

func (s *CreateDatasetRequest) GetSourceDatasetVersion() *string {
	return s.SourceDatasetVersion
}

func (s *CreateDatasetRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateDatasetRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateDatasetRequest) GetUri() *string {
	return s.Uri
}

func (s *CreateDatasetRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateDatasetRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *CreateDatasetRequest) GetVersionLabels() []*Label {
	return s.VersionLabels
}

func (s *CreateDatasetRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDatasetRequest) SetAccessibility(v string) *CreateDatasetRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateDatasetRequest) SetAccessibleRoleIdList(v []*string) *CreateDatasetRequest {
	s.AccessibleRoleIdList = v
	return s
}

func (s *CreateDatasetRequest) SetDataCount(v int64) *CreateDatasetRequest {
	s.DataCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDataSize(v int64) *CreateDatasetRequest {
	s.DataSize = &v
	return s
}

func (s *CreateDatasetRequest) SetDataSourceType(v string) *CreateDatasetRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateDatasetRequest) SetDataType(v string) *CreateDatasetRequest {
	s.DataType = &v
	return s
}

func (s *CreateDatasetRequest) SetDescription(v string) *CreateDatasetRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetRequest) SetEdition(v string) *CreateDatasetRequest {
	s.Edition = &v
	return s
}

func (s *CreateDatasetRequest) SetImportInfo(v string) *CreateDatasetRequest {
	s.ImportInfo = &v
	return s
}

func (s *CreateDatasetRequest) SetLabels(v []*Label) *CreateDatasetRequest {
	s.Labels = v
	return s
}

func (s *CreateDatasetRequest) SetMountAccessReadWriteRoleIdList(v []*string) *CreateDatasetRequest {
	s.MountAccessReadWriteRoleIdList = v
	return s
}

func (s *CreateDatasetRequest) SetName(v string) *CreateDatasetRequest {
	s.Name = &v
	return s
}

func (s *CreateDatasetRequest) SetOptions(v string) *CreateDatasetRequest {
	s.Options = &v
	return s
}

func (s *CreateDatasetRequest) SetProperty(v string) *CreateDatasetRequest {
	s.Property = &v
	return s
}

func (s *CreateDatasetRequest) SetProvider(v string) *CreateDatasetRequest {
	s.Provider = &v
	return s
}

func (s *CreateDatasetRequest) SetProviderType(v string) *CreateDatasetRequest {
	s.ProviderType = &v
	return s
}

func (s *CreateDatasetRequest) SetSourceDatasetId(v string) *CreateDatasetRequest {
	s.SourceDatasetId = &v
	return s
}

func (s *CreateDatasetRequest) SetSourceDatasetVersion(v string) *CreateDatasetRequest {
	s.SourceDatasetVersion = &v
	return s
}

func (s *CreateDatasetRequest) SetSourceId(v string) *CreateDatasetRequest {
	s.SourceId = &v
	return s
}

func (s *CreateDatasetRequest) SetSourceType(v string) *CreateDatasetRequest {
	s.SourceType = &v
	return s
}

func (s *CreateDatasetRequest) SetUri(v string) *CreateDatasetRequest {
	s.Uri = &v
	return s
}

func (s *CreateDatasetRequest) SetUserId(v string) *CreateDatasetRequest {
	s.UserId = &v
	return s
}

func (s *CreateDatasetRequest) SetVersionDescription(v string) *CreateDatasetRequest {
	s.VersionDescription = &v
	return s
}

func (s *CreateDatasetRequest) SetVersionLabels(v []*Label) *CreateDatasetRequest {
	s.VersionLabels = v
	return s
}

func (s *CreateDatasetRequest) SetWorkspaceId(v string) *CreateDatasetRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDatasetRequest) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VersionLabels != nil {
		for _, item := range s.VersionLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
