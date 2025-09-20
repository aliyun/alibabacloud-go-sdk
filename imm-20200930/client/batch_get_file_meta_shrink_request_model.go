// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFileMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *BatchGetFileMetaShrinkRequest
	GetDatasetName() *string
	SetProjectName(v string) *BatchGetFileMetaShrinkRequest
	GetProjectName() *string
	SetURIsShrink(v string) *BatchGetFileMetaShrinkRequest
	GetURIsShrink() *string
	SetWithFieldsShrink(v string) *BatchGetFileMetaShrinkRequest
	GetWithFieldsShrink() *string
}

type BatchGetFileMetaShrinkRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The array of object URIs. You can specify up to 100 object URIs in an array.
	//
	// This parameter is required.
	URIsShrink *string `json:"URIs,omitempty" xml:"URIs,omitempty"`
	// The fields to return. If you specify this parameter, only specified metadata fields are returned. You can use this parameter to control the size of the response.
	//
	// If you do not specify this parameter or leave this parameter empty, the operation returns all metadata fields.
	WithFieldsShrink *string `json:"WithFields,omitempty" xml:"WithFields,omitempty"`
}

func (s BatchGetFileMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFileMetaShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *BatchGetFileMetaShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchGetFileMetaShrinkRequest) GetURIsShrink() *string {
	return s.URIsShrink
}

func (s *BatchGetFileMetaShrinkRequest) GetWithFieldsShrink() *string {
	return s.WithFieldsShrink
}

func (s *BatchGetFileMetaShrinkRequest) SetDatasetName(v string) *BatchGetFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchGetFileMetaShrinkRequest) SetProjectName(v string) *BatchGetFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchGetFileMetaShrinkRequest) SetURIsShrink(v string) *BatchGetFileMetaShrinkRequest {
	s.URIsShrink = &v
	return s
}

func (s *BatchGetFileMetaShrinkRequest) SetWithFieldsShrink(v string) *BatchGetFileMetaShrinkRequest {
	s.WithFieldsShrink = &v
	return s
}

func (s *BatchGetFileMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
