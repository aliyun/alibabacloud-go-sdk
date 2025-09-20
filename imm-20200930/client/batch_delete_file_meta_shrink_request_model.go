// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteFileMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *BatchDeleteFileMetaShrinkRequest
	GetDatasetName() *string
	SetProjectName(v string) *BatchDeleteFileMetaShrinkRequest
	GetProjectName() *string
	SetURIsShrink(v string) *BatchDeleteFileMetaShrinkRequest
	GetURIsShrink() *string
}

type BatchDeleteFileMetaShrinkRequest struct {
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
	URIsShrink *string `json:"URIs,omitempty" xml:"URIs,omitempty"`
}

func (s BatchDeleteFileMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteFileMetaShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *BatchDeleteFileMetaShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchDeleteFileMetaShrinkRequest) GetURIsShrink() *string {
	return s.URIsShrink
}

func (s *BatchDeleteFileMetaShrinkRequest) SetDatasetName(v string) *BatchDeleteFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchDeleteFileMetaShrinkRequest) SetProjectName(v string) *BatchDeleteFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchDeleteFileMetaShrinkRequest) SetURIsShrink(v string) *BatchDeleteFileMetaShrinkRequest {
	s.URIsShrink = &v
	return s
}

func (s *BatchDeleteFileMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
