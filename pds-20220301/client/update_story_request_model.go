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
	SetCustomLabels(v map[string]*string) *UpdateStoryRequest
	GetCustomLabels() map[string]*string
	SetDriveId(v string) *UpdateStoryRequest
	GetDriveId() *string
	SetStoryId(v string) *UpdateStoryRequest
	GetStoryId() *string
	SetStoryName(v string) *UpdateStoryRequest
	GetStoryName() *string
}

type UpdateStoryRequest struct {
	// if can be null:
	// true
	Cover *UpdateStoryRequestCover `json:"cover,omitempty" xml:"cover,omitempty" type:"Struct"`
	// Deprecated
	CustomLabels map[string]*string `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// name1
	StoryName *string `json:"story_name,omitempty" xml:"story_name,omitempty"`
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

func (s *UpdateStoryRequest) GetCustomLabels() map[string]*string {
	return s.CustomLabels
}

func (s *UpdateStoryRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *UpdateStoryRequest) GetStoryId() *string {
	return s.StoryId
}

func (s *UpdateStoryRequest) GetStoryName() *string {
	return s.StoryName
}

func (s *UpdateStoryRequest) SetCover(v *UpdateStoryRequestCover) *UpdateStoryRequest {
	s.Cover = v
	return s
}

func (s *UpdateStoryRequest) SetCustomLabels(v map[string]*string) *UpdateStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *UpdateStoryRequest) SetDriveId(v string) *UpdateStoryRequest {
	s.DriveId = &v
	return s
}

func (s *UpdateStoryRequest) SetStoryId(v string) *UpdateStoryRequest {
	s.StoryId = &v
	return s
}

func (s *UpdateStoryRequest) SetStoryName(v string) *UpdateStoryRequest {
	s.StoryName = &v
	return s
}

func (s *UpdateStoryRequest) Validate() error {
	if s.Cover != nil {
		if err := s.Cover.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStoryRequestCover struct {
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 642a88dd06e49d9c0a14411ebae606f70edd9a59
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s UpdateStoryRequestCover) String() string {
	return dara.Prettify(s)
}

func (s UpdateStoryRequestCover) GoString() string {
	return s.String()
}

func (s *UpdateStoryRequestCover) GetFileId() *string {
	return s.FileId
}

func (s *UpdateStoryRequestCover) GetRevisionId() *string {
	return s.RevisionId
}

func (s *UpdateStoryRequestCover) SetFileId(v string) *UpdateStoryRequestCover {
	s.FileId = &v
	return s
}

func (s *UpdateStoryRequestCover) SetRevisionId(v string) *UpdateStoryRequestCover {
	s.RevisionId = &v
	return s
}

func (s *UpdateStoryRequestCover) Validate() error {
	return dara.Validate(s)
}
