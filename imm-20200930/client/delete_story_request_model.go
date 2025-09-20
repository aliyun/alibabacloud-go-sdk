// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *DeleteStoryRequest
	GetDatasetName() *string
	SetObjectId(v string) *DeleteStoryRequest
	GetObjectId() *string
	SetProjectName(v string) *DeleteStoryRequest
	GetProjectName() *string
}

type DeleteStoryRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The ID of the story to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// id1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteStoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteStoryRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *DeleteStoryRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *DeleteStoryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteStoryRequest) SetDatasetName(v string) *DeleteStoryRequest {
	s.DatasetName = &v
	return s
}

func (s *DeleteStoryRequest) SetObjectId(v string) *DeleteStoryRequest {
	s.ObjectId = &v
	return s
}

func (s *DeleteStoryRequest) SetProjectName(v string) *DeleteStoryRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteStoryRequest) Validate() error {
	return dara.Validate(s)
}
