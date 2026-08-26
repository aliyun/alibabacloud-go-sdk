// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPlaylistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int32) *StartPlaylistRequest
	GetOffset() *int32
	SetOwnerId(v int64) *StartPlaylistRequest
	GetOwnerId() *int64
	SetProgramId(v string) *StartPlaylistRequest
	GetProgramId() *string
	SetRegionId(v string) *StartPlaylistRequest
	GetRegionId() *string
	SetResumeMode(v string) *StartPlaylistRequest
	GetResumeMode() *string
	SetStartItemId(v string) *StartPlaylistRequest
	GetStartItemId() *string
}

type StartPlaylistRequest struct {
	// The start offset for the video file. This parameter is valid only for video files. Unit: milliseconds.
	//
	// A value greater than 0 specifies the start time relative to the first frame.
	//
	// example:
	//
	// 10000
	Offset  *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the playlist. If you add items to the playlist by calling the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation, use the value of the ProgramId parameter that is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The restart mode. Valid values:
	//
	// - **Restart**: Starts from the beginning.
	//
	// - **Continue**: Resumes playback from where it was stopped. The **StartItemId*	- parameter is required only when you set the **ResumeMode*	- parameter to **Custom**.
	//
	// - **Custom**: Custom start point.
	//
	// example:
	//
	// Custom
	ResumeMode *string `json:"ResumeMode,omitempty" xml:"ResumeMode,omitempty"`
	// The ID of the item to play first. When the carousel starts, this item is played.
	//
	// 	Notice:
	//
	// This parameter is required only when you set **ResumeMode*	- to **Custom**.
	//
	// example:
	//
	// asdfasdfasdf****
	StartItemId *string `json:"StartItemId,omitempty" xml:"StartItemId,omitempty"`
}

func (s StartPlaylistRequest) String() string {
	return dara.Prettify(s)
}

func (s StartPlaylistRequest) GoString() string {
	return s.String()
}

func (s *StartPlaylistRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *StartPlaylistRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartPlaylistRequest) GetProgramId() *string {
	return s.ProgramId
}

func (s *StartPlaylistRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartPlaylistRequest) GetResumeMode() *string {
	return s.ResumeMode
}

func (s *StartPlaylistRequest) GetStartItemId() *string {
	return s.StartItemId
}

func (s *StartPlaylistRequest) SetOffset(v int32) *StartPlaylistRequest {
	s.Offset = &v
	return s
}

func (s *StartPlaylistRequest) SetOwnerId(v int64) *StartPlaylistRequest {
	s.OwnerId = &v
	return s
}

func (s *StartPlaylistRequest) SetProgramId(v string) *StartPlaylistRequest {
	s.ProgramId = &v
	return s
}

func (s *StartPlaylistRequest) SetRegionId(v string) *StartPlaylistRequest {
	s.RegionId = &v
	return s
}

func (s *StartPlaylistRequest) SetResumeMode(v string) *StartPlaylistRequest {
	s.ResumeMode = &v
	return s
}

func (s *StartPlaylistRequest) SetStartItemId(v string) *StartPlaylistRequest {
	s.StartItemId = &v
	return s
}

func (s *StartPlaylistRequest) Validate() error {
	return dara.Validate(s)
}
