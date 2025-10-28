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
	// The visibility of the workspace. Valid values:
	//
	// 	- PRIVATE: The workspace is visible only to you and the administrator of the workspace.
	//
	// 	- PUBLIC: The workspace is visible to all users.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- OSS: Object Storage Service (OSS)
	//
	// 	- NAS: File Storage NAS (NAS)
	//
	// example:
	//
	// NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The data type. Valid values:
	//
	// 	- COMMON: common
	//
	// 	- PIC: picture
	//
	// 	- TEXT: text
	//
	// 	- VIDEO: video
	//
	// 	- AUDIO: audio
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Edition     *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
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
	// The dataset configurations to be imported to a storage, such as OSS, NAS, or CPFS.
	//
	// **OSS**
	//
	// {\\
	//
	// "region": "${region}",// The region ID\\
	//
	// "bucket": "${bucket}",// The bucket name\\
	//
	// "path": "${path}" // The file path\\
	//
	// }\\
	//
	//
	// **NAS**
	//
	// {\\
	//
	// "region": "${region}",// The region ID\\
	//
	// "fileSystemId": "${file_system_id}", // The file system ID\\
	//
	// "path": "${path}", // The file system path\\
	//
	// "mountTarget": "${mount_target}" // The mount point of the file system\\
	//
	// }\\
	//
	//
	// **CPFS**
	//
	// {\\
	//
	// "region": "${region}",// The region ID\\
	//
	// "fileSystemId": "${file_system_id}", // The file system ID\\
	//
	// "protocolServiceId":"${protocol_service_id}", // The file system protocol service\\
	//
	// "exportId": "${export_id}", // The file system export directory\\
	//
	// "path": "${path}", // The file system path\\
	//
	// }\\
	//
	//
	// **CPFS for Lingjun**
	//
	// {\\
	//
	// "region": "${region}",// The region ID\\
	//
	// "fileSystemId": "${file_system_id}", // The file system ID\\
	//
	// "path": "${path}", // The file system path\\
	//
	// "mountTarget": "${mount_target}" // The mount point of the file system, CPFS for Lingjun only\\
	//
	// "isVpcMount": boolean, // Whether the mount point is a VPC mount point, CPFS for Lingjun only\\
	//
	// }\\
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
	IsShared   *bool   `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The tags.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The latest version of the dataset.
	LatestVersion *DatasetVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The access permission on the dataset when the dataset is mounted. Valid values:
	//
	// 	- RO: read-only permissions
	//
	// 	- RW: read and write permissions
	//
	// example:
	//
	// RW
	MountAccess *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// The list of role names in the workspace that have read and write permissions on the mounted database. The names start with PAI are basic role names and the names start with role- are custom role names. If the list contains asterisks (\\*), all roles have read and write permissions.
	MountAccessReadWriteRoleIdList []*string `json:"MountAccessReadWriteRoleIdList,omitempty" xml:"MountAccessReadWriteRoleIdList,omitempty" type:"Repeated"`
	// The dataset name.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The extended fields of the dataset v1 (initial version). The value is a JSON string. When you use the dataset in Deep Learning Containers (DLC), you can use the mountPath field to specify the default mount path of the dataset.
	//
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The ID of the Alibaba Could account.
	//
	// example:
	//
	// 1631044****3440
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The property of the dataset of the initial version v1. Valid values:
	//
	// 	- FILE
	//
	// 	- DIRECTORY
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The dataset provider. If the value pai is returned, the dataset is a public dataset in PAI.
	//
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The type of the data source for the dataset. Valid values:
	//
	// 	- Ecs (default)
	//
	// 	- Lingjun
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
	RequestId     *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SharedFrom    *DatasetShareRelationship            `json:"SharedFrom,omitempty" xml:"SharedFrom,omitempty"`
	SharingConfig *GetDatasetResponseBodySharingConfig `json:"SharingConfig,omitempty" xml:"SharingConfig,omitempty" type:"Struct"`
	// The ID of the source dataset generated from a labeling job of iTAG.
	//
	// example:
	//
	// d-rcdg3wxxxxxhc5jk87
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	// The version of the source dataset generated from a labeling job of iTAG.
	//
	// example:
	//
	// v2
	SourceDatasetVersion *string `json:"SourceDatasetVersion,omitempty" xml:"SourceDatasetVersion,omitempty"`
	// The ID of the source for the dataset v1 (initial version). Valid values:
	//
	// 	- If SourceType is set to USER, the value of SourceId can be a custom string.
	//
	// 	- If SourceType is set to ITAG, the value of SourceId is the ID of the labeling job of iTAG.
	//
	// 	- If SourceType is set to PAI_PUBLIC_DATASET, SourceId is empty by default.
	//
	// example:
	//
	// jdnhf***fnrimv
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The type of the source for the dataset v1 (initial version). Valid values:
	//
	// 	- PAI-PUBLIC-DATASET: a public dataset of Platform for AI (PAI).
	//
	// 	- ITAG: a dataset generated from a labeling job of iTAG.
	//
	// 	- USER: a dataset registered by a user.
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The labeling template for the source dataset generated from a labeling job of iTAG.
	//
	// example:
	//
	// TextClassification
	TagTemplateType *string `json:"TagTemplateType,omitempty" xml:"TagTemplateType,omitempty"`
	// The URI of the initial version v1.
	//
	// 	- Sample format for the OSS data source: `oss://bucket.endpoint/object`
	//
	// 	- Sample formats for the NAS data source: `nas://<nasfisid>.region/subpath/to/dir/`: General-purpose NAS. `nas://<cpfs-fsid>.region/subpath/to/dir/`: Cloud Parallel File Storage (CPFS) 1.0. `nas://<cpfs-fsid>.region/<protocolserviceid>/`: CPFS 2.0. You can distinguish CPFS 1.0 and CPFS 2.0 file systems based on the format of the file system ID. The ID for CPFS 1.0 is in the cpfs-<8-bit ASCII characters> format. The ID for CPFS 2.0 is in the cpfs-<16-bit ASCII characters> format.
	//
	// example:
	//
	// nas://09f****f2.cn-hangzhou/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The ID of the user to which the dataset belongs.
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
