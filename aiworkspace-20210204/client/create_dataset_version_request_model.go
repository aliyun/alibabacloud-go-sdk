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
	// The number of dataset files.
	//
	// example:
	//
	// 300
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// The size of space occupied by dataset files. Unit: bytes.
	//
	// example:
	//
	// 19000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The data source type. Separate multiple values with commas (,). Valid values:
	//
	// - NAS: Alibaba Cloud Network Attached Storage (NAS).
	//
	// - OSS: Alibaba Cloud Object Storage Service (OSS).
	//
	// - CPFS
	//
	// > The DataSourceType of the version must be consistent with the DataSourceType of the dataset. Validation is performed against the dataset when a version is created.
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
	// The custom description of the dataset version, used to distinguish different dataset versions.
	//
	// example:
	//
	// This is a description of the dataset version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The storage import configuration of the dataset. OSS, NAS, and CPFS are supported.
	//
	// <details>
	//
	// <summary>OSS</summary>
	//
	// {<BR>
	//
	//   "region": "${region}",//Region ID<BR>
	//
	//   "bucket": "${bucket}",//Bucket name<BR>
	//
	//   "path": "${path}" //File path<BR>
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
	//   "region": "${region}",//Region ID<BR>
	//
	//   "fileSystemId": "${file_system_id}", //File system ID<BR>
	//
	//   "path": "${path}", //File system path<BR>
	//
	//   "mountTarget": "${mount_target}" //File system mount target<BR>
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
	//   "region": "${region}",//Region ID<BR>
	//
	//   "fileSystemId": "${file_system_id}", //File system ID<BR>
	//
	//   "protocolServiceId":"${protocol_service_id}", //File system protocol service<BR>
	//
	//   "exportId": "${export_id}", //File system export directory<BR>
	//
	//   "path": "${path}",  //File system path<BR>
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
	//   "region": "${region}",//Region ID<BR>
	//
	//   "fileSystemId": "${file_system_id}", //File system ID<BR>
	//
	//   "path": "${path}",  //File system path<BR>
	//
	//   "mountTarget": "${mount_target}" //File system mount target, specific to Lingjun edition<BR>
	//
	//   "isVpcMount": boolean, //Whether it is a VPC mount target, specific to Lingjun edition<BR>
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
	// The list of dataset version labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The extended field in JsonString format.
	//
	// When DLC uses a dataset, you can specify the default mount path of the dataset by configuring the mountPath field.
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
	// d-a0xbe5n03bhqof46ce
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The data source type. Default value: USER. Valid values:
	//
	// - PAI-PUBLIC-DATASET: PAI public dataset.
	//
	// - ITAG: dataset generated from iTAG annotation results.
	//
	// - USER: user-registered dataset.
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
	// General-purpose NAS format: `nas://<nasfisid>.region/subpath/to/dir/`;
	//
	// CPFS 1.0: `nas://<cpfs-fsid>.region/subpath/to/dir/`;
	//
	// CPFS 2.0: `nas://<cpfs-fsid>.region/<protocolserviceid>/`.
	//
	// CPFS 1.0 and CPFS 2.0 are distinguished by the format of the fsid: CPFS 1.0 format is cpfs-<8 ASCII characters>; CPFS 2.0 format is cpfs-<16 ASCII characters>.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/
	Uri                  *string                `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
