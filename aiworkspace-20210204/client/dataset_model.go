// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataset interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Dataset
	GetAccessibility() *string
	SetAccessibleRoleIdList(v []*string) *Dataset
	GetAccessibleRoleIdList() []*string
	SetDataSourceType(v string) *Dataset
	GetDataSourceType() *string
	SetDataType(v string) *Dataset
	GetDataType() *string
	SetDatasetId(v string) *Dataset
	GetDatasetId() *string
	SetDescription(v string) *Dataset
	GetDescription() *string
	SetEdition(v string) *Dataset
	GetEdition() *string
	SetGmtCreateTime(v string) *Dataset
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *Dataset
	GetGmtModifiedTime() *string
	SetImportInfo(v string) *Dataset
	GetImportInfo() *string
	SetIsShared(v bool) *Dataset
	GetIsShared() *bool
	SetLabels(v []*Label) *Dataset
	GetLabels() []*Label
	SetLatestVersion(v *DatasetVersion) *Dataset
	GetLatestVersion() *DatasetVersion
	SetMountAccess(v string) *Dataset
	GetMountAccess() *string
	SetMountAccessReadWriteRoleIdList(v []*string) *Dataset
	GetMountAccessReadWriteRoleIdList() []*string
	SetName(v string) *Dataset
	GetName() *string
	SetOptions(v string) *Dataset
	GetOptions() *string
	SetOwnerId(v string) *Dataset
	GetOwnerId() *string
	SetProperty(v string) *Dataset
	GetProperty() *string
	SetProviderType(v string) *Dataset
	GetProviderType() *string
	SetSharedFrom(v *DatasetShareRelationship) *Dataset
	GetSharedFrom() *DatasetShareRelationship
	SetSharingConfig(v *DatasetSharingConfig) *Dataset
	GetSharingConfig() *DatasetSharingConfig
	SetSourceDatasetId(v string) *Dataset
	GetSourceDatasetId() *string
	SetSourceDatasetVersion(v string) *Dataset
	GetSourceDatasetVersion() *string
	SetSourceId(v string) *Dataset
	GetSourceId() *string
	SetSourceType(v string) *Dataset
	GetSourceType() *string
	SetTagTemplateType(v string) *Dataset
	GetTagTemplateType() *string
	SetUri(v string) *Dataset
	GetUri() *string
	SetUserId(v string) *Dataset
	GetUserId() *string
	SetWorkspaceId(v string) *Dataset
	GetWorkspaceId() *string
}

type Dataset struct {
	// The workspace accessibility. Valid values:
	//
	// 	- PRIVATE (default): The dataset is accessible only to you and the administrator of the workspace.
	//
	// 	- PUBLIC: The dataset is accessible to all members in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility        *string   `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	AccessibleRoleIdList []*string `json:"AccessibleRoleIdList,omitempty" xml:"AccessibleRoleIdList,omitempty" type:"Repeated"`
	// The data source type.
	//
	// Valid values:
	//
	// 	- NAS
	//
	// 	- OSS
	//
	// example:
	//
	// NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The data type. Valid values:
	//
	// 	- COMMON (default)
	//
	// 	- PIC
	//
	// 	- TEXT
	//
	// 	- Video
	//
	// 	- AUDIO
	//
	// example:
	//
	// COMMON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// d-c0h44g3****j8o4348
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The dataset description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// BASIC
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The time when the dataset was created.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the dataset was modified.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The dataset import information, such as OSS, NAS, and CPFS.
	//
	// **OSS**
	//
	// { "region": "${region}",//The region ID. "bucket": "${bucket}",//The bucket name. "path": "${path}" //The file path. }
	//
	// **NAS**
	//
	// **CPFS**
	//
	// **CPFS for Lingjun**
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
	// The labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The latest dataset version.
	LatestVersion *DatasetVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// MountAccess
	//
	// example:
	//
	// RO RW
	MountAccess *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// The IDs of the roles that have read and write permissions on the dataset in the workspace. The IDs starting with PAI is the IDs of the basic roles, and the IDs starting with role- is the IDs of the custom roles. If the list contains "\\*", all roles have read and write permissions.
	MountAccessReadWriteRoleIdList []*string `json:"MountAccessReadWriteRoleIdList,omitempty" xml:"MountAccessReadWriteRoleIdList,omitempty" type:"Repeated"`
	// The dataset name.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The extended field that can be used as an option. The value is a JSON string. When you use the dataset in Deep Learning Containers (DLC), you can use the mountPath field to specify the default mount path of the dataset.
	//
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1631044****3440
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The dataset property. Valid values:
	//
	// 	- FILE
	//
	// 	- DIRECTORY
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The provider type of the dataset. Valid values:
	//
	// 	- Ecs (default)
	//
	// 	- Lingjun
	//
	// example:
	//
	// Ecs
	ProviderType  *string                   `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	SharedFrom    *DatasetShareRelationship `json:"SharedFrom,omitempty" xml:"SharedFrom,omitempty"`
	SharingConfig *DatasetSharingConfig     `json:"SharingConfig,omitempty" xml:"SharingConfig,omitempty" type:"Struct"`
	// The ID of the source dataset for the labeled dataset.
	//
	// example:
	//
	// d-bvfasdfxxxxj8o411
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	// The version of the source dataset for the labeled dataset.
	//
	// example:
	//
	// v2
	SourceDatasetVersion *string `json:"SourceDatasetVersion,omitempty" xml:"SourceDatasetVersion,omitempty"`
	// The source ID.
	//
	// example:
	//
	// jdnhf***fnrimv
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type.
	//
	// Valid values:
	//
	// 	- PAI_PUBLIC_DATASET
	//
	// 	- ITAG
	//
	// 	- USER
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The labeling template of the iTAG labeled dataset.
	//
	// example:
	//
	// text-classification
	TagTemplateType *string `json:"TagTemplateType,omitempty" xml:"TagTemplateType,omitempty"`
	// URI examples:
	//
	// 	- Object Storage Service (OSS) data source: `oss://bucket.endpoint/object`
	//
	// 	- File Storage NAS (NAS) data source: `nas://<nasfisid>.region/subpath/to/dir/`
	//
	// 	- Cloud Parallel File Storage (CPFS) 1.0 data source: `nas://<cpfs-fsid>.region/subpath/to/dir/`
	//
	// 	- CPFS 2.0 data source: `nas://<cpfs-fsid>.region/<protocolserviceid>/`
	//
	// >  You can distinguish CPFS 1.0 and CPFS 2.0 file systems based on the format of the file system ID: The ID of the CPFS 1.0 file system is in the cpfs-<8-bit ASCII characters> format. The ID of the CPFS 2.0 file system is in the cpfs-<16-bit ASCII characters> format.
	//
	// example:
	//
	// nas://09f****f2.cn-hangzhou/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The user ID.
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

func (s Dataset) String() string {
	return dara.Prettify(s)
}

func (s Dataset) GoString() string {
	return s.String()
}

func (s *Dataset) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Dataset) GetAccessibleRoleIdList() []*string {
	return s.AccessibleRoleIdList
}

func (s *Dataset) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *Dataset) GetDataType() *string {
	return s.DataType
}

func (s *Dataset) GetDatasetId() *string {
	return s.DatasetId
}

func (s *Dataset) GetDescription() *string {
	return s.Description
}

func (s *Dataset) GetEdition() *string {
	return s.Edition
}

func (s *Dataset) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Dataset) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Dataset) GetImportInfo() *string {
	return s.ImportInfo
}

func (s *Dataset) GetIsShared() *bool {
	return s.IsShared
}

func (s *Dataset) GetLabels() []*Label {
	return s.Labels
}

func (s *Dataset) GetLatestVersion() *DatasetVersion {
	return s.LatestVersion
}

func (s *Dataset) GetMountAccess() *string {
	return s.MountAccess
}

func (s *Dataset) GetMountAccessReadWriteRoleIdList() []*string {
	return s.MountAccessReadWriteRoleIdList
}

func (s *Dataset) GetName() *string {
	return s.Name
}

func (s *Dataset) GetOptions() *string {
	return s.Options
}

func (s *Dataset) GetOwnerId() *string {
	return s.OwnerId
}

func (s *Dataset) GetProperty() *string {
	return s.Property
}

func (s *Dataset) GetProviderType() *string {
	return s.ProviderType
}

func (s *Dataset) GetSharedFrom() *DatasetShareRelationship {
	return s.SharedFrom
}

func (s *Dataset) GetSharingConfig() *DatasetSharingConfig {
	return s.SharingConfig
}

func (s *Dataset) GetSourceDatasetId() *string {
	return s.SourceDatasetId
}

func (s *Dataset) GetSourceDatasetVersion() *string {
	return s.SourceDatasetVersion
}

func (s *Dataset) GetSourceId() *string {
	return s.SourceId
}

func (s *Dataset) GetSourceType() *string {
	return s.SourceType
}

func (s *Dataset) GetTagTemplateType() *string {
	return s.TagTemplateType
}

func (s *Dataset) GetUri() *string {
	return s.Uri
}

func (s *Dataset) GetUserId() *string {
	return s.UserId
}

func (s *Dataset) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Dataset) SetAccessibility(v string) *Dataset {
	s.Accessibility = &v
	return s
}

func (s *Dataset) SetAccessibleRoleIdList(v []*string) *Dataset {
	s.AccessibleRoleIdList = v
	return s
}

func (s *Dataset) SetDataSourceType(v string) *Dataset {
	s.DataSourceType = &v
	return s
}

func (s *Dataset) SetDataType(v string) *Dataset {
	s.DataType = &v
	return s
}

func (s *Dataset) SetDatasetId(v string) *Dataset {
	s.DatasetId = &v
	return s
}

func (s *Dataset) SetDescription(v string) *Dataset {
	s.Description = &v
	return s
}

func (s *Dataset) SetEdition(v string) *Dataset {
	s.Edition = &v
	return s
}

func (s *Dataset) SetGmtCreateTime(v string) *Dataset {
	s.GmtCreateTime = &v
	return s
}

func (s *Dataset) SetGmtModifiedTime(v string) *Dataset {
	s.GmtModifiedTime = &v
	return s
}

func (s *Dataset) SetImportInfo(v string) *Dataset {
	s.ImportInfo = &v
	return s
}

func (s *Dataset) SetIsShared(v bool) *Dataset {
	s.IsShared = &v
	return s
}

func (s *Dataset) SetLabels(v []*Label) *Dataset {
	s.Labels = v
	return s
}

func (s *Dataset) SetLatestVersion(v *DatasetVersion) *Dataset {
	s.LatestVersion = v
	return s
}

func (s *Dataset) SetMountAccess(v string) *Dataset {
	s.MountAccess = &v
	return s
}

func (s *Dataset) SetMountAccessReadWriteRoleIdList(v []*string) *Dataset {
	s.MountAccessReadWriteRoleIdList = v
	return s
}

func (s *Dataset) SetName(v string) *Dataset {
	s.Name = &v
	return s
}

func (s *Dataset) SetOptions(v string) *Dataset {
	s.Options = &v
	return s
}

func (s *Dataset) SetOwnerId(v string) *Dataset {
	s.OwnerId = &v
	return s
}

func (s *Dataset) SetProperty(v string) *Dataset {
	s.Property = &v
	return s
}

func (s *Dataset) SetProviderType(v string) *Dataset {
	s.ProviderType = &v
	return s
}

func (s *Dataset) SetSharedFrom(v *DatasetShareRelationship) *Dataset {
	s.SharedFrom = v
	return s
}

func (s *Dataset) SetSharingConfig(v *DatasetSharingConfig) *Dataset {
	s.SharingConfig = v
	return s
}

func (s *Dataset) SetSourceDatasetId(v string) *Dataset {
	s.SourceDatasetId = &v
	return s
}

func (s *Dataset) SetSourceDatasetVersion(v string) *Dataset {
	s.SourceDatasetVersion = &v
	return s
}

func (s *Dataset) SetSourceId(v string) *Dataset {
	s.SourceId = &v
	return s
}

func (s *Dataset) SetSourceType(v string) *Dataset {
	s.SourceType = &v
	return s
}

func (s *Dataset) SetTagTemplateType(v string) *Dataset {
	s.TagTemplateType = &v
	return s
}

func (s *Dataset) SetUri(v string) *Dataset {
	s.Uri = &v
	return s
}

func (s *Dataset) SetUserId(v string) *Dataset {
	s.UserId = &v
	return s
}

func (s *Dataset) SetWorkspaceId(v string) *Dataset {
	s.WorkspaceId = &v
	return s
}

func (s *Dataset) Validate() error {
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

type DatasetSharingConfig struct {
	SharedTo []*DatasetShareRelationship `json:"SharedTo,omitempty" xml:"SharedTo,omitempty" type:"Repeated"`
}

func (s DatasetSharingConfig) String() string {
	return dara.Prettify(s)
}

func (s DatasetSharingConfig) GoString() string {
	return s.String()
}

func (s *DatasetSharingConfig) GetSharedTo() []*DatasetShareRelationship {
	return s.SharedTo
}

func (s *DatasetSharingConfig) SetSharedTo(v []*DatasetShareRelationship) *DatasetSharingConfig {
	s.SharedTo = v
	return s
}

func (s *DatasetSharingConfig) Validate() error {
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
