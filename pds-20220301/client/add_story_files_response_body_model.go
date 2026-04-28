// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStoryFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *AddStoryFilesResponseBody
	GetDriveId() *string
	SetFiles(v []*AddStoryFile) *AddStoryFilesResponseBody
	GetFiles() []*AddStoryFile
	SetRequestId(v string) *AddStoryFilesResponseBody
	GetRequestId() *string
	SetStoryId(v string) *AddStoryFilesResponseBody
	GetStoryId() *string
}

type AddStoryFilesResponseBody struct {
	// example:
	//
	// 1
	DriveId   *string         `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Files     []*AddStoryFile `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	RequestId *string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s AddStoryFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddStoryFilesResponseBody) GoString() string {
	return s.String()
}

func (s *AddStoryFilesResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *AddStoryFilesResponseBody) GetFiles() []*AddStoryFile {
	return s.Files
}

func (s *AddStoryFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddStoryFilesResponseBody) GetStoryId() *string {
	return s.StoryId
}

func (s *AddStoryFilesResponseBody) SetDriveId(v string) *AddStoryFilesResponseBody {
	s.DriveId = &v
	return s
}

func (s *AddStoryFilesResponseBody) SetFiles(v []*AddStoryFile) *AddStoryFilesResponseBody {
	s.Files = v
	return s
}

func (s *AddStoryFilesResponseBody) SetRequestId(v string) *AddStoryFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddStoryFilesResponseBody) SetStoryId(v string) *AddStoryFilesResponseBody {
	s.StoryId = &v
	return s
}

func (s *AddStoryFilesResponseBody) Validate() error {
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
