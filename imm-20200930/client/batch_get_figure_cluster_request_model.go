// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFigureClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *BatchGetFigureClusterRequest
	GetDatasetName() *string
	SetObjectIds(v []*string) *BatchGetFigureClusterRequest
	GetObjectIds() []*string
	SetProjectName(v string) *BatchGetFigureClusterRequest
	GetProjectName() *string
}

type BatchGetFigureClusterRequest struct {
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
	ObjectIds []*string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty" type:"Repeated"`
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s BatchGetFigureClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFigureClusterRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFigureClusterRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *BatchGetFigureClusterRequest) GetObjectIds() []*string {
	return s.ObjectIds
}

func (s *BatchGetFigureClusterRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchGetFigureClusterRequest) SetDatasetName(v string) *BatchGetFigureClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchGetFigureClusterRequest) SetObjectIds(v []*string) *BatchGetFigureClusterRequest {
	s.ObjectIds = v
	return s
}

func (s *BatchGetFigureClusterRequest) SetProjectName(v string) *BatchGetFigureClusterRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchGetFigureClusterRequest) Validate() error {
	return dara.Validate(s)
}
