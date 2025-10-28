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
	// The workspace accessibility. Valid values:
	//
	// 	- PRIVATE: The workspace is accessible only to you and the administrator of the workspace. This is the default value.
	//
	// 	- PUBLIC: The workspace is accessible to all users.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The number of dataset files.
	//
	// example:
	//
	// 500
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// The size of the dataset file. Unit: bytes.
	//
	// example:
	//
	// 10000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The data source type. Valid values:
	//
	// 	- OSS: Object Storage Service (OSS).
	//
	// 	- NAS: File Storage NAS (NAS).
	//
	// This parameter is required.
	//
	// example:
	//
	// NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The type of the dataset. Default value: COMMON. Valid values:
	//
	// 	- COMMON: common
	//
	// 	- PIC: picture
	//
	// 	- TEXT: text
	//
	// 	- Video: video
	//
	// 	- AUDIO: audio
	//
	// example:
	//
	// COMMON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The description of the dataset. Descriptions are used to differentiate datasets.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Edition     *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The dataset configurations to be imported to a storage, such as OSS, NAS, or Cloud Parallel File Storage (CPFS).
	//
	// **OSS**
	//
	// {\\
	//
	// "region": "${region}",// The region ID.\\
	//
	// "bucket": "${bucket}",//The bucket name.\\
	//
	// "path": "${path}" // The file path.\\
	//
	// }\\
	//
	//
	// **NAS**
	//
	// {\\
	//
	// "region": "${region}",// The region ID.\\
	//
	// "fileSystemId": "${file_system_id}", // The file system ID.\\
	//
	// "path": "${path}", // The file system path.\\
	//
	// "mountTarget": "${mount_target}" // The mount point of the file system.\\
	//
	// }\\
	//
	//
	// **CPFS**
	//
	// {\\
	//
	// "region": "${region}",// The region ID.\\
	//
	// "fileSystemId": "${file_system_id}", // The file system ID.\\
	//
	// "protocolServiceId":"${protocol_service_id}", // The file system protocol service.\\
	//
	// "exportId": "${export_id}", // The file system export directory.\\
	//
	// "path": "${path}", // The file system path.\\
	//
	// }\\
	//
	//
	// **CPFS for Lingjun**
	//
	// {\\
	//
	// "region": "${region}",// The region ID.\\
	//
	// "fileSystemId": "${file_system_id}", // The file system ID.\\
	//
	// "path": "${path}", // The file system path.\\
	//
	// "mountTarget": "${mount_target}" // The mount point of the file system, CPFS for Lingjun only.\\
	//
	// "isVpcMount": boolean, // Whether the mount point is a virtual private cloud (VPC) mount point, CPFS for Lingjun only.\\
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
	// The tags.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The list of role names in the workspace that have read and write permissions on the mounted database. The names start with PAI are basic role names and the names start with role- are custom role names. If the list contains asterisks (\\*), all roles have read and write permissions.
	//
	// 	- If you set the value to ["PAI.AlgoOperator", "role-hiuwpd01ncrokkgp21"], the account of the specified role is granted the read and write permissions.
	//
	// 	- If you set the value to ["\\*"], all accounts are granted the read and write permissions.
	//
	// 	- If you set the value to [], only the creator of the dataset has the read and write permissions.
	MountAccessReadWriteRoleIdList []*string `json:"MountAccessReadWriteRoleIdList,omitempty" xml:"MountAccessReadWriteRoleIdList,omitempty" type:"Repeated"`
	// The dataset name. The name must meet the following requirements:
	//
	// 	- The name must start with a letter, digit, or Chinese character.
	//
	// 	- The name can contain underscores (_) and hyphens (-).
	//
	// 	- The name must be 1 to 127 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The extended field, which is a JSON string. When you use the dataset in Deep Learning Containers (DLC), you can configure the mountPath field to specify the default mount path of the dataset.
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
	// 	- FILE
	//
	// 	- DIRECTORY
	//
	// This parameter is required.
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The dataset provider. The value cannot be set to pai.
	//
	// example:
	//
	// Github
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The source type of the dataset. Valid values:
	//
	// 	- Ecs (default)
	//
	// 	- Lingjun
	//
	// example:
	//
	// Ecs
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
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
	// The data source ID.
	//
	// 	- If SourceType is set to USER, the value of SourceId is a custom string.
	//
	// 	- If SourceType is set to ITAG, the value of SourceId is the ID of the labeling job of iTAG.
	//
	// 	- If SourceType is set to PAI_PUBLIC_DATASET, SourceId is empty by default.
	//
	// example:
	//
	// jdnhf***fnrimv
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The type of the data source. Default value: USER.
	//
	// Valid values:
	//
	// 	- PAI_PUBLIC_DATASET: a public dataset of PAI.
	//
	// 	- ITAG: a dataset generated from a labeling job of iTAG.
	//
	// 	- USER: a dataset registered by a user.
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The URI of the data source.
	//
	// 	- Value format if DataSourceType is set to OSS: `oss://bucket.endpoint/object`.
	//
	// 	- Value formats if DataSourceType is set to NAS: General-purpose NAS: `nas://<nasfisid>.region/subpath/to/dir/`. CPFS 1.0: `nas://<cpfs-fsid>.region/subpath/to/dir/`. CPFS 2.0: `nas://<cpfs-fsid>.region/<protocolserviceid>/`. You can distinguish CPFS 1.0 and CPFS 2.0 file systems based on the format of the file system ID: The ID for CPFS 1.0 is in the cpfs-<8-bit ASCII characters> format. The ID for CPFS 2.0 is in the cpfs-<16-bit ASCII characters> format.
	//
	// This parameter is required.
	//
	// example:
	//
	// nas://09f****f2.cn-hangzhou/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The ID of the Alibaba Cloud account to which the dataset belongs. The workspace owner and administrator have permissions to create datasets for specified members in the workspace.
	//
	// example:
	//
	// 2485765****023475
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The description of the dataset of the initial version.
	//
	// example:
	//
	// The initial version
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// The list of tags to be added to the dataset of the initial version.
	VersionLabels []*Label `json:"VersionLabels,omitempty" xml:"VersionLabels,omitempty" type:"Repeated"`
	// The ID of the workspace to which the dataset belongs. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID. If you do not specify this parameter, the default workspace is used. If the default workspace does not exist, an error is reported.
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
