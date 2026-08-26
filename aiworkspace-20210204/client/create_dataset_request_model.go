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
	SetDatasetTaskRamRole(v string) *CreateDatasetRequest
	GetDatasetTaskRamRole() *string
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
	SetUserMetricsEndpoints(v []*UserMetricsEndpoint) *CreateDatasetRequest
	GetUserMetricsEndpoints() []*UserMetricsEndpoint
	SetVersionDescription(v string) *CreateDatasetRequest
	GetVersionDescription() *string
	SetVersionLabels(v []*Label) *CreateDatasetRequest
	GetVersionLabels() []*Label
	SetWorkspaceId(v string) *CreateDatasetRequest
	GetWorkspaceId() *string
}

type CreateDatasetRequest struct {
	// The visibility of the workspace. Valid values:
	//
	// - PRIVATE (default): visible only to yourself and administrators within the workspace.
	//
	// - PUBLIC: visible to all users in the workspace.
	//
	// - ROLE_PUBLIC: visible to specified workspace roles. For the role list, refer to AccessibleRoleIdList. Under this condition, the dataset owner and administrators always have visibility.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// Takes effect when Accessibility is set to ROLE_PUBLIC. The list of workspace role names that can view the dataset. IDs starting with PAI are basic role IDs, and IDs starting with role- are custom role IDs.
	AccessibleRoleIdList []*string `json:"AccessibleRoleIdList,omitempty" xml:"AccessibleRoleIdList,omitempty" type:"Repeated"`
	// The number of files in the dataset.
	//
	// example:
	//
	// 500
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// The size of space occupied by the dataset files. Unit: bytes.
	//
	// example:
	//
	// 10000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The data source type. Valid values:
	//
	// - OSS: Alibaba Cloud Object Storage Service (OSS).
	//
	// - NAS: Alibaba Cloud Apsara File Storage NAS General Purpose.
	//
	// - EXTREMENAS: Alibaba Cloud Apsara File Storage NAS Extreme.
	//
	// - CPFS: Alibaba Cloud Cloud Parallel File Storage (CPFS) General Purpose.
	//
	// - BMCPFS: Alibaba Cloud Cloud Parallel File Storage (CPFS) AI Edition.
	//
	// - MAXCOMPUTE: Alibaba Cloud MaxCompute.
	//
	// - URL: public HTTP/HTTPS URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The data type of the dataset. Default value: COMMON. Valid values:
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
	// DatasetTaskRamRole
	//
	// example:
	//
	// acs:ram::1234567890123456:role/role-name
	DatasetTaskRamRole *string `json:"DatasetTaskRamRole,omitempty" xml:"DatasetTaskRamRole,omitempty"`
	// The custom description of the dataset to distinguish it from other datasets.
	//
	// example:
	//
	// This is a description of the dataset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The dataset type. Default value: BASIC. Valid values:
	//
	// - BASIC: basic. Does not support dataset file metadata management.
	//
	// - ADVANCED: advanced. Only supported for OSS type. Each version supports up to 1 million file metadata entries.
	//
	// - LOGICAL: logical. Only supported for OSS type. Each version supports up to 3 million file metadata entries.
	//
	// example:
	//
	// ADVANCED
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
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
	// <summary>AI Edition CPFS</summary>
	//
	// {<BR>
	//
	// "region": "${region}",//Region ID<BR>
	//
	// "fileSystemId": "${file_system_id}", //File system ID<BR>
	//
	// "path": "${path}", //File system path<BR>
	//
	// "mountTarget": "${mount_target}" //File system mount target, specific to AI Edition<BR>
	//
	// "isVpcMount": boolean, //Whether it is a VPC mount target, specific to AI Edition<BR>
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
	// The list of labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The list of workspace role names that have read and write permissions when the dataset is mounted. IDs starting with PAI are basic role IDs, and IDs starting with role- are custom role IDs. If the list contains "*", all roles have read and write permissions.
	//
	// - Specified roles: ["PAI.AlgoOperator", "role-hiuwpd01ncrokkgp21"]
	//
	// - All accounts: ["*"]
	//
	// - Dataset creator only: []
	MountAccessReadWriteRoleIdList []*string `json:"MountAccessReadWriteRoleIdList,omitempty" xml:"MountAccessReadWriteRoleIdList,omitempty" type:"Repeated"`
	// The name of the dataset. Naming rules:
	//
	// - Must start with a lowercase letter, uppercase letter, digit, or Chinese character.
	//
	// - Can contain underscores (_) or hyphens (-).
	//
	// - Must be 1 to 127 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The extended field in JsonString format.
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
	// The property of the dataset. Valid values:
	//
	// - FILE: file.
	//
	// - DIRECTORY: folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The dataset provider. Cannot be set to pai.
	//
	// example:
	//
	// Github
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
	// The source dataset ID of the annotation dataset.
	//
	// example:
	//
	// d-bvfasdfxxxxj8o411
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	// The source dataset version of the annotation dataset.
	//
	// example:
	//
	// v2
	SourceDatasetVersion *string `json:"SourceDatasetVersion,omitempty" xml:"SourceDatasetVersion,omitempty"`
	// The data source ID.
	//
	// - If SourceType is USER, SourceId can be customized.
	//
	// - If SourceType is ITAG, which indicates a dataset generated from iTAG annotation results, SourceId is the iTAG task ID.
	//
	// - If SourceType is PAI_PUBLIC_DATASET, which indicates a dataset created from a PAI public dataset, SourceId is empty by default.
	//
	// example:
	//
	// jdnhf***fnrimv
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The data source type. Default value: USER.
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Examples of Uri configurations:
	//
	// - If the data source type is OSS: `oss://bucket.endpoint/object`
	//
	// - If the data source type is NAS:
	//
	// General Purpose NAS format: `nas://<nasfisid>.region/subpath/to/dir/`;
	//
	// CPFS 1.0: `nas://<cpfs-fsid>.region/subpath/to/dir/`;
	//
	// CPFS 2.0: `nas://<cpfs-fsid>.region/<protocolserviceid>/`.
	//
	// CPFS 1.0 and CPFS 2.0 are distinguished by the fsid format: CPFS 1.0 format is cpfs-<8 ASCII characters>; CPFS 2.0 format is cpfs-<16 ASCII characters>.
	//
	// This parameter is required.
	//
	// example:
	//
	// nas://09f****f2.cn-hangzhou/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The Alibaba Cloud account ID of the dataset owner. Workspace owners and administrators have permissions to create datasets for specified workspace members.
	//
	// example:
	//
	// 2485765****023475
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// UserMetricsEndpoints
	UserMetricsEndpoints []*UserMetricsEndpoint `json:"UserMetricsEndpoints,omitempty" xml:"UserMetricsEndpoints,omitempty" type:"Repeated"`
	// The description of the initial version of the dataset.
	//
	// example:
	//
	// This is a description of the first dataset version.
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// The list of labels for the initial version.
	VersionLabels []*Label `json:"VersionLabels,omitempty" xml:"VersionLabels,omitempty" type:"Repeated"`
	// The ID of the workspace where the dataset resides. For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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

func (s *CreateDatasetRequest) GetDatasetTaskRamRole() *string {
	return s.DatasetTaskRamRole
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

func (s *CreateDatasetRequest) GetUserMetricsEndpoints() []*UserMetricsEndpoint {
	return s.UserMetricsEndpoints
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

func (s *CreateDatasetRequest) SetDatasetTaskRamRole(v string) *CreateDatasetRequest {
	s.DatasetTaskRamRole = &v
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

func (s *CreateDatasetRequest) SetUserMetricsEndpoints(v []*UserMetricsEndpoint) *CreateDatasetRequest {
	s.UserMetricsEndpoints = v
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
	if s.UserMetricsEndpoints != nil {
		for _, item := range s.UserMetricsEndpoints {
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
