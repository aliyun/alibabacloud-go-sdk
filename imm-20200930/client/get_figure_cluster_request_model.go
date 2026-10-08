// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFigureClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *GetFigureClusterRequest
	GetDatasetName() *string
	SetObjectId(v string) *GetFigureClusterRequest
	GetObjectId() *string
	SetProjectName(v string) *GetFigureClusterRequest
	GetProjectName() *string
}

type GetFigureClusterRequest struct {
	// The name of the dataset. For more information about how to obtain the dataset name, see [CreateDataset](~~CreateDataset~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The object ID of the clustering group. You can obtain the object ID from the face group information returned by [QueryFigureClusters](~~QueryFigureClusters~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// Cluster-1f2e1a2c-d5ee-4bc5-84f6-fef94ea****
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the project. For more information about how to obtain the project name, see [CreateProject](~~CreateProject~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetFigureClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFigureClusterRequest) GoString() string {
	return s.String()
}

func (s *GetFigureClusterRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetFigureClusterRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetFigureClusterRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetFigureClusterRequest) SetDatasetName(v string) *GetFigureClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *GetFigureClusterRequest) SetObjectId(v string) *GetFigureClusterRequest {
	s.ObjectId = &v
	return s
}

func (s *GetFigureClusterRequest) SetProjectName(v string) *GetFigureClusterRequest {
	s.ProjectName = &v
	return s
}

func (s *GetFigureClusterRequest) Validate() error {
	return dara.Validate(s)
}
