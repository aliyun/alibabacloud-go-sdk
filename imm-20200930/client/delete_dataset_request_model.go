// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *DeleteDatasetRequest
	GetDatasetName() *string
	SetProjectName(v string) *DeleteDatasetRequest
	GetProjectName() *string
}

type DeleteDatasetRequest struct {
	// The name of the dataset. For information about how to create a dataset, see [CreateDataset](https://help.aliyun.com/document_detail/478160.html).
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
}

func (s DeleteDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *DeleteDatasetRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteDatasetRequest) SetDatasetName(v string) *DeleteDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *DeleteDatasetRequest) SetProjectName(v string) *DeleteDatasetRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteDatasetRequest) Validate() error {
	return dara.Validate(s)
}
