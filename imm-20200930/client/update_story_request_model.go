// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCover(v *UpdateStoryRequestCover) *UpdateStoryRequest
	GetCover() *UpdateStoryRequestCover
	SetCustomId(v string) *UpdateStoryRequest
	GetCustomId() *string
	SetCustomLabels(v map[string]interface{}) *UpdateStoryRequest
	GetCustomLabels() map[string]interface{}
	SetDatasetName(v string) *UpdateStoryRequest
	GetDatasetName() *string
	SetObjectId(v string) *UpdateStoryRequest
	GetObjectId() *string
	SetProjectName(v string) *UpdateStoryRequest
	GetProjectName() *string
	SetStoryName(v string) *UpdateStoryRequest
	GetStoryName() *string
}

type UpdateStoryRequest struct {
	// The cover image of the story.
	Cover *UpdateStoryRequestCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
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
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
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

func (s UpdateStoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateStoryRequest) GetCover() *UpdateStoryRequestCover {
	return s.Cover
}

func (s *UpdateStoryRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *UpdateStoryRequest) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *UpdateStoryRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateStoryRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *UpdateStoryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateStoryRequest) GetStoryName() *string {
	return s.StoryName
}

func (s *UpdateStoryRequest) SetCover(v *UpdateStoryRequestCover) *UpdateStoryRequest {
	s.Cover = v
	return s
}

func (s *UpdateStoryRequest) SetCustomId(v string) *UpdateStoryRequest {
	s.CustomId = &v
	return s
}

func (s *UpdateStoryRequest) SetCustomLabels(v map[string]interface{}) *UpdateStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *UpdateStoryRequest) SetDatasetName(v string) *UpdateStoryRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateStoryRequest) SetObjectId(v string) *UpdateStoryRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateStoryRequest) SetProjectName(v string) *UpdateStoryRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateStoryRequest) SetStoryName(v string) *UpdateStoryRequest {
	s.StoryName = &v
	return s
}

func (s *UpdateStoryRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateStoryRequestCover struct {
	// The URI of the cover image.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://bucket1/object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s UpdateStoryRequestCover) String() string {
	return dara.Prettify(s)
}

func (s UpdateStoryRequestCover) GoString() string {
	return s.String()
}

func (s *UpdateStoryRequestCover) GetURI() *string {
	return s.URI
}

func (s *UpdateStoryRequestCover) SetURI(v string) *UpdateStoryRequestCover {
	s.URI = &v
	return s
}

func (s *UpdateStoryRequestCover) Validate() error {
	return dara.Validate(s)
}
