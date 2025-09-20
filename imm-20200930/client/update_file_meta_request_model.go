// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *UpdateFileMetaRequest
	GetDatasetName() *string
	SetFile(v *InputFile) *UpdateFileMetaRequest
	GetFile() *InputFile
	SetProjectName(v string) *UpdateFileMetaRequest
	GetProjectName() *string
}

type UpdateFileMetaRequest struct {
	// The name of the dataset. You can obtain the name of the dataset from the response of the [CreateDataset](https://help.aliyun.com/document_detail/478160.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The file and its metadata items to be updated. The value must be in the JSON format.
	//
	// This parameter is required.
	File *InputFile `json:"File,omitempty" xml:"File,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s UpdateFileMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileMetaRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateFileMetaRequest) GetFile() *InputFile {
	return s.File
}

func (s *UpdateFileMetaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateFileMetaRequest) SetDatasetName(v string) *UpdateFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateFileMetaRequest) SetFile(v *InputFile) *UpdateFileMetaRequest {
	s.File = v
	return s
}

func (s *UpdateFileMetaRequest) SetProjectName(v string) *UpdateFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateFileMetaRequest) Validate() error {
	return dara.Validate(s)
}
