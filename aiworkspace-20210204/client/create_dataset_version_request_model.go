// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataCount(v int64) *CreateDatasetVersionRequest
	GetDataCount() *int64
	SetDataSize(v int64) *CreateDatasetVersionRequest
	GetDataSize() *int64
	SetDataSourceType(v string) *CreateDatasetVersionRequest
	GetDataSourceType() *string
	SetDatasetTaskRamRole(v string) *CreateDatasetVersionRequest
	GetDatasetTaskRamRole() *string
	SetDescription(v string) *CreateDatasetVersionRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateDatasetVersionRequest
	GetDryRun() *bool
	SetImportInfo(v string) *CreateDatasetVersionRequest
	GetImportInfo() *string
	SetLabels(v []*Label) *CreateDatasetVersionRequest
	GetLabels() []*Label
	SetOptions(v string) *CreateDatasetVersionRequest
	GetOptions() *string
	SetProperty(v string) *CreateDatasetVersionRequest
	GetProperty() *string
	SetSourceId(v string) *CreateDatasetVersionRequest
	GetSourceId() *string
	SetSourceType(v string) *CreateDatasetVersionRequest
	GetSourceType() *string
	SetUri(v string) *CreateDatasetVersionRequest
	GetUri() *string
	SetUserMetricsEndpoints(v []*UserMetricsEndpoint) *CreateDatasetVersionRequest
	GetUserMetricsEndpoints() []*UserMetricsEndpoint
}

type CreateDatasetVersionRequest struct {
	// The number of files in the dataset.
	//
	// example:
	//
	// 300
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// The size of the space occupied by the dataset files. Unit: bytes.
	//
	// example:
	//
	// 19000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The data source type. Separate multiple values with commas. Valid values:
	//
	// - NAS: Alibaba Cloud File Storage NAS.
	//
	// - OSS: Alibaba Cloud Object Storage Service (OSS).
	//
	// - CPFS
	//
	// Note: The DataSourceType of the version must be the same as the DataSourceType of the dataset. The system validates this when you create the version.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// UserMetricsEndpoints
	//
	// example:
	//
	// acs:ram::1234567890123456:role/role-name
	DatasetTaskRamRole *string `json:"DatasetTaskRamRole,omitempty" xml:"DatasetTaskRamRole,omitempty"`
	// The custom description of the dataset version. You can use this description to distinguish between different dataset versions.
	//
	// example:
	//
	// This is a description of the dataset version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run without actually creating the dataset version.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The storage import configuration of the dataset. OSS, NAS, and CPFS are supported.
	//
	// <details>
	//
	// <summary>OSS</summary>
	//
	// {<BR>
	//
	//   "region": "${region}",// The region ID.<BR>
	//
	//   "bucket": "${bucket}",// The bucket name.<BR>
	//
	//   "path": "${path}" // The file path.<BR>
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
	//   "region": "${region}",// The region ID.<BR>
	//
	//   "fileSystemId": "${file_system_id}", // The file system ID.<BR>
	//
	//   "path": "${path}", // The file system path.<BR>
	//
	//   "mountTarget": "${mount_target}" // The mount target of the file system.<BR>
	//
	// }<BR>
	//
	// </details>
	//
	//
	// <details>
	//
	// <summary>CPFS</summary>
	//
	// {<BR>
	//
	//   "region": "${region}",// The region ID.<BR>
	//
	//   "fileSystemId": "${file_system_id}", // The file system ID.<BR>
	//
	//   "protocolServiceId":"${protocol_service_id}", // The protocol service of the file system.<BR>
	//
	//   "exportId": "${export_id}", // The export directory of the file system.<BR>
	//
	//   "path": "${path}",  // The file system path.<BR>
	//
	// }<BR>
	//
	// </details>
	//
	// <details>
	//
	// <summary>CPFS for AI computing</summary>
	//
	// {<BR>
	//
	//   "region": "${region}",// The region ID.<BR>
	//
	//   "fileSystemId": "${file_system_id}", // The file system ID.<BR>
	//
	//   "path": "${path}",  // The file system path.<BR>
	//
	//   "mountTarget": "${mount_target}" // The mount target of the file system. This parameter is specific to the AI computing edition.<BR>
	//
	//   "isVpcMount": boolean, // Specifies whether the mount target is a VPC mount target. This parameter is specific to the AI computing edition.<BR>
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
	// The list of labels for the dataset version.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The extended field. The value is a JSON string. When DLC uses the dataset, you can configure the mountPath field to specify the default mount path for the dataset.
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
	// - FILE: file
	//
	// - DIRECTORY: directory
	//
	// This parameter is required.
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The source ID.
	//
	// - If SourceType is set to USER, you can specify a custom SourceId.
	//
	// - If SourceType is set to ITAG, the dataset is generated from the labeling results of the iTAG module, and SourceId is the task ID in iTAG.
	//
	// - If SourceType is set to PAI_PUBLIC_DATASET, the dataset is created from a PAI public dataset, and SourceId is empty by default.
	//
	// example:
	//
	// d-a0xbe5n03bhqof46ce
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type. Default value: USER. Valid values:
	//
	// - PAI-PUBLIC-DATASET: PAI public dataset.
	//
	// - ITAG: dataset generated from the labeling results of the iTAG module.
	//
	// - USER: user-registered dataset.
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The URI. Configuration examples:
	//
	// - If the data source type is OSS: `oss://bucket.endpoint/object`
	//
	// - If the data source type is NAS:
	//
	//   For general-purpose NAS, the format is `nas://<nasfisid>.region/subpath/to/dir/`.
	//
	//   For CPFS 1.0, the format is `nas://<cpfs-fsid>.region/subpath/to/dir/`.
	//
	//   For CPFS 2.0, the format is `nas://<cpfs-fsid>.region/<protocolserviceid>/`.
	//
	//   CPFS 1.0 and CPFS 2.0 are distinguished by the format of the fsid: the format for CPFS 1.0 is cpfs-<8 ASCII characters>, and the format for CPFS 2.0 is cpfs-<16 ASCII characters>.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The list of user metric collection endpoints.
	UserMetricsEndpoints []*UserMetricsEndpoint `json:"UserMetricsEndpoints,omitempty" xml:"UserMetricsEndpoints,omitempty" type:"Repeated"`
}

func (s CreateDatasetVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionRequest) GetDataCount() *int64 {
	return s.DataCount
}

func (s *CreateDatasetVersionRequest) GetDataSize() *int64 {
	return s.DataSize
}

func (s *CreateDatasetVersionRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CreateDatasetVersionRequest) GetDatasetTaskRamRole() *string {
	return s.DatasetTaskRamRole
}

func (s *CreateDatasetVersionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDatasetVersionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDatasetVersionRequest) GetImportInfo() *string {
	return s.ImportInfo
}

func (s *CreateDatasetVersionRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *CreateDatasetVersionRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateDatasetVersionRequest) GetProperty() *string {
	return s.Property
}

func (s *CreateDatasetVersionRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateDatasetVersionRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateDatasetVersionRequest) GetUri() *string {
	return s.Uri
}

func (s *CreateDatasetVersionRequest) GetUserMetricsEndpoints() []*UserMetricsEndpoint {
	return s.UserMetricsEndpoints
}

func (s *CreateDatasetVersionRequest) SetDataCount(v int64) *CreateDatasetVersionRequest {
	s.DataCount = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetDataSize(v int64) *CreateDatasetVersionRequest {
	s.DataSize = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetDataSourceType(v string) *CreateDatasetVersionRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetDatasetTaskRamRole(v string) *CreateDatasetVersionRequest {
	s.DatasetTaskRamRole = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetDescription(v string) *CreateDatasetVersionRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetDryRun(v bool) *CreateDatasetVersionRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetImportInfo(v string) *CreateDatasetVersionRequest {
	s.ImportInfo = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetLabels(v []*Label) *CreateDatasetVersionRequest {
	s.Labels = v
	return s
}

func (s *CreateDatasetVersionRequest) SetOptions(v string) *CreateDatasetVersionRequest {
	s.Options = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetProperty(v string) *CreateDatasetVersionRequest {
	s.Property = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetSourceId(v string) *CreateDatasetVersionRequest {
	s.SourceId = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetSourceType(v string) *CreateDatasetVersionRequest {
	s.SourceType = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetUri(v string) *CreateDatasetVersionRequest {
	s.Uri = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetUserMetricsEndpoints(v []*UserMetricsEndpoint) *CreateDatasetVersionRequest {
	s.UserMetricsEndpoints = v
	return s
}

func (s *CreateDatasetVersionRequest) Validate() error {
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
	return nil
}
