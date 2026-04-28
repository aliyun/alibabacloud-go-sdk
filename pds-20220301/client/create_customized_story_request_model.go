// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedStoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomLabels(v map[string]*string) *CreateCustomizedStoryRequest
	GetCustomLabels() map[string]*string
	SetDriveId(v string) *CreateCustomizedStoryRequest
	GetDriveId() *string
	SetStoryCover(v *CreateCustomizedStoryRequestStoryCover) *CreateCustomizedStoryRequest
	GetStoryCover() *CreateCustomizedStoryRequestStoryCover
	SetStoryFiles(v []*CreateCustomizedStoryRequestStoryFiles) *CreateCustomizedStoryRequest
	GetStoryFiles() []*CreateCustomizedStoryRequestStoryFiles
	SetStoryName(v string) *CreateCustomizedStoryRequest
	GetStoryName() *string
	SetStorySubType(v string) *CreateCustomizedStoryRequest
	GetStorySubType() *string
	SetStoryType(v string) *CreateCustomizedStoryRequest
	GetStoryType() *string
}

type CreateCustomizedStoryRequest struct {
	// Deprecated
	CustomLabels map[string]*string `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	StoryCover *CreateCustomizedStoryRequestStoryCover `json:"story_cover,omitempty" xml:"story_cover,omitempty" type:"Struct"`
	// This parameter is required.
	StoryFiles []*CreateCustomizedStoryRequestStoryFiles `json:"story_files,omitempty" xml:"story_files,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test_name
	StoryName *string `json:"story_name,omitempty" xml:"story_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_created
	StorySubType *string `json:"story_sub_type,omitempty" xml:"story_sub_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_created
	StoryType *string `json:"story_type,omitempty" xml:"story_type,omitempty"`
}

func (s CreateCustomizedStoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedStoryRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequest) GetCustomLabels() map[string]*string {
	return s.CustomLabels
}

func (s *CreateCustomizedStoryRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *CreateCustomizedStoryRequest) GetStoryCover() *CreateCustomizedStoryRequestStoryCover {
	return s.StoryCover
}

func (s *CreateCustomizedStoryRequest) GetStoryFiles() []*CreateCustomizedStoryRequestStoryFiles {
	return s.StoryFiles
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

func (s *CreateCustomizedStoryRequest) SetCustomLabels(v map[string]*string) *CreateCustomizedStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *CreateCustomizedStoryRequest) SetDriveId(v string) *CreateCustomizedStoryRequest {
	s.DriveId = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStoryCover(v *CreateCustomizedStoryRequestStoryCover) *CreateCustomizedStoryRequest {
	s.StoryCover = v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStoryFiles(v []*CreateCustomizedStoryRequestStoryFiles) *CreateCustomizedStoryRequest {
	s.StoryFiles = v
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
	if s.StoryCover != nil {
		if err := s.StoryCover.Validate(); err != nil {
			return err
		}
	}
	if s.StoryFiles != nil {
		for _, item := range s.StoryFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCustomizedStoryRequestStoryCover struct {
	// This parameter is required.
	//
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 642a88d4aff041ee68fd4fc89beb80e1119da343
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s CreateCustomizedStoryRequestStoryCover) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedStoryRequestStoryCover) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequestStoryCover) GetFileId() *string {
	return s.FileId
}

func (s *CreateCustomizedStoryRequestStoryCover) GetRevisionId() *string {
	return s.RevisionId
}

func (s *CreateCustomizedStoryRequestStoryCover) SetFileId(v string) *CreateCustomizedStoryRequestStoryCover {
	s.FileId = &v
	return s
}

func (s *CreateCustomizedStoryRequestStoryCover) SetRevisionId(v string) *CreateCustomizedStoryRequestStoryCover {
	s.RevisionId = &v
	return s
}

func (s *CreateCustomizedStoryRequestStoryCover) Validate() error {
	return dara.Validate(s)
}

type CreateCustomizedStoryRequestStoryFiles struct {
	// This parameter is required.
	//
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 642a88d4aff041ee68fd4fc89beb80e1119da343
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s CreateCustomizedStoryRequestStoryFiles) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedStoryRequestStoryFiles) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequestStoryFiles) GetFileId() *string {
	return s.FileId
}

func (s *CreateCustomizedStoryRequestStoryFiles) GetRevisionId() *string {
	return s.RevisionId
}

func (s *CreateCustomizedStoryRequestStoryFiles) SetFileId(v string) *CreateCustomizedStoryRequestStoryFiles {
	s.FileId = &v
	return s
}

func (s *CreateCustomizedStoryRequestStoryFiles) SetRevisionId(v string) *CreateCustomizedStoryRequestStoryFiles {
	s.RevisionId = &v
	return s
}

func (s *CreateCustomizedStoryRequestStoryFiles) Validate() error {
	return dara.Validate(s)
}
