// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v *Address) *CreateStoryRequest
	GetAddress() *Address
	SetCustomLabels(v map[string]*string) *CreateStoryRequest
	GetCustomLabels() map[string]*string
	SetDriveId(v string) *CreateStoryRequest
	GetDriveId() *string
	SetMaxImageCount(v int64) *CreateStoryRequest
	GetMaxImageCount() *int64
	SetMinImageCount(v int64) *CreateStoryRequest
	GetMinImageCount() *int64
	SetStoryEndTime(v string) *CreateStoryRequest
	GetStoryEndTime() *string
	SetStoryId(v string) *CreateStoryRequest
	GetStoryId() *string
	SetStoryName(v string) *CreateStoryRequest
	GetStoryName() *string
	SetStoryStartTime(v string) *CreateStoryRequest
	GetStoryStartTime() *string
	SetStorySubType(v string) *CreateStoryRequest
	GetStorySubType() *string
	SetStoryType(v string) *CreateStoryRequest
	GetStoryType() *string
}

type CreateStoryRequest struct {
	Address *Address `json:"address,omitempty" xml:"address,omitempty"`
	// Deprecated
	CustomLabels map[string]*string `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 30
	MaxImageCount *int64 `json:"max_image_count,omitempty" xml:"max_image_count,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	MinImageCount *int64 `json:"min_image_count,omitempty" xml:"min_image_count,omitempty"`
	// example:
	//
	// 2022-12-30T16:00:00Z
	StoryEndTime *string `json:"story_end_time,omitempty" xml:"story_end_time,omitempty"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId   *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
	StoryName *string `json:"story_name,omitempty" xml:"story_name,omitempty"`
	// example:
	//
	// 2016-12-30T16:00:00Z
	StoryStartTime *string `json:"story_start_time,omitempty" xml:"story_start_time,omitempty"`
	// example:
	//
	// Food
	StorySubType *string `json:"story_sub_type,omitempty" xml:"story_sub_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TagMemory
	StoryType *string `json:"story_type,omitempty" xml:"story_type,omitempty"`
}

func (s CreateStoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStoryRequest) GoString() string {
	return s.String()
}

func (s *CreateStoryRequest) GetAddress() *Address {
	return s.Address
}

func (s *CreateStoryRequest) GetCustomLabels() map[string]*string {
	return s.CustomLabels
}

func (s *CreateStoryRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *CreateStoryRequest) GetMaxImageCount() *int64 {
	return s.MaxImageCount
}

func (s *CreateStoryRequest) GetMinImageCount() *int64 {
	return s.MinImageCount
}

func (s *CreateStoryRequest) GetStoryEndTime() *string {
	return s.StoryEndTime
}

func (s *CreateStoryRequest) GetStoryId() *string {
	return s.StoryId
}

func (s *CreateStoryRequest) GetStoryName() *string {
	return s.StoryName
}

func (s *CreateStoryRequest) GetStoryStartTime() *string {
	return s.StoryStartTime
}

func (s *CreateStoryRequest) GetStorySubType() *string {
	return s.StorySubType
}

func (s *CreateStoryRequest) GetStoryType() *string {
	return s.StoryType
}

func (s *CreateStoryRequest) SetAddress(v *Address) *CreateStoryRequest {
	s.Address = v
	return s
}

func (s *CreateStoryRequest) SetCustomLabels(v map[string]*string) *CreateStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *CreateStoryRequest) SetDriveId(v string) *CreateStoryRequest {
	s.DriveId = &v
	return s
}

func (s *CreateStoryRequest) SetMaxImageCount(v int64) *CreateStoryRequest {
	s.MaxImageCount = &v
	return s
}

func (s *CreateStoryRequest) SetMinImageCount(v int64) *CreateStoryRequest {
	s.MinImageCount = &v
	return s
}

func (s *CreateStoryRequest) SetStoryEndTime(v string) *CreateStoryRequest {
	s.StoryEndTime = &v
	return s
}

func (s *CreateStoryRequest) SetStoryId(v string) *CreateStoryRequest {
	s.StoryId = &v
	return s
}

func (s *CreateStoryRequest) SetStoryName(v string) *CreateStoryRequest {
	s.StoryName = &v
	return s
}

func (s *CreateStoryRequest) SetStoryStartTime(v string) *CreateStoryRequest {
	s.StoryStartTime = &v
	return s
}

func (s *CreateStoryRequest) SetStorySubType(v string) *CreateStoryRequest {
	s.StorySubType = &v
	return s
}

func (s *CreateStoryRequest) SetStoryType(v string) *CreateStoryRequest {
	s.StoryType = &v
	return s
}

func (s *CreateStoryRequest) Validate() error {
	if s.Address != nil {
		if err := s.Address.Validate(); err != nil {
			return err
		}
	}
	return nil
}
