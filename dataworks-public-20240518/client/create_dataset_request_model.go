// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateDatasetRequest
	GetComment() *string
	SetDataType(v string) *CreateDatasetRequest
	GetDataType() *string
	SetInitVersion(v *CreateDatasetRequestInitVersion) *CreateDatasetRequest
	GetInitVersion() *CreateDatasetRequestInitVersion
	SetName(v string) *CreateDatasetRequest
	GetName() *string
	SetOrigin(v string) *CreateDatasetRequest
	GetOrigin() *string
	SetProjectId(v int64) *CreateDatasetRequest
	GetProjectId() *int64
	SetStorageType(v string) *CreateDatasetRequest
	GetStorageType() *string
}

type CreateDatasetRequest struct {
	// The description of the dataset. It must not exceed 1,024 characters in length.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The data type. Valid values:
	//
	// 	- COMMON: Common (Default)
	//
	// 	- PIC
	//
	// 	- TEXT
	//
	// 	- TABLE
	//
	// 	- VIDEO
	//
	// 	- AUDIO
	//
	// 	- INDEX
	//
	// example:
	//
	// COMMON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The initial version of the dataset.
	//
	// This parameter is required.
	InitVersion *CreateDatasetRequestInitVersion `json:"InitVersion,omitempty" xml:"InitVersion,omitempty" type:"Struct"`
	// The name of the dataset. It cannot be an empty string and must not exceed 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_oss_dataset
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source of the dataset. Currently, only DataWorks is supported.
	//
	// example:
	//
	// DataWorks
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The storage type. Currently supported values:
	//
	// 	- OSS
	//
	// 	- NAS: General-purpose NAS file systems
	//
	// 	- EXTREMENAS: Extreme NAS file systems
	//
	// 	- DLF_LANCE: Data Lake Formation
	//
	// Valid values:
	//
	// 	- NAS: General-purpose NAS file systems
	//
	// 	- MAXCOMPUTE: MaxCompute table
	//
	// 	- CPFS: Cloud Parallel File Storage
	//
	// 	- BMCPFS: CPFS for Lingjun
	//
	// 	- EXTREMENAS: Extreme NAS file systems
	//
	// 	- OSS: Object Storage Service
	//
	// 	- DLF_LANCE: Data Lake Formation.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDatasetRequest) GetDataType() *string {
	return s.DataType
}

func (s *CreateDatasetRequest) GetInitVersion() *CreateDatasetRequestInitVersion {
	return s.InitVersion
}

func (s *CreateDatasetRequest) GetName() *string {
	return s.Name
}

func (s *CreateDatasetRequest) GetOrigin() *string {
	return s.Origin
}

func (s *CreateDatasetRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDatasetRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDatasetRequest) SetComment(v string) *CreateDatasetRequest {
	s.Comment = &v
	return s
}

func (s *CreateDatasetRequest) SetDataType(v string) *CreateDatasetRequest {
	s.DataType = &v
	return s
}

func (s *CreateDatasetRequest) SetInitVersion(v *CreateDatasetRequestInitVersion) *CreateDatasetRequest {
	s.InitVersion = v
	return s
}

func (s *CreateDatasetRequest) SetName(v string) *CreateDatasetRequest {
	s.Name = &v
	return s
}

func (s *CreateDatasetRequest) SetOrigin(v string) *CreateDatasetRequest {
	s.Origin = &v
	return s
}

func (s *CreateDatasetRequest) SetProjectId(v int64) *CreateDatasetRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDatasetRequest) SetStorageType(v string) *CreateDatasetRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDatasetRequest) Validate() error {
	if s.InitVersion != nil {
		if err := s.InitVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDatasetRequestInitVersion struct {
	// The description. It must not exceed 1,024 characters in length.
	//
	// example:
	//
	// Initial Version
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The storage import configuration for the dataset. The required configuration information varies by storage type.
	//
	// **NAS**
	//
	// For valid values, refer to the response of the file storage API DescribeFileSystems.
	//
	// ```JSON
	//
	// {
	//
	// "fileSystemId": "3b6XXX89c9", // The file system ID.
	//
	// "fileSystemStorageType":  "Performance" // The storage specification of the file system.
	//
	// "vpcId": "vpc-uf66oxxxrqge1t2gson7s" // The VPC ID of the mount point.
	//
	// }
	//
	// ```
	ImportInfo map[string]*string `json:"ImportInfo,omitempty" xml:"ImportInfo,omitempty"`
	// The mount path. It must start with /mnt/. Default value: /mnt/data.
	//
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// URL
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-oss-bucket/test_dir/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDatasetRequestInitVersion) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestInitVersion) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestInitVersion) GetComment() *string {
	return s.Comment
}

func (s *CreateDatasetRequestInitVersion) GetImportInfo() map[string]*string {
	return s.ImportInfo
}

func (s *CreateDatasetRequestInitVersion) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateDatasetRequestInitVersion) GetUrl() *string {
	return s.Url
}

func (s *CreateDatasetRequestInitVersion) SetComment(v string) *CreateDatasetRequestInitVersion {
	s.Comment = &v
	return s
}

func (s *CreateDatasetRequestInitVersion) SetImportInfo(v map[string]*string) *CreateDatasetRequestInitVersion {
	s.ImportInfo = v
	return s
}

func (s *CreateDatasetRequestInitVersion) SetMountPath(v string) *CreateDatasetRequestInitVersion {
	s.MountPath = &v
	return s
}

func (s *CreateDatasetRequestInitVersion) SetUrl(v string) *CreateDatasetRequestInitVersion {
	s.Url = &v
	return s
}

func (s *CreateDatasetRequestInitVersion) Validate() error {
	return dara.Validate(s)
}
