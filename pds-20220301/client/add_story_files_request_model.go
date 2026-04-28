// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStoryFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *AddStoryFilesRequest
	GetDriveId() *string
	SetFiles(v []*AddStoryFilesRequestFiles) *AddStoryFilesRequest
	GetFiles() []*AddStoryFilesRequestFiles
	SetStoryId(v string) *AddStoryFilesRequest
	GetStoryId() *string
}

type AddStoryFilesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string                      `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Files   []*AddStoryFilesRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s AddStoryFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddStoryFilesRequest) GoString() string {
	return s.String()
}

func (s *AddStoryFilesRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *AddStoryFilesRequest) GetFiles() []*AddStoryFilesRequestFiles {
	return s.Files
}

func (s *AddStoryFilesRequest) GetStoryId() *string {
	return s.StoryId
}

func (s *AddStoryFilesRequest) SetDriveId(v string) *AddStoryFilesRequest {
	s.DriveId = &v
	return s
}

func (s *AddStoryFilesRequest) SetFiles(v []*AddStoryFilesRequestFiles) *AddStoryFilesRequest {
	s.Files = v
	return s
}

func (s *AddStoryFilesRequest) SetStoryId(v string) *AddStoryFilesRequest {
	s.StoryId = &v
	return s
}

func (s *AddStoryFilesRequest) Validate() error {
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

type AddStoryFilesRequestFiles struct {
	// This parameter is required.
	//
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 642a88dd06e49d9c0a14411ebae606f70edd9a59
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s AddStoryFilesRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s AddStoryFilesRequestFiles) GoString() string {
	return s.String()
}

func (s *AddStoryFilesRequestFiles) GetFileId() *string {
	return s.FileId
}

func (s *AddStoryFilesRequestFiles) GetRevisionId() *string {
	return s.RevisionId
}

func (s *AddStoryFilesRequestFiles) SetFileId(v string) *AddStoryFilesRequestFiles {
	s.FileId = &v
	return s
}

func (s *AddStoryFilesRequestFiles) SetRevisionId(v string) *AddStoryFilesRequestFiles {
	s.RevisionId = &v
	return s
}

func (s *AddStoryFilesRequestFiles) Validate() error {
	return dara.Validate(s)
}
