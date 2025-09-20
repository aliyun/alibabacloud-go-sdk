// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedStoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverShrink(v string) *CreateCustomizedStoryShrinkRequest
	GetCoverShrink() *string
	SetCustomLabelsShrink(v string) *CreateCustomizedStoryShrinkRequest
	GetCustomLabelsShrink() *string
	SetDatasetName(v string) *CreateCustomizedStoryShrinkRequest
	GetDatasetName() *string
	SetFilesShrink(v string) *CreateCustomizedStoryShrinkRequest
	GetFilesShrink() *string
	SetProjectName(v string) *CreateCustomizedStoryShrinkRequest
	GetProjectName() *string
	SetStoryName(v string) *CreateCustomizedStoryShrinkRequest
	GetStoryName() *string
	SetStorySubType(v string) *CreateCustomizedStoryShrinkRequest
	GetStorySubType() *string
	SetStoryType(v string) *CreateCustomizedStoryShrinkRequest
	GetStoryType() *string
}

type CreateCustomizedStoryShrinkRequest struct {
	// The cover image of the story. You can specify an image as the cover image of the custom story.
	//
	// This parameter is required.
	CoverShrink *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// The custom labels. You can specify labels to help you identify and retrieve the story.
	//
	// example:
	//
	// {"Bucket": "examplebucket"}
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The files of the story. You can specify up to 100 files in a custom story.
	//
	// This parameter is required.
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the story.
	//
	// This parameter is required.
	//
	// example:
	//
	// name1
	StoryName *string `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	// The subtype of the story. For information about valid subtypes, see [Story types and subtypes](https://help.aliyun.com/document_detail/2743998.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Solo
	StorySubType *string `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	// The type of the story. For information about valid types, see [Story types and subtypes](https://help.aliyun.com/document_detail/2743998.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// PeopleMemory
	StoryType *string `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
}

func (s CreateCustomizedStoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedStoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryShrinkRequest) GetCoverShrink() *string {
	return s.CoverShrink
}

func (s *CreateCustomizedStoryShrinkRequest) GetCustomLabelsShrink() *string {
	return s.CustomLabelsShrink
}

func (s *CreateCustomizedStoryShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateCustomizedStoryShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *CreateCustomizedStoryShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateCustomizedStoryShrinkRequest) GetStoryName() *string {
	return s.StoryName
}

func (s *CreateCustomizedStoryShrinkRequest) GetStorySubType() *string {
	return s.StorySubType
}

func (s *CreateCustomizedStoryShrinkRequest) GetStoryType() *string {
	return s.StoryType
}

func (s *CreateCustomizedStoryShrinkRequest) SetCoverShrink(v string) *CreateCustomizedStoryShrinkRequest {
	s.CoverShrink = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetCustomLabelsShrink(v string) *CreateCustomizedStoryShrinkRequest {
	s.CustomLabelsShrink = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetDatasetName(v string) *CreateCustomizedStoryShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetFilesShrink(v string) *CreateCustomizedStoryShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetProjectName(v string) *CreateCustomizedStoryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetStoryName(v string) *CreateCustomizedStoryShrinkRequest {
	s.StoryName = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetStorySubType(v string) *CreateCustomizedStoryShrinkRequest {
	s.StorySubType = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetStoryType(v string) *CreateCustomizedStoryShrinkRequest {
	s.StoryType = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
