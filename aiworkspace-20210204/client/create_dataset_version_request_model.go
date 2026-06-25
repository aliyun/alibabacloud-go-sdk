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
	SetDescription(v string) *CreateDatasetVersionRequest
	GetDescription() *string
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
	// The type of the data source. If you specify multiple types, separate them with commas (,). Valid values:
	//
	// - NAS: The data is stored in Alibaba Cloud File Storage (NAS).
	//
	// - OSS: The data is stored in Alibaba Cloud Object Storage Service (OSS).
	//
	// - CPFS
	//
	// Note: The DataSourceType of the version must be the same as the DataSourceType of the dataset. The system verifies this consistency when you create the version.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// A custom description for the dataset version. This helps distinguish different dataset versions.
	//
	// example:
	//
	// This is a description of the dataset version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The storage import configuration of the dataset. Supported storage types include OSS, NAS, and CPFS.
	//
	// <details>
	//
	// <summary>
	//
	// OSS
	//
	// </summary>
	//
	// {<br>
	//
	// "region": "${region}",// The region ID.<br>
	//
	// "bucket": "${bucket}",// The bucket name.<br>
	//
	// "path": "${path}" // The file path.<br>
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
	// {<br>
	//
	// "region": "${region}",// The region ID.<br>
	//
	// "fileSystemId": "${file_system_id}", // The file system ID.<br>
	//
	// "path": "${path}", // The file system path.<br>
	//
	// "mountTarget": "${mount_target}" // The mount target of the file system.<br>
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
	// {<br>
	//
	// "region": "${region}",// The region ID.<br>
	//
	// "fileSystemId": "${file_system_id}", // The file system ID.<br>
	//
	// "protocolServiceId":"${protocol_service_id}", // The protocol service of the file system.<br>
	//
	// "exportId": "${export_id}", // The exported directory of the file system.<br>
	//
	// "path": "${path}", // The file system path.<br>
	//
	// }
	//
	// </details>
	//
	// <details>
	//
	// <summary>
	//
	// Intelligent Computing CPFS
	//
	// </summary>
	//
	// {<br>
	//
	// "region": "${region}",// The region ID.<br>
	//
	// "fileSystemId": "${file_system_id}", // The file system ID.<br>
	//
	// "path": "${path}", // The file system path.<br>
	//
	// "mountTarget": "${mount_target}", // The mount target of the file system. This parameter is specific to the Intelligent Computing edition.<br>
	//
	// "isVpcMount": boolean, // Specifies whether the mount target is in a VPC. This parameter is specific to the Intelligent Computing edition.<br>
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
	// A list of tags for the dataset version.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The extended field, which is a JSON string.
	//
	// When DLC uses the dataset, you can configure the mountPath field to specify the default mount path for the dataset.
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
	// - DIRECTORY: A folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The ID of the data source.
	//
	// - If SourceType is set to USER, you can customize the SourceId.
	//
	// - If SourceType is set to ITAG, which indicates a dataset generated from the annotation results of the iTAG module, SourceId is the task ID from iTAG.
	//
	// - If SourceType is set to PAI_PUBLIC_DATASET, which indicates a dataset created from a public PAI dataset, SourceId is empty by default.
	//
	// example:
	//
	// d-a0xbe5n03bhqof46ce
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The type of the data source. The default value is USER. Valid values:
	//
	// - PAI-PUBLIC-DATASET: a public dataset from PAI.
	//
	// - ITAG: a dataset generated from the annotation results of the iTAG module.
	//
	// - USER: a dataset registered by a user.
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The following examples show how to configure the URI:
	//
	// - If the data source type is OSS: `oss://bucket.endpoint/object`
	//
	// - If the data source type is NAS:
	//
	//   The format for a general-purpose NAS file system is `nas://<nasfisid>.region/subpath/to/dir/`.
	//
	//   CPFS 1.0: `nas://<cpfs-fsid>.region/subpath/to/dir/`.
	//
	//   CPFS 2.0: `nas://<cpfs-fsid>.region/<protocolserviceid>/`.
	//
	//   CPFS 1.0 and CPFS 2.0 are distinguished by the format of the fsid. The format for CPFS 1.0 is cpfs-<8 ASCII characters>. The format for CPFS 2.0 is cpfs-<16 ASCII characters>.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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

func (s *CreateDatasetVersionRequest) GetDescription() *string {
	return s.Description
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

func (s *CreateDatasetVersionRequest) SetDescription(v string) *CreateDatasetVersionRequest {
	s.Description = &v
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
	return nil
}
