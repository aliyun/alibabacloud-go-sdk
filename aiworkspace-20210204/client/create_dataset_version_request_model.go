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
	// The number of dataset files.
	//
	// example:
	//
	// 300
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// The size of the dataset file. Unit: bytes.
	//
	// example:
	//
	// 19000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The type of the data source. Separate multiple types with commas (,). Valid values:
	//
	// 	- NAS: File Storage NAS (NAS).
	//
	// 	- OSS: Object Storage Service (OSS).
	//
	// 	- CPFS
	//
	// Note: The DataSourceType value of a dataset version must be the same as that of the dataset. When you create a dataset version, the system checks whether the values are the same.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The description of the dataset. Descriptions are used to differentiate datasets.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The dataset storage import configurations, such as OSS, NAS, and CPFS.
	//
	// **OSS**
	//
	// {\\
	//
	// "region": "${region}",// The region ID\\
	//
	// "bucket": "${bucket}",//The bucket name\\
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
	// "path": "${path}", // The ile system path\\
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
	// The tags of the dataset version.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The extended field, which is of the JsonString type. When you use the dataset in Deep Learning Containers (DLC), you can use the mountPath field to specify the default mount path of the dataset.
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
	// The ID of the data source.
	//
	// 	- If SourceType is set to USER, the value of SourceId can be a custom string.
	//
	// 	- If SourceType is set to ITAG, the value of SourceId is the ID of the labeling job of iTAG.
	//
	// 	- If SourceType is set to PAI_PUBLIC_DATASET, SourceId is empty by default.
	//
	// example:
	//
	// d-a0xbe5n03bhqof46ce
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The type of the data source. Default value: USER. Valid values:
	//
	// 	- PAI-PUBLIC-DATASET: a public dataset of Platform for AI (PAI).
	//
	// 	- ITAG: a dataset generated from a labeling job of iTAG.
	//
	// 	- USER: a dataset registered by a user.
	//
	// For each job type:
	//
	// 	- PAI_PUBLIC_DATASET: PAI_PUBLIC_DATASET.
	//
	// 	- ITAG: ITAG.
	//
	// 	- USER: USER.
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Example format:
	//
	// 	- Value format when DataSourceType is set to OSS: `oss://bucket.endpoint/object`.
	//
	// 	- Value formats when DataSourceType is set to NAS: General-purpose NAS: `nas://<nasfisid>.region/subpath/to/dir/`. CPFS 1.0: `nas://<cpfs-fsid>.region/subpath/to/dir/`. CPFS 2.0: `nas://<cpfs-fsid>.region/<protocolserviceid>/`. You can distinguish CPFS 1.0 and CPFS 2.0 file systems based on the format of the file system ID: The ID for CPFS 1.0 is in the cpfs-<8-bit ASCII characters> format. The ID for CPFS 2.0 is in the cpfs-<16-bit ASCII characters> format.
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
