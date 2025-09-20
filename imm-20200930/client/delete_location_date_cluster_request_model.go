// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLocationDateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *DeleteLocationDateClusterRequest
	GetDatasetName() *string
	SetObjectId(v string) *DeleteLocationDateClusterRequest
	GetObjectId() *string
	SetProjectName(v string) *DeleteLocationDateClusterRequest
	GetProjectName() *string
}

type DeleteLocationDateClusterRequest struct {
	// The name of the dataset. For information about how to create a dataset, see [CreateDataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The ID of the group to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// location-date-cluster-71dd4f32-9597-4085-a2ab-3a7b0fd0aff9
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteLocationDateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLocationDateClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteLocationDateClusterRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *DeleteLocationDateClusterRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *DeleteLocationDateClusterRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteLocationDateClusterRequest) SetDatasetName(v string) *DeleteLocationDateClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *DeleteLocationDateClusterRequest) SetObjectId(v string) *DeleteLocationDateClusterRequest {
	s.ObjectId = &v
	return s
}

func (s *DeleteLocationDateClusterRequest) SetProjectName(v string) *DeleteLocationDateClusterRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteLocationDateClusterRequest) Validate() error {
	return dara.Validate(s)
}
