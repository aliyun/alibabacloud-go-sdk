// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *GetFileMetaRequest
	GetDatasetName() *string
	SetProjectName(v string) *GetFileMetaRequest
	GetProjectName() *string
	SetURI(v string) *GetFileMetaRequest
	GetURI() *string
	SetWithFields(v []*string) *GetFileMetaRequest
	GetWithFields() []*string
}

type GetFileMetaRequest struct {
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
	WithFields []*string `json:"WithFields,omitempty" xml:"WithFields,omitempty" type:"Repeated"`
}

func (s GetFileMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileMetaRequest) GoString() string {
	return s.String()
}

func (s *GetFileMetaRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetFileMetaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetFileMetaRequest) GetURI() *string {
	return s.URI
}

func (s *GetFileMetaRequest) GetWithFields() []*string {
	return s.WithFields
}

func (s *GetFileMetaRequest) SetDatasetName(v string) *GetFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *GetFileMetaRequest) SetProjectName(v string) *GetFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *GetFileMetaRequest) SetURI(v string) *GetFileMetaRequest {
	s.URI = &v
	return s
}

func (s *GetFileMetaRequest) SetWithFields(v []*string) *GetFileMetaRequest {
	s.WithFields = v
	return s
}

func (s *GetFileMetaRequest) Validate() error {
	return dara.Validate(s)
}
