// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *GetFileMetaShrinkRequest
	GetDatasetName() *string
	SetProjectName(v string) *GetFileMetaShrinkRequest
	GetProjectName() *string
	SetURI(v string) *GetFileMetaShrinkRequest
	GetURI() *string
	SetWithFieldsShrink(v string) *GetFileMetaShrinkRequest
	GetWithFieldsShrink() *string
}

type GetFileMetaShrinkRequest struct {
	// The name of the dataset. For more information about how to obtain the dataset name, refer to [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The name of the project. For more information about how to obtain the project name, refer to [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URI of the file. Make sure that the file has been **indexed**.
	//
	// The OSS URI format is oss://${Bucket}/${Object}, where `${Bucket}` is the name of the OSS bucket that resides in the same region as the current project, and `${Object}` is the full path of the file including the file name extension.
	//
	// The PDS URI format is pds://domains/${domain}/drives/${drive}/files/${file}/revisions/${revision}.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// Specifies the specific fields to return, instead of all existing metadata fields. You can use this parameter to reduce the size of the returned struct.
	//
	// If you do not specify this parameter or leave it empty, all fields are returned.
	WithFieldsShrink *string `json:"WithFields,omitempty" xml:"WithFields,omitempty"`
}

func (s GetFileMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetFileMetaShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetFileMetaShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetFileMetaShrinkRequest) GetURI() *string {
	return s.URI
}

func (s *GetFileMetaShrinkRequest) GetWithFieldsShrink() *string {
	return s.WithFieldsShrink
}

func (s *GetFileMetaShrinkRequest) SetDatasetName(v string) *GetFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *GetFileMetaShrinkRequest) SetProjectName(v string) *GetFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *GetFileMetaShrinkRequest) SetURI(v string) *GetFileMetaShrinkRequest {
	s.URI = &v
	return s
}

func (s *GetFileMetaShrinkRequest) SetWithFieldsShrink(v string) *GetFileMetaShrinkRequest {
	s.WithFieldsShrink = &v
	return s
}

func (s *GetFileMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
