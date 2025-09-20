// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFileMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *BatchGetFileMetaRequest
	GetDatasetName() *string
	SetProjectName(v string) *BatchGetFileMetaRequest
	GetProjectName() *string
	SetURIs(v []*string) *BatchGetFileMetaRequest
	GetURIs() []*string
	SetWithFields(v []*string) *BatchGetFileMetaRequest
	GetWithFields() []*string
}

type BatchGetFileMetaRequest struct {
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
	URIs []*string `json:"URIs,omitempty" xml:"URIs,omitempty" type:"Repeated"`
	// The fields to return. If you specify this parameter, only specified metadata fields are returned. You can use this parameter to control the size of the response.
	//
	// If you do not specify this parameter or leave this parameter empty, the operation returns all metadata fields.
	WithFields []*string `json:"WithFields,omitempty" xml:"WithFields,omitempty" type:"Repeated"`
}

func (s BatchGetFileMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFileMetaRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *BatchGetFileMetaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchGetFileMetaRequest) GetURIs() []*string {
	return s.URIs
}

func (s *BatchGetFileMetaRequest) GetWithFields() []*string {
	return s.WithFields
}

func (s *BatchGetFileMetaRequest) SetDatasetName(v string) *BatchGetFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchGetFileMetaRequest) SetProjectName(v string) *BatchGetFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchGetFileMetaRequest) SetURIs(v []*string) *BatchGetFileMetaRequest {
	s.URIs = v
	return s
}

func (s *BatchGetFileMetaRequest) SetWithFields(v []*string) *BatchGetFileMetaRequest {
	s.WithFields = v
	return s
}

func (s *BatchGetFileMetaRequest) Validate() error {
	return dara.Validate(s)
}
