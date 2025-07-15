// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColorWordsShrink(v string) *AsyncCreateClipsTaskShrinkRequest
	GetColorWordsShrink() *string
	SetHeight(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetHeight() *int32
	SetMusicUrl(v string) *AsyncCreateClipsTaskShrinkRequest
	GetMusicUrl() *string
	SetMusicVolume(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetMusicVolume() *int32
	SetSubtitleFontSize(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetSubtitleFontSize() *int32
	SetTaskId(v string) *AsyncCreateClipsTaskShrinkRequest
	GetTaskId() *string
	SetVoiceStyle(v string) *AsyncCreateClipsTaskShrinkRequest
	GetVoiceStyle() *string
	SetVoiceVolume(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetVoiceVolume() *int32
	SetWidth(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetWidth() *int32
	SetWorkspaceId(v string) *AsyncCreateClipsTaskShrinkRequest
	GetWorkspaceId() *string
}

type AsyncCreateClipsTaskShrinkRequest struct {
	ColorWordsShrink *string `json:"ColorWords,omitempty" xml:"ColorWords,omitempty"`
	// example:
	//
	// 1920
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// http://music.mp4
	MusicUrl         *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
	MusicVolume      *int32  `json:"MusicVolume,omitempty" xml:"MusicVolume,omitempty"`
	SubtitleFontSize *int32  `json:"SubtitleFontSize,omitempty" xml:"SubtitleFontSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 17a299a9-f223-4707-b0dd-4c22519bddf5
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VoiceStyle  *string `json:"VoiceStyle,omitempty" xml:"VoiceStyle,omitempty"`
	VoiceVolume *int32  `json:"VoiceVolume,omitempty" xml:"VoiceVolume,omitempty"`
	// example:
	//
	// 1080
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-ipe7d81yq4sl5jmk
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncCreateClipsTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetColorWordsShrink() *string {
	return s.ColorWordsShrink
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetHeight() *int32 {
	return s.Height
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetMusicVolume() *int32 {
	return s.MusicVolume
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetSubtitleFontSize() *int32 {
	return s.SubtitleFontSize
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetVoiceStyle() *string {
	return s.VoiceStyle
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetVoiceVolume() *int32 {
	return s.VoiceVolume
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetWidth() *int32 {
	return s.Width
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetColorWordsShrink(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.ColorWordsShrink = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetHeight(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.Height = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetMusicUrl(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.MusicUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetMusicVolume(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.MusicVolume = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetSubtitleFontSize(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.SubtitleFontSize = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetTaskId(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetVoiceStyle(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.VoiceStyle = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetVoiceVolume(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.VoiceVolume = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetWidth(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.Width = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetWorkspaceId(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
