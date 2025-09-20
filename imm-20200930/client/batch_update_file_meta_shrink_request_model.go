// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFileMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *BatchUpdateFileMetaShrinkRequest
	GetDatasetName() *string
	SetFilesShrink(v string) *BatchUpdateFileMetaShrinkRequest
	GetFilesShrink() *string
	SetProjectName(v string) *BatchUpdateFileMetaShrinkRequest
	GetProjectName() *string
}

type BatchUpdateFileMetaShrinkRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The files whose metadata you want to update.
	//
	// This parameter is required.
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s BatchUpdateFileMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *BatchUpdateFileMetaShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *BatchUpdateFileMetaShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchUpdateFileMetaShrinkRequest) SetDatasetName(v string) *BatchUpdateFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchUpdateFileMetaShrinkRequest) SetFilesShrink(v string) *BatchUpdateFileMetaShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *BatchUpdateFileMetaShrinkRequest) SetProjectName(v string) *BatchUpdateFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchUpdateFileMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
