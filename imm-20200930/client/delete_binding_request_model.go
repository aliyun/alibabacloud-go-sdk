// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *DeleteBindingRequest
	GetDatasetName() *string
	SetProjectName(v string) *DeleteBindingRequest
	GetProjectName() *string
	SetURI(v string) *DeleteBindingRequest
	GetURI() *string
}

type DeleteBindingRequest struct {
	// The name of the dataset. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URI of the OSS bucket to which the dataset is bound.
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

func (s DeleteBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBindingRequest) GoString() string {
	return s.String()
}

func (s *DeleteBindingRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *DeleteBindingRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteBindingRequest) GetURI() *string {
	return s.URI
}

func (s *DeleteBindingRequest) SetDatasetName(v string) *DeleteBindingRequest {
	s.DatasetName = &v
	return s
}

func (s *DeleteBindingRequest) SetProjectName(v string) *DeleteBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteBindingRequest) SetURI(v string) *DeleteBindingRequest {
	s.URI = &v
	return s
}

func (s *DeleteBindingRequest) Validate() error {
	return dara.Validate(s)
}
