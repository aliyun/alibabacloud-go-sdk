// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFigureClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *UpdateFigureClusterRequest
	GetDatasetName() *string
	SetFigureCluster(v *FigureClusterForReq) *UpdateFigureClusterRequest
	GetFigureCluster() *FigureClusterForReq
	SetProjectName(v string) *UpdateFigureClusterRequest
	GetProjectName() *string
}

type UpdateFigureClusterRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The information about the cluster.
	//
	// This parameter is required.
	FigureCluster *FigureClusterForReq `json:"FigureCluster,omitempty" xml:"FigureCluster,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s UpdateFigureClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFigureClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateFigureClusterRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateFigureClusterRequest) GetFigureCluster() *FigureClusterForReq {
	return s.FigureCluster
}

func (s *UpdateFigureClusterRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateFigureClusterRequest) SetDatasetName(v string) *UpdateFigureClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateFigureClusterRequest) SetFigureCluster(v *FigureClusterForReq) *UpdateFigureClusterRequest {
	s.FigureCluster = v
	return s
}

func (s *UpdateFigureClusterRequest) SetProjectName(v string) *UpdateFigureClusterRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateFigureClusterRequest) Validate() error {
	return dara.Validate(s)
}
