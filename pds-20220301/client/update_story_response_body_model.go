// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *UpdateStoryResponseBody
	GetDriveId() *string
	SetStoryId(v string) *UpdateStoryResponseBody
	GetStoryId() *string
}

type UpdateStoryResponseBody struct {
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s UpdateStoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStoryResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *UpdateStoryResponseBody) GetStoryId() *string {
	return s.StoryId
}

func (s *UpdateStoryResponseBody) SetDriveId(v string) *UpdateStoryResponseBody {
	s.DriveId = &v
	return s
}

func (s *UpdateStoryResponseBody) SetStoryId(v string) *UpdateStoryResponseBody {
	s.StoryId = &v
	return s
}

func (s *UpdateStoryResponseBody) Validate() error {
	return dara.Validate(s)
}
