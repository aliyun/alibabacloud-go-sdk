// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *DeleteStoryRequest
	GetDriveId() *string
	SetStoryId(v string) *DeleteStoryRequest
	GetStoryId() *string
}

type DeleteStoryRequest struct {
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
}

func (s DeleteStoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteStoryRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *DeleteStoryRequest) GetStoryId() *string {
	return s.StoryId
}

func (s *DeleteStoryRequest) SetDriveId(v string) *DeleteStoryRequest {
	s.DriveId = &v
	return s
}

func (s *DeleteStoryRequest) SetStoryId(v string) *DeleteStoryRequest {
	s.StoryId = &v
	return s
}

func (s *DeleteStoryRequest) Validate() error {
	return dara.Validate(s)
}
