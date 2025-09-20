// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *GetBindingRequest
	GetDatasetName() *string
	SetProjectName(v string) *GetBindingRequest
	GetProjectName() *string
	SetURI(v string) *GetBindingRequest
	GetURI() *string
}

type GetBindingRequest struct {
	// The name of the dataset. You can obtain the name of the dataset from the response of the [CreateDataset](https://help.aliyun.com/document_detail/478160.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
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

func (s GetBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBindingRequest) GoString() string {
	return s.String()
}

func (s *GetBindingRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetBindingRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetBindingRequest) GetURI() *string {
	return s.URI
}

func (s *GetBindingRequest) SetDatasetName(v string) *GetBindingRequest {
	s.DatasetName = &v
	return s
}

func (s *GetBindingRequest) SetProjectName(v string) *GetBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *GetBindingRequest) SetURI(v string) *GetBindingRequest {
	s.URI = &v
	return s
}

func (s *GetBindingRequest) Validate() error {
	return dara.Validate(s)
}
