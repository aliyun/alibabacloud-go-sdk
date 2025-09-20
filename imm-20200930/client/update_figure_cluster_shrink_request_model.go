// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFigureClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *UpdateFigureClusterShrinkRequest
	GetDatasetName() *string
	SetFigureClusterShrink(v string) *UpdateFigureClusterShrinkRequest
	GetFigureClusterShrink() *string
	SetProjectName(v string) *UpdateFigureClusterShrinkRequest
	GetProjectName() *string
}

type UpdateFigureClusterShrinkRequest struct {
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
	FigureClusterShrink *string `json:"FigureCluster,omitempty" xml:"FigureCluster,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s UpdateFigureClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFigureClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFigureClusterShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateFigureClusterShrinkRequest) GetFigureClusterShrink() *string {
	return s.FigureClusterShrink
}

func (s *UpdateFigureClusterShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateFigureClusterShrinkRequest) SetDatasetName(v string) *UpdateFigureClusterShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateFigureClusterShrinkRequest) SetFigureClusterShrink(v string) *UpdateFigureClusterShrinkRequest {
	s.FigureClusterShrink = &v
	return s
}

func (s *UpdateFigureClusterShrinkRequest) SetProjectName(v string) *UpdateFigureClusterShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateFigureClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
