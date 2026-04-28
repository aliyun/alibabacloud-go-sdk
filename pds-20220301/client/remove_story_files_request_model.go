// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveStoryFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *RemoveStoryFilesRequest
	GetDriveId() *string
	SetFiles(v []*RemoveStoryFilesRequestFiles) *RemoveStoryFilesRequest
	GetFiles() []*RemoveStoryFilesRequestFiles
	SetStoryId(v string) *RemoveStoryFilesRequest
	GetStoryId() *string
}

type RemoveStoryFilesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string                         `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Files   []*RemoveStoryFilesRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s RemoveStoryFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveStoryFilesRequest) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *RemoveStoryFilesRequest) GetFiles() []*RemoveStoryFilesRequestFiles {
	return s.Files
}

func (s *RemoveStoryFilesRequest) GetStoryId() *string {
	return s.StoryId
}

func (s *RemoveStoryFilesRequest) SetDriveId(v string) *RemoveStoryFilesRequest {
	s.DriveId = &v
	return s
}

func (s *RemoveStoryFilesRequest) SetFiles(v []*RemoveStoryFilesRequestFiles) *RemoveStoryFilesRequest {
	s.Files = v
	return s
}

func (s *RemoveStoryFilesRequest) SetStoryId(v string) *RemoveStoryFilesRequest {
	s.StoryId = &v
	return s
}

func (s *RemoveStoryFilesRequest) Validate() error {
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

type RemoveStoryFilesRequestFiles struct {
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

func (s RemoveStoryFilesRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s RemoveStoryFilesRequestFiles) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesRequestFiles) GetFileId() *string {
	return s.FileId
}

func (s *RemoveStoryFilesRequestFiles) GetRevisionId() *string {
	return s.RevisionId
}

func (s *RemoveStoryFilesRequestFiles) SetFileId(v string) *RemoveStoryFilesRequestFiles {
	s.FileId = &v
	return s
}

func (s *RemoveStoryFilesRequestFiles) SetRevisionId(v string) *RemoveStoryFilesRequestFiles {
	s.RevisionId = &v
	return s
}

func (s *RemoveStoryFilesRequestFiles) Validate() error {
	return dara.Validate(s)
}
