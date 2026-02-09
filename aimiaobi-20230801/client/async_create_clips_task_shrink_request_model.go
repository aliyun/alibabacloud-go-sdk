// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloseMusic(v bool) *AsyncCreateClipsTaskShrinkRequest
	GetCloseMusic() *bool
	SetCloseSubtitle(v bool) *AsyncCreateClipsTaskShrinkRequest
	GetCloseSubtitle() *bool
	SetCloseVoice(v bool) *AsyncCreateClipsTaskShrinkRequest
	GetCloseVoice() *bool
	SetClosingCreditsUrl(v string) *AsyncCreateClipsTaskShrinkRequest
	GetClosingCreditsUrl() *string
	SetColorWordsShrink(v string) *AsyncCreateClipsTaskShrinkRequest
	GetColorWordsShrink() *string
	SetCosyVoiceAppKey(v string) *AsyncCreateClipsTaskShrinkRequest
	GetCosyVoiceAppKey() *string
	SetCosyVoiceToken(v string) *AsyncCreateClipsTaskShrinkRequest
	GetCosyVoiceToken() *string
	SetCustomVoiceStyle(v string) *AsyncCreateClipsTaskShrinkRequest
	GetCustomVoiceStyle() *string
	SetCustomVoiceUrl(v string) *AsyncCreateClipsTaskShrinkRequest
	GetCustomVoiceUrl() *string
	SetCustomVoiceVolume(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetCustomVoiceVolume() *int32
	SetHeight(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetHeight() *int32
	SetHighDefSourceVideosShrink(v string) *AsyncCreateClipsTaskShrinkRequest
	GetHighDefSourceVideosShrink() *string
	SetMusicStyle(v string) *AsyncCreateClipsTaskShrinkRequest
	GetMusicStyle() *string
	SetMusicUrl(v string) *AsyncCreateClipsTaskShrinkRequest
	GetMusicUrl() *string
	SetMusicVolume(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetMusicVolume() *int32
	SetOpeningCreditsUrl(v string) *AsyncCreateClipsTaskShrinkRequest
	GetOpeningCreditsUrl() *string
	SetStickersShrink(v string) *AsyncCreateClipsTaskShrinkRequest
	GetStickersShrink() *string
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
	CloseMusic        *bool   `json:"CloseMusic,omitempty" xml:"CloseMusic,omitempty"`
	CloseSubtitle     *bool   `json:"CloseSubtitle,omitempty" xml:"CloseSubtitle,omitempty"`
	CloseVoice        *bool   `json:"CloseVoice,omitempty" xml:"CloseVoice,omitempty"`
	ClosingCreditsUrl *string `json:"ClosingCreditsUrl,omitempty" xml:"ClosingCreditsUrl,omitempty"`
	ColorWordsShrink  *string `json:"ColorWords,omitempty" xml:"ColorWords,omitempty"`
	CosyVoiceAppKey   *string `json:"CosyVoiceAppKey,omitempty" xml:"CosyVoiceAppKey,omitempty"`
	CosyVoiceToken    *string `json:"CosyVoiceToken,omitempty" xml:"CosyVoiceToken,omitempty"`
	CustomVoiceStyle  *string `json:"CustomVoiceStyle,omitempty" xml:"CustomVoiceStyle,omitempty"`
	// example:
	//
	// http://xxx/xxx.mp4
	CustomVoiceUrl *string `json:"CustomVoiceUrl,omitempty" xml:"CustomVoiceUrl,omitempty"`
	// example:
	//
	// 0
	CustomVoiceVolume *int32 `json:"CustomVoiceVolume,omitempty" xml:"CustomVoiceVolume,omitempty"`
	// example:
	//
	// 1920
	Height                    *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	HighDefSourceVideosShrink *string `json:"HighDefSourceVideos,omitempty" xml:"HighDefSourceVideos,omitempty"`
	MusicStyle                *string `json:"MusicStyle,omitempty" xml:"MusicStyle,omitempty"`
	// example:
	//
	// http://music.mp4
	MusicUrl          *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
	MusicVolume       *int32  `json:"MusicVolume,omitempty" xml:"MusicVolume,omitempty"`
	OpeningCreditsUrl *string `json:"OpeningCreditsUrl,omitempty" xml:"OpeningCreditsUrl,omitempty"`
	StickersShrink    *string `json:"Stickers,omitempty" xml:"Stickers,omitempty"`
	SubtitleFontSize  *int32  `json:"SubtitleFontSize,omitempty" xml:"SubtitleFontSize,omitempty"`
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

func (s *AsyncCreateClipsTaskShrinkRequest) GetCloseMusic() *bool {
	return s.CloseMusic
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCloseSubtitle() *bool {
	return s.CloseSubtitle
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCloseVoice() *bool {
	return s.CloseVoice
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetClosingCreditsUrl() *string {
	return s.ClosingCreditsUrl
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetColorWordsShrink() *string {
	return s.ColorWordsShrink
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCosyVoiceAppKey() *string {
	return s.CosyVoiceAppKey
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCosyVoiceToken() *string {
	return s.CosyVoiceToken
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCustomVoiceStyle() *string {
	return s.CustomVoiceStyle
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCustomVoiceUrl() *string {
	return s.CustomVoiceUrl
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCustomVoiceVolume() *int32 {
	return s.CustomVoiceVolume
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetHeight() *int32 {
	return s.Height
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetHighDefSourceVideosShrink() *string {
	return s.HighDefSourceVideosShrink
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetMusicStyle() *string {
	return s.MusicStyle
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetMusicVolume() *int32 {
	return s.MusicVolume
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetOpeningCreditsUrl() *string {
	return s.OpeningCreditsUrl
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetStickersShrink() *string {
	return s.StickersShrink
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

func (s *AsyncCreateClipsTaskShrinkRequest) SetCloseMusic(v bool) *AsyncCreateClipsTaskShrinkRequest {
	s.CloseMusic = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCloseSubtitle(v bool) *AsyncCreateClipsTaskShrinkRequest {
	s.CloseSubtitle = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCloseVoice(v bool) *AsyncCreateClipsTaskShrinkRequest {
	s.CloseVoice = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetClosingCreditsUrl(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.ClosingCreditsUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetColorWordsShrink(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.ColorWordsShrink = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCosyVoiceAppKey(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.CosyVoiceAppKey = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCosyVoiceToken(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.CosyVoiceToken = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCustomVoiceStyle(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.CustomVoiceStyle = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCustomVoiceUrl(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.CustomVoiceUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCustomVoiceVolume(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.CustomVoiceVolume = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetHeight(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.Height = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetHighDefSourceVideosShrink(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.HighDefSourceVideosShrink = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetMusicStyle(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.MusicStyle = &v
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

func (s *AsyncCreateClipsTaskShrinkRequest) SetOpeningCreditsUrl(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.OpeningCreditsUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetStickersShrink(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.StickersShrink = &v
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
