// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStoryFilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *AddStoryFilesShrinkRequest
	GetDatasetName() *string
	SetFilesShrink(v string) *AddStoryFilesShrinkRequest
	GetFilesShrink() *string
	SetObjectId(v string) *AddStoryFilesShrinkRequest
	GetObjectId() *string
	SetProjectName(v string) *AddStoryFilesShrinkRequest
	GetProjectName() *string
}

type AddStoryFilesShrinkRequest struct {
	// The name of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The objects that you want to add.
	//
	// This parameter is required.
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	// The ID of the story.
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s AddStoryFilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddStoryFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddStoryFilesShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *AddStoryFilesShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *AddStoryFilesShrinkRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *AddStoryFilesShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *AddStoryFilesShrinkRequest) SetDatasetName(v string) *AddStoryFilesShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *AddStoryFilesShrinkRequest) SetFilesShrink(v string) *AddStoryFilesShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *AddStoryFilesShrinkRequest) SetObjectId(v string) *AddStoryFilesShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *AddStoryFilesShrinkRequest) SetProjectName(v string) *AddStoryFilesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *AddStoryFilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
