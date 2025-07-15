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
	// The offset of the position where the system starts the playback. This parameter takes effect only if the input source is a video file. Unit: milliseconds.
	//
	// A value greater than 0 indicates an offset from the first frame.
	//
	// example:
	//
	// 10000
	Offset  *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the episode list. If the episode list was created by calling the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation, check the value of the response parameter ProgramId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The method to resume the playback of the episode list. Valid values:
	//
	// 	- **Restart**: resumes the playback from the beginning.
	//
	// 	- **Continue**: resumes the playback from the position where the previous playback stops. The **StartItemId*	- parameter is required only if you set **ResumeMode*	- to **Custom**.
	//
	// 	- **Custom**: resumes the playback from a custom position.
	//
	// example:
	//
	// Custom
	ResumeMode *string `json:"ResumeMode,omitempty" xml:"ResumeMode,omitempty"`
	// The ID of the first episode to play. This episode is the first to play in carousel playback.
	//
	// >  This parameter is required only if you set ResumeMode to Custom.
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
