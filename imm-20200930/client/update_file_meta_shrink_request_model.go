// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *UpdateFileMetaShrinkRequest
	GetDatasetName() *string
	SetFileShrink(v string) *UpdateFileMetaShrinkRequest
	GetFileShrink() *string
	SetProjectName(v string) *UpdateFileMetaShrinkRequest
	GetProjectName() *string
}

type UpdateFileMetaShrinkRequest struct {
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
	FileShrink *string `json:"File,omitempty" xml:"File,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s UpdateFileMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileMetaShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateFileMetaShrinkRequest) GetFileShrink() *string {
	return s.FileShrink
}

func (s *UpdateFileMetaShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateFileMetaShrinkRequest) SetDatasetName(v string) *UpdateFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateFileMetaShrinkRequest) SetFileShrink(v string) *UpdateFileMetaShrinkRequest {
	s.FileShrink = &v
	return s
}

func (s *UpdateFileMetaShrinkRequest) SetProjectName(v string) *UpdateFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateFileMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
