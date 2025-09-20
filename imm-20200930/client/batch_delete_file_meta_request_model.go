// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteFileMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *BatchDeleteFileMetaRequest
	GetDatasetName() *string
	SetProjectName(v string) *BatchDeleteFileMetaRequest
	GetProjectName() *string
	SetURIs(v []*string) *BatchDeleteFileMetaRequest
	GetURIs() []*string
}

type BatchDeleteFileMetaRequest struct {
	// The name of the dataset. You can obtain the name of the dataset from the response of the [CreateDataset](https://help.aliyun.com/document_detail/478160.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URIs of the OSS buckets in which the files whose metadata you want to delete are stored. You can specify up to 100 URIs.
	//
	// This parameter is required.
	URIs []*string `json:"URIs,omitempty" xml:"URIs,omitempty" type:"Repeated"`
}

func (s BatchDeleteFileMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteFileMetaRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *BatchDeleteFileMetaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchDeleteFileMetaRequest) GetURIs() []*string {
	return s.URIs
}

func (s *BatchDeleteFileMetaRequest) SetDatasetName(v string) *BatchDeleteFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchDeleteFileMetaRequest) SetProjectName(v string) *BatchDeleteFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchDeleteFileMetaRequest) SetURIs(v []*string) *BatchDeleteFileMetaRequest {
	s.URIs = v
	return s
}

func (s *BatchDeleteFileMetaRequest) Validate() error {
	return dara.Validate(s)
}
