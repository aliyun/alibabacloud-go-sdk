// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveStoryFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *RemoveStoryFilesResponseBody
	GetDriveId() *string
	SetStoryId(v string) *RemoveStoryFilesResponseBody
	GetStoryId() *string
}

type RemoveStoryFilesResponseBody struct {
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s RemoveStoryFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveStoryFilesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *RemoveStoryFilesResponseBody) GetStoryId() *string {
	return s.StoryId
}

func (s *RemoveStoryFilesResponseBody) SetDriveId(v string) *RemoveStoryFilesResponseBody {
	s.DriveId = &v
	return s
}

func (s *RemoveStoryFilesResponseBody) SetStoryId(v string) *RemoveStoryFilesResponseBody {
	s.StoryId = &v
	return s
}

func (s *RemoveStoryFilesResponseBody) Validate() error {
	return dara.Validate(s)
}
