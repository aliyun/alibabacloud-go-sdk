// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedStoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *CreateCustomizedStoryResponseBody
	GetDriveId() *string
	SetStoryId(v string) *CreateCustomizedStoryResponseBody
	GetStoryId() *string
}

type CreateCustomizedStoryResponseBody struct {
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s CreateCustomizedStoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedStoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *CreateCustomizedStoryResponseBody) GetStoryId() *string {
	return s.StoryId
}

func (s *CreateCustomizedStoryResponseBody) SetDriveId(v string) *CreateCustomizedStoryResponseBody {
	s.DriveId = &v
	return s
}

func (s *CreateCustomizedStoryResponseBody) SetStoryId(v string) *CreateCustomizedStoryResponseBody {
	s.StoryId = &v
	return s
}

func (s *CreateCustomizedStoryResponseBody) Validate() error {
	return dara.Validate(s)
}
