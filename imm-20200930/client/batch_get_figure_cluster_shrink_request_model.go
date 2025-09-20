// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFigureClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *BatchGetFigureClusterShrinkRequest
	GetDatasetName() *string
	SetObjectIdsShrink(v string) *BatchGetFigureClusterShrinkRequest
	GetObjectIdsShrink() *string
	SetProjectName(v string) *BatchGetFigureClusterShrinkRequest
	GetProjectName() *string
}

type BatchGetFigureClusterShrinkRequest struct {
	// The name of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The cluster IDs.
	//
	// This parameter is required.
	ObjectIdsShrink *string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty"`
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s BatchGetFigureClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFigureClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFigureClusterShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *BatchGetFigureClusterShrinkRequest) GetObjectIdsShrink() *string {
	return s.ObjectIdsShrink
}

func (s *BatchGetFigureClusterShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchGetFigureClusterShrinkRequest) SetDatasetName(v string) *BatchGetFigureClusterShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchGetFigureClusterShrinkRequest) SetObjectIdsShrink(v string) *BatchGetFigureClusterShrinkRequest {
	s.ObjectIdsShrink = &v
	return s
}

func (s *BatchGetFigureClusterShrinkRequest) SetProjectName(v string) *BatchGetFigureClusterShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchGetFigureClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
