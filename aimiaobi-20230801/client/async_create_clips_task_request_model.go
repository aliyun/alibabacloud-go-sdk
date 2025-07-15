// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColorWords(v []*AsyncCreateClipsTaskRequestColorWords) *AsyncCreateClipsTaskRequest
	GetColorWords() []*AsyncCreateClipsTaskRequestColorWords
	SetHeight(v int32) *AsyncCreateClipsTaskRequest
	GetHeight() *int32
	SetMusicUrl(v string) *AsyncCreateClipsTaskRequest
	GetMusicUrl() *string
	SetMusicVolume(v int32) *AsyncCreateClipsTaskRequest
	GetMusicVolume() *int32
	SetSubtitleFontSize(v int32) *AsyncCreateClipsTaskRequest
	GetSubtitleFontSize() *int32
	SetTaskId(v string) *AsyncCreateClipsTaskRequest
	GetTaskId() *string
	SetVoiceStyle(v string) *AsyncCreateClipsTaskRequest
	GetVoiceStyle() *string
	SetVoiceVolume(v int32) *AsyncCreateClipsTaskRequest
	GetVoiceVolume() *int32
	SetWidth(v int32) *AsyncCreateClipsTaskRequest
	GetWidth() *int32
	SetWorkspaceId(v string) *AsyncCreateClipsTaskRequest
	GetWorkspaceId() *string
}

type AsyncCreateClipsTaskRequest struct {
	ColorWords []*AsyncCreateClipsTaskRequestColorWords `json:"ColorWords,omitempty" xml:"ColorWords,omitempty" type:"Repeated"`
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

func (s AsyncCreateClipsTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTaskRequest) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTaskRequest) GetColorWords() []*AsyncCreateClipsTaskRequestColorWords {
	return s.ColorWords
}

func (s *AsyncCreateClipsTaskRequest) GetHeight() *int32 {
	return s.Height
}

func (s *AsyncCreateClipsTaskRequest) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *AsyncCreateClipsTaskRequest) GetMusicVolume() *int32 {
	return s.MusicVolume
}

func (s *AsyncCreateClipsTaskRequest) GetSubtitleFontSize() *int32 {
	return s.SubtitleFontSize
}

func (s *AsyncCreateClipsTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncCreateClipsTaskRequest) GetVoiceStyle() *string {
	return s.VoiceStyle
}

func (s *AsyncCreateClipsTaskRequest) GetVoiceVolume() *int32 {
	return s.VoiceVolume
}

func (s *AsyncCreateClipsTaskRequest) GetWidth() *int32 {
	return s.Width
}

func (s *AsyncCreateClipsTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncCreateClipsTaskRequest) SetColorWords(v []*AsyncCreateClipsTaskRequestColorWords) *AsyncCreateClipsTaskRequest {
	s.ColorWords = v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetHeight(v int32) *AsyncCreateClipsTaskRequest {
	s.Height = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetMusicUrl(v string) *AsyncCreateClipsTaskRequest {
	s.MusicUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetMusicVolume(v int32) *AsyncCreateClipsTaskRequest {
	s.MusicVolume = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetSubtitleFontSize(v int32) *AsyncCreateClipsTaskRequest {
	s.SubtitleFontSize = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetTaskId(v string) *AsyncCreateClipsTaskRequest {
	s.TaskId = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetVoiceStyle(v string) *AsyncCreateClipsTaskRequest {
	s.VoiceStyle = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetVoiceVolume(v int32) *AsyncCreateClipsTaskRequest {
	s.VoiceVolume = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetWidth(v int32) *AsyncCreateClipsTaskRequest {
	s.Width = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetWorkspaceId(v string) *AsyncCreateClipsTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) Validate() error {
	return dara.Validate(s)
}

type AsyncCreateClipsTaskRequestColorWords struct {
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EffectColorStyle *string `json:"EffectColorStyle,omitempty" xml:"EffectColorStyle,omitempty"`
	FontSize         *int32  `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// example:
	//
	// 0
	TimelineIn *int32 `json:"TimelineIn,omitempty" xml:"TimelineIn,omitempty"`
	// example:
	//
	// 5
	TimelineOut *int32 `json:"TimelineOut,omitempty" xml:"TimelineOut,omitempty"`
	// example:
	//
	// 0.2
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.5
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s AsyncCreateClipsTaskRequestColorWords) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTaskRequestColorWords) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetContent() *string {
	return s.Content
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetEffectColorStyle() *string {
	return s.EffectColorStyle
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetFontSize() *int32 {
	return s.FontSize
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetTimelineIn() *int32 {
	return s.TimelineIn
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetTimelineOut() *int32 {
	return s.TimelineOut
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetX() *float32 {
	return s.X
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetY() *float32 {
	return s.Y
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetContent(v string) *AsyncCreateClipsTaskRequestColorWords {
	s.Content = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetEffectColorStyle(v string) *AsyncCreateClipsTaskRequestColorWords {
	s.EffectColorStyle = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetFontSize(v int32) *AsyncCreateClipsTaskRequestColorWords {
	s.FontSize = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetTimelineIn(v int32) *AsyncCreateClipsTaskRequestColorWords {
	s.TimelineIn = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetTimelineOut(v int32) *AsyncCreateClipsTaskRequestColorWords {
	s.TimelineOut = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetX(v float32) *AsyncCreateClipsTaskRequestColorWords {
	s.X = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetY(v float32) *AsyncCreateClipsTaskRequestColorWords {
	s.Y = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) Validate() error {
	return dara.Validate(s)
}
