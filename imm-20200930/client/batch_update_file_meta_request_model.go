// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFileMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *BatchUpdateFileMetaRequest
	GetDatasetName() *string
	SetFiles(v []*InputFile) *BatchUpdateFileMetaRequest
	GetFiles() []*InputFile
	SetProjectName(v string) *BatchUpdateFileMetaRequest
	GetProjectName() *string
}

type BatchUpdateFileMetaRequest struct {
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
	Files []*InputFile `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s BatchUpdateFileMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *BatchUpdateFileMetaRequest) GetFiles() []*InputFile {
	return s.Files
}

func (s *BatchUpdateFileMetaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchUpdateFileMetaRequest) SetDatasetName(v string) *BatchUpdateFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchUpdateFileMetaRequest) SetFiles(v []*InputFile) *BatchUpdateFileMetaRequest {
	s.Files = v
	return s
}

func (s *BatchUpdateFileMetaRequest) SetProjectName(v string) *BatchUpdateFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchUpdateFileMetaRequest) Validate() error {
	return dara.Validate(s)
}
