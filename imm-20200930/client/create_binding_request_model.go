// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateBindingRequest
	GetDatasetName() *string
	SetProjectName(v string) *CreateBindingRequest
	GetProjectName() *string
	SetURI(v string) *CreateBindingRequest
	GetURI() *string
}

type CreateBindingRequest struct {
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
	// The URI of the OSS bucket to which you bind the dataset.
	//
	// Specify the value in the oss://${Bucket} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://examplebucket
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBindingRequest) GoString() string {
	return s.String()
}

func (s *CreateBindingRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateBindingRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateBindingRequest) GetURI() *string {
	return s.URI
}

func (s *CreateBindingRequest) SetDatasetName(v string) *CreateBindingRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateBindingRequest) SetProjectName(v string) *CreateBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateBindingRequest) SetURI(v string) *CreateBindingRequest {
	s.URI = &v
	return s
}

func (s *CreateBindingRequest) Validate() error {
	return dara.Validate(s)
}
