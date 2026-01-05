// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateDatasetVersionShrinkRequest
	GetComment() *string
	SetDatasetId(v string) *CreateDatasetVersionShrinkRequest
	GetDatasetId() *string
	SetImportInfoShrink(v string) *CreateDatasetVersionShrinkRequest
	GetImportInfoShrink() *string
	SetMountPath(v string) *CreateDatasetVersionShrinkRequest
	GetMountPath() *string
	SetUrl(v string) *CreateDatasetVersionShrinkRequest
	GetUrl() *string
}

type CreateDatasetVersionShrinkRequest struct {
	// The description for this dataset version. Maximum length: 1,024 characters.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The dataset ID. Currently supports DataWorks datasets only.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataworks-dataset:3pXXXb8o0ngr07njhps1
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The storage import configuration for the dataset. Required configuration varies by storage type.
	//
	// **NAS**
	//
	// For valid values, see the response from the file storage API DescribeFileSystems.
	//
	// ```JSON
	//
	// {
	//
	// "fileSystemId": "3b6XXX89c9", // The file system ID.
	//
	// "fileSystemStorageType":  "Performance" // The file system storage type.
	//
	// "vpcId": "vpc-uf66oxxxrqge1t2gson7s" // The VPC ID for the mount point.
	//
	// }
	//
	// ```
	ImportInfoShrink *string `json:"ImportInfo,omitempty" xml:"ImportInfo,omitempty"`
	// The mount path, which must start with /mnt/. Default value: /mnt/data.
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

func (s CreateDatasetVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDatasetVersionShrinkRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateDatasetVersionShrinkRequest) GetImportInfoShrink() *string {
	return s.ImportInfoShrink
}

func (s *CreateDatasetVersionShrinkRequest) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateDatasetVersionShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateDatasetVersionShrinkRequest) SetComment(v string) *CreateDatasetVersionShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDatasetVersionShrinkRequest) SetDatasetId(v string) *CreateDatasetVersionShrinkRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateDatasetVersionShrinkRequest) SetImportInfoShrink(v string) *CreateDatasetVersionShrinkRequest {
	s.ImportInfoShrink = &v
	return s
}

func (s *CreateDatasetVersionShrinkRequest) SetMountPath(v string) *CreateDatasetVersionShrinkRequest {
	s.MountPath = &v
	return s
}

func (s *CreateDatasetVersionShrinkRequest) SetUrl(v string) *CreateDatasetVersionShrinkRequest {
	s.Url = &v
	return s
}

func (s *CreateDatasetVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
