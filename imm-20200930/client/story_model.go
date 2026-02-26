// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStory interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*Address) *Story
	GetAddresses() []*Address
	SetCover(v *File) *Story
	GetCover() *File
	SetCreateTime(v string) *Story
	GetCreateTime() *string
	SetCustomId(v string) *Story
	GetCustomId() *string
	SetCustomLabels(v map[string]interface{}) *Story
	GetCustomLabels() map[string]interface{}
	SetDatasetName(v string) *Story
	GetDatasetName() *string
	SetFigureClusterIds(v []*string) *Story
	GetFigureClusterIds() []*string
	SetFiles(v []*File) *Story
	GetFiles() []*File
	SetObjectId(v string) *Story
	GetObjectId() *string
	SetObjectType(v string) *Story
	GetObjectType() *string
	SetOwnerId(v string) *Story
	GetOwnerId() *string
	SetProjectName(v string) *Story
	GetProjectName() *string
	SetStoryEndTime(v string) *Story
	GetStoryEndTime() *string
	SetStoryName(v string) *Story
	GetStoryName() *string
	SetStoryStartTime(v string) *Story
	GetStoryStartTime() *string
	SetStorySubType(v string) *Story
	GetStorySubType() *string
	SetStoryType(v string) *Story
	GetStoryType() *string
	SetUpdateTime(v string) *Story
	GetUpdateTime() *string
}

type Story struct {
	// The addresses.
	Addresses []*Address `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The story cover.
	Cover *File `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// The time when the story was created.
	//
	// example:
	//
	// 2022-01-14T10:17:18.102700407+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The custom ID.
	//
	// example:
	//
	// test
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// The custom labels.
	//
	// example:
	//
	// {"key": "val"}
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The fluster IDs.
	FigureClusterIds []*string `json:"FigureClusterIds,omitempty" xml:"FigureClusterIds,omitempty" type:"Repeated"`
	// The story files.
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The ID of the story object.
	//
	// example:
	//
	// id1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The type of the object.
	//
	// example:
	//
	// story
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The ID of the owner to which the story belongs.
	//
	// example:
	//
	// 102321002****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The time when the story ends.
	//
	// example:
	//
	// 2022-01-14T10:17:18.102700407+08:00
	StoryEndTime *string `json:"StoryEndTime,omitempty" xml:"StoryEndTime,omitempty"`
	// The name of the story.
	//
	// example:
	//
	// name1
	StoryName *string `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	// The time when the story starts.
	//
	// example:
	//
	// 2022-01-14T10:17:18.102700407+08:00
	StoryStartTime *string `json:"StoryStartTime,omitempty" xml:"StoryStartTime,omitempty"`
	// The subtype of the story.
	//
	// example:
	//
	// ImportantPerson
	StorySubType *string `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	// The story type.
	//
	// example:
	//
	// PeopleMemory
	StoryType *string `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
	// The time when the story was updated.
	//
	// example:
	//
	// 2022-01-14T10:17:18.102700407+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Story) String() string {
	return dara.Prettify(s)
}

func (s Story) GoString() string {
	return s.String()
}

func (s *Story) GetAddresses() []*Address {
	return s.Addresses
}

func (s *Story) GetCover() *File {
	return s.Cover
}

func (s *Story) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Story) GetCustomId() *string {
	return s.CustomId
}

func (s *Story) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *Story) GetDatasetName() *string {
	return s.DatasetName
}

func (s *Story) GetFigureClusterIds() []*string {
	return s.FigureClusterIds
}

func (s *Story) GetFiles() []*File {
	return s.Files
}

func (s *Story) GetObjectId() *string {
	return s.ObjectId
}

func (s *Story) GetObjectType() *string {
	return s.ObjectType
}

func (s *Story) GetOwnerId() *string {
	return s.OwnerId
}

func (s *Story) GetProjectName() *string {
	return s.ProjectName
}

func (s *Story) GetStoryEndTime() *string {
	return s.StoryEndTime
}

func (s *Story) GetStoryName() *string {
	return s.StoryName
}

func (s *Story) GetStoryStartTime() *string {
	return s.StoryStartTime
}

func (s *Story) GetStorySubType() *string {
	return s.StorySubType
}

func (s *Story) GetStoryType() *string {
	return s.StoryType
}

func (s *Story) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *Story) SetAddresses(v []*Address) *Story {
	s.Addresses = v
	return s
}

func (s *Story) SetCover(v *File) *Story {
	s.Cover = v
	return s
}

func (s *Story) SetCreateTime(v string) *Story {
	s.CreateTime = &v
	return s
}

func (s *Story) SetCustomId(v string) *Story {
	s.CustomId = &v
	return s
}

func (s *Story) SetCustomLabels(v map[string]interface{}) *Story {
	s.CustomLabels = v
	return s
}

func (s *Story) SetDatasetName(v string) *Story {
	s.DatasetName = &v
	return s
}

func (s *Story) SetFigureClusterIds(v []*string) *Story {
	s.FigureClusterIds = v
	return s
}

func (s *Story) SetFiles(v []*File) *Story {
	s.Files = v
	return s
}

func (s *Story) SetObjectId(v string) *Story {
	s.ObjectId = &v
	return s
}

func (s *Story) SetObjectType(v string) *Story {
	s.ObjectType = &v
	return s
}

func (s *Story) SetOwnerId(v string) *Story {
	s.OwnerId = &v
	return s
}

func (s *Story) SetProjectName(v string) *Story {
	s.ProjectName = &v
	return s
}

func (s *Story) SetStoryEndTime(v string) *Story {
	s.StoryEndTime = &v
	return s
}

func (s *Story) SetStoryName(v string) *Story {
	s.StoryName = &v
	return s
}

func (s *Story) SetStoryStartTime(v string) *Story {
	s.StoryStartTime = &v
	return s
}

func (s *Story) SetStorySubType(v string) *Story {
	s.StorySubType = &v
	return s
}

func (s *Story) SetStoryType(v string) *Story {
	s.StoryType = &v
	return s
}

func (s *Story) SetUpdateTime(v string) *Story {
	s.UpdateTime = &v
	return s
}

func (s *Story) Validate() error {
	if s.Addresses != nil {
		for _, item := range s.Addresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Cover != nil {
		if err := s.Cover.Validate(); err != nil {
			return err
		}
	}
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
