// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateDatasetVersionRequest
	GetComment() *string
	SetDatasetId(v string) *CreateDatasetVersionRequest
	GetDatasetId() *string
	SetImportInfo(v map[string]*string) *CreateDatasetVersionRequest
	GetImportInfo() map[string]*string
	SetMountPath(v string) *CreateDatasetVersionRequest
	GetMountPath() *string
	SetUrl(v string) *CreateDatasetVersionRequest
	GetUrl() *string
}

type CreateDatasetVersionRequest struct {
	// The description of the dataset version. The description can be up to 1024 characters in length.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The dataset ID. Currently, only DataWorks datasets are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataworks-dataset:3pXXXb8o0ngr07njhps1
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The storage import configuration for the dataset. The required configuration varies depending on the storage type.
	//
	// <details>
	//
	// <summary>NAS</summary>
	//
	// The values can be obtained from the response of the File Storage API DescribeFileSystems operation.
	//
	// ```JSON
	//
	// {
	//
	//   "fileSystemId": "3b6XXX89c9", // The file system ID.
	//
	//   "fileSystemStorageType": "Performance", // The storage specification of the file system.
	//
	//   "vpcId": "vpc-uf66oxxxrqge1t2gson7s" // The VPC ID of the mount target.
	//
	// }
	//
	// ```
	//
	// </details>
	ImportInfo map[string]*string `json:"ImportInfo,omitempty" xml:"ImportInfo,omitempty"`
	// The mount path. The path must start with /mnt/. Default value: /mnt/data.
	//
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-oss-bucket/test_dir/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDatasetVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDatasetVersionRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateDatasetVersionRequest) GetImportInfo() map[string]*string {
	return s.ImportInfo
}

func (s *CreateDatasetVersionRequest) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateDatasetVersionRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateDatasetVersionRequest) SetComment(v string) *CreateDatasetVersionRequest {
	s.Comment = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetDatasetId(v string) *CreateDatasetVersionRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetImportInfo(v map[string]*string) *CreateDatasetVersionRequest {
	s.ImportInfo = v
	return s
}

func (s *CreateDatasetVersionRequest) SetMountPath(v string) *CreateDatasetVersionRequest {
	s.MountPath = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetUrl(v string) *CreateDatasetVersionRequest {
	s.Url = &v
	return s
}

func (s *CreateDatasetVersionRequest) Validate() error {
	return dara.Validate(s)
}
