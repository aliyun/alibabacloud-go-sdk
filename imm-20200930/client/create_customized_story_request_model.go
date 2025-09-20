// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedStoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCover(v *CreateCustomizedStoryRequestCover) *CreateCustomizedStoryRequest
	GetCover() *CreateCustomizedStoryRequestCover
	SetCustomLabels(v map[string]interface{}) *CreateCustomizedStoryRequest
	GetCustomLabels() map[string]interface{}
	SetDatasetName(v string) *CreateCustomizedStoryRequest
	GetDatasetName() *string
	SetFiles(v []*CreateCustomizedStoryRequestFiles) *CreateCustomizedStoryRequest
	GetFiles() []*CreateCustomizedStoryRequestFiles
	SetProjectName(v string) *CreateCustomizedStoryRequest
	GetProjectName() *string
	SetStoryName(v string) *CreateCustomizedStoryRequest
	GetStoryName() *string
	SetStorySubType(v string) *CreateCustomizedStoryRequest
	GetStorySubType() *string
	SetStoryType(v string) *CreateCustomizedStoryRequest
	GetStoryType() *string
}

type CreateCustomizedStoryRequest struct {
	// The cover image of the story. You can specify an image as the cover image of the custom story.
	//
	// This parameter is required.
	Cover *CreateCustomizedStoryRequestCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	// The custom labels. You can specify labels to help you identify and retrieve the story.
	//
	// example:
	//
	// {"Bucket": "examplebucket"}
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
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
	Files []*CreateCustomizedStoryRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
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

func (s CreateCustomizedStoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedStoryRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequest) GetCover() *CreateCustomizedStoryRequestCover {
	return s.Cover
}

func (s *CreateCustomizedStoryRequest) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *CreateCustomizedStoryRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateCustomizedStoryRequest) GetFiles() []*CreateCustomizedStoryRequestFiles {
	return s.Files
}

func (s *CreateCustomizedStoryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateCustomizedStoryRequest) GetStoryName() *string {
	return s.StoryName
}

func (s *CreateCustomizedStoryRequest) GetStorySubType() *string {
	return s.StorySubType
}

func (s *CreateCustomizedStoryRequest) GetStoryType() *string {
	return s.StoryType
}

func (s *CreateCustomizedStoryRequest) SetCover(v *CreateCustomizedStoryRequestCover) *CreateCustomizedStoryRequest {
	s.Cover = v
	return s
}

func (s *CreateCustomizedStoryRequest) SetCustomLabels(v map[string]interface{}) *CreateCustomizedStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *CreateCustomizedStoryRequest) SetDatasetName(v string) *CreateCustomizedStoryRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetFiles(v []*CreateCustomizedStoryRequestFiles) *CreateCustomizedStoryRequest {
	s.Files = v
	return s
}

func (s *CreateCustomizedStoryRequest) SetProjectName(v string) *CreateCustomizedStoryRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStoryName(v string) *CreateCustomizedStoryRequest {
	s.StoryName = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStorySubType(v string) *CreateCustomizedStoryRequest {
	s.StorySubType = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStoryType(v string) *CreateCustomizedStoryRequest {
	s.StoryType = &v
	return s
}

func (s *CreateCustomizedStoryRequest) Validate() error {
	return dara.Validate(s)
}

type CreateCustomizedStoryRequestCover struct {
	// The URI of the cover image.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket1/cover
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateCustomizedStoryRequestCover) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedStoryRequestCover) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequestCover) GetURI() *string {
	return s.URI
}

func (s *CreateCustomizedStoryRequestCover) SetURI(v string) *CreateCustomizedStoryRequestCover {
	s.URI = &v
	return s
}

func (s *CreateCustomizedStoryRequestCover) Validate() error {
	return dara.Validate(s)
}

type CreateCustomizedStoryRequestFiles struct {
	// The URIs of the files.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"URI":"oss://bucket1/file1"}]
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateCustomizedStoryRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedStoryRequestFiles) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequestFiles) GetURI() *string {
	return s.URI
}

func (s *CreateCustomizedStoryRequestFiles) SetURI(v string) *CreateCustomizedStoryRequestFiles {
	s.URI = &v
	return s
}

func (s *CreateCustomizedStoryRequestFiles) Validate() error {
	return dara.Validate(s)
}
