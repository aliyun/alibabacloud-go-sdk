// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveStoryFilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *RemoveStoryFilesShrinkRequest
	GetDatasetName() *string
	SetFilesShrink(v string) *RemoveStoryFilesShrinkRequest
	GetFilesShrink() *string
	SetObjectId(v string) *RemoveStoryFilesShrinkRequest
	GetObjectId() *string
	SetProjectName(v string) *RemoveStoryFilesShrinkRequest
	GetProjectName() *string
}

type RemoveStoryFilesShrinkRequest struct {
	// The name of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The files that you want to delete.
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
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s RemoveStoryFilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveStoryFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *RemoveStoryFilesShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *RemoveStoryFilesShrinkRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *RemoveStoryFilesShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *RemoveStoryFilesShrinkRequest) SetDatasetName(v string) *RemoveStoryFilesShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *RemoveStoryFilesShrinkRequest) SetFilesShrink(v string) *RemoveStoryFilesShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *RemoveStoryFilesShrinkRequest) SetObjectId(v string) *RemoveStoryFilesShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *RemoveStoryFilesShrinkRequest) SetProjectName(v string) *RemoveStoryFilesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *RemoveStoryFilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
