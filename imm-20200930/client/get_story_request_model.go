// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *GetStoryRequest
	GetDatasetName() *string
	SetObjectId(v string) *GetStoryRequest
	GetObjectId() *string
	SetProjectName(v string) *GetStoryRequest
	GetProjectName() *string
}

type GetStoryRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The ID of the story.
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
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetStoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStoryRequest) GoString() string {
	return s.String()
}

func (s *GetStoryRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetStoryRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetStoryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetStoryRequest) SetDatasetName(v string) *GetStoryRequest {
	s.DatasetName = &v
	return s
}

func (s *GetStoryRequest) SetObjectId(v string) *GetStoryRequest {
	s.ObjectId = &v
	return s
}

func (s *GetStoryRequest) SetProjectName(v string) *GetStoryRequest {
	s.ProjectName = &v
	return s
}

func (s *GetStoryRequest) Validate() error {
	return dara.Validate(s)
}
