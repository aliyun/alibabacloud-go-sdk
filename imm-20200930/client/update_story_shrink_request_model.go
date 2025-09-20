// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverShrink(v string) *UpdateStoryShrinkRequest
	GetCoverShrink() *string
	SetCustomId(v string) *UpdateStoryShrinkRequest
	GetCustomId() *string
	SetCustomLabelsShrink(v string) *UpdateStoryShrinkRequest
	GetCustomLabelsShrink() *string
	SetDatasetName(v string) *UpdateStoryShrinkRequest
	GetDatasetName() *string
	SetObjectId(v string) *UpdateStoryShrinkRequest
	GetObjectId() *string
	SetProjectName(v string) *UpdateStoryShrinkRequest
	GetProjectName() *string
	SetStoryName(v string) *UpdateStoryShrinkRequest
	GetStoryName() *string
}

type UpdateStoryShrinkRequest struct {
	// The cover image of the story.
	CoverShrink *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// The custom ID.
	//
	// example:
	//
	// test
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// The custom tags. You can specify up to 100 custom tags.
	//
	// example:
	//
	// {"key": "value"}
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdata
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
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
	// The name of the story.
	//
	// example:
	//
	// newstory
	StoryName *string `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
}

func (s UpdateStoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStoryShrinkRequest) GetCoverShrink() *string {
	return s.CoverShrink
}

func (s *UpdateStoryShrinkRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *UpdateStoryShrinkRequest) GetCustomLabelsShrink() *string {
	return s.CustomLabelsShrink
}

func (s *UpdateStoryShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateStoryShrinkRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *UpdateStoryShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateStoryShrinkRequest) GetStoryName() *string {
	return s.StoryName
}

func (s *UpdateStoryShrinkRequest) SetCoverShrink(v string) *UpdateStoryShrinkRequest {
	s.CoverShrink = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetCustomId(v string) *UpdateStoryShrinkRequest {
	s.CustomId = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetCustomLabelsShrink(v string) *UpdateStoryShrinkRequest {
	s.CustomLabelsShrink = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetDatasetName(v string) *UpdateStoryShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetObjectId(v string) *UpdateStoryShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetProjectName(v string) *UpdateStoryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetStoryName(v string) *UpdateStoryShrinkRequest {
	s.StoryName = &v
	return s
}

func (s *UpdateStoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
